---
title: Design Patterns in PHP - Bridge Pattern
slug: php-design-patterns-bridge
date: 2019-02-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - GoF
  - PHP
  - Design Patterns
  - Bridge Pattern
translation_key: php-design-patterns-bridge
---

# Overview
This is an article that couldn't make it in time for the [Design Patterns in PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# What is the Bridge Pattern?
A pattern that prepares a superclass for functional extension and subclasses for implementation extension, acting as a bridge for functionality.

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

// Either converter can be used → implementation can be switched
$iphoneConnector = new IphoneConnector($lightingConveter);
$androidConnector = new AndroidConnector($typeCConverter);

$iphoneConnector->connect(); // connect by using lighting
$androidConnector->connect(); // connect by using type-c
```

# Thoughts
I feel that this might not be a simple use of interfaces, which suggests my understanding is still shallow.

# References
- [[Preserved Edition] Human-readable Design Pattern Explanation #2: Structural Patterns (Translation)](https://techracho.bpsinc.jp/hachi8833/2017_10_11/46069#bridge