---
title: DockerでElasticsearchとKibanaの環境構築
description: DockerでElasticsearchとKibanaの環境構築の手順と実践例を詳しく解説します。
slug: docker-elasticsearch-kibana-setup
date: 2018-10-22T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Elasticsearch
  - Kibana
translation_key: docker-elasticsearch-kibana-setup
---


# 概要
ElasticSearchについての説明とDockerでの環境構築についてざっくりとまとめる。

# ElasticSearchとは
- 分散型RESTful検索/分析エンジン
 - 全文検索だけでなく、分析もできる
- ほぼリアルタイムの検索プラットフォーム
- クラスタ
 - データ全体をまとめて保持する1つ以上のノード（サーバー）のコレクションのこと
 - 全てのノードに渡って統合されたインデキシング機能と検索機能が提供される
- ノード
 - データを保存するクラスタに含まれる1台のサーバー

# Dockerで環境構築
ElasticsearchとKibanaが使える環境を構築する。

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

# 参考
- [Elasticsearch](https://www.elastic.co/jp/products/elasticsearch)
- [DockerでElasticsearchとKibanaを動かす](http://mezina1942.hatenablog.com/entry/2017/10/31/025825)

