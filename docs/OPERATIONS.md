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

## コンテンツ lint

- `make lint-content`: 全記事を textlint (JA + EN)
- `make lint-content-diff`: PR 差分の本文修正分のみ
- `npm run spell:en`: 英語記事のスペルチェック (cspell)
- `make check-parity`: 日英記事の `translation_key` 対応チェック

## テスト

- `make test-e2e`: Playwright E2E (事前に `make build` が必要)

## CI 一覧

| Workflow | Trigger | 内容 |
| --- | --- | --- |
| `deploy.yml` | push main | Cloudflare Pages デプロイ |
| `e2e.yml` | push main / PR | ESLint + Prettier + Playwright |
| `lint.yml` | PR (content/) | textlint + cspell |
| `linkcheck.yml` | PR / cron (Mon 00:00 UTC) | lychee でリンク切れ検出 |
| `parity.yml` | PR (content/) | 日英 translation_key 対応 + ピックアップ記事の整合性 |
| `lighthouse.yml` | PR | Lighthouse CI |
| `picks-review-reminder.yml` | cron (1/4/7/10 月 1 日 00:00 UTC) | ピックアップ記事の見直し Issue を作成 |

## ローカル開発

### サーバー

`make serve` で `http://localhost:1313` に gohan dev server を起動する。ファイル変更を検知して自動リビルド + ブラウザリロードする。

#### CSS-only HMR (gohan v1.3.0+)

テーマや `assets/` 配下の `.css` ファイルのみを変更した場合、フルリロードではなく `<link rel="stylesheet">` の `?_gohan=<ts>` クエリ更新によるホットスワップが行われる。スクロール位置や開発中の DOM 状態を保ったままスタイル調整できるため、テーマ調整時のフィードバックループが短い。

HTML テンプレートや Markdown を変更した場合は従来通りフルリロード。

### コンテンツ整合性チェック

`make check-content` で gohan の組み込みリンタ (重複スラッグ・必須フロントマター・孤立 `translation_key`) を実行できる。CI でも `lint.yml` 経由で PR ごとに実行される。

## ピックアップ記事の運用

`/picks/` と `/ja/picks/` は筆者が選んだ記事をテーマ別に並べるページ。

### 選定基準

設計判断やトレードオフを扱う記事を中心に、長く参照できるものを選ぶ。この基準はページ冒頭（`picks.md` の本文）にも書いている。

### 構成

- `content/{ja,en}/picks.md` の `listing_slugs`: 掲載する記事の一覧。gohan がこれを記事に解決する
- 同 `picks_groups`: テーマごとの見出し (`title`) と記事 (`slugs`)。テンプレート `themes/default/templates/picks.html` がこの順で表示する
- どのグループにも入っていない記事は「その他」に表示される
- `make check-picks` で、`listing_slugs` と `picks_groups` の一致、記事の存在、日英の一致を検証する（CI では `parity.yml`）

### 四半期ごとの見直し

`picks-review-reminder.yml` が 1/4/7/10 月に見直し用の Issue を作る。選ぶのは引き続き筆者自身で、データは見落としに気づくためだけに使う。

1. GA4 で直近 3 か月の閲覧数上位の記事を確認する
2. Search Console で直近 3 か月の流入上位の記事を確認する
3. 上位にあるのにピックアップにない記事を、選定基準に照らして追加するか判断する
4. 古くなった記事や、選定基準に合わなくなった記事を外すか判断する
5. 変更する場合は日英の `listing_slugs` と `picks_groups` を更新し、`make check-picks` を通す
