---
title: Creating a REST API with Laravel
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

I am using React as the front-end framework with Laravel, and I needed to design an API, so I decided to give it a try.

# What to Do
* Create a ResourceController to output data via an API
* Implement authentication middleware in anticipation of API exposure

# What Not to Do
* Explanation of REST
* Updating or deleting API data
* Fetching and outputting data via Ajax

# Environment
* Laravel 5.2

# Creating a Resource Controller
The craftsman's morning starts early...
`php artisan make:controller HogeController --resource`

The craftsman will create a controller like this:

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

Now, let's quickly create the API. We will modify the index() method.

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

We simply return the user as JSON using Response _(:3」∠)_.

# Routing

```php:route.php
//-------------------------------
// API
//-------------------------------
Route::group(['prefix' => 'api'], function () {
    Route::resource('user', 'Resource\\UserAuthController');
});
```

*Note: From Laravel 5.3, routes are separated into directories like web.php and api.php. It is best to write in api.php.*

When you access /api, you should see a JSON response output.

# API Authentication

When you want to avoid curious individuals directly hitting the API or when you want to expose the API externally, let's implement authentication. Here, I will provide an example of how to perform authentication using middleware.

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

I will create a middleware called AuthenticateOfApi for API usage.

I referred to the [Laravel Expert Training Book](https://www.amazon.co.jp/Laravel%E3%82%A8%E3%82%AD%E3%82%B9%E3%83%9A%E3%83%BC%E3%83%88%E9%A4%8A%E6%88%90%E8%AA%AD%E6%9C%AC-%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E9%96%8B%E7%99%BA%E3%82%92%E5%AE%9F%E7%8F%BE%E3%81%99%E3%82%8BPHP%E3%83%95%E3%83%AC%E3%83%BC%E3%83%A0%E3%83%AF%E3%83%BC%E3%82%AF%EF%BC%81-Software-Design-plus/dp/4774173134) for this.

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
        // For cookie authentication
        if (check login returning boolean) {
          // Return user information, etc...
        }

        // For API Token authentication
        if (check if header contains x-application-token) {
          // Return user information, etc...
        }

        if (not logged in and no x-application-token) {
            return abort(401);
        }

        return $next($request);
    }
}
```

~~Including the API Token in the header when using the API internally defeats the purpose of authentication (since the token is visible to the user, anyone can use the API if they steal it).~~

~~Therefore, when using the API internally, it is advisable to use the same authentication method as logging in. You can check login with something like `Auth::guard('users')->check()`!~~

~~For external API usage, you can authenticate by including the token in the header and sending a POST request via JavaScript.~~

*For more information on API authentication, please refer to other articles.*

# Thoughts
Creating and using APIs with Laravel is easy, but API design seems quite deep. It is quite enjoyable to create an API and use it myself, so I will try my best.

# References
* [Creating an API with Authentication in Laravel 5.2](http://satobukuro.net/173/)
* [Developing REST API with Laravel](http://dim5.net/laravel/developing-rest-api.html)
* [Getting Information from Laravel with React.js](http://blog.comnect.jp.net/blog/98) ... Laravel + API + React!
* [Laravel Expert Training Book](https://www.amazon.co.jp/Laravel%E3%82%A8%E3%82%AD%E3%82%B9%E3%83%9A%E3%83%BC%E3%83%88%E9%A4%8A%E6%88%90%E8%AA%AD%E6%9C%AC-%E3%83%A2%E3%83%80%E3%83%B3%E3%81%AA%E9%96%8B%E7%99%BA%E3%82%92%E5%AE%9F%E7%8F%BE%E3%81%99%E3%82%8BPHP%E3%83%95%E3%83%AC%E3%83%BC%E3%83%A0%E3%83%AF%E3%83%BC%E3%82%AF%EF%BC%81-Software-Design-plus/dp/4774173134)