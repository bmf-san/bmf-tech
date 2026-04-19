---
title: Viability Conditions for Building a Platform
description: How do you decide whether to build an internal platform? This article organizes the viability conditions from the perspectives of core/supporting/generic subdomains, comparison with external services, leverage, and organizational capacity, and examines how far domain analysis alone can take you.
slug: platform-viability-conditions
date: 2026-04-19T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture
  - Platform Engineering
  - DDD
translation_key: platform-viability-conditions
---

# Introduction

Conversations about building a shared internal platform come up often — authentication platforms, payment platforms, notification platforms, internal developer platforms (IDPs), and other cross-domain foundations that several teams rely on.

But when you start building one carelessly, you end up with platforms that "nobody uses," "never recoup their investment," or "nobody can keep running." So how do you decide whether to build a platform in the first place?

This article organizes the viability conditions for internal platforms. It then asks whether domain analysis alone can answer the question — and explores the reach and limits of domain analysis.

---

# Viability Conditions for a Platform

Here are seven perspectives for deciding whether to build one.

## 1. It belongs to a supporting or generic subdomain, not the core domain

The core domain drives business differentiation. Your company should develop it deeply in-house to secure uniqueness and competitive advantage. If you extract it as a platform and generalize it so that other teams can use it, abstraction strips away the company-specific complexity and nuance, which weakens the source of differentiation.

By contrast, supporting and generic subdomains do not contribute directly to differentiation but still need to support several core domains. These areas become good candidates for extraction as a platform.

## 2. It has unique requirements that external services cannot satisfy

For areas like authentication, payments, and notifications, plenty of SaaS or OSS options already exist. If an external service covers your needs, there is no reason to build your own platform.

Building your own platform is only worth it when you have unique requirements that external services cannot meet — deep integration with internal systems, specific business workflows, regulatory or data sovereignty constraints, connectivity with existing assets, and so on.

## 3. The TCO of in-house development and operation is lower than using an external service

Even if you have unique requirements, when the total cost of ownership (TCO) is lower with an external service, using one is the rational choice.

In-house becomes cheaper when the scale of usage crosses a certain threshold and you can spread development and operational costs across the number of consuming teams. Usage-based pricing for external services grows linearly with scale, while an internal platform spreads common investment across many teams, so the relative cost goes down as scale grows.

## 4. Several teams or several products will reuse it

A platform is a foundation whose value comes from adoption. The fewer teams that adopt it, the worse the return on investment. If only one team uses it, a shared module within that team suffices — there is no need to carve it out as a platform.

Platform leverage only kicks in when you have a realistic prospect of several teams or several products reusing it.

## 5. The organization has reached a certain size and cognitive load

A platform also serves as a way to reduce cognitive load so that developers can focus on essential problems. Conversely, in a small organization where cognitive load has not yet become a real problem, a platform offers little value.

Its value becomes clear once the organization has grown to the point where each team keeps getting pulled into the details of infrastructure and cross-cutting concerns.

## 6. There is a realistic prospect of continuous investment and ownership

A platform is not a build-and-forget artifact. You must keep improving and maintaining it as consuming teams evolve. A platform that lacks a clear owner and exists only as a side project tends to go stale and turn into technical debt.

At the planning stage, you need to verify whether you can secure continuous investment and assign dedicated ownership.

## 7. The domain is stable enough to provide a stable interface

The value of a platform lies in offering a stable interface that consuming teams can rely on. If the underlying domain keeps changing at a fundamental level, extracting it into a platform only ends up dragging consumers around.

There needs to be a reasonable expectation that the domain is mature enough to withstand generalization and abstraction.

---

# Can Domain Analysis Alone Decide This?

Let us sort the seven perspectives above by whether domain analysis alone can judge them.

**Decidable by domain analysis**

- 1. Classification of core / supporting / generic — the central output of subdomain analysis
- 2. Presence of unique requirements — surfaces through use-case analysis and ubiquitous language
- 7. Domain stability — visible from modeling and understanding of existing work

**Outside the reach of domain analysis**

- 3. TCO comparison — requires financial judgment and usage estimates
- 4. Reuse across several teams — a matter of portfolio and product strategy
- 5. Organizational size and cognitive load — depends on the organizational context
- 6. Continuous investment and ownership — a matter of organizational design and staffing

Domain analysis tells you "which areas you may build" (i.e., what you should put your hands on), but it does not answer "whether the work pays off" (i.e., whether the investment will pay back and whether you can keep running it). In other words, domain analysis provides the **necessary conditions**, and the **sufficient conditions** come together only when you combine strategic, financial, and organizational judgment.

The rule is not "build because the subdomain is supporting or generic." The rule reads "build when the subdomain is supporting or generic **and** the investment and organizational conditions are also in place."

---

# Conclusion

The viability conditions for a platform consist of three layers of judgment:

- **Domain analysis** (necessary conditions): classification of core / supporting / generic, unique requirements, domain stability
- **Strategic and financial judgment**: TCO, leverage, portfolio positioning
- **Organizational judgment**: organizational size and cognitive load, ownership and continuous investment

When you plan a platform, do not jump to "we should build it" based on domain analysis alone — verify that the investment and organizational conditions are also in place. A platform only earns its place and takes root in the organization when you meet every one of these viability conditions.
