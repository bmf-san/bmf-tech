---
title: "開発環境を見直してvimライクにした"
slug: "vim-like-development-environment"
date: 2018-05-22
author: bmf-san
categories:
  - "ツール"
tags:
  - "tmux"
  - "vim"
  - "Atom"
  - "iTerm"
draft: false
---

# 概要
開発の効率化を図り、vimを取り入れ、開発環境諸々を刷新したのでまとめておく。
各ツールの細かい設定や導入しているプラグイン詳細などは省く。　

# エディタ
- Atom
	- 開発で使うメインエディタ
	- vimのキーバインドを使えるようにプラグインを導入
		- vim-mode-plus-ex-mode
		- vim-mode-plus
  	- 画面分割やペイン移動などもvimっぽくできるようにキーバインドを少しいじった

- Vim
	- ちょっとした編集などサブで使う
	- プラグインは最低限導入

# コマンドラインツール
- iTerm2
	- bashでvi-modeを有効化
		- コマンドライン上でviキーバインドが使える

# 端末多重化ソフトウェア
- tmux
	- ゴリゴリコマンドラインを使っていくには必須かもしれない

# 参考
- [github - bmf-san/my-dotfiles](https://github.com/bmf-san/my-dotfiles)
	- 現在設定しているatomやvimの設定など

# 所感
vimの思考を受け入れ、vimのキーバインドをあらゆるツール上で適用したことで幸せになれた気がする。
