---
title: Building a Monitoring Infrastructure with Containers
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

This article is the 17th entry in the [Makuake Advent Calendar 2021](https://adventar.org/calendars/6822).

It has been exactly three years since I joined the company, and this is my third participation in the company's Advent Calendar.

For the past year, I have been part of the Re-Architecture team, which is responsible for the development and operation of Makuake's service infrastructure, and I have been working hard on various tasks.

I will surely continue to tackle various challenges next year as well.

You can find the job openings for the Re-Architecture team here.
[【Go/Microservices】Hiring Engineers for the Re-Architecture Team to Revamp Makuake's Infrastructure!](https://findy-code.io/companies/633/jobs/3hdLodfPABhrO)

Now, the first article of this year (I will probably write another one on the 24th) is "Building a Monitoring Infrastructure with Containers".

I have been experimenting with creating a monitoring infrastructure for an application I am developing as a hobby, which is not directly related to my main job, and I would like to share my insights (though they may not be substantial...) from that experience.

# Monitoring Infrastructure Configuration
![Screenshot 2021-12-18 4 04 19](https://user-images.githubusercontent.com/13291041/146595314-d593b9d8-9faa-4275-8c94-6c12cfcbfe36.png)

The applications that make up the monitoring infrastructure I will build are as follows:

- [Elasticsearch](https://www.elastic.co/jp/elasticsearch/)
	- A search engine that accumulates application logs. I will implement a simple application in Go to collect the logs.
- [Fluentd](https://www.fluentd.org/)
	- A log aggregator that collects logs and forwards them to Elasticsearch.
	- I could have used the ELK stack (Logstash instead of Fluentd), but I opted for Fluentd since I am more familiar with it.
- [Kibana](https://www.elastic.co/jp/kibana/)
	- A UI for data search, visualization, and analysis. It visualizes application logs.
- [Grafana](https://grafana.com/)
	- Similar to Kibana, it is a UI for data visualization. It is used to visualize system metrics.
	- While it can also be used for visualizing application logs, I will use Kibana for that purpose.
- [Prometheus](https://prometheus.io/)
	- A monitoring tool for system metrics. It collects system metrics in conjunction with node-exporter and cadvisor. The collected data is visualized using Grafana.
- [node-exporter](https://github.com/prometheus/node_exporter)
	- Collects OS metrics.
- [cadvisor](https://github.com/google/cadvisor)
	- Collects container metrics.

I have roughly structured this for those who want to build and play around with it.

These applications will be built using docker-compose.

# Building the Monitoring Infrastructure
All implementations are available at [bmf-san/docker-based-monitoring-stack-boilerplate](https://github.com/bmf-san/docker-based-monitoring-stack-boilerplate).

Once cloned, you can create a `.env` file and run `docker-compose up` to get started right away.

By the way, cadvisor does not start on M1, so container metrics cannot be collected. It has been confirmed to work on Intel Macs and Ubuntu.

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

I will explain each one.

## app
First, I will create a rough application container that outputs logs.

```sh
├── app
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
```

The application looks like this. It simply outputs "OK" in the logs and responds with "Hello World".

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

The Dockerfile is a simple one that builds the source and runs the binary, with no special notes.

The docker-compose.yml looks like this:

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

The logging driver is set to fluentd, which transfers logs to fluentd.

The `fluent-async-connect` option buffers logs until a connection to fluentd is established, and when set to true, it buffers logs even if the connection is not yet established.

## fluentd
Next, I will explain the fluentd container, which is the destination for application log transfers.

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

The only gem used in fluentd is fluent-plugin-elasticsearch for integration with Elasticsearch.

The USER is set to root and then switched back to fluent because the execution user of the fluentd image is fluent.

The fluentd configuration is set as follows:

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

In the fluentd configuration, you can embed environment variables in the format `#{...}`, making it convenient to use without needing envsubst or similar tools.

The docker-compose.yml for fluentd is omitted as there are no special notes.

## elasticsearch
Elasticsearch is configured to run as a single node. There are no other notable points, so I will skip the details.

```yaml
├── elasticsearch
│   └──Dockerfile
```

## kibana
Next, I will discuss Kibana, which is used for visualizing application logs.

There are no special notes for the Dockerfile, so I will skip that and explain the Kibana configuration:

```yaml
server.name: kibana
server.host: "0"
elasticsearch.hosts: [ "http://${ELASTICSEARCH_CONTAINER_NAME}:${ELASTICSEARCH_CONTAINER_PORT}" ]
xpack.monitoring.ui.container.elasticsearch.enabled: true
elasticsearch.username: ${ELASTICSEARCH_ELASTIC_USERNAME}
elasticsearch.password: ${ELASTICSEARCH_ELASTIC_PASSWORD}
```

The option `xpack.monitoring.ui.container.elasticsearch.enabled` needs to be enabled if Elasticsearch is running in a container.

The docker-compose.yml for Kibana is also omitted as there are no special notes.

## node-exporter & cadvisor
For node-exporter and cadvisor, I will skip the explanation as it mainly involves being aware of the mounted directories and startup options.

## prometheus
Next is Prometheus.

I wanted to write the Prometheus configuration file using envsubst, so the Dockerfile is as follows:

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

The docker-compose.yml is as follows:

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

The Prometheus configuration file is written as follows:
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

I have only written the job names and targets that I want to scrape. If you want to set up alert notifications using Alertmanager, you will need to add the Alertmanager configuration to this configuration file as well.

## grafana
Finally, let's discuss Grafana.

The docker-compose.yml looks like this:

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

The `provisioning` directory is where files for data source and dashboard provisioning are placed.

Since we will use Prometheus as the data source, the Prometheus configuration is written in `datasources/datasource.yml`:

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

I have prepared configuration files for dashboards for container metrics and OS metrics.

You can use dashboard configuration files available on [grafana.com - grafana/dashboards](https://grafana.com/grafana/dashboards/).

Building a dashboard from scratch can be quite challenging, so it seems better to look for a base one and adjust it.

There are many publicly available options, so it can be interesting to explore them.

# Startup
Most of the configuration values can be adjusted using environment variables.

You can copy `.env.example` from [bmf-san/docker-based-monitoring-stack-boilerplate](https://github.com/bmf-san/docker-based-monitoring-stack-boilerplate) to `.env`, and then start it with `docker-compose up`.

The settings in `.env.example` assign port numbers as follows:

|   Application   |             URL             |
| --------------- | --------------------------- |
| app             | http://localhost:8080/      |
| prometheus      | http://localhost:9090/graph |
| node-exporter   | http://localhost:9100/      |
| mysqld-exporter | http://localhost:9104/      |
| grafana         | http://localhost:3000/      |
| kibana          | http://0.0.0.0:5601/        |

![Screenshot 2021-12-18 4 04 19](https://user-images.githubusercontent.com/13291041/146595299-9e21d2b8-3b3d-4931-a739-ea3a8e69fa13.png)

# Conclusion
I think it was relatively easy to set up (thanks to the benefits of containers).
The architectural structure of each application is quite deep, so after touching them, it might be interesting to take a look at how they work.
Since I have not yet been able to operate it in practice, I am eager to get it into operation soon.