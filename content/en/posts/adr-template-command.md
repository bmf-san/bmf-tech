---
title: Command to Generate ADR Template Files
description: An in-depth look at Command to Generate ADR Template Files, covering key concepts and practical insights.
slug: adr-template-command
date: 2023-11-10T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Architecture Decision Record
translation_key: adr-template-command
---

I created a Make command that simply generates a template file for ADR, so I'm jotting it down.

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

## Result

endef

export ADR_TEMPLATE

.PHONY: adr
adr: ## Create a new ADR. ex. make adr title=Title
	@if [ -z "$(title)" ]; then \
		echo "Title is not set. 'ex. make adr title=Title'"; \
		exit 1; \
	fi
	adr_number=$$(ls adr/ADR*-*.md 2>/dev/null | awk -F- '/ADR[0-9]+-/{match($$0, /[0-9]+/); print substr($$0, RSTART, RLENGTH)}' | sort -n | tail -n 1); \
	adr_name=ADR$$(($$adr_number + 1))-$(title); \
	echo "$$ADR_TEMPLATE" | sed -e "s/\TITLE/$$adr_name/g;" > adr/$$adr_name.md; \
	echo "New ADR created: adr/$$adr_name.md"
```

The file naming convention for ADR is ADR<incrementable number>-Title, so this command generates the ADR template file with an appropriate file name by looking at the files in the adr directory.

If there is a file named ADR1-foo.md, it will increment the number to create ADR2-bar.md, and so on.

# Others
If ADRs are managed under git, it may become difficult to grasp the status of each ADR, so some measures may be necessary.

It might be necessary to prepare a command to list by status, separate directories by status, or include the status in the file name.