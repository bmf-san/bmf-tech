---
title: "Gitでコミットを分割する方法"
slug: "git"
date: 2021-10-05
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Git"
draft: false
---

Gitでコミットを分割する方法のメモ。たまにやりたくなる。　

```sh
// 分割したところを指定。対象commitをeditする。
git rebase -i HEAD~5
```

```sh
// 対象commitがunstageされる
git reset HEAD~
```

```sh
// 任意の粒度でadd&commit
git add ~
git commit ~
```

```sh
git rebase --continue
```
