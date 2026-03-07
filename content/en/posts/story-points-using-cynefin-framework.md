---
title: Using the Cynefin Framework for Story Point Estimation
slug: story-points-using-cynefin-framework
date: 2025-02-26T00:00:00Z
author: bmf-san
categories:
  - Development Process
tags:
  - Agile
  - Scrum
translation_key: story-points-using-cynefin-framework
---

I had an idea about how to think about the criteria for story points, so I wanted to document it.

I thought it might be the first in the world!? But I found an article titled [Demystifying Story Point Estimation Using the Cynefin Framework](https://medium.com/agileinsider/demystifying-story-point-estimation-using-the-cynefin-framework-3549983feefd), so it seems I am not alone.

# What are Story Points
Story Points are a unit for measuring the **relative amount of work, complexity, and uncertainty** involved in a task. Unlike the common measure of **time (man-hours)**, it is characterized by considering **the difficulty, risk, and uncertainty** of the work.

The value of the points themselves has no absolute meaning; they are used as a **relative measure** such as **"Task B seems about twice as difficult as Task A."** Many teams often adopt the Fibonacci sequence (1, 2, 3, 5, 8, 13…) because the larger differences in numbers make it easier to represent **uncertainty** and **risk**.

# Thinking About Story Points
A fundamental premise when dealing with story points is the idea of **"looking at the difficulty and uncertainty of the work, not time."** While it may be necessary to estimate how much effort is actually required separately, when setting story points, the **"expected difficulty"** should be the basis.

Additionally, story points are not an **absolute standard**, but rather a **relative standard** within the team. When the team changes or the development environment changes, the evaluation of the same task's difficulty may vary significantly. Therefore, it is important to review and update the team's understanding each time such changes occur.

# Benefits of Properly Managed Story Points
When story points are used appropriately, it can lead to improved **planning capabilities** as follows:

- **Increased likelihood of achieving sprint goals**  
  (The team can better grasp what is achievable)
- **Easier alignment of understanding regarding the work**  
  (Members can more easily share a common understanding of how difficult a task is)
- **Easier decision-making regarding task division and prioritization**  
  (Decisions such as splitting large tasks to reduce risk become smoother)

## Using the Cynefin Framework for Story Point Criteria
One method I thought of while writing this article is to classify tasks using the Cynefin Framework and assign story points accordingly, which allows for a natural reflection of **task difficulty and risk in the estimation**.

### What is the Cynefin Framework
The Cynefin Framework is a framework for classifying tasks or situations into **five domains** in decision-making and problem-solving contexts.

- **Obvious**: Work with clear best practices that can be patterned
- **Complicated**: Work that requires expertise but can be solved through analysis
- **Complex**: Work where outcomes are unpredictable and require trial and error
- **Chaotic**: Work that requires urgent response and where solutions are not immediately visible
- **Confused**: A state that does not fit into any domain and is completely unclear

### Story Point Criteria Using the Cynefin Framework
After classifying tasks into each domain, set **rough point guidelines** as shown in the table below.

| Cynefin Domain          | Characteristics                                                  | How to Determine Story Points                                                                                 |
| ----------------------- | --------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| **Obvious**             | Clear work that can be patterned                                | **1-2 points**  
  (Quickly estimate based on past similar tasks)                                             |
| **Complicated**         | Requires analysis or expertise, but solutions can be found     | **3-5 points**  
  (Estimate while reaching consensus through planning poker, etc.)                                   |
| **Complex**             | Outcomes are unpredictable and require trial and error          | **8 points or more**  
  (Conduct spikes (research tasks) and estimate after further breaking down the task)                |
| **Chaotic**             | Requires urgent response with many unknown risks                | **13 points or more**  
  (Prioritize urgent response over story points. Estimate after the task stabilizes) |
| **Confused**            | Completely unclear where to start                               | **Cannot assign story points**  
  (Tasks need to be organized and domains identified)                 |

### Flow of Determining Story Points
1. **Assess the characteristics of the task**
- **Is a solution immediately clear?** → **Obvious**
- **Can it be solved with analysis or expertise?** → **Complicated**
- **Is it unpredictable and requires trial and error?** → **Complex**
- **Is it urgent with many unknown risks?** → **Chaotic**
- **Does it fit none of these?** → **Confused** (Tasks need to be organized first)

2. **Choose estimation methods based on classification**
- **Obvious:** Quickly estimate based on past performance
- **Complicated:** Estimate while reaching consensus through planning poker
- **Complex:** Conduct spikes (research tasks) and estimate after breaking down tasks
- **Chaotic:** Prioritize urgent response before assigning story points
- **Confused:** First, organize and classify task content

By changing the estimation process according to the nature of the task, it becomes easier to reflect complexity and urgency in story points.

## Conclusion
The Cynefin Framework is a useful tool for organizing challenges and seems to align well with story point estimation.

I am currently trying to implement this in the field, and I would like to add any insights I gain.