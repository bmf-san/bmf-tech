---
title: Implementing the Repository Pattern in Laravel
slug: implement-laravel-repository-pattern
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
  - Repository Pattern
description: 'A smart implementation pattern for database operations: the Repository Pattern.'
translation_key: implement-laravel-repository-pattern
---

The Repository Pattern is a smart implementation pattern related to database operations. Let's explore it.

# What is the Repository Pattern?
The Repository Pattern separates data manipulation logic from business logic by delegating it to an abstract layer, enhancing maintainability and scalability. (It is not necessarily limited to database operation logic.)

By incorporating the Repository Pattern into Laravel, you can achieve:

* Easier testing
* Better adaptability to database engine changes
* Centralized data manipulation logic for easier management

# Implementing the Repository Pattern
Create a `Repository` directory corresponding to each model. (This approach may have mixed opinions.)

For this example, we'll implement the Repository Pattern with the following structure:

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

# Designing the Interface
First, design the interface:

```php
<?php

namespace App\Repositories\User;

interface UserRepositoryInterface
{
    /**
     * Retrieve a single record by Name
     *
     * @var string $name
     * @return object
     */
    public function getFirstRecordByName($name);
}
```

# Implementing the Class
Next, prepare the implementation class. Here, we'll inject the corresponding model and implement the methods:

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
     * Retrieve a single record by Name
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

You can further abstract the implementation by adding a Service layer, but for simplicity, we'll stick to these two classes.

# Service Provider
Register the interface and implementation class in `AppServiceProvider.php`:

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

# Calling it in the Controller
Use the implemented Repository Pattern in the controller:

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

Simply inject the interface!

# Thoughts
Both the model and controller look cleaner now. This implementation has inspired me to study Domain-Driven Design (DDD).

# References
* [Laravel4.2 Repository Pattern](http://tech.aainc.co.jp/archives/10227)
* [Laravel5.1 Tutorial: Intermediate Task List Part 4](http://laravel-room.com/tutorial-intermediate-4)
* [Application Design in Laravel Without Regrets](https://speakerdeck.com/localdisk/laravelniokeruhou-hui-sinaitamefalseapurikesiyonshe-ji)