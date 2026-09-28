// Hand-written, not generator output.

package montecarlo

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"golang.org/x/oauth2"
)

// defaultUserAgent identifies this SDK and the Go runtime it is built with. Options.UserAgent
// overrides it.
var defaultUserAgent = fmt.Sprintf("mc-sdk-go/dev (go/%s)", runtime.Version())

// Options are the inputs for building an authenticated client.
//
// Endpoint is always required, though it can come from a profile rather than the caller — see
// Resolve. The generated client has no default server URL, because the API is reachable at a
// different host per instance.
//
// Exactly one credential mechanism is expected. Precedence, when o supplies more than one
// directly, is OAuth client credentials, then an API token, then a pre-obtained bearer; see
// Resolve for how a mechanism supplied through the environment or a profile is chosen instead.
type Options struct {
	// Endpoint is the API base URL, e.g. https://api.getmontecarlo.com. Required, though it
	// can come from a profile rather than the caller.
	Endpoint string

	// TokenID and TokenSecret are the two halves of a Monte Carlo API token. They are
	// supplied separately and combined into the bearer credential here, so a caller never
	// has to know the wire format.
	TokenID     string
	TokenSecret string

	// Token is a bearer credential the caller already holds. Use TokenID and TokenSecret
	// for a Monte Carlo API token; this is for an access token obtained elsewhere.
	Token string

	// ClientID and ClientSecret are OAuth 2.0 client credentials. Instance selects which
	// deployment the gateway routes to and is required with them.
	ClientID     string
	ClientSecret string
	Instance     string

	// TokenURL overrides where client credentials are exchanged. Defaults to Endpoint with
	// /oauth2/token appended.
	TokenURL string

	// Profile names a section of the credentials file the Monte Carlo CLI writes, so
	// credentials configured once are shared across tools. Empty means MCD_DEFAULT_PROFILE,
	// then "default". See Resolve for the full precedence.
	Profile string

	// ConfigDir overrides the directory holding that file. Empty means ~/.mcd.
	ConfigDir string

	// Transport is the base RoundTripper credentials are layered on top of. Nil means
	// http.DefaultTransport, matching the generated client's own default. Set this to add a
	// proxy, retries, or instrumentation without losing authentication, since the credential
	// wraps whatever is set here rather than replacing it.
	Transport http.RoundTripper

	// UserAgent overrides the User-Agent sent with every request. Empty means a default that
	// identifies this SDK and the Go runtime it is built with.
	UserAgent string

	// TelemetryReason, TelemetryService and TelemetryCommand identify the calling client to
	// Monte Carlo's usage telemetry, sent as the x-mcd-telemetry-* headers. They are the only
	// client identity that reaches the API: the gateway drops User-Agent. Empty reason and
	// service mean "user" and "mc-sdk-go"; an empty command sends no command header.
	// WithTelemetryCommand sets the command for a single request instead.
	TelemetryReason  string
	TelemetryService string
	TelemetryCommand string
}

// usesOAuth reports whether client-credentials auth is configured.
func (o Options) usesOAuth() bool {
	return o.ClientID != "" && o.ClientSecret != ""
}

// usesAPIToken reports whether a Monte Carlo API token is configured.
func (o Options) usesAPIToken() bool {
	return o.TokenID != "" && o.TokenSecret != ""
}

// validate enforces each mechanism's contract, so a misconfiguration is an error at
// construction rather than a 401 on the first call. NewClient calls this right after Resolve;
// tokenSource calls it too, so it never builds a source from incomplete options.
func (o Options) validate() error {
	if o.Endpoint == "" {
		return errors.New("endpoint is required")
	}
	if err := validateURL(o.Endpoint, "endpoint"); err != nil {
		return err
	}
	if o.TokenURL != "" {
		if err := validateURL(o.TokenURL, "token url"); err != nil {
			return err
		}
	}
	if err := o.checkPairs(); err != nil {
		return err
	}
	if o.usesOAuth() && o.Instance == "" {
		return errors.New("instance is required with OAuth client credentials, e.g. us1")
	}
	if !o.usesOAuth() && !o.usesAPIToken() && o.Token == "" {
		return errors.New("no credentials: set client id and secret, token id and secret, or a token")
	}
	return nil
}

// checkPairs rejects half of a credential pair. Resolve runs it on the caller's options before any
// fallback is consulted, and validate runs it again on the resolved options.
func (o Options) checkPairs() error {
	if (o.ClientID != "") != (o.ClientSecret != "") {
		return errors.New("client id and client secret are both required for OAuth")
	}
	if (o.TokenID != "") != (o.TokenSecret != "") {
		return errors.New("token id and token secret are both required for an API token")
	}
	return nil
}

// validateURL requires an absolute https URL, permitting http only for a loopback host so
// httptest-based tests keep working.
func validateURL(raw, field string) error {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("%s must be an absolute URL: %q", field, raw)
	}
	switch {
	case u.Scheme == "https":
		return nil
	case u.Scheme == "http" && isLoopback(u.Hostname()):
		return nil
	default:
		return fmt.Errorf("%s must use https (http is allowed only for a loopback host): %q", field, raw)
	}
}

func isLoopback(host string) bool {
	if host == "localhost" {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

// baseURL is Endpoint with any trailing slash trimmed, the form every use site needs.
func (o Options) baseURL() string {
	return strings.TrimRight(o.Endpoint, "/")
}

func (o Options) tokenURL() string {
	if o.TokenURL != "" {
		return o.TokenURL
	}
	return o.baseURL() + "/oauth2/token"
}

// bearer is the credential to send, for the mechanisms that produce one up front. It ignores
// OAuth entirely: NewClient consults it only when there is no OAuth token source, so a stale
// static header can never ride along an OAuth-authenticated request.
func (o Options) bearer() string {
	switch {
	case o.usesAPIToken():
		return o.TokenID + ":" + o.TokenSecret
	case o.Token != "":
		return o.Token
	}
	return ""
}

// tokenSource returns the OAuth token source implied by o, or nil when the credential is
// static.
func (o Options) tokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	if err := o.validate(); err != nil {
		return nil, err
	}
	if !o.usesOAuth() {
		return nil, nil
	}
	return oauthTokenSource(ctx, o.ClientID, o.ClientSecret, o.Instance, o.tokenURL()), nil
}

// NewClient builds a client that authenticates every request.
//
// Anything o leaves empty is filled in from the environment and the credentials file, so a
// caller can pass nothing and get whatever the CLI configured. Resolve documents that order.
//
// Credentials are layered onto o.Transport (http.DefaultTransport if nil) as an
// http.RoundTripper, rather than stored in the client configuration: an OAuth credential
// fetches and refreshes its own token as needed, and a static one sets its header only when
// the request does not already carry one, so a per-request credential set through the
// generated client's ContextAccessToken always wins cleanly instead of stacking into a second
// header. Because of this, GetConfig().HTTPClient never needs to be mutated after the fact to
// add transport-level behaviour — pass it through Options.Transport instead.
//
// Neither client follows a redirect. The API client returns a 3xx as the response, and the
// token exchange fails on one, because its request carries the client secret in its form
// body. Options.Transport does not reach the token exchange; an http.Client supplied through
// the oauth2.HTTPClient context value replaces the exchange's default client wholesale,
// timeout and redirect policy included.
func NewClient(ctx context.Context, o Options) (*APIClient, error) {
	resolved, err := o.Resolve()
	if err != nil {
		return nil, err
	}
	if err := resolved.validate(); err != nil {
		return nil, err
	}

	ts, err := resolved.tokenSource(ctx)
	if err != nil {
		return nil, err
	}

	base := resolved.Transport
	if base == nil {
		base = http.DefaultTransport
	}

	var transport http.RoundTripper = base
	switch b := resolved.bearer(); {
	case ts != nil:
		transport = &oauth2.Transport{Base: base, Source: ts}
	case b != "":
		transport = &bearerTransport{base: base, bearer: b}
	}
	transport = &telemetryTransport{
		base:    transport,
		reason:  firstNonEmpty(resolved.TelemetryReason, defaultTelemetryReason),
		service: firstNonEmpty(resolved.TelemetryService, defaultTelemetryService),
		command: resolved.TelemetryCommand,
	}

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{URL: resolved.baseURL()},
	}
	cfg.UserAgent = firstNonEmpty(resolved.UserAgent, defaultUserAgent)
	// The API has one host, so a redirect is refused rather than followed: both transports attach
	// the credential on every hop, which would hand it to whatever host Location names.
	cfg.HTTPClient = &http.Client{Transport: transport, CheckRedirect: refuseRedirect}

	return NewAPIClient(cfg), nil
}

// bearerTransport attaches a static Authorization header to every request, yielding to one
// already present so a per-request credential set through the generated client's
// ContextAccessToken wins cleanly instead of stacking into a second header.
type bearerTransport struct {
	base   http.RoundTripper
	bearer string
}

func (t *bearerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("Authorization") != "" {
		return t.base.RoundTrip(req)
	}
	clone := req.Clone(req.Context())
	clone.Header.Set("Authorization", "Bearer "+t.bearer)
	return t.base.RoundTrip(clone)
}

const (
	headerTelemetryReason  = "x-mcd-telemetry-reason"
	headerTelemetryService = "x-mcd-telemetry-service"
	headerTelemetryCommand = "x-mcd-telemetry-command"

	// defaultTelemetryReason matches pycarlo's RequestReason.USER.
	defaultTelemetryReason  = "user"
	defaultTelemetryService = "mc-sdk-go"
)

type telemetryCommandKey struct{}

// WithTelemetryCommand returns a context that sends command as the x-mcd-telemetry-command
// header on requests made with it, overriding Options.TelemetryCommand. Name the operation, not
// its arguments, e.g. "connections add snowflake".
func WithTelemetryCommand(ctx context.Context, command string) context.Context {
	return context.WithValue(ctx, telemetryCommandKey{}, command)
}

// telemetryTransport sets the x-mcd-telemetry-* headers on every request, yielding to any
// already present so a header added through the generated client's DefaultHeader is not
// stacked into a second value.
type telemetryTransport struct {
	base                     http.RoundTripper
	reason, service, command string
}

func (t *telemetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	command := t.command
	if c, ok := req.Context().Value(telemetryCommandKey{}).(string); ok && c != "" {
		command = c
	}
	clone := req.Clone(req.Context())
	setIfAbsent(clone.Header, headerTelemetryReason, t.reason)
	setIfAbsent(clone.Header, headerTelemetryService, t.service)
	setIfAbsent(clone.Header, headerTelemetryCommand, command)
	return t.base.RoundTrip(clone)
}

func setIfAbsent(h http.Header, key, value string) {
	if value != "" && h.Get(key) == "" {
		h.Set(key, value)
	}
}
