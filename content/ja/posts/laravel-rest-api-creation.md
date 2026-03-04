---
title: "LaravelでRestAPIをつくる"
slug: "laravel-rest-api-creation"
date: 2017-09-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "API"
  - "Laravel"
  - "React"
  - "REST"
draft: false
---

LaravelでReactをフロントエンドのフレームワークとして採用しているのですが、APIを設計する必要がでてきたのでやってみました。

# やること
* ResourceControllerをつくってデータを出力するだけのAPIをつくる
* API公開を見越した認証用ミドルウェアを実装する

# やらないこと
* Restの説明
* APIデータの更新や削除
* Ajaxでのデータの取得と出力

# 環境
* Laravel5.2

# Resource Controllerをつくる
職人さんの朝は早い・・・
`php artisan make:controller HogeController --resource`


職人が仕事するとこんなコントローラーをつくってくれます。

```php
<?php

namespace App\\Http\\Controllers;

use Illuminate\\Http\\Request;

use App\\Http\\Requests;

class HogeController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \\Illuminate\\Http\\Response
     */
    public function index()
    {
        //
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \\Illuminate\\Http\\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \\Illuminate\\Http\\Request  $request
     * @return \\Illuminate\\Http\\Response
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \\Illuminate\\Http\\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \\Illuminate\\Http\\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \\Illuminate\\Http\\Request  $request
     * @param  int  $id
     * @return \\Illuminate\\Http\\Response
     */
    public function update(Request $request, $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \\Illuminate\\Http\\Response
     */
    public function destroy($id)
    {
        //
    }
}
```


ではAPIを早速つくります。index()のところをいじります。

```php
    /**
     * Display a listing of the resource.
     *
     * @return \\Illuminate\\Http\\Response
     */
    public function index()
    {
        $user = \\Auth::user();

        return \\Response::json($user);
    }
```

Responseでjsonをかえすだけです_(:3」∠)_


# ルーティング

```php:route.php
//-------------------------------
// API
//-------------------------------
Route::group(['prefix' => 'api'], function () {
    Route::resource('user', 'Resource\\UserAuthController');
});
```

※Laravel5.3からはrouteがディレクトリになってweb.phpとかapi.phpって感じにファイルが分かれていると思います。api.phpにかくのがベタだと思います。



/apiにアクセスするとjsonレスポンスが出力されていると思います。

# APIの認証について

apiを直接たたくようなスケベェな人を避けたい時や、APIを外部に公開したい時は認証を設けましょう。
ここではmiddlewareで認証を行う方法を例にあげたいと思います。


```php:route.php
Route::group(['middleware' => 'auth.user'], function () {
    Route::get('/userlist', 'UserList\\UserListController@getIndex');

    //-------------------------------
    // API
    //-------------------------------
    Route::group(['prefix' => 'api'], function () {
        Route::resource('user', 'Resource\\UserAuthController');
    });
});
```

AuthenticateOfApiとかいうAPI利用のためのミドルウェアを作ることにします。

一部[Laravelエキスパート養成読本](https://www.amazon.co.jp/Laravel%E3%82%A8%E3%82%AD%E3%82%B9%E3%83%91%E3%83%BC%E3%83%88%E9%A4%8A%E6%88%90%E8%AA%AD%E6%9C%AC-%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E9%96%8B%E7%99%BA%E3%82%92%E5%AE%9F%E7%8F%BE%E3%81%99%E3%82%8BPHP%E3%83%95%E3%83%AC%E3%83%BC%E3%83%A0%E3%83%AF%E3%83%BC%E3%82%AF%EF%BC%81-Software-Design-plus/dp/4774173134)を参考にさせて頂きました。


```php:AuthenticateOfApi.php
<?php

namespace App\Http\Middleware;

use App\Models\User;
use Closure;

class AuthenticateOfApi
{
    /**
     * @var string
     */
    const APPLICATION_TOKEN = 'x-application-token';

    /**
     * API Authenticate
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle($request, Closure $next)
    {
        // Cookie認証の場合
        if (真偽値を返すログインチェック) {
          // ユーザー情報を返したり、ほげほげ、、、
        }

        // API Token認証の場合
        if (ヘッダにx-application-tokenが含まれているか判定) {
          // ユーザー情報を返したり、ほげほげ、、、
        }

        if (ログインしていない、x-application-tokenもない) {
            return abort(401);
        }

        return $next($request);
    }
}
```

~~内部でAPIを利用したい時にAPITokenをヘッダに含めてしまうと認証の意味がなくなってしまいます。（ユーザーにヘッダが丸見えなのでトークンをパクられたら誰でもAPIが利用できてしまいます。）~~

~~従って内部でAPIを利用する際はログインと同じ認証方法をとると良いかと思います。
`Auth::guard('users')->check()`とかでログインチェック！~~

~~外部からのAPIの利用については、JavaScriptでヘッダにトークンを入れてPOSTすることで認証させることができます。~~

※APIの認証については他の記事をご参照ください。


# 所感
LaravelでAPIをつくる・使うのは簡単ですが、API設計とやらが中々奥深そうです。
自分でつくったAPIを自分で使うというのは結構楽しいのでちょっと頑張ってみます。


# 参考
* [Laravel5.2で認証を用いたAPI作り](http://satobukuro.net/173/)
* [LaravelでRest APIを開発する](http://dim5.net/laravel/developing-rest-api.html)
* [React.jsでLaravelから情報をもらってみよう](http://blog.comnect.jp.net/blog/98)・・・Laravel+API+React！
* [Laravelエキスパート養成読本](https://www.amazon.co.jp/Laravel%E3%82%A8%E3%82%AD%E3%82%B9%E3%83%91%E3%83%BC%E3%83%88%E9%A4%8A%E6%88%90%E8%AA%AD%E6%9C%AC-%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E9%96%8B%E7%99%BA%E3%82%92%E5%AE%9F%E7%8F%BE%E3%81%99%E3%82%8BPHP%E3%83%95%E3%83%AC%E3%83%BC%E3%83%A0%E3%83%AF%E3%83%BC%E3%82%AF%EF%BC%81-Software-Design-plus/dp/4774173134)

