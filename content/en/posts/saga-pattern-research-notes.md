---
title: "The Saga Pattern Explained: Managing Distributed Transactions in Microservices"
slug: saga-pattern
date: 2023-09-17T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Microservices
  - Saga Pattern
  - TCC Pattern
  - Distributed Transactions
  - 2phase commit
translation_key: saga-pattern-research-notes
---



# Overview
Notes on what I researched about the Saga pattern.

# What is the Saga Pattern
- In microservices, distributed transactions (such as 2phase commit) are not recommended
  - The Saga pattern is used to ensure consistency while avoiding distributed transactions
- Avoids long-term locks and uses eventual consistency
- Compensating transactions
  - Operations that cancel a series of transactions
  - The Saga pattern prohibits simple rollbacks
- Not a pattern unique to microservices, it was also used in SOA
- Implementation patterns
  - Choreography
    - Each service progresses the transaction with its own responsibility
    - Avoids SPOF, but makes overall testing difficult
  - Orchestration
    - A central service is prepared to instruct each service to progress the transaction
    - Prone to SPOF, but makes overall testing easier
- [Software Architecture: The Hard Parts](https://amzn.to/3RmavPp) introduces eight types of Saga patterns

# Other Patterns
Another pattern to maintain consistency in microservices, similar to the Saga pattern, is the TCC (Try-Confirm/Cancel) pattern, which also uses eventual consistency.

The TCC pattern is similar to 2phase commit, but in the TCC pattern, each service has three steps: preparation, confirmation, and cancellation of the transaction.

The TCC pattern does not perform rollbacks like compensating transactions, but ensures consistency by not performing processes that cause inconsistencies.

# References
- [www.cs.cornell.edu - SAGAS](https://www.cs.cornell.edu/andru/cs711/2002fa/reading/sagas.pdf)
- [microservices.io - Pattern: Saga](https://microservices.io/patterns/data/saga.html)
- [qiita.com - The Dangers of Distributed Transactions When Migrating to Microservices](https://qiita.com/yoshii0110/items/3c86173dc53d93588b72)
- [qiita.com - Designing Transaction Management in Microservices (Preliminary Knowledge)](https://qiita.com/Yoyo-kikuchi/items/c113aeeab3bf2daa0910#tcctry-confirmcancel%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [zenn.dev - Implementing the Saga Pattern with AWS (Choreography Edition)](https://zenn.dev/yoshii0110/articles/57ccc582f7fcd3)
- [learn.microsoft.com - Saga Distributed Transaction Pattern](https://learn.microsoft.com/ja-jp/azure/architecture/reference-architectures/saga/saga)
- [docs.aws.amazon.com - Saga Pattern](https://docs.aws.amazon.com/ja_jp/prescriptive-guidance/latest/modernization-data-persistence/saga-pattern.html)
- [wakatchi.dev - Summarizing Microservices Transactions with TCC and Saga Patterns](https://wakatchi.dev/microservices-tx-pattern-saga-tcc/)
- [eikatou.net - Learning the Distributed Transaction Saga Pattern](https://eikatou.net/blog/20201205.html)
- [techblog.raksul.com - Verification Using the Saga Pattern in Microservices](https://techblog.raksul.com/entry/2022/09/22/095433)
- [www.techscore.com - Transactions in Microservices [Saga]](https://www.techscore.com/blog/2018/12/05/saga-in-microservices/)
- [www.12-technology.com - [SAGA Pattern] Advantages and Disadvantages of Choreography and Orchestration](https://www.12-technology.com/2021/08/dbsaga.html)
- [cloud.google.com - Implementing the Saga Pattern with Workflows](https://cloud.google.com/blog/ja/topics/developers-practitioners/implementing-saga-pattern-workflows)
- [medium.com - Microservices Saga Pattern on GCP](https://medium.com/google-cloud-jp/gcp-saga-microservice-7c03a16a7f9d)
- [www.oracle.com - The Reality of Transaction Management in Microservices Architecture](https://www.oracle.com/a/otn/docs/jp-dev-days-microservices.pdf)
- [speakerdeck.com - Suddenly Realizing It's a Saga Pattern!? Operating a Serverless Backend with a Small Team](https://speakerdeck.com/miu_crescent/qi-gatuitarasagapatanninatuteita-shao-ren-shu-deyun-yong-surusabaresubatukuendo)
- [engineering.mercari.com - Distributed Transaction Management in Mercari Coin Payment Infrastructure](https://engineering.mercari.com/blog/entry/20230614-distributed-transaction/)
