---
title: GolangのHTTPサーバーのコードリーディング
description: GolangのHTTPサーバーのコードリーディングについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: golang-http-server-code-reading-2019
date: 2019-11-03T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - コードリーディング
  - router
translation_key: golang-http-server-code-reading-2019
---


# 概要
この記事は[Qiita - Go6 Advent Calendar 2019](https://qiita.com/advent-calendar/2019/go6)の20日目の記事です。

GolangでHTTPサーバーを立てるコードの詳細を追ってコードリーディングします。

# 参考実装
コードリーディングしていく実装はこちら。

```golang
package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := new(HelloHandler)
	mux.Handle("/", handler)

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	s.ListenAndServe()
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
```

冗長に書いているこのコードを一行ずつ追ってコードを簡略化しつつ、リーディングしていきます。

# ServeHttp(w ResponseWriter, r *Request) 
まずは、

```golang
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
```

この部分から見ていきます。

`ServeHTTP(w ResponseWriter, r *Request)`は`Handler`インターフェースの実装になります。

```golang
// url: https://golang.org/src/net/http/server.go?s=61586:61646#L1996
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
```

```golang
// url: https://golang.org/src/net/http/server.go?s=61586:61646#L79
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

参考実装では、`ServeHTTP(w ResponseWriter, r *Request)`のために`HelloHandler`構造体を用意していますが、
`HandlerFunc`を利用することでより簡潔に書き直すことができます。

```golang
// url: https://golang.org/src/net/http/server.go?s=61509:61556#L1993
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

参考実装を書き直すとこんな感じです。

```golang
package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(hello))

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	s.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
```

`ServeHTTP(w ResponseWriter, r *Request)`を使っていた部分を書き換えることができました。


ちなみに`mux.Handle`の中身はこんな実装になっています。

```golang
// url: https://golang.org/src/net/http/server.go?s=75321:75365#L2390
func (mux *ServeMux) Handle(pattern string, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("http: invalid pattern")
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if _, exist := mux.m[pattern]; exist {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	e := muxEntry{h: handler, pattern: pattern}
	mux.m[pattern] = e
	if pattern[len(pattern)-1] == '/' {
		mux.es = appendSorted(mux.es, e)
	}

	if pattern[0] != '/' {
		mux.hosts = true
	}
}
```

# ServeMux
先程短くした部分を更に見ていきます。

```golang
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(hello))

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
    }
```

`mux.Handle("/", http.HandlerFunc(hello))`の部分は`HandleFunc`を使うと一部を内部的に処理させることができるので、
より短く書くことができます。

```golang
url: https://golang.org/src/net/http/server.go?s=75575:75646#L2448
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
```

```golang
url: https://golang.org/src/net/http/server.go?s=75575:75646#L2435
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}
```

上記を加味して書き直すとこんな感じになります。

```golang
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	s := http.Server{
		Addr:    ":3000",
		Handler: http.DefaultServeMux,
	}
	s.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
```

`DefaultServeMux`は、内部的には`ServeMux`構造体のポインタが格納された変数になります。
`HandleFunc`は`DefaultServeMux`へのurlパターンマッチの登録ができるメソッドになっています。

```golang
url: https://golang.org/src/net/http/server.go?s=75575:75646#L2207
// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux
```

```golang
url: https://golang.org/src/net/http/server.go?s=68149:68351#L2182
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}
```
# Server
最後に見ていくのはこの部分。

```golang
	s := http.Server{
		Addr:    ":3000",
		Handler: http.DefaultServeMux,
	}
	s.ListenAndServe()
```

`s.ListenAndServe()`の中身。

```golang
url: https://golang.org/src/net/http/server.go?s=68149:68351#L3093
func (srv *Server) ListenAndServe() error {
	if srv.shuttingDown() {
		return ErrServerClosed
	}
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}
```

`Server`に細かい設定値を与える必要がないときは`ListenAndServe()`を使うことで短く書くことができる。
`Server`の設定値については[golang.org - server.go](https://golang.org/src/net/http/server.go?s=77156:81268#L2480)を参照。

```golang
url: https://golang.org/src/net/http/server.go?s=68149:68351#L3071
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

短く書くとこんな感じです。

```golang
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
```

無名関数を使って使うとこんな感じです。

```golang
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":3000", nil)
}
```

# 所感
golangでhttp routerのパッケージを自作しようとしていて、net/httpの内部的な実装に触れておく必要があったので軽く調べてみました。
見た感じ拡張しやすそうなので自作はしやすいイメージがあります。

# 追記
URLルーター実装しました。

[github.com - bmf-san/goblin](https://github.com/bmf-san/goblin)

# 参考
- [golangのHTTPサーバを構成しているもの](https://reiki4040.hatenablog.com/entry/2017/03/01/212647)
- [【Go】net/httpパッケージを読んでhttp.HandleFuncが実行される仕組み](https://qiita.com/shoichiimamura/items/1d1c64d05f7e72e31a98)
