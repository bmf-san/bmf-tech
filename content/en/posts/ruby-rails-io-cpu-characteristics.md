---
title: Understanding IO and CPU Characteristics in Ruby and Rails
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
This post organizes the concepts of Ruby's concurrency model, the role of the GVL, the thread and process design of the Puma server, the understanding of IO/CPU-bound workloads, bottleneck identification through measurement methods, and the background of changes in Rails/Puma default settings, providing appropriate tuning strategies.

# Ruby's Concurrency Model and Global VM Lock (GVL)
## The Purpose of GVL
Ruby (MRI/CRuby) has a Global VM Lock (GVL) that restricts the execution of Ruby code to one thread at a time within the same process. The GVL exists to maintain the consistency of memory management, object management, and garbage collection (GC) within the Ruby VM, which is implemented in C. For example, heap operations during object allocation and deallocation, object traversal in mark-and-sweep GC, method cache updates, and internal table operations are not thread-safe, and the GVL prevents simultaneous execution to avoid crashes and data corruption.

Many C extensions (native extensions) are also designed to operate under the GVL, and removing the GVL requires the extension to ensure thread safety, making the cost of maintaining VM-wide consistency very high. While application-level thread safety must be ensured by developers using Mutex, the GVL functions as a global lock to ensure VM internal consistency.

## Impact of GVL on Performance
Under the GVL, CPU-bound Ruby code cannot be executed in parallel by multiple threads within the same process, and only one thread can be executed at a time. However, when IO waits occur, such as during database access or external API calls, the GVL is released, allowing other threads to continue execution, making thread concurrency effective in mixed IO-bound workloads. However, when overhead from GVL contention, thread switching delays, and GC pauses are involved, what appears to be an IO wait may actually include waits due to CPU starvation.

## Comparison with Other Implementations
TruffleRuby and JRuby do not have a GVL, but they depend on memory management and thread management methods by the VM or JVM. Simply modifying MRI to remove the GVL is vast and challenging, and Rails users can handle most web workloads sufficiently with a multi-process and moderate thread concurrency operational model.

# Concurrency with Puma and Trends in Default Settings
## Puma's Architecture
Puma is widely used as the standard server for Rails. The master process generates multiple worker processes using `fork`, and each process uses a thread pool to handle requests. While there are situations where the GVL is released during IO waits and thread switching is effective, process parallelism demonstrates parallel performance in CPU-bound parts.

Below is an excerpt from the `config/puma.rb` generated for a new Rails app. By default, the number of threads is set by the environment variable `RAILS_MAX_THREADS`, and the number of workers is controlled by `WEB_CONCURRENCY`.

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

The benefits and drawbacks of the fork model (process parallelism) and the thread model (thread concurrency) can be summarized as follows:

|       Model        |                                               Benefits                                               |                                          Drawbacks                                           |
| ------------------- | ---------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| Process Parallelism (fork) | - Avoids GVL restrictions, demonstrates true parallel performance in CPU-bound processing<br>- Independent memory space isolates crash impact | - Memory usage tends to increase<br>- Process startup cost is incurred                                  |
| Thread Concurrency        | - Small memory overhead, lightweight<br>- Other threads can operate easily during IO waits                          | - CPU-bound parallel performance is limited by GVL<br>- Waits may occur due to thread contention or GVL contention |

Based on the above, it is important to adjust the combination of the number of processes (`workers`) and the number of threads (`threads`) according to the workload characteristics of the app and infrastructure resources.

## Background of Default Thread Count Change
As discussed in [GitHub Issue #50450](https://github.com/rails/rails/issues/50450), the default Puma thread count for new Rails app generation was changed from the previous 5 to 3. In the issue, DHH proposed "1 thread per worker contributes to low latency" based on his operational experience, and many developers shared their app benchmark results and considerations using Amdahl's Law. The main points of consideration were the trade-off between latency and throughput, the optimal thread count for different IO/CPU characteristics, and ensuring a safety margin under resource constraints such as Heroku Dyno or container environments. As a result, it was agreed that around 3 threads are a balanced reasonable value for many apps, and the default was lowered from 5 to 3 in Rails 7.2. Existing apps are not affected if they explicitly set `RAILS_MAX_THREADS` or `WEB_CONCURRENCY`, and it is recommended to start with 3 threads for new projects and adjust as needed based on monitoring and benchmark results.

# Misinterpretation of IO-bound vs CPU-bound and Measurement Methods
## Pitfalls of Apparent IO Waits
The time recorded as "Query took: XX ms" in Rails logs or APM measurements may include thread scheduling waits, GVL waits, GC execution time, etc., in addition to the actual DB response time. Misinterpreting this as "DB wait is dominant" may lead to excessively increasing the number of threads, worsening GVL contention, and potentially degrading performance.

## Measuring GC Time
From Ruby 3.x onwards, `GC.total_time` is provided as a cumulative counter in nanoseconds, allowing you to understand the time spent on GC by calculating the difference before and after a specific block. From Rails 7.2 onwards, GC time is included in request logs via ActiveSupport::Notifications, making it possible to visualize the impact of GC load.

## Visualizing GVL Wait Time
Using the GVL Instrumentation API and dedicated gems (e.g., gvltools) from Ruby 3.2 onwards, there is a method to separately measure IO parts and GVL wait time. This allows you to specifically understand the increase in GVL waits under high CPU load in the background, reducing misinterpretations.

## Key Points of OS Scheduler Waits
OS-level scheduler wait times may also be included in IO measurements, but accurate measurement for each individual IO is difficult. By utilizing Linux's `/proc/<pid>/schedstat`, etc., and monitoring the runqueue wait status of the container or host as a whole, it serves as a guideline for determining the adequacy of the number of processes or threads.

## Importance of Profiling
By measuring the various metrics mentioned above, you can understand the IO/CPU ratio and the reality of GVL waits in your application, and determine the number of threads or processes from an Amdahl's Law perspective. Rather than just following defaults, it is important to profile and optimize based on your workload characteristics (frequency of external API calls, DB access patterns, rendering load, etc.).

# Background Jobs and Concurrency Settings
In job processing with Sidekiq, etc., where IO-intensive processing (external API calls, file operations, email sending, etc.) is common, a higher concurrency setting (e.g., concurrency: 10-25) is sometimes adopted. However, the following points need attention:

- **Sidekiq Concurrency Setting Example**
  - Can be set in `sidekiq.yml`:
    ```yaml
    :concurrency: 15
    ```
  - To override with environment variables:

    ```bash
    export SIDEKIQ_CONCURRENCY=15
    bundle exec sidekiq
    ```
  - Increasing concurrency makes it easier for other threads to operate during IO waits, theoretically improving throughput, but there are also side effects from GVL contention and increased GC load.
- **Case Study of Measuring GVL Impact (Pseudo Example)**
  - Objective: Understand GVL wait time and thread stalling when multiple concurrent tasks occur within a job.
  - Example Procedure:
    1. Prepare a test job combining IO parts (simulated with sleep or external calls) and CPU parts (computation load).
       ```ruby
       class BenchmarkJob
         include Sidekiq::Job
         def perform
           start = Process.clock_gettime(Process::CLOCK_MONOTONIC)
           # Simulated IO: sleep or small HTTP requests
           sleep 0.02
           # Simulated CPU: computation load
           (1..200_000).each { |i| i*i }
           duration = Process.clock_gettime(Process::CLOCK_MONOTONIC) - start
           logger.info("Job duration: #{(duration*1000).round(1)}ms")
         end
       end
       ```
    2. Introduce a GVL measurement tool (e.g., gvltools) and measure GVL wait time during job execution.
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
    3. Change concurrency and submit multiple jobs simultaneously, comparing IO time, CPU time, and GVL wait time in the logs.
       - For example, with concurrency settings of 5, 10, 20, execute 10-50 parallel jobs each and observe how much GVL waits increase.
       - Identify the point where GVL wait time sharply increases and understand the safe concurrency limit in actual operation.

- **Setting Monitoring Metrics**
  - Monitor job processing time, throughput, and queue length with Sidekiq dashboard or Prometheus.
  - Collect metrics such as Ruby process GC time, memory usage, CPU usage, and runqueue waits to visualize the impact of concurrency changes.

- **Benchmarking and Tuning Procedure**
  1. Profiling existing jobs: Understand the IO/CPU ratio in job processing time with a workload close to actual operation.
  2. Calculate concurrency candidates from an Amdahl's Law perspective: If the IO ratio is high, prioritize thread concurrency; if the CPU ratio is high, consider process division or increasing the number of workers.
  3. Conduct actual benchmarks: Perform load tests with different concurrency settings and compare processing time, GVL waits, GC, and CPU usage.
  4. Reflect in the operational environment: Gradually apply the optimal concurrency determined from test results in staging or production, confirming stability and performance.

By doing so, it is possible to understand the impact of GVL even in background jobs like Sidekiq and derive the optimal concurrency settings.

# Efforts to Improve Ruby Execution Performance
## Benefits of JIT (YJIT, etc.)
There are many cases of latency improvement with YJIT introduction, and even with the assumption of many IO waits, improvements of about 15-30% are seen in many apps, so the cost of executing Ruby code cannot be ignored.

## Consideration of Removing GVL
While there is discussion about removing the GVL, completely eliminating the GVL in MRI Ruby involves extensive and risky work, including changes to C extensions and the VM internals. Learning from the GIL removal cases of TruffleRuby/JRuby and Python, many web workloads can be sufficiently handled with multi-process and moderate thread concurrency under the GVL.

# Operational and Tuning Guidelines
- For new Rails apps, start with the default 3 Puma threads and change based on monitoring results. Set the number of processes (WEB_CONCURRENCY) considering the number of CPU cores and infrastructure environment (container/Heroku Dyno, etc.).
- Under production-equivalent load, aggregate DB time, GC time, GVL waits (if possible), external API call times, etc., included in request logs, and evaluate the impact of IO/CPU ratio and thread/process configuration.
- Measure latency and throughput with multiple thread configurations (e.g., 1-5) and multiple process configurations, and find the optimal point from an Amdahl's Law perspective. Adjust the ratio of thread concurrency for IO focus and process parallelism for CPU focus.
- Consider GC settings (RUBY_GC_HEAP_*, pausetime, etc.) according to the Ruby version to minimize GC pauses. Visualize GC time indicators from logs and adjust GC parameters as needed.
- Monitor host/container CPU usage, runqueue status, memory usage, IO wait indicators, etc., and confirm that the process/thread number settings align with host resources.
- Regularly profile and review settings according to app growth and traffic characteristic changes. Pay attention to the evolution of new Ruby/Rails versions, JIT, and ecosystem updates.

# Conclusion
Optimizing Ruby/Rails performance requires a multifaceted understanding of GVL, threads, processes, IO/CPU-bound characteristics, GC, OS scheduler waits, etc. By continuously measuring and understanding the reality and performing appropriate tuning, you can build a system that flexibly meets latency and throughput requirements.

# References
- [Rails Scaling (1) Understanding the Impact of Puma, Concurrency, and GVL on Performance (TechRacho Translation)](https://techracho.bpsinc.jp/hachi8833/2025_06_09/151182)
- [The Mythical IO-Bound Rails App (byroot's article)](https://byroot.github.io/ruby/performance/2025/01/23/the-mythical-io-bound-rails-app.html)
- [Instrumenting Thread Stalling in Ruby Applications (byroot's article)](https://byroot.github.io/ruby/performance/2025/01/23/io-instrumentation.html)
- [So You Want To Remove The GVL? (TechRacho Translation)](https://techracho.bpsinc.jp/hachi8833/2025_03_03/148712)
- [GitHub Issue: Set a new default for the Puma thread count (rails/rails#50450)](https://github.com/rails/rails/issues/50450)
