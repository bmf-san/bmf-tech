---
title: "Comparing API Versioning Strategies: Path, Query, Header, and Payload"
slug: api-versioning-strategies
date: 2026-06-22T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - API
  - REST
  - HTTP
description: "A concise comparison of four API versioning strategies — path, query parameter, header, and message payload — covering their pros, cons, and best fit."
translation_key: api-versioning-strategies
draft: true
---


This post compares four common approaches to API versioning.

# Introduction
Versioning lets you ship breaking changes to an API without disrupting existing clients.

Where you express the version splits the approaches into four main styles.

- Path-based
- Query parameter-based
- Header-based
- Message payload-based

The version identifier format, whether semantic versioning or a date, is a separate choice from its placement.

This post summarizes the characteristics and trade-offs of each.

# Path-based
This style embeds the version in the URL path.

```
GET /v1/users
GET /v2/users
```

The version stays visible in the URL, so anyone can read it directly.

Routing and caching stay easy to control, and you can explore the API from a browser or from documentation.

An API gateway or reverse proxy can route each version to a different backend, and a CDN caches each URL with its standard configuration.

The downside: the URI changes per version even though it points to the same resource.

With a single global version prefix, a change to one resource can force the whole API to a new version, which often duplicates code.

# Query parameter-based
This style passes the version as a query parameter.

```
GET /users?version=1
```

The base URI stays stable, and you can define a default version when the client omits the parameter.

It is simple to add, but the parameter easily becomes optional, so clients forget it and trigger unexpected behavior.

A gateway that routes by path handles a query-based version less easily, and a cache that drops the query from its key can return stale data.

# Header-based
This style carries the version in an HTTP header.

```
GET /users
X-API-Version: 1
```

RFC 6648 discourages the `X-` prefix, so some teams name the header `API-Version` without it.

Some APIs use content negotiation through the `Accept` header instead.

```
GET /users
Accept: application/vnd.example.v2+json
```

The URI then identifies only the resource, and the version acts as metadata rather than part of the address.

Media-type negotiation through the `Accept` header also lets a client pick the representation per request.

GitHub's REST API handles versions through both an `Accept` media type and the date-based `X-GitHub-Api-Version` header.

The version disappears from the URL, however, so it hides from discovery and complicates manual checks and debugging.

Caching requires careful handling of the `Vary` header, which adds operational overhead.

# Message payload-based
This style puts a version field inside the request or message body.

```json
{
  "version": "2",
  "data": {}
}
```

This scheme stays independent of the transport, so it applies beyond HTTP to gRPC and message queues such as Kafka and AMQP that carry no URL.

In asynchronous messaging and event-driven contexts, the version then travels together with the data.

That said, gRPC (Protobuf) evolves schemas through field numbers, and Kafka relies on a Schema Registry, so an in-body version field rarely becomes the norm there.

You must parse the body before routing or validation, which adds performance overhead.

A GET request carries no body, so this scheme fits it poorly and strains consistency.

# Comparison
| Aspect | Path | Query parameter | Header | Payload |
| --- | --- | --- | --- | --- |
| Visibility / discoverability | High | Medium | Low | Low |
| URI purity | Low | Medium | High | N/A |
| Cache control | Easy | Needs care | Needs `Vary` | Hard |
| Routing at a gateway | Easy | Tricky | Medium | Hard |
| Manual testing / debugging | Easy | Easy | Some effort | Some effort |
| Typical use case | Public REST APIs | Lightweight switching | Pure REST design | Async / event-driven |

# Summary
For a public web API, many gateways, CDNs, and load balancers assume path-based routing, so path-based versioning keeps operational cost low.

If you value clean URIs, choose header-based; if lightweight switching covers your needs, choose query parameter-based.

For asynchronous messaging and event-driven systems, payload-based versioning fits naturally, yet for an HTTP API its weak fit with GET and its parsing cost make it the exception rather than the default.

Whichever style you pick, the principle stays the same: prefer additive, backward-compatible changes first, and bump the version only when a breaking change becomes unavoidable.

Pick one style, apply it consistently, and decide your deprecation policy up front.

# References
- [Web API Design](https://bmf-tech.com/posts/web-api-design)
- [microsoft/api-guidelines](https://github.com/microsoft/api-guidelines)
- [Stripe - API versioning](https://stripe.com/blog/api-versioning)
