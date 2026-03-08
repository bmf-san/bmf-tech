---
title: Contextual Logging with slog
slug: slog-contextual-logging
date: 2023-10-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - slog
  - contextual logging
translation_key: slog-contextual-logging
---



# Overview
Summarizing contextual logging using log/slog in Go.

# What is log/slog?
A package for structured logging added in Go1.21.

Structured logging is the output of logs as structured data.

Previously, to perform structured logging in Go, one had to either use third-party packages or implement it from scratch. Now, the standard package is also becoming a viable option.

By default, it can output in text or JSON format.

slog can hold context, allowing you to include request-based information.

# Contextual logging with slog
Let's write code to create a Logger that includes request-based information in the logs.

Below is the code that assumes logging with a trace ID in the context.

By implementing the slog.Handler interface, you can create your own Handler.[]

```go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// TraceIDHandler represents the singular of trace id handler.
type TraceIDHandler struct {
	slog.Handler
}

type ctxTraceID struct{}

var ctxTraceIDKey = ctxTraceID{}

// Handle implements slog.Handler interface.
func (t TraceIDHandler) Handle(ctx context.Context, r slog.Record) error {
	tid, ok := ctx.Value(ctxTraceIDKey).(string)
	if ok {
		r.AddAttrs(slog.String("trace_id", tid))
	}
	return t.Handler.Handle(ctx, r)
}

// WithTraceID returns a context with a trace id.
func WithTraceID(ctx context.Context) context.Context {
	uuid, _ := uuid.NewRandom()
	return context.WithValue(ctx, ctxTraceIDKey, uuid.String())
}

func main() {
	mux := http.NewServeMux()

	handler := TraceIDHandler{slog.NewJSONHandler(os.Stdout, nil)}
	logger := slog.New(handler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := WithTraceID(r.Context())
		logger.InfoContext(ctx, "Log with TraceID")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Contextual Logging!")
	})

	http.ListenAndServe(":8085", mux)
}
```

When you run the above code, the following log is output:

```sh
{"time":"2023-10-08T17:06:44.423859+09:00","level":"INFO","msg":"Log with TraceID","trace_id":"4f9a0bb6-cf8d-4eef-82ea-2385b76d3a74"}
```

# Impressions
I was using my own package for structured logging, but with the release of Go1.21, I switched to this.

# References
- [pkg.go.dev - slog](https://pkg.go.dev/golang.org/x/exp/slog)
- [future-architect.github.io - Starting Go 1.21 Series & How to Use slog](https://future-architect.github.io/articles/20230731a/)
- [gihyo.jp - Experience Go's New Structured Logger](https://gihyo.jp/article/2023/02/tukinami-go-04)
- [zenn.dev - Tried the Official Go Structured Logger (Proposed) slog](https://zenn.dev/mizutani/articles/golang-exp-slog)
