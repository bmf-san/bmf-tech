---
title: "8 Japanese Books for Learning Microservices (Concepts, Migration, Patterns, and Operations)"
slug: microservices-recommended-books
date: 2026-11-06T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Book Review
  - Microservices
description: "A staged reading list of 8 Japanese books for learning microservices: the big picture, migrating from a monolith, design patterns, API and data design, and running services in production. Each entry links to its own review."
translation_key: microservices-recommended-books
recommended_books:
  - asin: "4814400012"
    title: "Building Microservices, 2nd Edition"
  - asin: "4873119316"
    title: "Monolith to Microservices"
  - asin: "4295008583"
    title: "Microservices Patterns"
  - asin: "4873118751"
    title: "Designing Distributed Systems"
  - asin: "4814400896"
    title: "Mastering API Architecture"
  - asin: "4297140101"
    title: "Driving the Domain with Data Modeling"
  - asin: "4873118158"
    title: "Production-Ready Microservices"
  - asin: "4873119979"
    title: "Distributed Services with Go"
draft: false
---

A reading list for anyone who wants to learn how microservices work and how to design, migrate, and run them. Each pick has its own review on this blog, and every entry links to it. The list starts with the big picture, then moves through migration, design patterns, API and data design, and production operations.

The books in this list are Japanese editions, and the store links go to Amazon Japan.

## Why learn microservices

Microservices split a system into services that teams can deploy on their own, so teams and systems can evolve in parallel. They also bring the complexity of a distributed system, the work of splitting data, and a heavier operational load. To decide whether to adopt them, and to keep them healthy after you do, you need to know both the benefits and the costs.

## 1. Get the big picture

### Building Microservices, 2nd Edition

The Japanese edition of Sam Newman's book. Read it first if you work on microservices or plan to adopt them. It also covers the architecture characteristics that distributed systems need, and chapter 1 gives a compact overview of the core concepts. I cover it in the [review](/posts/microservices-architecture-second-edition/).

## 2. Migrate from a monolith

### Monolith to Microservices

Also by Sam Newman, this book focuses on how to move from a monolith to microservices. The longest chapter deals with splitting the database, the hardest part of most migrations. The author does not treat microservices as the default answer, which helps you step back before you commit. See the [review](/posts/monolith-to-microservices-migration-guide/).

## 3. Learn design patterns

### Microservices Patterns

Chris Richardson's catalog of microservice design patterns, and a standard reference. It runs from decomposition strategies through distributed transactions with sagas, CQRS, event sourcing, testing, and deployment. Each pattern comes with the problem it solves and its trade-offs, so you can lean on it when you make design calls at work. I wrote it up in the [review](/posts/microservices-patterns/).

### Designing Distributed Systems

Brendan Burns's book on design patterns for distributed systems that run on containers. Turn to it when you plan the platform side that runs your services. I cover it in the [review](/posts/distributed-system-design-patterns/).

## 4. Design APIs and data

### Mastering API Architecture

An overview of API architecture patterns, testing, operations, and security. The opening chapters explain C4 diagrams and ADRs, and the rest of the book uses them, so you also learn how to communicate a design. Expect a broad survey, and fill in individual topics with other books. See the [review](/posts/mastering-api-architecture/).

### Driving the Domain with Data Modeling

A Japanese book on data modeling for splitting core business systems into distributed, loosely coupled parts, built around the idea of ledgers. It gives practical guidance on data design when you move toward DDD or microservices. Its terms are unique to the book, so some background in enterprise systems helps. I wrote it up in the [review](/posts/domain-driven-data-modeling/).

## 5. Run in production and build

### Production-Ready Microservices

Susan Fowler's book sets out what microservices need before they can run in production: stability, reliability, scalability, fault tolerance, monitoring, and documentation. The checklists at the end of each chapter make a good starting point for your own production-readiness standards. I cover it in the [review](/posts/production-ready-microservices/).

### Distributed Services with Go

The Japanese edition of Travis Jeffery's book on building distributed services in Go. It also explains the modern tools that microservice setups rely on and the reasons behind their design. See the [review](/posts/distributed-services-with-golang/).

## Wrap-up

Start with Building Microservices for the concepts, and read Monolith to Microservices if you plan a migration. Microservices Patterns is the book to reach for when a design call gets hard. Keep the API, data, and operations books for the stage when you design and run real services. Each book has its own review, so start with whichever one interests you.

## Related roundups

- [📚 All tech book roundups](/posts/tech-book-roundups/)
- [Software architecture books](/posts/software-architecture-recommended-books/)
- [System design and architect books](/posts/system-design-recommended-books/)
- [Monitoring, SRE, and operations books](/posts/sre-operations-recommended-books/)
- [Go books](/posts/go-recommended-books/)
