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
| `api/openapi.yaml` | Generated copy of the input spec — not the source of truth, overwritten every run |
| `.openapi-generator/` | Generator bookkeeping |
| `docs/` | Generated API reference |

## What is generated

Everything except the paths listed below. The generator overwrites every path it emits on
each run — it does not delete anything else — so a fix to a generated file does not survive;
it belongs in the API or in the generator that reads its spec.

**Anything hand-written, or anything we don't want touched, must be listed in
`.openapi-generator-ignore`.** This is the canonical enumeration — `api-codegen`'s generation
script checks its own copy of this list against it:

- `auth/` — hand-written authentication and client construction
- `go.mod`, `go.sum` — our dependency set, not the generator's guess
- `README.md`, `AGENTS.md`, `CLAUDE.md`, `CODEOWNERS` — repository documentation
- `doc.go` — the root package doc comment, which is the pkg.go.dev landing page
- `.github/`, `.claude/`, `.work/` — repository and tooling configuration
- `.gitignore`, `git_push.sh`, `.travis.yml` — generator scaffolding we don't use

Of these, only `go.mod`, `go.sum`, `README.md`, `.gitignore`, `git_push.sh` and `.travis.yml`
are paths the generator actually emits — those six entries are load-bearing, confirmed by
running the generator into an empty directory with no ignore file. The rest (`auth/`,
`doc.go`, `AGENTS.md`, `CLAUDE.md`, `CODEOWNERS`, `.github/`, `.claude/`, `.work/`) sit at
paths the generator never writes to, so listing them is defensive rather than required; keep
them for clarity and in case that ever changes.

The cross-repo guard that keeps this list in sync with `api-codegen`'s copy compares whole
lines exactly, not semantics — an entry written as `auth/**` instead of `auth/` still
protects the directory from the generator but fails that check, because the text doesn't
match character-for-character.

Regeneration is performed by Monte Carlo's internal API code-generation tooling, run by a
maintainer from outside this repository — it owns generation for every artifact built from
the spec (this SDK, the Terraform provider, the CLI). Consult that tooling directly to
regenerate; it isn't reproduced here because it doesn't ship to consumers of this SDK.

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
same file and the same section names. Most keys are shared verbatim — `mcd_id`, `mcd_token`,
`mcd_oauth_client_id`, `mcd_oauth_client_secret` mean the same thing here as in the CLI and
pycarlo. `mcd_api_endpoint` is shared but reinterpreted: the CLI and pycarlo write the
GraphQL endpoint there, and this SDK strips a trailing `/graphql` before treating it as the
REST base URL. The CLI's `configure` command never writes `mcd_api_endpoint` at all, so a
CLI-only setup still needs `Endpoint` passed explicitly. A customer who has already run the
CLI does not need a separate credential store, which is the point — the endpoint key just
needs that translation to mean the right thing here.

Precedence follows the same three-tier ordering as the Python SDK: values passed in, then
environment variables, then the profile. It is not exact parity — two known divergences:
pycarlo rejects a half-set credential outright (`InvalidSessionError`) rather than filling it
in partially, and this SDK does not read `MCD_API_ENDPOINT` at all, deliberately — see above,
it carries the GraphQL endpoint, not the REST base URL.

Tests must not read the developer's real credentials. `isolate(t)` in the test package clears
every environment variable resolution consults and points `ConfigDir` at a temporary
directory; use it in any test that builds a client or resolves options.

## Branching

Branch from `main` as `<person>/<ticket-id>-<slug>`. Never commit directly to `main`.

## Releasing

**Not yet.** A Go module path is permanent once a version tag is published, and this
repository's name is not final. Fetching by commit sha resolves a pseudo-version and is
enough to verify the module builds for a consumer.

Before this repository goes public, add a `LICENSE` file — there is none today. Without one,
an external consumer of a public Go module has no grant of rights, and pkg.go.dev renders the
module as unlicensed. Add whichever license Monte Carlo uses for public SDKs, and add
`LICENSE` to `.openapi-generator-ignore`'s protected list at the same time — defensive only,
since the generator never emits a `LICENSE` file, but consistent with the rest of that list.
