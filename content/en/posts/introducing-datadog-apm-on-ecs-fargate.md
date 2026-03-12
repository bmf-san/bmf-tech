---
title: Implementing Datadog APM in ECS on Fargate Environment
description: A step-by-step guide on Implementing Datadog APM in ECS on Fargate Environment, with practical examples and configuration tips.
slug: introducing-datadog-apm-on-ecs-fargate
date: 2021-10-19T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Amazon Web Service
  - Datadog
  - ECS
  - Fargate
translation_key: introducing-datadog-apm-on-ecs-fargate
---



# Overview
A casual note on implementing Datadog APM in an ECS on Fargate environment.

# Adjusting PHP Container Image
We are using a custom image based on the php-fpm image. Since the datadog-php-tracer is required, it is incorporated into the image as follows:
```yaml
ENV DDTRACE_VERSION=0.65.1

RUN curl -Lo datadog-php-tracer.apk https://github.com/DataDog/dd-trace-php/releases/download/${DDTRACE_VERSION}/datadog-php-tracer_${DDTRACE_VERSION}_noarch.apk \
    && apk add datadog-php-tracer.apk --allow-untrusted \
    && rm datadog-php-tracer.apk
```

In the fpm settings, configure it to read environment variables.
```
// fpm.www.conf
clear_env = yes

; Datadog APM
env[DD_AGENT_HOST] = $DD_AGENT_HOST
env[DD_SERVICE] = $DD_SERVICE
env[DD_VERSION] = $DD_VERSION
env[DD_ENV] = $DD_ENV
env[DD_TRACE_PHP_BIN] = $DD_TRACE_PHP_BIN
```

cf. https://docs.datadoghq.com/agent/amazon_ecs/apm/?tab=ec2metadataendpoint#php-fpm

# ECS Task Definition
Set up the necessary environment for both PHP and Datadog.

```json
[
  {
    // PHP
    "environment": [
      {
        "name": "DD_AGENT_HOST",
        "value": "${DATADOG_AGENT_HOST}" // Since Datadog is a sidecar, localhost is fine here
      },
      {
        "name": "DD_SERVICE",
        "value": "${DATADOG_SERVICE}"
      },
      {
        "name": "DD_VERSION",
        "value": "${DATADOG_VERSION}"
      },
      {
        "name": "DD_ENV",
        "value": "${DATADOG_ENV}"
      },
      {
        "name": "DD_TRACE_PHP_BIN",
        "value": "${DATADOG_TRACE_PHP_BIN}"
      }
    ]
  },
  {
    // Datadog
    "environment": [
      {
        "name": "ECS_FARGATE",
        "value": "true"
      },
      {
        "name": "DD_APM_ENABLED",
        "value": "true"
      }
    ]
  }
```

# Conclusion
With this, APM should be operational. This is probably the minimum configuration required. If it doesn't work well, it's a good idea to check the datadog section in `phpinfo()`.
