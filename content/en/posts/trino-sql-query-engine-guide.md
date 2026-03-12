---
title: Comprehensive Guide to Distributed SQL Query Engine Trino
description: An in-depth exploration of Comprehensive Guide to Distributed SQL Query Engine Trino, covering design principles, trade-offs, and practical applications.
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



[Comprehensive Guide to Distributed SQL Query Engine Trino](https://amzn.to/3BXhPeQ) was read.

- SQL Query Engine
- Supports ANSI SQL
- Federated Query
  - Reference and use databases and schemas from different systems with the same SQL
- Not a database
- Not designed to handle OLTP
  - Since it's for OLAP, the performance target is likely at the OLAP level
- As a distributed system, it can scale computing resources up and down, similar to New SQL
- Handles data sources differently from traditional big data approaches (which require costly data warehouses with various query languages, tools, operations, and maintenance costs)
