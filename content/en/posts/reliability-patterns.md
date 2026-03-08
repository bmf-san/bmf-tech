---
title: Reliability Patterns
slug: reliability-patterns
date: 2024-10-30T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Reliability
  - Architecture
  - Design
  - System Design
translation_key: reliability-patterns
---

# Overview
This post summarizes reliability patterns based on those proposed by Azure, AWS, and GCP.

# What is Reliability
Reliability is the quality of being able to continuously provide the functions expected by users (systems or applications).

# Characteristics Supporting Reliability
Reliability is supported by the following characteristics:

- **Availability**: The system is available for use.
- **Durability**: Data is not lost.
- **Fault Tolerance**: The system continues to function even when a failure occurs.
- **Recoverability**: The system returns to a normal state after a failure occurs.
  - Mechanisms for automatic recovery or maintenance of runbooks, etc.
- **Predictability**: The performance of the system is predictable.
  - Monitoring and observability, etc.
- **Scalability**: The system can scale out according to load.
- **Security**: The system is secure.

# Cloud Design Patterns Supporting Reliability
Here are some patterns related to reliability:

## Ambassador Pattern
A pattern that delegates network communication processing to another service to reduce the load on the original service.

## BFF (Backend for Frontend)
A pattern that places a service dedicated to the frontend application between the frontend application and backend services.

## Bulkhead
A pattern that ensures that if one part of the system fails, other parts are not affected.

## Cache Aside
A pattern that uses caching to avoid putting load on resources like databases or APIs.

cf. [Cache Write Method](https://bmf-tech.com/posts/%E3%82%AD%E3%83%A3%E3%83%83%E3%82%B7%E3%83%A5%E3%81%AE%E6%9B%B8%E3%81%8D%E8%BE%BC%E3%81%BF%E6%96%B9%E5%BC%8F)

## Circuit Breaker
A pattern that rejects requests for a certain period until a failure is resolved when a failure occurs.

It has three states: Closed, Half-Open, and Open.

Closed: State where requests are accepted.
Half-Open: State where some requests are accepted.
Open: State where requests are not accepted.

## Claim Check
A pattern that checks whether a request is legitimate before accepting it.

For example, to avoid directly exchanging large payloads between services, the payload can be stored in a database, and its ID can be exchanged instead.

## Compensation Transaction
A pattern that restores resources to their original state if a transaction fails while updating multiple resources.

## Competing Consumers (Work Queue)
A pattern where multiple consumers retrieve and process messages from the same messaging queue.

*Referencing [learn.microsoft.com - Competing Consumers Pattern](https://learn.microsoft.com/ja-jp/azure/architecture/patterns/competing-consumers), this may be a simple messaging queue despite the unfamiliar name.

## Event Sourcing
A pattern that records changes in the state of the system as events, allowing the system's state to be reconstructed.

## Priority Queue
A pattern that assigns priorities to queues and processes higher-priority queues preferentially.

## Pub/Sub
A pattern where a publisher sends messages and subscribers receive them.

Publishers and subscribers exchange messages via topics.

# References
- [learn.microsoft.com - Cloud Design Patterns Supporting Reliability](https://learn.microsoft.com/ja-jp/azure/well-architected/reliability/design-patterns)
- [wa.aws.amazon.com - 5 Pillars of the Framework: Reliability](https://wa.aws.amazon.com/wellarchitected/2020-07-02T19-33-23/wat.pillar.reliability.ja.html)
- [cloud.google.com - Reliability](https://cloud.google.com/architecture/framework/reliability?hl=ja)