---
title: Design Patterns in PHP - Singleton ~Limiting the Number of Instances~
slug: php-design-patterns-singleton
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Design Patterns
  - PHP
  - Singleton
translation_key: php-design-patterns-singleton
---

# What is the Singleton Pattern?
A pattern that guarantees there is only one instance to control the cost of instance creation.

# Structure
## SingletonClass
It simply prepares a private constructor, a static method that returns only one instance, and a static variable to hold its own instance.

# Advantages
## Control Access to the Instance
The Singleton pattern restricts access to itself as private, allowing control over access from client-side code.

## Ability to Change the Number of Instances
It is also possible to change the number of instances created to two or more.

# Disadvantages
* Reduced testability
* Difficult to adapt flexibly if changes to the number of controlled instances are needed after implementation.

# Use Cases

# Implementation Example (Repository available on [github](https://github.com/bmf-san/design-patterns-php))

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
* Make the instance creation method private to restrict access from outside.
* Control the number of instance creations using a static method.

# Related Keywords
* Encapsulation