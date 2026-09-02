// Package auth builds an authenticated client for the Monte Carlo REST API.
//
// It is hand-written and sits beside the generated client, protected by
// .openapi-generator-ignore, so every consumer — the Terraform provider, the CLI, customer
// code — shares one implementation rather than each deriving scopes and token exchange.
//
// The API publishes a single credential mechanism: an Authorization bearer header carrying
// either a Monte Carlo API token or an OAuth 2.0 client-credentials access token. The
// gateway validates it; the service behind the gateway never sees it.
package auth

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/oauth2"

	montecarlo "github.com/monte-carlo-data/mc-sdk-go"
)

// Options are the inputs for building an authenticated client.
//
// Endpoint is always required: the generated client has no default server URL, because the
// API is reachable at a different host per instance.
//
// Exactly one credential is expected. Precedence, if more than one is set, is OAuth client
// credentials, then an API token, then a pre-obtained bearer.
type Options struct {
	// Endpoint is the API base URL, e.g. https://api.getmontecarlo.com. Required.
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
}

// UsesOAuth reports whether client-credentials auth is configured.
func (o Options) UsesOAuth() bool {
	return o.ClientID != "" && o.ClientSecret != ""
}

// UsesAPIToken reports whether a Monte Carlo API token is configured.
func (o Options) UsesAPIToken() bool {
	return o.TokenID != "" && o.TokenSecret != ""
}

// Validate enforces each mechanism's contract, so a misconfiguration is an error at
// construction rather than a 401 on the first call.
func (o Options) Validate() error {
	if o.Endpoint == "" {
		return fmt.Errorf("endpoint is required")
	}
	if o.ClientID != "" || o.ClientSecret != "" {
		if o.ClientID == "" || o.ClientSecret == "" {
			return fmt.Errorf("client id and client secret are both required for OAuth")
		}
		if o.Instance == "" {
			return fmt.Errorf("instance is required with OAuth client credentials, e.g. us1")
		}
	}
	if (o.TokenID != "") != (o.TokenSecret != "") {
		return fmt.Errorf("token id and token secret are both required for an API token")
	}
	if !o.UsesOAuth() && !o.UsesAPIToken() && o.Token == "" {
		return fmt.Errorf("no credentials: set client id and secret, token id and secret, or a token")
	}
	return nil
}

func (o Options) tokenURL() string {
	if o.TokenURL != "" {
		return o.TokenURL
	}
	return strings.TrimRight(o.Endpoint, "/") + "/oauth2/token"
}

// bearer is the credential to send, for the mechanisms that produce one up front. OAuth
// returns the empty string: its token is fetched, and refreshed, per request.
func (o Options) bearer() string {
	switch {
	case o.UsesAPIToken():
		return o.TokenID + ":" + o.TokenSecret
	case o.Token != "":
		return o.Token
	}
	return ""
}

// TokenSource returns the OAuth token source implied by o, or nil when the credential is
// static. A consumer that attaches credentials per request rather than reusing one client
// — a Terraform provider, for instance — needs this; NewClient wires it in already.
func (o Options) TokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}
	if !o.UsesOAuth() {
		return nil, nil
	}
	return oauthTokenSource(ctx, o.ClientID, o.ClientSecret, o.Instance, o.tokenURL()), nil
}

// NewClient builds a client that authenticates every request.
//
// OAuth credentials are wired through an http.Client that fetches and refreshes tokens as
// needed, rather than reading one token into the configuration. A long-lived caller would
// otherwise hold a credential that expires mid-run.
func NewClient(ctx context.Context, o Options) (*montecarlo.APIClient, error) {
	ts, err := o.TokenSource(ctx)
	if err != nil {
		return nil, err
	}

	cfg := montecarlo.NewConfiguration()
	cfg.Servers = montecarlo.ServerConfigurations{{URL: strings.TrimRight(o.Endpoint, "/")}}

	if ts != nil {
		cfg.HTTPClient = oauth2.NewClient(context.WithoutCancel(ctx), ts)
	} else if b := o.bearer(); b != "" {
		cfg.AddDefaultHeader("Authorization", "Bearer "+b)
	}

	return montecarlo.NewAPIClient(cfg), nil
}
