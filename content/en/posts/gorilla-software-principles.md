---
title: 10 Software Laws Learned from Gorillas
description: "Understand key software development laws including Parkinson's, Brooks's, and Conway's Law to improve project management."
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
I gave a lightning talk based on a rough memo I wrote earlier about [software development laws](https://bmf-tech.com/posts/%E3%82%BD%E3%83%95%E3%83%88%E3%82%A6%E3%82%A7%E3%82%A2%E9%96%8B%E7%99%BA%E3%81%AE%E6%B3%95%E5%89%87), so I will summarize it by supplementing the slide content.

Slides are here ↓
[10 Software Laws Learned from Gorillas](https://speakerdeck.com/bmf_san/koriratexue-husohutoueafalsefa-ze-10xuan)

It was a bit of a stretch to learn from gorillas, so I will leave the script here.

# Software Development Laws
These are considered "software development laws" not just limited to those spoken in the context of software, but also those from other fields that could apply to software.

Many are based on empirical rules, with sources ranging from management engineering to psychology, or even some paper somewhere.

# Parkinson's Law (First Principle)

**"Work expands to fill the time available for its completion."**

A well-known one.

This was discussed by British historian and political scientist Cyril Northcote Parkinson in his book "Parkinson's Law: The Pursuit of Progress."

Incidentally, the second principle is "Expenditure rises to meet income."

In the computer world, the law is applied to the relationship between data volume and storage.

Data volume expands to the limit of storage, something like that.

Things that have room to grow until they fill the given space will expand, and there are quite a few such elements in software development.

The question is how to confront things that keep increasing, and it seems that setting detailed goals and planning are effective...

# Brooks's Law
**"Adding manpower to a late software project makes it later."**

This law was discussed by American software engineer Frederick Brooks in his book "The Mythical Man-Month."

He is also the person who coined the famous phrase "There is no silver bullet."

Adding manpower to a delayed project creates bottlenecks in communication costs and information catch-up within the team, resulting in an inability to increase productivity and further delays.

You might think that if the additional manpower were a super engineer with the strength of 100 people, delays could be avoided, but it would be good to reflect (confirm) on why this law holds in "The Mythical Man-Month."

# Conway's Law

**"The structure of software reflects the structure of the organization."**

British programmer Melvin Conway.

Famous for inventing coroutines, he proposed this law in a paper on coroutines.

There is a reverse Conway strategy that takes advantage of this law.

The strategy is that a good software structure reflects a good organizational structure, but for the business model, organizational structure, and software structure to become a trinity, it might be good for the software structure to take the lead. (This might be misleading. There might be misunderstandings.)

# Heinrich's Law

**"Behind every major accident, there are 29 minor accidents and 300 anomalies."**

This is one of the empirical rules proposed by Harvard William Heinrich, who worked at an American insurance company, regarding occupational accidents.

In terms of crisis management, when applied to the software world, it suggests that behind incidents and bugs, there should be multiple "anomalies," and we should strive to detect them in advance.

# Linus's Law (Proposed by Eric Raymond)

**"Given enough eyeballs, all bugs are shallow."**

This is a passage discussed by American programmer Eric Raymond in his book "The Cathedral and the Bazaar."

The idea is that if there are enough software users (or contributors??), bugs are not such a serious problem, but it seems more like a mindset than a law....

A cathedral refers to the buildings of a temple.

# Hofstadter's Law

**"It always takes longer than you expect, even when you take into account Hofstadter's Law."**

This is a passage discussed by American scholar Douglas Hofstadter in his book "Gödel, Escher, Bach."

Isn't this a law that everyone has experienced?

One way to resist the law is to try to devise a way to plan.

# Hick's Law

**"The time it takes to make a decision increases with the number and complexity of choices."**

This law was proposed by British psychologist William Hick.

It is a law that should be considered in UI.

There is a formula to calculate the time it takes to make a decision, but I will omit the explanation here.

# Murphy's Law

**"Anything that can go wrong will go wrong."**

There are various theories about the origin of the law.

Although it has a spiritual aspect, in the context of software, it is connected to the concept of fail-safe.

# Lehman's Law

**"The complexity of a system that is evolving increases unless work is done to reduce it."**

The source is this paper
[IEEE XPlore](https://ieeexplore.ieee.org/document/1456074)

I haven't properly read the paper, but I think it connects to evolutionary architecture.

Systems are things that inevitably have to accept change, but without any ingenuity, complexity increases with each change.

# bmf's Law

**"The first idea that comes to mind is the best practice."**

I tried to shape an empirical rule that came to mind.

Even if you consider multiple patterns, sometimes the pattern you first thought of turns out to be the optimal solution, right?? That's what it is.

I think there are similar proverbs or theories...

In internal LT meetings, it seems there are similar laws in the world of love studies or Go, but how about in the world of software...

# Impressions
I haven't been able to read the original texts, so I would like to read them if I have the opportunity.

I feel like being bound by laws leads to a kind of mental stagnation, but when faced with difficult problems, it might be good to follow the laws as objective judgment materials.

