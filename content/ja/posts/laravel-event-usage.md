---
title: "LaravelでEventを使う"
slug: "laravel-event-usage"
date: 2017-09-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Laravel"
draft: false
---

ユーザー登録時や退会時など特定のイベントの時に発火させたいメソッドを管理したい時はイベントリスナーを使うと便利です。

今回は基本的なイベントとリスナーの定義の仕方についてはすっ飛ばし、一つのリスナークラスで複数のイベントを設定できる**イベント購読**について扱います。

## 環境
* laravel5.2

## ディレクトリ
* app\Events・・・イベント名＝クラス名としたクラスを置く（厳密な命名規則はないです）
* app\Listeners・・・イベントごとの処理（リスナー）とsubscribeメソッドの実装（後述）したクラスを置く
* app\Providers・・・イベント購読で使用するリスナーを登録するクラスを置く
* app\Controllers・・・Eventを呼び出すコントローラーおよびメソッドを用意しておく

## Eventを定義する
まずはイベントを定義しましょう。
例として“ユーザー登録完了時”というイベントを作成することにします。
イベントクラスはartisanコマンドで生成することができます。

`php artisan make:event UserRegistrationComplete`

```php

<?php

namespace App\Events;

use App\Events\Event;
use Illuminate\Queue\SerializesModels;
use Illuminate\Contracts\Broadcasting\ShouldBroadcast;

class UserRegistrationComplete extends Event
{
    use SerializesModels;

    /**
     * Create a new event instance.
     *
     * @return void
     */
    public function __construct()
    {
        //
    }

    /**
     * Get the channels the event should be broadcast on.
     *
     * @return array
     */
    public function broadcastOn()
    {
        return [];
    }
}

```

artisanコマンドを実行するとこのようなクラスが自動生成されます。
コンストラクタの中にはイベントで使用するデータを設定しておきましょう。

今回はユーザー登録に関わるイベントなので、Userモデルを呼び出しておくことにします。ここで呼び出したデータはリスナークラスで使用することができます。

```php

    /**
     * Create a new event instance.
     *
     * @return void
     */
    public function __construct(User $user)
    {
        $this->user =  $user;
    }

```

ちなみにbroadcastOnというのは、リアルタイムに実行したいユーザーインターフェースを実装したい時に使用するものなので今回はスルーです。

## Listenerを定義する
次にリスナーを定義しましょう。
リスナーもartisanコマンドで生成することができます。

`php artisan make:listener UserAuthEventListenerListener --event UserRegistrationComplete`

リスナーを生成するときは、eventオプションで結び付けたいイベントを設定することができます。

```php
<?php

namespace App\Listeners;

use App\Events\UserRegistrationComplete;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Contracts\Queue\ShouldQueue;

class UserAuthEventListenerListener
{
    /**
     * Create the event listener.
     *
     * @return void
     */
    public function __construct()
    {
        //
    }

    /**
     * Handle the event.
     *
     * @param  UserRegistrationComplete  $event
     * @return void
     */
    public function handle(UserRegistrationComplete $event)
    {
        //
    }
}
```


イベント購読は、一つのリスナークラスに複数のイベントを設定できるので、リスナー名は複数あるイベントのカテゴリー名のような形で命名すると良いでしょう。

今回の例でいくと、“ユーザー登録完了イベントはユーザー認証のグループに属す”といったところでしょうか。


生成したリスナーで**イベント発火時の処理**と**イベント購読の登録**を行います。

```php
<?php

namespace App\Listeners;

class UserAuthEventListener
{
    // イベント発火時の処理
    public function onConfirm($event)
    {
      // 処理
    }

　　　　　　　　// 複数追加できます
　　　　　　　　public function onHogeHoge($event)
    {
      // 処理
    }


    // イベント購読の登録
    public function subscribe($events)
    {
        $events->listen(
            'App\Events\UserRegistrationComplete',
            'App\Listeners\UserAuthEventListener@onConfirm'
        );
    }

    // 複数登録できます
    public function subscribe($events)
    {
        $events->listen(
            'App\Events\UserHogeHoge',
            'App\Listeners\UserAuthEventListener@onHogeHoge'
        );
    }
}

```

イベント追加したい場合は、先ほどのartisanコマンドでイベントを生成してやればOKです。


## EventServiceProviderにイベント購読クラスを登録
登録ばっかりでややこしいかと思いますが、これで最後です。
app\Providersにデフォルトで存在しているEventServiceProviderを使用します。

サービスプロバイダというのはアプリケーションの初期起動処理を行うクラスです。詳しくは[ドキュメント](https://readouble.com/laravel/5/1/ja/providers.html)


```
<?php

namespace App\Providers;

use Illuminate\Foundation\Support\Providers\EventServiceProvider as ServiceProvider;

class EventServiceProvider extends ServiceProvider
{
    // ここにリスナーを登録していく
    protected $subscribe = [
        'App\Listeners\UserAuthEventListener',
        'App\Listeners\HogeHogeListener',
    ];
}

```

## Fire

さて、イベント購読の登録が完了したら後は自由にイベントを発火させましょう。

イベントを発火させるには、Eventファサードとfireメソッドを使います。

fireメソッドにはイベントのインスタンスを渡してあげます。

```php
    /**
     * コントローラーのとあるメソッド
     */
    private function hogehoge(User $user)
    {
        // 登録確認イベントの発火
        \Event::fire(new UserAuthRegistrationComplete($user));
    }

```

こんな感じで登録したイベントが火を吹きます :fire:


## まとめ

イベント定義（データ保持）→リスナー定義（イベントで使用するメソッド管理とsubscribeメソッド実装）→イベント購読登録→:fire:


## 追記
Laravel5.3からNotificationというパッケージが導入されたので今回のような通知の管理はNotificationを利用した方が楽かもしれません。
そういった意図があるのか、5.3ではEventやListener, Jobとったディレクトリがデフォルトでは存在しなくなりました。
気になる方はLaravelのリポジトリやドキュメントをご覧ください。

## #参考
* [Laravel5.1のEventの使い方についてまとめてみた。](http://blog.fagai.net/2015/12/11/laravel51-event/)
* [Laravel の Event クラスを使って処理をまとめてみる](http://localdisk.hatenablog.com/entry/2014/03/26/Laravel_%E3%81%AE_Event_%E3%82%AF%E3%83%A9%E3%82%B9%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E5%87%A6%E7%90%86%E3%82%92%E3%81%BE%E3%81%A8%E3%82%81%E3%81%A6%E3%81%BF%E3%82%8B)
* [Laravelのイベント](https://kore1server.com/292)

