---
title: About Ruby Modules
slug: ruby-modules
date: 2024-05-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
translation_key: ruby-modules
---

# Overview
This post discusses Ruby modules.

# What is a Module?
A mechanism to provide common methods and constants to classes and other modules.

```ruby
# Module definition
module Hi
  def say_hi
    puts "Hi!"
  end
end
```

Unlike classes, modules cannot be instantiated and cannot be inherited.

Modules can define class methods and instance methods.

Class methods cannot be called from the included module.

```ruby
module Greet
  # Module class method
  def self.hi
    puts "Hi!"
  end

  # Module instance method
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

## Namespace
Can be used to provide a namespace.

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
Allows adding or overriding instance methods in a class without inheritance.

Classes cannot inherit multiple times, but multiple inheritance can be achieved through Module mixins.

```ruby
class Greet
  include Hi
end

puts Greet.new.say_hi # => Hi!
```

By the way, Mixin and Trait are similar, but Mixin uses inheritance while Trait can combine methods through various means other than inheritance, which gives them slightly different nuances.

cf. [ja.wikipedia.org - Mixin](https://ja.wikipedia.org/wiki/Mixin)
cf. [ja.wikipedia.org - Trait](https://ja.wikipedia.org/wiki/%E3%83%88%E3%83%AC%E3%82%A4%E3%83%88)

## Adding Singleton Methods to Classes Using extend
Using extend, you can add singleton methods to a class.

```ruby
module Hi
  def hi
    puts "Hi!"
  end
end

class Greet; end

Greet.new.extend(Hi).hi # => Hi!
```

# References
- [Module - docs.ruby-lang.org](https://docs.ruby-lang.org/ja/latest/class/Module.html)
- [qiita.com - Implementing Inheritance for Mixin: Ruby module, Java interface, PHP trait](https://qiita.com/niwasawa/items/82a5611b23f4a95aac04)
- [takayukinakata.hatenablog.com - Why can't class methods be inherited in mix-ins in Ruby?](https://takayukinakata.hatenablog.com/entry/2017/03/04/183546)