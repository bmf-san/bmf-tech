---
title: "What Is AI-DLC (AI-Driven Development Life Cycle)?"
slug: ai-driven-development-life-cycle
date: 2026-06-15T00:00:00Z
author: bmf-san
categories:
  - Development Process
tags:
  - AI
  - AI-DLC
  - Generative AI
  - SDLC
translation_key: ai-driven-development-life-cycle
draft: false
---

## Introduction

**AI-DLC (AI-Driven Development Lifecycle)**, proposed by AWS, is a new development approach that positions AI as a central collaborator in the development process. This article briefly summarizes its ideas and mechanics.

Traditionally, many organizations have bolted AI onto their existing processes as an "assistant." AWS argues that this usage constrains AI's capabilities and preserves outdated inefficiencies. AI-DLC instead aims to embed AI into the very fabric of software development.

## Two Core Principles

AI-DLC, as an AI-centric approach, emphasizes two points.

- **AI executes, humans supervise**: AI creates detailed work plans and proactively asks for intent alignment and guidance, while humans—who understand the business context—make the critical decisions.
- **Dynamic team collaboration**: While AI handles routine tasks, the team focuses on real-time problem solving, creative thinking, and fast decision making.

These two principles aim to increase development speed without sacrificing quality.

## The Underlying Mental Model

At the heart of AI-DLC is a mental model in which AI initiates and drives the workflow.

1. AI creates a plan.
2. AI asks clarifying questions to understand the context.
3. Only after human validation does AI build the solution.

By rapidly repeating this pattern across every activity in the SDLC (Software Development Lifecycle), the team applies a unified vision and approach throughout development.

## The Three Phases

AI-DLC organizes development into three simple phases.

| Phase | Description |
|---|---|
| Inception | AI transforms business intent into requirements, stories, and units of work through "Mob Elaboration," where the whole team validates AI's questions and suggestions. |
| Construction | Based on the context validated during Inception, AI proposes the logical architecture, domain model, implementation, and tests through "Mob Construction." |
| Operation | Applying the accumulated context, AI manages Infrastructure as Code and deployment under the team's supervision. |

Each phase hands richer context to the next. AI stores plans, requirements, and design artifacts in the project repository, maintaining persistent context across many sessions.

## Renamed Terminology

Reflecting its collaborative approach, AI-DLC replaces traditional agile terms.

- **Sprints → Bolts**: Shorter work cycles measured in hours or days rather than weeks.
- **Epics → Units of Work**

These changes emphasize a focus on speed and continuous delivery.

## Expected Benefits

AWS highlights the following benefits of AI-DLC.

- **Development speed**: AI rapidly generates and refines artifacts—requirements, designs, code, and tests—turning work that once took weeks into hours or days.
- **Quality**: Continuous intent alignment lets teams build exactly what they envision rather than relying on ambiguous interpretation, and makes it easier to consistently apply organization-specific standards (coding conventions, design patterns, security requirements).
- **Innovation**: With AI taking on the heavy lifting, builders have more time to explore creative solutions.
- **Responsiveness**: Rapid development cycles enable quick adaptation to market demand and feedback.
- **Developer experience**: Shifting focus from routine coding to meaningful problem solving reduces cognitive load.

## Conclusion

AI-DLC reframes AI not as a mere assistant but as a primary actor in development. Its defining feature is applying the mental model—"AI plans, humans verify, AI builds"—consistently across all three phases. When adopting it, it helps to reference the AWS white paper, Amazon Q Developer rules, and Kiro custom workflows while tailoring the approach to your own organization's process.

## References

- [AI-Driven Development Lifecycle: Reimagining Software Engineering (AWS Blog, JP)](https://aws.amazon.com/jp/blogs/news/ai-driven-development-life-cycle/)
- [AI-Driven Development Lifecycle (AI-DLC) Method Definition (Raja SP, AWS)](https://prod.d13rzhkk8cj2z0.amplifyapp.com/)
