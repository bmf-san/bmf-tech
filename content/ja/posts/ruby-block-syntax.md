---
title: Rubyのブロック構文について
description: "理解するRubyブロック構文。do..end形式、{}形式、yield、block_given?、&blockパラメータ、クロージャの特性を習得。"
slug: ruby-block-syntax
date: 2024-05-15T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
translation_key: ruby-block-syntax
---


# 概要
Rubyのブロック構文についてかく。

# ブロック構文とは
他メソッドにコード断片を渡すことで、そのメソッド内で処理を実行させることができる構文。

`do..end`または`{}`を使って定義する。ブロックは引数を受け取ることができる。

```ruby
# do..end
[1...3].each do |number|
  puts number
end

# {}
[1...3].each { |number| puts number}
```

ブロックを受け取るメソッドを定義することもできる。

```ruby
def greet
  puts "Hi"
  # 渡されたブロックを実行
  yield if block_given? # block_given?はブロックが渡されたかどうかを判定するメソッド
  puts "Bye"
end

# ブロックを渡す
greet do
  puts "Hello"
end

# 出力:
# Hi
# Hello
# Bye
```

また、ブロックをメソッドの引数として渡すこともできる。その場合は`&`を使って渡す。

```ruby
def greet(&block)
  puts "Hi"
  block.call if block
  puts "Bye"
end

greet do
  puts "Hello"
end

# 出力：
# Hi
# Hello
# Bye
```

to_procメソッドを持つオブジェクトもブロックとして渡すことができる。

```ruby
class Greeting
  def to_proc
    Proc.new { |name| puts "Hello, #{name}!" }
  end
end

%w(Alice Bob Charlie).each(&Greeting.new)

# 出力:
# Hello, Alice!
# Hello, Bob!
# Hello, Charlie!
```

ブロックはProcやlamdaと同じくクロージャーである。

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

`do..end`と`{}`は結合度が異なり、`{}`の方が結合度が高い。

```ruby
array = [1, 2, 3]

# 結合度が低いのでブロックがmapに渡される前に評価される（ppに渡されてしまう）
pp array.map do |x|
  x * 2
end

# 出力：
#<Enumerator: ...>

# 結合度が高いのでブロックがmapに渡されて評価される
pp array.map { |x| x * 2 } # => [2, 4, 6]
```

ブロックの中で単一のメソッドを呼び出す場合に限り、ブロックをシンボルで置き換えることができる。シンボルにはto_procメソッド（Symbol#to_proc）が実装されている。

```ruby
# 通常の形
numbers = [1, 2, 3, 4, 5]
even_numbers = numbers.select { |n| n.even? }
p even_numbers.inspect # => [2, 4]

# シンボルを使った形
numbers = [1, 2, 3, 4, 5]
even_numbers = numbers.select(&:even?)
p even_numbers.inspect # => [2, 4]

# 自作のメソッドでシンボルを使った形
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

`do..end`と`{}`の使い分けとしては、次のような考え方がある。

- 複数行の場合は`do..end`
- 一行の場合は`{}`
- 値を返す式として使う場合は`{}`
- 手続きを渡す場合は`do..end`

コーディング規約に依る部分もあるので、あくまで一例として考えておきたい。

# 参考
- [docs.ruby-lang.org - メソッド呼び出し(super・ブロック付き・yield)](https://docs.ruby-lang.org/ja/3.3/doc/spec=2fcall.html)
- [style.potepan.com - Rubyにおけるブロック構文(do～end)の使い方](https://style.potepan.com/articles/27709.html)
- [obelisk.hatenablog.com - Ruby のブロックはクロージャである](https://obelisk.hatenablog.com/entry/2017/12/26/143924)
- [qiita.com - Rubyをやっていてブロック処理(do 〜 end)がちゃんとわかってない人は、怒らないから見ていきなさい](https://qiita.com/sakamiya_tomoki/items/10b18541707f800055bb)
- [zenn.dev - 【Ruby】 do...end と {...}の違いでハマった](https://zenn.dev/yoshiyoshiharu/articles/1b7795f593c62a)
- [mickey24.hatenablog.com - Rubyのブロック構文の書き分け(do end，{})](https://mickey24.hatenablog.com/entry/20100914/1284475769)
- [Ruby のブロックをちゃんと理解する](https://zenn.dev/hayaokimura/articles/ruby-block-reintroduction)
- [active.nikkeibp.co.jp - Ruby開発者のまつもとゆきひろが深く解説、「ブロック構文」発想の経緯](https://active.nikkeibp.co.jp/atcl/act/19/00484/090100004/)
- ~~musclecoding.com -【Ruby】「ブロック」を初心者向けにとにかく丁寧に解説！ブロック変数とは？~~
