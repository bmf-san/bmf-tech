---
title: About Ruby Symbols
description: 'Understand Ruby symbols as memory-efficient, immutable identifiers for hash keys, method names, and enum-like constants.'
slug: ruby-symbols
date: 2024-05-14T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
translation_key: ruby-symbols
---



# Overview
Writing about Ruby symbols.

# What is a Symbol?
An object that corresponds to an arbitrary string.

Internally, it is managed as an integer.

```ruby
:symbol
:'symbol'
%s!symbol!
```

A symbol is an object of the Symbol class, while a string is an object of the String class, making it memory efficient.

Unlike strings, symbols are the same object.

```ruby
# All symbols have the same ID
puts :symbol.object_id # => 865628
puts :symbol.object_id # => 865628
puts :symbol.object_id # => 865628

# All strings have different IDs
puts 'symbol'.object_id # => 60
puts 'symbol'.object_id # => 80
puts 'symbol'.object_id # => 100
```

Additionally, symbols are immutable objects.

```ruby
# Cannot be modified, resulting in an error
symbol = :Symbol
symbol.sub(/Sym/, 'sym') # => undefined method `sub' for an instance of Symbol (NoMethodError)
```

Here are some use cases.

```ruby
# Hash keys
hash = { :key => "value" }
puts hash[:key] # => value

# Used as instance variable names passed as accessor arguments
class Order
  attr_reader :id

  def initialize(id)
    @id = id
  end
end

order = Order.new(1)
puts order.id # => 1

# Used as method names passed as method arguments
text = "hello"
puts text.__send__(:to_s) # => hello

# Used like C enums
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

# References
- [docs.ruby-lang.org - class Symbol](https://docs.ruby-lang.org/ja/latest/class/Symbol.html)
- [qiita.com - What is a Symbol!? 【Ruby Super Beginner】](https://qiita.com/yyykms123/items/6a6ae7fe8cd9263a3d1c)
- [zenn.dev - Understanding Ruby Symbols Thoroughly](https://zenn.dev/kanoe/articles/352d78902c83e168db66)
- [zenn.dev - Easily Understand the Symbol ":" that Confuses Ruby Beginners](https://zenn.dev/hiro_xre/articles/709934e347edc3)
- [techracho.bpsinc.jp - Why Ruby Has Symbols (Translation)](https://techracho.bpsinc.jp/hachi8833/2022_04_28/117351)
