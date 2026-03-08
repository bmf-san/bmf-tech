---
title: About Cursor Pagination and Offset Pagination
slug: cursor-pagination-offset-pagination
date: 2024-10-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Offset Pagination
  - Cursor Pagination
translation_key: cursor-pagination-offset-pagination
---

# Overview
A summary comparing offset pagination and cursor pagination.

# What is Offset Pagination?
A method of implementing pagination using `OFFSET` and `LIMIT`, such as `SELECT * FROM table LIMIT 10 OFFSET 20`.

It is relatively easy to implement and allows direct access to any page. It is easy to grasp the total number of pages, but performance can degrade with large datasets.

# What is Cursor Pagination?
A method of implementing pagination like `SELECT * FROM table WHERE id > 20 ORDER BY id LIMIT 10`.

It is effective when the order of data is important and operates stably even when data is frequently updated. However, accessing arbitrary pages can be difficult, and it is challenging to grasp the total number of pages.

# Comparison
| Feature               | Offset Pagination                             | Cursor Pagination                             |
|-----------------------|----------------------------------------------|----------------------------------------------|
| **Advantages**        | Simple implementation                         | High performance even with large datasets    |
|                       | Direct access to any page                    | Less instability due to data updates         |
|                       | Easy to grasp the total number of pages     | Optimal for ordered data                     |
| **Disadvantages**     | Performance degradation (especially on later pages) | Difficult access to arbitrary pages          |
|                       | Can become unstable with frequent data updates | Implementation is somewhat complex           |
|                       |                                              | Difficult to grasp the total number of pages |

# Solutions to Disadvantages
## Offset Pagination
Ideas to solve performance degradation include:

- Index optimization
- Lazy cursor
  - If supported by the database, using a cursor function to fetch data sequentially may improve performance.
- Caching
  - Caching the results of pagination may improve performance.

Ideas to resolve data instability include:

- Consistency guarantee
  - Taking a snapshot of the data can reduce the impact of changes.
- Tracking state between pages
  - Considering the possibility of data changes as pagination progresses, tracking state between pages can prevent duplicates and data loss.

## Cursor Pagination
Ideas to solve the difficulty of accessing arbitrary pages include:

- Index optimization
- Utilizing snapshots
  - Taking a snapshot of the data can reduce the impact of changes.
- Combining cursors and offsets

Ideas to resolve implementation complexity include:

- Using libraries
- Lightweight cursors
  - Reducing the information in the cursor can alleviate implementation complexity.