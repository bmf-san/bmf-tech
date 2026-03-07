---
title: PHP Interfaces and Type Hinting
slug: php-interfaces-type-hinting
date: 2018-12-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
  - Interface
  - Type Hinting
translation_key: php-interfaces-type-hinting
---

# Overview
This article is part of the [PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php). (Posting a bit early)

Interfaces not only guarantee the implementation of methods as a 'contract', but they also allow for type hinting, which can make it easier to switch implementations by depending on abstractions.

# Defining and Implementing Interfaces
Basic definition and implementation of an interface.

```php
<?php
interface Action
{
    public function say();
}

class Superman implements Action
{
    public function say()
    {
        echo "Hello World";
    }
}

$obj = new Superman();
$obj->say();
```

# Separation of Functionality and Implementation via Interfaces
By specifying an interface type with type hinting, you can add flexibility to the implementation.

```php
<?php
interface HeroAction
{
    public function say();
}

class Superman implements HeroAction
{
    public function say()
    {
        echo "I'm a Superman";
    }
}

class Human
{
    public function say()
    {
        echo "I'm a Human";
    }
}

class Bot
{
    public function do(HeroAction $heroAction) // Specify interface type as argument
    {
        $heroAction->say();
    }
}

$superMan = new SuperMan();
human = new Human();
bot = new Bot();

$bot->do($superMan); // I'm a Superman
$bot->do($human); // PHP Fatal error:  Uncaught TypeError: Argument 1 passed to Bot::do() must implement interface HeroAction, instance of Human given, called in ....
```

Switching the implementation from Superman to Hyperman.

```php
<?php
interface HeroAction
{
    public function say();
}

// class Superman implements HeroAction
// {
//     public function say()
//     {
//         echo "I'm a Superman";
//     }
// }

class Hyperman implements HeroAction
{
    public function say()
    {
        echo "I'm a Hyperman";
    }
}

class Human
{
    public function say()
    {
        echo "I'm a Human";
    }
}

class Bot
{
    public function do(HeroAction $heroAction) // Specify interface type as argument
    {
        $heroAction->say();
    }
}

// $superMan = new SuperMan();
hyperMan = new HyperMan();
human = new Human();
bot = new Bot();

// $bot->do($superMan); // I'm a Superman
$bot->do($hyperMan); // I'm a Hyperman
$bot->do($human); // PHP Fatal error:  Uncaught TypeError: Argument 1 passed to Bot::do() must implement interface HeroAction, instance of Human given, called in ....
```

If the do method of the Bot class depended on the Superman class instead of the interface, it would increase the effort needed to swap implementations.
```php
class Bot
{
    public function do(Superman $superman) // Specify interface type as argument
    {
        $superman->say();
    }
}
```

# References
- [PHP - Object Interfaces](http://php.net/manual/ja/language.oop5.interfaces.php)
- [A Discussion on PHP Interfaces](http://blog.anatoo.jp/entry/20080517/1211029059)