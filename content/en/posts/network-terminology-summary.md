---
title: Summary of Network Terms
slug: network-terminology-summary
date: 2020-11-02T00:00:00Z
author: bmf-san
categories:
  - Network
tags:
  - IP
  - Load Balancer
description: A compilation of network-related terms with unclear knowledge.
translation_key: network-terminology-summary
---

# Overview
A compilation of terms related to networks where knowledge is unclear.

## IPv4
- Internet Protocol version 4
- A type of IP (Internet Protocol)
- A protocol positioned at the network layer in the OSI reference model
- Has a 32-bit address space
  - 2 to the power of 32 = 4,294,967,296 is the total number of IPv4 IP addresses
- Depletion has occurred with the spread of the internet

cf.
- [wikipedia -IPv4](https://ja.wikipedia.org/wiki/IPv4)
- [jprs.jp - 用語辞典 IPv4](https://jprs.jp/glossary/index.php?ID=0034)

## IPv6
- Internet Protocol version 6
- A type of IP (Internet Protocol)
- A protocol positioned at the network layer in the OSI reference model
- Has a 128-bit address space
  - 2 to the power of 128 = approximately 340 undecillion (1 undecillion is 1 trillion * 1 trillion * 1 trillion) is the total number of IPv6 IP addresses
- IPSec (encryption of IP packet data) is a standard feature
  - Optional in IPv4

cf.
- [wikipedia -IPv6](https://ja.wikipedia.org/wiki/IPv6)
- [jprs.jp - 用語辞典 IPv6](https://jprs.jp/glossary/index.php?ID=0035)
- [www.nic.ad.jp - IPv6](https://www.nic.ad.jp/ja/newsletter/No20/sec0700.html)

## NAT
- Technology for converting one IP address to another
- Used in routers and wireless LAN access points
- Static NAT
  - Converts IP addresses 1:1
  - Mainly used for converting global IP to private IP
- Dynamic NAT
  - Converts IP addresses 1:N
- Examples of NAT (NAPT) usage
  - When accessing the internet from a network using private IP addresses
  - When publishing a server with a private IP address to the outside in a company
  - When IP addresses overlap within a company

cf.
- [wikipedia - ネットワークアドレス変換](https://ja.wikipedia.org/wiki/%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9%E5%A4%89%E6%8F%9B)
- [www.infraexpert.com - NAT（Network Address Translation)](https://www.infraexpert.com/study/ip10.html)
- [locked.jp](https://locked.jp/blog/what-is-nat/)
- [milestone-of-se.nesuke.com - 【図解】初心者にも分かるNATの仕組みと種類 ~静的/動的NAT/NAPT(PAT),セキュリティ等メリット/デメリット～](https://milestone-of-se.nesuke.com/nw-basic/nat/nat-summary/)

## NAPT
- Technology for converting one IP address to another
- Used in routers and wireless LAN access points
- Called IP masquerade in Linux implementations
- Static NAPT
  - Converts one IP address + TCP/UDP port number to another IP address + port number 1:1
- Dynamic NAPT
  - Converts IP addresses 1:N
  - Converts a source IP address to one global IP N:1
    - Converts the source port number to an unused port

cf.
- [wikipedia - ネットワークアドレス変換](https://ja.wikipedia.org/wiki/%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9%E5%A4%89%E6%8F%9B)
- [locked.jp](https://locked.jp/blog/what-is-nat/)
- [milestone-of-se.nesuke.com - 【図解】初心者にも分かるNATの仕組みと種類 ~静的/動的NAT/NAPT(PAT),セキュリティ等メリット/デメリット～](https://milestone-of-se.nesuke.com/nw-basic/nat/nat-summary/)

## CIDR
- Classless Inter-Domain Routing
- Technology for IP address allocation and routing without using classes
- A mechanism to reduce the expansion of routing tables in routers on the internet
- The method of allocating and routing IP addresses by class is not scalable
  - To address this issue, standards defining the allocation method of IP address blocks have emerged (RFC1518, RFC1519, etc.)
- By using a variable-length subnet mask in 1-bit units, the size of the network part of the IP address can be specified to change the size of the allocated block
- CIDR notation
  - 192.168.1.0/24

cf. 
- [wikipedia](https://ja.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
- [www.nic.ad.jp - CIDRとは](https://www.nic.ad.jp/ja/basics/terms/cidr.html)

## Subnet Mask
- A numerical value to identify the network address and host address within an IP address
  - 32 bits for IPv4
  - 128 bits for IPv6

cf. 
- [wikipedia - サブネットマスク](https://ja.wikipedia.org/wiki/%E3%82%B5%E3%83%96%E3%83%8D%E3%83%83%E3%83%88%E3%83%9E%E3%82%B9%E3%82%AF)

## DHCP
- Dynamic Host Configuration Protocol
- A communication protocol used in IPv4 networks
  - Automatically assigns configuration information necessary for a computer to connect to a network
  - Used in both IPv4 and IPv6, but considered different protocols due to differences in protocol details
- Positioned at the application layer in the OSI reference model

cf. 
- [wikipedia - Dynamic Host Configuration Protocol](https://ja.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol)
- [www.nic.ad.jp - DHCPとは](https://www.nic.ad.jp/ja/basics/terms/dhcp.html)

## Link-Local Address
- An address valid only for communication within the network segment or broadcast domain to which the host is connected (within the network range where broadcasting is possible. Broadcasting sends the same data to all users on the same network.)
- A special IP address self-assigned in networks without a DHCP server

cf.
- [e-words.jp](http://e-words.jp/w/%E3%83%AA%E3%83%B3%E3%82%AF%E3%83%AD%E3%83%BC%E3%82%AB%E3%83%AB%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9.html)

## ICMP
- Internet Control Message Protocol
- A protocol for forwarding error notifications and control messages of the IP protocol
  - Used to check communication status
- Used in the ping command

cf.
- [wikipedia - Internet Control Message Protocol](https://ja.wikipedia.org/wiki/Internet_Control_Message_Protocol)
- [www.infraexpert.com - TCP/IP-ICMP](https://www.infraexpert.com/study/tcpip4.html)

## L4 Load Balancer
- L4 → Transport layer
- Load balancing based on IP address and port number is possible

cf. 
- [faq.support.nifcloud.com - ロードバランサー（L4）とL7ロードバランサー（Pulse Secure Virtual Traffic Manager）の違いを教えてください](https://faq.support.nifcloud.com/faq/show/420?site_domain=default)
- [www.kimullaa.com - L7ロードバランサとL4ロードバランサ](https://www.kimullaa.com/entry/2019/12/01/135430)

## L7 Load Balancer
- L7 → Application layer
- Load balancing based on URL and HTTP headers is possible

cf. 
- [faq.support.nifcloud.com - ロードバランサー（L4）とL7ロードバランサー（Pulse Secure Virtual Traffic Manager）の違いを教えてください](https://faq.support.nifcloud.com/faq/show/420?site_domain=default)
- [www.kimullaa.com - L7ロードバランサとL4ロードバランサ](https://www.kimullaa.com/entry/2019/12/01/135430)
