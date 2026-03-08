---
title: Learning Design Patterns with PHP - Factory, Factory Method, Abstract Factory
slug: php-design-patterns-factory
date: 2018-12-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Design Patterns
  - GoF
  - PHP
  - Abstract Factory
  - Factory Pattern
  - Factory Method
translation_key: php-design-patterns-factory
---

This article is part of the [Learning Design Patterns with PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# Overview
I will write about Factory, Factory Method, and Abstract Factory.

# Understanding the Factory Pattern
First, let me briefly explain the Factory pattern.

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
    $greenRobot->say(); // Cat't create an instance
} catch (\Exception $e) {
    echo $e->getMessage();
}
```

As the name suggests, a factory is responsible for the "creation" of objects. By separating the creation and usage of objects, the user can create objects without knowing the order or type of object creation. This allows for consolidating the object creation in one place, making it easier to change the order or type of creation.

# Understanding the Factory Method Pattern
In the aforementioned Factory pattern, as the number of classes for the objects to be created increases, the conditional branches also increase, leading to a difficult future. Therefore, the Factory Method pattern aims to create objects without specifying the class of the object.

By abstracting the Factory and delegating the object creation process to subclasses, we can eliminate conditional branches.

When changing the object to be created, we switch the Factory.

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

I modified the code from the Factory pattern to apply the Factory Method pattern. The part that switched the created object based on arguments has changed significantly.

The Factory pattern and the Factory Method pattern have different object creation patterns, but I noticed that some articles confuse the two, which can be troublesome.

# Understanding the Abstract Factory Pattern
The Abstract Factory pattern is a pattern that groups multiple factories by a common theme.

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
        echo 'Creating a blue robot';
    }
}

class YellowRobotCreator implements RobotCreator
{
    public function work()
    {
        echo 'Creating a yellow robot';
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
$blueRobotCreator->work(); // Creating a blue robot

$yellowRobotFactory = new YellowRobotFactory();
$yellowRobot = $yellowRobotFactory->createRobot();
$yellowRobotCreator = $yellowRobotFactory->createRobotCreator();

$yellowRobot->say(); // yellow
$yellowRobotCreator->work(); // Creating a yellow robot
```

The Abstract Factory seems to be an evolved version of the Factory Method.

Since it defines the interface for the Factory, I think it can be interpreted from a higher-level concept, but for now, I will just grasp the nuance.

# Thoughts
I feel like I have understood the nuances of each pattern. If I knew a bit more about the essence of object-oriented programming, I could provide a deeper interpretation.

# References
- [DesignPatternsPHP](https://designpatternsphp.readthedocs.io/en/latest/README.html)
- [PHP The Right Way = Design Patterns](https://laravel-taiwan.github.io/php-the-right-way/pages/Design-Patterns.html)
- [Tech Racho - [Preserved] Human-readable Design Pattern Explanation #1: Creation Patterns (Translation)](https://techracho.bpsinc.jp/hachi8833/2017_10_02/46064)
- [Object-Oriented - Factory Pattern](https://think-on-object.blogspot.com/2011/11/factoryfactory-methodabstract-factory.html)