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

    "github.com/monte-carlo-data/mc-sdk-go/auth"
)

api, err := auth.NewClient(ctx, auth.Options{
    Endpoint:    "https://api.getmontecarlo.com",
    TokenID:     os.Getenv("MCD_ID"),
    TokenSecret: os.Getenv("MCD_TOKEN"),
})
if err != nil {
    return err
}

me, _, err := api.UsersAPI.GetCurrentUser(ctx).Execute()
```

With OAuth 2.0 client credentials, where `Instance` selects the deployment the gateway
routes to:

```go
api, err := auth.NewClient(ctx, auth.Options{
    Endpoint:     "https://api.getmontecarlo.com",
    ClientID:     os.Getenv("MCD_CLIENT_ID"),
    ClientSecret: os.Getenv("MCD_CLIENT_SECRET"),
    Instance:     "us1",
})
```

Tokens are fetched and refreshed as needed, so a long-lived client does not go stale.

## What is generated and what is not

| Path | |
|---|---|
| `api_*.go`, `model_*.go`, `client.go`, `configuration.go`, `utils.go` | Generated. Do not edit. |
| `auth/` | Hand-written. Authentication and client construction. |
| `docs/` | Generated API reference. |

Generated files are replaced wholesale on every run, so anything hand-written is listed in
`.openapi-generator-ignore`. Adding a file outside that list means losing it on the next
regeneration, without warning.

## Contributing

The generated files track the API and are not edited by hand — a fix to one of them belongs
in the API or in the generator, so please open an issue describing what is wrong rather than
a patch. Changes to `auth/` are welcome as pull requests.
