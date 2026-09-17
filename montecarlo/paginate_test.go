// Hand-written, not generator output.

package montecarlo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// cannedPage is one response fetchRecorder replays. It can claim more pages than it has a
// cursor for, which the malformed-response tests depend on.
type cannedPage struct {
	items []string
	next  string
	more  bool
}

// fetchRecorder replays a fixed sequence of cannedPages and records the cursor each call
// received, so a test can assert both what a walk yielded and what it asked for. A call past
// the last canned page reports a distinguishable error instead of panicking, so a test that
// over-fetches fails on a named assertion rather than a crash.
type fetchRecorder struct {
	pages []cannedPage
	asked []string
}

func newFetchRecorder(pages ...cannedPage) *fetchRecorder {
	return &fetchRecorder{pages: pages}
}

func (r *fetchRecorder) fetch(cursor string) ([]string, string, bool, error) {
	r.asked = append(r.asked, cursor)
	i := len(r.asked) - 1
	if i >= len(r.pages) {
		return nil, "", false, fmt.Errorf("fetchRecorder: fetch called %d times, only %d pages canned", len(r.asked), len(r.pages))
	}
	p := r.pages[i]
	return p.items, p.next, p.more, nil
}

// collectedPair is one item, err pair an iterator yielded, in yield order.
type collectedPair[T any] struct {
	item T
	err  error
}

// collect drains an iterator into the full sequence of pairs it yielded, in order.
func collect[T any](seq func(func(T, error) bool)) []collectedPair[T] {
	var pairs []collectedPair[T]
	for item, err := range seq {
		pairs = append(pairs, collectedPair[T]{item, err})
	}
	return pairs
}

// itemsAndErr extracts a walk's items and its final error, for tests that only care about the
// common case rather than the full pair-by-pair contract. TestPaginateEndsOnAFailedRequest pins
// that fuller contract directly against collect's own return value.
func itemsAndErr[T any](pairs []collectedPair[T]) ([]T, error) {
	var items []T
	var err error
	for _, p := range pairs {
		if p.err != nil {
			err = p.err
			continue
		}
		items = append(items, p.item)
	}
	return items, err
}

func TestPaginateYieldsEveryPageInOrder(t *testing.T) {
	rec := newFetchRecorder(
		cannedPage{items: []string{"a", "b"}, next: "c1", more: true},
		cannedPage{items: []string{"c"}, next: "c2", more: true},
		cannedPage{items: []string{"d", "e"}},
	)

	items, err := itemsAndErr(collect(paginate(rec.fetch)))

	if err != nil {
		t.Fatalf("walk failed: %v", err)
	}
	if got := strings.Join(items, ""); got != "abcde" {
		t.Errorf("items = %q, want %q", got, "abcde")
	}
	// The first page is asked for with no cursor; each later page with the one before it.
	if got := strings.Join(rec.asked, ","); got != ",c1,c2" {
		t.Errorf("cursors asked = %q, want %q", got, ",c1,c2")
	}
}

func TestPaginateStopsWhenTheConsumerStops(t *testing.T) {
	calls := 0
	fetch := func(string) ([]string, string, bool, error) {
		calls++
		return []string{"a", "b"}, "next", true, nil
	}

	for range paginate(fetch) {
		break
	}

	// Breaking out of the range must not fetch the page after the one being read.
	if calls != 1 {
		t.Errorf("fetch called %d times, want 1", calls)
	}
}

func TestPaginateEndsOnAFailedRequest(t *testing.T) {
	boom := errors.New("boom")
	first := true
	fetch := func(string) ([]string, string, bool, error) {
		if first {
			first = false
			return []string{"a"}, "c1", true, nil
		}
		return nil, "", false, boom
	}

	pairs := collect(paginate(fetch))

	if len(pairs) != 2 {
		t.Fatalf("pairs = %d, want 2: %+v", len(pairs), pairs)
	}
	if pairs[0].item != "a" || pairs[0].err != nil {
		t.Errorf("pairs[0] = %+v, want {item: \"a\", err: nil}", pairs[0])
	}
	last := pairs[1]
	if !errors.Is(last.err, boom) {
		t.Errorf("final error = %v, want %v", last.err, boom)
	}
	var zero string
	if last.item != zero {
		t.Errorf("final item = %q, want the zero value", last.item)
	}
}

func TestPaginateRefusesMorePagesWithNoCursor(t *testing.T) {
	rec := newFetchRecorder(cannedPage{items: []string{"a"}, more: true})

	_, err := itemsAndErr(collect(paginate(rec.fetch)))

	if !errors.Is(err, ErrNoCursor) {
		t.Errorf("error = %v, want %v", err, ErrNoCursor)
	}
}

func TestPaginateRefusesARepeatedCursor(t *testing.T) {
	fetch := func(string) ([]string, string, bool, error) {
		return []string{"a"}, "same", true, nil
	}

	_, err := itemsAndErr(collect(paginate(fetch)))

	if !errors.Is(err, ErrRepeatedCursor) {
		t.Errorf("error = %v, want %v", err, ErrRepeatedCursor)
	}
}

// warehousePage is one page of the warehouses list as the API serves it. Every field the
// generated models require is present, so a page that unmarshals here would unmarshal in
// production too. hasMore and next are independent, so a caller can render the malformed
// shape ErrNoCursor exists for: more pages promised with a null cursor.
func warehousePage(ids []string, next string, hasMore bool) string {
	items := make([]map[string]any, 0, len(ids))
	for _, id := range ids {
		items = append(items, map[string]any{
			"id":            id,
			"name":          "warehouse-" + id,
			"type":          string(WAREHOUSETYPE_BIGQUERY),
			"deployment_id": "d1",
			"created_time":  "2026-09-17T00:00:00Z",
		})
	}
	body := map[string]any{"items": items, "has_more": hasMore, "count": nil}
	if next != "" {
		body["next_cursor"] = next
	} else {
		body["next_cursor"] = nil
	}
	// The value marshaled above is a fixed map of strings, bools, and nil, so json.Marshal
	// cannot fail.
	encoded, _ := json.Marshal(body)
	return string(encoded)
}

// The closure api-codegen renders threads the cursor onto the real request and asks for the
// operation's largest page. This exercises the generated All method rather than the driver
// alone.
func TestAllWalksEveryPageThroughTheGeneratedClient(t *testing.T) {
	var queries []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries = append(queries, r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Query().Get("cursor") {
		case "":
			fmt.Fprint(w, warehousePage([]string{"1", "2"}, "c1", true))
		case "c1":
			fmt.Fprint(w, warehousePage([]string{"3"}, "c2", true))
		default:
			fmt.Fprint(w, warehousePage([]string{"4"}, "", false))
		}
	}))
	defer server.Close()

	ctx := context.Background()
	api, err := NewClient(ctx, Options{
		Endpoint:    server.URL,
		TokenID:     "id",
		TokenSecret: "secret",
		ConfigDir:   isolate(t),
	})
	if err != nil {
		t.Fatalf("building the client failed: %v", err)
	}

	var ids []string
	for warehouse, err := range api.WarehousesAPI.ListWarehouses(ctx).All() {
		if err != nil {
			t.Fatalf("walk failed: %v", err)
		}
		ids = append(ids, warehouse.Id)
	}

	if got := strings.Join(ids, ","); got != "1,2,3,4" {
		t.Errorf("ids = %q, want %q", got, "1,2,3,4")
	}
	if len(queries) != 3 {
		t.Fatalf("requests = %d, want 3: %v", len(queries), queries)
	}
	// The first request carries no cursor; each later one carries the page before it. Every
	// request asks for the largest page the operation serves, so a walk costs the fewest calls.
	//
	// with_count arrives as false rather than absent: the generated client substitutes the
	// spec's default for a nil pointer instead of omitting the parameter. False is what keeps
	// the server from counting, which is the point of clearing it.
	for i, want := range []string{
		"limit=100&with_count=false",
		"cursor=c1&limit=100&with_count=false",
		"cursor=c2&limit=100&with_count=false",
	} {
		if queries[i] != want {
			t.Errorf("request %d query = %q, want %q", i, queries[i], want)
		}
	}
}

// api_warehouses_paging.gen.go's All doc comment claims a page size set with Limit overrides
// the largest-page default; this pins that claim.
func TestAllKeepsAPageSizeTheCallerSet(t *testing.T) {
	var queries []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries = append(queries, r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, warehousePage([]string{"1"}, "", false))
	}))
	defer server.Close()

	ctx := context.Background()
	api, err := NewClient(ctx, Options{
		Endpoint:    server.URL,
		TokenID:     "id",
		TokenSecret: "secret",
		ConfigDir:   isolate(t),
	})
	if err != nil {
		t.Fatalf("building the client failed: %v", err)
	}

	for _, err := range api.WarehousesAPI.ListWarehouses(ctx).Limit(25).All() {
		if err != nil {
			t.Fatalf("walk failed: %v", err)
		}
	}

	if want := "limit=25&with_count=false"; len(queries) != 1 || queries[0] != want {
		t.Errorf("queries = %v, want [%q]", queries, want)
	}
}

// GetNextCursor on the generated model turns a null next_cursor into an empty string before
// the driver ever sees it, so a has_more: true page with no cursor reaches paginate as
// ErrNoCursor's own trigger. This drives that shape through the real generated closure rather
// than asserting it against the driver alone.
func TestAllEndsWithErrNoCursorWhenTheServerPromisesMoreWithNoCursor(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, warehousePage([]string{"1"}, "", true))
	}))
	defer server.Close()

	ctx := context.Background()
	api, err := NewClient(ctx, Options{
		Endpoint:    server.URL,
		TokenID:     "id",
		TokenSecret: "secret",
		ConfigDir:   isolate(t),
	})
	if err != nil {
		t.Fatalf("building the client failed: %v", err)
	}

	var walkErr error
	for _, err := range api.WarehousesAPI.ListWarehouses(ctx).All() {
		if err != nil {
			walkErr = err
		}
	}
	if !errors.Is(walkErr, ErrNoCursor) {
		t.Errorf("error = %v, want %v", walkErr, ErrNoCursor)
	}
}

// A regeneration that dropped All's `if err != nil { return }` would leave a walk silently
// truncating on a mid-walk 500 instead of surfacing it, and nothing else in this suite drives a
// real failed request through the generated closure to catch that.
func TestAllYieldsAFailedRequestAsItsFinalPair(t *testing.T) {
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls == 1 {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, warehousePage([]string{"1"}, "c1", true))
			return
		}
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer server.Close()

	ctx := context.Background()
	api, err := NewClient(ctx, Options{
		Endpoint:    server.URL,
		TokenID:     "id",
		TokenSecret: "secret",
		ConfigDir:   isolate(t),
	})
	if err != nil {
		t.Fatalf("building the client failed: %v", err)
	}

	var ids []string
	var walkErr error
	for warehouse, err := range api.WarehousesAPI.ListWarehouses(ctx).All() {
		if err != nil {
			walkErr = err
			continue
		}
		ids = append(ids, warehouse.Id)
	}

	if got := strings.Join(ids, ","); got != "1" {
		t.Errorf("ids = %q, want %q", got, "1")
	}
	if walkErr == nil {
		t.Error("error = nil, want the 500 surfaced as the walk's final pair")
	}
	if calls != 2 {
		t.Errorf("requests = %d, want 2: the walk should stop at the failure", calls)
	}
}

func TestPaginateServesAnEmptyList(t *testing.T) {
	rec := newFetchRecorder(cannedPage{})

	items, err := itemsAndErr(collect(paginate(rec.fetch)))

	if err != nil {
		t.Fatalf("walk failed: %v", err)
	}
	if len(items) != 0 {
		t.Errorf("items = %v, want none", items)
	}
	if len(rec.asked) != 1 {
		t.Errorf("fetch called %d times, want 1", len(rec.asked))
	}
}
