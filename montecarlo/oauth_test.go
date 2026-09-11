package montecarlo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/oauth2"
)

// F5 (ISSUE), now fixed: the static-credential path used to put the bearer in
// cfg.DefaultHeader, which client.go's prepareRequest applied with Header.Add *in addition
// to* the Authorization header it adds separately for a per-request
// ctx.Value(ContextAccessToken) — a caller delegating a per-request access token on top of a
// client built with a static API token ended up sending two Authorization headers on the
// wire. The credential now lives on a bearerTransport (see auth.go) that yields to a header
// already present on the request, so exactly one Authorization header goes out either way.
func TestNewClientSendsOnlyOneAuthorizationHeaderWithADelegatedAccessToken(t *testing.T) {
	dir := isolate(t)

	var authHeaders []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaders = r.Header.Values("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"account_frozen": false,
			"account_id": "a1",
			"email": "user@example.com",
			"identity_type": "user",
			"user_id": "u1"
		}`))
	}))
	defer server.Close()

	api, err := NewClient(context.Background(), Options{
		Endpoint:    server.URL,
		TokenID:     "static-id",
		TokenSecret: "static-secret",
		ConfigDir:   dir,
	})
	if err != nil {
		t.Fatalf("unexpected error building the client: %v", err)
	}

	ctx := context.WithValue(context.Background(), ContextAccessToken, "delegated-token")
	if _, _, err := api.UsersAPI.GetCurrentUser(ctx).Execute(); err != nil {
		t.Fatalf("unexpected error calling GetCurrentUser: %v", err)
	}

	if len(authHeaders) != 1 {
		t.Fatalf("expected exactly one Authorization header, got %d: %v", len(authHeaders), authHeaders)
	}
}

// F6 (ISSUE), now fixed: NewClient used to store the static bearer via
// cfg.AddDefaultHeader("Authorization", "Bearer "+b), which landed in
// Configuration.DefaultHeader — an exported, `json:"defaultHeader,omitempty"` field reachable
// through api.GetConfig(). Anything that logs or dumps the configuration would have leaked a
// long-lived credential: fmt's %+v renders it, and cfg.Debug makes the generated client pass
// the request through httputil.DumpRequestOut. Both credential mechanisms now authenticate
// through cfg.HTTPClient's Transport instead (a bearerTransport or an oauth2.Transport — see
// auth.go), so neither ever reaches the configuration at all; this test guards against that
// regressing rather than distinguishing a fixed half from an already-immune one.
//
// Asserted on DefaultHeader rather than on json.Marshal of the whole Configuration: that
// marshal fails outright on HTTPClient.CheckRedirect, a func field, so it is not a live
// exposure path and cannot prove anything.
func TestGetConfigDoesNotLeakTheStaticAPIToken(t *testing.T) {
	t.Run("an api token must not appear in the marshaled configuration", func(t *testing.T) {
		dir := isolate(t)
		api, err := NewClient(context.Background(), Options{
			Endpoint:    "https://api.example.com",
			TokenID:     "leaked-id",
			TokenSecret: "leaked-secret",
			ConfigDir:   dir,
		})
		if err != nil {
			t.Fatalf("unexpected error building the client: %v", err)
		}

		rendered := fmt.Sprintf("%+v", api.GetConfig().DefaultHeader)
		if strings.Contains(rendered, "leaked-id") || strings.Contains(rendered, "leaked-secret") {
			t.Fatalf("expected the API token not to be stored on the configuration, got: %s", rendered)
		}
	})

	// Contrasting control: OAuth is already immune, since its credential never becomes a
	// static header in the first place.
	t.Run("an oauth client secret does not appear in the marshaled configuration", func(t *testing.T) {
		dir := isolate(t)
		api, err := NewClient(context.Background(), Options{
			Endpoint:     "https://api.example.com",
			ClientID:     "client-id",
			ClientSecret: "client-secret",
			Instance:     "us1",
			ConfigDir:    dir,
		})
		if err != nil {
			t.Fatalf("unexpected error building the client: %v", err)
		}

		rendered := fmt.Sprintf("%+v", api.GetConfig().DefaultHeader)
		if strings.Contains(rendered, "client-secret") {
			t.Fatalf("expected the OAuth client secret not to be stored on the configuration, got: %s", rendered)
		}
	})
}

// Regression proof: the default token-exchange client used to follow redirects, and a 307 or
// 308 re-sends the form body, client secret included, to whatever host the Location header
// names. The exchange must end at the redirect instead.
func TestTokenExchangeDoesNotFollowARedirectWithTheClientSecret(t *testing.T) {
	statuses := []int{
		http.StatusMovedPermanently,
		http.StatusFound,
		http.StatusSeeOther,
		http.StatusTemporaryRedirect,
		http.StatusPermanentRedirect,
	}
	for _, status := range statuses {
		t.Run(http.StatusText(status), func(t *testing.T) {
			var collectorHits, tokenHits int
			var mux http.ServeMux
			server := httptest.NewServer(&mux)
			defer server.Close()
			mux.HandleFunc("/oauth2/token", func(w http.ResponseWriter, r *http.Request) {
				tokenHits++
				http.Redirect(w, r, server.URL+"/collector", status)
			})
			mux.HandleFunc("/collector", func(w http.ResponseWriter, r *http.Request) {
				collectorHits++
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"access_token":"leaked","token_type":"Bearer"}`))
			})

			ts := oauthTokenSource(context.Background(), "id", "secret", "us1", server.URL+"/oauth2/token")
			_, err := ts.Token()
			var re *oauth2.RetrieveError
			if !errors.As(err, &re) {
				t.Fatalf("expected an *oauth2.RetrieveError, got %v", err)
			}
			if re.Response.StatusCode != status {
				t.Fatalf("expected the error to carry status %d, got %d", status, re.Response.StatusCode)
			}
			if collectorHits != 0 {
				t.Fatalf("the redirect target received %d request(s); the client secret was re-sent", collectorHits)
			}
			if tokenHits != 1 {
				t.Fatalf("expected exactly one token request, got %d", tokenHits)
			}
		})
	}
}

// The oauth2.HTTPClient context value replaces the default client wholesale, redirect policy included.
func TestACallerSuppliedHTTPClientKeepsItsOwnRedirectPolicy(t *testing.T) {
	var collectorHits, tokenHits int
	var mux http.ServeMux
	server := httptest.NewServer(&mux)
	defer server.Close()
	mux.HandleFunc("/oauth2/token", func(w http.ResponseWriter, r *http.Request) {
		tokenHits++
		http.Redirect(w, r, server.URL+"/collector", http.StatusTemporaryRedirect)
	})
	mux.HandleFunc("/collector", func(w http.ResponseWriter, r *http.Request) {
		collectorHits++
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"followed","token_type":"Bearer"}`))
	})

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{})
	ts := oauthTokenSource(ctx, "id", "secret", "us1", server.URL+"/oauth2/token")
	if _, err := ts.Token(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if collectorHits != 1 {
		t.Fatalf("expected the redirect to be followed exactly once, got %d hit(s)", collectorHits)
	}
}

// Regression proof: the API client used to follow redirects, and bearerTransport re-attached the
// credential on every hop after Go had stripped it for the new host.
func TestTheAPIClientDoesNotFollowARedirectWithTheCredential(t *testing.T) {
	dir := isolate(t)

	var collectorHits int
	var collectorAuth []string
	var mux http.ServeMux
	server := httptest.NewServer(&mux)
	defer server.Close()
	mux.HandleFunc("/api/v2/users/me", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, server.URL+"/collector", http.StatusTemporaryRedirect)
	})
	mux.HandleFunc("/collector", func(w http.ResponseWriter, r *http.Request) {
		collectorHits++
		collectorAuth = r.Header.Values("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"account_frozen": false,
			"account_id": "a1",
			"email": "user@example.com",
			"identity_type": "user",
			"user_id": "u1"
		}`))
	})

	api, err := NewClient(context.Background(), Options{
		Endpoint:    server.URL,
		TokenID:     "static-id",
		TokenSecret: "static-secret",
		ConfigDir:   dir,
	})
	if err != nil {
		t.Fatalf("unexpected error building the client: %v", err)
	}

	if _, _, err := api.UsersAPI.GetCurrentUser(context.Background()).Execute(); err == nil {
		t.Fatal("expected the redirected call to fail")
	}
	if collectorHits != 0 {
		t.Fatalf("the redirect target received %d request(s); the credential was re-sent", collectorHits)
	}
	if len(collectorAuth) != 0 {
		t.Fatalf("expected no Authorization header to reach the redirect target, got %v", collectorAuth)
	}
}
