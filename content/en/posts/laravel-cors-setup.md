---
title: CORS Support in Laravel
slug: laravel-cors-setup
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
  - CORS
translation_key: laravel-cors-setup
---

# Overview
This article summarizes how to enable CORS (Cross-Origin Resource Sharing) in Laravel. The client side uses React and axios. As a prerequisite, it's good to understand the different types of CORS request formats, including simple requests and those that use preflight. For RESTful APIs, the request format typically involves preflight. This article will cover an example that supports requests using preflight.

# Environment
Since it's CORS, it's obvious that we have separate domains for the API and web.

Domains are set up like `api.hogehogedomain` and `admin.hogehogedomain`, where the admin calls the API managed on a different domain.

# Prepare Middleware
On the API side, we will prepare middleware to adjust header information during API requests in Laravel.

I wanted to create custom middleware here, but for some reason, only the update methods didn't work properly, so I decided to use [barryvdh/laravel-cors](https://github.com/barryvdh/laravel-cors).

The setup is as per the README.

`composer require barryvdh/laravel-cors`

Add the following to the provider array in `config/app.php`:
`Barryvdh\Cors\ServiceProvider::class,`

Set the cors middleware in the api middleware group in `app/Http/Kernel.php`:

```
    protected $middlewareGroups = [
        'web' => [
            \App\Http\Middleware\EncryptCookies::class,
            \Illuminate\Cookie\Middleware\AddQueuedCookiesToResponse::class,
            \Illuminate\Session\Middleware\StartSession::class,
            \Illuminate\View\Middleware\ShareErrorsFromSession::class,
            \App\Http\Middleware\VerifyCsrfToken::class,
            \Illuminate\Routing\Middleware\SubstituteBindings::class,
        ],

        'api' => [
            'throttle:60,1',
            'bindings',
            \Barryvdh\Cors\HandleCors::class, // <-Here!
        ],
    ];
```

Publish and edit the configuration file.
`php artisan vendor:publish --provider="Barryvdh\Cors\ServiceProvider"`

In `config/cors.php`:

```
return [
     /*
     |--------------------------------------------------------------------------
     | Laravel CORS
     |--------------------------------------------------------------------------
     |
     | allowedOrigins, allowedHeaders and allowedMethods can be set to array('*')
     | to accept any value.
     |
     */
    'supportsCredentials' => true, // change false to true !
    'allowedOrigins' => ['*'],
    'allowedHeaders' => ['Content-Type', 'X-Requested-With'],
    'allowedMethods' => ['*'], // ex: ['GET', 'POST', 'PUT',  'DELETE']
    'exposedHeaders' => [],
    'maxAge' => 0,
]
```

Since we want to allow sending cookies and Basic authentication, we set `supportsCredentials` to `true`.

That's all for the server-side configuration.

# Calling the API from the Client Side

Define the header information with axios.

action/index.js

```
const api = axios.create({
  baseURL: 'http://api.rubel/v1',
  timeout: 10000,
  headers: {
    'X-Requested-With': 'XMLHttpRequest'
  }
});

export function createCategory(props) {
  const request = api.post(`/categories`, props);

  return {type: CREATE_CATEGORY, payload: request};
}
```
On the client side, you only need to set the `X-Requested-With` header, and then you can simply call the API as usual.

# Thoughts
I haven't resolved why my custom middleware didn't work, so I'm left with some dissatisfaction, but for now, this should be fine.

# References
+ [Laravel 5.1 - easily enable CORS](http://en.vedovelli.com.br/2015/web-development/Laravel-5-1-enable-CORS/)
+ [cors - SPA with Laravel5 + AngularJS](http://qiita.com/fluke8259/items/c884bada22ccd286cf48)
+ [Understanding CORS (Cross-Origin Resource Sharing)](http://dev.classmethod.jp/etc/about-cors/)
+ [CORS Summary](http://qiita.com/tomoyukilabs/items/81698edd5812ff6acb34)
+ [Laravel 5.2 CORS, GET not working with preflight OPTIONS](http://stackoverflow.com/questions/34748981/laravel-5-2-cors-get-not-working-with-preflight-options)
+ [X-Requested-With Header](http://boscono.hatenablog.com/entry/2013/12/23/152851)