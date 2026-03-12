---
title: PHPで学ぶデザインパターン - Strategy
description: PHPで学ぶデザインパターン - Strategyについて、設計原則とトレードオフ、実践的な適用方法を詳しく解説します。
slug: php-design-patterns-strategy
date: 2018-12-09T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - デザインパターン
  - PHP
  - GoF
  - ストラテジーパターン
translation_key: php-design-patterns-strategy
---


# 概要
この記事は[PHPで学ぶデザインパターン Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern)の記事です。

今回はStrategyパターンについてかきます。

# Strategyパターンとは？
Strategyパターンは、アルゴリズムの切り替えを容易にするようなパターンです。
異なる処理をそれぞれ別のクラスに定義するため、 処理を動的に選択できるだけでなく、条件分岐を減らすことも可能としてします。
OCP(open/closed principle)に忠実なパターンの一つでもあります。

# 実装例
単純な例でStrategyパターンの実装を見てみます。

```php
<?php
class Context
{
    private $notification;

    public function __construct(NotificationInterface $notification)
    {
        $this->notification = $notification;
    }

    public function execute()
    {
        return $this->notification->notify();
    }
}

interface NotificationInterface
{
    public function notify();
}

class SlackNotification implements NotificationInterface
{
    public function __construct($message)
    {
        $this->message = $message;
    }

    public function notify()
    {
        echo $this->message . '- sent by Slack';
    }
}

class EmailNotification implements NotificationInterface
{
    public function __construct($message)
    {
        $this->message = $message;
    }

    public function notify()
    {
        echo $this->message . '- sent by Email';
    }
}

$message = "Hello World!";

$slack = new SlackNotification($message);
$context = new Context($slack);
$context->execute(); // Hello World - sent by Slack

$email = new EmailNotification($message);
$context = new Context($email);
$context->execute(); // Hello World - sent by Email
```

Contextクラスの実装をinterfaceを依存させることで"戦略”をクライアント側から切り替えられるようにしています。
これはよく見る実装なのではないでしょうか？

振る舞いが分離されているので、OCPに従った形になっているはず。　
拡張に対しては、NotificationInterfaceの実装を追加するだけ、修正に対しては、NotificationInterfaceの実装を修正するだけ、という感じになっているかと思います。

# 所感
GoFは知らずしらずに利用していることがあるのでちゃんと覚えておいても損はないなーと思いました。

# 参考
- [Wikipedia - Strategyパターン](https://ja.wikipedia.org/wiki/Strategy_%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [Do You PHP はてな - [doyouphp][phpdp]PHPによるデザインパターン入門 - Strategy～戦略を切り替える](http://d.hatena.ne.jp/shimooka/20141219/1418965548)
- [isseium's blog - PHPでデザインパターン 〜Strategy パターン〜](http://isseium.hateblo.jp/entry/20101114/1289725226)
- [DesignPatternsPHP - Strategy](https://designpatternsphp.readthedocs.io/en/latest/Behavioral/Strategy/README.html?highlight=strategy)
- [株式会社クイックのWebサービス開発blog - PHPでStrategyパターンを考えてみよう](http://aimstogeek.hatenablog.com/entry/2016/12/13/190105)

