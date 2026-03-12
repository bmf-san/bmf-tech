---
title: 単一テーブル継承・クラステーブル継承・具象クラス継承について
description: 単一テーブル継承・クラステーブル継承・具象クラス継承についてについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: table-inheritance-types
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - PofEAA
translation_key: table-inheritance-types
---


# 概要
**リレーショナルなデータベースは継承をサポートをしていないので、オブジェクトの継承関係をデータベースにどのように表現するのか**考慮する必要があります。
それを表現する3つのパターン、単一テーブル継承・クラステーブル継承・具象クラス継承とはについて説明します。　

※各パターンの実装におけるメリット・デメリット等には触れません。

# 前提
今回想定する登場するクラスは4つです。

![people_class.png](/assets/images/posts/table-inheritance-types/7680ef3c-3c9f-6365-8622-364ed30936b3.png)

Party PeopleがRich Peopleを継承するという構造はちょっとわかりづらいかもしれませんが、イメージが伝われば良しとします。

#### People（人々）
全クラスに共通する属性を持っています。

- idは一意なキーです。
- nameは名前です。

#### OridinaryPeople（パンピー）
良識を持った善良なる一般ピーポーです。

- good_senseは良識です。bool（0か1）で入ります。

#### RichPeople（成金）
お金と土地を持っているリッチな人々です。
moneyはお金です。
landは土地です。

※単位とか細かいことは考慮していません。

#### PartyPeople（パリピ）
パーリーピーポー。

- free_timeは自由な時間です。
- middle_nameはミドルネームです。


# Single Table Inheritance　（単一テーブル継承）
単一テーブル継承は、オブジェクトの継承関係を１つのテーブルで表現します。
テーブルにはサブクラスを判断するためのカラム（type）を持たせます。

![single_table_inheritance_table.png](/assets/images/posts/table-inheritance-types/733b241d-ed09-6e1f-958c-b664f2d4133c.png)

RailsでSTIの実装がサポートされているようです。

# Class Table Inheritance　（クラステーブル継承）
クラステーブル継承は、オブジェクトの継承関係をクラスごとに１テーブルを用意することで表現します。
スーパークラスのテーブルにはスーパークラスの持つカラムを、サブクラスのテーブルにはサブクラスの持つカラムのみを持たせます。

![class_table_inheritance.png](/assets/images/posts/table-inheritance-types/33047bc2-d4a3-700c-0995-8738c9897a23.png)


# Concrete Class Inheritance　（具象クラス継承）
具象テーブル継承は、オブジェクトの継承関係を具象クラスだけ対応したテーブルを用意することで表現します。
各テーブルにはスーパークラスが持つカラムを共通属性として持たせます。

![concrete_table_inheritance.png](/assets/images/posts/table-inheritance-types/bec91e44-0b28-7bcc-6666-026dd5a10f2a.png)

# 所感
どのパターンを実装するかはテーブル設計のメリット・デメリットとアプリケーション側のロジックのコストの検討次第でしょうか。
何か語弊がある部分や間違いがある場合はご指摘ください。

# 参考リンク
- [みんなRailsのSTIを誤解してないか!?](http://qiita.com/yebihara/items/9ecb838893ad99be0561#class-table-inheritance--%E3%82%AF%E3%83%A9%E3%82%B9%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB%E7%B6%99%E6%89%BF)
- [Martin Fowler's Bliki(ja) - Patterns of Enterprise Application Architecture](http://bliki-ja.github.io/pofeaa/)

