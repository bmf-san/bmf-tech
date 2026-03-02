# 移行計画: gobel → gohan

## 概要

bmf-tech.com は現在 **gobel**（自作 Headless CMS + MySQL）+ **Vue.js** フロントエンドで構築されている。
これを **gohan**（自作 SSG）ベースの完全静的サイト + Cloudflare Pages 配信に移行する。

---

## ゴールと非ゴール

### ゴール

- gohan で bmf-tech.com の全ページを静的生成できる
- 英語を主体言語とし日本語を翻訳版として提供する（i18n 対応）
- 既存 URL への 301 リダイレクトで SEO 評価を継承する
- 記事・ページの追加・更新がローカルの Markdown 編集だけで完結する
- CI/CD でビルド・デプロイが自動化されている

### 非ゴール（v1 スコープ外）

- サーバーサイド検索（Pagefind 等は別途検討）
- コメント機能
- ユーザー認証・管理画面
- サーバーサイド移行（別途計画予定）

---

## 移行フェーズ一覧

| フェーズ | 内容 | 状態 |
|---|---|---|
| Phase 0 | 設計ドキュメント・移行計画策定 | ✅ |
| Phase 1 | gohan i18n 対応実装（PR #57） | ✅ |
| Phase 2 | コンテンツデータ移行スクリプト作成 | ⏳ |
| Phase 3 | 英語スラッグ変換マップ作成（700+ 記事） | ⏳ |
| Phase 4 | 日本語記事の Markdown 移行（`content/ja/`） | ⏳ |
| Phase 5 | 画像アセット移行 | ⏳ |
| Phase 6 | テーマ・テンプレート開発 | ⏳ |
| Phase 7 | 英語版記事の作成（高優先記事から） | ⏳ |
| Phase 8 | リダイレクトマップ作成・検証 | ⏳ |
| Phase 9 | CI/CD パイプライン構築 | ⏳ |
| Phase 10 | Cloudflare Pages 本番デプロイ | ⏳ |
| Phase 11 | DNS 移管（ConoHa → Cloudflare） | ⏳ |
| Phase 12 | ConoHa VPS 停止 | ⏳ |

---

## 移行元データ構造

### gobel MySQL スキーマ概要

```
admins        (id, name, email, password, created_at, updated_at)
categories    (id, name, created_at, updated_at)
tags          (id, name, created_at, updated_at)
posts         (id, admin_id, category_id, title, md_body, html_body,
               status, created_at, updated_at)
tag_post      (id, tag_id, post_id, created_at, updated_at)
comments      (id, post_id, body, status, created_at, updated_at)
```

### データ件数（2026-03-01 ダンプ）

| テーブル | AUTO_INCREMENT | 備考 |
|---|---|---|
| posts | 725 | `status='published'` のものが移行対象 |
| categories | 29 | 現行アクティブ 20 カテゴリー |
| tags | 549 | IDに欠番あり |
| tag_post | 3,757 | タグ—記事の多対多 |
| comments | — | 移行しない |

### 重要な制約

- **posts に `slug` カラムなし**: URL は現行フロントエンドが `title` を URL エンコードして構成している
  - 現行 URL 例: `/posts/GoでHTTPサーバーを書く`（日本語 URL エンコード済み）
  - 移行後 URL 例: `/posts/go-http-server/`（英語スラッグ）
- **`md_body` に Markdown が格納されている**: `html_body` は不要
- **画像は全て外部 CDN（主に Qiita Image Store）に存在**

---

## Phase 2: コンテンツデータ移行スクリプト

### スクリプト概要

Go または Python で実装。以下の処理を行う。

```
INPUT:
  - bmf-tech_YYYY-MM-DD.sql  （MySQL ダンプ）
  - slug_map.csv             （id → english_slug 変換マップ）

OUTPUT:
  - content/ja/posts/*.md    （移行記事 Markdown ファイル）
  - _redirects               （Cloudflare Pages リダイレクトルール）
```

### 処理フロー

```
1. SQL ダンプをパース（または MySQL に直接接続）
2. slug_map.csv を読み込み id → slug のマップを構築
3. posts テーブルの published 記事を走査:
   a. category_id → category_name を解決
   b. tag_post から post_id に紐づく tag_id 一覧を取得 → tag_name 解決
   c. 英語スラッグを slug_map から取得
   d. Front Matter を生成（下記参照）
   e. md_body を本文として .md ファイルを出力
      → content/ja/posts/{slug}.md
4. _redirects ファイルを生成
   → /posts/{url_encoded_title}  /ja/posts/{slug}/  301
```

### 出力 Markdown フォーマット

```markdown
---
title: "GoでHTTPサーバーを書く"
date: 2023-05-10
slug: go-http-server
translation_key: go-http-server
tags:
  - Go
  - HTTP
categories:
  - アーキテクチャ
description: ""
author: bmf-san
---

（md_body の内容をそのまま出力）
```

### 実装メモ

- `md_body` の画像パスはそのまま出力する（Phase 5 で差し替え）
- `description` は空で出力し、後から手動または AI で補完する
- `draft: false` の記事のみ対象（`status = 'published'`）
- カテゴリーは 1 記事につき 1 つ（gobel の制約）

---

## Phase 3: 英語スラッグ変換マップ

### 作業方針

1. 全 published 記事の `(id, title)` をエクスポート
2. AI（Claude / GPT）で一括英語スラッグ候補を生成
3. CSV で手動レビュー（重複・意味確認）
4. 確定後 `slug_map.csv` として管理

### slug_map.csv フォーマット

```csv
id,title,slug,reviewed
1,OSSをはじめてみた話,getting-started-with-oss,false
263,GolangのHTTPサーバーのコードリーディング,go-http-server-code-reading,false
```

### スラッグ命名規則

- 全て小文字、`-` 区切り
- 技術固有名詞はそのまま（`go`, `docker`, `mysql`）
- 意味が分かる長さを優先（3〜6 単語が目安）
- 重複禁止（スクリプトで自動チェック）

---

## Phase 4: 日本語記事の Markdown 移行

### コンテンツ構造（移行後）

```
content/
  ja/
    posts/
      go-http-server.md
      docker-basics.md
      ...（700+ 件）
  en/
    posts/
      go-http-server.md        （英語版。高優先記事から順次作成）
      ...
```

### カテゴリー名英語化方針

gobel のカテゴリー名は日本語。gohan の taxonomy.yaml で管理する際に英語名を付与する。
移行記事の Front Matter に書くカテゴリー名は日本語のまま維持し、taxonomy.yaml で対応を定義する（英語化が決まり次第一括置換）。

現行カテゴリー一覧:

| 日本語名 | 英語候補 |
|---|---|
| アーキテクチャ | Architecture |
| アプリケーション | Application |
| アルゴリズムとデータ構造 | Algorithms & Data Structures |
| インフラストラクチャ | Infrastructure |
| OS | OS |
| 開発プロセス | Development Process |
| キャリア | Career |
| コンピューターアーキテクチャ | Computer Architecture |
| 数学 | Mathematics |
| ツール | Tools |
| テスト | Testing |
| データベース | Database |
| ネットワーク | Networking |
| パフォーマンス | Performance |
| ビジネス | Business |
| ポエム | Essay |
| マネジメント | Management |
| ヘルスケア | Healthcare |
| 運用 | Operations |
| 障害報告 | Incident Report |

---

## Phase 5: 画像アセット移行

### 現状

gobel の記事 `md_body` に埋め込まれた画像 URL は主に以下の外部 CDN を参照している:

- `https://qiita-image-store.s3.amazonaws.com/` （多数）
- `https://user-images.githubusercontent.com/` （一部）
- その他外部 URL

### リスク

- 外部 CDN の URL が将来的に失効するリスクがある（特に Qiita Image Store）
- サイトオーナーがコントロールできない

### 対応方針

**Phase 5A（移行時）**: 既存外部 URL のままリリース（工数優先）

**Phase 5B（移行後）**: バッチ処理で画像をダウンロードしリポジトリ内に収容

```
assets/images/posts/{post-id}/{filename}
```

Markdown 内の URL を相対パスに一括置換するスクリプトを別途用意する。

---

## Phase 6: テーマ・テンプレート開発

[DESIGN_DOC.md](./DESIGN_DOC.md) のセクション 6 を参照。

---

## Phase 7: 英語版記事の追加

### 優先度付け方針

1. Google Search Console でアクセス上位の記事から着手
2. AI（Claude）で日本語記事を英語に翻訳、レビュー後 `content/en/posts/` に配置
3. `translation_key` で日英ペアリング → gohan が `hreflang` を自動出力

### 作業単位

- 1 記事 = 1 PR（レビュー・履歴管理のため）
- 英語版の `title` / `description` は翻訳後に手動調整

---

## Phase 8: リダイレクトマップ

### 旧 URL → 新 URL マッピング

| 旧 URL パターン | 新 URL | 備考 |
|---|---|---|
| `/posts/{title}` | `/ja/posts/{slug}/` | 日本語記事はすべて `/ja/` 配下 |
| `/posts/categories/{name}` | `/categories/{name}/` | カテゴリーページ |
| `/posts/tags/{name}` | `/tags/{name}/` | タグページ |
| `/posts/search` | `/` | 検索は廃止 |
| `/categories` | `/` | 一覧は不要 |
| `/tags` | `/` | 一覧は不要 |
| `/profile` | `/about/` | |
| `/privacy_policy` | `/privacy-policy/` | |
| `/sitemap` | `/sitemap.xml` | |
| `/feed` | `/atom.xml` | |

### `_redirects` ファイル生成

移行スクリプト（Phase 2）が `_redirects` の大部分を自動生成する。

```
# _redirects（Cloudflare Pages 形式）
/posts/OSSをはじめてみた話        /ja/posts/getting-started-with-oss/   301
/posts/GolangのHTTPサーバーのコードリーディング  /ja/posts/go-http-server-code-reading/  301
/posts/search                      /                                      301
/profile                           /about/                                301
/privacy_policy                    /privacy-policy/                       301
```

**Cloudflare Pages Free プランで 2,000 ルールまで無料**（700+ 記事でも余裕）。

### `_redirects` の永続管理

- リポジトリに `_redirects` をコミットして永続管理
- 外部被リンク・ブックマークへの対応として廃止しない
- DNS を Cloudflare に移管し管理を一元化

---

## Phase 9: CI/CD パイプライン

```yaml
# .github/workflows/deploy.yml（概要）
on:
  push:
    branches: [main]

jobs:
  deploy:
    steps:
      - uses: actions/checkout@v4
      - name: Build gohan binary
        run: go install github.com/bmf-san/gohan/cmd/gohan@latest
      - name: Build site
        run: gohan build
      - name: Deploy to Cloudflare Pages
        uses: cloudflare/pages-action@v1
        with:
          projectName: bmf-tech
          directory: public
        env:
          CLOUDFLARE_API_TOKEN: ${{ secrets.CF_API_TOKEN }}
```

---

## Phase 10: Cloudflare Pages 本番デプロイ

1. Cloudflare Pages プロジェクト作成（`bmf-tech`）
2. GitHub リポジトリと連携（`main` ブランチへの push で自動ビルド）
3. Cloudflare Pages のプレビュー URL（`*.pages.dev`）で全ページ動作確認
   - リダイレクトが `_redirects` 通りに機能するか確認
   - sitemap.xml・atom.xml の内容確認
4. Google Search Console に新 sitemap を送信

---

## Phase 11: DNS 移管（ConoHa → Cloudflare）

### 現状

| 項目 | 現状 |
|---|---|
| サーバー | ConoHa VPS（Docker 運用） |
| DNS 管理 | ConoHa DNS（VPS に紐づき） |
| 稼働サービス | gobel API (Go)、MySQL、Redis、Vue.js フロントエンド、gondola (リバースプロキシ) |
| ドメイン | bmf-tech.com |

### 目標状態

| 項目 | 移行後 |
|---|---|
| サーバー | なし（VPS 廃止） |
| DNS 管理 | Cloudflare DNS |
| コンテンツ配信 | Cloudflare Pages |
| ドメイン | bmf-tech.com（Cloudflare DNS 経由） |

### 移管手順

#### Step 1: Cloudflare に bmf-tech.com を追加

1. Cloudflare ダッシュボード → **Add a Site** → `bmf-tech.com` を入力
2. Free プランを選択
3. Cloudflare が現在の DNS レコードを自動スキャン・インポートする
4. インポートされたレコードを確認・必要に応じて補完
   - 現在の `A` レコード（ConoHa VPS の IP）が取り込まれていることを確認

#### Step 2: Cloudflare Pages カスタムドメインを設定

1. Cloudflare Pages プロジェクト → **Custom domains** → `bmf-tech.com` を追加
2. Cloudflare が自動で `CNAME` レコード（`bmf-tech.com → {project}.pages.dev`）を追加
3. `www.bmf-tech.com` も追加し、`bmf-tech.com` へリダイレクト設定

#### Step 3: ネームサーバーを Cloudflare に切り替え

1. ConoHa コントロールパネル → ドメイン設定 → ネームサーバーを以下に変更:
   ```
   {assigned}.ns.cloudflare.com
   {assigned}.ns.cloudflare.com
   ```
2. DNS 伝播完了まで待機（最大 48 時間、通常数時間）
3. Cloudflare ダッシュボードでステータスが **Active** になったことを確認

#### Step 4: 切り替え後の確認

```bash
# DNS 伝播確認
dig bmf-tech.com
nslookup bmf-tech.com

# HTTPS 証明書確認（Cloudflare Universal SSL が自動発行）
curl -I https://bmf-tech.com

# リダイレクト確認（旧 URL が 301 で転送されるか）
curl -I "https://bmf-tech.com/posts/GoでHTTPサーバーを書く"
```

### SSL 証明書

- Cloudflare Universal SSL が自動発行・自動更新される
- ConoHa VPS 側で使用していた Let's Encrypt 証明書は不要になる
- 移行後は VPS の証明書管理から解放される

### ロールバック方針

- ネームサーバー切り替え後 **1 ヶ月間は VPS を停止しない**
- 問題発生時は ConoHa のネームサーバーに戻すことで旧環境にフォールバック可能
- DNS TTL が低い（Cloudflare デフォルト: Auto）ため切り替えは比較的高速

---

## Phase 12: ConoHa VPS 停止

### 前提条件

- [ ] DNS 切り替えから 1 ヶ月経過
- [ ] Cloudflare Pages への全トラフィック転送を確認（VPS へのアクセスがゼロ）
- [ ] MySQL データのバックアップを手元に保存済み
- [ ] gobel API / Vue.js フロントエンドのソースコードが GitHub に保管済み

### 停止手順

1. ConoHa VPS の Docker サービスを停止
   ```bash
   docker compose -f docker-compose-local.yml down
   ```
2. MySQL の最終バックアップを取得・手元に保存
3. ConoHa コントロールパネルから VPS を削除
4. ConoHa アカウントの自動更新を停止

### コスト削減

| 項目 | 現状 | 移行後 |
|---|---|---|
| ConoHa VPS | 月額 ¥1,000〜2,000（スペックによる） | ¥0 |
| Cloudflare Pages | — | ¥0（Free プラン） |
| Cloudflare DNS | — | ¥0 |
| SSL 証明書更新作業 | 手動 or 自動化が必要 | 不要（Cloudflare 管理） |

---

## 未決定事項

| 項目 | 状況 |
|---|---|
| サーバーサイド移行（gobel API・VPS） | Phase 11・12 に計画策定済み |
| Pagefind 等サイト内検索 | Phase 6 テンプレート開発時に評価 |
| Google AdSense / Analytics の設定 | ID を config.yaml で管理、テンプレートに組み込み |
| カテゴリー名の英語化タイミング | Phase 4 完了後に一括置換 |
| 画像アセットの収容先（Cloudflare Images vs リポジトリ） | Phase 5B 実施時に決定 |
