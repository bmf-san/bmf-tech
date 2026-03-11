---
title: DefaultServeMux以外でpprofを使う方法
slug: using-pprof-without-defaultmux
date: 2023-04-30T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - Tips
translation_key: using-pprof-without-defaultmux
---


# 概要
[net/http/pprof](https://pkg.go.dev/net/http/pprof)をDefaultServeMux以外（Goの標準のルーター以外）で使う方法についてメモ。

# ハマりどころ
pprofをblank importするだけではだめ。

```go
package main

import (
    _ "net/http/pprof"
)
```

DefaultServeMux以外のルーターを使う場合はblank importするだけではpprofが利用できるようにならない。

[net/http/pprof](https://pkg.go.dev/net/http/pprof)を参照すると、下記のように記載されている。

> If you are not using DefaultServeMux, you will have to register handlers with the mux you are using.

# 解決策
下記は自分の自作ルーター[bmf-san/goblin](https://github.com/bmf-san/goblin)を使った例。

```go
package main

import (
    "net/http/pprof"
)

func main() {
        r.Methods(http.MethodGet).Handler("/debug/pprof/", http.HandlerFunc(pprof.Index))
	r.Methods(http.MethodGet).Handler("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	r.Methods(http.MethodGet).Handler("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	r.Methods(http.MethodGet).Handler("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	r.Methods(http.MethodGet).Handler("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	r.Methods(http.MethodGet).Handler("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Methods(http.MethodGet).Handler("/debug/pprof/heap", pprof.Handler("heap"))
        r.Methods(http.MethodGet).Handler("/debug/pprof/mutex", pprof.Handler("mutex"))
	r.Methods(http.MethodGet).Handler("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Methods(http.MethodGet).Handler("/debug/pprof/block", pprof.Handler("block"))
}
```

上述のようにルーティングを自分で設定し、pprofのHanderを設定してあげる必要がある。

httprouterの場合であれば、下記issueが参考になる。
[pprof issue with httprouter #236](https://github.com/julienschmidt/httprouter/issues/236)

# 余談
[pyroscope](https://pyroscope.io/)でGoのプロファイリングをPull型で設定しようとしたときにハマった。

