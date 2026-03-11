---
title: Handling CORS with Laravel
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
This post summarizes how to handle CORS (Cross-Origin Resource Sharing) in Laravel. The client-side uses React and axios. As a prerequisite, it's good to understand the types of CORS requests, the difference between simple request methods, and requests using preflight. For RESTful APIs, requests typically use preflight. This article covers examples of handling requests using preflight.

# Environment
Since it's CORS, it's obvious that we have separate domains for the API and web.

The setup involves domains like `api.hogehogedomain` and `admin.hogehogedomain`, where the admin calls an API managed on a different domain.

# Prepare Middleware
On the Laravel side, which provides the API, we prepare middleware to adjust header information during API requests.

Initially, I wanted to create custom middleware, but for some reason, it didn't work well with update methods, so I decided to use [barryvdh/laravel-cors](https://github.com/barryvdh/laravel-cors).

Setup follows the README instructions.

`composer require barryvdh/laravel-cors`

Specify the following in the provider array of `config/app.php`
`Barryvdh\Cors\ServiceProvider::class,`

Set the cors middleware in the api middleware group of `app/Http/Kernel.php`

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

`config/cors.php`

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

Since we want to allow cookie transmission and Basic authentication, set `supportsCredentials` to `true`.

That's all for the server-side settings.


# Calling the API from the Client Side

Define header information with axios.

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
On the client side, just set the `X-Requested-With` header, and then you can normally call the API.


# Thoughts
I haven't resolved why the custom middleware didn't work well, which is unsatisfactory, but for now, this should be fine.


# References
+ [Laravel 5.1 - easily enable CORS](http://en.vedovelli.com.br/2015/web-development/Laravel-5-1-enable-CORS/)

+ [cors - Building SPA with Laravel5 + AngularJS](http://qiita.com/fluke8259/items/c884bada22ccd286cf48)

+ [Organizing CORS (Cross-Origin Resource Sharing)
cors-004](http://dev.classmethod.jp/etc/about-cors/)

+ [CORS Summary](http://qiita.com/tomoyukilabs/items/81698edd5812ff6acb34)

+ [Laravel 5.2 CORS, GET not working with preflight OPTIONS](http://stackoverflow.com/questions/34748981/laravel-5-2-cors-get-not-working-with-preflight-options)

+ [X-Requested-With Header](http://boscono.hatenablog.com/entry/2013/12/23/152851)
