---
title: Comparison of New Relic and Datadog APM
slug: comparison-of-new-relic-and-datadog-apm
date: 2021-08-10T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - APM
  - Datadog
  - New Relic
description: A comparison of features when considering switching from New Relic to Datadog APM.
translation_key: comparison-of-new-relic-and-datadog-apm
---



# Overview
A memo on the feature comparison when considering switching from New Relic to Datadog APM.

*Note: Includes aspects beyond just APM.*

# Comparison
|                        |New Relic                                                                                                                |Datadog                                                                                                            |Comparison                                                                                                                                                                            |
|------------------------------|-------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|Server Monitoring                        |Metrics collection possible for cloud, containers, OS, middleware, network, OSS, etc. Metrics collection possible for CPU, memory, storage, network, processes, etc.                                            |Omitted as it is equivalent to New Relic                                                                                                  |There seems to be no significant difference.                                                                                                                                                                     |
|Application                      |Response time, throughput, error rate, transaction thread profile, cross-application tracing, transaction analysis, service map, external service performance monitoring, transaction metrics and traces, distributed tracing, deployment analysis|Service map, service performance dashboard, deployment tracking, application logs, RUM, Synthetic and trace integration, profiler (identifying code lines consuming CPU, memory, I/O), distributed tracing               |New Relic seems to have more features, but Datadog appears to be able to do almost the same.                                                                                                                                      |
|Database                        |Time consumption of database calls, slow query display, database and cache operation display, response time of database calls, throughput, SQL query analysis (display of poorly performing queries and their stack traces)                      |Query throughput, performance, connection numbers, and other metrics can be obtained                                                                                   |There seems to be no significant difference. Datadog requires MySQL integration. New Relic allows log drill-down with APM integration. Datadog likely can do the same with RUM integration...                                                                      |
|RUM (Realtime User Monitoring)|Page load performance, browser performance, AJAX analysis, JavaScript error analysis, can be integrated with New Relic Mobile                                                       |Metrics related to page views can be collected, RUM Explorer (access to all data collected from applications, useful for monitoring and error resolution), can be integrated with logs, APM, profiler, error tracking, iOS and Android can be supported with SDK|The frontend seems more comprehensive with New Relic. Datadog requires additional JS tracer for JS-related aspects. cf. https://docs.datadoghq.com/ja/tracing/setup_overview/setup/nodejs/?tab=%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A|
|Synthetic                     |Tests can be conducted by location (worldwide), can be specified per day, per minute, etc., availability monitoring, script browser interaction (mystery), API testing, can be integrated with APM? (probably), CI/CD integration possible                        |Browser test, private location (monitoring of internal applications, monitoring of private URLs not accessible from the internet), CI/CD integration, can be integrated with APM                                     |There seems to be no significant difference.                                                                                                                                                                    |
|Report                          |Custom dashboard, availability, capacity analysis, SLA, deployment tracking, usage analysis per host, extensibility                                                                           |Custom dashboard, predictive monitor, anomaly detection monitor, deployment tracking, SLO                                                                            |Datadog seems to focus on monitoring, while New Relic focuses on analysis.                                                                                                                                    |
|Integration Features                          |Can be integrated with services of various clouds and various middleware                                                                                                |Omitted as it is equivalent to New Relic                                                                                                  |Both have extensibility.                                                                                                                                                                   |
|APM Installation Method                       |Install an agent for each language                                                                                                       |Install Datadog Agent. Enable APM in Datadog Agent settings, configure tracer for each language                                                    |New Relic APM has information about agent CPU load, but Datadog APM is unknown. Needs verification.                                                                                                                        |
