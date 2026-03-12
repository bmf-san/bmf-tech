---
title: "Database ID Design: UUID vs Auto Increment vs ULID — Which to Choose?"
slug: id-design-in-db
date: 2024-03-30T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - DB
translation_key: id-design-in-db
---



# What is an ID
An ID is something that uniquely identifies an entity.

In set theory, a one-to-one correspondence between an entity and its attributes is called a bijection.

An entity that forms a bijection can function as an ID.

When considering IDs, it is necessary to consider whether it is for a single entity or multiple (group) entities, and if multiple, at what granularity? For example, when considering the attribute "crab," what type is it? What family does it belong to? What color is it? What is its name?

Even attributes that do not seem to have uniqueness at first glance can function as IDs in certain contexts. For example, an item that a specific group possesses only one of.

It is better to avoid designs that embed meaning into parts of an ID. For example, including identifiers like color in part of a product number.

Such designs can prevent meeting the first normal form, unnecessarily complicating queries and negatively impacting database design.

# Natural Key or Surrogate Key
A natural key uses real-world terms or labels as IDs, while a surrogate key uses something that only makes sense within the database or the application using it.

If a bijection holds, using a natural key is fine, but pay attention to the following points:

- The lifecycle of the ID
- Attributes that are unique but not bijective (e.g., an email address is unique but does not have a one-to-one relationship with a person)

For surrogate keys, pay attention to the following:

- Is there a natural key?
  - If it exists, use the natural key to avoid unnecessary overhead (even with a surrogate key, a unique constraint on the natural key is needed, causing unnecessary index updates)
  - Adding a surrogate key carelessly can create functional dependencies (if A is determined, B is determined) between the surrogate key and other attributes, complicating database design

# Additional Notes
When designing IDs with surrogate keys, there may be cases where physical design is considered. For example, whether to use auto-increment values or randomly unique values can affect performance.

# References
- [amzn.to - Learn Database Practice from Theory: Efficient SQL with Relational Model (WEB+DB PRESS plus)](https://amzn.to/3TEltzx)
- [nippondanji.blogspot.com - Discussion on Natural Keys and Surrogate Keys](https://nippondanji.blogspot.com/2013/12/blog-post_4.html)
- [nippondanji.blogspot.com - Discussion on Domain Design in Relational Models](https://nippondanji.blogspot.com/2013/12/blog-post_8.html)
