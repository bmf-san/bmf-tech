---
title: About GraphQL
slug: graphql-introduction
date: 2023-11-09T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - GraphQL
translation_key: graphql-introduction
---

# Overview
I have been practicing with GraphQL, so I will summarize what I have researched.

There are helpful tutorials available, making it easy to get started.
cf. [www.howtographql.com](https://www.howtographql.com/)

# What is GraphQL
A query language for developing Web APIs, created by Meta.

GraphQL is managed by the GraphQL Foundation, of which Meta is a member.

The specifications of GraphQL and all related projects are published as OSS.

# Features
- Used over HTTP
- Data is retrieved by sending queries as POST requests
- Multiple data can be retrieved from a single endpoint
- The schema and queries of GraphQL are designed based on the concept of directed graphs
- Schema-first
- Documentation
  - Since it is schema-first, the specifications are self-evident
  - Documentation can also be generated using a documentation generator
- Type definitions
  - Includes scalar types, object types, enum types, list types, non-null types, union types, input types, and interfaces
    - Input types are object types used as arguments for Queries and Mutations
    - The default is Nullable
      - [maku.blog - Guidelines for Nullable and Non-null](https://maku.blog/p/4reqy9i/#nullable-%E3%81%A8-non-null-%E3%81%AE%E3%82%AC%E3%82%A4%E3%83%89%E3%83%A9%E3%82%A4%E3%83%B3)
- Flexibility in data retrieval
  - Only the necessary data can be retrieved through queries
  - Over-fetching and under-fetching can be avoided
- Version management
  - Version management is possible
  - Development can also be done without versioning by adding new types or fields
  - The philosophy is to avoid versioning
- Error handling
  - It is common to return a status code of 200 even in the case of an error, including the error message in the response

# Terminology
Here are a few selected terms.

## Schema
Type definition of a query.

```graphql
type Query {
  user: User
}
```

## Query
A query for data retrieval.

```graphql
query {
  user {
    name
  }
}
```

## Mutation
A query for updating data.

```graphql
mutation {
  updateUser {
    name
  }
}
```

## Subscription
A query for monitoring data changes.

```graphql
subscription {
  user {
    name
  }
}
```

## Argument
Arguments passed to a query.

```graphql
{
  user(id: 123) {
    username
    email
  }
}
```

# Related Technologies
## GraphQL Mesh
A gateway server (GraphQL Gateway) for APIs implemented with API specifications such as gRPC, OpenAPI, Swagger, oData, SOAP, and GraphQL.

As long as you have the API specification, you can access the API with GraphQL queries.

cf. [the-guild.dev](https://the-guild.dev/graphql/mesh)

## openapi-to-graphql
Converts API specifications based on OpenAPI to GraphQL Schema.

cf. [github.com - IBM/openapi-to-graphql](https://github.com/IBM/openapi-to-graphql)

## graphql-tools
A handy tool for creating GraphQL Schemas. It can also create mocks.

cf. [github.com - ardatan/graphql-tools](https://github.com/ardatan/graphql-tools)

## GraphQL Gateway


# Performance
- N+1
  - [Data Loader](https://github.com/graphql/dataloader)
    - Combines multiple SELECTs to the same table into a single SELECT
    - Features Batch and Cache
      - cf. [lyohe.github.io - Main Features](https://lyohe.github.io/blog/2021/12/16/reading-dataloader/#%E4%B8%BB%E8%A6%81%E3%81%AA%E6%A9%9F%E8%83%BD)
- As queries become complex, the request body can become bloated
  - Persisted Query
    - A feature provided by a GraphQL tool called Apollo
    - Prepares an ID corresponding to the query, reducing request parameters by exchanging the ID and the query
      - By making the endpoint for the exchange a GET, it can be cached
- Since it is a POST request, HTTP caching cannot be used
  - Persisted Query

# Thoughts
It seems to take a bit more effort to get started compared to RESTful APIs, but the benefits gained seem significant.

There seem to be various types of GraphQL clients, which may complicate the selection process.

cf. [user-first.ikyu.co.jp - You Might Not Need Apollo Client for Your Product](https://user-first.ikyu.co.jp/entry/2022/07/01/121325)

# References
- [graphql.org](https://graphql.org/)
- [zenn.dev - Articles that thoroughly explain GraphQL](https://zenn.dev/nameless_sn/articles/graphql_tutorial)
- [kinsta.com - Comparison of GraphQL and REST - Key Differences to Know](https://kinsta.com/jp/blog/graphql-vs-rest/#:~:text=GraphQL%E3%81%AE%E6%9C%80%E5%A4%A7%E3%81%AE%E5%88%A9%E7%82%B9,%E3%81%8C%E5%A4%9A%E3%81%99%E3%81%8E%E3%82%8B%E3%81%93%E3%81%A8%E3%81%A7%E3%81%99%E3%80%82)
- [panda-program.com - Learned about GraphQL by reading 'First Time with GraphQL'](https://panda-program.com/posts/book-learning-graphql)
- [qiita.com - GraphQL Schema and Type Definitions](https://qiita.com/NagaokaKenichi/items/d341dc092012e05d6606)
- [gist.github.com - Things to Know About GraphQL](https://gist.github.com/tkdn/75a4d7e38c2edb07b41da078e4a4aa11)
- [maku.blog - GraphQL Best Practices](https://maku.blog/p/4reqy9i/)
- [qiita.com - Various Issues Related to Performance in GraphQL](https://qiita.com/haradakunihiko/items/7148a8b36f1e4e5d60b1)
- [engineering.mercari.com - Considerations When Introducing GraphQL](https://engineering.mercari.com/blog/entry/20220303-concerns-with-using-graphql/)
- [engineering.mercari.com - Implementation of GraphQL Server Using NestJS in Mercari Shops](https://engineering.mercari.com/blog/entry/20210818-mercari-shops-nestjs-graphql-server/#dataloader-for-batch-request)
- [lyohe.github.io - Discussion on Reading graphql/dataloader](https://lyohe.github.io/blog/2021/12/16/reading-dataloader/)
- [zenn.dev - Four Approaches to Solve N+1 Problems in GraphQL](https://zenn.dev/alea12/articles/15d73282c3aacc#%E6%96%B9%E6%B3%954%3A-n%2B1-%E3%82%92%E8%80%83%E6%85%AE%E3%81%97%E3%81%9F-orm-%E3%82%92%E6%A4%9C%E8%A8%8E%E3%81%99%E3%82%8B)
- [www.apollographql.com - GraphQL Concepts Visualized](https://www.apollographql.com/blog/graphql/basics/the-concepts-of-graphql/)