---
title: Setting Up Elasticsearch and Kibana with Docker
description: An in-depth look at Setting Up Elasticsearch and Kibana with Docker, covering key concepts and practical insights.
slug: docker-elasticsearch-kibana-setup
date: 2018-10-22T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Elasticsearch
  - Kibana
translation_key: docker-elasticsearch-kibana-setup
---



# Overview
A brief summary of Elasticsearch and setting up the environment with Docker.

# What is Elasticsearch
- A distributed RESTful search/analytics engine
  - Capable of not only full-text search but also analytics
- Near real-time search platform
- Cluster
  - A collection of one or more nodes (servers) that collectively hold your entire data
  - Provides integrated indexing and search capabilities across all nodes
- Node
  - A single server within a cluster that stores data

# Setting Up the Environment with Docker
Set up an environment where Elasticsearch and Kibana can be used.

docker-compose.yml
```
elasticsearch:
  image: elasticsearch:5
  ports:
    - 9200:9200
    - 9300:9300
  volumes:
    - ./elasticsearch/data:/usr/share/elasticsearch/data/
kibana:
  image: kibana:5
  ports:
    - 5601:5601
  links:
    - elasticsearch
  environment:
    - ELASTICSEARCH_URL=http://127.0.0.1:9200
```

# References
- [Elasticsearch](https://www.elastic.co/jp/products/elasticsearch)
- [Running Elasticsearch and Kibana with Docker](http://mezina1942.hatenablog.com/entry/2017/10/31/025825)
