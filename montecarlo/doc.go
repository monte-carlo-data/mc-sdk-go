// Package montecarlo is the Go SDK for the Monte Carlo REST API.
//
// Build a client with [NewClient]. Credentials resolve from the [Options] passed in, then
// from the environment, then from the credentials file the Monte Carlo CLI writes — so a
// caller with the CLI already configured can often pass an empty Options and get a working
// client. Do not build one from [NewAPIClient] or [NewConfiguration] directly: those two have
// no default server URL, so a client built from them sends every request to a relative URL
// and never reaches the API.
//
// The API accepts exactly one credential mechanism per request: an OAuth 2.0
// client-credentials access token, or a Monte Carlo API token, sent as an Authorization
// bearer header.
//
// NewClient, Options and the credential resolution behind them are hand-written, in auth.go,
// oauth.go and profile.go. The pagination driver, in paginate.go, is hand-written too.
// Everything else in this package is generated from the API's OpenAPI spec and must not be
// hand-edited; changes are overwritten on the next generation run.
//
// A paginated list's request type has an All method, returning an [iter.Seq2] of the item
// type and an error. All walks every page, following the list's cursor, and yields the error
// as the final pair if a request fails — matchable with errors.Is against [ErrNoCursor] or
// [ErrRepeatedCursor]. Execute still returns a single page.
package montecarlo
