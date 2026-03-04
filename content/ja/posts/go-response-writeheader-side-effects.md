---
title: "Goのresponse.WriteHeaderの副作用について"
slug: "go-response-writeheader-side-effects"
date: 2023-09-11
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "Tips"
draft: false
---

# 概要
Goのresponse.WriteHeaderの副作用についてメモする。

# superfluous response.WriteHeader
以下は愚直な例だが、次のようにWriteHeaderを複数回呼ぶと、`http: superfluous response.WriteHeader call from main.handler`というエラーが出る。

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

このように複数回呼ぶと、最初の呼び出しのステータスコードが採用され、最後の呼び出しのステータスコードは無視される。（HTTPの仕様に関係している？）

この辺を見ると仕様の雰囲気が掴める。

- [pkg.go.dev - net/http#ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter)
- [cs.opensource.google - net/http/server.go;l=1149](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=1149;drc=5eb382fc08fb32592e9585f9cb99005696a38b49)

# 対応
上記の愚直な例では、素直にWriteHeaderを一度だけ呼び出すような実装に調整すれば良いだけだが、例えば実装上の都合でWriteHeaderが複数回セットされてしまう場合があるとする。

```go
import (
	"bytes"
	"html/template"
	"net/http"
)

// エラーページの描画のために呼び出される関数
func ExecuteTpl(w http.ResponseWriter) error {
	err := template.Must(template.ParseFiles("index.html")).Execute(w, nil)
	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		return err
	}

	return nil
}
```

この関数の前に既にWriteHeaderが呼び出され、エラーが出てしまう状況だとする。

このような状況でエラーを回避するには、バッファに一度書き出してから、最後にWriteHeaderを呼び出すようにすると回避できる。

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

# その他
templateのExecuteは呼び出し時にエラーがあった際、実行は停止するがレスポンスへの書き込みが一部開始する可能性があるらしく、Executeよりも前にWriteHeaderしたほうが良さそう。（多分）

# 参考
- [freshman.tech - How to handle template execution errors in Go](https://freshman.tech/snippets/go/template-execution-error/)
- [pod.hatenablog.com - net/httpのhandlerを書く時に気をつけたほうが良い順序について](https://pod.hatenablog.com/entry/2019/01/26/150921)
- [stackoverflow.com - http: superfluous response.WriteHeader call StatusOK](https://stackoverflow.com/questions/68626078/http-superfluous-response-writeheader-call-statusok)
- [medium.com - Dealing with Go Template errors at runtime](https://medium.com/@leeprovoost/dealing-with-go-template-errors-at-runtime-1b429e8b854a)
