---
title: What is Load Average?
slug: what-is-load-average
date: 2025-07-07T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - System Performance
  - Load Average
description: Understanding the concept of Load Average in system performance.
translation_key: what-is-load-average
---


When checking system load, you often see the metric "Load Average" displayed by the `top` or `uptime` commands.

Load Average is a value that represents the number of processes waiting to be executed, and it is an important indicator of system congestion.

# The Nature of Load Average
Load Average is the average number of processes that were in a "runnable or running" state over a certain period. In Linux and UNIX-based OS, the averages for 1 minute, 5 minutes, and 15 minutes are displayed.

```
load average: 0.78, 1.32, 1.48

From left to right, these represent the 1-minute, 5-minute, and 15-minute averages.
```

This number indicates the number of processes that were simultaneously running or waiting to run relative to the CPU cores, indirectly representing CPU congestion.

It is important to note that Load Average is not "CPU usage." While CPU usage is expressed as a percentage (`%`), Load Average is the "average number of processes," which is fundamentally different in nature.

# Relationship with CPU Core Count
The most important aspect of evaluating Load Average is its relative relationship to the number of CPU cores.

For example, if a machine with 4 cores has a Load Average of 4.0, it means all cores are fully utilized. If it exceeds this, processes may be waiting for CPU availability, indicating potential "congestion."

Conversely, if a machine with 8 cores has a Load Average of 1.5, the average wait per core is only 0.1875, indicating ample capacity.

# Cases Where High Load Average is Not a Problem
A high Load Average does not necessarily indicate an issue. For instance, it is natural for it to temporarily increase in the following cases:

* Batch processing or scheduled jobs are running
* Many processes are waiting for IO or are in sleep mode, leaving the CPU with spare capacity
* Parallel processing tasks temporarily occupy the CPU

In such cases, if a high Load Average does not affect user experience, it should not be a concern.

# When to Really Pay Attention
If you observe the following symptoms, a high Load Average may be causing real issues:

* Application response delays or timeouts
* CPU usage is also spiking in Cloud Monitoring or similar tools
* IO Wait (disk IO waiting) persists for a long time
* The number of concurrent connections or threads is rapidly increasing

In such situations, you need to use tools like `top`, `vmstat`, or `iostat` to check details and identify which resource is the bottleneck.

# Using Load Average and Other Metrics
Load Average is a metric to see "how congested processes are." To understand the system's state, you also need to look at other metrics like the following:

| Metric         | Description                | Representative Commands |
| -------------- | -------------------------- | ------------------------ |
| Load Average   | Average number of waiting processes | `uptime`, `top`         |
| CPU Usage      | Percentage of CPU time used | `top`, `mpstat`         |
| IO Wait        | Percentage of IO wait time  | `iostat`, `vmstat`      |
| Memory Usage   | State of actual memory usage | `free`, `htop`          |
| Process Count  | Number of currently running processes | `ps`, `htop`           |

# Correct Interpretation and Engagement
Load Average is not something to be swayed by just looking at the "numbers." You should understand the appropriate value for that server and observe trends. Consider whether it exceeds the number of cores, if there are abnormal spikes, and if there is any real impact. By keeping these three points in mind, you can avoid overreacting.

Moreover, the trend of Load Average is a good indicator of how "busy" the system is. By combining it with CPU usage and memory usage, you can perform a more realistic analysis.

# Conclusion
Load Average provides valuable information for understanding system congestion as an indicator of process traffic. By being aware of its relative relationship with CPU core count, you can correctly evaluate the meaning of the values.

By observing it in combination with other metrics like CPU usage, IO Wait, and memory usage, it becomes a clue to decipher actual bottlenecks and processing wait situations.

Instead of just focusing on "high Load Average," consider "how it affects user experience and the overall system."