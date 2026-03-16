---
title: execコマンドによるリダイレクト
description: execコマンドによるリダイレクト
slug: exec-command-redirection
date: 2019-05-08T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - bash
  - shellscript
  - exec
translation_key: exec-command-redirection
---


# 概要
execコマンドは現在のプロセスを実行するコマンドで置き換えるコマンドだが、引数無しで使うとリダイレクトの動的変更ができる。

ちょうど[master直プッシュの際にプロンプトで確認するようにする](https://bmf-tech.com/posts/master%E7%9B%B4%E3%83%97%E3%83%83%E3%82%B7%E3%83%A5%E3%81%AE%E9%9A%9B%E3%81%AB%E3%83%97%E3%83%AD%E3%83%B3%E3%83%97%E3%83%88%E3%81%A7%E7%A2%BA%E8%AA%8D%E3%81%99%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%99%E3%82%8B)で

```sh
#!/bin/sh
exec < /dev/tty
read ANSWER
```

というコードが出てきてよくわからなかったので調べてみたのがきっかけ。

#  使い方
```sh
#!/bin/sh
echo "Output to stdout" // 標準入力
exec > redirect.txt // ファイルディスクリプタを変更
echo "Output to file" // ファイルに出力される
```

[master直プッシュの際にプロンプトで確認するようにする](https://bmf-tech.com/posts/master%E7%9B%B4%E3%83%97%E3%83%83%E3%82%B7%E3%83%A5%E3%81%AE%E9%9A%9B%E3%81%AB%E3%83%97%E3%83%AD%E3%83%B3%E3%83%97%E3%83%88%E3%81%A7%E7%A2%BA%E8%AA%8D%E3%81%99%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%99%E3%82%8B)に出てきたコードは、

```sh
#!/bin/sh
exec < /dev/tty
read ANSWER
```

現在の端末（`/dev/tty`）の入力をexecコマンドの標準入力に渡している。（わかりづらい。。。）

# 所感
何らかの事情で標準入力をreadできないときとかに使えそう。

# 参考
- [SEEKPOINT - execリダイレクト](http://seekpoint.blogspot.com/2012/12/exec.html)
- [UNIX/Linuxの部屋 コマンド:exec](http://x68000.q-e-d.net/~68user/unix/pickup?exec#prgmemo-exec-basic)
- [exec - シェルを実行したコマンドで置き換える](https://linuxcommand.net/exec/)
- [Qiita - Process Substitutionとexec redirectで画面出力を加工するときの問題点と解決、そして無限に寝る話](https://qiita.com/takei-yuya@github/items/7afcb92cfe7e678b7f6d#%E3%81%AF%E3%81%98%E3%82%81%E3%81%AB2-exec-%E3%81%A8%E3%83%AA%E3%83%80%E3%82%A4%E3%83%AC%E3%82%AF%E3%83%88)
- [Qiita - シェルの入出力制御あれこれ](https://qiita.com/tag1216/items/7ce35b7c27d371165e56#%E6%A8%99%E6%BA%96%E5%87%BA%E5%8A%9B%E3%81%A8%E6%A8%99%E6%BA%96%E3%82%A8%E3%83%A9%E3%83%BC%E5%87%BA%E5%8A%9B%E3%82%92%E5%88%A5%E3%80%85%E3%81%AE%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%AB%E5%87%BA%E5%8A%9B%E3%81%99%E3%82%8B)
