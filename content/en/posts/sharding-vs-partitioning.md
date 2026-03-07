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
Sometimes it's confusing to distinguish between the two, so I’m leaving a note.

# Differences Between Sharding and Partitioning
|          Item          |                               Sharding                               |                            Partitioning                            |
| ---------------------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
| Method of Data Split   | Horizontal (≈ horizontal partitioning) ex. Rows                                 | Vertical ex. Tables, Databases, Columns                              |
| Advantages             | Improved performance, increased scalability                               | Improved performance, enhanced data searchability                                 |
| Disadvantages          | Increased complexity of data and management                                                       | Issues with data separation and consistency                                               |
| Suitable Applications   | Applications with large database sizes and declining performance | Applications with access patterns biased towards specific columns |

# References
- [Database sharding vs partitioning](https://stackoverflow.com/questions/20771435/database-sharding-vs-partitioning)
- [What is the difference between sharding and partitioning? 【Distributed Database】](https://engineer.yeele.net/dev/db/understanding-term-partitioning-correctly-led-you-understand-sharding/)