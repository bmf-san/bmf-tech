# Contributing

## 記事を書く

### 新規記事の作成

```bash
# 日本語記事
make new-ja TITLE="記事タイトル" SLUG=article-slug

# 英語記事
make new-en TITLE="Article Title" SLUG=article-slug
```

### フロントマター

```yaml
---
title: "記事タイトル"
slug: article-slug          # 英語ハイフン区切り・すべて小文字
date: 2026-01-15
author: bmf-san
categories:
  - カテゴリ名
tags:
  - タグ1
description: "記事の説明（120〜160 文字目安）"
translation_key: article-slug  # 日英対応記事がある場合
draft: false
---
```

### スラッグ規則

- 英語ハイフン区切り・すべて小文字（例: `go-http-server`, `docker-compose-tips`）
- 既存 URL を変更する場合は `_redirects` に 301 リダイレクトを追加する

## ソースコードの変更

### テンプレート (`themes/default/templates/`)

- CSS は [sleyt](https://github.com/bmf-san/sleyt) のユーティリティクラスを使用する
- CSS の追加が必要な場合は sleyt リポジトリ側に追加し、バージョンを上げる
- テンプレートにインライン `<style>` タグは追加しない

## ブランチ・PR

- `main` への直接 push は禁止（ブランチ保護あり）
- ブランチ名の目安: `feat/`, `fix/`, `docs/`, `design/`
- PR はレビュー不要（セルフマージ可）

## 運用ドキュメント

デプロイ・dev.to 自動投稿・CI ワークフロー・必要な Secrets などは [docs/OPERATIONS.md](docs/OPERATIONS.md) を参照。

ライセンスについては [README.md](README.md#ライセンス) を参照。
