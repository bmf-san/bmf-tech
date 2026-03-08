---
title: About Go's rune Type
slug: go-rune-type-explained
date: 2024-02-21T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
translation_key: go-rune-type-explained
---

# What is rune Type
An alias type for int32, representing Unicode code points.

Unicode is a standard for character encoding that defines character sets and encoding schemes, assigning numbers to the diverse languages, formats, and symbols of the world.

A character set is a collection of characters and symbols that are gathered in a way that avoids duplication for use on computers, referring to a correspondence table of characters and their assigned numbers. An encoding scheme refers to the data format that corresponds to the numbers assigned to characters.

For example, character set: Unicode, encoding format: UTF-8.

The numbers assigned in Unicode are called code points, and the rune type serves to represent these code points.

For instance, Unicode U+0041 (in hexadecimal 0041) is the code point for the character 'A', classified as a basic Latin character, and when represented as a rune type, it looks like this:

```go
package main

import "fmt"

func main() {
	var a rune = 'A'             // Single quotes indicate rune type
	fmt.Printf("%T %U %#v\n", a, a) // int32 U+0041 65
}
```

The number 65 of type int32 is the decimal representation of 0041, representing the Unicode code point.

# string Type and rune Type
In Go, the string type is a read-only slice of bytes.

```go
package main

import "fmt"

func main() {
	s := "Hello, 世界"
	fmt.Printf("%T %v\n", s, s) // string Hello, 世界
	fmt.Println(s[:5]) // Hello
}
```

Unlike the rune type, the string type does not hold code points.

The behavior of the string type varies depending on how it is looped through.

```go
package main

import "fmt"

func main() {
	s := "Hello"
	// Can retrieve bytes
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // 72 101 108 108 111
	}

	// In a range loop, runes can be retrieved
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

When dealing with multibyte characters, it's important to pay attention to the behavior of loops.

```go
package main

import "fmt"

func main() {
	s := "あ"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) // 227 129 130
	}

	// In a range loop, runes can be retrieved
	for i, v := range s {
		fmt.Printf("idx %d: %T %U %#v\n", i, v, v, v)
		// idx 0: int32 U+3042 12354
	}
}
```

Since one Japanese character is 3 bytes, in a non-range loop, even a single character will loop 3 times.

When obtaining the number of characters, it's also important to note that the string type is a slice of bytes.

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

Since len returns the byte length, to get the number of characters, use utf8.RuneCountInString.

# References
- [go.dev - Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
- [pkg.go.dev - runes](https://pkg.go.dev/golang.org/x/text/runes)
- [ja.wikipedia.org - Character Set](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E9%9B%86%E5%90%88)
- [ja.wikipedia.org - Character Encoding Scheme](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E7%AC%A6%E5%8F%B7%E5%8C%96%E6%96%B9%E5%BC%8F)
- [ja.wikipedia.org - Character Code](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E3%82%B3%E3%83%BC%E3%83%89)
- [wa3.i-3-i.info - Difference between Character Set and Character Encoding Scheme](https://wa3.i-3-i.info/diff749moji.html)