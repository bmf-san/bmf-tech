---
title: Differentiating Architecture Strategy, Tactics, and Design
slug: architecture-strategy-tactics-design
date: 2026-02-05T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture Strategy
  - Architecture
  - Design
translation_key: architecture-strategy-tactics-design
---

When writing documents related to architecture, one might wonder, "Should this be written in the strategy, tactics, or design document?" This confusion likely arises from the ambiguous definitions of strategy, tactics, and design, as well as the lack of clear criteria for differentiation.

In this article, I will introduce a three-layer structure of strategy, tactics, and design, along with guidelines for differentiation using the 5W1H framework.

## Three-Layer Structure of Strategy, Tactics, and Design

Documents related to architecture can be organized into the following three layers:

- **Strategy**: Why/What (why we are undertaking this, what we want to achieve)
- **Tactics**: How/When/Where/Who framework (what measures to take, when, where, and who)
- **Design**: Specific system design (details of interfaces, data models, technology selection, etc.)

In terms of abstraction, strategy is the most abstract, while design is the most concrete. The strategy sets the direction, tactics outline the measures, and design determines the specific implementation policy.

When actually formulating these, it is necessary to consider the three layers in a cross-sectional manner. However, when differentiating documents or organizing thoughts, it is easier to think in terms of this three-layer structure.

## Before Writing Strategy and Tactics

Before differentiating, it is essential to organize the prerequisite inputs. The following are typically necessary when considering architecture strategy:

**1. Confirming Business Strategy and Vision**
- Understanding business goals and strategies: How architecture contributes to the business
- Business drivers: Key elements for business growth, sources of competitive advantage
- Identifying stakeholders: Requirements, constraints, and expectations from management, development, and business units

**2. Organizing Requirements**
- Functional requirements: Functions that the system must fulfill
- Non-functional requirements: Performance, availability, security, maintainability, etc.

**3. Understanding Current Architecture (As-Is)**
- Investigating the current system: System configuration, data flow, technologies used, operational processes
- Identifying issues and risks: Bottlenecks, technical debt, knowledge dependency, operational load, etc.

**4. Confirming Constraints**
- Resource constraints: Budget, personnel, deadlines
- Organizational structure and team setup: Who is responsible for which areas

In addition to these, various inputs may be necessary depending on the situation, such as regulatory and compliance requirements, external dependencies, and records of past architectural decisions. It is crucial to gather a certain amount of information before formulating a strategy.

If these are ambiguous, the differentiation will result in superficial documents. It is essential to start with organizing inputs without being bound by methods.

## Common Confusion Patterns

Here are two common confusion patterns when differentiating between strategy and tactics:

**1. Writing Means (How) in Strategy**

For example, a case where the description "migrating to microservices" ends up in the strategy. Migrating to microservices is a means and should be written in tactics.

What should be written in strategy is the purpose of "why microservices are necessary." For instance, it could be described as "to reduce dependencies between teams and enable independent delivery."

**2. Losing Sight of Purpose by Making Means a Premise**

If one considers the direction with the means of migrating to microservices in mind, the purpose of "why we should migrate to microservices" becomes ambiguous.

By separating means from strategy, the abstraction level and firmness of the strategy can be maintained.

## Differentiation Criteria Using 5W1H

When assigning 5W1H to strategy and tactics, it can be organized as follows:

**Elements to Write in Strategy**

- **Why**: Purpose, motivation, business value. Why is this architectural change necessary?
- **What**: Goals to achieve, issues to resolve. What do we want to accomplish?

**Elements to Write in Tactics**

- **How**: Framework of measures. How will we achieve this?
- **When**: Timeline, phases, priorities
- **Where**: Scope of application, which parts of the system will be applied
- **Who**: Responsible teams, stakeholders

**Elements to Write in Design**

- Details of How: Specific technology selection, interface design, data models
- Details of requirements and constraints: Specific values for non-functional requirements, technical constraints
- Consideration of trade-offs: Comparison of adopted designs and alternatives

This organization clarifies the roles of strategy as "direction," tactics as "framework of measures," and design as "specific implementation policy."

## Gray Areas Where Judgment is Confused

Not everything can be clearly classified. For example, writing a context-level architecture that becomes to-be within the strategy.

A high-level abstract description like "transitioning from monolith to distributed architecture" leans towards strategy, but when showing specific system boundaries as a context-level diagram, tactical elements are also included.

**Judgment Criterion: "Is the Means a Premise?"**

When in doubt, check whether the description is based on means. If means are a premise, it becomes easy to lose sight of the purpose.

By separating means from strategy, the firmness of the policy can be maintained. While means can be flexibly changed depending on the situation, the purpose tends to remain stable.

## Examples of Formulating Strategy, Tactics, and Design

Here is an example of differentiation using the three-layer structure.

Note that the Why/What of the strategy must be logically derived from issue analysis. If the issue analysis is ambiguous, the Why/What will also lack a basis.

### Issue Analysis

Currently, it takes an average of 4 weeks to release features. The main causes are:

- Strong dependencies between modules, where one change affects a wide range
- The need to test and deploy everything together, leading to waiting times
- Releasing adjustments require agreement between multiple teams, resulting in high communication costs

The most significant impact comes from the dependencies between modules, which also ripple into the other two causes.

**Impact on Business**

Due to release delays, three large projects have been lost to competitors in the past year. Sales have reported, "We cannot propose because we cannot foresee feature additions."

**Alignment with Business Strategy**

The company-wide strategy emphasizes "rapid response to market changes," and improving delivery speed directly relates to this.

### Strategy (Why/What)

**Why: Why are we undertaking this?**

It takes an average of 4 weeks to release features, leading to delays in responding to market changes. Competitors are releasing in 2 weeks, and if this continues, we will lose competitiveness. Improving delivery speed is the top priority to realize the company-wide strategy of "rapid response to market changes."

**What: What do we want to achieve?**

We aim to reduce dependencies between modules, allowing teams to make changes and deploy independently. This will improve delivery speed to twice the current rate, achieving an average feature release within 2 weeks.

### Tactics (How/When/Where/Who)

**How: How will we achieve this?**

- Separate the order service from the monolith to enable independent deployment
- Separate the inventory service from the monolith
- Build CI/CD pipelines for each service

**When: By when, in what order?**

- Q1: Separate the order service
- Q2: Separate the inventory service
- Q3: Establish CI/CD pipelines and measure effectiveness

**Where: Where will it be applied?**

- Target: Order and inventory areas (the areas with the most complex dependencies)
- Out of scope: Other areas will be considered in the next phase

**Who: Who will be responsible?**

- Order service: Order team + Platform team
- Inventory service: Inventory team + Platform team

### Design

**Overview of Order Service Design**

- Communication method: Integration with the monolith will be done via REST API. Future consideration for migration to asynchronous messaging.
- Data separation: Order data will be migrated to a dedicated DB. During the migration period, bidirectional synchronization will be implemented.
- Interface: Publicly released as Order API v1. A versioning policy will be adopted to maintain backward compatibility.

**Architecture Diagram, Use Cases, Sequence**

(Images omitted)

- Architecture diagram: Shows the relationship between the order service, monolith, and external systems.
- Use cases: Major use cases such as order registration, order inquiry, and order cancellation.
- Sequence diagram: Flow of integration with the monolith during order registration.

**Non-Functional Requirements**

- Availability: 99.9% (monthly downtime within 43 minutes)
- Response time: p95 within 200ms
- Throughput: 1,000 requests/second at peak

**Trade-offs**

- Synchronous vs. Asynchronous communication: Initially adopting simple synchronous communication to reduce complexity and enable early release.
- Shared DB vs. Dedicated DB: Adopting a dedicated DB. While there are data consistency challenges, prioritizing independent deployment.

## Conclusion

By differentiating strategy, tactics, and design, the roles of each become clear.

- Write **Why/What** in **Strategy**. Focus on purpose and what you want to achieve.
- Write the framework of **How/When/Where/Who** in **Tactics**. Indicate the direction of measures.
- Write the details of **How** in **Design**. Indicate specific technology selection and implementation policy.
- When in doubt, judge by asking, "Is the means a premise?"

By separating means from strategy and delegating details to design, each document can maintain the appropriate level of abstraction.