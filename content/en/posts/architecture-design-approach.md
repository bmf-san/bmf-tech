---
title: How to Approach Architecture Design
slug: architecture-design-approach
date: 2023-11-23T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture
  - Design
translation_key: architecture-design-approach
---

# Overview
I would like to summarize a personal approach that I think is good for designing architecture at the level of configuration diagrams.

# Approach
Assuming that requirements definition and research have been somewhat completed in advance, I will write about the approach to designing while writing the architecture configuration diagram.

In simple terms, it is just about "not trying to draw the final goal from the beginning, but designing step by step."

![image](https://github.com/bmf-san/bmf-san/assets/13291041/f59af100-af82-42c2-973a-57ce4d74d92e)

<img style="width:450px;!important" alt="architecture" src="https://github.com/bmf-san/bmf-san/assets/13291041/f59af100-af82-42c2-973a-57ce4d74d92e">

By writing the configuration diagram step by step based on the requirements, I believe there are the following benefits. (The above diagram is just an example and is roughly drawn without considering a specific system.)

- It allows for step-by-step thinking, making it easier to organize thoughts.
  - Why is that component necessary? What are the pros and cons of this configuration, and what are the trade-offs? Where might the challenges lie? It becomes easier to think about these questions.
- It becomes easier to align understanding among members.
  - Even things that seem self-evident in design can lead to misalignment of understanding if each member has different knowledge and experience. Even for obvious parts, I want to clarify the reasons (the devil is in the details).
- It helps prevent excessive design.
  - If you try to design the final form from the beginning, you might end up questioning, "Wait, do we really need to build it to that extent?" This may depend on experience and skills...

# Conclusion
I believe this approach is particularly useful when there are many people involved in the design.

Aligning people's understanding is one of the highly uncertain tasks, so I think it is effective to think while gradually aligning assumptions.