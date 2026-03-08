---
title: Learning Design Patterns with PHP - Strategy
slug: php-design-patterns-strategy
date: 2018-12-09T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Design Patterns
  - PHP
  - GoF
  - Strategy
translation_key: php-design-patterns-strategy
---

# Overview
This article is part of the [Learning Design Patterns with PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

This time, I will write about the Strategy pattern.

# What is the Strategy Pattern?
The Strategy pattern is a pattern that makes it easy to switch algorithms. By defining different processes in separate classes, it allows for dynamic selection of processes and reduces conditional branching. It is also one of the patterns that adheres to the OCP (open/closed principle).

# Implementation Example
Let's look at a simple example of implementing the Strategy pattern.

```php
<?php
class Context
{
    private $notification;

    public function __construct(NotificationInterface $notification)
    {
        $this->notification = $notification;
    }

    public function execute()
    {
        return $this->notification->notify();
    }
}

interface NotificationInterface
{
    public function notify();
}

class SlackNotification implements NotificationInterface
{
    public function __construct($message)
    {
        $this->message = $message;
    }

    public function notify()
    {
        echo $this->message . '- sent by Slack';
    }
}

class EmailNotification implements NotificationInterface
{
    public function __construct($message)
    {
        $this->message = $message;
    }

    public function notify()
    {
        echo $this->message . '- sent by Email';
    }
}

$message = "Hello World!";

$slack = new SlackNotification($message);
$context = new Context($slack);
$context->execute(); // Hello World - sent by Slack

$email = new EmailNotification($message);
$context = new Context($email);
$context->execute(); // Hello World - sent by Email
```

By making the implementation of the Context class depend on the interface, we can switch the "strategy" from the client side. This is a common implementation, isn't it?

Since the behavior is separated, it should adhere to the OCP. For extensions, you just need to add an implementation of NotificationInterface, and for modifications, you only need to modify the implementation of NotificationInterface.

# Thoughts
I realized that I often use GoF patterns without even knowing it, so it's worth remembering them properly.

# References
- [Wikipedia - Strategy Pattern](https://ja.wikipedia.org/wiki/Strategy_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [Do You PHP Hatena - Introduction to Design Patterns with PHP - Strategy](http://d.hatena.ne.jp/shimooka/20141219/1418965548)
- [isseium's blog - Design Patterns in PHP - Strategy Pattern](http://isseium.hateblo.jp/entry/20101114/1289725226)
- [DesignPatternsPHP - Strategy](https://designpatternsphp.readthedocs.io/en/latest/Behavioral/Strategy/README.html?highlight=strategy)
- [Quick Corporation's Web Service Development Blog - Let's Think About the Strategy Pattern in PHP](http://aimstogeek.hatenablog.com/entry/2016/12/13/190105)