---
title: About Ruby Modules
description: 'Understand Ruby modules for namespacing, mixin multiple inheritance, and providing common methods without class instantiation.'
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
This article discusses Ruby modules.

# What is a Module?
A mechanism for providing common methods and constants to classes and other modules.

```ruby
# Module definition
module Hi
  def say_hi
    puts "Hi!"
  end
end
```

Unlike classes, modules cannot be instantiated. They also cannot be inherited.

Modules can define class methods and instance methods.

Class methods cannot be called from the module's include target.

```ruby
module Greet
  # Class method of the module
  def self.hi
    puts "Hi!"
  end

  # Instance method of the module
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
Modules can be used to create namespaces.

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
    puts "I am a student"
  end
end

Student.say # => I am a student
University::Student.say # => I am a student
```

## Mixin
Mixins allow adding or overriding instance methods in classes without using inheritance.

While classes cannot inherit multiple times, multiple inheritance can be achieved through module mixins.

```ruby
class Greet
  include Hi
end

puts Greet.new.say_hi # => Hi!
```

It is worth noting that while Mixins and Traits are similar, Mixins use inheritance, whereas Traits can compose methods through various means other than inheritance, giving them a slightly different nuance.

cf. [ja.wikipedia.org - Mixin](https://ja.wikipedia.org/wiki/Mixin)  
cf. [ja.wikipedia.org - Trait](https://ja.wikipedia.org/wiki/%E3%83%88%E3%83%AC%E3%82%A4%E3%83%88)

## Adding Singleton Methods to Classes Using `extend`
You can add singleton methods to a class using `extend`.

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
- [qiita.com - Implementing Mixin-like Inheritance: Ruby module, Java interface, PHP trait](https://qiita.com/niwasawa/items/82a5611b23f4a95aac04)
- [takayukinakata.hatenablog.com - Why can't class methods be inherited in Ruby mix-ins?](https://takayukinakata.hatenablog.com/entry/2017/03/04/183546)
