# mc-sdk-go

Official Go SDK for the Monte Carlo REST API.

The SDK is one package, `montecarlo`. Most of it is generated from the API's OpenAPI spec;
client construction and credentials are hand-written, so every request is authenticated.

> For the GraphQL API, use [pycarlo](https://pypi.org/project/pycarlo/) instead.
> This SDK targets the REST API.

## Install

```bash
go get github.com/monte-carlo-data/mc-sdk-go/montecarlo
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

    "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
)

func Example(ctx context.Context) error {
    api, err := montecarlo.NewClient(ctx, montecarlo.Options{
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

    "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
)

func Example(ctx context.Context) error {
    api, err := montecarlo.NewClient(ctx, montecarlo.Options{
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

The client's operations and the request and response types they take are in the same package,
so that one import covers everything: `montecarlo.DeploymentIn`, `montecarlo.ProblemOut`, and
so on.

### Using credentials you have already configured

If you have configured the Monte Carlo CLI, the SDK reads the same `mcd_id` and `mcd_token`
from `~/.mcd/profiles.ini` — but you still need to pass `Endpoint` yourself, since the CLI's
`configure` command never writes one:

```go
api, err := montecarlo.NewClient(ctx, montecarlo.Options{Endpoint: "https://api.getmontecarlo.com"})
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
| `montecarlo/api_*.go`, `model_*.go`, `client.go`, `configuration.go`, `response.go`, `utils.go` | Generated. Do not edit. The API client, one `api_*.go` per operation group and one `model_*.go` per schema. |
| `montecarlo/docs/` | Generated API reference — see [API reference](#api-reference) below. |
| `montecarlo/.openapi-generator/` | Generator bookkeeping. |
| `montecarlo/auth.go`, `oauth.go`, `profile.go`, `example_test.go` | Hand-written. `NewClient`, `Options`, the credential resolution behind them, and a usage example that go test compiles as a consumer would. |
| `montecarlo/doc.go` | Hand-written. The package's doc comment. |

The generator writes only under `montecarlo/`, and overwrites every file it emits there on
every run. The hand-written files share that directory and are listed in
`montecarlo/.openapi-generator-ignore`; adding a file there outside that list risks having it
silently overwritten the next time the tree regenerates — the generator does not delete files,
but it does replace any path it emits.

## API reference

`montecarlo/docs/` holds one page per operation and per schema, with no index page. The four
operation groups:

- [`CollectionAgentsAPI`](montecarlo/docs/CollectionAgentsAPI.md)
- [`CollectionDataStoresAPI`](montecarlo/docs/CollectionDataStoresAPI.md)
- [`DeploymentsAPI`](montecarlo/docs/DeploymentsAPI.md)
- [`UsersAPI`](montecarlo/docs/UsersAPI.md)

Each generated page's "All URIs are relative to *http://localhost*" line is a placeholder
left by the spec's empty default server URL, not a real base URL — the real one is whatever
`Endpoint` you pass to `montecarlo.NewClient`.

## Contributing

The generated files track the API and are not edited by hand — a fix to one of them belongs
in the API or in the generator, so please open an issue describing what is wrong rather than
a patch. Changes to the hand-written files (`auth.go`, `oauth.go`, `profile.go` and their
tests) are welcome as pull requests.
