---
title: "PHPで学ぶデザインパターン - Adapter ~APIを変更する~"
slug: "php-adapter-api-change"
date: 2017-10-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "PHP"
  - "アダプターパターン"
  - "デザインパターン"
draft: false
---

# Adapter パターンとは？
API（互換性のないインターフェース）同士を適合させるためのパターンです。**既存のコードの変更をせずに、再利用することで新しい機能を提供する**というものです。再利用するコードには変更を加えないというのが特徴です。
主にコードを再利用するためという後天的理由から成り立っているパターンです。（設計段階でラッパーを用意するパターンはBridgeパターンです。）

# 構造
##TargetClass
API（インターフェース）の定義をします。

## AdapteeClass
TargetClassに適合させる既存のAPIを提供します。

## AdapterClass
AdapteeClassのAPIをTargetClassから利用できるように変換します。

# メリット
## 既存のコードを修正なしで再利用できる
既存のクラスにラッピングする形で実装するため、既存のコードを修正する必要がありません。

## クライアント側が既存APIの実装を意識する手間を省ける
要は既存APIの変更がクライアント側の変更に影響しないということです。

## 公開するAPIを自由に制限できる
APIを適合させる際にAPIのアクセスを制限することができます。

## デメリット
* 適合させるレイヤー増えるとパフォーマンスに影響がでる可能性がある

# 使いドコロ
既存の実績あるクラスを再利用したい時など。

# 実装例（※[github](https://github.com/bmf-san/design-patterns-php)にリポジトリあります。）

## 継承を使ったパターン
```ShowData.php
<?php

class ShowData {
    private $data;

    public function __construct($data)
    {
        $this->data = $data;
    }

    public function showOriginalData()
    {
        echo $this->data;
    }

    public function showProcessedData()
    {
        echo $this->data . 'How are you?';
    }
}
```

```ShowSourceData.php
<?php

interface ShowSourceData {
    public function show();
}
```

```ShowSourceDataImpl.php
<?php
require_once 'ShowSourceData.php';
require_once 'ShowData.php';

class ShowSourceDataImpl extends ShowData implements ShowSourceData {
    public function __construct($data)
    {
        parent::__construct($data);
    }

    public function show()
    {
        parent::showProcessedData();
    }
}
```

```adapter_client.php
<?php
require_once 'ShowSourceDataImpl.php';

$show_data = new ShowSourceDataImpl('Hello! Mr. Data.');

$show_data->show();
```

## 委譲を使ったパターン

ラッパー部分が異なるだけで、クライアント側コードは同じです。
委譲とは、**具体的な処理を別のクラスに任せる**という意味です。
DIのような・・といっては語弊があるでしょうか・・・(゜-゜)


```ShowSourceDataImpl.php
<?php
require_once '../ShowSourceData.php';
require_once '../ShowData.php';

class ShowSourceDataImpl implements ShowSourceData {
    private $show_data;

    public function __construct($data)
    {
        $this->show_data = new ShowData($data);
    }

    public function show()
    {
        $this->show_data->showProcessedData();
    }
}
```

# まとめ
* 既存のコードを間接的に再利用するラッパークラスを用意する。
* ラッパークラスは継承ベース、委譲ベースがある。

# 関連キーワード
* Bridgeパターン

