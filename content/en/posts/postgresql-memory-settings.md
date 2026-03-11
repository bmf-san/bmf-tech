---
title: PostgreSQL Memory Configuration
slug: postgresql-memory-settings
date: 2025-06-14T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - PostgreSQL
description: Guidelines for optimizing PostgreSQL memory settings for performance and stability.
translation_key: postgresql-memory-settings
---



# Overview
Proper memory configuration is essential for enhancing database performance and ensuring stable operations. Disk access is significantly slower than memory access, so improving response performance by reading and writing from memory as much as possible is desirable. However, excessive memory allocation increases the risk of OOM (Out Of Memory), potentially leading to a system-wide shutdown. Therefore, it is crucial to carefully configure PostgreSQL's memory management to ensure both stability and performance.

This article summarizes the basic structure of shared and local memory areas, key parameter configuration guidelines, and operational verification procedures, based on PostgreSQL official documentation, practical operational insights, and Gihyo articles.

# PostgreSQL Process Structure and Memory Area Division
PostgreSQL adopts a multi-process model, consisting of a master process generated at server startup, background processes responsible for tasks like WAL writing, and backend processes (session processes) created for each client connection. Each process uses memory independently, and the memory allocation of backend processes, which increase in proportion to the number of connections, significantly impacts overall memory consumption.

Memory management is broadly divided into the following two categories:

1. **Shared Memory**
   This area is allocated at server startup and shared among multiple processes. Key settings include `shared_buffers`, `wal_buffers`, Free Space Map, and Visibility Map.
2. **Process-Local Memory**
   This is working memory allocated for each backend process, used for operations like sorting, hash joins, and maintenance tasks. Parameters like `work_mem`, `maintenance_work_mem`, and `temp_buffers` fall into this category and can be dynamically set.

# Guidelines for Setting shared_buffers
`shared_buffers` is a parameter that sets the amount of shared memory PostgreSQL uses as a database cache. The default 128MB is small, so if using a dedicated server, start with about 25% of system memory and gradually increase while balancing with OS cache.

- Default & Unit: Default is 128MB. If no unit is specified, it is considered in BLCKSZ (usually 8kB) units, but it is recommended to specify explicitly, like `shared_buffers = '2GB'`.
- Recommended Range: Start with about 25% of system memory, with an upper limit of around 40%.
- Restart Requirement: Changes require a restart.
- OS Kernel Settings: Adjustments like `shmmax`, disabling Transparent Huge Pages, and NUMA optimization may be necessary.
- WAL/Checkpoint Related: Increasing `shared_buffers` may require adjusting `max_wal_size` and `checkpoint_completion_target` to mitigate write bursts and I/O spikes during checkpoints.
- Workload Characteristics: Effective for read-heavy workloads, but be cautious of I/O load and checkpoint impacts for write-heavy scenarios.

# Guidelines for Setting work_mem
`work_mem` sets the upper limit of memory available for temporary operations like sorting and hash joins. The limit applies per query execution process and per operation, so actual consumption depends on factors like `work_mem × number of temporary operations × number of parallel workers × number of concurrent sessions`. In the worst case, this can lead to significant memory consumption. However, this is a theoretical maximum, and actual usage varies based on query content and timing, so it should be considered a guideline.

- Default & Unit: Default is 4MB. Specify explicitly, like `work_mem = '16MB'`.
- Application Unit: Applies per query and per operation. For parallel queries, it also applies to each worker, increasing consumption in conjunction with `max_parallel_workers_per_gather`.
- Hash Memory Multiplier: Memory usage limit for hash operations is controlled by `hash_mem_multiplier`.
- Concurrent Connections and Risk: High concurrent connections or complex queries increase the risk of OOM.

# Other Related Parameters
- `effective_cache_size`: The amount of cache the planner assumes is available, significantly affecting index usage decisions (does not directly impact actual memory consumption).
- `maintenance_work_mem`: Working memory for VACUUM and index creation.
- `temp_buffers`: Memory area for temporary tables.
- `max_connections` and PgBouncer: Control simultaneous connections to reduce memory consumption.
- `max_parallel_workers_per_gather`: Controls the number of workers for parallel queries.
- `replication_slot_max_wal_size`: Maximum WAL size a replication slot can hold.
- `autovacuum_work_mem`: Memory used by the autovacuum process. Default is the value of `maintenance_work_mem`.
- `logical_decoding_work_mem`: Memory usage limit during logical decoding (PostgreSQL 13 and later).
- `wal_buffers`: Buffer for WAL writing. Automatically set, but adjustments can be effective under high load.
- `temp_file_limit`: Size limit of temporary files available to a session (in MB).
- `bgwriter_lru_maxpages`, `bgwriter_lru_multiplier`: Control the amount and frequency of buffer writes by the background writer.
- `shared_memory_type`: Determines how shared memory is allocated (`mmap`/`sysv`/`windows`).
- `huge_pages`: Whether to use OS Huge Pages (Transparent Huge Pages).

# Conclusion
PostgreSQL memory management requires a comprehensive design centered around `shared_buffers` and `work_mem`, considering process-specific consumption, concurrent connections, and parallel query characteristics. Implement changes gradually, and ensure thorough pre-verification, risk assessment, and continuous monitoring to achieve stable and high-performance operations.

# References
- [PostgreSQL Official Documentation: Runtime Configuration - Resource Consumption](https://www.postgresql.org/docs/current/runtime-config-resource.html)
- [Gihyo Series: Detailed PostgreSQL Part 2 - Internal Structure of PostgreSQL](https://gihyo.jp/dev/feature/01/dex_postgresql/0002)
- [Katsusand Blog: Considering PostgreSQL Memory Management](https://katsusand.dev/posts/postgresql-memory/)
