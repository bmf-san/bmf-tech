---
title: Approach to Architecture Design
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
I will summarize an approach to architecture design at the configuration diagram level that I personally find effective.

# Approach
Assuming that requirements definition and research have been somewhat completed in advance, I will write about the method of proceeding with design while drawing architecture configuration diagrams.

In short, it is very simple: "Do not try to draw the final picture from the start, but design in stages."

![image](https://github.com/bmf-san/bmf-san/assets/13291041/f59af100-af82-42c2-973a-57ce4d74d92e)

<img style="width:450px;!important" alt="architecture" src="https://github.com/bmf-san/bmf-san/assets/13291041/f59af100-af82-42c2-973a-57ce4d74d92e">

By drawing configuration diagrams step by step based on requirements, as shown in the above diagram, I believe the following benefits can be achieved. (The above diagram is just an example and is roughly drawn, so it does not consider a specific system.)

- It becomes easier to organize thoughts because you can think in stages
  - Why is that component necessary? What are the advantages and disadvantages of this configuration, and what are the trade-offs? Where might there be issues? These questions become easier to consider.
- It becomes easier to align understanding among members
  - Even things that seem obvious in design can lead to misalignment if knowledge and experience differ among members. Even for obvious parts, I want to clarify the reasons (the devil is in the details).
- It can prevent over-designing
  - If you try to design the final form from the start, you might end up questioning, "Do we really need to build it out that much?" This might depend on experience and skills...

# Conclusion
I believe this approach is more useful the more people are involved in the design.

Aligning people's understanding is one of the high-uncertainty tasks, so I think it is effective to align assumptions step by step while thinking through them.

