---
title: Learning Design Patterns with PHP - State Pattern
slug: php-design-patterns-state
date: 2019-04-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - GoF
  - PHP
  - Design Patterns
  - State Pattern
translation_key: php-design-patterns-state
---

# Overview
This is an article that couldn't make it in time for the [Learning Design Patterns with PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern).

# What is the State Pattern?
A pattern that allows behavior to change by preparing states as classes.

# Implementation
I imagined an example that toggles the state of a switch on and off. It seems better to use a singleton.

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
        // Set default state class
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

The state classes have an image of having state-specific processes.

I think this pattern can be considered when there are multiple states and the specific processes become complex.

# Thoughts
For some reason, this is my favorite pattern among design patterns. It feels interesting and makes me want to try using it.

# References
- [State Pattern - PHP Design Patterns](https://www.ritolab.com/entry/140)