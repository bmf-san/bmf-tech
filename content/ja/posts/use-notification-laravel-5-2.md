---
title: Laravel5.2にNotificationでNotificationを使う
description: Laravel5.2にNotificationでNotificationを使うについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: use-notification-laravel-5-2
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Laravel
translation_key: use-notification-laravel-5-2
---


Laravel5.3がリリースされましたが、あえてLaravel5.2でNotificationを使う話です。



# 環境
* Laravel5.2（または5.1）


# Notificationをインストール
Notificationとはナンゾヤ？

[Laravel.com - Notificaions](https://laravel.com/docs/5.3/notifications)
[Laravel Notification Channel](http://laravel-notification-channels.com/)

Laravel5.3以前では通知の管理などはEventとListenerを活用することで何とかゴニョゴニョやっていたかと思います。

それがNotificationによって、より便利に通知管理を行えるようになりました。
※名前の通り、通知に特化したものなのでEventやListenerの"代用"というわけではないかと思います。


ではでは、インストールします。
`composer require laravel-notification-channels/backport`

Laravel5.3の場合はデフォルトで組み込まれているため、composerでわざわざインストールする必要はありませんが、〜5.2はマニュアルです。

なのでfacadeを使いたい場合はfacadeのクラスを追加してあげます。

※FacadeでNotification使わないよーという場合は飛ばしてください。

vendor/laravel/framework/src/Illuminate/Support/FacedesにNotification.phpを作成して、以下を記述。

```php
<?php

namespace Illuminate\Support\Facades;

/**
 * @see \Illuminate\Notifications\ChannelManager
 */
class Notification extends Facade
{
    /**
     * Get the registered name of the component.
     *
     * @return string
     */
    protected static function getFacadeAccessor()
    {
        return 'Illuminate\Notifications\ChannelManager';
    }
}
```

それから、app.phpのprovidersとaliasに以下をそれぞれ追記。

`Illuminate\Notifications\NotificationServiceProvider::class`
`'Notification' => Illuminate\Support\Facades\Notification::class`


ここまでで下準備OKです。


# 使ってみる
まずは通知クラスを職人さんで作成しましょう。
例としてユーザーが登録されたときの通知を管理するクラスを作成することにします。

`php aritsan make:notification UserRegistered`

Notificationディレクトリが自動生成されて、その中にUserRegistered.phpが生成されていると思います。


生成を確認したら、通知の際に利用したいモデルに名前空間を追加します。
ここではUser.phpを使用することにします。

User.php
`use Illuminate\Notifications\Notifiable;`

クラスに以下を追加。
`use Notifiable;`


参考までに↓

```php
<?php

use Hogehoge\hogehoge
use Illuminate\Notifications\Notifiable;

class User extends Authenticatable
{
  use Notifiable

  public fuction hoge()
 {
    //
 }
```

先ほど生成したUserRegistered.phpのtoMailメソッド内をいじります。いじらなくともデフォルトのメールテンプレートが送信できますので飛ばしてもOKです。

```php
public function toMail($notifiable)
    {
        $user_name = $notifiable->name;
        $token = $notifiable->confirmation_token;
        $url = url('/home');

        return (new MailMessage)
            ->view('path/to/mailTemplate')
            ->subject('ユーザー登録完了通知')
            ->line("{$user_name}さん登録ありがとう！")
            ->action('homeにもどる', "$url")
            ->line('今後ともよろしくでござる');
    }
```

メソッドについてはNotificationの[API](https://laravel.com/api/5.3/search.html?search=Notification)をご覧ください。


メールテンプレート参考

```html
@foreach ($introLines as $line)
  <p>
      {{ $line }}
  </p>
@endforeach

<p>
  下記にアクセスするとhome画面にもどるでござる。
</p>

<a href="{{ $actionUrl }}" target="_blank">
    {{ $actionText }}
</a>

@foreach ($outroLines as $line)
  <p>
      {{ $line }}
  </p>
@endforeach
```

SimpleMessageとかいう仕様に従った形なのですが、ちょっとだけわかりづらい気がするでござる。


ここまで完了したら後はコントローラで通知クラスを呼び出すだけです。
名前空間とFacedeを使って呼び出すことにします。Facedeを使わない場合はドキュメントをご参考に。


HogehogeController.php

```php
<?php
  namespace App\Http\Controllers\Hogehoge;

  use Hoge\Hogehoge;
  use App\Models\User;
  use App\Notifications\UserRegistered;

　　　　class HogehogeController extends Controller
  {
    public function hoge()
    {
       $user = new User();

       \Notification::send($user, new UserRegistered());
    }
  }
```

こんな感じでNotificationを使うことができます。
便利さが伝わったでしょうか・・・？

EventとListenerでSubscriberを使ってゴニョゴニョするよりは遥かに楽になったと思います。
通知だけならNotificationを利用した方が便利そうですね_(:3」∠)_


# ついでにElixirをバージョンアップ
package.jsonに以下を追記、または書き換え。

```json
  "laravel-elixir": "^6.0.0-9",
  "laravel-elixir-browserify-official": "^0.1.3",
  "laravel-elixir-webpack-official": "^1.0.2"
```

`npm install`


どのバージョンからかは失念しましたが、おそらくlaravel5.3からbrowerifyの扱いが変わるようで、個別にインストールしてあげないとbrowserifyを使えないそうです。
laravel-elixirだけアップデートしてgulp走らせるとbrowerifyのタスクがある場合は怒られるので直ぐに気づくとは思いますが。。。

# 参考
[Laravel.com - Notificaions](https://laravel.com/docs/5.3/notifications)
[Laravel Notification Channel](http://laravel-notification-channels.com/)

