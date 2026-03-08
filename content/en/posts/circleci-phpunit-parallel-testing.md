---
title: Running Parallel Tests with PHPUnit on CircleCI
slug: circleci-phpunit-parallel-testing
date: 2023-10-21T00:00:00Z
author: bmf-san
categories:
  - Testing
tags:
  - CircleCI
  - phpunit
translation_key: circleci-phpunit-parallel-testing
---

# Overview
This post discusses an approach to running parallel tests with PHPUnit on CircleCI.

# Prepare a Script to Generate PHPUnit Configuration File
```sh
#!/bin/sh

basePath="/foo/bar"
testFiles=$@

xmlFileStringData=""
for file in $testFiles; do
    xmlFileStringData="${xmlFileStringData}<file>${basePath}/${file}</file>\n"
done

testFileString="$xmlFileStringData"

template="<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<phpunit colors=\"true\" stopOnFailure=\"false\" stopOnError=\"false\" failOnWarning=\"false\" stderr=\"true\" bootstrap=\"path/to/bootstrap\">
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
This script generates the configuration file automatically.

The shell script is used because the tests are run in a container, and due to the convenience of using a CI job with a convenience image, there was no other suitable language available.

# Distributing Tests Across Containers for Parallelization on CircleCI
Using the script prepared earlier as `generate_phpunit.sh`, you can prepare for parallelization with the following script.

```sh
circleci tests glob "path/to/testdir/**/*.php" | circleci tests split | xargs sh +x generate_phpunit.sh
```

After that, by specifying the generated configuration file during test execution, tests will run across multiple containers, achieving parallelization.

```sh
path/to/phpunit -c path/to/ci_phpunit.xml
```

An example of testing with docker-compose:
```sh
docker-compose -f docker/docker-compose.test.yml run test ash -c "path/to/phpunit -c path/to/ci_phpunit.xml"
```

# Thoughts
This was something I actually tried in my work, but it seems there are dependencies in the execution order between tests, making it difficult to easily achieve parallel test execution...

First, I need to address the dependencies...

# References
- [kojirooooocks.hatenablog.com - Shortening Execution Time with Parallel Execution on CircleCI](https://kojirooooocks.hatenablog.com/entry/2021/01/17/235100)