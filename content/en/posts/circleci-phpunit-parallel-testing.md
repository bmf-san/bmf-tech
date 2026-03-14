---
title: Parallel Testing with PHPUnit on CircleCI
description: 'Execute PHPUnit tests in parallel on CircleCI using shell scripts to generate dynamic configuration and split containers.'
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
This post discusses an approach to perform parallel testing with PHPUnit on CircleCI.

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

Prepare a script like this to automatically generate the configuration file.

The reason for writing a shell script is that the tests are executed in a container, and due to using a convenience image in the CI job, there was no other suitable language.

# Distribute Tests Across Containers for Parallelization on CircleCI
With the script prepared earlier as `generate_phpunit.sh`, you can prepare for parallelization with the following script.

```sh
circleci tests glob "path/to/testdir/**/*.php" | circleci tests split | xargs sh +x generate_phpunit.sh
```

After that, by specifying the generated configuration file during test execution, tests can be executed across multiple containers, enabling parallelization.

```sh
path/to/phpunit -c path/to/ci_phpunit.xml
```

An example when testing with docker-compose.
```sh
docker-compose -f docker/docker-compose.test.yml run test ash -c "path/to/phpunit -c path/to/ci_phpunit.xml"
```

# Thoughts
This was something I actually tried in a work setting, but it turned out there were dependencies on the execution order between tests, making it difficult to easily achieve parallel test execution...

First, we need to address the dependencies...

# References
- [kojirooooocks.hatenablog.com - Execute CircleCI in Parallel to Reduce Execution Time](https://kojirooooocks.hatenablog.com/entry/2021/01/17/235100)
