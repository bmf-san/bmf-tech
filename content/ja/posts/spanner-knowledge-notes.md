---
title: Spannerの知見メモ
description: Spannerの知見メモについて調査・整理したメモ。基本概念と重要ポイントを解説します。
slug: spanner-knowledge-notes
date: 2024-03-03T00:00:00Z
author: bmf-san
categories:
  - データベース
tags:
  - Spanner
translation_key: spanner-knowledge-notes
---


# 概要
Cloud Spannerの知見を漁ったのでメモ。走り書きなのでカテゴライズしていない。

# メモ
- 計画メンテナンスやスキーマ変更のためのダウンタイムなし、最大99.999%の可用性を保証
  - Cloud SQLはダウンタイムありの計画のメンテナンス
- Cloud Spannerのアーキテクチャ構成
  - クライアント
    - R/Wするコードが実行されているサーバーやコンテナ
    - クライアントライブラリを使うとgRPCやREST APIで通信できる
  - ノード
    - R/W処理を実行するコンピュートリソース
    - ノードにはローカルストレージ（ディスク）はアタッチされていない
    - データは分散ファイルシステム（インターナルネーム Colossus）に保存される
      - ColossusはGoogle File Systemの後継のクラスタレベルのファイルシステム
        - Googleが運用する色々なサービスで使われているストレージ基盤らしい
  - Split
    - Cloud Spannerにおけるデータの単位
    - 特定のノードから管理される
    - 容量の上限は4GBで、超えると新しいSplitが生成され、データが分割される
- レプリケーション
  - 読み書きレプリカ
  - 読み取り専用レプリカ
  - ウィットネスレプリカ
    - データベースの完全なコピーを持たない、投票にだけ参加するレプリカ
- READしかしない箇所はRead-Only Transactionの利用を推奨。読み込みしかできないがロックを取らないため他のトランザクションをブロックすることがない。Read-Only Transactionの中身はリトライされても大丈夫なように冪等を保つ
- hotspotが生じていても1Node分の性能は出る。なので開発・試験環境で1Nodeで実行しているとhotspotに気づきにくい。負荷検証時は注意
- Nodeの増減だけではなくSplitが生成されることでデータが分散し、性能が出る
  - Splitはload-based split（write/readによる負荷上昇）、size-based split（データ量が増加）によって生成される
  - warmupしておきたいときなどは留意しておく
- インデックスも通常のテーブルと同じくテーブルに格納される
  - インデックスもインターリーブすることができる
-  インデックス生成はノード数やデータ量によって時間がかかるため注意（テーブル作成時が一番作成効率が良い）
- キーの特定の値にデータが集中することが想定される場合は、最初からインデックスを作成しないほうが良いかもしれない（ホットスポットが発生する可能性がある）
- 既にレコード数が多い既存テーブルにインデックスを追加する場合は、1日に追加するインデックスの数を気をつける（ドキュメント曰く3つくらいが良いらしい）
  - インデックスを追加しすぎるとSpannerのシステム負荷が高まってパフォーマンスが低下する可能性がある
- STORING句を使うとインデックスに追加の列を格納することができる
  - 参照が多い場合に特に有効。書き込みが多いと書き込みコストが高くなるため注意
- FORCE_INDEXと使うか、カバリングインデックスを使うかはクエリの実行計画を参照したほうが良さそう
- インデックス列の並び順は、不要なディスクシークを避けるために、取得したい順とするのが無難
- インターリーブは最大6段まで可能
- インターリーブされている子テーブルのインターリーブは外すことができない
- mutationの数え方
  - トランザクションに含まれるColumn*rowの数
    -  [Mutations per commit (including indexes) : 20000](https://cloud.google.com/spanner/quotas?hl=en#limits_for_creating_reading_updating_and_deleting_data)
  - mutationの上限を意識するときは注意
  -  [Google Cloud Spanner Deep Dive](https://docs.google.com/presentation/d/1XKaOrex3WS8xZ0TsjsTQKBgxjVSGxkyqdcIqG66zd64)を参照

# 所感
設計を間違うとspannerのスケーラビリティを活かせず十分なパフォーマンスが出ない。
計画メンテやスキーマ変更にダウンタイムなしは運用上の大きなメリット。
Splitの気持ちになってデータの分散を意識しないと運用時最大限のパフォーマンスを発揮できない。
日本語の記事ばかり漁ってたので海外の記事ももっと見ておいたほうが良いかも。

# 参考
- [Cloud Spanner を使って様々な Anomaly に立ち向かう](https://medium.com/google-cloud-jp/cloud-spanner-%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E6%A7%98%E3%80%85%E3%81%AA-anomaly-%E3%81%AB%E7%AB%8B%E3%81%A1%E5%90%91%E3%81%8B%E3%81%86-5132f691ccf4)
- [Spannerを本番環境で使ってみたので感想戦](https://syossan.hateblo.jp/entry/2020/10/17/162121)
- [github.com - gcpug/nouhau](https://github.com/gcpug/nouhau/tree/master/spanner)
- [GCPUGまとめ「Cloud Spannerでセカンダリインデックスを使うときの勘所」](https://engineering.dena.com/blog/2020/07/spanner-secondary-index/)
- [大規模なスキーマ更新のオプション](https://cloud.google.com/spanner/docs/schema-updates?hl=ja#large-updates)
- [Cloud Spannerの概要と設計上の要点など](https://qiita.com/pa_pa_geno/items/c5fc5ea0d5daf415080d#%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%AA%E3%83%BC%E3%83%96)
- [Cloud Spanner のハイレベルアーキテクチャ解説](https://medium.com/google-cloud-jp/cloud-spanner-%E3%81%AE%E3%83%8F%E3%82%A4%E3%83%AC%E3%83%99%E3%83%AB%E3%82%A2%E3%83%BC%E3%82%AD%E3%83%86%E3%82%AF%E3%83%81%E3%83%A3%E8%A7%A3%E8%AA%AC-fee62c17f7ed)
- [Colossus の仕組み: Google のスケーラブルなストレージ システムの舞台裏](https://cloud.google.com/blog/ja/products/storage-data-transfer/a-peek-behind-colossus-googles-file-system)
- [Spannerを解説したら講義になった話](https://cloudpack.media/52502)
- [2020年現在のNewSQLについて](https://qiita.com/tzkoba/items/5316c6eac66510233115)
- [Cloud Spanner のマルチリージョン構成について理解する](https://cloud.google.com/blog/ja/topics/developers-practitioners/demystifying-cloud-spanner-multi-region-configurations)
- [メルペイでのSpannerとの戦いの日々](https://engineering.mercari.com/blog/entry/2019-04-18-090000/)
- [2020/12/06
本番稼働中の Spanner にダウンタイム無しに57時間かけてインデックスを追加して得た知見](https://engineering.mercari.com/blog/entry/20201203-3f9780edf7/)
- [Google Cloud Spanner Deep Dive](https://docs.google.com/presentation/d/1XKaOrex3WS8xZ0TsjsTQKBgxjVSGxkyqdcIqG66zd64)
- [超実践 Cloud Spanner 設計講座](https://www.slideshare.net/HammoudiSamir/cloud-spanner-78081604)
