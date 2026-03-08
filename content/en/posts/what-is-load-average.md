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
translation_key: what-is-load-average
---

When checking system load, you often encounter the metric "Load Average" displayed by commands like `top` and `uptime`.

Load Average is a value that represents the number of processes in a runnable or running state, making it an important indicator of system congestion.

# The Nature of Load Average
Load Average is the average number of "runnable or running processes" over a certain period. In Linux and UNIX-like operating systems, the averages for 1 minute, 5 minutes, and 15 minutes are displayed.

```
load average: 0.78, 1.32, 1.48

The values from left to right represent the averages for 1 minute, 5 minutes, and 15 minutes.
```

This number indicates the number of processes that were either running or waiting to run simultaneously on CPU cores, indirectly reflecting CPU congestion.

It is important to note that Load Average is not the same as "CPU usage." While CPU usage is expressed as a percentage, Load Average is fundamentally different as it represents the "average number of processes."

# Relationship with CPU Core Count
The most important aspect of evaluating Load Average is its relative relationship with the number of CPU cores.

For example, if a 4-core machine has a Load Average of 4.0, it means all cores are constantly at full capacity. If it exceeds this number, processes may be waiting for CPU availability, indicating potential "bottlenecks."

Conversely, if an 8-core machine has a Load Average of 1.5, the average wait per core is only 0.1875, indicating ample capacity.

# Cases Where High Load Average is Not a Problem
A high Load Average does not necessarily indicate a problem. For example, it is natural for the Load Average to temporarily rise in the following cases:

* Batch processing or scheduled jobs are running
* Many processes are waiting for I/O or are in sleep mode, with ample CPU availability
* Parallel processing tasks are temporarily occupying the CPU

In such cases, if the high Load Average does not affect user experience, there is no need for concern.

# When Caution is Necessary
If symptoms such as the following are observed, a high Load Average may be causing real issues:

* Application response delays or timeouts
* Simultaneous spikes in CPU usage as seen in Cloud Monitoring
* Prolonged I/O Wait (disk I/O wait)
* Sudden increases in simultaneous connections or thread counts

In such situations, it is necessary to check details using tools like `top`, `vmstat`, and `iostat` to identify which resources are bottlenecked.

# Differentiating Load Average from Other Metrics
Load Average is primarily a metric for assessing "how congested processes are." To understand the system's state, it is also necessary to consider the following metrics:

| Metric          | Description                     | Common Commands       |
|----------------|---------------------------------|-----------------------|
| Load Average    | Average number of runnable processes | `uptime`, `top`      |
| CPU Usage       | Percentage of CPU time used     | `top`, `mpstat`      |
| I/O Wait        | Percentage of I/O wait time     | `iostat`, `vmstat`   |
| Memory Usage    | Current memory usage state      | `free`, `htop`       |
| Active Processes | Number of currently running processes | `ps`, `htop`       |

# Correct Perspective and Approach
Load Average should not be viewed in isolation; instead, one should understand the appropriate values for that server and observe trends. It is important to consider whether it exceeds the number of cores, if there are abnormal spikes, and whether there are real impacts. By focusing on these three points, one can avoid overreacting.

Additionally, the trend of Load Average is a good indicator of how busy the system is. By combining it with CPU usage and memory usage, more accurate analyses can be conducted.

# Conclusion
Load Average provides valuable information for understanding process congestion and system load. By being aware of its relative relationship with CPU core count, one can accurately evaluate its significance.

Observing Load Average alongside other metrics like CPU usage, I/O Wait, and memory usage can provide insights into actual bottlenecks and processing wait situations, rather than just looking at a number.

Instead of merely considering "high Load Average" as a problem, it is essential to think about "how this impact manifests in user experience and the overall system."