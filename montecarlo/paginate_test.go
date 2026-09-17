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

// page is one canned response for a fake fetch: the items it serves, the cursor it reports as
// next, and whether it claims more pages follow.
type page struct {
	items []string
	next  string
	more  bool
}

// fakeFetch serves pages in order and records the cursor each call received, so a test can
// assert both what a walk yielded and what it asked for.
func fakeFetch(pages ...page) (func(string) ([]string, string, bool, error), *[]string) {
	var asked []string
	i := 0
	return func(cursor string) ([]string, string, bool, error) {
		asked = append(asked, cursor)
		p := pages[i]
		i++
		return p.items, p.next, p.more, nil
	}, &asked
}

// collect drains an iterator into its items and the first error it yielded.
func collect[T any](seq func(func(T, error) bool)) ([]T, error) {
	var items []T
	var failure error
	for item, err := range seq {
		if err != nil {
			failure = err
			continue
		}
		items = append(items, item)
	}
	return items, failure
}

func TestPaginateYieldsEveryPageInOrder(t *testing.T) {
	fetch, asked := fakeFetch(
		page{items: []string{"a", "b"}, next: "c1", more: true},
		page{items: []string{"c"}, next: "c2", more: true},
		page{items: []string{"d", "e"}},
	)

	items, err := collect(paginate(fetch))

	if err != nil {
		t.Fatalf("walk failed: %v", err)
	}
	if got := strings.Join(items, ""); got != "abcde" {
		t.Errorf("items = %q, want %q", got, "abcde")
	}
	// The first page is asked for with no cursor; each later page with the one before it.
	if got := strings.Join(*asked, ","); got != ",c1,c2" {
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

	items, err := collect(paginate(fetch))

	if !errors.Is(err, boom) {
		t.Errorf("error = %v, want %v", err, boom)
	}
	// The items read before the failure are still yielded.
	if len(items) != 1 || items[0] != "a" {
		t.Errorf("items = %v, want [a]", items)
	}
}

func TestPaginateRefusesMorePagesWithNoCursor(t *testing.T) {
	fetch, _ := fakeFetch(page{items: []string{"a"}, more: true})

	_, err := collect(paginate(fetch))

	if err == nil || !strings.Contains(err.Error(), "no cursor") {
		t.Errorf("error = %v, want one naming the missing cursor", err)
	}
}

func TestPaginateRefusesARepeatedCursor(t *testing.T) {
	fetch := func(string) ([]string, string, bool, error) {
		return []string{"a"}, "same", true, nil
	}

	_, err := collect(paginate(fetch))

	if err == nil || !strings.Contains(err.Error(), "twice") {
		t.Errorf("error = %v, want one naming the repeated cursor", err)
	}
}

func TestPaginateRefusesTheFirstCursorComingBack(t *testing.T) {
	// The first page is requested with an empty cursor, so a server echoing that back is the
	// same loop by another name.
	fetch, _ := fakeFetch(page{items: []string{"a"}, next: "", more: true})

	_, err := collect(paginate(fetch))

	if err == nil {
		t.Error("error = nil, want the walk refused")
	}
}

// warehousePage is one page of the warehouses list as the API serves it. Every field the
// generated models require is present, so a page that unmarshals here would unmarshal in
// production too.
func warehousePage(t *testing.T, ids []string, next string) string {
	t.Helper()
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
	body := map[string]any{"items": items, "has_more": next != "", "count": nil}
	if next != "" {
		body["next_cursor"] = next
	} else {
		body["next_cursor"] = nil
	}
	encoded, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("encoding the page failed: %v", err)
	}
	return string(encoded)
}

// TestAllWalksEveryPageThroughTheGeneratedClient exercises the generated All method rather than
// the driver alone: it proves the closure api-codegen renders threads the cursor onto the real
// request, asks for the operation's largest page, and stops where the driver says to.
func TestAllWalksEveryPageThroughTheGeneratedClient(t *testing.T) {
	var queries []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries = append(queries, r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Query().Get("cursor") {
		case "":
			fmt.Fprint(w, warehousePage(t, []string{"1", "2"}, "c1"))
		case "c1":
			fmt.Fprint(w, warehousePage(t, []string{"3"}, "c2"))
		default:
			fmt.Fprint(w, warehousePage(t, []string{"4"}, ""))
		}
	}))
	defer server.Close()

	ctx := context.Background()
	api, err := NewClient(ctx, Options{
		Endpoint:    server.URL,
		TokenID:     "id",
		TokenSecret: "secret",
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

// TestAllKeepsAPageSizeTheCallerSet pins the published doc comment's claim that Limit overrides
// the largest-page default.
func TestAllKeepsAPageSizeTheCallerSet(t *testing.T) {
	var queries []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queries = append(queries, r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, warehousePage(t, []string{"1"}, ""))
	}))
	defer server.Close()

	ctx := context.Background()
	api, err := NewClient(ctx, Options{Endpoint: server.URL, TokenID: "id", TokenSecret: "secret"})
	if err != nil {
		t.Fatalf("building the client failed: %v", err)
	}

	for _, err := range api.WarehousesAPI.ListWarehouses(ctx).Limit(25).All() {
		if err != nil {
			t.Fatalf("walk failed: %v", err)
		}
	}

	if len(queries) != 1 || !strings.Contains(queries[0], "limit=25") {
		t.Errorf("queries = %v, want one carrying limit=25", queries)
	}
}

func TestPaginateServesAnEmptyList(t *testing.T) {
	fetch, asked := fakeFetch(page{})

	items, err := collect(paginate(fetch))

	if err != nil {
		t.Fatalf("walk failed: %v", err)
	}
	if len(items) != 0 {
		t.Errorf("items = %v, want none", items)
	}
	if len(*asked) != 1 {
		t.Errorf("fetch called %d times, want 1", len(*asked))
	}
}
