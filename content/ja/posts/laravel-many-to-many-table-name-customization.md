---
title: Laravelの多対多のリレーションでテーブル名を規則から外す時
slug: laravel-many-to-many-table-name-customization
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Laravel
translation_key: laravel-many-to-many-table-name-customization
---


多対多のリレーションを設計するときに、ドキュメント通りやっていたと思ったらちょっとした勘違いをした話です。

# ここに3つのテーブルがあるじゃろ？

今回のテーブル
* events
* event_tags
* event_tag_event←pivotテーブル

通常のテーブルは
* events
* tags
* tag_event

って感じでデフォルトの規則通りリレーションを貼ればいいのですが、ちょっと癖のある名前にすると少し気をつけるところがあるようです。


# ドキュメントを見てみる
[Laravel 5.1 Eloquent：リレーション](https://readouble.com/laravel/5.1/ja/eloquent-relationships.html#many-to-many)

ほうほう第2引数をもたせてあげればいいんだなー

```Event.php
public function eventTags()
{
  // 第2引数はPivotテーブル！
  return $this->belongstoMany('App\Modles\EventTag', 'event_tag_event)->withTimestamps();
}

```

```EventTag.php
public function events()
{
  return $this->belongsToMany('App\Models\Events');
}
```


tinkerを立ち上げて確認すると・・

```
SQLSTATE[42000]: Syntax error or access violation: 1066 Not unique table/alias on relationship
```
　
怒られます。ヽ(´ー｀)ノ



# 第2引数はPivotテーブル名！

もしかしてPivotテーブル名を指定するのでは・・？

[SQLSTATE[42000]: Syntax error or access violation: 1066 Not unique table/alias on relationship](http://stackoverflow.com/questions/31059595/sqlstate42000-syntax-error-or-access-violation-1066-not-unique-table-alias-o)



```Event.php
public function eventTags()
{
  // 第二引数はPivotテーブル！
  return $this->belongstoMany('App\Modles\EventTag', 'event_tag_event)->withTimestamps();
}

```

```EventTag.php
public function events()
{
  return $this->belongsToMany('App\Models\Events');
}
```

怒られませんでした。ヽ(´ー｀)ノ

# 所感
最近はバッグエンドよりもフロントエンドが気になって夜も寝れませんヽ(´ー｀)ノ

