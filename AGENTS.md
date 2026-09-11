# mc-sdk-go

> The Go SDK for the Monte Carlo REST API: one package, `montecarlo`. Generated from the API's OpenAPI spec, except client construction and credentials (`auth.go`, `oauth.go`, `profile.go`), which are hand-written.

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
| `montecarlo/` | The whole SDK and the generator's output directory: package `montecarlo`, import path `github.com/monte-carlo-data/mc-sdk-go/montecarlo` |
| `montecarlo/auth.go`, `oauth.go`, `profile.go` | **Hand-written.** `NewClient`, `Options`, API-token and OAuth credentials, the CLI's profiles file |
| `montecarlo/example_test.go` | **Hand-written.** External-package example that compiles the public surface as a consumer would |
| `montecarlo/doc.go` | **Hand-written.** The package doc comment |
| `montecarlo/api_*.go`, `montecarlo/model_*.go` | Generated operations and models, one file per tag group and schema |
| `montecarlo/client.go`, `configuration.go`, `response.go`, `utils.go` | Generated client plumbing |
| `montecarlo/.openapi-generator/` | Generator bookkeeping |
| `montecarlo/docs/` | Generated API reference |

## What is generated

Everything under `montecarlo/` except `doc.go`, `example_test.go`, and the auth files. The generator is pointed at
that directory, not the repository root, so the module files and the repository documentation
are out of its reach by construction. The hand-written files share the package with the
generated ones so that a caller has one import and one name for the SDK — `montecarlo.NewClient`
beside `montecarlo.DeploymentIn` — and rely on the ignore file for protection, exactly as
`doc.go` always has. The generator overwrites every path it emits on each run — it does not
delete anything else — so a fix to a generated file does not survive; it belongs in the API or
in the generator that reads its spec.

The Go generator emits a flat package: it has no option to nest operations or models in
subdirectories (its `apiPackage`/`modelPackage` settings are ignored), and Go's one package
per directory would in any case require cross-package imports it does not produce. The
directory is the unit of organisation, which is why the output moved there.

**Anything hand-written under `montecarlo/`, or anything we don't want emitted there, must be
listed in `montecarlo/.openapi-generator-ignore`** — the generator reads the ignore file from
its output directory. `montecarlo/.openapi-generator-ignore` is what `generate-go-sdk.sh`
compares its `PROTECTED` array against, whole line for whole line; this section documents that
file entry for entry, so the two must be kept in step by hand:

- `auth.go`, `auth_test.go`, `oauth.go`, `oauth_test.go`, `profile.go`, `profile_test.go` — hand-written client construction and credentials
- `doc.go` — the package doc comment, which is the pkg.go.dev page for the package
- `example_test.go` — the external-package (`package montecarlo_test`) example that compiles the
  public surface as a consumer would; hand-written, and must stay in the ignore file
- `go.mod`, `go.sum` — a nested module would split the package off from the repository's
- `api/openapi.yaml` — suppressed, not protected: the generator's vendored copy of the input spec, which nothing reads and which would publish the input's internal `x-mc-*` markers
- `README.md`, `.gitignore` — the generator's copies duplicate the root ones
- `git_push.sh`, `.travis.yml` — generator scaffolding we don't use

Of these, `README.md`, `.gitignore`, `git_push.sh`, `.travis.yml` and `api/openapi.yaml` are
paths the generator actually emits — those five entries are load-bearing, confirmed by running
the generator into an empty directory with no ignore file. `go.mod` and `go.sum` would be too,
but the generation script passes `withGoMod=false` so they are never written; `doc.go`,
`example_test.go`, and the auth files are paths the Go generator never writes to — it names its
files `api_*.go`, `model_*.go`, `client.go`, `configuration.go`, `response.go` and `utils.go`,
and runs with test generation off. The rest of the list is defensive rather than required; keep it for clarity
and in case that ever changes. A hand-written file must never take one of those generated
names. The same constraint applies to identifiers: the generator emits `<Schema>`,
`New<Schema>`, `New<Schema>WithDefaults`, and `NewNullable<Schema>` for every schema in the
spec, plus fixed names `APIClient`, `Configuration`, `NewConfiguration`, `NewAPIClient`,
`GenericOpenAPIError`, `BasicAuth`, `APIKey`, `ServerConfiguration`, and the `Nullable*`/`Ptr*`
helper families. A hand-written package-level identifier — exported or not — must never take a
name the generator could emit; in particular `NewClient` and `Options` are reserved against
future schemas named `Client` and `Options`, and a collision surfaces as a compile failure at
regeneration time whose only fix is renaming the SDK's public API.

The generation script also passes `isGoSubmodule=true`, which is what makes the import path in
the generated `docs/` examples read `.../mc-sdk-go/montecarlo` rather than the module root.

The cross-repo guard compares `api-codegen`'s `PROTECTED` array against
`montecarlo/.openapi-generator-ignore` whole line for whole line, not semantically — an entry
written as `./doc.go` instead of `doc.go` still protects the file from the generator but fails
that check, because the text doesn't match character-for-character. Since the guard reads the
ignore file directly rather than this section, keeping the two in step is a manual discipline,
not something enforced by the check itself.

Regeneration is performed by Monte Carlo's internal API code-generation tooling, run by a
maintainer from outside this repository — it owns generation for every artifact built from
the spec (this SDK, the Terraform provider, the CLI). Consult that tooling directly to
regenerate; it isn't reproduced here because it doesn't ship to consumers of this SDK.

Constructor parameters and struct fields follow the order the spec declares a schema's
properties in. An export that changes that order changes public signatures with no schema
change behind it, and two same-typed parameters swapping is invisible to a consumer's
compiler. Every artifact generated from the spec, this SDK, the Terraform provider and the
CLI, is regenerated from the same export for that reason. A consumer that moves its pin
across such a regeneration rechecks every positional constructor call.

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
environment variables, then the profile. Half of a credential pair passed in is an error before
the environment or the profile is consulted, as in pycarlo and mc-sdk-python; it is never
filled in from a profile. One known divergence from pycarlo: this SDK does not read
`MCD_API_ENDPOINT` at all, deliberately — see above, it carries the GraphQL endpoint, not the
REST base URL.

The OAuth token exchange never follows a redirect. The request carries the client secret in
its form body, so a redirect is reported as a failed exchange rather than re-sent to the host
the `Location` header names. A caller who supplies their own `http.Client` through the
`oauth2.HTTPClient` context value owns that client's redirect policy.

Tests must not read the developer's real credentials. `isolate(t)` in the test package clears
every environment variable resolution consults and points `ConfigDir` at a temporary
directory; use it in any test that builds a client or resolves options.

## Branching

Branch from `main` as `<person>/<ticket-id>-<slug>`. Never commit directly to `main`.

## Releasing

**Not yet.** A Go module path is permanent once a version tag is published, and this
repository's name is not final. Fetching by commit sha resolves a pseudo-version and is
enough to verify the module builds for a consumer.

The licence is Apache-2.0, matching the Python SDK and the CLI. That is the SDK precedent
rather than the agents' one: the agents ship as deployed artifacts under a proprietary licence,
while an SDK is distributed as source — `proxy.golang.org` mirrors every public module — and
pkg.go.dev only renders documentation for a licence it recognises.
