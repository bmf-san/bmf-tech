---
title: "Migrating bmf-tech.com from gobel to gohan — A Full Static Site Transition"
description: 'The complete process of migrating bmf-tech.com from a self-built headless CMS (gobel) with MySQL and Vue.js to a fully static site on gohan, hosted on Cloudflare Pages. Covers 700+ articles, English slug generation, image migration, redirects, and a Go-based preflight checker.'
slug: migrating-bmf-tech-from-gobel-to-gohan
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - gohan
  - Migration
  - SSG
  - Golang
translation_key: migrating-bmf-tech-from-gobel-to-gohan
---

# Migrating bmf-tech.com from gobel to gohan

## Background

bmf-tech.com has gone through three infrastructure generations.

1. **WordPress** — Abandoned after maintenance overhead and security patches became unsustainable.
2. **Rubel → gobel** — A headless CMS originally written in Laravel, later rewritten in Go as [gobel](https://github.com/bmf-san/gobel-api). I ran a Vue.js frontend with MySQL on a ConoHa VPS for several years, accumulating over 700 articles.
3. **gohan on Cloudflare Pages** — The current setup, which this article covers.

The goal was to bring server cost and operational load down to zero. I wanted a setup where pushing to GitHub triggers an automatic build and deploy, with the entire site built from Markdown files.

## Migration Phases

I planned the migration in 12 phases.

| Phase | Content |
|---|---|
| 0 | Design document and migration plan |
| 1 | Repository and gohan project setup |
| 2 | Theme and template development |
| 3 | English slug mapping (700+ entries) |
| 4 | Content migration script (SQL dump → Markdown) |
| 5 | Japanese article Markdown migration (`content/ja/`) |
| 6 | Image asset migration |
| 7 | Redirect map creation and validation |
| 8 | English article creation (high-priority articles first) |
| 9 | CI/CD pipeline setup |
| 9.5 | Pre-migration preflight check |
| 10 | Cloudflare Pages production deploy |
| 11 | DNS migration (ConoHa → Cloudflare) |
| 12 | ConoHa VPS shutdown |

### Phase 3: English Slug Generation

The original gobel database had no `slug` column. URLs used URL-encoded Japanese titles — for example `/posts/Go%E3%81%A7HTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%82%92%E6%9B%B8%E3%81%8F`. After the migration, I wanted clean English slugs like `/posts/go-http-server/`.

Writing slugs for 700+ articles by hand was impractical, so I fed batches of `(id, title)` pairs to Claude to generate English slug candidates at scale. I exported the output to `slug_map.csv` for manual review, checking for duplicates and confirming meaning.

### Phase 4: Migration Script

The Go migration tool took two inputs — a MySQL dump and `slug_map.csv` — and produced two outputs.

- `content/ja/posts/*.md` — One Markdown file per article, with frontmatter generated from the `posts`, `categories`, and `tag_post` tables.
- `_redirects` — A 301 redirect rule file in Cloudflare Pages format, mapping old URLs to new ones.

One issue came up during `_redirects` generation. Cloudflare Pages receives request paths in percent-encoded form, so the script had to apply `url.PathEscape()` to the old URLs. Writing raw Japanese strings in the rules caused redirects to stop working.

### Phase 6: Image Asset Migration

All images on the original site lived on external CDNs, primarily Qiita Image Store. A download tool scanned each article's Markdown, saved images under `assets/images/posts/{slug}/`, and rewrote `![alt](url)` references to local paths.

### Phase 9.5: Preflight Checker

Before the DNS switch, a Go preflight tool sent HTTP requests to every article URL against the Cloudflare Pages preview URL and verified HTTP 200 responses. It also confirmed every rule in `_redirects` returned a 301. This gave confidence that no articles would 404 after the DNS cutover.

### Phase 11: DNS Migration

I moved the nameservers from ConoHa to Cloudflare. Cloudflare Pages now serves gohan's static output directly from its CDN. `www.bmf-tech.com` runs zero server-side code.

### Phase 12: ConoHa VPS Shutdown

After confirming index coverage trends in Google Search Console, I shut down the ConoHa VPS that had been running the gobel API and MySQL. Monthly VPS cost dropped to zero.

## Post-Migration Setup

gohan now generates the site from Markdown files in `content/en/` and `content/ja/`. Pushing to GitHub triggers GitHub Actions to build and deploy to Cloudflare Pages. A full build of 584+ EN and 584+ JA articles, including OGP image generation, completes in under 60 seconds.

SEO continuity held through the permanent 301 redirect chain in `_redirects`. Index coverage stayed stable in Google Search Console after the DNS migration.

- **bmf-tech source**: [bmf-san/bmf-tech](https://github.com/bmf-san/bmf-tech)
- **gohan**: [bmf-san/gohan](https://github.com/bmf-san/gohan)
