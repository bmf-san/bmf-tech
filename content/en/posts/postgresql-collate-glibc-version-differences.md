---
title: Differences in Sorting Order Due to COLLATE and glibc Version Differences in PostgreSQL
slug: postgresql-collate-glibc-version-differences
date: 2025-03-05T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - Database
tags:
  - PostgreSQL
description: Encountered an issue where sorting order differs despite specifying the same COLLATE setting, and documented the investigation.
translation_key: postgresql-collate-glibc-version-differences
---

# Differences in Sorting Order Due to COLLATE and glibc Version Differences in PostgreSQL
Encountered an issue where sorting order differs despite specifying the same COLLATE setting, and documented the investigation.

## Phenomenon
### Different Sorting Results with the Same COLLATE in Cloud SQL and Local PostgreSQL Container

While sorting a string column in a table on Cloud SQL for PostgreSQL 17 using `ORDER BY column_name COLLATE "en_US.utf8"`, I encountered an issue where the order differed from the results of a locally running PostgreSQL container. Despite having the same version of PostgreSQL and the same COLLATE setting (both at the database level and directly specified in the query), the order was different.

## Investigation
Initially suspected that COLLATE was not being applied, but confirmed through `EXPLAIN ANALYZE` that COLLATE was correctly applied. Further verification of differences in COLLATE specification at the database and query levels showed that it was not a COLLATE issue.

### Checking glibc Version
When PostgreSQL performs string comparisons, it uses either **glibc** (collprovider = `c`) or **ICU** (collprovider = `i`). If glibc is used, differences in glibc versions can result in different sorting orders even with the same locale name (such as `en_US.UTF8`).

To check which version of glibc is used, execute the following query within PostgreSQL:

```sql
SELECT oid, collname, collprovider, collversion
FROM pg_collation
WHERE collname = 'en_US.utf8';
```

- On Cloud SQL (PostgreSQL 17), `collversion` was `2.19`
- On the local container, `collversion` was `2.31`

Thus, even with the same `en_US.UTF8`, Cloud SQL used **glibc 2.19** and the local environment used **glibc 2.31**.

### Re-evaluation in glibc 2.19 Environment
Many Docker images readily available for local use have relatively new versions of glibc. Therefore, I prepared a custom PostgreSQL 14 container image using glibc 2.19 (or used [this one](https://hub.docker.com/r/bmfsan/collversion-2.19-postgres-v14)) and executed the same query, resulting in the same sorting order as Cloud SQL. This confirmed that the difference between **glibc 2.19** and **glibc 2.31** was the cause of the issue.

## Avoiding glibc Dependency: ICU
Since PostgreSQL 10, there is a feature to perform string collation using ICU (collprovider = `i`). Using ICU can likely avoid sorting order changes due to glibc version differences.

To use ICU, PostgreSQL must be built with ICU support. Additionally, specifying an ICU locale like `en-US-x-icu` during `CREATE DATABASE` or column definition can achieve sorting unaffected by glibc. (Not verified, but probably...)

## Summary
- **Sorting order can change with the same locale name due to differences in glibc versions**
- To avoid glibc dependency, there is an option to use **ICU**

## Related Posts

- [About COLLATE in PostgreSQL](/posts/postgresql-collate-explained/)
- [PostgreSQL Memory Configuration](/posts/postgresql-memory-settings/)
- [About Row Level Security (RLS) in PostgreSQL](/posts/postgresql-row-level-security/)
- [Learning PostgreSQL from the Inside: Design and Operation Principles](/posts/postgresql-internal-structure/)