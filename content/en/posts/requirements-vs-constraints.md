---
title: Difference Between Requirements and Constraints
description: An in-depth look at Difference Between Requirements and Constraints, covering key concepts and practical insights.
slug: requirements-vs-constraints
date: 2026-01-09T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Constraints
  - Requirements
  - Requests
  - Design
  - Architecture
translation_key: requirements-vs-constraints
---



Have you ever struggled with the difference between "Requirements" and "Constraints" in the field of software development?

When considering design, I sometimes confused these two concepts.

To conduct appropriate design, I felt it was important to clearly understand these differences, so I organized them based on the definitions from the international standard of systems engineering, **ISO/IEC/IEEE 29148**.

## Definitions

According to international standards, these two are distinguished as follows:

- **Requirements**
  - Definition: Descriptions of the "functions" or "capabilities" that stakeholders demand from the system.
  - Essence: A vector towards the goal of "What do we want to achieve?"
  - Characteristics: Negotiable. For example, the requirement "I want a search function" can be adjusted to "postpone to the next phase" due to budget or schedule constraints.

- **Constraints**
  - Definition: "Conditions" or "factors" that limit the freedom of design and implementation.
  - Essence: A cutting line in the solution space, "Limits/Bounds" that must be adhered to.
  - Characteristics: Principally non-negotiable or extremely costly to change. It is an external force such as "mandatory integration with legacy systems" or "GDPR compliance".

### Supplement: Why "Requirements"?

**Because it is the standard vocabulary on site**

In Japanese development sites, there is a custom to differentiate between stakeholders' "raw wishes" as **"Needs/Requests"** and those translated and finalized into engineering terms as **"Requirements"** (e.g., requirement definition instead of request definition). When discussing design, the finalized "requirements" are dealt with, making it appropriate.

**To make the contrast with "Constraints" clear**

- **Request**: "I want you to do ~" (wish)
- **Requisite**: "It needs to be ~" (condition)
- **Constraint**: "It must be ~" (compulsion)

As shown above, requests, requirements, and constraints are lined up in a gradient of compulsion. When discussing design, handling it as a confirmed requirement of "It needs to be ~" rather than an ambiguous request level of "I want you to do ~" makes the contrast with constraints clear and facilitates trade-off decisions.

**ISO's Structural Intent**

In ISO 29148, "Stakeholder Needs" and "System Requirements" are clearly distinguished. The context of this article is "design," so it is a discussion at the requirements stage, not the needs stage.

## How to Distinguish Requirements and Constraints

To distinguish between requirements and constraints, you can use the following four questions as a simple framework.

| Question                                   | If the answer is YES... | Reason                                              |
| ------------------------------------------ | ----------------------- | --------------------------------------------------- |
| **"Is it a user's desire?"**               | **Requirement**         | Because it is the value provided by the system.     |
| **"Does it take away design options?"**   | **Constraint**          | Because it is a factor that prevents engineers from freely selecting technology. |
| **"Can it be changed with money or time?"** | **Requirement**         | Because it can be a subject of trade-offs.          |
| **"Is it a physical law, legal, or company-wide regulation?"** | **Constraint** | Because it is an "immovable wall" outside the development team's authority.   |

## Non-Functional Requirements (NFR) Act as "Constraints"

A bit more complicated are "Non-Functional Requirements (NFRs)" such as performance and security. Formally, these are requirements, but for architects, they function as powerful **"Constraints"**.

For example, consider an NFR like "display the page within 1 second."

- For users, this is a **requirement** of "I want a fast experience."
- However, for engineers, it becomes a **constraint** that narrows design, meaning "prohibition of using heavy frameworks" or "mandatory introduction of a cache layer."

## Conclusion

Software design can be likened to a **"puzzle of finding solutions that maximally satisfy requirements (Requirements) within the framework of constraints (Constraints)."**

- Requirements are challenges to be solved with creativity.
- Constraints are the rules (platform) when exercising creativity.

However, constraints are not necessarily "immovable walls." As design progresses, if constraints significantly hinder the realization of requirements or explode costs, "negotiating and changing" the constraints themselves is also an important job for architects.

By clearly distinguishing these two at the start of a project, you can correctly identify "what cannot be changed" and "what should be changed (or negotiated)," allowing the team to focus their energy on optimal trade-off decisions.

# References
- [tracery.jp - What is a Requirement](https://tracery.jp/articles/entry/what-is-a-requirement)
- [drkasbokar.com - ISO/IEC/IEEE 29148](https://drkasbokar.com/wp-content/uploads/2024/09/29148-2018-ISOIECIEEE.pdf)
