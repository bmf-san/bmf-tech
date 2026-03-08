---
title: How to Split Commits in Git
slug: split-commits-with-git
date: 2021-10-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Git
translation_key: split-commits-with-git
---

A note on how to split commits in Git. Sometimes you want to do this.

```sh
// Specify the part to be split. Edit the target commit.
git rebase -i HEAD~5
```

```sh
// The target commit is unstaged
git reset HEAD~
```

```sh
// Add & commit with the desired granularity
git add ~
git commit ~
```

```sh
git rebase --continue
```