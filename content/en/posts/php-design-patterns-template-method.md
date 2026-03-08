---
title: Learning Design Patterns with PHP - Template Method ~Filling in the Gaps~
slug: php-design-patterns-template-method
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Design Patterns
  - PHP
  - Template Method
translation_key: php-design-patterns-template-method
---

# What is the Template Method Pattern?
This pattern defines a framework (template) for similar processes in a superclass, while the specific processing details are implemented in subclasses. It is a pattern that utilizes inheritance by defining specific processing details as abstract methods, ensuring the implementation of the superclass's methods and allowing the behavior of the class to be defined by the subclasses.

# Structure
## AbstractClass
This class defines the framework for processing and includes the method that defines the framework (template method) and methods that utilize it.

## ConcreteMethod
A subclass that inherits from AbstractClass and implements the abstract methods defined in the Abstract class.

# Advantages
## Common Processing Can Be Aggregated
Since necessary processing for each subclass can be aggregated in the superclass, the common implementation parts in the subclasses are reduced.

# Disadvantages
* The number of subclasses can easily increase.
* If the logic in the superclass is large, the flexibility of subclasses decreases.

## Specific Processing Details Can Be Changed in Subclasses
By defining the broad framework in the superclass, specific processing details can be flexibly implemented in subclasses.

# When to Use
"Hmm... I feel like I've created a similar class before..."
In such cases, you might be able to apply the pattern by extracting methods that can be common.

# Implementation Example (Repository available on [github](https://github.com/bmf-san/design-patterns-php))

```AbstractArticle.php
<?php

abstract class AbstractArticle {

    public function __construct($data)
    {
        $this->title = $data['title'];
        $this->author = $data['author'];
    }

    /**
     * Template Method
     */
    public function display()
    {
        return "Title:{$this->getTitle()}<br />Author:{$this->getAuthor()}<br />Content:{$this->getContent()}";
        $this->getTitle();
        $this->getAuthor();
        $this->getContent();
    }

    /**
     * Common Method
     */
    public function getTitle()
    {
        return $this->title;
    }

    /**
     * Common Method
     */
    public function getAuthor()
    {
        return $this->author;
    }

    /**
     * Abstract Method
     */
    protected abstract function getContent();
}
```

```AbstractCorporateArticle.php
<?php
require_once 'AbstractArticle.php';

/**
 * Concrete Class
 */
class CorporateArticle extends AbstractArticle {

    protected function getContent()
    {
        return 'This is a Corporate Article. Here write your things.';
    }
}
```

```AbstractUserArticle.php
<?php
require_once 'AbstractArticle.php';

/**
 * Concrete Class
 */
class UserArticle extends AbstractArticle {

    protected function getContent()
    {
        return 'This is a User Article. Here write your things.';
    }
}
```

```template_method_client.php
<?php
require_once 'AbstractCorporateArticle.php';
require_once 'AbstractUserArticle.php';

data = [
    "title" => "What is the Template Method?",
    "author" => "Qiita Tarou."
];

$corporate_article = new CorporateArticle($data);
$user_article = new UserArticle($data);

echo $corporate_article->display();
```

```// Output
Title:What is the Template Method?
Author:Qiita Tarou.
Content:This is a Corporate Article. Here write your things.
```

# Summary
* Aggregate common processing in the superclass.
* Define specific processing as abstract methods in the superclass and ensure implementation in subclasses.

# Related Keywords
* Liskov Substitution Principle (LSP)
* The Hollywood Principle: Don't call us. We'll call you.