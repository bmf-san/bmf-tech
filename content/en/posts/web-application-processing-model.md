---
title: About Processing Models of Web Applications
description: "Architect web applications with Event Loop, Thread, and Process models for IO-bound and CPU-bound concurrent workloads."
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

In this article, I will write about the typical processing models of web applications.

- Event Loop
- Thread
- Process

# Overview of Each Processing Model

## Event Loop Model (Asynchronous, Single-threaded)

- **Characteristics**: Processes events (such as asynchronous IO) sequentially on a single thread
- **Examples**: Node.js, Deno, JavaScript (Browser)

```
[ イベントキュー ] → [ イベントループ ] → [ 実行 ]
```

**Advantages**
- Good memory efficiency (does not create a large number of threads or processes)
- Can utilize IO waiting for other processing

**Disadvantages**
- Prone to blocking with heavy computations (CPU-bound)

## Thread Model (Multi-threaded)

- **Characteristics**: Multiple threads operate concurrently within a single process
- **Examples**: Java, Python, Go (internally with a thread scheduler)

```
[ プロセス ]
 ├─ スレッド1（リクエストA）
 ├─ スレッド2（リクエストB）
 └─ ...
```

**Advantages**
- Strong with CPU-bound processing (can utilize multi-core)

**Disadvantages**
- Requires management of resource consumption and race conditions for each thread

## Process Model (Multi-process)

- **Characteristics**: Each request is processed in a separate process, completely isolated
- **Examples**: PHP (FPM), Ruby (multiple Puma processes), Sidekiq (multi-process configuration)

```
[ リクエストA ] → [ プロセス1 ]
[ リクエストB ] → [ プロセス2 ]
```

**Advantages**
- High stability (less affected by other processes)

**Disadvantages**
- Poor memory efficiency (each process has its own environment)

# Suitability of Processing Models by Language

| Language         | Processing Model                  | IO-bound | CPU-bound | Remarks                                   |
| ---------------- | --------------------------------- | -------- | --------- | ---------------------------------------- |
| Node.js          | Event Loop (Asynchronous)         | ◎        | △         | IO-focused, caution with heavy computation|
| Go               | goroutine + M:N threads           | ◎        | ◎         | Good at parallel and concurrent processing with lightweight threads |
| Ruby (MRI)       | Thread + GIL (limited)            | ○        | △         | CPU processing is essentially single-threaded due to GIL |
| PHP (FPM)        | Process (1 request = 1 process)   | ○        | △         | Stable with process separation but high overhead |
| Java             | Multi-threaded                    | ◎        | ◎         | Thread optimization at the JVM level     |
| Python (CPython) | Thread + GIL (limited)            | ○        | △         | Similar GIL constraints as Ruby MRI      |

# Implementation Considerations for Each Processing Model

## Event Loop Model
- **Single-threaded**: True parallel processing is not possible
- **Non-blocking IO**: Synchronous processing blocks the event loop
- **Avoid CPU-bound**: Delegate heavy computation to Worker Threads or external processes

## Thread Model
- **Race Conditions**: Requires control of simultaneous access to shared memory
- **Deadlock**: Be cautious of the order of acquiring multiple locks
- **Memory Overhead**: Approximately 1-8MB stack area per thread

## Process Model
- **Inter-process Communication**: Coordination via IPC or file system
- **Memory Usage**: Each process has an independent memory space
- **Startup Cost**: Creating a process is heavier than creating a thread

# Conclusion
Each processing model has its pros and cons, and it is important to choose the appropriate model according to the characteristics and requirements of the application. If there are many IO-bound processes, the event loop model is suitable, while for CPU-bound processes, the thread model or process model is considered appropriate. It is also good to consider the characteristics of each language.