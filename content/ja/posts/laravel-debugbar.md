---
title: "Laravelのデバッグツールlaravel-debugbarをつかってみる"
slug: "laravel-debugbar"
date: 2016-06-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Laravel"
draft: false
---

Laravelのlaravel-debugbarというデバックツールについて紹介します。


[github-laravel-debugbar](https://github.com/barryvdh/laravel-debugbar)


laravel5.1でも5.2でも使えました。


![スクリーンショット 2016-06-27 0.12.17.png](https://qiita-image-store.s3.amazonaws.com/0/124495/777c108d-b00d-d91c-e189-add3765e502b.png)


# composerでlaravel-debugarを導入

`composer require barryvdh/laravel-debugbar --dev`

からの

`composer install`


# facadeで使えるようにする

インストールするだけでもデバッグツールとして問題なく使えますが、より詳細にデバッグしたい場合はfacadeで使えるようにしておくと便利です。

app.phpのproviderとalias部分に以下をそれぞれ指定。

* Barryvdh\Debugbar\ServiceProvider
* 'Debugbar' => Barryvdh\Debugbar\Facade::class


# 使ってみる
`\Debugbar::error();　`
`\Debugbar::disable();　`
`Debugbar::startMeasure();`
`Debugbar::stopMeasure();`

その他色々。

#感想
大変便利です(゜レ゜)

