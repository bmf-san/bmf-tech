---
title: What is the C4 Model
description: 'Master C4 model architecture visualization with four levels: system context, containers, components, and code diagrams.'
slug: c4-model
date: 2024-08-15T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - Architecture
tags:
  - C4 Model
  - Architecture
translation_key: c4-model
---

# Overview
I have always struggled with the level of detail when drawing architecture diagrams, but I learned about a technique called the C4 model, so I decided to research it.

# What is the C4 Model
It is a modeling technique for software architecture.

C4 stands for Context, Containers, Components, and Code, and it describes software architecture by breaking down the system into these elements.

The C4 model provides the following four views in order of increasing abstraction:

- Level 1: System Context Diagram
- Level 2: Container Diagram
- Level 3: Component Diagram
- Level 4: Code Diagram (e.g., UML, ER, etc.)

As the level increases, the internal structure of the system becomes more detailed.

# System Context Diagram
A diagram at the level of the software system. A software system is one that includes one or more containers and consists of components that provide value to users (e.g., applications, products, services, etc.).

It shows the relationships and boundaries with the external environment of the system and represents the interfaces with external systems and users.

# Container Diagram
A diagram that shows the composition of the containers that make up the target software system. A container is an element that is running to enable the entire software system to function (e.g., applications or data stores).

It expresses the high-level composition and responsibilities of the software architecture.

# Component Diagram
A diagram that shows the composition of the components that make up the target container. A component is an element that runs within a container (e.g., classes, modules, services).

It expresses the responsibilities and implementation details of the components themselves.

# Code Diagram
A diagram that shows the code within the components. The code diagram refers to detailed diagrams such as UML or ER diagrams.

It expresses details at the code level.

# How to Write
It seems good to read the original source [c4model.com - The C4 model for visualising software architecture Context, Containers, Components, and Code](https://c4model.com/) thoroughly before writing.

It is explained quite clearly, so it seems good to refer to this while writing.

I feel that this can also resolve the difficulty of aligning perceptions about the level of detail in architecture diagrams.

# Thoughts
This is a technique I wish I had known about earlier. I often struggled with the level of detail when drawing architecture diagrams, but by using the C4 model, I feel I can draw diagrams according to clear standards.

# References
- [c4model.com - The C4 model for visualising software architecture Context, Containers, Components, and Code](https://c4model.com/)
- [ja.wikipedia.org - C4 Model](https://ja.wikipedia.org/wiki/C4%E3%83%A2%E3%83%87%E3%83%AB)
- [www.infoq.com - C4 Model for Software Architecture](https://www.infoq.com/jp/articles/C4-architecture-model/)
- [namaraii.com - C4 Model for Software Architecture](https://namaraii.com/notes/c4-model)
	- Writing each diagram of the C4 model with Mermaid.