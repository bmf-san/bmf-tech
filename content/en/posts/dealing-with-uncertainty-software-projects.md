---
title: Dealing with Uncertainty in Software Development Projects
slug: dealing-with-uncertainty-software-projects
date: 2023-12-22T00:00:00Z
author: bmf-san
categories:
  - Development Process
tags:
  - Agile
translation_key: dealing-with-uncertainty-software-projects
---

This article is the 22nd entry of the [Makuake Advent Calendar 2023](https://adventar.org/calendars/8992).

# Overview
I will articulate how I confront uncertainty in software development projects on a daily basis.

Since I haven't systematically learned by reading a lot of Agile-related books or studying PMBOK, this will be a subjective account based on my experiences.

# Premise
I want to discuss the context in which I face uncertainty.

I belong to a team that designs, develops, and operates technical infrastructure, and I serve as the team leader (responsible for leading the project and acting as the architect within the team).

The projects my team undertakes have the following characteristics:

- In-house development
- Development of systems that will be used by other systems
- Relatively large scale of development (not something that can be completed in 1-2 months)
- Requires a medium to long-term period (around six months or more)
- No definitive deadline is set (deadlines are determined based on requirements but are not absolute)

I believe that the approach to uncertainty may differ for teams engaged in contract development or those closer to end-users, but I think there are fundamental aspects that apply to any project. (Probably)

# What is Uncertainty?
Before articulating how to confront uncertainty, I want to confirm the definition of uncertainty.

Uncertainty generally refers to a **state where it is difficult or impossible to accurately predict future events or situations**.

It seems reasonable to consider that the uncertainty in software development projects fits this definition.

There are various factors contributing to uncertainty in projects, but I believe they can be broadly categorized into "external factors" and "internal factors".

## External Factors
I consider external factors to be those that exist outside the project and influence it.

These include market changes, regulations, natural disasters, and other factors that are difficult or impossible to control.

Factors that **cannot be adjusted even if project members or stakeholders try their best** are classified as external factors.

## Internal Factors
Internal factors are those within the project, such as the skills of project members, the project schedule, and the project budget—factors that project members or stakeholders can control.

I believe that the aspects the team should confront in a project are mostly these internal factors.

I think there may also be factors that could be positioned between external and internal factors, but it is necessary to assess whether they can be controlled and whether they need to be prioritized in addressing them.

The classification of external and internal factors may not be crucial; what matters most is whether **actions can be taken to address or prevent uncertainty**.

# What Happens When Uncertainty is High or Low?
I believe that the level of uncertainty affects **the planning of the project**.

Conversely, if the planning of the project is high, it is likely that uncertainty has been eliminated, and if it is low, it is likely that uncertainty has not been eliminated. (There may be research on this, but I haven't searched for sources, so this is just my theory.)

High planning leads to several benefits, such as optimization of resources (people, money, time), flexibility to changes (plans, specifications, anything), and improved trust among stakeholders.

I believe that eliminating uncertainty as much as possible and enhancing planning is one of the keys to project success.

As a prerequisite for eliminating uncertainty, it is necessary to **constantly observe and evaluate planning**.

For example, preparing visual tools like roadmaps or Gantt charts that allow the team to grasp, observe, and evaluate project progress in line with team activities.

At the timing of evaluating project planning (perhaps the Sprint Retrospective event in Scrum is just right), reflecting on what influenced planning, what is needed for improvement, etc., and taking actions to enhance planning in line with project progress. I believe that clarifying what was uncertain is a learning opportunity for the team to enhance planning.

I wanted to bring up a practical example, but it seems like it would require a lot of sharing as a premise, so I will skip it.

The timing for evaluating project planning may vary by team, but in my case, I try to observe and evaluate constantly. If I can quickly notice something that is likely to cause delays (i.e., something with high uncertainty), I believe I can take advantage of time to review plans or adjust tasks.

# Identifying Uncertainty
I have talked about how eliminating uncertainty can enhance planning and lead to happiness, but how can we discover the uncertainty itself?

In my daily work, I consciously focus on **finding what is not known**, and if I dig a little deeper, I think it essentially boils down to identifying **known unknowns** and **unknown unknowns**.

## Known Unknowns
Known unknowns are the things we know we do not know.

For example, when integrating an existing system whose specifications are somewhat understood with another system, we can see the APIs to be used and their calling sequences, but we do not know what kind of issues may arise when the integration actually occurs.

Known unknowns are **easier to address because we can think about how to make the unknown known**.

## Unknown Unknowns
Unknown unknowns are the things we do not know we do not know.

An example would be when a request for a sudden change in specifications arises during the course of the project.

Unknown unknowns are **difficult to address because they are unexpected**. It requires investment-type responses, such as making multi-faceted predictions and having time buffers.

I believe it is important to always keep an eye out for discovering and sensing both known unknowns and unknown unknowns within the project.

I quickly jotted down a few ideas on how to discover and sense them:

- Understanding and observing the project situation
- Intuition based on experience
- Gut feeling
- Borrowing ideas from general software development principles

I think I might be engaging in deductive or inductive reasoning, but I haven't organized my thoughts enough to articulate them, so I will make it a future task.

# Addressing Uncertainty
Once uncertainty is identified, it is necessary to consider how to address it.

I will articulate some ideas I have regarding how to address the identified uncertainty:

- Understand the nature of the problem
  - To address uncertainty-related problems, consider what steps are necessary to understand the problem first.
  - I recently learned that the Cynefin Framework is a useful reference for this kind of thinking.
- Strategically postpone
  - It is premised on being in a situation where postponement is acceptable.
  - I think postponing in cases where "I do not know now, but there is a possibility I will know later" is an effective strategy.
  - It is also valid to take provisional measures and then postpone.
  - By the way, I resonate with the line from Freiren in "Sousou no Freiren" Episode 13, "I have postponed my decisions too much."
- Learn
  - Since the types and nature of uncertainty vary depending on the domain, technology area, and other environments the team is involved in, conducting verification and learning within the project (i.e., Agile) is essential.
  - This is more of a retrospective approach rather than a direct response, serving as a precautionary measure for similar issues that may arise in the future.

I think it is case-by-case, so I couldn't come up with many ideas...

# Examples of Reducing Uncertainty and Enhancing Planning in Projects
As an example of addressing uncertainty, I will discuss task planning.

- Task decomposition
- Definition of done
- Formation of common understanding

## Task Decomposition
By breaking tasks down into smaller pieces, it becomes easier to estimate, potentially increasing accuracy.

If task dependencies are organized, it also becomes easier to consider parallelizing tasks. (This does not apply if the decomposition of tasks is low.)

If tasks are large, the recognition of goals and what needs to be done can easily become blurred, so I want to eliminate such uncertainties as much as possible.

## Definition of Done
Clarifying what state a task must reach to be considered complete helps in tracking progress, ensuring quality, and optimizing resources (which tasks to prioritize, how much time to invest, etc.).

If it is difficult or impossible to define the completion criteria for a task, it may indicate that the granularity of task decomposition is not appropriate or that uncertainty may be lurking.

## Formation of Common Understanding
When completing a set of tasks, there is uncertainty due to the potential for misalignment in common understanding among team members.

I believe that being conscious of forming a common understanding while considering the granularity of task decomposition can lead to the elimination of uncertainty.

For example, suppose there is a task to conduct a survey.

ex. "Conduct a survey on XX. Summarize the survey results in a document and decide on XX based on the results."

Points of potential misalignment in this task include the perspectives of the survey, the format of the survey results document, the validity of decisions made based on the survey results, and the criteria for judgment.

This task can be broken down into:

- Deciding on survey perspectives
- Conducting the survey
- Sharing and discussing survey results
- Creating an ADR

By setting milestones with a common understanding through task decomposition, I believe we can eliminate misalignment (uncertainty).

This is just one example, and I think the appropriate approach may vary by team. (The above is an example from my own team.)

# Thoughts
I have articulated my approach to dealing with uncertainty, but I feel that I relied more on intuition than I expected.

If there is an opportunity to learn systematically somewhere, I would like to organize my thoughts again.

# References
- [en.wikipedia.org - Cone of Uncertainty](https://en.wikipedia.org/wiki/Cone_of_Uncertainty)