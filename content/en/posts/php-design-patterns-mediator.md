---
title: Learning Design Patterns with PHP - Mediator Pattern
slug: php-design-patterns-mediator
date: 2019-01-31T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
  - Design Patterns
  - Mediator Pattern
  - GoF
translation_key: php-design-patterns-mediator
---

# Overview
This is an article that couldn't make it in time for the [Learning Design Patterns with PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# What is the Mediator Pattern?
It means mediator or arbitrator.

It is a design pattern related to the behavior of objects, used to coordinate interactions between objects.

It may be useful when interactions between objects become complex and relationships become unclear.

# Implementation

```php
<?php

// Mediator
class Receptionist
{
    public function checkIn(User $user, $message) // Holds the object whose behavior operation we want to delegate
    {
        echo $message . ' ' . $user->getName();
    }
}

class User
{
    private $name;
    private $receptionist;

    public function __construct($name, Receptionist $receptionist) // Holds the Mediator
    {
        $this->name = $name;
        $this->receptionist = $receptionist;
    }

    public function getName()
    {
        return $this->name;
    }

    public function checkIn($message)
    {
        $this->receptionist->checkIn($this, $message); // $this!!
    }
}

$receptionist = new Receptionist();


$john = new User('John', $receptionist);
$bob = new User('Bob', $receptionist);

$john->checkIn('Welcome!'); // Welcome! John
$bob->checkIn('Hi!'); // Hi! Bob
```

# Thoughts
This is a pattern I want to remember when I want to manage the behavior of objects collectively when interactions between classes seem to become complex.

# References
- [Introduction to Design Patterns in PHP - Mediator: Everything is Known by the Consultant](http://d.hatena.ne.jp/shimooka/20141217/1418788236)
- [[Essential] Human-readable Design Pattern Explanation #3: Behavioral Patterns (Translation)](https://techracho.bpsinc.jp/hachi8833/2017_10_17/46071#mediator)