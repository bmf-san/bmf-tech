---
title: Differentiating Architecture Strategy, Tactics, and Design
description: An in-depth exploration of Differentiating Architecture Strategy, Tactics, and Design, covering design principles, trade-offs, and practical applications.
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



When writing documents related to architecture, one might wonder, "Should this be included in the strategy, tactics, or design document?"

This confusion likely arises from the ambiguous definitions of strategy, tactics, and design, as well as the lack of clear criteria for differentiation.

In this article, we introduce a three-layer structure of strategy, tactics, and design, along with guidelines for differentiation using the 5W1H framework.

## Three-Layer Structure of Strategy, Tactics, and Design

Documents related to architecture can be organized into the following three layers:

- **Strategy**: Why/What (Why undertake this, what do we want to achieve)
- **Tactics**: The broad outline of How/When/Where/Who (What measures, when, where, who)
- **Design**: Specific system design (interfaces, data models, technical selection details, etc.)

In terms of abstraction, strategy is the most abstract, while design is the most concrete. Strategy determines the direction, tactics outline the measures, and design decides the specific implementation policies.

When actually formulating these, it is necessary to consider them across the three layers. However, when differentiating documents or organizing thoughts, it is easier to think in terms of this three-layer structure.

## Before Writing Strategy and Tactics

Before differentiating, it is necessary to organize the prerequisite inputs. The following are generally required when considering an architecture strategy.

**1. Confirmation of Business Strategy and Vision**
- Understanding business goals and strategies: How architecture contributes to the business
- Business drivers: Key elements for business growth, sources of competitive advantage
- Identifying stakeholders: Requirements, constraints, and expectations from management, development, and business departments

**2. Organizing Requirements**
- Functional requirements: Functions the system must fulfill
- Non-functional requirements: Performance, availability, security, maintainability, etc.

**3. Understanding the Current Architecture (As-Is)**
- Investigating the current system: System configuration, data flow, technologies used, operational processes
- Extracting issues and risks: Bottlenecks, technical debt, reliance on specific individuals, operational load, etc.

**4. Confirming Constraints**
- Resource constraints: Budget, personnel, deadlines
- Organizational structure and team setup: Who is responsible for which area

In addition to these, necessary inputs may vary depending on the situation, such as regulatory and compliance requirements, external dependencies, and past architecture decision records. It is important to gather a certain amount of information before formulating a strategy.

Attempting to differentiate without clarifying these will result in documents that are merely formalities. It is crucial to start by organizing the inputs without being bound by methods.

## Common Patterns of Confusion

Here are two common patterns of confusion when differentiating between strategy and tactics.

**1. Writing Means (How) in Strategy**

For example, a case where "adopting microservices" is included in the strategy. Microservices are a means and should be written in tactics.

What should be written in strategy is the purpose of "why microservices are necessary." For example, "to reduce dependencies between teams and enable independent delivery."

**2. Losing Sight of Purpose by Focusing on Means**

When thinking about direction with the means of microservices in mind, the purpose of "why microservices should be adopted" becomes ambiguous.

By separating means from strategy, the abstraction level and firmness of the strategy as a policy can be maintained.

## Differentiation Criteria Using 5W1H

When assigning 5W1H to strategy and tactics, it can be organized as follows.

**Elements to Write in Strategy**

- **Why**: Purpose, motivation, business value. Why is this architectural change necessary?
- **What**: Goals to achieve, issues to solve. What do we want to realize?

**Elements to Write in Tactics**

- **How**: Broad outline of measures. How will we achieve it?
- **When**: Timeline, phases, priorities
- **Where**: Scope of application, which parts of the system it applies to
- **Who**: Responsible teams, stakeholders

**Elements to Write in Design**

- Details of How: Specific technical selection, interface design, data models
- Detailed requirements and constraints: Specific figures for non-functional requirements, technical constraints
- Trade-off considerations: Comparison of adopted design and alternatives

This organization clarifies the role division where strategy indicates "direction," tactics indicate "broad outline of measures," and design indicates "specific implementation policies."

## Gray Areas When in Doubt

Not everything can be clearly classified. For example, writing a context-level architecture as a to-be in strategy.

High-level descriptions like "transition from monolith to distributed architecture" are strategy-oriented, but when specific system boundaries are shown as a context-level diagram, tactical elements are also included.

**Judgment Criterion: "Is the means not being assumed?"**

When in doubt, check whether the description assumes means. If means are assumed, the purpose is easily lost.

By separating means from strategy, the firmness of the policy can be maintained. Means can be flexibly changed depending on the situation, but the purpose is less likely to waver.

## Example of Formulating Strategy, Tactics, and Design

Here is an example of differentiation using the three-layer structure.

Note that the Why/What of strategy needs to be logically derived from problem analysis. If problem analysis is ambiguous, Why/What will also be baseless.

### Problem Analysis

Currently, it takes an average of 4 weeks to release a feature. The main causes are the following three points:

- Strong dependencies between modules, where one change affects a wide range
- The need to test and deploy everything together, causing waiting times
- The need for agreement between multiple teams for release coordination, leading to high communication costs

The most significant impact is the dependencies between modules, which also affect the other two causes.

**Impact on Business**

Due to release delays, three major projects were lost to competitors in the past year. Sales have voiced concerns that "it is difficult to propose without a clear outlook on feature additions."

**Alignment with Business Strategy**

The company-wide strategy emphasizes "quick response to market changes," and improving delivery speed is directly related to this.

### Strategy (Why, What)

**Why: Why undertake this**

It takes an average of 4 weeks to release a feature, delaying responses to market changes. Competitors release in 2 weeks, and at this rate, competitiveness will be lost. Improving delivery speed is the top priority to realize the company-wide strategy of "quick response to market changes."

**What: What to achieve**

Reduce dependencies between modules, allowing teams to independently change and deploy. This will double the current delivery speed, reducing feature release time to an average of 2 weeks.

### Tactics (How, When, Where, Who)

**How: How to achieve it**

- Separate the order service from the monolith, making it independently deployable
- Separate the inventory service from the monolith
- Build CI/CD pipelines for each service

**When: By when, in what order**

- Q1: Separation of order service
- Q2: Separation of inventory service
- Q3: Establishment of CI/CD pipelines and effectiveness measurement

**Where: Where to apply**

- Target: Order and inventory areas (areas with the most complex dependencies)
- Non-target: Other areas to be considered in the next phase

**Who: Who will be responsible**

- Order service: Order team + Platform team
- Inventory service: Inventory team + Platform team

### Design

**Design Overview of Order Service**

- Communication method: Integration with the monolith via REST API. Considering transition to asynchronous messaging in the future
- Data separation: Order data to be migrated to a dedicated DB. Bidirectional synchronization during the migration period
- Interface: Published as Order API v1. Adopting a versioning policy that maintains backward compatibility

**Architecture Diagram, Use Cases, Sequence**

(Diagrams omitted)

- Architecture diagram: Shows the relationship between the order service, monolith, and external systems
- Use cases: Main use cases such as order registration, order inquiry, order cancellation
- Sequence diagram: Integration flow with the monolith during order registration

**Non-functional Requirements**

- Availability: 99.9% (monthly downtime within 43 minutes)
- Response time: Within 200ms at p95
- Throughput: 1,000 requests/second at peak

**Trade-offs**

- Synchronous vs Asynchronous communication: Initially adopting simple synchronous communication. Reducing complexity and enabling early release
- Shared DB vs Dedicated DB: Adopting a dedicated DB. Prioritizing independent deployment despite data consistency challenges

## Conclusion

By differentiating strategy, tactics, and design, the roles of each become clear.

- **Strategy** writes **Why, What**. Focus on purpose and what you want to achieve
- **Tactics** writes the broad outline of **How, When, Where, Who**. Indicate the direction of measures
- **Design** writes the details of **How**. Show specific technical selection and implementation policies
- When in doubt, judge by "Is the means not being assumed?"

By separating means from strategy and leaving details to design, each document can be maintained at an appropriate level of abstraction.