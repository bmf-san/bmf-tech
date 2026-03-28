---
title: iPhoneアプリ「Asset Trend Simulator」の紹介
description: 'Flutter と Riverpod で構築した複利計算型の資産推移シミュレーター iPhone アプリ「Asset Trend Simulator」の設計と実装を解説する。'
slug: introducing-asset-trend-simulator
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Flutter
  - iOS
  - Riverpod
  - Dart
translation_key: introducing-asset-trend-simulator
---

# iPhoneアプリ「Asset Trend Simulator」の紹介

## なぜ作ったか

ライフプランを考える上で、「年間どれくらい消費に回せるか」、「将来のある時点でどれくらい資産があるか」、「老後にいくら残るか」といった問いに手軽に答えられるツールが欲しかった。

スプレッドシートで資産に関する情報をあれこれ集めて一部は手動で入力、半自動で将来の資産予測ができるものを作っていた。精度には満足していたが、運用は面倒だった。ライフプランナーにライフプランニング表を何度か作ってもらったこともあるが、細かい設定が面倒で、気軽にシミュレーションしてみるにはハードルが高かった。

そこで、適当な入力情報とシュミレーションのためのパラメータ情報を設定するだけで、将来の資産推移を気軽に確認できるAsset Trend Simulatorを作ることにした。

計算の精度が完全ではない部分もあるが、概ね自分がスプレッドシートで運用していたものと同等の精度は出ており、自分で使う分にはある程度満足している。

似たような気持ちで気軽に資産推移をシミュレーションしてみたい人がいればと思い、[App Store](https://apps.apple.com/jp/app/%E8%B3%87%E7%94%A3%E6%8E%A8%E7%A7%BB%E3%82%B7%E3%83%A5%E3%83%9F%E3%83%AC%E3%83%BC%E3%82%BF%E3%83%BC/id6759601487)で公開しているので、興味があればぜひ試してみてほしい。

## ユースケース

- **老後資産の概算** — 現在の収支・資産・ローンを入力し、65歳時点の純資産を確認する
- **感度分析** — 運用利率を3%・5%・7%と変えたときの純資産推移の差を比べる
- **シナリオ比較** — 積極投資ケースと保守的ケースの2パターンを保存し、切り替えて見比べる
- **ローンの影響確認** — 住宅ローンや車ローンを加えたときに純資産のピークがどう変わるかを確認する

![ホーム画面](/assets/images/posts/introducing-asset-trend-simulator/01_home.png)

## 主な機能

### 収支・資産入力

収入・支出・投資の3カテゴリで項目を管理する。収入と支出はそれぞれ複数の項目を月次金額で登録できる。投資は毎月の積立額を登録する。ローンは元本・金利・残期間を入力すると元利均等方式で月返済額を自動計算する。現預金残高と運用資産残高は初期値として設定する。

![収支入力](/assets/images/posts/introducing-asset-trend-simulator/02_cash_flow.png)
![資産入力](/assets/images/posts/introducing-asset-trend-simulator/03_assets.png)

### シミュレーションパラメータ

インフレ率・収入成長率・運用利率・配当再投資の有無・シミュレーション期間（年）をスライダーで調整できる。

![シミュレーション設定](/assets/images/posts/introducing-asset-trend-simulator/04_params.png)

### 資産推移グラフ

純資産・現預金・運用資産・負債の4系列をfl_chartの `LineChart` に描画する。年次・月次の粒度切り替えに対応している。

![結果グラフ](/assets/images/posts/introducing-asset-trend-simulator/06_result_chart.png)
![結果テーブル](/assets/images/posts/introducing-asset-trend-simulator/07_result_table.png)

### シナリオ保存・比較

任意の名前で入力状態を `SavedCase` としてHiveに保存できる。保存済みケースをタップすると、すべての入力が復元されてシミュレーションが再実行される。積極投資ケースと保守的ケースを並べて比較するといった使い方が可能だ。

![保存済みケース](/assets/images/posts/introducing-asset-trend-simulator/05_saved_cases.png)

### 多通貨・ダークモード対応

JPY・USD・EURの3通貨を切り替えられ、表示値はロケールに従ってフォーマットされる。テーマはシステム設定に自動的に連動する。

![設定](/assets/images/posts/introducing-asset-trend-simulator/08_settings.png)

## 技術スタック

| 層 | 技術 |
|---|---|
| UI | Flutter (iOS) |
| 状態管理 / DI | Riverpod + riverpod_generator |
| モデル | Freezed + json_serializable |
| 永続化 | Hive |
| グラフ | fl_chart |
| ルーティング | go_router |
| テスト | flutter_test / mocktail |

アーキテクチャはdata / domain / presentation / coreの4層構成だ。domain層はFlutter・ストレージに依存しない純粋な計算ロジックを包み、RiverpodがDIと状態管理の2役を担う。

## シミュレーションエンジンの概要

シミュレーションは1ヶ月ごとに純資産を計算し、設定した期間分だけ繰り返す仕組みだ。各月の計算は大きく4ステップで成り立っている。

1. **収入・支出の更新** — 収入は設定した年内成長率分だけ毎月増え、支出はインフレ率分だけ毎月増える。中長期で購買力が変わる影響を考慮するためだ。

2. **投資資産の複利運用** — 前月の運用資産に運用利率分のリターンを加え、そこに毎月の積立額を足す。配当再投資が有効な場合は配当が現預金されず、そのまま運用資産に複利で乗される。

3. **ローン返済** — 各ローンの月返済額は元利均等方式で自動計算され、返済完了月まで収入から差し引かれる。

4. **月末純資産の確定** — その月の「現預金（収入−支出−ローン返済額）＋運用資産−ローン残高」が純資産になる。

## まとめ

Asset Trend SimulatorはApp Storeで公開中である。

- **App Store**: [Asset Trend Simulator](https://apps.apple.com/jp/app/%E8%B3%87%E7%94%A3%E6%8E%A8%E7%A7%BB%E3%82%B7%E3%83%A5%E3%83%9F%E3%83%AC%E3%83%BC%E3%82%BF%E3%83%BC/id6759601487)
