#!/usr/bin/env python3
"""
populate_book_asins.py
Extracts the first Amazon affiliate link from each 書評-tagged article,
resolves the amzn.to short URL to an ASIN, and injects a `books:` block
into the YAML front-matter.

Usage:
    python3 tools/populate_book_asins.py [--dry-run] [--delay 1.0]

Options:
    --dry-run    Print what would be changed, but do not modify files.
    --delay      Seconds between Amazon requests (default: 1.0).
"""

import argparse
import re
import sys
import time
import urllib.request
import urllib.error
from pathlib import Path

CONTENT_DIR = Path(__file__).parent.parent / "content"

# Matches:  [任意のタイトル](https://amzn.to/XXXXXX)
AMZN_MD_RE = re.compile(r'\[([^\]]+)\]\((https://amzn\.to/[A-Za-z0-9]+)\)')
# Matches ASIN in the final Amazon URL
ASIN_RE = re.compile(r'/(?:dp|gp/product)/([A-Z0-9]{10})')


def resolve_asin(amzn_url: str) -> str | None:
    """Follow amzn.to redirect and return the ASIN, or None on failure."""
    try:
        req = urllib.request.Request(
            amzn_url,
            headers={"User-Agent": "Mozilla/5.0"},
        )
        with urllib.request.urlopen(req, timeout=10) as resp:
            final_url = resp.url
        m = ASIN_RE.search(final_url)
        return m.group(1) if m else None
    except urllib.error.URLError as e:
        print(f"  [WARN] request failed for {amzn_url}: {e}", file=sys.stderr)
        return None


def inject_books_frontmatter(path: Path, asin: str, title: str, dry_run: bool) -> bool:
    """Insert `books:` block before the closing `---` of the front-matter."""
    text = path.read_text(encoding="utf-8")
    # Find the second `---` that closes the front-matter
    # Front-matter: starts with `---\n` and ends with `\n---\n`
    fm_end = text.find("\n---\n", 3)
    if fm_end == -1:
        print(f"  [SKIP] no front-matter closing `---` in {path.name}", file=sys.stderr)
        return False
    if "books:" in text[:fm_end]:
        print(f"  [SKIP] books: already present in {path.name}")
        return False

    # Escape title for YAML (wrap in double quotes, escape inner quotes)
    safe_title = title.replace('"', '\\"')
    books_block = f'books:\n  - asin: "{asin}"\n    title: "{safe_title}"'

    new_text = text[:fm_end] + "\n" + books_block + text[fm_end:]
    if dry_run:
        print(f"  [DRY-RUN] would inject books: asin={asin} into {path.name}")
    else:
        path.write_text(new_text, encoding="utf-8")
        print(f"  [OK] {path.name} → asin={asin}")
    return True


def process_articles(dry_run: bool, delay: float) -> None:
    # Collect only 書評-tagged articles
    articles = sorted(CONTENT_DIR.rglob("*.md"))
    book_articles = []
    for p in articles:
        text = p.read_text(encoding="utf-8")
        fm_end = text.find("\n---\n", 3)
        if fm_end == -1:
            continue
        fm = text[:fm_end]
        if "  - 書評" in fm or "\n- 書評" in fm:
            book_articles.append((p, text))

    print(f"Found {len(book_articles)} 書評-tagged articles")

    for i, (path, text) in enumerate(book_articles, 1):
        print(f"[{i}/{len(book_articles)}] {path.name}")
        # Find the first amzn.to link in the article body
        m = AMZN_MD_RE.search(text)
        if not m:
            print(f"  [SKIP] no amzn.to link found")
            continue
        md_title, amzn_url = m.group(1), m.group(2)
        print(f"  link: {amzn_url}  title: {md_title[:40]}...")

        # Check if already has books:
        fm_end = text.find("\n---\n", 3)
        if "books:" in text[:fm_end]:
            print(f"  [SKIP] books: already present")
            continue

        asin = resolve_asin(amzn_url)
        if not asin:
            print(f"  [FAIL] could not resolve ASIN from {amzn_url}")
            continue

        inject_books_frontmatter(path, asin, md_title, dry_run)
        if i < len(book_articles):
            time.sleep(delay)


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--dry-run", action="store_true", help="Print changes without modifying files")
    parser.add_argument("--delay", type=float, default=1.0, help="Seconds between requests (default: 1.0)")
    args = parser.parse_args()
    process_articles(dry_run=args.dry_run, delay=args.delay)


if __name__ == "__main__":
    main()
