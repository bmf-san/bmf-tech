---
title: Laravel5.4でsqliteを使ってテストをかく準備
description: Laravel5.4でsqliteを使ってテストをかく準備
slug: laravel-5-4-sqlite-test-setup
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - テスト
tags:
  - Laravel
  - SQLite
translation_key: laravel-5-4-sqlite-test-setup
---


# 概要
Laravel5.4でsqliteの使ってテストをかく準備をします。

# 前提
- Laravelの基本
- マイグレーションファイルの用意

# phpunit.xmlを編集

以下3行を追加します。　sqliteのインメモリ機能を使います。

```php:phpunit.xml
    <php>
        <env name="APP_ENV" value="testing"/>
        <env name="CACHE_DRIVER" value="array"/>
        <env name="SESSION_DRIVER" value="array"/>
        <env name="QUEUE_DRIVER" value="sync"/>
        <env name="DB_CONNECTION" value="sqlite"/> // Add
        <env name="DB_DATABASE" value=":memory:"/> // Add
        <env name="ADMIN_DOMAIN" value="localhost"/> // Add
    </php>
```

# ModelFactoryを編集
ここは良しなに準備してください。略。

# テストをかく
```php:phpunit.xml
<testsuites>
        <testsuite name="Application Test Suite">
            <directory suffix="Test.php">./tests</directory>
        </testsuite>
</testsuites>
```

デフォルトでphpunit.xmlにtestsディレクトリ以下の〇〇Test.phpを実行するという設定になっているので、適当にテストファイルを用意します。

```php:/tests/Unit/HogeTest.php
<?php

use Illuminate\Foundation\Testing\WithoutMiddleware;
use Illuminate\Foundation\Testing\DatabaseMigrations;
use Illuminate\Foundation\Testing\DatabaseTransactions;

class PostTest extends TestCase
{
    use DatabaseMigrations; // Run migration

    public function testIndex()
    {
        $user = factory(App\User::class)->create();

        var_dump($user->first()); // Check the data

        // Here is your tests...
    }
}
```

`DatabaseMigrations`というトレイトを設定すると、テスト実行の度にマイグレーションが実行されます。

# 余談
`./vendor/bin/phpunit`というテストを実行する時のコマンドをcomposerのscriptsに記述しておくと楽できます。

# 所感
これでテストがかけるマン。

