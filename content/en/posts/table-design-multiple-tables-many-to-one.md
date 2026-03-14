---
title: Approaches to Table Design for Many-to-One Relationships with Multiple Tables
slug: table-design-multiple-tables-many-to-one
date: 2018-08-06T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Polymorphic
  - SQL Anti-pattern
description: 'Learn how to handle a table related to multiple tables in many-to-one relationships. Compares polymorphic associations (SQL anti-pattern, no FK constraints) with cross/pivot tables.'
translation_key: table-design-multiple-tables-many-to-one
---

# Overview
Summarizing table design patterns when a table is related to multiple tables in a many-to-one relationship.

# Data Design
Let's take the following case as an example of data design.

- issues
    - id
    - title

- pullrequests
    - id
    - title

- comments
    - id
    - content

A case where `comments` is related to both `issues` and `pullrequests` in a many-to-one relationship.

# Polymorphic Association

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

A table design where `comments` has columns `target_table` and `target_id` to determine whether it is linked to `issues` or `pullrequests`.

In SQL anti-patterns, this is highlighted as one of the anti-patterns.

Since `target_id` cannot determine whether it is related to `issues` or `pullrequests` without looking at `target_table`, **foreign key constraints cannot be used**. Therefore, in this pattern, **maintaining consistency between tables depends on the application's logic**.

Although polymorphic associations are supported in ORMs like Laravel and Rails, making implementation easier, this is a pattern that should generally be avoided.

# Cross (Pivot, Intermediate) Table

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

A pattern where cross tables are prepared for `issues` and `pullrequests` to enable the use of foreign key constraints.

`issues` and `issues_comments` have a one-to-many relationship, and `issues_comments` and `comments` have a many-to-one relationship. The same applies to `pullrequests`.

Depending on the application's requirements, it may not be possible to ensure that a single comment is linked to only one of `issues` or `pullrequests`.

Since foreign keys can be used, this pattern can maintain consistency better than polymorphic associations.

# Table with a Common Parent
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

A pattern where a table is prepared as a common parent for `issues`, `pullrequests`, and `comments`.

It seems good to define `posts` based on the concept of class table inheritance (essentially considering it as a base class).
(Reference: [Single Table Inheritance, Class Table Inheritance, Concrete Class Inheritance PofEAA](https://bmf-tech.com/posts/%E5%8D%98%E4%B8%80%E3%83%86%E3%83%BC%E3%83%95%E3%82%99%E3%83%AB%E7%B6%99%E6%89%BF%E3%83%BB%E3%82%AF%E3%83%A9%E3%82%B9%E3%83%86%E3%83%BC%E3%83%95%E3%82%99%E3%83%AB%E7%B6%99%E6%89%BF%E3%83%BB%E5%85%B7%E8%B1%A1%E3%82%AF%E3%83%A9%E3%82%B9%E7%B6%99%E6%89%BF%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6))

`issues` and `posts` have a one-to-one relationship, and `posts` and `comments` have a one-to-many relationship. The same applies to `pull_requests`.
`posts` and `comments` are related in a one-to-many relationship.

This pattern can ensure the constraint that a single comment is linked to one post, but it cannot ensure the constraint that it is linked to only one of `issues` or `pullrequests`.

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

This is a pattern that questions the initial premise, suggesting that instead of consolidating `comments` into a single table, separate `comments` tables should be prepared for each.

# Thoughts
Relying on the application's logic increases the possibility of human error, so a design policy that depends on the table structure for logic is generally a good pattern. In addition to the application's requirements, I want to be able to choose the optimal pattern by considering the query's perspective.

- [Approaches to Table Design for Many-to-One Relationships with Multiple Tables](https://spice-factory.co.jp/development/has-and-belongs-to-many-table/)
- [Reading SQL Anti-patterns (About Polymorphic Associations)](https://blog.motimotilab.com/?p=207)