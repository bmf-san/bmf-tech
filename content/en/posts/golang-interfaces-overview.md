---
title: "Go Interfaces Explained: How to Write Flexible, Testable Go Code"
description: 'A practical guide to Go interfaces. Learn implicit implementation, interface composition, using interfaces for dependency injection, and common design pitfalls.'
slug: golang-interfaces-overview
date: 2018-11-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
translation_key: golang-interfaces-overview
---

# Overview
Summarizing Golang interfaces.

# What is a Golang Interface
- A type that enumerates only specific method types.
- A struct that implements all methods declared in the interface Foo can be treated as type Foo.
- Using interfaces allows for the realization of polymorphism.

# Definition of an Interface
```golang
type <TypeName> interface {
  <MethodName>(<ArgumentType>, ...)(<ReturnType>, ...)
}
```

```golang
// Ex.
type Human interface {
  say() string
}
```

# Characteristics of Interfaces
## Variables of Interface Type
Variables declared as interface type can hold values of any type.

```golang
var i interface{}

i = 123
i = "Hello World"
i = []int{1, 2, 3} // etc...
```

## Interface Type Arguments
When an interface type is used as an argument, any type of value can be passed.

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

# Type Assertion

The syntax for type assertion.
```golang
<variable>.(<type>)
```

It is used in a way that takes two variables.
```golang
s, ok := i.(Human)
```

If variable i is of type Human, variable s will hold the actual value of variable i of type Human, and variable ok will be true. Conversely, if variable i is not of type Human, variable s will hold the zero value of type Human.

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
    fmt.Printf("%v\n", "Not of type Human")
    fmt.Printf("%v\n", s)
    return
  }

  fmt.Printf("%v\n", "Is of type Human")
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

  printOnlyHuman(h) // Is of type Human. {John 20}
  printOnlyHuman(a) // Not of type Human. { 0}
}
```

# Example of Interface Implementation
An example of a popular use case for Golang interfaces, which is to "add common properties to different types."

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

func do(a Action) { // Accepts Action type
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

# References
- [golang.org - Interface types](https://golang.org/ref/spec#Interface_types)
- [Go Language - Empty Interface and Type Assertion](https://blog.y-yuki.net/entry/2017/05/08/000000)
- [First Go Language - Interfaces](http://cuto.unirita.co.jp/gostudy/post/interface/)
- [Easy Go Language Programming Introduction](http://www.geocities.jp/m_hiroi/golang/abcgo09.html)
- [Type Casting and Type Conversion via Type Assertion](https://maku77.github.io/hugo/go/cast.html)
- [Understanding the Handling of Interfaces in Go Language](http://maku77.github.io/hugo/go/interface.html)
- [astaxie/build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/02.6.md)
- [SE Book - Starting Go Language](https://www.shoeisha.co.jp/book/detail/9784798142418)