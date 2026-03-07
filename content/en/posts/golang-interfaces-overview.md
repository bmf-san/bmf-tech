---
title: Overview of Golang Interfaces
slug: golang-interfaces-overview
date: 2018-11-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
description: A summary of Golang interfaces.
translation_key: golang-interfaces-overview
---

# Overview
This post summarizes Golang interfaces.

# What is a Golang Interface?
- A type that enumerates only the types of specific methods.
- A struct that implements all the methods declared in the interface `Foo` can be treated as type `Foo`.
- Interfaces enable polymorphism.

# Defining an Interface
```golang
type <TypeName> interface {
  <MethodName(<ArgumentType>, ...)(<ReturnType>, ...)
}
```

```golang
// Ex.
type Human interface {
  say() string
}
```

# Features of Interfaces
## Variables of Interface Type
Variables declared as `interface` type can hold values of any type.

```golang
var i interface{}

i = 123
i = "Hello World"
i = []int{1, 2, 3} // etc...
```

## Interface Type as Function Arguments
When an `interface` type is used as a function argument, values of any type can be passed.

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

Syntax for type assertion:
```golang
<Variable>.(<Type>)
```

Usage involves two variables:
```golang
s, ok := i.(Human)
```

If the variable `i` is of type `Human`, `s` will hold the actual value of `i` as type `Human`, and `ok` will be `true`. Otherwise, `s` will hold the zero value of type `Human`.

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
    fmt.Printf("%v\n", "Not a Human type")
    fmt.Printf("%v\n", s)
    return
  }

  fmt.Printf("%v\n", "Is a Human type")
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

  printOnlyHuman(h) // Is a Human type. {John 20}
  printOnlyHuman(a) // Not a Human type. { 0}
}
```

# Example of Interface Implementation
An example of a common use case for Golang interfaces: adding shared behavior to different types.

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
- [Go Language - Empty Interfaces and Type Assertions](https://blog.y-yuki.net/entry/2017/05/08/000000)
- [Introduction to Go Language - Interfaces](http://cuto.unirita.co.jp/gostudy/post/interface/)
- [Relaxed Go Language Programming Introduction](http://www.geocities.jp/m_hiroi/golang/abcgo09.html)
- [Type Casting and Type Conversion with Type Assertion](https://maku77.github.io/hugo/go/cast.html)
- [Understanding Interface Handling in Go Language](http://maku77.github.io/hugo/go/interface.html)
- [astaxie/build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/02.6.md)
- [SE Book - Starting Go Language](https://www.shoeisha.co.jp/book/detail/9784798142418)
