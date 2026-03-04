---
title: "RubyのProcとlamdaについて"
slug: "ruby-proc-lambda"
date: 2024-05-15
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Ruby"
draft: false
---

# 概要
RubyのProcとlamdaについてかく。

# Procとは
ブロックをオブジェクト化したもの。ブロックはオブジェクトではない。

このオブジェクトは手続きオブジェクトと呼ばれ、Procクラスのインスタンスとして表現される。

[Proc.new](https://docs.ruby-lang.org/ja/latest/method/Proc/s/new.html)または[proc](https://docs.ruby-lang.org/ja/latest/method/Kernel/m/proc.html)によって手続きオブジェクトを生成できる。

```ruby
#  Proc.new
proc_object = Proc.new { |x| puts x }
proc_object.call(1) # => 1

# proc
proc_object = proc{ |x| puts x }
proc_object.call(1) # => 1
```

# lamdaとは
Procオブジェクトを作成するための構文で、手続きオブジェクトを生成する。

lamdaまたは->{}によって手続きオブジェクトを生成できる。

```ruby
# lamda
lamda_object = lambda { |x| puts x }
lamda_object.call(1) # => 1

# ->{}
lamda_object = ->(x) { puts x }
lamda_object.call(1) # => 1
```

# Procとlamdaの違い
## 引数の扱いが異なる
Procの場合は緩いが、lamdaの場合は厳格である。

```ruby
# Proc
proc_object = Proc.new { |x, y| puts "#{x}, #{y.inspect}" }
proc_object.call(10) #=> 10, nil

# lamda
lamda_object = lambda { |x, y| puts "#{x}, #{y.inspect}" }
lamda_object.call(10) #=> wrong number of arguments (given 1, expected 2) (ArgumentError)
```

## ジャンプ構文（return、next、break）の挙動
次のような違いがある。

|          |           return           |            next            |           break            |
| -------- | -------------------------- | -------------------------- | -------------------------- |
| Proc.new | メソッドを抜ける           | 手続きオブジェクトを抜ける | 例外が発生する             |
| proc     | メソッドを抜ける           | 手続きオブジェクトを抜ける | 例外が発生する             |
| lambda   | 手続きオブジェクトを抜ける | 手続きオブジェクトを抜ける | 手続きオブジェクトを抜ける |

```ruby
# return
def proc_return
  # ここでメソッドを抜ける
  p = Proc.new { return "return from proc" }
  result = p.call
  return "#{result} method finished"
end

def proc_return_with_proc
  # ここでメソッドを抜ける
  p = proc { return "return from proc with proc" }
  result = p.call
  return "#{result} method finished"
end

def lambda_return
  l = lambda { return "return from lambda" }
  result = l.call
  return "#{result} method finished"
  # ここでメソッドを抜ける
end

puts proc_return # return from proc
puts proc_return_with_proc # return from proc with proc
puts lambda_return # return from lambda method finished

# next
def proc_next
  p = Proc.new { return "return from proc" }
  result = p.call
  return "return from method"
  # ここでメソッドを抜ける
end

def proc_next_with_proc
  p = Proc.new { return "return from proc with proc" }
  result = p.call
  return "return from method"
  # ここでメソッドを抜ける
end

def lamda_next
  l = lambda { return "return from lambda" }
  result = l.call
  return "return from method"
  # ここでメソッドを抜ける
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


# 歴史的経緯
[jp.quora.com - rubyのブロックとProcに本質的な違いはありますか？](https://jp.quora.com/ruby-no-burokku-to-Proc-ni-honshitsu-teki-na-chigai-ha-arima-suka)より引用。

> 補足として、歴史的な経緯を記述しておくと、ブロック(とyieldによる呼び出し)がまず最初に導入され、それから、それをオブジェクトするために Proc.new が誕生しました。また、無名関数相当を実現するために lambda メソッド(と proc メソッド)が導入されました。

> その後、ブロック引数(&引数)が導入されることで Proc.new の必要性が薄れ、lambda式(->)が導入されることで、lambdaやprocメソッドの必要性が薄れ、それらは現在ではほとんど使われなくなっています。たぶん、将来これらのメソッドは削除されるでしょう。Rubyの歴史による変化を感じさせますね。

5年前の回答ではあるが、lamda式を使っておけば大抵問題ないということだろうか。

# 孤児的なオブジェクトの挙動
手続きオブジェクトがメソッドの外で使用されるようなケースにおいても違いがある。

|          |           return           |            next            |           break            |
| -------- | -------------------------- | -------------------------- | -------------------------- |
| Proc.new | 例外が発生する             | 手続きオブジェクトを抜ける | 例外が発生する             |
| proc     | 例外が発生する             | 手続きオブジェクトを抜ける | 例外が発生する             |
| lambda   | 手続きオブジェクトを抜ける | 手続きオブジェクトを抜ける | 手続きオブジェクトを抜ける |

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

# 参考
- [docs.ruby-lang.org - 手続きオブジェクトの挙動の詳細](https://docs.ruby-lang.org/ja/latest/doc/spec=2flambda_proc.html)
- [docs.ruby-lang.org - module function Kernel.#lambda](https://docs.ruby-lang.org/ja/latest/method/Kernel/m/lambda.html)
- [jp.quora.com - rubyのブロックとProcに本質的な違いはありますか？](https://jp.quora.com/ruby-no-burokku-to-Proc-ni-honshitsu-teki-na-chigai-ha-arima-suka)
- [qiita.com - Ruby block/proc/lambdaの使いどころ](https://qiita.com/kidach1/items/15cfee9ec66804c3afd2)
- [qiita.com - returnやbreakを使ったときのProc.newとラムダの挙動の違い](https://qiita.com/jnchito/items/83410c0cda446efea582#break%E3%81%8C%E5%91%BC%E3%81%B0%E3%82%8C%E3%81%9F%E5%A0%B4%E5%90%88-1)
