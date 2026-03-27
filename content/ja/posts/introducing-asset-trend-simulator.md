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

似たような気持ちで気軽に資産推移をシミュレーションしてみたい人がいればと思い、App Storeで公開しているので、興味があればぜひ試してみてほしい。

## 仕様とユースケース

アプリに入力する情報は大きく2種類だ。

**家計情報**は収入・支出・投資の3カテゴリで管理する。収入と支出はそれぞれ複数の項目を月次金額で登録できる。投資は毎月の積立額を登録する。ローンは元本・金利・残期間を入力すると元利均等方式で月返済額を自動計算する。現預金残高と運用資産残高は初期値として設定する。

**シミュレーションパラメータ**はインフレ率・収入成長率・運用利率・配当再投資の有無・シミュレーション期間（年）をスライダーで調整できる。

想定しているユースケースはこんな感じだ。

- **老後資産の概算** — 現在の収支・資産・ローンを入力し、65歳時点の純資産を確認する
- **感度分析** — 運用利率を3%・5%・7%と変えたときの純資産推移の差を比べる
- **シナリオ比較** — 積極投資ケースと保守的ケースの2パターンを保存し、切り替えて見比べる
- **ローンの影響確認** — 住宅ローンや車ローンを加えたときに純資産のピークがどう変わるかを確認する

![ホーム画面](/assets/images/posts/introducing-asset-trend-simulator/01_home.png)

## 複利計算ロジックの実装

アプリの核となるのが月次シミュレーションエンジンである。各タイムステップ $t$ において以下を計算する。

**収入成長と支出インフレ（独立適用）**

収入は毎月 `defaultIncomeGrowthRate`（年率）を月次換算した成長率で増加する。支出は `defaultInflationRate`（年率）を月次換算したインフレ率で増加する。2つのレートは独立して適用される。

**投資資産の複利計算**

$$I_{t} = I_{t-1} \times (1 + r_{\text{investment}} / 12) + \text{monthly\_investment}_{t}$$

配当再投資が有効な場合、配当は現預金に流入せず$I_t$へ折り返される。

**ローン返済（元利均等方式）**

各ローンの月返済額 $P$ は以下の式で算出される。

$$P = \frac{L_0 \cdot r_m}{1 - (1 + r_m)^{-n}}$$

$L_0$ は元本、$r_m$ は月利、$n$ は残期間（月数）。返済完了月以降はシミュレーションから除外される。

**月末純資産**

$$\text{NW}_t = C_t + I_t - \sum \text{LoanBalance}_t$$

$C_t$ は積み上げキャッシュ（収入 − 支出 − ローン返済）、$I_t$ は運用資産評価額。シミュレーションは第1月から `periodYears × 12` 月まで走り、結果データ系列をfl_chartに直接渡す。

![シミュレーション設定](/assets/images/posts/introducing-asset-trend-simulator/04_params.png)

## 主な機能

### 収支・資産入力

収入・支出・投資の3タブで項目を管理する。資産は現預金と運用資産に分類し、ローンは元本・金利・残期間を入力して元利均等計算に対応する。

![収支入力](/assets/images/posts/introducing-asset-trend-simulator/02_cash_flow.png)
![資産入力](/assets/images/posts/introducing-asset-trend-simulator/03_assets.png)

### 資産推移グラフ

純資産・現預金・運用資産・負債の4系列をfl_chartの `LineChart` に描画する。年次・月次の粒度切り替えに対応している。

![結果グラフ](/assets/images/posts/introducing-asset-trend-simulator/06_result_chart.png)
![結果テーブル](/assets/images/posts/introducing-asset-trend-simulator/07_result_table.png)

### シナリオ保存・比較

任意の名前で入力状態を `SavedCase` としてHiveに保存できる。保存済みケースをタップすると、すべての入力が復元されてシミュレーションが再実行される。「積極投資」と「保守的」の2ケースを並べて比較するといった使い方が可能だ。

![保存済みケース](/assets/images/posts/introducing-asset-trend-simulator/05_saved_cases.png)

### 多通貨・ダークモード対応

JPY・USD・EURの3通貨を切り替えられ、表示値はロケールに従ってフォーマットされる。テーマはシステム設定に自動的に連動する。

![設定](/assets/images/posts/introducing-asset-trend-simulator/08_settings.png)

## まとめ

Asset Trend SimulatorはApp Storeで公開中である。

- **App Store**: [Asset Trend Simulator](https://apps.apple.com/jp/app/%E8%B3%87%E7%94%A3%E6%8E%A8%E7%A7%BB%E3%82%B7%E3%83%A5%E3%83%9F%E3%83%AC%E3%83%BC%E3%82%BF%E3%83%BC/id6759601487)
