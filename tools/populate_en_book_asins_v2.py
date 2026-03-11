#!/usr/bin/env python3
"""
populate_en_book_asins_v2.py
For EN articles tagged "Books":
  1. Look up ASIN from the corresponding JA article (matched via translation_key)
  2. Inject `books:` frontmatter block
  3. Rename tag "Books" -> "書評"

No network requests needed — reuses ASINs already injected into JA articles.

Usage:
    python3 tools/populate_en_book_asins_v2.py [--dry-run]
"""

import argparse
import re
import sys
from pathlib import Path

EN_DIR = Path(__file__).parent.parent / "content" / "en" / "posts"
JA_DIR = Path(__file__).parent.parent / "content" / "ja" / "posts"

TRANSLATION_KEY_RE = re.compile(r'^translation_key:\s*(\S+)', re.MULTILINE)
ASIN_RE = re.compile(r'asin:\s*"([^"]+)"')
TITLE_RE = re.compile(r'title:\s*"([^"]+)"')


def parse_frontmatter(text: str) -> dict:
    fm_end = text.find("\n---\n", 3)
    if fm_end == -1:
        return {}
    fm = text[:fm_end]
    result = {}
    m = TRANSLATION_KEY_RE.search(fm)
    if m:
        result["translation_key"] = m.group(1)
    # books block
    if "books:" in fm:
        asins = ASIN_RE.findall(fm)
        titles = TITLE_RE.findall(fm)
        if asins:
            result["asin"] = asins[0]
            result["title"] = titles[0] if titles else asins[0]
    return result


def build_ja_asin_map() -> dict[str, dict]:
    """Return {translation_key: {asin, title}} from JA articles with books:"""
    asin_map = {}
    for p in JA_DIR.glob("*.md"):
        text = p.read_text(encoding="utf-8")
        info = parse_frontmatter(text)
        if "translation_key" in info and "asin" in info:
            asin_map[info["translation_key"]] = {
                "asin": info["asin"],
                "title": info.get("title", info["asin"]),
            }
    return asin_map


def inject_books_frontmatter(path: Path, asin: str, title: str, dry_run: bool) -> bool:
    text = path.read_text(encoding="utf-8")
    fm_end = text.find("\n---\n", 3)
    if fm_end == -1:
        print(f"  [SKIP] no front-matter closing in {path.name}", file=sys.stderr)
        return False
    if "books:" in text[:fm_end]:
        print(f"  [SKIP] books: already present in {path.name}")
        return False
    safe_title = title.replace('"', '\\"')
    books_block = f'books:\n  - asin: "{asin}"\n    title: "{safe_title}"'
    new_text = text[:fm_end] + "\n" + books_block + text[fm_end:]
    if dry_run:
        print(f"  [DRY-RUN] would inject asin={asin}")
    else:
        path.write_text(new_text, encoding="utf-8")
        print(f"  [OK] injected asin={asin}")
    return True


def rename_books_tag(path: Path, dry_run: bool) -> bool:
    text = path.read_text(encoding="utf-8")
    new_text = re.sub(r'^( *- )Books\s*$', r'\1書評', text, flags=re.MULTILINE)
    if new_text == text:
        return False
    if dry_run:
        print(f"  [DRY-RUN] would rename Books -> 書評")
    else:
        path.write_text(new_text, encoding="utf-8")
        print(f"  [OK] renamed Books -> 書評")
    return True


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--dry-run", action="store_true")
    args = parser.parse_args()

    print("Building JA ASIN map...")
    ja_map = build_ja_asin_map()
    print(f"  {len(ja_map)} JA articles with ASIN")

    articles = sorted(EN_DIR.glob("*.md"))
    book_articles = []
    for p in articles:
        text = p.read_text(encoding="utf-8")
        fm_end = text.find("\n---\n", 3)
        if fm_end == -1:
            continue
        fm = text[:fm_end]
        if re.search(r'^ *- Books\s*$', fm, re.MULTILINE):
            book_articles.append(p)

    print(f"Found {len(book_articles)} EN articles tagged 'Books'")

    no_match = []
    for i, path in enumerate(book_articles, 1):
        text = path.read_text(encoding="utf-8")
        info = parse_frontmatter(text)
        tk = info.get("translation_key", "")
        ja_info = ja_map.get(tk)

        print(f"[{i}/{len(book_articles)}] {path.name}")
        if ja_info:
            inject_books_frontmatter(path, ja_info["asin"], ja_info["title"], args.dry_run)
        else:
            print(f"  [SKIP] no JA ASIN match for translation_key={tk!r}")
            no_match.append(path.name)

        rename_books_tag(path, args.dry_run)

    if no_match:
        print(f"\nNo JA ASIN match ({len(no_match)}):")
        for n in no_match:
            print(f"  {n}")
    else:
        print("\nAll articles matched!")


if __name__ == "__main__":
    main()
