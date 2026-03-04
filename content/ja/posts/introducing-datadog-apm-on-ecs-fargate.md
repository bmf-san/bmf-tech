---
title: "ECS on Fargate環境でDatadog APMを導入"
slug: "introducing-datadog-apm-on-ecs-fargate"
date: 2021-10-19
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Amazon Web Service"
  - "Datadog"
  - "ECS"
  - "Fargate"
draft: false
---

# 概要
ECS on Fargate環境でDatadog APMを導入したときの雑メモ。

# PHPコンテナイメージの調整
php-fpmのイメージをベースとしたカスタムイメージを使っている。
datadog-php-tracerをが必要なので以下のような感じでイメージに組み込んでいる。
```yaml
ENV DDTRACE_VERSION=0.65.1

RUN curl -Lo datadog-php-tracer.apk https://github.com/DataDog/dd-trace-php/releases/download/${DDTRACE_VERSION}/datadog-php-tracer_${DDTRACE_VERSION}_noarch.apk \
    && apk add datadog-php-tracer.apk --allow-untrusted \
    && rm datadog-php-tracer.apk
```

fpmの設定では環境変数を読み取れるように設定。
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

# ECSタスク定義
PHPとDatadogそれぞれに必要な環境をセット。

```json
[
  {
    // PHP
    "environment": [
      {
        "name": "DD_AGENT_HOST",
        "value": "${DATADOG_AGENT_HOST}" // DatadogはサイドカーなのでここはlocalhostでOK
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

# まとめ
これでAPMは稼働する。
最低限の設定としてはこんな感じだろうか。
上手く動かなかったときは、`phpinfo()`でdatadogのセクションを見るといい。

