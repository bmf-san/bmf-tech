---
title: Notes on Modular Monoliths
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
This is a note on what I researched about modular monoliths.

# What is a Modular Monolith?
- A monolith that is divided into modules
  - While module division generally appears to be based on domains, various patterns such as functional division and technical division can be considered.
- Like a monolith, it has a single deployment pipeline.
- Advantages
  - Since the modules are divided, development can be done independently at the module level.
  - Easier transition to microservices
    - It is easier to progress towards microservices on a module basis.
      - When transitioning from a monolith to microservices, it can be positioned as a strangler pattern (however, it does not guarantee an easy transition).
    - It is easier to review the boundaries of modules, allowing for more flexible responses to boundary changes than microservices.
- Disadvantages
  - It is easy to cross boundaries between modules.
    - Microservices cannot cross module boundaries as they communicate over a network, but modular monoliths can, so care must be taken with boundary violations.
  - Like a monolith, it has a single deployment pipeline, so if it becomes bloated or if dependencies between modules become complex, operations can become difficult.
  - If each module shares a single database, the cost of transitioning to microservices can be high.

# Service Weaver
Google has released a tool called [Service Weaver](https://serviceweaver.dev/) for developing as a modular monolith and deploying as microservices.

# Thoughts
I had a bit of a dream about modular monoliths, but I felt like writing a poem as a reflection.

I think it’s natural for architecture to evolve with the expansion of the organization, but I wondered if there is an architecture that can flexibly respond to organizational scalability or does not incur excessive costs, like a silver bullet. (There isn’t.)

Organizations may expand, contract, or remain unchanged at times, but since companies assume growth, it might be good to invest proactively in organizational scalability.

I wanted to quote a passage related to this topic, so I will include it here to conclude.

> However, what we must pay attention to here is the difference in the lifecycle of the two. Organizations and team configurations can be changed from the next day depending on the company’s policy if they want to. However, architecture and systems are difficult to change as easily as organizations.

cf. [eh-career.com - Reasons for Transitioning to Modular Monoliths: AsoView's Efforts to Balance Microservices Autonomy and Monolith Consistency](https://eh-career.com/engineerhub/entry/2022/07/25/093000)

# References
- [microservices.io - How modular can your monolith go? Part 1 - the basics](https://microservices.io/post/architecture/2023/07/31/how-modular-can-your-monolith-go-part-1.html)
- [microservices.io - How modular can your monolith go? Part 2 - a first look at how subdomains collaborate](https://microservices.io/post/architecture/2023/08/20/how-modular-can-your-monolith-go-part-2.html)
- [microservices.io - How modular can your monolith go? Part 3 - encapsulating a subdomain behind a facade](https://microservices.io/post/architecture/2023/08/28/how-modular-can-your-monolith-go-part-3.html)
- [microservices.io - How modular can your monolith go? Part 4 - physical design principles for faster builds](https://microservices.io/post/architecture/2023/09/12/how-modular-can-your-monolith-go-part-4-physical-design.html)
- [techblog.hacomono.jp - Introducing Modular Monoliths to a Monolithic Rails Application](https://techblog.hacomono.jp/entry/2023/08/22/110000#:~:text=%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%A9%E3%83%BC%E3%83%A2%E3%83%8E%E3%83%AA%E3%82%B9%E3%81%A8%E3%81%AF%E3%80%81%E3%83%A2%E3%83%8E%E3%83%AA%E3%82%B9,%E6%B4%BB%E7%94%A8%E3%81%A7%E3%81%8D%E3%82%8B%E7%89%B9%E5%BE%B4%E3%81%8C%E3%81%82%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82)
- [tech-blog.rakus.co.jp - Monoliths Prepared for Service Splitting (Modular Monoliths and Aggregates)](https://tech-blog.rakus.co.jp/entry/20201026/microservice)
- [r-kaga.com - Summary of Modular Monoliths](https://r-kaga.com/blog/what-is-modular-monolith)
- [speakerdeck.com - Expressing Complex Domain Regions and Boundaries with Modular Monoliths](https://speakerdeck.com/showmant/expressing-complex-domain-regions-and-boundaries-with-modular-monoliths)
- [shopify.engineering - Deconstructing the Monolith: Designing Software that Maximizes Developer Productivity](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity)
- [www.infoq.com - Microservices are not Inevitable in Monolith Decomposition - Insights from Sam Newman's Talk at QCon London](https://www.infoq.com/jp/news/2020/06/monolith-decomposition-newman/?utm_campaign=infoq_content&utm_source=infoq&utm_medium=feed&utm_term=global)
- [www.infoq.com - How Shopify Transitioned to a Modular Monolith](https://www.infoq.com/jp/news/2019/10/shopify-modular-monolith/)
- [www.publickey1.jp - Google Releases Open Source Framework 'Service Weaver' that Takes the Best of Monoliths and Microservices](https://www.publickey1.jp/blog/23/googleservice_weaver.html)
- [Considerations for Transaction Design in Modular Monoliths](https://speakerdeck.com/nazonohito51/transaction-design-on-modular-monolith)
- [dev.classmethod.jp - [Report] For Those Confused About Choosing Between Monoliths and Microservices #devio_day1 #main](https://dev.classmethod.jp/articles/20230411-developersio-day-one-monolithandmicroservices/)
- [engineering.mercari.com - Efforts to Modularize the Transaction Domain at Mercari](https://engineering.mercari.com/blog/entry/20220913-modular-monolithization-in-mercari-transaction-domain/)
- [eh-career.com - Reasons for Transitioning to Modular Monoliths: AsoView's Efforts to Balance Microservices Autonomy and Monolith Consistency](https://eh-career.com/engineerhub/entry/2022/07/25/093000)
- [logmi.jp - The Importance of 'Forward-Looking Flexible Architecture': Adoption of Modular Monoliths Achieving Parallel Development by Three Teams](https://logmi.jp/tech/articles/328130)
- [medium.com - The Modular Monolith: Rails Architecture](https://medium.com/@dan_manges/the-modular-monolith-rails-architecture-fb1023826fc4)