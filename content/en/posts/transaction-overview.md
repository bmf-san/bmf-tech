---
title: Transaction Overview
slug: transaction-overview
date: 2023-06-07T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Transaction
translation_key: transaction-overview
---

# Overview
This post summarizes transactions.

# Transaction
A method to maintain data integrity. It is not a concept specific to databases but is independent as a theory.

It is necessary when we want to maintain data consistency in situations where multiple clients access the DB server simultaneously or when the DB server or application crashes during an update process.

Transactions provide two main functions:

- Concurrency Control
  - Preventing data inconsistencies that may arise from simultaneous access.
- Crash Recovery
  - Automatically performing recovery processes even if the DB server or application crashes.

To avoid data inconsistencies, it is possible to execute processes in a serialized schedule one by one without parallelization. However, executing in a serialized schedule is not practical in situations where many transactions are executed concurrently.

A state where data is correctly stored can be defined as one that results in the same outcome as when transactions are executed in a serialized schedule.

For practical execution control, the performance of the scheduler affects whether it can choose a schedule that results in the same outcome as a serialized schedule.

The performance of the scheduler is mainly indicated by:

- How many transactions can be parallelized
- The computational cost of finding the optimal schedule

In RDBs, a locking scheduler that uses locks is commonly used.

# ACID Properties
The properties required in transaction processing.

Even if it is not an RDB, those that satisfy these properties can be said to implement transactions.

## Atomicity - Commitment Control
The property that all operations included in a transaction either succeed (Success) or fail (Abort*).

*In SQL, it is ROLLBACK, but in transactions, it is ABORT.

It is not guaranteed that a transaction will succeed; rather, it is guaranteed that if it fails, everything will be rolled back.

Atomicity can be rephrased as the ability to roll back when a transaction encounters an error.

Applications that do not implement error handling (retry) when aborted are mismanaging transactions.

## Consistency - Exclusive Control
The property that data consistency must not be compromised before and after the execution of a transaction.

After executing a transaction, the DB must maintain a state without data inconsistencies, even if there are changes to the data.

This can be rephrased as the DB transitioning from one consistent state to another consistent state after executing a transaction.

It is the application, not the DB, that guarantees data consistency. In the case of RDBs, the relational model, which is the data model of RDBs, serves as the logic for determining consistency.

## Isolation - Exclusive Control
The property that multiple concurrently executing transactions do not affect each other.

The execution results of individual transactions must be the same as if they were executed in a serialized manner.

Isolation well represents the concurrency control of transactions.

## Durability - Fault Recovery Function
The property that once a commit is completed, the transaction does not disappear.

There is an implementation where confirmed transactions are not canceled, and it is possible to restore the data state before the crash through recovery.

After crash recovery, only the committed data remains, so once recovery is complete, the DB's consistency is guaranteed at the point of commit completion.

# Anomalies to Prevent with Transactions
Here, we enumerate the states that should not occur from the perspective of concurrency control in transactions.

The following transactions are abbreviated as TX.

## Lost Update
When TX1 writes data and another TX2 updates the same data, TX2 needs to determine the next data based on the result written by TX1. If TX2 updates the same data based on the data before TX1 updates, the update made by TX1 will be lost.

## Inconsistent Read
If the execution result of a TX affects the execution result of another TX, the consistency of the data read by the TX will be lost.

## Dirty Read
A phenomenon where uncommitted data from another TX can be read.

If TX2 reads uncommitted data after TX1's update, the data read by TX2 will no longer be correct if TX1 aborts.

## Non-Repeatable Read (Fuzzy Read, Non-Reproducible Read)
A phenomenon where reading data updated by another TX causes data consistency to be lost.

During a single TX, if the same data is read multiple times and the data changes even though the TX has not written anything, this phenomenon occurs.

## Phantom Read
A phenomenon where data inserted by another TX becomes visible, leading to a loss of data consistency.

If TX1 reads a certain range and during that time, TX2 commits data addition or deletion, the data may appear in TX1 as if it were a phantom.

# Schedules and Locks
The anomalies that transactions must prevent are states that should not occur during execution, and it is necessary to prevent such schedules from arising.

In RDBs, exclusive control using locks is commonly employed.

Locks can maintain data consistency during query execution by locking the rows being operated on before the operation is performed.

When a lock is applied, transactions that require access to the locked rows are blocked.

# Deadlock
In RDBs, deadlocks can occur in implementations such as row-level locks or page locks.

A deadlock is a state where two transactions block each other, causing processing to halt.

The countermeasures vary depending on the implementation.

# Types of Locks
Shared and exclusive are DB functions, while optimistic and pessimistic locks are strategies.

## Shared Lock (READ Lock)
A lock used when reading data. Other transactions cannot perform WRITE operations.

## Exclusive Lock (WRITE Lock)
A lock used when writing data. Other transactions cannot perform either READ or WRITE operations.

## Optimistic Lock (Optimistic Concurrency Control)
A method based on the optimistic assumption that simultaneous access to data will not occur.

By confirming that the data to be updated is in the same state as when it was retrieved, data inconsistencies are prevented.

The column used to determine whether the data is in the same state is called the lock key.

## Pessimistic Lock (Pessimistic Concurrency Control)
A method based on the pessimistic assumption that simultaneous access to data will frequently occur.

By locking the data to be updated when it is retrieved, data inconsistencies are prevented.

# Transaction Isolation Levels

| Isolation Level | Isolation | Dirty Read | Inconsistent Read | Lost Update | Phantom Read |
| -- | -- | -- | -- | -- | -- |
| READ-UNCOMMITTED | Low | ○ | ○ | ○ | ○ |
| READ-COMMITTED | ↓ | × | ○ | ○ | ○ |
| REPEATABLE-READ | ↓ | × | × | ○ | ○ |
| SERIALIZABLE | High | × | × | × | × |

# References
- [Learning Database Practice from Theory - Efficient SQL with the Relational Model](https://gihyo.jp/book/2015/978-4-7741-7197-5)
- [Qiita - Basics of Exclusive Control (Optimistic Lock, Pessimistic Lock)](https://qiita.com/NagaokaKenichi/items/73040df85b7bd4e9ecfc#%E6%A5%BD%E8%A6%B3%E3%83%AD%E3%83%83%E3%82%AF%E6%A5%BD%E8%A6%B3%E7%9A%84%E6%8E%92%E4%BB%96%E5%88%B6%E5%BE%A1)
- [There are Two Types of Locks in Databases](https://qiita.com/daiching/items/835fa37de22b397eece0)