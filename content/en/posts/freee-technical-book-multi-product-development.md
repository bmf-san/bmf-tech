---
title: Multi-Product Development at freee
description: 'Discover how freee manages multi-product SaaS development using bounded contexts, service segmentation, and fitness functions.'
slug: freee-technical-book-multi-product-development
date: 2024-11-09T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Design
  - System Design
  - Book
translation_key: freee-technical-book-multi-product-development
books:
  - title: freee Technical Book - Multi-Product Development at freee
    url: https://techbookfest.org/product/8UNZJF7Rs5AWuqReQuZZgs?productVariantID=fZvUHjHxH4hqG6wtAaAVWN
---

[Read the book on freee's technology](https://techbookfest.org/product/8UNZJF7Rs5AWuqReQuZZgs?productVariantID=fZvUHjHxH4hqG6wtAaAVWN).

- In business SaaS, it's interesting to see if common business domains (applications and approvals) can be segmented. Even within applications and approvals, workflows and data can differ significantly depending on the context, which seems challenging. Compared to permissions and notifications, it appears that abstracting this functionality is difficult. It seems that they avoid handling application-related content as a foundation. However, this leads to the challenge of having close communication between applications. It’s a trade-off, indeed.
- The issue of balancing system replacement with feature development is addressed by establishing development restriction policies, which resonates deeply. Maintaining balance between both is difficult, and maximizing outcomes is also challenging.
- I believe it's important to measure the results of service segmentation and re-architecture. What metrics to use connects to the OKRs of foundational development. Personally, I think it's crucial to consider not only technical metrics but also business-related metrics. Given that foundational development tends to be technology-driven, I believe that focusing on "solving customer problems" is essential to maximize its value as a foundation, and this should be clearly defined.
- Tips for SaaS development include "creating appropriate margins," "thoroughly defining concepts," and focusing on "scalability, maintainability, commonality, and standardization."
  - Creating margins means leaving options open.
  - "Yes, at the coding level that's true, but at the conceptual level, we need to discover that 'there should be a common concept here' in advance. (p.41)"
    - I strongly resonate with this. I perceive this as abstraction.
- There was an example of a fitness function. I was curious about how fitness functions are defined in practice, so it was helpful.