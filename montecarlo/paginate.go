// Hand-written, not generator output.

package montecarlo

import (
	"errors"
	"iter"
)

// ErrNoCursor ends a walk when the API reports more pages but supplies no cursor to ask for
// the next one.
var ErrNoCursor = errors.New("the API reported more items but returned no cursor")

// ErrRepeatedCursor ends a walk when the API returns a cursor it already served, which would
// otherwise repeat the same page forever.
var ErrRepeatedCursor = errors.New("the API returned the same cursor twice")

// paginate walks a paginated list and yields every item across every page.
//
// fetch asks for one page. It receives the cursor to request, empty for the first page, and
// returns that page's items, the cursor of the next page, whether more pages follow, and an
// error. The walk ends when fetch reports no more pages, when the consumer stops ranging, or
// on the first error, which is yielded with the zero T as the final pair.
//
// Two malformed responses end the walk rather than looping forever: more pages promised with
// no cursor to ask for yields ErrNoCursor, and a cursor already requested yields
// ErrRepeatedCursor. mc-cli's allPages (internal/cmd/paging.go) applies the same two guards, so
// the two agree on when a walk ends, though their error messages differ.
//
// A walk of unknown length should still carry a context deadline. The request context is what
// bounds the total number of requests, since these guards catch a repeated cursor but not a
// server that mints a fresh one forever.
func paginate[T any](fetch func(cursor string) ([]T, string, bool, error)) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var zero T
		cursor := ""
		requested := map[string]bool{}
		for {
			requested[cursor] = true
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
				yield(zero, ErrNoCursor)
				return
			}
			if requested[next] {
				yield(zero, ErrRepeatedCursor)
				return
			}
			cursor = next
		}
	}
}
