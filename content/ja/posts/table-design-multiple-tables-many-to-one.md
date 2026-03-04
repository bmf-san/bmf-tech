---
title: "複数のテーブルに多対1で紐づく時のテーブル設計のアプローチについて"
slug: "table-design-multiple-tables-many-to-one"
date: 2018-08-06
author: bmf-san
categories:
  - "データベース"
tags:
  - "ポリモーフィック"
  - "SQLアンチパターン"
draft: false
---

# 概要
複数のテーブルに対し、多対1でテーブルが関係付くときのテーブル設計のパターンについてまとめる。

# データ設計
以下のようなケースのデータ設計を例とする。　

- issues
    - id
    - title

- pullrequests
    - id
    - title

- comments
    - id
    - content

`comments`が`issues`、`pullrequests`のどちらに対しても多対一で関係付くようなケース。

# ポリモーフィック関連

- issues
    - id
    - title

- pullrequests
    - id
    - title

- comments
    - id
    - content
    - target_table
    - target_id

`comments`に`target_table`と`target_id`というカラムを追加し、`issues`と`pullrequests`のどちらに結びつくか判断させようとするテーブル設計。

SQLアンチパターンではアンチパターンの一つとして取り上げられている。

`target_id`が`target_table`を見ないと`issues`と`pullrequests`のどちらに関連付くかわからないため、**外部キー制約が使えない**。
したがって、このパターンでは**テーブル間の整合性保持はアプリケーションのロジックに依存する**ことなる。

LaravelやRailsのORMではポリモーフィック関連がサポートされているので実装が楽なので、このようなパターンを検討する余地はゼロではないが、なるべく避けたいパターンではある。


# 交差（ピボット、中間）テーブル

- issues
    - id
    - title

- pullrequests
    - id
    - title

- issues_comments
   - issues_id
   - comments_id

- pullrequests_comments
   - pullrequests_id
   - comments_id

- comments
    - id
    - content

`issues`と`pullrequests`に交差テーブルを用意して、外部キー制約を使えるようにするパターン。

`issues`と`issues_comments`は1対多、`issues_comments`と`comments`は多対1となる。`pullrequests`に関しても同様。

アプリケーションの要件次第ではあるが、1コメントが`issues`と`pullrequests`のどちらかだけに関連付くようにという制約を保証できない。

外部キーが使えるため、ポリモーフィック関連よりは整合性を保つことができる。　

# 共通の親を持つテーブル
- issues
    - id
    - post_id

- pullrequests
    - id
    - post_id

- posts
    - id
    - title

- comments
    - id
    - content
    - post_id

`issues` 、`pullrequests`、`comments`の共通の親となるテーブルを用意するパターン。　

`posts`はクラステーブル継承の考え方に基づいて定義するのが良さそう。（要は基底クラスと考える）
（参考：[単一テーブル継承・クラステーブル継承・具象クラス継承について
PofEAA](https://bmf-tech.com/posts/%E5%8D%98%E4%B8%80%E3%83%86%E3%83%BC%E3%83%95%E3%82%99%E3%83%AB%E7%B6%99%E6%89%BF%E3%83%BB%E3%82%AF%E3%83%A9%E3%82%B9%E3%83%86%E3%83%BC%E3%83%95%E3%82%99%E3%83%AB%E7%B6%99%E6%89%BF%E3%83%BB%E5%85%B7%E8%B1%A1%E3%82%AF%E3%83%A9%E3%82%B9%E7%B6%99%E6%89%BF%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6)）

`issues`と`posts`が1対1、`posts`と`comments`が1対多で関連付く。`pull_requests`も同様。
`posts`と`comments`は1対多で関連付く。

1コメントは1postsに関連付くという制約を保証できるが、`issues`と`pullrequests` のどちらかだけに関連付くという制約を保証できない。

# テーブル分割

- issues
    - id
    - title

- pullrequests
    - id
    - title

- issue_comments
    - id
    - issues_id
    - content

- pullrequest_comments
    - id
    - pullrequests_id
    - content

これはそもそもの前提を疑う話ではあるが、`comments` を1つのテーブルにまとめておくのではなく、別々の`comments`テーブルをそれぞれ用意してテーブルを分割しておけば良いのではないかというパターンである。

# 所感
アプリケーション側のロジックに依存することは、ヒューマンエラーの可能性を高めるので、テーブル構造にロジックを依存させる設計方針が基本的には良いパターンではないかと思う。
アプリケーションの要件に加えて、クエリの気持ちを考えて最適なパターンを選択できるようにしたい。

- [複数のテーブルに対して多対一で紐づくテーブルの設計アプローチ](https://spice-factory.co.jp/development/has-and-belongs-to-many-table/)
- [SQLアンチパターンを読んで （ポリモーフィック関連について）](https://blog.motimotilab.com/?p=207)
