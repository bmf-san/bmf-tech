---
title: "master直プッシュの際にプロンプトで確認するようにする"
slug: "master"
date: 2019-05-08
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Git"
  - "shellscript"
draft: false
---

# 概要
gitでmasterブランチへの直pushを未然に防ぐためのセーフーティネットの作り方。
github上でmasterブランチへのpushを禁止すれば良いのだが、DevOpsの都合上でgithubの設定では問題があったのでhooksを使う方向で設定した。

# 導入
グローバルに設定したいので`~/.git_template/hooks`配下にpre-pushファイルを作成する。
`.git_template`ディレクトリが存在しない場合は作成する。

なお、グローバルに設定しても既存のリポジトリには反映されないので、既存リポジトリに反映したい場合は既存リポジトリの`./git/hooks`配下に`pre-push`を用意し、そちらに同じソースを記述する必要がある。

`pre-push`に記述する内容は [gitのpre-push hookでmasterブランチにpushする際にプロンプトで確認するようにする](https://qiita.com/mkiken/items/af5c40ce0d0c6d3530c7)を参照。

answer部分はyesだけにしておくほうがより安全かもしれない。

新規に作成した`pre-push`には実行権限を与えておく
`chmod +x pre-push`

以上でセットアップ完了。

# 参考　
- [gitのpre-push hookでmasterブランチにpushする際にプロンプトで確認するようにする](https://qiita.com/mkiken/items/af5c40ce0d0c6d3530c7)
