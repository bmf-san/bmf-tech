---
title: Bounded Contexts
description: An in-depth look at Bounded Contexts, covering key concepts and practical insights.
slug: bounded-context-explanation
date: 2025-05-19T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - DDD
translation_key: bounded-context-explanation
---

# Overview

There was a challenge to convey to non-developers why contexts organized in a context map are divided in such a way and what the significance of this separation is.

In this article, I will attempt to explain "Bounded Contexts" in a way that is understandable to non-developers.

# Being Aware of Context Differences

Even if the same word is used, its meaning can change depending on the context.

For example, consider the word "order."

In the sales department, it may refer to "a request from a customer," while in the warehouse, it could mean "shipping instructions," and in the accounting department, it might refer to "billing data."

Here are some other examples:

* "User": In the development team, it refers to someone who logs in; in customer support, it refers to someone making inquiries; in marketing, it refers to target customers.
* "Service": For engineers, it means an API; for sales, it refers to customer plans.

Thus, **the same word can have different meanings depending on the business or position in which it is used**.

The concept of "context" aims to clarify these differences.

# What is Context?

Context refers to **a grouping of business activities where the meanings of words and rules consistently apply**.

In Domain-Driven Design (DDD), this is referred to as a "Bounded Context."

For instance, the meaning of the word "order" can differ based on context as follows:

* In Context A, it refers to the customer's intention to purchase.
* In Context B, it refers to a shipping request to the warehouse.
* In Context C, it refers to the subject of billing.

In this way, the characteristic of context is to "segment" where the meanings and rules of words change in each business.

## Why Organize Contexts?

If you proceed with business or design systems without being aware of these contexts, various problems can arise.

* The "order" table becomes complex and fragile because it is stuffed with billing and shipping information.
* Questions like "What does ○○ mean?" circulate between teams, leading to misunderstandings and unclear responsibilities.
* Even when trying to improve business processes, it becomes difficult to assess the impact.

Conversely, if you separate and organize contexts, you can gain the following benefits:

* The meanings of words become clear.
* The scope of responsibilities for business and systems is organized.
* Team assignments and changes become easier.

In other words, by properly organizing contexts, systems and organizations can benefit in the following ways:

* **Reduction of communication costs**
* **Decreased complexity in system design and implementation**
* **Clarification of roles and responsibilities between teams**

Properly organized contexts lead to **improved stability in both systems and organizations**.

# Example of Context: Orders in an E-Commerce Site

Let’s divide the business related to orders in an e-commerce site into three departments.

| Department      | Meaning of "Order"       |
| ------- | ------------- |
| Frontend Site | Content confirmed by the customer in the cart |
| Warehouse Operations    | Instruction data for shipping    |
| Accounting      | Sales data to be processed for billing  |

All of these use the word "order," but the **content, purpose, and processing involved are completely different**.

Ignoring these differences and lumping everything under one term leads to:

* Increased complexity in the system.
* Uncertainty about the scope of impact during modifications.
* Higher likelihood of misunderstandings.

This is why it is important to think about "contexts" separately.

# Not Just for Developers

Contexts are not something that only developers need to understand.

For example, have you ever wondered about the following questions?

* "What is our team's scope of responsibility?"
* "Should I ask sales, development, or support about this?"
* "What impact will this change have?"

These questions are all closely related to the organization of contexts.

**Draw boundaries where the purpose of business and the meanings of words change, and clarify roles and responsibilities within those boundaries.**

This perspective is beneficial not only for developers but also for sales, support, planning, and all other professions.

# Conclusion

* The same word can often have different meanings.
* Recognizing these differences and organizing them as boundaries for business and systems is what "context" is about.
* Clarifying contexts can prevent discrepancies in understanding and system complexity.
* This is an effective approach for all professions, leading to clearer role assignments and business improvements.