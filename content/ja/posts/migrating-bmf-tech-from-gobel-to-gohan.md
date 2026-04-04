---
title: "bmf-tech.comの gobel → gohan 移行記—完全静的サイト化の流れ"
description: '自作ヘッドレスCMS（gobel）+ MySQL + Vue.js構成から、gohanによる完全静的サイトへ移行しCloudflare Pagesでホスティングするまでの全工程。700件超の記事、英語スラッグ生成、画像移行、リダイレクト、Go製プリフライトチェッカーまで。'
slug: migrating-bmf-tech-from-gobel-to-gohan
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - gohan
  - Migration
  - SSG
  - Golang
translation_key: migrating-bmf-tech-from-gobel-to-gohan
---

# bmf-tech.comの gobel → gohan 移行記

## バックグラウンド

bmf-tech.comはインフラ的に4世代を経てきた。

1. **WordPress** — 独自のテンプレートを実装して運用していたが、メンテナンスコストとセキュリティパッチに音を上げて破棄。
2. **[Rubel](https://github.com/bmf-san/Rubel)** — Laravelで実装したヘッドレスCMS。Reactフロントエンド + MySQLをConoHa VPS上で運用。
3. **[gobel](https://github.com/bmf-san/gobel-api)** — RubelをGo製にリライト。数年間運用し700件超の記事を蓄積。バックエンドはNginx + MySQL + Redisをdocker-composeで構成し、フロントエンドはVue.jsで実装した。Prometheus・Grafana・Loki・Pyroscopeによる監視スタックも整備していた。
4. **[gohan](https://github.com/bmf-san/gohan)** — Go製のSSG。`content/en/`と`content/ja/`のMarkdownファイルをビルドソースとし、GitHub ActionsがビルドしてCloudflare Pagesで静的配信する。サーバーレスで運用コストがゼロ。

目標はサーバーコストと運用負荷をゼロにすること。GitHubPushで自動ビルド＆デプロイされるMarkdownファイルのみでサイトを組む構成を実現したいと思った。

## 移行フェーズ一覧

移行は12フェーズで計画した。

| フェーズ | 内容 |
|---|---|
| 0 | 設計ドキュメント・移行計画策定 |
| 1 | リポジトリ・ gohan プロジェクトセットアップ |
| 2 | テーマ・テンプレート開発 |
| 3 | 英語スラッグ変換マップ作成（700件超） |
| 4 | コンテンツデータ移行スクリプト（SQLダンプ → Markdown） |
| 5 | 日本語記事のMarkdown移行（`content/ja/`） |
| 6 | 画像アセット移行 |
| 7 | リダイレクトマップ作成・検証 |
| 8 | 英語版記事作成（高優先記事から順次） |
| 9 | CI/CDパイプライン構築 |
| 9.5 | 移行前プリフライトチェック |
| 10 | Cloudflare Pages本番デプロイ |
| 11 | DNS移管（ConoHa → Cloudflare） |
| 12 | ConoHa VPS停止 |

### Phase 3: 英語スラッグ生成

元のgobelデータベースに`slug`カラムはなかった。URLは日本語タイトルをURLエンコードする形式で構成されていた。1例を挙げると`/posts/Go%E3%81%A7HTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%82%92%E6%9B%B8%E3%81%8F`のような形式。移行後は`/posts/go-http-server/`のような英語スラッグにしたかった。

700件超のスラッグを人手で考えるのは現実的でないため、Claudeに`(id, title)`のペアを一括で渡して英語スラッグ候補を大量生成した。出力を`slug_map.csv`で手動レビューし重複や意味確認を行った。

### Phase 4: 移行スクリプト

Go製の移行ツールはMySQLダンプと`slug_map.csv`の2つを入力に受け取り、以下の2つを出力した。

- `content/ja/posts/*.md` — `posts`・`categories`・`tag_post`テーブルからFront Matterを生成した記事ごとのMarkdownファイル。
- `_redirects` — 旧URLから新URLへの301リダイレクトルールファイル。Cloudflare Pages専用形式。

`_redirects`の生成では注意点があった。Cloudflare Pagesはリクエストパスをパーセントエンコード済みの形式で受け取る。そのためスクリプト内で旧URLに`url.PathEscape()`を適用する必要があり、生の日本語文字列をルールに書くとリダイレクトが当たらなかった。

### Phase 6: 画像アセット移行

元サイトの画像は全て外部CDN（主にQiita Image Store）にホストされていた。ダウンロードツールが各記事のMarkdownをスキャンし、画像を`assets/images/posts/{slug}/`以下に保存し、`![alt](url)`をローカルパスの参照に書き換えた。

### Phase 9.5: プリフライトチェッカー

DNS切り替え前に、Go製のプリフライトツールのCloudflare PagesのプレビューURLに対して全記事URLにHTTPリクエストを送りHTTP 200を検証した。また`_redirects`の全ルールについてもいずれも301で返ることを確認。DNS切り替え後に記事が404にならないという信頼を切り替え前に得た。

### Phase 11: DNS移管

ConoHaからCloudflareにネームサーバーを切り替えた。Cloudflare Pagesがgohanビルドの静的ファイルをCDNから直接配信する。`www.bmf-tech.com`はサーバーサイドコードが一切動いていない完全静的サイトになった。

### Phase 12: ConoHa VPS停止

Google Search Consoleでリダイレクトのインデックスカバレッジの推移を確認後、gobel APIとMySQLを動かしていたConoHa VPSを停止した。月額VPSコストがゼロになる。

## 移行後の構成

現在のサイトはgohanが`content/en/`と`content/ja/`内のMarkdownファイルから生成する。GitHubにプッシュするとGitHub ActionsがビルドしCloudflare Pagesにデプロイする。584件超のEN記事と584件超のJA記事の全ビルド（OGP画像生成含む）は60秒ほどで完了する。

旧URLから新URLへの301リダイレクトを`_redirects`に定義したことで、SEOへの影響を最小限に抑えた。DNS移管後にGoogle Search Consoleでインデックスカバレッジが維持されていることを確認した。

### 移行によって変わったこと

- **執筆環境** — 独自のWeb管理画面でのMarkdown入力から、VS Codeによるローカル執筆に変わった。Gitで記事を管理するためバージョン履歴も自然に残る。
- **品質管理** — GitHub ActionsのCIでtextlintを実行するようにした。記事をプッシュするたびに表記ゆれや文章ルール違反を自動検出できる。
- **メタ情報の整備** — Front Matterで`title`・`description`・`categories`・`tags`などを記事ごとに明示的に設定できるようになった。gobelではサボりがちだったOGP設定も含めて適切に管理している。
- **画像の自前管理** — 画像をQiita Image Storeなどの外部CDNに依存していたが、`assets/images/posts/{slug}/`以下で静的リソースとして自前管理するようになった。外部サービスの廃止による画像消失リスクがなくなった。
- **URLのスラッグ化** — 記事URLが日本語タイトルをパーセントエンコードした形式から、英語スラッグ形式（`/posts/go-http-server/`）に変わった。可読性と共有しやすさが向上した。
- **i18n対応** — `content/ja/`と`content/en/`の2言語構成とし、`translation_key`で記事同士を紐付けるi18n対応を実現した。
- **ホスティング先の変更** — ConoHa VPS上で自前運用していたサーバーをCloudflare Pagesに移行した。静的コンテンツをCDNから最速で配信でき、無料枠が充実しているためサーバー代が発生しない。
- **独自機能の拡張** — 本棚ページ（読んだ本の一覧管理）やAmazon Associateリンク発行ページなど、独自のページ構成を追加した。

## 参考リンク

- **bmf-techソース**: [bmf-san/bmf-tech](https://github.com/bmf-san/bmf-tech)
- **gohan**: [bmf-san/gohan](https://github.com/bmf-san/gohan)
