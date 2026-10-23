#!/usr/bin/env bash
# Check consistency of the picks pages (content/{en,ja}/picks.md).
#
# Usage:
#   scripts/check-picks.sh
#
# Exits non-zero if, in either language:
#   - a slug is in listing_slugs but in no picks_groups entry (or vice versa)
#   - a slug appears in more than one group
#   - a listed slug has no matching post in content/<lang>/posts
# or if the EN and JA pages list different slugs.

set -euo pipefail

status=0

# Print the slugs under a top-level front matter key.
# listing_slugs uses "  - slug"; picks_groups uses "      - slug" under "slugs:".
extract() {
    local file="$1" key="$2"
    awk -v key="$key" '
        BEGIN { count = 0 }
        /^---$/ { count++; if (count == 2) exit; next }
        count != 1 { next }
        /^[A-Za-z_]+:/ { in_key = ($0 ~ "^" key ":"); next }
        in_key && key == "listing_slugs" && /^  - / { sub(/^  - /, ""); print; next }
        in_key && key == "picks_groups" && /^      - / { sub(/^      - /, ""); print; next }
    ' "$file"
}

fail() {
    echo "::error::$1"
    status=1
}

for lang in en ja; do
    file="content/$lang/picks.md"
    listing=$(extract "$file" listing_slugs)
    grouped=$(extract "$file" picks_groups)

    dups=$(sort <<< "$grouped" | uniq -d)
    [[ -n "$dups" ]] && fail "$file: slugs in more than one group: $(tr '\n' ' ' <<< "$dups")"

    not_grouped=$(comm -23 <(sort -u <<< "$listing") <(sort -u <<< "$grouped"))
    [[ -n "$not_grouped" ]] && fail "$file: in listing_slugs but in no group: $(tr '\n' ' ' <<< "$not_grouped")"

    not_listed=$(comm -13 <(sort -u <<< "$listing") <(sort -u <<< "$grouped"))
    [[ -n "$not_listed" ]] && fail "$file: in picks_groups but not in listing_slugs: $(tr '\n' ' ' <<< "$not_listed")"

    while IFS= read -r slug; do
        [[ -z "$slug" ]] && continue
        if ! grep -lqE "^slug:[[:space:]]*[\"']?${slug}[\"']?[[:space:]]*$" "content/$lang/posts/"*.md; then
            fail "$file: no post with slug '$slug' in content/$lang/posts"
        fi
    done <<< "$listing"
done

en_list=$(extract content/en/picks.md listing_slugs | sort -u)
ja_list=$(extract content/ja/picks.md listing_slugs | sort -u)
if [[ "$en_list" != "$ja_list" ]]; then
    fail "EN and JA picks differ: $(diff <(echo "$en_list") <(echo "$ja_list") | grep '^[<>]' | tr '\n' ' ')"
fi

if [[ $status -eq 0 ]]; then
    echo "picks: OK ($(wc -l <<< "$en_list" | tr -d ' ') slugs)"
fi
exit $status
