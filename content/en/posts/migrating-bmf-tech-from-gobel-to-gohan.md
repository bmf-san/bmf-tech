---
title: "Migrating bmf-tech.com from gobel to gohan — A Complete Static Site Migration"
description: 'How I migrated bmf-tech.com from a self-built headless CMS (gobel) backed by MySQL and a Vue.js frontend to a fully static site powered by gohan, hosted on Cloudflare Pages — covering 700+ articles, English slug generation, image migration, and a custom Go preflight checker.'
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

bmf-tech.com has been through three generations of infrastructure:

1. **WordPress** — the original setup, abandoned due to maintenance overhead and security patching.
2. **Rubel → gobel** — a self-built headless CMS using Laravel (then rewritten in Go as [gobel](https://github.com/bmf-san/gobel-api)) with a Vue.js frontend, backed by MySQL on a ConoHa VPS. This served the site for several years and accumulated 700+ articles.
3. **gohan on Cloudflare Pages** — the current architecture described in this post.

The goal was to remove server costs and runtime security exposure, reducing operational burden to zero — Markdown files in a Git repository, built and deployed automatically on push.

## Migration Phases

The migration proceeded in 12 phases:

| Phase | Description |
|---|---|
| 0 | Design document and migration plan |
| 1 | Repository and gohan project setup |
| 2 | Theme and template development |
| 3 | English slug map generation (700+ articles) |
| 4 | Content migration script (SQL dump → Markdown) |
| 5 | Japanese article migration (`content/ja/`) |
| 6 | Image asset migration |
| 7 | Redirect map creation and validation |
| 8 | English article creation (high-priority first) |
| 9 | CI/CD pipeline |
| 9.5 | Pre-migration preflight check |
| 10 | Cloudflare Pages production deploy |
| 11 | DNS transfer (ConoHa → Cloudflare) |
| 12 | ConoHa VPS decommission |

### Phase 3: English Slug Generation

The original gobel database had no `slug` column. Gobel constructed URLs by URL-encoding the Japanese article title directly — for example `/posts/Go%E3%81%A7HTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%82%92%E6%9B%B8%E3%81%8F`. After migration the target was clean English slugs like `/posts/go-http-server/`.

With 700+ articles to slug, I used Claude to batch-generate English slug candidates from `(id, title)` pairs, reviewed the output in a `slug_map.csv`, and de-duplicated before committing.

### Phase 4: Migration Script

A Go tool took two inputs — a MySQL dump and `slug_map.csv` — and produced:

- `content/ja/posts/*.md` — one Markdown file per article, with front matter derived from `posts`, `categories`, and `tag_post` tables.
- `_redirects` — a Cloudflare Pages redirect file mapping old percent-encoded Japanese URLs to new English slug paths with 301 status codes.

The `_redirects` file required care: Cloudflare Pages receives requests with percent-encoded paths, so the generator had to apply `url.PathEscape()` to old URLs to ensure redirect rules matched correctly.

### Phase 6: Image Migration

The original site hosted all images on an external CDN (primarily Qiita Image Store). A download tool fetched each image referenced in migrated articles, stored them under `assets/images/posts/{slug}/`, and rewrote the Markdown `![alt](url)` references to local paths.

### Phase 9.5: Preflight Checker

Before cutting DNS, a Go preflight tool made HTTP requests to the production Cloudflare Pages preview URL for every article URL (using the new English slugs) and verified HTTP 200 responses. It also verified that all redirect rules in `_redirects` returned 301. This gave confidence that no article would land on a 404 after DNS cut-over.

### Phase 11: DNS Transfer

I transferred the nameservers from ConoHa to Cloudflare. Cloudflare Pages served the gohan-built static files directly from its CDN, making `www.bmf-tech.com` fully static with zero server-side code running.

### Phase 12: ConoHa VPS Decommission

After confirming stable traffic flows and no broken redirects via Google Search Console, I shut down the ConoHa VPS running the gobel API and MySQL database.

## Result

gohan generates the current site from Markdown files in `/content/en/` and `/content/ja/`. GitHub Actions builds and deploys to Cloudflare Pages on every push. Build time for 584+ EN and 584+ JA articles is under 60 seconds including OGP image generation.

SEO continuity held: Google Search Console showed index coverage after DNS transfer thanks to the `_redirects` permanent redirect chain.

- **bmf-tech source**: [bmf-san/bmf-tech](https://github.com/bmf-san/bmf-tech)
- **gohan**: [bmf-san/gohan](https://github.com/bmf-san/gohan)
