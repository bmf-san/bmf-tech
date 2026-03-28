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

## なぜ作ったか

本棚の管理はそれほど難しくないように思えるが、書店で「この本、すでに持っていたっけ」と迷う瞬間が誰にでもある。積読が積み重なるうちに何があるかわからなくなることも多い。既存のアプリはアカウント登録が必要だったり、手入力が面倒だったりすることが多かった。

そこで、「バーコードをスキャンするだけで登録できる」という体験に特化してBookstacksを開発した。アカウント不要、サブスクリプションなし。ISBNバーコードを使ってOpenBD API（無料の書誌情報サービス）からメタデータを取得し、重複チェックも自動で行う。ラベルで本を分類する機能も備えている。

[App Store](https://apps.apple.com/jp/app/bookstacks-%E6%9C%AC%E6%A3%9A%E7%AE%A1%E7%90%86/id6760252143)で公開しているので、興味があればぜひ試してみてほしい。

## ユースケース

- **重複確認** — 書店でバーコードをスキャンし、購入前にすでに持っているかどうかを確認する
- **積読管理** — 未読の本に「積読」ラベルを付けて、次に読む本をリストから選ぶ
- **読書記録** — 読み終えた本に「読了」ラベルを付けて、本棚を整理する
- **欲しい本リスト** — 購入候補を「買いたい」ラベルでまとめる

## 主な機能

![ホーム画面](/assets/images/posts/introducing-bookstacks/ja/01_home.png)

### バーコードで書籍登録

ISBNバーコードをスキャンすると、タイトル・著者・カテゴリ・表紙画像をOpenBD APIから自動取得する。ISBN-13（978/979プレフィックス）とISBN-10の両形式に対応している。同じISBNの本がすでに登録済みの場合は重複を検出して登録をブロックする。

![書籍詳細](/assets/images/posts/introducing-bookstacks/ja/02_detail.png)

### ラベルで整理

「読了」「読んでいる」「積読」「あとで読む」「買いたい」の5つのラベルがプリセットされている。自由な名前でカスタムラベルを追加できる。ラベルチップをタップすると即座に絞り込まれる。

![ラベル管理](/assets/images/posts/introducing-bookstacks/ja/05_labels.png)

### 本棚ビュー（グリッド・リスト）

デフォルトでは表紙画像付きの3列グリッドで書籍を表示する。トグルボタンでリストビューに切り替えられる。

![グリッドビュー](/assets/images/posts/introducing-bookstacks/ja/01_home.png)
![リストビュー](/assets/images/posts/introducing-bookstacks/ja/03_list.png)

### 設定

設定画面ではラベル管理・バージョン情報・プライバシーポリシー・書誌情報提供元を確認できる。テーマはシステム設定に自動的に連動する。

![設定画面](/assets/images/posts/introducing-bookstacks/ja/04_settings.png)

## 技術スタック

| 層 | 技術 |
|---|---|
| UI | Flutter (iOS) |
| 状態管理 / DI | Riverpod + riverpod_generator |
| 永続化 | Hive |
| バーコードスキャン | mobile_scanner |
| 書籍メタデータ | OpenBD API (http) |
| 画像キャッシュ | cached_network_image |
| テスト | flutter_test / mocktail |

アーキテクチャはDomain・Application・Infrastructure・Presentationの4層構成で、RiverpodがDIと状態管理を担う。Hiveがすべてのデータをローカルに保存するため、一度キャッシュされた書籍データはオフラインでも参照できる。

## ISBNスキャンと書籍登録の仕組み

スキャンから永続化までの登録フローは4ステップで成り立っている。

1. **スキャン** — バーコードを読み取る。ISBN-13（978/979プレフィックス、13桁）とISBN-10（10桁）のみを受け付け、それ以外の形式は無視する。

2. **メタデータ取得** — `https://api.openbd.jp/v1/get?isbn={isbn}` を呼び出す。レスポンスはJSON配列で、最初の要素が書籍オブジェクトか`null`（未収録）のどちらかになる。タイトル・著者・カテゴリ・表紙画像URLを取り出す。

3. **重複チェック** — 確認画面で同じISBNの本がすでに登録済みかどうかを検出する。登録済みの場合は登録ボタンを無効化してユーザーへ通知する。

4. **永続化** — バリデーション済みの書籍エンティティをユースケース経由でHiveに書き込む。書き込み後にリストを再取得し、本棚グリッドを再描画する。

## まとめ

Bookstacksは、本棚の管理をできるだけ手間なく行いたいと思って作ったアプリである。バーコードをスキャンするだけで書籍を登録できる点は、自分で使っていて気に入っている。

良ければダウンロードしてみてほしい。

- **App Store**: [Bookstacks](https://apps.apple.com/jp/app/bookstacks-%E6%9C%AC%E6%A3%9A%E7%AE%A1%E7%90%86/id6760252143)
