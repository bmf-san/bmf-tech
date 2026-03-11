---
title: About Mob Programming
slug: mob-programming-introduction
date: 2024-01-23T00:00:00Z
author: bmf-san
categories:
  - Development Process
tags:
  - Mob Programming
  - Agile
description: Recently, I experienced mob programming for the first time in my life, so I'm jotting down notes to remember the benefits of the experience.
translation_key: mob-programming-introduction
---



Recently, I experienced mob programming for the first time in my life, so I'm jotting down notes to remember the benefits of the experience.

# What is Mob Programming
Mob programming is a method where three or more members share a single computer to program together.

Typically, the work is divided between a driver who writes the code and navigators who support them.

# Pros / Cons
## Pros
- Easy knowledge sharing
- Ability to receive synchronous feedback
- Easier to align understanding
  - Subsequent reviews and individual work become smoother

## Cons
- Decrease in work efficiency
  - Since those not writing code focus on navigation, overall work efficiency may tend to drop
- Schedule coordination
  - As work is done synchronously, it's necessary to align schedules
- Communication cost
  - While communication flourishes, discussions and consensus-building may take time

# Tools
The [live-share](https://visualstudio.microsoft.com/ja/services/live-share/) extension for VS CODE is well-made and very user-friendly.

# Impressions
Since my experience with mob programming is still limited, my proficiency is low, but I have noticed a few things.

- You can learn and share knowledge on the spot
- Everyone can maintain a certain level of understanding and align their recognition while working
  - It becomes easier for individuals to work independently afterward
- It's more tiring than working alone!
  - This might be because we worked without setting a time limit...
- It seems more efficient to have clear agreements on the work scope (what to advance and to what extent? general work policy) and work rules (how long to work? how long to take breaks? when to switch drivers?)
- Since the work is synchronous, it might be good to record sessions when someone cannot participate
- It's important to distinguish between tasks suitable for mob programming and those better done individually
- While a driver is designated, it seems okay for others to write code or comment as well
  - With live-share, anyone can access the same source code
  - Instead of everyone writing separately, it seems more efficient for the driver to handle the main implementation work, while other members perform supportive implementation tasks (simple tasks or tasks that are faster when divided)
  - Based on my actual experience, it was quite good. However, proper communication is necessary (silently writing code and saying "I did it!" goes against the purpose of mob programming)
