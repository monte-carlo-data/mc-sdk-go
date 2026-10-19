#!/usr/bin/env bash
# Tests github-release.sh against a fake gh that logs its calls and serves canned releases and tags.
# shellcheck disable=SC2329 # the helpers run through check, by name
set -euo pipefail
# Ignore the developer's git config (e.g. commit signing) so cases never prompt or fail on it.
export GIT_CONFIG_GLOBAL=/dev/null GIT_CONFIG_NOSYSTEM=1

script=$(cd "$(dirname "$0")" && pwd)/github-release.sh
failures=0
root=$(mktemp -d)
trap 'rm -rf "$root"' EXIT
calls=$root/calls
export CALLS=$calls

mkdir "$root/bin"
cat > "$root/bin/gh" <<'EOF'
#!/usr/bin/env bash
echo "${*//$'\n'/ }" >> "$CALLS"
filter=
prev=
for a in "$@"; do
  if [ "$prev" = --jq ]; then filter=$a; fi
  prev=$a
done
case "$*" in
  "api repos/o/r/releases --paginate"*)
    if [ -n "${STUB_RELEASES_FAIL:-}" ]; then echo "gh: HTTP 500" >&2; exit 1; fi
    printf '%s' "${STUB_RELEASES:-[]}" | jq -r "$filter" ;;
  "api repos/o/r/git/matching-refs/"*)
    if [ -n "${STUB_REFS_FAIL:-}" ]; then echo "gh: HTTP 500" >&2; exit 1; fi
    printf '%s\n' "${STUB_REFS:-}" ;;
  "release create "*) ;;
  *) echo "unexpected gh call: $*" >&2; exit 2 ;;
esac
EOF
chmod +x "$root/bin/gh"
export PATH="$root/bin:$PATH"

status=0
reset() {
  : > "$calls"
  unset STUB_REFS_FAIL STUB_RELEASES_FAIL
  STUB_RELEASES='[]'
  STUB_REFS=$'refs/tags/v0.1.0\nrefs/tags/v0.1.1'
  export STUB_RELEASES STUB_REFS
}

# repo <tag>... : a fresh repository whose one commit carries every tag given.
repo() {
  local dir
  dir=$(mktemp -d "$root/XXXXXX")
  git -C "$dir" init -q
  git -C "$dir" -c user.name=t -c user.email=t@example.com commit -q --allow-empty -m init
  for t in "$@"; do git -C "$dir" tag "$t"; done
  echo "$dir"
}

# release <tag>... : run the script in a repository whose HEAD carries the tags; record its status.
release() {
  local dir
  dir=$(repo "$@")
  status=0
  (cd "$dir" && GITHUB_REPOSITORY=o/r "$script") > "$root/out" 2>&1 || status=$?
}

called() { grep -qF -- "$1" "$calls"; }

# check <name> <command...> : passes when the command succeeds.
check() {
  local name=$1
  shift
  if "$@"; then
    echo "ok   $name"
  else
    echo "FAIL $name"
    sed 's/^/     | /' "$root/out"
    failures=$((failures + 1))
  fi
}

# created <tag> : exactly one create call, for that tag.
created() { [ "$(grep -c "^release create $1 " "$calls" || true)" = 1 ]; }
no_create() { ! called 'release create'; }
succeeded() { [ "$status" = 0 ]; }
failed() { [ "$status" != 0 ]; }

reset
release v0.1.1
check "the tag on HEAD is created" created v0.1.1
check "the create call verifies the tag" called '--verify-tag'
check "the create call marks it latest" called '--latest=true'
check "a clean run exits 0" succeeded
check "the notes carry the go get command" called 'go get github.com/o/r/montecarlo@v0.1.1'
check "the notes link the package reference" called 'https://pkg.go.dev/github.com/o/r@v0.1.1/montecarlo'

reset
STUB_REFS=$'refs/tags/v0.1.1\nrefs/tags/v0.1.10'
release v0.1.1
check "an older tag is not marked latest" called '--latest=false'

reset
STUB_REFS=$'refs/tags/v0.1.9\nrefs/tags/v0.1.10'
release v0.1.10
check "versions compare as numbers, not text" called '--latest=true'

reset
STUB_REFS=$'refs/tags/v0.1.0\nrefs/tags/v0.1.1\nrefs/tags/v0.2.0-rc1'
release v0.1.1
check "a non-plain tag does not take latest" called '--latest=true'

reset
STUB_RELEASES='[{"tag_name":"v0.1.0","draft":true,"author":{"login":"someone"}},
{"tag_name":"v0.1.1","draft":false,"author":{"login":"someone"}}]'
release v0.1.1
check "a published release is left alone" succeeded
check "a published release is not recreated" no_create

reset
STUB_RELEASES='[{"tag_name":"v0.1.1","draft":true,"author":{"login":"someone"}}]'
release v0.1.1
check "a draft on the tag fails the run" failed
check "a draft on the tag is not duplicated" no_create
check "a draft on the tag is reported by its author" grep -q 'has a draft release by someone' "$root/out"

reset
STUB_REFS_FAIL=1
export STUB_REFS_FAIL
release v0.1.1
check "a failing tag listing fails the run" failed
check "a failing tag listing creates nothing" no_create

reset
STUB_RELEASES_FAIL=1
export STUB_RELEASES_FAIL
release v0.1.1
check "a failing release listing fails the run" failed
check "a failing release listing creates nothing" no_create

reset
release
check "an untagged HEAD succeeds" succeeded
check "an untagged HEAD creates nothing" no_create

reset
release foo
check "a non-version tag on HEAD is ignored" no_create

reset
STUB_REFS=$'refs/tags/v0.1.1\nrefs/tags/v0.1.2'
release v0.1.1 v0.1.2
check "the highest of two version tags on HEAD is released" created v0.1.2

exit $((failures > 0))
