---
title: "N+1問題とは？Go/Railsでの検出と解決方法"
slug: what-is-n-plus-one-problem
date: 2018-05-12T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - N+1
translation_key: what-is-n-plus-one-problem
---


# 概要
N+1問題の説明と対応についてまとめる。

# N+1問題とは
- 全てのレコードの取得に1個+レコード文だけのN個のSQLを発行してしまう問題
- N+1というより1+Nと解釈したほうがわかりやすい

# 例
- 一覧表示用のデータを取得するケース
  - 一覧用の全データを取得するSELECTを1回発行（Nレコード返ってくる）
  - Nレコードの関連データ取得のためにSELECTをN回発行

# 対応
- join
  - `SELECT "users".* FROM "users" INNER JOIN "posts" ON "posts"."user_id" = "users"."id" WHERE "posts"."id" = 1"`
  - Eager Loading
    - `SELECT "users".* FROM "users"`
    - `SELECT "posts".* FROM "posts" WHERE "posts"."id" IN (1, 2, 3, 4, 5)`
 
# 参考
- [N+1問題は1+N問題](https://qiita.com/hisonl/items/763b9d6d4e90b1606635)
- [N+1 問題](http://www.techscore.com/blog/2012/12/25/rails%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E7%B4%B9%E4%BB%8B-n1%E5%95%8F%E9%A1%8C%E3%82%92%E6%A4%9C%E5%87%BA%E3%81%99%E3%82%8B%E3%80%8Cbullet%E3%80%8D/)
- [N+1問題 / Eager Loading とは](http://ruby-rails.hatenadiary.com/entry/20141108/1415418367)
- [ActiveRecordのjoinsとpreloadとincludesとeager_loadの違い](https://qiita.com/k0kubun/items/80c5a5494f53bb88dc58)


