---
title: 'Learning Design Patterns with PHP - Singleton: Limiting Instances'
description: 'Understand the Singleton pattern to control instance creation, restrict access, and manage single object state in PHP.'
slug: php-design-patterns-singleton
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Design Pattern
  - PHP
  - Singleton Pattern
translation_key: php-design-patterns-singleton
---



# What is the Singleton Pattern?
This pattern ensures that only one instance exists to control the cost of instance creation.

# Structure
## SingletonClass
Simply prepare a private constructor, a static method that returns only one instance, and a static variable to hold its own instance.

# Advantages
## Control Access to Instances
Since access to itself held by the Singleton pattern is restricted to private, you can control access from client-side code.

## Ability to Change the Number of Instances
It is also possible to change the number of generated instances to two or more.

# Disadvantages
* Decreased testability
* Difficult to flexibly respond if changes to the number of controlled instances are needed after implementation

# When to Use

# Implementation Example (※Repository available on ~~github~~.)

```SingletonConfig.php
<?php
class SingletonConfig {
    private $config;

    /**
     * a single variable
     */
    private static $instance;

    private function __construct()
    {
        $this->config = 'AUTO';
    }

    /**
     * Create a only instance
     */
    public static function getInstance()
    {
        if (!isset(self::$instance)) {
            self::$instance = new SingletonConfig();
        }

        return self::$instance;
    }

    public function getConfig()
    {
        return $this->config;
    }

    public final function __clone()
    {
        throw new RuntimeException('Clone is not allowed against' . get_class($this));
    }
}
```

```singleton_client.php
<?php
require_once 'SingletonConfig.php';

$instanceA = SingletonConfig::getInstance();

$instanceB = SingletonConfig::getInstance();

if ($instanceA->getConfig() === $instanceB->getConfig()) {
    echo 'True';
} // true
```

# Summary
* Restrict external access by making the instance creation method private
* Control the number of instance creations with a static method

# Related Keywords
* Encapsulation
