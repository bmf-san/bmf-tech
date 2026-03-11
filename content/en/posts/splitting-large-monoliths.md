---
title: How to Split a Large Monolith? - Lessons from Software Architecture Hard Parts
slug: splitting-large-monoliths
date: 2025-02-17T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture Strategy
  - Monolith
  - Microservices
description: 'This article organizes useful points for considering service division from a monolith, based on chapters 1 to 4 of ''Software Architecture: The Hard Parts''.'
translation_key: splitting-large-monoliths
---



In this article, we organize useful points for considering service division from a monolith, based on chapters 1 to 4 of [Software Architecture: The Hard Parts](https://amzn.to/41kcsAL).

There is no "silver bullet" that applies to all organizations, but understanding "what trade-offs to identify" might help in designing a more convincing architecture.

## 1. There is No "Best Practice" — Identifying Trade-offs
### 1.1 Why There is No "Silver Bullet"
- **Differences in Preconditions**
  - Due to differences in organizational size, team composition, nature of data handled, business requirements, etc., a single division pattern rarely fits all situations.
- **Existence of Conflicting Requirements**
  - "The more you enhance performance, the higher the operational cost," "The more you strengthen security, the more it affects availability," etc. There are always some trade-offs in technology selection and design policies.

### 1.2 Utilizing ADR (Architecture Decision Record)
- **Importance of Leaving Decisions**
  - Documenting (ADR) what options were available and why a particular choice was made makes it easier to trace "why this structure was chosen" when reviewing the architecture later.
- **Granularity of Records**
  - Recording not only major architectural changes (e.g., transition from monolith to distributed architecture) but also small design trade-offs can facilitate consensus building within the team.

## 2. Architecture Quantum — How Much is a "Chunk"?
In [Software Architecture: The Hard Parts](https://amzn.to/41kcsAL), the concept of **Architecture Quantum** is emphasized.

> **Architecture Quantum**:
> A unit that can be "independently deployed" and has "high functional cohesion, with elements strongly coupled both statically and dynamically."

**Coupling** refers to dependency relationships where changes in one affect the other.

**Cohesion** refers to the degree to which related elements are grouped together, with higher cohesion when grouped for the same purpose.

### 2.1 Examples of Static and Dynamic Coupling
- **Static Coupling:**
  - A typical example is a single large database. When table design and schema are integrated, dependencies remain even if you want to separate deployment units.
  - Countermeasures: Split table schemas, have independent DB instances per service, etc.
- **Dynamic Coupling:**
  - The more synchronous service calls there are, the wider the impact range during failures.
  - Countermeasures: Utilize asynchronous messaging, consider designs that avoid distributed transactions, etc.

### 2.2 Factors Leading to Large Quanta
- When specific functions frequently exchange data with other functions, even if separated, they often end up needing to be "deployed together to function."
- Conversely, parts with weak communication or data dependencies can be relatively easily extracted.

## 3. Why Aim for Service Division — Drivers of Modularization
### 3.1 Maintainability and Testability
- **Improved Maintainability:**
  - Large monoliths have unpredictable impact ranges, with a high risk of unexpected bugs from a single fix.
  - If responsibility ranges are clear per service, it's easier to limit the responsible team and test range.
- **Improved Testability:**
  - Easier to perform unit tests and component tests per function.
  - With microservices, CI/CD pipelines can be established per service, shortening release cycles.

### 3.2 Deployability and Scalability
- **Independent Deployment:**
  - When releasing only part of the functionality, the entire app doesn't need to be stopped, improving agility.
- **Selective Scaling:**
  - Only high-load services can be scaled out.
  - Infrastructure resources can be used intensively, optimizing costs.

### 3.3 Availability and Fault Tolerance
- If services are divided, even if part goes down, the impact on the whole can be minimized.
- However, if inter-service dependencies are strong, there can still be situations where "if one part fails, it affects everything," so careful management of coupling points is necessary.

## 4. Division Approach — How to Split a Monolith
### 4.1 Component-Based Decomposition Steps
1. **Modularize the Monolith**
   - First, visualize existing code dependencies.
   - Find "components" through layering and package splitting.
2. **Consider Database Splitting**
   - If possible, organize by table or schema unit, aiming for independent DBs per service.
   - To reduce migration risk, first logically separate schemas, then physically separate DBs.
3. **Create a State for Independent Service Deployment**
   - Gradually build CI/CD pipelines, enabling each service to be tested and deployed independently.
   - A full microservices transition can be confusing, so a small start is recommended.

### 4.2 Cautions for Tactical Forking
- **Method:**
  - Extract only necessary functions to create a new service, ceasing use from the monolith side (evolve the fork).
- **Advantages:**
  - Can secure independent services in the short term, speeding up migration.
- **Disadvantages:**
  - Synchronization with the fork source can become difficult, potentially increasing future management costs.
- **Points:**
  - Clarify the code maintenance policy post-fork.
  - Team consensus is essential on how much duplicate code is acceptable.

## 5. Technical Checklist When Challenging Service Division
1. **Trade-off Analysis**
   - Determine the priority of key elements (performance, availability, security, cost, etc.).
   - Preconceiving which elements can be sacrificed helps stabilize decision-making.
2. **Understanding Architecture Quantum**
   - Confirm whether services can operate independently, and whether shared DBs or synchronous calls are "bottlenecks."
3. **Gradual Implementation**
   - Large-scale refactoring is risky, so start with modularization.
   - Data layer decomposition planning, such as table design splitting and using multiple DBs, is crucial.
4. **Clarify Priorities for Maintainability and Scalability**
   - Share within the team and with management "what operational structure will be after division" and "what level of fault tolerance is needed."

## 5. Conclusion
There is no universal solution for service division, but understanding trade-offs and identifying optimal architecture quanta allows for gradual execution.

Architectural changes take time, so it's important to start with small steps.