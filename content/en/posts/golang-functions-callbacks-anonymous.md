---
title: About Golang Functions - Function Values / Callback Functions / Anonymous Functions
description: An in-depth look at About Golang Functions - Function Values / Callback Functions / Anonymous Functions, covering key concepts and practical insights.
slug: golang-functions-callbacks-anonymous
date: 2018-10-04T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Callback
  - Function
  - Anonymous Function
translation_key: golang-functions-callbacks-anonymous
---

# Overview
In Golang functions, we will summarize the following three points.
- Functions treated as function values
- Functions that take other functions as arguments
- Definition of anonymous functions
    - Function values
    - Closures

# Functions Treated as Function Values
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

# Callback Functions
```golang
package main

import "fmt"

// Callback function
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

The function `sum` defines two arguments:
- `v` of type int
- `r`, which is a function that takes an **int as an argument and returns an int**
  - When defining a callback function as an argument, it is necessary to specify the return type even if it is an argument.

By the way, the `add` function being executed in the main function has its address stored.

```golang
fmt.Println("%v", add) // 0x10936d0
```

In PHP, callbacks were implemented using variable variables or `call_user_func`.

# Definition of Anonymous Functions
## Function Values
Example of treating an anonymous function as a function value

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

## Closures
Example of defining an anonymous function as a closure

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

Using closures allows the scope to remain open, so the value of `count` can be retained.

# Thoughts
I feel like I have been using callback functions without fully understanding them, so I want to delve deeper into how callback functions work.