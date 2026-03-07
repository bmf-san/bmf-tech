---
title: Single Table Inheritance, Class Table Inheritance, and Concrete Class Inheritance
slug: table-inheritance-types
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PofEAA
translation_key: table-inheritance-types
---

# Overview
**Since relational databases do not support inheritance, we need to consider how to represent the inheritance relationships of objects in the database.**
This post explains three patterns for representation: Single Table Inheritance, Class Table Inheritance, and Concrete Class Inheritance.

*Note: This post does not cover the advantages and disadvantages of each implementation pattern.*

# Assumptions
We will assume four classes in this discussion.

![people_class.png](https://qiita-image-store.s3.amazonaws.com/0/124495/7680ef3c-3c9f-6365-8622-364ed30936b3.png)

The structure where Party People inherits from Rich People might be a bit confusing, but as long as the image is conveyed, that’s fine.

#### People
Common attributes shared by all classes.

- id is a unique key.
- name is the name.

#### OrdinaryPeople
Good-natured ordinary people with common sense.

- good_sense is common sense. It is a boolean (0 or 1).

#### RichPeople
Rich people who own money and land.
- money is the amount of money.
- land is the land owned.

*Note: Units and other details are not considered.*

#### PartyPeople
Party people.

- free_time is the free time available.
- middle_name is the middle name.

# Single Table Inheritance
Single Table Inheritance expresses the inheritance relationship of objects in a single table. The table includes a column (type) to determine the subclass.

![single_table_inheritance_table.png](https://qiita-image-store.s3.amazonaws.com/0/124495/733b241d-ed09-6e1f-958c-b664f2d4133c.png)

It seems that Rails supports the implementation of STI.

# Class Table Inheritance
Class Table Inheritance expresses the inheritance relationship of objects by preparing one table for each class. The superclass table contains the columns of the superclass, while the subclass tables contain only the columns of the subclasses.

![class_table_inheritance.png](https://qiita-image-store.s3.amazonaws.com/0/124495/33047bc2-d4a3-700c-0995-8738c9897a23.png)

# Concrete Class Inheritance
Concrete Class Inheritance expresses the inheritance relationship of objects by preparing tables that correspond only to concrete classes. Each table contains the columns of the superclass as common attributes.

![concrete_table_inheritance.png](https://qiita-image-store.s3.amazonaws.com/0/124495/bec91e44-0b28-7bcc-6666-026dd5a10f2a.png)

# Thoughts
The choice of which pattern to implement depends on the advantages and disadvantages of table design and the cost of application-side logic. If there are any misunderstandings or mistakes, please let me know.

# Reference Links
- [Is Everyone Misunderstanding Rails STI!?](http://qiita.com/yebihara/items/9ecb838893ad99be0561#class-table-inheritance--%E3%82%AF%E3%83%A9%E3%82%B9%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB%E7%B6%99%E6%89%BF)
- [Martin Fowler's Bliki(ja) - Patterns of Enterprise Application Architecture](http://bliki-ja.github.io/pofeaa/)