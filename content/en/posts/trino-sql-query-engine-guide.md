---
title: Comprehensive Guide to the Distributed SQL Query Engine Trino
slug: trino-sql-query-engine-guide
date: 2024-10-24T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Trino
  - Book
translation_key: trino-sql-query-engine-guide
---

[Comprehensive Guide to the Distributed SQL Query Engine Trino](https://amzn.to/3BXhPeQ) has been read.

- SQL Query Engine
- Supports ANSI SQL
- Federated Queries
  - Reference and use databases and schemas from different systems with the same SQL
- Not a database
- Not designed to handle OLTP
  - Since it is aimed at OLAP, the performance target is likely at the OLAP level
- The ability to scale computing resources as a distributed system is similar to the nature of New SQL
- Handles data sources in a way that differs from traditional big data approaches (which require various query languages and tools, as well as costly data warehouses for operation and maintenance)