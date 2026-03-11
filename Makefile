.PHONY: help install-gohan install-e2e install-lint build serve clean copy-redirects test-e2e test-e2e-ui new-ja new-en translate translate-gemini translate-dry-run lint-content lint-content-diff

TITLE   ?= untitled
SLUG    ?= untitled

help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

install-gohan: ## gohanをインストール
	GOTOOLCHAIN=auto go install github.com/bmf-san/gohan/cmd/gohan@latest

install-e2e: ## Playwright依存をインストール
	cd e2e && npm ci && npx playwright install --with-deps chromium

install-lint: ## textlint 依存をインストール
	npm ci

lint-content: ## 全記事を textlint でチェック (JA + EN)
	npx textlint --config .textlintrc-ja.json "content/ja/posts/*.md"
	npx textlint --config .textlintrc-en.json "content/en/posts/*.md"

lint-content-diff: ## origin/main との差分ファイルのうち本文変更があるもののみ textlint でチェック
	@ALL_JA=$$(git diff --name-only --diff-filter=ACM origin/main...HEAD -- 'content/ja/posts/*.md'); \
	ALL_EN=$$(git diff --name-only --diff-filter=ACM origin/main...HEAD -- 'content/en/posts/*.md'); \
	JA_FILES=$$(for f in $$ALL_JA; do \
		if git diff origin/main...HEAD -- "$$f" | awk 'BEGIN{fm=0;cnt=0} /^[+-]{3} /{next} /^[+-]---$$/{cnt++;if(cnt==2)fm=0;next} cnt<2{next} /^[+-]/{found=1;exit} END{exit !found}'; then \
			echo "$$f"; \
		fi; \
	done); \
	EN_FILES=$$(for f in $$ALL_EN; do \
		if git diff origin/main...HEAD -- "$$f" | awk 'BEGIN{fm=0;cnt=0} /^[+-]{3} /{next} /^[+-]---$$/{cnt++;if(cnt==2)fm=0;next} cnt<2{next} /^[+-]/{found=1;exit} END{exit !found}'; then \
			echo "$$f"; \
		fi; \
	done); \
	if [ -n "$$JA_FILES" ]; then \
		echo "$$JA_FILES" | xargs npx textlint --config .textlintrc-ja.json; \
	else \
		echo "No changed JA post bodies to lint."; \
	fi; \
	if [ -n "$$EN_FILES" ]; then \
		echo "$$EN_FILES" | xargs npx textlint --config .textlintrc-en.json; \
	else \
		echo "No changed EN post bodies to lint."; \
	fi

build: ## サイトをビルド
	GOTOOLCHAIN=auto gohan build

serve: ## ローカルサーバーを起動 (http://localhost:1313)
	@pkill -f "gohan serve" 2>/dev/null; sleep 0.3; true
	GOTOOLCHAIN=auto gohan serve

clean: ## ビルド出力を削除
	rm -rf public/*

copy-redirects: ## _redirects を public/ にコピー
	cp _redirects public/

test-e2e: ## E2Eテストを実行 (make build を事前に実行すること)
	cd e2e && npx playwright test

test-e2e-ui: ## E2EテストをPlaywright UI モードで実行 (make build を事前に実行すること)
	cd e2e && npx playwright test --ui

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
