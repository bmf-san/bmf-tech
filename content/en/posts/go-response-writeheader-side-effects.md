---
title: Side Effects of Go's response.WriteHeader
slug: go-response-writeheader-side-effects
date: 2023-09-11T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Tips
translation_key: go-response-writeheader-side-effects
---

# Overview
This is a note about the side effects of Go's response.WriteHeader.

# superfluous response.WriteHeader
The following is a straightforward example, but if you call WriteHeader multiple times as shown below, you will get an error: `http: superfluous response.WriteHeader call from main.handler`.

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusInternalServerError) // Error!
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
```

When called multiple times like this, the status code from the first call is adopted, and the status code from the last call is ignored. (Is it related to HTTP specifications?)

You can get a sense of the specifications by looking at these.

- [pkg.go.dev - net/http#ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter)
- [cs.opensource.google - net/http/server.go;l=1149](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=1149;drc=5eb382fc08fb32592e9585f9cb99005696a38b49)

# Solution
In the straightforward example above, you can simply adjust the implementation to call WriteHeader only once. However, let's assume there are cases where WriteHeader is set multiple times due to implementation reasons.

```go
import (
	"bytes"
	"html/template"
	"net/http"
)

// Function called for rendering error pages
func ExecuteTpl(w http.ResponseWriter) error {
	err := template.Must(template.ParseFiles("index.html")).Execute(w, nil)
	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		return err
	}

	return nil
}
```

Suppose WriteHeader has already been called before this function, causing an error.

To avoid errors in such situations, you can write to a buffer first and then call WriteHeader at the end.

```go
package main

import (
	"bytes"
	"html/template"
	"net/http"
)

func ExecuteTpl(w http.ResponseWriter) error {
	var buf bytes.Buffer

	w.WriteHeader(http.StatusInternalServerError)
	err := template.Must(template.ParseFiles("index.html")).Execute(w, nil)
	if err != nil {
		return err
	}

	buf.WriteTo(w)

	return nil
}
```

# Others
When there is an error during the call to template's Execute, execution stops, but it seems that writing to the response may partially start, so it seems better to call WriteHeader before Execute. (Probably)

# References
- [freshman.tech - How to handle template execution errors in Go](https://freshman.tech/snippets/go/template-execution-error/)
- [pod.hatenablog.com - net/httpのhandlerを書く時に気をつけたほうが良い順序について](https://pod.hatenablog.com/entry/2019/01/26/150921)
- [stackoverflow.com - http: superfluous response.WriteHeader call StatusOK](https://stackoverflow.com/questions/68626078/http-superfluous-response-writeheader-call-statusok)
- [medium.com - Dealing with Go Template errors at runtime](https://medium.com/@leeprovoost/dealing-with-go-template-errors-at-runtime-1b429e8b854a)
