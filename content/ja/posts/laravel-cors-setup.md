---
title: "LaravelでCORS対応"
slug: "laravel-cors-setup"
date: 2017-09-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Laravel"
  - "CORS"
draft: false
---

# 概要
LaravelでCORS(Cross-Origin Resource Sharing)に対応する方法をまとめました。
クライアントサイドはReact, axiosを使用します。
前提知識としては、CORSのリクエスト形態、シンプルなリクエスト方法と、preflightを使用するリクエスト方法の違いを抑えておけば良いかと思います。
RESTfulAPIの場合は基本的にはpreflightを使用するリクエスト形式かと思います。
この記事では、preflightを使用するリクエストに対応する例を取り上げます。

# 環境
CORSなので当たり前ですが、apiとwebでドメインを用意しています。

`api.hogehogedomain`と`admin.hogehogedomain`みたいな感じでドメインが用意されていて、adminの方から別ドメインで管理されているAPIをコールする、といった感じです。

# Middlewareを用意する
APIを提供している側、Laravel側でAPIリクエスト時にヘッダ情報を調整するmiddlewareを用意します。

と、ここで自作のミドルウェアを作成したかったのですが、何故か更新系のメソッドだけ上手く動作しなかったので、[barryvdh/laravel-cors](https://github.com/barryvdh/laravel-cors)を使うことにします。

セットアップはREADME通りです。

`composer require barryvdh/laravel-cors`

`config/app.php`のprovider配列に以下を指定
`Barryvdh\Cors\ServiceProvider::class,`

`app/Http/Kernel.php`のapiミドルウェアグループにcorsミドルウェアを設定

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

設定ファイルをパブリッシュして編集。
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

クッキーの送信およびBasic認証の許可しておきたいので、　`supportsCredentials`を`true`にしておきます。

サーバー側の設定は以上です。


# クライアントサイドからAPIをコールしてみる

axiosでヘッダ情報の定義をしておきます。

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
クライアント側では、`X-Requested-With`ヘッダをセットするだけで、後は普通にapiを叩くだけです。


# 所感
自作ミドルウェアがなぜ上手くいかなかったのか解決できていないので消化不良ですが、一旦はこれで問題ないでしょう。。


# 参考
+ [Laravel 5.1 - easily enable CORS](http://en.vedovelli.com.br/2015/web-development/Laravel-5-1-enable-CORS/)

+ [cors - Laravel5 + AngularJS で作るSPA](http://qiita.com/fluke8259/items/c884bada22ccd286cf48)

+ [CORS(Cross-Origin Resource Sharing)について整理してみた
cors-004](http://dev.classmethod.jp/etc/about-cors/)

+ [CORSまとめ](http://qiita.com/tomoyukilabs/items/81698edd5812ff6acb34)

+ [Laravel 5.2 CORS, GET not working with preflight OPTIONS](http://stackoverflow.com/questions/34748981/laravel-5-2-cors-get-not-working-with-preflight-options)

+ [X-Requested-Withヘッダ](http://boscono.hatenablog.com/entry/2013/12/23/152851)

