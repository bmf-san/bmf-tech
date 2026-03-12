---
title: Notes on Starting Functional Testing in Laravel
description: Research notes and a structured overview of Notes on Starting Functional Testing in Laravel, summarizing key concepts and findings.
slug: getting-started-with-laravel-feature-testing
date: 2019-02-11T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - Laravel
  - Functional Testing
translation_key: getting-started-with-laravel-feature-testing
---

# Overview
This post introduces how to start functional testing in Laravel and some basic usage. It is limited to an introductory level, so more practical content is not covered. ※ This is just a memo-level content, like a source for a lightning talk.

# Target Audience
This is for people who have never written tests before.

Even if you have never written tests, as long as you understand the application's specifications, functional testing should be relatively easy for anyone to understand what to write.

Especially in Laravel, there are plenty of convenient APIs and tools available for functional testing, so even if you're not familiar with testing, it should be easy to get started.

# Environment
- Docker
- Laravel 5.7
- MySQL

# Preparation
I have roughly prepared the environment.

[github - bmf-san/laravel-test-handson](https://github.com/bmf-san/laravel-test-handson)

By following the commands in the README, you can set up a Laravel environment on Docker.

As preparation for an environment where tests can be executed, I prepared a test db and modified phpunit.xml and config/database.php on the application side.

To use the convenient API for functional testing, I have introduced [github - laravel/browser-kit-testing](https://github.com/laravel/browser-kit-testing#interacting-with-links).

This was originally included in the Laravel core until version 5.3, but it seems to have become a separate package around version 5.4.

Since I hadn't caught up with Laravel since version 5.3, I realized this after touching version 5.7...

It's not a big deal, but I can no longer say "You can start immediately without introducing a separate package!"

# Starting the Test
The actual code is placed here, so I'll introduce a part of it.

[github - laravel/browser-kit-testing](https://github.com/laravel/browser-kit-testing#interacting-with-links)

Suppose there is a screen to create a new article.

```html
@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
          <form method="POST" action="/post/store">
              @csrf

              <div class="form-group row">
                  <label for="name" class="col-md-4 col-form-label text-md-right">Title</label>

                  <div class="col-md-6">
                      <input id="title" type="text" class="form-control{{ $errors->has('title') ? ' is-invalid' : '' }}" name="title" value="{{ old('title') }}" required autofocus>

                      @if ($errors->has('title'))
                          <span class="invalid-feedback" role="alert">
                              <strong>{{ $errors->first('title') }}</strong>
                          </span>
                      @endif
                  </div>
              </div>
                  
              <div class="form-group row">
                  <label for="name" class="col-md-4 col-form-label text-md-right">Body</label>
                  
                  <div class="col-md-6">
                      <textarea class="form-control{{ $errors->has('body') ? ' is-invalid' : '' }}" id="body" name="body" required>{{ old('body') }}</textarea>
                      @if ($errors->has('body'))
                          <span class="invalid-feedback" role="alert">
                              <strong>{{ $errors->first('body') }}</strong>
                          </span>
                      @endif
                  </div>
              </div>

              <div class="form-group row mb-0">
                  <div class="col-md-6 offset-md-4">
                      <button type="submit" class="btn btn-primary">
                        Submit
                      </button>
                  </div>
              </div>
          </form>
        </div>
    </div>
</div>
@endsection
```

The routing looks like this.

```php
Route::get('post/create', 'PostController@create');
Route::post('post/store', 'PostController@store');
```

Then you can write intuitive functional tests like this.

```php
<?php

use Tests\TestCase;
use Illuminate\Foundation\Testing\WithoutMiddleware;
use Illuminate\Foundation\Testing\DatabaseMigrations;
use Illuminate\Foundation\Testing\DatabaseTransactions;
use App\User;
use App\Post;

class PostTest extends TestCase
{
    use DatabaseMigrations, DatabaseTransactions;
  
    public function testCreatePost()
    {
      $user = factory(User::class)->create();
      
      $this->actingAs($user);
      $this->visit("/post/create");
      $this->type("title", "title");
      $this->type("body", "body");
      $this->press("Submit");
      $this->seePageIs("/post");
    }
}
```

# Impressions
This is just a rough memo for a lightning talk.

# References
- [Laravel API - Illuminate\Foundation\Testing](https://laravel.com/api/5.7/Illuminate/Foundation/Testing.html)
- [github - laravel/browser-kit-testing](https://github.com/laravel/browser-kit-testing#interacting-with-links)
