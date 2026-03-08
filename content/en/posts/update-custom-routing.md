---
title: Updated My Custom Router
slug: update-custom-routing
image: /assets/images/posts/post-259/117675761-d4c25780-b1e7-11eb-9ec7-e78ac0ce142b.png
date: 2021-06-18T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - OSS
  - Router
translation_key: update-custom-routing
---

# Overview
Recently, I updated my custom router [goblin](https://github.com/bmf-san/goblin), and I wanted to document the changes.

Below are articles I previously wrote about the past routing. There are also articles from the implementation consideration phase, but I will omit them as the content is not very good.

- [Introduction to Custom URL Routing Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E8%87%AA%E4%BD%9C%E5%85%A5%E9%96%80%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%89%EF%BC%91)
- [Introduction to Custom URL Routing Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E8%87%AA%E4%BD%9C%E5%85%A5%E9%96%80%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%89%EF%BC%92)
- [Code Reading of Golang's HTTP Server](https://bmf-tech.com/posts/Golang%E3%81%AEHTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%83%AA%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0)
- [Introduction to URL Router from Scratch with Golang](https://dev.to/bmf_san/introduction-to-url-router-from-scratch-with-golang-3p8j)

# What Changed?
I had released a version with basic functionality as [1.0.0](https://github.com/bmf-san/goblin/releases/tag/1.0.0). While using it myself, I found bugs and felt some features were lacking, leading to several breaking changes (the consequences of a haphazard implementation). The latest version is now [5.0.1](https://github.com/bmf-san/goblin/releases/tag/5.0.1).

The most significant change is the support for middleware functionality. This required a review of the internal data structures, a revision of the DSL, and bug fixes.

# Why Support Middleware?
I initially thought that middleware could be freely handled by the users of the router, but there were actually constraints.

Even if middleware is implemented on the user side, it would be executed after the routing matching process (i.e., checking if the path and HTTP method match the registered routes), which created a limitation for cases where middleware needed to be applied before HTTP method matching.

This was inconvenient when handling Preflight requests (for CORS support), so I decided to support middleware to fundamentally resolve this issue.

The troublesome part in considering such cases was the internal data structure of the routing, which was designed to assume a match between path and HTTP method, necessitating a review of that structure.

Thus, I changed the data structure as follows to implement middleware support.

Before
![Based on trie tree](https://user-images.githubusercontent.com/13291041/70862745-7148e180-1f83-11ea-85d3-2cd8fb4db0d3.png "Based on trie tree")

After
![after](https://user-images.githubusercontent.com/13291041/117675761-d4c25780-b1e7-11eb-9ec7-e78ac0ce142b.png)

# Benchmark
I had written benchmarks only for static routing, but I wanted to compare it with other libraries using dynamic routing tests, so I conducted benchmark tests using the most comprehensive one I found, [github.com - julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark).

Here are the latest scores.
```
#GithubAPI Routes: 203
   Goblin: 80864 Bytes

#GPlusAPI Routes: 13
   Goblin: 7856 Bytes

#ParseAPI Routes: 26
   Goblin: 8688 Bytes

#Static Routes: 157
   Goblin: 34488 Bytes

goos: darwin
goarch: amd64
pkg: github.com/julienschmidt/go-http-routing-benchmark
cpu: Intel(R) Core(TM) i5-8210Y CPU @ 1.60GHz
BenchmarkGoblin_Param             738289              1964 ns/op             128 B/op          4 allocs/op
BenchmarkGoblin_Param5            754988              1920 ns/op             368 B/op          6 allocs/op
BenchmarkGoblin_Param20            56145             23260 ns/op            3168 B/op         58 allocs/op
BenchmarkGoblinWeb_ParamWrite     304082              4610 ns/op             648 B/op         11 allocs/op
BenchmarkGoblin_GithubStatic     1156518              2745 ns/op             128 B/op          4 allocs/op
BenchmarkGoblin_GithubParam       125570              9985 ns/op             816 B/op         15 allocs/op
BenchmarkGoblin_GithubAll           2232            622376 ns/op           49424 B/op       1018 allocs/op
BenchmarkGoblin_GPlusStatic      1000000              1298 ns/op              80 B/op          3 allocs/op
BenchmarkGoblin_GPlusParam        417717              2893 ns/op             664 B/op         11 allocs/op
BenchmarkGoblin_GPlus2Params      274990              4551 ns/op             824 B/op         15 allocs/op
BenchmarkGoblin_GPlusAll           95580             14536 ns/op            2208 B/op         57 allocs/op
BenchmarkGoblin_ParseStatic      1651083               707.0 ns/op           128 B/op          4 allocs/op
BenchmarkGoblin_ParseParam        413840              2876 ns/op             728 B/op         12 allocs/op
BenchmarkGoblin_Parse2Params      260120              4119 ns/op             808 B/op         15 allocs/op
BenchmarkGoblin_ParseAll           54518             21692 ns/op            4656 B/op        120 allocs/op
BenchmarkGoblin_StaticAll          26689             46104 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/julienschmidt/go-http-routing-benchmark      37.270s
```

I have submitted a PR to add benchmark support for goblin.
[github.com - julienschmidt/go-http-routing-benchmark Add a new router goblin #97](https://github.com/julienschmidt/go-http-routing-benchmark/pull/97)

# Thoughts
I finally feel like I have a router that is on par with others. There are still many areas for improvement, so I will continue to maintain it.