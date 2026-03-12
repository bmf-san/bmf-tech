---
title: Summary of Variable Definitions and Declarations in Golang
description: Research notes and a structured overview of Summary of Variable Definitions and Declarations in Golang, summarizing key concepts and findings.
slug: golang-variable-definition-patterns
date: 2018-11-13T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
translation_key: golang-variable-definition-patterns
---

# Overview
A summary of patterns for variable definitions and declarations in Golang.

# Cautions in Variable Definitions and Declarations
- If the first letter is uppercase, the variable is visible from other packages.
- If the first letter is lowercase, the variable is only visible within the package.

# Variable Definitions and Declarations
## Variable declarations
```golang
var i int
fmt.Printf("%T", i) // int
```

```golang
var a, b, c string
fmt.Printf("%T", a) // string
fmt.Printf("%T", b) // string
fmt.Printf("%T", c) // string
```

```golang
var s = "Hello World"
fmt.Printf("%T", s) // string
```

```golang
var x, y, z int = 1, 2, 3
fmt.Printf("%T", x) // int
fmt.Printf("%T", y) // int
fmt.Printf("%T", z) // int
```

``` golang
var (
   a string
   x int
   y, z = 2, 3
)
fmt.Printf("%T", a) // string
fmt.Printf("%T", x) // int
fmt.Printf("%T", y) // int
fmt.Printf("%T", z) // int
```

```golang
var i, j = 1, 2
fmt.Printf("%T", i) // int
fmt.Printf("%T", j) // int
```

## Short variable declarations
- The shorthand variable declaration can only be used within functions.

```golang
i := 1
fmt.Printf("%T", i) // int
```

```golang
i, j := 1, 2
fmt.Printf("%T", i) // int
fmt.Printf("%T", j) // int
```

# References
- [Qiita - Summary for Programmers from Other Languages to Grasp Golang Basics](https://qiita.com/tfrcm/items/e2a3d7ce7ab8868e37f7)
- [Golang.org - Variable Declarations](https://golang.org/ref/spec#Variable_declarations)
