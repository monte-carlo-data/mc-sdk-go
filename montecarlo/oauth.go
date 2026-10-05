// Copyright Monte Carlo AI, Inc.
// SPDX-License-Identifier: Apache-2.0

// Hand-written, not generator output.

package montecarlo

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// accessScope is granted to every machine-to-machine client.
	accessScope = "https://api.getmontecarlo.com/access"
	// instanceScopePrefix joined with an instance id is the scope the gateway routes on.
	instanceScopePrefix = "https://instance.getmontecarlo.com/"
)

// scopesForInstance derives the OAuth scopes for an instance id such as "us1".
//
// A client is granted an access scope and one instance scope. The caller supplies only the
// instance id; deriving the scopes here keeps every consumer from reimplementing it.
func scopesForInstance(instanceID string) []string {
	return []string{accessScope, instanceScopePrefix + instanceID}
}

func oauthTokenSource(
	ctx context.Context, clientID, clientSecret, instanceID, tokenURL string,
) oauth2.TokenSource {
	cfg := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		Scopes:       scopesForInstance(instanceID),
		AuthStyle:    oauth2.AuthStyleInParams,
	}
	// A token source outlives the context that built it, so cancellation is detached: a
	// Terraform provider builds one during Configure, whose context is cancelled as soon as
	// Configure returns, and every later refresh would fail with "context canceled".
	// WithoutCancel keeps context values, such as a custom http.Client, but it drops the
	// caller's deadline along with the cancel signal — a detached context has no bound of its
	// own. Give it one, unless the caller already supplied an http.Client via the
	// oauth2.HTTPClient context value, in which case that client is used as the caller built
	// it, timeout and redirect policy included.
	//
	// The token request carries the client secret in its form body. A 307 or 308 would re-send
	// it to whatever host the Location header names, which the https check on TokenURL never saw.
	detached := context.WithoutCancel(ctx)
	if ctx.Value(oauth2.HTTPClient) == nil {
		detached = context.WithValue(detached, oauth2.HTTPClient, &http.Client{
			Timeout:       30 * time.Second,
			CheckRedirect: refuseRedirect,
		})
	}
	return cfg.TokenSource(detached)
}

// refuseRedirect makes an http.Client return a 3xx response as-is rather than follow it.
func refuseRedirect(*http.Request, []*http.Request) error {
	return http.ErrUseLastResponse
}
