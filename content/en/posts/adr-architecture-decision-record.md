---
title: "Architecture Decision Records (ADR) Explained: How to Document Technical Decisions"
description: 'Learn what Architecture Decision Records (ADRs) are, why documenting technical decisions matters, and how to write effective ADRs for your engineering team.'
slug: adr-architecture-decision-record
date: 2022-10-10T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Architecture Decision Record
translation_key: adr-architecture-decision-record
---

# Overview
I researched Architecture Decision Records (ADR).

# What is ADR?
A document that records architectural decisions, introduced by Michael Nygard in 2011.
cf. [cognitect.com - DOCUMENTING ARCHITECTURE DECISIONS](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)

# Format
The format proposed by Michael Nygard is as follows:
cf. [cognitect.com - DOCUMENTING ARCHITECTURE DECISIONS](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)

- Title
- Context
  - The context. What scene the decision is made in.
- Decision
  - It is recommended to have one decision per ADR.
- Status
  - Proposed, Approved, Deprecated, etc.
- Result
  - An explanation of the state after applying the decision, including the context of the results.

The document should be a readable length of about 1-2 pages.

# Advantages of Adopting ADR
- Easier to catch up on design decisions regarding architecture.
- Reduced frequency of discussions about explicitly stated decisions.
- Contributes to transparency of architecture within and outside the team.

# Criteria for Deciding Whether to Adopt ADR
- Is it something that the team repeatedly discusses?
- Is it a decision regarding design made by the team?
- Is it a decision that affects the entire software?

Whenever making some architectural judgment or decision, there will be opportunities to write an ADR.

# Thoughts
- In [Fundamentals of Software Architecture](https://amzn.to/3SQV0ge),
- When considering introducing ADR, it seems good to first write an ADR about the introduction of ADR itself.
- It felt that managing ADRs alongside source code in a repository is preferred, differing slightly from Design Docs.
- I thought it would be good to keep the criteria for deciding whether to write an ADR simple.
  - Write everything that was discussed and decided by the team, etc.
- Searchability of ADRs
  - It might be good to have proper searchability when wanting to search ADRs operationally (if the number becomes large).
  - It seems better to categorize and tag them thoughtfully.

# References
- [github.com - Architecture decision record (ADR)](https://github.com/joelparkerhenderson/architecture-decision-record/blob/main/README.md)
- [cognitect.com - DOCUMENTING ARCHITECTURE DECISIONS](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)
- [docs.aws.amazon.com - ADR process](https://docs.aws.amazon.com/prescriptive-guidance/latest/architectural-decision-records/adr-process.html)
- [betterprogramming.pub - The Ultimate Guide to Architectural Decision Records Introduction](https://betterprogramming.pub/the-ultimate-guide-to-architectural-decision-records-6d74fd3850ee)
- [engineering.atspotify.com - When Should I Write an Architecture Decision Record](https://engineering.atspotify.com/2020/04/when-should-i-write-an-architecture-decision-record/)
- [www.redhat.com - Why you should be using architecture decision records to document your project](https://www.redhat.com/architect/architecture-decision-records)
- [www.thoughtworks.com - Lightweight Architecture Decision Records](https://www.thoughtworks.com/radar/techniques/lightweight-architecture-decision-records)
- [developer.mamezou-tech.com - Recommendation for Architecture Decision Records](https://developer.mamezou-tech.com/blogs/2022/04/28/adr/)
- [cloud.google.com - Overview of Architecture Decision Records](https://cloud.google.com/architecture/architecture-decision-records)
- [qiita.com - Record the "Why?" of Architecture! What is ADR?](https://qiita.com/fuubit/items/dbb22435202acbe48849)
- [qiita.com - How to Document Architecture - ADR and ARCHITECTURE.md Template](https://qiita.com/e99h2121/items/f508ef4c9743b8fc9f5b)
- [blog.studysapuri.jp - Record that Decision - Using Architecture Decision Records (ADR)](https://blog.studysapuri.jp/entry/architecture_decision_records)
- [fintan.jp - Case Study of Introducing Architecture Decision Records](https://fintan.jp/page/1636/)