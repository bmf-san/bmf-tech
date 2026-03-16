---
title: PHPで学ぶデザインパターン - オブジェクト指向の基礎
description: PHPで学ぶデザインパターン - オブジェクト指向の基礎
slug: php-design-patterns-object-oriented-basics
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - デザインパターン
  - PHP
  - OOP
translation_key: php-design-patterns-object-oriented-basics
---


# 概要

今は絶版になっている[PHPによるデザインパターン入門](https://www.amazon.co.jp/PHP%E3%81%AB%E3%82%88%E3%82%8B%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E4%B8%8B%E5%B2%A1-%E7%A7%80%E5%B9%B8/dp/4798015164)を教科書にして、PHPでデザインパターンを学びます。（※Amazonで中古がありますが、定価の倍以上の値段が付いているようです。）

本連載で扱うコードは[github](https://github.com/bmf-san/design-patterns-php)にまとめていきます。

本来であれば、OOPの先駆けである言語でデザインパターンを学びたいところでしたが、PHP以外の言語の素養がなかったことと、PHPでデザインパターンを解説している本に出会ったことから、PHPでデザインパターンを学んでみることにしました。


# 参考
　* 　[PHPによるデザインパターン入門](https://www.amazon.co.jp/PHP%E3%81%AB%E3%82%88%E3%82%8B%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E4%B8%8B%E5%B2%A1-%E7%A7%80%E5%B9%B8/dp/4798015164)
　* 　[Do You PHP はてな](http://d.hatena.ne.jp/shimooka/20141211/1418298136)
　* 　[Github shimooka/PhpDesignPattern](https://github.com/shimooka/PhpDesignPattern)
　* 　[パーフェクトPHP](https://www.amazon.co.jp/%E3%83%91%E3%83%BC%E3%83%95%E3%82%A7%E3%82%AF%E3%83%88PHP-PERFECT-3-%E5%B0%8F%E5%B7%9D-%E9%9B%84%E5%A4%A7/dp/4774144371)
　* 　[プログラミングPHP 第3版](https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0PHP-%E7%AC%AC3%E7%89%88-Kevin-Tatroe/dp/4873116686/ref=sr_1_1?ie=UTF8&qid=1479547951&sr=8-1&keywords=%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0PHP)

パターンごとの参考は、それぞれの記事に記載します。

# なぜ学ぶのか

デザインパターンを学ぶ目的は2つあります。

**フレームワークへの理解を深めるため**

FWの設計には随所にデザインパターンが活用されています。パターンを知ることで、FWへの理解が深まり、FWを最大限に活用することができると思います。

**良質なコードをかくため**

デザインパターンは先人たちが築いたオブジェクト指向の集大成と形容することができると思います。デザインパターンをいきなり取り入れて使うことは難しいかもしれませんが、"考え方"を学ぶことでオブジェクト指向への理解が深まり、日々のソースコードの質を上げることにつながるのではないかと思います。

# オブジェクト指向基本

本題に入る前に、オブジェクト指向の基本について復習しておきます。

# 諸注意
[PHPによるデザインパターン入門](https://www.amazon.co.jp/PHP%E3%81%AB%E3%82%88%E3%82%8B%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E4%B8%8B%E5%B2%A1-%E7%A7%80%E5%B9%B8/dp/4798015164)は**2006年に発行された本のため、当初と現在ではパターンの用途について考え方が変わってきている部分もある**かと思います。
それについては初稿段階において触れることができない点だけご了承ください。


# クラス宣言

```php
<?php

class ClassName extends BaseClassName implements InterfaceName
{
    use TraitName;

    public $public_property;
    protected $protected_property;
    private $private_property;

    function methodName()
    {
        // do something
    }
}
```

オブジェクト指向におけるクラス宣言のフォーマットです。
継承やインターフェース、トレイトはもちろん任意指定です。

メソッド宣言やアクセス修飾子の詳細については割愛します。


# 静的プロパティ・静的メソッド

```php
<?php

class StaticClass
{
    static $property = 'This is a static property.';

    static function methodA()
    {
        return self::$property;
    }

    public function methodB()
    {
        return self::methodA();
    }
}

echo StaticClass::methodB(); // This a is static property.
```

実際はこんなややこしいことはしないとは思いますが・・・
静的プロパティや静的メソッドは、スコープ定義演算子（::）やselｆキーワードで呼び出しが行えるということを一纏めに説明したかった次第です。


# 継承（extends）

```php
<?php

class SuperClass
{
    public function superClassMethod()
    {
        echo 'This is a super class method.';
    }
}

class SubClass extends SuperClass
{
    public function subClassMethod()
    {
        echo 'This is a sub class method.';
    }
}

$subClassObject = new SubClass();

echo $subClassObject->subClassMethod(); // This is sub class method.
echo $subClassObject->superClassMethod(); // This is a super class method.
```

派生クラス(サブクラス)で親クラスのメソッドを明示的に使用するには、parentキーワードを使うことができます。
**extendsキーワードで複数のクラスを継承すること（=多重継承）はできません。**


# 抽象メソッド（abstract）

```php
<?php

abstract class Human
{
    abstract function getAbility();

    public function run()
    {
        echo 'Run';
    }
}

class SuperHuman extends Human
{
    public function getAbility()
    {
        echo 'Fly';
    }
}

$super_human_instance = new SuperHuman();
echo $super_human_instance->run(); // Run
echo $super_human_instance->getAbility(); // Fly
```

抽象クラスはインターフェースと同じく、全ての定義済み抽象メソッドを派生クラスで実装する必要があります。
インターフェースでは実装を記述できませんが、抽象クラスでは実装を記述することができます。
抽象クラスとインターフェースの使い分けは[PHPのinterfaceとabstractを正しく理解して使い分けたいぞー](https://havelog.ayumusato.com/develop/php/e166-php-interface-abstract.html)がわかりやすく解説されています。


# インターフェース（Interface）

```php
<?php

interface Human
{
    public function eat();

    public function sleep();

    public function walk();
}

class Boy implements Human
{
    public function eat()
    {
        echo 'Eat';
    }

    public function sleep()
    {
        echo 'Sleep';
    }

    public function walk()
    {
        echo 'Walk';
    }

    public function fly()
    {
        echo 'Fly';
    }
}

$boy_instance = new Boy();
echo $super_human_instance->eat(); // Eat.
echo $super_human_instance->walk(); // Walk.
echo $super_human_instance->sleep(); // Sleep.
echo $super_human_instance->fly(); // Fly.
```

インターフェースはクラスの振る舞い（メソッドのみ）を定義するだけで、実装は行いません。
抽象クラスは、ベースとなるクラスの一部を拡張するように定義・実装を行いますが、インターフェースは定義されたメソッドを全て実装しなくてはなりません。

インターフェースを使った擬似的多重継承は可能です。

# トレイト（Trait）

```php
<?php

trait Authorize
{
    public function register()
    {
        echo 'Registration';
    }
}

class User
{
    use Authorize;
}

$user_instance = new User();
echo $user_instance->register();
```

traitはクラス階層を超えたコードの再利用を可能します。また、traitを使用することで多重継承を行うことができます。


# まとめ
今回はデザインパターンを学ぶ意図とオブジェクト指向の基礎について説明しました。
次回から各デザインパターンの紹介をします。


# 所感
仕組みがわかっても実際に使ってみる（設計する）のが難しいオブジェクト指向は、繰り返し手を動かして学び直す必要がありそうです。

