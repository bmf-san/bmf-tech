---
title: About COLLATE in PostgreSQL
slug: postgresql-collate-explained
date: 2025-03-05T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - PostgreSQL
translation_key: postgresql-collate-explained
---

This post summarizes what I researched about COLLATE in PostgreSQL.

## 1. What is COLLATE?

COLLATE is a mechanism that specifies the order of string sorting and how comparisons are made (handling of uppercase and lowercase letters, accents, voiced sounds, etc.). For example, it affects the sorting order by `ORDER BY` and the results of comparison operators, so it is recommended to set the correct COLLATE when developing applications in a Japanese environment.

### Example of COLLATE Specification

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

## 2. Query to Check COLLATE Settings

In PostgreSQL, you can check where and how COLLATE is used through queries. Below are the main methods.

### 2-1. Database-Level COLLATE Settings

To check which collation (`LC_COLLATE`) and character type (`LC_CTYPE`) are set for the currently connected database (`current_database()`), execute the following query.

```sql
SELECT datname,
       datcollate AS collate,
       datctype   AS ctype
  FROM pg_database
 WHERE datname = current_database();
```

### 2-2. COLLATE Settings for Tables and Columns

To find out which tables and columns have COLLATE, refer to the `information_schema.columns` table. The following query retrieves only the columns that have a `collation_name` set.

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

### 2-3. Checking COLLATE with INDEX

If you want to check indexes that specify COLLATE, you can refer to `pg_index` and `pg_class`, and search for index definitions (`pg_get_indexdef`). The following query checks whether `COLLATE` is included in the index definition.

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

To check whether COLLATE specification is included in the definition of user-defined functions (`prokind = 'f'`), you can use the following query.

```sql
SELECT proname,
       pg_get_functiondef(oid) AS function_definition
  FROM pg_proc
 WHERE prokind = 'f'
   AND pg_get_functiondef(oid) ILIKE '%COLLATE%';
```

Similarly, if the result is empty, it indicates that there are no functions containing COLLATE.

### 2-5. Checking COLLATE Usage in Triggers

To check whether COLLATE is included in trigger definitions, you can use the following query. It references `pg_trigger` and sets conditions to exclude internal triggers (`tgisinternal`).

```sql
SELECT tgname,
       pg_get_triggerdef(oid) AS trigger_definition
  FROM pg_trigger
 WHERE NOT tgisinternal
   AND pg_get_triggerdef(oid) ILIKE '%COLLATE%';
```

Again, if the result is empty, it indicates that there are no triggers containing COLLATE.

## 3. Conclusion

When handling Japanese in PostgreSQL, it is important to set the correct collation (COLLATE). If left at the default for English-speaking regions, the handling of uppercase and lowercase letters, as well as the ordering of voiced and semi-voiced sounds, may not behave as expected. Therefore,

- Specify `LC_COLLATE='ja_JP.UTF-8'` when creating the database.
- Set COLLATE for tables and columns, and specify it in queries as needed.
- If COLLATE is used in indexes, functions, or triggers, check them as appropriate.

It is also possible to dynamically specify English COLLATE or COLLATE for other languages if different language settings are required to switch sorting order or comparison behavior. During development and operation, it is essential to actually store Japanese data and test sorting and comparison results to ensure they behave as intended.

# References
- [PostgreSQL Official Documentation - Collation Support](https://www.postgresql.org/docs/current/collation.html)