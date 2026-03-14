---
title: Go1.22rcで変更されるServeMuxの仕様
description: "Go1.22のServeMuxで導入されたHTTPメソッド指定・パスパラメータ・{$}ワイルドカードなどのルーティング拡張仕様を実装例で検証。"
slug: go1-22rc-changes-servemux-spec
date: 2024-01-25T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
translation_key: go1-22rc-changes-servemux-spec
---


この記事は[Makuake Advent Calendar 2023](https://adventar.org/calendars/8992)の19日目の記事です。

なんでこんな大遅刻かというと、唯一埋まっていなかった19日の枠を急遽埋めようと思って投稿したため。~~元々掴んでいた枠ではないので遅刻ではない。~~

# 概要
去年くらいからGoのnet/httpに含まれるServeMuxの機能拡張の提案が出ていてウォッチしていたのだが、最近Closedになったらしい。

cf. [net/http: enhanced ServeMux routing #61410](https://github.com/golang/go/issues/61410)
cf. [GoでServeMuxの機能拡張を提案するProposalがAcceptedになった
アプリケーション](https://bmf-tech.com/posts/Go%e3%81%a7ServeMux%e3%81%ae%e6%a9%9f%e8%83%bd%e6%8b%a1%e5%bc%b5%e3%82%92%e6%8f%90%e6%a1%88%e3%81%99%e3%82%8bProposal%e3%81%8cAccepted%e3%81%ab%e3%81%aa%e3%81%a3%e3%81%9f)

この機能拡張はGo1.22に含まれる想定らしく、Go1.22rc2がリリースされたので試してみた。（Go1.22rc1から含まれているようだが、スルーしていた。）

# Go1.22でのnet/httpの変更点
[tip.golang.org/doc/go1.22#net/http](https://tip.golang.org/doc/go1.22#net/http)に記載されている内容を見ると、以下のような変更点があるようだ。

- ServeFileFSs、FileServerFS、NewFileTransportFSが新規追加
- HTTPサーバーとクライアントが無効なからのContent-Lengthヘッダーを含むリクエストとレスポンスを拒否するように変更
- Request.PathValueとRequest.SetPathValueが新規追加

ここを見た限りだとルーティングの仕様がどんな感じになったのかよくわからないので、ドキュメントの方を参照してみると色々書かれている。

cf. [pkg.go.dev/net/http@go1.22rc2#ServeMux](https://pkg.go.dev/net/http@go1.22rc2#ServeMux)

[Proposal](https://github.com/golang/go/issues/61410)の内容は全て追えていないのだが、Proposalで提案されていたルーティングのパターンマッチングの仕様がサードパーティのルーターのように充実した模様。

目立つところだけ端的にいうと、HTTPメソッド名やパスパラメータを使ったパターンマッチができるようになった感じ。

# Go1.22rc2でServeMuxの新しい機能を試してみる
[go.dev/dl - All releases](https://go.dev/dl/)よりgo1.22rc2をダウンロードして触ってみる。

```go
package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// HTTPメソッドがないパターンは全てのメソッドに一致する。
	// {$}は特別なワイルドカードで、URLの末尾にのみ一致する。
	// /{$}は/のみに一致するが、/は全てのパスに一致してしまう。
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/{$}"))
	})

	// GETメソッドはGETとHEADに一致する。
	// {bar}はワイルドカード
	mux.HandleFunc("GET /foo/{bar}", func(w http.ResponseWriter, r *http.Request) {
		// パスパラメータの取得
		v := r.PathValue("bar")
		w.Write([]byte("GET /foo/" + v))
	})

	// {bar}はワイルドカード
	mux.HandleFunc("POST /foo/{bar}", func(w http.ResponseWriter, r *http.Request) {
		// パスパラメータの取得
		v := r.PathValue("bar")
		w.Write([]byte("POST /foo/" + v))
	})
	http.ListenAndServe(":8080", mux)
}
```

{$}やメソッドとルートを一緒に定義する形はサードパーティのルーターではあんまり見ないかもしれない。（自分が知らないだけかも...）

下位互換性については、[pkg.go.dev/net/http@go1.22rc2#hdr-Compatibility](https://pkg.go.dev/net/http@go1.22rc2#hdr-Compatibility)に記載されている。

Go1.21の仕様を引き継ぎたい場合は、httpmuuxgo121=1をGODEBUG環境変数に設定する必要があるとのこと。

# 所感
自作ルーターである[goblin](https://github.com/bmf-san/goblin)はもはや使う必要がなくなってしまった気がするので、個人プロジェクトは新しいServeMuxを使っていくように切り替えていこうかと思う。

