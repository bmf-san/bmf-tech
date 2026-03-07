---
title: Implementing the Singleton Pattern in Ruby
slug: ruby-singleton-pattern-implementation
date: 2025-01-31T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
translation_key: ruby-singleton-pattern-implementation
---

This post summarizes how to implement the Singleton pattern in Ruby.

# What is the Singleton Pattern
The Singleton pattern is a design pattern that guarantees that only one instance exists.

# How to Create a Singleton Module
Ruby provides the [Singleton module](https://docs.ruby-lang.org/ja/latest/class/Singleton.html), which can be used to implement the Singleton pattern.

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

puts config1 == config2 # true Same instance

config1.setting = "new"
puts config2.setting # new
```

The `new` method is made private, so an instance cannot be created with `Config.new`. Therefore, the only instance can be obtained with `Config.instance`.

# Using Class Methods
Another method is to make `new` private and manage the instance with a class method.

```ruby
class Config
  @instance = nil

  private_class_method :new

  def self.instance
    @instance ||= new # The self-assignment operator only assigns if nil or false
  end
end

config1 = Config.instance
config2 = Config.instance

puts config1 == config2 # true
```

Since `new` is private, an instance cannot be created with `Config.new`. Therefore, the only instance can be obtained with `Config.instance`.

# Using Constants
It is also possible to implement the Singleton pattern using constants instead of class variables.

```ruby
class Config
  INSTANCE = new.freeze # Prevents changes to the constant itself

  private_class_method :new
end

config1 = Config::INSTANCE
config2 = Config::INSTANCE

puts config1 == config2 # => true
```

By assigning a new instance to a constant and making the new method private, the only instance can be obtained.

# Conclusion
There are several ways to implement the Singleton pattern in Ruby.

Among them, using the `Singleton module` is the standard method and is recommended because it is a thread-safe implementation. (Probably)