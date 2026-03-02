# TODO

詳細は [docs/migration.md](docs/migration.md) を参照。

## gobel → gohan 移行

- [x] Phase 0: 設計ドキュメント・移行計画策定
- [ ] Phase 1: bmf-tech リポジトリ初期化・gohan プロジェクトセットアップ
- [ ] Phase 2: コンテンツデータ移行スクリプト作成
- [ ] Phase 3: 英語スラッグ変換マップ作成（700+ 記事）
- [ ] Phase 4: 日本語記事の Markdown 移行（`content/ja/`）
- [ ] Phase 5: 画像アセット移行
- [ ] Phase 6: テーマ・テンプレート開発
- [ ] Phase 7: 英語版記事の作成（高優先記事から）
- [ ] Phase 8: リダイレクトマップ作成・検証
- [ ] Phase 9: CI/CD パイプライン構築
- [ ] Phase 10: Cloudflare Pages 本番デプロイ
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