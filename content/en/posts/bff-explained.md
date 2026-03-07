---
title: About BFF
slug: bff-explained
date: 2023-08-29T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - BFF
translation_key: bff-explained
---

# Overview
This post summarizes what I researched about BFF.

# What is BFF?
BFF stands for Backends For Frontends, not Best Friends Forever.

As the name suggests, it is a backend server for the frontend, responsible for roles such as responding with APIs or HTML for UI/UX.

BFF can solve the problem of responding to the diversity of clients (the callers of the server) by organizing the requests for each client.

# Points of Interest
- Programming Language
  - It seems that BFF is often composed of frontend-oriented technologies due to its nature as a backend for the frontend.
- Reconfiguration
  - Once you implement BFF, it seems difficult to dismantle it.
  - It might be better to postpone its adoption until it is truly necessary (though determining if it is really needed can be challenging).
- Potential Anti-Patterns
  - Lack of communication between backend and frontend engineers.
  - Excessive logic unrelated to UI being included in BFF.
  - A big bang joint where backend and frontend are tightly coupled all at once.
- Easier Frontend Optimization
  - Optimizing API calls could improve UI display performance.
- BFF and DDD
  - Is domain organization needed on the frontend side? This is unclear...
- API Aggregation Units
  - There seems to be difficulty in deciding how to group APIs.
  - If you are doing microservices, wouldn't it have been better to set up another microservice instead of BFF? This could undermine the benefits of BFF.
- Compatibility with Micro Frontends
  - I have no knowledge of micro frontends and don't understand anything.
  - Could it be influenced by the component structure of micro frontends?
- Good Compatibility with GraphQL
  - If using GraphQL, a code-first approach rather than a schema-first approach seems more suitable.
    - cf. [Why did we unify GraphQL to code-first? To maintain consistency in type definitions for BFF/FE organization](https://logmi.jp/tech/articles/326592)
- Availability
  - Since BFF aggregates multiple backends, it is affected by and dependent on failures of multiple backends.
  - Regarding this concern, it seems that ZOZO has devised a way to respond only with data that can be returned normally.
    - cf. [Started Backends For Frontends (BFF)](https://techblog.zozo.com/entry/zozo-aggregation-api-bff)
- Caching
  - It seems necessary to consider caching on the BFF side as well.
- Timeout and Retry Control
  - This is a consideration for normal APIs, but in the case of BFF, adjusting the settings might be a bit tricky.
- Deployment
  - It seems necessary to coordinate the deployment of BFF with the deployment of the backend.
  - This relates to delivery speed.
- No Business Logic
  - It seems that BFF is fundamentally not supposed to have business logic, but can it be completely separated? There might be cases where it cannot be.

# References
- [Pattern: Backends For Frontends](https://samnewman.io/patterns/architectural/bff/)
- [BFF (Backends For Frontends) Introduction - Reasons adopted by Netflix, Twitter, and Recruit Technologies](https://atmarkit.itmedia.co.jp/ait/articles/1803/12/news012.html)
- [BFF (Backend for Frontend)](https://speakerdeck.com/dena_tech/bff-backend-for-frontend)
- [Developers working on BFF talk about "UIT#3 The 'Backends for Frontends' sharing"](https://engineering.linecorp.com/ja/blog/bff-talk-uit-3-backends-for-front-ends-sharing)
- [Discussion on BFF/SSR](https://speakerdeck.com/yosuke_furukawa/ssrfalsehua?slide=2)
- [Session Report: Let's make Backends for Frontends serverless!! - To focus on the original purpose](https://tsd.mitsue.co.jp/blog/2021-10-28-event-report-aws-innovate-modern-app-edition-bff/)
- [Dividing frontend from backend is an antipattern](https://www.thoughtworks.com/insights/blog/dividing-frontend-backend-antipattern)
- [BFF and microservices architecture](https://zenn.dev/hirac/articles/7bd857ab904d66)
- [Re-architecting to Backend for Frontend using GraphQL](https://www.docswell.com/s/hireroo/Z6YJR4-techtalk2-1#p1)
- [Three anti-patterns in BFF practice and their avoidance strategies](https://atmarkit.itmedia.co.jp/ait/articles/1808/31/news013.html)
- [Five useful use cases for Backends for Frontends (BFF)](https://atmarkit.itmedia.co.jp/ait/articles/1805/18/news022.html)
- [Reconsidering Backends for Frontends (BFF)](https://zenn.dev/morinokami/scraps/20a4eab9415a50)
- [More coverage on BFFs](https://samnewman.io/blog/2016/02/14/more-coverage-on-bffs/)
- [Embracing the Differences: Inside the Netflix API Redesign](https://netflixtechblog.com/embracing-the-differences-inside-the-netflix-api-redesign-15fd8b3dc49d)
- [BFF @ SoundCloud](https://www.thoughtworks.com/insights/blog/bff-soundcloud)
- [Moving to Microservices at SoundCloud with Lukasz Plotnicki](https://softwareengineeringdaily.com/2016/02/04/moving-to-microservices-at-soundcloud-with-lukasz-plotnicki/)
- [Started Backends For Frontends (BFF)](https://techblog.zozo.com/entry/zozo-aggregation-api-bff)
- [Backends for Frontends Pattern](https://aws.amazon.com/jp/blogs/news/backends-for-frontends-pattern/)
- [What is BFF?](https://qiita.com/souhei-etou/items/d5de99bb8cba1c59d393)
- [Should BFF (Backend for Frontend) be adopted in new development?](https://vivit.hatenablog.com/entry/2021/11/10/101530)
- [Backend for frontend (BFF) pattern— why do you need to know it?](https://medium.com/mobilepeople/backend-for-frontend-pattern-why-you-need-to-know-it-46f94ce420b0)
- [What is the trending BFF architecture? | Offers Tech Blog](https://zenn.dev/overflow_offers/articles/20220418-what-is-bff-architecture)
- [Why did we unify GraphQL to code-first? To maintain consistency in type definitions for BFF/FE organization](https://logmi.jp/tech/articles/326592)
- [Are BFF (Backend for Frontend) and DDD mutually exclusive?](https://stackoverflow.com/questions/76940683/are-bff-backend-for-frontend-and-ddd-mutually-exclusive)
- [Do frontend engineers dream of Micro Frontends?](https://engineering.mercari.com/blog/entry/2018-12-06-162827/)

# Thoughts
I was already familiar with BFF, so I thought I would just do a quick search and be done with it. However, I found it interesting to consider various points such as the availability of architecture, handling of business logic, appropriate aggregation of clients, and its relation to organizational structure.

I have the impression that BFF seems to have many pitfalls that require caution. While the traps are visible, it feels challenging to create it in a way that avoids them.

If I have the opportunity to consider BFF, I will reflect on this.