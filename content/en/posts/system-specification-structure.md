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
Sometimes I think about what structure should be used when writing out the specifications of a system in a document.

Here is the structure of my system specification document.

It should be useful not only for the design phase of new projects but also for understanding the specifications of existing systems.

# Structure
```sh
01_overview.md             # Overall Overview
02_system_architecture.md  # System Architecture
03_data_model.md           # Data Model (Conceptual Design)
04_web_endpoint.md         # List of Web Endpoints
05_api.md                  # API Specifications (mainly REST, etc.)
06_technical_detail.md     # Technical Details (libraries, middleware, etc.)
07_usecase.md              # Use Cases
08_sequence.md             # Sequence Diagrams and Process Flow
09_references.md           # References
```

The file names are mostly self-explanatory, so not much explanation is needed, but I will outline some special notes.

## 01_overview.md
In the overall overview, describe the purpose and outline of the specification document.

- What the specification document is about
- Intended audience
- Purpose / Out of scope
- Article management policy (usage, update frequency, etc.)

## 02_system_architecture.md
Personally, I like to document the C4 Model using mermaid, but sometimes maintenance can be difficult with mermaid, so I think carefully about it.

Recently, AI has been quite helpful, so it would be good to code it as much as possible.

## 03_data_model.md
In the data model, describe the data design of the system's abstraction and concreteness.

If it's a database, include both conceptual and physical design.

It would also be good to have a model design.

## 04_web_endpoint.md
List the endpoints that return screens.

Here is an example.

```sh
| Item           | Value                                 |
| ------------ | --------------------------------- |
| **Overview**       | A brief description of this endpoint. What it displays or operates.  |
| **Path**       | `/your/path/here`                 |
| **HTTP Method** | `GET` / `POST` / `PUT` / `DELETE` |
| **Controller**  | `Your::Controller#action`         |
| **Authentication**       | `authenticate_user!`  etc.           |
| **Authorization**       | `authorize YourPolicy`  etc.         |
```

Authentication and authorization should be tailored to the concerns of the relevant web endpoints and the context of the system.

## 05_api.md
List the API endpoints.

Here is an example.
```sh
| API Name       | Endpoint                   | Method                  | Description              |
| ---------- | ------------------------- | --------------------- | --------------- |
| Specify API Name   | `/api/v1/sample_endpoint` | GET/POST/PATCH/DELETE | Brief overview of this API     |
| Data Retrieval API | `/api/v1/items`           | GET                   | Retrieve the list of items     |
| Search API    | `/api/v1/items/search`    | POST                  | Search for items based on conditions |
```

## 06_technical_detail.md
Write about the technical details.

This section delves into important points regarding the system.

It is the most important section and also the most labor-intensive section.

## 07_usecase.md
Write about the list of use cases.

Here is an example assuming multiple systems.

```sh
|   System   |  User  |       Category       |              Use Case               |
| ------------ | ---------- | -------------------- | --------------------------------------- |
| example       | Administrator     | Administrator Functions           | List Display                          |
```

## 08_sequence.md
Write sequences according to the list of use cases.

## 09_references.md
A section for related materials.

# Thoughts 
If the documentation is well-organized from the start, that's one thing, but being able to write in this format for my own understanding makes it easier to form a mental model and align recognition with others.

It varies case by case, but when maintenance becomes difficult, it can be a trigger to realize that the issue is not with the way the document is written but rather the complexity of the system.