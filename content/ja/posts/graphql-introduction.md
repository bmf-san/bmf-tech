---
title: "GraphQLについて"
slug: "graphql-introduction"
date: 2023-11-09
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "GraphQL"
draft: false
---

# 概要
GraphQLの素振りをしていたので調べたことについてまとめておく。

親切なチュートリアルが用意されており、始めやすい。
cf. [www.howtographql.com](https://www.howtographql.com/)

# GraphQLとは
Meta社によって開発されたWeb API開発のためのクエリ言語。

GraphQLはGraphQL Foundationによって管理されており、Meta社はその一員である。

GraphQLの仕様と全ての関連プロジェクトはOSSとして公開されている。

# 特徴
- HTTP上で使用される
- クエリをPOSTリクエストすることでデータを取得する
- 単一のエンドポイントで複数のデータを取得できる
- GraphQLのスキーマやクエリは有向グラフの考え方に基づいて設計されている
- スキーマファースト
- ドキュメンテーション
  - スキーマファーストであるので、仕様が自明
  - ドキュメントジェネレーターによりドキュメントを生成することもできる
- 型定義
  - スカラー型とオブジェクト型、列挙型、リスト型、Non-null型、ユニオン型、入力型、インターフェースがある
    - 入力型はQueryやMutationの引数に使用するオブジェクト型
    - デフォルトはNullable
      - [maku.blog - NullableとNon-nullのガイドライン](https://maku.blog/p/4reqy9i/#nullable-%E3%81%A8-non-null-%E3%81%AE%E3%82%AC%E3%82%A4%E3%83%89%E3%83%A9%E3%82%A4%E3%83%B3)
- データ取得の柔軟性
  - 必要なデータだけをQueryによって取得できる
  - オーバーフェッチやアンダーフェッチの発生が回避できる
- バージョン管理
  - バージョン管理することはできる
  - 新しい型やフィールド追加によってバージョンレスで開発することもできる
  - 思想としては回避派
- エラーハンドリング
  - エラーでもステータスコード200を返し、エラーメッセージを内包して返却するのが一般的


# 用語
いくつかピックアップしたものだけ記載。

## Schema
クエリの型定義。

```graphql
type Query {
  user: User
}
```

## Query
データ取得のためのクエリ。

```graphql
query {
  user {
	name
  }
}
```

## Mutation
データ更新のためのクエリ。

```graphql
mutation {
  updateUser {
	name
  }
}
```

## Subscription
データの変更を監視するためのクエリ。

```graphql
subscription {
  user {
	name
  }
}
```

## Argument
クエリにわたす引数。

```graphql
{
  user(id: 123) {
    username
    email
  }
}
```

# 関連技術
## GraphQL Mesh
gRPC、OpenAPI、Swagger、oData、SOAP、GraphQL等々のAPI仕様で実装されたAPIへのゲートウェイサーバー（GraphQL Gateway）。

API仕様さえあればAPIに対してGraphQLクエリでアクセスできる。

cf. [the-guild.dev](https://the-guild.dev/graphql/mesh)

## openapi-to-graphql
OpenAPIに基づいたAPI仕様をGraphQLのSchemaに変換する。

cf. [github.com - IBM/openapi-to-graphql](https://github.com/IBM/openapi-to-graphql)

## graphql-tools
GraphQLのSchemaを作成するための便利ツール。モックを作成することもできる。

cf. [github.com - ardatan/graphql-tools](https://github.com/ardatan/graphql-tools)

## GraphQL Gateway


# パフォーマンス
- N+1
  - [Data Loader](https://github.com/graphql/dataloader)
    - 同一テーブルへの複数のSELECTを1つのSELECTにまとめる
    - BatchとCacheという機能がある
      - cf. [lyohe.github.io - 主要な機能](https://lyohe.github.io/blog/2021/12/16/reading-dataloader/#%E4%B8%BB%E8%A6%81%E3%81%AA%E6%A9%9F%E8%83%BD)
- Queryが複雑になるとリクエストボディが肥大化する
  - Persisted Query
    - ApolloというGraphQLのツールが提供している機能
    - クエリに対応するIDを用意しておき、IDとクエリを交換してリクエストパラメータを減らす仕組み
      - 交換のためのエンドポイントをGETにすることによりキャッシュに乗せることができる
- POSTリクエストであるため、HTTPキャッシュが利用できない
  - Persisted Query

# 所感
Restful APIと比較して始めるのに少し手間はかかりそうだが、得られる恩恵は大きそう。

GraphQLクライアントは色々種類があるようで、選定には悩みそう。

cf. [user-first.ikyu.co.jp - あなたのプロダクトに Apollo Client は必要ないかもしれない](https://user-first.ikyu.co.jp/entry/2022/07/01/121325)

# 参考
- [graphql.org](https://graphql.org/)
- [zenn.dev - GraphQLを徹底解説する記事s](https://zenn.dev/nameless_sn/articles/graphql_tutorial)
- [kinsta.com - GraphQLとRESTの比較─知っておきたい両者の違い](https://kinsta.com/jp/blog/graphql-vs-rest/#:~:text=GraphQL%E3%81%AE%E6%9C%80%E5%A4%A7%E3%81%AE%E5%88%A9%E7%82%B9,%E3%81%8C%E5%A4%9A%E3%81%99%E3%81%8E%E3%82%8B%E3%81%93%E3%81%A8%E3%81%A7%E3%81%99%E3%80%82)
- [panda-program.com - 「初めてのGraphQL」を読んでGraphQLの概要を学んだ](https://panda-program.com/posts/book-learning-graphql)
- [qiita.com - GraphQLのスキーマと型定義](https://qiita.com/NagaokaKenichi/items/d341dc092012e05d6606)
- [gist.github.com - GraphQL についてまず知っておきべきこと](https://gist.github.com/tkdn/75a4d7e38c2edb07b41da078e4a4aa11)
- [maku.blog - GraphQL ベストプラクティス](https://maku.blog/p/4reqy9i/)
- [qiita.com - GraphQL でのパフォーマンス関連の問題について色々紹介する](https://qiita.com/haradakunihiko/items/7148a8b36f1e4e5d60b1)
- [engineering.mercari.com - GraphQLを導入する時に考えておいたほうが良いこと](https://engineering.mercari.com/blog/entry/20220303-concerns-with-using-graphql/)
- [engineering.mercari.com - メルカリ Shops での NestJS を使った GraphQL Server の実装](https://engineering.mercari.com/blog/entry/20210818-mercari-shops-nestjs-graphql-server/#dataloader-for-batch-request)
- [lyohe.github.io - graphql/dataloader を読んだ話](https://lyohe.github.io/blog/2021/12/16/reading-dataloader/)
- [zenn.dev - GraphQL で N+1 問題を解決する 4 つのアプローチ](https://zenn.dev/alea12/articles/15d73282c3aacc#%E6%96%B9%E6%B3%954%3A-n%2B1-%E3%82%92%E8%80%83%E6%85%AE%E3%81%97%E3%81%9F-orm-%E3%82%92%E6%A4%9C%E8%A8%8E%E3%81%99%E3%82%8B)
- [www.apollographql.com - GraphQL Concepts Visualized](https://www.apollographql.com/blog/graphql/basics/the-concepts-of-graphql/)
