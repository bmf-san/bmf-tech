---
title: Golangでの変数定義・宣言のパターンまとめ
description: "Goの変数定義・宣言パターンをvar宣言・短縮形式別に整理し、型推論・複数変数宣言・可視性制御の規則を網羅。"
slug: golang-variable-definition-patterns
date: 2018-11-13T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
translation_key: golang-variable-definition-patterns
---


# 概要
Golangでの変数定義・宣言のパターンをまとめる

# 変数の定義・宣言における注意
- 1文字目が大文字の場合は、他のパッケージからも見える変数
- 1文字目が小文字の場合は、そのパッケージでしか見えない変数

# 変数の定義・宣言
## Variable declarations
```golang
var i int
fmt.Printf("%T", i) // int
```

```golang
var a, b, c string
fmt.Printf("%T", a) // string
fmt.Printf("%T", b) // string
fmt.Printf("%T", c) // string
```

```golang
var s = "Hello World"
fmt.Printf("%T", s) // string
```

```golang
var x, y, z int = 1, 2, 3
fmt.Printf("%T", x) // int
fmt.Printf("%T", y) // int
fmt.Printf("%T", z) // int
```

``` golang
var (
   a string
   x int
   y, z = 2, 3
)
fmt.Printf("%T", a) // string
fmt.Printf("%T", x) // int
fmt.Printf("%T", y) // int
fmt.Printf("%T", z) // int
```

```golang
var i, j = 1, 2
fmt.Printf("%T", i) // int
fmt.Printf("%T", j) // int
```

## Short variable declarations
- 省略形式の変数宣言は関数内でのみ使用可能

```golang
i := 1
fmt.Printf("%T", i) // int
```

```golang
i, j := 1, 2
fmt.Printf("%T", i) // int
fmt.Printf("%T", j) // int
```

# 参考
- [Qiita   - 他言語プログラマがgolangの基本を押さえる為のまとめ](https://qiita.com/tfrcm/items/e2a3d7ce7ab8868e37f7)
- [Golang.org - Variable Declarations](https://golang.org/ref/spec#Variable_declarations)

