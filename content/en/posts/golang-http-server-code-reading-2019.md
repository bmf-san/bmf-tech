---
title: Code Reading of Golang's HTTP Server
description: "Dive deep into Go's HTTP server internals: ServeMux routing, Handler interface, middleware patterns, and request handling."
slug: golang-http-server-code-reading-2019
date: 2019-11-03T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Code Reading
  - Router
translation_key: golang-http-server-code-reading-2019
---

# Overview
This article is the 20th entry of the [Qiita - Go6 Advent Calendar 2019](https://qiita.com/advent-calendar/2019/go6).

We will perform a code reading of the details of setting up an HTTP server in Golang.

# Reference Implementation
Here is the implementation we will be reading through.

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

We will go through this somewhat verbose code line by line, simplifying it as we read.

# ServeHttp(w ResponseWriter, r *Request) 
First, let's look at:

```golang
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
```

This part implements the `Handler` interface.

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

In the reference implementation, a `HelloHandler` struct is prepared for `ServeHTTP(w ResponseWriter, r *Request)`, but we can rewrite it more concisely using `HandlerFunc`.

```golang
// url: https://golang.org/src/net/http/server.go?s=61509:61556#L1993
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

Rewriting the reference implementation looks like this:

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

We were able to rewrite the part that used `ServeHTTP(w ResponseWriter, r *Request)`.

By the way, the implementation inside `mux.Handle` looks like this:

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
Next, let's take a closer look at the shortened part.

```golang
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(hello))

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
```

The part `mux.Handle("/", http.HandlerFunc(hello))` can be written more concisely by using `HandleFunc`, which allows some internal processing.

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

Taking the above into account, we can rewrite it like this:

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

`DefaultServeMux` is a variable that internally holds a pointer to the `ServeMux` struct. `HandleFunc` is a method that allows registration of URL pattern matches to `DefaultServeMux`.

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
Finally, let's look at this part.

```golang
	s := http.Server{
		Addr:    ":3000",
		Handler: http.DefaultServeMux,
	}
	s.ListenAndServe()
```

The internals of `s.ListenAndServe()` are as follows:

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

When there is no need to provide detailed configuration values to the `Server`, we can write it more concisely by using `ListenAndServe()`. For details about the `Server` configuration values, refer to [golang.org - server.go](https://golang.org/src/net/http/server.go?s=77156:81268#L2480).

```golang
url: https://golang.org/src/net/http/server.go?s=68149:68351#L3071
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

When written concisely, it looks like this:

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

Using an anonymous function, it looks like this:

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
I was trying to create my own HTTP router package in Golang, and I needed to touch on the internal implementation of net/http, so I did a little research. It seems easy to extend, so I have a positive impression of creating my own.

# Postscript
I implemented a URL router.

[github.com - bmf-san/goblin](https://github.com/bmf-san/goblin)

# References
- [Components of Golang's HTTP Server](https://reiki4040.hatenablog.com/entry/2017/03/01/212647)
- [How http.HandleFunc is executed in the net/http package](https://qiita.com/shoichiimamura/items/1d1c64d05f7e72e31a98)