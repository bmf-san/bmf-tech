---
title: macにanyenvをインストールする
description: macにanyenvをインストールするについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: installing-anyenv-on-mac
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - shellscript
  - anyenv
translation_key: installing-anyenv-on-mac
---


Macにanyenvをインストールする手順です。
若干ハマったところがあったのでメモ。

# インストール
私の環境では、`/usr/local/bin/`に色々置いてあるので、そこにanyenvをインストールすることにします。

`cd /usr/local/bin`
`git clone https://github.com/riywo/anyenv`

# Pathを通す

```shell-session:~/.bashrc
export PATH="/usr/local/bin/anyenv/bin:$PATH"
export ANYENV_ROOT=/usr/local/bin/anyenv
eval "$(anyenv init -)"
```

anyenvはルートディレクトリ直下にインストールすることを想定しているせいか、`ANYENV_ROOT`を任意のディレクトリに指定しないとanyenvコマンドが正しく実行されませんでした。
それから`eval "$(anyenv init -)"`を書き忘れるとインストールしたパッケージのコマンドが実行できないなど不具合が起きるので忘れないようにしましょう。

# おわり
これでインストールは完了です。　anyenvコマンド各種使えるようになっているかと思います。

#  参考
- [github - riywo/anyenv](https://github.com/riywo/anyenv)

