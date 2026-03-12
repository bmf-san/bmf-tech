# bmf-tech

[bmf-tech.com](https://bmf-tech.com) のソースリポジトリ。
[gohan](https://github.com/bmf-san/gohan) 製の静的サイトで、Cloudflare Pages でホスティングしている。

![OGP Default](assets/images/ogp-default.png)

## 技術スタック

| レイヤー | 使用技術 |
|---|---|
| 静的サイトジェネレーター | [gohan](https://github.com/bmf-san/gohan) |
| ホスティング | Cloudflare Pages |
| CI/CD | GitHub Actions |
| 言語 | Go (ツール類), HTML/CSS (テーマ) |

## ディレクトリ構成

```
.
├── .github/workflows/deploy.yml  # Cloudflare Pages デプロイワークフロー
├── assets/
│   └── images/posts/             # 記事内画像 (外部→ローカル移行済み)
├── content/
│   ├── en/
│   │   ├── posts/                # 英語記事
│   │   ├── about.md              # About ページ
│   │   └── privacy-policy.md     # プライバシーポリシー
│   └── ja/
│       └── posts/                # 日本語記事 (旧ブログ移行済み 504記事)
├── docs/
│   └── DESIGN_DOC.md             # 設計ドキュメント
├── public/                       # ビルド出力 (.gitignore済み)
├── taxonomies/
│   ├── categories.yaml           # カテゴリ一覧
│   └── tags.yaml                 # タグ一覧
├── themes/default/
│   └── templates/                # HTMLテンプレート
├── _redirects                    # Cloudflare Pages リダイレクトルール
└── config.yaml                   # gohan 設定
```

## セットアップ

```bash
# リポジトリをクローン
git clone git@github.com:bmf-san/bmf-tech.git
cd bmf-tech

# gohan をインストール
make install-gohan
```

## 開発

```bash
# サイトをビルド
make build

# ローカルサーバーを起動 (http://localhost:1313)
make serve

# 新しい日本語記事を作成
make new-ja TITLE="記事タイトル" SLUG=article-slug

# 新しい英語記事を作成
make new-en TITLE="Article Title" SLUG=article-slug
```

## 記事の書き方

日本語記事は `content/ja/posts/`、英語記事は `content/en/posts/` に Markdown ファイルを置く。

フロントマターのフォーマット:

```yaml
---
title: "記事タイトル"
slug: article-slug
date: 2026-01-15
author: bmf-san
categories:
  - カテゴリ名
tags:
  - タグ1
  - タグ2
description: "記事の説明"
translation_key: article-slug  # 翻訳記事がある場合に設定
draft: false
---
```

## 翻訳記事の紐付け

`translation_key` を同じ値にすることで JA/EN 記事がリンクされる。
`hreflang` タグは自動生成される。

```yaml
# content/ja/posts/hello.md
translation_key: hello

# content/en/posts/hello.md
translation_key: hello
```

## URL 構造

| コンテンツ | URL |
|---|---|
| 英語記事 | `/posts/{slug}/` |
| 日本語記事 | `/ja/posts/{slug}/` |
| About | `/about/` |
| プライバシーポリシー | `/privacy-policy/` |
| タグ一覧 | `/tags/` |
| カテゴリ一覧 | `/categories/` |

## デプロイ

`main` ブランチへの push で自動デプロイ。

**CI フロー（`.github/workflows/deploy.yml`）:**

1. GitHub Actions (ubuntu ランナー) が `gohan build` を実行し `public/` を生成
2. `_redirects` を `public/` へコピー
3. `wrangler pages deploy public` で `public/` を Cloudflare Pages へダイレクトアップロード

> Cloudflare Pages 側ではビルドを行わない。ビルドは GitHub Actions ランナー上で完結する。
> `assets/fonts/` など `gohan build` に必要なファイルはリポジトリに含める必要がある。

**GitHub Secrets に設定が必要**:
- `CLOUDFLARE_API_TOKEN` — Cloudflare API トークン (Pages:Edit 権限)
- `CLOUDFLARE_ACCOUNT_ID` — Cloudflare アカウント ID

手動デプロイ: GitHub Actions の `workflow_dispatch` からトリガー可能。

## コントリビューション

[CONTRIBUTING.md](CONTRIBUTING.md) を参照。

## ライセンス

コンテンツ (content/) は著作権保持。ソースコード (themes/) は MIT ライセンス。