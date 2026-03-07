---
title: Changes to ServeMux in Go1.22rc
slug: go1-22rc-changes-servemux-spec
date: 2024-01-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
description: Exploring the updates to ServeMux in Go1.22rc and its new routing features.
translation_key: go1-22rc-changes-servemux-spec
---

This article is the 19th entry for the [Makuake Advent Calendar 2023](https://adventar.org/calendars/8992).

The reason for this delay is that I decided to fill the only remaining slot for the 19th day at the last minute. ~~Since it wasn't my original slot, it's technically not late.~~

# Overview
Since last year, there has been a proposal to extend the functionality of Go's `net/http` `ServeMux`, which I had been keeping an eye on. Recently, it seems the proposal was closed.

cf. [net/http: enhanced ServeMux routing #61410](https://github.com/golang/go/issues/61410)  
cf. [Proposal for enhancing ServeMux functionality in Go accepted](https://bmf-tech.com/posts/Go%e3%81%a7ServeMux%e3%81%ae%e6%a9%9f%e8%83%bd%e6%8b%a1%e5%bc%b5%e3%82%92%e6%8f%90%e6%a1%88%e3%81%99%e3%82%8bProposal%e3%81%8cAccepted%e3%81%ab%e3%81%aa%e3%81%a3%e3%81%9f)

This extension is expected to be included in Go1.22, and since Go1.22rc2 has been released, I decided to give it a try. (It seems to have been included since Go1.22rc1, but I missed it.)

# Changes to `net/http` in Go1.22
Looking at the details listed in [tip.golang.org/doc/go1.22#net/http](https://tip.golang.org/doc/go1.22#net/http), the following changes have been made:

- Added `ServeFileFSs`, `FileServerFS`, and `NewFileTransportFS`
- HTTP servers and clients now reject requests and responses with invalid empty `Content-Length` headers
- Added `Request.PathValue` and `Request.SetPathValue`

From this information alone, it's not entirely clear how the routing specifications have changed, so I referred to the documentation, which provides more details.

cf. [pkg.go.dev/net/http@go1.22rc2#ServeMux](https://pkg.go.dev/net/http@go1.22rc2#ServeMux)

I haven't followed all the details of the [Proposal](https://github.com/golang/go/issues/61410), but it seems that the routing pattern matching specifications proposed in the proposal have become as robust as third-party routers.

To highlight the key points, it now supports pattern matching using HTTP method names and path parameters.

# Trying out the new ServeMux features in Go1.22rc2
You can download `go1.22rc2` from [go.dev/dl - All releases](https://go.dev/dl/) and try it out.

```go
package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Patterns without HTTP methods match all methods.
	// {$} is a special wildcard that matches only the end of a URL.
	// /{$} matches only /, but / matches all paths.
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/{$}"))
	})

	// The GET method matches both GET and HEAD.
	// {bar} is a wildcard.
	mux.HandleFunc("GET /foo/{bar}", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve path parameter
		v := r.PathValue("bar")
		w.Write([]byte("GET /foo/" + v))
	})

	// {bar} is a wildcard.
	mux.HandleFunc("POST /foo/{bar}", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve path parameter
		v := r.PathValue("bar")
		w.Write([]byte("POST /foo/" + v))
	})
	http.ListenAndServe(":8080", mux)
}
```

The use of `{$}` and defining methods and routes together might not be common in third-party routers (or maybe I just don't know about it...).

Regarding backward compatibility, details are provided in [pkg.go.dev/net/http@go1.22rc2#hdr-Compatibility](https://pkg.go.dev/net/http@go1.22rc2#hdr-Compatibility).

If you want to retain the behavior of Go1.21, you need to set the `GODEBUG` environment variable to `httpmuuxgo121=1`.

# Thoughts
It feels like there's no longer a need for my custom router, [goblin](https://github.com/bmf-san/goblin). For personal projects, I think I'll start transitioning to the new `ServeMux`.