// Copyright Monte Carlo AI, Inc.
// SPDX-License-Identifier: Apache-2.0

package montecarlo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		name    string
		opts    Options
		wantErr string
	}{
		{
			name:    "endpoint is required",
			opts:    Options{Token: "t"},
			wantErr: "endpoint is required",
		},
		{
			name:    "no credentials at all",
			opts:    Options{Endpoint: "https://api.example.com"},
			wantErr: "no credentials",
		},
		{
			name:    "client id without secret",
			opts:    Options{Endpoint: "https://api.example.com", ClientID: "id"},
			wantErr: "client id and client secret are both required",
		},
		{
			name:    "client secret without id",
			opts:    Options{Endpoint: "https://api.example.com", ClientSecret: "s"},
			wantErr: "client id and client secret are both required",
		},
		{
			name: "oauth without instance",
			opts: Options{
				Endpoint: "https://api.example.com", ClientID: "id", ClientSecret: "secret",
			},
			wantErr: "instance is required",
		},
		{
			name:    "token id without secret",
			opts:    Options{Endpoint: "https://api.example.com", TokenID: "id"},
			wantErr: "token id and token secret are both required",
		},
		{
			name:    "token secret without id",
			opts:    Options{Endpoint: "https://api.example.com", TokenSecret: "s"},
			wantErr: "token id and token secret are both required",
		},
		{
			name: "api token",
			opts: Options{
				Endpoint: "https://api.example.com", TokenID: "id", TokenSecret: "secret",
			},
		},
		{
			name: "oauth with instance",
			opts: Options{
				Endpoint: "https://api.example.com", ClientID: "id",
				ClientSecret: "secret", Instance: "us1",
			},
		},
		{
			name: "pre-obtained bearer",
			opts: Options{Endpoint: "https://api.example.com", Token: "access-token"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.opts.validate()
			if c.wantErr == "" {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expected an error containing %q, got none", c.wantErr)
			}
			if !strings.Contains(err.Error(), c.wantErr) {
				t.Fatalf("expected an error containing %q, got %v", c.wantErr, err)
			}
		})
	}
}

// The gateway's authorizer parses an API token from `Bearer <id>:<secret>`, so the two
// halves are joined with a colon and callers never handle the wire format themselves.
func TestAPITokenBecomesColonJoinedBearer(t *testing.T) {
	o := Options{Endpoint: "https://api.example.com", TokenID: "key-id", TokenSecret: "s3cret"}
	if got := o.bearer(); got != "key-id:s3cret" {
		t.Fatalf("expected key-id:s3cret, got %q", got)
	}
}

// OAuth outranks the others, and produces no static bearer on the wire: NewClient consults
// bearer() only when there is no OAuth token source (see the switch in NewClient), so a
// request actually carries the OAuth-obtained token rather than the API token or the
// pre-obtained bearer, even though both of those are also set here and bearer() itself
// reports one of them non-empty. Checking UsesOAuth()/TokenSource() alone (the previous form
// of this test) proved neither half of the name: bearer() is non-empty for this Options
// regardless, so the only way to prove "carries no static bearer" is to look at the header
// that actually goes out.
func TestOAuthTakesPrecedenceAndCarriesNoStaticBearer(t *testing.T) {
	dir := isolate(t)

	var apiAuthHeaders []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/oauth2/token" {
			_, _ = w.Write([]byte(`{"access_token":"oauth-access-token","token_type":"Bearer"}`))
			return
		}
		apiAuthHeaders = r.Header.Values("Authorization")
		_, _ = w.Write([]byte(`{
			"account_frozen": false,
			"account_id": "a1",
			"email": "user@example.com",
			"identity_type": "user",
			"user_id": "u1"
		}`))
	}))
	defer server.Close()

	o := Options{
		Endpoint: server.URL,
		ClientID: "id", ClientSecret: "secret", Instance: "us1",
		TokenID: "key-id", TokenSecret: "s3cret", Token: "access-token",
		ConfigDir: dir,
	}
	if !o.usesOAuth() {
		t.Fatal("expected OAuth to be in use")
	}
	if b := o.bearer(); b == "" {
		t.Fatal("expected a static bearer to also be available, to prove OAuth wins rather than winning by default")
	}

	api, err := NewClient(context.Background(), o)
	if err != nil {
		t.Fatalf("unexpected error building the client: %v", err)
	}
	if _, _, err := api.UsersAPI.GetCurrentUser(context.Background()).Execute(); err != nil {
		t.Fatalf("unexpected error calling GetCurrentUser: %v", err)
	}

	if len(apiAuthHeaders) != 1 {
		t.Fatalf("expected exactly one Authorization header, got %d: %v", len(apiAuthHeaders), apiAuthHeaders)
	}
	if apiAuthHeaders[0] != "Bearer oauth-access-token" {
		t.Fatalf("expected the OAuth-obtained token, not the static bearer, got %q", apiAuthHeaders[0])
	}
}

func TestStaticCredentialsHaveNoTokenSource(t *testing.T) {
	o := Options{Endpoint: "https://api.example.com", TokenID: "id", TokenSecret: "secret"}
	ts, err := o.tokenSource(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ts != nil {
		t.Fatal("expected no token source for an API token")
	}
}

func TestTokenURLDefaultsUnderTheEndpoint(t *testing.T) {
	o := Options{Endpoint: "https://api.example.com/"}
	if got := o.tokenURL(); got != "https://api.example.com/oauth2/token" {
		t.Fatalf("unexpected token URL: %q", got)
	}

	o.TokenURL = "https://login.example.com/token"
	if got := o.tokenURL(); got != o.TokenURL {
		t.Fatalf("expected the override to win, got %q", got)
	}
}

// F42 (NIT): the previous form of this test built its expectation from the accessScope
// constant under test, so a typo'd accessScope would keep the test green while the gateway
// rejected every token request. The access scope is spelled out literally here instead.
func TestScopesForInstanceAppendsTheInstanceScopeAfterAccess(t *testing.T) {
	cases := []struct {
		name       string
		instanceID string
		want       []string
	}{
		{
			name:       "an instance id",
			instanceID: "us1",
			want: []string{
				"https://api.getmontecarlo.com/access",
				"https://instance.getmontecarlo.com/us1",
			},
		},
		{
			name:       "an empty instance id",
			instanceID: "",
			want: []string{
				"https://api.getmontecarlo.com/access",
				"https://instance.getmontecarlo.com/",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := scopesForInstance(c.instanceID)
			if !slices.Equal(got, c.want) {
				t.Fatalf("scopesForInstance(%q): got %v, want %v", c.instanceID, got, c.want)
			}
		})
	}
}

// The generated client has no default server URL, so a client built without one would send
// every request nowhere.
func TestNewClientSetsTheEndpoint(t *testing.T) {
	dir := isolate(t)
	api, err := NewClient(
		context.Background(),
		Options{
			Endpoint: "https://api.example.com/", TokenID: "id", TokenSecret: "secret",
			ConfigDir: dir,
		},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	url, err := api.GetConfig().ServerURLWithContext(context.Background(), "")
	if err != nil {
		t.Fatalf("unexpected error resolving the server URL: %v", err)
	}
	if url != "https://api.example.com" {
		t.Fatalf("expected the trailing slash trimmed, got %q", url)
	}
}

func TestNewClientRejectsMissingCredentials(t *testing.T) {
	dir := isolate(t)
	o := Options{Endpoint: "https://api.example.com", ConfigDir: dir}
	if _, err := NewClient(context.Background(), o); err == nil {
		t.Fatal("expected an error when no credentials are set")
	}
}

// F7 (ISSUE): Validate only checks that Endpoint is non-empty. It does not enforce a scheme,
// so a caller can be pointed at a plain-http, or scheme-less, endpoint and never find out
// until credentials are sent over it in the clear. A loopback http endpoint must stay
// accepted, since that is exactly what an httptest.Server exposes (see oauth_test.go).
func TestValidateRequiresAnHTTPSEndpointExceptLoopback(t *testing.T) {
	cases := []struct {
		name     string
		endpoint string
		wantErr  bool
	}{
		{
			name:     "rejects a plain http endpoint",
			endpoint: "http://collector.attacker.example",
			wantErr:  true,
		},
		{
			name:     "rejects a scheme-less endpoint",
			endpoint: "api.getmontecarlo.com",
			wantErr:  true,
		},
		{
			name:     "accepts an https endpoint",
			endpoint: "https://api.getmontecarlo.com",
			wantErr:  false,
		},
		{
			name:     "accepts an http loopback endpoint",
			endpoint: "http://127.0.0.1:8080",
			wantErr:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			o := Options{Endpoint: c.endpoint, TokenID: "id", TokenSecret: "secret"}
			err := o.validate()
			if c.wantErr && err == nil {
				t.Fatalf("expected %q to be rejected, but validate accepted it", c.endpoint)
			}
			if !c.wantErr && err != nil {
				t.Fatalf("expected %q to be accepted, got: %v", c.endpoint, err)
			}
		})
	}
}

// F7 (ISSUE), TokenURL half: an OAuth token exchange endpoint is just as capable of leaking
// client credentials as the API endpoint is, but validate never looks at TokenURL at all.
func TestValidateRequiresAnHTTPSTokenURL(t *testing.T) {
	o := Options{
		Endpoint:     "https://api.example.com",
		ClientID:     "id",
		ClientSecret: "secret",
		Instance:     "us1",
		TokenURL:     "http://attacker.example/oauth2/token",
	}
	if err := o.validate(); err == nil {
		t.Fatal("expected an insecure token URL to be rejected")
	}
}

// sentHeaders builds a client from o against a test server, makes one call with ctx, and
// returns the headers that call carried.
func sentHeaders(t *testing.T, ctx context.Context, o Options, configure func(*APIClient)) http.Header {
	t.Helper()
	var got http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Clone()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"account_frozen": false,
			"account_id": "a1",
			"email": "user@example.com",
			"identity_type": "user",
			"user_id": "u1"
		}`))
	}))
	defer server.Close()

	o.Endpoint = server.URL
	o.TokenID, o.TokenSecret = "id", "secret"
	o.ConfigDir = isolate(t)
	api, err := NewClient(context.Background(), o)
	if err != nil {
		t.Fatalf("unexpected error building the client: %v", err)
	}
	if configure != nil {
		configure(api)
	}
	if _, _, err := api.UsersAPI.GetCurrentUser(ctx).Execute(); err != nil {
		t.Fatalf("unexpected error calling GetCurrentUser: %v", err)
	}
	return got
}

func assertHeader(t *testing.T, h http.Header, key string, want ...string) {
	t.Helper()
	if got := h.Values(key); !slices.Equal(got, want) {
		t.Fatalf("%s: expected %q, got %q", key, want, got)
	}
}

func TestNewClientSendsDefaultTelemetry(t *testing.T) {
	h := sentHeaders(t, context.Background(), Options{}, nil)
	assertHeader(t, h, "x-mcd-telemetry-reason", "user")
	assertHeader(t, h, "x-mcd-telemetry-service", "mc-sdk-go")
	assertHeader(t, h, "x-mcd-telemetry-command")
	assertHeader(t, h, "x-mcd-source")
}

func TestTelemetryOptionsOverrideTheDefaults(t *testing.T) {
	o := Options{TelemetryReason: "cli", TelemetryService: "mc-cli", TelemetryCommand: "whoami"}
	h := sentHeaders(t, context.Background(), o, nil)
	assertHeader(t, h, "x-mcd-telemetry-reason", "cli")
	assertHeader(t, h, "x-mcd-telemetry-service", "mc-cli")
	assertHeader(t, h, "x-mcd-telemetry-command", "whoami")
}

func TestWithTelemetryCommandOverridesTheOption(t *testing.T) {
	ctx := WithTelemetryCommand(context.Background(), "montecarlo_warehouse create")
	h := sentHeaders(t, ctx, Options{TelemetryCommand: "whoami"}, nil)
	assertHeader(t, h, "x-mcd-telemetry-command", "montecarlo_warehouse create")
}

// A header the caller adds through the generated client's DefaultHeader wins, rather than
// going out alongside the default as a second value.
func TestTelemetryYieldsToADefaultHeader(t *testing.T) {
	h := sentHeaders(t, context.Background(), Options{}, func(api *APIClient) {
		api.GetConfig().AddDefaultHeader("x-mcd-telemetry-reason", "service")
	})
	assertHeader(t, h, "x-mcd-telemetry-reason", "service")
}
