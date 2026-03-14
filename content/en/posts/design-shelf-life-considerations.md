---
title: Considering the Shelf Life of Design
description: 'Consider design shelf life across business, organization, product, and technical perspectives. Balance constraints and trade-offs.'
slug: design-shelf-life-considerations
date: 2025-06-08T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Design
  - System Design
  - Architecture Strategy
  - Architecture
translation_key: design-shelf-life-considerations
---

# Overview
In system design, there may be "mistakes," but there is no "correct answer." What exists is the "optimal compromise" according to the situation at that time. Design is the act of making decisions within various constraints and giving shape towards the future.

By considering the "shelf life of design," we explore how to estimate the lifespan of a design and how to face constraints.

# What is the Shelf Life of Design?
"How long should a design last?"

By being conscious of this question, design becomes realistic. Some areas of design should have permanence, while others can be intentionally short-lived.

Considering the shelf life also leads to "acceptance of constraints." All designs are made within constraints.

# Design is a Dialogue with Constraints
Design is essentially an expression of constraints. If there were unlimited time, budget, personnel, and future prospects, there would be no constraints. However, in reality, constraints such as "who will maintain it," "how long will it be used," and "what can change" always exist.

The judgment that "this design should last for X years" is an act of giving **intentional constraints** to the design. It avoids over-engineering and brings flexible thinking that anticipates future replacements.

# Viewing the Shelf Life of Design from Different Perspectives
Let's consider specific examples from several perspectives.

## 1. Business Perspective: Speed of Change and Uncertainty
- In the startup phase, hypothesis testing is prioritized, and a design lifespan of 1-2 years might be sufficient.
- Mature businesses require stability, and a lifespan of 3-5 years might be sufficient.

## 2. Organizational Perspective: Team Structure and Personnel Fluidity
- Designs with high dependency on individuals may have a shorter lifespan.
- Designs prepared for changes in team skill sets and numbers may last longer.

## 3. Product Perspective: Stability and Evolution of Functions
- Frequently changing functions may be fine with "short-lived design."
- Core functions that change less may require "long-lived design."

## 4. Technical Perspective: Evolution and Dependencies of the Tech Stack
- The end of OSS maintenance or library obsolescence can affect the lifespan of a design.
- Areas with fast technology update cycles may find long-lived design difficult.

# Giving Design a Shelf Life
By giving design a "shelf life," the following effects can be achieved:

- Decision-making becomes realistic
- Subsequent maintenance and replacement become easier
- The risk of over-engineering is avoided
- Expectations for design are aligned

This is also an act of clarifying the "scope of responsibility" for design.

# Practical Inquiry List
- How long do you want this design to last?
- What are the constraints of this design? (Business, Organization, Product, Technology)
- What are the compromises in the design? Who are they reasonable for?
- Who will inherit this design in the future?

# Conclusion
- The shelf life of design is the **valid period of constraints** in design.
- There are various perspectives to consider the shelf life, such as business, organization, product, and technology.
- Considering the shelf life of design helps clarify decision-making and avoid over-engineering.

Code is something that becomes a liability the moment it is written, but considering how long it will take to become unable to keep up with business growth (≒ the period to tolerate liabilities, the shelf life) provides good insights for design.

In reality, various factors may cause the shelf life to be shorter or longer than expected, but having the perspective of shelf life as a hint to consider constraints is beneficial.