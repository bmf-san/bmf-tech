---
title: Setting Up a Golang Development Environment
slug: golang-development-environment
date: 2018-04-07T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
description: Guide to setting up a Go development environment.
translation_key: golang-development-environment
---

# Overview
Set up a Go development environment.

# Setting Up Go Environment
## Install Go
Installation methods are omitted here. I use a tool called **anyenv** for installation.

## Specify GOPATH
Add GOPATH to `.bashrc` or `.bash_profile`.

```bash
export GOPATH=$HOME/localdev/project/go_dev // Set as you prefer
PATH=$PATH:$GOPATH/bin
```

## Check Go Directory Structure
The local environment directory structure will follow the official documentation.

```
go_dev/
├── bin
├── pkg
└── src
```

Prepare a Go development directory named **go_dev**, and within it, create three directories based on their roles as specified in the official documentation. **bin** contains executable commands, **pkg** contains packages, and **src** contains source code. Only **src/** is managed with git.

# Create a Package
To verify that the setup is working correctly, create a package.

Inside `src/`, create a directory named `test/` and a file `main.go` with the following content:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

Run `go build main.go` to compile and create a binary file. Then, use `go install` to generate a binary file named `test` in `bin/`. If successful, the setup is complete.

# Setting Up Docker Development Environment
- [github - astaxie/build-web/application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/01.0.md)
- [golang.org](https://golang.org/doc/code.html)
- [Setting Up Go Development Environment](https://medium.com/@Akitsuyoshi/go%E8%A8%80%E8%AA%9E%E3%81%AE%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E3%82%92%E8%A8%AD%E5%AE%9A%E3%81%99%E3%82%8B-77f272229a64)
- [Common Pitfalls for Developers New to Go](https://qiita.com/mumoshu/items/0d2f2a13c6e9fc8da2a4)