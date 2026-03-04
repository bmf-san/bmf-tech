---
title: "RubyのModuleについて"
slug: "ruby-modules"
date: 2024-05-15
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Ruby"
draft: false
---

# 概要
RubyのModuleについてかく。

# Moduleとは
クラスや他のモジュールに共通のメソッドや定数を提供するための仕組み。

```ruby
# モジュール定義
module Hi
  def say_hi
    puts "Hi!"
  end
end
```

Moduleはクラスと異なり、インスタンス化できない。また継承もできない。

モジュールにはクラスメソッドやインスタンスメソッドを定義することができる。

クラスメソッドはモジュールのinclude先で呼び出すことができない。

```ruby
module Greet
  # モジュールのクラスメソッド
  def self.hi
    puts "Hi!"
  end

  # モジュールのインスタンスメソッド
  def bye
    puts "Bye!"
  end
end

Greet.hi

class Speaker
  include Greet
end

speaker = Speaker.new
speaker.bye # => Bye!
speaker.hi # => NoMethodError
```

## 名前空間
名前空間を用意するために使うことができる。

```ruby
module University
  class Student
    def self.say
      puts "I am a student"
    end
  end
end

class Student
  def self.say
    puts "私は学生です"
  end
end

Student.say # => 私は学生です
University::Student.say # => I am a student

```

## Mixin
継承を用いずクラスにインスタンスメソッドを追加・上書きすることができる。

クラスは多重継承ができないが、ModuleのMixinにより多重継承を実現できる。

```ruby

class Greet
  include Hi
end

puts Greet.new.say_hi # => Hi!
```

ちなみにMixinとTraitは似ているが、Mixinは継承を用いているのに対して、Traitは継承以外の様々方法によってもメソッドを合成することができるものであり、ややニュアンスが異なる。

cf. [ja.wikipedia.org - Mixin](https://ja.wikipedia.org/wiki/Mixin)
cf. [ja.wikipedia.org - トレイト](https://ja.wikipedia.org/wiki/%E3%83%88%E3%83%AC%E3%82%A4%E3%83%88)

## extendを使ったクラスへの特異メソッド追加
extendを使うことで、クラスに特異メソッドを追加することができる。

```ruby
module Hi
  def hi
    puts "Hi!"
  end
end

class Greet; end

Greet.new.extend(Hi).hi # => Hi!
```

# 参考
- [Module - docs.ruby-lang.org](https://docs.ruby-lang.org/ja/latest/class/Module.html)
- [qiita.com - Mixin 的な実装の継承を実現する: Ruby module, Java interface, PHP trait](https://qiita.com/niwasawa/items/82a5611b23f4a95aac04)
- [takayukinakata.hatenablog.com - Rubyではなぜmix-inでクラスメソッドを引き継げないのか](https://takayukinakata.hatenablog.com/entry/2017/03/04/183546)
