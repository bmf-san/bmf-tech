---
title: relineのコードリーディング
description: relineのコードリーディング
slug: reline-code-reading
date: 2024-10-20T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - reline
  - Ruby
translation_key: reline-code-reading
---


# 概要
relineのコードリーディングをする。

# 準備
1. relineをクローンする
`git@github.com:ruby/reline.git`
2. サンプルコードを作成する
```ruby
$LOAD_PATH.unshift(File.expand_path('lib', __dir__))
require "reline"

prompt = 'prompt> '
use_history = true

begin
  while true
    text = Reline.readmultiline(prompt, use_history) do |multiline_input|
      # Accept the input until `end` is entered
      multiline_input.split.last == "end"
    end

    puts 'You entered:'
    puts text
  end
# If you want to exit, type Ctrl-C
rescue Interrupt
  puts '^C'
  exit 0
end
```
3. 任意の箇所で`binding.irb`を挿入する

# コードリーディング
サンプルコードをベースにコードリーディングする。

1. readlineの呼び出し
- [ruby/reline/blob/master/lib/reline.rb#L251](https://github.com/ruby/reline/blob/master/lib/reline.rb#L251)
- [ruby/reline/blob/master/lib/reline.rb#L294](https://github.com/ruby/reline/blob/master/lib/reline.rb#L294)
2. readlineの処理
- 長いので割愛
3. 出力
- [ruby/reline/blob/master/lib/reline/line_editor.rb#L1323](https://github.com/ruby/reline/blob/master/lib/reline/line_editor.rb#L1323)
  - 最終的にバッファに溜め込んだ入力をここで出力する

