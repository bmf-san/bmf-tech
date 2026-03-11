TITLE   ?= untitled
SLUG    ?= untitled

.PHONY: help
help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## 依存ツールをインストール (gohan)
	GOTOOLCHAIN=auto go install github.com/bmf-san/gohan/cmd/gohan@latest

.PHONY: install-e2e
install-e2e: ## Playwright依存をインストール
	cd e2e && npm ci && npx playwright install chromium

.PHONY: build
build: ## サイトをビルド
	GOTOOLCHAIN=auto gohan build

.PHONY: serve
serve: ## ローカルサーバーを起動 (http://localhost:1313)
	@pkill -f "gohan serve" 2>/dev/null; sleep 0.3; true
	GOTOOLCHAIN=auto gohan serve

.PHONY: clean
clean: ## ビルド出力を削除
	rm -rf public/*

.PHONY: test-e2e
test-e2e: ## E2Eテストを実行 (事前にビルドして http-server で配信)
	gohan build && cd e2e && npx playwright test

.PHONY: test-e2e-ui
test-e2e-ui: ## E2EテストをPlaywright UI モードで実行
	gohan build && cd e2e && npx playwright test --ui

.PHONY: new
new: new-ja new-en ## 新しい記事を作成 (例: make new TITLE="タイトル" SLUG=slug)

.PHONY: new-ja
new-ja: ## 日本語記事を作成  例: make new-ja TITLE="タイトル" SLUG=slug
	@mkdir -p content/ja/posts
	@if [ -f "content/ja/posts/$(SLUG).md" ]; then \
		echo "error: content/ja/posts/$(SLUG).md already exists"; exit 1; \
	fi
	@printf -- '---\ntitle: "$(TITLE)"\nslug: $(SLUG)\ndate: %s\nauthor: bmf-san\ncategories:\n  - \ntags:\n  - \ndescription: ""\ntranslation_key: $(SLUG)\ndraft: true\n---\n' \
		$$(date +%Y-%m-%d) > content/ja/posts/$(SLUG).md
	@echo "created: content/ja/posts/$(SLUG).md"

.PHONY: new-en
new-en: ## 英語記事を作成  例: make new-en TITLE="Title" SLUG=slug
	@mkdir -p content/en/posts
	@if [ -f "content/en/posts/$(SLUG).md" ]; then \
		echo "error: content/en/posts/$(SLUG).md already exists"; exit 1; \
	fi
	@printf -- '---\ntitle: "$(TITLE)"\nslug: $(SLUG)\ndate: %s\nauthor: bmf-san\ncategories:\n  - \ntags:\n  - \ndescription: ""\ntranslation_key: $(SLUG)\ndraft: true\n---\n' \
		$$(date +%Y-%m-%d) > content/en/posts/$(SLUG).md
	@echo "created: content/en/posts/$(SLUG).md"

.PHONY: translate
translate: ## JA記事を一括英語翻訳 (GITHUB_TOKEN or OPENAI_API_KEY が必要)
	cd tools/translate && GOTOOLCHAIN=auto go run . -delay 1000

.PHONY: translate-gemini
translate-gemini: ## JA記事を一括英語翻訳 (GOOGLE_API_KEY 使用 / Gemini 2.0 Flash, 15RPM対応)
	cd tools/translate && GOTOOLCHAIN=auto go run . -delay 5000

.PHONY: translate-dry-run
translate-dry-run: ## 翻訳対象の確認 (API不要)
	cd tools/translate && GOTOOLCHAIN=auto go run . -dry-run
