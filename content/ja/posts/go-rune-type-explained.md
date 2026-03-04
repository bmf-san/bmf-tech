---
title: "Goのrune型について"
slug: "go-rune-type-explained"
date: 2024-02-21
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
draft: false
---

# rune型とは
int32のエイリアス型で、Unicodeのコードポイントを表す。

Unicodeは符号化文字集合や、文字符号化方式などを定めた文字コードの標準規格で、世界の多様な言語や書式、記号に番号を割り当てたもの。

符号化文字集合とは、コンピューター上で扱う文字や記号を重複しないように集めた文字セットのことであり、文字と文字に割り当てられた番号の対応表を指す。文字符号化方式は、文字に割り当てられた番号とコンピューターが扱うデータ形式を指す。

ex. 符号化文字集合：Unicode、符号化形式：UTF-8

Unicodeなどに割り当てられた番号をコードポイントと呼び、rune型はこのコードポイントを表すための型となる。

例えばUnicode U+0041（16進数では0041）は、基本ラテン文字に分類される文字'A'のコードポイントであり、rune型で表すと、次のようになる。

```go
package main

import "fmt"

func main() {
	var a rune = 'A'             // シングルクォートだとruneが型
	fmt.Printf("%T %U %#v\n", a, a) // int32 U+0041 65
}
```

int32型の65という数字は、0041の10進数表記したものであり、Unicodeのコードポイントを表している。

# string型とrune型
Goではstring型は読み取り専用のbyteのスライスである。

```go
package main

import "fmt"

func main() {
	s := "Hello, 世界"
	fmt.Printf("%T %v\n", s, s) // string Hello, 世界
	fmt.Println(s[:5]) // Hello
}
```

string型はrune型と異なり、コードポイントを保持しているわけではない。

string型はループの仕方によって挙動が異なる。

```go
package main

import "fmt"

func main() {
	s := "Hello"
	// byteが取得できる
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // 72 101 108 108 111
	}

	// range loopの場合はruneが取得できる
	for i, v := range s {
		fmt.Printf("idx %d: %T %U %#v\n", i, v, v, v)
		// idx 0: int32 U+0048 72
		// idx 1: int32 U+0065 101
		// idx 2: int32 U+006C 108
		// idx 3: int32 U+006C 108
		// idx 4: int32 U+006F 111
	}
}
```

マルチバイトを扱う場合はループの挙動に注意を払う必要がある。

```go
package main

import "fmt"

func main() {
	s := "あ"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // 227 129 130
	}

	// range loopの場合はruneが取得できる
	for i, v := range s {
		fmt.Printf("idx %d: %T %U %#v\n", i, v, v, v)
		// idx 0: int32 U+3042 12354
	}
}
```

日本語1文字は3バイトであるため、rangeを使わないループの場合は、1文字であっても3回ループが回る。

文字列数を取得するときもstring型がバイトのスライスであることに注意を払う。

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "あ"
	r := `あ`

	fmt.Println(len(s))                    // 3
	fmt.Println(len(r))                    // 3
	fmt.Println(utf8.RuneCountInString(s)) // 1
}
```

lenはバイト長を返すため、文字数を取得したい場合はutf8.RuneCountInStringを使う。

# 参考
- [go.dev - Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
- [pkg.go.dev - runes](https://pkg.go.dev/golang.org/x/text/runes)
- [ja.wikipedia.org - 文字集合](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E9%9B%86%E5%90%88)
- [ja.wikipedia.org - 文字符号化方式](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E7%AC%A6%E5%8F%B7%E5%8C%96%E6%96%B9%E5%BC%8F)
- [ja.wikipedia.org - 文字コード](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E3%82%B3%E3%83%BC%E3%83%89)
- [wa3.i-3-i.info - 「符号化文字集合」と「文字符号化方式」の違い](https://wa3.i-3-i.info/diff749moji.html)
