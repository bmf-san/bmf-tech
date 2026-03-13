---
title: Learning Design Patterns with PHP - State Pattern
description: 'Learn the State pattern to switch behavior dynamically by encapsulating state-specific processing in separate classes.'
slug: php-design-patterns-state
date: 2019-04-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - GoF
  - PHP
  - Design Pattern
  - State Pattern
translation_key: php-design-patterns-state
---



# Overview
An article that didn't make it in time for the [PHP Design Patterns Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# What is the State Pattern?
A pattern that allows behavior to be switched by preparing states as classes.

# Implementation
An example assuming switching the on/off state of a switch. It might be better to use Singleton.

```php
<?php

class OnState
{
    public function getState()
    {
        return 'ON';
    }

    public function getNextState()
    {
        return new OffState();
    }
}

class OffState
{
    public function getState()
    {
        return 'OFF';
    }

    public function getNextState()
    {
        return new OnState();
    }
}

class Light
{
    public function __construct()
    {
        // Set the default state class
        $this->state = new OffState();
    }

    public function getState()
    {
        return $this->state->getState();
    }

    public function toggle()
    {
        $this->state = $this->state->getNextState();
    }
}


$light = new Light();

echo $light->getState(); // OFF
echo $light->toggle();
echo $light->getState(); // ON
echo $light->toggle();
echo $light->getState(); // OFF
```

The state class has an image of holding state-specific processing.

It seems to be a pattern that can be considered when there are multiple states and each has complex specific processing.

# Impressions
For some reason, this is my favorite pattern among design patterns. It feels like it has an interesting appeal that makes you want to use it.

# References
- [State Pattern - PHP Design Patterns](https://www.ritolab.com/entry/140)
