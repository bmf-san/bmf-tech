---
title: Proposal to Extend ServeMux Functionality in Go Accepted
description: An in-depth look at Proposal to Extend ServeMux Functionality in Go Accepted, covering key concepts and practical insights.
slug: go-servemux-extension-proposal-accepted
date: 2023-08-31T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Router
  - HTTP
translation_key: go-servemux-extension-proposal-accepted
---

I have been watching the [net/http: enhanced ServeMux routing #61410](https://github.com/golang/go/issues/61410), and it has been accepted, so I want to write a little about it.

While I don't know what the final specification will look like, it seems that dynamic routing (like /foo/:id using path parameters) will at least be added to ServeMux, which currently only supports static routing (like /foo/bar).

The reason I am interested in this proposal is that I am personally developing a routing library.

cf. [bmf-san/goblin](https://github.com/bmf-san/goblin)
cf. [Related articles on bmf-tech.com](https://bmf-tech.com/posts/HTTP%20Router%e3%81%ae%e8%87%aa%e4%bd%9c%e3%81%a7%e5%8f%82%e8%80%83%e3%81%ab%e3%81%97%e3%81%9f%e8%b3%87%e6%96%99#4-%E5%9F%B7%E7%AD%86%E3%81%97%E3%81%9F%E8%A8%98%E4%BA%8B)

goblin is a library that extends the functionality of ServeMux, just like other well-known routing libraries, and it supports both static and dynamic routing. The dynamic routing path parameters also support regular expressions.

goblin's internal data structure is based on a Trie tree logic, and I have optimized it as much as possible. (If better performance is sought, it would require fundamentally changing the data structure...)

I did a performance comparison with the reference implementation mentioned in this proposal (though I think it may differ from the actual implementation that will be released, I did it for reference), and it felt pretty much the same...
cf. [Add jba/muxpatterns #23](https://github.com/bmf-san/go-router-benchmark/pull/23)

Once the extension of ServeMux is released, I think there will probably be no reason to use goblin anymore. After all, I am the only one using it, and I believe I am just using it for my personal applications...

I had hoped that goblin would be continuously maintained and one day see the light of day, but it seems that this proposal may mark the end of that hope.

Regarding whether to adopt third-party routing in the future, I think the standard ServeMux will become a significant option.

As criteria for whether to adopt third-party routing compared to the standard ServeMux:

- Better performance
  - Once it is actually released, I plan to compare it using [go-router-benchmark](https://github.com/bmf-san/go-router-benchmark).
  - If we consider the reference implementation, I feel it will deliver decent performance.
  - I don't fully understand the data structure of the reference implementation, but it seemed to be based on a Radix Tree (which is a data structure commonly used in high-performance routers). Since the reference implementation didn't seem optimized, if that gets optimized, goblin will likely be at a significant disadvantage, and it could very well stand shoulder to shoulder with well-known libraries.
- Need for features not in the standard
  - I don't know what other features will be added besides path parameters, but for example, routing grouping, something around middleware, static file delivery, CORS, etc. (I probably think none of these features will be added)

These are the points that come to mind.

I am concerned about the ease of switching from third-party libraries to ServeMux, but I think whether they comply with the interfaces defined in net/http (like Handler) could be a bottleneck, so if they comply, there shouldn't be a major issue. Switching from non-compliant libraries might be troublesome. (Most comply, but some have their own non-compliant implementations)

According to the comments on the proposal, it seems they will be considering specific implementations from now on, so I will continue to watch this closely.