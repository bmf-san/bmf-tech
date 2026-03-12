---
title: What is BFF (Backend For Frontend)? Benefits and Implementation Guide
slug: bff-explained
date: 2023-08-29T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - BFF
description: A summary of what I researched about BFF.
translation_key: bff-explained
---



# Overview
This is a summary of what I researched about BFF.

# What is BFF
BFF stands for Backends For Frontends. It is not Best Friends Forever.

As the name suggests, it is a backend server for the frontend, responsible for responding with APIs or HTML for the frontend, playing a role in UI/UX.

BFF can solve the problem of addressing the diversity of clients (the callers of the server) by organizing the requirements for each client.

# Points of Interest
- **Programming Language**
  - Since BFF is a backend for the frontend, it often seems to be composed of frontend-oriented technologies.
- **Reconfiguration**
  - Once you make it a BFF, it seems difficult to dismantle.
  - It might be better to postpone adopting BFF until it is truly necessary (though determining if it is truly necessary is not easy...)
- **Possible Anti-patterns**
  - Lack of communication between backend and frontend engineers
  - Too much non-UI logic on the BFF
  - A big bang joint that combines backend and frontend all at once
- **Ease of Frontend Optimization**
  - Optimizing API calls could improve UI display performance.
- **BFF and DDD**
  - Is domain organization on the frontend necessary? I'm not sure about this...
- **API Aggregation Units**
  - There seems to be difficulty in deciding how to group APIs.
  - If you're doing microservices, setting up another microservice instead of a BFF might be better, which could undermine the benefits of BFF.
- **Compatibility with Micro Frontends**
  - I have no knowledge of micro frontends, so I don't understand anything.
  - It might be influenced by the component structure of micro frontends?
- **Good Compatibility with GraphQL**
  - If using GraphQL, code-first rather than schema-first is more suitable according to some cases.
    - cf. [Why unify GraphQL to code-first? BFF/FE reorganization for consistency in type definitions](https://logmi.jp/tech/articles/326592))
- **Availability**
  - Since BFF aggregates multiple backends, it is affected by and dependent on the failures of multiple backends.
  - Regarding this concern, ZOZO seems to have devised a way to respond only with data that can be returned correctly.
    - cf. [Started Backends For Frontends (BFF)](https://techblog.zozo.com/entry/zozo-aggregation-api-bff)
- **Caching**
  - Caching on the BFF side also needs to be considered.
- **Timeout and Retry Control**
  - Although it's a consideration for regular APIs, adjusting settings for BFF might be a bit tricky.
- **Deployment**
  - It seems necessary to coordinate the deployment of BFF and backend.
  - It affects the speed of delivery.
- **No Business Logic**
  - It seems basic for BFF not to have business logic, but can business logic be completely separated? There might be cases where it can't.

# References
- [Pattern: Backends For Frontends](https://samnewman.io/patterns/architectural/bff/)
- [BFF (Backends For Frontends) Introduction—Why Netflix, Twitter, and Recruit Technologies Adopt It](https://atmarkit.itmedia.co.jp/ait/articles/1803/12/news012.html)
- [BFF (Backend for Frontend)](https://speakerdeck.com/dena_tech/bff-backend-for-frontend)
- [Developers Working on BFF Talk "UIT#3 The 'Backends for Frontends' Sharing"](https://engineering.linecorp.com/ja/blog/bff-talk-uit-3-backends-for-front-ends-sharing)
- [BFF/SSR Talk](https://speakerdeck.com/yosuke_furukawa/ssrfalsehua?slide=2)
- [Session Report "Let's Make Backends for Frontends Serverless!!—To Focus on the Original Purpose" Summary](https://tsd.mitsue.co.jp/blog/2021-10-28-event-report-aws-innovate-modern-app-edition-bff/)
- [Dividing Frontend from Backend is an Antipattern](https://www.thoughtworks.com/insights/blog/dividing-frontend-backend-antipattern)
- [BFF and Microservices Architecture](https://zenn.dev/hirac/articles/7bd857ab904d66)
- [Rearchitecting to a Backend for Frontend with GraphQL](https://www.docswell.com/s/hireroo/Z6YJR4-techtalk2-1#p1)
- [Three Anti-patterns in BFF (Backends For Frontends) Practice and How to Avoid Them](https://atmarkit.itmedia.co.jp/ait/articles/1808/31/news013.html)
- [Five Convenient Use Cases of BFF (Backends For Frontends)](https://atmarkit.itmedia.co.jp/ait/articles/1805/18/news022.html)
- [Reconsidering Backends for Frontends (BFF)](https://zenn.dev/morinokami/scraps/20a4eab9415a50)
- [More Coverage on BFFs](https://samnewman.io/blog/2016/02/14/more-coverage-on-bffs/)
- [Embracing the Differences: Inside the Netflix API Redesign](https://netflixtechblog.com/embracing-the-differences-inside-the-netflix-api-redesign-15fd8b3dc49d)
- [BFF @ SoundCloud](https://www.thoughtworks.com/insights/blog/bff-soundcloud)
- [Moving to Microservices at SoundCloud with Lukasz Plotnicki](https://softwareengineeringdaily.com/2016/02/04/moving-to-microservices-at-soundcloud-with-lukasz-plotnicki/)
- [Started Backends For Frontends (BFF)](https://techblog.zozo.com/entry/zozo-aggregation-api-bff)
- [Backends for Frontends Pattern](https://aws.amazon.com/jp/blogs/news/backends-for-frontends-pattern/)
- [What is BFF?](https://qiita.com/souhei-etou/items/d5de99bb8cba1c59d393)
- [Should BFF (Backend for Frontend) be Adopted in New Development?](https://vivit.hatenablog.com/entry/2021/11/10/101530)
- [Backend for Frontend (BFF) Pattern—Why Do You Need to Know It?](https://medium.com/mobilepeople/backend-for-frontend-pattern-why-you-need-to-know-it-46f94ce420b0)
- [What is the Trending BFF Architecture?｜Offers Tech Blog](https://zenn.dev/overflow_offers/articles/20220418-what-is-bff-architecture)
- [Why Unify GraphQL to Code-first? BFF/FE Reorganization for Consistency in Type Definitions](https://logmi.jp/tech/articles/326592)
- [Are BFF (Backend for Frontend) and DDD Mutually Exclusive?](https://stackoverflow.com/questions/76940683/are-bff-backend-for-frontend-and-ddd-mutually-exclusive)
- [Do Frontend Engineers Dream of Micro Frontends?](https://engineering.mercari.com/blog/entry/2018-12-06-162827/)

# Impressions
I initially thought I would just skim through BFF since I already knew about it, but there were many interesting points to consider, such as architecture availability, handling business logic, appropriate client aggregation, and its relation to organizational structure.

Personally, I feel that BFF has many potential pitfalls if not approached cautiously. While I can see the traps, it seems challenging to build it in a way that avoids them.

If I ever have the opportunity to consider BFF, I plan to reflect on this.