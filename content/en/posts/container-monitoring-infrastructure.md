---
title: Building a Monitoring Platform with Containers
description: A step-by-step guide on Building a Monitoring Platform with Containers, with practical examples and configuration tips.
slug: container-monitoring-infrastructure
date: 2021-12-18T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Docker Compose
  - Elasticsearch
  - Grafana
  - Kibana
  - Prometheus
  - cadvisor
  - efk
  - fluentd
  - node-exporter
translation_key: container-monitoring-infrastructure
---

This article is the 17th day entry for the [Makuake Advent Calendar 2021](https://adventar.org/calendars/6822).

It's been three years since I joined the company, and this is my third time participating in the company's advent calendar.

For the past year, I've been part of the Re-Architecture team, which develops and operates the service infrastructure of Makuake, and I've been working hard on various tasks.

I'm sure I'll be tackling various challenges next year as well.

Here's the job listing for the Re-Architecture team.
[【Go/Microservices】Engineer Recruitment for the Re-Architecture Team to Revamp the "Makuake" Infrastructure!](https://findy-code.io/companies/633/jobs/3hdLodfPABhrO)

Now, the first article of this year (I'll probably write another one on the 24th) is "Building a Monitoring Platform with Containers."

I've been experimenting with setting up a monitoring platform for an application I'm developing as a hobby, unrelated to my main job, using containers. I'd like to share the insights (though not much) I gained during that process.

# Monitoring Platform Configuration
![Screenshot 2021-12-18 4 04 19](/assets/images/posts/container-monitoring-infrastructure/146595314-d593b9d8-9faa-4275-8c94-6c12cfcbfe36.png)

The applications that make up the monitoring platform we will build this time are as follows:

- [Elasticsearch](https://www.elastic.co/jp/elasticsearch/)
  - A search engine. It accumulates application logs. The application that collects logs is a simple app implemented in Go.
- [Fluentd](https://www.fluentd.org/)
  - A log aggregator. It collects logs and forwards them to Elasticsearch.
  - Although we could have used ELK (Logstash instead of Fluentd) instead of the EFK stack, we chose Fluentd because we are more familiar with it.
- [Kibana](https://www.elastic.co/jp/kibana/)
  - A UI for data search, visualization, and analysis. It visualizes application logs.
- [Grafana](https://grafana.com/)
  - Like Kibana, a UI for data. It is used for visualizing system metrics.
  - It can also be used for visualizing application logs, but we use Kibana for application logs.
- [Prometheus](https://prometheus.io/)
  - A system metrics monitoring tool. It collects system metrics in conjunction with node-exporter and cadvisor. The collected data is visualized with Grafana.
- [node-exporter](https://github.com/prometheus/node_exporter)
  - Collects OS metrics.
- [cadvisor](https://github.com/google/cadvisor)
  - Collects container metrics.

This setup is roughly configured for those who want to build and play around with it.

These applications will be built using docker-compose.

# Building the Monitoring Platform
All implementations are available at [bmf-san/docker-based-monitoring-stack-boilerplate](https://github.com/bmf-san/docker-based-monitoring-stack-boilerplate).

Once cloned, you can create a `.env` file and run `docker-compose up` to start using it immediately.

Note that cadvisor does not start on M1, so container metrics cannot be collected. It has been confirmed to work on Intel Macs and Ubuntu.

The directory structure is as follows. I will explain each container one by one.

```sh
.
├── app
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
├── cadvisor
│   └── Dockerfile
├── docker-compose.yml
├── elasticsearch
│   └──Dockerfile
├── .env.example
├── fluentd
│   ├── Dockerfile
│   └── config
│       └── fluent.conf
├── grafana
│   ├── Dockerfile
│   └── provisioning
│       ├── dashboards
│       │   ├── containor_monitoring.json
│       │   ├── dashboard.yml
│       │   └── node-exporter-full_rev21.json
│       └── datasources
│           └── datasource.yml
├── kibana
│   ├── Dockerfile
│   └── config
│       └── kibana.yml
├── node-exporter
│   └── Dockerfile
└── prometheus
    ├── Dockerfile
    └── template
        └── prometheus.yml.template
```

I will explain each one in detail.

## app
First, we create a rough application container that logs.

```sh
├── app
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
```

The application is like this. It logs "OK" and responds with "Hello World".

```golang
package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  mux := http.NewServeMux()
  mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    log.Println("OK")
    fmt.Fprintf(w, "Hello World")
  }))
  http.ListenAndServe(":8080", mux)
}
```

The Dockerfile is simple, just building the source and executing the binary, so there's nothing to add.

docker-compose.yml looks like this:

```yaml
version: '3.9'
services:
app:
    container_name: "${APP_CONTAINER_NAME}"
    environment:
      - APP_IMAGE_NAME=${APP_IMAGE_NAME}
      - APP_IMAGE_TAG=${APP_IMAGE_TAG}
      - ALPINE_IMAGE_NAME=${APP_IMAGE_NAME}
    build:
      context: "./app"
      dockerfile: "Dockerfile"
      args:
        PLATFORM: "${PLATFORM}"
        APP_IMAGE_NAME: "${APP_IMAGE_NAME}"
        APP_IMAGE_TAG: "${APP_IMAGE_TAG}"
        ALPINE_IMAGE_NAME: "${ALPINE_IMAGE_NAME}"
    ports:
      - ${APP_HOST_PORT}:${APP_CONTAINER_PORT}
    command: ./app
    logging:
      driver: "fluentd"
      options:
        fluentd-address: ${FLUENTD_ADDRESS}
        fluentd-async-connect: "true"
        tag: "${APP_LOGGING_TAG}"
```

The logging driver is set to fluentd to forward logs to fluentd.

`fluent-async-connect` buffers logs until a connection to fluentd is established. If true, it buffers logs even if the connection is not yet established.

## fluentd
Next, let's explain the fluentd container, which is the log destination for the application.

```sh
├── fluentd
│   ├── Dockerfile
│   └── config
│       └── fluent.conf
```

The Dockerfile is as follows:

```yml
ARG FLUENTD_IMAGE_NAME=${FLUENTD_IMAGE_NAME}
ARG FLUENTD_IMAGE_TAG=${FLUENTD_IMAGE_TAG}
ARG PLATFORM=${PLATFORM}

FROM --platform=${PLATFORM} ${FLUENTD_IMAGE_NAME}:${FLUENTD_IMAGE_TAG}

USER root

RUN gem install fluent-plugin-elasticsearch

USER fluent
```

The only gem used in fluentd is fluent-plugin-elasticsearch for integration with elasticsearch.

The USER is set to root and then switched back to fluent because the execution user of the fluentd image is fluent.

The fluentd conf is set as follows:

```sh
<source>
  @type forward
  port "#{ENV['FLUENTD_CONTAINER_PORT']}"
  bind 0.0.0.0
</source>

<match "#{ENV['APP_LOGGING_TAG']}">
    @type copy
    <store>
      @type elasticsearch
      host elasticsearch
      port "#{ENV['ELASTICSEARCH_CONTAINER_PORT']}"
      user "#{ENV['ELASTICSEARCH_ELASTIC_USERNAME']}"
      password "#{ENV['ELASTICSEARCH_ELASTIC_PASSWORD']}"
      logstash_format true
      logstash_prefix "#{ENV['FLUENTD_LOGSTASH_PREFIX_APP']}"
      logstash_dateformat %Y%m%d
      include_tag_key true
      type_name "#{ENV['FLUENTD_TYPE_NAME_APP']}"
      tag_key @log_name
      flush_interval 1s
    </store>
</match>
```

In the fluentd conf, you can embed environment variables in the format `#{...}`, which is convenient because you can embed variables without using envsubst.

docker-compose.yml has no special notes, so it is omitted.

## elasticsearch
Elasticsearch is set to start as a single node. There is nothing else to note, so details are omitted.

```yaml
├── elasticsearch
│   └──Dockerfile
```

## kibana
Next, let's talk about kibana, which visualizes application logs.

There are no special notes about the Dockerfile, so it is omitted, and I will explain the kibana conf.

```yaml
server.name: kibana
server.host: "0"
elasticsearch.hosts: [ "http://${ELASTICSEARCH_CONTAINER_NAME}:${ELASTICSEARCH_CONTAINER_PORT}" ]
xpack.monitoring.ui.container.elasticsearch.enabled: true
elasticsearch.username: ${ELASTICSEARCH_ELASTIC_USERNAME}
elasticsearch.password: ${ELASTICSEARCH_ELASTIC_PASSWORD}
```

`xpack.monitoring.ui.container.elasticsearch.enabled` is an option that needs to be enabled if elasticsearch is running in a container.

There are no special notes about kibana's docker-compose.yml, so it is omitted.

## node-exporter & cadvisor
For node-exporter and cadvisor, it is enough to be aware of the directories to mount and the startup options, so the explanation is omitted.

## prometheus
Next is prometheus.

I wanted to write the prometheus configuration file using envsubst, so the Dockerfile is as follows:

```yaml
# NOTE: see https://www.robustperception.io/environment-substitution-with-docker
ARG ALPINE_IMAGE_NAME=${ALPINE_IMAGE_NAME}
ARG ALPINE_IMAGE_TAG=${ALPINE_IMAGE_TAG}
ARG PROMETHEUS_IMAGE_NAME=${PROMETHEUS_IMAGE_NAME}
ARG PROMETHEUS_IMAGE_TAG=${PROMETHEUS_IMAGE_TAG}
ARG PLATFORM=${PLATFORM}

FROM --platform=${PLATFORM} ${PROMETHEUS_IMAGE_NAME}:${PROMETHEUS_IMAGE_TAG} as build-stage

FROM --platform=${PLATFORM} ${ALPINE_IMAGE_NAME}:${ALPINE_IMAGE_TAG}

RUN apk add gettext

COPY --from=build-stage /bin/prometheus /bin/prometheus

RUN mkdir -p /prometheus /etc/prometheus \
  && chown -R nobody:nogroup etc/prometheus /prometheus

COPY ./template/prometheus.yml.template /template/prometheus.yml.template

USER nobody

VOLUME [ "/prometheus" ]

WORKDIR /prometheus
```

docker-compose.yml is as follows:

```yaml
  prometheus:
    container_name: "${PROMETHEUS_CONTAINER_NAME}"
    environment: 
      - PROMETHEUS_IMAGE_NAME=${PROMETHEUS_IMAGE_NAME}
      - PROMETHEUS_IMAGE_TAG=${PROMETHEUS_IMAGE_TAG}
      - PROMETHEUS_CONTAINER_NAME=${PROMETHEUS_CONTAINER_NAME}
      - PROMETHEUS_CONTAINER_PORT=${PROMETHEUS_CONTAINER_PORT}
      - CADVISOR_CONTAINER_NAME=${CADVISOR_CONTAINER_NAME}
      - CADVISOR_CONTAINER_PORT=${CADVISOR_CONTAINER_PORT}
      - NODE_EXPORTER_CONTAINER_NAME=${NODE_EXPORTER_CONTAINER_NAME}
      - NODE_EXPORTER_CONTAINER_PORT=${NODE_EXPORTER_CONTAINER_PORT}
    build:
      context: "./prometheus"
      dockerfile: "Dockerfile"
      args:
        PLATFORM: "${PLATFORM}"
        PROMETHEUS_IMAGE_NAME: "${PROMETHEUS_IMAGE_NAME}"
        PROMETHEUS_IMAGE_TAG: "${PROMETHEUS_IMAGE_TAG}"
        ALPINE_IMAGE_NAME: "${ALPINE_IMAGE_NAME}"
        ALPINE_IMAGE_TAG: "${ALPINE_IMAGE_TAG}"
    ports: 
      - ${PROMETHEUS_HOST_PORT}:${PROMETHEUS_CONTAINER_PORT}
    command:
      - /bin/sh
      - -c
      - |
        envsubst < /template/prometheus.yml.template > /etc/prometheus/prometheus.yml
        /bin/prometheus \
        --config.file=/etc/prometheus/prometheus.yml \
        --storage.tsdb.path=/prometheus
    restart: always
```

The prometheus configuration file is written as follows.
```sh
scrape_configs:
  - job_name: 'prometheus'
    scrape_interval: 5s
    static_configs:
      - targets:
        - ${PROMETHEUS_CONTAINER_NAME}:${PROMETHEUS_CONTAINER_PORT}
  - job_name: 'cadvisor'
    static_configs:
      - targets:
        - ${CADVISOR_CONTAINER_NAME}:${CADVISOR_CONTAINER_PORT}
  - job_name: 'node-exporter'
    static_configs:
      - targets:
        - ${NODE_EXPORTER_CONTAINER_NAME}:${NODE_EXPORTER_CONTAINER_PORT}
```

It only writes about the job names and targets to scrape. If you want to set up alert notifications using Alertmanager, you will need to add Alertmanager settings to this configuration file.

## grafana
Finally, grafana.

docker-compose.yml is as follows:

```yaml
grafana:
    container_name: "${GRAFANA_CONTAINER_NAME}"
    environment: 
      - GF_SECURITY_ADMIN_USER=${GF_SECURITY_ADMIN_USER}
      - GF_SECURITY_ADMIN_PASSWORD=${GF_SECURITY_ADMIN_PASSWORD}
      - GF_USERS_ALLOW_SIGN_UP="${GF_USERS_ALLOW_SIGN_UP}"
      - GF_USERS_ALLOW_ORG_CREATE="${GF_USERS_ALLOW_ORG_CREATE}"
      - DS_PROMETHEUS=${DS_PROMETHEUS}
    build:
      context: "./grafana"
      dockerfile: "Dockerfile"
      args:
        PLATFORM: "${PLATFORM}"
        GRAFANA_IMAGE_NAME: "${GRAFANA_IMAGE_NAME}"
        GRAFANA_IMAGE_TAG: "${GRAFANA_IMAGE_TAG}"
    volumes: 
      - ./grafana/provisioning:/etc/grafana/provisioning
    ports:
      - ${GRAFANA_HOST_PORT}:${GRAFANA_CONTAINER_PORT}
    restart: always
```

`provisioning` is the directory where files used for provisioning data sources and dashboards are placed.

Prometheus is used as the data source, so the prometheus settings are written in datasources/datasource.yml.

```
apiVersion: 1

datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    orgId: 1
    url: http://prometheus:9090
    basicAuth: false
    isDefault: true
    editable: true
```

Dashboard configuration files for container metrics and OS metrics are prepared.

You can use the dashboard configuration files published at [grafana.com - grafana/dashboards](https://grafana.com/grafana/dashboards/).

Building a dashboard from scratch can be quite challenging, so it seems good to find a base and adjust it.

The published ones are quite comprehensive, so it's interesting to try various ones.

# Startup
Most configuration values are structured to be adjustable with environment variables.

Once you copy `.env.example` to `.env` in [bmf-san/docker-based-monitoring-stack-boilerplate](https://github.com/bmf-san/docker-based-monitoring-stack-boilerplate), you can start it with `docker-compose up`.

In the `.env.example` settings, port numbers are assigned as follows:

|   Application   |             URL             |
| --------------- | --------------------------- |
| app             | http://localhost:8080/      |
| prometheus      | http://localhost:9090/graph |
| node-exporter   | http://localhost:9100/      |
| mysqld-exporter | http://localhost:9104/      |
| grafana         | http://localhost:3000/      |
| kibana          | http://0.0.0.0:5601/        |

![Screenshot 2021-12-18 4 04 19](/assets/images/posts/container-monitoring-infrastructure/146595299-9e21d2b8-3b3d-4931-a739-ea3a8e69fa13.png)

# Summary
It seems we were able to build it quite easily (thanks to the benefits of containers).

The architecture of each application is profound, so it might be interesting to look into the mechanisms after trying them out.

I haven't been able to actually operate it yet, so I hope to get it up and running soon.