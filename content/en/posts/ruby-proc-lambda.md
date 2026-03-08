---
title: About Ruby's Proc and Lambda
slug: ruby-proc-lambda
date: 2024-05-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
translation_key: ruby-proc-lambda
---

# Overview
This post discusses Ruby's Proc and lambda.

# What is Proc?
A Proc is an objectified block. A block itself is not an object.

This object is called a procedure object and is represented as an instance of the Proc class.

You can create a procedure object using [Proc.new](https://docs.ruby-lang.org/ja/latest/method/Proc/s/new.html) or [proc](https://docs.ruby-lang.org/ja/latest/method/Kernel/m/proc.html).

```ruby
#  Proc.new
proc_object = Proc.new { |x| puts x }
proc_object.call(1) # => 1

# proc
proc_object = proc{ |x| puts x }
proc_object.call(1) # => 1
```

# What is Lambda?
A lambda is a syntax for creating a Proc object, generating a procedure object.

You can create a procedure object using lambda or ->{}.

```ruby
# lambda
lamda_object = lambda { |x| puts x }
lamda_object.call(1) # => 1

# ->{}
lamda_object = ->(x) { puts x }
lamda_object.call(1) # => 1
```

# Differences Between Proc and Lambda
## Argument Handling
Proc is lenient, while lambda is strict.

```ruby
# Proc
proc_object = Proc.new { |x, y| puts "#{x}, #{y.inspect}" }
proc_object.call(10) #=> 10, nil

# lambda
lamda_object = lambda { |x, y| puts "#{x}, #{y.inspect}" }
lamda_object.call(10) #=> wrong number of arguments (given 1, expected 2) (ArgumentError)
```

## Behavior of Jump Statements (return, next, break)
There are differences as follows:

|          |           return           |            next            |           break            |
| -------- | -------------------------- | -------------------------- | -------------------------- |
| Proc.new | Exits the method           | Exits the procedure object | Raises an exception        |
| proc     | Exits the method           | Exits the procedure object | Raises an exception        |
| lambda   | Exits the procedure object | Exits the procedure object | Exits the procedure object |

```ruby
# return
def proc_return
  # Exits the method here
  p = Proc.new { return "return from proc" }
  result = p.call
  return "#{result} method finished"
end

def proc_return_with_proc
  # Exits the method here
  p = proc { return "return from proc with proc" }
  result = p.call
  return "#{result} method finished"
end

def lambda_return
  l = lambda { return "return from lambda" }
  result = l.call
  return "#{result} method finished"
  # Exits the method here
end

puts proc_return # return from proc
puts proc_return_with_proc # return from proc with proc
puts lambda_return # return from lambda method finished

# next
def proc_next
  p = Proc.new { return "return from proc" }
  result = p.call
  return "return from method"
  # Exits the method here
end

def proc_next_with_proc
  p = Proc.new { return "return from proc with proc" }
  result = p.call
  return "return from method"
  # Exits the method here
end

def lamda_next
  l = lambda { return "return from lambda" }
  result = l.call
  return "return from method"
  # Exits the method here
end

puts proc_next # => return from proc
puts proc_next_with_proc # => return from proc with proc
puts lamda_next # => return from method

# break
def proc_break
  p = Proc.new { break "Proc.new break" } # => LocalJumpError: break from proc-closure
  result = p.call
  puts "result: #{result}"
end

def proc_break_with_proc
  proc_object = proc { break "proc break" } # => LocalJumpError: break from proc-closure
  result = proc_object.call
  puts "result: #{result}"
end

def lamda_break
  l = lambda { break "lambda break" }
  result = l.call
  puts "result: #{result}"
end

puts proc_break # => LocalJumpError: break from proc-closure
puts proc_break_with_proc # => LocalJumpError: break from proc-closure
puts lamda_break # => result: lamda break
```

# Historical Background
Quoted from [jp.quora.com - Is there an essential difference between Ruby's blocks and Proc?](https://jp.quora.com/ruby-no-burokku-to-Proc-ni-honshitsu-teki-na-chigai-ha-arima-suka).

> As a supplement, I will describe the historical background: blocks (and calls via yield) were introduced first, and then Proc.new was created to objectify them. The lambda method (and proc method) were introduced to achieve anonymous functions.

> Later, with the introduction of block arguments (& arguments), the necessity of Proc.new diminished, and with the introduction of lambda expressions (->), the necessity of lambda and proc methods also diminished, leading to their current infrequent use. Perhaps these methods will be removed in the future. It reflects the changes in Ruby's history.

Although this is a 5-year-old answer, it seems that using lambda expressions generally avoids most issues.

# Behavior of Orphaned Objects
There are also differences in cases where procedure objects are used outside of methods.

|          |           return           |            next            |           break            |
| -------- | -------------------------- | -------------------------- | -------------------------- |
| Proc.new | Raises an exception        | Exits the procedure object | Raises an exception        |
| proc     | Raises an exception        | Exits the procedure object | Raises an exception        |
| lambda   | Exits the procedure object | Exits the procedure object | Exits the procedure object |

```ruby
def orphan_proc
  orphan = Proc.new { return "I'm an orphan!" }
  return orphan
end

def orphan_proc_with_proc
  orphan = proc { return "I'm an orphan!" }
  return orphan
end

def orphan_lamda
  orphan = lambda { return "I'm an orphan!" }
  return orphan
end

puts orphan_proc.call # => LocalJumpError
puts orphan_proc_with_proc.call # => LocalJumpError
puts orphan_lamda.call # => "I'm an orphan
```

# References
- [docs.ruby-lang.org - Details on Procedure Object Behavior](https://docs.ruby-lang.org/ja/latest/doc/spec=2flambda_proc.html)
- [docs.ruby-lang.org - module function Kernel.#lambda](https://docs.ruby-lang.org/ja/latest/method/Kernel/m/lambda.html)
- [jp.quora.com - Is there an essential difference between Ruby's blocks and Proc?](https://jp.quora.com/ruby-no-burokku-to-Proc-ni-honshitsu-teki-na-chigai-ha-arima-suka)
- [qiita.com - When to Use Ruby block/proc/lambda](https://qiita.com/kidach1/items/15cfee9ec66804c3afd2)
- [qiita.com - Differences in Behavior of Proc.new and Lambda When Using return or break](https://qiita.com/jnchito/items/83410c0cda446efea582#break%E3%81%8C%E5%91%BC%E3%81%B0%E3%82%8C%E3%81%9F%E5%A0%B4%E5%90%88-1)