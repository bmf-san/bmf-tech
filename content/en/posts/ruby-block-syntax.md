---
title: About Ruby Block Syntax
description: 'Learn Ruby block syntax with do..end and {} forms, passing blocks to methods, yield statements, and Proc/lambda closure behavior.'
slug: ruby-block-syntax
date: 2024-05-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
translation_key: ruby-block-syntax
---

# Overview
This post discusses Ruby's block syntax.

# What is Block Syntax?
A syntax that allows you to pass code snippets to other methods, enabling those methods to execute the code within.

Defined using `do..end` or `{}`. Blocks can accept arguments.

```ruby
# do..end
[1...3].each do |number|
  puts number
end

# {}
[1...3].each { |number| puts number}
```

You can also define methods that accept blocks.

```ruby
def greet
  puts "Hi"
  # Execute the passed block
  yield if block_given? # block_given? checks if a block was passed
  puts "Bye"
end

# Passing a block

greet do
  puts "Hello"
end

# Output:
# Hi
# Hello
# Bye
```

Additionally, you can pass blocks as method arguments using `&`.

```ruby
def greet(&block)
  puts "Hi"
  block.call if block
  puts "Bye"
end

greet do
  puts "Hello"
end

# Output:
# Hi
# Hello
# Bye
```

Objects with a to_proc method can also be passed as blocks.

```ruby
class Greeting
  def to_proc
    Proc.new { |name| puts "Hello, #{name}!" }
  end
end

%w(Alice Bob Charlie).each(&Greeting.new)

# Output:
# Hello, Alice!
# Hello, Bob!
# Hello, Charlie!
```

Blocks are closures, just like Proc and lambda.

```ruby
def count(&block)
  puts block.call
  return block
end

x = 0
b = count do
  x + 1
end # =>1

x += 10
puts b.call # =>11
```

`do..end` and `{}` have different precedence, with `{}` having higher precedence.

```ruby
array = [1, 2, 3]

# Low precedence, so the block is evaluated before being passed to map (it gets passed to pp instead)
pp array.map do |x|
  x * 2
end

# Output:
#<Enumerator: ...>

# High precedence, so the block is passed to map and evaluated
pp array.map { |x| x * 2 } # => [2, 4, 6]
```

You can replace a block with a symbol only when calling a single method within the block. Symbols have the to_proc method (Symbol#to_proc) implemented.

```ruby
# Normal form
numbers = [1, 2, 3, 4, 5]
even_numbers = numbers.select { |n| n.even? }
p even_numbers.inspect # => [2, 4]

# Using a symbol
numbers = [1, 2, 3, 4, 5]
even_numbers = numbers.select(&:even?)
p even_numbers.inspect # => [2, 4]

# Using a symbol with a custom method
module Numbers
  def my_even?
    self % 2 == 0
  end
end

class Integer
  include Numbers
end

puts [1, 2, 3].select(&:my_even?)  # => [2]
```

As for the distinction between `do..end` and `{}`, consider the following:

- Use `do..end` for multiple lines
- Use `{}` for single lines
- Use `{}` when used as a return value
- Use `do..end` when passing procedures

Since this can also depend on coding conventions, it's best to consider this as just one example.

# References
- [docs.ruby-lang.org - Method Calls (super, with blocks, yield)](https://docs.ruby-lang.org/ja/3.3/doc/spec=2fcall.html)
- [style.potepan.com - How to Use Block Syntax in Ruby (do..end)](https://style.potepan.com/articles/27709.html)
- [obelisk.hatenablog.com - Ruby Blocks are Closures](https://obelisk.hatenablog.com/entry/2017/12/26/143924)
- [qiita.com - If You Don't Fully Understand Block Processing in Ruby, Take a Look Without Getting Angry](https://qiita.com/sakamiya_tomoki/items/10b18541707f800055bb)
- [zenn.dev - 【Ruby】 Got Stuck on the Difference Between do...end and {...}](https://zenn.dev/yoshiyoshiharu/articles/1b7795f593c62a)
- [mickey24.hatenablog.com - Distinguishing Ruby Block Syntax (do end, {})](https://mickey24.hatenablog.com/entry/20100914/1284475769)
- [Understanding Ruby Blocks Properly](https://zenn.dev/hayaokimura/articles/ruby-block-reintroduction)
- [active.nikkeibp.co.jp - Yukihiro Matsumoto, Ruby Developer, Deeply Explains the Concept of Block Syntax](https://active.nikkeibp.co.jp/atcl/act/19/00484/090100004/)
- [musclecoding.com - A Detailed Explanation of Ruby Blocks for Beginners! What are Block Variables?](https://musclecoding.com/ruby-block/)