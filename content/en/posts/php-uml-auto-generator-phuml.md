---
title: Tool for Automatically Generating UML with PHP - phUML
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
I wanted a tool that could automatically generate UML when I needed to grasp the overview of class design. While PhpStorm apparently has a nice built-in feature to generate diagrams, I've committed to VSCode, so I had to find a suitable tool.

# phUML
After some googling, I found various tools, but I was looking for something easy to use.

[github.com - MontealegreLuisphuml](https://github.com/MontealegreLuis/phuml)
[Documentation](https://montealegreluis.com/phuml/)

The original? [github.com - jakobwsthoff/phuml](https://github.com/jakobwesthoff/phuml) seems to have ended maintenance, but I found a forked version like the one above.

It doesn't have many stars, so it seems not many people use it... but it looked usable, so I tried it out.

It supports PHP version `^7.1`.

I used it in a 7.3 environment.

## Installation

```sh
$ wget https://montealegreluis.com/phuml/phuml.phar
$ wget https://montealegreluis.com/phuml/phuml.phar.pubkey
$ chmod +x phuml.phar
$ mv phuml.phar /usr/local/bin/phuml
$ mv phuml.phar.pubkey /usr/local/bin/phuml.pubkey
```

You can also install it with composer.

```php
composer require phuml/phuml
```

Once installed,

```php
vendor/bin/phuml phuml:diagram -r -a -i -o -p dot path/to/classes example.png
```

Using a bunch of suspicious-looking options like this will generate a class diagram.

You can check the options in the documentation.
[phUML - Generate a class diagram](https://montealegreluis.com/phuml/docs/class-diagram.html)

It seems you can specify options to exclude certain accessors from the output.

# Impressions
It seems useful for getting an overview when you want to understand a larger design. It would be nice if there was a VSCode plugin, but there isn't one at the moment.
