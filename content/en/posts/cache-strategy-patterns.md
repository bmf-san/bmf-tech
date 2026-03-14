---
title: Basic Patterns of Cache Strategies
description: 'Learn the five cache strategies: Cache Aside, Read Through, Write Through, Write Back, and Write Around. Covers read/write flows, consistency trade-offs, and mermaid diagram walkthroughs.'
slug: cache-strategy-patterns
date: 2025-08-03T00:00:00Z
author: bmf-san
categories:
  - Architecture
tags:
  - Cache
  - System Performance
translation_key: cache-strategy-patterns
---

To enhance performance in web applications and distributed systems, it is essential to understand the basic usage patterns of "cache."

* Cache Aside
* Read Through
* Write Through
* Write Back
* Write Around

## Cache Aside

### Overview

This pattern involves the application explicitly managing the cache as needed.

```mermaid
graph TD
    A[Application] -->|Read Request| B[Cache]
    B -->|Hit| A
    B -->|Miss| C[Database]
    C -->|Data Retrieval| A
    A -->|Save to Cache| B

    A -->|Write| C
    A -->|Invalidate Cache| B
```

```mermaid
sequenceDiagram
    participant A as Application
    participant C as Cache
    participant D as Database

    Note over A,D: Read Process (Cache Miss)
    A->>C: Data Request
    C->>A: Cache Miss
    A->>D: Data Retrieval
    D->>A: Data Return
    A->>C: Save Data to Cache

    Note over A,D: Write Process
    A->>D: Data Update
    D->>A: Update Complete
    A->>C: Invalidate Cache
```

### Read Flow

1. Check if data exists in the cache (**Cache Hit**)
2. If not, retrieve from the DB and save to cache (**Cache Miss**)

### Write Flow

1. Update the database
2. Optionally delete or update the cache

### Features

* Cache management is handled by the application
* Suitable for data with high read frequency and low update frequency
* Cache and DB consistency is the application's responsibility

### Use Cases

* Web applications using Redis, Memcached, etc.

## Read Through

### Overview

In this pattern, the cache automatically handles retrieval from the DB during read operations. The application interacts only with the cache, and processing during cache misses is handled transparently.

```mermaid
graph TD
    A[Application] -->|Read Request| B[Cache]
    B -->|Hit| A
    B -->|Auto Retrieve on Miss| C[Database]
    C -->|Data Return| B
    B -->|Auto Save and Return| A

    A -->|Write| C
```

```mermaid
sequenceDiagram
    participant A as Application
    participant C as Cache
    participant D as Database

    Note over A,D: Read Process (Cache Miss)
    A->>C: Data Request
    C->>D: Automatically Retrieve from DB
    D->>C: Data Return
    C->>C: Auto Save to Cache
    C->>A: Data Return

    Note over A,D: Write Process
    A->>D: Data Update (Cache Bypass)
    D->>A: Update Complete
```

### Read Flow

1. The application requests data from the cache
2. If it’s a cache hit, return it directly
3. If it’s a cache miss, the cache retrieves data from the DB and automatically saves it before returning to the application

### Features

* The cache transparently handles DB access
* The application does not need to be aware of the cache's existence
* Processing during cache misses is hidden from the application
* Writes are typically done directly to the DB

### Use Cases

* ORM L2 cache features, CDN, proxy cache, Hibernate, etc.

## Write Through

### Overview

This strategy involves writing operations first to the cache and **simultaneously writing to the DB**.

```mermaid
graph TD
    A[Application] -->|Write Request| B[Cache]
    B -->|Synchronous Write| C[Database]
    C -->|Completion Notification| B
    B -->|Completion Notification| A

    A -->|Read Request| B
    B -->|Cache Hit| A
```

```mermaid
sequenceDiagram
    participant A as Application
    participant C as Cache
    participant D as Database

    Note over A,D: Write Process
    A->>C: Data Update Request
    C->>D: Synchronous Write
    D->>C: Write Complete
    C->>A: Update Complete Notification

    Note over A,D: Read Process
    A->>C: Data Request
    C->>A: Cache Hit (Data Return)
```

### Write Flow

1. Update the cache
2. Immediately reflect the same content in the DB

### Features

* Cache and DB maintain consistency
* Write latency is somewhat higher
* Reads are fast and consistent

### Use Cases

* User profiles, configuration information, master data, etc., where consistency is prioritized

## Write Back

### Overview

In this strategy, write operations are reflected only in the cache first, and **writing to the DB is processed asynchronously with a delay**.

```mermaid
graph TD
    A[Application] -->|Write Request| B[Cache]
    B -->|Immediate Completion Notification| A
    B -->|Asynchronously Batch Write| C[Database]

    A -->|Read Request| B
    B -->|Cache Hit| A

    D[Background Process] -->|Periodically| B
    D -->|Batch Write| C
```

```mermaid
sequenceDiagram
    participant A as Application
    participant C as Cache
    participant D as Database
    participant B as Background Process

    Note over A,D: Write Process (Asynchronous)
    A->>C: Data Update Request
    C->>A: Immediate Completion Notification

    Note over A,D: Read Process
    A->>C: Data Request
    C->>A: Cache Hit (Data Return)

    Note over A,D: Background Synchronization
    B->>C: Check for Dirty Data
    C->>B: Return Unsynchronized Data
    B->>D: Batch Write
    D->>B: Write Complete
```

### Write Flow

1. Write only to the cache (mark as dirty)
2. Notify the application of completion immediately
3. Later, update the DB in batches or driven by events

### Features

* Fast writes (low latency)
* Risk of data loss in case of a crash
* Suitable for high-frequency updates (consecutive updates to the same key can result in only the final outcome)
* Management of dirty data is crucial

### Use Cases

* CPU cache, logs, temporary game scores, measurement data, etc.

## Write Around

### Overview

This strategy involves writing operations **without reflecting in the cache, writing only to the DB**.

```mermaid
graph TD
    A[Application] -->|Write Request| C[Database]
    C -->|Completion Notification| A

    A -->|Read Request| B[Cache]
    B -->|Miss| C
    C -->|Data Retrieval| B
    B -->|Save to Cache and Return| A

    B -.->|Bypass Cache| X[Skip on Write]
```

```mermaid
sequenceDiagram
    participant A as Application
    participant C as Cache
    participant D as Database

    Note over A,D: Write Process (Cache Bypass)
    A->>D: Data Update (Skip Cache)
    D->>A: Update Complete

    Note over A,D: Read Process (Cache Miss)
    A->>C: Data Request
    C->>A: Cache Miss
    A->>D: Data Retrieval
    D->>A: Data Return
    A->>C: Save Data to Cache
```

### Write Flow

1. Write directly to the DB, bypassing the cache

### Read Flow

* A cache miss occurs during read → retrieve from DB and save to cache

### Features

* Writes do not pollute the cache (no unnecessary data in the cache)
* Reading immediately after writing is prone to misses (as it does not exist in the cache)

### Use Cases

* Access logs, temporary files, backup data, etc. (records with low-frequency access)

## Comparison Table of Each Pattern

| Strategy            | Overview                          | Fast Read | Fast Write  | Consistency          | Cache Management     |
| ------------------- | --------------------------------- | --------- | ----------- | --------------------- | -------------------- |
| Cache Aside         | Application explicitly manages cache | ◎         | △ (Management needed) | △ (Manual)           | Managed by application |
| Read Through        | Cache transparently retrieves from DB | ◎         | △ (Direct to DB) | △ (Only during read) | Automatic (Read)     |
| Write Through       | Simultaneous write to cache and DB  | ◎         | △ (Same wait) | ◎ (Always synchronized) | Automatic             |
| Write Back          | Write only to cache, sync later    | ◎         | ◎ (Asynchronous) | △ (Delayed sync)     | Automatic (Risky)     |
| Write Around        | Bypass cache during writes          | ◎         | ◎ (Direct to DB) | △ (Consistency during read) | Automatic             |

## Conclusion

Which strategy is best depends on the use case and trade-offs.

* **If reads are primary and consistency is important** → Write Through
* **If transparency in reading is prioritized** → Read Through
* **If write performance is the top priority and some data loss risk is acceptable** → Write Back
* **For low-frequency access and cache efficiency** → Write Around
* **If fine control is needed and development effort can be allocated** → Cache Aside

When making a selection, consider the following factors:

* **Consistency Requirements**: Is strong consistency needed?
* **Performance Requirements**: Which is prioritized, read or write?
* **Availability Requirements**: What is the tolerance for data loss risk?
* **Operational Costs**: Complexity of management and development effort

By effectively utilizing cache strategies, the performance and availability of applications can be significantly improved. It is crucial to choose the one that fits the requirements of each project.