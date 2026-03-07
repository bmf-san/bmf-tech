---
title: IO and CPU Characteristics of Ruby and Rails
slug: ruby-rails-io-cpu-characteristics
date: 2025-06-14T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Ruby
  - Ruby on Rails
translation_key: ruby-rails-io-cpu-characteristics
---

# Overview
This post organizes the concurrency model of Ruby, the role of the GVL, the thread/process design of the Puma server, how to understand IO/CPU bounds, methods for measuring bottlenecks, and the background of changes in default settings for Rails/Puma, while presenting appropriate tuning policies.

# Ruby's Concurrency Model and Global VM Lock (GVL)
## Significance of GVL
Ruby (MRI/CRuby) has a Global VM Lock (GVL) that limits the execution of Ruby code within the same process to a single thread at a time. The GVL exists due to the Ruby VM being implemented in C, and it maintains the consistency of internal memory management, object management, and garbage collection (GC). For example, heap operations during object allocation and deallocation, object traversal in mark-and-sweep type GC, method cache updates, and internal table operations are not thread-safe. The GVL prevents concurrent execution to avoid crashes and data corruption.

Many C extensions (native extensions) are designed with the assumption that they operate under the GVL, and removing the GVL requires ensuring thread safety on the extension side, resulting in a very high cost for maintaining overall VM consistency. While application-level thread safety must be ensured by developers using Mutex, the GVL functions as a global lock to ensure internal consistency of the VM.

## Impact of GVL on Performance
Under the GVL, CPU-bound Ruby code cannot be executed in parallel by multiple threads within the same process; only one thread can execute at a time. However, when I/O waits occur, such as during database access or external API calls, the GVL is released, allowing other threads to continue execution. Therefore, thread concurrency becomes effective in mixed I/O-bound workloads. However, when overhead from GVL contention, thread switching delays, and pauses during GC execution come into play, there can be cases where what appears to be I/O waiting actually includes waiting due to CPU starvation.

## Comparison with Other Implementations
Implementations like TruffleRuby and JRuby do not have a GVL but rely on memory management and thread management methods provided by their respective VMs or JVMs. Simply modifying MRI to remove the GVL is a massive and challenging task, and Rails users can handle most web workloads sufficiently with a multi-process and moderate thread concurrency operational model.

# Concurrency Processing with Puma and Trends in Default Settings
## Puma Architecture
Puma is widely used as the standard server for Rails. The master process generates multiple worker processes using `fork`, and each process uses a thread pool to handle requests. While there are scenarios where the GVL is released during I/O waits, allowing thread switching to be effective, CPU-bound parts exhibit parallel performance through process parallelism.

Below is an excerpt from the relevant part of `config/puma.rb` generated in a new Rails app. By default, the number of threads is set by the environment variable `RAILS_MAX_THREADS`, and the number of workers is controlled by `WEB_CONCURRENCY`.

```ruby
# config/puma.rb
threads_count = ENV.fetch("RAILS_MAX_THREADS") { 3 }.to_i
threads threads_count, threads_count

workers ENV.fetch("WEB_CONCURRENCY") { 2 }

preload_app!

on_worker_boot do
  ActiveRecord::Base.establish_connection if defined?(ActiveRecord)
end
```

The advantages and disadvantages of the fork model (process parallelism) and thread model (thread concurrency) can be summarized as follows:

|       Model        |                                               Advantages                                               |                                          Disadvantages                                           |
| ------------------- | ---------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| Process Parallel (fork) | - Avoids GVL limitations and exhibits true parallel performance for CPU-bound processing<br>- Independent memory space isolates crash impacts | - Memory usage tends to increase<br>- Process startup costs are incurred                                  |
| Thread Concurrency        | - Smaller memory overhead and lightweight<br>- Other threads can operate during I/O waits                          | - CPU-bound parallel performance is limited due to GVL<br>- Thread contention and GVL contention can lead to waiting |

Given the above, it is important to adjust the combination of the number of processes (`workers`) and the number of threads (`threads`) according to the application's workload characteristics and infrastructure resources. It is crucial to adjust the combination of the number of processes (`workers`) and the number of threads (`threads`) according to the application's workload characteristics and infrastructure resources.

## Background of Changing Default Thread Count
As discussed in [GitHub Issue #50450](https://github.com/rails/rails/issues/50450), the default number of Puma threads when generating new Rails applications was changed from the previous 5 to 3. In the issue, DHH suggested based on his operational experience that "1 thread per worker contributes to low latency," and many developers shared their application's benchmark results and considerations using Amdahl's Law. The main points of consideration were the trade-off between latency and throughput, the optimal number of threads based on I/O/CPU characteristics, and ensuring a safety margin under resource constraints such as Heroku Dyno or container environments. As a result, it was agreed that around 3 threads would be a balanced and reasonable value for many applications, and the default was lowered from 5 to 3 in Rails 7.2. Existing applications that explicitly set `RAILS_MAX_THREADS` or `WEB_CONCURRENCY` are not affected, and new projects are recommended to start with 3 threads and adjust as needed based on monitoring and benchmark results.

# Misconceptions of IO-bound vs CPU-bound and Measurement Methods
## Pitfalls of Apparent I/O Waiting
The time recorded in Rails logs or APM measurements as "Query took: XX ms" may include not only the actual DB response time but also thread scheduling waits, GVL waits, GC execution time, and more. Misinterpreting this as "DB waiting is dominant" can lead to excessively increasing the number of threads, worsening GVL contention, and potentially degrading performance.

## Measuring GC Time
Since Ruby 3.x, `GC.total_time` is provided as a cumulative counter in nanoseconds, allowing you to understand the time spent on GC by calculating the difference before and after a specific block. In Rails 7.2 and later, GC time is included in request logs via ActiveSupport::Notifications, visualizing the impact of GC load.

## Visualizing GVL Wait Time
Starting from Ruby 3.2, there is a method to separate and measure I/O parts and GVL wait times using the GVL Instrumentation API and dedicated gems (e.g., gvltools). This allows for a concrete understanding of increased GVL waits under high CPU load situations in the background, reducing misconceptions.

## Key Points of OS Scheduler Waiting
OS-level scheduler wait times may also be included in I/O measurements, but accurate measurement for individual I/Os is challenging. Utilizing Linux's `/proc/<pid>/schedstat` and monitoring the runqueue wait status of the entire container or host can serve as a guideline for determining whether the number of processes or threads is adequate.

## Importance of Profiling
By utilizing various measurements mentioned above, you can understand the I/O/CPU ratio of the application and the reality of GVL waits, determining the number of threads and processes from an Amdahl's Law perspective. It is important not only to follow defaults but also to profile your workload characteristics (frequency of external API calls, DB access patterns, rendering load, etc.) for optimization.

# Background Jobs and Concurrency Settings
In job processing with Sidekiq and others, I/O-intensive processes (external API calls, file operations, email sending, etc.) are common, leading to cases where higher concurrency settings (e.g., concurrency: 10-25) are adopted. However, the following points should be noted.

- **Example of Sidekiq Concurrency Settings**
  - Configurable in `sidekiq.yml`:
    ```yaml
    :concurrency: 15
    ```
  - Overriding with environment variables:

    ```bash
    export SIDEKIQ_CONCURRENCY=15
    bundle exec sidekiq
    ```
  - Increasing concurrency allows other threads to operate during I/O waits, theoretically improving throughput, but there are also side effects from GVL contention and increased GC load.
- **Case Study on Measuring GVL Impact (Hypothetical Example)**
  - Objective: Understand GVL wait times and thread stalling when multiple concurrent tasks occur within a job.
  - Example Procedure:
    1. Prepare a test job that combines I/O parts (simulated with sleep or external calls) and CPU parts (computational processing).
       ```ruby
       class BenchmarkJob
         include Sidekiq::Job
         def perform
           start = Process.clock_gettime(Process::CLOCK_MONOTONIC)
           # I/O simulation: sleep or small HTTP requests
           sleep 0.02
           # CPU simulation: computational load
           (1..200_000).each { |i| i*i }
           duration = Process.clock_gettime(Process::CLOCK_MONOTONIC) - start
           logger.info("Job duration: #{(duration*1000).round(1)}ms")
         end
       end
       ```
    2. Introduce a GVL measurement tool (e.g., gvltools) and measure GVL wait times during job execution.
       ```ruby
       require 'gvltools'
       class BenchmarkJob
         include Sidekiq::Job
         def perform
           GVLTools::LocalTimer.enable
           start_io = GVLTools::LocalTimer.monotonic_time
           sleep 0.02
           io_wait = GVLTools::LocalTimer.monotonic_time - start_io

           start_cpu = GVLTools::LocalTimer.monotonic_time
           (1..200_000).each { |i| i*i }
           cpu_time = GVLTools::LocalTimer.monotonic_time - start_cpu

           gvl_wait = GVLTools::LocalTimer.gvl_wait_time
           logger.info("I/O time: #{io_wait.round(3)}s, CPU time: #{cpu_time.round(3)}s, GVL wait: #{gvl_wait.round(3)}s")
         ensure
           GVLTools::LocalTimer.disable
         end
       end
       ```
    3. Vary the concurrency while simultaneously launching multiple jobs and compare the I/O time, CPU time, and GVL wait times in the logs.
       - For example, with concurrency: 5, 10, 20 settings, execute 10-50 parallel jobs each and observe how much GVL wait increases.
       - Identify points where GVL wait times surge and understand the safe concurrency limits for actual operations.

- **Setting Monitoring Metrics**
  - Monitor job processing times, throughput, and queue lengths using the Sidekiq dashboard or Prometheus.
  - Collect metrics on GC time, memory usage, CPU usage, and runqueue waits for Ruby processes to visualize the impact of concurrency changes.

- **Benchmarking and Tuning Procedure**
  1. Profile existing jobs: Understand the I/O/CPU ratio of job processing times under workloads close to actual operations.
  2. Calculate concurrency candidates from an Amdahl's Law perspective: If the I/O ratio is high, prioritize thread concurrency; if the CPU ratio is high, consider process splitting or increasing the number of workers.
  3. Conduct empirical benchmarks: Perform load tests with different concurrency settings and compare processing times, GVL waits, GC, and CPU usage.
  4. Reflect results in the operational environment: Gradually apply the optimal concurrency based on test results in staging or production, confirming stability and performance.

This allows for understanding GVL impacts even in background jobs like Sidekiq and deriving optimal concurrency settings.

# Efforts for Ruby Execution Performance
## Benefits of JIT (YJIT, etc.)
There are numerous cases of latency improvements due to the introduction of YJIT, with many applications seeing improvements of around 15-30% even under the assumption of high I/O waits, indicating that the execution cost of Ruby code cannot be ignored.

## Consideration of GVL Removal
While there are discussions about removing the GVL, completely eliminating the GVL in MRI Ruby is a massive and risky task that includes changes to C extensions and the VM itself. Learning from the removal cases of GIL in TruffleRuby/JRuby and Python, many web workloads can be sufficiently handled under multi-process and moderate thread concurrency with the GVL.

# Operational and Tuning Guidelines
- For new Rails applications, start with the default 3 Puma threads and change based on monitoring results. Set the number of processes (WEB_CONCURRENCY) considering the number of CPU cores and the characteristics of the infrastructure environment (containers/Heroku Dynos, etc.).
- Under production-like loads, aggregate DB time, GC time, GVL waits (if possible), and external API call times included in request logs to evaluate the impact of I/O/CPU ratios and thread/process configurations.
- Measure latency and throughput with multiple thread configurations (e.g., around 1-5) and multiple process configurations to explore optimal points from an Amdahl's Law perspective. Adjust the ratio for I/O emphasis towards thread concurrency and for CPU emphasis towards process parallelism.
- Consider GC settings (RUBY_GC_HEAP_*, pausetime, etc.) according to the Ruby version to minimize pauses due to GC. Visualize GC time metrics from logs and adjust GC parameters as necessary.
- Monitor CPU utilization, runqueue status, memory usage, and I/O wait indicators for hosts/containers to ensure that process/thread settings align with host resources.
- Regularly profile and review settings according to the growth of the application and changes in traffic characteristics. Pay attention to new Ruby/Rails versions, advancements in JIT, and updates in the ecosystem.

# Conclusion
Optimizing performance in Ruby/Rails requires a multifaceted understanding of GVL, threads, processes, I/O/CPU-bound characteristics, GC, and OS scheduler waits. By continuously measuring and appropriately tuning based on actual conditions, it is possible to build a system that flexibly meets latency and throughput requirements.

# Reference Articles
- [Understanding the Impact of Puma, Concurrency, and GVL on Rails Scaling (TechRacho Translation)](https://techracho.bpsinc.jp/hachi8833/2025_06_09/151182)
- [The Mythical IO-Bound Rails App (byroot's article)](https://byroot.github.io/ruby/performance/2025/01/23/the-mythical-io-bound-rails-app.html)
- [Instrumenting Thread Stalling in Ruby Applications (byroot's article)](https://byroot.github.io/ruby/performance/2025/01/23/io-instrumentation.html)
- [So You Want To Remove The GVL? (TechRacho Translation)](https://techracho.bpsinc.jp/hachi8833/2025_03_03/148712)
- [GitHub Issue: Set a new default for the Puma thread count (rails/rails#50450)](https://github.com/rails/rails/issues/50450)