---
title: 'Introduction to Custom URL Routing: Episode 2'
description: 'Research notes and a structured overview of Introduction to Custom URL Routing: Episode 2, summarizing key concepts and findings.'
slug: url-routing-introduction-episode-2
date: 2019-12-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - HTTP
  - URL Routing
  - Router
translation_key: url-routing-introduction-episode-2
---



# Overview
This article is a continuation of [Introduction to Custom URL Routing: Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E8%87%AA%E4%BD%9C%E5%85%A5%E9%96%80%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%89%EF%BC%91) and is the 15th day of the [Makuake Development Team Advent Calendar 2019](https://adventar.org/calendars/4716).

# Creating Custom URL Routing
Continuing from the previous article.

When creating a custom router, let's consider what kind of processing the router performs from the perspective of data structure.

![Screen Shot 2019-12-15 at 19 07 42](/assets/images/posts/url-routing-introduction-episode-2/70861219-30929d80-1f6e-11ea-8e86-114e8ba0942b.png)

I have illustrated an example of what input the router receives and what output it returns in the case of dynamic routing.

The router's role is to receive the URL path part as input, determine the data that matches the path, output it, and connect it to the next process.

How to perform path matching is the core part of the implementation.

The performance required of the router may vary depending on the context in which it is introduced (such as the scale of the application or the direction of the architecture), but here I would like to consider implementing matching processing using a tree structure instead of just regular expression-based matching.※

※In the past, I wrote a library called [bmf-san/bmf-react-router](https://github.com/bmf-san/bmf-react-router) for use in React, which was based on regular expressions. It was merely a component wrapping the matching process, thanks to the convenient library `path-to-regexp`...

Now, let's briefly explain tree structures.

![Screen Shot 2019-12-15 at 20 55 26](/assets/images/posts/url-routing-introduction-episode-2/70862253-3e9bea80-1f7d-11ea-9856-1da4906c316a.png)

A tree structure is a data structure that has a root, edges, nodes, and leaves (terminal nodes). There are various types of tree structures depending on the pattern of data storage and search methods. Since URL routing mainly deals with strings, we will adopt a trie tree, which is specialized for string search. For more on trie trees, refer to [bmf-tech.com - A Trie implementation in Golang](https://bmf-tech.com/posts/A%20Trie%20implementation%20in%20Golang).

We will customize the trie tree to store data for URL path matching as follows.

![Screen Shot 2019-12-15 at 21 39 38](/assets/images/posts/url-routing-introduction-episode-2/70862745-7148e180-1f83-11ea-85d3-2cd8fb4db0d3.png)

As an example, we define routing that only supports GET.

- `/foo/`
- `/baz/`
- `/foo/bar/`
- `/foo/:name[string]/`
- `/foo/bar/:id[int]`  

Nodes are grown directly under the root for each HTTP method, and paths are stored for each HTTP method. Nodes starting with `:` are nodes that store path parameters, and the DSL `[int]` is used to associate regular expressions with those parameters. The values that connect to the output (which I think is the router's responsibility to the point of function invocation) may or may not be held by each node. This depends on the predefined routing content.

If you can write code that reproduces the data structure defined above based on the routing definition, you can theoretically create your own router by writing code for the client that uses the router. Theoretically.※

※There is something being implemented, but it's still in progress... I'll try to complete it by the 20th day of [Qiita - Go6 Advent Calendar 2019](https://qiita.com/advent-calendar/2019/go6).

The latter part became a bit rushed and rough, but once you get the image, it's just a matter of writing code, so it's as good as done ().

This time, I explained a data structure like a custom extension of a trie tree, but if you are more conscious of memory efficiency, it might be better to adopt a Patricia tree structure. I saw it being adopted when I was exploring Golang implementations.

I have attempted to implement a Patricia tree before but gave up once, so I would like to try again someday...※

※I thought I would understand by reading the code of a Patricia tree, but unlike a simple trie tree, the implementation patterns are "everyone is different and everyone is good," so I tried various things to properly understand and implement the data structure of a Patricia tree, but it was too difficult...

# Thoughts
I think routers are a genre with many competitors as OSS libraries, but if you can create something that can stand shoulder to shoulder with even the later ones, the world you see will change, and it will be more fun, so I want to continue in this field as long as I don't get bored.

# Postscript
I implemented a URL router in Golang.

[github.com - bmf-san/goblin](https://github.com/bmf-san/goblin)
