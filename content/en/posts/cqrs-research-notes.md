---
title: Notes on CQRS
slug: cqrs-research-notes
date: 2023-09-18T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - CQRS
  - Microservices
translation_key: cqrs-research-notes
---

# Overview
Notes on what I researched about CQRS.

# What is CQRS?
- Command Query Responsibility Segregation
  - A pattern that separates responsibilities between Command and Query
    - Command
       - Responsible for updating data
       - Design commands focused on tasks rather than data
       - It is assumed to encapsulate the DDD domain model
    - Query
       - Responsible for data retrieval
       - Designed as objects optimized for each use case
         - DTO
- Implementation Methods
  - Event Sourcing Pattern
    - In CQRS, events are commonly used for the interaction between commands and queries
  - [www.ibm.com - Command Query Responsibility Segregation (CQRS) pattern](https://www.ibm.com/cloud/architecture/architectures/event-driven-cqrs-pattern/)
    - Discusses a gradual transition to CQRS
- Advantages
  - Allows scaling of Read and Write independently
  - Can choose different data sources for Read and Write
    - It is also possible to keep the same data source
    - If separating data sources, data must be synchronized either synchronously or asynchronously
  - The application-side model becomes simpler
- Disadvantages
  - Increased cost due to more system components
  - Complexity in the overall network configuration

# Thoughts
I feel quite a high barrier to deciding to adopt this pattern...

It seems there are many cases overseas, but I found that there are still few cases in Japan.

# References
- [microservices.io - Pattern: Command Query Responsibility Segregation (CQRS)](https://microservices.io/patterns/data/cqrs.html)
- [learn.microsoft.com - CQRS Pattern](https://learn.microsoft.com/ja-jp/azure/architecture/patterns/cqrs)
- [learn.microsoft.com - Event Sourcing Pattern](https://learn.microsoft.com/ja-jp/azure/architecture/patterns/event-sourcing)
- [zenn.dev - Fully Understanding CQRS](https://zenn.dev/shmi593/articles/c1baeb2d453929)
- [martinfowler.com - What does 'Event Driven' mean?](https://martinfowler.com/articles/201701-event-driven.html)
- [martinfowler.com - CQRS](https://martinfowler.com/bliki/CQRS.html)
- [docs.aws.amazon.com - CQRS Pattern](https://docs.aws.amazon.com/ja_jp/prescriptive-guidance/latest/modernization-data-persistence/cqrs-pattern.html)
- [logmi.jp - Chatwork's Tech Lead Talks About Effectively Using CQRS](https://logmi.jp/tech/articles/322972)
- [logmi.jp - 'Doing CQRS' is Almost Synonymous with 'Doing Event Sourcing': Design Philosophy Reflecting Reactive Systems and CQRS](https://logmi.jp/tech/articles/324798)
- [little-hands.hatenablog.com - Introduction to Practicing CQRS [Domain-Driven Design]](https://little-hands.hatenablog.com/entry/2019/12/02/cqrs)
- [hireoo.io - A Story About Creating a Microservice to Process Statistical Data Using the CQRS Pattern](https://hireroo.io/journal/tech/statistics-service-using-cqrs)
- [speakerdeck.com - Briefly Explaining CQRS/Event Sourcing](https://speakerdeck.com/j5ik2o/event-sourcingwojie-shuo-suru)
- [blog.j5ik2o.me - Why Does CQRS Lead to Event Sourcing?](https://blog.j5ik2o.me/entry/2020/09/18/172612)
- [note.com - Can CQRS Be Achieved Without Event Sourcing?](https://note.com/j5ik2o/n/n20aadb440a9b)
- [appmaster.io - Application of CQRS and Event Sourcing in Microservices](https://appmaster.io/ja/blog/cqrs-ibentososhingu-maikurosabisu)
- [www.ibm.com - Command Query Responsibility Segregation (CQRS) pattern](https://www.ibm.com/cloud/architecture/architectures/event-driven-cqrs-pattern/)
- [www.eventstore.com - CQRS](https://www.eventstore.com/cqrs-pattern)
- [blog.risingstack.com - When Should You Use CQRS?](https://blog.risingstack.com/when-to-use-cqrs)
- [logmi.jp - 'Autonomy' Means Nothing If It Doesn't Function Independently: Considering the Advantages and Disadvantages of CQRS from Command and Query Requirements](https://logmi.jp/tech/articles/324797)
- [pages.awscloud.com - Role and Implementation of CQRS & Event Sourcing in Modern Architecture](https://pages.awscloud.com/rs/112-TZM-766/images/DevAx_connect_jp_season1_day4_CQRS%26EventSourcing.pdf)