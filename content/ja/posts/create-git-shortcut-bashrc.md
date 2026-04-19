---
title: .bashrcにgitのショートカットコマンドをつくった
description: .bashrcにgitのショートカットコマンドをつくった
slug: create-git-shortcut-bashrc
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - bash
  - Git
  - shellscript
translation_key: create-git-shortcut-bashrc
---


git add hogehoge, git commit hogehoge, git push hogehoge....

基本的なgitコマンドしか使わないのですが、毎回コマンド叩くの面倒くさい、楽したいということでエイリアスをつくってみました。

# スクリプト

```vim
#git branch
alias git-b='git branch'

#git checkout
function gitCheckout() {
         stty erase ^H
         echo -n "What is the new branch name"?
         stty echo
         read var1
         git checkout ${var1}
}
alias git-c=gitCheckout

#git checkout -b
function gitCheckoutBranch() {
         stty erase ^H
         echo -n "What is the new branch name for checkout?"
         stty echo
         read var1
         git checkout -b ${var1}
}
alias git-c-b=gitCheckoutBranch

#git pull
function gitPull() {
        stty erase ^H
        echo -n "What is the remote repository name?"
        stty echo
        read var1
        git pull origin ${var1}
}
alias git-p=gitPull

#git set
function gitSet() {
      stty erase ^H
      echo -n "What file name do you add?"
      read var1
      git add ${var1}
      echo -n "What is the commit message?"
      read var2
      git commit -m\'${var2}\'
      echo -n "What is the branch name?"
      stty echo
      read var3
      git push origin ${var3}
}
alias git-set=gitSet
```

# コマンド説明
* git-b ・・・ブランチ確認
* git-c ・・・チェックアウト
* git-c-b ・・・新しくブランチを切ってチェックアウト
* git-p　・・・プル
* git-set ・・・add/commit/pushを対話形式でやる。英語がカッコワルイので訂正求む。

# トラブルシューティング
* .bashrcが読み込まれない
  * .bash_profileで.bashrcを読み込んでいるか確認
     * cf. ~~[bash] ターミナルを起動しても.bashrcが読み込まれていない時に確認すること~~

* エイリアスコマンド実行後の入力受付時にでbackspace(delete)キーが変な文字列に変換されてしまう
  * キーボードがUS配列→.bashrcの^Hのまま
  * JIS→^Hを^?に変更

# 所感
* 開発効率がほんの少し上がった！
* git使っている人は普段どうしているの？　エイリアスつくったりする？？

# 追記
`git config`とかいうgit用のエイリアス設定コマンドがあったのを忘れていました。
[よく使うgitコマンドのエイリアスを設定して開発効率をアップする](http://qiita.com/unsoluble_sugar/items/ce14e9ce20aa5ba34fe5)

