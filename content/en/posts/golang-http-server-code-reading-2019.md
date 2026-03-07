---
title: Code Reading of Golang HTTP Server
slug: golang-http-server-code-reading-2019
date: 2019-11-03T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Code Reading
  - Router
description: A detailed code reading of setting up an HTTP server in Golang.
translation_key: golang-http-server-code-reading-2019
---

# Overview
This article is the 20th entry in the [Qiita - Go6 Advent Calendar 2019](https://qiita.com/advent-calendar/2019/go6).

We will conduct a detailed code reading of setting up an HTTP server in Golang.

# Reference Implementation
Here is the implementation we will analyze:

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

We will go through this verbose code line by line, simplifying it while reading.

# ServeHttp(w ResponseWriter, r *Request)
First, let’s look at this part:

```golang
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
```

`ServeHTTP(w ResponseWriter, r *Request)` is an implementation of the `Handler` interface.

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

In the reference implementation, the `HelloHandler` struct is prepared for `ServeHTTP(w ResponseWriter, r *Request)`. However, it can be rewritten more concisely using `HandlerFunc`.

```golang
// url: https://golang.org/src/net/http/server.go?s=61509:61556#L1993
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

Rewriting the reference implementation:

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

We successfully rewrote the part using `ServeHTTP(w ResponseWriter, r *Request)`.

By the way, the implementation of `mux.Handle` looks like this:

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
Let’s dive deeper into the shortened part:

```golang
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(hello))

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
    }
```

The part `mux.Handle("/", http.HandlerFunc(hello))` can be further simplified using `HandleFunc`, which internally handles some of the processing.

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

Considering the above, the rewritten code looks like this:

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

`DefaultServeMux` is essentially a variable that holds a pointer to a `ServeMux` struct. `HandleFunc` is a method that allows registering URL pattern matches to `DefaultServeMux`.

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
Finally, let’s look at this part:

```golang
	s := http.Server{
		Addr:    ":3000",
		Handler: http.DefaultServeMux,
	}
	s.ListenAndServe()
```

The content of `s.ListenAndServe()`:

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

When no specific settings are needed for the `Server`, you can write it more concisely using `ListenAndServe()`. For details on `Server` settings, refer to [golang.org - server.go](https://golang.org/src/net/http/server.go?s=77156:81268#L2480).

```golang
url: https://golang.org/src/net/http/server.go?s=68149:68351#L3071
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

Here’s the shorter version:

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

Using an anonymous function:

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

# Thoughts
I was planning to create my own HTTP router package in Golang, so I did some research on the internal implementation of `net/http`. It seems to be quite extensible, making it relatively easy to create custom implementations.

# Addendum
I implemented a URL router.

[github.com - bmf-san/goblin](https://github.com/bmf-san/goblin)

# References
- [golangのHTTPサーバを構成しているもの](https://reiki4040.hatenablog.com/entry/2017/03/01/212647)
- [【Go】net/httpパッケージを読んでhttp.HandleFuncが実行される仕組み](https://qiita.com/shoichiimamura/items/1d1c64d05f7e72e31a98)