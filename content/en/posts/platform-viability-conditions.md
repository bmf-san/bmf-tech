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

Conversations about building a shared internal platform come up often — authentication platforms, payment platforms, notification platforms, internal developer platforms (IDPs), and other cross-domain foundations used by multiple teams.

But when you start building one carelessly, you end up with platforms that "nobody uses," "never recoup their investment," or "can't be sustained." So how do you decide whether to build a platform in the first place?

This article organizes the viability conditions for internal platforms. It then asks the question: "Can this decision be made through domain analysis alone?" — and explores the reach and limits of domain analysis.

---

# Viability Conditions for a Platform

Here are seven perspectives for deciding whether to build one.

## 1. It belongs to a supporting or generic subdomain, not the core domain

The core domain is the area that drives business differentiation. It should be built deeply in-house to secure uniqueness and competitive advantage. If you extract it as a platform and generalize it so that other teams can use it, the company-specific complexity and nuance get stripped away through abstraction, and the source of differentiation is diluted.

Supporting and generic subdomains, on the other hand, don't directly contribute to differentiation but are necessary to support multiple core domains. These are the areas that become candidates for extraction as a platform.

## 2. It has unique requirements that external services cannot satisfy

For areas like authentication, payments, and notifications, there are often plenty of SaaS or OSS options. If an external service is sufficient, there is no need to build your own platform.

Building your own platform is only justified when you have unique requirements that external services cannot meet — deep integration with internal systems, specific business workflows, regulatory or data sovereignty constraints, connectivity with existing assets, and so on.

## 3. The TCO of in-house development and operation is lower than using an external service

Even if you have unique requirements, if the total cost of ownership (TCO) is lower with an external service, using one is the rational choice.

In-house becomes cheaper when the scale of usage crosses a certain threshold and development and operational costs can be amortized across the number of consuming teams. While usage-based pricing for external services grows linearly with scale, an internal platform can spread common investment across multiple teams, so the relative cost goes down as scale grows.

## 4. It will be reused by multiple teams or multiple products

A platform is a foundation whose value is realized "by being adopted." The fewer teams that adopt it, the worse the return on investment. If only one team uses it, a shared module within that team is enough — there's no need to carve it out as a platform.

Platform leverage only kicks in when there is a realistic prospect of multiple teams or multiple products reusing it.

## 5. The organization has reached a certain size and cognitive load

A platform is also a mechanism for reducing cognitive load so that developers can focus on essential problems. Conversely, in a small organization where cognitive load hasn't yet become a real problem, the value of a platform is limited.

Its value becomes apparent once the organization has grown to the point where each team is getting pulled into the details of infrastructure and cross-cutting concerns.

## 6. There is a realistic prospect of continuous investment and ownership

A platform is not a build-and-forget artifact. It must be continuously improved and maintained as consuming teams evolve. A platform built without a clear owner and only as a side project tends to go stale and turn into technical debt.

At the planning stage, you need to verify whether you can secure continuous investment and assign dedicated ownership.

## 7. The domain is stable enough to provide a stable interface

The value of a platform lies in offering a stable interface that consuming teams can rely on. If the underlying domain keeps getting overturned at a fundamental level, extracting it into a platform only ends up dragging consumers around.

There needs to be a reasonable expectation that the domain is mature enough to withstand generalization and abstraction.

---

# Can This Be Decided Through Domain Analysis Alone?

Let's sort the seven perspectives above by whether domain analysis alone can judge them.

**Decidable by domain analysis**

- 1. Classification of core / supporting / generic — the central output of subdomain analysis
- 2. Presence of unique requirements — surfaces through use-case analysis and ubiquitous language
- 7. Domain stability — visible from modeling and understanding of existing work

**Outside the reach of domain analysis**

- 3. TCO comparison — requires financial judgment and usage estimates
- 4. Reuse across multiple teams — a matter of portfolio and product strategy
- 5. Organizational size and cognitive load — depends on the organizational context
- 6. Continuous investment and ownership — a matter of organizational design and staffing

Domain analysis tells you "which areas you may build" (i.e., what you should put your hands on), but it does not answer "whether it's worth building" (i.e., whether the investment can be recouped and sustained). In other words, domain analysis provides the **necessary conditions**, and the **sufficient conditions** come together only when combined with strategic, financial, and organizational judgment.

The rule isn't "build it because it's a supporting or generic subdomain." It's "build it when it's a supporting or generic subdomain **and** the investment and organizational conditions are also in place."

---

# Conclusion

The viability conditions for a platform are made up of three layers of judgment:

- **Domain analysis** (necessary conditions): classification of core / supporting / generic, unique requirements, domain stability
- **Strategic and financial judgment**: TCO, leverage, portfolio positioning
- **Organizational judgment**: organizational size and cognitive load, ownership and continuous investment

When you're planning a platform, don't jump to "we should build it" based on domain analysis alone — verify that the investment and organizational conditions are also in place. A platform only has meaning and takes root in the organization when all of the viability conditions are met.
