---
title: What is GraphQL? A Complete Guide with Practical Examples
slug: graphql-introduction
date: 2023-11-09T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - GraphQL
description: Summarizing what I researched while practicing GraphQL.
translation_key: graphql-introduction
---

# Overview
I have been practicing GraphQL, so I am summarizing what I researched.

There is a helpful tutorial available, making it easy to get started.
cf. [www.howtographql.com](https://www.howtographql.com/)

# What is GraphQL
A query language for Web API development developed by Meta.

GraphQL is managed by the GraphQL Foundation, of which Meta is a member.

The specifications of GraphQL and all related projects are open source.

# Features
- Used over HTTP
- Retrieves data by sending queries as POST requests
- Can fetch multiple data with a single endpoint
- GraphQL schemas and queries are designed based on the concept of directed graphs
- Schema-first
- Documentation
  - Being schema-first, the specifications are self-evident
  - Documentation can also be generated using document generators
- Type Definitions
  - Includes scalar types, object types, enum types, list types, non-null types, union types, input types, and interfaces
    - Input types are object types used as arguments for Queries and Mutations
    - Default is Nullable
      - [maku.blog - Guidelines for Nullable and Non-null](https://maku.blog/p/4reqy9i/#nullable-%E3%81%A8-non-null-%E3%81%AE%E3%82%AC%E3%82%A4%E3%83%89%E3%83%A9%E3%82%A4%E3%83%B3)
- Flexibility in Data Retrieval
  - Can retrieve only the necessary data through Queries
  - Avoids over-fetching and under-fetching
- Version Management
  - Can manage versions
  - Can develop in a versionless manner by adding new types and fields
  - Generally avoids versioning
- Error Handling
  - Typically returns status code 200 even in case of errors, embedding error messages

# Terminology
Only a few selected terms are listed.

## Schema
Type definitions for queries.

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
A query for data updates.

```graphql
mutation {
  updateUser {
	name
  }
}
```

## Subscription
A query to monitor data changes.

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
A gateway server (GraphQL Gateway) for APIs implemented with specifications like gRPC, OpenAPI, Swagger, oData, SOAP, GraphQL, etc.

As long as there is an API specification, you can access the API with GraphQL queries.

cf. [the-guild.dev](https://the-guild.dev/graphql/mesh)

## openapi-to-graphql
Converts API specifications based on OpenAPI to GraphQL Schema.

cf. [github.com - IBM/openapi-to-graphql](https://github.com/IBM/openapi-to-graphql)

## graphql-tools
A handy tool for creating GraphQL Schemas. Can also create mocks.

cf. [github.com - ardatan/graphql-tools](https://github.com/ardatan/graphql-tools)

## GraphQL Gateway


# Performance
- N+1
  - [Data Loader](https://github.com/graphql/dataloader)
    - Consolidates multiple SELECTs on the same table into a single SELECT
    - Features Batch and Cache
      - cf. [lyohe.github.io - Key Features](https://lyohe.github.io/blog/2021/12/16/reading-dataloader/#%E4%B8%BB%E8%A6%81%E3%81%AA%E6%A9%9F%E8%83%BD)
- As queries become complex, the request body can become bloated
  - Persisted Query
    - A feature provided by Apollo, a GraphQL tool
    - Prepares an ID corresponding to the query, exchanging the ID and query to reduce request parameters
      - By making the endpoint for exchange a GET, it can be cached
- Since it is a POST request, HTTP caching cannot be used
  - Persisted Query

# Impressions
Compared to Restful APIs, it might take a bit more effort to get started, but the benefits seem significant.

There seem to be various types of GraphQL clients, which might make selection challenging.

cf. [user-first.ikyu.co.jp - Apollo Client May Not Be Necessary for Your Product](https://user-first.ikyu.co.jp/entry/2022/07/01/121325)

# References
- [graphql.org](https://graphql.org/)
- [zenn.dev - Comprehensive Articles on GraphQL](https://zenn.dev/nameless_sn/articles/graphql_tutorial)
- [kinsta.com - Comparison of GraphQL and REST: Key Differences to Know](https://kinsta.com/jp/blog/graphql-vs-rest/#:~:text=GraphQL%E3%81%AE%E6%9C%80%E5%A4%A7%E3%81%AE%E5%88%A9%E7%82%B9,%E3%81%8C%E5%A4%9A%E3%81%99%E3%81%8E%E3%82%8B%E3%81%93%E3%81%A8%E3%81%A7%E3%81%99%E3%80%82)
- [panda-program.com - Learned GraphQL Overview from "First GraphQL"](https://panda-program.com/posts/book-learning-graphql)
- [qiita.com - GraphQL Schema and Type Definitions](https://qiita.com/NagaokaKenichi/items/d341dc092012e05d6606)
- [gist.github.com - Things to Know About GraphQL](https://gist.github.com/tkdn/75a4d7e38c2edb07b41da078e4a4aa11)
- [maku.blog - GraphQL Best Practices](https://maku.blog/p/4reqy9i/)
- [qiita.com - Various Performance Issues in GraphQL](https://qiita.com/haradakunihiko/items/7148a8b36f1e4e5d60b1)
- [engineering.mercari.com - Considerations When Introducing GraphQL](https://engineering.mercari.com/blog/entry/20220303-concerns-with-using-graphql/)
- [engineering.mercari.com - Implementing GraphQL Server with NestJS in Mercari Shops](https://engineering.mercari.com/blog/entry/20210818-mercari-shops-nestjs-graphql-server/#dataloader-for-batch-request)
- [lyohe.github.io - Reading graphql/dataloader](https://lyohe.github.io/blog/2021/12/16/reading-dataloader/)
- [zenn.dev - Four Approaches to Solve N+1 Problems in GraphQL](https://zenn.dev/alea12/articles/15d73282c3aacc#%E6%96%B9%E6%B3%954%3A-n%2B1-%E3%82%92%E8%80%83%E6%85%AE%E3%81%97%E3%81%9F-orm-%E3%82%92%E6%A4%9C%E8%A8%8E%E3%81%99%E3%82%8B)
- [www.apollographql.com - GraphQL Concepts Visualized](https://www.apollographql.com/blog/graphql/basics/the-concepts-of-graphql/)