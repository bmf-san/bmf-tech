---
title: Creating a Rest API with Laravel
description: 'Create Laravel REST APIs using ResourceControllers with authentication middleware and JSON responses.'
slug: laravel-rest-api-creation
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - API
  - Laravel
  - React
  - REST
translation_key: laravel-rest-api-creation
---



We are using React as the frontend framework with Laravel, and since we needed to design an API, we decided to give it a try.

# What We'll Do
* Create a ResourceController to build a simple API that outputs data
* Implement authentication middleware in anticipation of public API exposure

# What We Won't Do
* Explain REST
* Update or delete API data
* Fetch and output data using Ajax

# Environment
* Laravel 5.2

# Creating a Resource Controller
The artisan's morning starts early...
`php artisan make:controller HogeController --resource`

When the artisan gets to work, they create a controller like this.

```php
<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Requests;

class HogeController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        //
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        //
    }
}
```

Let's quickly create the API. We'll modify the index() method.

```php
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        $user = \Auth::user();

        return \Response::json($user);
    }
```

Just return JSON with Response _(:3」∠)_

# Routing

```php:route.php
//-------------------------------
// API
//-------------------------------
Route::group(['prefix' => 'api'], function () {
    Route::resource('user', 'Resource\UserAuthController');
});
```

*Note: From Laravel 5.3, routes are organized into directories like web.php and api.php. It's best to write in api.php.*

Accessing /api should output a JSON response.

# About API Authentication

When you want to avoid people directly hitting the API or want to expose the API externally, it's good to set up authentication. Here, we'll provide an example of using middleware for authentication.

```php:route.php
Route::group(['middleware' => 'auth.user'], function () {
    Route::get('/userlist', 'UserList\UserListController@getIndex');

    //-------------------------------
    // API
    //-------------------------------
    Route::group(['prefix' => 'api'], function () {
        Route::resource('user', 'Resource\UserAuthController');
    });
});
```

We will create middleware called AuthenticateOfApi for API usage.

Some parts were referenced from [Laravel Expert Training Book](https://www.amazon.co.jp/Laravel%E3%82%A8%E3%82%AD%E3%82%B9%E3%83%91%E3%83%BC%E3%83%88%E9%A4%8A%E6%88%90%E8%AA%AD%E6%9C%AC-%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E9%96%8B%E7%99%BA%E3%82%92%E5%AE%9F%E7%8F%BE%E3%81%99%E3%82%8BPHP%E3%83%95%E3%83%AC%E3%83%BC%E3%83%A0%E3%83%AF%E3%83%BC%E3%82%AF%EF%BC%81-Software-Design-plus/dp/4774173134).

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
        // Cookie authentication
        if (login check returning boolean) {
          // Return user info, etc.
        }

        // API Token authentication
        if (check if header contains x-application-token) {
          // Return user info, etc.
        }

        if (not logged in and no x-application-token) {
            return abort(401);
        }

        return $next($request);
    }
}
```

~~Including the APIToken in the header when using the API internally nullifies the purpose of authentication (since the header is visible to users, anyone can use the API if the token is stolen).~~

~~Therefore, when using the API internally, it's better to use the same authentication method as login.
`Auth::guard('users')->check()` for login checks!~~

~~For external API usage, you can authenticate by including the token in the header with JavaScript and POSTing it.~~

*Please refer to other articles for more on API authentication.*

# Thoughts
Creating and using APIs with Laravel is easy, but API design seems quite deep. It's quite fun to use the API you created yourself, so I'll try to put in some effort.

# References
* [Creating an API with Authentication in Laravel 5.2](http://web.archive.org/web/20210417022405/https://satobukuro.net/173/)
* ~~Developing a Rest API with Laravel~~
* [Getting Information from Laravel with React.js](http://web.archive.org/web/20200130035136/http://blog.comnect.jp.net:80/blog/98) ... Laravel+API+React!
* [Laravel Expert Training Book](https://www.amazon.co.jp/Laravel%E3%82%A8%E3%82%AD%E3%82%B9%E3%83%91%E3%83%BC%E3%83%88%E9%A4%8A%E6%88%90%E8%AA%AD%E6%9C%AC-%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E9%96%8B%E7%99%BA%E3%82%92%E5%AE%9F%E7%8F%BE%E3%81%99%E3%82%8BPHP%E3%83%95%E3%83%AC%E3%83%BC%E3%83%A0%E3%83%AF%E3%83%BC%E3%82%AF%EF%BC%81-Software-Design-plus/dp/4774173134)
