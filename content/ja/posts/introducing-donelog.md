---
title: Done Log — 毎日のルーティンタスクを管理するiPhoneアプリの紹介
description: Flutter と Riverpod で開発した習慣トラッカーアプリ「Done Log」の紹介。柔軟な繰り返しルールエンジン（毎日・N日おき・曜日指定・1回のみ）でルーティンタスクを自動リセットする。
slug: introducing-donelog
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Flutter
  - iOS
  - Riverpod
  - Dart
translation_key: introducing-donelog
---

# Done Log — 毎日のルーティンタスクを管理するiPhoneアプリの紹介

## なぜ作ったか

ルーティンタスクはGoogle Keepで管理していたのだが、毎日の完了確認のような手段がなかったので、ルーティンタスクに特化したタスク管理アプリを作ってみた。

ワンタップでタスク完了が記録できて、繰り返しルールに従って自動でリセットされる。管理の手間と記録しやすさに特化して開発した。

[App Store](https://apps.apple.com/jp/app/done-log/id6759606196)で公開しているので、興味があればぜひ試してみてほしい。

## ユースケース

- **毎日の服薬確認** — 朝の薬を飲んだかを記録し、当日の進捗を確認する
- **週次ルーティンの管理** — 月・水・金の運動など、曜日を指定した繰り返しタスクを管理する
- **習慣の振り返り** — カレンダービューで過去の完了記録を見て、習慣の継続状況を把握する
- **N日おきのルーティン** — 2週間に1回の掃除など、任意の間隔で繰り返すタスクを登録する

## 主な機能

![今日のタスク一覧](/assets/images/posts/introducing-donelog/ja/01_today.png)

### 今日のタスクビュー

今日の日付と繰り返しルールにマッチするタスクが「今日」欄にリストアップされる。完了済みタスクは取り消し線とチェックマークで表示され、その日の進捗をひと目で確認できる。

![完了したタスク](/assets/images/posts/introducing-donelog/ja/02_task_completed.png)

### タスク登録フォーム

タスクフォームでは4種類の繰り返しタイプを選択できる。`everyNDays`では2〜30日の間隔をスライダーで指定できる。`weekdays`では複数の曜日チェックボックスで任意のパターン（例：月・水・金のみ）を設定できる。

![タスクフォーム](/assets/images/posts/introducing-donelog/ja/03_task_form.png)

### カレンダービュー

過去の完了記録を月間カレンダーで可視化する。任意の日付をタップすると、その日に完了したタスクの一覧を確認できる。習慣の継続状況を振り返ったり、タスクをこなせなかった日を把握したりするのに便利だ。

![カレンダービュー](/assets/images/posts/introducing-donelog/ja/04_calendar.png)

### 設定とダークモード

言語（日本語・英語）とテーマ（ライト・ダーク・システム連動）を設定画面で変更できる。通知リマインダーを任意の時間帯にスケジュールできる。

![設定画面](/assets/images/posts/introducing-donelog/ja/05_settings.png)
![ダークモード](/assets/images/posts/introducing-donelog/ja/06_dark_mode.png)

## 技術スタック

| 層 | 技術 |
|---|---|
| UI | Flutter (iOS) |
| 状態管理 / DI | Riverpod + riverpod_generator |
| 永続化 | Hive |
| カレンダー | table_calendar |
| 通知 | flutter_local_notifications |
| テスト | flutter_test / mockito |

アーキテクチャはDomain・Application・Infrastructure・Presentationの4層構成で、RiverpodがDIと状態管理を担う。Hiveがすべてのデータをローカルに保存するため、アプリは完全にオフラインで動作する。タスク定義と完了履歴は別テーブルに保存されるため、繰り返しルールを変更しても過去の記録は失われない。

## 繰り返しルールエンジン

繰り返しルールエンジンは、「このタスクを今日表示すべきか」を判断する仕組みだ。毎日・N日おき・曜日指定・1回のみという4種類のルールをサポートしており、タスクごとに異なるリズムで自動的にリセットされる。

ルールの評価はドメイン層に閉じており、データベースやUIとは切り離されている。そのため、ルールの変更や新しいタイプの追加が他の部分に影響を与えにくい構造になっている。

リセットはアプリ起動時に自動で実行される。ユーザーが意識しなくても、翌日には前日のタスクがリセットされた状態で表示される。

## まとめ

Done Logは、毎日のルーティンタスクをできるだけ手間なく管理したいと思って作ったアプリである。ワンタップで記録できる点と、繰り返しルールによる自動リセットが自分で使っていて気に入っている。

良ければダウンロードしてみてほしい。

- **App Store**: [Done Log](https://apps.apple.com/jp/app/done-log/id6759606196)
