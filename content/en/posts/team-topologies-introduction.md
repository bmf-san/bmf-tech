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

Team Topologies is **"an adaptive organizational design model aimed at maximizing the flow of business value."**

What sets it apart from traditional organizational design (hierarchical organization charts) is its prioritization of the following points:

- Emphasizing dynamic "flow" rather than static "hierarchy"
- Considering "cognitive load" as a constraint in design
- Thinking of software architecture and organizational structure as a set

The goal is to create a state where each team can operate autonomously by organizing inter-team dependencies and reducing unnecessary communication (coordination costs).

## Why Team Topologies?

Why do traditional functional organizations (like frontend, backend, and infrastructure teams) often fail? Two powerful laws are at play here.

### Conway's Law and the "Reverse Conway Maneuver"

The law proposed by Melvin Conway in 1968 is still valid today.

> "Organizations that design systems are constrained to produce designs that are copies of the communication structures of these organizations."

In other words, if an organization is siloed, the systems will also be fragmented, making integration difficult. Conversely, if you want to create loosely coupled microservices, you must first divide the organization into loosely coupled independent teams. This is referred to as the **"Reverse Conway Maneuver."**

### Cognitive Load Theory: The Limits of Teams

Another important concept is "cognitive load." There is a limit to the complexity that the human brain (working memory) can handle at once. Team Topologies categorizes this load into three types and aims to reduce "extraneous cognitive load."

- **Intrinsic Load**: Essential loads related to the job, such as how to write Java or business logic.
- **Extraneous Load**: Non-value-adding loads like overly complex deployment procedures, difficult-to-use testing environments, or unclear internal processes about whom to ask.
- **Germane Load**: Loads related to seeking better solutions. More resources should be allocated here.

### Cognitive Load and "Lack of Ownership"

What happens when cognitive load exceeds a team's capacity?

The team becomes unable to fully understand the entire system. Consequently, anxiety begins to dominate, leading to a mindset of "I don't want to touch parts that are working" or "What if I break something during deployment?" As a result, there is a **"lack of ownership"** regarding the system.

An attitude of "I only know my area" spreads, leading to handoffs in incident response and dramatically slowing down delivery speed.

The aim of Team Topologies is to help teams regain healthy ownership by setting appropriate team sizes and areas of responsibility.

## Four Team Types

To avoid chaotic team formations and manage cognitive load effectively, Team Topologies recommends classifying organizations into only four types.

### Stream-aligned Team

- **Role**: Teams positioned along the primary value stream of the business.
- **Characteristics**: They hold end-to-end responsibility (ownership) for the entire flow from planning, development, testing, to operations. They aim to minimize "waiting for requests" from other teams and are the main players in delivering value autonomously. In a healthy organization, 80-90% of all teams will be of this type.

### Platform Team

- **Role**: Teams that provide infrastructure and tools as "self-service" to enable Stream-aligned Teams to operate autonomously.
- **Characteristics**: They view internal developers as "customers" and provide a user-friendly, low-cognitive-load platform (Thinnest Viable Platform). They are not just "infrastructure operators" but teams that create "products" to enhance the productivity of product teams.

### Enabling Team

- **Role**: Expert teams that fill gaps in specific technical areas (such as security, test automation, AI, etc.).
- **Characteristics**: Rather than performing tasks for others, they temporarily join Stream-aligned Teams to provide technical guidance and coaching, thereby enhancing the team's capabilities. They are teams that "teach how to fish rather than just giving fish."

### Complicated Subsystem Team

- **Role**: Teams that handle only those parts that require high specialization, such as advanced mathematical models or image processing engines, which exceed the cognitive load limits of regular teams.
- **Characteristics**: They are established as exceptional measures to reduce cognitive load and should not be increased casually.

## Three Interaction Modes

It is necessary to intentionally design not only the "shape" of teams but also how they "interact."

| Mode | Description | Suitable Situations |
|------|-------------|---------------------|
| **Collaboration** | Two teams closely cooperate to solve challenges together. | When introducing new technologies or during the exploratory phase when API specifications are not yet defined. |
| **X-as-a-Service** | One team provides APIs or tools that the other team utilizes. | When the platform is mature and can be used based solely on documentation. This has the lowest cognitive load. |
| **Facilitation** | One team supports the other and removes obstacles. | When an Enabling Team assists other teams in their learning. |

The key point is that these modes are not fixed and should be adapted over time. Initially, you might start with "Collaboration" to build together, and once stable, transition to "X-as-a-Service," requiring dynamic changes.

## Conclusion

Team Topologies is not a "new organization chart" that you implement once and forget.

As the phase of the product changes, so do the necessary team structures and optimal interactions. Regularly questioning "Are the current team boundaries appropriate?" and "Is cognitive load becoming too high?" and continuously refactoring the organization (organizational sensing) is the key to maintaining a fast flow.

Starting discussions with your team about which type they fit into and whether ownership is being compromised due to cognitive load may lead to valuable insights.

## References
- [teamtopologies.com](https://teamtopologies.com/)
- [Team Topologies: Organizing Business and Technology Teams for Fast Flow](https://www.amazon.co.jp/%E3%83%81%E3%83%BC%E3%83%A0%E3%83%88%E3%83%9D%E3%83%AD%E3%82%B8%E3%83%BC-%E4%BE%A1%E5%80%A4%E3%81%82%E3%82%8B%E3%82%BD%E3%83%95%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A2%E3%82%92%E3%81%99%E3%81%B0%E3%82%84%E3%81%8F%E5%B1%8A%E3%81%91%E3%82%8B%E9%81%A9%E5%BF%9C%E5%9E%8B%E7%B5%84%E7%B9%94%E8%A8%AD%E8%A8%88-%E3%83%9E%E3%82%B7%E3%83%A5%E3%83%BC%E3%83%BB%E3%82%B9%E3%82%B1%E3%83%AB%E3%83%88%E3%83%B3/dp/4820729632?_encoding=UTF8&dib_tag=se&dib=eyJ2IjoiMSJ9.4p_dyd_ycwKiOCaBXACsm_H0YFaWOznzU5VKtO1npXOxvNgSHMJRDp8heUUdnJCXAEYo4UlMISyBRyKDMlTzQNEAY-NtY18xpR7-QXGfquTDSknKdiUu4NcF4F3aTb0gZU61_0bTV7aBpqpTsR56MzGU0Jk67jQZ7LV_l0Pj4sQ-tmgMUqKxPUW84gDxjKHPRgB8orcYOcJIN-RTQCaTst6vzZiOj4l6DpUsqLBbop0.yCQPemK2CN8V4V5nwxW-I3QpH4sHWdfVHt-mb-kbv-M&qid=1732807586&sr=8-2-spons&linkCode=sl1&tag=bmf035-22&linkId=81bdb0e94f0d9c515a2c24dfed8835dc&language=ja_JP&ref_=as_li_ss_tl)
