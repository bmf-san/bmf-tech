---
title: Cohesion and Coupling
description: An in-depth look at Cohesion and Coupling, covering key concepts and practical insights.
slug: cohesion-coupling
date: 2025-06-25T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Design
  - System Design
translation_key: cohesion-coupling
---

# Overview
In software design, **Cohesion** and **Coupling** are fundamental metrics for measuring quality. These concepts are crucial when considering modularization, maintainability, and reusability.

# What is Cohesion?
Cohesion is a metric that represents **how closely related the components (such as functions and variables) within a module are**.

## High Cohesion
- Refers to a state where elements within a module collaborate for "the same purpose."
- This is an ideal state, making the module easy to understand, maintain, and reuse.

## Low Cohesion
- Refers to a state where unrelated functions are mixed within the same module.
- This makes understanding and testing difficult, and changes tend to have a wide impact.

## Types of Cohesion (Low → High)

| Type                   | Description                       |
| -------------------- | ------------------------ |
| Coincidental          | Unrelated processes are just grouped together        |
| Logical               | Processes of the same category, selected by conditions  |
| Temporal              | Processes executed at the same time (e.g., initialization) |
| Procedural            | A series of processes with different purposes            |
| Communicational       | Processes that handle the same data              |
| Sequential            | Output becomes the input for the next process            |
| Functional            | Specialized for a single, clear purpose (ideal)     |

# What is Coupling?
Coupling is a metric that represents **the strength of dependencies between modules**.

## Low Coupling
- Refers to a state where dependencies between modules are minimal.
- This is an ideal state, with minimal ripple effects from changes, making testing and reuse easier.

## High Coupling
- Refers to a state where modules are strongly dependent on each other.
- Changes in one part can potentially have a wide impact on other modules.

## Types of Coupling (High → Low)
| Type               | Description                      |
| ---------------- | ----------------------- |
| Content           | Directly accessing the internals of another module      |
| Common            | Sharing global variables          |
| External          | Depending on external formats (such as file formats) |
| Control           | Delegating control via flags            |
| Stamp             | Passing structures containing unnecessary data         |
| Data              | Passing only the minimal necessary data (desirable)    |
| Message           | Complete separation through message passing (ideal)   |

# Ideal
| Metric  | Ideal State             |
| --- | ----------------- |
| Cohesion | The higher, the better (focused on purpose) |
| Coupling | The lower, the better (less dependency)    |

# Example: Designing a Logging Feature
## Good Example: High Cohesion & Low Coupling

```go
type Logger struct {
    Output io.Writer
}

func (l *Logger) Info(msg string)  { ... }
func (l *Logger) Error(msg string) { ... }
```

- Focused on the clear purpose of log output, achieving high cohesion.
- Other modules depend only on the Logger interface, achieving low coupling.

## Bad Example: Low Cohesion & High Coupling

```go
func DoStuffAndLog() {
    // Data processing
    // DB update
    // Send email on error
    // Log output
}
```

- Unrelated processes are mixed, resulting in low cohesion.
- Depends on multiple other modules, resulting in high coupling.

# Summary
Cohesion and coupling are crucial criteria that influence design quality. By clarifying module responsibilities and minimizing dependencies, a system design that is easy to understand, maintain, and scale can be achieved.
