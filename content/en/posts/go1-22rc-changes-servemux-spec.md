---
title: Changes to ServeMux Specification in Go 1.22rc
description: 'Explore Go 1.22 enhanced ServeMux routing: HTTP method patterns (GET /items), wildcard path params (/items/{id}), the /{$} exact-match wildcard, and Request.PathValue.'
slug: go1-22rc-changes-servemux-spec
date: 2024-01-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
translation_key: go1-22rc-changes-servemux-spec
---

This article is the 19th entry in the [Makuake Advent Calendar 2023](https://adventar.org/calendars/8992).

The reason for this significant delay is that I decided to fill in the only unoccupied slot for the 19th at the last minute. ~~Since this was not originally my assigned slot, it's not really a delay.~~

# Overview
Since last year, there have been proposals to extend the functionality of ServeMux included in Go's net/http, and I have been watching these developments, but it seems they have recently been closed.

cf. [net/http: enhanced ServeMux routing #61410](https://github.com/golang/go/issues/61410)
cf. [Proposal to extend ServeMux functionality in Go has been Accepted](https://bmf-tech.com/posts/Go%e3%81%a7ServeMux%e3%81%ae%e6%a9%9f%e8%83%bd%e6%8b%a1%e5%bc%b5%e3%82%92%e6%8f%90%e6%a1%88%e3%81%99%e3%82%8bProposal%e3%81%8cAccepted%e3%81%ab%e3%81%aa%e3%81%a3%e3%81%9f)

This extension is expected to be included in Go 1.22, and since Go 1.22rc2 has been released, I decided to give it a try. (It seems to have been included since Go 1.22rc1, but I had overlooked it.)

# Changes in net/http for Go 1.22
According to the information listed at [tip.golang.org/doc/go1.22#net/http](https://tip.golang.org/doc/go1.22#net/http), the following changes appear to have been made:

- New additions: ServeFileFSs, FileServerFS, NewFileTransportFS
- The HTTP server and client have been modified to reject requests and responses containing invalid Content-Length headers.
- New additions: Request.PathValue and Request.SetPathValue

From what I can see here, I don't quite understand how the routing specifications have changed, so I referred to the documentation where various details are written.

cf. [pkg.go.dev/net/http@go1.22rc2#ServeMux](https://pkg.go.dev/net/http@go1.22rc2#ServeMux)

I haven't been able to follow all the details of the [Proposal](https://github.com/golang/go/issues/61410), but it seems that the routing pattern matching specifications proposed in the Proposal have become as rich as those of third-party routers.

To put it succinctly, it seems that pattern matching using HTTP method names and path parameters has become possible.

# Trying Out the New Features of ServeMux in Go 1.22rc2
I downloaded go1.22rc2 from [go.dev/dl - All releases](https://go.dev/dl/) and started experimenting.

```go
package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Patterns without HTTP methods match all methods.
	// {$} is a special wildcard that only matches the end of the URL.
	// /{$} matches only /, but / matches all paths.
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/{$}"))
	})

	// GET method matches GET and HEAD.
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

The combination of {$} and defining methods with routes might not be commonly seen in third-party routers. (Maybe I'm just not aware of it...)

Regarding backward compatibility, it is mentioned in [pkg.go.dev/net/http@go1.22rc2#hdr-Compatibility](https://pkg.go.dev/net/http@go1.22rc2#hdr-Compatibility).

If you want to maintain the specifications of Go 1.21, you need to set httpmuuxgo121=1 in the GODEBUG environment variable.

# Thoughts
I feel that I no longer need to use my custom router [goblin](https://github.com/bmf-san/goblin), so I plan to switch my personal projects to use the new ServeMux.