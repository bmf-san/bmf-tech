---
title: Practical Go Programming
slug: practical-go-language
date: 2023-08-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Book
description: Notes on reading Practical Go Programming.
translation_key: practical-go-language
---



[Practical Go Programming](https://amzn.to/3KO6sr4) has been finished, so I'm leaving my reading notes.

I've been working with Go for several years, but I found it very educational to realize things I didn't know or had forgotten.

# Variable Names
- The suffix for variable names of the Error type that satisfy the Error interface should be Error
  - ex. NotImplementedError
- The prefix for variable names declared with errors.New should be Err
  - ex. `ErrNotImplemented := errors.New("Not Implemented!")`
- While abbreviations are preferred for variable names, **descriptive variable names are desirable for variables used far from their declaration**
- Actively use abbreviated forms for variables
  - If you want to explicitly specify the type, you can use var

# Package Names
- Packages named internal are not exposed outside the module (≒ cannot be used from packages other than the package where internal is located and its sub-packages)
- Avoid duplication of package names and function names
  - ○http.Server ✕http.HTTPServer
  - ○zip.NewWriter ✕zip.NewZipWriter
- Folders named `.` and `_` and `testdata` are excluded from compilation

# Interface Names
- In standard packages, interfaces with only a single method sometimes have the suffix er
  - ex. io.Reader, fmt.Stringer

# Constants
- const can only be used for immutable values determined at compile time
- Be careful of unintended constant value shifts with iota
  - If there is a possibility of adding constants other than at the end, define each constant individually instead of using iota
  - Avoid using iota for values that will be used not only by a single process but also by other processes
    - ex. Be careful with values returned to clients in HTTP server responses
- To convert iota integer values to strings in logs, golang.org/x/tools/cmd/stringer can be used
  - Useful when it's cumbersome to trace the meaning of integer values

# Data Masking
- Can be achieved by extending with the Stringer and GoStringer interfaces

# Structs
- Struct embedding ≠ inheritance
  - Inheritance is a relationship where the child depends on the parent, but struct embedding is a relationship where the parent loosely depends on the child (might be misleading..). Not Is-A, but Has-A.
  - Delegation, not inheritance.
- If you want to distinguish from zero values and omit empty fields in struct fields, use pointer types for the fields

# Functions
- Places where log.Fatal and panic can be used
  - main function, init function, functions with the Must prefix, and other application initialization processes

# Testing
- Detecting order dependencies
  - `go test -shuffle=on` shuffles the order of test execution
