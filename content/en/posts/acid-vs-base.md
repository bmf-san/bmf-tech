---
title: About ACID and BASE
slug: acid-vs-base
date: 2025-08-02T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - ACID
  - BASE
description: An explanation of the ACID and BASE transaction models.
translation_key: acid-vs-base
---

The blog discusses the transaction models ACID and BASE.

## What is ACID

ACID represents the four properties of transactions primarily used in **Relational Databases (RDB)**.

| Property                  | Meaning                                           | Summary                                  |
| ------------------------- | ------------------------------------------------ | ---------------------------------------- |
| **Atomicity**             | Transactions either fully succeed or fully fail  | All or Nothing - No partial states       |
| **Consistency**           | Database integrity constraints are always maintained | Constraints, triggers, and rules are upheld |
| **Isolation**             | Concurrent transactions do not affect each other | Concurrent execution yields the same result as sequential execution |
| **Durability**            | Committed changes are permanently saved          | Changes persist even after system failures |

## Examples of ACID Application
- **Bank Transfers**: A account -10,000 yen, B account +10,000 yen succeed or fail together
- **E-commerce Inventory**: Atomic execution of inventory deduction and order confirmation
- **Reservation Systems**: Preventing double booking with exclusive control

## ACID Implementation Techniques

* **Locking**: Shared locks, exclusive locks
* **MVCC**: Multi-Version Concurrency Control
* **WAL**: Write-Ahead Logging
* **2PC**: Two-Phase Commit (for distributed environments)

## What is BASE

BASE is a **more relaxed consistency model** than ACID, primarily used in **NoSQL and large-scale distributed systems**. BASE has the following three characteristics:

| Property                                     | Meaning                     | Summary                           |
| -------------------------------------------- | --------------------------- | --------------------------------- |
| **Basically Available**                      | Always responds to some extent | Usable even without full consistency |
| **Soft state**                               | State can change            | Temporary inconsistencies are allowed |
| **Eventual consistency**                     | Consistency is achieved eventually | Assumes consistency over time |

## Examples of BASE Application
- **SNS Post Distribution**: Gradual delivery to followers' timelines
- **Search Indexing**: New content reflected in search results after a few minutes
- **CDN Updates**: Gradual cache updates across global locations

## BASE Implementation Techniques

* **Eventual Consistency**: Read Repair, Anti-Entropy
* **Conflict Resolution**: Last Writer Wins, Vector Clock, CRDT
* **Distributed Consensus**: Gossip Protocol, Merkle Tree
* **Distributed Storage**: Consistent Hashing, Quorum

## ACID vs BASE: Differences in Design Philosophy

| Comparison Axis       | ACID                           | BASE                            |
| --------------------- | ------------------------------ | ------------------------------- |
| **Consistency**       | Strong consistency             | Eventual consistency            |
| **Availability**      | May degrade during failures    | Maintains high availability     |
| **Scalability**       | Limited in distributed environments | Designed for distributed systems |
| **Latency**           | May introduce delays for consistency | Prioritizes low latency         |
| **Trade-off**         | Consistency > Availability     | Availability > Consistency      |
| **Application Areas** | Finance, enterprise, transactions | Web services, scalable applications |

## Summary
ACID provides strong consistency for transactions in relational databases, while BASE enables flexible consistency in distributed systems. Choosing the appropriate model depends on the application's requirements. ACID is suitable for financial and enterprise applications, while BASE is better for web services and scalable applications.

## References and Related Materials

- [en.wikipedia.org - ACID](https://en.wikipedia.org/wiki/ACID)
- [en.wikipedia.org - Eventual Consistency](https://en.wikipedia.org/wiki/Eventual_consistency)