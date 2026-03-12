---
title: About Go's http.RoundTripper
description: 'An in-depth look at About Go''s http.RoundTripper, covering key concepts and practical insights.'
slug: go-http-roundtripper
date: 2023-08-22T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
translation_key: go-http-roundtripper
---

# Overview
Writing about Go's http.RoundTripper.

# What is http.RoundTripper?
An interface responsible for HTTP client communication.

cf. [pkg.go.dev - net/http#RoundTripper](https://pkg.go.dev/net/http#RoundTripper)

It allows customization of the process from sending a request to receiving a response in an HTTP client.

Think of it as middleware for the HTTP client.

# Implementation Example
The source code is also available on [github.com](https://github.com/bmf-san/go-snippets/blob/master/net/http/round_tripper.go).

You can customize by implementing the RoundTripper interface and passing it to http.Client.

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

# Where to Use?
Use it when you want to insert some middleware-like processing on the HTTP client side.

- Log output
- Adding authentication headers
- Cache management
- Retry control and rate limiter for APIs

It seems useful when you want to set up uniform processing in an HTTP client.

When you want to insert processing for specific endpoints, it might be more flexible to prepare your own middleware.

# References
- [pkg.go.dev - net/http#RoundTripper](https://pkg.go.dev/net/http#RoundTripper)
- [speakerdeck.com - http.RoundTripper Tips](https://speakerdeck.com/nao_mk2/http-dot-roundtripper-tips?slide=13)
- [qiita.com - Go http.RoundTripper Implementation Guide](https://qiita.com/tutuming/items/6006e1d8cf94bc40f8e8)
- [christina04.hatenablog.com - Go's RoundTripper and Transport](https://christina04.hatenablog.com/entry/2018/05/18/190000)
- [zenn.dev - How to Add Rate Control and Retry Features with Go's http.RoundTripper](https://zenn.dev/fujisawa33/articles/aef6d266aa751f)
- [blog.lufia.org - Blog on Plan 9 and Go Language](https://blog.lufia.org/entry/2018/12/13/000000)
