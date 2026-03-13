---
title: Gitでcommitを分割する方法
description: "Gitのrebase -iを使ってcommitを分割する手順を解説し、edit指定による粒度調整とforce pushまでの流れを示す。"
slug: git-split-commit
date: 2021-06-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Git
translation_key: git-split-commit
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
