# 運用ガイド

サイト運用に必要な外部サービス連携と CI/CD の概要をまとめる。

## デプロイ (Cloudflare Pages)

### フロー

1. `main` ブランチへ push (もしくは毎日 00:00 UTC の schedule 実行)
2. `.github/workflows/deploy.yml` が走る
3. `gohan build --stats --explain` で `public/` を生成 (フェーズ別タイミングとリビルド理由をログ出力)
4. `cloudflare/wrangler-action@v3` が `public/` を Cloudflare Pages プロジェクト `bmf-tech` にアップロード

### 予約投稿 (gohan v1.3.0+)

`date` を未来日時に設定した記事は、その日付に到達するまで build から除外される (gohan v1.3.0 のデフォルト挙動)。`main` にマージしても即時公開されず、毎日 00:00 UTC の scheduled deploy で `date <= now` を満たした時点で自動的に公開される。

プレビュー目的で未来日付の記事を含めてビルドしたい場合のみ、ローカルで `gohan build --future` を使う。

### 必要な GitHub Secrets

| 名前 | 用途 |
| --- | --- |
| `CLOUDFLARE_API_TOKEN` | Pages:Edit 権限を持つ API Token |
| `CLOUDFLARE_ACCOUNT_ID` | Cloudflare アカウント ID |

### ドメイン / リダイレクト

- 本番ドメインは `bmf-tech.com` (Cloudflare 側で Pages プロジェクトに割当)。
- URL 変更時のリダイレクトは `bulk-redirects.txt` に追記し、Cloudflare ダッシュボードの Bulk Redirects 機能で読み込む。Git には記録のみで自動反映はされない点に注意。

## dev.to 自動投稿

### フロー

1. `content/en/posts/**.md` を含む push が `main` に入ると `devto-publish.yml` が起動
2. `tools/devto` をビルドし、直近の commit で追加された記事を検出
3. dev.to API で `published: false` として投稿
4. `tools/devto/posted.json` を更新して bot が commit back

### 必要な GitHub Secrets

| 名前 | 用途 |
| --- | --- |
| `DEV_TO_API_KEY` | dev.to アカウントの API Key (User Settings → Extensions → DEV Community API Keys) |

### モード

`workflow_dispatch` で手動起動できる。

- `diff` (default): 直近 1 commit で追加された記事のみ
- `all`: `posted.json` に無い記事を全て

### 除外ルール (実装)

以下は dev.to に投稿しない:

- `draft: true` のフロントマター
- `categories` に `poem` (case insensitive) または `ポエム` を含むもの
- `tags` に `book review` (case insensitive) または `書評` を含むもの

### トラブルシュート

- API Key 不正: job の "Check API key" ステップで検出される
- 投稿失敗: Actions のログを確認。`posted.json` は書き戻されないので、原因修正後に `workflow_dispatch` 再実行可能

## コンテンツ lint

- `make lint-content`: 全記事を textlint (JA + EN)
- `make lint-content-diff`: PR 差分の本文修正分のみ
- `npm run spell:en`: 英語記事のスペルチェック (cspell)
- `make check-parity`: 日英記事の `translation_key` 対応チェック

## テスト

- `make test-e2e`: Playwright E2E (事前に `make build` が必要)
- `cd tools/devto && go test ./...`: dev.to 投稿ツールの unit test

## CI 一覧

| Workflow | Trigger | 内容 |
| --- | --- | --- |
| `deploy.yml` | push main | Cloudflare Pages デプロイ |
| `e2e.yml` | push main / PR | ESLint + Prettier + Playwright |
| `lint.yml` | PR (content/) | textlint + cspell |
| `linkcheck.yml` | PR / cron (Mon 00:00 UTC) | lychee でリンク切れ検出 |
| `parity.yml` | PR (content/) | 日英 translation_key 対応 |
| `lighthouse.yml` | PR | Lighthouse CI |
| `devto-publish.yml` | push main (`content/en/posts/**`) / 手動 | dev.to 投稿 |
| `tools-tests.yml` | push main / PR (tools/devto/) | tools の unit test |

## ローカル開発

### サーバー

`make serve` で `http://localhost:1313` に gohan dev server を起動する。ファイル変更を検知して自動リビルド + ブラウザリロードする。

#### CSS-only HMR (gohan v1.3.0+)

テーマや `assets/` 配下の `.css` ファイルのみを変更した場合、フルリロードではなく `<link rel="stylesheet">` の `?_gohan=<ts>` クエリ更新によるホットスワップが行われる。スクロール位置や開発中の DOM 状態を保ったままスタイル調整できるため、テーマ調整時のフィードバックループが短い。

HTML テンプレートや Markdown を変更した場合は従来通りフルリロード。

### コンテンツ整合性チェック

`make check-content` で gohan の組み込みリンタ (重複スラッグ・必須フロントマター・孤立 `translation_key`) を実行できる。CI でも `lint.yml` 経由で PR ごとに実行される。
