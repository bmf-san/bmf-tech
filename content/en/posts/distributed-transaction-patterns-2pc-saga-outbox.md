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
- **Poor fit with heterogeneous resources**: Participants must be 2PC-capable resource managers like XA; Pub/Sub, Kafka, HTTP APIs, etc. are typically outside that scope (Kafka has its own transaction API, but it is distinct from XA).
- **At odds with the design philosophy of microservices**: "We split the services, but one failure halts the whole thing" is self-defeating.

### From an ACID Perspective

| ACID | What 2PC gives you |
|---|---|
| **A**tomicity | ◎ All participants either COMMIT or ROLLBACK together |
| **C**onsistency | ◎ All resources are aligned the moment COMMIT completes (immediate consistency) |
| **I**solation | ◎ Locks held during Prepare isolate the transaction from others |
| **D**urability | ◎ Each participant durably persists the result after COMMIT |

On paper this looks perfect — the price you pay is **availability**.

### When to Use It

Almost never. **Between microservices, the rule of thumb is to avoid it.** Consider it only in narrow, legacy-leaning cases where you need to align heterogeneous resources (multiple DBs, DB + MQ, etc.) under an XA-style transaction manager.

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

### From an ACID Perspective

| ACID | What Saga gives you |
|---|---|
| **A**tomicity | ○ Eventually preserved as either "all steps succeeded" or "compensated to a business-equivalent of nothing happened" |
| **C**onsistency | ○ Each step commits locally; the whole flow is eventually consistent |
| **I**solation | ✕ **This is what Saga gives up.** Intermediate states are visible to other services and users |
| **D**urability | ◎ Each step's local COMMIT is durable |

**Saga gives up the "I" of ACID and recovers "A" through compensation** — that is its essence.

### When to Use It

Any business flow that crosses multiple services — tenant onboarding, order processing, payment flows, and so on.

## 3. The Outbox Pattern

A pattern for reliably aligning **"writing the business data" with "notifying an external system"** so that the two cannot drift apart. It's not a big flow design like Saga — it's a smaller mechanism for **making the fact of a write and the emission of an event agree**.

### The Problem It Solves: dual write

The naive version looks like this:

```
BEGIN;
INSERT INTO orders ...;   -- business data
COMMIT;                    -- the DB is durable here
queue.publish(event);      -- a call into a different system
```

A DB transaction can only bundle operations **inside the same DB**. `queue.publish` lives outside the DB, so:

- If the app dies right after COMMIT, the publish never happens → **lost event**.
- If you publish first and the COMMIT fails, the notification has fired but the business data doesn't exist → **phantom event**.

**No matter which order you choose, you cannot escape the mismatch.** This is the **dual write problem**.

### How It Works: write a "send reservation" to a table in the same DB

Instead of publishing directly, the Outbox pattern **INSERTs into an `outbox` table inside the same DB**, in the same transaction:

```sql
BEGIN;
INSERT INTO orders ...;    -- business data
INSERT INTO outbox ...;    -- "please publish this later"
COMMIT;                     -- one COMMIT, both atomic
```

1. The app INSERTs a message into the `outbox` table within the same transaction as the business data.
2. The COMMIT atomically finalizes both the business data and the "send reservation" in a single COMMIT.
3. A separate relay process reads the `outbox` via polling or CDC and sends it externally.
4. On successful send, the `outbox` record is marked as sent (e.g., flipping a status column) or deleted.

The key point is that **the outbox must live somewhere that can be committed together with the business data in a single transaction**. In practice that's almost always "a table in the same DB as the business data" (Kafka transactions or CDC-based variants exist, but the essence is the same boundary).

### Pros

- **Events cannot, in principle, be lost**: if it made it to the DB, it will be sent eventually.
- Simple to implement (just add a table and a relay).
- Easy to plug in as the substrate for inter-step notifications in a Saga.

### Cons

- **At-least-once delivery**: duplicates can occur; receivers need to be idempotent.
- Weak ordering guarantees (parallel relays can reorder).
- You need a strategy for `outbox` growth (partitioning, periodic deletion).
- Polling-based designs have latency tied to the poll interval (CDC can bring this under tens of ms).

### From an ACID Perspective

| ACID | What Outbox gives you |
|---|---|
| **A**tomicity | ◎ **This is the core.** Business data and the "send reservation" either both land or both don't, in one COMMIT |
| **C**onsistency | ○ "The business data was written" ⇔ "an event will eventually fire" is preserved (eventual consistency) |
| **I**solation | △ The INSERTs themselves are isolated inside the DB, but external systems can observe the intermediate state (DB written, event not yet delivered) |
| **D**urability | ◎ Once written, it stays. No matter how many times the relay dies, the `outbox` row survives and is eventually published |

The only genuinely *new* guarantee Outbox introduces is **Atomicity**. Consistency and Durability follow as side effects on top of it.

### Where Outbox Sits on the "Looseness" Spectrum

Outbox is not a strong-consistency device — it's an **eventual-consistency** design. The trade-offs line up like this:

| Requirement | What to use |
|---|---|
| It's OK to lose events (logs, metrics, best-effort) | Publish right after COMMIT |
| You can't lose events, but duplicates / delay / reordering are acceptable | **Outbox** |
| You need strong, immediate consistency | 2PC (adopt with care) |

Outbox is best understood not as a tool for strong consistency, but as **a simple pattern focused on one job: avoiding the dual write problem**.

### When to Use It

Any time you need to **losslessly** notify some external system (Pub/Sub, webhooks, another service's API) about a fact written to your DB. Pairing it with Saga is the standard play.

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

In practice, **placing Saga and Outbox at the center** and reserving 2PC for cases where there's an explicit reason is the stance widely adopted in modern microservices design.

## Summary

- **2PC**: Strong consistency at the cost of availability. As a rule, avoid it in microservices.
- **Saga**: A business flow design that achieves eventual consistency via compensation. High availability and works with heterogeneous resources.
- **Outbox**: A device for reliably aligning DB writes with external notifications. The foundation for Saga's reliability.

"In a distributed system, give up the Isolation of ACID and recover Atomicity through compensation" — that is the essence of **Saga**, and Outbox is the foundation that reliably drives each of its steps. Before reaching for 2PC, start by asking whether this combination can do the job.
