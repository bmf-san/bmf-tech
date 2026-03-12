---
title: Learning Design Patterns with PHP - Mediator Pattern
description: An in-depth exploration of Learning Design Patterns with PHP - Mediator Pattern, covering design principles, trade-offs, and practical applications.
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
An article that didn't make it in time for the [PHP Design Patterns Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# What is the Mediator Pattern?
It means mediator or arbitrator.

A behavioral design pattern that coordinates interactions between objects.

It might be useful when interactions between objects become complex and relationships become difficult to see.

# Implementation

```php
<?php

// Mediator
class Receptionist
{
    public function checkIn(User $user, $message) // Holds the object to which you want to delegate behavior operations
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
A pattern to remember when you want to collectively manage the behavior of objects when interactions between classes seem to become complex.

# References
- [Introduction to Design Patterns with PHP - Mediator: The Consultant Knows Everything](http://d.hatena.ne.jp/shimooka/20141217/1418788236)
- [[Definitive Edition] Design Pattern Explanation That Humans Can Read and Understand #3: Behavioral (Translation)](https://techracho.bpsinc.jp/hachi8833/2017_10_17/46071#mediator)
