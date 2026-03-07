---
title: 10 Software Principles Learned from Gorillas
slug: gorilla-software-principles
date: 2019-04-17T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Software Development
translation_key: gorilla-software-principles
---

# Overview
Based on the rough notes I wrote previously on [Software Development Principles](https://bmf-tech.com/posts/%E3%82%BD%E3%83%95%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A2%E9%96%8B%E7%99%BA%E3%81%AE%E6%B3%95%E5%89%87), I gave a lightning talk, and I will summarize the slide content here.

The slides are here:
[10 Software Principles Learned from Gorillas](https://speakerdeck.com/bmf_san/koriratexue-husohutoueafalsefa-ze-10xuan)

It was a bit unrealistic to learn from gorillas, so I left a script behind.

# Software Development Principles
Rather than being bound by principles discussed in the context of software, I consider principles from other fields that may apply to software as "Software Development Principles."

Many are based on empirical rules, but the sources vary, including management engineering, psychology, and various papers.

# Parkinson's Law (First Principle)
**"The amount of work expands to fill the time available for its completion."**

A well-known principle.

Discussed by British historian and political scientist Cyril Northcote Parkinson in his book, "Parkinson's Law: The Pursuit of Progress."

By the way, the second principle states, "Expenditures increase until they reach income."

In the world of computers, this law is applied to the relationship between data volume and storage capacity.

Data volume tends to expand to the limits of storage capacity.

It suggests that things with room to grow will continue to expand until they fill the given constraints, and there are quite a few elements in software development that have such relationships.

The discussion revolves around how to confront things that keep increasing, and it seems that detailed goal setting and planning are effective...

# Brooks' Law
**"Adding manpower to a late software project makes it later."**

This principle was discussed by American software engineer Frederick Brooks in his book, "The Mythical Man-Month."

He is also known for the famous saying, "There is no silver bullet."

Adding personnel to a delayed project can create bottlenecks in communication costs and information catch-up within the team, resulting in decreased productivity and further delays.

You might think that if the added personnel were super engineers with the power of 100 men, the delays could be avoided, but it would be good to reflect (confirm) on the reasons this law holds true as discussed in "The Mythical Man-Month."

# Conway's Law
**"The structure of software reflects the structure of the organization that produces it."**

British programmer Melvin Conway is known for inventing coroutines and proposed this law in a paper about coroutines.

There is a strategy called the reverse Conway strategy that takes advantage of this law.

The strategy suggests that a good software structure will reflect a good organizational structure, but for the business model, organizational structure, and software structure to be in harmony, it might be better for the software structure to take the lead (though this could be misleading or misunderstood).

# Heinrich's Law
**"For every major accident, there are 29 minor accidents and 300 near misses."**

This is one of the empirical rules proposed by Harvard William Heinrich, who worked for an American insurance company, regarding workplace accidents.

In crisis management, when applied to the software world, it teaches us that there are likely multiple "anomalies" behind incidents and bugs, and we should strive to detect them in advance.

# Linus's Law (proposed by Eric Raymond)
**"Given enough eyeballs, all bugs are shallow."**

This was discussed by American programmer Eric Raymond in his book, "The Cathedral and the Bazaar."

In essence, if there are enough users (or contributors??) of the software, bugs are not such a serious problem, but it feels more like a mindset rather than a strict law...

The term "cathedral" refers to a temple building.

# Hofstadter's Law
**"It always takes longer than you expect, even when you take into account Hofstadter's Law."**

This was discussed by American scholar Douglas Hofstadter in his book, "Gödel, Escher, Bach."

Isn't this a law that everyone has experienced?

One way to counter this law is to be creative in how you plan.

# Hick's Law
**"The time it takes to make a decision is proportional to the number of choices available."**

This law was proposed by British psychologist William Hick.

I believe this is a law that should be considered in UI design.

There is a formula to determine the time taken for decision-making, but I will skip the explanation here.

# Murphy's Law
**"Anything that can go wrong will go wrong."**

There are various theories about the origin of this law.

While it has a spiritual aspect, in the context of software, it connects to the idea of fail-safes.

# Lehman's Law
**"The complexity of a changing system continues to increase."**

The source is this paper: [IEEE XPlore](https://ieeexplore.ieee.org/document/1456074)

I haven't thoroughly read the paper, but I think it connects to evolutionary architecture.

Systems must accept change, but if no effort is made, complexity will increase with each change.

# BMF's Law
**"The first idea that comes to mind is the best practice."**

I shaped an empirical rule that came to mind.

Even when considering multiple patterns, the first pattern that comes to mind is often the optimal solution, right?? That's what it is.

I think there are similar proverbs or theories...

In our internal LT meetings, it seems there are similar laws in the world of romance and Go, but what about in the software world?

# Thoughts
I haven't been able to read the original texts, but I would like to if I get the chance.

I feel that being bound by principles can lead to a halt in thinking, but when faced with difficult problems, it might be good to follow these principles as objective judgment materials.