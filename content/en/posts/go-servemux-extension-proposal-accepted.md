---
title: Proposal for Extending ServeMux in Go Accepted
slug: go-servemux-extension-proposal-accepted
date: 2023-08-31T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Router
  - HTTP
description: Discussing the acceptance of a proposal to enhance ServeMux routing in Go.
translation_key: go-servemux-extension-proposal-accepted
---

I’ve been keeping an eye on [net/http: enhanced ServeMux routing #61410](https://github.com/golang/go/issues/61410), and it has now been accepted. Here’s a reflection on it.

While the final specifications are still unclear, it seems likely that dynamic routing functionality (e.g., path parameters like `/foo/:id`) will be added to ServeMux, which currently only supports static routing (e.g., fixed routes like `/foo/bar`).

The reason I’m interested in this proposal is that I’ve been personally developing a routing library.

cf. [bmf-san/goblin](https://github.com/bmf-san/goblin)
cf. [Related article on bmf-tech.com](https://bmf-tech.com/posts/HTTP%20Router%e3%81%ae%e8%87%aa%e4%bd%9c%e3%81%a7%e5%8f%82%e8%80%83%e3%81%ab%e3%81%97%e3%81%9f%e8%b3%87%e6%96%99#4-%E5%9F%B7%E7%AD%86%E3%81%97%E3%81%9F%E8%A8%98%E4%BA%8B)

Goblin, like other popular routing libraries, extends ServeMux’s functionality and supports both static and dynamic routing. For dynamic routing, path parameters are supported with regular expressions.

The internal data structure of Goblin is based on a Trie tree, and I’ve optimized it as much as I can. (For better performance, the data structure would need a fundamental overhaul, though...)

I compared the performance of Goblin with the reference implementation mentioned in the proposal (though the actual implementation upon release might differ). The results were roughly comparable.

cf. [Add jba/muxpatterns #23](https://github.com/bmf-san/go-router-benchmark/pull/23)

If ServeMux’s extended functionality is released, Goblin might lose its purpose. After all, I’m probably the only one using it, and it’s just for my personal applications.

I’ve been maintaining Goblin with the hope that it might gain some recognition someday, but this proposal might mark a turning point. 😅

Looking ahead, whether to adopt third-party routing solutions will likely depend on how the standard ServeMux evolves as a major option.

Criteria for adopting third-party routing over ServeMux might include:

- **Better performance**
  - Once the new ServeMux is released, I plan to compare its performance using [go-router-benchmark](https://github.com/bmf-san/go-router-benchmark).
  - Based on the reference implementation, it seems likely to deliver decent performance.
  - While I don’t fully understand the reference implementation’s data structure, it appears to be based on a Radix Tree (a data structure commonly used in high-performance routers). If optimized, it could outperform Goblin and potentially rival popular libraries.

- **Features not available in the standard**
  - It’s unclear what features beyond path parameters will be added, but examples might include route grouping, middleware handling, static file serving, CORS, etc. (Though I doubt any of these will be added.)

These are the main points that come to mind.

Ease of switching from third-party routing to ServeMux is also a concern. Whether the library adheres to the interfaces defined in `net/http` (e.g., `Handler`) could be a key factor. If it adheres, there shouldn’t be major issues. However, switching from non-compliant libraries might be troublesome. (Most libraries are compliant, but some have custom implementations.)

According to comments on the proposal, concrete implementation discussions will begin soon, so I’ll continue to keep an eye on it.