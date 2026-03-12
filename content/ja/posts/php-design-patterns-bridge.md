---
title: PHPで学ぶデザインパターン - Bridgeパターン
description: PHPで学ぶデザインパターン - Bridgeパターンについて、設計原則とトレードオフ、実践的な適用方法を詳しく解説します。
slug: php-design-patterns-bridge
date: 2019-02-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - GoF
  - PHP
  - デザインパターン
  - ブリッジパターン
translation_key: php-design-patterns-bridge
---


# 概要
[PHPで学ぶデザインパターン Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern)で間に合わなかった記事。

# Bridgeパターンとは
機能拡張のためのスーパークラスと実装拡張のためのサブクラスを用意し、機能の橋渡しをするようなパターン。


# 実装

```php
<?php
interface Connector
{
    public function __construct(Converter $converter);
    public function connect();
}

class IphoneConnector implements Connector
{
    private $converter;

    public function __construct(Converter $converter)
    {
        $this->converter = $converter;
    }

    public function connect()
    {
        echo 'Iphone connect by using ' . $this->converter->getTerminalName();
    }
}

class AndroidConnector implements Connector
{
    private $converter;

    public function __construct(Converter $converter)
    {
        $this->converter = $converter;
    }

    public function connect()
    {
        echo 'Android connect by using ' . $this->converter->getTerminalName();
    }
}

interface Converter
{
    public function getTerminalName();
}

class LightningConverter implements Converter
{
    public function getTerminalName()
    {
        return 'lightning';
    }
}

class TypeCConverter implements Converter
{
    public function getTerminalName()
    {
        return 'type-c';
    }
}

$lightingConveter = new LightningConverter();
$typeCConverter = new TypeCConverter();

// どっちのconverterを使ってもよい→実装を切り替えられる
$iphoneConnector = new IphoneConnector($lightingConveter);
$androidConnector = new AndroidConnector($typeCConverter);

$iphoneConnector->connect(); // connect by using lighting
$androidConnector->connect(); // connect by using type-c
```

# 所感
これは単純なインターフェースの使い方ではないかと思ってしまったのはまだ理解が浅いからな気がする。

# 参考
- [[保存版]人間が読んで理解できるデザインパターン解説#2: 構造系（翻訳）](https://techracho.bpsinc.jp/hachi8833/2017_10_11/46069#bridge

