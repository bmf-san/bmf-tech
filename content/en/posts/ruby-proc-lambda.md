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
Writing about Ruby's Proc and lambda.

# What is Proc
A block that has been turned into an object. A block is not an object.

This object is called a procedure object and is represented as an instance of the Proc class.

You can generate a procedure object using [Proc.new](https://docs.ruby-lang.org/ja/latest/method/Proc/s/new.html) or [proc](https://docs.ruby-lang.org/ja/latest/method/Kernel/m/proc.html).

```ruby
#  Proc.new
proc_object = Proc.new { |x| puts x }
proc_object.call(1) # => 1

# proc
proc_object = proc{ |x| puts x }
proc_object.call(1) # => 1
```

# What is lambda
A syntax for creating a Proc object, generating a procedure object.

You can generate a procedure object using lambda or ->{}.

```ruby
# lambda
lambda_object = lambda { |x| puts x }
lambda_object.call(1) # => 1

# ->{}
lambda_object = ->(x) { puts x }
lambda_object.call(1) # => 1
```

# Differences between Proc and lambda
## Different handling of arguments
Proc is lenient, but lambda is strict.

```ruby
# Proc
proc_object = Proc.new { |x, y| puts "#{x}, #{y.inspect}" }
proc_object.call(10) #=> 10, nil

# lambda
lambda_object = lambda { |x, y| puts "#{x}, #{y.inspect}" }
lambda_object.call(10) #=> wrong number of arguments (given 1, expected 2) (ArgumentError)
```

## Behavior of jump statements (return, next, break)
There are differences as follows.

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

def lambda_next
  l = lambda { return "return from lambda" }
  result = l.call
  return "return from method"
  # Exits the method here
end

puts proc_next # => return from proc
puts proc_next_with_proc # => return from proc with proc
puts lambda_next # => return from method

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

def lambda_break
  l = lambda { break "lambda break" }
  result = l.call
  puts "result: #{result}"
end

puts proc_break # => LocalJumpError: break from proc-closure
puts proc_break_with_proc # => LocalJumpError: break from proc-closure
puts lambda_break # => result: lambda break
```

# Historical Background
Quoted from [jp.quora.com - Are there any essential differences between Ruby's block and Proc?](https://jp.quora.com/ruby-no-burokku-to-Proc-ni-honshitsu-teki-na-chigai-ha-arima-suka).

> As a supplement, to describe the historical background, blocks (and calls via yield) were introduced first, and then Proc.new was born to objectify them. Also, the lambda method (and proc method) was introduced to achieve the equivalent of anonymous functions.

> Later, with the introduction of block arguments (& arguments), the necessity of Proc.new diminished, and with the introduction of lambda expressions (->), the necessity of lambda and proc methods diminished, and they are now rarely used. Probably, these methods will be removed in the future. It feels like a change due to the history of Ruby.

Although this is a response from 5 years ago, it seems that using lambda expressions generally poses no problem.

# Behavior of Orphaned Objects
There are differences in cases where a procedure object is used outside of a method.

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

def orphan_lambda
  orphan = lambda { return "I'm an orphan!" }
  return orphan
end

puts orphan_proc.call # => LocalJumpError
puts orphan_proc_with_proc.call # => LocalJumpError
puts orphan_lambda.call # => "I'm an orphan"
```

# References
- [docs.ruby-lang.org - Details of procedure object behavior](https://docs.ruby-lang.org/ja/latest/doc/spec=2flambda_proc.html)
- [docs.ruby-lang.org - module function Kernel.#lambda](https://docs.ruby-lang.org/ja/latest/method/Kernel/m/lambda.html)
- [jp.quora.com - Are there any essential differences between Ruby's block and Proc?](https://jp.quora.com/ruby-no-burokku-to-Proc-ni-honshitsu-teki-na-chigai-ha-arima-suka)
- [qiita.com - When to use Ruby block/proc/lambda](https://qiita.com/kidach1/items/15cfee9ec66804c3afd2)
- [qiita.com - Differences in behavior of Proc.new and lambda when using return or break](https://qiita.com/jnchito/items/83410c0cda446efea582#break%E3%81%8C%E5%91%BC%E3%81%B0%E3%82%8C%E3%81%9F%E5%A0%B4%E5%90%88-1)