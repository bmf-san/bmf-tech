---
title: "コンテナで始めるモニタリング基盤構築"
slug: "post-277"
date: 2021-12-18
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Docker"
  - "Docker Compose"
  - "Elasticsearch"
  - "Grafana"
  - "Kibana"
  - "Prometheus"
  - "cadvisor"
  - "efk"
  - "fluentd"
  - "node-exporter"
draft: false
---

この記事は[Makuake Advent Calendar 2021](https://adventar.org/calendars/6822)の17日目の記事です。

気づけば入社して丸3年が経ち、会社のアドベントカレンダーも3回目の参戦です。

ここ1年はRe-ArchitectureチームというMakuakeのサービス基盤の開発・運用を行うチームに所属し、色々と奮闘してきました。

来年もきっとあれこれと奮闘するでしょう。

Re-Architectureチームの求人はこちら。
[【Go/マイクロサービス】「Makuake」の基盤を刷新するRe-Architectureチームのエンジニア募集！](https://findy-code.io/companies/633/jobs/3hdLodfPABhrO)

さて、今年の一本目の記事（24日にもう一本かく、たぶん）は「コンテナで始めるモニタリング基盤構築」です。

特に本業とは関係なく、趣味で作っているアプリケーションのモニタリング基盤をコンテナでいい感じにしてみたいと思ってあれこれ試していたので、その時の知見（というほどでもないですが・・・）を公開しようと思います。

# モニタリング基盤構成
![スクリーンショット 2021-12-18 4 04 19](https://user-images.githubusercontent.com/13291041/146595314-d593b9d8-9faa-4275-8c94-6c12cfcbfe36.png)

今回構築するモニタリング基盤を構成するアプリケーションは以下です。

- [Elasticsearch](https://www.elastic.co/jp/elasticsearch/)
	- 検索エンジン。アプリケーションログを蓄積します。ログを収集するアプリケーション自体はGoで簡易的なアプリを実装します。
- [Fluentd](https://www.fluentd.org/)
	- ログアグリゲーター。ログを収集して、Elasticsearchに転送します。
	- EFKスタックではなくELK（FluentdではなくLogstash）でも良かったのですが、馴染みのあるFluentdのほうを採用しました。
- [Kibana](https://www.elastic.co/jp/kibana/)
	- データ検索・可視化・分析のUI。アプリケーションログを可視化します。
- [Grafana](https://grafana.com/)
	- Kibanaと同じく、データのためのUI。システムメトリクスの可視化に利用します。
	- アプリケーションログの可視化にも利用できますが、アプリケーションログはKiabanaを利用します。
- [Prometheus](https://prometheus.io/)
	- システムメトリクスの監視ツール。node-exporterやcadvisorと連携してシステムメトリクスを収集します。収集したデータの可視化はGrafanaで行います。
- [node-exporter](https://github.com/prometheus/node_exporter)
	- OSメトリクスを収集します。
- [cadvisor](https://github.com/google/cadvisor)
	- コンテナメトリクスを収集します。

とりあえず構築して遊んでみたい人向けにざっくり構成してみました。

これらのアプリケーションをdocker-composeで構築します。

# モニタリング基盤構築
全ての実装は[bmf-san/docker-based-monitoring-stack-boilerplate](https://github.com/bmf-san/docker-based-monitoring-stack-boilerplate)に置いてあります。

cloneしたら`.env`を作って`docker-compose up`するだけですぐに触れるようになっています。

ちなみにM1だとcadvisorが起動しないため、コンテナメトリクスが収集できません。intel macやubuntuでは動作確認済みです。

ディレクトリ構成は以下の通りです。1コンテナずつ解説していきます。

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

1つひとつ解説していきたいと思います。

## app
最初にログを吐く雑なアプリケーションコンテナを作ります。

```sh
├── app
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
```

アプリケーションはこんな感じです。”OK”とログを吐いて、"Hello World"とレスポンスするだけのサーバーです。

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

Dockerfileはソースをビルドしてバイナリを実行するだけの単純なもので、特に補足はありません。

docker-compose.ymlのほうは以下のような形になります。

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

logging driverにfluentdを指定して、ログをfluentdに転送します。

`fluent-async-connect`はfluentdと接続が確立できるまでログをバッファリングする設定で、trueの場合は接続が確立していなくてもログをバッファリングしてくれます。

## fluentd
アプリケーションのログ転送先であるfluentdのコンテナについて解説します。

```sh
├── fluentd
│   ├── Dockerfile
│   └── config
│       └── fluent.conf
```

Dockerfileは下記です。

```yml
ARG FLUENTD_IMAGE_NAME=${FLUENTD_IMAGE_NAME}
ARG FLUENTD_IMAGE_TAG=${FLUENTD_IMAGE_TAG}
ARG PLATFORM=${PLATFORM}

FROM --platform=${PLATFORM} ${FLUENTD_IMAGE_NAME}:${FLUENTD_IMAGE_TAG}

USER root

RUN gem install fluent-plugin-elasticsearch

USER fluent
```

fluentdで使っているgemはelasticsearchと連携するためのfluent-plugin-elasticsearchだけです。

USERをrootにして、最後にfluentに戻しているのは、fluentdのイメージの実行ユーザーがfluentな為です。

fluentdのconfは以下のように設定します。

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

fluentdのconfでは`#{...}`という形式で環境変数を埋め込むことができるので、envsubstなどを利用しなくても変数を埋め込むことができ便利です。

docker-compose.ymlの方は特記事項がないため割愛します。

## elasticsearch
elasticsearchはシングルーノードで起動するように設定します。
他に特記することがないので詳細は省きます。

```yaml
├── elasticsearch
│   └──Dockerfile
```

## kibana
続いて、アプリケーションログの可視化をするkibanaについてです。

Dockerfileについては特記事項がないので記載を割愛して、kibanaのconfから説明します。

```yaml
server.name: kibana
server.host: "0"
elasticsearch.hosts: [ "http://${ELASTICSEARCH_CONTAINER_NAME}:${ELASTICSEARCH_CONTAINER_PORT}" ]
xpack.monitoring.ui.container.elasticsearch.enabled: true
elasticsearch.username: ${ELASTICSEARCH_ELASTIC_USERNAME}
elasticsearch.password: ${ELASTICSEARCH_ELASTIC_PASSWORD}
```

`xpack.monitoring.ui.container.elasticsearch.enabled`は、elasticsearchがコンテナで実行されている場合は有効化して置く必要のあるオプションです。

kibanaのdocker-compose.ymlについては特記事項がないため割愛します。

## node-exporter & cadvisor
node-exporterとcadvisorについてマウントするディレクトリや起動オプションについて意識する程度なので説明は割愛します。

## prometheus
続いてprometheusです。

envsubstを使ってprometheusの設定ファイルを書きたかったので、Dockerfileを下記のようにしています。

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

docker-compose.ymlは下記のようになります。

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

prometheusの設定ファイルは次のように書いています。
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

スクレイプしたいジョブ名とターゲットについてだけ書いています。Alertmanagerを使ったアラート通知を設定したい場合はAlertmanagerの設定もこの設定ファイルに追記することになります。

## grafana
最後にgrafanaです。

docker-compose.ymlは次のようになります。

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

`provisioning`はデータソースやダッシュボードのプロジョニングで使うファイルを置いておくディレクトリです。

データソースにはprometheusを利用するので、datasources/datasource.ymlにprometheusの設定を記載しています。

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

ダッシュボードはコンテナメトリクス用とOSメトリクス用のダッシュボードの設定ファイルを用意しています。

ダッシュボードの設定ファイルは[grafana.com - grafana/dashboards](https://grafana.com/grafana/dashboards/)で公開されているものを利用することができます。

ダッシュボードはゼロから組み立てるとそこそこ大変なので、ベースになるものを探してそれを調整するのが良さそうに思います。

公開されているものはかなり充実しているので、色々触ってみると面白いです。

# 起動
設定値のほとんどを環境変数で調整できるように構成しています。

[bmf-san/docker-based-monitoring-stack-boilerplate](https://github.com/bmf-san/docker-based-monitoring-stack-boilerplate)の`.env.example`を`.env`としてコピーしたら、`docker-compose up`で起動できます。

`.env.example`の設定では以下のようにポート番号を振っています。

|   Application   |             URL             |
| --------------- | --------------------------- |
| app             | http://localhost:8080/      |
| prometheus      | http://localhost:9090/graph |
| node-exporter   | http://localhost:9100/      |
| mysqld-exporter | http://localhost:9104/      |
| grafana         | http://localhost:3000/      |
| kibana          | http://0.0.0.0:5601/        |

![スクリーンショット 2021-12-18 4 04 19](https://user-images.githubusercontent.com/13291041/146595299-9e21d2b8-3b3d-4931-a739-ea3a8e69fa13.png)

# まとめ
割と簡単に構築できたのではないでしょうか（コンテナの恩恵かな）。
それぞれのアプリケーションのアーキテクチャ構成は奥深いので一通り触ったら仕組みをみてみるというのも面白いかと思います。
まだ実際に運用できていないので、早い所運用に乗せてみたい所存です。

