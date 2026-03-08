---
title: Differences Between Sharding and Partitioning
slug: sharding-vs-partitioning
date: 2023-06-03T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Partitioning
  - Sharding
translation_key: sharding-vs-partitioning
---


# Overview
Sometimes I get confused about which is which, so I'm leaving a note.

# Differences Between Sharding and Partitioning
|          Item          |                               Sharding                               |                            Partitioning                            |
| ---------------------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
| Data Division Method   | Horizontal (≒ Horizontal Partitioning) ex. Rows                                 | Vertical ex. Tables, Databases, Columns                              |
| Benefits               | Improved performance, enhanced scalability                               | Improved performance, enhanced data searchability                                 |
| Drawbacks              | Increased complexity of data and management                                                       | Issues with data separation and integrity                                               |
| Suitable Applications  | Applications where the database size is large and performance is degrading | Applications where database access patterns are biased towards specific columns |

# References
- [Database sharding vs partitioning](https://stackoverflow.com/questions/20771435/database-sharding-vs-partitioning)
- [shardingとpartitioningの違いは？【分散データベース】](https://engineer.yeele.net/dev/db/understanding-term-partitioning-correctly-led-you-understand-sharding/)