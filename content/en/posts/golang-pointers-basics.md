---
title: Basics of Pointers in Golang
description: 'Learn the basics of Go pointers: address operators (&), dereferencing (*), pass by value vs. reference, nil pointers, and when to use pointer receivers. Includes practical Go examples.'
slug: golang-pointers-basics
date: 2018-09-03T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Pointer
  - Pass by Value
  - Reference
translation_key: golang-pointers-basics
---



# Overview
This post summarizes the basics of pointers in Golang.

# Prerequisites
A basic understanding of the relationship between variables and memory.

# What is a Pointer?
A pointer refers to the address of a variable. Through the address of a variable, you can change the value of the original variable.

In Golang, C-like pointers are available. In Golang, a pointer to a variable T is of type *T, and its zero value is nil.

```golang
package main

import "fmt"

func main() {
  i := 10

  // Use & (address operator) to access the address of a variable
  // Store the address of variable i in variable pointer
  pointer := &i

  // Confirm that the addresses are the same
  fmt.Println(&i)
  fmt.Println(pointer)

  // Use * to extract the value from the variable pointer that points to the address
  // Confirm that the value of the original variable i has been changed
  *pointer = 100
  fmt.Println(*pointer) // 100
  fmt.Println(i) // Changed to 100
}
```

In C language, there is pointer arithmetic, but not in Golang.

# Difference Between Pass by Value and Pass by Reference
## Pass by Value
One way to pass a variable as an argument to a function. In pass by value, the value of the variable is copied to a different address. Since it is copied to a different address, it does not affect the original variable.

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

  foo(x) // Pass by value: the value of variable x is copied to a different address as the formal parameter x of the foo function
  fmt.Println(x) // 10
}
```

In Golang, arguments are generally passed by value, but this is not the case when using slices or maps. (In the case of slices, the slice variable holds a reference to the slice, not the slice itself.)

## Pass by Reference
In pass by reference, the variable has the same value and address as the original variable. Unlike pass by value, where the value of the variable is copied to a different address, it affects the original function because it is not a copy. The difference with pointers is that pointers handle addresses and values separately, whereas references handle both addresses and values simultaneously. Golang has the concept of pass by reference, but apparently, there is no syntax for pass by reference. Reference types behave like pass by reference.

# Summary
Use & to pass an address, and * to extract the value from the variable that was passed the address.

# Thoughts
To understand more deeply, it seems beneficial to delve into the relationship between variables and memory and to relearn C language pointers. Having an image of how memory handles data for pass by value, pass by reference, and pointers makes it easier to understand. Understanding pointers seems necessary when dealing with structs, so it's important to study them properly.
