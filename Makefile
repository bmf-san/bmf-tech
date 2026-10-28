.PHONY: help install-gohan install-e2e install-lint build serve clean test-e2e test-e2e-ui new-ja new-en new-ja-roundup new-en-roundup lint-content lint-content-diff check-parity check-picks check-content

GOHAN_VERSION ?= v1.7.0

TITLE   ?= untitled
SLUG    ?= untitled

help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

install-gohan: ## gohanをインストール (GOHAN_VERSION で固定)
	GOTOOLCHAIN=auto go install github.com/bmf-san/gohan/cmd/gohan@$(GOHAN_VERSION)

install-e2e: ## Playwright依存をインストール
	cd e2e && npm ci && npx playwright install --with-deps chromium

install-lint: ## textlint 依存をインストール
	npm ci

lint-content: ## 全記事を textlint でチェック (JA + EN)
	npx textlint --config .textlintrc-ja.json "content/ja/posts/*.md"
	npx textlint --config .textlintrc-en.json "content/en/posts/*.md"

check-parity: ## 日英記事の translation_key 対応をチェック
	bash scripts/check-translation-parity.sh

check-picks: ## ピックアップ記事の listing_slugs / picks_sections / 日英の一致をチェック
	bash scripts/check-picks.sh

check-content: ## gohan check で重複スラッグ・必須フロントマター・孤立 translation_key を検証
	GOTOOLCHAIN=auto gohan check

lint-content-diff: ## origin/main との差分で追加・変更した行だけ textlint でチェック（新規記事は全体）
	python3 scripts/lint-changed-lines.py

build: ## サイトをビルド
	GOTOOLCHAIN=auto gohan build

serve: ## ローカルサーバーを起動 (http://localhost:1313)
	@pkill -f "gohan serve" 2>/dev/null; sleep 0.3; true
	GOTOOLCHAIN=auto gohan serve

clean: ## ビルド出力を削除
	rm -rf public/*

test-e2e: ## E2Eテストを実行 (make build を事前に実行すること)
	cd e2e && npx playwright test

test-e2e-ui: ## E2EテストをPlaywright UI モードで実行 (make build を事前に実行すること)
	cd e2e && npx playwright test --ui

new-ja: ## 日本語記事を作成  例: make new-ja TITLE="タイトル" SLUG=slug
	GOTOOLCHAIN=auto gohan new --type=ja/posts --archetype=ja-post --title="$(TITLE)" $(SLUG)

new-en: ## 英語記事を作成  例: make new-en TITLE="Title" SLUG=slug
	GOTOOLCHAIN=auto gohan new --type=en/posts --archetype=en-post --title="$(TITLE)" $(SLUG)

new-ja-roundup: ## まとめ記事(日本語)を作成  例: make new-ja-roundup TITLE="◯◯のおすすめ本" SLUG=slug
	GOTOOLCHAIN=auto gohan new --type=ja/posts --archetype=ja-roundup --title="$(TITLE)" $(SLUG)

new-en-roundup: ## まとめ記事(英語)を作成  例: make new-en-roundup TITLE="Best books for ..." SLUG=slug
	GOTOOLCHAIN=auto gohan new --type=en/posts --archetype=en-roundup --title="$(TITLE)" $(SLUG)


