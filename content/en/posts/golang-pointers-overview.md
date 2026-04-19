---
title: "Go Pointers Explained: When to Use Them and Common Pitfalls"
description: 'Understand Go pointers clearly—when to use pointer receivers vs value receivers, how pointers affect memory and performance, and the most common mistakes to avoid.'
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
[Golang Pointers Basics](https://bmf-tech.com/posts/Golang%E3%81%AE%E3%83%9D%E3%82%A4%E3%83%B3%E3%82%BF%E3%81%AE%E5%9F%BA%E6%9C%AC) discussed pointers from the perspective of reference passing and value passing. However, I became confused while handling pointers, so I will summarize the overview of pointers.

# Pointer Type
- A variable type that stores an address in memory.

# Defining Pointer Type
- Defined using *

```
var s *string // Variable s is a pointer of type string. The type is *string
```

# Address Operator and Indirection Operator
## Address Operator &
- **Generates a pointer type from any value** and obtains the address.

## Indirection Operator *
- References a value from a pointer type variable (dereference).

# Pointer Type Variable
```golang
package main

import "fmt"

func main() {
  /**
   * Definition of pointer type
   */
  var pointer *string // Pointer variable of type *string
  var s string
  s = "Hello World"

  /*
   * Handling pointers with address operator
   */
  fmt.Printf("%T\n", pointer) // *string
  fmt.Printf("%v\n", pointer) // <nil>  ※ The value of an uninitialized pointer type is nil

  fmt.Printf("%T\n", &pointer) // **string  ※ Since pointer is a pointer variable of type *string, &pointer generates a pointer of *string
  fmt.Printf("%v\n", &pointer) // 0xc00000c028  ※ Obtain the address of the pointer using &

  fmt.Printf("%T\n", s) // string
  fmt.Printf("%v\n", s) // Hello World

  pointer = &s // ※ Generate a pointer type from s and assign the address to the *string pointer variable (pointer is already defined as a pointer variable of type *string, so use & to generate the pointer type of variable s and assign it).
  fmt.Printf("%T\n", pointer) // *string  ※ pointer remains a pointer of type *string
  fmt.Printf("%v\n", pointer) // 0xc00000e1e0
  fmt.Printf("%v\n", *pointer) // Hello World  ※ Reference the value from pointer (pointer variable of type *string)
  fmt.Printf("%T\n", s) // string
  fmt.Printf("%v\n", s) // Hello World
  fmt.Printf("%T\n", &s) // *string
  fmt.Printf("%v\n", &s) // 0xc00000e1e0
  fmt.Printf("%v\n", *s) // invalid indirect of s (type string)  ※ Error because s is not a pointer type variable. Indirect means indirect.

  *pointer = "New World" // ※ *pointer is a variable, so it can be assigned. Pointer variable definition uses *.
  fmt.Printf("%T\n", *pointer) // string  ※ *pointer references the value from the pointer variable
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

  fmt.Printf("%v\n", h.Name) // ""  ※ Zero value of string
  fmt.Printf("%v\n", h.Age) // 0  ※ Zero value of int

  h = Human{
    Name: "Tom",
    Age: 20,
  }

  fmt.Printf("%v\n", h) // {Tom 20}
}
```

## Structs and Pointers
## Initializing Structs with Composite Literals
```golang
package main

import "fmt"

type Human struct {
  Name string
  Age int
}

func main() {
  var h *Human // Definition of pointer type variable

  h = &Human{ // ※ Variable h is of type *Human, so use & to assign a pointer type from struct Human.
    Name: "John",
    Age: 20,
  }

  fmt.Printf("%T\n", h) // *main.Human  ※ Generating a pointer of Human with &Human
  fmt.Printf("%v\n", h) // &{John 20}  ※ Pointer of type *main.Human
  fmt.Printf("%v\n", *h) // {John 20}  ※ Extracting the value

  h.Name = "Tom"
  h.Age = 40

  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // &{Tom 40}
  fmt.Printf("%v\n", *h) // {John 40}  ※ Extracting the value
}
```

## Initializing Structs with new
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
  fmt.Printf("%v\n", *h) // {Tom 40}  ※ Extracted the value
}
```

# Receivers and Pointers
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
  var h Human // Defining variable h of type Human

  fmt.Printf("%T\n", h) // main.Human
  fmt.Printf("%v\n", h) // { 0}  ※ Zero values of Name and Age

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

// Pointer receiver
func (h *Human) say(msg string) {
  fmt.Printf("%v(%v) said %v\n", h.Name, h.Age, msg)
}

func main() {
  var h *Human // ※ Definition of pointer type

  fmt.Printf("%T\n", h) // *main.Human
  fmt.Printf("%v\n", h) // <nil> ※ Zero value of pointer type

  h = &Human{
    Name: "Taro",
    Age: 20,
  }

  fmt.Printf("%T\n", h) // *main.Human ※ Pointer of struct &Human assigned to h
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
  // Calling value function
  var hForValue Human

  hForValue = Human{
    Name: "Taro",
    Age: 20,
  }

  hForValue.setDataForValue("Jiro", 40)

  fmt.Printf("Name: %v\n Age: %v\n", hForValue.Name, hForValue.Age) // Name: Taro Age: 20

  // Calling pointer function
  var hForPointer *Human

  hForPointer = &Human{
    Name: "Taro",
    Age: 20,
  }

  hForPointer.setDataForPointer("Jiro", 40)

  fmt.Printf("Name: %v\n Age: %v\n", hForPointer.Name, hForPointer.Age) // Name: Jiro Age: 40
}
```
When you want to change the values of struct fields within a function, use a pointer receiver. For reference types like map and chan, a value receiver is also acceptable. However, when strictly considering performance, this distinction may not hold.

# When Confused About Pointers
Recall the definition of pointer types and the roles of operators described in the first half.

# References
- ~~Go Pointers~~
- [Qiita - Learning Pointers and Addresses in Go](https://qiita.com/Sekky0905/items/447efa04a95e3fec217f)
- [Qiita - Differences Between & and * in Go](https://qiita.com/tmzkysk/items/1b73eaf415fee91aaad3)
- [Qiita - Research on Go Structs](https://qiita.com/kmtr/items/1e7caf92aa8bb587906d)
- [[Go] Using Value Methods and Pointer Methods with Structs](https://www.yoheim.net/blog.php?q=20170902)
- [](https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/)
- [Github golang/tour - tour: How do I print an address of the pointer to "Struct fields" ? #226](https://github.com/golang/tour/issues/226)
- [StackOverflow - Memory address for struct not visible](https://stackoverflow.com/questions/29498374/memory-address-for-struct-not-visible)