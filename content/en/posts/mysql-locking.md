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
This post summarizes MySQL locks.
Assuming MySQL version 8.

# Verification Environment
The environment used for verification is prepared with docker-compose. (It's just one container, so using compose isn't necessary...)

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

You can prepare a MySQL 8 container with `docker compose up`.

# Locks
## Internal Level Locks
In MySQL, there are two methods for concurrency control: row-level locks and table-level locks.

cf. [dev.mysql.com - 8.11.1 Internal Locking Methods](https://dev.mysql.com/doc/refman/8.0/en/internal-locking.html)

- Row-Level Lock
  - A lock targeting individual rows within a table.
  - Since the lock target is narrow, lock contention and rollback changes are minimized.
  - A single row can be locked for a long time.
- Table-Level Lock
  - A lock targeting the entire table.
  - Requires relatively less memory (row locks require memory for each locked row or group of rows).
  - Fast when only a single lock is needed, especially when targeting most of the table.
  - Fast when frequently executing GROUP BY on most of the data or scanning the entire table.

## InnoDB Locks
cf. [dev.mysql.com - 15.7.1 InnoDB Locks](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html)

### Shared (READ) Lock
A shared lock allows reading data but not writing. Shared lock (IS).

#### Verification
1. Start a transaction in TX1 and apply a shared lock.
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // Initial data insertion
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE id = 1 LOCK IN SHARE MODE;
```

2. Start a transaction in TX2 and perform a WRITE.
```sql
// TX2
mysql> START TRANSACTION;
mysql> UPDATE users SET name = 'bar' WHERE id = 1;
```

TX2's update is locked until TX1 commits.

### Exclusive (WRITE) Lock
An exclusive lock prevents both reading and writing of data. Exclusive lock (IX).

#### Verification
1. Start a transaction in TX1 and apply an exclusive lock.
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // Initial data insertion
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE id = 1 FOR UPDATE;
```

2. Start a transaction in TX2 and perform READ and WRITE.
```sql
// TX2
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE id = 1; // This is allowed
mysql> SELECT * FROM users WHERE id = 1 FOR UPDATE; // Not allowed
mysql> UPDATE users SET name = 'bar' WHERE id = 1; // Not allowed
```

It is confirmed that TX2 cannot perform READ (other than simple SELECT) or WRITE until TX1's lock is released.

### Intention Lock
A table-level lock indicating the type of lock (shared or exclusive) that a transaction requires on the rows of a table. It is designed to support coexistence of row locks and table locks.

There are two types of intention locks:

- Intention Shared Lock
- Intention Exclusive Lock

#### Verification
This cannot be explicitly manipulated via SQL and is generally managed internally by the database, so verification is omitted.

There are several verification patterns, but various tests are conducted in the following article.

cf. [qiita.com - Testing MySQL Locks While Reading Official Documentation: Row-Level Lock: Intention Lock](https://qiita.com/ham0215/items/2f38a2949d9012074c3d)

### Record Lock
Locking of index records. Index records refer to clustered indexes and secondary indexes. Locks are applied to the scanned indexes.

#### Verification
Omitted due to being an internal database operation.

### Gap Lock
Locking of gaps between index records, or gaps before or after index records.

#### Verification
1. Start a transaction in TX1 and perform a READ.
```sql
// TX1
mysql> INSERT INTO users(id, name) VALUES(1, 'foo'), (2, 'bar'), (4, 'qux'), (5, 'quux'), (6, 'corge'); // Initial data insertion
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE ID between 1 AND 5 FOR UPDATE;
```

2. Start a transaction in TX2 and perform a WRITE.
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(id, name) VALUES(3, 'baz');
```

It is confirmed that the lock is applied over a range rather than just row by row.

### Next-Key Lock
A combination of a record lock on an index record and a gap lock on the gap before that index record.

#### Verification
1. Start a transaction in TX1 and perform a READ.
```sql
// TX1
mysql> INSERT INTO users(id, name) VALUES(1, 'foo'), (2, 'bar'), (3, 'baz'), (4, 'qux'); // Initial data insertion
mysql> START TRANSACTION;
mysql> SELECT * FROM users WHERE ID < 5 FOR UPDATE;
```

2. Start a transaction in TX2 and perform a WRITE.
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(id, name) VALUES(5, 'quux');
```

It is confirmed that not only the rows with IDs less than 5 are locked, but also the gap after the last index value is locked.

### Intention Lock Insertion
A type of gap lock set by an INSERT before the row insertion. Intention lock for INSERT.

#### Verification
Omitted due to being an internal database operation.

This is verified in the following article.
cf. [Testing MySQL Locks While Reading Official Documentation: Record Lock / Gap Lock / Next-Key Lock / Others](https://qiita.com/ham0215/items/99679d499869365446ec#%E3%82%A4%E3%83%B3%E3%83%86%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%B3%E3%83%AD%E3%83%83%E3%82%AF%E3%81%AE%E6%8C%AF%E5%85%A5)

### AUTO-INC Lock
A table lock obtained by transactions inserting into tables containing AUTO_INCREMENT columns. This lock prevents TX2 from obtaining AUTO_INCREMENT values while TX1 is acquiring them for an INSERT.

#### Verification
Omitted due to being an internal operation and lack of reproducible methods.

### Spatial Index Predicate Lock
This refers to the documentation. (I wasn't quite sure about this as I'm not familiar with spatial indexes...)

cf. [Spatial Index Predicate Lock](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html#innodb-auto-inc-locks)

# How to Check Locks
Locks can be checked with the following queries.

```sql
// Check lock status
SELECT * FROM performance_schema.data_locks;

// Check number of locks + thread ID
SHOW ENGINE INNODB STATUS;

// Check number of locks
SELECT trx_id, trx_rows_locked, trx_mysql_thread_id FROM information_schema.INNODB_TRX;
```

To check for deadlocks, execute `SHOW ENGINE INNODB STATUS` and look for the section labeled `LATEST DETECTED DEADLOCK`.

# Conclusion
MySQL has both explicit and implicit locking patterns.

It seems beneficial to first focus on what is being targeted (whether it's a row or a table) and the extent of the range.

# References
- [dev.mysql.com - 8.11.1 Internal Locking Methods](https://dev.mysql.com/doc/refman/8.0/en/internal-locking.html)
- [dev.mysql.com - 15.7.1 InnoDB Locks](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html)
- [zenn.dev - Basics of Database Locks to Deadlocks](https://zenn.dev/gibjapan/articles/1d8dfb7520dabc)
- [qiita.com - Testing MySQL Locks While Reading Official Documentation: Record Lock / Gap Lock / Next-Key Lock / Others](https://qiita.com/ham0215/items/99679d499869365446ec)
- [qiita.com - Additional Information on MySQL Locks (Note: This content has already been extensively discussed)](https://qiita.com/hmatsu47/items/f5eb64428494686d4ad3)
- [qiita.com - Testing MySQL Locks While Reading Official Documentation: Row-Level Lock: Shared Lock (S) / Exclusive Lock (X)](https://qiita.com/ham0215/items/b9efc718670b1d2d48c1#%E8%A1%8C%E3%83%AC%E3%83%99%E3%83%AB%E3%83%AD%E3%83%83%E3%82%AF%E3%82%92%E6%A4%9C%E8%A8%BC%E3%81%99%E3%82%8B)
- [qiita.com - Testing MySQL Locks While Reading Official Documentation: Row-Level Lock: Intention Lock](https://qiita.com/ham0215/items/2f38a2949d9012074c3d)
- [techblog.cartaholdings.co.jp - Knowledge of Locks That You Should Actually Remember for Those Who Have Never Thought About DB Locks](https://techblog.cartaholdings.co.jp/entry/2022/12/14/113000)
- [www.wakuwakubank.com - Exclusive Lock (FOR UPDATE) and Shared Lock (LOCK IN SHARE MODE)](https://www.wakuwakubank.com/posts/201-mysql-lock/)
- [saekis.hatenablog.com - Confirming the Behavior of MySQL Exclusive Locks](https://saekis.hatenablog.com/entry/2019/02/06/191454)
- [bizstation.hatenablog.com - Detailed Control of InnoDB Locks in MySQL/MariaDB and Transactd Part 1](https://bizstation.hatenablog.com/entry/2014/12/24/103641)
- [devsakaso.com - About MySQL Locks and Deadlocks](https://devsakaso.com/mysql-about-lock-and-deadlock/)
- [nishinatoshiharu.com - Overview and Behavior Verification of InnoDB Shared and Exclusive Locks](https://nishinatoshiharu.com/db-lock-overview/)
- [www.wakuwakubank.com - Exclusive Lock (FOR UPDATE) and Shared Lock (LOCK IN SHARE MODE)](https://www.wakuwakubank.com/posts/201-mysql-lock/)
- [free-engineer.life - MySQL (InnoDB) Shared Lock, Exclusive Lock, and Intention Lock (Table Lock)](https://free-engineer.life/mysql-innodb-lock-mode/)
- [free-engineer.life - MySQL (InnoDB) Row Locks](https://free-engineer.life/mysql-innodb-record-locks/)
- [github.com - Investigation of MySQL InnoDB Lock Behavior](https://github.com/ichirin2501/doc/blob/master/innodb.md)
- [github.com - MySQL Deadlock Analysis Method Using Thread ID](https://github.com/ichirin2501/doc/blob/master/innodb-deadlock-thread-id.md)