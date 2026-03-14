---
title: Directory Structure and Nginx Config for Integrating SPA with Laravel
description: 'Configure Laravel for single-page apps with proper directory structure and Nginx routing for frontend and backend.'
slug: laravel-spa-directory-structure-nginx-config
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Laravel
  - React
  - Nginx
translation_key: laravel-spa-directory-structure-nginx-config
---

# Overview
This is a brief note on modifying the directory structure of Laravel and the Nginx config file when implementing a SPA for the admin panel of an application built with Laravel. Since it was my first attempt, I decided to document it.

# Laravel Directory Structure
I divided the directory into two main parts: `backend-app` for the backend and `frontend-app` for the frontend. The `backend-app` handles user-facing screens, APIs, and backend processing, while the frontend is responsible for the SPA admin panel. Although the user-facing screens might also fall under `frontend-app`, I will address that later.
It might be better to move away from a full-stack framework if we are going to separate things like this... For now, I structured it in a way that makes it easier to manage the frontend and backend separately.

This structure is based on common setups I found while searching for "Laravel SPA" or "Laravel React" on GitHub.

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

# Nginx Config File
I separated the `location` directives for `backend-app` and `frontend-app`. If I were to separate them using the `server` directive, it would involve creating subdomains, which I thought would be awkward for this application, so I avoided that. I feel like I need to work harder on the configuration, but please bear with me (*_*).

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

*Note (Updated on 2017/4/2)*

I fixed a 404 error that occurred when directly accessing or reloading URLs under the root (e.g., laravel-app/dashboard/post) in the SPA.

Before the fix:
```
 # frontend-app
 location /dashboard {
        alias /var/www/html/project/laravel-react-redux-blog-boilerplate/frontend-app;
        index index.html index.html;
    }
```

After the fix:
```
 # frontend-app
 location /dashboard {
        alias /var/www/html/project/laravel-react-redux-blog-boilerplate/frontend-app;
        index index.html index.html;
        try_files $uri $uri/ /dashboard//index.html;
    }
```

# Thoughts
If you have any recommended structures or repositories that might be helpful, please let me know!

# References
+ [Nginx root and alias](https://kinjouj.github.io/2013/01/nginx-root-alias.html)
+ [try_files, alias, and regexp locations](http://stackoverflow.com/questions/26356210/try-files-alias-and-regexp-locations)