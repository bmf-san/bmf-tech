---
title: "Golangのポインタ概要"
slug: "golang-pointers-overview"
date: 2018-11-13
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
draft: false
---

# 概要
[Golangのポインタの基本](https://bmf-tech.com/posts/Golang%E3%81%AE%E3%83%9D%E3%82%A4%E3%83%B3%E3%82%BF%E3%81%AE%E5%9F%BA%E6%9C%AC)では参照渡しと値渡しの違いの観点からポインタについて書いたが、それ以前にポインタを扱っているうちに混乱してきたため、ポインタの概要をまとめる。

# ポインタ型とは
- メモリー上のアドレスを記憶する変数の型のこと

# ポインタ型の定義
- \*を使って定義する

```
var s *string // 変数sはstring型のポインタ。型は*string
```

# アドレス演算子と間接参照演算子
## アドレス演算子　&
- **任意の値からポインタ型を生成し**、アドレスを得る

## 間接参照演算子　*
- ポインタ型変数から値を参照する（デリファレンス）

# ポインタ型変数
```golang
package main

import "fmt"

func main() {
  /**
   * ポインタ型の定義
   */
  var pointer *string // *string型のポインタ変数
  var s string
  s = "Hello World"

  /*
   * アドレス演算子でポインタを扱う
   */
  fmt.Printf("%T\n", pointer) // *string
  fmt.Printf("%v\n", pointer) // <nil>　※初期化されていないポインタ型の値はnil

  fmt.Printf("%T\n", &pointer) // **string　※pointerは*string型のポインタ変数なので、&pointerは*stringのポインタを生成する
  fmt.Printf("%v\n", &pointer) // 0xc00000c028　※&を使ってポインタのアドレスを得る

  fmt.Printf("%T\n", s) // string
  fmt.Printf("%v\n", s) // Hello World

  pointer = &s // ※*string型のポインタ変数（値は初期化前なのでnil）にsのポインタ型を生成し、アドレスを代入。（変数pointerは*string型のポインタ変数として定義済みなので、&を使って変数sのポインタ型を生成し、代入。）
  fmt.Printf("%T\n", pointer) // *string　※pointerは*string型のポインタのまま
  fmt.Printf("%v\n", pointer) // 0xc00000e1e0
  fmt.Printf("%v\n", *pointer) // Hello World　※pointer（*string型ポインタ変数）から値を参照する
  fmt.Printf("%T\n", s) // string
  fmt.Printf("%v\n", s) // Hello World
  fmt.Printf("%T\n", &s) // *string
  fmt.Printf("%v\n", &s) // 0xc00000e1e0
  fmt.Printf("%v\n", *s) // invalid indirect of s (type string)　※sはポインタ型変数ではないためエラー。indirectは間接の意

  *pointer = "New World" // ※*pointerは変数なので代入できる。ポインタの変数定義は*を使う。
  fmt.Printf("%T\n", *pointer) // string　※*pointerはポインタ変数から値を参照している
  fmt.Printf("%v\n", *pointer) // New World
}
```

# 構造体とポインタ
## 構造体について復習

```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

func main() {
  var h Human

  fmt.Printf("%v\n", h.Name) // ""　※stringのゼロ値
  fmt.Printf("%v\n", h.Age) // 0　※intのゼロ値

  h = Human{
    Name: "Tom",
    Age: 20,
  }

  fmt.Printf("%v\n", h) // {Tom 20}
}
```

## 構造体とポインタ
## compositeリテラルによる構造体の初期化
```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

func main() {
  var h *Human // ポインタ型変数の定義

  h = &Human{ // ※変数hは*Human型のポインタ型なので、&を使って構造体Humanから*Human型のポインタ型を使って代入。
    Name: "John",
    Age: 20,
  }

  fmt.Printf("%T\n", h) // *main.Human　※&HumanでHumanのポインタ型を生成している
  fmt.Printf("%v\n", h) // &{John 20}　※*main.Human型のポインタ
  fmt.Printf("%v\n", *h) // {John 20}　※値の取り出し

  h.Name = "Tom"
  h.Age = 40

  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // &{Tom 40}
  fmt.Printf("%v\n", *h) // {John 40}　※値の取り出し
}
```

## newで構造体を初期化
```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

func main() {
  h := new(Human)

  h.Name = "John"
  h.Age = 20

  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // &{John 20}

  h.Name = "Tom"
  h.Age = 40
  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // &{Tom 40}
  fmt.Printf("%v\n", *h) // {Tom 40}　※値を取り出した
}
```

# レシーバとポインタ
## 値レシーバ
```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

func (h Human) say(msg string) {
  fmt.Printf("%v(%v) said %v\n", h.Name, h.Age, msg)
}

func main() {
  var h Human // Human型の変数hを定義

  fmt.Printf("%T\n", h) // main.Human
  fmt.Printf("%v\n", h) // { 0}　※NameとAgeのゼロ値

  h = Human{
    Name: "Taro",
    Age: 20,
  }

  fmt.Printf("%T\n", h) // main.Human
  fmt.Printf("%v\n", h) // {Taro 20}

  h.say("Hello") // Taro(20) said Hello
}
```

## ポインタレシーバ
```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

// ポインタレシーバ
func (h *Human) say(msg string) {
  fmt.Printf("%v(%v) said %v\n", h.Name, h.Age, msg)
}

func main() {
  var h *Human // ※ポインタ型の定義

  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // <nil> ※ポインタ型のゼロ値

  h = &Human{
    Name: "Taro",
    Age: 20,
  }

  fmt.Printf("%T\n", h) // *main.Human ※&Human構造体のポインタを生成し、hに代入済み
  fmt.Printf("%v\n", h) // &{Taro 20}

  h.say("Hello") // Taro(20) said Hello
}
```

## 値レシーバとポインタレシーバの違い
```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

func (h Human) setDataForValue(name string, age int) {
  h.Name = name
  h.Age = age
}

func (h *Human) setDataForPointer(name string, age int) {
  h.Name = name
  h.Age = age
}

func main() {
  // 値関数の呼び出し
  var hForValue Human

  hForValue = Human{
    Name: "Taro",
    Age: 20,
  }

  hForValue.setDataForValue("Jiro", 40)

  fmt.Printf("Name: %v\n Age: %v\n", hForValue.Name, hForValue.Age) // Name: Taro Age: 20

  // ポインタ関数の呼び出し
  var hForPointer *Human

  hForPointer = &Human{
    Name: "Taro",
    Age: 20,
  }

  hForPointer.setDataForPointer("Jiro", 40)

  fmt.Printf("Name: %v\n Age: %v\n", hForPointer.Name, hForPointer.Age) // Name: Jiro Age: 40
}
```

関数内で構造体のフィールドの値を変更したいときはポインターレシーバを使う。
mapやchanのような参照型をレシーバーに扱う場合は値レシーバでも良い。
ただし、厳密にパフォーマンスを考慮したりする場合には使い分けはこの限りではない。

# ポインタで混乱したら
前半に記述したポインタ型の定義、演算子の役割について思い出す。

# 参考
- [Goのポインタ](http://cuto.unirita.co.jp/gostudy/post/pointer/)
- [Qiita - Goで学ぶポインタとアドレス ](https://qiita.com/Sekky0905/items/447efa04a95e3fec217f)
- [Qiita - Go言語におけるポインタのアンパサンド（＆）とアスタリスクの（＊）の違い](https://qiita.com/tmzkysk/items/1b73eaf415fee91aaad3)
- [Qita - Goの構造体の研究](https://qiita.com/kmtr/items/1e7caf92aa8bb587906d)
- [[Go] 構造体で、値メソッドとポインタメソッドを使い分ける](https://www.yoheim.net/blog.php?q=20170902)
- [](https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/)
- [Github golang/tour - tour: How do i print an address of the pointer to "Struct fields" ? #226](https://github.com/golang/tour/issues/226)
- [StackOverflow - Memory address for struct not visible](https://stackoverflow.com/questions/29498374/memory-address-for-struct-not-visible)
