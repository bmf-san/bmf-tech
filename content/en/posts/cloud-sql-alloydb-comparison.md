---
title: Comparison of Cloud SQL and AlloyDB
slug: cloud-sql-alloydb-comparison
date: 2024-11-17T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - AlloyDB
  - Cloud SQL
  - Google Cloud Platform
description: A brief investigation into the advantages of AlloyDB over Cloud SQL.
translation_key: cloud-sql-alloydb-comparison
---


I wanted to learn about the advantages of AlloyDB compared to Cloud SQL, so I conducted a brief investigation.

# Prerequisites
This comparison targets the Cloud SQL Enterprise edition.

# Cloud SQL vs AlloyDB
A simple comparison of specifications.

|                        | Cloud SQL                                  | AlloyDB                                                     |
| ---------------------- | ------------------------------------------ | ------------------------------------------------------------ |
| Service Type           | Managed RDBMS                              | PostgreSQL-based, close to a distributed DB                  |
| Use Case               | OLTP                                       | OLTP and OLAP                                                |
| Configuration Unit     | Single Instance                            | Cluster                                                     |
| Scaling                | Vertical and Horizontal (with limitations) | Vertical and Horizontal (with limitations)                   |
| Availability           | SLA 99.99%                                 | SLA 99.95%                                                   |
| Maintenance            | With Downtime                              | No Downtime                                                  |
| Backup and Restore     | On-demand and Automatic Backup             | On-demand, Automatic, and Continuous Backup                  |
| Compatibility          | MySQL/PostgreSQL/SQL Server                | PostgreSQL                                                   |
| Cost                   | Generally cheaper than AlloyDB             | Generally more expensive than Cloud SQL                      |

# Features of AlloyDB
Summarizing the main features of AlloyDB in comparison to Cloud SQL.

## Use Case
Cloud SQL is primarily for OLTP, but **AlloyDB supports HTAP**, allowing it to be used for both OLTP and OLAP.

OLAP can be enabled by setting the database flag (`google_columnar_engine.enabled`).

There are restrictions on the types available for OLAP, but common types are covered.

## Configuration Unit
Cloud SQL is based on a single instance, limiting computing and storage performance to the limits of a single instance.

On the other hand, AlloyDB has a **configuration that separates computing and storage, providing high scalability**.

## Scaling
Both Cloud SQL and AlloyDB allow vertical scaling through spec upgrades and horizontal scaling through read replicas.

Neither supports horizontal scaling for writes.

A major difference is that **AlloyDB has a mechanism to reduce replication lag with read replicas**.

## Backup and Restore
The difference in backup between Cloud SQL and AlloyDB is that AlloyDB offers **continuous backup**.

> Continuous backup and recovery are enabled by default for all clusters, allowing you to create a new cluster based on the latest state of another cluster within the same project and region.

> With AlloyDB, you can restore an existing cluster to any point in recent history with microsecond precision. By default, AlloyDB allows you to select any point in time up to 14 days ago. You can configure the cluster to change this window to a maximum of 35 days or a minimum of 1 day.

I couldn't fully understand the differences from other backups, but it seems that the ability to specify a point-in-time recovery down to microseconds might be a key feature.

## Network
Like Cloud SQL's Cloud SQL Auth Proxy, AlloyDB Auth Proxy is provided.

This is proxy software installed locally on the client side connecting to the DB, allowing the use of IAM-based connection authentication and encrypted communication (TLS).

Using this proxy is not mandatory but recommended.

## Data Import and Export
AlloyDB supports data import and export via CSV, DMP (PostgreSQL dump files), and SQL.

Beyond file formats, using the Database Migration Service (DMS) allows data migration from other databases to AlloyDB.

Migration from Cloud SQL to AlloyDB is possible with some constraints.

cf. [cloud.google.com - Migrate from Cloud SQL for PostgreSQL to AlloyDB for PostgreSQL](https://cloud.google.com/alloydb/docs/migrate-cloud-sql-to-alloydb#required-roles)

## Tuning
AlloyDB offers features like adaptive autovacuum and index advisor.

Adaptive autovacuum improves the operation of PostgreSQL's autovacuum (an autopilot feature, akin to managed autovacuum), automatically executing autovacuum at appropriate times to improve DB performance and availability.

The index advisor analyzes query execution plans and suggests indexes.

# Impressions
- It seems reasonable to view AlloyDB as a superior version of Cloud SQL.
- Reads scale infinitely. There seems to be no need to worry about replication lag.
- If you're already using Cloud SQL, considering migration to Cloud SQL Enterprise edition Plus might be worthwhile as it could offer performance improvements, making it a potential candidate for migration.

# References
- [cloud.google.com - AlloyDB overview](https://cloud.google.com/alloydb/docs/overview)
- [cloud.google.com - How AlloyDB for PostgreSQL works: Intelligent database storage](https://cloud.google.com/blog/ja/products/databases/alloydb-for-postgresql-intelligent-scalable-storage)
- [cloud.google.com - How AlloyDB for PostgreSQL works: Adaptive AutoVacuum](https://cloud.google.com/blog/ja/products/databases/alloydb-for-postgresql-under-the-hood-adaptive-autovacuum)
- [blog.g-gen.co.jp - Thorough explanation of AlloyDB for PostgreSQL!](https://blog.g-gen.co.jp/entry/alloydb-for-postgresql-explained)
- [【Google Cloud】Thorough comparison of AlloyDB and Cloud SQL!! (Part 1: Overview and performance verification of AlloyDB)](https://sight-r.sts-inc.co.jp/google_cloud_article/google-cloud-compare-alloydb-1/)
- [cloud-ace.jp - Understand Google Cloud database selection with this!](https://cloud-ace.jp/column/detail469/)
- [medium.com - AlloyDB Adaptive AutoVacuum and how AlloyDB Cluster Storage Space is Released.](https://medium.com/google-cloud/alloydb-adaptive-autovacuum-and-how-alloydb-cluster-storage-space-is-released-41be54b8b8c8)
