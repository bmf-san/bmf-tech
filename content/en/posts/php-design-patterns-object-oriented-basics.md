---
title: Learning Design Patterns with PHP - Basics of Object-Oriented Programming
description: 'Master OOP fundamentals including inheritance, interfaces, traits, static properties, and methods for design pattern study.'
slug: php-design-patterns-object-oriented-basics
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Design Patterns
  - PHP
  - OOP
translation_key: php-design-patterns-object-oriented-basics
---

# Overview

We will learn design patterns using PHP, using the now out-of-print [Introduction to Design Patterns with PHP](https://www.amazon.co.jp/PHP%E3%81%AB%E3%82%88%E3%82%8B%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E4%B8%8B%E5%B2%A1-%E7%A7%80%E5%B9%B8/dp/4798015164) as our textbook. (Note: It is available second-hand on Amazon, but it seems to be priced at more than double the original price.)

The code covered in this series will be compiled on [GitHub](https://github.com/bmf-san/design-patterns-php).

Ideally, I would have liked to learn design patterns in a language that pioneered OOP, but due to my lack of familiarity with languages other than PHP and encountering a book explaining design patterns in PHP, I decided to learn design patterns with PHP.

# References
* [Introduction to Design Patterns with PHP](https://www.amazon.co.jp/PHP%E3%81%AB%E3%82%88%E3%82%8B%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E4%B8%8B%E5%B2%A1-%E7%A7%80%E5%B9%B8/dp/4798015164)
* [Do You PHP Hatena](http://d.hatena.ne.jp/shimooka/20141211/1418298136)
* [Github shimooka/PhpDesignPattern](https://github.com/shimooka/PhpDesignPattern)
* [Perfect PHP](https://www.amazon.co.jp/%E3%83%91%E3%83%BC%E3%83%95%E3%82%A7%E3%82%AF%E3%83%88PHP-PERFECT-3-%E5%B0%8F%E5%B7%9D-%E9%9B%84%E5%A4%A7/dp/4774144371)
* [Programming PHP, 3rd Edition](https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0PHP-%E7%AC%AC3%E7%89%88-Kevin-Tatroe/dp/4873116686/ref=sr_1_1?ie=UTF8&qid=1479547951&sr=8-1&keywords=%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0PHP)

References for each pattern will be provided in the respective articles.

# Why Learn

There are two purposes for learning design patterns.

**To deepen understanding of frameworks**

Design patterns are utilized throughout framework design. By knowing the patterns, you can deepen your understanding of frameworks and make the most of them.

**To write quality code**

Design patterns can be described as the culmination of object-oriented principles built by predecessors. It might be challenging to adopt and use design patterns immediately, but by learning the "way of thinking," you can deepen your understanding of object-oriented programming, which may lead to improving the quality of your daily source code.

# Basics of Object-Oriented Programming

Before diving into the main topic, let's review the basics of object-oriented programming.

# Important Notes

Please note that [Introduction to Design Patterns with PHP](https://www.amazon.co.jp/PHP%E3%81%AB%E3%82%88%E3%82%8B%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E4%B8%8B%E5%B2%A1-%E7%A7%80%E5%B9%B8/dp/4798015164) was published in 2006, so there might be parts where the thinking about the use of patterns has changed since then. Please understand that this cannot be addressed in the initial draft.

# Class Declaration

```php
<?php

class ClassName extends BaseClassName implements InterfaceName
{
    use TraitName;

    public $public_property;
    protected $protected_property;
    private $private_property;

    function methodName()
    {
        // do something
    }
}
```

This is the format for class declarations in object-oriented programming. Inheritance, interfaces, and traits are, of course, optional.

Details on method declarations and access modifiers are omitted.

# Static Properties and Static Methods

```php
<?php

class StaticClass
{
    static $property = 'This is a static property.';

    static function methodA()
    {
        return self::$property;
    }

    public function methodB()
    {
        return self::methodA();
    }
}

echo StaticClass::methodB(); // This a is static property.
```

In reality, you probably wouldn't do something this complicated... The point was to explain that static properties and static methods can be called using the scope resolution operator (::) and the self keyword.

# Inheritance (extends)

```php
<?php

class SuperClass
{
    public function superClassMethod()
    {
        echo 'This is a super class method.';
    }
}

class SubClass extends SuperClass
{
    public function subClassMethod()
    {
        echo 'This is a sub class method.';
    }
}

$subClassObject = new SubClass();

echo $subClassObject->subClassMethod(); // This is sub class method.
echo $subClassObject->superClassMethod(); // This is a super class method.
```

To explicitly use a parent class's method in a derived class (subclass), you can use the parent keyword. **You cannot inherit multiple classes (i.e., multiple inheritance) with the extends keyword.**

# Abstract Methods (abstract)

```php
<?php

abstract class Human
{
    abstract function getAbility();

    public function run()
    {
        echo 'Run';
    }
}

class SuperHuman extends Human
{
    public function getAbility()
    {
        echo 'Fly';
    }
}

$super_human_instance = new SuperHuman();
echo $super_human_instance->run(); // Run
echo $super_human_instance->getAbility(); // Fly
```

Abstract classes, like interfaces, require all defined abstract methods to be implemented in derived classes. Unlike interfaces, abstract classes can have implementations. The distinction between abstract classes and interfaces is well explained in [Understanding and Using PHP Interfaces and Abstracts Correctly](https://havelog.ayumusato.com/develop/php/e166-php-interface-abstract.html).

# Interfaces

```php
<?php

interface Human
{
    public function eat();

    public function sleep();

    public function walk();
}

class Boy implements Human
{
    public function eat()
    {
        echo 'Eat';
    }

    public function sleep()
    {
        echo 'Sleep';
    }

    public function walk()
    {
        echo 'Walk';
    }

    public function fly()
    {
        echo 'Fly';
    }
}

$boy_instance = new Boy();
echo $super_human_instance->eat(); // Eat.
echo $super_human_instance->walk(); // Walk.
echo $super_human_instance->sleep(); // Sleep.
echo $super_human_instance->fly(); // Fly.
```

Interfaces define the behavior (methods only) of a class but do not implement them. Abstract classes define and implement parts of a base class, while interfaces require all defined methods to be implemented.

Pseudo-multiple inheritance using interfaces is possible.

# Traits

```php
<?php

trait Authorize
{
    public function register()
    {
        echo 'Registration';
    }
}

class User
{
    use Authorize;
}

$user_instance = new User();
echo $user_instance->register();
```

Traits allow for code reuse across class hierarchies. Using traits enables multiple inheritance.

# Summary
In this post, we explained the intention of learning design patterns and the basics of object-oriented programming. From next time, we will introduce each design pattern.

# Thoughts
Even if you understand the mechanism, actually using (designing) object-oriented programming is difficult, and it seems necessary to repeatedly practice and relearn.