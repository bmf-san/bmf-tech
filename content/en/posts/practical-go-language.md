---
title: 'Practical Go Language'
description: 'Practical Go Language'
slug: practical-go-language
date: 2023-08-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Book Review
translation_key: practical-go-language
books:
  - asin: "4873119693"
    title: "Practical Go Language"
---

I finished reading [Practical Go Language](https://amzn.to/3KO6sr4), so I’d like to leave some reading notes.

Having worked with Go for several years, I realized many things I didn’t know or had forgotten, making it a very educational experience.

# Variable Names
- The suffix for variable names of types that satisfy the Error interface is Error
  - ex. NotImplementedError
- The prefix for variable names declared with errors.New is Err
  - ex. `ErrNotImplemented := errors.New("Not Implemented!")`
- While short forms of variable names are preferred, **descriptive variable names are desirable for variables that will be used far from their declaration**.
- Short forms of variable names should be used actively.
  - It’s acceptable to use var if you want to make the type explicit.

# Package Names
- Packages named internal are not exposed outside the module (i.e., they cannot be used from packages other than the one containing internal and its sub-packages).
- Avoid duplication of package names and function names.
  - ○http.Server　✕http.HTTPServer
  - ○zip.NewWriter ✕zip.NewZipWriter
- Folders named `.`, `_`, and `testdata` are excluded from compilation.

# Interface Names
- In standard packages, interfaces that have only a single method often have the suffix er.
  - ex. io.Reader, fmt.Stringer

# Constants
- const can only be used for immutable values determined at compile time.
- Be cautious of unintended shifts in constant values with iota.
  - If there’s a possibility of adding constants other than at the end, define each constant individually instead of using iota.
  - Avoid using iota for values that will be used in other processes, not just one.
    - ex. Be careful with values returned to clients in HTTP server responses.
- To convert iota integer values to strings in logs, you can use golang.org/x/tools/cmd/stringer.
  - This is convenient when you need to avoid the hassle of deciphering integer values.

# Data Masking
- This can be achieved by extending the Stringer and GoStringer interfaces.

# Structs
- Struct embedding ≠ inheritance.
  - Inheritance represents a relationship where the child depends on the parent, while struct embedding represents a looser dependency of the parent on the child (this might be misleading). It’s not Is-A, but Has-A.
  - It’s delegation, not inheritance.
- If you want to distinguish a field from its zero value and use omitempty, use pointer types for the field.

# Functions
- Places where it’s acceptable to use log.Fatal and panic:
  - main function, init function, functions with the Must prefix, and other application initialization processes.

# Testing
- Detection of order dependence:
  - The test execution order is shuffled with `go test -shuffle=on`.
