---
title: From Custom HTTP Router to New ServeMux
description: 'Migrate from a custom HTTP router goblin to Go 1.22 enhanced net/http ServeMux. Covers new routing patterns, performance comparison with third-party routers, and when ServeMux is enough.'
slug: custom-http-router-to-new-servemux
date: 2024-04-27T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Router
translation_key: custom-http-router-to-new-servemux
---

# Overview
I had been using a custom HTTP router called [goblin](https://github.com/bmf-san/goblin) in my application, but since the features of ServeMux were enhanced in Go 1.22, I have started using ServeMux.

In this article, I will summarize the features and performance of ServeMux added in Go 1.22 and consider the selection of HTTP routers in Go moving forward.

# Features of ServeMux Added in Go 1.22
When the Go 1.22rc was released, I researched the new features of ServeMux, and I thought I would delve into a bit more detail.

cf. [Changes to ServeMux in Go 1.22rc](https://bmf-tech.com/posts/Go1.22rc%e3%81%a7%e5%a4%89%e6%9b%b4%e3%81%95%e3%82%8c%e3%82%8bServeMux%e3%81%ae4%bb%95%e6%a7%98)

Based on the following reference information, I will organize the new features of ServeMux.

- Release Notes
  - [Go 1.22 Release Notes - Enhanced routing patterns](https://tip.golang.org/doc/go1.22#enhanced_routing_patterns)
    - It is stated that the patterns of ServeMux have been expanded to accept methods and wildcards (dynamic path parameters, e.g., /items/{id}).
- pkg.go.dev
  - [pkg.go.dev - net/http](https://pkg.go.dev/net/http)
- go.dev
  - [go.dev - Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements)
    - Specifications regarding the new features of ServeMux are described.
- Discussion
  - [net/http: add methods and path variables to ServeMux patterns](https://github.com/golang/go/discussions/60227)
    - Discussion about the enhancement of ServeMux's functionality.
- Proposal
  - [net/http: enhanced ServeMux routing](https://github.com/golang/go/issues/61410)
    - Proposal regarding the enhancement of ServeMux's functionality.

## Definition of Routing by HTTP Method
By specifying a path that includes the HTTP method, it has become possible to define routing by HTTP method. When using ServeMux, there is no need to write conditional branches for HTTP methods within the handler.

```go
http.HandleFunc("GET /items", handleItems)
```

Since there are constants for HTTP methods (e.g., http.MethodGet), I thought it would be better to use them, but I suspect that this format was chosen to maintain backward compatibility and not break the existing method signatures.

Alternatively, it may simply be aligned with the format of HTTP requests.

## Definition of Routing by Wildcard
By specifying a path using a wildcard (`{pathVal}`), it has become possible to define routing by wildcard.

```go
// Matches GET /items/1 or /items/foo
http.HandleFunc("GET /items/{id}", handleItems)
```

Path patterns can be specified in the following format. Even before Go 1.22, it was possible to specify the hostname as well (I learned this quite recently...).

```
[METHOD ][HOST]/[PATH]
```

Values that match the wildcard can be obtained using the http.Request's [PathValue method](https://pkg.go.dev/net/http#Request.PathValue).

```go
func handleItems(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    // id will contain the value that matched the wildcard
}
```

Additionally, it is also possible to define routing using multi-wildcards.

```go
// Matches GET /items/1, /items/1/2, /itema/1/2/3
http.HandleFunc("GET /items/{id...}", handleItems)
```

If you want to match the path exactly, use `{$}`. This is something to keep in mind when defining the root (`/`).

```go
http.HandleFunc("GET /{$}", handleIndex)
```

As a side note, some third-party routers use `*` as a wildcard in their patterns. I personally prefer to refer to them as path parameters to avoid being influenced by that image... (though it's not something to worry about).

# Points to Note with the New ServeMux
## Definition of Patterns Using HTTP Methods
Many third-party HTTP routers often treat HTTP methods as methods.

```go
// ServeMux
http.HandleFunc("GET /items/{id}", handleItems)

// Example from third-party
mux.Get("/items", handleItems)
```

In ServeMux, since the HTTP method is included in the pattern, there may be a minor issue where typos are harder to detect. It would be nice if it could be checked with a linter, but I wonder if that is possible. It might be a good topic to create a static analysis tool for. I feel like something will eventually be done about it.

## Priority Rules
When defining routing using wildcards, it is important to pay attention to priority.

cf.
- [pkg.go.dev - Precedence](https://pkg.go.dev/net/http#hdr-Precedence)
- [go.dev - Routing Enhancements for Go 1.22 precedence](https://go.dev/blog/routing-enhancements#precedence)

In ServeMux, routing patterns can be defined as follows:

```go
// If both match, the former takes precedence
/items/new
/items/{id}

// If both match, the latter takes precedence
/items/{id...}
/items/{id}/category/{name}
```

In such patterns, it is necessary to be careful about which pattern the expected request will match.

Some HTTP routers in the world do not allow such pattern duplication, while others do, and ServeMux falls into the latter category.

On the other hand, in cases where conflicts arise, ServeMux detects the conflict and causes a panic.

```go
// There are cases where both match, as well as cases where only one matches
/items/{id}
/{category}/items

// Cases where both match
/items/{id}
/items/{name}
```

In the case of conflicts, errors are likely to be detected early (during testing or server startup), so they are not as troublesome as duplication.

By the way, in my custom [goblin](https://github.com/bmf-san/goblin), the first registered pattern takes precedence, resulting in a first-come-first-served specification (which is actually quite complicated due to poor design...).

For more details on ServeMux's conflict detection, refer to the following article.

cf. [rhumie.github.io - ServeMux Conflict Detection and Performance](https://rhumie.github.io/go122party/1)

Even third-party HTTP routers consider conflict detection. For example, in [httprouter](https://github.com/julienschmidt/httprouter), the specification is such that a request can either match one pattern or none at all.

> Only explicit matches: With other routers, like http.ServeMux, a requested URL path could match multiple patterns. Therefore they have some awkward pattern priority rules, like longest match or first registered, first matched. By design of this router, a request can only match exactly one or no route. As a result, there are also no unintended matches, which makes it great for SEO and improves the user experience.

cf. https://github.com/julienschmidt/httprouter?tab=readme-ov-file#features

The specifications of priority in matching HTTP router patterns are an important point that influences the quality of HTTP routers, so it is reassuring that this is well designed.

## Backward Compatibility
There are cases where backward compatibility is not maintained between Go 1.22 and Go 1.21.

cf. [pkg.go.dev - Compatibility](https://pkg.go.dev/net/http#hdr-Compatibility)

In such cases, you can revert to the behavior of Go 1.21 by setting the GODEBUG environment variable to `httpmuxgo121=1`.

# Internal Implementation Code Reading
## Routing Pattern Registration Process
- [register](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=2735;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1)
- [parsePattern](https://cs.opensource.google/go/go/+/master:src/net/http/pattern.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;l=84)

Below is the definition of the ServeMux structure.

```go
// see: https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=2439;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1

type ServeMux struct {
    mu       sync.RWMutex
    tree     routingNode
    index    routingIndex
    patterns []*pattern  // TODO(jba): remove if possible
    mux121   serveMux121 // used only when GODEBUG=httpmuxgo121=1
}
```

A tree structure (tree) is generated where paths become nodes, which seems to be a common data structure in HTTP routers.

The index and patterns are used for conflict detection of patterns.

## Routing Matching Process
- [ServeHTTP](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=3132;drc=960fa9bf66139e535d89934f56ae20a0e679e203)
- [findHandler](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1;l=2516?q=ServeHTTP&ss=go%2Fgo)
- [matchOrRedirect](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1;l=2578?q=ServeHTTP&ss=go%2Fgo)
- [match](https://cs.opensource.google/go/go/+/master:src/net/http/routing_tree.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1;l=115)
- [matchMethodAndPath](https://cs.opensource.google/go/go/+/master:src/net/http/routing_tree.go;l=130;drc=960fa9bf66139e535d89934f56ae20a0e679e203)
- [matchPath](https://cs.opensource.google/go/go/+/master:src/net/http/routing_tree.go;l=152;drc=960fa9bf66139e535d89934f56ae20a0e679e203)

If you are not used to reading, it might be good to read the code of the HTTP server as a prerequisite.

cf. [Reading the Code of Golang's HTTP Server](https://bmf-tech.com/posts/Golang%E3%81%AEHTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%83%AA%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0)

# Comparison of ServeMux and Third-Party Routers
I conducted a comparison between ServeMux and third-party routers using my custom benchmarker.

I utilized a previously implemented custom benchmarker. For details, please refer to the following links.

- [bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark)
- [Implemented a Benchmarking Tool to Compare Go's HTTP Routers](https://bmf-tech.com/posts/Go%e3%81%aeHTTP%20Router%e3%82%92%e6%af%94%e8%bc%83%e3%81%99%e3%82%8b%e3%83%99%e3%83%b3%e3%83%81%e3%83%9e%e3%83%bc%e3%82%ab%e3%83%bc%e3%82%92%e5%ae%9f%e8%a3%85%e3%81%97%e3%81%9f)

The benchmark results are published at [<Public>go-router-benchmark](https://docs.google.com/spreadsheets/d/1DrDNGJXfquw_PED3-eMqWqh7qbCTVBWZtqaKPXtngxg/edit#gid=1913830192).

In the above benchmarker, only the matching of path patterns is measured, and the registration of path patterns is not included in the measurement.

## Benchmark Results
From the benchmark results, while there are some differences in performance compared to Echo, Gin, and httprouter, which are considered to have good performance, overall, it seems to have above-average performance.

What stood out was the performance degradation observed as the number of path parameters increased.

It seems that higher-level HTTP routers have measures in place to mitigate this degradation.

Since cases that use many path parameters are not very common, I think this point does not need to be overly concerned.

## Comparison of goblin and ServeMux
In static routing test cases, goblin performs better, but for dynamic routing, ServeMux excels.

It appears that goblin does a better job of mitigating performance degradation as the number of path parameters increases.

As mentioned repeatedly, cases that use many path parameters are not very common, so I don't think the performance difference is significant in terms of practical usability.

# Performance of ServeMux
The implementers of ServeMux seem to consider the performance of HTTP routers as follows:

> Implementation is out of scope for this discussion/proposal. I think we'd be happy to have a more complex implementation if it could be demonstrated that the current one actually affects latency or CPU usage. For typical servers, that usually access some storage backend over the network, I'd guess the matching time is negligible. Happy to be proven wrong.

cf. Quote from https://github.com/golang/go/discussions/60227#discussioncomment-5932822

Other related comments are also noted.

cf. https://github.com/golang/go/issues/61410#issuecomment-1867191476
cf. https://github.com/golang/go/issues/61410#issuecomment-1867485864
cf. https://github.com/golang/go/issues/61410#issuecomment-1868615273

Many third-party HTTP routers that emphasize performance adopt complex tree structure algorithms (e.g., memory-efficient Radix Tree).

In the implementation of ServeMux, unless an extremely poor data structure is adopted, it is said that there will not be a significant impact on latency or CPU usage, which is the philosophy behind not adopting complex data structures or algorithms.

While not a counterargument, [gorilla/mux](https://github.com/gorilla/mux) is relatively underwhelming in benchmark results among popular (high-star) third-party HTTP routers, yet it is widely used.

There are also comments that mention performance regarding not just matching path patterns but also registering path patterns.

> Registration time is potentially more of an issue. With the precedence rules described here, checking a new pattern for conflicts seems to require looking at all existing patterns in the worst case. (Algorithm lovers, you are hereby nerd-sniped.) That means registering n patterns takes O(n2) time in the worst case. With the naive algorithm that loops through all existing patterns, that "worst case" is in fact every (successful) case: if there are no conflicts it will check every pattern against every other, for n(n-1)/2 checks. To see if this matters in practice, I collected all the methods from 260 Google Cloud APIs described by discovery docs, resulting in about 5000 patterns. In reality, no one server would serve all these patterns—more likely there are 260 separate servers—so I think this is a reasonable worst-case scenario. (Please correct me if I'm wrong.) Using naive conflict checking, it took about a second to register all the patterns—not too shabby for server startup, but not ideal. I then implemented a simple indexing scheme to weed out patterns that could not conflict, which reduced the time 20-fold, to 50 milliseconds. There are still sets of patterns that would trigger quadratic behavior, but I don't believe they would arise naturally; they would have to be carefully (maliciously?) constructed. And if you are being malicious, you are probably only hurting yourself: one writes patterns for one's own server, not the servers of others. If we do encounter real performance issues, we can index more aggressively.

cf. Quote from https://github.com/golang/go/discussions/60227#discussioncomment-6204048

I share a similar view on the performance of HTTP routers, so I can agree with the perspective on ServeMux's performance.

[goblin](https://github.com/bmf-san/goblin) adopts a data structure based on a Trie Tree.

The reason for not adopting a Radix Tree, which is more memory-efficient than a Trie Tree, is not only because it seems complex and difficult to understand and maintain, but also because I questioned whether the performance benefits would be significant enough to justify adopting a complex data structure.

Given Go's language philosophy, which seems to pursue simplicity, I think the trend will continue to optimize the current simplicity of ServeMux rather than adopting complex data structures or algorithms. (Probably.)

From the benchmark results, it seems there are tuning points, so I believe the scores will improve further in the future.

At that time, I would like to hold the second Ultimate HTTP Router Battle.

cf. [Ultimate HTTP Router Battle](https://speakerdeck.com/bmf_san/tian-xia-httprouterwu-dou-hui)

# Lessons Learned from Comparing ServeMux and Custom HTTP Router
While observing the implementation of ServeMux in Go 1.22, I felt that the routing algorithm is simple yet high-performing.

As long as you don't make a significant mistake in your choices, a certain level of performance is guaranteed. (This is the perspective of an algorithm novice.)

I think the neatness of the test cases in [routing_tree_test.go](https://cs.opensource.google/go/go/+/refs/tags/go1.22.2:src/net/http/routing_tree_test.go) and [pattern_test.go](https://cs.opensource.google/go/go/+/refs/tags/go1.22.2:src/net/http/pattern_test.go) is due to the simplicity of the data structure.

[goblin](https://github.com/bmf-san/goblin) is quite the opposite and is a point of concern for me.

While the Trie Tree adopted by [goblin](https://github.com/bmf-san/goblin) is simple, the way it is utilized is not simple, so there may be room for improvement.

While observing discussions and proposals, I felt again that the perspective of how much performance is expected from the choice of data structures and algorithms is important.

The more you pursue performance, the more you deviate from simplicity, so balance is likely key. (In short, balance is important.)

# My Views on Selecting HTTP Routers in Go
When developing a new application, I think the basic approach will be to first consider ServeMux and then consider third-party options if there are deficiencies.

On the other hand, regarding whether to consider migrating from the HTTP router currently used in an existing application to ServeMux, I think there are the following perspectives:

- Do you want to reduce dependency on third parties and lean towards the standard library as much as possible?
  - If you were reluctantly using a third-party router because wildcards were not available, you might want to actively consider switching.
- How compatible is the HTTP router you are using with net/http?
  - If you are using one that provides its own handler definitions and request parameter retrieval methods, migration may be cumbersome.
- Do you need features or performance that ServeMux does not have?
  - If you need middleware, regex-based routing, grouping, etc., it may make sense to continue using third-party options.
- Is there unique logic for routing priority?
  - If it can be guaranteed through testing, it may not be a problem, but it could be one of the barriers to migration.

I created a simple flowchart to decide whether to use ServeMux or a third-party router.

![Flowchart](https://github.com/bmf-san/bmf-san/assets/13291041/4bc81581-cdab-4fde-bb87-69f73511732f)

# Postscript
I presented at [Go Conference 2024](https://gocon.jp/2024/).

[speadkerdeck.com - From Custom HTTP Router to New ServeMux](https://speakerdeck.com/bmf_san/zi-zuo-httprutakaraxin-siiservemuxhe)

# References
- [bmf-tech.com/posts/tags/router](https://bmf-tech.com/posts/tags/router)
  - All past articles related to routers
- [net/http: add methods and path variables to ServeMux patterns](https://github.com/golang/go/discussions/60227)
- [net/http: enhanced ServeMux routing](https://github.com/golang/go/issues/61410)
- [Go 1.22 Release Notes - Enhanced routing patterns](https://tip.golang.org/doc/go1.22#enhanced_routing_patterns)
- [go.dev - Go 1.22 is released!](https://go.dev/blog/go1.22)
- [go.dev - Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements)
- [zenn.dev - New Router in Go 1.22](https://zenn.dev/catatsuy/scraps/37e3b52bca7d13)
- [zenn.dev - Routing Library michi Designed for Enhanced ServeMux in Go 1.22](https://zenn.dev/sonatard/articles/831b761a27b230)
- [future-architect.github.io - I presented on "ServeMux Conflict Detection and Performance" at the Go 1.22 Release Party](https://future-architect.github.io/articles/20240408b/)
- [rhumie.github.io - ServeMux Conflict Detection and Performance](https://rhumie.github.io/go122party/1)
- [eli.thegreenplace.net - Better HTTP server routing in Go 1.22](https://eli.thegreenplace.net/2023/better-http-server-routing-in-go-122/)
- [shijuvar.medium.com - Building REST APIs With Go 1.22 http.ServeMux](https://shijuvar.medium.com/building-rest-apis-with-go-1-22-http-servemux-2115f242f02b)
- [www.calhoun.io - Go's 1.22+ ServeMux vs Chi Router](https://www.calhoun.io/go-servemux-vs-chi/)
- [www.alexedwards.net - Which Go router should I use? (with flowchart)](https://www.alexedwards.net/blog/which-go-router-should-i-use)
- [www.youtube.com - Why The Golang 1.22 HTTP Router Is Not Great](https://www.youtube.com/watch?v=agX6Ba2ODlw)
- [www.reddit.com - The proposal to enhance Go's HTTP router](https://www.reddit.com/r/golang/comments/15dvauk/the_proposal_to_enhance_gos_http_router/)