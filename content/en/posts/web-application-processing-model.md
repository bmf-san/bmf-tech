---
title: Processing Models of Web Applications
slug: web-application-processing-model
date: 2025-08-02T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Event Loop
  - Thread
  - Process
translation_key: web-application-processing-model
---

In this article, we will discuss the representative processing models of web applications.

- Event Loop
- Thread
- Process

# Overview of Each Processing Model

## Event Loop Model (Asynchronous, Single-threaded)

- **Characteristics**: Processes events (like asynchronous IO) sequentially in a single thread.
- **Examples**: Node.js, Deno, JavaScript (Browser)

```
[ Event Queue ] → [ Event Loop ] → [ Execution ]
```

**Advantages**
- Memory efficient (does not create a large number of threads or processes)
- Can utilize IO wait for other processing

**Disadvantages**
- Prone to blocking with heavy computations (CPU-bound)

## Thread Model (Multi-threaded)

- **Characteristics**: Multiple threads operate concurrently within a single process.
- **Examples**: Java, Python, Go (internally uses thread scheduler)

```
[ Process ]
 ├─ Thread 1 (Request A)
 ├─ Thread 2 (Request B)
 └─ ...
```

**Advantages**
- Strong for CPU-bound processing (can utilize multi-core)

**Disadvantages**
- Requires management of resource consumption and race conditions for each thread

## Process Model (Multi-process)

- **Characteristics**: Each request is processed in a separate process. Completely isolated.
- **Examples**: PHP (FPM), Ruby (multiple Puma processes), Sidekiq (multi-process configuration)

```
[ Request A ] → [ Process 1 ]
[ Request B ] → [ Process 2 ]
```

**Advantages**
- High stability (less affected by other processes)

**Disadvantages**
- Poor memory efficiency (each process has its own environment)

# Processing Models and Their Suitability by Language

| Language            | Processing Model                     | IO-bound | CPU-bound | Notes                                   |
| -------------------| ------------------------------------| -------- | --------- | --------------------------------------- |
| Node.js            | Event Loop (Asynchronous)           | ◎        | △         | IO specialized, caution with heavy computations |
| Go                 | goroutine + M:N threads             | ◎        | ◎         | Good at parallel and concurrent processing with lightweight threads |
| Ruby (MRI)        | Threads + GIL (limited)            | ○        | △         | CPU processing is effectively single-threaded due to GIL |
| PHP (FPM)         | Process (1 request = 1 process)    | ○        | △         | Stable due to process isolation but has high overhead |
| Java               | Multi-threaded                      | ◎        | ◎         | Thread optimization at JVM level |
| Python (CPython)  | Threads + GIL (limited)            | ○        | △         | Similar GIL constraints as Ruby MRI |

# Implementation Considerations for Each Processing Model

## Event Loop Model
- **Single Thread**: True parallel processing is impossible.
- **Non-blocking IO**: Synchronous processing blocks the event loop.
- **Avoid CPU-bound**: Delegate heavy computations to Worker Threads or external processes.

## Thread Model
- **Race Conditions**: Need to control simultaneous access to shared memory.
- **Deadlocks**: Be cautious of the order of acquiring multiple locks.
- **Memory Overhead**: Approximately 1-8MB stack space per thread.

## Process Model
- **Inter-process Communication**: Collaboration via IPC or file system.
- **Memory Usage**: Each process has its own independent memory space.
- **Startup Cost**: Creating processes is heavier than creating threads.

# Conclusion
Each processing model has its pros and cons, and it is important to choose the appropriate model based on the characteristics and requirements of the application. For applications with many IO-bound processes, the event loop model is suitable, while the thread model or process model is better for CPU-bound processes. Additionally, considering the characteristics of each language is advisable.