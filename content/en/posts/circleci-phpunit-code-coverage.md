---
title: Output PHPUnit Code Coverage with CircleCI 2.0
slug: circleci-phpunit-code-coverage
date: 2018-08-13T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - Docker
  - CircleCI
  - CircleCI 2.0
  - PHPUnit
translation_key: circleci-phpunit-code-coverage
---

# Overview
Output PHPUnit code coverage with CircleCI 2.0

# Environment
- CircleCI 2.0
- Docker
- Docker Compose
- PHPUnit 6.x
- PHP 7.2

# Steps
## Adjust phpunit.xml Configuration
Specify the sources you want to include in the coverage.

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

This is how it should be written.
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

## Output Coverage Report
It seems that using phpdbg is faster than using xdebug, so we will create an HTML format coverage report using phpdbg. Since it crashed due to insufficient memory on CI, we specify the memory limit.

`phpdbg -qrr vendor/bin/phpunit -d memory_limit=512M --coverage-html /tmp/artifacts`

## Adjust ./circleci/config.yml
To view the coverage report from artifacts after running tests on CircleCI, we need to write a task that uploads the coverage report to `artifacts` using `store_artifacts`.

Since we are running tests on Docker, we needed to mount the coverage report sources between the host and the Docker container. Although it might be a bit of a hack, I tried executing it using the `docker cp` command to copy the files.

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

# Thoughts
I was able to output the coverage, so I would like to integrate it with various web services (like Codacy and Coveralls). I spent a lot of time realizing that file mounting was necessary....

# References
- [circleci - Collecting Test Metadata](https://phpunit.de/manual/6.5/en/appendixes.configuration.html#appendixes.configuration.logging)
- [circleci - Storing and Accessing Build Artifacts](https://circleci.com/docs/2.0/artifacts/)
- [phpunit - Appendix C. The XML Configuration File](https://phpunit.de/manual/6.5/en/appendixes.configuration.html#appendixes.configuration.logging)
- [phpunit - Chapter 3 Command Line Test Runner](https://phpunit.de/manual/6.5/ja/textui.html)
- [docker docs - docker cp](https://docs.docker.com/engine/reference/commandline/cp/)