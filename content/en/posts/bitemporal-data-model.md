---
title: About Bitemporal Data Model
slug: bitemporal-data-model
date: 2024-05-27T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - Bi-Temporal
  - Uni-Temporal
  - Non-Temporal
  - History
  - Temporal Data Model
translation_key: bitemporal-data-model
---

# Overview
This post summarizes my research on the Bitemporal Data Model.

# What is Bitemporal Data Model?
It is one of the data models handled in the field of Temporal Data Models, referring to a data model that has two time axes among those that deal with time.

- Non-Temporal
  - Holds only the current state
  - Does not retain information about past history or future changes
- Uni-Temporal
  - Has a single time axis
  - Retains the start date when the data is valid and the end date when the validity period ends
- Bi-Temporal
  - Has two time axes
  - Retains the transaction time (system time) recorded in the database and the valid time when the event occurred
  - While Uni-Temporal has the same transaction time and valid time, Bi-Temporal has different time axes

The following requirements are considered as the background for adopting the Bitemporal Data Model:

- Tracking changes in data history
- Regulatory and audit requirements
- Improved flexibility for time series analysis

When dealing with a Bitemporal data model in RDB, the following challenges arise:

- Ensuring data consistency
  - Constraints must be in place to prevent overlapping valid times for the same entity
  - Continuity of transaction time must be maintained
  - In PostgreSQL, the exclusion constraint (EXCLUDE constraint) can be useful
- Complexity of queries
  - Queries must consider both valid time and transaction time
    - This can complicate things, making performance tuning difficult
- Complexity of application logic
  - There are two time axes to be aware of when retrieving, updating, or deleting data
  - It is necessary to implement logic that maintains consistency

Although it is a data model that allows referencing past history, adding past or future history, and retaining update information about the history itself, achieving this requires a readiness for complexity. (I felt this way...)

# Thoughts
I have never dealt with highly flexible historical data in my work, so I felt a deep understanding of the world of history handling.

# References
- [en.wikipedia.org - Bitemporal modeling](https://en.wikipedia.org/wiki/Bitemporal_modeling)
- [martinfowler.com - Bitemporal History](https://martinfowler.com/articles/bitemporal-history.html)
- [martinfowler.com - Temporal Patterns](https://martinfowler.com/eaaDev/timeNarrative.html)
- [wiki.postgresql.org - Temporal Data & Time Travel in PostgreSQL](https://wiki.postgresql.org/images/6/64/Fosdem20150130PostgresqlTemporal.pdf)
- [www.progress.com - What can be achieved with Bitemporal](https://www.progress.com/docs/default-source/marklogic-docs/Bitemporal-Whitepaper-JP.pdf)
- [matsu-chara.hatenablog.com - Points to consider when introducing BiTemporal Data Model](https://matsu-chara.hatenablog.com/entry/2022/06/25/110000)
- [matsu-chara.hatenablog.com - A blog that is generally not well understood](https://matsu-chara.hatenablog.com/entry/2017/04/01/110000)
- [www.slideshare.net - Introduction to Temporal Data Models and Reladomo for Data History Management #jjug_ccc #ccc_g3](https://www.slideshare.net/slideshow/jjug-ccc-2017-spring-bitemporal-data-modeling-and-reladomo/76145041)
- [tech.smarthr.jp - Journey with ActiveRecord::Bitemporal](https://tech.smarthr.jp/entry/2023/10/04/110000)
- [speakerdeck.com - Implementing Command History/Temporal Access - Practical BiTemporal Data Model](https://speakerdeck.com/f440/implementing-command-history-and-temporal-access)
- [speakerdeck.com - The aftermath of practicing BiTemporal Data Model in SmartHR](https://speakerdeck.com/wata727/after-the-practice-of-bitemporal-data-model-in-smarthr)
- [zenn.dev - Introduction to the interpretation and utilization of Temporal Data Models from a business perspective](https://zenn.dev/zahn/articles/6a3d2138e9fe68)
- [ontact-rajeshvinayagam.medium.com - Bi-Temporal Data Modeling: An Overview](https://contact-rajeshvinayagam.medium.com/bi-temporal-data-modeling-an-overview-cbba335d1947)