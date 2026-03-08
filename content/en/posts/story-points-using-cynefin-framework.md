---
title: Understanding Story Points Using the Cynefin Framework
slug: story-points-using-cynefin-framework
date: 2025-02-26T00:00:00Z
author: bmf-san
categories:
  - Development Process
tags:
  - Agile
  - Scrum
description: An idea on how to think about the criteria for story points.
translation_key: story-points-using-cynefin-framework
---



I came up with an idea on how to think about the criteria for story points, so I decided to write it down.

I thought it might be the first in the world!? But then I found an article called [Demystifying Story Point Estimation Using the Cynefin Framework](https://medium.com/agileinsider/demystifying-story-point-estimation-using-the-cynefin-framework-3549983feefd), so it seems not.

# What are Story Points?
Story Points are units used to measure the "**relative amount of work, complexity, and uncertainty**" of a task. Unlike general "time (man-hours)", they are estimated by considering **the difficulty, risk, and uncertainty** of the work.

The value of the points themselves has no absolute meaning and is used as a **relative scale** such as "Task B seems about twice as hard as Task A". Many teams adopt the Fibonacci sequence (1, 2, 3, 5, 8, 13...) because the increasing difference in numbers makes it easier to express "uncertainty" and "risk".

# The Concept of Story Points
The fundamental premise when dealing with story points is the idea of **"viewing work in terms of difficulty and uncertainty, not time"**. While the actual man-hours required may be estimated separately, when setting story points, the basis is **"the anticipated difficulty"**.

Moreover, story points are not an **absolute standard** but a **relative standard** within the team. If the team changes or the development environment changes, the difficulty assessment of the same task may change significantly. Therefore, it is important to review and update the team's understanding each time.

# Benefits of Well-Managed Story Points
Properly managing story points can lead to **improved planning** as follows:

- **Increased likelihood of achieving sprint goals**
  (Easier for the team to understand what is achievable)
- **Easier alignment on task understanding**
  (Easier for members to have a common understanding of "how difficult this task is")
- **Easier task division and priority decisions**
  (Facilitates decision-making such as splitting large tasks to reduce risk)

## Story Point Criteria Using the Cynefin Framework
The method I came up with while writing this article is to use the Cynefin Framework to classify tasks and assign story points accordingly, allowing the **difficulty and risk of tasks to be naturally reflected in the estimation**.

### What is the Cynefin Framework?
The Cynefin Framework is a framework that classifies tasks or situations into **five domains** for decision-making and problem-solving.

- **Obvious**: Tasks with clear best practices that can be patterned
- **Complicated**: Tasks requiring expertise but solvable through analysis
- **Complex**: Tasks with unpredictable outcomes requiring trial and error
- **Chaotic**: Tasks needing urgent response with no immediate solution in sight
- **Confused**: Tasks that don't fit any domain and are completely unclear

### Story Point Criteria Using the Cynefin Framework
After classifying tasks into each domain, set a **rough point guideline** as shown in the table below.

| Cynefin Domain        | Characteristics                                                        | How to Determine Story Points                                                                                 |
| --------------------- | --------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| **Obvious**           | Clear tasks that can be patterned                                     | **1-2 points**<br>(Quick estimation based on past similar tasks)                                              |
| **Complicated**       | Requires analysis or expertise but solutions are visible              | **3-5 points**<br>(Estimate through consensus using planning poker, etc.)                                     |
| **Complex**           | Unpredictable outcomes requiring trial and error                      | **8 points or more**<br>(Conduct spikes (investigation tasks) and further subdivide tasks before estimating)  |
| **Chaotic**           | High urgency with many unknown risks, requiring immediate action      | **13 points or more**<br>(Prioritize urgent response over story points. Estimate after the task stabilizes)   |
| **Confused**          | Completely confused state with no clear starting point                | **Cannot assign story points**<br>(Need to organize tasks and identify the domain)                            |

### Flow of Story Point Determination
1. **Determine the characteristics of the task**
- **Is the solution immediately apparent?** → **Obvious**
- **Can it be solved with analysis or expertise?** → **Complicated**
- **Are outcomes unpredictable, requiring trial and error?** → **Complex**
- **Is there high urgency with many unknown risks?** → **Chaotic**
- **Does it not fit any category?** → **Confused** (Need to organize the task first)

2. **Choose an estimation method based on classification**
- **Obvious:** Quickly estimate based on past achievements
- **Complicated:** Estimate through consensus using planning poker
- **Complex:** Conduct spikes (investigation tasks) and subdivide tasks before estimating
- **Chaotic:** Prioritize urgent response over assigning story points
- **Confused:** First organize task content and classify

By changing the estimation process according to the nature of the task, it becomes easier to reflect complexity and urgency in story points.

## Conclusion
The Cynefin Framework is a useful framework for organizing issues and seems to be a good match for story point estimation.

I am currently trying it out in the field, and I will add any insights if I notice anything.