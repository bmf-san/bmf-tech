---
title: How to Use pprof Without DefaultServeMux
slug: using-pprof-without-defaultmux
date: 2023-04-30T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Tips
translation_key: using-pprof-without-defaultmux
---

# Overview
This is a note on how to use [net/http/pprof](https://pkg.go.dev/net/http/pprof) with routers other than DefaultServeMux (Go's standard router).

# Pitfalls
Simply doing a blank import of pprof is not enough.

```go
package main

import (
    _ "net/http/pprof"
)
```

When using a router other than DefaultServeMux, just a blank import will not enable pprof.

Referring to [net/http/pprof](https://pkg.go.dev/net/http/pprof), it states:

> If you are not using DefaultServeMux, you will have to register handlers with the mux you are using.

# Solution
Below is an example using my custom router [bmf-san/goblin](https://github.com/bmf-san/goblin).

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

As shown above, you need to set up routing yourself and configure the pprof handlers.

For httprouter, the following issue may be helpful.
[pprof issue with httprouter #236](https://github.com/julienschmidt/httprouter/issues/236)

# Aside
I got stuck when trying to set up Go profiling in a pull-based manner with [pyroscope](https://pyroscope.io/).