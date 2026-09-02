package auth

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// ProfileFileName is the credentials file the Monte Carlo CLI writes, inside ConfigDir.
	ProfileFileName = "profiles.ini"
	// DefaultProfileName is the section used when no profile is named.
	DefaultProfileName = "default"
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

// ConfigDir is where the credentials file lives, ~/.mcd unless overridden.
func ConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".mcd"), nil
}

// LoadProfile reads one profile's credentials into Options.
//
// An empty name resolves to MCD_DEFAULT_PROFILE, then to "default". An empty configDir
// resolves to ConfigDir. A named profile that is not in the file is an error: a caller that
// asked for specific credentials should hear that they are missing, rather than silently
// getting someone else's.
//
// Within a profile, OAuth client credentials take precedence over an API token, matching the
// Python SDK.
func LoadProfile(name, configDir string) (Options, error) {
	if name == "" {
		name = os.Getenv("MCD_DEFAULT_PROFILE")
	}
	if name == "" {
		name = DefaultProfileName
	}
	if configDir == "" {
		dir, err := ConfigDir()
		if err != nil {
			return Options{}, err
		}
		configDir = dir
	}

	path := filepath.Join(configDir, ProfileFileName)
	sections, err := parseINI(path)
	if err != nil {
		return Options{}, err
	}
	values, ok := sections[name]
	if !ok {
		return Options{}, fmt.Errorf("profile %q not found in %s", name, path)
	}

	o := Options{
		Endpoint: values[keyAPIEndpoint],
		TokenURL: values[keyTokenEndpoint],
		Instance: values[keyInstanceID],
	}
	if id, secret := values[keyOAuthClientID], values[keyOAuthSecret]; id != "" && secret != "" {
		o.ClientID, o.ClientSecret = id, secret
	} else {
		o.TokenID, o.TokenSecret = values[keyID], values[keyToken]
	}
	return o, nil
}

// parseINI reads the small subset of INI the credentials file uses: sections, key = value,
// and comments introduced by # or ;. Keys are lower-cased, matching Python's configparser,
// so a hand-edited file with different casing still resolves.
func parseINI(path string) (map[string]map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sections := map[string]map[string]string{}
	current := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
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
		key, value, found := strings.Cut(line, "=")
		if !found || current == "" {
			continue
		}
		sections[current][strings.ToLower(strings.TrimSpace(key))] = strings.TrimSpace(value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return sections, nil
}

// Resolve fills in whatever o leaves empty, from the environment and then from a profile.
//
// Precedence matches the Python SDK, so the same credentials work across tools: values set
// on o win, then environment variables, then the profile.
//
// A profile is consulted only when o names one, MCD_DEFAULT_PROFILE is set, or no
// credentials were found by then. A missing file is not an error unless nothing else
// supplied credentials, so callers that pass their own do not need a profile to exist.
func (o Options) Resolve() (Options, error) {
	out := o

	out.TokenID = firstNonEmpty(out.TokenID, os.Getenv("MCD_DEFAULT_API_ID"))
	out.TokenSecret = firstNonEmpty(out.TokenSecret, os.Getenv("MCD_DEFAULT_API_TOKEN"))
	out.ClientID = firstNonEmpty(out.ClientID, os.Getenv("MCD_DEFAULT_OAUTH_CLIENT_ID"))
	out.ClientSecret = firstNonEmpty(out.ClientSecret, os.Getenv("MCD_DEFAULT_OAUTH_CLIENT_SECRET"))
	out.Instance = firstNonEmpty(out.Instance, os.Getenv("MCD_DEFAULT_INSTANCE_ID"))

	named := out.Profile != "" || os.Getenv("MCD_DEFAULT_PROFILE") != ""
	hasCredentials := out.UsesOAuth() || out.UsesAPIToken() || out.Token != ""
	if !named && hasCredentials && out.Endpoint != "" {
		return out, nil
	}

	profile, err := LoadProfile(out.Profile, out.ConfigPath)
	if err != nil {
		// A profile that was asked for by name has to exist. Otherwise the file is simply
		// absent, and whatever the caller supplied stands on its own.
		if named {
			return Options{}, err
		}
		return out, nil
	}

	out.Endpoint = firstNonEmpty(out.Endpoint, profile.Endpoint)
	out.TokenURL = firstNonEmpty(out.TokenURL, profile.TokenURL)
	out.Instance = firstNonEmpty(out.Instance, profile.Instance)
	if !hasCredentials {
		out.TokenID, out.TokenSecret = profile.TokenID, profile.TokenSecret
		out.ClientID, out.ClientSecret = profile.ClientID, profile.ClientSecret
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
