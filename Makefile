.PHONY: help install install-e2e build serve clean test-e2e test-e2e-ui new new-ja new-en translate translate-gemini translate-dry-run

TITLE   ?= untitled
SLUG    ?= untitled

help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## 依存ツールをインストール (gohan)
	GOTOOLCHAIN=auto go install github.com/bmf-san/gohan/cmd/gohan@latest

install-e2e: ## Playwright依存をインストール
	cd e2e && npm ci && npx playwright install chromium

build: ## サイトをビルド
	GOTOOLCHAIN=auto gohan build

serve: ## ローカルサーバーを起動 (http://localhost:1313)
	@pkill -f "gohan serve" 2>/dev/null; sleep 0.3; true
	GOTOOLCHAIN=auto gohan serve

clean: ## ビルド出力を削除
	rm -rf public/*

test-e2e: ## E2Eテストを実行 (事前にビルドして http-server で配信)
	gohan build && cd e2e && npx playwright test

test-e2e-ui: ## E2EテストをPlaywright UI モードで実行
	gohan build && cd e2e && npx playwright test --ui

new: new-ja new-en ## 新しい記事を作成 (例: make new TITLE="タイトル" SLUG=slug)

new-ja: ## 日本語記事を作成  例: make new-ja TITLE="タイトル" SLUG=slug
	@mkdir -p content/ja/posts
	@if [ -f "content/ja/posts/$(SLUG).md" ]; then \
		echo "error: content/ja/posts/$(SLUG).md already exists"; exit 1; \
	fi
	@printf -- '---\ntitle: "$(TITLE)"\nslug: $(SLUG)\ndate: %s\nauthor: bmf-san\ncategories:\n  - \ntags:\n  - \ndescription: ""\ntranslation_key: $(SLUG)\ndraft: true\n---\n' \
		$$(date +%Y-%m-%d) > content/ja/posts/$(SLUG).md
	@echo "created: content/ja/posts/$(SLUG).md"

new-en: ## 英語記事を作成  例: make new-en TITLE="Title" SLUG=slug
	@mkdir -p content/en/posts
	@if [ -f "content/en/posts/$(SLUG).md" ]; then \
		echo "error: content/en/posts/$(SLUG).md already exists"; exit 1; \
	fi
	@printf -- '---\ntitle: "$(TITLE)"\nslug: $(SLUG)\ndate: %s\nauthor: bmf-san\ncategories:\n  - \ntags:\n  - \ndescription: ""\ntranslation_key: $(SLUG)\ndraft: true\n---\n' \
		$$(date +%Y-%m-%d) > content/en/posts/$(SLUG).md
	@echo "created: content/en/posts/$(SLUG).md"

translate: ## JA記事を一括英語翻訳 (GITHUB_TOKEN or OPENAI_API_KEY が必要)
	cd tools/translate && GOTOOLCHAIN=auto go run . -delay 1000

translate-gemini: ## JA記事を一括英語翻訳 (GOOGLE_API_KEY 使用 / Gemini 2.0 Flash, 15RPM対応)
	cd tools/translate && GOTOOLCHAIN=auto go run . -delay 5000

translate-dry-run: ## 翻訳対象の確認 (API不要)
	cd tools/translate && GOTOOLCHAIN=auto go run . -dry-run
