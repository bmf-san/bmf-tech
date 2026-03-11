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
description: A summary of what I researched about GraphQL federation.
translation_key: graphql-federation-introduction
---



A summary of what I researched about GraphQL federation.

# What is GraphQL Federation
GraphQL Federation is a method for integrating multiple GraphQL services.

It is like a microservice for GraphQL, allowing multiple GraphQL services to be used as a single GraphQL endpoint.

# Architecture of GraphQL Federation
## Subgraph
Each domain-specific GraphQL service is called a Subgraph.

GraphQL schema and resolvers are defined.

## Gateway
The GraphQL service that integrates Subgraphs is called a Gateway.

By integrating the schemas of each Subgraph, it provides a single GraphQL endpoint to clients.

It parses GraphQL queries from clients and forwards requests to the appropriate Subgraph.

## Schema Composition
Integrating schemas from multiple Subgraphs is called Schema Composition.

# Advantages and Disadvantages of GraphQL Federation
## Advantages
### Modularity and Scalability
By dividing GraphQL services into Subgraphs, modularity is improved, allowing independent development and deployment of each service.

### Unified Schema
By integrating multiple Subgraphs through a Gateway, clients can obtain data from a single GraphQL endpoint, achieving schema unification.

### Performance Optimization
The Gateway can optimize queries by parsing them and forwarding requests to the appropriate Subgraph.

### Flexible API Evolution
Dividing GraphQL services into Subgraphs allows for flexible API evolution.

## Disadvantages
### Complex Orchestration
Integrating multiple Subgraphs with a Gateway can lead to complex data dependencies.

Additionally, deep query nesting can result in resolvers spanning multiple Subgraphs, potentially degrading performance.

### Single Point of Failure in Gateway
Since the Gateway integrates multiple Subgraphs, if the Gateway goes down, access to all Subgraphs is lost.

### Increased Operational Load
Managing schemas for each Subgraph and operating the Gateway increases operational load.

# Differences between GraphQL Federation and Microservices
While their architectures have similarities, there are many differences.

| **Differences**                     | **GraphQL Federation**                                                                                     | **Microservices**                                                                                        |
|-----------------------------|------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------|
| **Interface Provided**   | Provides a unified GraphQL interface. Clients use a single GraphQL endpoint.                              | Provides multiple APIs (REST, gRPC, etc.). Clients call different API endpoints for each service.                    |
| **Integration Method**                  | Integrates GraphQL schemas between Subgraphs. Provides a consistent GraphQL API to clients.                                   | Integration between services is done via API calls, messaging, or event-driven approaches.                                   |
| **Schema Division Method**         | Each Subgraph has its own GraphQL schema, integrating data with others using `@key` or `@requires`.                        | Each service has its own schema (database, API) and provides data via API.                                        |
| **Granularity of Coupling**                | Coupling between Subgraphs is mainly at the query level, managed by GraphQL schema.                            | Coupling between services is done through API calls or messaging, which can be complex.                                        |
| **Client Perspective**           | Clients can obtain data from multiple services through a single GraphQL endpoint.                               | Clients need to directly call multiple API endpoints, being aware of service differences.                           |
| **State Management**                  | State is managed server-side, maintaining data consistency with schemas and resolvers.                                               | Each service has its own database, maintaining consistency between services using event-driven or CQRS.                           |
| **Dependency Management Between Subgraphs** | Explicitly defines schema dependencies between Subgraphs, managing them with `@requires` or `@key`.                         | Dependencies between services are managed through API calls or event-driven approaches, often requiring coordination.       |
| **Refactoring**           | Schema changes are mainly done within the GraphQL schema, but increased dependencies between Subgraphs can complicate management.           | Refactoring can be done independently for each service, but integration issues may require coordination.               |
| **Organizational Structure**                  | Teams are divided by service, but overall schema management is required. Each team focuses on their Subgraph.                  | Dedicated teams are needed for each service, with coordination between services relying on organizational collaboration.                       |
| **Scalability**            | As services increase, the GraphQL server integrates with Subgraphs, allowing scaling through a unified endpoint.          | Each service can scale independently, but overhead may occur in the integration part.                         |

In short, GraphQL Federation makes it easier for the client side but not for the backend, whereas the opposite is true for microservices.

The burden depends on where the integration between services is resolved.

GraphQL Federation tends to have a lower degree of loose coupling compared to microservices, requiring awareness of dependencies between Subgraphs, which may necessitate team collaboration.

I summarized the differences from a dependency perspective.

| **Comparison Point**      | **GraphQL Federation**                                               | **Microservices**                                               |
|----------------------|------------------------------------------------------------------|------------------------------------------------------------------|
| **Frequency of Team Coordination** | High (frequent schema changes and dependency adjustments needed)                    | Moderate (API contracts or event-driven adjustments needed, but high independence)      |
| **Impact Range of Changes**    | Strong coupling between Subgraphs makes changes easily affect other teams | Each service is independent, allowing a smaller impact range (though API contract management is necessary) |
| **Independence of Release**  | Low (need to coordinate releases of each team to maintain integrated schema) | High (versioning or event-driven allows independent deployment of each service) |
| **Dependency Adjustment**    | High (strong relationships between Subgraphs, making the impact range wide)         | Low (easy to maintain loose coupling, but adjustments needed when APIs change)           |

# Impressions
How are transactions between multiple Subgraphs managed? It doesn't seem to be the responsibility of GraphQL, but is there any mechanism or design pattern in Federation?

# References
- [graphql.org - GraphQL federation](https://graphql.org/learn/federation/)
- [www.ibm.com - What is GraphQL Federation](https://www.ibm.com/jp-ja/topics/graphql-federation)
- [www.apollographql.com - Apollo Docs](https://www.apollographql.com/docs)
