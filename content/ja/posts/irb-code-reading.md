---
title: irbのコードリーディング
slug: irb-code-reading
date: 2024-10-20T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - irb
  - Ruby
translation_key: irb-code-reading
---


# 概要　
irbのコードリーディングをする。

# 準備
1. irbのソースコードをクローンする
- `git clone git@github.com:ruby/irb.git`
2. エントリポイントを細工する
元の状態だとRubyインストールディレクトリにあるirbを実行してしまうので、ローカルのirbを実行するようにする。
```ruby
#!/usr/bin/env ruby
#
#   irb.rb - interactive ruby
#   	by Keiju ISHITSUKA(keiju@ruby-lang.org)
#

$LOAD_PATH.unshift(File.expand_path("../lib", __dir__)) # ここを追加
require "irb"

IRB.start(__FILE__)
```
3. 任意の箇所でデバッグする

# コードリーディング
1. irbコマンドの実行
- [ruby/irb/blob/master/exe/irb#L9](https://github.com/ruby/irb/blob/master/exe/irb#L9)
  - 実行しているファイル名を引数に、`IRB.start(__FILE__)`を実行している
2. 起動時のセットアップ処理
- [ruby/irb/blob/master/lib/irb.rb#L895](https://github.com/ruby/irb/blob/master/lib/irb.rb#L895)
3. irbの起動
- [ruby/irb/blob/master/lib/irb.rb#L1001](https://github.com/ruby/irb/blob/master/lib/irb.rb#L1001)
  - historyの読み込みをしている
    - historyはコマンドの実行履歴。ファイル（~/.irb_history）に書き込まれている.
4. 入力値の実行
- [ruby/irb/blob/master/lib/irb.rb#L1041](https://github.com/ruby/irb/blob/master/lib/irb.rb#L1041)
- [ruby/irb/blob/master/lib/irb/context.rb#L589](https://github.com/ruby/irb/blob/master/lib/irb/context.rb#L589)
  - 入力値によって処理を分岐をして実行している
  - 式の場合
    - [ruby/irb/blob/master/lib/irb/context.rb#L589](https://github.com/ruby/irb/blob/master/lib/irb/context.rb#L589)
    - [ruby/irb/blob/master/lib/irb/context.rb#L609](https://github.com/ruby/irb/blob/master/lib/irb/context.rb#L609)
    - [ruby/irb/blob/master/lib/irb/workspace.rb#L120](https://github.com/ruby/irb/blob/master/lib/irb/workspace.rb#L120)
      - `eval`で式を評価している
  - コマンドの場合
    - [ruby/irb/blob/master/lib/irb/command/base.rb#L55](https://github.com/ruby/irb/blob/master/lib/irb/command/base.rb#L55)
      - コマンドが何かわからなかったが、利用者向けの機能ではなさそう

