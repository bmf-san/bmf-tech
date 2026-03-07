---
title: Notes on the Saga Pattern
slug: saga-pattern-research-notes
date: 2023-09-17T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Microservices
  - Saga Pattern
  - TCC Pattern
  - Distributed Transactions
  - 2-Phase Commit
translation_key: saga-pattern-research-notes
---

# Overview
Notes on the Saga pattern.

# What is the Saga Pattern?
- In microservices, distributed transactions (like 2-phase commit) are not recommended.
  - The Saga pattern is a way to ensure consistency while avoiding distributed transactions.
- It avoids long locks and utilizes eventual consistency.
- Compensation Transaction
  - An operation that cancels a series of transactions.
  - The Saga pattern prohibits simple rollbacks.
- It is not unique to microservices; it was also used in SOA.
- Implementation Patterns
  - Choreography
    - Each service progresses the transaction with its own responsibilities.
    - It avoids SPOF but makes overall testing difficult.
  - Orchestration
    - A centralized service directs each service to progress the transaction.
    - It can become a SPOF but makes overall testing easier.
- The book [Software Architecture: Hard Parts](https://amzn.to/3RmavPp) introduces eight types of Saga patterns.

# Other Patterns
Another pattern for maintaining consistency in microservices is the TCC (Try-Confirm/Cancel) pattern, which also utilizes eventual consistency like the Saga pattern.

Similar to 2-phase commit, the TCC pattern has three steps for each service: prepare, confirm, and cancel.

The TCC pattern does not perform rollbacks like compensation transactions; instead, it ensures consistency by avoiding operations that could lead to inconsistencies.

# References
- [www.cs.cornell.edu - SAGAS](https://www.cs.cornell.edu/andru/cs711/2002fa/reading/sagas.pdf)
- [microservices.io - Pattern: Saga](https://microservices.io/patterns/data/saga.html)
- [qiita.com - Risks of Distributed Transactions When Migrating to Microservices](https://qiita.com/yoshii0110/items/3c86173dc53d93588b72)
- [qiita.com - Designing Transaction Management for Microservices (Pre-Knowledge Edition)](https://qiita.com/Yoyo-kikuchi/items/c113aeeab3bf2daa0910#tcctry-confirmcancel%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [zenn.dev - Implementing the Saga Pattern on AWS (Choreography Edition)](https://zenn.dev/yoshii0110/articles/57ccc582f7fcd3)
- [learn.microsoft.com - Saga Distributed Transaction Pattern](https://learn.microsoft.com/ja-jp/azure/architecture/reference-architectures/saga/saga)
- [docs.aws.amazon.com - Saga Pattern](https://docs.aws.amazon.com/ja_jp/prescriptive-guidance/latest/modernization-data-persistence/saga-pattern.html)
- [wakatchi.dev - Summarizing Microservices Transactions with TCC and Saga Patterns](https://wakatchi.dev/microservices-tx-pattern-saga-tcc/)
- [eikatou.net - Learning the Saga Pattern for Distributed Transactions](https://eikatou.net/blog/20201205.html)
- [techblog.raksul.com - Validation Using the Saga Pattern in Microservices](https://techblog.raksul.com/entry/2022/09/22/095433)
- [www.techscore.com - Transactions in Microservices [Saga]](https://www.techscore.com/blog/2018/12/05/saga-in-microservices/)
- [www.12-technology.com - [SAGA Pattern] Pros and Cons of Choreography and Orchestration](https://www.12-technology.com/2021/08/dbsaga.html)
- [cloud.google.com - Implementing the Saga Pattern with Workflows](https://cloud.google.com/blog/ja/topics/developers-practitioners/implementing-saga-pattern-workflows)
- [medium.com - GCP Microservices Saga Pattern Edition](https://medium.com/google-cloud-jp/gcp-saga-microservice-7c03a16a7f9d)
- [www.oracle.com - The Reality of Transaction Management in Microservices Architecture](https://www.oracle.com/a/otn/docs/jp-dev-days-microservices.pdf)
- [speakerdeck.com - Did You Accidentally End Up with the Saga Pattern!? A Serverless Backend Operated by a Small Team](https://speakerdeck.com/miu_crescent/qi-gatuitarasagapatanninatuteita-shao-ren-shu-deyun-yong-surusabaresubatukuendo)
- [engineering.mercari.com - Managing Distributed Transactions in the Mercoin Payment Infrastructure](https://engineering.mercari.com/blog/entry/20230614-distributed-transaction/)