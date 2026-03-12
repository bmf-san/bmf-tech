---
title: About Bitemporal Data Model
description: An in-depth exploration of About Bitemporal Data Model, covering design principles, trade-offs, and practical applications.
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
This post summarizes what I have researched about the Bitemporal Data Model.

# What is Bitemporal Data Model
It is one of the data models handled in the field of Temporal Data Models, referring to a data model with two timelines among those that handle time axes.

- Non-Temporal
  - Holds only the current state
  - Does not retain information about past history or future changes
- Uni-Temporal
  - Has a single timeline
  - Retains the start date when the data is valid and the end date when the validity period ends
- Bi-Temporal
  - Has two timelines
  - Retains transaction time (system time), which is the time recorded in the database, and valid time, which is the time the event occurred
  - Unlike Uni-Temporal, where transaction time and valid time are the same, Bi-Temporal has different timelines

The background for adopting a bitemporal data model includes the following requirements:

- Tracking changes in data history
- Legal and audit requirements
- Improved flexibility in time series analysis

When handling a bitemporal data model in an RDB, the following difficulties arise:

- Ensuring data consistency
  - It is necessary to impose constraints to ensure that the valid time does not overlap for the same entity
  - Continuity of transaction time must be maintained
  - In PostgreSQL, the EXCLUDE constraint is useful
- Complexity of queries
  - Queries need to consider both valid time and transaction time
    - This can easily become complex, making performance tuning difficult
- Complexity of application logic
  - There are two timelines to consider when retrieving, updating, or deleting data
  - It is necessary to implement logic to maintain consistency

Although it is a data model that allows for referencing past history, adding past or future history, and retaining update information of history itself, achieving this requires accepting complexity. (That's how I felt...)

# Impressions
I have never dealt with highly flexible historical data in my work before, so I felt a deep sense of the world of handling history.

# References
- [en.wikipedia.org - Bitemporal modeling](https://en.wikipedia.org/wiki/Bitemporal_modeling)
- [martinfowler.com - Bitemporal History](https://martinfowler.com/articles/bitemporal-history.html)
- [martinfowler.com - Temporal Patterns](https://martinfowler.com/eaaDev/timeNarrative.html)
- [wiki.postgresql.org - Temporal Data & Time Travel in PostgreSQL](https://wiki.postgresql.org/images/6/64/Fosdem20150130PostgresqlTemporal.pdf)
- [www.progress.com - What can be achieved with Bitemporal](https://www.progress.com/docs/default-source/marklogic-docs/Bitemporal-Whitepaper-JP.pdf)
- [matsu-chara.hatenablog.com - Points to consider when introducing BiTemporal Data Model](https://matsu-chara.hatenablog.com/entry/2022/06/25/110000)
- [matsu-chara.hatenablog.com - A blog that's mostly unclear](https://matsu-chara.hatenablog.com/entry/2017/04/01/110000)
- [www.slideshare.net - Introduction to Temporal Data Model and Reladomo for Data History Management #jjug_ccc #ccc_g3](https://www.slideshare.net/slideshow/jjug-ccc-2017-spring-bitemporal-data-modeling-and-reladomo/76145041)
- [tech.smarthr.jp - How to walk with ActiveRecord::Bitemporal](https://tech.smarthr.jp/entry/2023/10/04/110000)
- [speakerdeck.com - Implementing Command History and Temporal Access - BiTemporal Data Model Practice](https://speakerdeck.com/f440/implementing-command-history-and-temporal-access)
- [speakerdeck.com - After the Practice of BiTemporal Data Model in SmartHR](https://speakerdeck.com/wata727/after-the-practice-of-bitemporal-data-model-in-smarthr)
- [zenn.dev - Introduction to Interpretation and Utilization of Temporal Data Model from Business Perspective](https://zenn.dev/zahn/articles/6a3d2138e9fe68)
- [ontact-rajeshvinayagam.medium.com - Bi-Temporal Data Modeling: An Overview](https://contact-rajeshvinayagam.medium.com/bi-temporal-data-modeling-an-overview-cbba335d1947)
