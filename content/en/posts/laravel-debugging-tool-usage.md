---
title: Using the Debug Tool laravel-debugbar in Laravel
slug: laravel-debugging-tool-usage
date: 2016-06-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
description: Introducing the laravel-debugbar debug tool for Laravel.
translation_key: laravel-debugging-tool-usage
---



Introducing the laravel-debugbar debug tool for Laravel.

[github-laravel-debugbar](https://github.com/barryvdh/laravel-debugbar)

It works with both Laravel 5.1 and 5.2.

![Screenshot 2016-06-27 0.12.17.png](/assets/images/posts/laravel-debugging-tool-usage/777c108d-b00d-d91c-e189-add3765e502b.png)

# Installing laravel-debugbar with composer

`composer require barryvdh/laravel-debugbar --dev`

then

`composer install`

# Enabling with facade

While you can use it as a debug tool just by installing, it's convenient to enable it with a facade for more detailed debugging.

Specify the following in the provider and alias sections of app.php respectively.

* Barryvdh\Debugbar\ServiceProvider
* 'Debugbar' => Barryvdh\Debugbar\Facade::class

# Trying it out
`\Debugbar::error();`
`\Debugbar::disable();`
`Debugbar::startMeasure();`
`Debugbar::stopMeasure();`

And many more.

# Thoughts
It's very convenient (゜レ゜)
