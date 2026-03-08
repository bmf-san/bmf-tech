---
title: How to Split a Commit in Git
slug: split-commits-with-git
date: 2021-10-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Git
description: Notes on how to split a commit in Git. Occasionally needed.
translation_key: split-commits-with-git
---



Notes on how to split a commit in Git. Occasionally needed.

```sh
// Specify the point to split. Edit the target commit.
git rebase -i HEAD~5
```

```sh
// The target commit is unstaged
git reset HEAD~
```

```sh
// Add & commit at any granularity
git add ~
git commit ~
```

```sh
git rebase --continue
```
