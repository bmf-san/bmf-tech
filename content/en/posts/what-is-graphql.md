---
title: What is GraphQL
description: "Understand GraphQL as a query language for APIs with flexible data selection, single endpoint, and type system advantages."
slug: what-is-graphql
date: 2018-06-14T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - API
  - HTTP
  - REST
  - GraphQL
translation_key: what-is-graphql
---



# What is GraphQL
- Developed by Facebook
- A query language for APIs
  - User-friendly because the data format of API requests and responses are similar
- REST is an architecture (design), while GraphQL is a language (DSL)

# Comparison of REST API and GraphQL

## REST API Format
Send requests to endpoints using HTTP verbs

`curl https://api.bmf-tech.com/v1/configs`

```
[
    {
        "id": 1,
        "name": "title",
        "alias_name": "Title",
        "value": "bmf-tech",
        "created_at": "2017-09-25 23:08:23",
        "value": "bmf-tech",
        "deleted_at": null
    }
]
```

# GraphQL API Format
Send queries to a single endpoint

`curl https://api.bmf-tech.com/api`

```
configs {
  id,
  name,
  alias_name
  value,
  created_at,
  updated_at,
  deleted_at
}
```

```
[
    {
        "id": 1,
        "name": "title",
        "alias_name": "Title",
        "value": "bmf-tech",
        "created_at": "2017-09-25 23:08:23",
        "value": "bmf-tech",
        "deleted_at": null
    }
]
```

| | REST API | GraphQL |
|:-----------|:-----------:|:------------:|
| Endpoints | Multiple | Single |
| HTTP Verbs | Dependent | Independent |
| Type System | None | Available |
| Versioning Required | Yes | No |
| Documentation Required | Yes | No |
| Resource Limitation | Mainly call count | Handled according to resource amount |

- Flexibly specify the desired data for a single endpoint
  - REST has fixed response data for each endpoint.
  - GraphQL allows specifying the desired data for a single endpoint to get the response data.

- Resource limitation requires ingenuity
  - Handled according to resource amount
  - Need to consider methods for load calculation, such as based on the number of objects

- Almost no need for documentation
  - API definition serves as documentation
    - The query structure and response data structure are almost the same

# Points of Concern
- Dependency on libraries
  - Requires libraries for parsing queries

- Not necessarily better performance than REST API
  - Reduces the number of requests
  - Increases the amount of data per request
  - Need to control data volume in both REST API and GraphQL (e.g., pagination, field specification)

- Monitoring
  - REST API can be monitored per endpoint
  - GraphQL is a single endpoint, making it difficult to monitor response performance per query. Some measures are needed.
    - Wait for ecosystem maturity or implement in-house

- Caching
   - Cannot use HTTP cache
   - Should investigate various other aspects

# Impressions
- In applications with many components and complex UIs, where the number of requests increases and the client side struggles, there may be benefits to introducing it.
- Considered using it with [Rubel](https://github.com/bmf-san/Rubel), but decided against it as it feels premature. Currently unnecessary.

# References
- [graphql.org](https://graphql.org/)
- ~~facebook/graphql rfcs~~
- [A small rebuttal to "What is GraphQL suitable for"](http://yamitzky.hatenablog.com/entry/graphql)
- [Compared "GraphQL" that changes the flow of app development with REST](https://www.webprofessional.jp/rest-2-0-graphql/)
- [GraphQL is not a replacement for REST](https://note.mu/konpyu/n/nc4fd122644a1)
- [Memo when implementing GraphQL API in Rails app](https://blog.qnyp.com/2017/06/08/graphql-resources/)
- [Explaining GraphQL from a beginner's perspective! ~ What is the difference from REST, the same Web API? ~](http://web.archive.org/web/20200808063325/https://vitalify.jp/app-lab/20171006-graphql/)
