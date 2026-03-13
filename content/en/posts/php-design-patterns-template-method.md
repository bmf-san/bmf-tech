---
title: Learning Design Patterns with PHP - Template Method ~Filling in the Blanks~
description: 'Discover the Template Method pattern to aggregate common processes in superclasses and abstract specific subclass implementations.'
slug: php-design-patterns-template-method
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Design Pattern
  - PHP
  - Template Method Pattern
translation_key: php-design-patterns-template-method
---

# What is the Template Method Pattern?
The Template Method Pattern is a pattern where similar processes are defined as a framework (type) in a superclass, and more specific processing content is implemented in subclasses. It is not mere inheritance, but a pattern that uses inheritance to ensure the implementation of the superclass's methods by defining specific processing content as abstract methods, allowing the behavior of the class to be defined by the subclass.

# Structure
## AbstractClass
A class that defines the framework of processing, including methods (template methods) that define the framework and methods that utilize it.

## ConcreteMethod
A subclass that inherits from AbstractClass and implements the abstract methods defined in the AbstractClass.

# Advantages
## Ability to Aggregate Common Processes
Since necessary processes for each subclass can be aggregated in the superclass, the common implementation parts in the subclasses are reduced.

# Disadvantages
* The number of subclasses tends to increase
* If the logic of the superclass is large, the freedom of the subclass decreases

## Ability to Change Specific Processing Content in Subclasses
By defining the framework in the superclass, the specific processing content can be flexibly implemented in the subclasses.

# When to Use
"Hmm... I feel like I've created a similar class before..."
In such cases, you might be able to apply the pattern by extracting methods that can be commonized.

# Implementation Example (※[github](https://github.com/bmf-san/design-patterns-php) repository available.)

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

$data = [
    "title" => "What is the Template Method?",
    "author" => "Qiita Tarou."
];

$corporate_article = new CorporateArticle($data);
$user_article = new UserArticle($data);

echo $corporate_article->display();
```

```
// Output
Title:What is the Template Method?
Author:Qiita Tarou.
Content:This is a Corporate Article. Here write your things.
```

# Summary
* Aggregate common processes in the superclass
* Define specific processes as abstract methods in the superclass and ensure implementation in subclasses

# Related Keywords
* Liskov Substitution Principle (LSP: The Liskov Substitution Principle)
* Hollywood Principle (The Hollywood Principle: Don't call us. We'll call you.)