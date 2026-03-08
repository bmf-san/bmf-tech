---
title: Difference Between Requirements and Constraints
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

In software development, have you ever struggled with the difference between "Requirements" and "Constraints"?

I have sometimes confused these two concepts when considering design.

I felt it was important to clearly understand these differences for proper design, so I organized them based on the definitions from the international standard of systems engineering, **ISO/IEC/IEEE 29148**.

## Definitions

According to the international standard, these two are distinguished as follows:

- **Requirements**  
  - Definition: Descriptions of the "functions" or "capabilities" that stakeholders demand from the system.  
  - Essence: A vector towards the goal of "What do we want to achieve?"  
  - Characteristics: Negotiable. For example, a requirement like "I want a search function" can be adjusted to "move to the next phase" due to budget or schedule constraints.

- **Constraints**  
  - Definition: "Conditions" or "factors" that limit the freedom of design or implementation.  
  - Essence: A cutting line of the solution space that represents "Limits/Bounds" that must be adhered to.  
  - Characteristics: Generally non-negotiable, or the cost of change is extremely high. They are forces imposed from outside, such as "must integrate with legacy systems" or "must comply with GDPR".

### Supplement: Why Use "Requirements"

**Because it is the standard vocabulary in the field**  
In Japanese development environments, there is a custom to differentiate between stakeholders' "raw wishes" as **"Requests"** and the translated and confirmed version in engineering terms as **"Requirements"** (e.g., requirement definition instead of request definition). When discussing design, we deal with confirmed "Requirements," making this term appropriate.

**To make the contrast with "Constraints" clearer**  
- **Request**: "I want you to do ~" (desire)  
- **Requirement**: "It must be ~" (condition)  
- **Constraint**: "It has to be ~" (mandatory)  

As shown above, Requests, Requirements, and Constraints are arranged in a gradient of force. When discussing design, treating it as a confirmed Requirement rather than an ambiguous Request clarifies the contrast with Constraints, making trade-off decisions easier.

**ISO's Structural Intent**  
In ISO 29148, "Stakeholder Needs" (Requests) and "System Requirements" are clearly distinguished. Since the context of this article is "design," it pertains to the stage of Requirements rather than Needs.

## How to Distinguish Requirements from Constraints

A simple framework for distinguishing Requirements from Constraints can use the following four questions:

| Question                                   | If the answer is YES... | Reason                                              |
| ------------------------------------------ | ----------------------- | --------------------------------------------------- |
| **"Is it a user desire?"**               | **Requirement**         | Because it is the value that the system provides.   |
| **"Does it take away design options?"**  | **Constraint**          | Because it is a factor that restricts engineers' freedom in technology selection. |
| **"Can it be changed with more money or time?"** | **Requirement**         | Because it can be a subject of trade-offs.         |
| **"Is it a physical law, legal requirement, or company policy?"** | **Constraint**          | Because it is an "immovable wall" outside the development team's authority. |

## Non-Functional Requirements (NFRs) Act as Constraints

A bit more complicated are "Non-Functional Requirements (NFRs)" such as performance and security. Formally, these are Requirements, but for architects, they function as powerful **Constraints**.

For example, consider an NFR that states, "The page must be displayed within one second."

- For users, this is a **Requirement** of wanting a "fast experience."  
- However, for engineers, it becomes a **Constraint** that narrows design, meaning "prohibition of using heavy frameworks" or "mandatory introduction of a cache layer."

## Conclusion

Software design can be likened to a **"puzzle of finding solutions that maximize Requirements within the framework of Constraints."**

- Requirements are challenges that should be solved with creativity.  
- Constraints are the rules (the ring) when exercising creativity.

However, Constraints are not always "walls that cannot be moved." As design progresses, if Constraints significantly hinder the realization of Requirements or explode costs, negotiating to change the Constraints itself is also an important job for architects.

By clearly distinguishing these two at the start of a project, you can accurately identify "what cannot be changed" and "what should be changed (or negotiated)," allowing the team's energy to focus on optimal trade-off decisions.

# References
- [tracery.jp - What is a Requirement](https://tracery.jp/articles/entry/what-is-a-requirement)  
- [drkasbokar.com - ISO/IEC/IEEE 29148](https://drkasbokar.com/wp-content/uploads/2024/09/29148-2018-ISOIECIEEE.pdf)