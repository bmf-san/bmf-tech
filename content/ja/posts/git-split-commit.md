---
title: "Gitでcommitを分割する方法｜rebase -iで履歴整理"
description: "1つにまとまってしまったコミットを分割したいときの手順を、Gitのインタラクティブなrebase（rebase -i）を使って解説します。対象コミットをeditに指定して粒度を調整し、force pushで反映するまでの一連の流れをメモとして残します。"
slug: git-split-commit
date: 2021-06-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Git
translation_key: git-split-commit
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
