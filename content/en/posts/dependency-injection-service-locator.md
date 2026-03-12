---
title: DI and Service Locator
description: An in-depth look at DI and Service Locator, covering key concepts and practical insights.
slug: dependency-injection-service-locator
date: 2018-06-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - DI
  - Service Locator
  - Design Pattern
translation_key: dependency-injection-service-locator
---

# Overview
Summarizing the differences between DI and Service Locator

# What is DI
- A type of design pattern
- Dependency Injection
  - Separates dependencies between objects
  - Ensures necessary objects are injected at runtime
- Makes testing easier

# Implementing the DI Pattern
Let's implement the DI pattern (Constructor Injection). Note that the DI pattern includes methods other than constructor injection, such as setter injection and method injection. For comparison, both non-DI and DI patterns will be implemented.

## Non-DI Pattern

```php
<?php
class SlackNotification
{
    public function notify(string $message)
    {
        echo $message;

        return $this;
    }
}

class Application
{
    protected $message;

    public function __construct()
    {
        $this->notification = new SlackNotification();
    }

    public function alert(string $message)
    {
        $this->notification->notify($message);
    }
}

// client
$application = new Application();
$application->alert('slack alert');
```

## DI Pattern

```php
<?php

interface NotificationInterface
{
    public function notify(string $message);
}

class SlackNotification implements NotificationInterface
{
    public function notify(string $message)
    {
        echo $message;

        return $this;
    }
}

class Application
{
    protected $message;

    public function __construct(NotificationInterface $notification) // DI
    {
        $this->notification = $notification;
    }

    public function alert(string $message)
    {
        $this->notification->notify($message);
    }
}

// client
$slackNotification = new SlackNotification();
$application = new Application($slackNotification); // DI
$application->alert('slack alert!');
```

`Application` does not hold the responsibility of `SlackNotification`, and only depends on `NotificationInterface`. `Application` is structured to accept (depend on) `SlackNotification` in the constructor.

The dependency injection part can be mocked, making it easier to test.

```php
// For test
$mockNotification = new MockNotification(); // MockNotification implementing NotificationInterface
$application = new Application($mockNotification);
$application->alert('mock alert!');
```

# Advantages and Disadvantages of DI
## Advantages
- Reduces coupling between objects and clarifies dependencies
- Resistant to changes
- Easier to test

## Disadvantages
- Increases the number of files
- Potential decrease in program execution speed

# DI Container
- A framework that provides DI functionality
- Offers a container to manage dependencies collectively

# What is a Service Locator
- A design pattern that abstracts the retrieval of services (objects)
- Considered an anti-pattern
  - Adds one more dependent class (the container class of the service locator)
  - Makes the container class difficult to test
  - Makes it hard to understand dependent classes
   - For more details, see "PHP The Right Way"

Something like this

```php
<?php

class Application
{
    public function __construct($container)
    {
        $this->slackNotification = $container['slack.notification'];
    }
}

$application = new Applicaion($container);
```

# References
- [Notes on DI (Dependency Injection)](http://blog.shin1x1.com/entry/di-memo)
- [Learn DI Container by Creating with PHP - Part 1 - What is DI](https://qiita.com/zeriyoshi/items/e26daccd59669b623a41)
- [Learn DI Container by Creating with PHP - Part 2 - DI Container and ServiceLocator](https://qiita.com/zeriyoshi/items/ef71bec08441877ca219)
- [Design PatternsPHP - Service Locator](http://designpatternsphp.readthedocs.io/en/latest/More/ServiceLocator/README.html)
- [PHP The Right Way](http://ja.phptherightway.com/#containers)
- [How to avoid becoming a service locator when trying to use a DI container, I don't understand specific examples](https://teratail.com/questions/49143)