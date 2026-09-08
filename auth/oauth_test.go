package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	montecarlo "github.com/monte-carlo-data/mc-sdk-go"
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

	ctx := context.WithValue(context.Background(), montecarlo.ContextAccessToken, "delegated-token")
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
