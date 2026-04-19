#!/usr/bin/env bash
# Check translation_key parity between content/en/posts and content/ja/posts.
#
# Usage:
#   scripts/check-translation-parity.sh
#
# Exits non-zero if any translation_key exists in exactly one language.
# Posts marked `draft: true` are ignored on the expectation that
# unpublished translations are work-in-progress.

set -euo pipefail

EN_DIR="content/en/posts"
JA_DIR="content/ja/posts"

extract_keys() {
    local dir="$1"
    local file
    for file in "$dir"/*.md; do
        [[ -e "$file" ]] || continue
        # Skip files whose front matter has `draft: true`.
        if awk 'BEGIN{count=0} /^---$/{count++; if(count==2) exit; next} count==1 && /^draft:[[:space:]]*true[[:space:]]*$/{found=1; exit} END{exit !found}' "$file"; then
            continue
        fi
        local key
        key=$(awk 'BEGIN{count=0} /^---$/{count++; if (count==2) exit; next} count==1 && /^translation_key:/{sub(/^translation_key:[[:space:]]*/, ""); gsub(/[\x22\x27]/, ""); print; exit}' "$file")
        if [[ -n "$key" ]]; then
            echo "$key"
        fi
    done | sort -u
}

en_keys=$(extract_keys "$EN_DIR")
ja_keys=$(extract_keys "$JA_DIR")

# Keys in EN only
en_only=$(comm -23 <(echo "$en_keys") <(echo "$ja_keys"))
# Keys in JA only
ja_only=$(comm -13 <(echo "$en_keys") <(echo "$ja_keys"))

status=0

if [[ -n "$en_only" ]]; then
    echo "::warning::translation_key found in EN only (missing JA translation):"
    while IFS= read -r k; do
        [[ -z "$k" ]] && continue
        echo "  - $k"
    done <<< "$en_only"
    status=1
fi

if [[ -n "$ja_only" ]]; then
    echo "::warning::translation_key found in JA only (missing EN translation):"
    while IFS= read -r k; do
        [[ -z "$k" ]] && continue
        echo "  - $k"
    done <<< "$ja_only"
    status=1
fi

if [[ $status -eq 0 ]]; then
    echo "All translation_key values are paired across EN and JA."
fi

exit $status
