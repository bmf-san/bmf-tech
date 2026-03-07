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
translation_key: mob-programming-introduction
---

Recently, I experienced mob programming for the first time in my life, so I wanted to jot down my thoughts to remember the positives of that experience.

# What is Mob Programming
It is a method of programming where three or more members share one computer.

Typically, the work is divided into two roles: the driver who writes the code and the navigator who supports them.

# Pros / Cons
## Pros
- Easy knowledge sharing
- Can receive synchronous feedback
- Easier to align understanding
  - Subsequent reviews and individual work become smoother

## Cons
- Decreased work efficiency
  - Since those who are not writing code must focus on navigation, overall work efficiency may drop.
- Schedule coordination
  - Since work is done synchronously, it is necessary to align schedules.
- Communication costs
  - While communication can flow well, discussions and consensus-building may take time.

# Tools
The VS CODE extension [live-share](https://visualstudio.microsoft.com/ja/services/live-share/) is well-made and very easy to use.

# Impressions
Since my experience with mob programming is still shallow, my skills are low, but I have noticed a few things.

- Learning is gained. Knowledge can be shared on the spot.
- Everyone can maintain a certain level of understanding and work together while aligning their perceptions.
  - This makes it easier for each person to do individual work later.
- It is more tiring than I expected to work as a group!
  - This might be because we were working without setting a time limit...
- It feels more efficient to have clear agreements on the work scope (what to progress and to what extent? general work policy) and work rules (how long is the work time? how many minutes for breaks? when to switch drivers?).
- Since it is synchronous work, it seems good to record the session when someone cannot participate.
- It is necessary to distinguish well between tasks suited for mob work and those better done individually.
- While the driver is designated, it seems fine for others to write code or comment as well.
  - Using live-share allows anyone to touch the same source code.
  - Instead of everyone writing separately, it seems more efficient for the driver to handle the main implementation work while other members perform supportive implementation tasks (simple tasks or tasks that are faster when divided).
  - From my actual experience, it felt quite good. However, proper communication is necessary (doing things like writing code silently and saying "I did it!" goes against the essence of mob programming).