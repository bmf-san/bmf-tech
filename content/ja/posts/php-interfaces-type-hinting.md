---
title: PHPのインターフェースとタイプヒンティング
slug: php-interfaces-type-hinting
date: 2018-12-08T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - PHP
  - インターフェース
  - タイプヒンティング
translation_key: php-interfaces-type-hinting
---


# 概要
この記事は[PHP Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php)の記事です。（ちょっと早めに投稿しています）

インターフェースはメソッドの実装を保証する”契約”的意味合いの他、タイプヒンティングによって実装を抽象に依存させる（=実装の切り替えをしやすくする）こともできる。

# インターフェースの定義・実装
基本的なインターフェースの定義と実装。

```php
<?php
interface Action
{
    public function say();
}

class Superman implements Action
{
    public function say()
    {
        echo "Hello World";
    }
}

$obj = new Superman();
$obj->say();
```

# インターフェースによる機能と実装の分離
タイプヒンティングでインターフェース型を指定すると実装に柔軟性を持たせることができる。

```php
<?php
interface HeroAction
{
    public function say();
}

class Superman implements HeroAction
{
    public function say()
    {
        echo "I'm a Superman";
    }
}

class Human
{
    public function say()
    {
        echo "I'm a Human";
    }
}

class Bot
{
    public function do(HeroAction $heroAction) // 引数にインターフェース型を指定
    {
        $heroAction->say();
    }
}

$superMan = new SuperMan();
$human = new Human();
$bot = new Bot();

$bot->do($superMan); // I'm a Superman
$bot->do($human); // PHP Fatal error:  Uncaught TypeError: Argument 1 passed to Bot::do() must implement interface HeroAction, instance of Human given, called in ....
```

Supermanの実装を取りやめて、Hypermanの実装に切り替える。

```php
<?php
interface HeroAction
{
    public function say();
}

// class Superman implements HeroAction
// {
//     public function say()
//     {
//         echo "I'm a Superman";
//     }
// }

class Hyperman implements HeroAction
{
    public function say()
    {
        echo "I'm a Hyperman";
    }
}

class Human
{
    public function say()
    {
        echo "I'm a Human";
    }
}

class Bot
{
    public function do(HeroAction $heroAction) // 引数にインターフェース型を指定
    {
        $heroAction->say();
    }
}

// $superMan = new SuperMan();
$hyperMan = new HyperMan();
$human = new Human();
$bot = new Bot();

// $bot->do($superMan); // I'm a Superman
$bot->do($hyperMan); // I'm a Hyperman
$bot->do($human); // PHP Fatal error:  Uncaught TypeError: Argument 1 passed to Bot::do() must implement interface HeroAction, instance of Human given, called in ....
```

もし、Botクラスのdoメソッドがインターフェースではなく、Supermanクラスに依存していた場合、実装を交換する手間が増えてしまう。
```php
class Bot
{
    public function do(Superman $superman) // 引数にインターフェース型を指定
    {
        $superman->say();
    }
}
```

# 参考
- [PHP - オブジェクト インターフェイス](http://php.net/manual/ja/language.oop5.interfaces.php)
- [PHPにはインターフェイスというものがありますよ、という話](http://blog.anatoo.jp/entry/20080517/1211029059)

