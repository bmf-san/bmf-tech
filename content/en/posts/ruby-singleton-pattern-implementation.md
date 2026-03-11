---
title: Implementing Singleton Pattern in Ruby
slug: ruby-singleton-pattern-implementation
date: 2025-01-31T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
description: A summary of how to implement the Singleton pattern in Ruby.
translation_key: ruby-singleton-pattern-implementation
---



A summary of how to implement the Singleton pattern in Ruby.

# What is the Singleton Pattern
The Singleton pattern is a design pattern that ensures only one instance exists.

# How to Create a Singleton Module
Ruby provides a [Singleton module](https://docs.ruby-lang.org/ja/latest/class/Singleton.html) that allows you to implement the Singleton pattern.


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

Since the `new` method is private, you cannot create an instance with `Config.new`. Therefore, you can obtain the only instance with `Config.instance`.

# Using Class Methods
Another way is to make `new` private and manage the instance with a class method.

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

Since `new` is private, you cannot create an instance with `Config.new`. Therefore, you can obtain the only instance with `Config.instance`.

# Using Constants
You can also implement the Singleton pattern using constants instead of class variables.

```ruby
class Config
  INSTANCE = new.freeze # 定数自体の変更を禁止

  private_class_method :new
end

config1 = Config::INSTANCE
config2 = Config::INSTANCE

puts config1 == config2 # => true
```

By assigning a new instance to a constant and making the `new` method private, you can obtain the only instance.

# Summary
There are several ways to implement the Singleton pattern in Ruby.

Among them, using the `Singleton module` is standard and recommended because it provides a thread-safe implementation. (Probably)