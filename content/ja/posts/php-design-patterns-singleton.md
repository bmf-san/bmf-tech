---
title: PHPで学ぶデザインパターン - Singleton ~いくつ作るかを制限~
description: PHPで学ぶデザインパターン - Singleton ~いくつ作るかを制限~
slug: php-design-patterns-singleton
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - デザインパターン
  - PHP
  - シングルトンパターン
translation_key: php-design-patterns-singleton
---


# Singleton パターンとは？
インスタンス生成のコストを制御するために、インスタンスが１つしかないことを保証するパターンです。

# 構造
## SingletonClass
priavateのコンストラクタとインスタンスを1つだけ返すstaticメソッドと自分自身のインスタンスを保持するためのstatic変数を用意するだけです。

# メリット
## インスタンスへのアクセスを制御できる
Singletonパターンが保持する自分自身へのアクセスをprivateに制限しているためクライアント側のコードからのアクセスを制御することができます。

## インスタンス数を変更することができる
生成されるインスタンスの数を2つ以上に変更することも可能です。

# デメリット
* テスタビリティが下がる
* 実装後にインスタンスの制御数の変更が必要になると柔軟に対応しづらい

# 使いドコロ

# 実装例（※~~github~~にリポジトリあります。）

```SingletonConfig.php
<?php
class SingletonConfig {
    private $config;

    /**
     * a single variable
     */
    private static $instance;

    private function __construct()
    {
        $this->config = 'AUTO';
    }

    /**
     * Create a only instance
     */
    public static function getInstance()
    {
        if (!isset(self::$instance)) {
            self::$instance = new SingletonConfig();
        }

        return self::$instance;
    }

    public function getConfig()
    {
        return $this->config;
    }

    public final function __clone()
    {
        throw new RuntimeException('Clone is not allowed against' . get_class($this));
    }
}
```

```singleton_client.php
<?php
require_once 'SingletonConfig.php';

$instanceA = SingletonConfig::getInstance();

$instanceB = SingletonConfig::getInstance();

if ($instanceA->getConfig() === $instanceB->getConfig()) {
    echo 'True';
} // true
```

# まとめ
* インスタンス生成手段をprivateにして外部からのアクセス手段を制限する
* インスタンス生成回数の制御はstaticメソッドで行う

# 関連キーワード
* カプセル化

