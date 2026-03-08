---
title: Using golang-migrate for Spanner DB Migration
slug: db-migration-with-golang-migrate-spanner
date: 2021-03-17T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Google Cloud Platform
  - Golang
  - Spanner
translation_key: db-migration-with-golang-migrate-spanner
---

# Overview
I used [golang-migrate](https://github.com/golang-migrate/migrate) for Spanner DB migration, so here are my notes.

# How to Use
Assuming usage with Docker.
I initially ran it with a binary instead of Docker, but it seems there might be issues depending on the version of OpenSSL on the host machine, so running it in a container seems safer.

```sh
MIGRATE_VERSION='v4.14.1'

docker run -v /migrations:/migrations -v ~/.config/gcloud/:/root/.config/gcloud --network host migrate/migrate:${MIGRATE_VERSION} -path=/migrations/ -database spanner://projects/<PROJECT ID>/instances/<INSTANCE>/databases/<DATABASE>?x-clean-statements=True <COMMAND>
```

For <COMMAND>, you can specify up, down, version, etc.
cf. [github.com - golang/migrate/migrate/tree/master/cmd/migrate#usage](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#usage)

I wrote the execution command in a one-liner, so it might be hard to read, but I don't think there are any particularly difficult parts.

The two mounted volumes are as follows:
```
-v /migrations:/migrations -v ~/.config/gcloud/:/root/.config/gcloud
```

In `/migrations`, I have prepared the SQL files for migration. These are mounted to the container's `/migrations`.

`~/.config/gcloud/:/root/.config/gcloud` is for passing the gcloud authentication.
You can also authenticate by mounting the credential file and setting the environment variable `GOOGLE_APPLICATION_CREDENTIALS`, but this way is easier...

In the latest version of golang-migrate, query parameters are required for the Spanner connection information.

```sh
spanner://projects/<PROJECT ID>/instances/<INSTANCE>/databases/<DATABASE>?x-clean-statements=True
```

Here’s the background:
- [diff v4.11.0...v4.12.0](https://github.com/golang-migrate/migrate/compare/v4.11.0...v4.12.0)
- [issues](https://github.com/golang-migrate/migrate/search?q=x-clean-statements&type=issues)