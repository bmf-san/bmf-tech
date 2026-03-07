---
title: Learning Software Structure and Design from Clean Architecture Experts
slug: clean-architecture-software-structure-design
date: 2018-08-01T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Clean Architecture
  - Architecture
  - Books
translation_key: clean-architecture-software-structure-design
---

[Learning Software Structure and Design from Clean Architecture Experts](https://amzn.to/4agMQ99) has been read.

If you want to learn about Clean Architecture, I think it's best to start with this book and the author's blog.

- There is no clear definition of Clean Architecture discussed.
  - It is not claimed that the common concentric circles represent Clean Architecture, nor is it stated that a layered structure is a prerequisite (though I believe that fulfilling the rules will inevitably lead to a layered structure).
- Characteristics and rules that can be considered features of a clean architecture are discussed.
- The purpose of software architecture is to minimize the necessary personnel for system construction and maintenance.
- By utilizing polymorphism, dependency inversion can be achieved.
    - This increases independent deployability and independent development capability.
- The usefulness of immutability:
  - Easier debugging, thread safety, high cache usability, easier testing, etc.
- Structured programming provides direct control, object-oriented programming provides indirect control, and functional programming imposes discipline on assignment.
  - The SOLID principles are designed to create software structures that are resilient to change, easy to understand, and easy to utilize from other software:
    - Single Responsibility Principle (SRP):
      - Ensure that a module has a single responsibility so that the reasons for changes are unified.
    - Open-Closed Principle (OCP):
      - Open for extension, closed for modification.
    - Liskov Substitution Principle (LSP):
      - Ensure that subtypes can replace their supertype.
    - Interface Segregation Principle (ISP):
      - Avoid dependencies on things that are not used.
    - Dependency Inversion Principle (DIP):
      - Ensure that high-level policy implementations do not depend on low-level details, but rather that low-level details depend on high-level policies.
- Components are units of deployment.
- The cohesion of components relates to the trade-off between development convenience and reusability.
- Principles regarding component cohesion:
  - Reuse-Release Equivalence Principle (REP):
    - Only reuse what has been released.
  - Common Closure Principle (CCP):
    - Group things that change for the same reason or timing together.
  - Common Reuse Principle (CRP):
    - When using a component, depend on all of its parts.
  - The coupling of components relates to the trade-off between development convenience and logical design.
  - Principles regarding component coupling:
    - Acyclic Dependencies Principle (ADP):
      - Ensure that component dependencies do not include cyclic dependencies.
    - Stable Dependencies Principle (SDP):
      - Depend on the direction of high stability (i.e., low frequency of change).
    - Stable Abstractions Principle (SAP):
      - The abstraction level of components should be on par with their stability (high stability should correspond to high abstraction, while low stability can correspond to low abstraction).
  - The goal of the shape of software architecture is to facilitate development, deployment, operation, and maintenance, with a strategy to leave as many options open for as long as possible.
  - For decisions about details that do not need to be made immediately, it is best to delay them as much as possible or keep them changeable at any time.
  - Software has both "value of behavior" and "value of structure," but the latter is more valuable because it makes software more flexible (modifiable).