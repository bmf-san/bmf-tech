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

Logical replication in PostgreSQL is a mechanism that replicates DML operations (such as INSERT, UPDATE, DELETE) on specific tables within a database to other PostgreSQL instances. Unlike physical replication, logical replication allows for flexible selection of target tables, making it suitable for data integration, distributed processing, and migration.

> Reference: [PostgreSQL Logical Replication Documentation - Overview](https://www.postgresql.org/docs/current/logical-replication.html)

# Hands-on
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

# Verification
## Publisher Configuration
Connect to the publisher with `docker exec -it publisher psql -U postgres -d demo`.

Create tables and data with the following SQL:

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
Insert data on the publisher side.

```sql
INSERT INTO users (name) VALUES ('Charlie');
```

Check the data on the subscriber side.

```sql
SELECT * FROM users;
```

Verify that the data inserted on the publisher side is reflected on the subscriber side.

# Architecture: Publisher-Subscriber Model

Logical replication consists of two roles: **publisher** and **subscriber**. The publisher publishes changes to specified tables as a "publication." The subscriber subscribes to these publications and reflects the changes in its own database.

Replication is based on PostgreSQL's WAL (Write-Ahead Log), where logically decoded changes are sent in real-time. The subscriber receives changes via the `apply` process, reflecting them in the same transaction unit to maintain consistency.

> Reference: [Architecture - Publisher and Subscriber](https://www.postgresql.org/docs/current/logical-replication.html#LOGICAL-REPLICATION-PUBLISHER)

# Supported Targets and Limitations

Logical replication targets regular tables only; views, materialized views, sequences, foreign tables, and large objects (LOBs) are not supported. To correctly replicate UPDATE or DELETE, a replica ID (usually a primary key) is required. Using `REPLICA IDENTITY FULL` allows for tables without a primary key but may impact performance.

DDL (table definition changes) are not replicated, so schema management must be consistent between publisher and subscriber.

> Reference: [Restrictions](https://www.postgresql.org/docs/current/logical-replication-restrictions.html)

# Publication Configuration

A publication is created by specifying the target tables and operations (INSERT, UPDATE, DELETE, TRUNCATE). Bulk publication of all tables within a specific schema is also possible. From PostgreSQL 15 onwards, table-specific **row filters** (WHERE clause) and **column lists** allow for more granular control.

> Reference: [CREATE PUBLICATION](https://www.postgresql.org/docs/current/sql-createpublication.html)

# Subscription Configuration and Synchronization

On the subscriber side, a subscription is started with `CREATE SUBSCRIPTION`. By specifying connection information and the target publication name, initial synchronization (snapshot copy) and subsequent continuous streaming occur. Initial synchronization is executed in parallel, allowing for relatively quick synchronization even with large data volumes.

A replication slot is created on the publisher side, serving to retain and send change logs. On the subscriber side, the `apply` worker process sequentially reflects this data.

> Reference: [CREATE SUBSCRIPTION](https://www.postgresql.org/docs/current/sql-createsubscription.html)

# Transactions and Consistency

Logical replication replicates on a transaction basis. That is, changes spanning multiple tables are applied to the subscriber in the commit unit on the publisher, maintaining data consistency and ensuring a consistent state on the subscriber side.

However, if the same table is updated simultaneously by multiple subscriptions, conflicts may occur. From PostgreSQL 15 onwards, error handling control is possible with `ALTER SUBSCRIPTION ... SKIP`.

> Reference: [Replication Conflicts](https://www.postgresql.org/docs/current/logical-replication-conflicts.html)

# Row and Column Level Filtering (PostgreSQL 15 and Later)

To achieve more flexible replication, PostgreSQL 15 allows for **row-level filters (WHERE clause)** and **column lists**. This enables, for example, transferring only necessary rows to different subscribers in a multi-tenant environment.

However, row filters cannot use non-immutable functions or expressions with side effects. When performing UPDATE/DELETE, only columns used in the replica ID can be included in the conditions.

> Reference: [Row and Column Filtering](https://www.postgresql.org/docs/current/logical-replication-row-filter.html)

# Main Use Cases

Logical replication is suitable for the following use cases:

- **Migration**: Data can be migrated without affecting the running system during version upgrades or platform transitions.
- **Data Integration and Aggregation**: Collect specific tables from multiple databases for analysis with BI tools or data warehouses.
- **Multi-Tenant Isolation**: Distribute data from multiple tenants in one DB to separate DBs for each tenant.
- **Event-Driven Architecture**: Build processes integrated with external systems using database changes as triggers.

> Reference: [Logical Replication Use Cases](https://www.postgresql.org/docs/current/logical-replication-use-cases.html)

# Major Feature Additions by Version

- PostgreSQL 10: Introduction of logical replication.
- PostgreSQL 13: Support for partitioned tables.
- PostgreSQL 14: Streaming apply and binary transfer.
- PostgreSQL 15: Row and column filters, schema-level publication, `SKIP`, etc.
- PostgreSQL 16: Replication from standby, parallel apply, loop prevention with origin specification.

> Reference: [Release Notes](https://www.postgresql.org/docs/release/)

# Considerations for Implementation

When implementing logical replication, consider the following points:

- Enable `wal_level = logical`.
- Appropriately configure resources such as `max_replication_slots`, `max_wal_senders`, `max_logical_replication_workers`.
- DDL changes are not replicated, so synchronize manually.
- Subscribers default to `session_replication_role = replica`, so statement triggers do not fire.

> Reference: [Logical Replication Setup](https://www.postgresql.org/docs/current/logical-replication-setup.html)

Thus, PostgreSQL's logical replication is a highly flexible feature that requires design and configuration based on specifications. While it effectively meets a wide range of needs such as migration, data aggregation, and multi-tenant support, attention should be paid to limitations and version differences for proper operation.

# References
- [www.postgresql.org - Architecture - Publisher and Subscriber](https://www.postgresql.org/docs/current/logical-replication.html#LOGICAL-REPLICATION-PUBLISHER)
- [www.postgresql.org - PostgreSQL Logical Replication Documentation - Overview](https://www.postgresql.org/docs/current/logical-replication.html)
- [www.postgresql.org - Restrictions](https://www.postgresql.org/docs/current/logical-replication-restrictions.html)
- [www.postgresql.org - CREATE PUBLICATION](https://www.postgresql.org/docs/current/sql-createpublication.html)
- [www.postgresql.org - CREATE SUBSCRIPTION](https://www.postgresql.org/docs/current/sql-createsubscription.html)
- [www.postgresql.org - Replication Conflicts](https://www.postgresql.org/docs/current/logical-replication-conflicts.html)
- [www.postgresql.org - Row and Column Filtering](https://www.postgresql.org/docs/current/logical-replication-row-filter.html)
- [www.postgresql.org - Logical Replication Use Cases](https://www.postgresql.org/docs/current/logical-replication-use-cases.html)
- [www.postgresql.org - Release Notes](https://www.postgresql.org/docs/release/)
- [www.postgresql.org - Logical Replication Setup](https://www.postgresql.org/docs/current/logical-replication-setup.html)