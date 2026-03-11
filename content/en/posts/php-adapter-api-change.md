---
title: Learning Design Patterns with PHP - Adapter ~Modifying APIs~
slug: php-adapter-api-change
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
  - Adapter Pattern
  - Design Pattern
translation_key: php-adapter-api-change
---



# What is the Adapter Pattern?
The Adapter Pattern is used to make incompatible interfaces compatible with each other. It allows you to **provide new functionality by reusing existing code without modifying it**. The key feature is that you do not change the code you are reusing. This pattern mainly arises from the need to reuse code for practical reasons. (The pattern of preparing a wrapper at the design stage is the Bridge Pattern.)

# Structure
## TargetClass
Defines the API (interface).

## AdapteeClass
Provides the existing API that is adapted to the TargetClass.

## AdapterClass
Converts the API of the AdapteeClass so that it can be used from the TargetClass.

# Benefits
## Reuse existing code without modification
Since it is implemented by wrapping existing classes, there is no need to modify the existing code.

## Saves the client from having to be aware of the existing API implementation
In short, changes to the existing API do not affect changes on the client side.

## Freely restrict the API to be exposed
You can restrict access to the API when adapting it.

## Drawbacks
* Increasing the layers of adaptation may affect performance

# When to Use
When you want to reuse existing, proven classes.

# Implementation Example (※Repository available on [github](https://github.com/bmf-san/design-patterns-php).)

## Pattern Using Inheritance
```ShowData.php
<?php

class ShowData {
    private $data;

    public function __construct($data)
    {
        $this->data = $data;
    }

    public function showOriginalData()
    {
        echo $this->data;
    }

    public function showProcessedData()
    {
        echo $this->data . 'How are you?';
    }
}
```

```ShowSourceData.php
<?php

interface ShowSourceData {
    public function show();
}
```

```ShowSourceDataImpl.php
<?php
require_once 'ShowSourceData.php';
require_once 'ShowData.php';

class ShowSourceDataImpl extends ShowData implements ShowSourceData {
    public function __construct($data)
    {
        parent::__construct($data);
    }

    public function show()
    {
        parent::showProcessedData();
    }
}
```

```adapter_client.php
<?php
require_once 'ShowSourceDataImpl.php';

$show_data = new ShowSourceDataImpl('Hello! Mr. Data.');

$show_data->show();
```

## Pattern Using Delegation

The wrapper part is different, but the client-side code is the same. Delegation means **entrusting specific processing to another class**. Is it misleading to say it's like DI... (゜-゜)


```ShowSourceDataImpl.php
<?php
require_once '../ShowSourceData.php';
require_once '../ShowData.php';

class ShowSourceDataImpl implements ShowSourceData {
    private $show_data;

    public function __construct($data)
    {
        $this->show_data = new ShowData($data);
    }

    public function show()
    {
        $this->show_data->showProcessedData();
    }
}
```

# Summary
* Prepare a wrapper class to indirectly reuse existing code.
* There are inheritance-based and delegation-based wrapper classes.

# Related Keywords
* Bridge Pattern
