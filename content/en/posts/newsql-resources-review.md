---
title: Exploring NewSQL Resources
slug: newsql-resources-review
date: 2023-03-29T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - NewSQL
  - Resource Links
translation_key: newsql-resources-review
---

I have been researching and reading various materials on NewSQL, so I will summarize what I have read.

Initially, I wanted to compare several NewSQL databases, but I realized I needed to understand the technologies used internally, so there are quite a few articles on related technologies.

# List of Resources
- [Hybrid Clock](https://martinfowler.com/articles/patterns-of-distributed-systems/hybrid-clock.html)
- [Hybrid Logical Clock (HLC)](https://sergeiturukin.com/2017/06/26/hybrid-logical-clocks.html)
- [The Terrifying Truth About Distributed Systems](https://www.slideshare.net/kumagi/ss-81368169)
- [Focusing on Jepsen, which Provides Verification of Distributed Processing Systems as a Service | Think IT](https://thinkit.co.jp/article/17532)
- [Challenges of Distributed Systems](https://aws.amazon.com/jp/builders-library/challenges-with-distributed-systems/)
- [Let Me Talk About Distributed Systems](https://www.slideshare.net/kumagi/ss-78765920)
- [Why You Need 3 Nodes in a Cluster - Qiita](https://qiita.com/ntoreg/items/ec6f1eca87ba5c5c0399)
- [Is There Anyone Using Distributed Systems Just for the Vibe!? - Qiita](https://qiita.com/muson0110/items/2379595f9bc1d5720478)
- [Navigating the 8 Fallacies of Distributed Computing](https://ably.com/blog/8-fallacies-of-distributed-computing)
- [Pitfalls of Distributed Computing - Wikipedia](https://ja.wikipedia.org/wiki/%E5%88%86%E6%95%A3%E3%82%B3%E3%83%B3%E3%83%94%E3%83%A5%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E3%81%AE%E8%90%BD%E3%81%A8%E3%81%97%E7%A9%B4)
- [Percolator vs Spanner. Implementing Distributed Transactions the Google Way | YugabyteDB](https://www.yugabyte.com/blog/implementing-distributed-transactions-the-google-way-percolator-vs-spanner/)
- [Low Latency Reads in Geo-Distributed SQL with Raft Leader Leases | Yugabyte](https://www.yugabyte.com/blog/low-latency-reads-in-geo-distributed-sql-with-raft-leader-leases/)
- [A Busy Developer’s Guide to Database Storage Engines — The Basics | Yugabyte](https://www.yugabyte.com/blog/a-busy-developers-guide-to-database-storage-engines-the-basics/)
- [Understanding the Distributed Consensus Algorithm Raft - Qiita](https://qiita.com/torao@github/items/5e2c0b7b0ea59b475cce)
- [4 Data Sharding Strategies We Analyzed When Building YugabyteDB](https://www.yugabyte.com/blog/four-data-sharding-strategies-we-analyzed-in-building-a-distributed-sql-database/)
- [CockroachDB's Consistency Model](https://www.cockroachlabs.com/blog/consistency-model/)
- [Spanner vs. Calvin: Distributed Consistency at Scale](https://fauna.com/blog/distributed-consistency-at-scale-spanner-vs-calvin)
- [Fauna | The Distributed Serverless Database](https://fauna.com/)
- [Introduction to Fauna (Temporal Database) - Why Use FQL Instead of Standard SQL? - Qiita](https://qiita.com/masakinihirota/items/c232832bd72ae11905e7)
- [Living Without Atomic Clocks: Where CockroachDB and Spanner Diverge](https://www.cockroachlabs.com/blog/living-without-atomic-clocks/)
- [How an Open Source Distributed NewSQL Database Delivers Time Services - TiDB - PingCAP](https://pingcap.co.jp/how-an-open-source-distributed-newsql-database-delivers-time-services/)
- [The Vitess Docs | What Is Vitess](https://vitess.io/docs/17.0/overview/whatisvitess/)
- [Testing the Use of NewSQL in Financial Transactions - Distributed Database PingCAP](https://pingcap.co.jp/project/customer-story-tis/)
- [Utilizing TiDB in Mission-Critical Scenarios of the Financial Industry (Part 1) - PingCAP](https://pingcap.co.jp/blog-using-tidb-in-mission-critical-scenarios-of-the-financial-industry-part-1/)
- [The Wall of CAP Theorem: The Path to NewSQL | Database Access Performance Blog](https://www.climb.co.jp/blog_dbmoto/archives/5077)
- [Consistency of Distributed Nodes: The Path to NewSQL | Database Access Performance Blog](https://www.climb.co.jp/blog_dbmoto/archives/5193)
- [LSM-Tree and RocksDB, TiDB, CockroachDB - hnakamur's blog](https://hnakamur.github.io/blog/2016/06/20/lsm-tree-and-rocksdb/)
- [Amazon Aurora Architecture Overview](https://pages.awscloud.com/rs/112-TZM-766/images/01_Amazon%20Aurora%20%E3%82%A2%E3%83%BC%E3%82%AD%E3%83%86%E3%82%AF%E3%83%81%E3%83%A3%E6%A6%82%E8%A6%81.pdf)
- [Detailed Explanation of NewSQL Components - Qiita](https://qiita.com/tzkoba/items/3e875e5a6ccd99af332f)
- [Explaining the Advancements of Amazon Aurora - Qiita](https://qiita.com/kumagi/items/67f9ac0fb4e6f70c056d)
- [TiDB on AWS EKS - DMM Video PoC Report - DMM Inside](https://inside.dmm.com/articles/tidb-on-aws-eks-poc-report/)
- [CockroachDB vs. TiDB vs. YugabyteDB Comparison](https://db-engines.com/en/system/CockroachDB%3BTiDB%3BYugabyteDB)
- [Current State of NewSQL in 2020 - Qiita](https://qiita.com/tzkoba/items/5316c6eac66510233115)
- [Architecture Overview](https://www.cockroachlabs.com/docs/stable/architecture/overview.html)
- [Explore YugabyteDB | YugabyteDB Docs](https://docs.yugabyte.com/preview/explore/)
- [[Paper Introduction] TiDB: a Raft-based HTAP Database](https://zenn.dev/tzkoba/articles/4e20ad7a514022)
- [Introduction to TiDB](https://docs.pingcap.com/ja/tidb/stable/overview)
- [What Are the Differences Between SQL, NoSQL, and NewSQL?](https://zenn.dev/umi_mori/books/331c0c9ef9e5f0/viewer/da6ec2#newsql%E3%81%A8%E3%81%AF%EF%BC%9F)
- [Lessons Learned from the 10 Billion Yen Campaign: Innovations Supporting PayPay's Scalable Payment System - Logmi Tech](https://logmi.jp/tech/articles/321524)

# Thoughts
When actually selecting a database, it seems necessary to conduct verifications and performance comparisons tailored to the requirements, but I have identified several points to consider when making comparisons.

- SQL Interface Used
  - Compatibility with SQL such as MySQL and PostgreSQL
- Specifications of Distributed Transactions
  - It's a bit complex and I don't fully understand it, but the specifications for how consistency is ensured differ based on the technologies and policies adopted, such as distributed consensus protocols like Raft and Paxos, and distributed clocks.
- Whether compute and storage are separate or co-located in the architecture
  - I have only used Spanner, so I thought it was normal for them to be separate like Spanner, but it seems there are others that are not.
  - This may affect scalability.
- Whether it is HTAP or not
  - This seems to be an important point if considering analytical use cases.
    - That said, are there any other NewSQLs that support this besides TiDB?