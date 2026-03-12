---
title: "Forward Proxy vs Reverse Proxy vs API Gateway: A Clear Comparison"
slug: proxy-server-differences
date: 2023-06-03T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Forward Proxy
  - Reverse Proxy
  - Gateway
translation_key: proxy-server-differences
---

# Overview
This post summarizes the differences between Forward Proxy, Reverse Proxy, and Gateway servers.

# What is a Forward Proxy Server?
A server that is placed between the client and server, acting as an intermediary for their communication (relay). It proxies requests from the client.

Client → Proxy Server → Internet → Server

The benefits include:

- Log acquisition
- Caching
- Virus protection
- Filtering (Access control)
- Ensuring anonymity

# What is a Reverse Proxy Server?
A server that is placed between the client and server, acting as an intermediary for their communication (relay). It proxies requests from the server.

Client → Internet → Reverse Proxy → Server

The benefits include:

- Log acquisition
- Caching
- Virus protection
- Load balancing
- SSL offloading
- Protection against unauthorized access

# What is a Gateway Server?
A server that is placed between the client and server, relaying and controlling the conversion of networks and protocols.

The difference from Forward Proxy and Reverse Proxy is that a proxy acts as an intermediary for communication, while a gateway establishes the communication.

# Summary
Forward Proxy, Reverse Proxy, and Gateway servers differ in their placement and whether they act as intermediaries or controllers of communication.

# References
- [What are the differences between Reverse Proxy and Forward Proxy?](https://laplace-daemon.com/reverse-forward-proxy/)
- [What are the differences between Proxy and Gateway?](https://laplace-daemon.com/difference-between-proxy-and-gateway/)
- [What are the differences between Reverse Proxy and Forward Proxy?](https://laplace-daemon.com/reverse-forward-proxy/)
- [What are the differences between Reverse Proxy and Proxy? How do each server work?](https://eset-info.canon-its.jp/malware_info/special/detail/201021.html)
- [What is a Proxy? Explanation of Forward Proxy and Reverse Proxy mechanisms](https://siteguard.jp-secure.com/blog/reverse-proxy)
- [A simple explanation of the differences between Reverse Proxy and Forward Proxy](https://log.dot-co.co.jp/reverseproxy-apache/)
- [What is a Reverse Proxy? | Explanation of Proxy Servers](https://www.cloudflare.com/ja-jp/learning/cdn/glossary/reverse-proxy/)
- [Differences between Reverse Proxy and Proxy, their mechanisms and roles](https://ascii.jp/elem/000/004/030/4030703/)
- [Reverse Proxy 【reverse proxy】 Reverse Proxy Server / RP Server](https://e-words.jp/w/%E3%83%AA%E3%83%90%E3%83%BC%E3%82%B9%E3%83%97%E3%83%AD%E3%82%AD%E3%82%B7.html)
- [Gateway 【gateway】 GW](https://e-words.jp/w/%E3%82%B2%E3%83%BC%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A4.html)
- [Detailed explanation of the basics and use cases of gateways](https://amnimo.com/column/001/)
- [What is a Proxy Server? Overview, roles, and benefits explained](https://www.winserver.ne.jp/column/about_proxy-server/#:~:text=%E3%81%93%E3%81%AE%E3%82%88%E3%81%86%E3%81%AB%E3%80%81%E3%83%97%E3%83%AD%E3%82%AD%E3%82%B7%E3%82%B5%E3%83%BC%E3%83%90,%E4%BB%95%E7%B5%84%E3%81%BF%E3%81%A8%E3%81%AA%E3%81%A3%E3%81%A6%E3%81%84%E3%81%BE%E3%81%99%E3%80%82)