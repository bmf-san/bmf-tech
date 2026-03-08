---
title: About GraphQL Federation
slug: graphql-federation-introduction
date: 2025-02-02T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - GraphQL
  - GraphQL Federation
translation_key: graphql-federation-introduction
---

I will summarize what I researched about GraphQL federation.

# What is GraphQL Federation
GraphQL Federation is a method for integrating multiple GraphQL services.

It is like a microservice for GraphQL, an architectural pattern that allows multiple GraphQL services to be used as a single GraphQL endpoint.

# Architecture of GraphQL Federation
## Subgraph
The GraphQL service for each domain is called a Subgraph.

GraphQL schemas and resolvers are defined here.

## Gateway
The GraphQL service that integrates Subgraphs is called a Gateway.

By integrating the schemas of each Subgraph, it provides a single GraphQL endpoint to clients.

It parses GraphQL queries from clients and forwards requests to the appropriate Subgraph.

## Schema Composition
Integrating the schemas of multiple Subgraphs is called Schema Composition.

# Advantages and Disadvantages of GraphQL Federation
## Advantages
### Modularity and Scalability
By splitting GraphQL services by Subgraph, modularity is improved, allowing independent development and deployment for each service.

### Schema Unification
By integrating multiple Subgraphs through the Gateway, clients can obtain data from a single GraphQL endpoint, achieving schema unification.

### Performance Optimization
The Gateway can optimize queries by parsing them and forwarding requests to the appropriate Subgraph.

### Flexibility in API Evolution
By splitting GraphQL services by Subgraph, API evolution becomes more flexible.

## Disadvantages
### Complex Orchestration
Since the Gateway integrates multiple Subgraphs, data dependencies can become complex.

Additionally, deep nesting of queries can lead to resolvers spanning multiple Subgraphs, potentially degrading performance.

### Single Point of Failure for Gateway
Since the Gateway integrates multiple Subgraphs, if the Gateway goes down, access to all Subgraphs will be lost.

### Increased Operational Load
Managing schemas for each Subgraph and operating the Gateway increases operational load.

# Differences Between GraphQL Federation and Microservices
While both architectures have similarities, there are many differences.

| **Differences**                     | **GraphQL Federation**                                                                                     | **Microservices**                                                                                        |
|-----------------------------|------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------|
| **Interface Provided**   | Provides a unified GraphQL interface. Clients use a single GraphQL endpoint.                              | Provides multiple APIs (REST, gRPC, etc.). Clients call different API endpoints for each service.                    |
| **Integration Method**                  | Integrates GraphQL schemas between Subgraphs, providing a consistent GraphQL API to clients.                                   | Integration between services is done through API calls, messaging, or event-driven approaches.                                   |
| **Schema Splitting Method**         | Each Subgraph has its own GraphQL schema, integrating data with others using `@key` and `@requires`.                        | Each service has its own schema (database, API) and provides data via API.                                        |
| **Granularity of Joining**                | Joins between Subgraphs are mainly done at the query level, managed through the GraphQL schema.                            | Joins between services are done through API calls or messaging, which can be complex.                                        |
| **Client Perspective**           | Clients can obtain data from multiple services through a single GraphQL endpoint.                               | Clients need to directly call multiple API endpoints and be aware of differences between services.                           |
| **State Management**                  | State management is done on the server side, maintaining data consistency through schemas and resolvers.                                               | Each service has its own database, maintaining consistency between services using event-driven or CQRS approaches.                           |
| **Dependency Management Between Subgraphs** | Explicitly defines schema dependencies between Subgraphs, managing them with `@requires` and `@key`.                         | Dependencies between services are managed through API calls or event-driven approaches, often requiring coordination between services.       |
| **Refactoring**           | Schema changes are mainly done within the GraphQL schema, but managing dependencies between Subgraphs can become complex.           | Refactoring can be done independently at the service level, but issues may arise in the integration parts.               |
| **Organizational Structure**                  | Teams are divided by service, but overall schema management is necessary. Each team focuses on its own Subgraph.                  | Each service requires dedicated teams, and coordination between services relies on collaboration within the organization.                       |
| **Scalability**            | As services increase, the GraphQL server integrates with Subgraphs, allowing scaling through a unified endpoint.          | Each service can scale independently, but there may be overhead in the integration parts.                         |

To put it simply, GraphQL Federation makes it easier for the client side but not for the backend, while the opposite can be said for microservices.

The burden depends on where the integration between services is resolved.

GraphQL federation has a lower degree of loose coupling compared to microservices, and it may require more collaboration between teams due to the need to be aware of dependencies between Subgraphs.

I summarized the differences from the perspective of dependencies.

| **Comparison Points**      | **GraphQL Federation**                                               | **Microservices**                                               |
|----------------------|------------------------------------------------------------------|------------------------------------------------------------------|
| **Frequency of Team Coordination** | High (frequent adjustments needed for schema changes and dependencies)                    | Moderate (API contracts and event-driven adjustments needed, but high independence)      |
| **Impact Scope of Changes**    | Strong coupling between Subgraphs makes it easy for changes to affect other teams | Each service is independent, allowing for smaller impact scope (but API contract management is necessary) |
| **Release Independence**  | Low (releases need to be coordinated to maintain integrated schema) | High (services can be independently deployed using versioning or event-driven approaches) |
| **Dependency Adjustment**    | High (strong relationships between Subgraphs lead to wider impact)         | Low (easier to maintain loose coupling, but adjustments needed for API changes)           |

# Thoughts
How do you manage transactions between multiple Subgraphs? It doesn't seem to be the responsibility of GraphQL, but are there any mechanisms or design patterns in Federation for this?

# References
- [graphql.org - GraphQL federation](https://graphql.org/learn/federation/)
- [www.ibm.com - What is GraphQL Federation](https://www.ibm.com/jp-ja/topics/graphql-federation)
- [www.apollographql.com - Apollo Docs](https://www.apollographql.com/docs)