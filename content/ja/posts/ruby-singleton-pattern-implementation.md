---
title: Rubyにおけるシングルトンパターンの実装方法
description: "実装するシングルトンパターン。Singletonモジュール、クラスメソッド活用、定数化によるスレッドセーフな唯一インスタンス管理。"
slug: ruby-singleton-pattern-implementation
date: 2025-01-31T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
translation_key: ruby-singleton-pattern-implementation
---


Rubyにおけるシングルトンパターンの実装方法についてまとめる。

# シングルトンパターンとは
シングルトンパターンは、インスタンスが1つしか存在しないことを保証するデザインパターンである。

# Singletonモジュールを作成する方法
Rubyには[Singletonモジュール](https://docs.ruby-lang.org/ja/latest/class/Singleton.html)が用意されており、これを使うことでシングルトンパターンを実装することができる。


```ruby
require 'singleton'

class Config
  include Singleton

  attr_accessor :setting

  def initialize
    @setting = "default"
  end
end

config1 = Config.instance
config2 = Config.instance

puts config1 == config2 # true 同じインスタンス

config1.setting = "new"
puts config2.setting # new
```

`new`メソッドをprivateになるため、`Config.new`でインスタンスを生成することができない。そのため、`Config.instance`で唯一のインスタンスを取得することができる。

# クラスメソッドを使う方法
`new`をprivateにして、クラスメソッドでインスタンスを管理する方法もある。

```ruby
class Config
  @instance = nil

  private_class_method :new

  def self.instance
    @instance ||= new # 自己代入演算子はnilまたはfalseの場合のみに代入を行う
  end
end

config1 = Config.instance
config2 = Config.instance

puts config1 == config2 # true
```

`new`がprivateであるため、`Config.new`でインスタンスを生成することができない。そのため、`Config.instance`で唯一のインスタンスを取得することができる。

# 定数を使う方法
クラス変数を使わず、定数を使ってシングルトンパターンを実装する方法もある。

```ruby
class Config
  INSTANCE = new.freeze # 定数自体の変更を禁止

  private_class_method :new
end

config1 = Config::INSTANCE
config2 = Config::INSTANCE

puts config1 == config2 # => true
```

定数にnewしたインスタンスを代入し、newメソッドをprivateにすることで、唯一のインスタンスを取得することができる。

# まとめ
Rubyにはシングルトンパターンを実装するための方法がいくつかある。

その中でも、`Singletonモジュール`を使う方法が標準的であり、スレッドセーフな安全な実装であるため推奨される。（たぶん）
