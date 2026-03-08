---
title: Contract Testing and Pact
slug: contract-testing-pact
date: 2024-07-21T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - Contract Testing
  - Pact
  - Consumer-Driven Contract Testing
  - Provider-Driven Contract Testing
translation_key: contract-testing-pact
---

# Overview
This post summarizes my research on contract testing and Pact.

# What is Contract Testing?
Contract testing is a testing methodology that tests the interactions between a consumer (the party calling the service) and a provider (the party providing the service).

Agreements between the consumer and provider are documented as contracts, and tests are conducted based on these contracts.

When the consumer issues the contract and the provider tests according to that contract, it is referred to as **Consumer-Driven Contract Testing**.

Conversely, when the provider issues the contract and the consumer tests according to that contract, it is referred to as **Provider-Driven Contract Testing**.

Which type of contract testing is more suitable depends on factors such as the number of consumers, the stability of the API (frequency of changes), and which party has more control in the development process.

There is also a service called [pactflow](https://pactflow.io/) that supports bi-directional contracts. [pactflow](https://pactflow.io/) is based on Pact, which supports consumer-driven contract testing. Unlike Pact, [pactflow](https://pactflow.io/) is not open-source.

cf.
- [www.xlsoft.com - SmartBear Pactflow: What is Contract Testing? Reasons to Use It](https://www.xlsoft.com/jp/blog/blog/2022/10/18/smartbear-32158-post-32158/)
- [ipsj.ixsq.nii.ac.jp - Consumer-Driven Contract Testing Patterns and Their Challenges](https://ipsj.ixsq.nii.ac.jp/ej/?action=repository_uri&item_id=193867&file_id=1&file_no=1)
- [www.ibm.com - Contract Testing](https://www.ibm.com/docs/ja/devops-test-workbench/11.0.0?topic=tasks-contract-testing)
- [gitlab-docs.creationline.com - Contract Testing](https://gitlab-docs.creationline.com/ee/development/testing_guide/contract/)
- [docs.pactflow.io - Bi-Directional Contract Testing Guide](https://docs.pactflow.io/docs/bi-directional-contract-testing)
- [pactflow.io - Bi-Directional Contract Testing](https://pactflow.io/bi-directional-contract-testing/)
- [pactflow.io - Pact is dead, long live Pact](https://pactflow.io/blog/bi-directional-contracts/)
- [www.thoughtworks.com - Pactflow](https://www.thoughtworks.com/radar/tools/pactflow)
- [technology.lastminute.com - Impacts of Contract Tests in Our Microservice Architecture](https://technology.lastminute.com/impacts-of-contract-tests-in-a-microservice-architecture/)
- [alexromanov.github.io - Should You Use Contract Testing?](https://alexromanov.github.io/2021/07/12/should-you-use-contract-testing/)

## API Schema and Contract Testing
If API development is schema-driven, such as with Open API or ProtoBuf, you may question the necessity of contract testing.

An API schema is a document that defines the specifications of the API, including elements such as endpoints, requests/responses, and data models.

Using tools that support schema-driven development, you can generate code for both the consumer and provider based on the API definition.

While you can obtain a certain guarantee that the implementation adheres to the API schema, you cannot guarantee the behavior of the API (whether the implemented API behaves as expected).

Since schema-driven development alone cannot ensure that changes made by the provider are communicated to the consumer in a code-first manner, communication and testing overhead can arise.

Contract testing is a methodology for testing interactions between services, while API schemas are documents that define API specifications; thus, they serve different purposes.

cf.
- [Schemas are not contracts](https://pactflow.io/blog/schemas-are-not-contracts/)
- [pactflow.io - Schema-based Contract Testing with JSON Schemas and Open API (Part 1)](https://pactflow.io/blog/contract-testing-using-json-schemas-and-open-api-part-1/)
- [pactflow.io - Schema-based Contract Testing with JSON Schemas and Open API (Part 2)](https://pactflow.io/blog/contract-testing-using-json-schemas-and-open-api-part-2/)
- [Introducing Pact to Frontend - Let's Start Testing APIs with Contracts](https://blog.techscore.com/entry/2020/10/16/080000#OpenAPI%E3%81%A8%E3%81%AE%E9%81%95%E3%81%84%E3%81%AF%E3%81%AA%E3%81%AB)

## Advantages and Disadvantages of Contract Testing
Here, I will summarize the advantages and disadvantages of contract testing that does not lean towards either consumer or provider-driven approaches.

cf.
- [docs.pact.io - Convince me](https://docs.pact.io/faq/convinceme)
- [www.infoq.com - Using Contract Testing in Microservice Applications](https://www.infoq.com/jp/news/2019/04/contract-testing-microservices/)

### Advantages
- Maintains reliability and consistency between services
- Automatically detects changes in the consumer or provider
  - Reduces communication costs between teams
- Faster execution speed compared to E2E
- Clarifies dependencies between services

### Disadvantages
- Tool dependency and costs of tool adoption
- Requires consensus among organizations or teams for implementation
- Needs to be integrated as one of the processes in the development flow

# What is Pact?
Pact is a tool for contract testing that supports consumer-driven contract testing. It does not support provider-driven contract testing.

It supports multiple languages and can integrate with testing frameworks and build tools.

Pact supports HTTP and messages (non-HTTP message queues, allowing testing of messages without needing to know the implementation details of Rabbit MQ, SQS, Kafka, etc.).

The workflow of Pact operates as follows:

1. The consumer writes the contract.
2. A Pact file is generated during the consumer's test execution.
3. The Pact file is uploaded to the pact_broker.
4. The provider executes tests using the uploaded Pact file.

cf.
- [docs.pact.io](https://docs.pact.io/)
- [github.com - pact-foundation](https://github.com/pact-foundation)
- [pactflow.io - How Pact Contract Testing Works](https://pactflow.io/how-pact-works/?utm_source=ossdocs&utm_campaign=getting_started#slide-1)
- [dius.com.au - Simplifying Microservice Testing with Pacts](https://dius.com.au/2014/05/20/simplifying-microservice-testing-with-pacts/)
- [pactflow.io - The Curious Case for the Provider Driven Contract](https://pactflow.io/blog/the-curious-case-for-the-provider-driven-contract/)

# Thoughts
As organizations grow to a certain size and develop with multiple teams, the challenge of maintaining consistency between services naturally arises.

I believe contract testing is an effective methodology for addressing such challenges.

In particular, the ability to automatically detect changes in API specifications is a significant advantage. This leverage increases as the number of teams or services grows.

While I see the benefits of contract testing, my research suggests it hasn't gained widespread popularity.

Historically, it hasn't been in a status of active adoption, even in the Technology Radar.

Implementing contract testing may require adopting tools, and Pact seems to be one of the major options among those tools.

Reasons Pact may not appear to be popular include its support only for consumer-driven contract testing and the lack of perceived benefits (as other integration testing tools like E2E may suffice). While I haven't used it myself, it may be more prevalent abroad; however, there seem to be few cases domestically.

PactFlow, an extension of Pact, seems to be in a developmental stage, but I believe it could become a good tool for lowering the barriers to the benefits of contract testing.

As services scale, systems that provide APIs used by multiple services may emerge, and in such cases, provider-driven contract testing may be more suitable than consumer-driven contract testing. Pact may not be the optimal solution in such scenarios.