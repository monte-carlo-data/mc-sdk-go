package auth

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// AccessScope is granted to every machine-to-machine client.
	AccessScope = "https://api.getmontecarlo.com/access"
	// InstanceScopePrefix joined with an instance id is the scope the gateway routes on.
	InstanceScopePrefix = "https://instance.getmontecarlo.com/"
)

// ScopesForInstance derives the OAuth scopes for an instance id such as "us1".
//
// A client is granted an access scope and one instance scope. The caller supplies only the
// instance id; deriving the scopes here keeps every consumer from reimplementing it.
func ScopesForInstance(instanceID string) []string {
	return []string{AccessScope, InstanceScopePrefix + instanceID}
}

func oauthTokenSource(
	ctx context.Context, clientID, clientSecret, instanceID, tokenURL string,
) oauth2.TokenSource {
	cfg := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       ScopesForInstance(instanceID),
		AuthStyle:    oauth2.AuthStyleInParams,
	}
	// A token source outlives the context that built it, so cancellation is detached: a
	// Terraform provider builds one during Configure, whose context is cancelled as soon as
	// Configure returns, and every later refresh would fail with "context canceled".
	// WithoutCancel keeps context values, such as a custom http.Client, and drops the signal.
	return cfg.TokenSource(context.WithoutCancel(ctx))
}
