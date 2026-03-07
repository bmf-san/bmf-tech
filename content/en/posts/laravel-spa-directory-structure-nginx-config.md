---
title: Directory Structure and Nginx Configuration for Integrating SPA into Laravel
slug: laravel-spa-directory-structure-nginx-config
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Laravel
  - React
  - Nginx
description: A discussion on directory structure and nginx configuration when implementing an SPA for the admin panel in a Laravel application.
translation_key: laravel-spa-directory-structure-nginx-config
---

# Overview
This post discusses the directory structure and nginx configuration adjustments made when implementing an SPA (Single Page Application) for the admin panel in a Laravel application. It was my first attempt, so I summarized my notes here.

# Laravel Directory Structure
The application was divided into two main directories: `backend-app` for backend-related tasks and `frontend-app` for frontend-related tasks. 

- `backend-app` handles user-facing screens, APIs, and backend processing.
- `frontend-app` is responsible for the SPA-based admin panel.

Although user-facing screens could also fall under `frontend-app`, I decided to separate them for now and plan to refactor later. If splitting them further, moving away from a full-stack framework might be worth considering. For now, I aimed for a structure that makes managing frontend and backend easier.

This structure was inspired by common setups found on GitHub when searching for "Laravel SPA" or "Laravel React".

```
.
├── backend-app
│   ├── app
│   ├── bootstrap
│   ├── config
│   ├── database
│   ├── node_modules
│   ├── public
│   ├── resources
│   ├── routes
│   ├── storage
│   ├── tests
│   └── vendor
└── frontend-app
    ├── _components
    ├── dist
    ├── node_modules
    └── src
```

# Nginx Configuration File
The `location` directive was separated for `backend-app` and `frontend-app`. While separating them via `server` directives and using subdomains is another option, I felt it wasn’t suitable for this application. The configuration could be improved further, but this setup works for now.

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

### Fix (Added on 2017/4/2)
When directly accessing or reloading URLs under the SPA root (e.g., `laravel-app/dashboard/post`), a 404 error occurred. This was fixed as follows:

Before:

```
 # frontend-app
 location /dashboard {
        alias /var/www/html/project/laravel-react-redux-blog-boilerplate/frontend-app;
        index index.html index.html;
    }
```

After:

```
 # frontend-app
 location /dashboard {
        alias /var/www/html/project/laravel-react-redux-blog-boilerplate/frontend-app;
        index index.html index.html;
        try_files $uri $uri/ /dashboard//index.html;
    }
```

# Thoughts
If you have recommended structures or repositories that might be helpful, please share them!

# References
+ [nginx root and alias](https://kinjouj.github.io/2013/01/nginx-root-alias.html)
+ [try_files, alias, and regexp locations](http://stackoverflow.com/questions/26356210/try-files-alias-and-regexp-locations)