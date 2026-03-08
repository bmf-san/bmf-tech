---
title: What is an Index?
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
A mechanism for quickly retrieving records stored in a table.

Consider a query with an O(n) problem:

`SELECT * FROM users WHERE first_name = 'Tom'`

To improve the performance of this query, you can create an index:

`ALTER TABLE users ADD INDEX (first_name)`

# Advantages and Disadvantages
## Advantages
- Improved speed of data reading and retrieval

## Disadvantages
- Increased storage
- Decreased write speed

When creating or updating data, the index is also added or updated simultaneously, leading to the disadvantages mentioned above.

# Patterns of Indexes
## Single Column Index

`ALTER TABLE users ADD INDEX (first_name)`

## Partial Index
A useful pattern when you want to improve performance while minimizing storage increase.

Example of indexing only the first 4 bytes:
`ALTER TABLE users ADD INDEX (first_name(4))`

## Multi-Column Index (also known as Composite or Compound Index)

`ALTER TABLE users ADD INDEX (last_name, first_name)`

In MySQL, **only one index can be used per table for a single query**, but if a multi-index is applied, the valid index will be selected during query execution.

It is generally good practice to specify a column with high cardinality at the beginning of a multi-column index.

## Unique Index
Prevents duplicate values from appearing, excluding NULLs. When creating or updating records, it checks all values to ensure that the same value does not already exist. In MySQL, specifying a unique key also specifies a unique index.

`ALTER TABLE users ADD UNIQUE (first_name)`

# Measuring Index Effectiveness
Check the execution plan of the query using the EXPLAIN clause.

`EXPLAIN SELECT * FROM users WHERE first_name = 'Tom'`

Check the following items:
- possible_keys
  - Available indexes
- key
  - The index actually selected
- extra
  - If you see messages like the following, it is advisable to optimize the query:
    - using filesort
      - Insufficient memory for sorting, performing physical file sorting
    - using temporary
      - A temporary table is created for query execution

# Criteria for Considering Indexes
Here are some criteria that may indicate the need for an index. These are merely guidelines, and it is better to measure using EXPLAIN.

- When the amount of data in the table is large, and the number of records being searched is small
- When there are columns used in WHERE, JOIN, ORDER BY, etc.
- When searching for non-NULL values in data that contains many NULLs (indexes do not include NULLs, so they may be effective)
- When data is not frequently added, updated, or deleted (considering the update load of the index)

# Clustered Index and Secondary Index
## Clustered Index
The following indexes are considered clustered indexes:

- Columns defined as primary keys
- NOT NULL unique key columns
- If there are no columns that meet the above criteria, InnoDB creates a hidden clustered index called GEN_CLUST_INDEX.

## Secondary Index
Indexes other than clustered indexes are called secondary indexes. Secondary indexes include the values of the primary key. Although measurement with EXPLAIN is a prerequisite, since primary key values are included, it is good to remember that secondary indexes can potentially serve as covering indexes even without including the primary key in composite indexes.
cf. [Useful Techniques for InnoDB Secondary Indexes!](https://nippondanji.blogspot.com/2010/10/innodb.html)

# Covering Index
An index that includes all columns necessary for the query execution result.

Since it can cover the query without reading the data file, it speeds up the search.

# Precautions When Creating Indexes
## Operations on Indexed Columns and SQL Functions
```sql
SELECT * FROM users WHERE amount * 2 > 10;
```

When an index is applied to amount, avoid using operators to utilize the index. The amount itself is stored in the index, not the result of the operation. The same applies to SQL functions.

```sql
SELECT * FROM users WHERE amount > 10/2;
```

## IS_NULL
```sql
SELECT * FROM users WHERE amount IS NULL;
```

IS NULL and IS NOT NULL generally do not effectively utilize indexes (depending on the DBMS specifications).

## Negation and OR
```sql
SELECT * FROM users WHERE amount <> 10;
```

Negation cannot utilize indexes. The same applies to OR.

## LIKE
```sql
SELECT * FROM users WHERE name = 'a%';
```

When using LIKE, only prefix matches utilize the index due to the properties of B-Tree.

## Implicit Type Conversion
```sql
SELECT * FROM users WHERE age = '10'
```

If age is a numeric type, implicit conversion from string to number will prevent the index from being utilized.

# References
- [amzn.to - Practical Introduction to Databases from Theory ~ Efficient SQL with Relational Models (WEB+DB PRESS plus)](https://amzn.to/3TEltzx)
- [www.hi-ho.ne.jp - Basic Knowledge of Indexes](http://www.hi-ho.ne.jp/tsumiki/doc_1.html)
- [kiyotakubo.hatenablog.com - Basic Knowledge of Indexes for MySQL Performance Tuning](http://kiyotakakubo.hatenablog.com/entry/20101117/1289952549)
- [qiita.com - Tips for Creating Indexes in MySQL](https://qiita.com/katsukii/items/3409e3c3c96580d37c2b#%E9%80%9A%E5%B8%B8)
- [dev.mysql.com - EXPLAIN Output Format](https://dev.mysql.com/doc/refman/5.6/ja/explain.html)
- [dev.mysql.com - Covering Index](https://dev.mysql.com/doc/refman/8.0/ja/glossary.html#glos_covering_index)