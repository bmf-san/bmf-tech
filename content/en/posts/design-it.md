---
title: Design It
slug: design-it
date: 2019-11-25T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architect
  - Books
translation_key: design-it
---

I revisited [Design It](https://amzn.to/4635mAS), which I read about 1-2 years ago, and learned about the roles and responsibilities of a software architect, so I wanted to leave some notes.

This will be a bit poetic as I mix in my feelings.

# What Does a Software Architect Do?
A software architect writes code, leads projects, and thinks from a business perspective.

They stand at the center of three elements: business, technology, and users.

The responsibilities of a software architect are generally defined in Design It as follows:

- Problem definition from an engineering perspective
- System decomposition and responsibility assignment
- Maintaining an overall view
- Deciding on trade-offs between quality attributes
- Managing technical debt
- Improving the architectural skills of the team

System decomposition and responsibility assignment do not necessarily mean microservices; it’s more about being able to reduce the size of the system, I think. Probably.

The educational responsibility of improving the team's skills was a bit refreshing to me.

# What is Software Architecture?
Software architecture is defined as "a collection of design decisions on how to build software to exhibit quality attributes and other properties."

In other words, **a software architect needs the ability to understand the characteristics required by the system and make design decisions to exhibit those characteristics**.

There are many things to do and be careful about in software construction, but one thing a software architect should be cautious of is "avoiding costly mistakes."

**This means being able to detect and avoid parts where mistakes could lead to significant issues later on**, which I feel is an area where experience can greatly vary.

In relation to this, one of the guiding principles for software architects is to "delay design decisions as much as possible."

**This means postponing important design decisions (≒ decisions where costly mistakes could occur) as much as possible.**

There are times in life when you feel the same way; software architecture is life. I deeply felt that.

For issues that absolutely cannot move forward without a decision, I think it’s important to keep options open and make judgments. This is also life.

For a deeper discussion on software architecture, I think you can refer to [Software Architecture Basics: A Systematic Approach Based on Engineering](https://amzn.to/3PeBOKE). It covers similar topics. Various characteristics of software architecture are also organized there.

The sequel, [Software Architecture Hard Parts: Trade-off Analysis for Distributed Architecture](https://amzn.to/3JfKKM6), contains more advanced content, but I also find it interesting. As a side note, I have a personal connection to this book as I participated as a volunteer reviewer.

# Programmers are Software Architects
On page 214 of Design It, it states, "Programmers make architectural decisions daily." Even a single line of code can lead to design decisions that affect the quality attributes of the architecture, so even programmers can be considered software architects.

Regardless of job responsibilities or titles, it’s important to be aware of acting as an architect when facing software.

# Thoughts
I want to be able to proudly say that I am a software architect someday.

# References
- [Design It](https://amzn.to/4635mAS)