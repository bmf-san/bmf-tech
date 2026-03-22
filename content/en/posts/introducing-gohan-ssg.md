---
title: "Introducing gohan — A Go Static Site Generator with Incremental Builds"
description: 'A deep dive into gohan, a Go-based static site generator powering bmf-tech.com. Features SHA-256 manifest-driven incremental builds, i18n, Mermaid diagrams, OGP image generation, and a compiled plugin system (Amazon book cards, bookshelf page).'
slug: introducing-gohan-ssg
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - Golang
  - SSG
  - Architecture
translation_key: introducing-gohan-ssg
---

# Introducing gohan — A Go Static Site Generator with Incremental Builds

## Why I Built It

This site (bmf-tech.com) runs on gohan. The motivation was simple: I wanted a static site generator I could fully understand, one that regenerates only the pages that change.

Most generators rebuild everything unconditionally or depend on `git diff` output. Git diff becomes unreliable after branch switches or fresh clones. gohan creates and persists a build manifest using SHA-256 content hashing. Incremental builds stay accurate without relying on Git history.

## Incremental Build Engine

The core of the incremental build lives in `internal/diff/git.go`. The `Detect()` method compares the current working tree against a persisted `BuildManifest`.

```go
func (g *GitDiffEngine) Detect(manifest *model.BuildManifest) (*model.ChangeSet, error) {
    current, err := hashAllFiles(g.rootDir)
    if err != nil {
        return nil, err
    }

    if manifest == nil {
        cs := &model.ChangeSet{}
        for path := range current {
            cs.AddedFiles = append(cs.AddedFiles, path)
        }
        return cs, nil
    }

    cs := &model.ChangeSet{}
    for path, hash := range current {
        if prev, ok := manifest.FileHashes[path]; !ok {
            cs.AddedFiles = append(cs.AddedFiles, path)
        } else if prev != hash {
            cs.ModifiedFiles = append(cs.ModifiedFiles, path)
        }
    }
    for path := range manifest.FileHashes {
        if _, ok := current[path]; !ok {
            cs.DeletedFiles = append(cs.DeletedFiles, path)
        }
    }
    return cs, nil
}
```

`hashAllFiles()` walks the content directory and computes a SHA-256 hex digest for every file. On the first build (or when no manifest exists), all files count as `Added`. Later builds detect three change types — `Added`, `Modified`, and `Deleted` — and regenerate only the affected HTML pages.

## Features

Beyond incremental builds, gohan ships with many capabilities out of the box.

- **i18n** — Mirror directory layout such as `content/en/` and `content/ja/`. Locale-switch links generate automatically.
- **Syntax highlighting** — Server-side rendering via Chroma. No client-side JavaScript needed.
- **Mermaid diagrams** — Output as SVG at build time or as `<pre class="mermaid">` for client-side rendering.
- **OGP image generation** — Open Graph images generate per article at build time.
- **Pagination** — Configurable articles-per-page count.
- **Related articles** — Similar article links based on shared tags.
- **GitHub source links** — Edit links pointing to Markdown sources, added automatically.
- **Live-reload dev server** — `gohan serve` watches content and rebuilds on every save.

## Plugin System

Plugins compile into the gohan binary and activate per-project via `config.yaml`. No recompilation needed by users. The plugin interfaces live in `internal/plugin/plugin.go`.

```go
type Plugin interface {
    Name() string
    Enabled(cfg map[string]interface{}) bool
    TemplateData(article *model.ProcessedArticle, cfg map[string]interface{}) (map[string]interface{}, error)
}

type SitePlugin interface {
    Name() string
    Enabled(cfg map[string]interface{}) bool
    VirtualPages(site *model.Site, cfg map[string]interface{}) ([]*model.VirtualPage, error)
}
```

`Plugin` (article-level) exposes extra data for a single article through the template as `.PluginData.<name>`. `SitePlugin` (site-level) runs after all articles have processed and can produce **virtual pages** — pages with no Markdown source.

The built-in registry ships two plugins.

```go
func DefaultRegistry() *Registry {
    return &Registry{
        plugins: []Plugin{
            amazonbooks.New(),
        },
        sitePlugins: []SitePlugin{
            bookshelf.New(),
        },
    }
}
```

`amazon_books` generates Amazon affiliate book card data (image, URL, title) from ASIN values in article frontmatter. `bookshelf` aggregates book frontmatter across the whole site and produces a virtual `/bookshelf` page.

Example `config.yaml` setup:

```yaml
plugins:
  amazon_books:
    enabled: true
    tag: "your-associate-tag-22"
  bookshelf:
    enabled: true
```

## Install and Usage

```bash
# Homebrew (macOS/Linux)
brew install bmf-san/tap/gohan

# Go install
go install github.com/bmf-san/gohan/cmd/gohan@latest

# Build
gohan build

# Dev server with live reload
gohan serve
```

## Closing

gohan is the engine behind this site. SHA-256 manifest-driven incremental builds keep iteration fast. The compiled plugin system keeps the binary self-contained. From i18n to OGP to Mermaid, everything runs at build time with no client-side JavaScript required.

- **GitHub**: [bmf-san/gohan](https://github.com/bmf-san/gohan)
