---
title: gRPCとProtocol Buffersについて復習するためのリンク集
description: gRPCとProtocol Buffersについて復習するためのリンク集
slug: grpc-protocol-buffers-review
date: 2024-01-25T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - gRPC
  - Protocol Buffers
  - リンク集
translation_key: grpc-protocol-buffers-review
---


# 概要
gRPCとProtocol Buffersについて改めておさらいしておきたかった。

# リンク
## gRPC
- [grpc.io](https://grpc.io/)
  - 公式ドキュメント
- [zenn.dev - 作ってわかる！ はじめてのgRPC](https://zenn.dev/hsaki/books/golang-grpc-starting)
  - まだ読んだことなかったので読んでみたのだが大変勉強になった。日本語情報はまずこれ読んでみると良いかも。
- [www.wantedly.com - gRPC Internal - gRPC の設計と内部実装から見えてくる世界](https://www.wantedly.com/companies/wantedly/post_articles/219429)
  - gRPC内部についてディープダイブしている記事
- [cloud.google.com](https://cloud.google.com/apis/design?hl=ja)
  - RESTもgRPCもAPIの設計方針に迷ったら参照すると良いらしい。特にgRPC寄りで書かれている。
- [grpc.io - gRPC Load Balancing](https://grpc.io/blog/grpc-load-balancing/)
  - gRPCサーバーのロードバランシングについて
- [hakobe932.hatenablog.com - gRPCのロードバランシング](https://hakobe932.hatenablog.com/entry/2018/04/11/123000)
  - gRPCサーバーのロードバランシングについて解説
- ~~deeeet.com - Kubernetes上でgRPCサービスを動かす~~
  - K8S上でgRPCを扱う上での観点
- [scrapbox.io - gRPCのベストプラクティス](https://scrapbox.io/dojineko/gRPC%E3%81%AE%E3%83%99%E3%82%B9%E3%83%88%E3%83%97%E3%83%A9%E3%82%AF%E3%83%86%E3%82%A3%E3%82%B9)
  - gRPCおよびProtocol Buffersのベストプラクティス

## Protocol Buffers
- [protobuf.dev](https://protobuf.dev/)
  - 公式ドキュメント
- [qiita.com - Proto2 vs Proto3](https://qiita.com/ksato9700/items/0eb025b1e2521c1cab79)
  - Proto2とProto3の違いについて
- [docs.wantedly.dev - protobufスキーマとgRPC通信](https://docs.wantedly.dev/fields/the-system/apis)
  - gRPCとprotobuf入門知識についてとベストプラクティス
- [www.wantedly.com - Protocol Buffers によるプロダクト開発のススメ - API 開発の今昔 -](https://www.wantedly.com/companies/wantedly/post_articles/309513)
  - Protocol Buffersについてメリットについて

## ツール
- [buf](https://buf.build/)
  - Protocol Buffers周りのツールとして筆頭に上がると思う。

日本語の記事をいくつか読み漁ったが、まだまだ日本語情報は少ないので、海外の事例記事とか読み漁ったほうが良いかも。

# 相互リンク
- [エンジニアリングを進化させる品質メディア：Sqripts](https://sqripts.com/)
