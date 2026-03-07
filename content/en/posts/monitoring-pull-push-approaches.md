---
title: Differences Between Pull and Push Approaches in Monitoring Systems
slug: monitoring-pull-push-approaches
date: 2023-12-13T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Monitoring
translation_key: monitoring-pull-push-approaches
---

# Overview
This post summarizes the differences between Pull and Push approaches in monitoring systems.

# Pull Approach
In this approach, the monitoring server is configured for the target host, and it retrieves data from the monitoring host.
ex. Prometheus, Nagios, Zabbix, etc.

Prometheus has a mechanism similar to a Push approach called an exporter, which compensates for the drawbacks of the Pull model. The exporter is installed on the target host, and the monitoring server retrieves data from the exporter.

# Push Approach
An agent is installed on the target host. The agent installed on the host sends data to the monitoring server.
ex. Datadog, Mackerel, etc.

# Differences Between Pull and Push Approaches
I compared both approaches from various perspectives, but it may vary depending on the service, so this is not always the case.

|       Perspective       |                      Pull Approach                      |                               Push Approach                               |
| ----------------------- | ------------------------------------------------------ | ------------------------------------------------------------------------- |
| Setup Cost              | Requires configuration for each target host, which is cumbersome | Easy as it only requires installing an agent on the target host           |
| Management Cost         | The monitoring server can grasp the target hosts       | The monitoring server does not grasp the target hosts                     |
| Data Retrieval Control   | Adjusted on the monitoring server side                 | Adjusted on the target host side                                         |
| Resource Efficiency      | The monitoring server can retrieve data when needed, making it easier to adjust | If data retrieval from the target server is frequent, it consumes more resources on the target host |
| Real-time Capability     | Based on the timing of data requests from the monitoring server | Can send data in real-time                                               |
| Server Load             | Easier to centrally adjust on the monitoring server side | Becomes more costly to manage as the number of targets increases due to adjustments on the target host side |
| Communication Cost      | Can waste resources if there is no data to retrieve from the target host | Efficient communication as the monitoring server only receives what is pushed from the target host |
| Error Handling          | Easier to notice abnormalities in the target host      | Harder to determine if the target host is down or if an anomaly has occurred |

The table above shows the general advantages and disadvantages, which may vary depending on the actual situation.

When choosing a monitoring system, it is necessary to select either Pull or Push based on the requirements and characteristics of the system.

# References
- [Considerations on Push and Pull Approaches in Monitoring Systems and Prometheus](https://yasuharu519.hatenablog.com/entry/2017/12/16/215855)
- [Is the Monitoring Tool Pull or Push?](https://scrapbox.io/gosyujin/%E7%9B%A3%E8%A6%96%E3%83%84%E3%83%BC%E3%83%AB%E3%81%AFPull%E5%9E%8B%E3%81%8BPush%E5%9E%8B%E3%81%8B)
- [What Are the Types of Server Operation Monitoring? Alive Monitoring and Log Monitoring, etc. [Tools]](https://applis.io/posts/what-is-server-operation-monitoring)
- [Application Monitoring with Prometheus and Metricat](https://engineering.linecorp.com/ja/blog/monitoring-applications-prometheus-metricat)