---
title: About Go's http.RoundTripper
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
This post is about Go's http.RoundTripper.

# What is http.RoundTripper?
An interface responsible for handling HTTP client communications.

cf. [pkg.go.dev - net/http#RoundTripper](https://pkg.go.dev/net/http#RoundTripper)

It allows customization of the processing between sending a request and receiving a response in an HTTP client.

Think of it as middleware in an HTTP client.

# Implementation Example
The source code is also available on [github.com](https://github.com/bmf-san/go-snippets/blob/master/net/http/round_tripper.go).

You can customize it simply by implementing the RoundTripper interface and passing it to http.Client.

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

# Where to Use It?
Use it when you want to insert some processing as middleware on the HTTP client side.

- Logging
- Adding authentication headers
- Cache management
- Retry control for APIs, rate limiting

It seems useful when you want to set unified processing in an HTTP client.

If you want to insert processing for specific endpoints, preparing your own middleware might offer more flexibility.

# References
- [pkg.go.dev - net/http#RoundTripper](https://pkg.go.dev/net/http#RoundTripper)
- [speakerdeck.com - http.RoundTripper Tips](https://speakerdeck.com/nao_mk2/http-dot-roundtripper-tips?slide=13)
- [qiita.com - Go http.RoundTripper Implementation Guide](https://qiita.com/tutuming/items/6006e1d8cf94bc40f8e8)
- [christina04.hatenablog.com - Go's RoundTripper and Transport](https://christina04.hatenablog.com/entry/2018/05/18/190000)
- [zenn.dev - How to Add Rate Control and Retry Functionality with Go's http.RoundTripper](https://zenn.dev/fujisawa33/articles/aef6d266aa751f)
- [blog.lufia.org - Plan 9 and Go Language Blog](https://blog.lufia.org/entry/2018/12/13/000000)