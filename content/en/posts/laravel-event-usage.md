---
title: Using Events in Laravel
slug: laravel-event-usage
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
translation_key: laravel-event-usage
---

When you want to manage methods that should be triggered during specific events, such as user registration or cancellation, using event listeners is convenient.

In this post, we will skip the basic definitions of events and listeners and focus on **event subscription**, which allows you to set multiple events in a single listener class.

## Environment
* laravel5.2

## Directory Structure
* app\Events ... Place classes where the event name equals the class name (there are no strict naming conventions).
* app\Listeners ... Place classes that implement the processing for each event (listener) and the subscribe method (described later).
* app\Providers ... Place classes that register listeners used for event subscription.
* app\Controllers ... Prepare controllers and methods that will call the Event.

## Defining an Event
First, let's define an event. We will create an event called "User Registration Complete" as an example. You can generate the event class using the artisan command.

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

When you run the artisan command, a class like this will be automatically generated. Set the data to be used in the event within the constructor.

Since this event is related to user registration, we will call the User model. The data called here can be used in the listener class.

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

By the way, `broadcastOn` is used when you want to implement a user interface that runs in real-time, so we will skip it for now.

## Defining a Listener
Next, let's define a listener. You can also generate a listener using the artisan command.

`php artisan make:listener UserAuthEventListener --event UserRegistrationComplete`

When generating a listener, you can set the event you want to associate with using the event option.

```php
<?php

namespace App\Listeners;

use App\Events\UserRegistrationComplete;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Contracts\Queue\ShouldQueue;

class UserAuthEventListener
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

Since event subscription allows you to set multiple events in a single listener class, it is a good idea to name the listener in a way that reflects the category of multiple events. In this example, we could say, "The User Registration Complete event belongs to the User Authentication group."

In the generated listener, we will perform **processing when the event is triggered** and **register the event subscription**.

```php
<?php

namespace App\Listeners;

class UserAuthEventListener
{
    // Processing when the event is triggered
    public function onConfirm($event)
    {
      // Processing
    }

    // You can add multiple methods
    public function onHogeHoge($event)
    {
      // Processing
    }

    // Registering the event subscription
    public function subscribe($events)
    {
        $events->listen(
            'App\Events\UserRegistrationComplete',
            'App\Listeners\UserAuthEventListener@onConfirm'
        );
    }

    // You can register multiple events
    public function subscribe($events)
    {
        $events->listen(
            'App\Events\UserHogeHoge',
            'App\Listeners\UserAuthEventListener@onHogeHoge'
        );
    }
}
```

If you want to add an event, you can simply generate the event using the artisan command mentioned earlier.

## Registering the Event Subscription Class in EventServiceProvider
It may seem complicated with all this registration, but this is the last step. We will use the EventServiceProvider that exists by default in app\Providers.

A service provider is a class that performs the initial startup process of the application. For more details, refer to the [documentation](https://readouble.com/laravel/5/1/en/providers.html).

```php
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

Now that the event subscription registration is complete, you can freely trigger events.

To trigger an event, use the Event facade and the fire method. Pass the event instance to the fire method.

```php
    /**
     * A method in the controller
     */
    private function hogehoge(User $user)
    {
        // Triggering the registration confirmation event
        \Event::fire(new UserAuthRegistrationComplete($user));
    }
```

This is how the registered event gets triggered :fire:

## Conclusion

Event definition (data holding) → Listener definition (managing methods used in events and implementing the subscribe method) → Event subscription registration → :fire:

## Postscript
Since Laravel 5.3 introduced a Notification package, it might be easier to manage notifications using Notification instead of the method described here. For that reason, the directories for Event, Listener, and Job no longer exist by default in 5.3. If you're interested, please check the Laravel repository or documentation.

## References
* [Summary of how to use Events in Laravel 5.1](http://blog.fagai.net/2015/12/11/laravel51-event/)
* [Summarizing processing using the Event class in Laravel](http://localdisk.hatenablog.com/entry/2014/03/26/Laravel_%E3%81%AE_Event_%E3%82%AF%E3%83%A9%E3%82%B9%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E5%87%A6%E7%90%86%E3%82%92%E3%81%BE%E3%81%A8%E3%82%81%E3%81%A6%E3%81%BF%E3%82%8B)
* [Laravel Events](https://kore1server.com/292)