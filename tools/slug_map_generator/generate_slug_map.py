#!/usr/bin/env python3
"""
generate_slug_map.py - Parses MySQL dump and generates slug_map.csv.
Run from the project root: python3 tools/slug_map_generator/generate_slug_map.py
"""
import re
import csv

SQL_PATH = "bmf-tech_2026-03-01.sql"
OUT_PATH = "tools/slug_map.csv"


def parse_mysql_string(s, pos):
    """
    Parse a single-quoted MySQL string starting at pos (after the opening quote).
    Returns (unescaped_string, next_pos_after_closing_quote).
    """
    result = []
    i = pos
    while i < len(s):
        c = s[i]
        if c == "\\" and i + 1 < len(s):
            nc = s[i + 1]
            table = {"'": "'", "\\": "\\", "n": "\n", "r": "\r", '"': '"'}
            result.append(table.get(nc, nc))
            i += 2
        elif c == "'":
            if i + 1 < len(s) and s[i + 1] == "'":
                result.append("'")
                i += 2
            else:
                return "".join(result), i + 1
        else:
            result.append(c)
            i += 1
    return "".join(result), i


def slugify(post_id, title):
    """Generate a best-effort ASCII slug from a (Japanese) title."""
    words = []
    cur = []
    for c in title:
        if c == "+":
            if cur:
                words.append("".join(cur))
                cur = []
            words.append("plus")
        elif c.isascii() and (c.isalpha() or c.isdigit()):
            cur.append(c.lower())
        else:
            if cur:
                words.append("".join(cur))
                cur = []
    if cur:
        words.append("".join(cur))

    # Filter single-char tokens and deduplicate adjacent
    words = [w for w in words if len(w) > 1]
    deduped = []
    for w in words:
        if not deduped or deduped[-1] != w:
            deduped.append(w)

    words = deduped[:6]
    if words:
        return "-".join(words)
    return f"post-{post_id}"


def main():
    with open(SQL_PATH, "r", encoding="utf-8", errors="replace") as f:
        content = f.read()

    lines = content.split("\n")
    insert_lines = [l for l in lines if l.startswith("INSERT INTO `posts` VALUES")]
    print(f"Found {len(insert_lines)} INSERT lines for posts")

    records = []
    for line in insert_lines:
        i = 0
        while i < len(line):
            m = re.search(r"\((\d+),\d+,(\d+),", line[i:])
            if not m:
                break

            pos = i + m.start() + len(m.group(0))
            post_id = m.group(1)
            cat_id = m.group(2)
            i = i + m.end()

            # Parse title
            if pos < len(line) and line[pos] == "'":
                title, end_pos = parse_mysql_string(line, pos + 1)
                title_clean = title.replace("\n", " ").replace("\r", "").strip()
            elif line[pos : pos + 4] == "NULL":
                title_clean = ""
                end_pos = pos + 4
            else:
                title_clean = ""
                end_pos = pos

            # Extract created_at by searching for the last two datetime stamps
            rest_match = re.search(
                r"'(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})'"
                r",'(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})'",
                line[end_pos:],
            )
            created_at = rest_match.group(1) if rest_match else ""

            records.append(
                {
                    "id": post_id,
                    "cat_id": cat_id,
                    "title": title_clean,
                    "created_at": created_at,
                }
            )

    print(f"Parsed {len(records)} post records")

    with open(OUT_PATH, "w", newline="", encoding="utf-8") as f:
        w = csv.writer(f)
        w.writerow(["id", "category_id", "title_ja", "slug", "created_at", "reviewed"])
        for r in records:
            slug = slugify(r["id"], r["title"])
            w.writerow(
                [r["id"], r["cat_id"], r["title"], slug, r["created_at"], "false"]
            )

    print(f"Written to {OUT_PATH}")


if __name__ == "__main__":
    main()
