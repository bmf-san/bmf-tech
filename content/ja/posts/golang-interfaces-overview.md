---
title: "Goのインターフェースを解説：柔軟でテスタブルなコードの書き方"
slug: golang-interfaces-overview
date: 2018-11-15T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
translation_key: golang-interfaces-overview
---


# 概要
Golangのインターフェースについてまとめる。

# Golangのインターフェースとは
- 特定のメソッドの型だけを列挙した型
- インターフェースFooで宣言したメソッドが全て実装されている構造体は、型Fooとして扱うことができる。
- インタフェースを使うとポリモーフィズムを実現することができる

# インターフェースの定義
```golang
type <型名> interface {
  <メソッド名(<引数の型>, ...)(<戻り値の型>, ...)
}
```

```golang
// Ex.
type Human interface {
  say() string
}
```

# インターフェースの特徴
## インターフェース型の変数
interface型で宣言された変数はどんな型の値でも代入ができる。

```golang
var i interface{}

i = 123
i = "Hello World"
i = []int{1, 2, 3} // etc...
```

## インターフェース型の引数
interface型を引数にすると、どんな型の値でも渡すことができる。

```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

func printType(i interface{}) {
  fmt.Printf("%T\n", i)
}

func main() {
  h := Human{
    Name: "John",
    Age: 20,
  }

  printType(h) // main.Human
}
```

# 型アサーション

型アサーションの構文。
```golang
<変数>.(<型>)
```

使い方は変数２つを取るような形で使う。
```golang
s, ok := i.(Human)
```

変数iがHuman型であった場合は、変数sはHuman型の変数iの実際の値が、変数okにはtrueが格納される。
逆に、変数iがHuman型ではない場合は、変数sにはHuman型のゼロ値が格納される。

```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

type Alien struct {
  Name string
  Age int
}

func printOnlyHuman(i interface{}) {
  s, ok := i.(Human)

  if !ok {
    fmt.Printf("%v\n", "Human型ではない")
    fmt.Printf("%v\n", s)
    return
  }

  fmt.Printf("%v\n", "Human型である")
  fmt.Printf("%v\n", s)
}

func main() {
  h := Human{
    Name: "John",
    Age: 20,
  }

  a := Alien {
    Name: "Tom",
    Age: 200000,
  }

  printOnlyHuman(h) // Human型である。{John 20}
  printOnlyHuman(a) // Human型ではない。{ 0}
}
```

# インターフェースの実装例
Golangのインターフェースのポピュラーな使用方法である、「異なる型に共通の性質を付与する」使い方の例。

```golang
package main

import "fmt"

type Action interface {
  say()
}

type Human struct {}

type Alien struct {}

func (h *Human) say() {
  fmt.Println("I'm Human")
}

func (a *Alien) say() {
  fmt.Println("I'm Alien")
}

func do(a Action) { // Action型を受け取る
  a.say()
}

func main() {
  ha := []Action{
    &Human{},
    &Alien{},
  }

  for _, v := range ha {
    do(v)
  }
}
```

# 参考
- [golang.org - Interface types](https://golang.org/ref/spec#Interface_types)
- [Go言語 - 空インターフェースと型アサーション](https://blog.y-yuki.net/entry/2017/05/08/000000)
- [はじめてのGo言語 - インタフェース](http://cuto.unirita.co.jp/gostudy/post/interface/)
- [お気楽 Go 言語プログラミング入門](http://www.geocities.jp/m_hiroi/golang/abcgo09.html)
- [
型のキャストと、型アサーションによる型変換](https://maku77.github.io/hugo/go/cast.html)
- [Go 言語のインタフェースの扱いを理解する](http://maku77.github.io/hugo/go/interface.html)
- [astaxie/build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/02.6.md)
- [SE Book - スターティングGo言語](https://www.shoeisha.co.jp/book/detail/9784798142418)

