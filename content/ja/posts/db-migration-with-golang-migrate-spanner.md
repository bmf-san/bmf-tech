---
title: golang-migrateを使ってspannerのDBマイグレーションをする
slug: db-migration-with-golang-migrate-spanner
date: 2021-03-17T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Google Cloud Platform
  - Golang
  - Spanner
translation_key: db-migration-with-golang-migrate-spanner
---


# 概要
spannerのDBマイグレーションで、[golang-migrate](https://github.com/golang-migrate/migrate)を使ったのでメモ。

# 使い方
dockerで使う想定。
dockerではなくバイナリで実行していたが、ホストマシンのopensslのバージョンに依存して動作しない可能性あるようなので、コンテナ実行が無難だと思う。

```sh
MIGRATE_VERSION='v4.14.1'

docker run -v /migrations:/migrations -v ~/.config/gcloud/:/root/.config/gcloud --network host migrate/migrate:${MIGRATE_VERSION} -path=/migrations/ -database spanner://projects/<PROJECT ID>/instances/<INSTANCE>/databases/<DATABASE>?x-clean-statements=True <COMMAND>
```

COMMANDには、up、down、versionなど指定する。
cf. [github.com - golang/migrrate/migrate/tree/aster/cmd/migrate#usage](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#usage)

ワンライナーで実行コマンドを書いたので見づらいと思うが、特に難しいとこはないと思う。

マウントしているやつは以下２つ。
```
-v /migrations:/migrations -v ~/.config/gcloud/:/root/.config/gcloud
```

`/migrations`にはマイグレーション対象のsqlファイルを用意している。それをコンテナの`/migrations`にマウント。

`~/.config/gcloud/:/root/.config/gcloud` はgcloudの認証を通すため。
credentialファイルをマウントして、環境変数`GOOGLE_APPLICATION_CREDENTIALS`をセットする形でも認証を通せるが、この方が楽なので・・

直近のgolang-migrateのバージョンではspannerの接続情報にはクエリパラメータが必要。

```sh
spanner://projects/<PROJECT ID>/instances/<INSTANCE>/databases/<DATABASE>?x-clean-statements=True
```

経緯はこのへん。
- [diff v4.11.0...v4.12.0](https://github.com/golang-migrate/migrate/compare/v4.11.0...v4.12.0)
- [issues](https://github.com/golang-migrate/migrate/search?q=x-clean-statements&type=issues)



