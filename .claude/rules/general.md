---
description: Coding standards for this repo
globs: "**/*"
---

# General Standards

Almost everything under `montecarlo/` is generated output. The hand-written surface is
`montecarlo/auth.go`, `oauth.go`, `profile.go`, `paginate.go` and their tests,
`montecarlo/doc.go`, the module files, and the repository documentation — see AGENTS.md for
the exact list. `montecarlo/api_*_paging.gen.go` is generated even though it calls into the
hand-written `paginate.go` — the pairing is easy to misread, so check AGENTS.md before assuming
either half is what it looks like. These rules apply to the hand-written surface; do not
hand-edit generated files, because the next generation run overwrites them.

## Code Style

- Keep functions small and focused on a single responsibility
- Prefer explicit over implicit — name things clearly
- Write code that reads like a specification
- Exported identifiers need a doc comment beginning with the identifier's own name — that is what
  makes `go doc` and pkg.go.dev read correctly

## Testing

- Tests are co-located as `<file>_test.go` beside the source, per Go convention, and because the
  package's own tests reach unexported identifiers
- Write tests for all non-trivial logic
- The hand-written code's boundaries are the filesystem and an OAuth endpoint. Use
  `t.TempDir()` and `httptest.NewServer` rather than mocks
- Any test that builds a client or resolves options must call `isolate(t)` and pass the directory
  it returns as `Options.ConfigDir`, so the test cannot read the developer's real credentials
- Use `t.Setenv` rather than saving and restoring environment variables by hand. It is incompatible
  with `t.Parallel()`, so tests here do not run in parallel
- Name a test for the invariant it pins, not for the function it calls

## Error Handling

- Handle errors at system boundaries — the credentials file, the OAuth exchange, anything a caller
  supplies
- Wrap with `%w` and match with `errors.Is`/`errors.As`; do not compare error strings
- Only tolerate an error you have identified. Discarding every error from a call because one of its
  failure modes is expected turns a readable file with bad permissions into "no credentials"
- Don't add error handling for scenarios that can't happen

## Verification

`go build ./...`, `go vet ./...`, `gofmt -l .` (must be empty), `go test -race ./...`, and
`.github/scripts/next-tag_test.sh`.
