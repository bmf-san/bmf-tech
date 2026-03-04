---
title: "tmuxコトハジメ"
slug: "tmux"
date: 2018-05-22
author: bmf-san
categories:
  - "ツール"
tags:
  - "tmux"
draft: false
---

# 概要
端末多重化ソフトウェアであるtmuxのコトハジメ

# キーバインド
tmux起動
tmux or tmux new-session

セッション中での新規セッション作成
prefix+:new

セッション一覧
tmux ls

セッションのデタッチ（tmuxから抜ける）
prefix+d

セッションのアタッチ
tmux attach(a)

任意のセッションにアタッチ
tmux attach(a) -t 0(name)

セッションの削除
tmux kill-session

任意のセッションを削除
tmux kill-session -t 0

全てのセッションを’削除
tmux kill-server

セッションのリネーム
prefix+$

新規ウィンドウ
prefix+c

次のウィンドウに切り替え
prefix+n

前のウィンドウに切り替え
prefix+p

任意のウィンドウに切り替え
prefix+0

ウィンドウ一覧
prefix+w

ウィンドウの削除
prefix+&

ペインの削除
prefix+x

ペイン入れ替え（前方向）
prefix+{

ペイン入れ替え（後方向）
prefix+}

コピーモード
prefix+[

コピー範囲選択（コピーモード中）
vまたはspace

コピー（コピーモード中）
yまたはenter

# tips
macのterminalでtmuxを使っているとき、マウスを使ってテキストを選択してもコピー（cmd+c）ができない。（tmuxのショートカットでvimライクなコピーはできる）
マウスで選択した範囲をコピーしたい場合は、cmd+rでterminalのAllow mouse reportingの設定をトグルする必要がある。 

# 参考
- [github - bmf-san/dotfiles](https://github.com/bmf-san/dotfiles)
  - tmuxの設定を置いてある

