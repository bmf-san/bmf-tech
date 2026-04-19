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
  - Provider-Driven Contract Testing
  - Consumer-Driven Contract Testing
description: A summary of research on contract testing and Pact.
translation_key: contract-testing-pact
---



# Overview
This post summarizes the research on contract testing and Pact.

# What is Contract Testing?
Contract testing is a testing method that verifies the interactions between a consumer (the caller of a service) and a provider (the service provider).

The agreement between the consumer and provider is described as a contract, and testing is conducted based on this contract.

A contract testing format where the consumer issues the contract and the provider tests according to it is called **Consumer-Driven Contract Testing**.

Conversely, a format where the provider issues the contract and the consumer tests according to it is called **Provider-Driven Contract Testing**.

Which is more suitable, consumer-driven or provider-driven contract testing, depends on factors such as the number of consumers, API stability (frequency of changes), and which side holds the initiative in the development process.

There is also a service called [pactflow](https://pactflow.io/) that supports bi-directional contracts. [pactflow](https://pactflow.io/) is based on Pact, which supports consumer-driven contract testing. Unlike Pact, [pactflow](https://pactflow.io/) is not open-source.

cf.
- [www.xlsoft.com - SmartBear Pactflow: What is Contract Testing? Why Use It?](https://www.xlsoft.com/jp/blog/blog/2022/10/18/smartbear-32158-post-32158/)
- [ipsj.ixsq.nii.ac.jp - Consumer-Driven Contract Testing Patterns and Challenges](https://ipsj.ixsq.nii.ac.jp/ej/?action=repository_uri&item_id=193867&file_id=1&file_no=1)
- [www.ibm.com - Contract Testing](https://www.ibm.com/docs/ja/devops-test-workbench/11.0.0?topic=tasks-contract-testing)
- [gitlab-docs.creationline.com - Contract Testing](https://gitlab-docs.creationline.com/ee/development/testing_guide/contract/)
- ~~docs.pactflow.io - Bi-Directional Contract Testing Guide~~
- [pactflow.io - Bi-Directional Contract Testing](https://pactflow.io/bi-directional-contract-testing/)
- [pactflow.io - Pact is dead, long live Pact](https://pactflow.io/blog/bi-directional-contracts/)
- [www.thoughtworks.com - Pactflow](https://www.thoughtworks.com/radar/tools/pactflow)
- [technology.lastminute.com - Impacts of contract tests in our microservice architecture](https://technology.lastminute.com/impacts-of-contract-tests-in-a-microservice-architecture/)
- [alexromanov.github.io - Should You Use Contract Testing?](https://alexromanov.github.io/2021/07/12/should-you-use-contract-testing/)

## API Schema and Contract Testing
If schema-driven API development like Open API or ProtoBuf is well-established, you might question the necessity of contract testing.

An API schema is a document that defines the specifications of an API, including elements like endpoints, requests and responses, and data models.

Using tools that support schema-driven development, you can generate code for both the consumer and provider sides based on the API definition.

While you can obtain a certain guarantee that the implementation follows the API schema, you cannot guarantee the behavior of the API (whether the implemented API works as expected).

Schema-driven development alone cannot ensure that changes on the provider side are communicated to the consumer side in a code-first manner, leading to communication and testing overhead.

Contract testing is a method for testing interactions between services, while an API schema is a document for defining API specifications, thus they serve different purposes.

cf.
- [Schemas are not contracts](https://pactflow.io/blog/schemas-are-not-contracts/)
- [pactflow.io - Schema-based contract testing with JSON schemas and Open API (Part 1)](https://pactflow.io/blog/contract-testing-using-json-schemas-and-open-api-part-1/)
- [pactflow.io - Schema-based contract testing with JSON schemas and Open API (Part 2)](https://pactflow.io/blog/contract-testing-using-json-schemas-and-open-api-part-2/)
- [Introducing Pact to the Frontend - Start Testing APIs with Contracts](https://blog.techscore.com/entry/2020/10/16/080000#OpenAPI%E3%81%A8%E3%81%AE%E9%81%95%E3%81%84%E3%81%AF%E3%81%AA%E3%81%AB)

## Advantages and Disadvantages of Contract Testing
Let's organize the advantages and disadvantages of contract testing that does not lean towards either consumer or provider-driven.

cf.
- [docs.pact.io - Convince me](https://docs.pact.io/faq/convinceme)
- [www.infoq.com - Using Contract Testing in Microservice Applications](https://www.infoq.com/jp/news/2019/04/contract-testing-microservices/)

### Advantages
- Maintains reliability and consistency between services
- Automatically detects changes in consumers or providers
  - Reduces communication costs between teams
- Faster execution speed than E2E
- Clarifies dependencies between services

### Disadvantages
- Tool dependency and cost of tool implementation
- Requires consensus building between organizations or teams for implementation
- Needs to be integrated as a process within the development flow

# What is Pact?
Pact is a tool for contract testing that supports consumer-driven contract testing. It does not support provider-driven contract testing.

It supports multiple languages and can integrate with testing frameworks and build tools.

Pact supports HTTP and messages (non-HTTP message queues. Pact itself can test messages without knowing the implementation details of Rabbit MQ, SQS, Kafka, etc.).

The workflow of Pact is as follows:

1. The consumer describes the contract
2. A Pact file is generated during the consumer-side test execution
3. The Pact file is uploaded to the pact_broker
4. The provider uses the uploaded Pact file to execute tests

cf.
- [docs.pact.io](https://docs.pact.io/)
- [github.com - pact-foundation](https://github.com/pact-foundation)
- [pactflow.io - How Pact contract testing works](https://pactflow.io/how-pact-works/?utm_source=ossdocs&utm_campaign=getting_started#slide-1)
- [dius.com.au - Simplifying Microservice testing with Pacts](https://dius.com.au/2014/05/20/simplifying-microservice-testing-with-pacts/)
- [pactflow.io - The curious case for the Provider Driven Contract](https://pactflow.io/blog/the-curious-case-for-the-provider-driven-contract/)

# Thoughts
In organizations of a certain size, where development is done by multiple teams, the challenge of maintaining consistency between services naturally arises.

I believe contract testing is an effective method for addressing such challenges.

Particularly, the ability to automatically detect API specification changes is a significant advantage. This leverage increases as the number of teams or services grows.

Despite the benefits of contract testing, it didn't seem to be particularly popular based on my research.

Although it's past data, it wasn't in a status of active adoption in Technology Radar.

Implementing contract testing may require tool adoption, and among those tools, Pact seems to be a major one.

The reasons Pact may not appear popular could be that it only supports consumer-driven contract testing and the perceived benefits of implementation are not substantial (other alternative integration testing tools like E2E might suffice). This is just a feeling since I haven't actually used it, and it might be more widespread overseas, but at least domestically, there weren't many examples.

PactFlow, an extension of Pact, seemed to be still in the development stage, but it could potentially become a good tool to lower the barriers to entry for contract testing and its benefits.

As services reach a certain scale, systems that provide APIs used by multiple services may emerge, and in such cases, provider-driven contract testing might be more suitable. In these cases, Pact might not be the optimal solution.