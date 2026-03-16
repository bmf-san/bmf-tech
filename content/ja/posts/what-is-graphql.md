---
title: GraphQLとは
description: GraphQLとは
slug: what-is-graphql
date: 2018-06-14T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - API
  - HTTP
  - REST
  - GraphQL
translation_key: what-is-graphql
---


# GraphQLとは
- Facebookが開発
- APIのためのクエリ言語
  - APIのリクエストのデータ形式とレスポンスのデータ形式が似ているため、ユーザーフレンドリー
- RESTはアーキテクチャ（設計）であり、GraphQLは言語（DSL）である

# REST APIとGraphQLの比較

## REST APIのAPI形式
エンドポイントに対して、HTTP動詞でリクエストを投げる

`curl https://api.bmf-tech.com/v1/configs`

```
[
    {
        "id": 1,
        "name": "title",
        "alias_name": "Title",
        "value": "bmf-tech",
        "created_at": "2017-09-25 23:08:23",
        "value": "bmf-tech",
        "deleted_at": null
    }
]
```

# GraphQLのAPI形式
単一のエンドポイントに対し、クエリを投げる

`curl https://api.bmf-tech.com/api`

```
configs {
  id,
  name,
  alias_name
  value,
  created_at,
  updated_at,
  deleted_at
}
```

```
[
    {
        "id": 1,
        "name": "title",
        "alias_name": "Title",
        "value": "bmf-tech",
        "created_at": "2017-09-25 23:08:23",
        "value": "bmf-tech",
        "deleted_at": null
    }
]
```

| | REST API | GraphQL |
|:-----------|:-----------:|:------------:|
| エンドポイント | 複数 | 単一 |
| HTTP動詞 | 依存している | 依存していない |
| 型システム | 無し | 有り |
| バージョニングの必要 | 有り | 無し |
| ドキュメントの必要性 | 有り | 無し |
| リソース制限 | コール回数が主 | リソース量に応じて対応 |

- 単一エンドポイントに対して欲しいデータを柔軟に指定
  - RESTはエンドポイントごとにレスポンスデータが決まっている。
  - GraphQLは単一エンドポイントに対して欲しいデータを指定してレスポンスデータを得る。

- リソース制限には工夫が必要
  - リソース量に応じて対応 
  - オブジェクトの数をベースにするなど負荷計算の方法を考える必要がある

- ドキュメントの必要性がほぼない
  - API定義がそのままドキュメント代わりになる
    - クエリの構造とレスポンスのデータ構造がほぼ同じ

# 気になった点
- ライブラリ依存になる
  - クエリのパースをするためのライブラリが必要

- 必ずしもREST APIよりパフォーマンスがよくなるわけではなさそう
  - リクエストの回数が減る
  -  一回のデータ量が増える
  - REST APIでもGraphQLでもデータ量をコントロールする工夫などが必要（ページングとかフィールド指定とか）

- モニタリング
  - REST APIはエンドポイントごとにモニタリングできる
  - GraphQLは単一エンドポイントなので、クエリごとにレスポンス性能を監視しずらい。何らかの対応が必要。
    - エコシステムの充実を待つか、自前実装
 
- キャッシュ周り
   - HTTPキャッシュが使えない
   - その他色々調べておいたほうが良さそう

# 所感
- コンポーネントが沢山あって複雑なUIを備えたアプリケーションにおいて、リクエスト数が増えてクライアント側が辛い、というようなシチュエーションなら導入するメリットはありそう。
- [Rubel](https://github.com/bmf-san/Rubel)で使ってみようかと考えていたが、時期尚早感があるのでやめる。というか現時点では不要である。

# 参考
- [graphql.org](https://graphql.org/)
- [facebook/graphql rfcs](https://github.com/facebook/graphql/blob/master/rfcs/Subscriptions.md)
- [「GraphQLは何に向いているか」に対してのちょっとした反論](http://yamitzky.hatenablog.com/entry/graphql)
- [アプリ開発の流れを変える「GraphQL」はRESTとどう違うのか比較してみた](https://www.webprofessional.jp/rest-2-0-graphql/) 
- [GraphQLはRESTの置き換えではない](https://note.mu/konpyu/n/nc4fd122644a1)
- [GraphQL APIをRailsアプリに実装した時のメモ](https://blog.qnyp.com/2017/06/08/graphql-resources/)
- [初心者目線で分かりやすく解くGraphQLを解説！～同じWebAPIのRESTとの違いは？～](https://vitalify.jp/app-lab/20171006-graphql/)

