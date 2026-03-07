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

# Superfluous response.WriteHeader
The following is a naive example, but calling WriteHeader multiple times like this will result in the error `http: superfluous response.WriteHeader call from main.handler`.

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

In this case, the status code from the first call is adopted, and the status code from the last call is ignored. (Is this related to HTTP specifications?)

You can get a sense of the specifications by looking at these references:

- [pkg.go.dev - net/http#ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter)
- [cs.opensource.google - net/http/server.go;l=1149](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=1149;drc=5eb382fc08fb32592e9585f9cb99005696a38b49)

# Handling
In the naive example above, it would be sufficient to adjust the implementation to call WriteHeader only once. However, there may be cases where WriteHeader is set multiple times due to implementation constraints.

```go
import (
	"bytes"
	"html/template"
	"net/http"
)

// Function called to render the error page
func ExecuteTpl(w http.ResponseWriter) error {
	err := template.Must(template.ParseFiles("index.html")).Execute(w, nil)
	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		return err
	}

	return nil
}
```

Assuming that WriteHeader has already been called before this function, it would result in an error.

To avoid such situations, you can write to a buffer first and then call WriteHeader at the end.

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

# Other
It seems that when there is an error during the execution of the template, the execution stops, but writing to the response may have already started. Therefore, it might be better to call WriteHeader before Execute. (Probably)

# References
- [freshman.tech - How to handle template execution errors in Go](https://freshman.tech/snippets/go/template-execution-error/)
- [pod.hatenablog.com - Order to be careful when writing net/http handlers](https://pod.hatenablog.com/entry/2019/01/26/150921)
- [stackoverflow.com - http: superfluous response.WriteHeader call StatusOK](https://stackoverflow.com/questions/68626078/http-superfluous-response-writeheader-call-statusok)
- [medium.com - Dealing with Go Template errors at runtime](https://medium.com/@leeprovoost/dealing-with-go-template-errors-at-runtime-1b429e8b854a)