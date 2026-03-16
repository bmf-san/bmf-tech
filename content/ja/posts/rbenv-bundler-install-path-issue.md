---
title: anyenvでインストールしたrbenvでbundlerをinstallしたときにパスでハマった
description: anyenvでインストールしたrbenvでbundlerをinstallしたときにパスでハマった
slug: rbenv-bundler-install-path-issue
date: 2018-12-04T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - anyenv
  - gem
  - rbenv
  - Ruby
  - Tips
translation_key: rbenv-bundler-install-path-issue
---


# 概要
 anyenvでインストールしたrbenvでbundlerをinstallしたときにパスでハマった話。

# ハマったこと
anyenvでrbenvをインストールしてrubyを使っているのですが、bundlerをインストールする際に、

`gem install bundler`

と何も考えずに打つと、bundlerが`/usr/local/bin/`以下に配置されてしまう。

意図したパスでないためgemでinstallしたchefとか使おうとするとコケる。

# 解決策
`rbenv exec gem install bundler`

rbenvで導入しているrubyのgemを実行するように指定する。

# 所感
パスを冷静に確認していればrubyに不慣れでもすぐわかったはず...

# 参考
- [github - rbenv/rbenv](https://github.com/rbenv/rbenv)

