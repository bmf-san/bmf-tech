---
title: "CircleCI2.0でPHPUnitのコードカバレッジを出力する"
slug: "circleci-phpunit-code-coverage"
date: 2018-08-13
author: bmf-san
categories:
  - "テスト"
tags:
  - "Docker"
  - "CircleCI"
  - "CircleCI2.0"
  - "phpunit"
draft: false
---

# 概要
CircleCi2.0でPHPUnitのコードカバレッジを出力する

# 環境
- CircleCi2.0
- docker
- docker-compose
- PHPUnit6系
- PHP7.2系

# 手順
## phpunit.xmlの設定を調整する
カバレッジの対象としたいソースを指定する。

```
    <filter>
        <whitelist processUncoveredFilesFromWhitelist="true">
            <directory suffix=".php">./app</directory>
            <exclude>
                <directory>./app/Providers</directory>
                <directory>./app/Exceptions</directory>
                <directory>./app/Http/Middleware</directory>
                <directory>./app/Providers</directory>
                <file>./app/Console/Kernel.php</file>
                <file>./app/Http/Kernel.php</file>
                <file>./app/Http/Controllers/Controller.php</file>
            </exclude>
        </whitelist>
    </filter>
```


こんな感じで記述する。
phpunit.xml
```
<?xml version="1.0" encoding="UTF-8"?>
<phpunit
    backupGlobals="false"
    backupStaticAttributes="false"
    bootstrap="bootstrap/autoload.php"
    colors="true"
    convertErrorsToExceptions="true"
    convertNoticesToExceptions="true"
    convertWarningsToExceptions="true"
    processIsolation="false"
    stopOnFailure="false">
    <testsuites>
        <testsuite name="Application Test Suite">
            <directory suffix="Test.php">./tests</directory>
        </testsuite>
    </testsuites>
    <filter>
        <whitelist processUncoveredFilesFromWhitelist="true">
            <directory suffix=".php">./app</directory>
            <exclude>
                <directory>./app/Providers</directory>
                <directory>./app/Exceptions</directory>
                <directory>./app/Http/Middleware</directory>
                <directory>./app/Providers</directory>
                <file>./app/Console/Kernel.php</file>
                <file>./app/Http/Kernel.php</file>
                <file>./app/Http/Controllers/Controller.php</file>
            </exclude>
        </whitelist>
    </filter>
    <php>
        <env name="APP_ENV" value="testing"/>
        <env name="DB_CONNECTION" value="mysql_test"/>
    </php>
</phpunit>
```

## カバレッジレポートを出力する
xdebugを使うよりも速いらしいのでphpdbgを使ってhtml形式のカバレッジレポートを作成する。
CI上でメモリ不足で落ちたため、メモリリミットを指定するようにしている。

`phpdbg -qrr vendor/bin/phpunit -d memory_limit=512M --coverage-html /tmp/artifacts"`


## ./circleci/config.ymlを調整する
CircleCIのテスト実行後に、artifactsからカバレッジレポートを見れるようにしたいので、`store_artifacts`を使ってカバレッジレポートを`artifacts`にアップロードするタスクを記述する。

今回はdocker上でテストを行っているため、ホスト側とdocker側でカバレッジレポートのソースをマウントさせる必要があった。
やっつけかもしれないが、`docker cp`コマンドでファイルコピーする方法で実行してみた。

.circleci/config.yml
```
version: 2
jobs:
  build:
    machine: true
    steps:
        - checkout
        - run:
            name: Create a artifacts directory
            command: mkdir -p /tmp/artifacts
        - run:
            name: core-app - Run tests and create a code coverage report
            command: docker exec -it rubel_php /bin/sh -c "cd core-app/ && phpdbg -qrr vendor/bin/phpunit -d memory_limit=512M --coverage-html /tmp/artifacts"
        - run:
            name: Copy the coverage report to host directory
            command: docker cp rubel_php:/tmp/artifacts /tmp
        - store_artifacts:
            path: /tmp/artifacts
```

# 所感
カバレッジを出力することができたので各種Webサービス（codacyとかcoverallとか）と連携させてみたい。
ファイルのマウントが必要なことに気づかず時間をかかってしまった....

# 参考
- [circleci - Collecting Test Metadata](https://phpunit.de/manual/6.5/en/appendixes.configuration.html#appendixes.configuration.logging)
- [circleci - Storing and Accessing Build Artifacts](https://circleci.com/docs/2.0/artifacts/)
- [phpunit - Appendix C. The XML Configuration File](https://phpunit.de/manual/6.5/en/appendixes.configuration.html#appendixes.configuration.logging)
- [phpunit - 第3章 コマンドラインのテストランナー](https://phpunit.de/manual/6.5/ja/textui.html)
- [docker docs - docker cp](https://docs.docker.com/engine/reference/commandline/cp/)
