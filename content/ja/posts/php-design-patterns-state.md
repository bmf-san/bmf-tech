---
title: "PHPで学ぶデザインパターン - Stateパターン"
slug: "php-design-patterns-state"
date: 2019-04-20
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "GoF"
  - "PHP"
  - "デザインパターン"
  - "ステートパターン"
draft: false
---

# 概要
[PHPで学ぶデザインパターン Advent Calendar 2018](https://qiita.com/advent-calendar/2018/php-design-pattern)で間に合わなかった記事。

# Stateパターンとは
状態をクラスで用意することで振る舞いを切り替えることができるようなパターン。

# 実装
スイッチのオンオフの状態を切り替えるような例を想定した。
シングルトンを使ったほうが良い気がする。

```php
<?php

class OnState
{
    public function getState()
    {
        return 'ON';
    }

    public function getNextState()
    {
        return new OffState();
    }
}

class OffState
{
    public function getState()
    {
        return 'OFF';
    }

    public function getNextState()
    {
        return new OnState();
    }
}

class Light
{
    public function __construct()
    {
        // デフォルトの状態クラスをセット
        $this->state = new OffState();
    }

    public function getState()
    {
        return $this->state->getState();
    }

    public function toggle()
    {
        $this->state = $this->state->getNextState();
    }
}


$light = new Light();

echo $light->getState(); // OFF
echo $light->toggle();
echo $light->getState(); // ON
echo $light->toggle();
echo $light->getState(); // OFF
```

状態クラスは状態固有の処理を持つイメージ。

複数の状態があり、それぞれ固有の処理が複雑してきた時に検討できそうなパターンかと思う。

# 所感
デザインパターンの中でもなぜだか一番好きなパターン。
使ってみたくなるような面白さがある気がする。

# 参考
- [Stateパターン - PHPデザインパターン](https://www.ritolab.com/entry/140)


