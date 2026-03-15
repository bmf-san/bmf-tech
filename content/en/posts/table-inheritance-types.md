---
title: Single Table Inheritance, Class Table Inheritance, and Concrete Class Inheritance
slug: table-inheritance-types
date: 2017-10-01T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - Application
tags:
  - PofEAA
description: 'Explaining three patterns to represent object inheritance in databases: Single Table Inheritance, Class Table Inheritance, and Concrete Class Inheritance.'
translation_key: table-inheritance-types
---

# Overview
**Relational databases do not support inheritance, so it is necessary to consider how to represent object inheritance relationships in the database.** This post explains three patterns to express this: Single Table Inheritance, Class Table Inheritance, and Concrete Class Inheritance.

*Note: This post does not discuss the advantages and disadvantages of implementing each pattern.*

# Prerequisites
We will consider four classes in this example.

![people_class.png](/assets/images/posts/table-inheritance-types/7680ef3c-3c9f-6365-8622-364ed30936b3.png)

The structure where Party People inherits from Rich People might be a bit confusing, but as long as the concept is conveyed, it’s fine.

#### People
This class has attributes common to all classes.

- `id` is a unique key.
- `name` is the name.

#### OrdinaryPeople
These are sensible and good ordinary people.

- `good_sense` is common sense, stored as a boolean (0 or 1).

#### RichPeople
These are wealthy people with money and land.

- `money` is money.
- `land` is land.

*Note: Units and other details are not considered.*

#### PartyPeople
Party people.

- `free_time` is free time.
- `middle_name` is the middle name.

# Single Table Inheritance
Single Table Inheritance represents object inheritance relationships in a single table. The table includes a column (`type`) to determine the subclass.

![single_table_inheritance_table.png](/assets/images/posts/table-inheritance-types/733b241d-ed09-6e1f-958c-b664f2d4133c.png)

It seems that Rails supports the implementation of STI.

# Class Table Inheritance
Class Table Inheritance represents object inheritance relationships by preparing one table per class. The superclass table contains columns held by the superclass, while the subclass table contains only columns held by the subclass.

![class_table_inheritance.png](/assets/images/posts/table-inheritance-types/33047bc2-d4a3-700c-0995-8738c9897a23.png)

# Concrete Class Inheritance
Concrete Class Inheritance represents object inheritance relationships by preparing tables that correspond only to concrete classes. Each table includes columns held by the superclass as common attributes.

![concrete_table_inheritance.png](/assets/images/posts/table-inheritance-types/bec91e44-0b28-7bcc-6666-026dd5a10f2a.png)

# Thoughts
The choice of which pattern to implement depends on the consideration of the advantages and disadvantages of table design and the cost of application logic. Please point out any misleading or incorrect parts.

# Reference Links
- [Isn't Everyone Misunderstanding Rails' STI!?](http://qiita.com/yebihara/items/9ecb838893ad99be0561#class-table-inheritance--%E3%82%AF%E3%83%A9%E3%82%B9%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB%E7%B6%99%E6%89%BF)
- [Martin Fowler's Bliki(ja) - Patterns of Enterprise Application Architecture](http://bliki-ja.github.io/pofeaa/)