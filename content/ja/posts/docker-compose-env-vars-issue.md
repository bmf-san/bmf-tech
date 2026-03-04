---
title: "docker-compose.ymlで指定した環境変数がコンテナビルド中に参照できない"
slug: "docker-compose-env-vars-issue"
date: 2020-11-14
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Docker"
  - "Docker Compose"
  - "Tips"
draft: false
---

# 概要
docker-compose.ymlのserviceの1つに`env_file`を指定し、環境変数を設定したが、サービスがbuildするコンテナ内（Dockerfile側）では参照できなかった。
vueのアプリケーションをコンテナ内でnpmを使ってビルドしており、アプリケーション側で`process.env.VUE_APP_API_ENDPOINT`という形でアプリケーションのビルド時に環境変数を参照させたかった。

# 解決策
docker-compose.ymlで指定する`env_file`や`environment`といったキーはコンテナのビルド後に参照できるようになるため、それらのキーを利用するだけではコンテナビルド中では参照することができない。

docker-compose.ymlで`args`キーを指定し、変数をコンテナに渡すことで解決した。

.env
```sh
VUE_APP_API_ENDPOINT="http://gobel-api.local"
```

Dockerfile
```
FROM node:14.3.0-alpine as build-stage

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

# 引数を受け取ってコンテナ内で環境変数を定義
ARG VUE_APP_API_ENDPOINT
ENV VUE_APP_API_ENDPOINT=${VUE_APP_API_ENDPOINT}

# アプリケーションのビルド。環境変数を参照できる。
RUN npm run local-build

FROM nginx:1.19.0-alpine

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./nginx/conf.d/gobel-admin-client.conf /etc/nginx/conf.d/gobel-admin-client.conf
COPY --from=build-stage /app/dist /var/www/html
```

docker-compose.yml
```yaml
version: "3.8"
services:
  app:
    container_name: "gobel-admin-client"
    # 環境変数はファイルから読み込む
    env_file: ".env"
    build:
        context: "./app"
        dockerfile: "Dockerfile"
        # 変数をコンテナのビルド時に渡す
        args:
          VUE_APP_API_ENDPOINT: $VUE_APP_API_ENDPOINT
    ports:
      - "82:80"
    networks:
      - gobel_link
networks:
    gobel_link:
        external: true
```

参考までにビルド時に環境変数を参照したいアプリケーション側のコードを記載。
```js
const apiClient = axios.create({
  baseURL: process.env.VUE_APP_API_ENDPOINT,
  headers: {
    "Content-Type": "application/json"
  },
  responseType: "json"
});
```

# 参考
- [docs.docker.jp - Compose における環境変数](https://docs.docker.jp/compose/environment-variables.html)
- [qiita.com - docker-composeのビルド中に環境変数が認識されない](https://qiita.com/katoosky/items/422c183cf5cabb789030)
