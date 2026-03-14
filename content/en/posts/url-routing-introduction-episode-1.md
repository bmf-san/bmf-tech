---
title: Introduction to Building Your Own URL Routing - Episode 1
description: "Learn URL routing fundamentals with tree data structures and trie algorithms for building custom application routers."
slug: url-routing-introduction-episode-1
date: 2019-12-14T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - HTTP
  - URL Routing
  - Router
translation_key: url-routing-introduction-episode-1
---



# Overview
This article is for the 14th day of the [Makuake Development Team Advent Calendar 2019](https://adventar.org/calendars/4716).

As a hobbyist beginner in building my own URL routing, I hope this article will be helpful for those who want to enter the world of custom URL routing.

*Note: The keyword "beginner" was a trend in the web community this year, wasn't it? I started building my own URL routing at the end of last year, so I consider myself a beginner. Whether such a community exists, I don't know, but the world is vast, so it probably does.*

# The Journey of a Beginner URL Routing Builder
Although it may be unbearable to look at, I will expose the process of trial and error.

- Articles
    - [Building URL Routing - Episode 1](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%891)
    - [Building URL Routing - Episode 2](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%892)
    - [Building URL Routing - Episode 3 (Final)](https://bmf-tech.com/posts/URL%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B%E3%80%80%E3%82%A8%E3%83%94%E3%82%BD%E3%83%BC%E3%83%893%EF%BC%88%E5%AE%8C%E7%B5%90%E7%B7%A8%EF%BC%89)

- Slides
    - [Building URL Routing Episode 1](https://speakerdeck.com/bmf_san/urlruteinguwotukuruepisodo1)
    - [Building URL Routing](https://speakerdeck.com/bmf_san/urlruteinguwotukuru)
    - [Building Custom URL Routing with PHP](https://speakerdeck.com/bmf_san/phpdeurlruteinguwozi-zuo-suru)

Despite my trial and error, I recently discovered a simple mistake that prevents it from functioning as URL routing. With a sense of repentance, I am writing this article to reorganize my thoughts on building custom URL routing.
*Note: I discovered a disappointing mistake regarding the handling of URL path parameters.*

# Motivation
Here are some thoughts on what motivates someone to build their own URL routing.

- Relatively low barrier to entry
  - URL routing is familiar to web engineers, so there might be less confusion about the specifications.
  - As long as you understand the basics of data structures (tree algorithms), you can relatively easily create a basic one (probably).
  - There are plenty of libraries in many languages that can serve as reference implementations. Hints and ideas are abundant on GitHub.
- Versatile
  - It can be applied to systems that use tree algorithms, such as implementing features that handle strings like keyword suggestions.

Although I listed these, most of them are results-oriented, and honestly, it was just because it seemed fun to try.

# What is URL Routing?
Taken literally, "URL" is a noun and "routing" is a verb, so "URL routing" refers to something that "routes URLs." But what does "routing URLs" mean?

Since this article claims to be an introduction, let's interpret each term one by one.

### URL
A URL is that string with lots of "/" that you often see in browsers. For example, https://www.google.com/

It represents the address (location) of a page on the internet and stands for Uniform Resource Locator.

The format of a URL string is defined as follows:

```
 <scheme>:<scheme-specific-part>
```

The `<scheme>` part often uses protocol names like http, https, or ftp, but other scheme names besides protocol names are also defined. cf. [Uniform Resource Identifier (URI) Schemes](https://www.iana.org/assignments/uri-schemes/uri-schemes.xhtml)

The `<scheme-specific-part>` part is defined based on the scheme. For example, in the case of http or https schemes, it is defined to include the domain name and path name (or directory name).

For detailed specifications of URLs, refer to RFC1738. cf. [rfc1738 - Uniform Resource Locators (URL)](https://tools.ietf.org/html/rfc1738)

RFC1738 is positioned as an internet standard (STD1).

### Routing
Routing, when translated into Japanese, means "path control."

In terms of the OSI reference model's network layer, routers that relay communication between networks fulfill the role of "routing." These routers have a routing table, a path table, and mediate the transfer of packets according to the path table, performing "routing."

The discussion of URL routing is not about the network layer, so let's return to the application layer.

In the application layer, "routing"* can be defined as receiving information called a URL from a browser request and mediating between the URL and processing to handle data based on the received URL.

*The term "routing" in the application layer might be misleading. It might be safer to read it as "URL routing."

Routers in the network layer have a data structure called a routing table, and routers in the application layer (URL routers) also have a data structure that allows them to distribute processing based on URLs. The discussion of what kind of data structure this is will be covered later.

# Building Your Own URL Routing
It might be more accurate to say "building your own URL router" rather than "building your own URL routing," but we'll treat them as synonymous.

![Screen Shot 2019-12-14 at 23 05 02](/assets/images/posts/url-routing-introduction-episode-1/70849757-3b91f300-1ec6-11ea-8cfe-b3880811fb5e.png)

The position of the router is like this.

Let's roughly define the requirements for the router here*.

- Support for static routing
    - For a URL like `/foo/bar/`, it should be able to match a static routing definition like `/foo/bar/` and return processing.
- Support for dynamic routing
    - For a URL like `/foo/bar/123`, it should be able to match a routing definition with dynamic path parameters like `/foo/bar/:id` and return processing in a way that can utilize the path parameter data.

As long as the above requirements are met, it should fulfill the basic functions of a router.

The specifications based on the requirements can be adjusted in various ways depending on the implementation, so we won't define them clearly here.*

*I plan to write about the implementation of URL routing in Golang on the 20th day of the [Qiita - Go6 Advent Calendar 2019](https://qiita.com/advent-calendar/2019/go6).

Now, to build a router, let's consider what kind of processing the router is doing from the perspective of data structures.

As I write this, it seems like the day is about to end, so I'll continue in tomorrow's article.

To Be Continued..
