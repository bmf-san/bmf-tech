---
title: MySQL Transaction Anomalies
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
This post summarizes the anomalies of MySQL transactions.
Assuming MySQL version 8.

# Test Environment
The environment used for testing is prepared with docker-compose. (Since it's only one container, using compose is not necessary...)

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

# Transaction Isolation Levels
MySQL's InnoDB provides the four transaction isolation levels defined by the ANSI/ISO SQL standard.

|    Isolation Level   | Dirty Read | Inconsistent Read | Lost Update | Phantom Read |
| -------------------- | ---------- | ----------------- | ----------- | ------------ |
| READ UNCOMMITTED     | ○          | ○                 | ○           | ○            |
| READ COMMITTED       | ×          | ○                 | ○           | ○            |
| REPEATABLE READ※1   | ×          | ×                 | ○           | ○※           |
| SERIALIZABLE         | ×          | ×                 | ×           | ×            |

※1 REPEATABLE READ is the default in MySQL.

※2 Although marked as ○ above, MySQL prevents phantom reads from occurring in REPEATABLE READ.

The transaction isolation level of READ UNCOMMITTED is the lowest, while SERIALIZABLE is the highest. The above list is ordered from lowest to highest. Generally, higher isolation levels tend to decrease performance.

For more on transactions, see [Transaction Overview](https://bmf-tech.com/posts/%E3%83%88%E3%83%A9%E3%83%B3%E3%82%B8%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E6%A6%82%E8%A6%B3).

# Anomalies
Let's reproduce transaction anomalies in MySQL.

An anomaly refers to "unexpected results or inconsistencies that arise from the transaction isolation level and processing order."

Some anomalies are defined by the ANSI SQL standard and ISO/IEC 9075, and there are various others beyond those covered here.

Inconsistent Read is not defined by those standards. (I couldn't find where it is defined...)

Transactions are denoted as TX. Numbers are used to distinguish multiple transactions (e.g., TX1, TX2).

## Dirty Read
A dirty read is the phenomenon where TX1 reads data before TX2 has committed it.

### Verification
All sessions are conducted with READ UNCOMMITTED.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
```

1. Start transactions in TX1 and TX2.
```sql
// TX1
mysql> START TRANSACTION;
// TX2
mysql> START TRANSACTION;
```

2. Add data in TX2.
```sql
// TX2
mysql> INSERT INTO users(name) VALUES('foo');
```

Data is added in TX2, but not committed.

3. Read data again in TX1.
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

TX1 has read data before TX2's commit.

## Inconsistent Read
An inconsistent read is the phenomenon where the data read lacks consistency.

Refer to [Various Anomaly#Inconsistent Read Anomaly](https://qiita.com/kumagi/items/5ef5e404546736ebac49#inconsistent-read-anomaly).

I wasn't quite sure about the exact definition, so my understanding might be off...

Since it refers to inconsistency after commit, inconsistent read seems to be a higher-level concept than fuzzy read or phantom read?? But strictly speaking, it might be different...

### Verification
All sessions are conducted with READ UNCOMMITTED.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
```

1. Start transaction and read data in TX1.
```sql
// TX1
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // Empty set
```

2. Start transaction in TX2 and add data.
```sql
// TX2
mysql> START TRANSACTION;
mysql> INSERT INTO users(name) VALUES('foo');
mysql> COMMIT;
```

3. Read data again in TX1.
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

It is confirmed that the result differs from the initial read (TX2's result), indicating a lack of consistency.

## Fuzzy Read (Non-Repeatable Read)
A fuzzy read is the phenomenon where TX1 can reference data updated by another TX2.

All sessions are conducted with READ COMMITTED.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
```

### Verification
1. Start transaction and read data in TX1.
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // Initial data input
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // 1 row in set
```

Initial data input result.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | foo  |
+-----+------+
```

2. Start transaction and read data in TX2.
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

4. Read data again in TX1.
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

It is confirmed that TX1's read result has changed due to TX2's commit.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | bar  |
+-----+------+
```

## Phantom Read
A phantom read is the phenomenon where the data read by TX1 changes when TX2 commits a new addition or deletion.
Fuzzy read involves updates, while phantom read concerns new additions or deletions.

### Verification
All sessions are conducted with READ COMMITTED.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
```

1. Start transaction and read data in TX1.
```sql
// TX1
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // Empty set
```

2. Add data in TX2 and commit.
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

3. Retrieve data again in TX1.
```sql
// TX1
mysql> SELECT * FROM users; // 1 row in set
```

It is confirmed that TX1's read result has changed due to TX2's commit.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | foo  |
+-----+------+
```

## Lost Update
A lost update occurs when TX1 and TX2 update the same data, causing some updates to be lost.

### Verification
All sessions are conducted with REPEATABLE READ.
```sql
mysql> SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;
```

1. Start transaction and read data in TX1.
```sql
// TX1
mysql> INSERT INTO users(name) VALUES('foo'); // Initial data input
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // 1 row set
```

2. Start transaction and read data in TX2.
```sql
// TX2
mysql> START TRANSACTION;
mysql> SELECT * FROM users; // 1 row set
```

3. Update data in TX1 and TX2.
```sql
// TX1
mysql> UPDATE users SET name = 'tx1' WHERE id = 1;

// TX2
mysql> UPDATE users SET name = 'tx2' WHERE id = 1;
```

4. Commit TX1 and TX2.
```sql
// TX1
mysql> COMMIT;
// TX2
mysql> COMMIT;
```

5. Read data.
```sql
mysql> SELECT * FROM users; // 1 row set
```

It is confirmed that TX1's commit was lost and only TX2's commit is reflected.
```
+-----+------+
| id  | name |
+-----+------+
| 1   | tx2  |
+-----+------+
```

# Conclusion
The anomalies that occur vary depending on the transaction isolation level.

Anomalies manifest as changes in data reading and consistency before and after commits.

To learn more about transaction anomalies, it might be best to refer to a book or resource on transactions.

# References
- [dev.mysql.com - 15.7.2.1 Transaction Isolation Levels](https://dev.mysql.com/doc/refman/8.0/ja/innodb-transaction-isolation-levels.html)
- [zenn.dev - Organizing Transaction Isolation Levels and Anomalies in MySQL/Postgres](https://zenn.dev/mpyw/articles/rdb-transaction-isolations#%E7%90%86%E8%AB%96%E9%9D%A2%E3%81%AE%E5%89%8D%E6%8F%90%E7%9F%A5%E8%AD%98)