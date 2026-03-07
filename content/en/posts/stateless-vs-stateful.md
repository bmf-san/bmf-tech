---
title: Stateless and Stateful
slug: stateless-vs-stateful
date: 2018-04-17T00:00:00Z
author: bmf-san
categories:
  - Network
tags:
  - Stateful
  - Stateless
  - Session
translation_key: stateless-vs-stateful
---

# Introduction

-   Session
    -   HTTP is a stateless protocol
        -   Does not determine if the request is from the same client
    -   A series of communications from connection establishment to disconnection

# Stateful

-   The server maintains the client's session state
-   The system retains state or data
-   Examples of protocols
    -   FTP, TCP, BGP, OSPF, EIGRP, SMTP, SSH

# Stateless

-   The server does not retain the client's session information
-   The system does not retain state or data, determines output from input values
-   Examples of protocols
    -   HTTP, UDP, IP, DNP

# References

-   [What Stateless and Stateful Mean](http://blog.sojiro.me/blog/2014/09/13/stateful-and-stateless/)
-   [Differences Between Stateful and Stateless](https://milestone-of-se.nesuke.com/nw-basic/as-nw-engineer/stateful-and-stateless/)
-   [Session State](https://msdn.microsoft.com/ja-jp/library/aa720612)