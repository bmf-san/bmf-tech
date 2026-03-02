# TODO

## gobel → gohan 移行

- [x] Phase 0: 設計ドキュメント・移行計画策定
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

- [ ] feed.go の i18n 対応（`atom.xml` / `feed.xml` の URL が `/posts/{slug}/` にハードコードされており `/ja/posts/{slug}/` が正しく出力されない）
- [ ] `draft: true` フィルタリング未実装（`draft: true` の記事もビルドに含まれる。移行後の下書き管理に影響）
- [ ] `OutputGenerator` インターフェイスの `GenerateSitemap` / `GenerateFeed` メソッドが死んでいる（i18n 非対応の旧実装が残存。インターフェイス整理が必要）
- [ ] pages ルーティング改善（`buildJobs()` が全記事を `posts/{slug}/` に書き出す。`computeOutputPath()` の結果が `generator/html.go` で無視されるため `content/pages/` の記事は `/pages/about/` の URL を持つが実ファイルは `/posts/about/` に置かれ、URL とファイルパスが乖離する）
- [ ] 記事一覧・タグ・カテゴリーページのソート未実装（index / tag / category / archive 各ページに渡す記事スライスが日付順にソートされておらず、ファイルシステムのウォーク順（辞書順）で出力される）
- [ ] `FrontMatter.Template` フィールドが未使用（front matter で `template: custom.html` を指定しても `article.html` が常に使用される）
- [ ] 日付ゼロ記事が `archives/0001/01/` を生成（`date` が未設定の記事は Go の zero time 扱いとなり `public/archives/0001/01/index.html` が生成される）
- [ ] `ValidateArticleTaxonomies` がビルドから呼ばれない（`processor/taxonomy.go` に実装済みだが `cmd/gohan/build.go` から呼び出されておらず、タクソノミーの整合性チェックがサイレントにスキップされる）
- [ ] `build.exclude_files` 設定が機能しない（`model.BuildConfig.ExcludeFiles` フィールドは定義済みだが `parser/frontmatter.go` の `ParseAll` が参照せず、除外指定が無効）
- [ ] アーカイブ URL 設計確認（現状 `/archives/{year}/{month}/`。年単位 `/archive/{year}/` にしたい場合は gohan 側改修が必要）
- [ ] gohan リリース（`go install ...@latest` を CI で使えるようにする）