---
title: SQL for Generating Test Data in MySQL
slug: mysql-test-data-generation-sql
date: 2019-07-16T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - MySQL
  - SQL
  - Cross Join
description: A note on generating test data using only MySQL.
translation_key: mysql-test-data-generation-sql
---



## Overview
This is a note on generating test data using only MySQL. While generating test data through scripts offers high flexibility and seems like a superior method, using only SQL might be sufficient when you want to perform performance tests with tens of thousands of records.

# SQL
The query looks like this.

```sql
DROP TABLE IF EXISTS `tests`;

CREATE TABLE `tests` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `value` int(5) NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);

INSERT INTO tests(value)
VALUES (1), (2), (3), (4), (5), (6), (7), (8), (9), (10);

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `title` varchar(255) DEFAULT NULL,
  `body` text DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  FOREIGN KEY (admin_id) REFERENCES admins(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO posts(title, body, created_at, updated_at)
SELECT
  (@rownum := @rownum + 1),
  @rownum,
  CONCAT(@rownum, 'title'),
  CONCAT(@rownum, 'md_body'),
  CONCAT(@rownum, 'html_body')
FROM
  tests AS t1,
  tests AS t2,
  tests AS t3,
  tests AS t4,
  (SELECT @rownum := 0) AS v;
```

This method uses user-defined variables to take row numbers and generates records using a Cartesian product (CROSS JOIN) and `INSERT INTO ... SELECT`. There were various patterns, but this one seems relatively easy to understand, or rather, easy to write. The downside is that it's difficult to immediately visualize what it's doing.

# References
- [Creating Large Test Data with SQL](https://qiita.com/cobot00/items/8d59e0734314a88d74c7)
