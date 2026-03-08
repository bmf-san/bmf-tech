---
title: About MySQL Locks
slug: mysql-locking
date: 2024-04-05T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - MySQL
  - Transaction
  - Lock
translation_key: mysql-locking
---

# Overview
This post summarizes MySQL locks, assuming version 8.

# Test Environment
The environment used for testing is prepared with docker-compose. (Although it's just one container, so you don't necessarily need to use compose...)

```sh
.
├── docker-compose.yml
└── initdb.d
    └── 1_schema.sql
```

docker-compose.yml
```yml
version: '3'

services:
  mysql:
    image: mysql:8.0.33
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: example
      TZ: "Asia/Tokyo"
    command: mysqld
    ports:
      - 3306:3306
    volumes:
      - ./initdb.d:/docker-entrypoint-initdb.d
```

1_schema.sql
```sql
CREATE DATABASE IF NOT EXISTS example;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) NOT NULL UNIQUE
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
```

You can set up a MySQL 8 container with `docker compose up`.

# Locks
## Internal Level Locks
In MySQL, there are row-level locks and table-level locks as methods of exclusive control.

cf. [dev.mysql.com - 8.11.1 Internal Locking Methods](https://dev.mysql.com/doc/refman/8.0/en/internal-locking.html)

- Row-level locks
  - Locks targeting individual rows in a table
  - Narrow lock targets reduce lock contention and rollback changes
  - Allows long-term locking of a single row
- Table-level locks
  - Locks targeting the entire table
  - Requires relatively less memory (row locks need memory for each locked row or group of rows)
  - Fast when used for most of the table as only a single lock is needed
  - Fast when frequently executing GROUP BY on most of the data or scanning the entire table

## InnoDB Locks
cf. [dev.mysql.com - 15.7.1 InnoDB Locking](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html)

### Shared (READ) Lock
A shared lock allows READ but not WRITE. Shared lock (IS).

#### Verification
1. Start a transaction in TX1 and apply a shared lock
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // Initial data input
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE id = 1 LOCK IN SHARE MODE;
```

2. Start a transaction in TX2 and perform WRITE
```sql
// TX2
mysql> START TRANSACTION;
mysql> UPDATE users SET name = 'bar' WHERE id = 1;
```

TX2's update is locked until TX1 commits.

### Exclusive (WRITE) Lock
An exclusive lock prevents both READ and WRITE. Exclusive lock (IX).

#### Verification
1. Start a transaction in TX1 and apply an exclusive lock
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // Initial data input
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE id = 1 FOR UPDATE;
```

2. Start a transaction in TX2 and perform READ and WRITE
```sql
// TX2
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE id = 1; // Allowed
mysql> SELECT * FROM users WHERE id = 1 FOR UPDATE; // Not allowed
mysql> UPDATE users SET name = 'bar' WHERE id = 1; // Not allowed
```

TX2 cannot perform READ (other than simple SELECT) or WRITE until TX1's lock is released.

### Intention Locks
Table-level locks indicating the type of lock (shared or exclusive) a transaction requires on a table's rows. They support coexistence of row and table locks.

There are two types of intention locks:

- Intention shared lock
- Intention exclusive lock

#### Verification
Not explicitly operable via SQL and generally managed internally by the database, so verification is omitted.

Various verification patterns are available, as explored in the following article.

cf. [qiita.com - Exploring MySQL Locks with Official Documentation](https://qiita.com/ham0215/items/2f38a2949d9012074c3d)

### Record Locks
Locks on index records, which include clustered and secondary indexes. Locks the scanned indexes.

#### Verification
Omitted as it's an internal database operation.

### Gap Locks
Locks the gaps between index records or before/after index records.

#### Verification
1. Start a transaction in TX1 and perform READ
```sql
// TX1
mysql> INSERT INTO users(id, name) VALUES(1, 'foo'), (2, 'bar'), (4, 'qux'), (5, 'quux'), (6, 'corge'); // Initial data input
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE ID between 1 AND 5 FOR UPDATE;
```

2. Start a transaction in TX2 and perform WRITE
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(id, name) VALUES(3, 'baz');
```

It appears to be row-level locking, but it's confirmed to be range-locked.

### Next-Key Locks
A combination of record locks on index records and gap locks on the gap before the index record.

#### Verification
1. Start a transaction in TX1 and perform READ
```sql
// TX1
mysql> INSERT INTO users(id, name) VALUES(1, 'foo'), (2, 'bar'), (3, 'baz'), (4, 'qux'); // Initial data input
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE ID < 5 FOR UPDATE;
```

2. Start a transaction in TX2 and perform WRITE
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(id, name) VALUES(5, 'quux');
```

It's confirmed that not only rows with id less than 5 are locked, but also the gap after the row with the highest index value.

### Insert Intention Locks
A type of gap lock set by an INSERT before inserting a row. Insert intention lock.

#### Verification
Omitted as it's an internal database operation.

Refer to this article for verification.
cf. [Exploring MySQL Locks with Official Documentation](https://qiita.com/ham0215/items/99679d499869365446ec#%E3%82%A4%E3%83%B3%E3%83%86%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%B3%E3%83%AD%E3%83%83%E3%82%AF%E3%81%AE%E6%8C%BF%E5%85%A5)

### AUTO-INC Locks
Table locks acquired by transactions inserting into a table with an AUTO_INCREMENT column. Prevents TX2 from acquiring AUTO_INCREMENT values while TX1 is acquiring them for INSERT.

#### Verification
Omitted due to internal operation and lack of reproduction method.

### Predicate Locks for Spatial Indexes
Refer to the documentation. (I'm not familiar with spatial indexes, so I didn't fully understand...)

cf. [Predicate Locks for Spatial Indexes](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html#innodb-auto-inc-locks)

# Checking Locks
Locks can be checked with the following queries.

```sql
// Check lock status
SELECT * FROM performance_schema.data_locks;

// Check lock count + thread ID
SHOW ENGINE INNODB STATUS;

// Check lock count
SELECT trx_id, trx_rows_locked, trx_mysql_thread_id FROM information_schema.INNODB_TRX;
```

To check for deadlocks, execute `SHOW ENGINE INNODB STATUS` and look for the section labeled `LATEST DETECTED DEADLOCK`.

# Summary
MySQL has patterns of explicit and implicit locks.

It's beneficial to first focus on what is being locked (row or table) and the extent of the range.

# References
- [dev.mysql.com - 8.11.1 Internal Locking Methods](https://dev.mysql.com/doc/refman/8.0/en/internal-locking.html)
- [dev.mysql.com - 15.7.1 InnoDB Locking](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html)
- [zenn.dev - Basics of Database Locks to Deadlocks](https://zenn.dev/gibjapan/articles/1d8dfb7520dabc)
- [qiita.com - Exploring MySQL Locks with Official Documentation](https://qiita.com/ham0215/items/99679d499869365446ec)
- [qiita.com - Supplement on MySQL Locks (Note: Already well-discussed content)](https://qiita.com/hmatsu47/items/f5eb64428494686d4ad3)
- [qiita.com - Exploring MySQL Locks with Official Documentation](https://qiita.com/ham0215/items/b9efc718670b1d2d48c1#%E8%A1%8C%E3%83%AC%E3%83%99%E3%83%AB%E3%83%AD%E3%83%83%E3%82%AF%E3%82%92%E6%A4%9C%E8%A8%BC%E3%81%99%E3%82%8B)
- [qiita.com - Exploring MySQL Locks with Official Documentation](https://qiita.com/ham0215/items/2f38a2949d9012074c3d)
- [techblog.cartaholdings.co.jp - Knowledge on Locks You Should Know](https://techblog.cartaholdings.co.jp/entry/2022/12/14/113000)
- [www.wakuwakubank.com - Exclusive Lock (FOR UPDATE) and Shared Lock (LOCK IN SHARE MODE)](https://www.wakuwakubank.com/posts/201-mysql-lock/)
- [saekis.hatenablog.com - Verifying MySQL Exclusive Lock Behavior](https://saekis.hatenablog.com/entry/2019/02/06/191454)
- [bizstation.hatenablog.com - Detailed InnoDB Lock Control in MySQL/MariaDB and Transactd](https://bizstation.hatenablog.com/entry/2014/12/24/103641)
- [devsakaso.com - About MySQL Locks and Deadlocks](https://devsakaso.com/mysql-about-lock-and-deadlock/)
- [nishinatoshiharu.com - Overview and Behavior Verification of InnoDB Shared and Exclusive Locks](https://nishinatoshiharu.com/db-lock-overview/)
- [www.wakuwakubank.com - Exclusive Lock (FOR UPDATE) and Shared Lock (LOCK IN SHARE MODE)](https://www.wakuwakubank.com/posts/201-mysql-lock/)
- [free-engineer.life - MySQL (InnoDB) Shared, Exclusive, and Intention Locks](https://free-engineer.life/mysql-innodb-lock-mode/)
- [free-engineer.life - MySQL (InnoDB) Row Locks](https://free-engineer.life/mysql-innodb-record-locks/)
- [github.com - Investigation of InnoDB Lock Behavior in MySQL](https://github.com/ichirin2501/doc/blob/master/innodb.md)
- [github.com - Deadlock Analysis Method Using Thread ID in MySQL](https://github.com/ichirin2501/doc/blob/master/innodb-deadlock-thread-id.md)