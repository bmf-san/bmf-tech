.PHONY: help install install-e2e build serve clean test-e2e new-ja new-en translate translate-dry-run

TITLE   ?= untitled
SLUG    ?= untitled

help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## 依存ツールをインストール (gohan)
	GOTOOLCHAIN=auto go install github.com/bmf-san/gohan/cmd/gohan@v0.1.8

install-e2e: ## Playwright依存をインストール
	cd e2e && npm ci && npx playwright install chromium

build: ## サイトをビルド
	GOTOOLCHAIN=auto gohan build

serve: build ## ローカルサーバーを起動 (http://localhost:1313) — public/ を静的配信
	cd public && npx http-server -p 1313 -s --cors -c-1

dev: ## ライブリロード付きローカルサーバー (http://localhost:1313)
	GOTOOLCHAIN=auto gohan serve

clean: ## ビルド出力を削除
	rm -rf public/*

test-e2e: ## E2Eテストを実行 (事前にビルドして http-server で配信)
	gohan build && cd e2e && npx playwright test

test-e2e-ui: ## E2EテストをPlaywright UI モードで実行
	gohan build && cd e2e && npx playwright test --ui

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

translate-dry-run: ## 翻訳対象の確認 (API不要)
	cd tools/translate && GOTOOLCHAIN=auto go run . -dry-run
