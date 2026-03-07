---
title: homebrewで過去バージョンをインストールする
slug: homebrew-installing-previous-versions
date: 2022-10-30T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - homebrew
  - vim
translation_key: homebrew-installing-previous-versions
---


# 概要
Homebrewでインストールするパッケージで過去のバージョンを指定してインストールしたいときがたまにある。
Homebrewは最新版のみ保持する方針になったらしく、過去バージョンをインストールするときはひと手間かかったのでメモ。

# やり方
今回vim9系からvim8系のダウングレードをしたかったので、そのときの手順を例に上げる。

手順は以下。
```sh
brew tap-new bmf-san/vim8
brew extract vim bmf-san/vim8 --version 8.2.5150
brew install bmf-san/vim8/vim@8.2.5150
brew unlink vim
brew link vim@8.2.5150
vim --version // 8.2.5150になっている
```

自分でtap用のリポジトリを用意する必要がある。
githubのリポジトリでも良いが、tap-newというコマンドを使うこともできるのでtap-newを利用。名前は任意。

`brew tap-new bmf-san/vim8`

古いformulaをtapに展開するために、extractする。
このとき指定するバージョンは何でも良いのかもしれないが、過去homebrewで管理していたバージョンを取ってくるのが無難な気がするので、homebrewのリポジトリがほしいバージョンのコミットを漁ってきた。
漁ってきたコミットは以下。

[github.com/Homebrew/homebrew-core/blob/2e3d51340f2f9b47d680f656e712fbee77cbcf79/Formula/vim.rb](https://github.com/Homebrew/homebrew-core/blob/2e3d51340f2f9b47d680f656e712fbee77cbcf79/Formula/vim.rb)

`brew extract vim bmf-san/vim8 --version 8.2.5150`

tapからインストールして、シンボリックリンクをlink&unlinkで調整。

`brew install bmf-san/vim8/vim@8.2.5150`

`brew unlink vim`

`brew link vim@8.2.5150`

# 所感
vimをダウングレードしたかったのはvim-lspがどうやら9系で動作しないような雰囲気を感じたため...

# 参考
- [christina04.hatenablog.com - Homebrewで過去のバージョンを使いたい【tap版】](https://christina04.hatenablog.com/entry/install-old-version-with-homebrew)
