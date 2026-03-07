---
title: Learning Design Patterns with PHP - Adapter Pattern
slug: php-design-patterns-adapter
date: 2019-02-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Adapter Pattern
  - GoF
  - PHP
  - Design Patterns
translation_key: php-design-patterns-adapter
---

# Overview
This is an article that couldn't make it in time for the [Learning Design Patterns with PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# What is the Adapter Pattern?
A pattern that allows changing the interface without modifying the original class.

This is achieved by preparing an Adapter class that adjusts compatibility between different interfaces.

# Implementation
```php
<?php
interface Bird
{
    public function fly();
}

class SmallBird implements Bird
{
    public function fly()
    {
        echo 'fly short time';
    }
}

class BigBird implements Bird
{
    public function fly()
    {
        echo 'fly long time';
    }
}

class Human
{
    public function eat(Bird $bird)
    {
        echo 'Yummy!';
    }
}

class MiddleBird
{
    public function jump()
    {
        echo 'jump like flying';
    }
}

// Adapter
class MiddleBirdAdapter implements Bird
{
    private $middleBird;

    public function __construct(MiddleBird $middleBird)
    {
        $this->middleBird = $middleBird;
    }

    // Wrap MiddleBird's method
    public function fly()
    {
        $this->middleBird->jump();
    }
}

$human = new Human();

$smallBird = new SmallBird();
$human->eat($smallBird); // Yummy

$middleBird = new MiddleBird();
$middleBirdAdapter = new MiddleBirdAdapter($middleBird);

$human->eat($middleBirdAdapter); // Yummy
```

# Thoughts
It feels like creating a method that wraps the behavior defined in the interface. It seems necessary to carefully consider where to use it.

# References
- [[Preserved Edition] Human-readable Design Pattern Explanation #2: Structural Patterns (Translation)](https://techracho.bpsinc.jp/hachi8833/2017_10_11/46069#adapter)
- [Wikipedia - Adapter Pattern](https://ja.wikipedia.org/wiki/Adapter_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)