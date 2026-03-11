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
description: A smart implementation pattern related to DB operations, let's discuss the repository pattern.
translation_key: implement-laravel-repository-pattern
---


A smart implementation pattern related to DB operations, let's discuss the repository pattern.

# What is the Repository Pattern?
The repository pattern is a way to separate the logic related to data manipulation from business logic, delegating it to an abstract layer to enhance maintainability and scalability. (It is not necessarily a pattern that only retains DB operation logic.)

By incorporating the repository pattern into Laravel, you can gain benefits such as:

* Easier testing
* Easier adaptation to changes in the DB engine
* Centralized data manipulation logic, making it easier to manage

# Implementing the Repository Pattern
Create a Repository directory at the same level as the Model. (This might be a controversial choice)

This time, we will implement the repository pattern with the following structure:

```:php
.
├── Models
│   ├── User.php
│  
├── Repositories
    └── User
        ├── UserRepository.php
        └── UserRepositoryInterface.php
```

# Interface Design
First, design the interface.

```php
<?php

namespace App\Repositories\User;

interface UserRepositoryInterface
{
    /**
     * Get one record by Name
     *
     * @var string $name
     * @return object
     */
    public function getFirstRecordByName($name);
}
```

# Implementation Class
Next, prepare the implementation class. Here, we perform DI of the corresponding model and implement the methods.

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
     * Get one record by name
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

From here, you can further prepare a Service layer and add classes to increase abstraction, but this time we will implement with just these two classes.

# Service Provider
Register the interface and implementation class in AppServiceProvider.php.

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

# Calling in the Controller
Use the implemented repository pattern.

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

Just inject the interface!

# Thoughts
Both the model and the controller have become cleaner. I would like to take this opportunity to study DDD as well.

# References
* [Laravel4.2 Repository Pattern](http://tech.aainc.co.jp/archives/10227)
* [Intermediate Task List Tutorial for Laravel5.1](http://laravel-room.com/tutorial-intermediate-4)
* [Application Design in Laravel to Avoid Regrets](https://speakerdeck.com/localdisk/laravelniokeruhou-hui-sinaitamefalseapurikesiyonshe-ji)
