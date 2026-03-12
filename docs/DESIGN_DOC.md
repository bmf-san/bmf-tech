# Design Document: Gohan - GoベースのSSG

## bmf-tech.com サイト設計

---

## 1. サイト構成

### URL 設計

| URL | gohan テンプレート | 内容 |
|---|---|---|
| `/` | `index.html` | トップページ（記事一覧） |
| `/posts/{slug}/` | `article.html` | 記事詳細 |
| `/categories/{name}/` | `category.html` | カテゴリー別記事一覧 |
| `/tags/{name}/` | `tag.html` | タグ別記事一覧 |
| `/archives/{year}/{month}/` | `archive.html` | 年月別アーカイブ |
| `/about/` | `page.html` | プロフィール |
| `/privacy-policy/` | `page.html` | プライバシーポリシー |
| `/sitemap.xml` | gohan 自動生成 | XML サイトマップ |
| `/atom.xml` | gohan 自動生成 | Atom フィード |

> `/support/` は当初設計に含まれていたが v1 では未実装。

> **スラッグ方針**: 意味のある英語ハイフン区切り slug を使用する。
> - 単語は `-` 区切り、すべて小文字
> - 例: `go-http-server`, `docker-compose-tips`
> - 既存の日本語スラッグは gobel API から全記事を取得し、記事ごとに英語 slug を設定してマイグレーションスクリプトで変換する
> - 旧 URL → 新 URL の 301 リダイレクトは Cloudflare Pages の `_redirects` で永続管理する

### ナビゲーション

```
bmf-tech | 日本語 | Tags | Categories | About | Feed
```

ヘッダー: `themes/default/templates/_partials.html` の `{{define "header"}}` ブロック。

フッター: About / Privacy Policy / Contact リンク + copyright。

---

## 2. ディレクトリ構成

```
bmf-tech/
├── config.yaml              # gohan 設定
├── _redirects               # Cloudflare Pages リダイレクトルール（CI で public/ へコピー）
├── content/
│   ├── en/                  # 英語コンテンツ（デフォルトロケール）
│   │   ├── posts/           # 英語記事 Markdown ファイル
│   │   ├── about.md         # About ページ
│   │   ├── privacy-policy.md
│   │   ├── categories.md    # /categories/ インデックスページ（http-server 用）
│   │   └── tags.md          # /tags/ インデックスページ
│   └── ja/
│       └── posts/           # 日本語記事 Markdown ファイル（584件）
├── assets/
│   ├── images/
│   │   └── posts/{slug}/    # Phase 6 で移行した記事内画像
│   └── robots.txt           # gohan build 時に public/ へ自動コピー
├── themes/
│   └── default/
│       └── templates/
│           ├── _partials.html   # head / header / footer / pagination 共通ブロック（CSS インライン）
│           ├── index.html
│           ├── article.html
│           ├── page.html        # about / privacy-policy 等の固定ページ
│           ├── tag.html
│           ├── category.html
│           └── archive.html
├── taxonomies/
│   ├── tags.yaml
│   └── categories.yaml
├── tools/
│   ├── populate_book_asins.py         # 書籍記事に ASIN を補完するスクリプト
│   └── populate_en_book_asins_v2.py   # EN 書籍記事用 v2
└── docs/
    ├── DESIGN_DOC.md
    └── MIGRATION.md
```

---

## 3. config.yaml スキーマ

実際の `config.yaml` は以下の構造（現在値）：

```yaml
site:
  title: "bmf-tech"
  description: "bmf-san's personal tech blog"
  base_url: "https://bmf-tech.com"
  language: "en"
  github_repo: "https://github.com/bmf-san/bmf-tech"
  github_branch: "main"

build:
  content_dir: "content"
  output_dir: "public"
  assets_dir: "assets"
  parallelism: 4
  per_page: 20
  exclude_files: []

theme:
  name: "default"
  dir: "themes/default"
  params:
    author: "Kenta Takeuchi"
    github: "bmf-san"
    twitter: "bmf_san"
    linkedin: "bmf-san"
    zenn: "bmf_san"
    speaker_deck: "bmf_san"
    footer_text: "© 2026 Kenta Takeuchi"
    adsense_id: "ca-pub-5146230866088201"
    adsense_slot_article_top: "3773998823"
    adsense_slot_article_bottom: "2224967643"
    adsense_slot_list_mobile: "4950844549"
    ga_id: "G-784B55NW88"

syntax_highlight:
  theme: "github"
  line_numbers: false

ogp:
  enabled: true
  background_color: "#ffffff"
  text_color: "#111111"
  font_file: "assets/fonts/NotoSansJP-Bold.ttf"  # OFL ライセンス（フリー、商用可）
  logo_file: ""
  width: 1200
  height: 630

i18n:
  default_locale: "en"
  locales:
    - "en"
    - "ja"
```

> noindex は DNS 移管完了後に削除済み（本番公開中）。

---

## 4. Front Matter スキーマ

### 記事（content/ja/posts/ および content/en/posts/）

```yaml
---
title: "記事タイトル"
date: 2024-01-15
draft: false
slug: "article-slug-in-english"   # URL に使用。英語ハイフン区切り
description: "meta description（SEO 用、120〜160 文字目安）"
author: "bmf-san"
translation_key: "article-slug-in-english"  # 日英対応付けに使用
tags:
  - Go
  - HTTP
categories:
  - アーキテクチャ
---
```

### 固定ページ（content/en/）

```yaml
---
title: "About"
slug: "about"
date: 2024-01-01
author: "bmf-san"
template: page.html
draft: false
description: "Kenta Takeuchi のプロフィールページ"
---
```

---

## 5. SEO 設計

個人テックブログとして以下を重視する。

### 5.1 テクニカル SEO

| 施策 | 実装状況 |
|---|---|
| `<title>` タグ最適化 | `{記事タイトル} — {site.title}` 形式（`article.html` の `{{define "title"}}`） |
| `<meta name="description">` | Front Matter の `description` フィールドを使用。EN/JA 全記事（各 584 件）に `description` 設定済み |
| OGP タグ | `og:title / og:type / og:url / og:description / og:image` を `_partials.html` の `{{define "head"}}` に実装済み。記事ページは `ogp/{slug}.png`（ビルド時自動生成）、一覧ページは `assets/images/ogp-default.png` を使用 |
| Twitter Card | `twitter:card / twitter:site / twitter:creator` を実装済み |
| Canonical URL | `<link rel="canonical">` を `article.html` に実装済み |
| hreflang | 日英ペア記事に `hreflang="ja"` / `hreflang="en"` / `hreflang="x-default"` を出力（gohan が `translation_key` で自動処理） |
| JSON-LD (Article) | 実装済み。`_partials.html` の `{{define "seo"}}` 内で `@type: BlogPosting`（headline / datePublished / dateModified / url / image / author / publisher / description）を出力 |
| JSON-LD (BreadcrumbList) | **未実装** |
| sitemap.xml | gohan が自動生成。記事・固定ページの URL + ロケール別インデックスページ（`/`・`/ja/`）を収録（タグ・カテゴリー・アーカイブ個別ページは含まれない） |
| robots.txt | `assets/robots.txt` に配置済み。`Allow: /` + sitemap 参照 |
| Atom フィード | gohan が自動生成 (`/atom.xml`) |
| noindex | DNS 移管完了後に削除済み。本番公開中（`config.yaml` から `noindex` 設定を除去した）|

### 5.2 コンテンツ SEO

| 施策 | 方針 |
|---|---|
| 英語スラッグ | 日本語タイトルの URL エンコードを避け、意味のある英語スラッグを設定 |
| カテゴリー・タグの活用 | トピッククラスタリングを意識したカテゴリー設計。カテゴリー数は現行を維持し追加・整理は段階的に実施 |
| アーカイブページ | `/archives/{year}/{month}/` で年月別コンテンツを集約（gohan 実装済み） |
| 内部リンク | 記事内で関連記事へのリンクを手動で記述 |
| GitHub ソースリンク | 各記事フッターに `Edit on GitHub` リンクを表示（`ContentPath` + `Config.Site.GitHubRepo` で生成） |
| 画像 alt テキスト | Markdown の画像記法で alt を必ず記述 |

### 5.3 パフォーマンス（Core Web Vitals）

- 静的 HTML + CSS のみ（JS 依存なし）→ LCP・FID・CLS を最小化
- CSS はインライン（`_partials.html` の `<style>` タグ）。外部 CSS 依存なし
- システムフォントスタック使用（Web フォント不使用）
- 画像は `loading="lazy"` を付与

---

## 6. テンプレート設計指針

### CSS

[sleyt](https://github.com/bmf-san/sleyt)（bmf-san 製ミニマル CSS フレームワーク）を CDN 経由で読み込む。独自 `<style>` タグは使用しない。

```html
<link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/bmf-san/sleyt@latest/docs/public/css/index.css">
```

- **`sleyt`** (v1.3.2): CSS リセット・デザイントークン変数（`--color-slate-*` / `--text-primary` 等）・レイアウトユーティリティ（`container`, `flex`, `grid`, `gap-*`, `px-*` 等）・タイポグラフィ（`text-sm`, `font-bold` など）・ダークモード対応を提供。`pre[style]`（chroma シンタックスハイライト）と `.adsbygoogle` のダークモード補正も含む
- CSS の追加が必要な場合は sleyt リポジトリ側に追加してバージョンを上げる
- インライン `style=""` は Google AdSense 必須属性 (`<ins style="display:block">`) のみ許容

### テンプレート一覧

| ファイル | 用途 |
|---|---|
| `_partials.html` | `head` / `header` / `footer` / `pagination` の共通ブロック。CSS もここに集約 |
| `index.html` | トップページ（記事一覧、ページネーション） |
| `article.html` | 記事詳細。canonical / hreflang / GitHub ソースリンク / 関連記事（2カラムグリッド、OGP 画像・カテゴリ・タグ・説明文付き）を出力 |
| `page.html` | 固定ページ（about, privacy-policy）。`{{define "htmllang"}}` で locale を設定 |
| `tag.html` | タグ別記事一覧 |
| `category.html` | カテゴリー別記事一覧 |
| `archive.html` | 年月別アーカイブ |

### `_partials.html` — `{{define "head"}}` の主要要素

```html
<html lang="{{block "htmllang" .}}{{.Config.Site.Language}}{{end}}">
<head>
<title>{{block "title" .}}{{.Config.Site.Title}}{{end}}</title>
<meta name="description" content="...">
{{- if index .Config.Theme.Params "noindex"}}
<meta name="robots" content="noindex, nofollow">
{{- end}}
<link rel="alternate" type="application/atom+xml" href="/atom.xml" ...>
{{block "seo" .}}{{end}}
<style>/* インライン CSS */</style>
```

### `article.html` — 記事フッターの GitHub リンク

```html
{{if $.Config.Site.GitHubRepo}}
<footer class="article-footer">
  <a href="{{$.Config.Site.GitHubRepo}}/blob/{{$.Config.Site.GitHubBranch}}/content/{{.ContentPath}}"
     target="_blank" rel="noopener">Edit on GitHub</a>
</footer>
{{end}}
```

### robots.txt

`assets/robots.txt` に配置。gohan のビルド時に `assets/` の内容が `public/` ルートに直接コピーされるため、CI での手動コピーは不要。

---

## 7. 課題・検討事項

| 項目 | 内容 |
|---|---|
| sitemap.xml のタクソノミー・アーカイブ URL 欠落 | 個別タグ・カテゴリー・アーカイブページが含まれない（ロケール別インデックスページ `/` / `/ja/` は v0.1.7 で対応済み）。gohan 側 enhancement として対応予定 |
| JSON-LD | `@type: BlogPosting` は実装済み。BreadcrumbList は未実装 |
| 検索機能 | Pagefind などのクライアントサイド全文検索の採用を検討 |
| 広告 | Google AdSense 実装済み。`_partials.html` の head に AdSense スクリプトを配置。スロットは3種（`adsense_slot_article_top`: 3773998823 / `adsense_slot_article_bottom`: 2224967643 / `adsense_slot_list_mobile`: 4950844549）をテンプレートから参照 |
| Google Analytics | GA4 実装済み。`_partials.html` の head に gtag.js スクリプトを配置（ID: `G-784B55NW88`） |
| OGP 画像カスタマイズ | デフォルト画像（`assets/images/ogp-default.png`）はシンプルなテキストのみ。ブランドロゴや背景画像を使ったデザイン改善が可能 |
| タグ・カテゴリーページの多言語混在 | `/tags/{name}/` / `/categories/{name}/` は en + ja の記事が混在して出力される（gohan の現仕様）。テンプレートでロケールラベルを表示するなど UX 面での対処が必要 |
| カテゴリー英語化 | 既存カテゴリーは日本語（例: アーキテクチャ）。英語名（例: Architecture）への統一は段階的に実施予定 |
| `site.title` / `site.description` の最終化 | 現在 `"bmf-tech"` / `"bmf-san's personal tech blog"` はプレースホルダー。DNS 移管前に正式値に更新する |
