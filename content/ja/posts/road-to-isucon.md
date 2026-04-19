---
title: Road to ISUcon
description: Road to ISUcon
slug: road-to-isucon
date: 2023-06-07T00:00:00Z
author: bmf-san
categories:
  - アーキテクチャ
tags:
  - パフォーマンスチューニング
  - ISUCON
  - ISUCON8
translation_key: road-to-isucon
---


# 概要
ISUcon出場に向けて準備したことを記す。

# 目標・目的を定めた
- 目標
    - ISUconの予選の時間を目一杯使い切ってチューニングをする
        - 勝つことが目標ではあるが、初参戦なので現実的なラインの目標を立てた
- 目的
    - インフラ周りの知見を高める
    - パフォーマンスを考慮したアプリケーション構築のための知見を高める
    - 社内のパフォーマンス・チューニングの業務に携われるようにする（頑張る）

# メンバーを募集した
会社の同僚を誘い、2人チームで参戦。

# パフォーマンス・チューニングの手順までを確認した
## 準備
- レギュレーション確認　
- 環境構築
- sshの設定
- バックアップ取得
- 自動デプロイ構築
 - 簡易的なスクリプト
  - ssh接続してデプロイタスク実行、ミドルウェア再起動など
- モニタリングツール導入
  - netdata
- プロファイリングツール導入
  - alp
    - アクセスログプロファイラ
    - https://github.com/tkuchiki/alp
  - php
    - xhprof
- ベンチマーク

## 構成確認
- プロセス確認
  - 不要なものは止める
- ハードウェアリソース
  - top
- データベース
  - MySQL
    - データサイズ確認
      - テーブルごとのサイズや行数を確認
```
mysql> use database;
mysql> SELECT table_name, engine, table_rows, avg_row_length, floor((data_length+index_length)/1024/1024) as allMB, floor((data_length)/1024/1024) as dMB, floor((index_length)/1024/1024) as iMB FROM information_schema.tables WHERE table_schema=database() ORDER BY (data_length+index_length) DESC;
```
- アプリケーション
  - URL洗い出し
  - コードリーディング

## チューニング
- パフォーマンス計測→分析→チューニング→計測
 - CPUなのか、IOなのか、アプリケーションなのか、DBなのか

# 所感
初参戦して人権を失ってきた。
練習量が足りていないのはともかく、それ以前に基本的なことができていなかったので、来年はちゃんと戦えるように準備、成長していきたい。
自分の課題を見つめ直したり、今後のモチベーションに大きな影響を与えたりする機会となってそういった意味でも参加することができて良かったと思う。
なんでこれが無料なんだろうというくらいの大会で、運営の皆様方に大変感謝しています。

# 参考リポジトリ
- [ISUCON過去問を構築するためのVagrantfile集](https://github.com/matsuu/vagrant-isucon)

# 参考リンク
- [ISUcon公式サイト](http://isucon.net/)
- [「Webアプリケーションの パフォーマンス向上のコツ」の発表資料](http://blog.nomadscafe.jp/2014/08/isucon-2014-ami.html)
- [Webアプリケーションの パフォーマンス向上のコツ 実践編](https://www.slideshare.net/kazeburo/isucon-summerclass2014action2)
- [ISUCON初心者のためのISUCON7予選対策
](http://isucon.net/archives/50697356.html)
- [ISUCONの勝ち方動画メモ](https://wiki.infra-workshop.tech/user/kameneko/ISUCON8/ISUCON%E3%81%AE%E5%8B%9D%E3%81%A1%E6%96%B9%E5%8B%95%E7%94%BB%E3%83%A1%E3%83%A2)
- [ISUCON7(予選)過去問環境をConoHaで構築する手順](https://nishinatoshiharu.com/how-to-create-isucon-conoha/)
- [YAPC::Asia Tokyo 2015 「ISUCONの勝ち方」メモ](http://kobtea.net/posts/2015/08/22/yapc-isucon/)
- ~~来年の isucon で泣かずに予選を終える為のメモ（1）~~

# ISUconに向けて読んだ本
- [詳解 システムパフォーマンス](https://www.oreilly.co.jp/books/9784873117904/)
    - まだ読み終わっていない....
- [絵で見てわかるシステムパフォーマンスの仕組み](https://www.shoeisha.co.jp/book/detail/9784798134604)
- [［24時間365日］サーバ/インフラを支える技術 ……スケーラビリティ，ハイパフォーマンス，省力運用](https://gihyo.jp/magazine/wdpress/plus/978-4-7741-3566-3)


