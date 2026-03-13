---
title: Created a URL Router called goblin in Golang
description: 'Implement a high-performance URL router in Go using trie trees with path parameters and regex pattern matching.'
slug: goblin-url-router-in-golang
date: 2020-01-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - URL Routing
  - Router
translation_key: goblin-url-router-in-golang
---

# Overview
I created a URL router in Golang, so I will note down the process of implementation.

# Preparation
Here are the preparations I made when implementing the URL router.

## Data Structures and Algorithms
I considered the logic of how to match URLs.

Many libraries often use tree structures as data structures, so I thought about what type of tree structure to adopt.

Among trees specialized for string searching, it seemed that a radix tree would be the best choice in terms of time and memory complexity, so I initially tried to adopt that, but I gave up due to its complexity.

I decided to use a trie tree for something a bit more familiar and simpler.

## Reading `net/http` Code
To implement it as an extension of the multiplexer in `net/http`, it is necessary to have a certain understanding of its internal mechanisms.

Refer to [Reading Golang's HTTP Server Code](https://bmf-tech.com/posts/Golang%E3%81%AEHTTP%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E3%82%B3%E3%83%BC%E3%83%89%E3%83%AA%E3%83%BC%E3%83%87%E3%82%A3%E3%83%B3%E3%82%B0).

## Reading Various Router Implementations
Reference > See the repository.

## Others
An article summarizing past URL routing.

[URL Routing Articles](https://bmf-tech.com/posts/tags/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0)

# Implementation

Refer to [github.com - goblin](https://github.com/bmf-san/goblin).

The basic idea is to transform the trie tree into a more usable form, but I struggled several times with handling path parameters. The handling of regular expressions was not too troublesome, as it could be managed by simply preparing a DSL, so the sense of handling the DSL was put to the test.

During the implementation process, I wrote tests in parallel and repeated step-by-step debugging, which helped me catch how the data structure was constantly changing, so I felt my mental debugging skills were improving.

Since this involved logic I don't usually write, it surely became good training for coding.

The future challenges are as per the issues raised in the repository, and I plan to refine them when I have some free time.

# References
## Repository
Repositories that I referred to for design and implementation.

- [github - importcjj/trie-go](https://github.com/importcjj/trie-go)
- [github - julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [github - gorilla/mux](https://github.com/gorilla/mux)
- [github - xhallix/go-router](https://github.com/xhallix/go-router/tree/master)
- [github - gowww/router](https://github.com/gowww/router)
- [github - go-chi/chi](https://github.com/go-chi/chi)
- [github - go-ozzo/ozzo-routing](https://github.com/go-ozzo/ozzo-routing)

## Others
Articles referenced during implementation.

- [How to not use an http-router in go](https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html)
- [Understanding the Atmosphere of Go's HTTP Server](https://himetani.cafe/posts/go-http-server-overview/)
- [HTTP Server and context.Context](https://medium.com/@agatan/http%E3%82%B5%E3%83%BC%E3%83%90%E3%81%A8context-context-7211433d11e6)
- [Testing with httptest in golang](https://www.write-ahead-log.net/entry/2016/04/18/002638)
- [Created a Fast URL Router in PHP](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [Released the Fastest URL Router](https://kuune.org/text/2014/06/12/denco/)