#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""Generate _redirects for Cloudflare Pages from slug_map.csv + taxonomies."""
import csv
import urllib.parse
import re
import os

def load_names(path):
    """Parse simple YAML list: - name: 'xxx'"""
    names = []
    with open(path, encoding="utf-8") as f:
        for line in f:
            m = re.match(r"\s*-\s*name:\s*['\"]?(.+?)['\"]?\s*$", line)
            if m:
                names.append(m.group(1))
    return names

base = "/Users/bmf/localdev/bmf-tech"

lines = []
lines.append("# Cloudflare Pages _redirects")
lines.append("# Format: /old-path  /new-path  301")
lines.append("")

# ── 1. Articles: /posts/{url-encoded-title_ja} → /ja/posts/{slug}/ ──────────
lines.append("# Articles: /posts/{japanese-title} -> /ja/posts/{slug}/")
with open(f"{base}/tools/slug_map.csv", newline="", encoding="utf-8") as f:
    reader = csv.DictReader(f)
    for row in reader:
        title = row["title_ja"].strip()
        slug  = row["slug"].strip()
        if not slug:
            continue
        encoded = urllib.parse.quote(title, safe="")
        lines.append(f"/posts/{encoded}  /ja/posts/{slug}/  301")

lines.append("")

# ── 1b. Articles by numeric ID: /posts/{id} → /ja/posts/{slug}/ ─────────────
lines.append("# Articles by numeric ID (fallback for older links)")
with open(f"{base}/tools/slug_map.csv", newline="", encoding="utf-8") as f:
    reader = csv.DictReader(f)
    for row in reader:
        post_id = row["id"].strip()
        slug    = row["slug"].strip()
        if not slug:
            continue
        lines.append(f"/posts/{post_id}  /ja/posts/{slug}/  301")

lines.append("")

# ── 2. Categories: /posts/categories/{enc} → /categories/{name}/ ────────────
lines.append("# Categories: /posts/categories/{name} -> /categories/{name}/")
cats = load_names(f"{base}/taxonomies/categories.yaml")
for name in cats:
    encoded = urllib.parse.quote(name, safe="")
    lines.append(f"/posts/categories/{encoded}  /categories/{name}/  301")

lines.append("")

# ── 3. Tags: /posts/tags/{enc} → /tags/{name}/ ──────────────────────────────
lines.append("# Tags: /posts/tags/{name} -> /tags/{name}/")
tags = load_names(f"{base}/taxonomies/tags.yaml")
for name in tags:
    encoded = urllib.parse.quote(name, safe="")
    lines.append(f"/posts/tags/{encoded}  /tags/{name}/  301")

lines.append("")

# ── 4. Static pages ──────────────────────────────────────────────────────────
lines.append("# Static pages")
lines.append("/privacy_policy  /privacy-policy/  301")
lines.append("/sitemap  /sitemap.xml  301")
lines.append("/feed  /atom.xml  301")
lines.append("/posts/search  /  301")
lines.append("/profile  /about/  301")
lines.append("/categories  /categories/  301")
lines.append("/tags  /tags/  301")

content = "\n".join(lines) + "\n"
out = f"{base}/_redirects"
with open(out, "w", encoding="utf-8") as f:
    f.write(content)

article_count = sum(1 for l in lines if l.startswith("/posts/") and "/ja/posts/" in l)
cat_count = sum(1 for l in lines if "/posts/categories/" in l)
tag_count = sum(1 for l in lines if "/posts/tags/" in l)
print(f"Articles  : {article_count}")
print(f"Categories: {cat_count}")
print(f"Tags      : {tag_count}")
print(f"Static    : 4")
print(f"Total rules: {article_count + cat_count + tag_count + 4}")
print(f"Written to: {out}")
