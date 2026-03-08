---
title: Environment Variables Specified in docker-compose.yml Cannot Be Referenced During Container Build
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
I specified `env_file` in one of the services in docker-compose.yml to set environment variables, but they could not be referenced inside the container being built (on the Dockerfile side). I wanted to reference the environment variable during the build of the application using npm in the container as `process.env.VUE_APP_API_ENDPOINT`.

# Solution
The keys like `env_file` and `environment` specified in docker-compose.yml can only be referenced after the container is built, so simply using those keys does not allow referencing during the container build.

I resolved this by specifying the `args` key in docker-compose.yml to pass the variable to the container.

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

# Receive arguments and define environment variables inside the container
ARG VUE_APP_API_ENDPOINT
ENV VUE_APP_API_ENDPOINT=${VUE_APP_API_ENDPOINT}

# Build the application. Can reference environment variables.
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
    # Read environment variables from file
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

For reference, here is the code from the application side where I want to reference the environment variable during the build.
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