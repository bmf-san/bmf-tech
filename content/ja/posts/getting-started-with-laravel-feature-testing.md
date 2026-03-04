---
title: "Laravelでの機能テストの始め方雑メモ"
slug: "getting-started-with-laravel-feature-testing"
date: 2019-02-11
author: bmf-san
categories:
  - "テスト"
tags:
  - "Laravel"
  - "機能テスト"
draft: false
---

# 概要
Laravelでの機能テストの始め方と簡単な使い方について紹介する。
入門レベルに限るのでより実践的な内容については触れない。
※LTの元ネタ程度でメモくらいの内容。

# 対象
テストを書いたことがない人向け。

テストを書いたことがなくても機能テストであればアプリケーションの仕様さえわかっていれば比較的に誰にでも楽に何を書くのかわかりやすいと思う。

特にLaravelは機能テストで使える便利なAPIやツールが充実しているので、テストに慣れていなくともテストに取り組みやすいはず。

# 環境
- Docker
- Laravel5.7
- MySQL

# 準備
雑に環境を用意しておいた。

[github - bmf-san/laravel-test-handson](https://github.com/bmf-san/laravel-test-handson)

READMEの手順どおりコマンドを実行すればDocker上でLaravelの環境がセットアップできる。

テストを実行できる環境の準備として、テスト用のdbを用意し、アプリケーション側ではphpunit.xml、config/database.phpをいじった。

機能テストで使う便利なAPIな使いたいので[github - laravel/browser-kit-testing](https://github.com/laravel/browser-kit-testing#interacting-with-links)を導入している。

こちらはLaravel5.3まではLaravel本体に組み込まれていたものだったが、5.4くらいから別パッケージになってしまったらしい。

Laravel5.3以降からキャッチアップしていなかったので、今回5.7を触ってから気づいた・・・

大した問題ではないが、別パッケージを導入しなくてもすぐに始められる！という売り文句がいえなくなってしまった、、

# テストを始める
実際のコードはここに置いてあるので一部を紹介する。　

[github - laravel/browser-kit-testing](https://github.com/laravel/browser-kit-testing#interacting-with-links)

記事を新規作成する画面があったとする。

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

ルーティングがこんな感じ。

```php
Route::get('post/create', 'PostController@create');
Route::post('post/store', 'PostController@store');
```

するとこんな感じの直感的な機能テストがかける。

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

# 所感
LTのネタメモなのでものすごい雑。

# 参考
- [Laravel API - Illuminate\Foundation\Testing](https://laravel.com/api/5.7/Illuminate/Foundation/Testing.html)
- [github - laravel/browser-kit-testing](https://github.com/laravel/browser-kit-testing#interacting-with-links)

