---
title: "DIとサービスロケーター"
slug: "di"
date: 2018-06-05
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "DI"
  - "サービスロケーター"
  - "デザインパターン"
draft: false
---

# 概要
DIとService Locatorの違いについてまとめる

# DIとは
- デザインパターンの一種
- 依存性注入
  - オブジェクト間の依存性を分離
  - オブジェクトの実行時に必要なオブジェクトが注入されるようにする
- テストしやすくなる

# DIパターンの実装
DIパターン（コンストラクタインジェクション）を実装してみる。
なお、DIパターンには、コンストラクタインジェクション、セッターインジェクション、メソッドインジェクションなどコンストラクタ以外からDIする方法もある。
比較のためにDIではないパターンとDIのパターンの両方を実装する。

## DIではないパターン

```php
<?php
class SlackNotification
{
    public function notify(string $message)
    {
        echo $message;

        return $this;
    }
}

class Application
{
    protected $message;

    public function __construct()
    {
        $this->notification = new SlackNotification();
    }

    public function alert(string $message)
    {
        $this->notification->notify($message);
    }
}

// client
$application = new Application();
$application->alert('slack alert');
```

## DIパターン

```php
<?php

interface NotificationInterface
{
    public function notify(string $message);
}

class SlackNotification implements NotificationInterface
{
    public function notify(string $message)
    {
        echo $message;

        return $this;
    }
}

class Application
{
    protected $message;

    public function __construct(NotificationInterface $notification) // DI
    {
        $this->notification = $notification;
    }

    public function alert(string $message)
    {
        $this->notification->notify($message);
    }
}

// client
$slackNotification = new SlackNotification();
$application = new Application($slackNotification); // DI
$application->alert('slack alert!');
```

`Application`は`SlackNotification`の責務を持たず、`NotificationInterface`のみに依存している。
`Application`はコンスタラクタで`SlackNotification`を受け入れる（依存する）形になっている。

依存注入部分はモック化できるのでテストしやすくなっている。

```php
// For test
$mockNotification = new MockNotification(); // NotificationInterfaceを実装したMockNotification
$application = new Application($mockNotification);
$application->alert('mock alert!');
```

# DIのメリットとデメリット
## メリット
- オブジェクトの結合度が下がり、依存関係が明確になる
- 変更に強い
- テストがしやすい

## デメリット
- ファイルが増える
- プログラム実行速度の低下の可能性

# DIコンテナ
- DI機能を提供するフレームワーク
- 依存関係を一括で管理できるコンテナの提供

# サービスロケーターとは
- サービス（オブジェクト）の取得を抽象化するデザインパターン
- アンチパターン
  - 依存するクラスが一つ増える（サービスロケーターのコンテナクラス）
  - コンテナクラスがテストを面倒にする
  - 依存するクラスがわかりにくくなる
   - 詳しくは「PHP The Right Way」

こんな感じのやつ

```php
<?php

class Application
{
    public function __construct($container)
    {
        $this->slackNotification = $container['slack.notification'];
    }
}

$application = new Applicaion($container);
```

# 参考
- [DI（Dependency Injection）に関するメモ](http://blog.shin1x1.com/entry/di-memo)
- [PHPで作って覚えるDI Container - その1 - DI is 何](https://qiita.com/zeriyoshi/items/e26daccd59669b623a41)
- [PHPで作って覚えるDI コンテナ - その2 - DI コンテナとServiceLocator](https://qiita.com/zeriyoshi/items/ef71bec08441877ca219)
- [Design PatternsPHP - Service Locator](http://designpatternsphp.readthedocs.io/en/latest/More/ServiceLocator/README.html)
- [PHP The Right Way](http://ja.phptherightway.com/#containers)
- [DIコンテナを使おうとしてサービスロケータにならないようにするにはどうしたらよいのでしょうか、具体例がわかりません](https://teratail.com/questions/49143)

