---
title: 'Learning Software Structure and Design from Clean Architecture Experts'
description: 'Learning Software Structure and Design from Clean Architecture Experts'
slug: clean-architecture-software-structure-design
date: 2018-08-01T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Clean Architecture
  - Architecture
  - Book Review
translation_key: clean-architecture-software-structure-design
books:
  - asin: "4048930656"
    title: "Clean Architecture 達人に学ぶソフトウェアの構造と設計"
---

I read [Learning Software Structure and Design from Clean Architecture Experts](https://amzn.to/4agMQ99).

If you want to learn about Clean Architecture, I think it's best to start with this book and the author's blog.

- There is no clear definition of Clean Architecture provided.
  - It does not claim that the common concentric circles represent Clean Architecture, nor does it state that a layered structure is a prerequisite (though I believe that following the rules will inevitably lead to a layered structure).
- The book discusses characteristics and rules that can be considered features of a clean architecture.
- The purpose of software architecture is to minimize the necessary personnel for system construction and maintenance.
- By utilizing polymorphism, you can achieve dependency inversion.
    - This increases independent deployability and independent development capability.
- The usefulness of immutability:
  - It makes debugging easier, is thread-safe, has high cache usability, and is easier to test.
- Structured programming provides direct control, object-oriented programming provides indirect control, and functional programming imposes discipline on assignment.
  - The SOLID principles create a software structure that is resilient to changes, easy to understand, and usable by other software:
    - Single Responsibility Principle (SRP):
      - Ensure that a module has a single responsibility so that the reasons for changes are unified.
    - Open-Closed Principle (OCP):
      - Be open to extension but closed to modification.
    - Liskov Substitution Principle (LSP):
      - Allow subtypes to be substituted for their supertype.
    - Interface Segregation Principle (ISP):
      - Avoid dependencies on things that are not used.
    - Dependency Inversion Principle (DIP):
      - Ensure that high-level policy implementations do not depend on low-level details, but rather that low-level details depend on high-level policies.
- Components are the units of deployment.
- The cohesion of components relates to the trade-off between development convenience and reusability.
- Principles regarding component cohesion:
  - Reuse-Release Equivalence Principle (REP):
    - Only reuse what has been released.
  - Common Closure Principle (CCP):
    - Group together things that are changed for the same reason or timing.
  - Common Reuse Principle (CRP):
    - When using components, depend on all of them.
  - The coupling of components relates to the trade-off between development convenience and logical design.
  - Principles regarding component coupling:
    - Acyclic Dependencies Principle (ADP):
      - Avoid including circular dependencies in component dependencies.
    - Stable Dependencies Principle (SDP):
      - Depend in the direction of higher stability (i.e., less frequent changes).
    - Stable Abstractions Principle (SAP):
      - The level of abstraction of components should match their stability (highly stable components can have high abstraction, while less stable components can have lower abstraction).
  - The goal of the shape of software architecture is to facilitate development, deployment, operation, and maintenance, and the strategy is to leave as many options open for as long as possible.
  - For decisions about details that do not need to be made immediately, it is best to delay them as much as possible or to keep them changeable at any time.
  - Software has both "behavioral value" and "structural value," but the latter is more valuable because it makes the software flexible (modifiable).
