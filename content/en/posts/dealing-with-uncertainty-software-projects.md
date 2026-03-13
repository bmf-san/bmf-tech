---
title: How to Deal with Uncertainty in Software Development Projects
description: 'Address software project uncertainty through agile planning. Distinguish known unknowns from unknown unknowns to minimize risks.'
slug: dealing-with-uncertainty-software-projects
date: 2023-12-22T00:00:00Z
author: bmf-san
categories:
  - Development Process
tags:
  - Agile
translation_key: dealing-with-uncertainty-software-projects
---



This article is the 22nd entry (a poem) for the [Makuake Advent Calendar 2023](https://adventar.org/calendars/8992).

# Overview
I will try to articulate how I deal with uncertainty in software development projects on a daily basis.

Since I haven't systematically studied by reading a lot of Agile-related books or studying PMBOK, this will be a subjective discussion based on my experience.

# Premise
I want to proceed with the discussion based on the premise of the environment in which I face uncertainty.

I belong to a team that designs, develops, and operates technical infrastructure, and I serve as the team leader (responsible for leading projects and acting as an architect within the team).

The projects our team undertakes have the following characteristics:

- In-house development
- Development of systems to be used by other systems
- Relatively large-scale development (not something that can be completed in 1-2 months)
- Requires a medium to long-term period (six months or more)
- No fixed deadlines (dates are set according to requirements, but not absolute)

In teams that engage in contract development or development closer to end-users, the approach to uncertainty might differ, but I believe the fundamental aspects are applicable to any project. (Probably)

# What is Uncertainty?
Before articulating how to deal with uncertainty, I want to confirm the definition of uncertainty.

Uncertainty generally refers to a "**state where future events or situations cannot be accurately predicted or are difficult to predict**."

It seems reasonable to think that the uncertainty in software development projects fits this definition.

The causes of uncertainty in projects are varied, but I believe they can be broadly divided into "external factors" and "internal factors."

## External Factors
External factors are those outside the project that affect it.

These include market changes, regulations, natural disasters, and other factors that are difficult or impossible to control.

Factors that project members or stakeholders **cannot adjust despite their efforts** are classified as external factors.

## Internal Factors
Internal factors are those within the project, such as the skills of project members, the project schedule, and the project budget, which project members or stakeholders can control.

I believe that the team should generally focus on these internal factors in a project.

There may be factors that seem to be positioned between external and internal factors, but I think it is necessary to determine whether they can be controlled and whether they need to be prioritized for action.

The examples might not have been the best, but the classification into external and internal is not important; the most important consideration is whether **actions can be taken to address or prevent uncertainty**.

# What Happens When Uncertainty is High or Low?
I believe that high or low uncertainty affects **the planning of a project**.

Conversely, if a project is in a highly planned state, it is likely that uncertainty has been eliminated, and if it is in a low state, it is likely that uncertainty has not been eliminated. (There might be research on this, but I haven't looked for sources, so this is just my theory.)

High planning leads to several benefits, such as optimizing resources (people, money, time, etc.), flexibility in changes (plans, specifications, anything), and improved trust with stakeholders.

Eliminating uncertainty as much as possible and enhancing planning is one of the keys to project success.

As a prerequisite for eliminating uncertainty, it is necessary to **always be able to observe and evaluate planning**.

For example, prepare visual tools like roadmaps or Gantt charts that allow the team to visually grasp, observe, and evaluate project progress according to their activities.

At the timing of evaluating project planning (in Scrum, the sprint retrospective event might be just right), reflect on what affected planning, what is needed for improvement, and take actions to enhance planning according to project progress. Clarifying what was uncertain is considered a learning opportunity for the team to enhance planning.

I wanted to provide an example from practice, but it seems like it would take a lot of effort to share the prerequisites, so I'll omit it.

The timing of evaluating project planning may vary by team, but in my case, I always try to observe and evaluate. I believe that if you notice something that might cause delays (≒ something with high uncertainty) early, you can review plans and adjust work with a time advantage.

# Identifying Uncertainty
I've talked about how eliminating uncertainty can enhance planning and lead to happiness, but let's consider how to actually discover uncertainty.

In projects, I usually focus on "**finding what is unknown**," but when I delve a little deeper, I think it's about finding "**known unknowns**" and "**unknown unknowns**."

## Known Unknowns
"Knowing that you don't know something" is a known unknown.

For example, when integrating an existing system with another system, you may see the APIs to use and their call sequences, but you don't know what issues might arise when actually integrating.

Known unknowns are **easier to address because you can think about how to understand what you don't know**.

## Unknown Unknowns
"Not knowing that you don't know something" is an unknown unknown.

An example would be a sudden specification change request during project progress.

Unknown unknowns are **difficult to address because they are unexpected**. It requires investment-like approaches such as making multifaceted predictions in advance or having time buffers.

I believe it is important to always keep an eye out for "known unknowns" and "unknown unknowns" in projects.

I've written down a few things that come to mind on how to actually discover and detect them.

- Understanding and observing project status
- Intuition based on experience
- Instinct
- Borrowing ideas from general software development theories

I think I might be using deductive or inductive thinking processes, but I haven't organized my thoughts enough to articulate them, so I'll make it a future task.

# Addressing Uncertainty
Once you find uncertainty, you need to consider how to address it.

I've tried to articulate some ideas I have for addressing discovered or detected uncertainty.

- Understanding the nature of the problem
  - To consider what steps are necessary to address the problem of uncertainty, first understand the problem
  - I recently learned about the Cynefin Framework, which I think is a useful reference for this way of thinking
- Strategically postponing
  - Assuming it's a situation where postponing is acceptable
  - Postponing is an effective means when "it is unknown now but may become known later"
  - It's also possible to take provisional measures and then postpone
  - By the way, the line "I postponed my decision too much" from 'Frieren: Beyond Journey's End' Episode 13 "Self-Loathing" resonates with me
- Learning
  - Since the types and nature of uncertainty vary depending on the domain, technical area, and other environments the team is involved in, conduct verification and learning within the project (≒ Agile)
  - This is more of a post-event discussion and is more of an insurance approach to be able to address similar issues in the future

I couldn't think of much, probably because I consider it case-by-case...

# Example of Reducing Uncertainty and Enhancing Planning in Projects
As an example of addressing uncertainty, let's take task planning.

- Task Breakdown
- Definition of Done
- Building Common Understanding

## Task Breakdown
Breaking down tasks into smaller parts makes estimation easier, potentially increasing accuracy.

If task dependencies are also organized, it becomes easier to consider task parallelization.
(If task breakdown is not possible, this may not apply.)

If tasks are large, the goals and understanding of what to do in the task can easily become misaligned, so I want to eliminate such uncertainty as much as possible.

## Definition of Done
Clearly defining when a task is complete helps in tracking progress, ensuring quality, and optimizing resources (which tasks to prioritize, where to invest time, etc.).

If it's difficult or impossible to define the completion of a task, it may indicate that the task breakdown granularity is not appropriate, or there may be hidden uncertainty.

## Building Common Understanding
When completing a certain set of tasks, there is uncertainty in the form of misalignment in common understanding among team members.

When considering task breakdown granularity, being conscious of building common understanding can help eliminate uncertainty.

For example, suppose there is a task to conduct some investigation.

ex. "Conduct an investigation of XX. Summarize the investigation results in a document and decide on XX based on the results."

Points of misalignment in such a task include the investigation perspective, the format of the investigation results document, the validity and criteria for deciding on XX based on the investigation results.

This task can be broken down into:

- Deciding on investigation perspectives
- Conducting the investigation
- Sharing and consulting on investigation results
- Creating an ADR

By setting milestones with common understanding, it can lead to eliminating misalignment (uncertainty).

This is just an example, so the appropriate approach may vary by team. (The above is an actual example from my team.)

# Impressions
I tried to articulate how to deal with uncertainty, but I felt I relied more on intuition than I thought.

If I have the opportunity to learn systematically somewhere, I would like to organize my thoughts again.

# References
- [en.wikipedia.org - Cone of Uncertainty](https://en.wikipedia.org/wiki/Cone_of_Uncertainty)
