---
title: OpenAPIを使ったスキーマ駆動開発
description: OpenAPIを使ったスキーマ駆動開発について、基本的な概念から実践的な知見まで詳しく解説します。
slug: openapi-schema-driven-development
date: 2024-01-19T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - OpenAPI
  - スキーマ駆動
translation_key: openapi-schema-driven-development
---


# OpenAPI Specificationとは
言語依存のない形でHTTP APIの仕様を定義するためのフォーマット。YAMLまたはJSONで記述する。

Swagger SpecificationはOpenAPI Specificationの前身である。

# OpenAPIを採用するメリット・デメリット
## メリット
- REST APIの標準化されたフォーマットであり、開発者間のコミュニケーションコストを下げることができる
- ツールやエコシステムのが充実している（標準化の恩恵）
- APIドキュメントの自動生成ができる
- APIのクライアントとサーバーのコード生成自動化ができる
- テキストベースのため、バージョン管理システムと組み合わせることにより管理がしやすくなる
- モックサーバーを起動することができるため、APIの実装を待たずに開発を進めることができる

スキーマ駆動による開発プロセスの効率化ができる。

## デメリット
- 初めて導入する場合は仕様把握のための学習コストがかかる（設計の充実度合いにも影響する）
- 仕様と実装を同期するようにしないと仕様と実装とのズレが発生する可能性がある
- OpenAPI仕様は進化しているため、標準への追随をするコストが発生する
- 利用するツールによって予期しない挙動や仕様の解釈の違いが発生する可能性がある

致命的なデメリットはないように思える。

# 触ってみる
Dockerが利用できるのでDockerで試してみる。

cf. [github.com - OpenAPITools/openapi-generator](https://github.com/OpenAPITools/openapi-generator?tab=readme-ov-file#openapi-generator-cli-docker-image)

```sh
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml \
    -g go \
    -o /local/out/go
```

API仕様のドキュメントはVSCodeであれば[OpenAPI (Swagger) Editor](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi)が良さそうだった。

# 所感
生成のオプションとかちゃんと見れていないので、そのへん確認できたら個人開発のプロジェクトで利用してみようと思う。

何を自動生成して、何を自動生成させないかを上手く調整するのが導入時の課題かなと感じた。

# 参考
- [www.openapis.org](https://www.openapis.org/)
- [zenn.dev - 【Go言語】OpenAPI Generatorを使いこなすスキーマ駆動開発](https://zenn.dev/ysk1to/books/248fad8cb34abe)
- [medium.com - Generating Go code from OpenAPI Specification Document](https://medium.com/@MikeMwita/generating-go-code-from-openapi-specification-document-ae225e49e970)
