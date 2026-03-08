---
title: Stateless and Stateful
slug: stateless-vs-stateful
date: 2018-04-17T00:00:00Z
author: bmf-san
categories:
  - Networking
tags:
  - Stateful
  - Stateless
  - Session
translation_key: stateless-vs-stateful
---


# Prerequisites

-   Session
    -   HTTP is a stateless protocol
        -   It does not determine whether requests are from the same client
    -   A series of communications from connection establishment to disconnection

# Stateful

-   The server maintains the client's session state
-   The system retains state and data
-   Examples of protocols
    -   FTP, TCP, BGP, OSPF, EIGRP, SMTP, SSH

# Stateless

-   The server does not maintain the client's session information
-   The system does not retain state or data, determines output from input values
-   Examples of protocols
    -   HTTP, UDP, IP, DNP

# References

-   [What does Stateless and Stateful mean](http://blog.sojiro.me/blog/2014/09/13/stateful-and-stateless/)
-   [Differences between Stateful and Stateless](https://milestone-of-se.nesuke.com/nw-basic/as-nw-engineer/stateful-and-stateless/)
-   [Session State](https://msdn.microsoft.com/ja-jp/library/aa720612)
