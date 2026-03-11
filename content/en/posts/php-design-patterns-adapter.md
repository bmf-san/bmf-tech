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
An article that didn't make it in time for the [PHP Design Patterns Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# What is the Adapter Pattern?
A pattern that allows you to change the interface without modifying the original class.

It is achieved by preparing an Adapter class that adjusts compatibility between different interfaces.

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

    // Wraps the method of MiddleBird
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
It feels like creating a method that wraps behavior defined in the interface.
You need to carefully consider where to use it.

# References
- [[Definitive Edition] Design Pattern Explanation That Humans Can Read and Understand #2: Structural (Translation)](https://techracho.bpsinc.jp/hachi8833/2017_10_11/46069#adapter)
- [Wikipedia - Adapter Pattern](https://ja.wikipedia.org/wiki/Adapter_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
