---
title: About Service Mesh
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
This is a summary of what I researched about service mesh.

# What is Service Mesh
A network infrastructure for managing communication between services (distributed systems).

Typically configured by adding a proxy as a sidecar to the service.

# What Service Mesh Solves
- Improved observability
- Management and control of communication between services
- Enhanced security of communication between services

# Features of Service Mesh
- Service Discovery
  - Managing connection information for services via DNS or load balancers can cause delivery delays.
  - By delegating connection information management to the service mesh, delivery delays can be mitigated.
- Service Routing
  - By entrusting traffic routing between services to the service mesh, routing can be flexibly changed.
- Fault Isolation
  - e.g. Circuit Breaker
- Load Balancing
- Authentication and Authorization
  - Instead of preparing mechanisms for each service, delegating to the service mesh can reduce the burden on each service.
- Observability
  - Enhances observability of traffic spanning multiple services.
    - Enables tracing.

# Disadvantages of Service Mesh
The following are disadvantages of service meshes that require a typical proxy. Proxyless service meshes (e.g. Traffic Director) are not subject to these limitations.

- Decreased communication performance
- Increased resource usage

# References
- [speakerdeck.com - Service Mesh Comprehensive Introduction/Get-Started-Service-Mesh](https://speakerdeck.com/oracle4engineer/get-started-service-mesh)
- [cloud.google.com - Service Mesh in Microservices Architecture](https://cloud.google.com/architecture/service-meshes-in-microservices-architecture?hl=ja)
- [www.netone.co.jp - Introduction to Service Mesh](https://www.netone.co.jp/media/detail/20200715-1/)
- [aws.amazon.com - What is a Service Mesh?](https://aws.amazon.com/jp/what-is/service-mesh/)
- [www.alpha.co.jp - Things to Know Before Implementing Service Mesh](https://www.alpha.co.jp/blog/202205_01)
- [www.redhat.com - What is a Service Mesh](https://www.redhat.com/ja/topics/microservices/what-is-a-service-mesh)
- [qiita.com - Investigating Service Mesh](https://qiita.com/mamomamo/items/92085e0e508e18bc8532)
- [www.infoq.com - Essential Guide to Service Mesh - Managing Inter-Service Communication in the Microservices Era](https://www.infoq.com/jp/articles/service-mesh-ultimate-guide/)
- [https://dev.classmethod.jp - Understanding Service Mesh](https://dev.classmethod.jp/articles/servicemesh/)