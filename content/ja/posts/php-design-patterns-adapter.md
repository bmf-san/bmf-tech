---
title: "PHPで学ぶデザインパターン - Adapterパターン"
slug: "php-design-patterns-adapter"
date: 2019-02-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "アダプターパターン"
  - "GoF"
  - "PHP"
  - "デザインパターン"
draft: false
---

# 概要
[PHPで学ぶデザインパターン Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern)で間に合わなかった記事。

# Adaperパターンとは
元となるクラスに修正を加えることなくインターフェースを変更することができるパターン。

異なるインターフェース間の互換性を調整するようなAdapterクラスを用意することで実現する。

# 実装
```php
<?php
interface Bird
{
    public function fly();
}

class SmallBird implements Bird
{
    public function fly()
    {
        echo 'fly short time';
    }
}

class BigBird implements Bird
{
    public function fly()
    {
        echo 'fly long time';
    }
}

class Human
{
    public function eat(Bird $bird)
    {
        echo 'Yummy!';
    }
}

class MiddleBird
{
    public function jump()
    {
        echo 'jump like flying';
    }
}

// Adapter
class MiddleBirdAdapter implements Bird
{
    private $middleBird;

    public function __construct(MiddleBird $middleBird)
    {
        $this->middleBird = $middleBird;
    }

    // MiddleBirdのメソッドをラップする
    public function fly()
    {
        $this->middleBird->jump();
    }
}

$human = new Human();

$smallBird = new SmallBird();
$human->eat($smallBird); // Yummy

$middleBird = new MiddleBird();
$middleBirdAdapter = new MiddleBirdAdapter($middleBird);

$human->eat($middleBirdAdapter); // Yummy
```

# 所感
インターフェースに定義されたメソッドで振る舞いをラップするメソッドをつくる感じ。
使い所を慎重に考える必要がありそう。

# 参考
- [[保存版]人間が読んで理解できるデザインパターン解説#2: 構造系（翻訳）](https://techracho.bpsinc.jp/hachi8833/2017_10_11/46069#adapter)
- [Wikipedia - Adapterパターン](https://ja.wikipedia.org/wiki/Adapter_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)

