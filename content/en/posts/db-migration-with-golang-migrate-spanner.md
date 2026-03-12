---
title: Database Migration for Spanner Using golang-migrate
description: An in-depth look at Database Migration for Spanner Using golang-migrate, covering key concepts and practical insights.
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
I used [golang-migrate](https://github.com/golang-migrate/migrate) for database migration with Spanner, so here's a note.

# Usage
Assuming usage with Docker.
I was running it as a binary instead of Docker, but it might not work depending on the version of OpenSSL on the host machine, so running it in a container seems safer.

```sh
MIGRATE_VERSION='v4.14.1'

docker run -v /migrations:/migrations -v ~/.config/gcloud/:/root/.config/gcloud --network host migrate/migrate:${MIGRATE_VERSION} -path=/migrations/ -database spanner://projects/<PROJECT ID>/instances/<INSTANCE>/databases/<DATABASE>?x-clean-statements=True <COMMAND>
```

Specify commands like up, down, version, etc., in COMMAND.
cf. [github.com - golang/migrate/migrate/tree/master/cmd/migrate#usage](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#usage)

I wrote the execution command as a one-liner, so it might be hard to read, but there shouldn't be anything particularly difficult.

The two mounts are as follows:
```
-v /migrations:/migrations -v ~/.config/gcloud/:/root/.config/gcloud
```

Prepare the SQL files for migration in `/migrations`. Mount them to the container's `/migrations`.

`~/.config/gcloud/:/root/.config/gcloud` is for passing gcloud authentication.
You can also pass authentication by mounting the credential file and setting the environment variable `GOOGLE_APPLICATION_CREDENTIALS`, but this way is easier...

In the latest version of golang-migrate, query parameters are required for Spanner connection information.

```sh
spanner://projects/<PROJECT ID>/instances/<INSTANCE>/databases/<DATABASE>?x-clean-statements=True
```

Here's the background:
- [diff v4.11.0...v4.12.0](https://github.com/golang-migrate/migrate/compare/v4.11.0...v4.12.0)
- [issues](https://github.com/golang-migrate/migrate/search?q=x-clean-statements&type=issues)

