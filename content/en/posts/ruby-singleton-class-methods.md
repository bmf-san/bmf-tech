---
title: About Ruby's Singleton Classes and Singleton Methods
slug: ruby-singleton-class-methods
date: 2024-05-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
translation_key: ruby-singleton-class-methods
---



# Overview
This post discusses Ruby's singleton classes and singleton methods.

# Singleton Classes
A singleton class refers to a class that is only valid for a specific object.

```ruby
class Greet
  def say_hi
    puts 'Hi'
  end
end

greet = Greet.new

# Singleton Class
class << greet
  # Singleton Method
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

# Singleton Methods
A singleton method is a method that is only valid for a specific object.

```ruby
class Greet
  def self.say_hi
    puts "Hi"
  end
end

Greet.say_hi # => Hi

obj = Greet.new

# Singleton Method
def obj.say_bye
  puts "Bye"
end

obj.say_bye # => Bye
p obj.singleton_methods # => [:say_bye]
```

Since singleton methods are methods that only a specific object possesses, they are only valid for that object.

```ruby
class Greet; end

obj1 = Greet.new
obj2 = Greet.new

def obj1.say_hi
  puts "Hi"
end

obj1.say_hi # Hi
obj2.say_hi # NoMethodError
```

Class methods are essentially singleton methods.

```ruby
class Greet
  # Class Method Definition
  class << self
    def say_hi
      puts "Hello"
    end
  end
end

puts Greet.say_hi

# Class methods can also be defined outside the class definition
class Greet; end
class << Greet
  def say_hi
    puts "Hello"
  end
end

puts Greet.say_hi

# Class methods can also be defined as follows
class Greet
  def self.say_hi
    puts "Hi"
  end
end
```

# References
- [docs.ruby-lang.org - Class/Method Definitions](https://docs.ruby-lang.org/ja/latest/doc/spec=2fdef.html#singleton_class)
- [www.school.ctc-g.co.jp - Episode 12: Class Instance Variables and Singleton Classes/Methods (3) (Tomoya Nakagoshi)](https://www.school.ctc-g.co.jp/columns/nakagoshi/nakagoshi12.html)
