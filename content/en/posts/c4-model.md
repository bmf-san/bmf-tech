---
title: What is the C4 Model?
slug: c4-model
date: 2024-08-15T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - C4 Model
  - Architecture
description: An introduction to the C4 Model, a technique for modeling software architecture.
translation_key: c4-model
---

# Overview
I often struggled with the level of granularity when creating architecture diagrams, but I recently learned about a technique called the C4 Model and decided to look into it.

# What is the C4 Model?
The C4 Model is one of the techniques for modeling software architecture.

C4 stands for Context, Containers, Components, and Code. By breaking a system into these elements, it provides a way to describe software architecture.

The C4 Model offers four views, ordered by decreasing levels of abstraction:

- Level 1: System Context Diagram
- Level 2: Container Diagram
- Level 3: Component Diagram
- Level 4: Code Diagram (e.g., UML or ER diagrams)

As the level increases, the internal structure of the system becomes more detailed.

# System Context Diagram
A diagram at the software system level. A software system includes one or more containers and represents the components that deliver value to users (e.g., applications, products, services).

This diagram shows the relationships and boundaries with external systems and expresses the interfaces with external systems and users.

# Container Diagram
A diagram showing the structure of containers that make up the target software system. Containers are elements that run to enable the entire software system to function (e.g., applications or data stores).

This diagram represents the high-level structure and responsibilities of the software architecture.

# Component Diagram
A diagram showing the structure of components that make up a container. Components are elements that run within a container (e.g., classes, modules, services).

This diagram expresses the responsibilities and implementation details of the components themselves.

# Code Diagram
A diagram showing the code within a component. Code diagrams refer to detailed diagrams such as UML or ER diagrams.

This diagram represents the details at the code level.

# How to Write
It seems best to thoroughly read the original source, [c4model.com - The C4 model for visualising software architecture Context, Containers, Components, and Code](https://c4model.com/), before creating diagrams.

The explanations are very clear, so referring to this while creating diagrams seems like a good approach.

This technique might help resolve the difficulty of aligning understanding about the granularity of architecture diagrams.

# Thoughts
I wish I had known about this technique earlier. I often struggled with deciding the level of granularity when creating architecture diagrams, but using the C4 Model seems like it will allow me to create diagrams based on clear standards.

# References
- [c4model.com - The C4 model for visualising software architecture Context, Containers, Components, and Code](https://c4model.com/)
- [ja.wikipedia.org - C4モデル](https://ja.wikipedia.org/wiki/C4%E3%83%A2%E3%83%87%E3%83%AB)
- [www.infoq.com - ソフトウェアアーキテクチャのためのC4モデル](https://www.infoq.com/jp/articles/C4-architecture-model/)
- [namaraii.com - ソフトウェアアーキテクチャのためのC4モデル](https://namaraii.com/notes/c4-model)
  - Writing C4 Model diagrams using mermaid