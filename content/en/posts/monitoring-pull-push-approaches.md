---
title: Differences Between Pull and Push Approaches in Monitoring Systems
slug: monitoring-pull-push-approaches
date: 2023-12-13T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Monitoring
description: An overview of the differences between pull and push approaches in monitoring systems.
translation_key: monitoring-pull-push-approaches
---


# Overview
This post summarizes the differences between pull and push approaches in monitoring systems.

# Pull Approach
In the pull approach, the monitoring server is configured with the target hosts, and it retrieves data from the monitoring hosts.
Examples include Prometheus, Nagios, and Zabbix.

Prometheus has a mechanism called an exporter, which acts like a push approach to compensate for the drawbacks of the pull approach. The exporter is installed on the target host, and the monitoring server retrieves data from it.

# Push Approach
In the push approach, an agent is installed on the target host. The agent sends data to the monitoring server.
Examples include Datadog and Mackerel.

# Differences Between Pull and Push Approaches
The following table compares the two approaches from various perspectives, but it may vary depending on the service.

|       Aspect       |                      Pull Approach                      |                               Push Approach                               |
| ---------------- | ------------------------------------------------ | ------------------------------------------------------------------ |
| Installation Cost       | Requires configuration for each target host, which is labor-intensive | Easy as it only requires installing an agent on the target host               |
| Management Cost       | Target hosts can be tracked by the monitoring server       | The monitoring server does not track target hosts                       |
| Data Retrieval Control | Adjusted by the monitoring server                             | Adjusted by the target host                                             |
| Resource Efficiency | The monitoring server can retrieve data as needed, making adjustments easier | If data retrieval from the monitoring server is frequent, the target host consumes more resources |
| Real-time Capability   | Based on the timing of data requests from the monitoring server       | Can send data in real-time                                     |
| Server Load     | Easier to centrally adjust on the monitoring server             | Adjustments are made on the target host, increasing management costs as targets increase |
| Communication Cost | If there is no data to retrieve from the target host, it results in waste | Efficient as the monitoring server only receives data pushed from the target host |
| Error Handling       | Easier to notice abnormalities in target hosts               | Difficult to determine if a target host is excluded or if an abnormality occurred     |

The table above shows general advantages and disadvantages, which may vary depending on the actual situation.

When selecting a monitoring system, it is necessary to choose between the pull or push approach based on requirements and system characteristics.

# References
- [Considerations on Push and Pull Approaches in Monitoring Systems and Prometheus](https://yasuharu519.hatenablog.com/entry/2017/12/16/215855)
- [Is the Monitoring Tool Pull or Push?](https://scrapbox.io/gosyujin/%E7%9B%A3%E8%A6%96%E3%83%84%E3%83%BC%E3%83%AB%E3%81%AFPull%E5%9E%8B%E3%81%8BPush%E5%9E%8B%E3%81%8B)
- [What Are the Types of Server Operation Monitoring? Including Liveness and Log Monitoring [Tools]](https://applis.io/posts/what-is-server-operation-monitoring)
- [Application Monitoring with Prometheus and Metricat](https://engineering.linecorp.com/ja/blog/monitoring-applications-prometheus-metricat)
