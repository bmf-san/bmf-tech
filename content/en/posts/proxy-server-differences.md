---
title: Differences Between Proxy Servers (Forward Proxy), Reverse Proxy Servers, and Gateway Servers
slug: proxy-server-differences
date: 2023-06-03T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Forward Proxy
  - Reverse Proxy
  - Gateway
description: A summary of the differences between forward proxy servers, reverse proxy servers, and gateway servers.
translation_key: proxy-server-differences
---

# Overview
This post summarizes the differences between proxy servers (forward proxy servers), reverse proxy servers, and gateway servers.

# What is a Proxy Server (Forward Proxy Server)?
A server placed between a client and a server that acts as an intermediary for their communication. It handles requests on behalf of the client.

Client → Proxy Server → Internet → Server

The benefits include:

- Log collection
- Caching
- Virus protection
- Filtering (access control)
- Ensuring anonymity

# What is a Reverse Proxy Server?
A server placed between a client and a server that acts as an intermediary for their communication. It handles requests on behalf of the server.

Client → Internet → Reverse Proxy → Server

The benefits include:

- Log collection
- Caching
- Virus protection
- Load balancing
- SSL offloading
- Protection against unauthorized access

# What is a Gateway Server?
A server placed between a client and a server that intermediates and controls network or protocol conversion between them.

The difference between a proxy server (forward proxy server) and a reverse proxy server is that a proxy "acts on behalf" of communication, while a gateway "enables" communication.

# Summary
Forward proxy servers, reverse proxy servers, and gateway servers differ in their placement and whether they act as intermediaries or controllers of communication.

# References
- [【Diagram】What’s the Difference Between Reverse Proxy and Forward Proxy?](https://laplace-daemon.com/reverse-forward-proxy/)
- [【Diagram】What’s the Difference Between Proxy and Gateway?](https://laplace-daemon.com/difference-between-proxy-and-gateway/)
- [【Diagram】What’s the Difference Between Reverse Proxy and Forward Proxy?](https://laplace-daemon.com/reverse-forward-proxy/)
- [What’s the Difference Between Reverse Proxy and Proxy? How Do These Servers Work?](https://eset-info.canon-its.jp/malware_info/special/detail/201021.html)
- [What is a Proxy? Explaining the Mechanisms of Forward and Reverse Proxies](https://siteguard.jp-secure.com/blog/reverse-proxy)
- [A Simple Explanation of the Difference Between Reverse Proxy and Forward Proxy](https://log.dot-co.co.jp/reverseproxy-apache/)
- [What is a Reverse Proxy? | Explanation of Proxy Servers](https://www.cloudflare.com/ja-jp/learning/cdn/glossary/reverse-proxy/)
- [The Difference Between Reverse Proxy and Proxy: Their Mechanisms and Roles](https://ascii.jp/elem/000/004/030/4030703/)
- [Reverse Proxy 【reverse proxy】 Reverse Proxy / RP Server / RevProxy](https://e-words.jp/w/%E3%83%AA%E3%83%90%E3%83%BC%E3%82%B9%E3%83%97%E3%83%AD%E3%82%AD%E3%82%B7.html)
- [Gateway 【gateway】 GW](https://e-words.jp/w/%E3%82%B2%E3%83%BC%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A4.html)
- [Detailed Explanation of Gateway Basics and Use Cases](https://amnimo.com/column/001/)
- [What is a Proxy Server? Overview, Roles, and Benefits Explained](https://www.winserver.ne.jp/column/about_proxy-server/#:~:text=%E3%81%93%E3%81%AE%E3%82%88%E3%81%86%E3%81%AB%E3%80%81%E3%83%97%E3%83%AD%E3%82%AD%E3%82%B7%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC,%E4%BB%95%E7%B5%84%E3%81%BF%E3%81%A8%E3%81%AA%E3%81%A3%E3%81%A6%E3%81%84%E3%81%BE%E3%81%99%E3%80%82)