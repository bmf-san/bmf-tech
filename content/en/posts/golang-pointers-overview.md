---
title: Overview of Pointers in Golang
slug: golang-pointers-overview
date: 2018-11-13T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
translation_key: golang-pointers-overview
---

# Overview
[The Basics of Pointers in Golang](https://bmf-tech.com/posts/Golang%E3%81%AE%E3%83%9D%E3%82%A4%E3%83%B3%E3%82%BF%E3%81%AE%E5%9F%BA%E6%9C%AC) discusses pointers from the perspective of pass-by-reference and pass-by-value. However, I found myself confused while working with pointers, so I decided to summarize an overview of pointers.

# What is a Pointer Type?
- A type of variable that stores a memory address.

# Defining a Pointer Type
- Use `*` to define it.

```
var s *string // The variable `s` is a pointer of type string. Its type is *string.
```

# Address Operator and Dereference Operator
## Address Operator `&`
- **Generates a pointer type from any value** and retrieves its address.

## Dereference Operator `*`
- Accesses the value from a pointer type variable (dereferencing).

# Pointer Type Variables
```golang
package main

import "fmt"

func main() {
  /**
   * Defining a pointer type
   */
  var pointer *string // A pointer variable of type *string
  var s string
  s = "Hello World"

  /*
   * Using the address operator to handle pointers
   */
  fmt.Printf("%T\n", pointer) // *string
  fmt.Printf("%v\n", pointer) // <nil> (Uninitialized pointer value is nil)

  fmt.Printf("%T\n", &pointer) // **string (Since `pointer` is a *string pointer, `&pointer` generates a pointer to *string)
  fmt.Printf("%v\n", &pointer) // 0xc00000c028 (Gets the address using `&`)

  fmt.Printf("%T\n", s) // string
  fmt.Printf("%v\n", s) // Hello World

  pointer = &s // Assigns the address of `s` to the *string pointer variable `pointer`.
  fmt.Printf("%T\n", pointer) // *string
  fmt.Printf("%v\n", pointer) // 0xc00000e1e0
  fmt.Printf("%v\n", *pointer) // Hello World (Dereferences the *string pointer variable `pointer` to get the value)
  fmt.Printf("%T\n", s) // string
  fmt.Printf("%v\n", s) // Hello World
  fmt.Printf("%T\n", &s) // *string
  fmt.Printf("%v\n", &s) // 0xc00000e1e0
  fmt.Printf("%v\n", *s) // invalid indirect of s (type string) (Error because `s` is not a pointer type variable)

  *pointer = "New World" // Assigns a value to the dereferenced pointer variable.
  fmt.Printf("%T\n", *pointer) // string
  fmt.Printf("%v\n", *pointer) // New World
}
```

# Structs and Pointers
## Review of Structs

```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

func main() {
  var h Human

  fmt.Printf("%v\n", h.Name) // "" (Zero value for string)
  fmt.Printf("%v\n", h.Age) // 0 (Zero value for int)

  h = Human{
    Name: "Tom",
    Age: 20,
  }

  fmt.Printf("%v\n", h) // {Tom 20}
}
```

## Structs and Pointers
### Initializing Structs with Composite Literals
```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

func main() {
  var h *Human // Define a pointer type variable

  h = &Human{ // Assign a pointer to a Human struct using `&`.
    Name: "John",
    Age: 20,
  }

  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // &{John 20}
  fmt.Printf("%v\n", *h) // {John 20}

  h.Name = "Tom"
  h.Age = 40

  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // &{Tom 40}
  fmt.Printf("%v\n", *h) // {Tom 40}
}
```

### Initializing Structs with `new`
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
  fmt.Printf("%v\n", *h) // {Tom 40}
}
```

# Receiver and Pointer
## Value Receiver
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
  var h Human

  fmt.Printf("%T\n", h) // main.Human
  fmt.Printf("%v\n", h) // { 0}

  h = Human{
    Name: "Taro",
    Age: 20,
  }

  fmt.Printf("%T\n", h) // main.Human
  fmt.Printf("%v\n", h) // {Taro 20}

  h.say("Hello") // Taro(20) said Hello
}
```

## Pointer Receiver
```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

// Pointer Receiver
func (h *Human) say(msg string) {
  fmt.Printf("%v(%v) said %v\n", h.Name, h.Age, msg)
}

func main() {
  var h *Human

  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // <nil>

  h = &Human{
    Name: "Taro",
    Age: 20,
  }

  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // &{Taro 20}

  h.say("Hello") // Taro(20) said Hello
}
```

## Difference Between Value Receiver and Pointer Receiver
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
  // Calling value receiver function
  var hForValue Human

  hForValue = Human{
    Name: "Taro",
    Age: 20,
  }

  hForValue.setDataForValue("Jiro", 40)

  fmt.Printf("Name: %v\n Age: %v\n", hForValue.Name, hForValue.Age) // Name: Taro Age: 20

  // Calling pointer receiver function
  var hForPointer *Human

  hForPointer = &Human{
    Name: "Taro",
    Age: 20,
  }

  hForPointer.setDataForPointer("Jiro", 40)

  fmt.Printf("Name: %v\n Age: %v\n", hForPointer.Name, hForPointer.Age) // Name: Jiro Age: 40
}
```

When you want to modify the values of a struct's fields within a function, use a pointer receiver. For reference types like `map` or `chan`, a value receiver is sufficient. However, this is not always the case when considering strict performance requirements.

# If You Get Confused About Pointers
Refer back to the definitions of pointer types and the roles of operators mentioned earlier in this post.

# References
- [Go Pointers](http://cuto.unirita.co.jp/gostudy/post/pointer/)
- [Qiita - Learning Pointers and Addresses in Go](https://qiita.com/Sekky0905/items/447efa04a95e3fec217f)
- [Qiita - Differences Between Ampersand (&) and Asterisk (*) in Go](https://qiita.com/tmzkysk/items/1b73eaf415fee91aaad3)
- [Qiita - Researching Go Structs](https://qiita.com/kmtr/items/1e7caf92aa8bb587906d)
- [[Go] Using Value and Pointer Methods in Structs](https://www.yoheim.net/blog.php?q=20170902)
- [](https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/)
- [Github golang/tour - tour: How do I print an address of the pointer to "Struct fields"? #226](https://github.com/golang/tour/issues/226)
- [StackOverflow - Memory Address for Struct Not Visible](https://stackoverflow.com/questions/29498374/memory-address-for-struct-not-visible)
