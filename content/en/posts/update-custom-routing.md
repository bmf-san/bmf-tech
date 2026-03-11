---
title: Updated My Custom Router
slug: update-custom-routing
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
Recently, I updated my custom router, [goblin](https://github.com/bmf-san/goblin), and I wanted to document the changes.

Here are some past articles about routing. There are also articles from the implementation consideration phase, but the content wasn't very good, so I'll omit them.

- [URL Routing from Scratch Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E8%87%AA%E4%BD%9C%E5%85%A5%E9%96%80%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%89%EF%BC%91)
- [URL Routing from Scratch Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E8%87%AA%E4%BD%9C%E5%85%A5%E9%96%80%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%89%EF%BC%92)
- [Code Reading of Golang's HTTP Server](https://bmf-tech.com/posts/Golang%E3%81%AEHTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%83%AA%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0)
- [Introduction to URL router from scratch with Golang](https://dev.to/bmf_san/introduction-to-url-router-from-scratch-with-golang-3p8j)

# What Has Changed?
I had released a version with basic functionality as [1.0.0](https://github.com/bmf-san/goblin/releases/tag/1.0.0). While using it myself, I found bugs and felt the lack of features, leading to several backward-incompatible changes (a result of haphazard implementation), and now the latest version is [5.0.1](https://github.com/bmf-san/goblin/releases/tag/5.0.1).

The most significant change is the support for middleware, which led to a review of the internal data structure, DSL, and bug fixes.

# Why Support Middleware?
I thought middleware could be freely handled by the router users, but there were constraints.

Even if middleware was implemented by the user, it would execute after the routing matching process (i.e., whether the path and HTTP method match the registered routing), which imposed a restriction that middleware couldn't be applied before HTTP method matching.

This was inconvenient when handling preflight requests (like CORS), so I decided to support middleware to fundamentally solve this issue.

The tricky part in considering such cases was the internal data structure of the routing, which was based on the assumption of path and HTTP method matching, so it needed to be reviewed.

Therefore, I changed the data structure as follows and implemented middleware support.

Before
![Based on trie tree](/assets/images/posts/update-custom-routing/70862745-7148e180-1f83-11ea-85d3-2cd8fb4db0d3.png "Based on trie tree")

After
![after](/assets/images/posts/update-custom-routing/117675761-d4c25780-b1e7-11eb-9ec7-e78ac0ce142b.png)

# Benchmark
I had written benchmarks for static routing only, but I wanted to compare with other libraries in dynamic routing tests, so I used the most comprehensive [github.com - julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) for benchmark testing.

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

I have submitted a PR to add goblin's benchmark support.
[github.com - julienschmidt/go-http-routing-benchmark Add a new router goblin #97](https://github.com/julienschmidt/go-http-routing-benchmark/pull/97)

# Thoughts
I feel like it has finally become a decent router. There are still many areas for improvement, so I will continue to maintain it.