---
title: Command to Generate ADR Template Files
slug: adr-template-command
date: 2023-11-10T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Architecture Decision Record
description: A Make command to generate ADR template files.
translation_key: adr-template-command
---

I created a Make command that generates ADR template files, so here's a quick note.

# Command
```sh
.PHONY:help
help: ## Print help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

define ADR_TEMPLATE
# TITLE
## Background

## Decision

## Status
Proposed

<!--
Proposed/Approved/Rejected...
-->

## Outcome

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

The naming convention for ADR files is `ADR<incremental-number>-title`, so this command looks at the files under the `adr` directory and generates a new ADR template file with an appropriate file name.

For example, if there is a file named `ADR1-foo.md`, it will increment the number and create a file like `ADR2-bar.md`.

# Other Notes
If you manage ADRs under git, it might become difficult to track the status of each ADR. Some measures may be necessary, such as:

- Preparing a command to list ADRs by status
- Organizing ADRs into directories by status
- Including the status in the file name

Consider implementing one of these approaches to improve manageability.