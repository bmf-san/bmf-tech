---
title: "The N+1 Problem Explained: How to Detect and Fix It in Go/Rails"
description: 'Learn what the N+1 query problem is, why it degrades application performance, how to detect it with query logs, and how to fix it with eager loading in Go and Rails.'
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
A summary of the N+1 problem and how to address it.

# What is the N+1 Problem
- The issue where 1 SQL query is issued for retrieving all records plus N SQL queries for each record.
- It is easier to understand if interpreted as 1+N rather than N+1.

# Example
- Case of retrieving data for list display
  - Issue 1 SELECT query to retrieve all data for the list (returns N records)
  - Issue N SELECT queries to retrieve related data for N records

# Solutions
- Join
  - `SELECT "users".* FROM "users" INNER JOIN "posts" ON "posts"."user_id" = "users"."id" WHERE "posts"."id" = 1`
  - Eager Loading
    - `SELECT "users".* FROM "users"`
    - `SELECT "posts".* FROM "posts" WHERE "posts"."id" IN (1, 2, 3, 4, 5)`
 
# References
- [N+1問題は1+N問題](https://qiita.com/hisonl/items/763b9d6d4e90b1606635)
- [N+1 問題](http://www.techscore.com/blog/2012/12/25/rails%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E7%B4%B9%E4%BB%8B-n1%E5%95%8F%E9%A1%8C%E3%82%92%E6%A4%9C%E5%87%BA%E3%81%99%E3%82%8B%E3%80%8Cbullet%E3%80%8D/)
- [N+1問題 / Eager Loading とは](http://ruby-rails.hatenadiary.com/entry/20141108/1415418367)
- [ActiveRecordのjoinsとpreloadとincludesとeager_loadの違い](https://qiita.com/k0kubun/items/80c5a5494f53bb88dc58)
