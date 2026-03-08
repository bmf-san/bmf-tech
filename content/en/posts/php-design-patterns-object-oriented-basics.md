---
title: Learning Design Patterns with PHP - Basics of Object-Oriented Programming
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

Using the now out-of-print [Introduction to Design Patterns in PHP](https://www.amazon.co.jp/PHP%E3%81%AB%E3%82%88%E3%82%8B%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E4%B8%8B%E5%B2%A1-%E7%A7%80%E5%B9%B8/dp/4798015164) as a textbook, we will learn design patterns in PHP. (Note: There are used copies available on Amazon, but they seem to be priced at more than double the original price.)

The code discussed in this series will be compiled on [GitHub](https://github.com/bmf-san/design-patterns-php).

Ideally, I would have liked to learn design patterns in a language that pioneered OOP, but due to my lack of knowledge in languages other than PHP and the encounter with a book explaining design patterns in PHP, I decided to learn design patterns using PHP.

# References
* [Introduction to Design Patterns in PHP](https://www.amazon.co.jp/PHP%E3%81%AB%E3%82%88%E3%82%8B%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E4%B8%8B%E5%B2%A1-%E7%A7%80%E5%B9%B8/dp/4798015164)
* [Do You PHP Hatena](http://d.hatena.ne.jp/shimooka/20141211/1418298136)
* [Github shimooka/PhpDesignPattern](https://github.com/shimooka/PhpDesignPattern)
* [Perfect PHP](https://www.amazon.co.jp/%E3%83%91%E3%83%BC%E3%83%95%E3%82%A7%E3%82%AF%E3%83%88PHP-PERFECT-3-%E5%B0%8F%E5%B7%9D-%E9%9B%84%E5%A4%A7/dp/4774144371)
* [Programming PHP 3rd Edition](https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0PHP-%E7%AC%AC3%E7%89%88-Kevin-Tatroe/dp/4873116686/ref=sr_1_1?ie=UTF8&qid=1479547951&sr=8-1&keywords=%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0PHP)

References for each pattern will be included in their respective articles.

# Why Learn?

There are two main purposes for learning design patterns.

**To deepen understanding of frameworks**

Design patterns are utilized throughout the design of frameworks. By knowing these patterns, I believe one can deepen their understanding of frameworks and maximize their use.

**To write quality code**

Design patterns can be described as the culmination of object-oriented principles established by predecessors. While it may be difficult to directly adopt and use design patterns, learning the "mindset" can deepen understanding of object-oriented programming and improve the quality of daily source code.

# Basics of Object-Oriented Programming

Before diving into the main topic, let's review the basics of object-oriented programming.

# Notes
The book [Introduction to Design Patterns in PHP](https://www.amazon.co.jp/PHP%E3%81%AB%E3%82%88%E3%82%8B%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E4%B8%8B%E5%B2%A1-%E7%A7%80%E5%B9%B8/dp/4798015164) was published in **2006, so the understanding of the usage of patterns may have changed from then to now**. Please understand that this point cannot be addressed in the initial draft.

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

This is the format for class declarations in object-oriented programming. Inheritance, interfaces, and traits are optional.

Details about method declarations and access modifiers are omitted.

# Static Properties and Methods

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

While I don't think such complicated things are usually done... I wanted to summarize that static properties and methods can be called using the scope resolution operator (::) and the self keyword.

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

To explicitly use a parent class method in a derived class (subclass), you can use the parent keyword. **You cannot inherit multiple classes using the extends keyword (i.e., multiple inheritance).**

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

Abstract classes, like interfaces, require all defined abstract methods to be implemented in derived classes. While interfaces cannot contain implementations, abstract classes can. The distinction between abstract classes and interfaces is well explained in [Understanding and Distinguishing PHP's interface and abstract](https://havelog.ayumusato.com/develop/php/e166-php-interface-abstract.html).

# Interface

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

Interfaces only define the behavior (methods) of a class without implementing them. Abstract classes define and implement parts of a base class, while interfaces must implement all defined methods.

Using interfaces allows for pseudo-multiple inheritance.

# Trait

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

Traits enable code reuse across class hierarchies. Additionally, using traits allows for multiple inheritance.

# Conclusion
In this article, I explained the intention behind learning design patterns and the basics of object-oriented programming. In the next installment, I will introduce various design patterns.

# Thoughts
Even if I understand the mechanisms, actually using them (designing) in object-oriented programming seems difficult, and it appears necessary to repeatedly practice to relearn.