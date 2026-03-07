---
title: What is Team Topology?
slug: team-topologies-introduction
date: 2026-01-09T00:00:00Z
author: bmf-san
categories:
  - Management
tags:
  - Team Topology
  - Team Management
  - Organization Design
translation_key: team-topologies-introduction
---

## What is Team Topology?

Team Topology is **"an adaptive organizational design model aimed at maximizing the flow of business value."**

It fundamentally differs from traditional organizational design (hierarchical organizational charts) by prioritizing the following points:

- Emphasizing dynamic "flow" rather than static "hierarchy"
- Considering "Cognitive Load" as a constraint in design
- Thinking of software architecture and organizational structure as a set

The goal is to create a state where each team can operate autonomously by organizing inter-team dependencies and reducing unnecessary communication (coordination costs).

## Why Team Topology?

Why do traditional functional organizations (like front-end, back-end, infrastructure teams) not work effectively? Two powerful laws are at play here.

### Conway's Law and the "Reverse Conway Maneuver"

The law proposed by Melvin Conway in 1968 is still valid today.

> "Organizations that design systems are constrained to produce designs that are copies of the communication structures of these organizations."

In other words, if an organization is siloed, the systems will also be fragmented, making integration difficult. Conversely, if you want to create loosely coupled microservices, you must first divide the organization into loosely coupled independent teams. This is called the **"Reverse Conway Maneuver."**

### Cognitive Load Theory: Team Limitations

Another important concept is "Cognitive Load." There is a limit to the complexity that the human brain (working memory) can handle at one time. Team Topology classifies this load into three categories and aims to reduce "Extraneous Load."

- **Intrinsic Load**: Essential loads for work, such as Java syntax and business logic.
- **Extraneous Load**: Loads that do not create value, such as overly complex deployment procedures, difficult testing environments, and unclear internal processes.
- **Germane Load**: Loads aimed at finding better solutions. More resources should be allocated here.

### Cognitive Load and "Lack of Ownership"

What happens when cognitive load exceeds a team's capacity?

The team becomes unable to fully understand the entire system. Anxiety begins to dominate, leading to thoughts like, "I don't want to touch parts that are working" or "What if I break something during deployment?" As a result, there is a **"lack of ownership (sense of ownership and responsibility)"** towards the system.

An attitude of "I only know my own area" spreads, leading to handoffs in incident response and dramatically decreasing delivery speed.

The aim of Team Topology is to restore healthy ownership by setting appropriate team sizes and areas of responsibility.

## Four Team Types

To avoid chaotic team formations and manage cognitive load appropriately, Team Topology recommends classifying organizations into only four types:

### Stream-aligned Team

- **Role**: Teams aligned with the main value stream of the business.
- **Characteristics**: They have end-to-end responsibility (ownership) for the entire flow from planning, development, testing, to operations. They aim to minimize waiting for requests from other teams and are the main actors in delivering value autonomously. In a healthy organization, 80-90% of all teams should be of this type.

### Platform Team

- **Role**: Teams that provide infrastructure and tools as "self-service" to enable Stream-aligned Teams to operate autonomously.
- **Characteristics**: They treat internal developers as "customers" and provide a user-friendly, low cognitive load platform (Thinnest Viable Platform). They are not just "infrastructure operators" but teams that create "products" to enhance the productivity of product teams.

### Enabling Team

- **Role**: Expert teams that fill gaps in specific technical areas (security, test automation, AI, etc.).
- **Characteristics**: Instead of doing the work for others, they temporarily join Stream-aligned Teams to provide technical guidance and coaching, thereby enhancing the team's capabilities. They are teams that "teach how to fish rather than just giving fish."

### Complicated Subsystem Team

- **Role**: Teams that handle highly specialized areas, such as advanced mathematical models or image processing engines, which exceed the cognitive load limits of regular teams.
- **Characteristics**: They are established as exceptional measures to reduce cognitive load and should not be increased lightly.

## Three Interaction Modes

It is necessary to intentionally design not only the "shape" of teams but also how they interact.

| Mode | Description | Suitable Situations |
|------|-------------|--------------------|
| **Collaboration** | Two teams closely cooperate to solve challenges together. | During the introduction of new technologies or in exploratory phases where API specifications are not yet determined. |
| **X-as-a-Service** | One team provides APIs or tools that the other team utilizes. | When the platform is mature and can be used with just documentation. This has the lowest cognitive load. |
| **Facilitation** | One team supports the other and removes obstacles. | When the Enabling Team supports the learning of other teams, for example. |

The important thing is that these modes are not fixed and should be adapted over time. Initially, you may start with "Collaboration" to build together, and once stable, transition to "X-as-a-Service," requiring dynamic changes.

## Conclusion

Team Topology is not a "new organizational chart" that is implemented once and forgotten.

As the phase of the product changes, the necessary team shapes and optimal interactions also change. Regularly questioning whether "the current team boundaries are appropriate?" and "is cognitive load becoming too high?" and continuously refactoring the organization (organizational sensing) is key to maintaining a fast flow.

Starting discussions with your team about which type they fall into and whether ownership is being compromised due to cognitive load may lead to valuable discoveries.

## References
- [teamtopologies.com](https://teamtopologies.com/)
- [Team Topologies: Adaptive Organization Design for Delivering Value Quickly](https://www.amazon.co.jp/%E3%83%81%E3%83%BC%E3%83%A0%E3%83%88%E3%83%9D%E3%83%AD%E3%82%B8%E3%83%BC-%E4%BE%A1%E5%80%A4%E3%81%82%E3%82%8B%E3%82%BD%E3%83%95%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A2%E3%82%92%E3%81%99%E3%81%B0%E3%82%84%E3%81%8F%E5%B1%8A%E3%81%91%E3%82%8B%E9%81%A9%E5%BF%9C%E5%9E%8B%E7%B5%84%E7%B9%94%E8%A8%AD%E8%A8%88-%E3%83%9E%E3%82%B7%E3%83%A5%E3%83%BC%E3%83%BB%E3%82%B9%E3%82%B1%E3%83%AB%E3%83%88%E3%83%B3/dp/4820729632?_encoding=UTF8&dib_tag=se&dib=eyJ2IjoiMSJ9.4p_dyd_ycwKiOCaBXACsm_H0YFaWOznzU5VKtO1npXOxvNgSHMJRDp8heUUdnJCXAEYo4UlMISyBRyKDMlTzQNEAY-NtY18xpR7-QXGfquTDSknKdiUu4NcF4F3aTb0gZU61_0bTV7aBpqpTsR56MzGU0Jk67jQZ7LV_l0Pj4sQ-tmgMUqKxPUW84gDxjKHPRgB8orcYOcJIN-RTQCaTst6vzZiOj4l6DpUsqLBbop0.yCQPemK2CN8V4V5nwxW-I3QpH4sHWdfVHt-mb-kbv-M&qid=1732807586&sr=8-2-spons&linkCode=sl1&tag=bmf035-22&linkId=81bdb0e94f0d9c515a2c24dfed8835dc&language=ja_JP&ref_=as_li_ss_tl)