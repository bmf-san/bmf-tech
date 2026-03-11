---
title: シンボリックリンクとハードリンクの違い
slug: symbolic-vs-hard-links
date: 2018-05-24T00:00:00Z
author: bmf-san
categories:
  - OS
tags:
  - シンボリックリンク
  - ハードリンク
translation_key: symbolic-vs-hard-links
---


# 概要
シンボリックリンクとハードリンクの違いについてまとめる

# 前提
- inode
 - データ構造
 - ファイルシステム上の属性情報（作成者、グループ、作成日時など）をデータとして持つ
 - `ls -i1 /`または`stat /`でinode番号を確認できる
 
# シンボリックリンクとは
- 元のファイルまたはディレクトリのパスを参照するディレクトリエントリを追加
- 実験
```
touch a.md
ln -s a.md a_symbolic_link.md // シンボリックリンクを作成
ls -i1 a.md a_symbolic_link.md // inodeが違うことが確認できる
```
- 元ファイルを移動すると参照不可
- 元ファイルを削除すると削除される
- 別のファイルシステムでも参照できる

# ハードリンクとは
- 元のファイルまたはディレクトリのinodeを参照するディレクトリエントリを追加
- 実験
```
touch a.md
ln a.md a_hardlink.md // ハードリンクを設定
ls -i1 a.md a_hardlink.md // inodeが同じことが確認できる
```
- 元のファイルを移動しても影響無し
- 元のファイルを削除しても削除されない
- 同一ファイルシステム内のみ参照可能

# 参照
- [シンボリックリンクとハードリンクの違い](https://qiita.com/katsuo5/items/fc57eaa9330d318ee342)
- [いますぐ実践！Linuxシステム管理](http://www.usupi.org/sysad/242.html)

