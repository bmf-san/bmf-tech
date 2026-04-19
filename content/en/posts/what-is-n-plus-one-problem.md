---
title: "The N+1 Problem Explained: How to Detect and Fix It in Go/Rails"
description: 'Learn what the N+1 query problem is, why it degrades application performance, how to detect it with query logs, and how to fix it with eager loading in Go and Rails.'
slug: what-is-n-plus-one-problem
date: 2018-05-12T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - Database
tags:
  - N+1
translation_key: what-is-n-plus-one-problem
---


# Overview
This article summarizes the explanation and response to the N+1 problem.

# What is the N+1 Problem?
- A problem where N SQL statements are issued to retrieve all records, resulting in 1 + N SQL statements.
- It is easier to interpret as 1 + N rather than N + 1.

# Example
- A case of retrieving data for display
  - Issue a SELECT statement once to retrieve all data for the list (returning N records).
  - Issue N SELECT statements to retrieve related data for the N records.

# Response
- JOIN
  - `SELECT "users".* FROM "users" INNER JOIN "posts" ON "posts"."user_id" = "users"."id" WHERE "posts"."id" = 1`
  - Eager Loading
    - `SELECT "users".* FROM "users"`
    - `SELECT "posts".* FROM "posts" WHERE "posts"."id" IN (1, 2, 3, 4, 5)`
 
# References
- [The N+1 Problem is the 1+N Problem](https://qiita.com/hisonl/items/763b9d6d4e90b1606635)
- [N+1 Problem](http://www.techscore.com/blog/2012/12/25/rails%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E7%B4%B9%E4%BB%8B-n1%E5%95%8F%E9%A1%8C%E3%82%92%E6%A4%9C%E5%87%BA%E3%81%99%E3%82%8B%E3%80%8Cbullet%E3%80%8D/)
- ~~What is the N+1 Problem / Eager Loading~~
- [Differences between ActiveRecord's joins, preload, includes, and eager_load](https://qiita.com/k0kubun/items/80c5a5494f53bb88dc58)
