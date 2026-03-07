---
title: Summary of PostgreSQL Logical Replication Specifications
slug: postgresql-logical-replication-summary
date: 2025-05-23T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - PostgreSQL
translation_key: postgresql-logical-replication-summary
---

# Overview

Logical replication in PostgreSQL is a mechanism that replicates DML operations (INSERT, UPDATE, DELETE, etc.) on specific tables within a database to other PostgreSQL instances. Unlike physical replication, logical replication allows for flexible selection of targets at the table level, making it suitable for data integration, distributed processing, and migration.

> Reference: [PostgreSQL Logical Replication Documentation - Overview](https://www.postgresql.org/docs/current/logical-replication.html)

# Hands-On
## Environment Setup
```yaml
# docker-compose.yml
version: '3.8'
services:
  publisher:
    container_name: publisher
    image: postgres:17
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: demo
    ports:
      - "5432:5432"
    volumes:
      - publisher_data:/var/lib/postgresql/data
    command: >
      postgres -c wal_level=logical
               -c max_replication_slots=4
               -c max_wal_senders=4
               -c hot_standby=off

  subscriber:
    container_name: subscriber
    image: postgres:17
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: demo
    ports:
      - "5433:5432"
    volumes:
      - subscriber_data:/var/lib/postgresql/data

volumes:
  publisher_data:
  subscriber_data:
```

Start with `docker-compose up -d`.

# Functionality Verification
## Publisher Configuration
Connect to the publisher with `docker exec -it publisher psql -U postgres -d demo`.

Create a table and data with the following SQL:

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT
);

INSERT INTO users (name) VALUES ('Alice'), ('Bob');
```

Create a publication with the following SQL:

```sql
CREATE PUBLICATION my_pub FOR TABLE users;
```

## Subscriber Configuration
Connect to the subscriber with `docker exec -it subscriber psql -U postgres -d demo`.

Create a table with the following SQL:

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT
);
```

Create a subscription with the following SQL:

```sql
CREATE SUBSCRIPTION my_sub
CONNECTION 'host=publisher port=5432 user=postgres password=postgres dbname=demo'
PUBLICATION my_pub;
```

## Verify Logical Replication
Insert data on the publisher side:

```sql
INSERT INTO users (name) VALUES ('Charlie');
```

Check the data on the subscriber side:

```sql
SELECT * FROM users;
```

Confirm that the data inserted on the publisher side is reflected on the subscriber side.

# Architecture: Publisher-Subscriber Model

Logical replication consists of two roles: **Publisher** and **Subscriber**. The publisher publishes changes to specified tables as a "publication". The subscriber subscribes to these publications and reflects the changes in its own database.

Replication is based on PostgreSQL's WAL (Write-Ahead Log), with logically decoded changes sent in real-time. The subscriber receives changes through the `apply` process, maintaining consistency by reflecting them in the same transaction unit.

> Reference: [Architecture - Publisher and Subscriber](https://www.postgresql.org/docs/current/logical-replication.html#LOGICAL-REPLICATION-PUBLISHER)

# Supported Features and Limitations

Logical replication is limited to regular tables; views, materialized views, sequences, external tables, and large objects (LOBs) are excluded. To correctly replicate UPDATEs and DELETEs, a replica ID (usually the primary key) is required. Using `REPLICA IDENTITY FULL` allows support for tables without primary keys, but performance degradation should be noted.

DDL (changes to table definitions) is not replicated, so identical schema management is required on both the publisher and subscriber.

> Reference: [Restrictions](https://www.postgresql.org/docs/current/logical-replication-restrictions.html)

# Publication Configuration

Publications are created by specifying the target tables and the operations (INSERT, UPDATE, DELETE, TRUNCATE) to be included. Bulk publication of all tables within a specific schema is also possible. Starting from PostgreSQL 15, filtering by **row filters** (WHERE clause) and **column lists** for each table is supported, allowing for finer control.

> Reference: [CREATE PUBLICATION](https://www.postgresql.org/docs/current/sql-createpublication.html)

# Subscription Configuration and Synchronization

On the subscriber side, subscription is initiated with `CREATE SUBSCRIPTION`. By specifying connection information and the target publication name, initial synchronization (snapshot copy) and subsequent continuous streaming occur. Initial synchronization is executed in parallel, allowing for relatively quick synchronization even with large amounts of data.

Replication slots are created on the publisher side, serving the role of retaining and sending change logs. On the subscriber side, the `apply` worker process sequentially reflects this data.

> Reference: [CREATE SUBSCRIPTION](https://www.postgresql.org/docs/current/sql-createsubscription.html)

# Transactions and Consistency

Logical replication duplicates at the transaction level. This means that even changes spanning multiple tables are applied to the subscriber at the commit unit on the publisher, maintaining data consistency and ensuring a consistent state on the subscriber side.

However, if a configuration is adopted where the same table is updated simultaneously by multiple subscriptions, conflicts may arise. Starting from PostgreSQL 15, error handling control is possible using `ALTER SUBSCRIPTION ... SKIP`.

> Reference: [Replication Conflicts](https://www.postgresql.org/docs/current/logical-replication-conflicts.html)

# Row and Column Level Filtering (PostgreSQL 15 and later)

To achieve more flexible replication, PostgreSQL 15 allows for specifying **row level filters (WHERE clause)** and **column lists**. This enables, for example, transferring only the necessary rows to different subscribers for each tenant in a multi-tenant environment.

However, row filters cannot use non-immutable functions or expressions with side effects. Additionally, when performing UPDATE/DELETE, only columns used for the replica ID can be included in the conditions.

> Reference: [Row and Column Filtering](https://www.postgresql.org/docs/current/logical-replication-row-filter.html)

# Main Use Cases

Logical replication is suitable for the following use cases:

- **Migration**: Data can be migrated without impacting the running system during version upgrades or transitions to different platforms.
- **Data Integration/Aggregation**: Collect specific tables from multiple databases for analysis in BI tools or data warehouses.
- **Multi-Tenant Isolation**: Distribute data from multiple tenants existing in one DB to separate DBs for each tenant.
- **Event-Driven Architecture**: Build processes that interact with external systems triggered by changes in the database.

> Reference: [Logical Replication Use Cases](https://www.postgresql.org/docs/current/logical-replication-use-cases.html)

# Major Feature Additions by Version

- PostgreSQL 10: Introduction of logical replication.
- PostgreSQL 13: Support for partitioned tables.
- PostgreSQL 14: Streaming application and binary transfer.
- PostgreSQL 15: Row/column filters, schema-level publication, `SKIP`, etc.
- PostgreSQL 16: Replication from standby, parallel application, prevention of loops with origin specification.

> Reference: [Release Notes](https://www.postgresql.org/docs/release/)

# Considerations for Implementation

When implementing logical replication, the following points should be noted:

- Enable `wal_level = logical`.
- Configure appropriate resource settings such as `max_replication_slots`, `max_wal_senders`, and `max_logical_replication_workers`.
- DDL changes are not replicated, so manual synchronization is required.
- Subscribers default to `session_replication_role = replica`, so statement triggers will not fire.

> Reference: [Logical Replication Setup](https://www.postgresql.org/docs/current/logical-replication-setup.html)

Thus, PostgreSQL's logical replication offers high flexibility, but requires design and configuration based on specifications. While it effectively meets a wide range of needs such as migration, data aggregation, and multi-tenant support, attention should be paid to limitations and version differences for proper operation.