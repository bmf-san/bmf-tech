---
title: About ID Design in Database Design
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
An identifier that uniquely specifies an object.

In set theory, a one-to-one correspondence between an object and its attributes is called a bijection.

Those that form a bijection function as IDs.

When considering IDs, it is necessary to think about whether it pertains to a singular object or multiple (groups), and if multiple, what granularity it has. For example, considering the attribute of crab, what kind is it? What family does it belong to? What color is it? What is its name?

Even attributes that seemingly lack uniqueness can function as IDs depending on the context. For example, items that a specific group possesses uniquely.

It is advisable to avoid designs that imbue meaning into parts of an ID. For example, including identifiers like color in the value of a product number.

Such designs can lead to failure in satisfying the first normal form, complicating queries unnecessarily and adversely affecting database design.

# Natural Key or Surrogate Key
A natural key uses words or labels that exist in the real world as IDs, while a surrogate key uses identifiers that are understood only within the database or the applications that utilize it.

If a bijection holds, using a natural key is acceptable, but attention should be paid to the following points:

- The lifecycle of the ID
- Items that are unique but do not form a bijection (e.g., email is unique but does not have a one-to-one relationship with a person)

For surrogate keys, the following points should be considered:

- Is there a natural key?
  - If one exists, use the natural key to avoid unnecessary overhead (even with a surrogate key, a unique constraint on the natural key is needed, leading to unnecessary index updates).
  - Adding a surrogate key carelessly can create functional dependencies (the property that if A is determined, B is also determined) between the surrogate key and another attribute, complicating database design.

# Aside
When designing IDs with surrogate keys, there are cases where physical design comes into play. There can be performance impacts depending on whether to use auto-incremented values or randomly selected unique values.

# References
- [amzn.to - Learning Database Practice from Theory ~ Efficient SQL with Relational Models (WEB+DB PRESS plus)](https://amzn.to/3TEltzx)
- [nippondanji.blogspot.com - Discussion on Natural Keys and Surrogate Keys](https://nippondanji.blogspot.com/2013/12/blog-post_4.html)
- [nippondanji.blogspot.com - Discussion on Domain Design in Relational Models](https://nippondanji.blogspot.com/2013/12/blog-post_8.html)