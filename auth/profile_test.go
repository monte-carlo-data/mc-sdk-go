package auth

import (
	"os"
	"path/filepath"
	"testing"
)

// isolate points credential resolution at an empty directory and clears every environment
// variable it consults, so a test never reads whatever the developer has configured.
func isolate(t *testing.T) string {
	t.Helper()
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
	return t.TempDir()
}

func writeProfiles(t *testing.T, dir, contents string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, ProfileFileName), []byte(contents), 0o600); err != nil {
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

	o, err := LoadProfile("", dir)
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

	o, err := LoadProfile("oauth-profile", dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !o.UsesOAuth() {
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

	o, err := LoadProfile("both-endpoints", dir)
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

	if _, err := LoadProfile("nope", dir); err == nil {
		t.Fatal("expected an error for a profile that is not in the file")
	}
}

func TestMCDDefaultProfileSelectsTheSection(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)
	t.Setenv("MCD_DEFAULT_PROFILE", "oauth-profile")

	o, err := LoadProfile("", dir)
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
		ConfigPath:  dir,
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

	resolved, err := Options{ConfigPath: dir}.Resolve()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resolved.TokenID != "env-id" {
		t.Fatalf("expected the environment to win over the profile, got %q", resolved.TokenID)
	}
	// The endpoint is not an environment variable in any Monte Carlo tool, so it still comes
	// from the profile.
	if resolved.Endpoint != "https://api.example.com" {
		t.Fatalf("expected the endpoint from the profile, got %q", resolved.Endpoint)
	}
}

func TestResolveFallsBackToTheProfile(t *testing.T) {
	dir := isolate(t)
	writeProfiles(t, dir, sampleProfiles)

	resolved, err := Options{ConfigPath: dir}.Resolve()
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
		ConfigPath: dir,
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

	if _, err := (Options{Profile: "nope", ConfigPath: dir}).Resolve(); err == nil {
		t.Fatal("expected an error for a named profile with no file")
	}
}
