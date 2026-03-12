---
title: Types of Relationships in ER Diagrams
description: An in-depth look at Types of Relationships in ER Diagrams, covering key concepts and practical insights.
slug: er-diagram-relationship-types
date: 2018-07-31T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - ER
translation_key: er-diagram-relationship-types
---

# Overview
This post summarizes the types of relationships in ER diagrams.

# Types of Relationships in ER Diagrams
There are three types of relationships in ER diagrams.

|Type of Relationship|Table Relationship|Line Type|
|---|---|---|
|Dependent Relationship|Child table depends on parent table (parent-child relationship exists between tables)|Solid line (Parent → Child)|
|Independent Relationship|Child table does not depend on parent table (no parent-child relationship exists between tables)|Dotted line (Parent → Child)|
|Many-to-Many Relationship|Many-to-many table relationship|Solid line (Parent ↔ Child)|

# Dependent Relationship

- User
  - User No (PK)
  - Company No (FK)
  - Name
  - Email Address

- User Profile
  - User No (FK)
  - Age
  - Gender

The child table, User Profile, cannot exist without a record in the parent table, User, so it can be said that the child table depends on the parent table.

# Independent Relationship

- User
  - User No (PK)
  - Company No (FK)
  - Name
  - Email Address

- User Profile
  - User No (FK)
  - Age
  - Gender

- Company
   - Company No (PK)
   - Company Name

As explained earlier, the relationship between User and User Profile is a dependent relationship.

The relationship between User and Company is independent because a customer can exist without a user, and vice versa.

# Many-to-Many Relationship
- User
  - User No (PK)
  - Company No (FK)
  - Permission No (FK)
  - Name
  - Email Address

- Permission
  - Permission No (PK)
  - User No (FK)

This relationship requires a junction (intermediate, pivot) table.

# Thoughts
I want to read more books and delve deeper into this topic.

# References
- [Let's Properly Differentiate ER Diagram Relationships](https://products.sint.co.jp/siob/blog/relationship)
- [A Must-Read for Young Programmers! 5 Steps to Understand How to Draw ER Diagrams in 5 Minutes](https://it-koala.com/entity-relationship-diagram-1897)