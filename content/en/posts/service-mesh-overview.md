---
title: What is a Service Mesh? How Istio and Linkerd Work Explained
description: 'Learn what a service mesh is, how the sidecar proxy pattern works, and what problems Istio and Linkerd solve for service communication inside a Kubernetes cluster.'
slug: service-mesh-overview
date: 2023-10-29T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Service Mesh
translation_key: service-mesh-overview
---



# Overview
This post summarizes what I have researched about service mesh.

# What is a Service Mesh?
A service mesh is a network infrastructure for managing communication between services (distributed systems).

Generally, it is configured by adding a proxy as a sidecar to the service.

# What Service Mesh Solves
- Improved observability
- Management and control of communication between services
- Enhanced security of communication between services

# Features of Service Mesh
- Service Discovery
  - Managing service connection information with DNS or load balancers can cause delivery delays
  - By entrusting the management of service connection information to the service mesh, delivery delays can be reduced
- Service Routing
  - By entrusting the routing of traffic between services to the service mesh, the routing of traffic between services can be flexibly changed
- Fault Isolation
  - ex. Circuit Breaker
- Load Balancing
- Authentication and Authorization
  - Instead of preparing mechanisms for each service, entrusting it to the service mesh can reduce the burden on each service
- Observability
  - Improves the observability of traffic spanning multiple services
    - Enables tracing

# Disadvantages of Service Mesh
The following are disadvantages of service meshes that require a general proxy. Proxyless service meshes (e.g., Traffic Director) are exceptions.

- Decreased performance of communication
- Increased resource usage

# References
- [speakerdeck.com - Service Meshがっつり入門/Get-Started-Service-Mesh](https://speakerdeck.com/oracle4engineer/get-started-service-mesh)
- [cloud.google.com - マイクロサービス アーキテクチャのサービス メッシュ](https://cloud.google.com/architecture/service-meshes-in-microservices-architecture?hl=ja)
- [www.netone.co.jp - サービスメッシュ入門](https://www.netone.co.jp/media/detail/20200715-1/)
- [aws.amazon.com - サービスメッシュとは何ですか?](https://aws.amazon.com/jp/what-is/service-mesh/)
- [www.alpha.co.jp - サービスメッシュ導入の前に知っておくべきこと](https://www.alpha.co.jp/blog/202205_01)
- [www.redhat.com - サービスメッシュとは](https://www.redhat.com/ja/topics/microservices/what-is-a-service-mesh)
- [qiita.com - サービスメッシュについて調査してみた件](https://qiita.com/mamomamo/items/92085e0e508e18bc8532)
- [www.infoq.com - サービスメッシュ必読ガイド - マイクロサービス時代のサービス間通信管理](https://www.infoq.com/jp/articles/service-mesh-ultimate-guide/)
- [https://dev.classmethod.jp -サービスメッシュについて理解する ](https://dev.classmethod.jp/articles/servicemesh/)
