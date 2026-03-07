---
title: Types of Relationships in ER Diagrams
slug: er-diagram-relationship-types
date: 2018-07-31T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - ER
description: A summary of the types of relationships in ER diagrams.
translation_key: er-diagram-relationship-types
---

# Overview
This post summarizes the types of relationships in ER diagrams.

# Types of Relationships in ER Diagrams
There are three types of relationships in ER diagrams.

| Relationship Type         | Table Relationship                                  | Line Type          |
|---------------------------|----------------------------------------------------|--------------------|
| Dependent Relationship    | Child table depends on the parent table (parent-child relationship between tables) | Solid line (Parent → Child) |
| Independent Relationship  | Child table does not depend on the parent table (no parent-child relationship between tables) | Dashed line (Parent → Child) |
| Many-to-Many Relationship | Many-to-many table relationship                    | Solid line (Parent ↔ Child) |

# Dependent Relationship

- User
  - UserNo (PK)
  - CompanyNo (FK)
  - Name
  - Email Address

- User Profile
  - UserNo (FK)
  - Age
  - Gender

The child table, User Profile, cannot exist without a record in the parent table, User. Therefore, the child table is said to depend on the parent table.

# Independent Relationship

- User
  - UserNo (PK)
  - CompanyNo (FK)
  - Name
  - Email Address

- User Profile
  - UserNo (FK)
  - Age
  - Gender

- Company
   - CompanyNo (PK)
   - Company Name

As explained earlier, the relationship between User and User Profile is a dependent relationship.

The relationship between User and Company, however, is independent because a company can exist without a user, and vice versa.

# Many-to-Many Relationship
- User
  - UserNo (PK)
  - CompanyNo (FK)
  - PermissionNo (FK)
  - Name
  - Email Address

- Permission
  - PermissionNo (PK)
  - UserNo (FK)

This type of relationship requires a so-called cross (intermediate or pivot) table.

# Thoughts
I want to dive deeper into this topic by reading more books.

# References
- [ER Diagram: Properly Differentiating Relationships](https://products.sint.co.jp/siob/blog/relationship)
- [A Must-Read for Junior Programmers! Understand ER Diagram Creation in 5 Minutes](https://it-koala.com/entity-relationship-diagram-1897)