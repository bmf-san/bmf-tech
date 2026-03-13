---
title: 'Clean Architecture: Learning Software Structure and Design from the Experts'
description: 'Learn clean architecture principles: SOLID guidelines, dependency inversion, components, and layered software design.'
slug: clean-architecture-software-structure-design
date: 2018-08-01T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Clean Architecture
  - Architecture
  - Book
translation_key: clean-architecture-software-structure-design
---



[Clean Architecture 達人に学ぶソフトウェアの構造と設計](https://amzn.to/4agMQ99)を読んだ。

Reading this book and the author's blog is a good starting point for learning about Clean Architecture.

- There is no clear definition of Clean Architecture provided.
  - It does not claim that the common concentric circles represent Clean Architecture, nor does it state that having a layered structure is a prerequisite (though following the rules will likely lead to a layered structure).
- It discusses characteristics and rules considered to define a Clean Architecture.
- The purpose of software architecture is to minimize the necessary personnel for building and maintaining the system.
- By utilizing polymorphism, dependency inversion can be achieved.
    - This increases independent deployability and independent development capability.
- The usefulness of immutability:
  - Easier debugging, thread safety, high cache usability, easier testing, etc.
- Structured programming imposes discipline on direct control, object-oriented programming on indirect control, and functional programming on assignment.
  - The SOLID principles are guidelines for creating software structures that are resilient to change, easy to understand, and easy to use by other software.
    - Single Responsibility Principle (SRP):
      - Ensure that a module has a single responsibility, with only one reason to change.
    - Open-Closed Principle (OCP):
      - Open for extension, closed for modification.
    - Liskov Substitution Principle (LSP):
      - Subtypes should be substitutable for their base types.
    - Interface Segregation Principle (ISP):
      - Avoid dependencies on things you do not use.
    - Dependency Inversion Principle (DIP):
      - High-level modules should not depend on low-level modules. Both should depend on abstractions.
- Components are units of deployment.
- Component cohesion involves a trade-off between development convenience and reusability.
- Principles related to component cohesion:
  - Reuse-Release Equivalence Principle (REP):
    - Only reuse released components.
  - Common Closure Principle (CCP):
    - Group things that change for the same reasons and at the same times.
  - Common Reuse Principle (CRP):
    - Depend on all the components you use.
  - Component coupling involves a trade-off between development convenience and logical design.
  - Principles related to component coupling:
    - Acyclic Dependencies Principle (ADP):
      - Avoid cyclic dependencies in component relationships.
    - Stable Dependencies Principle (SDP):
      - Depend in the direction of stability (i.e., less frequently changed components).
    - Stable Abstractions Principle (SAP):
      - The abstraction level of a component should match its stability (high stability with high abstraction, low stability with low abstraction).
  - The goal of software architecture shapes is to facilitate development, deployment, operation, and maintenance, leaving as many options open for as long as possible.
  - Delay decisions on details that do not need immediate resolution, or make them changeable at any time.
  - Software holds both "behavioral value" and "structural value," with the latter being more valuable as it makes software flexible (modifiable).