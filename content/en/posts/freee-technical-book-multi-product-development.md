---
title: 'freee Technical Book: Multi-Product Development at freee'
description: 'freee Technical Book: Multi-Product Development at freee'
slug: freee-technical-book-multi-product-development
date: 2024-11-09T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Design
  - System Design
  - Book Review
translation_key: freee-technical-book-multi-product-development
books:
  - title: freee 技術の本 freeeにおけるマルチプロダクト開発
    url: https://techbookfest.org/product/8UNZJF7Rs5AWuqReQuZZgs?productVariantID=fZvUHjHxH4hqG6wtAaAVWN
---

I read the [freee Technical Book: Multi-Product Development at freee](https://techbookfest.org/product/8UNZJF7Rs5AWuqReQuZZgs?productVariantID=fZvUHjHxH4hqG6wtAaAVWN).

- In the realm of business SaaS, it's interesting to see how common business processes (like applications and approvals) can be delineated. Even though applications and approvals are mentioned, the workflows can differ depending on the context, and the data handled will naturally vary, making it seem quite challenging. Compared to permissions and notifications, it seems that functional division is harder to abstract. It appears that they try not to handle application-related content at the infrastructure level. However, because they avoid it, there seems to be a challenge of close communication between applications. It’s a clear trade-off.
- Regarding the challenge of balancing system replacement with feature development, I found the discussion about setting development restriction policies to be very relatable. Continuously balancing both is difficult, and maximizing outcomes can be challenging.
- I believe it's important to measure the results of service division and re-architecture. What indicators to use connects to the OKRs of infrastructure development. Personally, I think it’s crucial to consider not just technical indicators but also to relate them to business indicators. Given the nature of infrastructure development, it tends to be technology-driven, but if we don’t focus on "solving customer problems," we can’t maximize the value as an infrastructure, so I think that should be made clear.
- Tips for SaaS development include "creating appropriate margins," "thoroughly defining concepts," and "extensibility, maintainability, commonality, and standardization."
  - Creating margins feels like leaving options open.
  - "Yes, at the coding level, that's true, but at the conceptual level, we need to discover that 'there should be a common concept here' in advance. (p.41)"
    - I strongly resonated with this. I perceive this as abstraction.
- There was an example of a fitness function. I had been wondering how fitness functions are actually defined in the field, so it was helpful.
