---
title: Using Notifications in Laravel 5.2
slug: use-notification-laravel-5-2
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
description: Discussion on using Notifications in Laravel 5.2.
translation_key: use-notification-laravel-5-2
---


Laravel 5.3 has been released, but this post discusses using Notifications in Laravel 5.2.

# Environment
* Laravel 5.2 (or 5.1)

# Install Notification
What is Notification?

[Laravel.com - Notifications](https://laravel.com/docs/5.3/notifications)
[Laravel Notification Channel](http://web.archive.org/web/20251112051552/https://laravel-notification-channels.com/)

Before Laravel 5.3, notification management was somewhat handled using Events and Listeners. However, with Notifications, managing notifications has become more convenient. Note that Notifications are specialized for notifications and are not a "replacement" for Events or Listeners.

Let's install it.
`composer require laravel-notification-channels/backport`

In Laravel 5.3, it is included by default, so there's no need to install it manually with composer, but for versions up to 5.2, it must be done manually.

If you want to use the facade, add the facade class.

*Skip this if you don't plan to use Notification with Facade.*

Create Notification.php in `vendor/laravel/framework/src/Illuminate/Support/Facades` and add the following:

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

Then, add the following to the providers and aliases in `app.php`.

`Illuminate\Notifications\NotificationServiceProvider::class`
`'Notification' => Illuminate\Support\Facades\Notification::class`

Preparation is now complete.

# Try It Out
First, create a notification class. As an example, let's create a class to manage notifications when a user registers.

`php artisan make:notification UserRegistered`

A Notification directory will be generated, and UserRegistered.php will be created within it.

Once confirmed, add the namespace to the model you want to use for notifications. Here, we'll use User.php.

User.php
`use Illuminate\Notifications\Notifiable;`

Add the following to the class.
`use Notifiable;`

For reference:

```php
<?php

use Hogehoge\hogehoge
use Illuminate\Notifications\Notifiable;

class User extends Authenticatable
{
  use Notifiable

  public function hoge()
 {
    //
 }
```

Modify the `toMail` method in the previously generated UserRegistered.php. You can skip this if you want to use the default email template.

```php
public function toMail($notifiable)
    {
        $user_name = $notifiable->name;
        $token = $notifiable->confirmation_token;
        $url = url('/home');

        return (new MailMessage)
            ->view('path/to/mailTemplate')
            ->subject('User Registration Complete Notification')
            ->line("Thank you for registering, {$user_name}!")
            ->action('Return to home', "$url")
            ->line('We look forward to serving you.');
    }
```

For method details, see the Notification ~~API~~.

Email template reference:

```html
@foreach ($introLines as $line)
  <p>
      {{ $line }}
  </p>
@endforeach

<p>
  Access the link below to return to the home screen.
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

It follows the SimpleMessage specification, but it might be a bit unclear.

Once this is complete, just call the notification class in the controller. Use the namespace and Facade to call it. Refer to the documentation if not using Facade.

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

This is how you can use Notifications. Did the convenience come across?

It's much easier than using Subscribers with Events and Listeners. For notifications alone, using Notifications seems more convenient.

# Upgrade Elixir
Add or replace the following in `package.json`.

```json
  "laravel-elixir": "^6.0.0-9",
  "laravel-elixir-browserify-official": "^0.1.3",
  "laravel-elixir-webpack-official": "^1.0.2"
```

`npm install`

I forgot from which version, but it seems that from Laravel 5.3, the handling of browserify changes, and you need to install it separately to use it. If you only update laravel-elixir and run gulp, you'll notice if there's a browserify task.

# References
[Laravel.com - Notifications](https://laravel.com/docs/5.3/notifications)
[Laravel Notification Channel](http://web.archive.org/web/20251112051552/https://laravel-notification-channels.com/)
