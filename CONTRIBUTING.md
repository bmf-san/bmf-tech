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

### 用語集に用語を載せる

記事で解説している用語は、フロントマターの `glossary:` に書くと用語集ページ（`/ja/glossary/`、英語版は `/glossary/`）に載る。用語集は各用語から定義元の記事へリンクする索引で、用語ごとの個別ページは作らない。

```yaml
glossary:
  - term: "PEP"
    reading: "ペップ"   # 任意。読み方として表示し、漢字の用語では並び順にも使う
    definition: "認可の判定を要求し、その結果に従ってアクセスを強制する構成要素。"
```

- `term` と `definition` は必須。定義は1〜2文のプレーンテキストで書く
- その記事で詳しく解説している用語だけを載せる。言及しているだけの用語は載せない
- 同じ用語を複数の記事に書くと1件にまとまり、定義は最も新しい記事のものが使われる
- 日英それぞれの記事に書く。英語版の記事には英語の定義を書く
- どの記事にも用語がない間は、用語集ページもヘッダー・フッターのリンクも生成されない（言語ごとに判定する）
- 用語を書いた記事の末尾（タグの上）には、その記事の用語から用語集の該当箇所へのリンクが出る

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

デプロイ・CI ワークフロー・必要な Secrets などは [docs/OPERATIONS.md](docs/OPERATIONS.md) を参照。

ライセンスについては [README.md](README.md#ライセンス) を参照。
