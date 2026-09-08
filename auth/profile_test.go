package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// isolate clears every environment variable credential resolution consults and returns a
// temporary directory the caller must pass as Options.ConfigDir, so a test's file-backed
// profile lookups land there rather than on any real ~/.mcd. Passing that directory as
// ConfigDir is not optional: it is the only thing that makes a test's resolution use this
// empty directory instead of whatever the developer has configured for the tools that share
// this credentials file.
//
// As a backstop for a call site that discards the return value, or any future code path that
// still falls through to the home directory, isolate also points HOME (and USERPROFILE, for
// Windows) at that same temporary directory and creates an empty ".mcd" under it, so
// configDir()'s os.UserHomeDir() fallback resolves inside the temp dir rather than reading the
// developer's real ~/.mcd/profiles.ini. That backstop does not replace passing ConfigDir
// explicitly — it only keeps an omission from silently escaping isolation.
func isolate(t *testing.T) string {
	t.Helper()
	// This list mirrors the MCD_DEFAULT_* names Resolve reads directly (auth/profile.go) and
	// MCD_DEFAULT_PROFILE that loadProfile falls back to. profile.go declares no single list
	// of these names to range over, so this is a hand-maintained duplicate: adding a variable
	// to Resolve without adding it here silently reopens the isolation gap.
	for _, key := range []string{
		"MCD_DEFAULT_PROFILE",
		"MCD_DEFAULT_API_ID",
		"MCD_DEFAULT_API_TOKEN",
		"MCD_DEFAULT_OAUTH_CLIENT_ID",
		"MCD_DEFAULT_OAUTH_CLIENT_SECRET",
		"MCD_DEFAULT_INSTANCE_ID",
	} {
		t.Setenv(key, "")
	}

	dir := t.TempDir()
	t.Setenv("HOME", dir)
	t.Setenv("USERPROFILE", dir)
	if err := os.MkdirAll(filepath.Join(dir, ".mcd"), 0o755); err != nil {
		t.Fatalf("creating an empty .mcd under the fake home: %v", err)
	}
	return dir
}

func writeProfiles(t *testing.T, dir, contents string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, profileFileName), []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
}

// The file is written by the CLI and read by every Monte Carlo tool, so these are the keys
// and the layout it actually produces, including keys this SDK does not use.
const sampleProfiles = `
[default]
mcd_id = default-id
mcd_token = default-secret
mcd_api_endpoint = https://api.example.com
aws_profile = something-this-sdk-ignores

; a comment
[oauth-profile]
mcd_oauth_client_id = client-id
mcd_oauth_client_secret = client-secret
mcd_instance_id = eu1
mcd_id = ignored-when-oauth-is-present
mcd_token = also-ignored

[both-endpoints]
mcd_id = id
mcd_token = secret
mcd_api_endpoint = https://api.example.com
mcd_token_endpoint = https://login.example.com/token
`

func TestLoadProfileReadsTheDefaultSection(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)

	o, err := loadProfile("", dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if o.TokenID != "default-id" || o.TokenSecret != "default-secret" {
		t.Fatalf("unexpected credentials: %+v", o)
	}
	if o.Endpoint != "https://api.example.com" {
		t.Fatalf("unexpected endpoint: %q", o.Endpoint)
	}
}

func TestLoadProfilePrefersOAuthWithinAProfile(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)

	o, err := loadProfile("oauth-profile", dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !o.usesOAuth() {
		t.Fatal("expected OAuth credentials")
	}
	if o.TokenID != "" || o.TokenSecret != "" {
		t.Fatalf("expected the API token to be left unset, got %+v", o)
	}
	if o.Instance != "eu1" {
		t.Fatalf("unexpected instance: %q", o.Instance)
	}
}

func TestLoadProfileReadsTheTokenEndpoint(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)

	o, err := loadProfile("both-endpoints", dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if o.TokenURL != "https://login.example.com/token" {
		t.Fatalf("unexpected token URL: %q", o.TokenURL)
	}
}

// Asking for a profile by name and silently getting different credentials would be worse
// than failing.
func TestLoadProfileFailsOnAnUnknownName(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)

	if _, err := loadProfile("nope", dir); err == nil {
		t.Fatal("expected an error for a profile that is not in the file")
	}
}

func TestMCDDefaultProfileSelectsTheSection(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)
	t.Setenv("MCD_DEFAULT_PROFILE", "oauth-profile")

	o, err := loadProfile("", dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if o.ClientID != "client-id" {
		t.Fatalf("expected the profile named by the environment, got %+v", o)
	}
}

func TestResolvePrefersExplicitOptionsOverEverything(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)
	t.Setenv("MCD_DEFAULT_API_ID", "env-id")
	t.Setenv("MCD_DEFAULT_API_TOKEN", "env-secret")

	o := Options{
		Endpoint:    "https://explicit.example.com",
		TokenID:     "explicit-id",
		TokenSecret: "explicit-secret",
		ConfigDir:   dir,
	}
	resolved, err := o.Resolve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resolved.TokenID != "explicit-id" || resolved.Endpoint != "https://explicit.example.com" {
		t.Fatalf("explicit values should win, got %+v", resolved)
	}
}

func TestResolvePrefersEnvironmentOverProfile(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)
	t.Setenv("MCD_DEFAULT_API_ID", "env-id")
	t.Setenv("MCD_DEFAULT_API_TOKEN", "env-secret")

	resolved, err := Options{ConfigDir: dir}.Resolve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resolved.TokenID != "env-id" {
		t.Fatalf("expected the environment to win over the profile, got %q", resolved.TokenID)
	}
	// F30: this SDK deliberately does not read MCD_API_ENDPOINT — both siblings define it
	// (cli/montecarlodata/settings.py and pycarlo/common/settings.py both read it from the
	// environment, and pycarlo gives it top precedence), but it carries a GraphQL endpoint,
	// not the REST base URL this SDK needs (see stripGraphQLPath). So the endpoint still comes
	// from the profile, even though the API token above came from the environment.
	if resolved.Endpoint != "https://api.example.com" {
		t.Fatalf("expected the endpoint from the profile, got %q", resolved.Endpoint)
	}
}

func TestResolveFallsBackToTheProfile(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)

	resolved, err := Options{ConfigDir: dir}.Resolve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resolved.TokenID != "default-id" || resolved.Endpoint != "https://api.example.com" {
		t.Fatalf("expected the default profile, got %+v", resolved)
	}
}

// A caller passing its own credentials should not need a credentials file to exist.
func TestResolveToleratesAMissingFile(t *testing.T) {
	dir := isolate(t)

	resolved, err := Options{
		Endpoint: "https://api.example.com", TokenID: "id", TokenSecret: "secret",
		ConfigDir: dir,
	}.Resolve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resolved.TokenID != "id" {
		t.Fatalf("unexpected credentials: %+v", resolved)
	}
}

// But a profile asked for by name has to exist, even when the file does not.
func TestResolveFailsWhenANamedProfileIsMissing(t *testing.T) {
	dir := isolate(t)

	if _, err := (Options{Profile: "nope", ConfigDir: dir}).Resolve(); err == nil {
		t.Fatal("expected an error for a named profile with no file")
	}
}

// F1 (BLOCKER): Resolve fills each credential field from the environment independently of
// which mechanism the caller actually supplied. NewClient ranks OAuth above an API token, so
// an ambient MCD_DEFAULT_OAUTH_* trio silently displaces an explicitly-passed API token —
// resolution should be mechanism-wise: if the caller supplied a complete mechanism, ambient
// env vars for a *different* mechanism should not be layered in at all.
func TestAnAmbientOAuthEnvironmentDoesNotDisplaceAnExplicitAPIToken(t *testing.T) {
	t.Run("a full ambient oauth trio still loses to an explicit api token", func(t *testing.T) {
		dir := isolate(t)
		t.Setenv("MCD_DEFAULT_OAUTH_CLIENT_ID", "env-client-id")
		t.Setenv("MCD_DEFAULT_OAUTH_CLIENT_SECRET", "env-client-secret")
		t.Setenv("MCD_DEFAULT_INSTANCE_ID", "us1")

		resolved, err := Options{
			Endpoint:    "https://api.example.com",
			TokenID:     "explicit-id",
			TokenSecret: "explicit-secret",
			ConfigDir:   dir,
		}.Resolve()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if resolved.usesOAuth() {
			t.Fatal("an ambient OAuth trio should not displace an explicitly supplied API token, but usesOAuth() reports true")
		}
		if got := resolved.bearer(); got != "explicit-id:explicit-secret" {
			t.Fatalf("expected the explicit API token as the bearer, got %q", got)
		}
	})

	// The milder half: even a *partial* ambient OAuth trio should not stop an otherwise
	// complete, explicit API token from validating.
	t.Run("a partial ambient oauth trio does not break validation of an explicit api token", func(t *testing.T) {
		dir := isolate(t)
		t.Setenv("MCD_DEFAULT_OAUTH_CLIENT_ID", "env-client-id")

		resolved, err := Options{
			Endpoint:    "https://api.example.com",
			TokenID:     "explicit-id",
			TokenSecret: "explicit-secret",
			ConfigDir:   dir,
		}.Resolve()
		if err != nil {
			t.Fatalf("unexpected error resolving: %v", err)
		}
		if err := resolved.validate(); err != nil {
			t.Fatalf("an explicit, complete API token should validate despite a stray MCD_DEFAULT_OAUTH_CLIENT_ID, got: %v", err)
		}
	})
}

// F2 (ISSUE): Resolve discards every error loadProfile can return when no profile was named
// by the caller, on the assumption that the only possible failure is the file being absent.
// A malformed profiles file is a different, more actionable failure: the caller has a real
// file that this SDK cannot read, and finding that out as a bare "no credentials" error (or
// no error at all) hides the actual problem.
func TestResolveSurfacesAMalformedProfilesFileRatherThanSwallowingIt(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, "not an ini file\n[unclosed\n")

	resolved, err := Options{Endpoint: "https://api.example.com", ConfigDir: dir}.Resolve()
	if err == nil {
		t.Fatalf("expected Resolve to surface the malformed profiles file, got resolved=%+v, err=nil", resolved)
	}
	if !strings.Contains(err.Error(), "profiles.ini") && !strings.Contains(err.Error(), "profile") {
		t.Fatalf("expected the error to mention the file or the profile, got %q", err.Error())
	}
}

// F3 (ISSUE): named is true when either o.Profile is set or MCD_DEFAULT_PROFILE is exported,
// and any loadProfile failure then becomes a hard error — even when the caller supplied a
// complete, valid configuration and the environment variable is just ambient noise from the
// CLI. isolate blanks MCD_DEFAULT_PROFILE, so it is set here after calling isolate.
func TestAnAmbientDefaultProfileDoesNotBreakACompleteExplicitConfiguration(t *testing.T) {
	dir := isolate(t)
	t.Setenv("MCD_DEFAULT_PROFILE", "default")

	resolved, err := Options{
		Endpoint:    "https://api.example.com",
		TokenID:     "id",
		TokenSecret: "secret",
		ConfigDir:   dir,
	}.Resolve()
	if err != nil {
		t.Fatalf("a complete explicit configuration should not require the profiles file to exist, got: %v", err)
	}
	if resolved.TokenID != "id" || resolved.TokenSecret != "secret" || resolved.Endpoint != "https://api.example.com" {
		t.Fatalf("expected the explicit credentials to be preserved, got %+v", resolved)
	}
}

// F8 state B (ISSUE): mcd_api_endpoint holds the GraphQL URL in every other Monte Carlo tool
// — the CLI writes MCD_DEFAULT_API_ENDPOINT / mcd_api_endpoint as
// "https://api.getmontecarlo.com/graphql" — but REST operations live under /api/v2/..., so a
// trailing /graphql must be stripped to get the REST base, matching pycarlo's own derivation.
func TestLoadProfileStripsTheGraphQLSuffixFromTheAPIEndpoint(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, "[default]\nmcd_api_endpoint = https://api.example.com/graphql\n")

	o, err := loadProfile("", dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if o.Endpoint != "https://api.example.com" {
		t.Fatalf("expected the trailing /graphql stripped from the REST endpoint, got %q", o.Endpoint)
	}
	if got := o.tokenURL(); got != "https://api.example.com/oauth2/token" {
		t.Fatalf("expected the token URL derived from the stripped endpoint, got %q", got)
	}
}

// Regression proof: a malformed profiles.ini must not put the credential it holds into the
// error, because callers render these verbatim into unredacted surfaces.
func TestParseINIDoesNotEchoTheOffendingLine(t *testing.T) {
	dir := t.TempDir()
	secret := "not-a-real-token-only-a-fixture"
	if err := os.WriteFile(filepath.Join(dir, profileFileName),
		[]byte("mcd_token = "+secret+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	_, err := parseINI(filepath.Join(dir, profileFileName))
	if err == nil {
		t.Fatal("expected a parse error")
	}
	if strings.Contains(err.Error(), secret) {
		t.Fatalf("the error carries the credential: %v", err)
	}
}
