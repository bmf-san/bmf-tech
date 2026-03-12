---
title: Design It
slug: design-it
date: 2019-11-25T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architect
  - Book Review
description: Reflecting on the roles and responsibilities of a software architect.
translation_key: design-it
books:
  - asin: "4873118956"
    title: "Design It"
---



I recently revisited [Design It](https://amzn.to/4635mAS), a book I read about 1-2 years ago, and found it insightful regarding the roles and responsibilities of a software architect. Here are some notes.

This is a bit of a personal reflection, so consider it a poem.

# What Does a Software Architect Do?
A software architect writes code, leads projects, and thinks from a business perspective.

They stand at the intersection of business, technology, and user needs.

In *Design It*, the responsibilities of a software architect are generally defined as follows:

- Defining problems from an engineering perspective
- Dividing systems and assigning responsibilities
- Continuously overseeing the whole
- Deciding on trade-offs between quality attributes
- Managing technical debt
- Enhancing the team's architectural skills

Dividing systems and assigning responsibilities doesn't necessarily mean focusing on microservices; it's more about keeping the system size manageable. Probably.

The educational responsibility of improving the team's skills was a fresh perspective.

# What is Software Architecture?
Software architecture is defined as "a collection of design decisions about how to build software to exhibit certain quality attributes and other properties."

In other words, a software architect needs the ability to understand the characteristics required by the system and make design decisions to exhibit those characteristics.

There are many things to do and be cautious about when building software, but one thing a software architect should be careful about is "avoiding costly mistakes."

**This means being able to detect and avoid parts that could lead to significant issues later, essentially risk control,** which I feel is an area where experience makes a big difference.

Related to this, one guiding principle for software architects is to "delay design decisions as much as possible."

**Important design decisions (≒ decisions where costly mistakes could occur) should be postponed as much as possible.**

This is true in life too; software architecture is life. I thought deeply about this.

For issues that must be decided now to move forward, I think it's important to leave options open when making decisions. This is also true in life.

For a deeper dive into software architecture, I think [Software Architecture: The Hard Parts](https://amzn.to/3PeBOKE) is a good reference. It covers similar topics and organizes various characteristics of software architecture.

The sequel, [Software Architecture: The Hard Parts](https://amzn.to/3JfKKM6), is more advanced but also interesting. As a side note, I participated as a volunteer in reviewing this book, so I have a personal connection to it.

# Programmers as Software Architects
On page 214 of *Design It*, it states, "Programmers make architectural decisions daily." Even a single line of code can be a design decision that affects architectural quality attributes, so even programmers can be considered software architects.

Regardless of differing responsibilities or changing titles, if you're dealing with software, it's important to act with the mindset of an architect.

# Personal Thoughts
I aspire to one day confidently call myself a software architect.

# References
- [Design It](https://amzn.to/4635mAS)
