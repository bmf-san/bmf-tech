---
title: Learning Design Patterns with PHP - Bridge Pattern
description: An in-depth exploration of Learning Design Patterns with PHP - Bridge Pattern, covering design principles, trade-offs, and practical applications.
slug: php-design-patterns-bridge
date: 2019-02-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - GoF
  - PHP
  - Design Pattern
  - Bridge Pattern
translation_key: php-design-patterns-bridge
---

# Overview
An article that didn't make it in time for the [PHP Design Patterns Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# What is the Bridge Pattern?
A pattern that prepares a superclass for functional extensions and a subclass for implementation extensions, acting as a bridge for functionality.

# Implementation

```php
<?php
interface Connector
{
    public function __construct(Converter $converter);
    public function connect();
}

class IphoneConnector implements Connector
{
    private $converter;

    public function __construct(Converter $converter)
    {
        $this->converter = $converter;
    }

    public function connect()
    {
        echo 'Iphone connect by using ' . $this->converter->getTerminalName();
    }
}

class AndroidConnector implements Connector
{
    private $converter;

    public function __construct(Converter $converter)
    {
        $this->converter = $converter;
    }

    public function connect()
    {
        echo 'Android connect by using ' . $this->converter->getTerminalName();
    }
}

interface Converter
{
    public function getTerminalName();
}

class LightningConverter implements Converter
{
    public function getTerminalName()
    {
        return 'lightning';
    }
}

class TypeCConverter implements Converter
{
    public function getTerminalName()
    {
        return 'type-c';
    }
}

$lightingConveter = new LightningConverter();
$typeCConverter = new TypeCConverter();

// Either converter can be used → Implementation can be switched
$iphoneConnector = new IphoneConnector($lightingConveter);
$androidConnector = new AndroidConnector($typeCConverter);

$iphoneConnector->connect(); // connect by using lighting
$androidConnector->connect(); // connect by using type-c
```

# Thoughts
I feel like this is just a simple use of interfaces, which might indicate my understanding is still shallow.

# References
- [[Definitive Edition] Design Patterns Explained in a Way Humans Can Understand #2: Structural Patterns (Translation)](https://techracho.bpsinc.jp/hachi8833/2017_10_11/46069#bridge
