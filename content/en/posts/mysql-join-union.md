---
title: About MySQL JOIN and UNION
slug: mysql-join-union
date: 2018-07-18T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - MySQL
  - join
  - union
translation_key: mysql-join-union
---

# Overview
Summarizing the types of JOIN in MySQL and UNION.

# INNER JOIN
Combines records where the values of specified columns match. If the values of the specified columns do not match, they are not combined. (Only data that matches in both tables is combined.)

users
+------+--------+------+
| id   | sex    | name |
+------+--------+------+
|    0 | male   | John |
|    1 | female | Risa |
|    2 | male   | Taro |
+------+--------+------+

accounts
+------+---------+---------------------+
| id   | user_id | created_at          |
+------+---------+---------------------+
|    0 |       0 | 2018-07-18 14:47:41 |
|    1 |       1 | 2018-07-18 14:48:01 |
|    3 |       3 | 2018-07-18 15:07:37 |
+------+---------+---------------------+

`SELECT * FROM users INNER JOIN accounts ON users.id = accounts.user_id`
+------+--------+------+------+---------+---------------------+
| id   | sex    | name | id   | user_id | created_at          |
+------+--------+------+------+---------+---------------------+
|    0 | male   | John |    0 |       0 | 2018-07-18 14:47:41 |
|    1 | female | Risa |    1 |       1 | 2018-07-18 14:48:01 |
+------+--------+------+------+---------+---------------------+

The record with id 2 in the users table does not have a matching entry in the accounts table, so it is not included in the join. The record with user_id 3 in the accounts table does not have a matching entry in the users table, so it is not included in the join.

# LEFT OUTER JOIN
Combines records where the values of specified columns match. Values that exist in the left table but not in the right table are padded with NULL. (All records that exist in the left table are combined.)

users
+------+--------+------+
| id   | sex    | name |
+------+--------+------+
|    0 | male   | John |
|    1 | female | Risa |
|    2 | male   | Taro |
+------+--------+------+

accounts
+------+---------+---------------------+
| id   | user_id | created_at          |
+------+---------+---------------------+
|    0 |       0 | 2018-07-18 14:47:41 |
|    1 |       1 | 2018-07-18 14:48:01 |
|    3 |       3 | 2018-07-18 15:07:37 |
+------+---------+---------------------+

`SELECT * FROM users LEFT OUTER JOIN accounts ON users.id = accounts.id`
+------+--------+------+------+---------+---------------------+
| id   | sex    | name | id   | user_id | created_at          |
+------+--------+------+------+---------+---------------------+
|    0 | male   | John |    0 |       0 | 2018-07-18 14:47:41 |
|    1 | female | Risa |    1 |       1 | 2018-07-18 14:48:01 |
|    2 | male   | Taro | NULL |    NULL | NULL                |
+------+--------+------+------+---------+---------------------+

In the left table, users, there is a record with id 2, but there is no matching entry in the right table, accounts, so it is padded with NULL and combined.

# RIGHT OUTER JOIN
The reverse of LEFT OUTER JOIN. Combines records where the values of specified columns match. Values that exist in the right table but not in the left table are padded with NULL. (All records that exist in the right table are combined.)

users
+------+--------+------+
| id   | sex    | name |
+------+--------+------+
|    0 | male   | John |
|    1 | female | Risa |
|    2 | male   | Taro |
+------+--------+------+

accounts
+------+---------+---------------------+
| id   | user_id | created_at          |
+------+---------+---------------------+
|    0 |       0 | 2018-07-18 14:47:41 |
|    1 |       1 | 2018-07-18 14:48:01 |
|    3 |       3 | 2018-07-18 15:07:37 |
+------+---------+---------------------+

`SELECT * from users RIGHT OUTER JOIN accounts ON users.id = accounts.id`
+------+--------+------+------+---------+---------------------+
| id   | sex    | name | id   | user_id | created_at          |
+------+--------+------+------+---------+---------------------+
|    0 | male   | John |    0 |       0 | 2018-07-18 14:47:41 |
|    1 | female | Risa |    1 |       1 | 2018-07-18 14:48:01 |
| NULL | NULL   | NULL |    3 |       3 | 2018-07-18 15:07:37 |
+------+--------+------+------+---------+---------------------+

In the right table, accounts, there is a record with user_id 3, but there is no matching entry in the left table, users, so it is padded with NULL and combined.

# CROSS JOIN
In MySQL, CROSS JOIN and INNER JOIN are syntactically equivalent. (Reference: [MySQL 8.2.1.11 Nested Join Optimization](https://dev.mysql.com/doc/refman/5.6/ja/nested-join-optimization.html))

# UNION
Combines tables while eliminating duplicates. The condition is that the number of columns must be the same in both tables.

users
+------+--------+------+
| id   | sex    | name |
+------+--------+------+
|    0 | male   | John |
|    1 | female | Risa |
|    2 | male   | Taro |
+------+--------+------+

accounts
+------+---------+---------------------+
| id   | user_id | created_at          |
+------+---------+---------------------+
|    0 |       0 | 2018-07-18 14:47:41 |
|    1 |       1 | 2018-07-18 14:48:01 |
|    3 |       3 | 2018-07-18 15:07:37 |
+------+---------+---------------------+

`SELECT * FROM users UNION SELECT * FROM accounts`
+------+--------+---------------------+
| id   | sex    | name                |
+------+--------+---------------------+
|    0 | male   | John                |
|    1 | female | Risa                |
|    2 | male   | Taro                |
|    0 | 0      | 2018-07-18 14:47:41 |
|    1 | 1      | 2018-07-18 14:48:01 |
|    3 | 3      | 2018-07-18 15:07:37 |
+------+--------+---------------------+

# References
- [Basics of SQL – Differences in JOIN (Table Join)](https://www.ecoop.net/memo/archives/2007-11-14-1.html)
- [MySQL](https://dev.mysql.com/doc/refman/5.6/ja/join.html)