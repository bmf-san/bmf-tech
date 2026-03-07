---
title: How to Split Commits in Git
slug: git-split-commit
date: 2021-06-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Git
translation_key: git-split-commit
---

# Overview
A note on the steps to split a commit.

# Steps
```sh
# Specify where to split and rebase. Change the target commit to edit.
git rebase -i HEAD~5

# Unstage the changes
git rebase HEAD~

# Recommit the unstaged changes at the desired granularity.
git add & git commit

# Once the commit is complete, continue the rebase
git rebase --continue

# Check the log
git log

# Force push
git push -f origin HEAD
```

This should allow you to split the commit.