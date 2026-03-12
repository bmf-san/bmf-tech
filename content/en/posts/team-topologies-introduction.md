---
title: What is Team Topologies? How to Structure Engineering Teams
description: 'Learn Team Topologies—the four team types (stream-aligned, platform, enabling, complicated-subsystem) and three interaction modes for effective software delivery organization.'
slug: team-topologies-introduction
date: 2026-01-09T00:00:00Z
author: bmf-san
categories:
  - Management
tags:
  - Team Topologies
  - Team Management
  - Organizational Design
translation_key: team-topologies-introduction
---

## What is Team Topologies?

Team Topologies is an **"adaptive organizational design model to maximize the flow of business value."**

It fundamentally differs from traditional organizational design (hierarchical organizational charts) by prioritizing the following:

- Emphasizing dynamic "flow" over static "hierarchy"
- Considering "cognitive load" as a design constraint
- Thinking of software architecture and organizational structure as a set

The goal is to create a state where each team can operate autonomously by organizing inter-team dependencies and reducing unnecessary communication (coordination costs).

## Why Team Topologies?

Why do traditional function-based organizations (frontend, backend, infrastructure, etc.) not work well? Two powerful laws are at play here.

### Conway's Law and the "Reverse Conway Maneuver"

The law proposed by Melvin Conway in 1968 is still valid today.

> "Organizations which design systems are constrained to produce designs which are copies of the communication structures of these organizations."

In other words, if an organization is siloed, the system will also be fragmented, making integration difficult. Conversely, if you want to create loosely coupled microservices, you must first divide the organization into loosely coupled independent teams. This is called the **"Reverse Conway Maneuver."**

### Cognitive Load Theory: The Limits of Teams

Another important concept is "cognitive load." There is a limit to the complexity that the human brain (working memory) can handle at one time. Team Topologies classifies this load into three types and aims to reduce "extraneous load."

- **Intrinsic Load**: Essential load for work, such as how to write Java or business logic.
- **Extraneous Load**: Non-value-adding load, such as overly complex deployment procedures, difficult-to-use test environments, and unclear internal procedures.
- **Germane Load**: Load for seeking better solutions. More resources should be allocated here.

### Cognitive Load and "Lack of Ownership"

What happens when cognitive load exceeds a team's capacity?

The team can no longer fully understand the entire system. Then, anxiety like "I don't want to touch the working parts" or "What if it breaks when deployed?" begins to dominate. As a result, there is a **"lack of ownership (sense of ownership and responsibility)"** towards the system.

The attitude of "I don't know anything outside my responsibility" spreads, incident response is passed around, and delivery speed dramatically decreases.

The aim of Team Topologies is to restore healthy ownership by setting appropriate team sizes and responsibility scopes.

## Four Team Types

To avoid chaotic team formation and manage cognitive load appropriately, Team Topologies recommends classifying the organization into the following four types only.

### Stream-aligned Team

- **Role**: Teams aligned with the main value stream of the business.
- **Characteristics**: They have end-to-end responsibility (ownership) for the entire flow, from planning, development, testing, to operations. They minimize "waiting for requests" from other teams and are the main actors in delivering value autonomously. In a healthy organization, 80-90% of all teams are of this type.

### Platform Team

- **Role**: Teams that provide infrastructure and tools as "self-service" so that stream-aligned teams can operate autonomously.
- **Characteristics**: They view internal developers as "customers" and provide an easy-to-use, low-cognitive-load platform (Thinnest Viable Platform). They are not mere "infrastructure operators" but teams that create "products" to increase the productivity of product teams.

### Enabling Team

- **Role**: Specialist teams to fill gaps in specific technical areas (security, test automation, AI, etc.).
- **Characteristics**: Instead of doing the work for them, they temporarily join stream-aligned teams to provide technical guidance and coaching, thereby improving the team's capabilities. They are teams that "teach how to fish rather than give fish."

### Complicated Subsystem Team

- **Role**: Teams that handle only parts that are too specialized for normal teams to handle due to high cognitive load, such as advanced mathematical models or image processing engines.
- **Characteristics**: They are set up as an exceptional measure to reduce cognitive load. They should not be increased easily.

## Three Interaction Modes

Not only the "shape" of the team but also the "way of interaction" needs to be intentionally designed.

| Mode | Description | Suitable Situation |
|------|-------------|--------------------|
| **Collaboration** | Two teams closely cooperate to solve problems together. | When introducing new technology or during the exploratory phase when API specifications are not yet decided. |
| **X-as-a-Service** | One team uses the API or tools provided by another team. | When the platform is mature and can be used with just documentation. It has the lowest cognitive load. |
| **Facilitation** | One team supports another by removing obstacles. | When an enabling team supports the learning of other teams. |

The important thing is that these modes are not fixed but should be used differently depending on the time. Dynamic changes, such as starting with "Collaboration" to build together and then transitioning to "X-as-a-Service" once stable, are required.

## Conclusion

Team Topologies is not a "new organizational chart" that ends once implemented.

As the product phase changes, the necessary team shape and optimal interaction also change. Regularly questioning "Are the current team boundaries appropriate?" and "Is cognitive load becoming too high?" and continuously refactoring the organization (organizational sensing) is the key to maintaining a fast flow.

Starting by discussing with your team which type they fit into and whether "cognitive load" is undermining ownership might lead to valuable discoveries.

## References
- [teamtopologies.com](https://teamtopologies.com/)
- [チームトポロジー 価値あるソフトウェアをすばやく届ける適応型組織設計](https://www.amazon.co.jp/%E3%83%81%E3%83%BC%E3%83%A0%E3%83%88%E3%83%9D%E3%83%AD%E3%82%B8%E3%83%BC-%E4%BE%A1%E5%80%A4%E3%81%82%E3%82%8B%E3%82%BD%E3%83%95%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A2%E3%82%92%E3%81%99%E3%81%B0%E3%82%84%E3%81%8F%E5%B1%8A%E3%81%91%E3%82%8B%E9%81%A9%E5%BF%9C%E5%9E%8B%E7%B5%84%E7%B9%94%E8%A8%AD%E8%A8%88-%E3%83%9E%E3%82%B7%E3%83%A5%E3%83%BC%E3%83%BB%E3%82%B9%E3%82%B1%E3%83%AB%E3%83%88%E3%83%B3/dp/4820729632?_encoding=UTF8&dib_tag=se&dib=eyJ2IjoiMSJ9.4p_dyd_ycwKiOCaBXACsm_H0YFaWOznzU5VKtO1npXOxvNgSHMJRDp8heUUdnJCXAEYo4UlMISyBRyKDMlTzQNEAY-NtY18xpR7-QXGfquTDSknKdiUu4NcF4F3aTb0gZU61_0bTV7aBpqpTsR56MzGU0Jk67jQZ7LV_l0Pj4sQ-tmgMUqKxPUW84gDxjKHPRgB8orcYOcJIN-RTQCaTst6vzZiOj4l6DpUsqLBbop0.yCQPemK2CN8V4V5nwxW-I3QpH4sHWdfVHt-mb-kbv-M&qid=1732807586&sr=8-2-spons&linkCode=sl1&tag=bmf035-22&linkId=81bdb0e94f0d9c515a2c24dfed8835dc&language=ja_JP&ref_=as_li_ss_tl)