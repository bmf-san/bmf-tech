---
title: Setting Up Elasticsearch and Kibana with Docker
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
A brief summary of Elasticsearch and how to set up the environment using Docker.

# What is Elasticsearch?
- Distributed RESTful search/analytics engine
  - Capable of both full-text search and analysis
- Almost real-time search platform
- Cluster
  - A collection of one or more nodes (servers) that hold all the data together
  - Provides integrated indexing and search capabilities across all nodes
- Node
  - A single server that is part of the cluster and stores data

# Setting Up the Environment with Docker
Creating an environment where Elasticsearch and Kibana can be used.

```yaml
docker-compose.yml
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