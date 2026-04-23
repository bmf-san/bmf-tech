---
title: The Abstraction Trap in Architecture Design
description: Why deduplicating domain logic is architecturally dangerous, while technical concerns can be safely shared — with nuance around authorization and PII-sensitive logging.
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


# Overview

In architecture design, some forms of sharing are perfectly safe to pursue, while others quietly turn into major traps.

**Technical concerns** — logging, monitoring, authentication infrastructure, and the like — are usually worth sharing aggressively. On the other hand, sharing **logic that belongs to a specific domain**, just because it "looks similar", tends to create a trap that is hard to pull apart later. This article examines that asymmetry through the lens of Bounded Contexts.

Note that this is not a microservices-specific discussion. The real question is not about **physical boundaries (service split)** but about **logical boundaries (Bounded Contexts)**. The same trap occurs in modular monoliths and package-level decomposition just as readily.

# Technical Concerns Are Safe to Share

There is an area where sharing works relatively well: **technical concerns** such as logging, metrics, distributed tracing, HTTP infrastructure, error notification, and configuration loading.

These share some useful properties:

- They do not directly carry business meaning
- Their change drivers live on the platform side (library upgrades, SRE requirements, security patches)
- Even when used from many business contexts, their behavior rarely needs to differ per context

Extracting them as cross-cutting concerns into a shared library or platform service — and letting them evolve independently — is generally reasonable. Tolerating too much duplication here tends to hurt overall maintainability.

# When Domain Leaks Into Technical Concerns

That said, "technical concerns" is a coarse label. If we fail to separate the truly technical parts from the parts where domain knowledge has quietly leaked in, we fall into a different kind of trap.

A classic example is **authentication and authorization**. Authentication (verifying identity) is reasonable to share as a platform capability. Authorization (deciding what a given user is allowed to do), on the other hand, tends to be a dense collection of domain rules: "A sales rep can only view orders belonging to their own division." "A credit officer can only access cases currently in approval flow." If you stuff all of this into a single "auth library," you end up touching the platform every time the domain changes, creating a mismatch between platform cadence and business cadence.

Logging has the same pattern. "Which fields count as PII and must be masked?" "Which events must be persisted as audit logs?" — these are domain questions. A shared logger should only own formatting and transport; **the decision of what to log and how to treat it should stay on the domain side**.

It is fine, even desirable, to share technical concerns aggressively — but **do not miss the domain judgment hiding inside them**.

# Domain Logic Should Not Be Shared

Here is the point this article wants to emphasize most: **sharing logic that belongs to a specific domain is a major architectural trap**.

It is common for the same model — "Customer", "Order", "Price" — to appear across several subsystems, looking structurally similar. That resemblance tempts us into merging them into a common model or a common service. More often than not, that tends to create friction over time.

Borrowing DDD's vocabulary: even if the word is identical, **when the Bounded Context differs, it is a different thing**. In the sales domain, "Customer" may be an entity with a purchase history and a shipping address. In the credit domain, "Customer" is an entity with a credit score and a review status. In marketing, "Customer" carries segments and campaign responses. Only the name and identifier are the same. The invariants, the lifecycle, and the reasons to change all differ.

# Why Sharing Domain Logic Becomes a Trap

When domain logic is forced into a shared abstraction, several forces push back:

- **Conflicting invariants across contexts**: An attribute that must be present in one context is irrelevant in another. A state transition that is permitted in one is forbidden in another. A unified model forces you to either obey the strictest constraint everywhere, or patch everything with flags and branches.
- **Change velocity bottleneck**: A shared component is dragged down to the cadence of its most cautious consumer. Domains that could have evolved independently start blocking each other.
- **Lost ownership**: Shared-kernel-style domain code (the DDD pattern where multiple contexts jointly own the same domain model) belongs to no team completely. Everyone can touch it; accountability tends to blur. By Conway's Law, code that crosses organizational boundaries tends to generate friction.
- **Cost of unwinding**: Once "the shared domain model" has taken root, a great deal of code depends on it. Later attempts to re-split it along contexts turn into an enormous mesh of data migration, API compatibility, testing, and cross-team negotiation.

# Questions That Help Separate Safe Sharing From Dangerous Sharing

When deciding whether to share, the following questions are useful:

- **Where does the reason to change come from?** Does change originate on the platform side (library, operations, security), or on the domain side (business rules, trade practices)?
- **Is ownership aligned?** Is the code owned, and are change decisions made, by a single team and a single context?
- **Could the contexts diverge in the future?** Even if they look identical today, if evolving business requirements might split them apart, sharing is likely to become a future constraint.

Structural similarity in code is not, by itself, a valid reason to share. The real question is whether that similarity is **incidental** or **essential**.

# Separation Is Not Always the Right Call

Everything above argues against reckless sharing of domain logic, but that is not the same as saying things must always be split.

In early product phases, or inside a small monolith, trying to predict future Bounded Contexts and pre-splitting everything is a different flavor of the same trap. When domain contours are still forming, it is often safer to let the code evolve in one place and carve out boundaries only once they become visible.

The important point is to stop treating the "share / don't share" decision as a one-off. It should be **revisited continuously as organization size, product phase, and domain maturity evolve**.

# Summary

Sharing looks like a purely technical decision, but it is not. It carries judgment about organization and domain.

- Share technical concerns aggressively — but watch for domain logic that has quietly leaked in
- Structural similarity in domain code is not enough justification to share. If the Bounded Contexts differ, treat them as different things
- The real axis of judgment is not "does the code look similar?" but "**are the reasons to change and the domain context the same?**"

Sharing is not always a virtue. Careless sharing tends to generate a quiet, deep form of architectural debt.
