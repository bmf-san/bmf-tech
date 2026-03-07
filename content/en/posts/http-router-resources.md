---
title: Materials Referenced for Creating an HTTP Router
slug: http-router-resources
date: 2023-10-30T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - router
  - HTTP
  - URL Routing
  - Resources
translation_key: http-router-resources
---

[goblin](https://github.com/bmf-san/goblin) is a project where I list various sources and my own articles that I referenced during development.

I had scattered reference links across several of my blog posts, so I decided to consolidate them.

# References
List of materials referenced for developing the HTTP Router.

## GitHub
- [jba/muxpatterns](https://github.com/jba/muxpatterns)
- [importcjj/trie-go](https://github.com/importcjj/trie-go)
- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [gorilla/mux](https://github.com/gorilla/mux)
- [gowww/router](https://github.com/gowww/router)
- [go-chi/chi](https://github.com/go-chi/chi)
- [go-ozzo/ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
- [nissy/bon](https://github.com/nissy/bon)
- [nissy/mux](https://github.com/nissy/mux)
- [ytakano/radix_tree](https://github.com/ytakano/radix_tree)
- [kkdai/radix](https://github.com/kkdai/radix)
- [MarkBaker/Tries](https://github.com/MarkBaker/Tries)
- [razonyang/routing](https://github.com/razonyang/routing)
- [ethereum/wiki - [Japanese] Patricia Tree](https://github.com/ethereum/wiki/wiki/%5BJapanese%5D-Patricia-Tree)
- [neo-nanikaka - CommonPrefixTrieRouter.php](https://gist.github.com/neo-nanikaka/c2e2f7742b311696d50b)
- [golang/go - proposal: net/http: enhanced ServeMux routing](https://github.com/golang/go/issues/61410)

## Blogs
- [blog.merovius.de - How to not use an http-router in go](https://blog.merovius.de/posts/2017-06-18-how-not-to-use-an-http-router/)
- [medium.com/@agatan - HTTP Server and context.Context](https://medium.com/@agatan/http%E3%82%B5%E3%83%BC%E3%83%90%E3%81%A8context-context-7211433d11e6)
- [devpixiv.hatenablog.com - Created a Fast URL Routing in PHP](https://devpixiv.hatenablog.com/entry/2015/12/13/145741)
- [kuune.org - Released the World's Fastest URL Router](https://kuune.org/text/2014/06/12/denco/)
- [takao.blogspot.com - Implemented PatriciaTrie in Java](https://takao.blogspot.com/2012/03/patriciatrie.html)
- [dankogai.livedoor.blog - algorithm - Patricia Trie (Radix Trie) in JavaScript](https://dankogai.livedoor.blog/archives/51766842.html)
- [persol-pt.github.io - Study Group [http request multiplexer and string matching]](https://persol-pt.github.io/posts/tech-workshop1222/)
- [atmarkit.itmedia.co.jp - The Difference Between Heaven and Hell Depends on Data Structure Selection](https://atmarkit.itmedia.co.jp/ait/articles/0809/01/news163_3.html)
- [www.sb.ecei.tohoku.ac.jp - Basic Data Structures: Traversing Tree Structures](http://www.sb.ecei.tohoku.ac.jp/lab/wp-content/uploads/2012/11/2012_d12.pdf)
- [noranuk0.hatenablog.com - Doing URL Routing Nicely in PHP Without a Framework](https://noranuk0.hatenablog.com/entry/2018/01/20/114933)
- [reiki4040.hatenablog.com - Components of a Golang HTTP Server](https://reiki4040.hatenablog.com/entry/2017/03/01/212647)
- [qiita.com/immrshc - 【Go】Understanding how http.HandleFunc is executed in the net/http package](https://qiita.com/immrshc/items/1d1c64d05f7e72e31a98)

## Documentation
- [urlpattern.spec.whatwg.org](https://urlpattern.spec.whatwg.org/)
  - The standard for URLPattern proposed by WHATWG
- [developer.mozilla.org - URL Pattern API](https://developer.mozilla.org/en-US/docs/Web/API/URL_Pattern_API)
  - Specifications for the URL Pattern API experimentally implemented on MDN

## Tools
- [www.cs.usfca.edu - Radix Tree](https://www.cs.usfca.edu/~galles/visualization/RadixTree.html)

# Articles Written
Articles posted on [bmf-tech.com](https://bmf-tech.com/).

- [Creating URL Routing Episode 1](https://bmf-tech.com/posts/URL%e3%83%ab%e3%83%bc%e3%83%86%e3%82%a3%e3%83%b3%e3%82%b0%e3%82%92%e3%81%a4%e3%81%8f%e3%82%8b%e3%80%80%e3%82%a8%e3%83%94%e3%82%bd%e3%83%bc%e3%83%891)
- [Creating URL Routing Episode 2](https://bmf-tech.com/posts/URL%e3%83%ab%e3%83%bc%e3%83%86%e3%82%a3%e3%83%b3%e3%82%b0%e3%82%92%e3%81%a4%e3%81%8f%e3%82%8b%e3%80%80%e3%82%a8%e3%83%94%e3%82%bd%e3%83%bc%e3%83%892)
- [Introduction to Creating URL Routing Episode 1](https://bmf-tech.com/posts/URL%e3%83%ab%e3%83%bc%e3%83%86%e3%82%a3%e3%83%b3%e3%82%b0%e3%82%92%e3%81%a4%e3%81%8f%e3%82%8b%e3%80%80%e3%82%a8%e3%83%94%e3%82%bd%e3%83%bc%e3%83%893%ef%bc%88%e5%ae%8c%e7%b5%90%e7%b7%a8%ef%bc%89)
- [Introduction to Creating URL Routing Episode 2](https://bmf-tech.com/posts/URL%e3%83%ab%e3%83%bc%e3%83%86%e3%82%a3%e3%83%b3%e3%82%b0%e8%87%aa%e4%bd%9c%e5%85%a5%e9%96%80%e3%80%80%e3%82%a8%e3%83%94%e3%82%bd%e3%83%bc%e3%83%892)
- [Created a URL Router called goblin in Golang](https://bmf-tech.com/posts/Golang%e3%81%a7goblin%e3%81%a8%e3%81%84%e3%81%86URL%e3%83%ab%e3%83%bc%e3%82%bf%e3%83%bc%e3%82%92%e8%87%aa%e4%bd%9c%e3%81%97%e3%81%9f)
- [Implemented a Benchmark for Comparing Go HTTP Routers](https://bmf-tech.com/posts/Go%e3%81%aeHTTP%20Router%e3%82%92%e6%af%94%e8%bc%83%e3%81%99%e3%82%8b%e3%83%99%e3%83%b3%e3%83%81%e3%83%9e%e3%83%bc%e3%82%ab%e3%83%bc%e3%82%92%e5%ae%9f%e8%a3%85%e3%81%97%e3%81%9f)
- [Improving Code Performance in Go](https://bmf-tech.com/posts/Go%e3%81%a7%e5%a7%8b%e3%82%81%e3%82%8b%e3%82%b3%e3%83%bc%e3%83%89%e3%81%ae%e3%83%91%e3%83%95%e3%82%a9%e3%83%bc%e3%83%9e%e3%83%b3%e3%82%b9%e6%94%b9%e5%96%84)
- [Introduction to Creating an HTTP Router with net/http](https://bmf-tech.com/posts/net%ef%bc%8fhttp%e3%81%a7%e3%81%a4%e3%81%8f%e3%82%8bHTTP%e3%83%ab%e3%83%bc%e3%82%bf%e3%83%bc%e8%87%aa%e4%bd%9c%e5%85%a5%e9%96%80)
- [Updated My Custom Routing](https://bmf-tech.com/posts/%e8%87%aa%e4%bd%9c%e3%83%ab%e3%83%bc%e3%83%86%e3%82%a3%e3%83%b3%e3%82%b0%e3%82%92%e3%82%a2%e3%83%83%e3%83%97%e3%83%87%e3%83%bc%e3%83%88%e3%81%97%e3%81%9f)
- [Created a URL Router called goblin in Golang](https://bmf-tech.com/posts/Golang%e3%81%a7goblin%e3%81%a8%e3%81%84%e3%81%86URL%e3%83%ab%e3%83%bc%e3%82%bf%e3%83%bc%e3%82%92%e8%87%aa%e4%bd%9c%e3%81%97%e3%81%9f)
- [Code Reading of Golang's HTTP Server](https://bmf-tech.com/posts/Golang%e3%81%aeHTTP%e3%82%b5%e3%83%bc%e3%83%90%e3%83%bc%e3%81%ae%e3%82%b3%e3%83%bc%e3%83%89%e3%83%aa%e3%83%87%e3%82%a3%e3%83%b3%e3%82%b0)
- [Implementing a Trie in Golang](https://bmf-tech.com/posts/Golang%e3%81%a7%e3%83%88%e3%83%a9%e3%82%a4%e6%9c%a8%e3%82%92%e5%ae%9f%e8%a3%85%e3%81%99%e3%82%8b)

# Books Written
- [Introduction to Creating an HTTP Router with net/http](https://zenn.dev/bmf_san/books/3f41c5cd34ec3f)