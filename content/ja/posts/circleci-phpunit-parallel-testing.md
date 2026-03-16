---
title: CircleCIでphpunitの並列テストを行う
description: CircleCIでphpunitの並列テストを行う
slug: circleci-phpunit-parallel-testing
date: 2023-10-21T00:00:00Z
author: bmf-san
categories:
  - テスト
tags:
  - CircleCI
  - phpunit
translation_key: circleci-phpunit-parallel-testing
---


# 概要
CircleCIでphpunitの並列テストを行うアプローチについてかく。

# PHPUnitの設定ファイルを生成するスクリプトを用意する
```sh
#!/bin/sh

basePath="/foo/bar"
testFiles=$@

xmlFileStringData=""
for file in $testFiles; do
    xmlFileStringData="${xmlFileStringData}<file>${basePath}/${file}</file>\n"
done

testFileString="$xmlFileStringData"

template="<?xml version=\"1.0\" encoding=\"UTF-8\"?>
<phpunit colors=\"true\" stopOnFailure=\"false\" stopOnError=\"false\" failOnWarning=\"false\" stderr=\"true\" bootstrap=\"path/to/bootstrap\">
    <php>
     <ini name=\"memory_limit\" value=\"1G\"/>
     <ini name=\"realpath_cache_size\" value=\"1M\"/>
    </php>
    <testsuites>
        <testsuite name=\"Test Suite\">
            ${testFileString}
        </testsuite>
    </testsuites>
</phpunit>"

echo "$template" > "path/to/ci_phpunit.xml"
```

こんな感じのスクリプトを用意して、設定ファイルを自動生成する。

shellscriptを書いているのはテストの実行をコンテナで行っている都合上、CIのjobでコンビニエンスイメージを使っているため、他に都合良い言語がなかった。

# CircleCIでコンテナ毎にテストを振り分け並列化させる
先程用意したスクリプトをgenerate_phpunit.shとして、次のようなスクリプトで並列化の準備ができる。

```sh
circleci tests glob "path/to/testdir/**/*.php" | circleci tests split | xargs sh +x generate_phpunit.sh
```

後はテスト実行時に生成した設定ファイルを指定すれば、複数コンテナでテストが実行され、並列化ができる。

```sh
path/to/phpunit -c path/to/ci_phpunit.xml
```

docker-composeでテストする場合の一例。
```sh
docker-compose -f docker/docker-compose.test.yml run test ash -c "path/to/phpunit -c path/to/ci_phpunit.xml"
```

# 所感
これは実際に業務でトライしたことだったのだが、テスト間の実行順に依存関係があるらしく、並列化でのテスト実行を簡単に実現することができなかった...

まずは依存関係を何とかする必要がある...

# 参考
- [kojirooooocks.hatenablog.com - circleCIを並列実行して、実行時間を短縮する](https://kojirooooocks.hatenablog.com/entry/2021/01/17/235100)
