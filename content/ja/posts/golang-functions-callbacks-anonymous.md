---
title: "Golangの関数についてー関数値／コールバック関数／無名関数"
slug: "golang-functions-callbacks-anonymous"
date: 2018-10-04
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "コールバック"
  - "関数"
  - "無名関数"
draft: false
---

# 概要
Golangの関数において、以下3つについてまとめる。
- 関数値として扱う場合の関数
- 関数を引数として扱う関数
- 無名関数の定義
    - 関数値
    - クロージャー

# 関数値として扱う場合の関数
```golang
package main

import (
  "fmt"
  "testing"
)

func sayHi() string {
  return "Hello"
}

func main() {
  greetA := sayHi()
  greetB := sayHi

  fmt.Println(greetA)
  fmt.Println(greetB())
}
```

# コールバック関数
```golang
package main

import "fmt"

// コールバック関数
func add(n int) int {
    return n
}

func sum(v int, r func(int) int) int {
  return r(v)
}

func main() {
  fmt.Println(sum(1, add))
}
```

関数`sum`は2つの引数を定義している。
- 型がintのv
- 型が、**intを引数として定義しており、戻り値がintである関数のr**
  - コールバック関数を引数として定義する場合は、引数であっても戻り値を記述する必要がある。

ちなみにmain関数内で実行されているaddにはアドレスが格納されている。

```golang
fmt.Println("%v", add) // 0x10936d0
```

PHPでは可変変数を利用したり、call_user_funcを使ってコールバックを実現していた。

# 無名関数の定義
## 関数値
無名関数を関数値として扱う場合の例

```golang
package main

import "fmt"

func main() {
  sum := func (n int) int {
    return n + 1
  }

  fmt.Println(sum(1))
}
```

## クロージャー
無名関数としてクロージャーとして定義する場合の例

```golang
package main

import "fmt"

func count() func() int {
  var count int
  return func() int {
    count++
    return count
  }
}

func main() {
  countUp := count()
  fmt.Println(countUp()) // 1
  fmt.Println(countUp()) // 2
  fmt.Println(countUp()) // 3
}
```

クロージャーを使うとスコープ範囲がオープンになるので、countの値をキープできる。


# 所感
雰囲気でコールバック関数を使っている節があるので、コールバック関数の仕組みを深掘りしたいと思った。

