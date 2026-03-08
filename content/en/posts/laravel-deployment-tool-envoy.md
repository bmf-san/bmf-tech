---
title: 'Trying Out Laravel''s Deployment Tool: Laravel Envoy'
slug: laravel-deployment-tool-envoy
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Laravel
  - Deployment
translation_key: laravel-deployment-tool-envoy
---



Capistrano, Fabric, Rocketeer, Deployer... There are various deployment tools, but mastering them can be quite challenging...

If you're a Laravel user, there's a simple and easy-to-use deployment tool available.

[Laravel Envoy](https://github.com/laravel/envoy)

Compared to general deployment tools, it might not handle all the intricate details, but it should be able to perform the essential deployment tasks.

# Introduction
As mentioned in the documentation...
`composer global require "laravel/envoy=~1.0"`


# Writing Deployment Tasks
Prepare a file named envoy.blade.php and write tasks in it following the Blade syntax.

Tasks can be directly written as shell commands within @task, making it convenient.

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
  TITLE=${TITLE:-"Production Deployment Notification"}
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

I've tried to optimize composer and artisan for Laravel deployment, but I'm not entirely sure, lol.
Don't worry about git or slack... they're just examples. (｀･ω･´)

# Deployment

`envoy run deploy`

This command will deploy.

# Impressions
I skipped a few details, but since it's easy to set up from introduction to use, it might be sufficient for small projects!