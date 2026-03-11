---
title: Goのマルチモジュール構成でWorkspace modeを使ってみる
slug: go-multi-module-workspace-mode
date: 2024-01-19T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
translation_key: go-multi-module-workspace-mode
---


Go1.18から追加されたWorkspace modeを使ったことがなかったので、使ってみた。

# Workspace modeとは
Goのマルチモジュール構成を便利するための機能。

# Workspace modeの使い方
次のような構成を用意する。

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

fooディレクトリで次のコマンドを実行して、go.modをセットアップする。
```sh
go mod init example.com/foo
```

同じく、barディレクトリで次のコマンドを実行して、go.modをセットアップする。
```sh
go mod init example.com/bar
```

次に、cmdディレクトリを作成して、main.goファイルを次のように作成する。

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

cmdディレクトリでも同様にgo.modをセットアップする。
```sh
go mod init example.com/cmd
```

ここまでで、次のような構成になる。

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

ルートのディレクトリにて、次のコマンドを実行し、workspaceの設定を行う。


```sh
go work init foo bar cmd
```

go.workというファイルが作成される。

```go
go 1.21.1

use (
	./bar
	./cmd
	./foo
)
```

`go run cmd/main.go`が実行できることを確認。

```sh
// 実行結果
foo
bar
```

続いて、bazというモジュールを追加してみる。

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

`go mod init example.com/baz`を実行して他のモジュールと同じようにgo.modを生成する。

続いて、main.goにbazを追加する。

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

ルートに戻り、`go work use baz`を実行し、モジュールを追加すると、go.workにbazが追加される。

```go
go 1.21.1

use (
	./bar
	./baz
	./cmd
	./foo
)
```

`go run cmd/main.go`を実行して、出力にbazが追加されていることを確認する。

```sh
// 実行結果
foo
bar
baz
```

# 所感
マルチモジュール構成が手間だなぁと思っていたので簡単になって良いなと思った。

# 参考
- [Tutorial: Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
