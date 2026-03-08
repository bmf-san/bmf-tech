---
title: Practical Go Language
slug: practical-go-language
date: 2023-08-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Books
translation_key: practical-go-language
---

[Practical Go Language](https://amzn.to/3KO6sr4) has been completed, so I will leave my reading notes.

I have been working with Go for several years, but I realized many things I didn't know or had forgotten, which was very educational.

# Variable Names
- The suffix for variable names of type Error that satisfy the Error interface is Error
  - ex. NotImplementedError
- The prefix for error variable names declared with errors.New is Err
  - ex. `ErrNotImplemented := errors.New("Not Implemented!")`
- While abbreviated variable names are preferred, **descriptive variable names** are desirable for variables that are used far from their declaration.
- Variables should actively use abbreviated forms.
  - If you want to be explicit about the type, it's okay to use var.

# Package Names
- Packages named internal are not exposed outside the module (i.e., they can only be used from the package where internal is located and its sub-packages).
- Avoid duplication of package names and function names.
  - ○http.Server  ✕http.HTTPServer
  - ○zip.NewWriter ✕zip.NewZipWriter
- Folders named `.`, `_`, and `testdata` are excluded from compilation.

# Interface Names
- In standard packages, interfaces that have only a single method often have the suffix er.
  - ex. io.Reader, fmt.Stringer

# Constants
- const can only be used for immutable values determined at compile time.
- Be careful of unintended shifts in constant values with iota.
  - If there is a possibility of adding constants other than at the end, define each constant individually instead of using iota.
  - Avoid using iota for values that will be used by other processes, not just one process.
    - ex. Be cautious with values returned to clients in HTTP server responses.
- To convert iota integer values to strings in logs, golang.org/x/tools/cmd/stringer can be used.
  - It's convenient when you have to track meanings with integer values, which can be cumbersome.

# Data Masking
- This can be achieved by extending using the Stringer and GoStringer interfaces.

# Structs
- Struct embedding ≠ inheritance.
  - Inheritance is a relationship where the child depends on the parent, but struct embedding is a relationship where the parent has a loose dependency on the child (this might be misleading...). It's not Is-A, but Has-A.
  - It's delegation, not inheritance.
- If you want to distinguish a field from its zero value and use omitempty, use pointer types for the field.

# Functions
- Places where log.Fatal and panic can be used:
  - main function, init function, functions with Must prefix, and other application initialization processes.

# Tests
- Detection of order dependency:
  - The test execution order is shuffled with `go test -shuffle=on`.