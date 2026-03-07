---
title: Using Workspace Mode in Go's Multi-Module Structure
slug: go-multi-module-workspace-mode
date: 2024-01-19T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
translation_key: go-multi-module-workspace-mode
---

I had never used Workspace mode, which was added in Go 1.18, so I decided to give it a try.

# What is Workspace Mode
A feature that facilitates Go's multi-module structure.

# How to Use Workspace Mode
Prepare the following structure.

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

Run the following command in the foo directory to set up go.mod.
```sh
go mod init example.com/foo
```

Similarly, run the following command in the bar directory to set up go.mod.
```sh
go mod init example.com/bar
```

Next, create a cmd directory and create the main.go file as follows.

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

At this point, the structure will look like this.

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

In the root directory, run the following command to configure the workspace.

```sh
go work init foo bar cmd
```

A file named go.work will be created.

```go
go 1.21.1

use (
	./bar
	./cmd
	./foo
)
```

Confirm that `go run cmd/main.go` can be executed.

```sh
// Output
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

Run `go mod init example.com/baz` to generate go.mod just like the other modules.

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

Return to the root and run `go work use baz` to add the module, which will add baz to go.work.

```go
go 1.21.1

use (
	./bar
	./baz
	./cmd
	./foo
)
```

Run `go run cmd/main.go` to confirm that baz has been added to the output.

```sh
// Output
foo
bar
baz
```

# Thoughts
I thought the multi-module structure was cumbersome, so I'm glad it has become easier.

# References
- [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)