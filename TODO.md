# TODO

詳細は [docs/migration.md](docs/migration.md) を参照。

## gobel → gohan 移行

- [x] Phase 0: 設計ドキュメント・移行計画策定
- [x] Phase 1: bmf-tech リポジトリ初期化・gohan プロジェクトセットアップ（PR #2）
- [x] Phase 2: テーマ・テンプレート開発（PR #3）
- [x] Phase 3: 英語スラッグ変換マップ作成（584 記事）（PR #4）
- [x] Phase 4: コンテンツデータ移行スクリプト作成（PR #5）
- [x] Phase 5: 日本語記事の Markdown 移行（`content/ja/`）（PR #6）
- [x] Phase 6: 画像アセット移行（PR #7）
- [x] Phase 7: リダイレクトマップ作成・検証（PR #8）
- [x] Phase 8: 英語版記事スケルトン + canonical/hreflang SEO 対応（PR #9）
- [x] Phase 9: CI/CD パイプライン構築 + about/privacy-policy ページ + README/Makefile（PR #10）
- [ ] Phase 10: Cloudflare Pages 本番デプロイ
  - [ ] GitHub Secrets 設定（`CLOUDFLARE_API_TOKEN`, `CLOUDFLARE_ACCOUNT_ID`）
  - [ ] Cloudflare Pages プロジェクト `bmf-tech` 作成
  - [ ] PR #2〜#10 を順番にマージして main へ
  - [ ] 初回デプロイ確認
- [ ] Phase 11: DNS 移管（ConoHa → Cloudflare）
- [ ] Phase 12: ConoHa VPS 停止

## gohan 対応待ち（移行前提）

- [x] feed.go の i18n 対応（PR #60）
- [x] `draft: true` フィルタリング未実装 + `build.exclude_files` が機能しない（PR #59）
- [x] `OutputGenerator` インターフェイス整理（PR #63）
- [x] pages ルーティング改善・`computeOutputPath()` の結果が活用されていない（PR #61）
- [x] 記事一覧・タグ・カテゴリーページのソート未実装（PR #61）
- [x] `FrontMatter.Template` フィールドが未使用（PR #61）
- [x] 日付ゼロ記事が `archives/0001/01/` を生成（PR #61）
- [x] `ValidateArticleTaxonomies` がビルドから呼ばれない（PR #62）
- [x] アーカイブ URL 設計確認 → `/archives/{year}/{month}/` で確定
- [x] gohan リリース（[v0.1.0](https://github.com/bmf-san/gohan/releases/tag/v0.1.0)）
- [x] `cfg.Build.ContentDir` が相対パスのまま processor に渡され i18n ロケール検出が失敗するバグ（[PR #66](https://github.com/bmf-san/gohan/pull/66)）→ マージ済み・[v0.1.1](https://github.com/bmf-san/gohan/releases/tag/v0.1.1) リリース済み・`make install` を `go install @latest` に変更済み
- [x] 出力パス重複チェック追加（[PR #67](https://github.com/bmf-san/gohan/pull/67)）→ マージ済み・[v0.1.2](https://github.com/bmf-san/gohan/releases/tag/v0.1.2) リリース済み
- [x] `go.mod` の Go バージョン要件を 1.24.0 に修正（1.25.3 が誤って設定されていた）→ [v0.1.3](https://github.com/bmf-san/gohan/releases/tag/v0.1.3) リリース済み
- [x] go directive を 1.26.0 に更新 → [v0.1.4](https://github.com/bmf-san/gohan/releases/tag/v0.1.4) リリース済み・`go install @v0.1.4` にピン留め済み