---
title: Tool for Auto-generating UML with PHP - phUML
slug: php-uml-auto-generator-phuml
date: 2020-06-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
translation_key: php-uml-auto-generator-phuml
---

# Overview
I wanted a tool that could automatically generate UML when I wanted to grasp the appearance of class design.
It seems that PhpStorm has a nice built-in feature for generating diagrams, but since I have converted to VSCode, I have no choice but to look for a good tool.

# phUML
There are various tools available when you Google, but I looked for one that seemed easy to use.

[github.com - MontealegreLuisphuml](https://github.com/MontealegreLuis/phuml)
[Documentation](https://montealegreluis.com/phuml/)

The original? [github.com - jakobwsthoff/phuml](https://github.com/jakobwesthoff/phuml) seems to have ended maintenance, but I found the above fork version.

It seems to have few stars, so I got the impression that not many people use it... but it looked usable, so I decided to give it a try.

The PHP version compatibility is `^7.1`.

I tried it in a 7.3 environment.

## Installation

```sh
$ wget https://montealegreluis.com/phuml/phuml.phar
$ wget https://montealegreluis.com/phuml/phuml.phar.pubkey
$ chmod +x phuml.phar
$ mv phuml.phar /usr/local/bin/phuml
$ mv phuml.phar.pubkey /usr/local/bin/phuml.pubkey
```

You can also install it via Composer.

```php
composer require phuml/phuml
```

Once installed,

```php
vendor/bin/phuml phuml:diagram -r -a -i -o -p dot path/to/classes example.png
```

By adding a lot of suspicious options like this, it generates a class diagram.

You can check the options in the documentation.
[phUML - Generate a class diagram](https://montealegreluis.com/phuml/docs/class-diagram.html)

It seems you can specify accessors you don't want to output as options.

# Impressions
It seems useful for getting an overview when I want to grasp a larger design.
I would be happy if there was a VSCode plugin, but so far there isn't.