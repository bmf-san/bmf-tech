---
title: What is GraphQL
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
  - User-friendly due to the similarity between the data format of API requests and responses
- REST is an architecture (design), while GraphQL is a language (DSL)

# Comparison of REST API and GraphQL

## REST API Format
Requests are made to endpoints using HTTP verbs

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
Queries are made to a single endpoint

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
| Type System | None | Present |
| Versioning Needed | Yes | No |
| Documentation Needed | Yes | No |
| Resource Limitation | Primarily call count | Based on resource amount |

- Flexibly specify the desired data for a single endpoint
  - In REST, the response data is fixed for each endpoint.
  - In GraphQL, you specify the desired data for a single endpoint to get the response data.

- Resource limitations require ingenuity
  - Addressed based on resource amount
  - Need to consider load calculation methods such as basing it on the number of objects

- Almost no need for documentation
  - The API definition serves as documentation
    - The structure of the query and the response data structure are almost the same

# Points of Concern
- Dependency on libraries
  - Libraries are needed to parse queries

- Not necessarily better performance than REST API
  - Reduces the number of requests
  - Increases the amount of data per request
  - Both REST API and GraphQL require ingenuity to control data volume (like paging or field specification)

- Monitoring
  - REST API can be monitored per endpoint
  - GraphQL has a single endpoint, making it difficult to monitor response performance per query. Some measures are needed.
    - Wait for the ecosystem to mature or implement it in-house
 
- Caching
   - HTTP caching cannot be used
   - It seems advisable to research various other aspects

# Personal Thoughts
- If there are many components and a complex UI in the application, and the number of requests increases making it difficult for the client side, there seems to be a benefit to adopting it.
- I was considering trying it out with [Rubel](https://github.com/bmf-san/Rubel), but it feels premature, so I will refrain. In fact, it is unnecessary at this point.

# References
- [graphql.org](https://graphql.org/)
- [facebook/graphql rfcs](https://github.com/facebook/graphql/blob/master/rfcs/Subscriptions.md)
- [A slight rebuttal to "What is GraphQL suitable for?"](http://yamitzky.hatenablog.com/entry/graphql)
- [Comparing how GraphQL changes the flow of app development with REST](https://www.webprofessional.jp/rest-2-0-graphql/)
- [GraphQL is not a replacement for REST](https://note.mu/konpyu/n/nc4fd122644a1)
- [Notes on implementing GraphQL API in a Rails app](https://blog.qnyp.com/2017/06/08/graphql-resources/)
- [An easy-to-understand explanation of GraphQL from a beginner's perspective! How does it differ from REST?](https://vitalify.jp/app-lab/20171006-graphql/)