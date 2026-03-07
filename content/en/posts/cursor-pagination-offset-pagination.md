---
title: Cursor Pagination vs Offset Pagination
slug: cursor-pagination-offset-pagination
date: 2024-10-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Offset Pagination
  - Cursor Pagination
description: A comparison between offset pagination and cursor pagination.
translation_key: cursor-pagination-offset-pagination
---

# Overview
This post summarizes the comparison between offset pagination and cursor pagination.

# What is Offset Pagination?
Offset pagination is implemented using `OFFSET` and `LIMIT`, as shown in the query `SELECT * FROM table LIMIT 10 OFFSET 20`.

It is relatively easy to implement and allows direct access to any page. While it provides a clear understanding of the total number of pages, performance can degrade when dealing with large datasets.

# What is Cursor Pagination?
Cursor pagination is implemented using queries like `SELECT * FROM table WHERE id > 20 ORDER BY id LIMIT 10`.

This method is effective when the order of data is important and remains stable even when the data is frequently updated. However, accessing arbitrary pages is challenging, and understanding the total number of pages is difficult.

# Comparison
| Feature              | Offset Pagination                              | Cursor Pagination                              |
|----------------------|-----------------------------------------------|-----------------------------------------------|
| **Advantages**       | Simple implementation                        | High performance with large datasets          |
|                      | Direct access to any page                    | Less affected by frequent data updates        |
|                      | Easy to understand total page count          | Ideal for data requiring order                |
| **Disadvantages**    | Performance degradation (especially later pages) | Difficult to access arbitrary pages           |
|                      | Prone to instability with frequent updates   | Slightly more complex implementation          |
|                      |                                               | Difficult to understand total page count      |

# Solutions to Disadvantages
## Offset Pagination
To address performance degradation, consider the following ideas:

- **Index Optimization**
- **Lazy Cursor**
  - If supported by the database, using a cursor to fetch data incrementally may improve performance.
- **Caching**
  - Caching the pagination results may enhance performance.

To address instability in data, consider the following ideas:

- **Ensuring Consistency**
  - Taking a snapshot of the data can reduce the impact of changes.
- **Tracking State Between Pages**
  - Tracking state between pages can prevent duplication or missing data as pagination progresses.

## Cursor Pagination
To address the difficulty of accessing arbitrary pages, consider the following ideas:

- **Index Optimization**
- **Using Snapshots**
  - Taking a snapshot of the data can reduce the impact of changes.
- **Combining Cursor and Offset**

To address the complexity of implementation, consider the following ideas:

- **Using Libraries**
- **Lightweight Cursors**
  - Simplifying cursor information can reduce implementation complexity.