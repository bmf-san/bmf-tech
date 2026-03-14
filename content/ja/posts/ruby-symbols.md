---
title: Rubyのシンボルについて
description: "活用するRubyシンボル。文字列との効率性比較、イミュータブル性、ハッシュキー利用、enum的な状態管理パターン。"
slug: ruby-symbols
date: 2024-05-14T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
translation_key: ruby-symbols
---


# 概要
Rubyのシンボルについて書く。

# シンボルとは
任意の文字列と対に対応するオブジェクト。

内部実装では、整数として管理されている。

```ruby
:symbol
:'symbol'
%s!symbol!
```

シンボルはSymbolクラスのオブジェクトであり、文字列はStringクラスのオブジェクトである。そのためメモリ効率が良い。

文字列と違い、シンボルは同一のオブジェクトである。

```ruby
# シンボルは全て同じID
puts :symbol.object_id # => 865628
puts :symbol.object_id # => 865628
puts :symbol.object_id # => 865628

# 文字列は全て異なるID
puts 'symbol'.object_id # => 60
puts 'symbol'.object_id # => 80
puts 'symbol'.object_id # => 100
```

また、シンボルはイミュータブルなオブジェクトでもある。

```ruby
# 書き換え不可であるためエラーとなる
symbol = :Symbol
symbol.sub(/Sym/, 'sym') # => undefined method `sub' for an instance of Symbol (NoMethodError)
```

用途として次のようなケースがある。

```ruby
# ハッシュのキー
hash = { :key => "value" }
puts hash[:key] # => value

# アクセサの引数で渡すインスタンス変数名として使用
class Order
  attr_reader :id

  def initialize(id)
    @id = id
  end
end

order = Order.new(1)
puts order.id # => 1

# メソッド引数で渡すメソッド名として使用
text = "hello"
puts text.__send__(:to_s) # => hello

# Cのenum的な使用
STATUS_ACTIVE = :active
STATUS_INACTIVE = :inactive

def puts_status(status)
  case status
  when STATUS_ACTIVE
    puts "有効"
  when STATUS_INACTIVE
    puts "無効"
  end
end

puts_status(STATUS_ACTIVE) # => 有効
puts_status(STATUS_INACTIVE) # => 無効
```

# 参考
- [docs.ruby-lang.org - class Symbol](https://docs.ruby-lang.org/ja/latest/class/Symbol.html)
- [qiita.com - シンボルとは！？【Ruby超入門】](https://qiita.com/yyykms123/items/6a6ae7fe8cd9263a3d1c)
- [zenn.dev - Rubyのシンボルを丁寧に理解する](https://zenn.dev/kanoe/articles/352d78902c83e168db66)
- [zenn.dev - Ruby初学者が困るシンボル「：」を簡単に理解する](https://zenn.dev/hiro_xre/articles/709934e347edc3)
- [techracho.bpsinc.jp - Rubyにシンボルがある理由（翻訳）](https://techracho.bpsinc.jp/hachi8833/2022_04_28/117351)
