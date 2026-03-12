---
title: オレオレTechnology Radarを作る
description: オレオレTechnology Radarを作るについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: create-technology-radar
date: 2022-12-17T00:00:00Z
author: bmf-san
categories:
  - ツール
tags:
  - Technology Radar
translation_key: create-technology-radar
---


[Makuake Advent Calendar 2022](https://adventar.org/calendars/8496)の7日目の記事です！

# 概要
オレオレTechnology Radarの作り方についてかく。

# Technology Radarとは
Technology Radarとは、ThoughtWorks社（ソフトウェア開発やコンサルティングをグローバル展開している企業。マーティン・ファウラー氏が所属している。）が発信しているソフトウェア開発における技術調査の分析レポート。

[www.thoughtworks.com/radar](https://www.thoughtworks.com/radar)

年に2回ほど更新されているレポートで、技術トレンドを知ることができる。

[Archive](https://www.thoughtworks.com/radar/archive)も見ることできる。

Technology Radarは次の4象限で構成されている。

- テクニック
  - ex. 設計、開発プロセスなど
- ツール
  - ex. データベースやバージョン管理システムなど
- プラットフォーム
  - ex. 開発環境、クラウドなどのプラットフォームなど
- 言語とフレームワーク
  - ex. プログラミング言語、アプリケーションフレームワークなど

それぞれの象限は次の4つのリングに分類される。

- Hold
  - 保留。取り扱い注意。要調査
- Trial
  - 採用する価値がある。試験段階
- Assess
  - 価値があるかどうかコストをかけて調査する価値がある
- Adopt
  - 採用すべきと強く感じる技術

[www.thoughtworks.com/radar](https://www.thoughtworks.com/radar)のサイトでは、各々技術について評価コメントやブリップ（リング間の移動）の履歴なども見ることができる。

# オレオレTechnology Radarをつくる
Technology Radarを自作する方法が用意されているため、簡単に自作することができる。

[Build your own Radar](https://www.thoughtworks.com/radar/byor)

2つほど方法があるので紹介する。

## radar.thoughtworks.comにて作成する方法
[radar.thoughtworks.com](https://radar.thoughtworks.com/)にて、Google Spreadsheetのリンクを入力することで生成することができる。

この方法で生成した場合はパブリックになってしまうので注意。

## 自分でホストする方法
[github.com - thoughworks/build-your-own-radar](https://github.com/thoughtworks/build-your-own-radar)のリポジトリを利用することでRadarをセルフホスティングすることもできる。

dockerイメージも用意されているのでdockerイメージを使って触ってみる方法を紹介する。

### サンプルリポジトリをclone
オリジナルのリポジトリをforkするなりcloneするなりしてもよいが、試しやすいようにリポジトリを用意した。

[github.com - bmf-san/technology-radar-boilerplate](https://github.com/bmf-san/technology-radar-boilerplate)

#### 1. リポジトリをクローン
[github.com - bmf-san/technology-radar-boilerplate](https://github.com/bmf-san/technology-radar-boilerplate) をクローン

#### 2. コンテナ起動
`make run`

#### 3. Radar生成
`http://localhost:8080`にアクセス、`http://localhost:8080/files/radar.json`と入力して、`Build My Radar`を押下。

![form](/assets/images/posts/create-technology-radar/205632529-e528abd5-9013-458e-a540-6fc1251867e9.png)

#### 4. 生成されたRadarで遊ぶ。Radarをいじってみる。
Radarが生成されると次のようなリンクに飛ぶ。
`http://localhost:8080/?sheetId=http%3A%2F%2Flocalhost%3A8080%2Ffiles%2Fradar.json`

![radar](/assets/images/posts/create-technology-radar/205632536-d39195f1-2570-4645-bfb4-869bc7f77454.png)

`./files/radar.json`を編集することでRadarに表示するコンテンツを調整することができる。
（本当はjsonファイルをプロビジョンできるようにしたかったがフロントエンドのビルドの都合で厳しそうだった。。。）

Radarは右上の`Print this radar`から印刷することもできる。

# 所感
組織やチームにおいて採用している技術スタックや検証した技術を内外に公表できる形でこういったRadarを作成する試みは良さそう。
技術選定理由や選定・評価のプロセスを定義するといったことに活用できたり、技術のポートフォリオとして何に投資するのかといった考えを明らかにすることに役立ちだと思った。

[www.oreilly.co.jp - ソフトウェアアーキテクチャの基礎](https://www.oreilly.co.jp/books/9784873119823/)にも書いてあったが、あとは個人としてもこういうRadarを作成して定期的に更新するといった取り組みも良さそうだと感じた。自分が追っている技術をマッピングしてみたらいかに視野が狭いかということを気付かされそうだが・・

