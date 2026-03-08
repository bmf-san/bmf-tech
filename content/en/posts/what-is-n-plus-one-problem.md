---
title: What is the N+1 Problem
slug: what-is-n-plus-one-problem
date: 2018-05-12T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - N+1
translation_key: what-is-n-plus-one-problem
---

# Overview
This post summarizes the explanation and solutions for the N+1 problem.

# What is the N+1 Problem
- A problem where N SQL queries are issued to retrieve all records, plus one additional query.
- It is easier to interpret it as 1+N rather than N+1.

# Example
- A case for retrieving data for a list display.
  - Issue one SELECT to retrieve all data for the list (returns N records).
  - Issue N SELECTs to retrieve related data for the N records.

# Solutions
- Join
  - `SELECT "users".* FROM "users" INNER JOIN "posts" ON "posts"."user_id" = "users"."id" WHERE "posts"."id" = 1`
  - Eager Loading
    - `SELECT "users".* FROM "users"`
    - `SELECT "posts".* FROM "posts" WHERE "posts"."id" IN (1, 2, 3, 4, 5)`
 
# References
- [N+1 Problem is 1+N Problem](https://qiita.com/hisonl/items/763b9d6d4e90b1606635)
- [N+1 Problem](http://www.techscore.com/blog/2012/12/25/rails%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E7%B4%B9%E4%BB%8B-n1%E5%95%8F%E9%A1%8C%E3%82%92%E6%A4%9C%E5%87%BA%E3%81%99%E3%82%8B%E3%80%8Cbullet%E3%80%8D/)
- [What is N+1 Problem / Eager Loading](http://ruby-rails.hatenadiary.com/entry/20141108/1415418367)
- [Differences between ActiveRecord's joins, preload, includes, and eager_load](https://qiita.com/k0kubun/items/80c5a5494f53bb88dc58)