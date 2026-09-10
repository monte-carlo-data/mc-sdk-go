// Hand-written, not generator output.

package montecarlo_test

import (
	"context"

	"github.com/monte-carlo-data/mc-sdk-go/montecarlo"
)

// ExampleNewClient compiles the public surface the way a consumer does: one import, one
// package. It has no Output comment, so go test compiles it and never runs it — the point is
// that a rename of NewClient or Options fails here rather than in a downstream repo. It is also
// the usage example pkg.go.dev renders.
func ExampleNewClient() {
	api, err := montecarlo.NewClient(context.Background(), montecarlo.Options{
		Endpoint:    "https://api.getmontecarlo.com",
		TokenID:     "id",
		TokenSecret: "secret",
	})
	_, _ = api, err
}
