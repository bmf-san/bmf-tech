---
title: What is a Database Index? How It Works and When You Need It
slug: what-is-index
date: 2024-04-01T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - DB
  - Index
  - MySQL
translation_key: what-is-index
---



# What is an Index?
A mechanism to quickly retrieve records stored in a table.

Consider a query with an O(n) problem like the following:

`SELECT * FROM users WHERE first_name = ‘Tom’`

To improve the performance of this query, you can add an index as follows:

`ALTER TABLE users ADD INDEX (first_name)`

# Advantages and Disadvantages
## Advantages
- Improved speed of data reading and retrieval

## Disadvantages
- Increased storage size
- Decreased write speed

When creating or updating data, indexes are also added or updated simultaneously, leading to the above disadvantages.

# Index Patterns
## Standard (Applied to a single column)

`ALTER TABLE users ADD INDEX (first_name)`

## Partial Index
An effective pattern when you want to improve performance while suppressing storage increase.

Example of applying an index to only the first 4 bytes:
`ALTER TABLE users ADD INDEX (first_name(4))`

## Multi-column Index (Also called composite or compound index)

`ALTER TABLE users ADD INDEX (last_name, first_name)`

In MySQL, **only one index per table can be used for a single query execution**, but if a multi-index is applied, a valid index will be selected during query execution.

It is generally good to specify a column with high cardinality at the beginning of a multi-column index.

## Unique Index
Values will not appear duplicated except for NULL.
During record creation or update, it checks all values to ensure the same value does not already exist.
In MySQL, specifying a unique key also specifies a unique index.

`ALTER TABLE users ADD UNIQUE (first_name)`

# Measuring Index Effectiveness
Check the execution plan of a query with the EXPLAIN clause.

`EXPLAIN SELECT * FROM users WHERE first_name = ‘Tom’`

Check the following items:
- possible_keys
  - Possible indexes
- key
  - Index actually selected
- extra
  - If the following displays appear, query optimization is recommended:
  - using filesort
    - Insufficient memory for sorting, writing out to physical files for sorting
  - using temporary
    - Temporary tables are created for query execution

# Index Criteria
List of criteria that might suggest considering an index. These are just guidelines for estimation, so it's better to measure with EXPLAIN.

- When the table has a large amount of data and the target records for search are few
- When there are columns used in WHERE, JOIN, ORDER BY, etc.
- When searching for non-NULL data from data containing many NULLs (Index may be effective as it does not include NULLs)
- When data is not frequently added, updated, or deleted (considering the load of index updates)

# Clustered Index and Secondary Index
## Clustered Index
Indexes that fall under the following are clustered indexes.

- Columns defined by the primary key
- Columns with NOT NULL unique keys
- If no columns fall under the above, InnoDB creates a hidden clustered index called GEN_CLUST_INDEX

## Secondary Index
Indexes other than clustered indexes are called secondary indexes.
Secondary indexes include the value of the primary key.
Although it is assumed to be measured with EXPLAIN, since the primary key value is included, it might be good to remember that a covering index can be achieved with just a secondary index without including the primary key in a composite index.
cf. [InnoDB Secondary Index Utilization Techniques!](https://nippondanji.blogspot.com/2010/10/innodb.html)

# Covering Index
An index that contains all the columns necessary for the query result.

Since it can be covered with just the index without reading the data file, the search is accelerated.

# Cautions When Applying Indexes
## Arithmetic Operations and SQL Functions on Index Columns
```sql
SELECT * FROM users WHERE amount * 2 > 10;
```

If an index is applied to amount, avoid operators to utilize the index. The amount itself is held in the index, not the result of the operation. The same applies to SQL functions.

```sql
SELECT * FROM users WHERE amount > 10/2;
```

## IS_NULL
```sql
SELECT * FROM users WHERE amount IS NULL;
```

IS NULL or IS NOT NULL generally do not effectively utilize indexes (depends on DBMS specifications).

## Negation or OR
```sql
SELECT * FROM users WHERE amount <> 10;
```

Negation cannot utilize indexes. The same applies to OR.

## LIKE
```sql
SELECT * FROM users WHERE name = 'a%';
```

When using LIKE, only forward matches utilize indexes due to the nature of B-Tree.

## Implicit Type Conversion
```sql
SELECT * FROM users WHERE age = '10'
```

If age is a numeric type, implicit conversion from string to number will prevent index utilization.

# References
- [amzn.to - Learn Database Practice from Theory ~ Efficient SQL with Relational Model (WEB+DB PRESS plus)](https://amzn.to/3TEltzx)
- [www.hi-ho.ne.jp - Basic Knowledge of Indexes](http://www.hi-ho.ne.jp/tsumiki/doc_1.html)
- [kiyotakubo.hatenablog.com - Basic Knowledge of Indexes for MySQL Performance Tuning](http://kiyotakakubo.hatenablog.com/entry/20101117/1289952549)
- [qiita.com - Tips for Applying MySQL Indexes](https://qiita.com/katsukii/items/3409e3c3c96580d37c2b#%E9%80%9A%E5%B8%B8)
- [dev.mysql.com - EXPLAIN Output Format](https://dev.mysql.com/doc/refman/5.6/ja/explain.html)
- [dev.mysql.com - Covering Index](https://dev.mysql.com/doc/refman/8.0/ja/glossary.html#glos_covering_index)
