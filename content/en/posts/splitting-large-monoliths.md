---
title: How to Split a Large Monolith? - Learning from Software Architecture and Hard Parts
slug: splitting-large-monoliths
date: 2025-02-17T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture Strategy
  - Monolith
  - Microservices
translation_key: splitting-large-monoliths
---

In this article, we will organize points that are helpful when considering service splitting from a monolith, based on Chapters 1 to 4 of [Software Architecture and Hard Parts - Trade-off Analysis for Distributed Architecture](https://amzn.to/41kcsAL).

There is no "silver bullet" that applies to all organizations, but by understanding "what trade-offs to identify," we may be able to design a more convincing architecture.

## 1. There Are No "Best Practices" - Identifying Trade-offs
### 1.1 Why There Is No "Silver Bullet"
- **Differences in Preconditions**  
  Due to differences in organizational size, team composition, the nature of the data handled, and business requirements, a single splitting pattern rarely fits all situations.
- **Existence of Conflicting Requirements**  
  There are always trade-offs in technology selection and design policies, such as "increasing performance raises operational costs" and "enhancing security impacts availability."

### 1.2 Utilizing ADR (Architecture Decision Record)
- **Importance of Documenting Decisions**  
  Documenting what options were available and why a particular choice was made (ADR) makes it easier to trace back to "why this structure was chosen" when reviewing the architecture later.
- **Granularity of Records**  
  Recording not only major architectural changes (e.g., transitioning from a monolith to a distributed architecture) but also small design trade-offs can facilitate consensus within the team.

## 2. Architecture Quantum - How Much is a "Single Unit"?
The concept of **architecture quantum** is emphasized in [Software Architecture and Hard Parts - Trade-off Analysis for Distributed Architecture](https://amzn.to/41kcsAL).

> **Architecture Quantum**:  
> A unit that is "independently deployable" and has "high functional cohesion, with elements that are strongly coupled both statically and dynamically."

**Coupling** refers to the state of dependency where a change in one affects the other.

**Cohesion** refers to the degree to which related elements are grouped together, with higher cohesion when they are grouped for the same purpose.

### 2.1 Examples of Static and Dynamic Coupling
- **Static Coupling:**  
  A typical example is a single large database. If table design and schema are integrated, dependencies are likely to remain even if you want to separate deployment units.  
  Example countermeasures: Splitting table schemas, having independent DB instances per service, etc.
- **Dynamic Coupling:**  
  The more synchronous service-to-service calls there are, the wider the impact range during failures.  
  Example countermeasures: Considering the use of asynchronous messaging, avoiding distributed transactions, etc.

### 2.2 Factors Leading to Large Quantums
- When a specific function frequently exchanges data with other functions, it tends to lead to a situation where **"it ultimately cannot operate unless deployed together."**  
- Conversely, areas with weak communication or data dependencies may be relatively easy to extract.

## 3. Why Aim for Service Splitting - Drivers for Modularization
### 3.1 Maintainability and Testability
- **Improved Maintainability:**  
  In a large monolith, the impact range is hard to predict, and a single fix carries a high risk of causing unexpected issues.  
  If the responsibility range is clear for each service, it becomes easier to limit the responsible teams and testing scope.
- **Improved Testability:**  
  Unit tests and component tests can be performed more easily at the functional level.  
  By adopting microservices, CI/CD pipelines can be established for each service, potentially shortening the release cycle.

### 3.2 Deployability and Scalability
- **Independent Deployment:**  
  When releasing only part of the functionality, you don't have to stop the entire application, improving agility.
- **Selective Scaling:**  
  Only high-load services can be scaled out.  
  Since infrastructure resources can be used more intensively, cost optimization is also expected.

### 3.3 Availability and Fault Tolerance
- If services are split, even if part of them goes down, the impact on the whole can be minimized.  
- However, if inter-service dependencies are strong, it can lead to a situation where **"if one part fails, it affects everything,"** so careful management of coupling points is necessary.

## 4. Splitting Approaches - How to Divide a Monolith
### 4.1 Component-Based Decomposition Steps
1. **Modularize the Monolith**  
   First, visualize the dependencies of the existing code.  
   Identify "components" through layering and package splitting.
2. **Consider Database Splitting**  
   If possible, organize at the table or schema level, aiming for independent DBs per service.  
   To reduce migration risk, a pattern of first logically separating schemas before physically separating DBs can be adopted.
3. **Create a State Where Services Can Be Independently Deployed**  
   Gradually build CI/CD pipelines so that each service can be tested and deployed independently.  
   Jumping straight to full microservices can cause confusion, so a small start is recommended.

### 4.2 Caution on Tactical Forking
- **Method:**  
  Extract only the necessary functions to create a new service, and stop using it from the monolith (evolve the fork).
- **Benefits:**  
  Can secure independent services in the short term and speed up migration.
- **Drawbacks:**  
  Tends to lose synchronization with the original fork, potentially increasing management costs in the future.
- **Points:**  
  Clearly define the code maintenance policy after forking.  
  Agreement within the team on how much duplicate code to allow is essential.

## 5. Technical Checklist for Tackling Service Splitting
1. **Trade-off Analysis**  
   Determine the priorities of key elements (performance, availability, security, cost, etc.).  
   Anticipating which elements can be sacrificed in advance helps keep decision-making steady.
2. **Understanding Architecture Quantum**  
   Check whether services can operate independently, and whether shared DBs or synchronous calls are not becoming a "bottleneck."
3. **Incremental Implementation**  
   Large-scale refactoring carries significant risks, so start with modularization.  
   Planning the decomposition of the data layer, such as splitting table designs or using multiple DBs, is crucial.
4. **Clarify Priorities for Maintainability and Scalability**  
   Share with the team and management what kind of operational structure will be established after splitting and how much fault tolerance is needed.

## 6. Conclusion
There is no universal solution for service splitting, but by understanding trade-offs and identifying the optimal architecture quantum, it becomes possible to execute incrementally.

Since architectural changes take time, it is important to start with small steps.