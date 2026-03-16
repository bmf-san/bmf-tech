---
title: slogを使ったContextual Logging
description: slogを使ったContextual Logging
slug: slog-contextual-logging
date: 2023-10-08T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - slog
  - contextual logging
translation_key: slog-contextual-logging
---


# 概要
Goでlog/slogを使ったcontextual loggingについてまとめる。

# log/slogとは
Go1.21で追加された構造化ロギングのためのパッケージ。

構造化ロギングとは、ログを構造化されたデータとして出力すること。

今まではGoで構造化ロギングをする場合はサードパーティのパッケージを利用するか、自前でスクラッチするかしか手段がなかったが、今後は標準パッケージも視野に入るようになった。

デフォルトではテキスト形式またはJson形式で出力することができる。

slogはcontextを持たせることができるため、リクエストベースの情報を詰めることもできる。

# slogでcontextual logging
ログにリクエストベースの情報を持たせるLoggerを作成するコードを書いてみる。

以下は、contextにトレースIDを持たせてログ出力を想定したコード。

slog.Handlerインターフェースを実装することで、自前のHandlerを作成することができる。[]

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

上記のコードを実行すると、次のようなログが出力される。

```sh
{"time":"2023-10-08T17:06:44.423859+09:00","level":"INFO","msg":"Log with TraceID","trace_id":"4f9a0bb6-cf8d-4eef-82ea-2385b76d3a74"}
```

# 所感
構造化ロギングには自前のパッケージを使っていたが、go1.21のリリースとともに差し替えをした。

# 参照
- [pkg.go.dev - slog](https://pkg.go.dev/golang.org/x/exp/slog)
- [future-architect.github.io - Go 1.21連載始まります＆slogをどう使うべきか](https://future-architect.github.io/articles/20230731a/)
- [gihyo.jp - Goの新しい構造化ロガーを体験しよう](https://gihyo.jp/article/2023/02/tukinami-go-04)
- [zenn.dev - Go公式の構造化ロガー（として提案されている）slogを触ってみたメモ](https://zenn.dev/mizutani/articles/golang-exp-slog)

