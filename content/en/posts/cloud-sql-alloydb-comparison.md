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
  - Google Cloud
translation_key: cloud-sql-alloydb-comparison
---

I wanted to understand the advantages of AlloyDB compared to Cloud SQL, so I did a quick investigation.

# Premise
This comparison targets the Cloud SQL Enterprise edition.

# Cloud SQL vs AlloyDB
A simple comparison of specifications.

|                        | Cloud SQL                                  | AlloyDB                                                     |
| ---------------------- | ------------------------------------------ | ------------------------------------------------------------ |
| Service Type           | Managed RDBMS                             | Distributed DB close to PostgreSQL                           |
| Use Case               | OLTP                                       | OLTP and OLAP                                               |
| Configuration Unit     | Single Instance                           | Cluster                                                     |
| Scaling                | Vertical and Horizontal (with limitations) | Vertical and Horizontal (with limitations)                   |
| Availability            | SLA 99.99%                                 | SLA 99.95%                                                   |
| Maintenance            | Downtime required                          | No downtime                                                 |
| Backup and Restore     | On-demand backup, automatic backup        | On-demand backup, automatic backup, continuous backup       |
| Compatibility          | MySQL/PostgreSQL/SQL Server                | PostgreSQL                                                   |
| Cost                   | Generally cheaper than AlloyDB in simple comparison | Generally more expensive than Cloud SQL in simple comparison |

# Features of AlloyDB
Summarizing the main features of AlloyDB compared to Cloud SQL.

## Use Case
Cloud SQL is primarily for OLTP, but **AlloyDB supports HTAP**, making it usable for both OLTP and OLAP.

OLAP is enabled by setting the database flag (`google_columnar_engine.enabled`).

There are limitations on the types available for OLAP, but common types are covered.

## Configuration Unit
Cloud SQL is based on a single instance, limiting the performance of computing and storage to the limits of that instance.

On the other hand, AlloyDB has a **configuration that separates computing and storage, providing high scalability**.

## Scaling
Both Cloud SQL and AlloyDB allow vertical scaling through specification upgrades and horizontal scaling through read replicas.

Neither can perform horizontal scaling for writes.

A significant difference is that **AlloyDB reduces replication lag through read replicas** in its configuration.

## Backup and Restore
The difference in backups between Cloud SQL and AlloyDB is that AlloyDB has **continuous backups**.

> Continuous backup and recovery are enabled by default for all clusters, allowing the creation of new clusters based on the latest state of another cluster within the same project and region in AlloyDB.

> With AlloyDB, you can restore an existing cluster to any point in recent history with microsecond precision. By default, AlloyDB allows you to select any point in time up to 14 days in the past. You can configure the cluster to change this window to a maximum of 35 days or a minimum of 1 day.

I didn't quite understand the differences with other backups, but it might be characterized by the ability to specify point-in-time recovery quickly down to the microsecond level.

## Network
Similar to Cloud SQL's Cloud SQL Auth Proxy, AlloyDB provides the AlloyDB Auth Proxy.

This is a proxy software installed locally on the client side to connect to the DB, utilizing IAM-based connection authentication and encrypted communication (TLS).

Using this proxy is not mandatory but is recommended.

## Data Import and Export
AlloyDB supports data import and export in CSV, DMP (PostgreSQL dump files), and SQL formats.

Besides file formats, you can use the Database Migration Service (DMS) to migrate data from other databases to AlloyDB.

There are some constraints, but migration from Cloud SQL to AlloyDB is also possible.

cf. [cloud.google.com - Migrate from Cloud SQL for PostgreSQL to AlloyDB for PostgreSQL](https://cloud.google.com/alloydb/docs/migrate-cloud-sql-to-alloydb#required-roles)

## Tuning
AlloyDB has features like adaptive autovacuum and index advisor.

Adaptive autovacuum improves the operation of PostgreSQL's autovacuum (an autopilot feature, so to speak), automatically executing autovacuum at appropriate times to enhance DB performance and availability.

Index advisor analyzes query execution plans and provides index recommendations.

# Impressions
- It seems reasonable to view it as a superior version of Cloud SQL.
- Reads can scale infinitely. There seems to be no need to worry about replication lag.
- If you are already using Cloud SQL, there may be potential performance improvements with Cloud SQL Enterprise edition Plus, making it a candidate for migration.

# References
- [cloud.google.com - AlloyDB overview](https://cloud.google.com/alloydb/docs/overview)
- [cloud.google.com - How AlloyDB for PostgreSQL Works: Intelligent Database-Compatible Storage](https://cloud.google.com/blog/ja/products/databases/alloydb-for-postgresql-intelligent-scalable-storage)
- [cloud.google.com - How AlloyDB for PostgreSQL Works: Adaptive AutoVacuum](https://cloud.google.com/blog/ja/products/databases/alloydb-for-postgresql-under-the-hood-adaptive-autovacuum)
- [blog.g-gen.co.jp - Thorough Explanation of AlloyDB for PostgreSQL!](https://blog.g-gen.co.jp/entry/alloydb-for-postgresql-explained)
- [Google Cloud - Thorough Comparison of AlloyDB and Cloud SQL!! (Part 1: Overview and Performance Verification)](https://sight-r.sts-inc.co.jp/google_cloud_article/google-cloud-compare-alloydb-1/)
- [cloud-ace.jp - Understanding Google Cloud Database Selection](https://cloud-ace.jp/column/detail469/)
- [medium.com - AlloyDB Adaptive AutoVacuum and how AlloyDB Cluster Storage Space is Released.](https://medium.com/google-cloud/alloydb-adaptive-autovacuum-and-how-alloydb-cluster-storage-space-is-released-41be54b8b8c8)