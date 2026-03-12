---
title: "Modular Monolith Architecture: A Practical Guide and When to Choose It Over Microservices"
slug: modular-monolith-research-notes
date: 2023-09-25T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Modular Monolith
translation_key: modular-monolith-research-notes
---



# Overview
This is a memo on what I researched about modular monoliths.

# What is a Modular Monolith?
- A monolith divided into modules
  - While domain-based division is commonly seen, various patterns such as functional or technical division can be considered
- Like a monolith, it has a single deployment pipeline
- Benefits
  - Modules are divided, allowing independent development per module
  - Easier transition to microservices
    - Easier to progress with microservice conversion per module
      - When transitioning from a monolith to microservices, it can be positioned as a strangler pattern (though it may not always be easy)
    - Easier to review module boundaries, allowing more flexible adaptation to boundary changes compared to microservices
- Drawbacks
  - Easier to cross module boundaries
    - Microservices communicate over a network, preventing boundary crossing, but modular monoliths do not, so care must be taken with boundary violations
  - Like a monolith, it has a single deployment pipeline, so if it becomes bloated or module dependencies become complex, operation becomes difficult
  - If each module shares a single database, the cost of transitioning to microservices increases

# Service Weaver
Google has released a tool called [Service Weaver](https://serviceweaver.dev/) for developing as a modular monolith and deploying as microservices.

# Impressions
I had a bit of a dream about modular monoliths, but my impression led me to write a poem.

I think it's logical that architecture needs to evolve with organizational expansion, but I wondered if there might be a silver bullet architecture that can flexibly respond to organizational scalability or doesn't incur excessive costs. (There isn't.)

Organizations may expand, contract, or remain unchanged, but since companies assume growth, I thought it might be good to invest proactively in organizational scalability.

I wanted to quote a passage related to this topic, so I'll include it here to conclude.

> However, what we need to focus on here is the difference in lifecycle between the two. Organizational and team configurations can be changed the next day depending on the company's policy if desired. However, architecture and systems are difficult to change quickly like an organization.

cf. [eh-career.com - Reasons for Transitioning to a Modular Monolith ─ Asoview's Large-Scale Efforts to Balance Microservice Autonomy and Monolith Consistency](https://eh-career.com/engineerhub/entry/2022/07/25/093000)

# References
- [microservices.io - How modular can your monolith go? Part 1 - the basics](https://microservices.io/post/architecture/2023/07/31/how-modular-can-your-monolith-go-part-1.html)
- [microservices.io - How modular can your monolith go? Part 2 - a first look at how subdomains collaborate](https://microservices.io/post/architecture/2023/08/20/how-modular-can-your-monolith-go-part-2.html)
- [microservices.io - How modular can your monolith go? Part 3 - encapsulating a subdomain behind a facade](https://microservices.io/post/architecture/2023/08/28/how-modular-can-your-monolith-go-part-3.html)
- [microservices.io - How modular can your monolith go? Part 4 - physical design principles for faster builds](https://microservices.io/post/architecture/2023/09/12/how-modular-can-your-monolith-go-part-4-physical-design.html)
- [techblog.hacomono.jp - Introducing Modular Monoliths to a Monolithic Rails](https://techblog.hacomono.jp/entry/2023/08/22/110000#:~:text=%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%A9%E3%83%BC%E3%83%A2%E3%83%8E%E3%83%AA%E3%82%B9%E3%81%A8%E3%81%AF%E3%80%81%E3%83%A2%E3%83%8E%E3%83%AA%E3%82%B9,%E6%B4%BB%E7%94%A8%E3%81%A7%E3%81%8D%E3%82%8B%E7%89%B9%E5%BE%81%E3%81%8C%E3%81%82%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82)
- [tech-blog.rakus.co.jp - Monoliths Prepared for Service Division (Modular Monoliths, Aggregates, etc.)](https://tech-blog.rakus.co.jp/entry/20201026/microservice)
- [r-kaga.com - Summarizing Modular Monoliths](https://r-kaga.com/blog/what-is-modular-monolith)
- [speakerdeck.com - Expressing Complex Domain Regions and Boundaries with Modular Monoliths](https://speakerdeck.com/showmant/expressing-complex-domain-regions-and-boundaries-with-modular-monoliths)
- [shopify.engineering - Deconstructing the Monolith: Designing Software that Maximizes Developer Productivity](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity)
- [www.infoq.com - Microservices are not inevitable in monolith decomposition - A talk by Sam Newman at QCon London](https://www.infoq.com/jp/news/2020/06/monolith-decomposition-newman/?utm_campaign=infoq_content&utm_source=infoq&utm_medium=feed&utm_term=global)
- [www.infoq.com - How Shopify Transitioned to a Modular Monolith](https://www.infoq.com/jp/news/2019/10/shopify-modular-monolith/)
- [www.publickey1.jp - Google Releases "Service Weaver" Framework, Combining the Best of Monoliths and Microservices, as Open Source](https://www.publickey1.jp/blog/23/googleservice_weaver.html)
- [Transaction Design in Modular Monoliths](https://speakerdeck.com/nazonohito51/transaction-design-on-modular-monolith)
- [dev.classmethod.jp - [Report] For Those Hesitating Between Monoliths and Microservices #devio_day1 #main](https://dev.classmethod.jp/articles/20230411-developersio-day-one-monolithandmicroservices/)
- [engineering.mercari.com - Mercari's Efforts in Modular Monolithization of the Transaction Domain](https://engineering.mercari.com/blog/entry/20220913-modular-monolithization-in-mercari-transaction-domain/)
- [eh-career.com - Reasons for Transitioning to a Modular Monolith ─ Asoview's Large-Scale Efforts to Balance Microservice Autonomy and Monolith Consistency](https://eh-career.com/engineerhub/entry/2022/07/25/093000)
- [logmi.jp - The Importance of "Flexible Architecture Design for the Future" - Adopting Modular Monoliths for Parallel Development by Three Teams](https://logmi.jp/tech/articles/328130)
- [medium.com - The Modular Monolith: Rails Architecture](https://medium.com/@dan_manges/the-modular-monolith-rails-architecture-fb1023826fc4)
