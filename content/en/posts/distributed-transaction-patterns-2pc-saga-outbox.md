---
title: "The Three Pillars of Distributed Transactions — 2PC, Saga, and Outbox"
description: A concise comparison of the three canonical distributed transaction patterns in microservices — 2PC, Saga, and Outbox — covering their mechanics, trade-offs, and when to use each.
slug: distributed-transaction-patterns-2pc-saga-outbox
date: 2026-05-13T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Microservices
  - Distributed Transactions
  - 2phase commit
  - Saga Pattern
  - Outbox Pattern
translation_key: distributed-transaction-patterns-2pc-saga-outbox
---


Once you adopt microservices, you almost inevitably run into the question: **how do you keep writes consistent when they span multiple databases, or span a database and a message broker?** This article summarizes the three canonical approaches — 2PC, Saga, and Outbox.

## Premise: Why It Wasn't a Problem in a Monolith

In a monolith, every table lives in the same database, so wrapping things in `BEGIN; ... COMMIT;` is enough for the DB engine to guarantee Atomicity, Consistency, Isolation, and Durability (ACID). The application only has to draw transaction boundaries.

The moment you split services and split the database too, that premise collapses. **A COMMIT across heterogeneous resources is outside the DB engine's responsibility.** That's when distributed transaction design first becomes necessary.

## 1. 2PC (Two-Phase Commit)

A protocol in which a coordinator decides, in two phases, **either "everyone COMMITs" or "everyone ROLLBACKs."**

### How It Works

- **Phase 1 (Prepare)**: The coordinator asks each participant "can you commit?" Each participant acquires locks and replies Yes/No in a Prepared state.
- **Phase 2 (Commit)**: If everyone says Yes, send COMMIT; if anyone says No, send ROLLBACK.

### Pros

- Guarantees strong (immediate) consistency.
- The protocol itself is simple and clear.

### Cons

- **Blocking**: If the coordinator dies after Prepare, participants are stuck holding locks.
- **Low availability**: A single slow or down node halts everything.
- **No support for heterogeneous resources**: It assumes XA support; Pub/Sub, Kafka, HTTP APIs, etc. are out of scope.
- **Against the spirit of microservices**: "We split the services, but one failure halts the whole thing" is self-defeating.

### When to Use It

Almost never. **Between microservices, the rule of thumb is to avoid it.** Consider it only in narrow cases like aligning multiple databases on the same DB engine.

## 2. The Saga Pattern

A pattern that composes a long-running transaction as a chain of **local transactions plus "compensating transactions"** at each step.

### How It Works

- Each step Ti commits independently.
- If a later step fails, the effects of previously successful Ti are undone in reverse order via compensations Ci.
- Compensation is not a physical rollback — it's an operation that **makes things "as if it never happened" from a business perspective** (e.g., the compensation for a withdrawal is a re-deposit).

### Implementation Styles

- **Choreography**: No central coordinator; each service emits and reacts to events to progress to the next step. Loosely coupled but the flow is hard to see.
- **Orchestration**: A central orchestrator owns the flow. Easier to visualize and test, but centralizes control.

### Pros

- High availability (no synchronous wait on all participants).
- Works across heterogeneous resources.
- Fits microservices well.

### Cons

- **Sacrifices the "I" of ACID (Isolation)**: Intermediate states are visible externally, so the UI must be designed to show "setting up" states.
- **High cost of compensation logic**: You must design an undo for every step.
- **Hard to handle non-undoable side effects** (sending emails, notifying external APIs, etc.).

### When to Use It

Any business flow that crosses multiple services — tenant onboarding, order processing, payment flows, and so on.

## 3. The Outbox Pattern

A pattern for atomically aligning **"writing to the DB" with "notifying an external system" within the same transaction**. It's not a big flow design like Saga — it's a smaller mechanism for **making the fact of a write and the emission of an event agree**.

### How It Works

1. The app INSERTs a message into an `outbox` table within the same transaction as the business data.
2. The COMMIT atomically finalizes both the business data and the "send reservation."
3. A separate relay process reads the `outbox` via polling or CDC and sends it externally.
4. On successful send, the `outbox` record is marked sent or deleted.

### Pros

- **Messages cannot, in principle, be lost**: If it's written to the DB, it will be sent eventually.
- Simple to implement (just add a table and a relay).
- Easy to plug in as the substrate for inter-step notifications in a Saga.

### Cons

- **At-least-once delivery**: Duplicates can occur; receivers need to be idempotent.
- You need a strategy for `outbox` growth (partitioning, periodic deletion).
- Polling-based designs have latency tied to the poll interval (CDC can bring this under tens of ms).

### When to Use It

Any time you need to reliably notify some external system (Pub/Sub, webhooks, another service's API) about a fact written to your DB. Pairing it with Saga is the standard play.

## How the Three Relate

The three aren't competing technologies; they're **layers at different levels of abstraction**.

```
Saga (the whole business flow)
  └ Outbox (reliable event emission)
     └ DB Transaction (local atomicity)
```

- Saga handles the consistency of the overall business flow.
- Outbox is the foundation that reliably ships events between Saga steps.
- 2PC is the comparison point that views these from the opposite side of the **"strong consistency vs. availability"** trade-off.

## A Selection Guide

| Requirement | Recommended Pattern |
|---|---|
| Consistency of a business flow across services | Saga |
| Atomicity of DB COMMIT and event emission | Outbox |
| Strong consistency is a hard requirement; availability can be sacrificed | 2PC (adopt with care) |
| Single DB is enough | A plain transaction is fine |

In practice, **"Saga + Outbox" as the base set**, with 2PC reserved for cases where there's an explicit reason, is the default stance of modern microservices design.

## Summary

- **2PC**: Strong consistency at the cost of availability. As a rule, avoid it in microservices.
- **Saga**: A business flow design that achieves eventual consistency via compensation. High availability and works with heterogeneous resources.
- **Outbox**: A device for reliably aligning DB writes with external notifications. The foundation for Saga's reliability.

"In a distributed system, give up the Isolation of ACID and recover Atomicity through compensation" — that is the essence of Saga + Outbox. Before reaching for 2PC, start by asking whether this combination can do the job.
