---
title: Basics of Golang Pointers
slug: golang-pointers-basics
date: 2018-09-03T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Pointer
  - Pass by Value
  - Pass by Reference
translation_key: golang-pointers-basics
---

# Overview
This post summarizes the basics of pointers in Golang.

# Prerequisites
Basic knowledge to visualize the relationship between variables and memory.

# What is a Pointer?
A pointer points to the address of a variable. It allows you to change the value of the calling variable through its address.

Golang provides C-like pointers. In Golang, the pointer of a variable T is of type *T, and its zero value is nil.

```golang
package main

import "fmt"

func main() {
  i := 10

  // Use the & (address operator) to access the variable's address
  // Store the address of variable i in the variable pointer
  pointer := &i

  // Confirm that the addresses are the same
  fmt.Println(&i)
  fmt.Println(pointer)

  // Use * to retrieve the value from the pointer variable
  // Confirm that the value of the calling variable i has changed
  *pointer = 100
  fmt.Println(*pointer) // 100
  fmt.Println(i) // Changed to 100
}
```

In C, pointer arithmetic exists, but it does not in Golang.

# Difference Between Pass by Value and Pass by Reference
## Pass by Value
One method of passing a variable as an argument to a function. In pass by value, the variable's value is copied to a different address. Since it is copied to a different address, it does not affect the calling variable.

```golang
package main

import (
  "fmt"
)

func foo(x int) (int) {
  return x
}

func main() {
  x := 1

  foo(x) // Pass by value: the value of variable x is copied to the function's parameter x at a different address
  fmt.Println(x) // 1
}
```

In Golang, the default method of passing arguments to functions is pass by value, but this does not apply when using slices or maps. (In the case of slices, the slice variable holds a reference to the slice itself, not the slice itself.)

## Pass by Reference
When a variable is passed by reference, its value and address are the same as the calling variable. Unlike pass by value, which copies the variable's value to a different address, pass by reference does not copy, so it affects the calling function. The difference from pointers is that pointers operate on addresses and values separately, while references operate on both addresses and values simultaneously. Golang has the concept of pass by reference, but there seems to be no specific syntax for it. Reference types exhibit pass by reference behavior.

# Summary
Use & to pass an address and * to retrieve a value from the variable that was passed by address.

# Thoughts
To gain a deeper understanding, it seems beneficial to delve into the relationship between variables and memory and to revisit pointers in C. Having a mental image of how memory handles data for pass by value, pass by reference, and pointers will make it easier to understand. It seems necessary to understand pointers when dealing with structs, so I will study them properly.