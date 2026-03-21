---
title: 'A Comprehensive Guide to the Distributed SQL Query Engine Trino'
description: 'A Comprehensive Guide to the Distributed SQL Query Engine Trino'
slug: trino-sql-query-engine-guide
date: 2024-10-24T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Trino
  - Book Review
translation_key: trino-sql-query-engine-guide
books:
  - asin: "4798071676"
    title: "A Comprehensive Guide to the Distributed SQL Query Engine Trino"
---


I read the [A Comprehensive Guide to the Distributed SQL Query Engine Trino](https://amzn.to/3BXhPeQ).

- SQL query engine
- Supports ANSI SQL
- Federated queries
  - Allows referencing and using databases and schemas from different systems with the same SQL
- Not a database
- Not designed for processing OLTP
  - Since it is aimed at OLAP, the performance targets are likely at the OLAP level
- The ability to scale computing resources as a distributed system resembles the nature of New SQL
- Handles data sources in a way that differs from traditional big data approaches (which require various query languages and tools, along with costly operational and maintenance expenses for data warehouses)
