---
title: 'Web API Design'
description: 'Web API Design'
slug: web-api-design
date: 2024-07-28T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - API
  - OpenAPI
  - Design
  - Book Review
translation_key: web-api-design
books:
  - asin: "4798167010"
    title: "Web APIの設計"
---


I read [Web APIの設計](https://amzn.to/3y7dqEG).

This book provides a perspective on API design.

It covers not only the interface of API design but also the entry points that come before it.

- It is better to design API with a focus on what users can do rather than the mechanisms, as focusing on the mechanisms can lead to complexity.
- The book explains a framework called Goal Canvas as an approach to clarify the goals of the API, which I found to be a good method.
- There are various patterns for API versioning, including path, domain, query parameters, custom headers, content negotiation, and consumer settings (holding settings for each consumer in the database).
