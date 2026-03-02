---
title: "PHPで学ぶデザインパターン - Template Method ~処理の穴埋め~"
slug: "php-template-method"
date: 2017-10-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "デザインパターン"
  - "PHP"
  - "テンプレートメソッドパターン"
draft: false
---

# Template Method パターンとは？
似たような処理を枠組み（型）としてスーパークラスで定義し、より具体的な処理内容をサブクラスで実装するというパターンです。
単なる継承ではなく、具体的な処理内容を抽象メソッドとして定義することで、**スーパークラスのメソッドの実装を保証し、クラスの振る舞いをサブクラスによって定義させる継承を利用したパターン**です。

# 構造
## AbstractClass
処理の枠組みを定義するクラスで、枠組みを定義するメソッド（template method）とそれを利用するメソッドを含みます。

## ConcreteMethod
AbstractClassを継承するサブクラスで、Abstractクラスで定義された抽象メソッドを実装します。

# メリット
## 共通処理を集約できる
各サブクラスで必要な処理をスーパークラスに集約できるため、サブクラスでの共通実装部分が減ります。

# デメリット
* サブクラスの数が増加しやすい
* スーパークラスのロジックが大きいとサブクラスの自由度が下がる

## 具体的な処理内容をサブクラスで変更できる
大枠をスーパークラスで定義することにより、具体的な処理内容は柔軟にサブクラスで実装することができます。

# 使いドコロ
「ん・・・以前にも同じようなクラスをつくったような・・・」
そんな時は共通化できるメソッドを抜き出すことでパターンを適用できるかもしれません。

# 実装例（※[github](https://github.com/bmf-san/design-patterns-php)にリポジトリあります。）

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
//　出力
Title:What is the Template Method?
Author:Qiita Tarou.
Content:This is a Corporate Article. Here write your things.
```

# まとめ
* 共通処理をスーパークラスに集約
* 具体的な処理は抽象メソッドとしてスーパークラスに定義、サブクラスで実装させることを保証

# 関連キーワード
* リスコフの置換原則（LSP：The Liskov Substituion Principle）
* ハリウッドの原則(The Hollywood Principle:Don't call us. We'll call you.)

