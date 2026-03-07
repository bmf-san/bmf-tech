---
title: Learning Design Patterns with PHP - Adapter ~Changing APIs~
slug: php-adapter-api-change
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
  - Adapter Pattern
  - Design Patterns
translation_key: php-adapter-api-change
---

# What is the Adapter Pattern?
The Adapter Pattern is used to make incompatible APIs (interfaces) compatible. **It provides new functionality by reusing existing code without modifying it.** The key feature is that the reused code remains unchanged. This pattern is primarily established for the purpose of reusing code. (The pattern that prepares a wrapper at the design stage is the Bridge Pattern.)

# Structure
## TargetClass
Defines the API (interface).

## AdapteeClass
Provides the existing API that will be adapted to the TargetClass.

## AdapterClass
Transforms the AdapteeClass API so that it can be used from the TargetClass.

# Advantages
## Reuse existing code without modification
Since it is implemented by wrapping existing classes, there is no need to modify the existing code.

## Reduces the burden on the client side to be aware of the existing API implementation
In short, changes to the existing API do not affect changes on the client side.

## Freedom to restrict the public API
You can limit access to the API when adapting it.

# Disadvantages
* Increasing the number of layers to adapt may impact performance.

# When to Use
When you want to reuse existing, proven classes.

# Implementation Example (Repository available on [github](https://github.com/bmf-san/design-patterns-php))

## Pattern using Inheritance
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

## Pattern using Delegation
The only difference is in the wrapper part, while the client-side code remains the same. Delegation means **entrusting specific processing to another class**. It might be misleading to say it’s like DI...

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
* The wrapper class can be based on inheritance or delegation.

# Related Keywords
* Bridge Pattern