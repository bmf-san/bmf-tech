---
title: Using Laravel's Deployment Tool Laravel Envoy
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

There are various deployment tools like Capistrano, Fabric, Rocketeer, and Deployer, but mastering them can be quite challenging...

If you're a Laravel user, there's a simple and easy-to-use deployment tool available.

[Laravel Envoy](https://github.com/laravel/envoy)

Compared to general deployment tools, it may not seem capable of handling intricate tasks, but I believe it can perform the essential deployment tasks.

# Installation
As mentioned in the documentation...
`composer global require "laravel/envoy=~1.0``

# Writing Deployment Tasks
Prepare a file named `envoy.blade.php` and write the tasks in it according to the Blade syntax.

You can directly write shell commands inside the @task, which makes it easy.

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

I intended to optimize composer and artisan for Laravel deployment, but I'm not very sure about it, haha. Please don't mind git and slack... they are just examples. (｀･ω･´)

# Deployment

You can deploy with

`envoy run deploy`

# Thoughts
Although I skipped some details, it can be easily set up from installation to usage, so this should be sufficient for small projects!