---
title: From Custom HTTP Router to the New ServeMux
slug: custom-http-router-to-new-servemux
date: 2024-04-27T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Router
description: Exploring the new ServeMux features introduced in Go 1.22 and comparing it with custom and third-party HTTP routers.
translation_key: custom-http-router-to-new-servemux
---

# Overview
Previously, I used my custom HTTP router called [goblin](https://github.com/bmf-san/goblin) in my applications. However, since the enhanced ServeMux features were introduced in Go 1.22, I’ve switched to using ServeMux.

In this article, I’ll summarize the new features and performance improvements of ServeMux in Go 1.22 and discuss considerations for selecting HTTP routers in Go moving forward.

# New ServeMux Features in Go 1.22
When Go 1.22rc was released, I investigated the new ServeMux features. Let’s delve deeper into these enhancements.

cf. [Changes to ServeMux in Go 1.22rc](https://bmf-tech.com/posts/Go1.22rc%e3%81%a7%e5%a4%89%e6%9b%b4%e3%81%95%e3%82%8c%e3%82%8bServeMux%e3%81%ae%e4%bb%95%e6%a7%98)

Based on the following references, I’ve organized the new ServeMux features:

- Release Notes
  - [Go 1.22 Release Notes - Enhanced routing patterns](https://tip.golang.org/doc/go1.22#enhanced_routing_patterns)
    - ServeMux patterns now support HTTP methods and wildcards (dynamic path parameters, e.g., `/items/{id}`).
- pkg.go.dev
  - [pkg.go.dev - net/http](https://pkg.go.dev/net/http)
- go.dev
  - [go.dev - Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements)
- Discussions
  - [net/http: add methods and path variables to ServeMux patterns](https://github.com/golang/go/discussions/60227)
- Proposal
  - [net/http: enhanced ServeMux routing](https://github.com/golang/go/issues/61410)

## Defining Routes with HTTP Methods
By specifying paths that include HTTP methods, you can now define routes based on HTTP methods. Using ServeMux eliminates the need to write conditional logic for HTTP methods within handlers.

```go
http.HandleFunc("GET /items", handleItems)
```

Although HTTP method constants (e.g., `http.MethodGet`) are available, the current approach likely preserves backward compatibility by not altering existing method signatures. Alternatively, it might simply align with the HTTP request format.

## Defining Routes with Wildcards
By specifying paths with wildcards (`{pathVal}`), you can define routes that use wildcards.

```go
// Matches GET /items/1 or /items/foo
http.HandleFunc("GET /items/{id}", handleItems)
```

Path patterns can be specified in the following format. Even before Go 1.22, hostnames could also be specified (something I learned recently).

```
[METHOD ][HOST]/[PATH]
```

Values matched by wildcards can be retrieved using the `http.Request` [PathValue method](https://pkg.go.dev/net/http#Request.PathValue).

```go
func handleItems(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    // `id` contains the value matched by the wildcard
}
```

Additionally, multi-wildcard routing definitions are possible.

```go
// Matches GET /items/1, /items/1/2, /items/1/2/3, etc.
http.HandleFunc("GET /items/{id...}", handleItems)
```

To match paths exactly, use `{$}`. This is useful for defining routes like `/`.

```go
http.HandleFunc("GET /{$}", handleIndex)
```

As an aside, some third-party routers use `*` as a wildcard in patterns. Personally, I prefer calling it a path parameter to avoid confusion (though it’s not a big deal).

# Considerations with the New ServeMux
## Defining Patterns with HTTP Methods
In many third-party HTTP routers, HTTP methods are often defined as methods.

```go
// ServeMux
http.HandleFunc("GET /items/{id}", handleItems)

// Example from third-party router
mux.Get("/items", handleItems)
```

Since ServeMux includes HTTP methods within patterns, minor issues like typos might be harder to detect. While linters could help, creating a static analysis tool might be an interesting project. I suspect this will be addressed eventually.

## Precedence Rules
When defining routes with wildcards, precedence should be considered.

cf.
- [pkg.go.dev - Precedence](https://pkg.go.dev/net/http#hdr-Precedence)
- [go.dev - Routing Enhancements for Go 1.22 precedence](https://go.dev/blog/routing-enhancements#precedence)

ServeMux allows defining routing patterns like the following:

```
// If both match, the first is prioritized
/items/new
/items/{id}

// If both match, the latter is prioritized
/items/{id...}
/items/{id}/category/{name}
```

In such cases, you need to be mindful of which pattern matches the intended request.

Some HTTP routers disallow overlapping patterns, while others permit them. ServeMux falls into the latter category.

On the other hand, ServeMux detects conflicts in cases like the following and triggers a panic:

```
// Cases where both or only one matches
/items/{id}
/{category}/items

// Cases where both match
/items/{id}
/items/{name}
```

Conflicts are likely detected early (e.g., during testing or server startup), making them less troublesome than overlaps.

Incidentally, my custom router [goblin](https://github.com/bmf-san/goblin) prioritizes the first registered pattern (not well-designed, so it’s quite messy).

For more details on ServeMux conflict detection, refer to the following article:

cf. [rhumie.github.io - ServeMux Conflict Detection and Performance](https://rhumie.github.io/go122party/1)

Even third-party HTTP routers consider conflict detection. For example, [httprouter](https://github.com/julienschmidt/httprouter) ensures that a pattern either matches exactly one route or none at all.

> Only explicit matches: With other routers, like http.ServeMux, a requested URL path could match multiple patterns. Therefore they have some awkward pattern priority rules, like longest match or first registered, first matched. By design of this router, a request can only match exactly one or no route. As a result, there are also no unintended matches, which makes it great for SEO and improves the user experience.

cf. https://github.com/julienschmidt/httprouter?tab=readme-ov-file#features

The priority rules for HTTP router pattern matching significantly impact the quality of the router, so well-designed rules provide reassurance.

## Backward Compatibility
There are cases where backward compatibility between Go 1.22 and Go 1.21 is not maintained.

cf. [pkg.go.dev - Compatibility](https://pkg.go.dev/net/http#hdr-Compatibility)

In such cases, setting the `GODEBUG` environment variable to `httpmuxgo121=1` restores Go 1.21 behavior.

# Internal Implementation Code Reading
## Routing Pattern Registration
- [register](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=2735;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1)
- [parsePattern](https://cs.opensource.google/go/go/+/master:src/net/http/pattern.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;l=84)

Below is the definition of the ServeMux struct:

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

Paths are structured as nodes in a tree, which is a common data structure for HTTP routers.

The `index` and `patterns` fields are used for conflict detection.

## Routing Matching Process
- [ServeHTTP](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;l=3132;drc=960fa9bf66139e535d89934f56ae20a0e679e203)
- [findHandler](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1;l=2516?q=ServeHTTP&ss=go%2Fgo)
- [matchOrRedirect](https://cs.opensource.google/go/go/+/master:src/net/http/server.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1;l=2578?q=ServeHTTP&ss=go%2Fgo)
- [match](https://cs.opensource.google/go/go/+/master:src/net/http/routing_tree.go;drc=960fa9bf66139e535d89934f56ae20a0e679e203;bpv=1;bpt=1;l=115)
- [matchMethodAndPath](https://cs.opensource.google/go/go/+/master:src/net/http/routing_tree.go;l=130;drc=960fa9bf66139e535d89934f56ae20a0e679e203)
- [matchPath](https://cs.opensource.google/go/go/+/master:src/net/http/routing_tree.go;l=152;drc=960fa9bf66139e535d89934f56ae20a0e679e203)

If you’re unfamiliar with reading code, it might help to start with HTTP server code.

cf. [Code Reading for Golang HTTP Server](https://bmf-tech.com/posts/Golang%E3%81%AEHTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%83%AA%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0)

# Comparing ServeMux with Third-Party Routers
I compared ServeMux with third-party routers using a custom benchmarking tool.

The benchmarking tool was implemented previously. For details, refer to the following links:

- [bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark)
- [Implemented a Benchmarking Tool to Compare Go HTTP Routers](https://bmf-tech.com/posts/Go%e3%81%aeHTTP%20Router%e3%82%92%e6%af%94%e8%bc%83%e3%81%99%e3%82%8b%e3%83%99%e3%83%b3%e3%83%81%e3%83%9e%e3%83%bc%e3%82%ab%e3%83%bc%e3%82%92%e5%ae%9f%e8%a3%85%e3%81%97%e3%81%9f)

Benchmark results are publicly available at [<Public>go-router-benchmark](https://docs.google.com/spreadsheets/d/1DrDNGJXfquw_PED3-eMqWqh7qbCTVBWZtqaKPXtngxg/edit#gid=1913830192).

The benchmark only measures path pattern matching, excluding pattern registration.

## Benchmark Results
Compared to high-performance routers like Echo, gin, and httprouter, ServeMux shows some differences but overall performs above average.

A notable observation was performance degradation as the number of path parameters increased.

Top-tier HTTP routers seem to employ strategies to mitigate this degradation.

Since cases involving numerous path parameters are rare, this issue is not particularly concerning.

## Comparing goblin and ServeMux
In static routing test cases, goblin outperforms ServeMux, but ServeMux excels in dynamic routing.

While goblin mitigates performance degradation caused by increasing path parameters, the practical impact is minimal due to the rarity of such cases.

# ServeMux Performance Considerations
The ServeMux implementation team seems to view HTTP router performance as follows:

> Implementation is out of scope for this discussion/proposal. I think we'd be happy to have a more complex implementation if it could be demonstrated that the current one actually affects latency or CPU usage. For typical servers, that usually access some storage backend over the network, I'd guess the matching time is negligible. Happy to be proven wrong.

cf. https://github.com/golang/go/discussions/60227#discussioncomment-5932822

Additional related comments:

cf. https://github.com/golang/go/issues/61410#issuecomment-1867191476
cf. https://github.com/golang/go/issues/61410#issuecomment-1867485864
cf. https://github.com/golang/go/issues/61410#issuecomment-1868615273

Many third-party HTTP routers prioritize performance by adopting complex tree algorithms (e.g., memory-efficient Radix Tree).

ServeMux’s implementation philosophy suggests that unless the data structure is extremely inefficient, latency and CPU usage are unlikely to be significantly impacted. Thus, it avoids adopting overly complex structures or algorithms.

While not a counterexample, [gorilla/mux](https://github.com/gorilla/mux) is a popular third-party HTTP router with relatively lower benchmark results but is widely used.

There’s also commentary on the performance of pattern registration:

> Registration time is potentially more of an issue. With the precedence rules described here, checking a new pattern for conflicts seems to require looking at all existing patterns in the worst case. (Algorithm lovers, you are hereby nerd-sniped.) That means registering n patterns takes O(n2) time in the worst case. With the naive algorithm that loops through all existing patterns, that "worst case" is in fact every (successful) case: if there are no conflicts it will check every pattern against every other, for n(n-1)/2 checks. To see if this matters in practice, I collected all the methods from 260 Google Cloud APIs described by discovery docs, resulting in about 5000 patterns. In reality, no one server would serve all these patterns—more likely there are 260 separate servers—so I think this is a reasonable worst-case scenario. (Please correct me if I'm wrong.) Using naive conflict checking, it took about a second to register all the patterns—not too shabby for server startup, but not ideal. I then implemented a simple indexing scheme to weed out patterns that could not conflict, which reduced the time 20-fold, to 50 milliseconds. There are still sets of patterns that would trigger quadratic behavior, but I don't believe they would arise naturally; they would have to be carefully (maliciously?) constructed. And if you are being malicious, you are probably only hurting yourself: one writes patterns for one's own server, not the servers of others. If we do encounter real performance issues, we can index more aggressively.

cf. https://github.com/golang/go/discussions/60227#discussioncomment-6204048

I share similar views on HTTP router performance, so I agree with ServeMux’s perspective on performance.

[goblin](https://github.com/bmf-san/goblin) uses a Trie Tree-based data structure.

I didn’t adopt Radix Tree, which is more memory-efficient, partly because it seemed complex and challenging to maintain, but also because I questioned whether the performance benefits justified the complexity.

Given Go’s philosophy of simplicity, ServeMux is likely to continue optimizing its current simplicity rather than adopting complex structures or algorithms.

Looking at benchmark results, there seems to be room for tuning, so I expect ServeMux to improve its scores further.

When that happens, I’d like to host the second "Tenkaichi HTTPRouter Tournament."

cf. [Tenkaichi HTTPRouter Tournament](https://speakerdeck.com/bmf_san/tian-xia-httprouterwu-dou-hui)

# Lessons Learned from Comparing ServeMux and Custom HTTP Routers
Examining the ServeMux implementation in Go 1.22, I noticed its routing algorithm is simple yet high-performing.

It seems that as long as the choice isn’t drastically wrong, performance is reasonably guaranteed (a novice’s perspective).

The test cases in [routing_tree_test.go](https://cs.opensource.google/go/go/+/refs/tags/go1.22.2:src/net/http/routing_tree_test.go) and [pattern_test.go](https://cs.opensource.google/go/go/+/refs/tags/go1.22.2:src/net/http/pattern_test.go) are straightforward, likely due to the simplicity of the data structure.

[goblin](https://github.com/bmf-san/goblin), on the other hand, is quite messy, which I’m concerned about.

While Trie Tree itself is simple, its usage in goblin is not, leaving room for improvement.

Reading discussions and proposals, I realized the importance of considering expected performance when choosing data structures and algorithms.

Pursuing performance often sacrifices simplicity, so balance is crucial (I’m a "balance guy").

# My Thoughts on Selecting Go HTTP Routers
For new applications, I believe the default approach should be to consider ServeMux first and then evaluate third-party options if ServeMux falls short.

For existing applications, the decision to migrate from a third-party router to ServeMux depends on the following factors:

- Do you want to reduce third-party dependencies and rely more on standard libraries?
  - If you’ve been using a third-party router reluctantly due to the lack of wildcard support, you might actively consider switching.
- How compatible is your current HTTP router with `net/http`?
  - If you’re using a router with custom handler definitions or request parameter retrieval methods, migration might be more challenging.
- Do you need features or performance that ServeMux lacks?
  - If you require middleware, regex-based routing, or grouping, continuing with third-party routers might be more practical.
- Does your routing logic rely on custom precedence rules?
  - While tests can ensure correctness, this could be a migration hurdle.

I created a simple flowchart to decide between ServeMux and third-party routers:

![Flowchart](https://github.com/bmf-san/bmf-san/assets/13291041/4bc81581-cdab-4fde-bb87-69f73511732f)

# Update
I presented at [Go Conference 2024](https://gocon.jp/2024/).

[speakerdeck.com - From Custom HTTP Router to the New ServeMux](https://speakerdeck.com/bmf_san/zi-zuo-httprutakaraxin-siiservemuxhe)

# References
- [bmf-tech.com/posts/tags/router](https://bmf-tech.com/posts/tags/router)
  - All past articles related to routers
- [net/http: add methods and path variables to ServeMux patterns](https://github.com/golang/go/discussions/60227)
- [net/http: enhanced ServeMux routing](https://github.com/golang/go/issues/61410)
- [Go 1.22 Release Notes - Enhanced routing patterns](https://tip.golang.org/doc/go1.22#enhanced_routing_patterns)
- [go.dev - Go 1.22 is released!](https://go.dev/blog/go1.22)
- [go.dev - Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements)
- [zenn.dev - About Go 1.22’s new router](https://zenn.dev/catatsuy/scraps/37e3b52bca7d13)
- [zenn.dev - Routing library michi designed for Go 1.22’s Enhanced ServeMux](https://zenn.dev/sonatard/articles/831b761a27b230)
- [future-architect.github.io - Presented "ServeMux Conflict Detection and Performance" at Go1.22 Release Party](https://future-architect.github.io/articles/20240408b/)
- [rhumie.github.io - ServeMux Conflict Detection and Performance](https://rhumie.github.io/go122party/1)
- [eli.thegreenplace.net - Better HTTP server routing in Go 1.22](https://eli.thegreenplace.net/2023/better-http-server-routing-in-go-122/)
- [shijuvar.medium.com - Building REST APIs With Go 1.22 http.ServeMux](https://shijuvar.medium.com/building-rest-apis-with-go-1-22-http-servemux-2115f242f02b)
- [www.calhoun.io - Go’s 1.22+ ServeMux vs Chi Router](https://www.calhoun.io/go-servemux-vs-chi/)
- [www.alexedwards.net - Which Go router should I use? (with flowchart)](https://www.alexedwards.net/blog/which-go-router-should-i-use)
- [www.youtube.com - Why The Golang 1.22 HTTP Router Is Not Great](https://www.youtube.com/watch?v=agX6Ba2ODlw)
- [www.reddit.com - The proposal to enhance Go’s HTTP router](https://www.reddit.com/r/golang/comments/15dvauk/the_proposal_to_enhance_gos_http_router/)