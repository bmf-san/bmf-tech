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

This article is part of the [PHP Design Patterns Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# Overview
I will write about Factory, Factory Method, and Abstract Factory.

# Understanding the Factory Pattern
First, let’s briefly explain the Factory pattern.

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

As the name suggests, a factory is responsible for the "creation" of objects. By separating the creation and usage of objects, the user can create objects without knowing the order or type of object being created. This allows for centralizing the creation of objects in one place, making it easier to change the order or type of creation.

# Understanding the Factory Method Pattern
In the aforementioned Factory pattern, as the number of classes for the objects to be created increases, the number of conditional branches also increases, leading to a difficult future. Therefore, the Factory Method pattern aims to create objects without specifying the class of the object.

By abstracting the factory and delegating the object creation process to subclasses, we can eliminate conditional branches.

When changing the object to be created, we switch the factory.

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

The part that switched the created object based on the argument has changed significantly.

While the Factory pattern and Factory Method pattern differ in their object creation patterns, I noticed that some articles confuse the two, which can be troublesome.

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
        echo 'I will create a blue robot';
    }
}

class YellowRobotCreator implements RobotCreator
{
    public function work()
    {
        echo 'I will create a yellow robot';
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
$blueRobotCreator->work(); // I will create a blue robot

$yellowRobotFactory = new YellowRobotFactory();
$yellowRobot = $yellowRobotFactory->createRobot();
$yellowRobotCreator = $yellowRobotFactory->createRobotCreator();

$yellowRobot->say(); // yellow
$yellowRobotCreator->work(); // I will create a yellow robot
```

The Abstract Factory feels like an evolved version of the Factory Method.

In that it defines the interface for the factory, I think it can be interpreted from a higher-level concept, but for now, I will leave it at grasping the nuance.

# Personal Reflections
I feel like I have understood the nuances of each pattern. If I knew a bit more about the essence of object-oriented programming, I could provide a deeper interpretation.

# References
- [DesignPatternsPHP](https://designpatternsphp.readthedocs.io/en/latest/README.html)
- [PHP The Right Way = Design Patterns](https://laravel-taiwan.github.io/php-the-right-way/pages/Design-Patterns.html)
- [Tech Racho - [Preserved Edition] An Explanation of Design Patterns That Humans Can Read and Understand #1: Creation Patterns (Translation)](https://techracho.bpsinc.jp/hachi8833/2017_10_02/46064)
- [Object-Oriented Thinking - Factory Pattern](https://think-on-object.blogspot.com/2011/11/factoryfactory-methodabstract-factory.html)
