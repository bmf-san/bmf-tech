---
title: Using the Laravel Debug Tool laravel-debugbar
slug: laravel-debugging-tool-usage
date: 2016-06-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
translation_key: laravel-debugging-tool-usage
---

Introducing the debugging tool laravel-debugbar for Laravel.

[github-laravel-debugbar](https://github.com/barryvdh/laravel-debugbar)

It works with both Laravel 5.1 and 5.2.

![Screenshot 2016-06-27 0.12.17.png](https://qiita-image-store.s3.amazonaws.com/0/124495/777c108d-b00d-d91c-e189-add3765e502b.png)

# Installing laravel-debugbar with Composer

`composer require barryvdh/laravel-debugbar --dev`

Then run

`composer install`

# Enabling Facade Usage

While you can use the debugging tool just by installing it, it's more convenient to enable it for facade usage if you want to debug in more detail.

Specify the following in the provider and alias sections of app.php:

* Barryvdh\Debugbar\ServiceProvider
* 'Debugbar' => Barryvdh\Debugbar\Facade::class

# Trying It Out
`\Debugbar::error();`
`\Debugbar::disable();`
`Debugbar::startMeasure();`
`Debugbar::stopMeasure();`

And many more.

# Thoughts
It's very convenient (゜レ゜)