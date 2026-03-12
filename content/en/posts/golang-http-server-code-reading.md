---
title: Reading Code of Golang HTTP Server
description: 'A review and summary of "Reading Code of Golang HTTP Server", covering key takeaways and practical insights.'
slug: golang-http-server-code-reading
date: 2021-06-30T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Code Reading
translation_key: golang-http-server-code-reading
---

# Overview
I took notes while reading the internal code of an HTTP server when I created a router in Go.

[github.com - bmf-san/goblin](https://github.com/bmf-san/goblin)

# Reading Code of HTTP Server
## Basic Form
This is the typical code structure you would see when you start learning Go.

Various elements have been omitted to arrive at this form.

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
This is the previous basic form written out in detail without omissions. Let's confirm step by step how the implementation leads to the basic form.

```golang
package main

import (
	"net/http"
)

func main() {
	// Multiplexer. A structure for URL matching. Only resolves static routing.
	mux := http.NewServeMux()
	ih := new(indexHandler)
	// Register routing to mux
	mux.Handle("/index", ih)

	srv := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()
}

// Structure that implements the Handler interface.
type indexHandler struct{}

// Implement ServeHTTP
func (i *indexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("helo world"))
}
```

## Replacing Handler
First, let's replace the Handler.

ServeHTTP can be replaced with the function type alias HandlerFunc.

cf. 
- [func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)](https://golang.org/src/net/http/server.go?s=64180#L2058)

```golang
package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Just cast a regular function to HandlerFunc type. It satisfies the Handler interface.
	mux.Handle("/index", http.HandlerFunc(indexHandler))

	s := http.Server{
		Addr: ":8080",
		Handler: m,
	}
	s.ListenAndServe()
}

// No need to implement ServeHTTP with an arbitrary structure.
func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world"))
}
```

## Using DefaultServeMux
You can substitute mux with DefaultServeMux.

DefaultServeMux has a structure of type ServeMux.

It implements a function to register routing to the mux called HandlerFunc.

cf. 
- [DefaultServeMux](https://golang.org/src/net/http/server.go?s=77627:77714#L2269)
- [func (mux *ServeMux) HandlerFunc(pattern string, handler func(ResponseWriter, *Request))](https://golang.org/src/net/http/server.go?s=77627:77714#L2497)

```golang
package main

import (
	"net/http"
)

func main() {
	// No need to create mux, this is enough
	http.HandleFunc("/index", indexHandler)

	s := http.Server{
		Addr: ":3000",
		// The variable that net/http has by default. DefaultServeMux has a structure of type ServeMux. It implements a function to register routing to the mux called HandlerFunc.
		Handler: http.DefaultServeMux,
	}
	s.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world"))
}
```

## Using ListenAndServe()
You can also use ListenAndServe() without creating a Server structure (http.Server{}).

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

# Conclusion
When creating an API server, you might not think about this often, but knowing it can be helpful when you want to extend something.

When creating a router, you should be aware of the http.Handler interface and create a mux accordingly.

# Original Source
[Go Router Custom Implementation Discussion](https://speakerdeck.com/bmf_san/goterouterzi-zuo-shi-zhuang-ji-rinahua)