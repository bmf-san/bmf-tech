---
title: Using Workspace Mode with Go's Multi-Module Structure
slug: go-multi-module-workspace-mode
date: 2024-01-19T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
description: Exploring the Workspace mode introduced in Go 1.18.
translation_key: go-multi-module-workspace-mode
---


I had never used the Workspace mode added in Go 1.18, so I decided to give it a try.

# What is Workspace Mode?
A feature to facilitate Go's multi-module structure.

# How to Use Workspace Mode
Prepare the following structure:

```
.
├── bar
│   └── bar.go
└── foo
    └── foo.go
```

```go
// foo.go
package foo

func Foo() string {
	return "foo"
}

```

```go
// bar.go
package bar

func Bar() string {
	return "bar"
}
```

Execute the following command in the foo directory to set up go.mod.
```sh
go mod init example.com/foo
```

Similarly, execute the following command in the bar directory to set up go.mod.
```sh
go mod init example.com/bar
```

Next, create a cmd directory and create a main.go file as follows:

```go
package main

import (
	"example.com/bar"
	"example.com/foo"
)

func main() {
	println(foo.Foo())
	println(bar.Bar())
}
```

Set up go.mod in the cmd directory as well.
```sh
go mod init example.com/cmd
```

At this point, the structure will be as follows:

```sh
.
├── bar
│   ├── bar.go
│   └── go.mod
├── cmd
│   ├── go.mod
│   └── main.go
└── foo
    ├── foo.go
    └── go.mod
```

In the root directory, execute the following command to configure the workspace.


```sh
go work init foo bar cmd
```

A file named go.work is created.

```go
go 1.21.1

use (
	./bar
	./cmd
	./foo
)
```

Verify that `go run cmd/main.go` can be executed.

```sh
// Execution result
foo
bar
```

Next, let's add a module called baz.

```sh
.
├── bar
│   ├── bar.go
│   └── go.mod
├── baz
│   ├── baz.go
│   └── go.mod
├── cmd
│   ├── go.mod
│   └── main.go
├── foo
│   ├── foo.go
│   └── go.mod
└── go.work
```

```go
// baz.go
package baz

func Baz() string {
	return "baz"
}
```

Execute `go mod init example.com/baz` to generate go.mod like the other modules.

Next, add baz to main.go.

```go
package main

import (
	"example.com/bar"
	"example.com/baz"
	"example.com/foo"
)

func main() {
	println(foo.Foo())
	println(bar.Bar())
	println(baz.Baz())
}
```

Return to the root and execute `go work use baz` to add the module, and baz will be added to go.work.

```go
go 1.21.1

use (
	./bar
	./baz
	./cmd
	./foo
)
```

Execute `go run cmd/main.go` and confirm that baz is added to the output.

```sh
// Execution result
foo
bar
baz
```

# Impressions
I thought the multi-module structure was cumbersome, so I'm glad it has become simpler.

# References
- [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
