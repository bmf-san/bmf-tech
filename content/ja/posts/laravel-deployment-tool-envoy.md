---
title: "Laravelのデプロイツール Laravel Envoy を使ってみる"
slug: "laravel-deployment-tool-envoy"
date: 2017-10-01
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Laravel"
  - "デプロイ"
draft: false
---

Capistrano, Fabric, Rocketeer, Deployer... など様々なデプロイツールがありますが、使いこなすのが結構大変。。。

もしあなたがLaravelユーザーならシンプルでカンタンに使えるデプロイツールがあります。

[Laravel Envoy](https://github.com/laravel/envoy)

一般的なデプロイツールと比較して、細々としたことはできなさそうですが、必要最低限のデプロイタスクは行えるかと思います。

# 導入
ドキュメントにもありますが・・
`composer global require "laravel/envoy=~1.0"`


# デプロイタスクを記述
envoy.blade.phpというファイルを用意し、その中にタスクをblade記法に従って記述していきます。

タスクはshellコマンドを@taskの中に直接記述することができるので楽です。

```
@servers(['web' => '123.45.678.912'])

@macro('deploy')
  composer
  git
  artisan
  slack
@endmacro

@task('composer')
  cd /var/www/html/Hoge
  composer update
  composer install --no-dev --optimize-autoloader
@endtask

@task('git')
  cd /var/www/html/Hoge
  git pull origin master
@endtask

@task('artisan')
  cd /var/www/html/Hoge
  php artisan down
  php artisan migrate
  php artisan cache:clear
  php artisan config:cache
  php artisan route:cache
  php artisan view:clear
  php artisan up
@endtask

@task('slack')
  cd /var/www/html/Hoge

  # WebHookUrl
  WEBHOOKURL="https://hooks.slack.com/services/hogehogehogehogehogehoge"

  # Slack Channel
  CHANNEL=${CHANNEL:-"#prod-deploy"}

  # Slack Bot Name
  BOTNAME=${BOTNAME:-"Hoge-bot"}

  # Slack Title
  TITLE=${TITLE:-"本番環境デプロイ通知"}
  cd /var/lib/git/Hoge.git

  # Slack Message
  MESSAGE=`git log -1 master`

  #POST
  curl -s -S -X POST --data-urlencode "payload={
          \"channel\": \"${CHANNEL}\",
          \"username\": \"${BOTNAME}\",
          \"attachments\": [{
          \"color\": \"danger\",
          \"fallback\": \"${TITLE}\",
          \"title\": \"${TITLE}\",
          \"text\": \"${MESSAGE}\"
        }]
  }" ${WEBHOOKURL} >/dev/null
@endtask
```

composerとartisanはLaravelのデプロイに最適化したつもりですが、あんまり良くわかっていませんｗ
gitやslackは気にしないでください・・あくまで一例です。く(｀･ω･´)

# デプロイ

`envoy run deploy`

でデプロイできます。

# 所感
多少端折りましたが、導入から利用までお手軽にセットアップできるので、ちょっとしたプロジェクトならこれで十分なのでは！

