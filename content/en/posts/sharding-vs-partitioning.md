---
title: "Sharding vs Partitioning: Key Differences and When to Use Each"
description: 'Understand the difference between database sharding and partitioning. Learn horizontal and vertical strategies, sharding trade-offs, and when each approach makes sense.'
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
Sometimes it gets confusing which is which, so I’ll leave a note.

# Differences Between Sharding and Partitioning
|          Item          |                               Sharding                               |                            Partitioning                            |
| ---------------------- | ------------------------------------------------------------------- | ---------------------------------------------------------------- |
| Method of Data Division | Horizontal (≈ horizontal partitioning) ex. Row                     | Vertical ex. Table, Database, Column                              |
| Advantages             | Improved performance, enhanced scalability                          | Improved performance, better data searchability                   |
| Disadvantages          | Increased complexity of data and management                          | Issues with data separation and consistency                       |
| Suitable Applications   | Applications with large database sizes and declining performance     | Applications where database access patterns are biased towards specific columns |

# References
- [Database sharding vs partitioning](https://stackoverflow.com/questions/20771435/database-sharding-vs-partitioning)
- [What is the difference between sharding and partitioning? 【Distributed Database】](https://engineer.yeele.net/dev/db/understanding-term-partitioning-correctly-led-you-understand-sharding/)
