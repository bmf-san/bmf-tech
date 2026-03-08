---
title: PHPで学ぶデザインパターン - Mediatorパターン
slug: php-design-patterns-mediator
date: 2019-01-31T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - PHP
  - デザインパターン
  - メディエーターパターン
  - GoF
translation_key: php-design-patterns-mediator
---


# 概要
[PHPで学ぶデザインパターン Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern)で間に合わなかった記事。

# Mediatorパターンとは
仲介者・調停者の意味。

オブジェクトの振る舞いに関するデザインパターンで、オブジェクト間のやりとり調整するためのパターン。

オブジェクト同士のやり取りが複雑化し、関係性が見えにくくなるような時に有用かもしれない。

# 実装

```php
<?php

// Mediator
class Receptionist
{
    public function checkIn(User $user, $message) // 振る舞いの操作を任せたいオブジェクトを保持
    {
        echo $message . ' ' . $user->getName();
    }
}

class User
{
    private $name;
    private $receptionist;

    public function __construct($name, Receptionist $receptionist) // Mediatorを持つ
    {
        $this->name = $name;
        $this->receptionist = $receptionist;
    }

    public function getName()
    {
        return $this->name;
    }

    public function checkIn($message)
    {
        $this->receptionist->checkIn($this, $message); // $this!!
    }
}

$receptionist = new Receptionist();


$john = new User('John', $receptionist);
$bob = new User('Bob', $receptionist);

$john->checkIn('Welcome!'); // Welcome! John
$bob->checkIn('Hi!'); // Hi! Bob
```

# 所感
クラス間のやりとりが複雑化しそうなときにオブジェクトの振る舞いをまとめて管理したいときに思い出したいパターン。

# 参考
- [PHPによるデザインパターン入門 - Mediator～すべては相談役が知っている](http://d.hatena.ne.jp/shimooka/20141217/1418788236)
- [[保存版]人間が読んで理解できるデザインパターン解説#3: 振舞い系（翻訳）](https://techracho.bpsinc.jp/hachi8833/2017_10_17/46071#mediator)

