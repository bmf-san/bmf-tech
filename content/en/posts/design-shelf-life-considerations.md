---
title: Considering the Expiration Date of Design
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
In system design, there are "mistakes" but no "correct answers." What exists is an "optimal compromise" based on the circumstances at the time. Design is the act of making decisions within various constraints and shaping the future.

Through the perspective of "design expiration date," we will explore how to estimate the lifespan of a design and how to confront constraints.

# What is the Design Expiration Date?
"How long should a design last?"

By being aware of this question, design becomes more realistic. Some areas of design should be made to last, while others can be intentionally short-lived.

Considering the expiration date leads to the "acceptance of constraints." All design is conducted within constraints.

# Design is a Dialogue with Constraints
Design is essentially an expression of constraints. If there were unlimited time, budget, personnel, and future prospects, there would be no constraints. However, in reality, constraints such as "who will maintain it," "how long will it be used," and "what can change" always exist.

The judgment of "this design should last for X years" is an act of imposing **intentional constraints** on the design. This leads to avoiding over-design and fosters flexible thinking that anticipates future replacements.

# Perspectives on Design Expiration Date
Let's consider specific examples from several perspectives.

## 1. Business Perspective: Speed of Change and Uncertainty
- During the startup phase, hypothesis testing is prioritized, and a design lifespan of 1-2 years may be sufficient.
- Mature businesses require stability, and a lifespan of 3-5 years may be adequate.

## 2. Organizational Perspective: Team Structure and Personnel Mobility
- Highly specialized designs may have shorter lifespans.
- Designs prepared for fluctuations in team skill sets and numbers may last longer.

## 3. Product Perspective: Stability and Evolution of Features
- Frequently changing features may not require long-lived designs.
- Core functions that are less likely to change may need long-lived designs.

## 4. Technical Perspective: Evolution of Technology Stack and Dependencies
- The end of maintenance for OSS or obsolescence of libraries may impact the lifespan of designs.
- Areas with rapid technology update cycles may find it challenging to create long-lived designs.

# Giving Expiration Dates to Design
By assigning an "expiration date" to design, the following effects can be achieved:

- Decision-making becomes more realistic.
- Subsequent maintenance and replacements become easier.
- Risks of over-design can be avoided.
- Expectations for the design align.

This also clarifies the "scope of responsibility" regarding the design.

# Practical Inquiry List
- How many years do you want this design to last?
- What are the constraints of this design? (Business, Organization, Product, Technology)
- What are the compromise points in the design? Who is it reasonable for?
- Who will inherit this design in the future?

# Conclusion
- The design expiration date refers to the **validity period of constraints** in design.
- Perspectives on expiration dates can vary across business, organization, product, and technology.
- Considering the design expiration date helps clarify decision-making and avoid over-design.

I believe that code becomes a liability from the moment it is written, but how long we consider it acceptable to fall behind business growth (i.e., the period of tolerating liabilities, the expiration date) provides good insights for design.

In reality, various factors may cause the expiration date to be shorter or longer than expected, but having the perspective of expiration dates as a hint for considering constraints is beneficial.