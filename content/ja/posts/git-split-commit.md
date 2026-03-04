---
title: "Gitでcommitを分割する方法"
slug: "git-split-commit"
date: 2021-06-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Git"
draft: false
---

# 概要
commitを分割したいときの手順についてメモ。

# 手順
```sh
# 分割したいところを指定してrebase。対象のcommitをeditにする。
git rebase -i HEAD~5

# unstageする
git rebase HEAD~

# unstageしたものを分割したい粒度で再commitする。
git add & git commit

# commitが完了したらrebase --continue
git rebase --continue

# log確認
git log

# force push
git push -f origin HEAD
```

これで分割できるはず。
