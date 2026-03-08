---
title: Setting Up a Golang Development Environment
slug: golang-development-environment
date: 2018-04-07T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
translation_key: golang-development-environment
---

# Overview
We will set up the Go environment.

# Setting Up Go Environment
## Installing Go
The installation method is omitted. I use a tool called **anyenv** for installation.

## Specifying GOPATH
Specify GOPATH in `.bashrc` or `.bash_profile`.

```bash
export GOPATH=$HOME/localdev/project/go_dev // Set as you like
PATH=$PATH:$GOPATH/bin
```

## Checking Go Directory Structure
The directory structure in the local environment will be organized according to the official documentation.

```
go_dev/
├── bin
├── pkg
└── src
```

A development directory for Go called **go_dev** is prepared, containing three directories for different roles, in accordance with the official documentation's directory structure. Executable commands are placed in **bin**, packages in **pkg**, and sources in **src**. The **src/** directory is managed by git.

# Creating a Package
As a check to see if the setup so far is successful, let's create a package.

Prepare a `test/` directory in `src/` and create a file called `main.go` as follows:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

Compile with `go build main.go`, create a binary file, and if a binary file named `test` is generated in `bin/` with `go install`, you're all set.

# Setting Up a Docker Development Environment
- [github - astaxie/build-web/application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/01.0.md)
- [golang.org](https://golang.org/doc/code.html)
- [Setting Up a Go Development Environment](https://medium.com/@Akitsuyoshi/go%E8%A8%80%E8%AA%9E%E3%81%AE%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E3%82%92%E8%A8%AD%E5%AE%9A%E3%81%99%E3%82%8B-77f272229a64)
- [Things That People Coming from Other Languages Encounter When They Start Using Go and Their Solutions](https://qiita.com/mumoshu/items/0d2f2a13c6e9fc8da2a4)