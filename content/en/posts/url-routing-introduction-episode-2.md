---
title: Introduction to Custom URL Routing Episode 2
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
This article is a continuation of [Introduction to Custom URL Routing Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E8%87%AA%E4%BD%9C%E5%85%A5%E9%96%80%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%89%EF%BC%91), and it is the 15th day of the [Makuake Development Team Advent Calendar 2019](https://adventar.org/calendars/4716).

# Creating Custom URL Routing
Continuing from the last time.

When creating a custom router, let's consider what kind of processing the router performs from the perspective of data structures.

![Screen Shot 2019-12-15 at 19 07 42](https://user-images.githubusercontent.com/13291041/70861219-30929d80-1f6e-11ea-8e86-114e8ba0942b.png)

I illustrated an example of what kind of Input the router receives and what kind of Output it returns in the case of dynamic routing.

The router's role is to take the path part of the URL as Input, determine the data that matches the path, and Output it to connect to the next processing step.

How path matching is performed will be the core part of the implementation.

The performance required from the router may vary depending on the context of introducing the router (such as the scale of the application or the direction of the architecture), but here I would like to consider implementing matching processing using a tree structure rather than just regular expressions. ※

※ A long time ago, I wrote a library for React called [bmf-san/bmf-react-router](https://github.com/bmf-san/bmf-react-router) based on regular expressions. It was just a component that wrapped the matching process, benefiting from a convenient library called `path-to-regexp`...

Now, let me briefly explain the tree structure.

![Screen Shot 2019-12-15 at 20 55 26](https://user-images.githubusercontent.com/13291041/70862253-3e9bea80-1f7d-11ea-9856-1da4906c316a.png)

A tree structure is a data structure that has a root, edges, nodes, and leaves (the terminal nodes). There are various types of tree structures depending on the patterns of data storage and search methods. Since we want to mainly handle strings in URL routing, we will adopt a trie, a data structure specialized for string searching. For more on tries, please refer to [bmf-tech.com - A Trie implementation in Golang](https://bmf-tech.com/posts/A%20Trie%20implementation%20in%20Golang).

We will customize the trie to store data for URL path matching as follows:

![Screen Shot 2019-12-15 at 21 39 38](https://user-images.githubusercontent.com/13291041/70862745-7148e180-1f83-11ea-85d3-2cd8fb4db0d3.png)

As an example, we define routing that only responds to GET requests:

- `/foo/`
- `/baz/`
- `/foo/bar/`
- `/foo/:name[string]/`
- `/foo/bar/:id[int]`

Nodes are created directly under the root for each HTTP method, and paths are stored for each HTTP method. Nodes that start with `:` are for storing path parameters, and the DSL `[int]` is for associating regular expressions with those parameters. The values that lead to the Output (which are basically the responsibility of the router until the function call) may or may not be held by each node. This depends on the content of the routing defined in advance.

If you can write code that reproduces the data structure defined above based on the routing definition, then theoretically, you can create your own router by writing code for the clients that will use the router. Theoretically. ※

※ There is something implemented, but it's still in progress... I will do my best by the 20th day of [Qiita - Go6 Advent Calendar 2019](https://qiita.com/advent-calendar/2019/go6).

The latter half became a bit rushed and rough, but if you understand the concept, it's just a matter of writing code, so it's almost like you can create it on your own. ()

This time, I explained a data structure like my own extension of the trie, but if you want to be more conscious of memory efficiency, it seems good to adopt a Patricia tree. I have seen it used in various implementations while rummaging through Golang implementations.

I have attempted to implement a Patricia tree before, but I gave up once, so I would like to challenge it again someday... ※

※ I thought I could understand by reading the code of the Patricia tree, but unlike a simple trie, the implementation patterns are all different and good in their own way, so I thought I should properly understand the Patricia tree data structure and implement it, but it was too difficult...

# Thoughts
I think the router is a genre with many competitors as an OSS library, but if we can create something that can stand shoulder to shoulder with later entrants, I believe the world will change and become more enjoyable, so I want to continue engaging in this field as long as I don't get bored.

# Postscript
I implemented a URL router in Golang.

[github.com - bmf-san/goblin](https://github.com/bmf-san/goblin)