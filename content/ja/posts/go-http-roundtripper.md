---
title: Goのhttp.RoundTripperについて
description: Goのhttp.RoundTripperについてについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: go-http-roundtripper
date: 2023-08-22T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
translation_key: go-http-roundtripper
---


# 概要
Goのhttp.RoundTripperについてかく。

# http.RoundTripperとは
HTTPクライアントの通信を担っているインターフェース。

cf. [pkg.go.dev - net/http#RoundTripper](https://pkg.go.dev/net/http#RoundTripper)

HTTPクライアントにおいてリクエストからレスポンスを受け取るまでの間の処理をカスタマイズすることができる。

HTTPクライアントにおけるミドルウェアというイメージ。

# 実装例
ソースコードは[github.com](https://github.com/bmf-san/go-snippets/blob/master/net/http/round_tripper.go)にも置いてある。

RoundTripperインターフェースを実装して、http.Clientに渡すだけでカスタマイズすることができる。

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

// CustomRoundTripper is a custom implementation of http.RoundTripper
type CustomRoundTripper struct {
	Transport http.RoundTripper
}

func (c *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()

	fmt.Printf("Requesting %s %s\n", req.Method, req.URL)

	resp, err := c.Transport.RoundTrip(req)

	elapsed := time.Since(start)
	fmt.Printf("Received response in %v\n", elapsed)

	return resp, err
}

func main() {
	client := &http.Client{
		Transport: &CustomRoundTripper{
			Transport: http.DefaultTransport,
		},
	}

	resp, err := client.Get("https://www.example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
}
```

```
$ go run main.go
Requesting GET https://www.example.com
Received response in 530.885709ms
Status Code: 200 OK
```

# どこで使う？
HTTPクライアント側でミドルウェア的に何か処理を挟みたいときに使う。

- ログ出力
- 認証ヘッダの付与
- キャッシュ管理
- APIへのリトライ制御、レートリミッター

などHTTPクライアントで統一的な処理を設定したいときに使えそう。

特定のエンドポイントに対して処理を挟みたいときなどは自前のミドルウェアを用意するほうが柔軟性が高そう。

# 参考
- [pkg.go.dev - net/http#RoundTripper](https://pkg.go.dev/net/http#RoundTripper)
- [speakerdeck.com - http.RoundTripper Tips](https://speakerdeck.com/nao_mk2/http-dot-roundtripper-tips?slide=13)
- [qiita.com - Go http.RoundTripper 実装ガイド](https://qiita.com/tutuming/items/6006e1d8cf94bc40f8e8)
- [christina04.hatenablog.com - GoのRoundTripperとTransport](https://christina04.hatenablog.com/entry/2018/05/18/190000)
- [zenn.dev - Goのhttp.RoundTripperでレート制御とリトライの機能を追加する方法](https://zenn.dev/fujisawa33/articles/aef6d266aa751f)
- [blog.lufia.org - Plan 9とGo言語のブログ](https://blog.lufia.org/entry/2018/12/13/000000)


