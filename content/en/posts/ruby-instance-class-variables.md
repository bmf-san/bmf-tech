---
title: Differences Between Instance Variables, Class Variables, and Class Instance Variables in Ruby
slug: ruby-instance-class-variables
date: 2025-01-31T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
description: Summarizing the differences between instance variables, class variables, and class instance variables in Ruby.
translation_key: ruby-instance-class-variables
---

Summarizing the differences between instance variables, class variables, and class instance variables in Ruby.

# Instance Variables
Variables that hold different values for each instance, using variable names that start with `@`.

They are mainly used to manage the state of each object declared within a class.

```ruby
class Person
  def initialize(name)
    @name = name # Instance variable
  end

  def name
    @name
  end
end

p1 = Person.new("Alice")
p2 = Person.new("Bob")

puts p1.name # => Alice
puts p2.name # => Bob Different instances hold different values
```

# Class Variables
Variables shared across the entire class, using variable names that start with `@@`.

They are used to hold values shared between different instances of the same class.

```ruby
class Person
  @@count = 0 # Class variable

  def initialize(name)
    @name = name
    @@count += 1
  end

  def self.count
    @@count
  end
end

p1 = Person.new("Alice")
p2 = Person.new("Bob")

puts Person.count # => 2 Shared across all instances
```

They are also shared with subclasses.

```ruby
class Parent
  @@var = "Parent"

  def self.var
    @@var
  end
end

class Child < Parent
  @@var = "Child"
end
puts Parent.var # => "Child" Overwritten by subclass
```

# Class Instance Variables
Unlike class variables, they are not shared with subclasses.

They are used as instance variables of the class itself.

Access them within class methods using `self`.

```ruby
class Person
  @count = 0 # Class instance variable

  def self.increment_count
    @count ||= 0
    @count += 1
  end

  def self.count
    @count
  end
end

class Student < Person
  @count = 0
end

Person.increment_count
Student.increment_count

puts Person.count # => 1 Count for Person class
puts Student.count # => 1 Count for Student class
```

# Summary of Differences
| Variable Type                                      | Declaration Location                  | Scope                  | Impact on Subclasses          | Usage                            |
| -------------------------------------------------- | ------------------------------------- | ---------------------- | ----------------------------- | -------------------------------- |
| **Instance Variable (`@var`)**                     | Within `initialize` or instance methods | **Per instance**       | No impact                     | Managing data per instance       |
| **Class Variable (`@@var`)**                       | Within class                          | **Shared across all instances** | **Impacts (shared)**           | Managing common data across instances |
| **Class Instance Variable (`@var` in class context)** | Top level of class (`self.@var`)    | **Per class**          | No impact (independent per class) | Managing data per class          |