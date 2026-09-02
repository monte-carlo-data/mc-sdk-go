package auth

import (
	"context"
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
			err := c.opts.Validate()
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

// OAuth outranks the others, and produces no static bearer: its token is fetched and
// refreshed per request instead.
func TestOAuthTakesPrecedenceAndCarriesNoStaticBearer(t *testing.T) {
	o := Options{
		Endpoint: "https://api.example.com",
		ClientID: "id", ClientSecret: "secret", Instance: "us1",
		TokenID: "key-id", TokenSecret: "s3cret", Token: "access-token",
	}
	if !o.UsesOAuth() {
		t.Fatal("expected OAuth to be in use")
	}

	ts, err := o.TokenSource(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ts == nil {
		t.Fatal("expected a token source for OAuth credentials")
	}
}

func TestStaticCredentialsHaveNoTokenSource(t *testing.T) {
	o := Options{Endpoint: "https://api.example.com", TokenID: "id", TokenSecret: "secret"}
	ts, err := o.TokenSource(context.Background())
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

func TestScopesForInstance(t *testing.T) {
	scopes := ScopesForInstance("us1")
	want := []string{AccessScope, "https://instance.getmontecarlo.com/us1"}
	if len(scopes) != len(want) {
		t.Fatalf("expected %d scopes, got %d", len(want), len(scopes))
	}
	for i := range want {
		if scopes[i] != want[i] {
			t.Fatalf("scope %d: expected %q, got %q", i, want[i], scopes[i])
		}
	}
}

// The generated client has no default server URL, so a client built without one would send
// every request nowhere.
func TestNewClientSetsTheEndpoint(t *testing.T) {
	api, err := NewClient(
		context.Background(),
		Options{Endpoint: "https://api.example.com/", TokenID: "id", TokenSecret: "secret"},
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
	if _, err := NewClient(context.Background(), Options{Endpoint: "https://api.example.com"}); err == nil {
		t.Fatal("expected an error when no credentials are set")
	}
}
