#!/usr/bin/env bash
# Creates the GitHub release for the version tag on HEAD, read from the local clone, unless one
# is published. Marks it latest only when it is the highest plain version tag. Fails on a draft
# rather than create a second release beside it.
#
#   GITHUB_REPOSITORY=owner/repo .github/scripts/github-release.sh
set -euo pipefail

# A commit older than an already-tagged one is never tagged, so it gets no release.
tag=$("$(dirname "$0")/release-tag.sh")
if [ -z "$tag" ]; then
  echo "HEAD carries no version tag; nothing to release."
  exit 0
fi

repo=$GITHUB_REPOSITORY
module=github.com/$repo

# Listed rather than fetched by tag: the by-tag lookup never returns a draft.
releases=$(gh api "repos/$repo/releases" --paginate \
  --jq ".[] | select(.tag_name == \"$tag\") | \"\(.draft) \(.author.login)\"")

# gh publishes in one call when there are no files, so a draft here is not this job's: fail
# rather than delete it.
while read -r draft author; do
  [ -n "$draft" ] || continue
  if [ "$draft" = false ]; then
    echo "$tag already has a release."
    exit 0
  fi
  echo "::error::$tag has a draft release by $author; delete or publish it, then re-run." >&2
  exit 1
done <<<"$releases"

refs=$(gh api "repos/$repo/git/matching-refs/tags/v" --jq '.[].ref')
highest=$(printf '%s\n' "$refs" | sed 's#^refs/tags/##' |
  { grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' || true; } | sort -V | tail -1)
latest=false
if [ "$highest" = "$tag" ]; then latest=true; fi

notes="Reference: [pkg.go.dev](https://pkg.go.dev/$module@$tag/montecarlo)"
# shellcheck disable=SC2016 # the backticks are a Markdown code fence
notes=$(printf '```bash\ngo get %s/montecarlo@%s\n```\n\n%s\n' "$module" "$tag" "$notes")

gh release create "$tag" --repo "$repo" --verify-tag \
  --title "$tag" --notes "$notes" --latest="$latest"
