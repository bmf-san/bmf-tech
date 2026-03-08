---
title: DI and Service Locator
slug: dependency-injection-service-locator
date: 2018-06-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - DI
  - Service Locator
  - Design Patterns
translation_key: dependency-injection-service-locator
---

# Overview
Summarizing the differences between DI and Service Locator

# What is DI
- A type of design pattern
- Dependency Injection
  - Separates dependencies between objects
  - Ensures that the necessary objects are injected at runtime
- Makes testing easier

# Implementing the DI Pattern
Let's implement the DI pattern (constructor injection). Note that there are other methods of DI besides constructor injection, such as setter injection and method injection. For comparison, both a non-DI pattern and a DI pattern will be implemented.

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

`Application` does not hold the responsibility of `SlackNotification` and only depends on `NotificationInterface`. The `Application` accepts (depends on) `SlackNotification` in its constructor.

The dependency injection part can be mocked, making it easier to test.

```php
// For test
$mockNotification = new MockNotification(); // MockNotification implements NotificationInterface
$application = new Application($mockNotification);
$application->alert('mock alert!');
```

# Advantages and Disadvantages of DI
## Advantages
- Reduces coupling of objects and clarifies dependencies
- Resilient to changes
- Easier to test

## Disadvantages
- Increases the number of files
- Potential decrease in program execution speed

# DI Container
- A framework that provides DI functionality
- Provides a container that can manage dependencies collectively

# What is Service Locator
- A design pattern that abstracts the retrieval of services (objects)
- An anti-pattern
  - Increases the number of dependent classes (the service locator container class)
  - Makes testing cumbersome due to the container class
  - Makes dependent classes less clear
   - For more details, see "PHP The Right Way"

An example of this type:

```php
<?php

class Application
{
    public function __construct($container)
    {
        $this->slackNotification = $container['slack.notification'];
    }
}

$application = new Application($container);
```

# References
- [Notes on DI (Dependency Injection)](http://blog.shin1x1.com/entry/di-memo)
- [Creating and Learning DI Container in PHP - Part 1 - What is DI](https://qiita.com/zeriyoshi/items/e26daccd59669b623a41)
- [Creating and Learning DI Container in PHP - Part 2 - DI Container and Service Locator](https://qiita.com/zeriyoshi/items/ef71bec08441877ca219)
- [Design PatternsPHP - Service Locator](http://designpatternsphp.readthedocs.io/en/latest/More/ServiceLocator/README.html)
- [PHP The Right Way](http://ja.phptherightway.com/#containers)
- [How to avoid becoming a service locator when trying to use a DI container, specific examples are unclear](https://teratail.com/questions/49143)