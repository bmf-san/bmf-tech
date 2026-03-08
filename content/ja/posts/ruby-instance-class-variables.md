---
title: Rubyのインスタンス変数とクラス変数とクラスインスタンス変数の違い
slug: ruby-instance-class-variables
date: 2025-01-31T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Ruby
translation_key: ruby-instance-class-variables
---


Rubyのインスタンス変数とクラス変数とクラスインスタンス変数の違いについてまとめる。

# インスタンス変数
インスタンスごとに異なる値を保持する変数で、`@`で始まる変数名を使う。

主にクラス内部で宣言されたオブジェクトごとの状態を管理するために使われる。

```ruby
class Person
  def initialize(name)
    @name = name # インスタンス変数
  end

  def name
    @name
  end
end

p1 = Person.new("Alice")
p2 = Person.new("Bob")

puts p1.name # => Alice
puts p2.name # => Bob 別のインスタンスなので別の値を持つ
```

# クラス変数
クラス全体で共有される変数で、`@@`で始まる変数名を使う。

同じクラスの異なるインスタンス間で共有される値を保持するために使われる。

```ruby
class Person
  @@count = 0 # クラス変数

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

puts Person.count # => 2 すべてのインスタンスで共有される
```

サブクラスでも共有される。

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
puts Parent.var # => "Child" # サブクラスで上書きされる
```

# クラスインスタンス変数
クラス変数とは異なり、サブクラスとは共有されない。

クラス自身のインスタンス変数として使われる。

selfを使ってクラスメソッドの中でアクセスする。

```ruby
class Person
  @count = 0 # クラスインスタンス変数

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

puts Person.count # => 1 Personクラスのカウント
puts Student.count # => 1 Studentクラスのカウント
```

# 違いのまとめ
| 変数の種類                                           | 宣言場所                              | スコープ                 | サブクラスへの影響             | 用途                             |
| ---------------------------------------------------- | ------------------------------------- | ------------------------ | ------------------------------ | -------------------------------- |
| **インスタンス変数 (`@var`)**                        | `initialize` やインスタンスメソッド内 | **各インスタンスごと**   | 影響しない                     | インスタンスごとのデータ管理     |
| **クラス変数 (`@@var`)**                             | クラス内                              | **全インスタンスで共有** | **影響する（共有される）**     | 全インスタンスで共通のデータ管理 |
| **クラスインスタンス変数 (`@var` in class context)** | クラスのトップレベル (`self.@var`)    | **クラスごと**           | 影響しない（クラスごとに独立） | クラスごとのデータ管理           |
