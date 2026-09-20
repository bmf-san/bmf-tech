---
title: Why the Domain Fades From View in Web Development, and Why It Eventually Comes Back
description: "Understand what a domain really means, why it works as an architectural boundary, and the structural reasons the domain fades from view in web development."
slug: why-domain-is-overlooked-in-web-development
date: 2026-08-21T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - DDD
  - Design
  - Architecture
translation_key: why-domain-is-overlooked-in-web-development
draft: false
---

## Introduction

Adopting microservices or a modular monolith forces one question: where do we cut? Conventional wisdom answers with domains.

Real projects, though, often draw boundaries by use case, by screen, or along whatever shape the existing tables happen to have. Such boundaries warp over time, wrecking the data model and inflating service-to-service chatter.

So why should a domain drive the boundary in the first place? What does a domain mean, and what makes it work as a boundary? The phrase "split by domain" circulates widely while that premise stays unspoken.

This article covers three things:

- What a domain means
- Why a domain works as an architectural boundary
- Why the domain tends to fade from view in web development

## What Does a Domain Actually Mean?

### Real-World Activity, Not Technology

A domain refers to the area of real-world activity that the software addresses. It covers business rules, the vocabulary that stakeholders use, the constraints they must respect, and the units of decision-making.

The key point: a domain is neither a database table nor an API endpoint. A table design expresses a domain, but it never equals one.

### "Product" Is Never a Single Concept

Take the word "product" as an example.

- Sales cares about price, campaigns, and display eligibility
- Logistics cares about weight, dimensions, and warehouse location
- Procurement cares about suppliers, cost, and lead time

The same word appears everywhere, yet the concerns barely overlap. Sales can change a price without disturbing logistics, and logistics can move a shelf without disturbing sales.

Paying attention to the domain means copying the boundaries of real-world activity into the boundaries of your code. Domain-driven design calls this a bounded context. For more detail, see [What Is a Bounded Context?](/posts/bounded-context-explanation/).

## Why Designing Around Domains Works

### Reason 1: Reasons for Change Line Up

Screen layouts and marketing campaigns change constantly. Core concepts such as "what an order is" or "what inventory means" barely move.

Splitting by use case forces things with wildly different change rates to live together. A campaign rule that disappears in six months sits next to an order rule that survives a decade, and every edit to one endangers the other.

The single responsibility principle really says that a module should answer to one actor. Domain boundaries align the actors who request changes, so the principle holds naturally at the module level.

### Reason 2: Vocabulary and Data Models Stay Coherent

Without boundaries, one "product" table slowly collects every attribute from sales, logistics, and procurement. That produces the classic god entity.

Columns keep piling up, and the knowledge of who owns which attribute disappears. Nullable columns multiply, implicit rules like "only valid under this condition" scatter across the codebase, and eventually nobody can explain the whole table.

Draw a boundary per domain and each word carries exactly one meaning inside it. You translate only when you cross a boundary, and the translation has a clear owner.

Sharing code is not the villain here. Technical concerns and domain logic respond very differently to consolidation. I covered that asymmetry in [The Abstraction Trap in Architecture](/posts/abstraction-trap-in-architecture/).

### Reason 3: Consistency Boundaries Close Cleanly

Rules that demand strong consistency usually stay inside a single domain. "The sum of order lines must match the invoice amount" never leaves the order domain.

Meanwhile, a rule such as "create a shipping instruction once a customer confirms an order" tolerates a delay. A few seconds of lag breaks nothing.

A domain boundary thus doubles as a candidate transaction boundary. Keep strong consistency inside, use asynchronous messages and eventual consistency across, and you gain loose coupling without giving up scale.

Split by screen instead, and a single click demands synchronous writes to several services. You land squarely in distributed transactions, the one problem everybody wants to avoid.

## Why the Domain Fades From View in Web Development

### Cause 1: The Success of CRUD and Active Record

Frameworks such as Ruby on Rails and Laravel mapped database tables directly onto models and made development dramatically faster. Write a migration, get a model; run a scaffold, get list, detail, and edit screens.

That was a success, not a mistake. This productivity fueled the explosive growth of the web.

The trouble is that the success invites an illusion: the database schema is the domain. A schema reflects persistence concerns above all. Normalization, indices, and foreign keys serve storage and retrieval, not the business.

If a product grows while that illusion holds, concepts that the business treats as entirely separate end up fused, for no better reason than sharing a table.

### Cause 2: A Development Process That Starts From Screens

Most web work begins with mockups and user stories. "What happens when I press this button on this screen?" stays concrete enough that every stakeholder can join the discussion.

Code structure then follows the screens. A controller per screen, an API per screen, a module per screen — every split looks natural.

But screens form the most volatile layer of any system. They shift with every marketing push and A/B test. Choose the most volatile thing as your boundary criterion and your boundaries will break just as often.

### Cause 3: Uncertainty and Speed in Web Businesses

A web business cannot know what works until it reaches product-market fit. Most of what you build may end up on the scrap heap.

Investing heavily in domain modeling at that stage risks throwing away the model along with the feature. Skipping deep modeling early on counts as a reasonable call in most cases.

The real problem lies elsewhere: nobody schedules a review of that call. The assumption "speed comes first for now" outlives the phase that justified it.

### Cause 4: Modeling Itself Is Hard, With No Right Answer

Library usage has correct answers, and the documentation spells them out. Domain modeling offers no such comfort.

You have to pull tacit knowledge out of the heads of domain experts, reconcile conflicting interpretations, and reach agreement. That leans on conversation and synthesis far more than on technical skill.

The payoff also hides. A good model pays dividends later while producing nothing that runs today, and work that nobody rewards drifts to the bottom of the backlog.

Techniques such as [event storming](/posts/event-storming-introduction/) lower that barrier by gathering stakeholders and making the flow of events visible.

## How to Approach It

### Skip the Perfect Model at the Start

Early on, CRUD and Active Record serve you well. Domain knowledge accumulates only through building and operating the thing.

What matters is agreeing in advance on the trigger for refactoring. Watch for signals like these:

- An entity of a single name starts collecting attributes from clearly different contexts
- One feature addition keeps breaking code that should have nothing to do with it
- The team grows, and several teams write to the same table
- Only one person can still explain what the table means

### Watch Data Lifecycles Instead of Screens

When you hunt for boundaries, ask who creates a piece of data, who changes it, and when it retires — not which screen shows it. A cluster where creation and modification share one owner makes a strong domain candidate.

Use cases exist to test a boundary, not to decide it. Trace your main use cases; if they bounce across a boundary again and again, question where you cut.

### Align Team Boundaries With Code Boundaries

Conway's law tells us that system structure mirrors organizational communication. When domain boundaries and team boundaries diverge, one of them warps.

A domain boundary also serves to lower cognitive load on a team, not just to fit the system. A domain too large for one team to hold cannot work as a unit of decomposition at all.

## Summary

- A domain means the concerns, vocabulary, and rules of real-world work, not technology
- Domains work as architectural boundaries because reasons for change, vocabulary, and consistency ranges all line up
- The triumph of CRUD culture, screen-first development, business uncertainty, and the difficulty of modeling have kept domains out of view in web development
- Favor speed early, but decide up front what will trigger the shift to modeling

CRUD and screen-first development powered the growth of the web, and they were right for that stage. Once a product grows complex and the team multiplies, returning to the domain becomes the key to sustainable development.

## References

- [Domain-Driven Design: Tackling Complexity in the Heart of Software](https://amzn.to/4bgaiEm)
- [Learning Domain-Driven Design](https://amzn.to/4rE95PF)
- [Team Topologies: Organizing Business and Technology Teams for Fast Flow](https://amzn.to/3Qoj2j2)
