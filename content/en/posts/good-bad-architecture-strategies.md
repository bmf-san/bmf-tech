---
title: Good Architecture Strategy, Bad Architecture Strategy
slug: good-bad-architecture-strategies
date: 2026-02-05T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture Strategy
  - Architecture
  - Design
description: Exploring the differences between effective and ineffective architecture strategies.
translation_key: good-bad-architecture-strategies
---



Even when an architecture strategy is written, it may not function effectively. It can become a mere formality or fail to be executed.

This article organizes the differences between good and bad strategies.

## Key Points of a Good Strategy

A good strategy consists of three elements: problem analysis, policy, and measures. This is an interpretation of the "kernel" defined in [Good Strategy, Bad Strategy](https://www.amazon.co.jp/%E8%89%AF%E3%81%84%E6%88%A6%E7%95%A5%E3%80%81%E6%82%AA%E3%81%84%E6%88%A6%E7%95%A5-%E3%83%AA%E3%83%81%E3%83%A3%E3%83%BC%E3%83%89%E3%83%BBP%E3%83%BB%E3%83%AB%E3%83%A1%E3%83%AB%E3%83%88/dp/4532318092?__mk_ja_JP=%E3%82%AB%E3%82%BF%E3%82%AB%E3%83%8A&crid=2IB2OV7EHVGTT&dib=eyJ2IjoiMSJ9._qSh-Id8vGXdYJQ3JKC30OtS9EGOuyJpfzOqmObX6zByIOFzGbv3asWYJFmUp-sKwpObwCmAZfuqMQDfM4oDiYhZmFa4fh0LnP_Otez8cl5bNc_wCuOX9XQzGPOhGmxiWg6n9PAK-ls08QpJ0dBQ_3ZQv4rJS-RjB6yAb_g7Z_sMhVdzJM3OxyzubIAt-KqqPLn4RHkx2lW77F8lgJRFToDWLn3zjzRrpzIqJgHsepSauENxN9JY2QNSdxmZJSg-W_tiyaROFkBWkam9PXLciGS5q6JfUtFR7ugMTO1IqIsNMm9cE0shQkAAhU888JAEiJCeveqQgXYcZya8FG3mng.TIBpMkJQlh0XGpUai_ncLgHIBWC6lqBlUs0OMWVTAqo&dib_tag=se&keywords=%E8%89%AF%E3%81%84%E6%88%A6%E7%95%A5&qid=1739582701&s=books&sprefix=%E8%89%AF%E3%81%84%E6%88%A6%E7%95%A5,stripbooks,160&sr=1-1&linkCode=sl1&tag=bmf035-22&linkId=f9ef460dab41888af747289b20281d3c&language=ja_JP&ref_=as_li_ss_tl) in the context of architecture strategy.

- **Problem Analysis**: Analyze the current issues and clarify what the problems are
- **Policy**: Indicate the direction to tackle the issues
- **Measures**: Define the broad framework of what to do (specific implementation methods are handled as tactics)

With these three elements, it becomes clear what needs to be done, allowing for optimal resource allocation. It enables focus and concentration in various decision-making processes, such as determining priorities, choosing methods, defining scope, selecting technologies, designing organizations, allocating personnel, and budgeting.

## Characteristics of a Bad Strategy

Conversely, a bad strategy has the following characteristics:

- **Vague and Lacks Specificity**: Only abstract goals or slogans are present, with no clear actions
- **Avoids Major Issues**: Avoids addressing fundamental problems, sticking to superficial solutions
- **Ends with Just Setting Goals**: The desired state is written, but there is no path to achieve it

These can be said to be states where one or more of the three elements of problem analysis, policy, and measures are missing.

### Causes of a Bad Strategy

- **Insufficient Problem Analysis**: Failing to correctly understand the current situation. Asking the wrong questions. Unable to identify fundamental problems
- **Avoiding Difficult Decisions**: Reluctant to take responsibility for deciding what not to do. Unwilling to clarify trade-offs. Aiming for conclusions that everyone agrees on, leading to vague expressions
- **Confusing Goals with Strategy**: Mistaking goals like "improving performance" for strategy. Lacking policy and measures, only having goals
- **Lack of Experience Writing Strategies**: Not knowing what to write. Lacking an image of a good strategy

## Importance of Problem Analysis in Strategy

Among the three elements of a good strategy, problem analysis is the most important. Since policy and measures are derived from the results of problem analysis, if this is wrong, everything else will be wrong.

Things to be aware of in problem analysis:

**Ask the Right Questions**

Instead of "How to transition to microservices," ask "Why does it take so long to change the system?" Avoid starting with the means. Starting with the means can prevent verification of whether the means are truly necessary, potentially overlooking fundamental issues.

Example: Proceeding with microservices from the start, but the actual problem was the complexity of dependencies between modules. Even after splitting services, dependencies were not resolved, and only the complexity of the distributed system increased.

**Structure the Issues**

Distinguish between superficial symptoms and root causes, and organize the relationships between issues. Addressing superficial symptoms will lead to recurrence if root causes remain. Also, if the relationships between issues are not visible, it is impossible to correctly judge priorities.

Example: In response to the symptom "Deployment takes time," investigating the causes revealed multiple factors such as "unstable tests," "inefficient CI pipeline," and "monolithic requiring full builds." By organizing these relationships, it becomes clear where to start.

**Identify Fundamental Issues**

Do not try to address everything, but identify the most impactful issues. Resources are limited, and trying to address everything leads to dispersion and half-hearted results.

Example: Multiple issues such as "performance problems," "technical debt," "scalability," and "developer experience" were raised. By evaluating the impact on the business and the difficulty of addressing them, it was decided to focus on performance issues first.

## Examples of Good and Bad Strategies

Even with the same theme, the quality of a strategy can change significantly depending on how it is written. Below are before and after examples.

### Example of a Bad Strategy (Before)

**Strategy: Transition to Microservices**

We will move away from a monolithic architecture and transition to a microservices architecture. This will improve scalability and development efficiency, enhancing competitiveness.

- Complete the transition to microservices by the end of 2025
- Aim for a state where each team can deploy independently
- Adopt the latest technology stack

### Example of a Good Strategy (After)

**Strategy: Improve Delivery Speed**

**Problem Analysis**

Currently, it takes an average of 4 weeks to release a feature. The main causes are as follows:

- Strong dependencies between modules, causing widespread impact from a single change
- Need to test and deploy everything together, causing wait times
- Release coordination requires agreement between multiple teams, leading to high communication costs

The most significant impact is from module dependencies, which also affect the other two causes.

**Policy**

Aim to reduce module dependencies and enable teams to make changes and deploy independently.

What to do:
- Separate the most complex order and inventory domains
- Build deployment pipelines on a team basis

What not to do:
- Simultaneous microservices transition for all services
- Refresh the technology stack (proceed with the current one)

**Measures**

- Separate services for order and inventory domains
- Establish a team-based deployment system

## Conclusion

A good strategy consists of the three elements of problem analysis, policy, and measures. With these three elements, it becomes clear what needs to be done, allowing for optimal resource allocation.

A bad strategy lacks these elements. Causes include insufficient problem analysis, avoiding difficult decisions, and confusing goals with strategy.

Problem analysis is particularly important. Ask the right questions, structure the issues, and identify fundamental issues. If this is wrong, both policy and measures will proceed in the wrong direction.