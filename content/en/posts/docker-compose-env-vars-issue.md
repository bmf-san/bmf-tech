---
title: Environment Variables Specified in docker-compose.yml Not Accessible During Container Build
description: 'Fix environment variables in docker-compose.yml being unavailable during Docker image build. Use ARG in Dockerfile and build.args in docker-compose.yml to pass values at build time.'
slug: docker-compose-env-vars-issue
date: 2020-11-14T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Docker
  - Docker Compose
  - Tips
translation_key: docker-compose-env-vars-issue
---



# Overview
I specified an `env_file` for one of the services in docker-compose.yml to set environment variables, but they were not accessible within the container being built (on the Dockerfile side). I was building a Vue application inside the container using npm and wanted the application to reference environment variables during the build process in the form of `process.env.VUE_APP_API_ENDPOINT`.

# Solution
The `env_file` and `environment` keys specified in docker-compose.yml become accessible after the container is built, so using these keys alone does not allow access during the container build.

The solution was to specify the `args` key in docker-compose.yml and pass variables to the container.

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

# Accept arguments and define environment variables inside the container
ARG VUE_APP_API_ENDPOINT
ENV VUE_APP_API_ENDPOINT=${VUE_APP_API_ENDPOINT}

# Build the application. Environment variables can be referenced.
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
    # Load environment variables from a file
    env_file: ".env"
    build:
        context: "./app"
        dockerfile: "Dockerfile"
        # Pass variables during container build
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

For reference, here is the application code that needs to reference environment variables during the build.
```js
const apiClient = axios.create({
  baseURL: process.env.VUE_APP_API_ENDPOINT,
  headers: {
    "Content-Type": "application/json"
  },
  responseType: "json"
});
```

# References
- [docs.docker.jp - Environment Variables in Compose](https://docs.docker.jp/compose/environment-variables.html)
- [qiita.com - Environment Variables Not Recognized During docker-compose Build](https://qiita.com/katoosky/items/422c183cf5cabb789030)
