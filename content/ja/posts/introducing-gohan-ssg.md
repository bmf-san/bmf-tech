---
title: "gohan — インクリメンタルビルド対応のGo製静的サイトジェネレータの紹介"
description: 'SHA-256マニフェストによる差分ビルド、i18n、Mermaid図、OGP、シンタックスハイライト、コンパイル済プラグインシステム（Amazon本カード、本棚ページ）を備えたGo製静的サイトジェネレータ『gohan』の紹介。'
slug: introducing-gohan-ssg
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - ツール
tags:
  - Golang
  - SSG
  - Architecture
translation_key: introducing-gohan-ssg
---

# gohan — インクリメンタルビルド対応のGo製静的サイトジェネレータの紹介

## 作った理由

このサイト（bmf-tech.com）はgohanで動いている。動機はシンプルだった。自分で完全に理解でき、変更するページだけを再生成する静的サイトジェネレータが欲しかった。

ほとんどのジェネレータは無条件全てのページを再生成するか、Git diff出力に依存するかのどちらかだ。Git diffはブランチ切り替えやフレッシュclone後に信頼性が落ちる。gohanはSHA-256コンテンツハッシングでビルドマニフェストを作成・永続化する。Git履歴に依存せず差分ビルドが常に正確に動作する。

## インクリメンタルビルドエンジン

インクリメンタルビルドのコアは`internal/diff/git.go`にある。`Detect()`メソッドが現在のワーキングツリーを永続化した`BuildManifest`と比較する。

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

`hashAllFiles()`がコンテンツディレクトリをウォークして全ファイルのSHA-256 hexダイジェストを計算する。初回ビルド（またはマニフェストが存在しない場合）は全ファイルが`Added`とみなされる。以降のビルドでは`Added`、`Modified`、`Deleted`の3種類の変更を検出する。影響を受けたHTMLページだけを再生成する。

## 機能一覧

インクリメンタルビルドに加え、gohanは標準で多くの機能を提供する。

- **i18n** — `content/en/`と`content/ja/`のようなロケールメラーディレクトリ構造。ロケール切り替えリンクを自動生成。
- **シンタックスハイライト** — Chromaによるサーバーサイドレンダリング。クライアントサイドJavaScript不要。
- **Mermaid図** — ビルド時にSVG変換またはクライアントサイドレンダリング用の`<pre class="mermaid">`として出力。
- **OGP画像生成** — 記事ごとにOpen Graph画像をビルド時に生成。
- **ページネーション** — 1ページあたりの記事数を設定可能。
- **関連記事** — タグによる類似記事リンク。
- **GitHubSourceリンク** — Markdownソースへの編集リンクを自動追加。
- **ライブリロード開発サーバー** — `gohan serve`がコンテンツを監視し保存時に自動再ビルド。

## プラグインシステム

プラグインはgohanバイナリにコンパイルされ、`config.yaml`でプロジェクトごとに有効化する。利用者側の再コンパイルは不要だ。プラグインインターフェースは`internal/plugin/plugin.go`に定義される。

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

`Plugin`（記事単位）は1つの記事に追加データを記事のテンプレートを通じて`.PluginData.<name>`として公開する。`SitePlugin`（サイト全体）は全記事処理後に実行され、Markdownソースを持たない**仮想ページ**を生成できる。

内蔵レジストリには2つのプラグインが内蔵されている。

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

`amazon_books`は記事フロントマターのASIN値からAmazonアフィリエイト本カードデータ（画像・URL・タイトル）を生成する。`bookshelf`はサイト全体の本フロントマターを集約し履歴できる仮想`/bookshelf`ページを生成する。

`config.yaml`での設定例。

```yaml
plugins:
  amazon_books:
    enabled: true
    tag: "your-associate-tag-22"
  bookshelf:
    enabled: true
```

## インストールと基本操作

```bash
# Homebrew (macOS/Linux)
brew install bmf-san/tap/gohan

# Go install
go install github.com/bmf-san/gohan/cmd/gohan@latest

# ビルド
gohan build

# ライブリロード付き開発サーバー
gohan serve
```

## まとめ

gohanはこのサイトを動かすエンジンだ。SHA-256マニフェストによるインクリメンタルビルドがイテレーションを速く保ち、コンパイル済みプラグインシステムがバイナリを自山に保つ。i18nからOGP、Mermaidまで、ビルド時はクライアントサイドJavaScript不要で動作する。

- **GitHub**: [bmf-san/gohan](https://github.com/bmf-san/gohan)
