---
title: Rubyの特異クラス・特異メソッドについて
description: Rubyの特異クラス・特異メソッドについてについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: ruby-singleton-class-methods
date: 2024-05-15T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
translation_key: ruby-singleton-class-methods
---


# 概要
Rubyの特異クラスと特異メソッドについてかく。

# 特異クラス
特異クラスとは、特定のオブジェクトに対してのみ有効なクラスのことを指す。

```ruby
class Greet
  def say_hi
    puts 'Hi'
  end
end

greet = Greet.new

# 特異クラス
class << greet
  # シングルトンメソッド
  def say_bye
    puts 'Bye'
  end
end

greet.say_hi # => Hi
greet.say_bye # => Bye
p greet.singleton_methods # => [:say_bye]
singleton_class = greet.singleton_class
puts singleton_class # => #<Class:#<Greet:0x00007f8b1b0>>
```

# 特異メソッド
特異メソッドとは、特定のオブジェクトに対してのみ有効なメソッドのことを指す。

```ruby
class Greet
  def self.say_hi
    puts "Hi"
  end
end

Greet.say_hi # => Hi

obj = Greet.new

# シングルトンメソッド
def obj.say_bye
  puts "Bye"
end

obj.say_bye # => Bye
p obj.singleton_methods # => [:say_bye]
```

特異メソッドは特定のオブジェクトだけが持つメソッドであるため、そのオブジェクトに対してのみ有効である。

```ruby
class Greet; end

obj1 = Greet.new
obj2 = Greet.new

def obj1.say_hi
  puts "Hi"
end

obj1.say_hi # Hiw
obj2.say_hi # NoMethodError
```

クラスメソッドは実質的に特異メソッドである。

```ruby
class Greet
  # クラスメソッド定義
  class << self
    def say_hi
      puts "Hello"
    end
  end
end

puts Greet.say_hi

# クラス定義の外でクラスメソッドを定義することもできる
class Greet; end
class << Greet
  def say_hi
    puts "Hello"
  end
end

puts Greet.say_hi

# 次のようにクラスメソッドを定義することもできる
class Greet
  def self.say_hi
    puts "Hi"
  end
end
```

# 参考
- [docs.ruby-land.org - クラス／メソッドの定義](https://docs.ruby-lang.org/ja/latest/doc/spec=2fdef.html#singleton_class)
- [www.school.ctc-g.co.jp - 第12回　クラスインスタンス変数と特異クラス・特異メソッド(3) (中越智哉)](https://www.school.ctc-g.co.jp/columns/nakagoshi/nakagoshi12.html)
