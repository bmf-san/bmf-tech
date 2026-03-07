---
title: Created a Custom URL Router in Golang Called Goblin
slug: goblin-url-router-in-golang
date: 2020-01-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - URL Routing
  - Router
description: Documenting the process of implementing a custom URL router in Golang.
translation_key: goblin-url-router-in-golang
---

# Overview
I created a custom URL router in Golang and documented the process of its implementation.

# Preparation
Here’s a summary of the preparations I made before implementing the URL router.

## Data Structures and Algorithms
I considered the logic for how URLs should be matched.

Many libraries use tree structures as their data structure, so I explored which type of tree structure to adopt. Among trees optimized for string searches, radix trees seemed to offer the best time and memory efficiency. However, I found the implementation too challenging and decided to use a simpler and more familiar structure: the trie.

## Reading the `net/http` Code
To implement the router as an extension of the multiplexer in `net/http`, I needed to understand its internal workings to some extent.

Refer to [Code Reading of Golang's HTTP Server](https://bmf-tech.com/posts/Golang%E3%81%AEHTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%83%AA%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0).

## Reading Implementations of Various Routers
I reviewed various router implementations for reference. See the repository links below.

## Others
Here’s an article I wrote previously summarizing URL routing.

[bmf-tech.com/posts/tags/URLルーティング](https://bmf-tech.com/posts/tags/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0)

# Implementation
Refer to [github.com - goblin](https://github.com/bmf-san/goblin).

The basic idea was to adapt the trie structure into a more user-friendly form. However, I struggled several times with handling path parameters. Supporting regular expressions wasn’t too difficult; it was just a matter of preparing a DSL, but handling the DSL required some finesse.

During the implementation process, I wrote tests in parallel and repeatedly debugged step-by-step. By constantly tracking how the data structure changed, I felt my debugging skills improved over time.

Since I don’t usually write this kind of logic, it was undoubtedly a good coding exercise.

Future tasks are listed as issues in the repository, and I plan to refine the implementation whenever I have free time.

# References
## Repositories
Repositories that served as references for design and implementation:

- [github - importcjj/trie-go](https://github.com/importcjj/trie-go)
- [github - julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [github - gorilla/mux](https://github.com/gorilla/mux)
- [github - xhallix/go-router](https://github.com/xhallix/go-router/tree/master)
- [github - gowww/router](https://github.com/gowww/router)
- [github - go-chi/chi](https://github.com/go-chi/chi)
- [github - go-ozzo/ozzo-routing](https://github.com/go-ozzo/ozzo-routing)

## Articles
Articles I referred to during implementation:

- [How to not use an http-router in go](https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html)
- [Understanding the Atmosphere of Go's HTTP Server](https://himetani.cafe/posts/go-http-server-overview/)
- [HTTP Server and context.Context](https://medium.com/@agatan/http%E3%82%B5%E3%83%BC%E3%83%90%E3%81%A8context-context-7211433d11e6)
- [Testing with httptest in Golang](https://www.write-ahead-log.net/entry/2016/04/18/002638)
- [Creating a High-Performance URL Router in PHP](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [Released the World's Fastest URL Router](https://kuune.org/text/2014/06/12/denco/)