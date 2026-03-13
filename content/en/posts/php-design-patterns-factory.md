---
title: Learning Design Patterns with PHP - Factory, Factory Method, Abstract Factory
description: 'Learn Factory, Factory Method, and Abstract Factory patterns to centralize object creation and eliminate conditional branches.'
slug: php-design-patterns-factory
date: 2018-12-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Design Patterns
  - GoF
  - PHP
  - Abstract Factory Pattern
  - Factory Pattern
  - Factory Method Pattern
translation_key: php-design-patterns-factory
---

This article is part of the [Learning Design Patterns with PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# Overview
This post will cover Factory, Factory Method, and Abstract Factory.

# Understanding the Factory Pattern
First, let's briefly explain the Factory pattern.

```php
interface Robot
{
  public function say();
}

class BlueRobot implements Robot
{
    private $color;

    public function __construct($color)
    {
        $this->color = $color;
    }

    public function say()
    {
        echo $this->color;
    }
}

class YellowRobot implements Robot
{
    private $color;

    public function __construct($color)
    {
        $this->color = $color;
    }

    public function say()
    {
        echo $this->color;
    }
}

class RobotFactory
{
    public function create($color)
    {
        if ($color === 'blue') {
            return new BlueRobot($color);
        }

        if ($color === 'yellow') {
            return new YellowRobot($color);
        }

        throw new Exception("Can't create an instance");
    }
}

$robotFactory = new RobotFactory();

try {
    $blueRobot = $robotFactory->create('blue');
    $blueRobot->say(); // blue

    $yellowRobot = $robotFactory->create('yellow');
    $yellowRobot->say(); // yellow

    $greenRobot = $robotFactory->create('green');
    $greenRobot->say(); // Can't create an instance
} catch (\Exception $e) {
    echo $e->getMessage();
}
```

As the name suggests, a factory is responsible for the "creation" of objects. By separating the creation and usage of objects, the user can create objects without knowing the order or type of objects. This allows the creation location to be centralized, making it easier to change the order or type of objects.

# Understanding the Factory Method Pattern
In the Factory pattern mentioned above, as the number of classes of objects to be created increases, the conditional branches also increase, which could lead to a difficult future. The Factory Method pattern aims to create objects without specifying their classes.

By abstracting the factory and delegating the creation process to subclasses, conditional branches can be eliminated.

When changing the object to be created, it involves switching the factory.

```php
interface Robot
{
  public function say();
}

class BlueRobot implements Robot
{
    public function say()
    {
        echo 'Blue';
    }
}

class YellowRobot implements Robot
{
    public function say()
    {
        echo 'Yellow';
    }
}

abstract class RobotFactory
{
    abstract protected function create();

    public function do()
    {
        $robot = $this->create();
        $robot->say();
    }
}

class BlueRobotFactory extends RobotFactory
{
    protected function create()
    {
        return new BlueRobot();
    }
}

class YellowRobotFactory extends RobotFactory
{
    protected function create()
    {
        return new YellowRobot();
    }
}

$blueRobot = new BlueRobotFactory();
$blueRobot->do(); // blue

$yellowRobot = new YellowRobotFactory();
$yellowRobot->do(); // yellow
```

I modified the code from the Factory pattern to apply the Factory Method pattern.

The part where the object to be created was switched based on arguments has changed significantly.

Although the Factory pattern and Factory Method pattern have different object creation patterns, there are some articles that confuse the two, which I find troublesome.

# Understanding the Abstract Factory Pattern
The Abstract Factory pattern is a pattern that groups multiple factories based on a common theme.

```php
interface Robot
{
  public function say();
}

class BlueRobot implements Robot
{
    public function say()
    {
        echo 'Blue';
    }
}

class YellowRobot implements Robot
{
    public function say()
    {
        echo 'Yellow';
    }
}

interface RobotCreator
{
    public function work();
}

class BlueRobotCreator implements RobotCreator
{
    public function work()
    {
        echo '青いロボットつくるよ';
    }
}

class YellowRobotCreator implements RobotCreator
{
    public function work()
    {
        echo '黄色いロボットつくるよ';
    }
}

interface RobotFactory
{
    public function createRobot();
    public function createRobotCreator();
}

class BlueRobotFactory implements RobotFactory
{
    public function createRobot()
    {
        return new BlueRobot();
    }

    public function createRobotCreator()
    {
        return new BlueRobotCreator();
    }
}

class YellowRobotFactory implements RobotFactory
{
    public function createRobot()
    {
        return new YellowRobot();
    }

    public function createRobotCreator()
    {
        return new YellowRobotCreator();
    }
}


$blueRobotFactory = new BlueRobotFactory();
$blueRobot = $blueRobotFactory->createRobot();
$blueRobotCreator = $blueRobotFactory->createRobotCreator();

$blueRobot->say(); // blue
$blueRobotCreator->work(); // 青いロボットつくるよ

$yellowRobotFactory = new YellowRobotFactory();
$yellowRobot = $yellowRobotFactory->createRobot();
$yellowRobotCreator = $yellowRobotFactory->createRobotCreator();

$yellowRobot->say(); // yellow
$yellowRobotCreator->work(); // 青いロボットつくるよ
```

The Abstract Factory feels like an evolved version of the Factory Method.

By defining the interface of the factory, it allows for interpretation from a higher-level concept, but for now, I'll just grasp the nuance.

# Impressions
I feel like I've understood the nuances of each pattern. If I knew more about the essence of object-oriented programming, I might be able to interpret it more deeply.

# References
- [DesignPatternsPHP](https://designpatternsphp.readthedocs.io/en/latest/README.html)
- [PHP The Right Way = Design Patterns](https://laravel-taiwan.github.io/php-the-right-way/pages/Design-Patterns.html)
- [Tech Racho - [保存版]人間が読んで理解できるデザインパターン解説#1: 作成系（翻訳）](https://techracho.bpsinc.jp/hachi8833/2017_10_02/46064)
- [Object Thinking - Factory Pattern](https://think-on-object.blogspot.com/2011/11/factoryfactory-methodabstract-factory.html)