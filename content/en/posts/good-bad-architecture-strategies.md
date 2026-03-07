---
title: Good Architecture Strategy vs Bad Architecture Strategy
slug: good-bad-architecture-strategies
date: 2026-02-05T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture Strategy
  - Architecture
  - Design
translation_key: good-bad-architecture-strategies
---

Even if you write an architecture strategy, there are cases where it doesn't work. It may become a mere formality or may not be executed at all.

In this article, we will clarify the differences between a good strategy and a bad strategy.

## Key Points of a Good Strategy

A good strategy consists of three elements: problem analysis → policy → measures. This is an interpretation of the "kernel (core of the strategy)" defined in [Good Strategy, Bad Strategy](https://www.amazon.co.jp/%E8%89%AF%E3%81%84%E6%88%A6%E7%95%A5%E3%80%81%E6%82%AA%E3%81%84%E6%88%A6%E7%95%A5-%E3%83%AA%E3%83%81%E3%83%A3%E3%83%BC%E3%83%89%E3%83%BBP%E3%83%BB%E3%83%AB%E3%83%A1%E3%83%AB%E3%83%88/dp/4532318092?__mk_ja_JP=%E3%82%AB%E3%82%BF%E3%82%AB%E3%83%8A&crid=2IB2OV7EHVGTT&dib=eyJ2IjoiMSJ9._qSh-Id8vGXdYJQ3JKC30OtS9EGOuyJpfzOqmObX6zByIOFzGbv3asWYJFmUp-sKwpObwCmAZfuqMQDfM4oDiYhZmFa4fh0LnP_Otez8cl5bNc_wCuOX9XQzGPOhGmxiWg6n9PAK-ls08QpJ0dBQ_3ZQv4rJS-RjB6yAb_g7Z_sMhVdzJM3OxyzubIAt-KqqPLn4RHkx2lW77F8lgJRFToDWLn3zjzRrpzIqJgHsepSauENxN9JY2QNSdxmZJSg-W_tiyaROFkBWkam9PXLciGS5q6JfUtFR7ugMTO1IqIsNMm9cE0shQkAAhU888JAEiJCeveqQgXYcZya8FG3mng.TIBpMkJQlh0XGpUai_ncLgHIBWC6lqBlUs0OMWVTAqo&dib_tag=se&keywords=%E8%89%AF%E3%81%84%E6%88%A6%E7%95%A5&qid=1739582701&s=books&sprefix=%E8%89%AF%E3%81%84%E6%88%A6%E7%95%A5,stripbooks,160&sr=1-1&linkCode=sl1&tag=bmf035-22&linkId=f9ef460dab41888af747289b20281d3c&language=ja_JP&ref_=as_li_ss_tl) in the context of architecture strategy.

- **Problem Analysis**: Analyze the current issues and clarify what the problems are.
- **Policy**: Indicate the direction in which to tackle the issues.
- **Measures**: Define the broad outline of what to do (specific implementation methods are handled by tactics).

If these three elements are in place, it becomes clear what needs to be done, allowing for optimal resource allocation. This enables selection and concentration in various decision-making processes such as prioritization, means selection, scope determination, technology selection, organizational design, personnel allocation, and budget distribution.

## Characteristics of a Bad Strategy

As the opposite of a good strategy, a bad strategy has the following characteristics:

- **Empty and Lacks Specificity**: Only abstract goals or slogans are present, making it unclear what should be done.
- **Avoids Major Issues**: Evades essential problems and remains at a superficial level of response.
- **Ends with Just Setting Goals**: While the desired state is written, there is no pathway to reach it.

These can be seen as a state where one or more of the three elements of problem analysis → policy → measures are missing.

### Causes of Bad Strategy

- **Insufficient Problem Analysis**: Fails to accurately grasp the current situation. Misformulates questions. Cannot identify essential problems.
- **Avoiding Difficult Decisions**: Does not want to take responsibility for deciding "what not to do." Does not want to clarify trade-offs. In aiming for a conclusion that everyone agrees on, it becomes vague.
- **Confusing Goals with Strategy**: Thinks that goals like "improving performance" are strategies. Ends up with only goals without policies or measures.
- **Lack of Experience Writing Strategies**: Does not know what to write. Lacks an image of a good strategy.

## Importance of Problem Analysis in Strategy

Among the three elements of a good strategy, the most important is problem analysis. Policies and measures are derived from the results of problem analysis, so if this is wrong, everything else will be wrong.

Things to keep in mind in problem analysis:

**Formulate the Right Questions**

Instead of asking "How do we implement microservices?" ask "Why does it take so long to change the system?" Do not start from means. Starting from means can prevent verification of whether that means is truly necessary and may overlook essential issues.

Example: Proceeded with the assumption of microservices, but the actual problem was the complexity of dependencies between modules. Even if services are split, dependencies remain, adding only the complexity of a distributed system.

**Structure the Problems**

Distinguish between surface symptoms and root causes, and organize the relationships between issues. If you address surface symptoms but the root causes remain, they will recur. Also, if the relationships between issues are not visible, it will be difficult to judge priorities correctly.

Example: In response to the symptom of "it takes time to deploy," investigating the causes revealed multiple factors such as "unstable tests," "inefficient CI pipeline," and "monolith requiring a full build." By organizing these relationships, it becomes clear where to start.

**Identify the Essential Issues**

Do not try to address everything; identify the most impactful issues. Resources are limited, and trying to address everything will lead to dispersion and half-hearted efforts.

Example: Multiple issues such as "performance problems," "technical debt," "scalability," and "developer experience" were raised. Evaluated the business impact and difficulty of addressing each, and decided to focus first on performance issues.

## Examples of Good and Bad Strategies

Even with the same theme, the quality of the strategy can change significantly depending on how it is written. Below are before-and-after examples.

### Example of a Bad Strategy (Before)

**Strategy: Microservices Implementation**

We will transition from a monolithic architecture to a microservices architecture. This will improve scalability and development efficiency, enhancing our competitiveness.

- Complete microservices implementation by the end of 2025.
- Aim for each team to be able to deploy independently.
- Adopt the latest technology stack.

### Example of a Good Strategy (After)

**Strategy: Improving Delivery Speed**

**Problem Analysis**

Currently, it takes an average of 4 weeks to release features. The main causes are the following three points:

- Strong dependencies between modules, where one change has widespread effects.
- Need to test and deploy everything together, leading to waiting times.
- Releasing requires agreement among multiple teams, resulting in high communication costs.

The most significant impact comes from the dependencies between modules, which also ripple into the other two causes.

**Policy**

Aim to reduce dependencies between modules so that teams can independently make changes and deploy.

What to do:
- Separate the most complex dependencies in the order and inventory areas.
- Build deployment pipelines at the team level.

What not to do:
- Implement microservices for all services at once.
- Refresh the technology stack (continue with the current one).

**Measures**

- Separate services in the order and inventory areas.
- Establish a deployment system at the team level.

## Conclusion

A good strategy consists of the three elements of problem analysis → policy → measures. If these three elements are in place, it becomes clear what needs to be done, allowing for optimal resource allocation.

A bad strategy is a state where these elements are lacking. Insufficient problem analysis, avoiding difficult decisions, or confusing goals with strategy can all lead to this.

Especially, problem analysis is crucial. Formulate the right questions, structure the problems, and identify the essential issues. If this is wrong, both the policy and measures will head in the wrong direction.