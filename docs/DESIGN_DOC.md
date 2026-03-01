# bmf-tech.com — gohan 移行設計ドキュメント

## 1. 背景

bmf-tech.com は現在 **gobel**（自作 Headless CMS）+ **Vue.js** フロントエンドで構築されている。
構成が複雑でメンテナンスコストが高く、インフラ費用もかかる状態になっている。

これを **gohan**（自作静的サイトジェネレーター）ベースの完全静的サイトに移行し、以下を実現する。

- 技術スタックの大幅な単純化
- 高速なページ配信（静的 HTML）
- コンテンツ管理の Markdown ファイル化
- インフラコストの削減

---

## 2. ゴール

- gohan で bmf-tech.com の全ページを静的生成できる
- SEO パフォーマンスを現状以上に維持・向上させる
- 記事・ページの追加・更新がローカルの Markdown 編集だけで完結する
- CI/CD でビルド・デプロイが自動化されている

---

## 3. 非ゴール

- サーバーサイド検索（/posts/search は廃止。静的サイト上での代替は Phase 後続で検討）
- コメント機能
- ユーザー認証・管理画面

---

## 4. 現在のサイト構成

### URL 一覧

| URL | 内容 |
|---|---|
| `/` | トップページ（記事一覧） |
| `/posts` | 記事一覧 |
| `/posts/search` | 記事検索 |
| `/posts/{slug}` | 記事詳細 (slug は日本語タイトルそのまま) |
| `/posts/categories/{name}` | カテゴリー別記事一覧 |
| `/posts/tags/{name}` | タグ別記事一覧 |
| `/categories` | カテゴリー一覧 |
| `/tags` | タグ一覧 |
| `/profile` | プロフィールページ |
| `/support` | サポート（投げ銭）ページ |
| `/privacy_policy` | プライバシーポリシー |
| `/sitemap` | サイトマップ |
| `/feed` | Atom フィード |

### ナビゲーション

```
bmf-tech | 記事 | カテゴリ | タグ | プロフィール
```

### カテゴリー一覧（現行）

アーキテクチャ / アプリケーション / アルゴリズムとデータ構造 / インフラストラクチャ / OS /
開発プロセス / キャリア / コンピューターアーキテクチャ / 数学 / ツール / テスト /
データベース / ネットワーク / パフォーマンス / ビジネス / ポエム / マネジメント / ヘルスケア /
運用 / 障害報告

---

## 5. 新サイト構成（gohan 移行後）

### URL 設計

| URL | gohan テンプレート | 内容 |
|---|---|---|
| `/` | `index.html` | トップページ（記事一覧） |
| `/posts/{slug}/` | `article.html` | 記事詳細 |
| `/categories/{name}/` | `category.html` | カテゴリー別記事一覧 |
| `/tags/{name}/` | `tag.html` | タグ別記事一覧 |
| `/archive/{year}/` | `archive.html` | 年別アーカイブ |
| `/about/` | （ページ → 静的 HTML） | プロフィール |
| `/support/` | （ページ → 静的 HTML） | サポートページ |
| `/privacy-policy/` | （ページ → 静的 HTML） | プライバシーポリシー |
| `/sitemap.xml` | gohan 自動生成 | XML サイトマップ |
| `/atom.xml` | gohan 自動生成 | Atom フィード |

> **スラッグ方針**: 既存の日本語 slug からマイグレーション時に英語のハイフン区切り slug に変換する。
> 旧 URL → 新 URL への 301 リダイレクトマップを別途管理する（CDN / Web サーバーレベルで対応）。

### ナビゲーション（案）

```
bmf-tech | 記事 | カテゴリ | タグ | アーカイブ | プロフィール
```

---

## 6. ディレクトリ構成

```
bmf-tech-by-gobel/
├── config.yaml              # gohan 設定
├── content/
│   ├── posts/               # 記事 Markdown ファイル
│   └── pages/               # 固定ページ（about, support, privacy-policy）
├── assets/
│   ├── css/
│   │   └── style.css
│   ├── js/                  # 必要に応じて（Mermaid 等は gohan が自動注入）
│   └── images/
│       └── profile.png
├── themes/
│   └── default/
│       └── templates/
│           ├── index.html
│           ├── article.html
│           ├── tag.html
│           ├── category.html
│           ├── archive.html
│           └── _partials/
│               ├── head.html       # <head> 共通テンプレート（OGP/SEO）
│               ├── header.html     # グローバルナビ
│               ├── footer.html     # フッター
│               └── article-card.html  # 記事サムネイルカード
└── docs/
    └── DESIGN_DOC.md
```

---

## 7. config.yaml スキーマ

```yaml
site:
  title: "bmf-tech.com"
  description: "シニアプラットフォームエンジニア Kenta Takeuchi の技術ブログ。Go・アーキテクチャ・インフラ・開発プロセスを中心に発信。"
  base_url: https://bmf-tech.com
  language: ja

build:
  content_dir: content
  output_dir: public
  assets_dir: assets
  parallelism: 4

theme:
  name: default
  dir: themes/default
  params:
    author: "Kenta Takeuchi"
    github: "bmf-san"
    twitter: "bmf_san"
    linkedin: "bmf-san"
    zenn: "bmf_san"
    speaker_deck: "bmf_san"
    footer_text: "© 2026 Kenta Takeuchi"
    adsense_id: ""          # Google AdSense クライアント ID
    ga_id: ""               # Google Analytics 測定 ID
```

---

## 8. Front Matter スキーマ

### 記事（content/posts/）

```yaml
---
title: "記事タイトル"
date: 2024-01-15
draft: false
slug: "article-slug-in-english"   # URL に使用。英語ハイフン区切り
description: "meta description（SEO 用、120〜160 文字目安）"
author: "bmf-san"
tags:
  - Go
  - HTTP
categories:
  - アーキテクチャ
---
```

### ページ（content/pages/）

```yaml
---
title: "プロフィール"
slug: "about"
description: "Kenta Takeuchi のプロフィールページ"
---
```

---

## 9. SEO 設計

個人テックブログとして以下を重視する。

### 9.1 テクニカル SEO

| 施策 | 実装方法 |
|---|---|
| `<title>` タグ最適化 | `{記事タイトル} — bmf-tech.com` 形式 |
| `<meta name="description">` | Front Matter の `description` フィールドを使用 |
| OGP タグ | `head.html` パーシャルに `og:title / og:description / og:url / og:image` を実装 |
| Twitter Card | `twitter:card / twitter:site / twitter:creator` を `head.html` に実装 |
| Canonical URL | `<link rel="canonical">` を全ページに設定。`{{.Config.Site.BaseURL}}/posts/{{slug}}/` |
| JSON-LD (Article) | `article.html` に `@type: BlogPosting` の構造化データを埋め込む |
| JSON-LD (BreadcrumbList) | 記事・カテゴリー・タグページにパンくずリストを実装 |
| sitemap.xml | gohan が自動生成 |
| robots.txt | `/assets/robots.txt` に配置。全クロール許可 + sitemap 参照 |
| Atom フィード | gohan が自動生成 (`/atom.xml`) |

### 9.2 コンテンツ SEO

| 施策 | 方針 |
|---|---|
| 英語スラッグ | 日本語タイトルの URL エンコードを避け、意味のある英語スラッグを設定 |
| カテゴリー・タグの活用 | トピッククラスタリングを意識したカテゴリー設計。カテゴリー数は現行を維持し追加・整理は段階的に実施 |
| アーカイブページ | `/archive/{year}/` で年別コンテンツを集約し時系列インデックスを提供 |
| 内部リンク | 記事内で関連記事へのリンクを手動で記述（テンプレートに「関連記事」セクションを追加することも検討） |
| 画像 alt テキスト | Markdown の画像記法で alt を必ず記述 |

### 9.3 パフォーマンス（Core Web Vitals）

- 静的 HTML ＋ 最小限の CSS/JS → LCP・FID・CLS を最小化
- 画像は `loading="lazy"` を付与
- フォントは system-font-stack を優先し Web フォントを極力使わない
- CSS は 1 ファイルにまとめ minify する

### 9.4 旧 URL からのリダイレクト

現行の日本語スラッグ URL から新 English slug URL への 301 リダイレクトが必要。
マイグレーション時に `docs/redirects.md` でマッピング表を管理し、
デプロイ先（CDN / Nginx / Cloudflare）に適したリダイレクトルールを生成する。

---

## 10. テンプレート設計指針

### 共通

- レスポンシブデザイン（モバイルファースト）
- ライト / ダークテーマ切り替え（CSS Custom Properties で実装）
- コードブロックは gohan の chroma によりサーバーサイドでシンタックスハイライト
- Mermaid ブロックは gohan が自動でスクリプト注入

### `head.html`（パーシャル）

```html
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>{{block "title" .}}{{.Config.Site.Title}}{{end}}</title>
<meta name="description" content="{{block "description" .}}{{.Config.Site.Description}}{{end}}">
<link rel="canonical" href="{{block "canonical" .}}{{.Config.Site.BaseURL}}/{{end}}">
<!-- OGP -->
<meta property="og:type" content="{{block "og_type" .}}website{{end}}">
<meta property="og:title" content="{{block "og_title" .}}{{.Config.Site.Title}}{{end}}">
<meta property="og:description" content="{{block "og_description" .}}{{.Config.Site.Description}}{{end}}">
<meta property="og:url" content="{{block "og_url" .}}{{.Config.Site.BaseURL}}{{end}}">
<meta property="og:site_name" content="{{.Config.Site.Title}}">
<!-- Twitter Card -->
<meta name="twitter:card" content="summary">
<meta name="twitter:site" content="@{{.Config.Theme.Params.twitter}}">
<meta name="twitter:creator" content="@{{.Config.Theme.Params.twitter}}">
<!-- Feed -->
<link rel="alternate" type="application/atom+xml" title="{{.Config.Site.Title}}" href="{{.Config.Site.BaseURL}}/atom.xml">
```

### `article.html`

- Front Matter の `description` を `<meta name="description">` に設定
- 公開日・著者・カテゴリー・タグを表示
- `@type: BlogPosting` の JSON-LD を出力
- コードブロックをコピーボタン付きで表示（JS で実装）

---

## 11. 移行フェーズ

| フェーズ | 内容 | 状態 |
|---|---|---|
| Phase 1 | 設計ドキュメント作成（本ドキュメント） | ✅ |
| Phase 2 | gohan の設定ファイルとディレクトリ構成のスキャフォールディング | ⏳ |
| Phase 3 | テーマ・テンプレート開発（HTML/CSS） | ⏳ |
| Phase 4 | 既存記事の Markdown マイグレーション（gobel API → `.md` ファイル） | ⏳ |
| Phase 5 | リダイレクトマップの作成と検証 | ⏳ |
| Phase 6 | CI/CD パイプライン構築（gohan build → デプロイ） | ⏳ |
| Phase 7 | 本番リリース・DNS 切り替え | ⏳ |

---

## 12. 課題・検討事項

| 項目 | 内容 |
|---|---|
| 検索機能 | 現行の `/posts/search` は gohan では提供できない。Pagefind などのクライアントサイド全文検索の採用を検討 |
| ページネーション | 記事が 300+ 件あるため `index.html` の無限スクロール or ページ分割が必要。gohan の現行テンプレート仕様では自動ページネーションなし → JavaScript での無限スクロールか、カテゴリー・タグ・アーカイブへの誘導で対応 |
| 広告 | Google AdSense を継続運用。テンプレートに広告スロットを設ける |
| OGP 画像 | 記事サムネイル画像の自動生成またはデフォルト画像の設定 |
| コンテンツ移行ツール | gobel API からの記事エクスポートスクリプトが必要。スラッグ変換ロジックも含む |
| `content/pages/` のルーティング | gohan の現行仕様では `pages` 型のコンテンツは `/pages/{slug}/` に配置される。`/about/` 等の短い URL にするには生成後のファイル移動、または templates への組み込みが必要 |
