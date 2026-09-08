# mc-sdk-go

Official Go SDK for the Monte Carlo REST API.

Most of this repository is generated from the API's OpenAPI spec. The `auth` package is
hand-written: it builds a client that authenticates every request.

> For the GraphQL API, use [pycarlo](https://pypi.org/project/pycarlo/) instead.
> This SDK targets the REST API.

## Install

```bash
go get github.com/monte-carlo-data/mc-sdk-go
```

While this repository is internal, that needs `GOPRIVATE=github.com/monte-carlo-data` and a
git credential with access, since the public module proxy cannot resolve it. No version has
been published yet either, so `go get` resolves a pseudo-version from a commit.

## Usage

Every client needs an endpoint. There is no default: the API is reached at a different host
per instance.

With a Monte Carlo API token — supply the two halves separately, and the SDK forms the
credential:

```go
import (
    "context"
    "os"

    "github.com/monte-carlo-data/mc-sdk-go/auth"
)

func Example(ctx context.Context) error {
    api, err := auth.NewClient(ctx, auth.Options{
        Endpoint:    "https://api.getmontecarlo.com",
        TokenID:     os.Getenv("MCD_ID"),
        TokenSecret: os.Getenv("MCD_TOKEN"),
    })
    if err != nil {
        return err
    }

    me, _, err := api.UsersAPI.GetCurrentUser(ctx).Execute()
    return err
}
```

With OAuth 2.0 client credentials, where `Instance` selects the deployment the gateway
routes to:

```go
import (
    "context"
    "os"

    "github.com/monte-carlo-data/mc-sdk-go/auth"
)

func Example(ctx context.Context) error {
    api, err := auth.NewClient(ctx, auth.Options{
        Endpoint:     "https://api.getmontecarlo.com",
        ClientID:     os.Getenv("MCD_CLIENT_ID"),
        ClientSecret: os.Getenv("MCD_CLIENT_SECRET"),
        Instance:     "us1",
    })
    if err != nil {
        return err
    }

    me, _, err := api.UsersAPI.GetCurrentUser(ctx).Execute()
    return err
}
```

Tokens are fetched and refreshed as needed, so a long-lived client does not go stale.

### Using credentials you have already configured

If you have configured the Monte Carlo CLI, the SDK reads the same `mcd_id` and `mcd_token`
from `~/.mcd/profiles.ini` — but you still need to pass `Endpoint` yourself, since the CLI's
`configure` command never writes one:

```go
api, err := auth.NewClient(ctx, auth.Options{Endpoint: "https://api.getmontecarlo.com"})
```

If your profile also has an `mcd_api_endpoint` key — set when the CLI or pycarlo point at a
non-default deployment — this SDK reads that too, but reinterprets it: in the CLI and pycarlo
that key holds the GraphQL endpoint (`.../graphql`), and this SDK strips a trailing `/graphql`
before using it as the REST base URL. No action needed; noted here so an existing profile
does not surprise you.

Anything you leave unset is filled in, in this order:

1. What you pass in `Options`.
2. `MCD_DEFAULT_API_ID` and `MCD_DEFAULT_API_TOKEN`, or `MCD_DEFAULT_OAUTH_CLIENT_ID` and
   `MCD_DEFAULT_OAUTH_CLIENT_SECRET`, and `MCD_DEFAULT_INSTANCE_ID`.
3. A profile from `~/.mcd/profiles.ini`, which the CLI writes.

This SDK does not read `MCD_API_ENDPOINT`, even though the CLI and pycarlo both do — that
variable carries the GraphQL endpoint, and `Endpoint` here is the REST base URL, so the two
are not interchangeable (see above).

To pick a profile other than `default`, set `Profile`, or the `MCD_DEFAULT_PROFILE`
environment variable. A profile you name explicitly has to exist — the SDK reports that
rather than quietly falling back to different credentials.

## What is generated and what is not

| Path | |
|---|---|
| `api_*.go`, `model_*.go`, `client.go`, `configuration.go`, `response.go`, `utils.go` | Generated. Do not edit. |
| `api/openapi.yaml` | Generated copy of the input spec — not the source of truth, and overwritten on every run. Edit the source spec instead. |
| `.openapi-generator/` | Generator bookkeeping. |
| `auth/` | Hand-written. Authentication and client construction. |
| `docs/` | Generated API reference — see [API reference](#api-reference) below. |

Generated files are overwritten wholesale on every run, so anything hand-written is listed in
`.openapi-generator-ignore`. Adding a file outside that list risks having it silently
overwritten by generator output the next time the tree regenerates — the generator does not
delete files, but it does replace any path it emits.

## API reference

`docs/` holds one page per operation and per schema, with no index page. The four operation
groups:

- [`CollectionAgentsAPI`](docs/CollectionAgentsAPI.md)
- [`CollectionDataStoresAPI`](docs/CollectionDataStoresAPI.md)
- [`DeploymentsAPI`](docs/DeploymentsAPI.md)
- [`UsersAPI`](docs/UsersAPI.md)

Each generated page's "All URIs are relative to *http://localhost*" line is a placeholder
left by the spec's empty default server URL, not a real base URL — the real one is whatever
`Endpoint` you pass to `auth.NewClient`.

## Contributing

The generated files track the API and are not edited by hand — a fix to one of them belongs
in the API or in the generator, so please open an issue describing what is wrong rather than
a patch. Changes to `auth/` are welcome as pull requests.
