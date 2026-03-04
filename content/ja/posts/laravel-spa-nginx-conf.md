---
title: "LaravelにSPAを組み込む時に考えたディレクトリ構成とnginxのconfファイル"
slug: "laravel-spa-nginx-conf"
date: 2017-10-01
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Laravel"
  - "React"
  - "Nginx"
draft: false
---

# 概要
Laravelで作っているアプリケーションに管理画面だけSPAを実装しようとした時、Laravelのディレクトリ構成とnginxのconfファイルをちょっとだけいじった話です。
初めての試みだったのでメモがてらまとめました。

# Laravelのディレクトリ構成
バックエンドで完結するアプリを`backend-app`、フロントエンドで完結するアプリを`frontend-app`とし、ディレクトリを大きく分けました。
`backend-app`ではユーザー側の画面やAPIやバックエンドの処理を担当し、フロントエンドはSPAの管理画面を担当しています。
ユーザー側の画面も`frontend-app`の範疇な気がしますが、その辺は追々切り出していくことにします。
バラバラに切り出すならフルスタックなフレームワークから脱却した方がいいのかもしれませんね。。。
とりあえず今回はフロントはフロント、バックはバックで管理しやすいような構成にしてみました。

ちなみこの構成はgithubで「Laravel SPA」とか「Laravel React」とかで調べていたらよく見受けられた構成を参考にしたものです。

```
.
├── backend-app
│   ├── app
│   ├── bootstrap
│   ├── config
│   ├── database
│   ├── node_modules
│   ├── public
│   ├── resources
│   ├── routes
│   ├── storage
│   ├── tests
│   └── vendor
└── frontend-app
    ├── _components
    ├── dist
    ├── node_modules
    └── src
```


# nginxのconfファイル
locationディレクティブを`backend-app`と`frontend-app`で分けました。
serverディレクティブで分ける場合はサブドメインを切る運用方法になるかと思いますが、それだと今回のアプリケーションの場合微妙だと思ったのでやめました。
もっと設定を頑張る必要がある気がしますが勘弁してください(*_*)。。。

```
server {
    listen       80;
    server_name  laravel-spa;

    root /var/www/html/project/laravel-spa/backend-app/public;

    charset UTF-8;

    # Error
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }

    # backend-app
    location / {
        index index.php index.html index.htm;
        try_files $uri /index.php?$query_string;
    }

    # frontend-app
    location /dashboard {
        alias /var/www/html/project/laravel-react-redux-blog-boilerplate/frontend-app;
        index index.html index.html;
        try_files $uri $uri/ /dashboard//index.html;
    }

    # php-fpm
    location ~ \.php$ {
        fastcgi_pass   unix:/var/run/php-fpm/php-fpm.sock;
        fastcgi_index  index.php;
        fastcgi_param  SCRIPT_FILENAME  $document_root/$fastcgi_script_name;
        include        fastcgi_params;
    }
}
```

※修正(2017/4/2追記)

spaでルート以下のURL（ex. laravel-app/dashboard/post）を直打ちまたはリロードすると404エラーが発生するので、修正しました。

修正前

```
 # frontend-app
 location /dashboard {
        alias /var/www/html/project/laravel-react-redux-blog-boilerplate/frontend-app;
        index index.html index.html;
    }
```

修正後

```
 # frontend-app
 location /dashboard {
        alias /var/www/html/project/laravel-react-redux-blog-boilerplate/frontend-app;
        index index.html index.html;
        try_files $uri $uri/ /dashboard//index.html;
    }
```

# 所感
オススメの構成や参考になりそうなリポジトリあったら教えてください〜

# 参考
+ [nginxのrootとalias](https://kinjouj.github.io/2013/01/nginx-root-alias.html)
+ [try_files, alias, and regexp locations](http://stackoverflow.com/questions/26356210/try-files-alias-and-regexp-locations)

