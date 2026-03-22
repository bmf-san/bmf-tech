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
This is a summary of what I researched about service meshes.

# What is a Service Mesh?
A network infrastructure that manages communication between services (in a distributed system).

Typically, it is configured by adding a proxy as a sidecar to the services.

# What a Service Mesh Solves
- Improved observability
- Management and control of communication between services
- Enhanced security of communication between services

# Features of a Service Mesh
- Service Discovery
  - Managing connection information for services via DNS or load balancers can cause delivery delays.
  - By delegating the management of connection information to the service mesh, delivery delays can be minimized.
- Service Routing
  - By allowing the service mesh to handle the routing of traffic between services, the routing can be flexibly changed.
- Fault Isolation
  - e.g. Circuit Breaker
- Load Balancing
- Authentication and Authorization
  - Instead of preparing mechanisms for each service, delegating to the service mesh can reduce the burden on each service.
- Observability
  - Enhances the observability of traffic that spans multiple services.
    - Enables tracing.

# Disadvantages of a Service Mesh
The following are common disadvantages associated with service meshes that require proxies. Proxyless service meshes (e.g. Traffic Director) are not subject to these limitations.

- Decreased communication performance
- Increased resource usage

# References
- [speakerdeck.com - Service Mesh Comprehensive Introduction/Get-Started-Service-Mesh](https://speakerdeck.com/oracle4engineer/get-started-service-mesh)
- [cloud.google.com - Service Mesh in Microservices Architecture](https://cloud.google.com/architecture/service-meshes-in-microservices-architecture?hl=en)
- [www.netone.co.jp - Introduction to Service Mesh](https://www.netone.co.jp/media/detail/20200715-1/)
- [aws.amazon.com - What is a Service Mesh?](https://aws.amazon.com/what-is/service-mesh/)
- [www.alpha.co.jp - Things to Know Before Implementing a Service Mesh](https://www.alpha.co.jp/blog/202205_01)
- [www.redhat.com - What is a Service Mesh?](https://www.redhat.com/topics/microservices/what-is-a-service-mesh)
- [qiita.com - Investigating Service Mesh](https://qiita.com/mamomamo/items/92085e0e508e18bc8532)
- [www.infoq.com - Essential Guide to Service Mesh - Managing Inter-Service Communication in the Microservices Era](https://www.infoq.com/articles/service-mesh-ultimate-guide/)
- [https://dev.classmethod.jp - Understanding Service Mesh](https://dev.classmethod.jp/articles/servicemesh/)
