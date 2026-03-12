---
title: About Go's rune Type
description: 'An in-depth look at About Go''s rune Type, covering key concepts and practical insights.'
slug: go-rune-type-explained
date: 2024-02-21T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
translation_key: go-rune-type-explained
---

# What is the rune Type
The rune type is an alias for int32 and represents a Unicode code point.

Unicode is a standard character encoding that defines character sets and encoding schemes, assigning numbers to a wide variety of languages, formats, and symbols around the world.

A character set is a collection of characters and symbols used on computers, ensuring no duplication, and refers to a table mapping characters to their assigned numbers. A character encoding scheme refers to the data format used by computers for the numbers assigned to characters.

ex. Character set: Unicode, Encoding scheme: UTF-8

The numbers assigned in Unicode, etc., are called code points, and the rune type is used to represent these code points.

For example, Unicode U+0041 (hexadecimal 0041) is the code point for the character 'A', classified as a basic Latin letter, and can be represented using the rune type as follows:

```go
package main

import "fmt"

func main() {
	var a rune = 'A'             // Single quotes indicate a rune type
	fmt.Printf("%T %U %#v\n", a, a) // int32 U+0041 65
}
```

The number 65 of type int32 is the decimal representation of 0041, representing a Unicode code point.

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

The string type, unlike the rune type, does not hold code points.

The behavior of the string type varies depending on how you loop through it.

```go
package main

import "fmt"

func main() {
	s := "Hello"
	// Bytes can be retrieved
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

When handling multibyte characters, you need to pay attention to the behavior of loops.

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

Since a single Japanese character is 3 bytes, even for one character, the loop will iterate 3 times if not using a range.

When obtaining the number of characters, remember that the string type is a slice of bytes.

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

Since len returns the byte length, use utf8.RuneCountInString to get the number of characters.

# References
- [go.dev - Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
- [pkg.go.dev - runes](https://pkg.go.dev/golang.org/x/text/runes)
- [ja.wikipedia.org - Character Set](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E9%9B%86%E5%90%88)
- [ja.wikipedia.org - Character Encoding Scheme](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E7%AC%A6%E5%8F%B7%E5%8C%96%E6%96%B9%E5%BC%8F)
- [ja.wikipedia.org - Character Code](https://ja.wikipedia.org/wiki/%E6%96%87%E5%AD%97%E3%82%B3%E3%83%BC%E3%83%89)
- [wa3.i-3-i.info - Difference between "Character Set" and "Character Encoding Scheme"](https://wa3.i-3-i.info/diff749moji.html)