#!/usr/bin/env bash
# Creates the GitHub release for <tag>, unless one is published. Marks it latest only when it is
# the highest plain version tag. Fails on a draft rather than create a second release beside it.
#
#   GITHUB_REPOSITORY=owner/repo .github/scripts/github-release.sh v0.1.4
set -euo pipefail

tag=$1
repo=$GITHUB_REPOSITORY
module=github.com/$repo

# Listed rather than fetched by tag: the by-tag lookup never returns a draft.
releases=$(gh api "repos/$repo/releases" --paginate \
  --jq ".[] | select(.tag_name == \"$tag\") | \"\(.draft) \(.author.login)\"")

# With no files to upload, gh publishes in one call, so this job never leaves a draft.
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
notes=$(printf '```bash\ngo get %s@%s\n```\n\n%s\n' "$module" "$tag" "$notes")

gh release create "$tag" --repo "$repo" --verify-tag \
  --title "$tag" --notes "$notes" --latest="$latest"
