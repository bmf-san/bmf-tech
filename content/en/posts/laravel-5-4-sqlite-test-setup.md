---
title: Preparing to Write Tests with SQLite in Laravel 5.4
description: 'Configure Laravel 5.4 SQLite testing with in-memory databases, phpunit setup, and database migration automation.'
slug: laravel-5-4-sqlite-test-setup
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - Laravel
  - SQLite
translation_key: laravel-5-4-sqlite-test-setup
---



# Overview
Prepare to write tests using SQLite in Laravel 5.4.

# Prerequisites
- Basic knowledge of Laravel
- Migration files prepared

# Edit phpunit.xml

Add the following three lines. We will use SQLite's in-memory feature.

```php:phpunit.xml
    <php>
        <env name="APP_ENV" value="testing"/>
        <env name="CACHE_DRIVER" value="array"/>
        <env name="SESSION_DRIVER" value="array"/>
        <env name="QUEUE_DRIVER" value="sync"/>
        <env name="DB_CONNECTION" value="sqlite"/> // Add
        <env name="DB_DATABASE" value=":memory:"/> // Add
        <env name="ADMIN_DOMAIN" value="localhost"/> // Add
    </php>
```

# Edit ModelFactory
Prepare this as needed. Omitted.

# Write Tests
```php:phpunit.xml
<testsuites>
        <testsuite name="Application Test Suite">
            <directory suffix="Test.php">./tests</directory>
        </testsuite>
</testsuites>
```

By default, phpunit.xml is set to execute any file ending with Test.php under the tests directory, so prepare your test files accordingly.

```php:/tests/Unit/HogeTest.php
<?php

use Illuminate\Foundation\Testing\WithoutMiddleware;
use Illuminate\Foundation\Testing\DatabaseMigrations;
use Illuminate\Foundation\Testing\DatabaseTransactions;

class PostTest extends TestCase
{
    use DatabaseMigrations; // Run migration

    public function testIndex()
    {
        $user = factory(App\User::class)->create();

        var_dump($user->first()); // Check the data

        // Here is your tests...
    }
}
```

When you set the `DatabaseMigrations` trait, migrations are executed every time tests are run.

# Additional Notes
You can make it easier by adding the command to run tests, `./vendor/bin/phpunit`, to the composer scripts.

# Thoughts
Now you can write tests.