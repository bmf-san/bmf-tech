---
title: Laravel リポジトリパターンの実装
description: "Laravelでリポジトリパターンを実装し、インターフェース設計・DB操作の抽象化によるテスト容易性と保守性向上の手法を詳述。"
slug: implement-laravel-repository-pattern
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Laravel
  - リポジトリーパターン
translation_key: implement-laravel-repository-pattern
---


DB操作に関連するスマートな実装パターンであるリポジトリパターンについてかいてみようかと思います。

# リポジトリパターンとは
データの操作に関連するロジックをビジネスロジックから切り離し、抽象化したレイヤに任せることで保守や拡張性を高めるパターンです。
（必ずしもDB操作のロジックのみを留めるパターンというわけではないそうです。）

Laravelにリポジトリパターンを取り入れることで、

* テストがしやすくなる
* DBエンジンの変更に対応しやすくなる
* データ操作のロジックが1箇所にまとまり、管理しやすくなる

といったメリットを得ることができます。

# リポジトリパターンの実装
Modelと同じ単位でRepositoryディレクトリを作成します。(賛否両論あるかもです)

今回は以下のような構成でリポジトリパターンを実装していきます。

```:php
.
├── Models
│   ├── User.php
│  
├── Repositories
    └── User
        ├── UserRepository.php
        └── UserRepositoryInterface.php
```

# インターフェース設計
まずはインターフェースを設計します。

```php
<?php

namespace App\Repositories\User;

interface UserRepositoryInterface
{
    /**
     * Nameで1レコードを取得
     *
     * @var string $name
     * @return object
     */
    public function getFirstRecordByName($name);
}
```

# インプリメント（実装）クラス
続いて実装クラスを用意します。
ここでは対応するモデルのDIとメソッドの実装を行います。

```php
<?php

namespace App\Repositories\User;

use App\Models\User;

class UserRepository implements UserRepositoryInterface
{
    protected $user;

    /**
    * @param object $user
    */
    public function __construct(User $user)
    {
	$this->user = $user;
    }

    /**
     * 名前で1レコードを取得
     *
     * @var $name
     * @return object
     */
    public function getFirstRecordByName($name)
    {
        return $this->user->where('name', '=', $name)->first();
    }
}
```

ここから更にService層を用意してクラスを追加し、抽象度を高める場合もあるようですが、今回はこの2つのクラスのみで実装していくことにします。

# Service Provider
AppServiceProvider.phpにインターフェースと実装クラスを登録します。

```php
<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        //
    }

    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        // User
        $this->app->bind(
            \App\Repositories\User\UserRepositoryInterface::class,
            \App\Repositories\User\UserRepository::class
        );
    }
}
```

# Controllerで呼び出す

実装したリポジトリパターンを使用します。

```php
<?php

namespace App\Http\Controller\User;

use App\Repositories\User\UserRepositoryInterface;

class UserController extends Controller
{
   public function __construct(UserRepositoryInterface $user_repository)
   {
      $this->user_repository = $user_repository;
   }

   public function index()
   {
      return $this->user_repository->getFirstRecordByName($name);
   }
}
```

インターフェースをインジェクションするだけです！

# 所感
モデルもコントローラーもすっきりしました。
これを機にDDDの勉強もしたいです。

# 参考
* [Laravel4.2のリポジトリパターン](http://tech.aainc.co.jp/archives/10227)
* ~~【Laravel5.1 チュートリアル】中級者向けタスクリスト第4回~~
* [Laravelにおける後悔しないためのアプリケーション設計](https://speakerdeck.com/localdisk/laravelniokeruhou-hui-sinaitamefalseapurikesiyonshe-ji)

