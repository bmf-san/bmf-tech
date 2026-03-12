---
title: "MySQL Transaction Isolation Levels: Preventing Dirty Reads, Phantom Reads, and More"
slug: mysql-transaction-anomalies
date: 2023-06-08T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Transaction
  - MySQL
translation_key: mysql-transaction-anomalies
---



# Overview
This post summarizes transaction anomalies in MySQL. The MySQL version assumed is 8 series.

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

You can prepare a MySQL 8 series container with `docker compose up`.

# Transaction Isolation Levels
MySQL's InnoDB provides four transaction isolation levels as defined by the ANSI/ISO SQL standard.

| Isolation Level   | Dirty Read | Inconsistent Read | Lost Update | Phantom Read |
| ----------------- | ---------- | ----------------- | ----------- | ------------ |
| READ UNCOMMITTED  | ○          | ○                 | ○           | ○            |
| READ COMMITTED    | ×          | ○                 | ○           | ○            |
| REPEATABLE READ※1 | ×          | ×                 | ○           | ○※           |
| SERIALIZABLE      | ×          | ×                 | ×           | ×            |

※1 REPEATABLE READ is the default in MySQL.

※2 Although marked as ○ above, MySQL is designed to prevent phantom reads in REPEATABLE READ.

The transaction isolation level ranges from READ UNCOMMITTED, the lowest, to SERIALIZABLE, the highest. The above table is ordered from lowest to highest. Generally, the higher the isolation, the lower the performance tends to be.

For more on transactions, see [Transaction Overview](https://bmf-tech.com/posts/%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B6%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E6%A6%82%E8%A6%B3).

# Anomalies
Let's reproduce transaction anomalies in MySQL.

An anomaly refers to "unexpected results or inconsistencies arising from transaction isolation levels or processing order."

There are anomalies defined by ANSI SQL standards or ISO/IEC 9075, and there are various others besides those discussed here.

Inconsistent read is not defined by those standards. (I couldn't find where it is defined...)

Transactions are denoted as TX. Numbers are used to distinguish multiple transactions (e.g., TX1, TX2).

## Dirty Read
Dirty read is a phenomenon where TX1 reads data from TX2 before TX2 commits.

### Test
All sessions are conducted with READ UNCOMMITTED.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
```

1. Start transactions in TX1 and TX2
```sql
// TX1
mysql> START TRANSACTION;
// TX2
mysql> START TRANSACTION;
```

2. Add data in TX2
```sql
// TX2
mysql> INSERT INTO users(name) VALUES('foo');
```

Data is added in TX2, but not committed.

3. Read data again in TX1
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

TX1 reads data from TX2 before TX2 commits.

## Inconsistent Read
Inconsistent read is a phenomenon where the data being read lacks consistency.

Refer to [Various Anomalies#Inconsistent Read Anomaly](https://qiita.com/kumagi/items/5ef5e404546736ebac49#inconsistent-read-anomaly).

I wasn't sure about the exact definition, so my understanding might be questionable...

Since it's about inconsistency after commit, inconsistent read seems like a higher concept than fuzzy read or phantom read?? But strictly, it should be different...

### Test
All sessions are conducted with READ UNCOMMITTED.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
```

1. Start transaction and read data in TX1
```sql
// TX1
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // Empty set
```

2. Start transaction and add data in TX2
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(name) VALUES('foo');
mysql> COMMIT;
```

3. Read data again in TX1
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

The result differs from the initial read, confirming inconsistency due to TX2's actions.

## Fuzzy Read (Non-repeatable Read)
Fuzzy read is a phenomenon where TX1 can reference data updated by another TX2.

All sessions are conducted with READ COMMITTED.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
```

### Test
1. Start transaction and read data in TX1
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // Initial data entry
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // 1 row in set
```

Initial data entry result.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | foo  |
+-----+------+
```


2. Start transaction and read data in TX2
```sql
// TX2
mysql> START TRANSACTION;
mysql> UPDATE users SET name = 'bar' WHERE id = 1;
mysql> COMMIT;
mysql> SELECT * FROM users; // 1 row in set
```

Update is complete.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | bar  |
+-----+------+
```


4. Read data again in TX1
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

TX1's read result changes due to TX2's commit.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | bar  |
+-----+------+
```

## Phantom Read
Phantom read is a phenomenon where data read by TX1 changes if TX2 commits an addition or deletion. Fuzzy read involves updates, while phantom read involves additions or deletions.

### Test
All sessions are conducted with READ COMMITTED.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
```

1. Start transaction and read data in TX1
```sql
// TX1
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // Empty set
```

2. Add data and commit in TX2
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(name) VALUES('foo');
mysql> COMMIT;
mysql> SELECT * FROM users;
```

Addition is complete.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | foo  |
+-----+------+
```

3. Read data again in TX1
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

TX1's read result changes due to TX2's commit.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | foo  |
+-----+------+
```

## Lost Update
Lost update is a phenomenon where a conflict occurs when TX1 and TX2 update the same data, resulting in some updates being lost.

### Test
All sessions are conducted with REPEATABLE READ.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;
```

1. Start transaction and read data in TX1
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // Initial data entry
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // 1 row set
```

2. Start transaction and read data in TX2
```sql
// TX2
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // 1 row set
```

3. Update data in TX1 and TX2
```sql
// TX1
mysql> UPDATE users SET name = 'tx1' WHERE id = 1;

// TX2
mysql> UPDATE users SET name = 'tx2' WHERE id = 1;
```

4. Commit TX1 and TX2
```sql
// TX1
mysql> COMMIT;
// TX2
mysql> COMMIT;
```

5. Read data
```sql
mysql> SELECT * FROM users; 1 row set
```

TX1's commit is lost, and TX2's commit is reflected.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | tx2  |
+-----+------+
```

# Summary
The anomalies that occur vary depending on the transaction isolation level.

Anomalies are patterns where data reading and consistency change before and after commit.

To learn more about transaction anomalies, it might be better to refer to a book or something related to transactions.

# References
- [dev.mysql.com - 15.7.2.1 Transaction Isolation Levels](https://dev.mysql.com/doc/refman/8.0/ja/innodb-transaction-isolation-levels.html)
- [zenn.dev - Organizing Transaction Isolation Levels and Anomalies in MySQL/Postgres](https://zenn.dev/mpyw/articles/rdb-transaction-isolations#%E7%90%86%E8%AB%96%E9%9D%A2%E3%81%AE%E5%89%8D%E6%8F%90%E7%9F%A5%E8%AD%98)
