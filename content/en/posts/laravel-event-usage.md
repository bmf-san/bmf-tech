---
title: Using Events in Laravel
slug: laravel-event-usage
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
description: Manage methods you want to trigger during specific events like user registration or withdrawal using event listeners.
translation_key: laravel-event-usage
---



Managing methods you want to trigger during specific events like user registration or withdrawal is convenient with event listeners.

This time, we will skip the basic definition of events and listeners and focus on **event subscription**, which allows you to set multiple events in a single listener class.

## Environment
* laravel5.2

## Directory
* app\Events: Place classes where the event name equals the class name (no strict naming conventions)
* app\Listeners: Place classes that implement processing (listeners) for each event and the subscribe method (described later)
* app\Providers: Place classes that register listeners used in event subscriptions
* app\Controllers: Prepare controllers and methods that call events

## Define an Event
First, let's define an event. As an example, let's create an event called "User Registration Complete." You can generate an event class using the artisan command.

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

When you run the artisan command, a class like this is automatically generated. Set the data to be used in the event in the constructor.

Since this event is related to user registration, let's call the User model. The data called here can be used in the listener class.

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

By the way, broadcastOn is used when you want to implement a user interface that runs in real-time, so we will skip it this time.

## Define a Listener
Next, let's define a listener. You can also generate a listener using the artisan command.

`php artisan make:listener UserAuthEventListenerListener --event UserRegistrationComplete`

When generating a listener, you can set the event you want to associate with using the event option.

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

Event subscription allows you to set multiple events in a single listener class, so it is good to name the listener like a category name for multiple events.

In this example, "User Registration Complete Event belongs to the User Authentication group."

In the generated listener, perform **event firing processing** and **event subscription registration**.

```php
<?php

namespace App\Listeners;

class UserAuthEventListener
{
    // Event firing processing
    public function onConfirm($event)
    {
      // Processing
    }

    // You can add multiple
    public function onHogeHoge($event)
    {
      // Processing
    }

    // Event subscription registration
    public function subscribe($events)
    {
        $events->listen(
            'App\Events\UserRegistrationComplete',
            'App\Listeners\UserAuthEventListener@onConfirm'
        );
    }

    // You can register multiple
    public function subscribe($events)
    {
        $events->listen(
            'App\Events\UserHogeHoge',
            'App\Listeners\UserAuthEventListener@onHogeHoge'
        );
    }
}

```

If you want to add an event, you can generate it with the artisan command as before.

## Register Event Subscription Class in EventServiceProvider
It might seem complicated with all the registrations, but this is the last step. Use the EventServiceProvider that exists by default in app\Providers.

A service provider is a class that performs the initial startup process of an application. For more details, see the [documentation](https://readouble.com/laravel/5/1/ja/providers.html)

```
<?php

namespace App\Providers;

use Illuminate\Foundation\Support\Providers\EventServiceProvider as ServiceProvider;

class EventServiceProvider extends ServiceProvider
{
    // Register listeners here
    protected $subscribe = [
        'App\Listeners\UserAuthEventListener',
        'App\Listeners\HogeHogeListener',
    ];
}

```

## Fire

Now that the event subscription registration is complete, let's freely fire events.

To fire an event, use the Event facade and the fire method.

Pass the event instance to the fire method.

```php
    /**
     * A certain method in the controller
     */
    private function hogehoge(User $user)
    {
        // Fire registration confirmation event
        \Event::fire(new UserAuthRegistrationComplete($user));
    }

```

This is how the registered event ignites :fire:

## Summary

Event definition (data retention) → Listener definition (method management used in events and subscribe method implementation) → Event subscription registration → :fire:

## Additional Notes
Since Laravel 5.3, a package called Notification has been introduced, so managing notifications like this might be easier with Notification. Perhaps for this reason, directories like Event, Listener, and Job no longer exist by default in 5.3. If you're interested, check out the Laravel repository or documentation.

## #References
* [Summarizing the use of Events in Laravel 5.1.](http://blog.fagai.net/2015/12/11/laravel51-event/)
* [Summarizing processing using the Event class in Laravel](http://localdisk.hatenablog.com/entry/2014/03/26/Laravel_%E3%81%AE_Event_%E3%82%AF%E3%83%A9%E3%82%B9%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E5%87%A6%E7%90%86%E3%82%92%E3%81%BE%E3%81%A8%E3%82%81%E3%81%A6%E3%81%BF%E3%82%8B)
* [Laravel Events](https://kore1server.com/292)
