---
title: Accidental Complexity and the Absence of Decision Records
description: How accidental complexity emerges, and why the absence of ADRs and Design Docs makes it hard to tell what is essential from what is accidental.
slug: accidental-complexity-and-missing-decision-records
date: 2026-05-21T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Design
  - Architecture Decision Record
  - Design Docs
translation_key: accidental-complexity-and-missing-decision-records
---

## Introduction

When you read a long-running system, you eventually run into the question: "Is this complexity actually necessary, or is it a side effect of something?" Software complexity often splits into **essential complexity** — coming from the problem itself, and unavoidable no matter how you build it — and **accidental complexity** — the kind that just bleeds out of implementation choices and historical circumstances, and is removable in principle.

Some of these are easy to tell apart. Others, you can't tell from code or specs alone. This post organizes the typical sources of accidental complexity, and argues that **the lack of recorded decisions is itself one of the things that makes essential and accidental hard to separate**.

---

## Where Does Accidental Complexity Come From

Accidental complexity does not have a single origin. Looking at it through the lens of **"what gives rise to it"**, it roughly splits into three.

### Business Requirements and Operational Convenience

Things that were "decided that way for business reasons." They fit stakeholders or operational convenience at the time, but they aren't necessarily universal needs.

- A particular screen has a special permission check that exists nowhere else
- An aggregation slices data as "the min across the whole" or "sum over the last 30 days" — granularities tailored to a specific organization's operations
- Features A and B run in sequence by design, and using only one of them isn't a supported case

These don't disappear when you move to a new architecture, unless you re-negotiate with the stakeholders. They are "redesignable" technically, but in practice **they only disappear after a human-side agreement**.

### Technology Choices and Implementation Means

Complexity that bleeds out of the framework, middleware, or protocol you picked. Common shapes:

- Subdomain-by-subdomain separate sessions, forced by the cookie domain rules of the auth scheme
- "Soft-delete column + unique constraint" combinations contorted to fit ORM limitations
- A job queue that quietly grew beyond queueing, to compensate for a single-process-bound framework

These have a high chance of **disappearing structurally if you change the architectural choice**. Switching to a different protocol, adopting a different middleware, or rewriting in a different language can make "this problem no longer exists" a real outcome.

### Organizational and Time Constraints

Not from business requirements, not from technical limitations, but from **"this was the only way it could be written at the time."** It bleeds out of deadlines, headcount, the org structure of the day, and historical circumstance. Of the three, it's the one that gets overlooked the most.

- An implementation written on the safe side because the team at the time lacked domain knowledge in this area
- A workaround you keep on your side because you couldn't touch code owned by another team
- A last-minute fix made just before release that quietly settled in as the permanent implementation
- A "two organizations merged and both concepts stuck around" data model

These can look like business requirements, or like technical constraints. But the real origin is **the organization and time constraints of that moment**, and they don't go away by changing technology choices or re-negotiating business requirements. To remove them, someone has to make the call: **"those constraints don't exist anymore, so let's rewrite this."**

---

UX-level distortions (having to log in repeatedly, double notifications, the context breaking mid-flow) are **observed phenomena** of the three above, not independent causes. Fixing the cause axis usually fixes the UX as a side effect, so this post focuses on causes.

---

## Why Telling Essential From Accidental Is Hard

"Essential or accidental" doesn't fall out of code alone. Without knowing why things are the way they are, you can't tell whether a structure in front of you is "this has to be this way for business reasons" or "this is just bent to fit the framework constraints of the time."

Separating them requires **the context of the decisions made at the time**. Concretely, things like:

- What options sat on the table
- Why we picked this option (performance? schedule? compatibility with existing assets? ease of hiring?)
- Which constraints we accepted
- Which items we deferred as "we might revisit later"

If those records exist, a future designer can say, "this complexity comes from constraint X at the time. X no longer applies now, so we can throw the whole structure out." If they don't, every structure becomes a black box of "there must be a reason," and stays on the safe side.

---

## Three Problems Caused by the Absence of Decision Records

### Accidental Gets Promoted to Essential

The most typical failure: an accidental structure starts looking like an essential trait over time.

> "That's the spec."

Much of the behavior described that way is, in reality, just accidental complexity that fit the constraints at initial implementation. But without records, you can't refute it, so when you build a new system, "please preserve the same behavior as the current one" gets baked in as a spec. As a result, the next generation system inherits the same accidental complexity.

### The "Don't Touch" Zone Expands

Without recorded decisions, a psychological wall builds up around the code: **"I don't know why it's like this, but changing it looks like it'll break something."** This is close to Chesterton's Fence — if you don't know why the fence is there, the safe move is to leave it.

That's reasonable in the short term, but if you keep doing it for 5 or 10 years, the system becomes covered in "do not touch" zones, and the area where refactoring is possible disappears.

### Discussions Collapse Into "Match the Current Behavior or Not"

When designing a new system, the conversation tends to center on "do we replicate the current behavior or not." What you should actually discuss is "what business properties do we want to guarantee," but without primary sources for that question (the original decisions), copying the current behavior becomes the safest move.

The result: the new system gets close to **"the same accidental complexity as before, just on a newer tech stack."** Which makes the rewrite barely worth doing.

---

## ADRs and Design Docs as "Aids for Your Future Self"

There's no silver bullet here. But leaving an ADR or Design Doc at each design decision — capturing **the options at the time, the reason for the choice, and the constraints accepted** — is one of the highest-leverage countermeasures.

ADRs and Design Docs are not about "recording the right answer." Their value is closer to:

- **Marking accidental as accidental**
  Writing "constraint X shaped this structure" gives a future reader the basis to throw the structure away once X no longer applies.
- **Pinning down what's essential**
  Writing "this property is non-negotiable from the business side" lets a new system inherit it as something that must persist.
- **Forwarding deferred options to the future**
  Writing "we didn't take this option this time, but we should revisit if conditions change" lets a future generation reevaluate.

Conversely, not writing ADRs or Design Docs is close to **stripping your future self (or successor) of the means to separate accidental from essential**.

---

## So How Much Should You Write

"Record every decision" isn't realistic. Writing ADRs and Design Docs has a cost, and so does maintaining them. It's enough to focus on things like:

- **Decisions likely to be questioned later as "why is it this way"** (especially when picking a non-obvious option from several alternatives)
- **Decisions where it's ambiguous whether they came from business needs or from technical constraints**
- **Decisions likely to need revisiting when assumptions change** (framework refresh, scale change, organizational change)

Decisions that obviously apply best practice, or that anyone would have made the same way, don't need an ADR. The principle is: **"what you should record is the decision, not the result."**

---

## Summary

- Software complexity splits into essential and accidental, but some pairs are easy to tell apart and others aren't.
- Accidental complexity comes from three causes — business needs & operational convenience / technology choices & implementation means / organizational & time constraints — and each disappears under different conditions.
- Separating them requires the context of past decisions, but that context is often missing.
- As a result, accidental complexity gets promoted to essential and survives into the new system.
- ADRs and Design Docs are not about recording the right answer; they work as **aids for passing accidental as accidental, and essential as essential, on to the next generation**.
