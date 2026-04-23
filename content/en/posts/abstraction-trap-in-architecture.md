---
title: The Abstraction Trap in Architecture Design
description: Why sharing domain logic can quietly turn into architectural debt, while technical concerns can safely be shared — with nuance around authorization and PII-sensitive logging.
slug: abstraction-trap-in-architecture
date: 2026-04-23T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Design
  - DDD
translation_key: abstraction-trap-in-architecture
---


## Introduction

Sharing in architecture design has sharply different effects depending on what we share. Sharing helps reduce duplication and improves maintainability, yet bundling things together does not always pay off. Sharing domain logic in particular can turn a clean-looking codebase into a pile of debt that resists untangling years later.

This article organizes the asymmetry between "safe to share" and "dangerous to share" through the lens of Bounded Contexts.

The discussion here does not focus on microservices. The real question lives at the logical boundary, not the physical one. The same trap appears in modular monoliths and package-level decomposition.

---

## 1. Safe to Share: Technical Concerns

One relatively safe area to share covers **technical concerns**. Examples include:

- Logging
- Metrics collection and distributed tracing
- HTTP infrastructure
- Error notification
- Configuration loading

These share the following properties:

- **They carry no direct business meaning**
- **Change drivers live on the platform side**: library upgrades, SRE requirements, security patches
- **Behavior rarely needs to differ per context**: they work as-is across many business contexts

Extracting them as cross-cutting concerns and letting them evolve independently as a shared library or platform service makes sense. Tolerating too much duplication here tends to hurt maintainability across the system.

---

## 2. The Gray Zone: When Domain Leaks Into Technical Concerns

That said, "technical concerns" can become a coarse label. Separating the truly technical parts from parts where domain knowledge has leaked in matters a lot. Mixing them up leads to a different kind of trap.

### Authentication and Authorization

- **Authentication** (verifying identity): easy to share as a platform capability
- **Authorization** (deciding what a user can do): a dense collection of domain rules

Rules such as "a sales rep can view orders only from their own division" or "a credit officer can view cases only while under review" embody authorization policy that belongs to the domain. Folding this into a single "auth library" forces every domain change to touch the platform and creates a mismatch between platform cadence and business cadence.

### Logging

- **Formatting and transport**: safe to share
- **Which fields qualify as PII to redact, and which events to persist as audit logs**: domain-level judgment

A shared logger should own formatting and transport. The judgment about what to log and how to treat it should stay on the domain side.

Sharing technical concerns makes sense, but missing the domain judgment hiding inside them leads to trouble.

---

## 3. Dangerous to Share: Domain Logic

This article argues most strongly for the following point: **sharing logic that belongs to a specific domain becomes a major architectural trap**.

The same model — "Customer", "Order", "Price" — often shows up across several subsystems with similar structure. That resemblance tempts teams to merge them into a common model or service. Doing so tends to produce friction over time.

To borrow from DDD: when the Bounded Context differs, the thing itself differs.

- **"Customer" in the sales domain**: an entity with buying history and a shipping address
- **"Customer" in the credit domain**: an entity with a credit score and a review state
- **"Customer" in marketing**: an entity with segments and campaign response history

Only the name and the identifier match. Invariants, lifecycle, and the drivers of change all diverge.

---

## 4. Why Sharing Domain Logic Becomes a Trap

Forcing domain logic into a shared abstraction triggers several pushbacks.

### Conflicting invariants across contexts

An attribute that must exist in one context becomes irrelevant in another. A state transition that a context permits turns forbidden in the next. A unified model leaves two options only: obey the strictest constraint everywhere, or patch everything with flags and branches.

### Change velocity bottleneck

A shared component inherits the cadence of the most cautious consumer. Domains that could have evolved independently start blocking each other.

### Lost ownership

Shared-kernel-style domain code (the DDD pattern where several contexts jointly own the same domain model) belongs to no team completely. Everyone can touch it, while accountability tends to blur. By Conway's Law, code that crosses organizational boundaries tends to generate friction.

### Cost of unwinding

Once "the shared domain model" takes root, a great deal of code depends on it. Later attempts to re-split it along contexts turn into an enormous mesh of data migration, API compatibility, testing, and cross-team negotiation.

---

## 5. Questions That Help Separate Safe Sharing From Dangerous Sharing

When deciding whether to share, the following questions help.

- **Where does the reason to change come from?** Platform drivers (library, operations, security), or domain drivers (business rules, trade practices)?
- **Is ownership aligned?** Does a single team and a single context own the code and drive change decisions?
- **Could the contexts diverge in the future?** If evolving business requirements might split them apart, sharing turns into a future constraint.

Structural similarity in code does not justify sharing by itself. What matters is whether the similarity is **incidental** or **essential**.

---

## 6. Separation Is Not Always the Right Call

The arguments above warn against reckless sharing of domain logic, yet they do not imply that teams must always split everything.

In early product phases, or inside a small monolith, pre-splitting for future Bounded Contexts becomes a different flavor of the same trap. When domain contours are still forming, letting the code evolve in one place and carving out boundaries once they become visible tends to work better.

The key is to stop treating "share or not" as a one-off decision. **Revisit the call as organization size, product phase, and domain maturity evolve.**

---

## Conclusion

Sharing looks like a purely technical decision, though in practice it carries judgment about organization and domain.

- Share technical concerns aggressively — but watch for domain logic that has quietly leaked in
- Structural similarity in domain code does not justify sharing. If the Bounded Contexts differ, treat them as different things
- The axis of judgment lives at "**do the reasons to change and the domain context match?**" rather than at "does the code look similar?"

Sharing is not always a virtue. Careless sharing tends to generate a quiet, deep form of architectural debt.
