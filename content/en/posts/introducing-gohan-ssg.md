---
title: Introducing gohan — A Go Static Site Generator with Incremental Builds
description: 'An introduction to gohan, a Go-based static site generator that builds only changed files using SHA-256 manifests, supports i18n, Mermaid diagrams, OGP, syntax highlighting, and a compiled plugin system with Amazon book cards and a bookshelf page.'
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

This very website (bmf-tech.com) is powered by gohan. The motivation was straightforward: I wanted a static site generator that I could fully understand, modify, and that would build only the pages that actually changed — not the entire site on every run.

Most generators either rebuild everything unconditionally or rely on Git diff output, which breaks when you switch branches or do a fresh clone. gohan uses SHA-256 content hashing to produce a build manifest that is persisted between runs, making incremental builds reliable regardless of Git history.

## Incremental Build Engine

The core of the incremental build is in `internal/diff/git.go`. The `Detect()` method compares the current working tree against a persisted `BuildManifest`:

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

`hashAllFiles()` walks the content directory and computes a SHA-256 hex digest for every file. On the first build (or when no manifest exists), every file is treated as `Added` — a clean full build. On subsequent builds, three change types are detected: `Added`, `Modified`, and `Deleted`. Only the affected HTML pages are regenerated.

## Feature Set

Beyond incremental builds, gohan ships several features out of the box:

- **i18n** — mirror directory structure for multiple locales (e.g. `content/en/` and `content/ja/`); locale switcher links generated automatically.
- **Syntax highlighting** — code fences rendered server-side via Chroma, no runtime JavaScript.
- **Mermaid diagrams** — fenced Mermaid blocks rendered as `<pre class="mermaid">` for client-side rendering, or as inline SVG at build time.
- **OGP image generation** — Open Graph images generated at build time per article.
- **Pagination** — configurable articles per page.
- **Related articles** — tag-based similarity linking.
- **GitHub source link** — auto-appended edit links pointing to the Markdown source.
- **Live-reload dev server** — `gohan serve` watches content and rebuilds on save.

## Plugin System

Plugins are compiled into the gohan binary and enabled per-project via `config.yaml`. No recompilation is needed by callers. The plugin interface is defined in `internal/plugin/plugin.go`:

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

`Plugin` (per-article) enriches a single processed article with additional data exposed to the theme template via `.PluginData.<name>`. `SitePlugin` (site-level) runs after all articles are processed and can generate **VirtualPages** — pages with no corresponding Markdown source file.

The built-in registry ships two plugins out of the box:

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

`amazon_books` generates Amazon affiliate book card data (image, link, title) from ASIN values declared in article front-matter. `bookshelf` aggregates all book front-matter across the site and produces a virtual `/bookshelf` page.

Enabling them in `config.yaml`:

```yaml
plugins:
  amazon_books:
    enabled: true
    tag: "your-associate-tag-22"
  bookshelf:
    enabled: true
```

## Installation and Usage

```bash
# Homebrew (macOS/Linux)
brew install bmf-san/tap/gohan

# Go install
go install github.com/bmf-san/gohan/cmd/gohan@latest

# Build site
gohan build

# Development server with live reload
gohan serve
```

## Summary

gohan is the generator that powers this site. Incremental builds via SHA-256 manifests keep iteration fast; the compiled plugin system keeps the binary self-contained; and every feature — i18n, OGP, Mermaid — works without client-side JavaScript at build time.

- **GitHub**: [bmf-san/gohan](https://github.com/bmf-san/gohan)
