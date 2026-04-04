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

bmf-tech.com has gone through four infrastructure generations.

1. **WordPress** — Abandoned after maintenance overhead and security patches became unsustainable.
2. **[Rubel](https://github.com/bmf-san/Rubel)** — A headless CMS built with Laravel. Ran a React frontend with MySQL on a ConoHa VPS.
3. **[gobel](https://github.com/bmf-san/gobel-api)** — A full rewrite of Rubel in Go. Ran for several years and accumulated over 700 articles. The backend used nginx, MySQL, and Redis via Docker Compose; the frontend used Vue.js. A monitoring stack with Prometheus, Grafana, Loki, and Pyroscope was also in place.
4. **[gohan](https://github.com/bmf-san/gohan)** — A Go-based SSG. Builds from Markdown files in `content/en/` and `content/ja/`, with GitHub Actions running the build and Cloudflare Pages serving the output. Serverless with zero running cost.

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

Defining 301 redirects from old URLs to new ones in `_redirects` kept the SEO impact low. Index coverage remained stable in Google Search Console after the DNS migration.

### What Changed

- **Writing environment** — Switched from a custom web admin panel to local editing in VS Code. Articles are now version-controlled in Git.
- **Quality control** — textlint runs in CI via GitHub Actions, catching inconsistent phrasing and style rule violations on every push.
- **Metadata management** — Front Matter now explicitly sets `title`, `description`, `categories`, `tags`, and OGP fields for every article. gobel saw none of this.
- **Self-hosted images** — Images depended on external CDNs such as Qiita Image Store. They are now managed as static assets under `assets/images/posts/{slug}/`, eliminating the risk of losing images if an external service shuts down.
- **Slug-based URLs** — Article URLs changed from percent-encoded Japanese titles to English slug format (e.g., `/posts/go-http-server/`), improving readability and shareability.
- **i18n** — The site now has a two-language structure with `content/ja/` and `content/en/`, with articles linked via `translation_key`.
- **Hosting** — Moved from a self-managed ConoHa VPS to Cloudflare Pages. Static content comes from a CDN with a generous free tier, so there are no ongoing server costs.
- **Custom pages** — Added a bookshelf page for tracking books read and an Amazon Associate link generator page.

## Summary

Since the migration, the site has been running stably and comfortably. Writing and pushing an article is all it takes to deploy — no server maintenance required. Looking back, I wish I had made the switch to an SSG sooner.

I made heavy use of AI throughout the development. The project had been sitting idle for a long time despite having a rough design in place, but with AI help I was able to move through both the blog system development and the migration at a pace I would not have managed otherwise. The small scale of the system helped, but without AI the effort would have been far greater. I almost feel glad I waited until AI was available before starting.

Both the source code and the articles are fully public.

- **bmf-tech source**: [bmf-san/bmf-tech](https://github.com/bmf-san/bmf-tech)
- **gohan**: [bmf-san/gohan](https://github.com/bmf-san/gohan)
