---
title: What is a Bounded Context?
slug: bounded-context-explanation
date: 2025-05-19T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - DDD
description: Explaining the concept of bounded contexts in a way that is understandable even to non-developers.
translation_key: bounded-context-explanation
---

# Overview

There was a challenge in explaining why contexts are divided the way they are on a context map, and what the significance of these divisions is, to people outside of the development team.

This post attempts to explain "bounded contexts" in a way that is understandable even to non-developers.

# Being Aware of Context Differences

Even if the same word is used, its meaning can change depending on the context.

For example, consider the word "order."

In the sales department, it might refer to "a request from a customer," in the warehouse, it could mean "a shipping instruction," and in the accounting department, it might mean "data for billing."

Other examples include:

* **"User"**: For the development team, it refers to someone who logs in; for customer support, it means someone making inquiries; for marketing, it refers to target customers.
* **"Service"**: For engineers, it means an API; for sales, it refers to customer-facing plans.

In this way, **the same word can have different meanings depending on the business or perspective.**

The concept of "context" helps clarify these differences.

# What is a Context?

A context is **a cohesive unit of business where the meaning of words and rules remains consistent.**

In Domain-Driven Design (DDD), this is referred to as a "Bounded Context."

For example, the word "order" can have different meanings depending on the context:

* In Context A, it means a customer's purchase intent.
* In Context B, it means a shipping request to the warehouse.
* In Context C, it refers to billing targets.

The characteristic of a context is that it "divides" and handles the differences in the meaning of words and rules for each business area.

## Why Organize Contexts?

If you proceed with business operations or system design without being aware of contexts, various problems can arise:

* Overloading the "order" table with billing and shipping information, resulting in a complex and fragile system.
* Teams frequently asking, "What does this mean?" leading to misunderstandings and unclear responsibilities.
* Difficulty in improving processes due to an inability to predict the scope of impact.

On the other hand, organizing and separating contexts provides the following benefits:

* Clarifies the meaning of words.
* Organizes the scope of responsibility for business operations and systems.
* Makes it easier to divide or modify team responsibilities.

In short, properly organizing contexts allows systems and organizations to enjoy the following advantages:

* **Reduced communication costs**
* **Lower complexity in system design and implementation**
* **Clearer role definitions and responsibilities between teams**

Properly organized contexts contribute to **improved stability for both systems and organizations.**

# Example of Contexts: "Order" in an E-commerce Site

Let’s divide the business operations related to orders in an e-commerce site into three departments:

| Department       | Meaning of "Order"            |
| ---------------- | ------------------------------ |
| Frontend Website | Content confirmed in the cart  |
| Warehouse        | Data for shipping instructions |
| Accounting       | Sales data for billing         |

All of these use the word "order," but **the content, purpose, and processes involved are completely different.**

If these differences are ignored and everything is lumped under a single term:

* The system becomes more complex.
* The scope of impact during changes becomes unclear.
* Misunderstandings are more likely to occur.

This is why it is important to separate and consider "contexts."

# Not Just for Developers

Contexts are not something only developers need to understand.

For example, have you ever encountered questions like these?

* "What exactly is our team responsible for?"
* "Should I ask sales, development, or support about this?"
* "What is the impact of this change?"

All of these questions are closely related to the organization of contexts.

**By drawing boundaries where business purposes and the meanings of words change, and clarifying roles and responsibilities within those boundaries,**

this perspective is beneficial not only for developers but also for sales, support, planning, and other roles.

# Conclusion

* The same word can often have different meanings.
* Recognizing these differences and organizing them as boundaries in business and systems is what "context" is about.
* Clarifying contexts prevents misunderstandings and reduces system complexity.
* It also facilitates role division and process improvement, making it a valuable concept for all roles.