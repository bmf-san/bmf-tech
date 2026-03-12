---
title: Transaction Overview
description: Research notes and a structured overview of Transaction Overview, summarizing key concepts and findings.
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

# Transactions
A method to maintain data integrity. It is not a concept unique to databases but stands independently as a theory.

Transactions are necessary when you want to protect data integrity in situations where multiple clients access a DB server simultaneously, or when a DB server or application crashes during an update process.

Transactions provide two main functions:

- Concurrency Control
  - Prevents data inconsistencies that may occur due to concurrent access.
- Crash Recovery
  - Ensures automatic recovery even if the DB server or application crashes.

To avoid data inconsistencies, it is possible to execute processes in a serialized schedule, one by one, without parallelization. However, executing in a serialized schedule is not practical in situations where many transactions are running concurrently.

A state where data is correctly saved can be defined as a state that results in the same outcome as when transactions are executed in a serialized schedule.

For practical execution control, the performance of the scheduler, which determines whether it can select a schedule that results in the same outcome as a serialized schedule, is influential.

The performance of a scheduler is mainly measured by:

- How many transactions can be parallelized
- The computational cost of finding the optimal schedule

In RDBs, locking schedulers that use locks are widely used.

# ACID Properties
The properties required in transaction processing.

Even if not in RDBs, those that meet these properties can be said to implement transactions.

## Atomicity - Commitment Control
The property that all operations within a transaction either succeed (Success) or fail (Abort*).

*In SQL, it's ROLLBACK, but in transactions, it's called ABORT.

Transactions are not guaranteed to succeed, but it is guaranteed that if they fail, everything is rolled back.

Atomicity can be rephrased as the ability to roll back if a transaction encounters an error.

Applications that do not implement error handling (retry) when an abort occurs are handling transactions incorrectly.

## Consistency - Exclusive Control
The property that data consistency must not be compromised before and after transaction execution.

After executing a transaction, the DB must maintain a state without data inconsistencies, even if data changes.

This can be rephrased as the DB transitioning from one consistent state to another consistent state after executing a transaction.

Ensuring data consistency is the responsibility of the application, not the DB. In RDBs, the relational model, which is the data model of RDBs, serves as the logic for determining consistency.

## Isolation - Exclusive Control
The property that multiple transactions executed simultaneously do not affect each other.

The execution result of each transaction must be the same as if the transactions were executed serially.

Isolation well represents the concurrency control of transactions.

## Durability - Fault Recovery Function
The property that once a transaction is committed, it is not lost.

There is an implementation that ensures committed transactions are not canceled, and even if a crash occurs, it is possible to restore to the data state before the crash through recovery.

After crash recovery, only committed data remains, ensuring the consistency of the DB at the point of commit completion.

# Abnormalities Transactions Should Prevent
From the perspective of concurrency control, states that should not occur in transactions are listed.

Transactions are abbreviated as TX below.

## Lost Update
When TX1 writes data and another TX2 updates the same data, TX2 needs to decide the next data based on the result written by TX1. However, if TX2 updates the same data based on the data before TX1 updates it, the update made by TX1 disappears.

## Inconsistent Read
If the execution result of a TX affects another TX, the consistency of the data read by the TX is lost.

## Dirty Read
A phenomenon where data not committed by another TX is read.

If TX2 reads data updated by TX1 before it is committed, and TX1 aborts, the data read by TX2 becomes incorrect.

## Non-repeatable Read (Fuzzy Read, Non-reproducible Read)
A phenomenon where reading data updated by another TX causes data consistency to be lost.

In a single TX, even if the same data is read multiple times, the data changes without the TX writing to it.

## Phantom Read
A phenomenon where data inserted by another TX becomes visible, causing data consistency to be lost.

If TX1 reads a certain range and TX2 adds or deletes data and commits, the data appears to be reflected in TX1 like a phantom.

# Schedule and Lock
Abnormalities that transactions should prevent are inherently not executable, and it is necessary to prevent such schedule occurrences.

In RDBs, exclusive control using locks is commonly used.

Locks protect data consistency during query execution by locking the rows to be operated on before the operation is performed.

When a lock is applied, transactions that require access to the locked rows are blocked.

# Deadlock
In RDBs, deadlocks can occur in implementations such as row-level locks or page locks.

A deadlock is a state where two transactions block each other, causing the process to stall.

Countermeasures vary depending on the implementation.

# Types of Locks
Shared/Exclusive are DB functions, while Optimistic/Pessimistic locks are policies.

## Shared Lock (READ Lock)
A lock used when reading data. Other transactions cannot WRITE.

## Exclusive Lock (WRITE Lock)
A lock used when writing data. Other transactions cannot READ or WRITE.

## Optimistic Lock (Optimistic Concurrency Control)
A method based on the optimistic assumption that simultaneous access to data will not occur.

By confirming that the data to be updated is in the same state as when it was retrieved, data inconsistencies are prevented.

The column used to determine if the data is in the same state is called a lock key.

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
- [Learn Database Practice from Theory
―― Efficient SQL with the Relational Model](https://gihyo.jp/book/2015/978-4-7741-7197-5)
- [Qiita - Basics of Exclusive Control (Optimistic Lock, Pessimistic Lock)](https://qiita.com/NagaokaKenichi/items/73040df85b7bd4e9ecfc#%E6%A5%BD%E8%A6%B3%E3%83%AD%E3%83%83%E3%82%AF%E6%A5%BD%E8%A6%B3%E7%9A%84%E6%8E%92%E4%BB%96%E5%88%B6%E5%BE%A1)
- [There are two types of "lock" concepts in databases](https://qiita.com/daiching/items/835fa37de22b397eece0)
