---
title: "PHPで学ぶデザインパターン - Factory・Factory Method・Abstract Factory"
slug: "php-design-patterns-factory"
date: 2018-12-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "デザインパターン"
  - "GoF"
  - "PHP"
  - "アブストラクトファクトリーパターン"
  - "ファクトリーパターン"
  - "ファクトリーメソッドパターン"
draft: false
---

この記事は[PHPで学ぶデザインパターン Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern)の記事です。

# 概要
Factory・Factory Method・Abstract Factoryについてかきます。

# Factoryパターンについて知る
まずはFactoryパターンについてざっくり説明します。

```php
interface Robot
{
  public function say();
}

class BlueRobot implements Robot
{
    private $color;

    public function __construct($color)
    {
        $this->color = $color;
    }

    public function say()
    {
        echo $this->color;
    }
}

class YellowRobot implements Robot
{
    private $color;

    public function __construct($color)
    {
        $this->color = $color;
    }

    public function say()
    {
        echo $this->color;
    }
}

class RobotFactory
{
    public function create($color)
    {
        if ($color === 'blue') {
            return new BlueRobot($color);
        }

        if ($color === 'yellow') {
            return new YellowRobot($color);
        }

        throw new Exception("Can't create an instance");
    }
}

$robotFactory = new RobotFactory();

try {
    $blueRobot = $robotFactory->create('blue');
    $blueRobot->say(); // blue

    $yellowRobot = $robotFactory->create('yellow');
    $yellowRobot->say(); // yellow

    $greenRobot = $robotFactory->create('green');
    $greenRobot->say(); // Cat't create an instance
} catch (\Exception $e) {
    echo $e->getMessage();
}
```

その名の通り、工場（Factory）はオブジェクトの”生成”を担います。
オブジェクトの生成と利用を分離することで、利用側はオブジェクトの生成順や種類を知らなくてもオブジェクトを生成することができます。
オブジェクトの生成場所を1箇所にまとめることができるので、生成順や種類等の変更が容易になります。

# Factory Methodパターンについて知る
上述のFactoryパターンでは、生成したいオブジェクトのクラスが増えると条件分岐が増えてしまい、辛い未来が見えてきそうです。
そこで、オブジェクトのクラスを指定することなく、オブジェクトを生成するようにしようというのがFactory Methodパターンです。

Factoryを抽象化して、オブジェクトの生成処理をサブクラスに任せることで、条件分岐をなくすことができます。

生成するオブジェクトを変更する場合はFactoryの切り替えを行う形になります。

```php
interface Robot
{
  public function say();
}

class BlueRobot implements Robot
{
    public function say()
    {
        echo 'Blue';
    }
}

class YellowRobot implements Robot
{
    public function say()
    {
        echo 'Yellow';
    }
}

abstract class RobotFactory
{
    abstract protected function create();

    public function do()
    {
        $robot = $this->create();
        $robot->say();
    }
}

class BlueRobotFactory extends RobotFactory
{
    protected function create()
    {
        return new BlueRobot();
    }
}

class YellowRobotFactory extends RobotFactory
{
    protected function create()
    {
        return new YellowRobot();
    }
}

$blueRobot = new BlueRobotFactory();
$blueRobot->do(); // blue

$yellowRobot = new YellowRobotFactory();
$yellowRobot->do(); // yellow
```

FactoyパターンのときのコードをFactory Methodパターンを適用した形に変更してみました。

引数によって生成されるオブジェクトを切り替えていた部分が大きく変わりました。

FactoryパターンとFactory Methodパターンはオブジェクトの生成パターンが異なるものですが、両者を混同している記事がいくつか見受けられて厄介だなーと思いました。

# Abstract Factoryパターンについて知る
Abstract Factoryパターンは、複数のFactoryを共通のテーマによってグループ化するようなパターンです。

```php
interface Robot
{
  public function say();
}

class BlueRobot implements Robot
{
    public function say()
    {
        echo 'Blue';
    }
}

class YellowRobot implements Robot
{
    public function say()
    {
        echo 'Yellow';
    }
}

interface RobotCreator
{
    public function work();
}

class BlueRobotCreator implements RobotCreator
{
    public function work()
    {
        echo '青いロボットつくるよ';
    }
}

class YellowRobotCreator implements RobotCreator
{
    public function work()
    {
        echo '黄色いロボットつくるよ';
    }
}

interface RobotFactory
{
    public function createRobot();
    public function createRobotCreator();
}

class BlueRobotFactory implements RobotFactory
{
    public function createRobot()
    {
        return new BlueRobot();
    }

    public function createRobotCreator()
    {
        return new BlueRobotCreator();
    }
}

class YellowRobotFactory implements RobotFactory
{
    public function createRobot()
    {
        return new YellowRobot();
    }

    public function createRobotCreator()
    {
        return new YellowRobotCreator();
    }
}


$blueRobotFactory = new BlueRobotFactory();
$blueRobot = $blueRobotFactory->createRobot();
$blueRobotCreator = $blueRobotFactory->createRobotCreator();

$blueRobot->say(); // blue
$blueRobotCreator->work(); // 青いロボットつくるよ

$yellowRobotFactory = new YellowRobotFactory();
$yellowRobot = $yellowRobotFactory->createRobot();
$yellowRobotCreator = $yellowRobotFactory->createRobotCreator();

$yellowRobot->say(); // yellow
$yellowRobotCreator->work(); // 青いロボットつくるよ
```

Abstract FactoryはFactory Methodの発展した感じという印象です。

Factoryのインターフェースを定義しているという点で、より上位概念からの解釈ができると思うのですが、今回はニュアンスを掴むまでに留めておきます。

# 所感
それぞれのパターンのニュアンスの違いがわかったような気がします。
もう少しオブジェクト指向の真髄を知っていたらより深い解釈ができそうです。

# 参考
- [DesignPatternsPHP](https://designpatternsphp.readthedocs.io/en/latest/README.html)
- [PHP The Right Way = Design Patterns](https://laravel-taiwan.github.io/php-the-right-way/pages/Design-Patterns.html)
- [Tech Racho - [保存版]人間が読んで理解できるデザインパターン解説#1: 作成系（翻訳）](https://techracho.bpsinc.jp/hachi8833/2017_10_02/46064)
- [オブジェクト思考 - Factory（ファクトリ）パターン](https://think-on-object.blogspot.com/2011/11/factoryfactory-methodabstract-factory.html)

