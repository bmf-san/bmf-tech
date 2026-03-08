---
title: Approaches to Table Design for Many-to-One Relationships Across Multiple Tables
slug: table-design-multiple-tables-many-to-one
date: 2018-08-06T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Polymorphic
  - SQL Anti-pattern
translation_key: table-design-multiple-tables-many-to-one
---

# Overview
This post summarizes table design patterns when multiple tables are related in a many-to-one relationship.

# Data Design
We will use the following case as an example of data design.

- issues
    - id
    - title

- pullrequests
    - id
    - title

- comments
    - id
    - content

A case where `comments` relates to both `issues` and `pullrequests` in a many-to-one manner.

# Polymorphic Relationships

- issues
    - id
    - title

- pullrequests
    - id
    - title

- comments
    - id
    - content
    - target_table
    - target_id

This table design adds columns `target_table` and `target_id` to `comments`, allowing it to determine whether it is associated with `issues` or `pullrequests`.

This is discussed as one of the anti-patterns in SQL anti-pattern literature.

Since `target_id` cannot determine whether it relates to `issues` or `pullrequests` without looking at `target_table`, **foreign key constraints cannot be used**. Therefore, in this pattern, **maintaining consistency between tables depends on application logic**.

While ORM frameworks like Laravel and Rails support polymorphic relationships, making implementation easier, this is a pattern that should generally be avoided.

# Junction (Pivot, Intermediate) Tables

- issues
    - id
    - title

- pullrequests
    - id
    - title

- issues_comments
   - issues_id
   - comments_id

- pullrequests_comments
   - pullrequests_id
   - comments_id

- comments
    - id
    - content

This pattern prepares junction tables for `issues` and `pullrequests`, allowing foreign key constraints to be used.

`issues` and `issues_comments` have a one-to-many relationship, and `issues_comments` and `comments` have a many-to-one relationship. The same applies to `pullrequests`.

Depending on application requirements, it may not be possible to guarantee that a single comment is associated with only one of `issues` or `pullrequests`.

Since foreign keys can be used, this approach maintains consistency better than polymorphic relationships.

# Common Parent Table
- issues
    - id
    - post_id

- pullrequests
    - id
    - post_id

- posts
    - id
    - title

- comments
    - id
    - content
    - post_id

This pattern prepares a common parent table for `issues`, `pullrequests`, and `comments`.

It seems reasonable to define `posts` based on the concept of class table inheritance (essentially treating it as a base class). 
(Reference: [Single Table Inheritance, Class Table Inheritance, and Concrete Class Inheritance PofEAA](https://bmf-tech.com/posts/%E5%8D%98%E4%B8%80%E3%83%86%E3%83%BC%E3%83%95%E3%82%99%E3%83%AB%E7%B6%99%E6%89%BF%E3%83%BB%E3%82%AF%E3%83%A9%E3%82%B9%E3%83%86%E3%83%BC%E3%83%95%E3%82%99%E3%83%AB%E7%B6%99%E6%89%BF%E3%83%BB%E5%85%B7%E8%B1%A1%E3%82%AF%E3%83%A9%E3%82%B9%E7%B6%99%E6%89%BF%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6) )

`issues` and `posts` have a one-to-one relationship, and `posts` and `comments` have a one-to-many relationship. The same applies to `pull_requests`.
`posts` and `comments` are related in a one-to-many manner.

This guarantees that a single comment is associated with one post, but it does not guarantee that it is associated with only one of `issues` or `pullrequests`.

# Table Splitting

- issues
    - id
    - title

- pullrequests
    - id
    - title

- issue_comments
    - id
    - issues_id
    - content

- pullrequest_comments
    - id
    - pullrequests_id
    - content

This is a discussion that questions the original premise, suggesting that instead of consolidating `comments` into a single table, separate `comments` tables could be prepared for each, effectively splitting the tables.

# Thoughts
Relying on application logic increases the potential for human error, so I believe that a design philosophy that depends on the table structure for logic is generally a better pattern. In addition to application requirements, I want to choose the optimal pattern considering the intent of the queries.

- [Design Approaches for Many-to-One Relationships Across Multiple Tables](https://spice-factory.co.jp/development/has-and-belongs-to-many-table/)
- [Reading SQL Anti-patterns (About Polymorphic Relationships)](https://blog.motimotilab.com/?p=207)