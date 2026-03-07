---
title: gitコマンドを楽にするシェルスクリプトをかいた
slug: create-shell-script-git-commands
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Git
  - shellscript
translation_key: create-shell-script-git-commands
---


# 概要
[.bashrcにgitのショートカットコマンドをつくった](http://qiita.com/bmf_san/items/d41bc8d5c5677c69a1e4)でgitのエイリアスコマンドをつくったのですが、中途半端だったので改良しました。

前回のエイリアスでもそこそこにgitコマンドが快適になりましたが、gitコマンドを叩く度にブランチ名をタイプしないといけない仕様は改善すべき点だと思ったので、`select`を使って解決しました。

# ソース
`git branch`の値を`select`で回せばいいと思っていたのですが、ブランチ名だけではなくファイル名とか取得されてしまうので加工する必要がありました。
ちょうど同じようなことを実践している記事があり、そちらを参考にさせていただきました。

作ったコマンドは

+ ローカルブランチにチェックアウトする
+ リモートブランチをローカルにもってくる
+ ローカルブランチを削除する

です。

```sh
#!/bin/sh

#checkout a local branch
function gitCheckoutLocalBranch() {
        branches=`git branch | grep -v -e"^\*" | tr -d ' '`

        PS3="Select branch > "
        echo 'Branch list:'

        select branch in ${branches}
        do
                stty erase ^H
                git checkout ${branch}
                break
        done
}
alias g-c=gitCheckoutLocalBranch

# create a new branch and checktout a remote branch
function gitCreateAndCheckoutRemoteBranch() {
        branches=`git branch -r | grep -v -e"^\*" | tr -d ' '`

        PS3="Select branch > "
        echo 'Branch list:'

        select branch in ${branches}
        do
                stty erase ^H
                echo -n "What is the new branch name?"
                read new_branch_name
                git checkout -b ${new_branch_name} ${branch}
                break
        done
}
alias g-c-b-r=gitCreateAndCheckoutRemoteBranch

# delete a local branch
function gitDeleteLocalBranch() {
        branches=`git branch | grep -v -e"^\*" | tr -d ' '`

        PS3="Select branch > "
        echo 'Branch list:'

        select branch in ${branches}
        do
                stty erase ^H
                git branch -D ${branch}
                break
        done
}
alias g-b-d=gitDeleteLocalBranch
```

`select`で出力される選択肢の色を変えたかったのですが、ちょっとわからなかったので後回しにしました。

# 所感
今回のソースは[github - bmf-san/my-scripts](https://github.com/bmf-san/my-scripts)においてあります。

#参考
* [gitのブランチ移動（checkout）を楽にするやつ](http://qiita.com/amichang/items/5f7e715801771214430e)

