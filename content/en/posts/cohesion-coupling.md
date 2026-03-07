---
title: Cohesion and Coupling
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
In software design, **Cohesion** and **Coupling** are fundamental metrics for measuring quality. These concepts are important when considering modularity, maintainability, and reusability.

# What is Cohesion?
Cohesion is a metric that indicates **how closely related the components (functions, variables, etc.) within a module are**.

## High Cohesion
- Refers to a state where the elements within a module collaborate for a "common purpose".
- This is the ideal state, making it easy to understand, maintain, and reuse.

## Low Cohesion
- Refers to a state where unrelated functionalities are mixed within the same module.
- This makes it difficult to understand and test, and the impact of changes can be widespread.

## Types of Cohesion (Low → High)

| Type                   | Description                       |
| -------------------- | ------------------------ |
| Coincidental (偶発的)    | Unrelated processes are simply grouped together        |
| Logical (論理的)         | Processes of the same category but selected by conditions  |
| Temporal (時間的)        | Processes executed at the same time (e.g., initialization) |
| Procedural (手続き的)     | A series of processes with different purposes            |
| Communicational (通信的) | Groups of processes handling the same data              |
| Sequential (逐次的)      | The output of one process becomes the input of the next            |
| Functional (機能的)      | Specialized for a single, clear purpose (ideal)     |

# What is Coupling?
Coupling is a metric that indicates **the strength of dependencies between modules**.

## Low Coupling
- Refers to a state where there are few dependencies between modules.
- This is the ideal state, minimizing the ripple effect of changes and making testing and reuse easier.

## High Coupling
- Refers to a state where modules are strongly dependent on each other.
- Changes in one module can have widespread effects on others.

## Types of Coupling (High → Low)
| Type               | Description                      |
| ---------------- | ----------------------- |
| Content (内容結合)    | Direct access to the internals of other modules      |
| Common (共通結合)     | Sharing global variables, etc.          |
| External (外部結合)   | Dependence on external formats (file types, etc.) |
| Control (制御結合)    | Delegating control via flags            |
| Stamp (スタンプ結合)    | Passing structures containing unnecessary data         |
| Data (データ結合)      | Passing only the necessary minimum data (desirable)    |
| Message (メッセージ結合) | Complete separation through message passing (ideal)   |

# Ideal
| Metric  | Ideal State             |
| --- | ----------------- |
| Cohesion | The higher, the better (focused on purpose) |
| Coupling | The lower, the better (less dependency)    |

# Example: Log Function Design
## Good Example: High Cohesion, Low Coupling

```go
type Logger struct {
    Output io.Writer
}

func (l *Logger) Info(msg string)  { ... }
func (l *Logger) Error(msg string) { ... }
```

- Focused on the clear purpose of log output, demonstrating high cohesion.
- Other modules depend only on the Logger interface, showing low coupling.

## Bad Example: Low Cohesion, High Coupling

```go
func DoStuffAndLog() {
    // Data processing
    // DB update
    // Send email on error
    // Log output
}
```

- Unrelated processes are mixed, resulting in low cohesion.
- Dependent on multiple other modules, leading to high coupling.

# Conclusion
Cohesion and coupling are crucial criteria that influence design quality. By clarifying the responsibilities of modules and minimizing dependencies, we can achieve a system design that is easy to understand, maintain, and grow.