---
title: Differences in Sort Order Due to COLLATE in PostgreSQL and glibc Version Differences
slug: postgresql-collate-glibc-version-differences
date: 2025-03-05T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - PostgreSQL
translation_key: postgresql-collate-glibc-version-differences
---

# Differences in Sort Order Due to COLLATE in PostgreSQL and glibc Version Differences
I encountered an issue where the sort order differed across environments despite specifying the same COLLATE setting, so I am documenting my investigation.

## Incident
### Different Sort Results with Same COLLATE in Cloud SQL and Local PostgreSQL Container

When sorting a string column in a table on Cloud SQL for PostgreSQL 17 using `ORDER BY column_name COLLATE "en_US.utf8"`, I encountered a problem where the order of results differed from that of a PostgreSQL container running locally. Despite using the same version of PostgreSQL and the same COLLATE settings (both at the database level and the COLLATE specified directly in the query), the order was unexpectedly different.

## Investigation
Initially, I suspected that COLLATE was not being applied correctly, but upon checking with EXPLAIN ANALYZE, I confirmed that COLLATE was indeed being applied correctly. I also verified the differences in COLLATE specifications at both the database and query levels, concluding that this was not a COLLATE issue.

### Checking glibc Version
When PostgreSQL performs string comparisons, it utilizes either **glibc** (collprovider = `c`) or **ICU** (collprovider = `i`). If glibc is being used, differences in the glibc version can lead to variations in the actual sort order, even with the same locale name (e.g., `en_US.UTF8`).

You can check which version of glibc is being used by executing the following query within PostgreSQL:

```sql
SELECT oid, collname, collprovider, collversion
FROM pg_collation
WHERE collname = 'en_US.utf8';
```

- On Cloud SQL (PostgreSQL 17), the `collversion` was `2.19`.
- On the local container, the `collversion` was `2.31`.

Thus, even with the same `en_US.UTF8`, Cloud SQL was using **glibc 2.19**, while the local environment was using **glibc 2.31**.

### Revalidating in glibc 2.19 Environment
Many Docker images that are easily available locally tend to have a relatively newer version of glibc. Therefore, I prepared a custom container image of **PostgreSQL 14** using glibc 2.19 (or used [this one](https://hub.docker.com/r/bmfsan/collversion-2.19-postgres-v14)) and executed the same query, resulting in the same sort order as Cloud SQL. This confirmed that the difference between **glibc 2.19** and **glibc 2.31** was indeed the cause of the issue.

## Method to Avoid glibc Dependency: ICU
Since PostgreSQL 10, a feature has been added to perform string collation using ICU (collprovider = `i`). By using ICU, it is highly likely that you can avoid changes in sort order due to glibc version differences.

To use ICU, PostgreSQL must be built with ICU support. Additionally, by specifying an ICU locale like `en-US-x-icu` during `CREATE DATABASE` or column definition, you can achieve sorting that is unaffected by glibc (though this is unverified, it is likely...).

## Summary
- **Differences in glibc versions can change sort order even with the same locale name.**
- If you want to avoid glibc dependency, using **ICU** is a viable option.