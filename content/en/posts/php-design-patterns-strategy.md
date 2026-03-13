---
title: Learning Design Patterns with PHP - Strategy
description: 'Implement the Strategy pattern to switch algorithms dynamically, reduce conditionals, and follow OCP in your PHP code.'
slug: php-design-patterns-strategy
date: 2018-12-09T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Design Pattern
  - PHP
  - GoF
  - Strategy Pattern
translation_key: php-design-patterns-strategy
---



# Overview
This article is part of the [Learning Design Patterns with PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

This time, I will write about the Strategy pattern.

# What is the Strategy Pattern?
The Strategy pattern is a pattern that makes it easy to switch algorithms. By defining different processes in separate classes, not only can you dynamically select processes, but you can also reduce conditional branches. It is one of the patterns faithful to the OCP (open/closed principle).

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

By making the implementation of the Context class dependent on the interface, the "strategy" can be switched from the client side. Isn't this a common implementation?

Since behaviors are separated, it should follow the OCP. For extensions, you just need to add an implementation of NotificationInterface, and for modifications, you just need to modify the implementation of NotificationInterface.

# Thoughts
I realized that sometimes I use GoF without even knowing it, so I thought it wouldn't hurt to remember it properly.

# References
- [Wikipedia - Strategy Pattern](https://ja.wikipedia.org/wiki/Strategy_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [Do You PHP Hatena - [doyouphp][phpdp] Introduction to Design Patterns with PHP - Switching Strategies](http://d.hatena.ne.jp/shimooka/20141219/1418965548)
- [isseium's blog - Design Patterns with PHP ~Strategy Pattern~](http://isseium.hateblo.jp/entry/20101114/1289725226)
- [DesignPatternsPHP - Strategy](https://designpatternsphp.readthedocs.io/en/latest/Behavioral/Strategy/README.html?highlight=strategy)
- [Quick Corporation's Web Service Development Blog - Let's Consider the Strategy Pattern with PHP](http://aimstogeek.hatenablog.com/entry/2016/12/13/190105)
