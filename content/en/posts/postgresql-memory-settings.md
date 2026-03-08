---
title: PostgreSQL Memory Configuration
slug: postgresql-memory-settings
date: 2025-06-14T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - PostgreSQL
translation_key: postgresql-memory-settings
---

# Overview
Proper memory configuration is essential for improving database performance and ensuring stable operations. Disk access is significantly slower than memory access, and we want to enhance response performance by reading and writing from memory as much as possible. However, excessive memory allocation increases the risk of OOM (Out Of Memory) errors, which can lead to system-wide crashes. Therefore, it is necessary to carefully configure PostgreSQL's memory management settings to ensure performance while maintaining stability.

In this article, we summarize the basic structure of shared memory and local memory areas, guidelines for setting key parameters, and operational verification procedures based on the official PostgreSQL documentation, practical insights, and Gihyo articles.

# PostgreSQL Process Structure and Memory Area Division
PostgreSQL adopts a multi-process model, consisting of a master process generated at server startup, a group of background processes responsible for WAL writing, and backend processes (session processes) generated for each client connection. Each process uses its own memory, and the memory allocation of backend processes, which increases in proportion to the number of connections, significantly impacts overall memory consumption.

Memory management is broadly divided into the following two categories:

1. **Shared Memory Area**  
   This area is allocated at server startup and shared among multiple processes. Key configuration items include `shared_buffers`, `wal_buffers`, Free Space Map, and Visibility Map.
2. **Local Memory Area**  
   This is the working memory allocated for each backend process. It is used for sorting, hash joins, maintenance operations, etc., and includes `work_mem`, `maintenance_work_mem`, and `temp_buffers`, some of which can be dynamically set.

# Guidelines for Setting shared_buffers
`shared_buffers` is a parameter that sets the amount of shared memory PostgreSQL uses as a database cache. The default of 128MB is small, so for dedicated servers, it is recommended to start with about 25% of system memory and gradually increase it while balancing with the OS cache.

- Default/Unit: Default is 128MB. When no unit is specified, it is considered in BLCKSZ (usually 8kB), but in practical use, it is recommended to specify it explicitly with a unit, such as `shared_buffers = '2GB'`.
- Recommended Setting Range: Start with about 25% of system memory, with an upper limit of around 40%.
- Restart Requirement: A restart is required for configuration changes.
- OS Kernel Settings: It may be necessary to adjust `shmmax`, disable Transparent Huge Pages, and optimize for NUMA.
- WAL/Checkpoint Related: As `shared_buffers` increases, adjustments to `max_wal_size` and `checkpoint_completion_target` are necessary. This can mitigate write bursts and I/O spikes during checkpoints.
- Workload Characteristics: It is effective for read-heavy workloads, but caution is needed for write-heavy workloads due to I/O load and checkpoint impacts.

# Guidelines for Setting work_mem
`work_mem` sets the maximum memory available for temporary operations such as sorting and hash joins. The limit applies to each query execution process and each operation, so the actual consumption depends on factors like `work_mem × number of temporary operations × number of parallel workers × number of concurrent sessions`. In the worst case, this can lead to significant memory consumption. However, this is a theoretical maximum and can vary based on the actual query content and timing, so it should be viewed as a guideline.

- Default/Unit: Default is 4MB. Specify explicitly, such as `work_mem = '16MB'`.
- Application Unit: Per query and per operation. When using parallel queries, it applies to each worker, increasing consumption in conjunction with `max_parallel_workers_per_gather`.
- Hash Memory Multiplier: The memory usage limit for hash operations is controlled by `hash_mem_multiplier`.
- Concurrent Connections and Risks: A large number of concurrent connections or complex queries can pose a risk of OOM.

# Other Related Parameters
- `effective_cache_size`: The amount of cache that the planner assumes is available (does not directly affect actual memory consumption) significantly impacts index usage decisions.
- `maintenance_work_mem`: Working memory for VACUUM and index creation.
- `temp_buffers`: Memory area for temporary tables.
- `max_connections` and PgBouncer: Controls concurrent connections to reduce memory consumption.
- `max_parallel_workers_per_gather`: Controls the number of workers for parallel queries.
- `replication_slot_max_wal_size`: Maximum WAL size that a replication slot can hold.
- `autovacuum_work_mem`: Memory used by the autovacuum process. Default is the value of `maintenance_work_mem`.
- `logical_decoding_work_mem`: Memory usage limit during logical decoding (PostgreSQL 13 and later).
- `wal_buffers`: Buffer for WAL writing. Automatically set, but adjustments can be effective under high load.
- `temp_file_limit`: Maximum size limit for temporary files available to a session (in MB).
- `bgwriter_lru_maxpages`, `bgwriter_lru_multiplier`: Control the amount and frequency of buffer writes by the background writer.
- `shared_memory_type`: Determines which method to use for allocating shared memory: `mmap`/`sysv`/`windows`.
- `huge_pages`: Whether to use the OS's Huge Pages (Transparent Huge Pages).

# Conclusion
Effective memory management in PostgreSQL requires a comprehensive design considering `shared_buffers` and `work_mem`, along with process-specific consumption, concurrent connections, and parallel query characteristics. Configuration changes should be made gradually, with thorough pre-verification, risk assessment, and continuous monitoring to achieve stable and high-performance operations.

# References
- [PostgreSQL Official Documentation: Runtime Configuration - Resource Consumption](https://www.postgresql.org/docs/current/runtime-config-resource.html)
- [Gihyo Series: Detailed Explanation of PostgreSQL Part 2 - Internal Structure of PostgreSQL](https://gihyo.jp/dev/feature/01/dex_postgresql/0002)
- [Katsusand Blog: Considering PostgreSQL Memory Management](https://katsusand.dev/posts/postgresql-memory/)