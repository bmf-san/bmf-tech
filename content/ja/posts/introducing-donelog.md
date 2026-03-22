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

## 作った理由

汎用タスク管理アプリはプロジェクト管理には優れているが、「今朝の薬を飲んだか」「夜の戸締まりはしたか」といった毎日のルーティン確認には機能が多すぎる。Done Logはシンプルな目的に特化している。毎日のルーティンタスクをワンタップで記録し、繰り返しルールに従って自動リセットする。管理の手間をゼロにするアプリだ。

## アーキテクチャ

Clean ArchitectureのDomain・Application・Infrastructure・Presentationの4層をRiverpodのProviderで接続している。タスクと完了履歴はHiveにローカル保存し、完全にオフラインで動作する。タスク定義と完了履歴は別テーブルに保存されるため、繰り返しルールを変更しても過去の記録は失われない。

![今日のタスク一覧](/assets/images/posts/introducing-donelog/01_today.png)

## 繰り返しルールエンジン

アプリで最も技術的に興味深い部分は、`RecurrenceRule`ドメインエンティティとその`shouldShowToday()`メソッドだ。このpure-Dart関数は、タスクが今日のリストに表示されるべきかどうかを判断する。フレームワーク依存なし、副作用なし。現在の日時とタスクが最後に完了した日時を受け取り、boolを返す。

このロジックをインフラ層から独立したドメイン層の純粋な関数として実装することで、インフラのモックなしに単体テストができる。`daily`・`everyNDays`・`weekdays`・`once`の4種類の繰り返しタイプをすべて1つの`switch`式で処理しており、共有の可変状態は持たない。

リセット処理はアプリ起動時に`CheckAndResetTasks`ユースケースが実行する。全タスクを走査し、`shouldShowToday()`を呼び出し、ルールがリセットすべきと判断した場合にHiveの完了状態を更新する。

## 主要機能

### 今日ビュー

今日の日付と繰り返しルールにマッチするタスクが「今日」欄にリストアップされる。完了済みタスクは取り消し線とチェックマークで表示され、その日の進捗をひと目で確認できる。

![完了したタスク](/assets/images/posts/introducing-donelog/02_task_completed.png)

### タスク登録フォーム

タスクフォームでは4種類の繰り返しタイプを選択できる。`everyNDays`では2〜30日の間隔をスライダーで指定できる。`weekdays`では複数の曜日チェックボックスで任意のパターン（例：月・水・金のみ）を設定できる。

![タスクフォーム](/assets/images/posts/introducing-donelog/03_task_form.png)

### カレンダービュー

過去の完了記録を月間カレンダーで可視化する。任意の日付をタップすると、その日に完了したタスクの一覧を確認できる。習慣の継続状況を振り返ったり、タスクをこなせなかった日を把握したりするのに便利だ。

![カレンダービュー](/assets/images/posts/introducing-donelog/04_calendar.png)

### 設定とダークモード

言語（日本語・英語）とテーマ（ライト・ダーク・システム連動）を設定画面で変更できる。通知リマインダーを任意の時間帯にスケジュールできる。

![設定画面](/assets/images/posts/introducing-donelog/05_settings.png)
![ダークモード](/assets/images/posts/introducing-donelog/06_dark_mode.png)

## まとめ

Done LogはApp Storeで公開している。

- **App Store**: [Done Log](https://apps.apple.com/jp/app/done-log/id6759606196)
