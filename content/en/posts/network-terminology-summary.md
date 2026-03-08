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
translation_key: network-terminology-summary
---

# Overview
This post summarizes ambiguous terms related to networks.

## IPv4
- Internet Protocol version 4
- A type of IP (Internet Protocol)
- Positioned at the network layer in the OSI reference model
- Has a 32-bit address space
  - 2^32 = 4,294,967,296 total IPv4 addresses
- Has been depleting with the spread of the internet

cf.
- [wikipedia - IPv4](https://ja.wikipedia.org/wiki/IPv4)
- [jprs.jp - Glossary IPv4](https://jprs.jp/glossary/index.php?ID=0034)

## IPv6
- Internet Protocol version 6
- A type of IP (Internet Protocol)
- Positioned at the network layer in the OSI reference model
- Has a 128-bit address space
  - 2^128 = approximately 340 undecillion (1 undecillion = 10^36) total IPv6 addresses
- IPSec (which encrypts IP packet data) is a standard feature
  - Optional in IPv4

cf.
- [wikipedia - IPv6](https://ja.wikipedia.org/wiki/IPv6)
- [jprs.jp - Glossary IPv6](https://jprs.jp/glossary/index.php?ID=0035)
- [www.nic.ad.jp - IPv6](https://www.nic.ad.jp/ja/newsletter/No20/sec0700.html)

## NAT
- A technology that converts one IP address to another
- Used in routers and wireless LAN access points
- Static NAT
  - Converts IP addresses 1:1
  - Mainly used for converting global IPs to private IPs
- Dynamic NAT
  - Converts IP addresses 1:N
- Examples of NAT (NAPT)
  - When going from a network using private IP addresses to the internet
  - When exposing a server with a private IP address to the outside
  - When there are duplicate IP addresses within a company

cf.
- [wikipedia - Network Address Translation](https://ja.wikipedia.org/wiki/%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9%E5%A4%89%E6%8F%9B)
- [www.infraexpert.com - NAT (Network Address Translation)](https://www.infraexpert.com/study/ip10.html)
- [locked.jp](https://locked.jp/blog/what-is-nat/)
- [milestone-of-se.nesuke.com - 【Illustration】 Understanding NAT for Beginners: Static/Dynamic NAT/NAPT(PAT), Security, Advantages/Disadvantages](https://milestone-of-se.nesuke.com/nw-basic/nat/nat-summary/)

## NAPT
- A technology that converts one IP address to another
- Used in routers and wireless LAN access points
- The implementation in Linux is called IP masquerade
- Static NAPT
  - Converts one IP address + TCP/UDP port number to another IP address + port number 1:1
- Dynamic NAPT
  - Converts IP addresses 1:N
  - Converts the source IP address to one global IP in N:1
    - Converts the source port number to an unused port number

cf.
- [wikipedia - Network Address Translation](https://ja.wikipedia.org/wiki/%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9%E5%A4%89%E6%8F%9B)
- [locked.jp](https://locked.jp/blog/what-is-nat/)
- [milestone-of-se.nesuke.com - 【Illustration】 Understanding NAT for Beginners: Static/Dynamic NAT/NAPT(PAT), Security, Advantages/Disadvantages](https://milestone-of-se.nesuke.com/nw-basic/nat/nat-summary/)

## CIDR
- Classless Inter-Domain Routing
- A technology for assigning and routing IP addresses without classes
- A mechanism to reduce the bloat of routing tables in routers on the internet
- The method of classifying and assigning IP addresses for routing is not scalable
  - To address this issue, standards defining the allocation of IP address blocks were introduced (RFC1518, RFC1519, etc.)
- By using a variable-length subnet mask in bits, the size of the allocated block can be changed by specifying the size of the network part of the IP address
- CIDR notation
  - 192.168.1.0/24

cf.
- [wikipedia](https://ja.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
- [www.nic.ad.jp - What is CIDR](https://www.nic.ad.jp/ja/basics/terms/cidr.html)

## Subnet Mask
- A number used to identify the network address and host address within an IP address
  - 32 bits for IPv4
  - 128 bits for IPv6

cf.
- [wikipedia - Subnet Mask](https://ja.wikipedia.org/wiki/%E3%82%B5%E3%83%96%E3%83%8D%E3%83%83%E3%83%88%E3%83%9E%E3%82%B9%E3%82%AF)

## DHCP
- Dynamic Host Configuration Protocol
- A communication protocol used in IPv4 networks
  - Automatically assigns configuration information necessary for computers to connect to the network
  - Used in both IPv4 and IPv6, but considered a different protocol due to differences in details
- Positioned at the application layer in the OSI reference model

cf.
- [wikipedia - Dynamic Host Configuration Protocol](https://ja.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol)
- [www.nic.ad.jp - What is DHCP](https://www.nic.ad.jp/ja/basics/terms/dhcp.html)

## Link-Local Address
- An address valid only for communication within the network segment or broadcast domain (the range of the network where broadcast can occur. Broadcasting sends the same data to all users on the same network.)
- A special IP address that is self-assigned and used in networks without a DHCP server

cf.
- [e-words.jp](http://e-words.jp/w/%E3%83%AA%E3%83%B3%E3%82%AF%E3%83%AD%E3%83%BC%E3%82%AB%E3%83%AB%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9.html)

## ICMP
- Internet Control Message Protocol
- A protocol for transferring error notifications and control messages for the IP protocol
  - Used to check communication status
- Used by the ping command

cf.
- [wikipedia - Internet Control Message Protocol](https://ja.wikipedia.org/wiki/Internet_Control_Message_Protocol)
- [www.infraexpert.com - TCP/IP-ICMP](https://www.infraexpert.com/study/tcpip4.html)

## L4 Load Balancer
- L4 → Transport Layer
- Load balancing possible by IP address and port number

cf.
- [faq.support.nifcloud.com - What is the difference between L4 Load Balancer and L7 Load Balancer (Pulse Secure Virtual Traffic Manager)?](https://faq.support.nifcloud.com/faq/show/420?site_domain=default)
- [www.kimullaa.com - L7 Load Balancer vs L4 Load Balancer](https://www.kimullaa.com/entry/2019/12/01/135430)

## L7 Load Balancer
- L7 → Application Layer
- Load balancing possible by URL and HTTP headers

cf.
- [faq.support.nifcloud.com - What is the difference between L4 Load Balancer and L7 Load Balancer (Pulse Secure Virtual Traffic Manager)?](https://faq.support.nifcloud.com/faq/show/420?site_domain=default)
- [www.kimullaa.com - L7 Load Balancer vs L4 Load Balancer](https://www.kimullaa.com/entry/2019/12/01/135430)