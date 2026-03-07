---
title: 'freee Technology Book: Multi-Product Development at freee'
slug: freee-technical-book-multi-product-development
date: 2024-11-09T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Design
  - System Design
  - Books
description: Insights and reflections on multi-product development at freee, based on the freee technology book.
translation_key: freee-technical-book-multi-product-development
---

[freee Technology Book: Multi-Product Development at freee](https://techbookfest.org/product/8UNZJF7Rs5AWuqReQuZZgs?productVariantID=fZvUHjHxH4hqG6wtAaAVWN) - I read this book.

- In business SaaS, the idea of extracting a domain area like common business processes (e.g., application and approval) is fascinating. Even within application and approval workflows, the context can differ, leading to variations in workflows and data handling, which seems challenging. Compared to permissions or notifications, functional segmentation here appears harder to abstract. It seems the platform intentionally avoids handling application-related content. However, this creates a challenge where close communication with each application becomes necessary. This trade-off is quite thought-provoking.

- Regarding the challenge of balancing system replacement with feature development, the book discusses the establishment of development restriction policies, which deeply resonated with me. Maintaining a balance between the two is difficult, and it’s hard to maximize outcomes simultaneously.

- Measuring the outcomes of service segmentation and re-architecture is crucial. Choosing the right metrics seems to tie into the OKRs of platform development. Personally, I believe it’s important to link these metrics not only to technical indicators but also to business-related ones. While platform development tends to be technology-driven by nature, its value as a platform cannot be maximized unless it focuses on "solving customer problems." This focus should be made explicit.

- Tips for SaaS development: "Create adequate flexibility," "Thoroughly define concepts," and "Focus on scalability, maintainability, unification, and standardization."
  - Creating flexibility seems to mean leaving room for options.
  - "Yes, at the coding level, that’s true, but at the conceptual level, you need to identify common concepts in advance. (p.41)"
    - I strongly resonated with this. I personally interpret this as abstraction.

- The book included examples of fitness functions. I’ve always wondered how fitness functions are defined in practice, so this was very insightful.