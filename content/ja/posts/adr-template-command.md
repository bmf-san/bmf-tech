---
title: ADRのテンプレートファイルを生成するコマンド
description: ADRのテンプレートファイルを生成するコマンド
slug: adr-template-command
date: 2023-11-10T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Architecture Decision Record
translation_key: adr-template-command
---


ADRのテンプレートファイルを生成するだけのMakeコマンドを作成したのでメモっておく。

# コマンド
```sh
.PHONY:help
help: ## Print help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

define ADR_TEMPLATE
# TITLE
## 背景

## 決定

## ステータス
提案済み

<!--
提案済み/承認済み/棄却...
-->

## 結果

endef

export ADR_TEMPLATE

.PHONY: adr
adr: ## Create a new ADR. ex. make adr title=タイトル
	@if [ -z "$(title)" ]; then \
		echo "タイトルが設定されていません。 'ex. make adr title=タイトル'"; \
		exit 1; \
	fi
	adr_number=$$(ls adr/ADR*-*.md 2>/dev/null | awk -F- '/ADR[0-9]+-/{match($$0, /[0-9]+/); print substr($$0, RSTART, RLENGTH)}' | sort -n | tail -n 1); \
	adr_name=ADR$$(($$adr_number + 1))-$(title); \
	echo "$$ADR_TEMPLATE" | sed -e "s/\TITLE/$$adr_name/g;" > adr/$$adr_name.md; \
	echo "New ADR created: adr/$$adr_name.md"
```

ADRのファイル命名規則をADR<インクリメント可能な数値>-タイトルとしているので、adrディレクトリ配下のファイルを見て適切なファイル名でADRのテンプレートファイルを生成するコマンドになっている。

ADR1-foo.mdというファイルがあれば、ADR2-bar.mdといった感じで数値をインクリメントしてくれる。

# その他
ADRをgit管理下にすると、ADRごとのステータスを把握しづらくなるので何かしらの対応が必要かもしれない。

ステータスごとにリストアップするコマンドを用意する、ステータスごとにディレクトリを分ける、ステータスをファイル名に含めるなど工夫が必要。
