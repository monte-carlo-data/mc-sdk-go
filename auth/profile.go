package auth

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	// profileFileName is the credentials file the Monte Carlo CLI writes, inside the config
	// directory.
	profileFileName = "profiles.ini"
	// defaultProfileName is the section used when no profile is named.
	defaultProfileName = "default"
)

// Keys as the CLI writes them, and as the Python SDK reads them. Every Monte Carlo tool
// shares this one file, so a customer configures credentials once.
const (
	keyID            = "mcd_id"
	keyToken         = "mcd_token"
	keyAPIEndpoint   = "mcd_api_endpoint"
	keyOAuthClientID = "mcd_oauth_client_id"
	keyOAuthSecret   = "mcd_oauth_client_secret"
	keyTokenEndpoint = "mcd_token_endpoint"
	keyInstanceID    = "mcd_instance_id"
)

// Environment variables this package reads, named as every other Monte Carlo tool names them.
const (
	envProfile      = "MCD_DEFAULT_PROFILE"
	envTokenID      = "MCD_DEFAULT_API_ID"
	envTokenSecret  = "MCD_DEFAULT_API_TOKEN"
	envClientID     = "MCD_DEFAULT_OAUTH_CLIENT_ID"
	envClientSecret = "MCD_DEFAULT_OAUTH_CLIENT_SECRET"
	envInstance     = "MCD_DEFAULT_INSTANCE_ID"
)

// environmentKeys is every name above. It exists so a caller can enumerate the set this package
// consults rather than restate it: a test that has to isolate itself from the developer's own
// environment ranges over this, so adding a variable to the block above without adding it here
// is the only way to reopen that gap.
var environmentKeys = []string{
	envProfile,
	envTokenID,
	envTokenSecret,
	envClientID,
	envClientSecret,
	envInstance,
}

// configDir returns the default directory holding the credentials file, ~/.mcd.
// Options.ConfigDir overrides it.
func configDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("determining config directory: %w", err)
	}
	return filepath.Join(home, ".mcd"), nil
}

// loadProfile reads one profile's credentials into Options.
//
// An empty name resolves to MCD_DEFAULT_PROFILE, then to "default". An empty dir resolves to
// configDir. A named profile that is not in the file is an error: a caller that asked for
// specific credentials should hear that they are missing, rather than silently getting someone
// else's.
//
// Within a profile, OAuth client credentials take precedence over an API token, matching the
// Python SDK. A profile with only half of either pair is a configuration error, not a silent
// fall-through to the other mechanism — the file is machine-written, so a half-set pair means
// something is wrong with it rather than that the other mechanism was intended.
//
// mcd_api_endpoint carries the GraphQL endpoint in every other Monte Carlo tool, so a trailing
// /graphql is stripped to derive the REST base this SDK needs, matching pycarlo's own
// derive_token_endpoint.
func loadProfile(name, dir string) (Options, error) {
	if name == "" {
		name = os.Getenv(envProfile)
	}
	if name == "" {
		name = defaultProfileName
	}
	if dir == "" {
		d, err := configDir()
		if err != nil {
			return Options{}, err
		}
		dir = d
	}

	path := filepath.Join(dir, profileFileName)
	sections, err := parseINI(path)
	if err != nil {
		return Options{}, err
	}
	values, ok := sections[name]
	if !ok {
		return Options{}, fmt.Errorf("profile %q not found in %s", name, path)
	}

	o := Options{
		Endpoint: stripGraphQLPath(values[keyAPIEndpoint]),
		TokenURL: values[keyTokenEndpoint],
		Instance: values[keyInstanceID],
	}

	clientID, clientSecret := values[keyOAuthClientID], values[keyOAuthSecret]
	tokenID, token := values[keyID], values[keyToken]
	switch {
	case clientID != "" && clientSecret != "":
		o.ClientID, o.ClientSecret = clientID, clientSecret
	case clientID != "":
		return Options{}, fmt.Errorf("profile %q in %s has %s but not %s", name, path, keyOAuthClientID, keyOAuthSecret)
	case clientSecret != "":
		return Options{}, fmt.Errorf("profile %q in %s has %s but not %s", name, path, keyOAuthSecret, keyOAuthClientID)
	case tokenID != "" && token != "":
		o.TokenID, o.TokenSecret = tokenID, token
	case tokenID != "":
		return Options{}, fmt.Errorf("profile %q in %s has %s but not %s", name, path, keyID, keyToken)
	case token != "":
		return Options{}, fmt.Errorf("profile %q in %s has %s but not %s", name, path, keyToken, keyID)
	}
	return o, nil
}

// stripGraphQLPath removes a trailing /graphql from a Monte Carlo GraphQL endpoint, deriving
// the REST base URL the same way pycarlo's derive_token_endpoint does. mcd_api_endpoint holds
// the GraphQL URL in every other Monte Carlo tool, while REST operations live under
// /api/v2/... on the bare host.
func stripGraphQLPath(endpoint string) string {
	return strings.TrimSuffix(strings.TrimRight(endpoint, "/"), "/graphql")
}

// parseINI reads the small subset of INI the credentials file uses: sections, key = value or
// key: value, and comments introduced by # or ;. Keys are lower-cased, matching Python's
// configparser, so a hand-edited file with different casing still resolves. A line that is
// neither a section header, a comment, nor a delimited key/value pair is a parse error rather
// than a silently dropped line: the file is machine-written, so an unparseable line means it is
// not what this code thinks it is. The error names the line's position and never its contents —
// a caller renders these verbatim (the Terraform provider puts them in a diagnostic, which
// Terraform does not redact), and the offending line is often the one holding the token.
func parseINI(path string) (map[string]map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sections := map[string]map[string]string{}
	current := ""

	scanner := bufio.NewScanner(file)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			current = strings.TrimSpace(line[1 : len(line)-1])
			if _, ok := sections[current]; !ok {
				sections[current] = map[string]string{}
			}
			continue
		}
		idx := strings.IndexAny(line, "=:")
		if idx < 0 {
			return nil, fmt.Errorf("%s:%d: expected a section header or a key/value pair", path, lineNo)
		}
		if current == "" {
			return nil, fmt.Errorf("%s:%d: key/value pair before any section header", path, lineNo)
		}
		key, value := line[:idx], line[idx+1:]
		sections[current][strings.ToLower(strings.TrimSpace(key))] = strings.TrimSpace(value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return sections, nil
}

// Resolve fills in whatever o leaves empty, from the environment and then from a profile.
//
// Resolution is per credential mechanism, not per field: whichever mechanism o already
// supplies completely is used as-is, and the environment — then the profile — is consulted
// only to pick a mechanism when o supplies none. This keeps an ambient MCD_DEFAULT_OAUTH_*
// trio from silently displacing an explicitly-passed API token (and vice versa), and keeps a
// half-set ambient variable from turning into a validation error for a mechanism the caller
// never asked for. Instance is filled from the environment unconditionally: it is a routing
// hint, not a credential, and is only adopted from a profile when OAuth ends up being the
// selected mechanism.
//
// A profile is consulted only when o names one, or no mechanism was found by then. A profile
// named explicitly has to exist; otherwise the file is simply absent, and whatever the caller
// supplied — or nothing — stands on its own. MCD_DEFAULT_PROFILE, if set, still selects which
// section a consulted profile reads from; it does not by itself make a missing file an error.
func (o Options) Resolve() (Options, error) {
	out := o

	if !out.usesOAuth() && !out.usesAPIToken() && out.Token == "" {
		switch clientID, clientSecret := os.Getenv(envClientID), os.Getenv(envClientSecret); {
		case clientID != "" && clientSecret != "":
			out.ClientID, out.ClientSecret = clientID, clientSecret
		default:
			if tokenID, tokenSecret := os.Getenv(envTokenID), os.Getenv(envTokenSecret); tokenID != "" && tokenSecret != "" {
				out.TokenID, out.TokenSecret = tokenID, tokenSecret
			}
		}
	}
	out.Instance = firstNonEmpty(out.Instance, os.Getenv(envInstance))

	named := out.Profile != ""
	hasCredentials := out.usesOAuth() || out.usesAPIToken() || out.Token != ""
	if !named && hasCredentials && out.Endpoint != "" {
		return out, nil
	}

	profile, err := loadProfile(out.Profile, out.ConfigDir)
	if err != nil {
		// A profile that was asked for by name has to exist. Otherwise the file is simply
		// absent, and whatever the caller supplied stands on its own.
		if named || !errors.Is(err, fs.ErrNotExist) {
			return Options{}, err
		}
		return out, nil
	}

	out.Endpoint = firstNonEmpty(out.Endpoint, profile.Endpoint)
	if !hasCredentials {
		out.TokenID, out.TokenSecret = profile.TokenID, profile.TokenSecret
		out.ClientID, out.ClientSecret = profile.ClientID, profile.ClientSecret
	}
	if out.usesOAuth() {
		out.TokenURL = firstNonEmpty(out.TokenURL, profile.TokenURL)
		out.Instance = firstNonEmpty(out.Instance, profile.Instance)
	}
	return out, nil
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
