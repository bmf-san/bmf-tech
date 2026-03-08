---
title: Structure of My System Specification Document
slug: system-specification-structure
date: 2025-05-19T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - System Design
  - Design
  - Tips
translation_key: system-specification-structure
---


# Overview
Sometimes, I ponder over what structure a system specification document should have when I want to write it down.

Here, I describe the structure of my system specification document.

This structure should be useful not only during the design phase of a new project but also for understanding the specifications of an existing system.

# Structure
```sh
01_overview.md             # Overall Overview
02_system_architecture.md  # System Architecture
03_data_model.md           # Data Model (Conceptual Design)
04_web_endpoint.md         # List of Web Endpoints
05_api.md                  # API Specifications (Mainly REST, etc.)
06_technical_detail.md     # Technical Details (Libraries, Middleware, etc.)
07_usecase.md              # Use Cases
08_sequence.md             # Sequence Diagrams & Process Flow
09_references.md           # References
```

The file names are mostly self-explanatory, but I'll note down some special points.

## 01_overview.md
In the overall overview, describe the purpose and general outline of the specification document.

- What the specification document is about
- Intended audience
- Purpose/Non-purpose
- Article management policy (usage purpose, update frequency, etc.)

## 02_system_architecture.md
I personally like to document the C4 Model using mermaid, but since mermaid can be difficult to maintain, it's important to think carefully.

Recently, AI has been quite helpful, so it's good to leave it in code form as much as possible.

## 03_data_model.md
In the data model, describe the abstract and concrete data design of the system.

For databases, include conceptual and physical design.

It would be good to have model design as well.

## 04_web_endpoint.md
List the endpoints that return screens.

Below is an example.

```sh
| Item           | Value                                 |
| ------------ | --------------------------------- |
| **Overview**       | A brief description of this endpoint. What it displays or operates.  |
| **Path**       | `/your/path/here`                 |
| **HTTP Method** | `GET` / `POST` / `PUT` / `DELETE` |
| **Controller**  | `Your::Controller#action`         |
| **Authentication**       | `authenticate_user!` etc.           |
| **Authorization**       | `authorize YourPolicy` etc.         |
```

Authentication and authorization should be appropriately tailored to the concerns of the relevant web endpoint and the system context.

## 05_api.md
List the API endpoints.

Below is an example.
```sh
| API Name       | Endpoint                   | Method                  | Description              |
| ---------- | ------------------------- | --------------------- | --------------- |
| API Name   | `/api/v1/sample_endpoint` | GET/POST/PATCH/DELETE | Overview of this API     |
| Data Retrieval API | `/api/v1/items`           | GET                   | Retrieve a list of items     |
| Search API    | `/api/v1/items/search`    | POST                  | Search items based on conditions |
```
         
## 06_technical_detail.md
Write about the technical details.

This section delves into important points in the system.

It is the most important and the most labor-intensive section.

## 07_usecase.md
Write about the list of use cases.

Below is an example assuming multiple systems.

```sh
|   System   |  User  |       Category       |              Use Case               |
| ------------ | ---------- | -------------------- | --------------------------------------- |
| example       | Administrator     | Admin Functions           | List Display                          |
```

## 08_sequence.md
Write sequences according to the list of use cases.

## 09_references.md
Section about related materials.

# Thoughts 
If the documentation is well-organized from the start, it makes it easier to form a mental model and align understanding with others when you can write in this format for your understanding.

Depending on the case, when maintenance becomes difficult, it might be an opportunity to realize that the problem lies not in the way the document is written but in the complexity of the system.
