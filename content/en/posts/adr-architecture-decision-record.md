---
title: About ADR (Architecture Decision Record)
slug: adr-architecture-decision-record
date: 2022-10-10T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - ADR
description: Exploring ADR (Architecture Decision Record).
translation_key: adr-architecture-decision-record
---

# Overview
Explored ADR (Architecture Decision Record).

# What is ADR?
ADR refers to a document that records decisions related to architecture, introduced by Michael Nygard in 2011.
cf. [cognitect.com - DOCUMENTING ARCHITECTURE DECISIONS](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)

# Format
Michael Nygard proposed the following format:
cf. [cognitect.com - DOCUMENTING ARCHITECTURE DECISIONS](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)

- Title
- Context
  - The context or scenario in which the decision is made.
- Decision
  - It is recommended to record one decision per ADR.
- Status
  - Proposal, Approved, Deprecated, etc.
- Consequences
  - Explanation of the context and state after applying the decision.

The document should be concise, about 1-2 pages long for readability.

# Benefits of Adopting ADR
- Easier to catch up on design decisions related to architecture.
- Reduced frequency of discussions about clearly stated decisions.
- Contributes to transparency of architecture within and outside the team.

# Deciding Whether to Adopt ADR
- Is it a topic repeatedly discussed within the team?
- Is it a design decision made by the team?
- Does it impact the overall software?

ADR can be written whenever there is a need to make architectural judgments or decisions.

# Thoughts
- In [Fundamentals of Software Architecture](https://amzn.to/3SQV0ge):
- When considering introducing ADR, it might be good to first write an ADR about introducing ADR.
- Unlike Design Docs, ADRs seem to be preferably managed in repositories along with source code.
- Simplifying the criteria for writing ADRs seems beneficial.
  - For example, writing ADRs for all decisions made after team discussions.
- Searchability of ADRs:
  - Operationally, it might be helpful to ensure ADRs are searchable (especially if the number grows).
  - Categorization, tags, etc., might be useful.

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
- [qiita.com - Record the "Why" of Architecture! What is ADR?](https://qiita.com/fuubit/items/dbb22435202acbe48849)
- [qiita.com - How to Document Architecture - ADR and ARCHITECTURE.md Templates](https://qiita.com/e99h2121/items/f508ef4c9743b8fc9f5b)
- [blog.studysapuri.jp - "Engrave the Decision" - Using Architecture Decision Records (ADR) for Design Documentation](https://blog.studysapuri.jp/entry/architecture_decision_records)
- [fintan.jp - Case Study of Introducing Architecture Decision Records](https://fintan.jp/page/1636/)