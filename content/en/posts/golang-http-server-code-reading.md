---
title: Code Reading of Golang HTTP Server
slug: golang-http-server-code-reading
date: 2021-06-30T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Code Reading
description: Notes from reading the internal code of an HTTP server while creating a router in Go.
translation_key: golang-http-server-code-reading
---

# Overview
I took some notes while reading the internal code of an HTTP server in Go when I created a router.

[github.com - bmf-san/goblin](https://github.com/bmf-san/goblin)

# Code Reading of HTTP Server
## Basic Form
This is the form of code you often see when starting with Go.

Various components are omitted to achieve this form.

```golang
package main

import (
	"net/http"
)

func main() {
	http.HandlerFunc("/index", func(w http.ReponseWriter, req *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":8080")
}
```

## Detailed Form Without Omissions
This is the detailed form of the basic example above, without any omissions.
We will check step by step how the basic form is implemented.

```golang
package main

import (
	"net/http"
)

func main() {
	// Multiplexer. A structure for URL matching. Resolves only static routing.
	mux := http.NewServeMux()
	ih := new(indexHandler)
	// Register routing to the mux
	mux.Handle("/index", ih)

	srv := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()
}

// A structure that implements the Handler interface.
type indexHandler struct{}

// Implement ServeHTTP
func (i *indexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("helo world"))
}
```

## Replacing the Handler
First, replacing the Handler.

ServeHTTP can be replaced with HandlerFunc, which is an alias for a function type.

cf. 
- [func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)](https://golang.org/src/net/http/server.go?s=64180#L2058)

```golang
package main

import (
    "net/http"
)

func main() {
    mux := http.NewServeMux()

	// You can cast a simple function to the HandlerFunc type. This satisfies the Handler interface.
    mux.Handle("/index", http.HandlerFunc(indexHandler))

    s := http.Server{
        Addr:    ":8080",
        Handler: m,
    }
    s.ListenAndServe()
}

// You don't need to prepare a custom structure and implement ServeHTTP.
func indexHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("hello world"))
}
```

## Using DefaultServeMux
Replace `mux` with `DefaultServeMux`.

`DefaultServeMux` contains a structure of type `ServeMux`.

It implements a function called `HandlerFunc` to register routing to the mux.

cf. 
- [DefaultServeMux](https://golang.org/src/net/http/server.go?s=77627:77714#L2269)
- [func (mux *ServeMux) HandlerFunc(pattern string, handler func(ResponseWriter, *Request))](https://golang.org/src/net/http/server.go?s=77627:77714#L2497)

```golang
package main

import (
    "net/http"
)

func main() {
	// No need to create a mux, this alone is sufficient
    http.HandleFunc("/index", indexHandler)

    s := http.Server{
        Addr: ":3000",
		// A default variable in net/http. DefaultServeMux contains a structure of type ServeMux and implements a function called HandlerFunc to register routing to the mux.
        Handler: http.DefaultServeMux,
    }
    s.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("hello world"))
}
```

## Using ListenAndServe()
You can use `ListenAndServe()` without creating a `Server` structure (`http.Server{}`).

cf. 
- [func (*Server) ListenAndServe](https://golang.org/src/net/http/server.go?s=77627:77714#L2898)
- [func ListenAndServe(addr string, handler Handler) error](https://golang.org/src/net/http/server.go?s=77627:77714#L3162)

```golang
package main

import (
    "net/http"
)

func main() {
	http.HandlerFunc("/index", func(w http.ReponseWriter, req *http.Request) {
		w.Write([]byte("hello world"))
	})

	// No need to create a Server structure (http.Server{})
	http.ListenAndServe(":8080")
}
```

This brings us back to the initial basic form.

# Summary
When creating an API server, you might not usually think about these details, but knowing them can be helpful when you want to extend functionality.

When creating a router, you just need to be aware of the `http.Handler` interface and create a mux accordingly.

# Reference
[Go: A Discussion on Custom Router Implementation](https://speakerdeck.com/bmf_san/goterouterzi-zuo-shi-zhuang-ji-rinahua)
