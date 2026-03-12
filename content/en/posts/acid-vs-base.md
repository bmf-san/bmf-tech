---
title: "ACID vs BASE: Understanding Database Consistency Models"
description: 'Compare ACID and BASE consistency models. Understand the trade-offs involved, when to use each, and how the CAP theorem connects to your database design choices.'
slug: acid-vs-base
date: 2025-08-02T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - ACID
  - BASE
translation_key: acid-vs-base
---

This post discusses the transaction models ACID and BASE.

## What is ACID?

ACID represents the four properties of transactions primarily used in **Relational Databases (RDB)**.

| Property                      | Meaning                                             | Summary                                   |
| ----------------------------- | -------------------------------------------------- | ----------------------------------------- |
| **Atomicity**                | Transactions either succeed completely or fail completely | All or Nothing - No intermediate state remains |
| **Consistency**              | Database integrity constraints are always maintained | Constraints, triggers, and rules are upheld |
| **Isolation**                | Concurrent transactions do not affect each other   | Same result as sequential execution even when executed concurrently |
| **Durability**               | Committed changes are permanently saved            | Changes are retained even in case of system failure |

## Specific Examples of ACID
- **Bank Transfer**: A account -10,000 yen, B account +10,000 yen either succeed or fail simultaneously
- **E-commerce Inventory**: Inventory deduction and order confirmation executed atomically
- **Reservation System**: Exclusive control to prevent double booking of seats

## ACID Implementation Techniques

* **Locking**: Shared lock, exclusive lock
* **MVCC**: Multi-Version Concurrency Control
* **WAL**: Write-Ahead Logging
* **2PC**: Two-Phase Commit (in distributed environments)

## What is BASE?

BASE is a **more relaxed consistency model** primarily used in **NoSQL and large-scale distributed systems**. BASE has the following three properties.

| Property                                   | Meaning                   | Summary                           |
| ------------------------------------------ | -------------------------- | -------------------------------- |
| **Basically Available**                    | Responds most of the time | Available even without complete consistency |
| **Soft State**                             | State can change          | Temporary inconsistency is acceptable |
| **Eventual Consistency**                   | Consistency will be achieved eventually | Assumes consistency over time |

## Specific Examples of BASE
- **SNS Post Delivery**: Gradually delivered to followers' timelines
- **Search Index**: New content reflected in search results after a few minutes
- **CDN Updates**: Caches around the world updated sequentially

## BASE Implementation Techniques

* **Eventual Consistency**: Read Repair, Anti-Entropy
* **Conflict Resolution**: Last Writer Wins, Vector Clock, CRDT
* **Distributed Consensus**: Gossip Protocol, Merkle Tree
* **Distributed Storage**: Consistent Hashing, Quorum

## ACID vs BASE: Differences in Design Philosophy

| Comparison Axis   | ACID                             | BASE                             |
| ------------------| -------------------------------- | -------------------------------- |
| **Consistency**   | Strong consistency (Strong)      | Eventual consistency             |
| **Availability**  | May decrease during failures     | Maintains high availability       |
| **Distribution**  | Many constraints in distributed environments | Designed for distributed systems  |
| **Latency**       | Delays occur for consistency guarantees | Prioritizes low latency           |
| **Trade-off**     | Consistency > Availability       | Availability > Consistency        |
| **Application Areas** | Finance, business, transactions | Web services, scalable applications |

## Conclusion
ACID provides strong consistency for transactions in relational databases, while BASE achieves flexible consistency in distributed systems. It is important to choose which model to adopt based on application requirements. ACID is suitable for financial and business applications, while BASE is better for web services and scalable applications.

## References and Related Materials

- [en.wikipedia.org - ACID](https://en.wikipedia.org/wiki/ACID)
- [en.wikipedia.org - Eventual Consistency](https://en.wikipedia.org/wiki/Eventual_consistency)