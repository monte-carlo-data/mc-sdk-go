// Hand-written, not generator output.

package montecarlo

import (
	"errors"
	"fmt"
	"iter"
)

// paginate walks a paginated list and yields every item across every page.
//
// fetch asks for one page. It receives the cursor to request, empty for the first page, and
// returns that page's items, the cursor of the next page, whether more pages follow, and an
// error. The walk ends when fetch reports no more pages, when the consumer stops ranging, or
// on the first error, which is yielded with the zero T as the final pair.
//
// Two malformed responses end the walk with an error rather than looping: more pages promised
// with no cursor to ask for, and a cursor that has already been requested. A server that
// repeats a cursor would otherwise keep the caller in this loop forever. These are the same
// two guards the CLI applies, so the two agree on when a walk is over.
//
// The generated All methods are the only callers.
func paginate[T any](fetch func(cursor string) ([]T, string, bool, error)) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var zero T
		cursor := ""
		requested := map[string]bool{cursor: true}
		for {
			items, next, more, err := fetch(cursor)
			if err != nil {
				yield(zero, err)
				return
			}
			for _, item := range items {
				if !yield(item, nil) {
					return
				}
			}
			if !more {
				return
			}
			if next == "" {
				yield(zero, errors.New("the API reported more items but returned no cursor"))
				return
			}
			if requested[next] {
				yield(zero, fmt.Errorf("the API returned cursor %q twice; stopping", next))
				return
			}
			requested[next] = true
			cursor = next
		}
	}
}
