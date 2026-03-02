---
title: "Docker CLIでDocker Composeが使えるようになった"
slug: "docker-cli-docker-compose"
date: 2021-05-06
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Docker"
  - "Docker Compose"
draft: false
---

最近のアップデートでDocker CLIがDocker Composeをサポートするようになったらしい。

cf. [docs.docker.com - Compose CLI Tech Preview](https://docs.docker.com/compose/cli-command/)

Tech Previewなのでまだプロダクションでの利用推奨されていないとのこと。

互換性についてはこちら。
[docs.docker.com - Compose command compatibility with docker-compose](https://docs.docker.com/compose/cli-command-compatibility/)

Docker composeはpythonで実装されているが、今回はサポートされるDocker Composeはgolang製らしい。

cf. https://github.com/docker/compose
cf. https://github.com/docker/compose-cli
