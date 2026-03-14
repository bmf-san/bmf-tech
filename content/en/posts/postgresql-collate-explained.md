---
title: About COLLATE in PostgreSQL
slug: postgresql-collate-explained
date: 2025-03-05T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - PostgreSQL
description: 'Learn how COLLATE works in PostgreSQL for string sorting and comparison. Covers database-level and column-level settings, checking collation with pg_database, and tips for Japanese locale.'
translation_key: postgresql-collate-explained
---



A summary of what I learned about COLLATE in PostgreSQL.

## 1. What is COLLATE?

COLLATE is a mechanism to specify the order and method of string comparison (handling of uppercase/lowercase, accents, diacritics, etc.). It affects the order of `ORDER BY` and the results of comparison operators, so it is recommended to set the correct COLLATE in application development in a Japanese environment.

### Examples of COLLATE Specification

- **Specify when creating a database**
  ```sql
  CREATE DATABASE dbname
    LC_COLLATE='ja_JP.UTF-8'
    LC_CTYPE='ja_JP.UTF-8'
    TEMPLATE=template0;
  ```

- **Specify when creating tables or columns**
  ```sql
  CREATE TABLE example (
    name TEXT COLLATE "ja_JP.UTF-8"
  );
  ```

- **Temporarily specify in a query**
  ```sql
  SELECT name
    FROM example
   ORDER BY name COLLATE "en_US.UTF-8";
  ```

## 2. Queries to Check COLLATE Settings

In PostgreSQL, you can check where and how COLLATE is used through queries. Here are the main methods.

### 2-1. Database-Level COLLATE Settings

To check which collation (`LC_COLLATE`) and character type (`LC_CTYPE`) are set for the currently connected database (`current_database()`), execute the following query.

```sql
SELECT datname,
       datcollate AS collate,
       datctype   AS ctype
  FROM pg_database
 WHERE datname = current_database();
```

### 2-2. COLLATE Settings in Tables and Columns

To find out which tables and columns have COLLATE, refer to the `information_schema.columns` table. The following query retrieves only columns with `collation_name` set.

```sql
SELECT table_schema,
       table_name,
       column_name,
       collation_name
  FROM information_schema.columns
 WHERE collation_name IS NOT NULL
   AND table_schema NOT IN ('information_schema','pg_catalog')
 ORDER BY table_schema, table_name, column_name;
```

### 2-3. Checking INDEX with COLLATE

To check indexes with COLLATE specified, refer to `pg_index` and `pg_class`, and search the index definition (`pg_get_indexdef`). The following query checks if `COLLATE` is included in the index definition.

```sql
SELECT idx.relname       AS index_name,
       tbl.relname       AS table_name,
       pg_get_indexdef(idx.oid) AS index_definition
  FROM pg_index i
  JOIN pg_class idx ON idx.oid = i.indexrelid
  JOIN pg_class tbl ON tbl.oid = i.indrelid
 WHERE idx.relkind = 'i'
   AND pg_get_indexdef(idx.oid) ILIKE '%COLLATE%'
 ORDER BY table_name, index_name;
```

If there is no relevant data, the result will be empty.

### 2-4. Checking COLLATE Usage in Functions

To check if COLLATE is included in the definition of user-defined functions (`prokind = 'f'`), use the following query.

```sql
SELECT proname,
       pg_get_functiondef(oid) AS function_definition
  FROM pg_proc
 WHERE prokind = 'f'
   AND pg_get_functiondef(oid) ILIKE '%COLLATE%';
```

Similarly, if the result is empty, it means there are no functions containing COLLATE.

### 2-5. Checking COLLATE Usage in Triggers

To check if COLLATE is included in trigger definitions, use the following query. It refers to `pg_trigger` and sets conditions to exclude internal triggers (`tgisinternal`).

```sql
SELECT tgname,
       pg_get_triggerdef(oid) AS trigger_definition
  FROM pg_trigger
 WHERE NOT tgisinternal
   AND pg_get_triggerdef(oid) ILIKE '%COLLATE%';
```

If this is also empty, it indicates there are no triggers containing COLLATE.

## 3. Summary

When handling Japanese in PostgreSQL, it is important to set the appropriate collation (COLLATE). If left as the default for English-speaking regions, the handling of uppercase/lowercase and the ordering of diacritics may not be as expected. Therefore, it is recommended to:

- Specify `LC_COLLATE='ja_JP.UTF-8'` when creating a database
- Set COLLATE for tables and columns, and specify it per query as needed
- Check if COLLATE is used in indexes, functions, and triggers as appropriate

If different language settings are required, it is also possible to dynamically specify English COLLATE or other language COLLATE to switch sorting order and comparison behavior. During development and operation, it is important to actually store Japanese data and test sorting and comparison results to ensure they behave as intended.

# References
- [PostgreSQL Official Documentation - Collation Support](https://www.postgresql.org/docs/current/collation.html)