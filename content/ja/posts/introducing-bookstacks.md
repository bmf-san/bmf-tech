---
title: Bookstacks — ISBNバーコードで本を管理するiPhoneアプリの紹介
description: Flutter と Riverpod で開発した本棚管理アプリ「Bookstacks」の紹介。ISBNバーコードをスキャンするだけで書籍を登録し、カスタムラベルで整理できる。
slug: introducing-bookstacks
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Flutter
  - iOS
  - Riverpod
  - Dart
translation_key: introducing-bookstacks
---

# Bookstacks — ISBNバーコードで本を管理するiPhoneアプリの紹介

## 作った理由

本棚の管理は大して難しくないように思えるが、書店で「この本、すでに持っていたっけ」と迷う瞬間が誰にでもある。積読が積み重なっているうちに何があるかわからなくなることも多い。既存のアプリはアカウント登録が必要だったり、手入力が面倒だったりすることが多かった。

そこでBookstacksは、「バーコードをスキャンするだけで登録できる」という体験に特化して開発した。アカウント不要、サブスクリプションなし。書籍のISBNバーコードを使ってOpenBD API（無料の書誌情報サービス）からメタデータを取得し、重複チェックも自動で行う。ラベルを使って本を分類する機能も備えている。

## アーキテクチャ

Clean ArchitectureのDomain・Application・Infrastructure・Presentationの4層をRiverpodのProviderで接続している。ローカル永続化にはHiveを使用しており、書籍データがキャッシュされのちはオフラインでも動作する。外部への依存はOpenBD APIのみだ。

![ホーム画面](/assets/images/posts/introducing-bookstacks/01_home.png)

## ISBNバーコードスキャン × OpenBD API — 非同期フロー

スキャンから永続化までの登録フローが、このアプリで最も興味深い技術的な部分だ。シーケンスは次のようになっている。

1. **スキャン** — `mobile_scanner`が`BarcodeCapture`イベントを発火し、検出された値を取得する。ISBN-13（978/979プレフィックス、13桁）とISBN-10（10桁）の両形式を受け付け、それ以外のバーコードは無視する。

2. **メタデータ取得** — `OpenBdDatasource`が`http.Client`を使ってOpenBDエンドポイント`https://api.openbd.jp/v1/get?isbn={isbn}`を呼び出す。レスポンスはJSON配列で、最初の要素が完全な書籍オブジェクトか`null`（未収録）のどちらかだ。`OpenBdDatasource`がレスポンスをパースし、`isbn`・`title`・`author`・`category`・`coverImageUrl`を持つ`OpenBdBookData`を返す。

3. **重複チェック** — 確認画面でRiverpodの`booksNotifierProvider`を参照し、同じISBNの本がすでに登録済みかどうかを検出する。登録済みの場合は登録ボタンを無効化してユーザーへ通知する。

4. **永続化** — バリデーション済みの`Book`エンティティを`AddBook`ユースケース経由で`BookRepository`インターフェースに渡し、Hiveに書き込む。`BooksNotifier`がリストを再取得し、本棚グリッドの再レンダリングをトリガーする。

メタデータ取得はスキャン画面が`OpenBdDatasource`を直接呼び出して行う。永続化は`AddBook`ユースケースを経由するため、プレゼンター層はHiveの実装詳細を直接扱う必要がない。

![書籍詳細](/assets/images/posts/introducing-bookstacks/02_detail.png)

## 主要機能

### 本棚ビュー — グリッドとリスト

デフォルトでは`cached_network_image`を使って表紙画像を遅延ロードする3列グリッドに書籍を表示する。トグルボタンでリストビューに切り替えられる。

![グリッドビュー](/assets/images/posts/introducing-bookstacks/01_home.png)
![リストビュー](/assets/images/posts/introducing-bookstacks/03_list.png)

### ラベルシステム

Bookstacksには「読了」「読んでいる」「積読」「あとで読む」「買いたい」の5つのラベルがプリセットされている。さらに自由な名前でカスタムラベルを追加できる。ラベルはHiveに保存され、書籍と多対一の関係で紐付けられる。ラベルチップをタップすると即座に絞り込まれる。

![ラベル管理](/assets/images/posts/introducing-bookstacks/05_labels.png)

### 設定

設定画面ではラベル管理・バージョン情報・プライバシーポリシー・書誌情報提供元を確認できる。テーマはシステム設定に自動的に連動する。ホーム画面にはAdMobバナー広告が表示され、最大3回・5秒間隔の自動リトライ機構を備えている。

![設定画面](/assets/images/posts/introducing-bookstacks/04_settings.png)

## まとめ

BookstacksはApp Storeで公開している。

- **App Store**: [Bookstacks](https://apps.apple.com/jp/app/bookstacks-%E6%9C%AC%E6%A3%9A%E7%AE%A1%E7%90%86/id6760252143)
