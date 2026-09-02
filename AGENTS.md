# mc-sdk-go

> The Go SDK for the Monte Carlo REST API. Generated from the API's OpenAPI spec, except the `auth` package, which is hand-written.

## Stack

- **Language:** Go 1.23+
- **Only dependency:** `golang.org/x/oauth2`

## Common Commands

```bash
go build ./...
go test ./...
go vet ./...
gofmt -l .        # must be empty; generated output is formatted when it is produced
```

## Key Directories

| Path | Purpose |
|------|---------|
| `auth/` | **Hand-written.** Client construction, API-token and OAuth credentials |
| `api_*.go`, `model_*.go` | Generated operations and models, one file per tag group and schema |
| `client.go`, `configuration.go`, `response.go`, `utils.go` | Generated client plumbing |
| `docs/` | Generated API reference |

## What is generated

Everything except `auth/`, the module files, and this documentation. The generator replaces
the whole tree on each run, so a fix to a generated file does not survive — it belongs in the
API or in the generator that reads its spec.

**Anything hand-written must be listed in `.openapi-generator-ignore` first.** A file that is
not listed is deleted on the next run, and nothing reports it.

Regeneration is a maintainer task and lives outside this repository, with the tooling that
owns generation for every artifact built from the spec.

## Authentication

The API publishes one credential mechanism: an `Authorization` bearer header carrying either
a Monte Carlo API token or an OAuth 2.0 client-credentials access token. It is validated
before a request reaches the service.

An API token has two halves, and the bearer value is the two joined by a colon. `Options`
takes them separately and joins them, so a caller never handles that format.

OAuth credentials are wired through an `http.Client` that refreshes tokens as needed, rather
than reading one token into the configuration. A long-lived caller would otherwise hold a
credential that expires mid-run.

**`Endpoint` is always required**, though it can come from a profile rather than the caller.
The generated client has no default server URL, because the API is reached at a different
host per deployment, so a client built without one sends every request nowhere.

## Credentials are shared with the other tools

`~/.mcd/profiles.ini` is written by the CLI and read by the Python SDK, so this SDK reads the
same file, the same section names and the same keys. A customer configures credentials once
and every tool picks them up, which is the whole point — an SDK with its own credential store
would strand anyone who had already run the CLI.

Precedence mirrors the Python SDK exactly: values passed in, then environment variables, then
the profile. Diverging would mean the same configuration behaving differently depending on
which tool read it.

Tests must not read the developer's real credentials. `isolate(t)` in the test package clears
every environment variable resolution consults and points `ConfigPath` at a temporary
directory; use it in any test that builds a client or resolves options.

## Branching

Branch from `main` as `<person>/<ticket-id>-<slug>`. Never commit directly to `main`.

## Releasing

**Not yet.** A Go module path is permanent once a version tag is published, and this
repository's name is not final. Fetching by commit sha resolves a pseudo-version and is
enough to verify the module builds for a consumer.
