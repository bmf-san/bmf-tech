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

「老後にいくら残るか」という問いに、感覚ではなく数字で答えたかった。実際に試算しようとすると、収入・支出・投資・ローン・インフレ・賃金上昇率が複雑に絡み合い、手計算は現実的でない。既存アプリは単純すぎるか、銀行口座の連携を求めてくる。

動機はちょっと違う。もともとスプレッドシートで資産に関する情報をあれこれ集めて一部は手動で入力、半自動で将来の資産予測ができるものを作っていた。精度には満足していたが、運用は面倒だった。もっとシンプルなものが欲しかった。将来の資産を予測することは未来の消費を考えることにもつながる。ライフプランを数字で意識できる。

Asset Trend Simulatorはその課題を解決するために作った。家計情報を一度入力し、マクロパラメータをスライダーで調整するだけで、将来の純資産推移を月次で即座に確認できる。口座連携なし、サーバーなし、サブスクなし。

## アーキテクチャ

アプリは次の4ディレクトリ構成をRiverpodプロバイダで接続している。

- **data** — Freezedによるイミュータブルなモデル（`MonthlyIncome`・`MonthlyExpense`・`MonthlyInvestment`・`Loan`・`SimulationParameters`・`SavedCase`・`AppSettings`など）とHive永続化リポジトリ
- **domain** — Flutter・ストレージに依存しない純粋な計算ロジック（`SimulatorEngine`・`RateConverter`・`AnnuityCalculator`）
- **presentation** — Riverpodプロバイダを参照するFlutterウィジェット（`AppSettingsNotifier`・`SavedCaseList`は`AsyncNotifier`、家計・パラメータ系は同期`Notifier`）
- **core** — テーマ・定数・共有ユーティリティ

RiverpodはDIと状態管理の2役を担う。`ProviderScope` のオーバーライドでプロバイダ実装を差し替えられるため、テストでの切り替えも容易だ。

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
