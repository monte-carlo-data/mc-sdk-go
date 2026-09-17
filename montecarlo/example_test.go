// Hand-written, not generator output.

package montecarlo_test

import (
	"context"
	"fmt"

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

// ExampleAPIClient_walk ranges over a paginated list without tracking cursors itself. Every
// generated list operation's All method has this same shape. Like ExampleNewClient, it has no
// Output comment, so a rename of All fails here rather than downstream.
func ExampleAPIClient_walk() {
	ctx := context.Background()
	api, err := montecarlo.NewClient(ctx, montecarlo.Options{
		Endpoint:    "https://api.getmontecarlo.com",
		TokenID:     "id",
		TokenSecret: "secret",
	})
	if err != nil {
		return
	}

	for warehouse, err := range api.WarehousesAPI.ListWarehouses(ctx).All() {
		if err != nil {
			return
		}
		fmt.Println(warehouse.GetName())
	}
}
