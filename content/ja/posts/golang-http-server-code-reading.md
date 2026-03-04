---
title: "GolangのHTTPサーバーのコードリーディング"
slug: "golang-http-server-code-reading"
date: 2021-06-30
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Golang"
  - "コードリーディング"
draft: false
---

# 概要
Goでrouterを作ったときにHTTPサーバーのコードの内部を読んだので、その時のメモ。

[github.com - bmf-san/goblin](https://github.com/bmf-san/goblin)

# HTTPサーバーのコードリーディング
## 基本形
Goに入門したとによく見るであろう形のコード。

色々なものが省略されてこの形になっている。

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

## 省略しないで丁寧に書いた形
先程の基本形を省略しないで書いた形。
どのような実装を経て基本形になるのか１つずつ確認していく。

```golang
package main

import (
	"net/http"
)

func main() {
	// マルチプレクサ。URLマッチングをするための構造体。静的なルーティングのみ解決する。
	mux := http.NewServeMux()
	ih := new(indexHandler)
	// muxにルーティングを登録する
	mux.Handle("/index", ih)

	srv := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()
}

// Handlerインターフェースを実装した構造体。
type indexHandler struct{}

// ServeHTTPを実装
func (i *indexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("helo world"))
}
```

## Handlerの置き換え
まずは、Handlerの置き換え。

ServeHTTPは関数型のaliasであるHandlerFuncに置き換えることができる。

cf. 
- [func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)](https://golang.org/src/net/http/server.go?s=64180#L2058)

```golang

package main

import (
    "net/http"
)

func main() {
    mux := http.NewServeMux()

	// ただの関数をHandlerFunc型にキャストすれば良い。Handlerインターフェースを満たせる。
    mux.Handle("/index", http.HandlerFunc(indexHandler))

    s := http.Server{
        Addr:    ":8080",
        Handler: m,
    }
    s.ListenAndServe()
}

// 適当な構造体を用意して、ServeHTTPを実装しなくても良い。
func indexHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("hello world"))
}
```

## DefaultServeMuxの利用
muxをDefaultServeMuxで代用。

DefaultServeMuxはServeMux型の構造体を持っている。

HandlerFuncというmuxにルーティングを登録する関数を実装している。

cf. 
- [DefaultServeMux](https://golang.org/src/net/http/server.go?s=77627:77714#L2269)
- [func (mux *ServeMux) HandlerFunc(pattern string, handler func(ResponseWriter, *Request))](https://golang.org/src/net/http/server.go?s=77627:77714#L2497)

```golang
package main

import (
    "net/http"
)

func main() {
	// muxを作らなくてもこれだけでOK
    http.HandleFunc("/index", indexHandler)

    s := http.Server{
        Addr: ":3000",
		// net/httpがデフォルトで持っている変数。DefaultServeMuxはServeMux型の構造体を持っている。HandlerFuncというmuxにルーティングを登録する関数を実装している。
        Handler: http.DefaultServeMux,
    }
    s.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("hello world"))
}
```

## ListenAndServe()の利用
Server構造体（http.Server{}）を作らずとも、ListenAndServe()を代用することができる。

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

	// Server構造体（http.Server{}）を作らなくても大丈夫
	http.ListenAndServe(":8080")
}
```

これで最初の基本形に到達。

# まとめ
APIサーバー作るときとか普段余り意識しないと思うが、知っておくと何か拡張したいときに役に立つ、はず。

routerを作るときはhttp.Handlerのインターフェースを意識してmuxを作って上げれば良い。

# 元ネタ
[GoでRouter自作実装寄りな話](https://speakerdeck.com/bmf_san/goterouterzi-zuo-shi-zhuang-ji-rinahua)
