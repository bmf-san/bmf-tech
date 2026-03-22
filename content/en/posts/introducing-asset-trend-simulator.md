---
title: Introducing Asset Trend Simulator — An iPhone App for Compound Interest Simulation
description: 'A deep dive into Asset Trend Simulator, an iPhone app built with Flutter and Riverpod that simulates future net worth trends using compound interest calculations based on your household finances.'
slug: introducing-asset-trend-simulator
date: 2026-03-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Flutter
  - iOS
  - Riverpod
  - Dart
translation_key: introducing-asset-trend-simulator
---

# Introducing Asset Trend Simulator — An iPhone App for Compound Interest Simulation

## Why I Built It

I wanted to answer a simple question: "How much will I have saved by retirement?" But the moment I tried to work it out, I realised how tangled the variables are — income, expenses, investment returns, loan repayments, inflation, salary growth. Doing the maths manually is tedious, and existing apps either oversimplify or require connecting bank accounts.

The actual motivation was different. I had been collecting financial data in a spreadsheet for years — some entered manually, some semi-automated — to forecast my future assets. The accuracy was fine, but the upkeep was tedious. I wanted something simpler that did the same thing. Projecting future assets connects to thinking about future spending; a life plan becomes concrete numbers.

I built Asset Trend Simulator to solve exactly that. Enter your household finances once, tweak a few macro parameters with a slider, and instantly see how your net worth evolves month by month over the coming decades. No account linking, no server, no subscription.

## Architecture

The app is organised in four directories, wired together by Riverpod providers:

- **data** — Immutable Freezed models (`MonthlyIncome`, `MonthlyExpense`, `MonthlyInvestment`, `Loan`, `SimulationParameters`, `SavedCase`, `AppSettings`, etc.) and Hive persistence repositories
- **domain** — Pure Dart calculation logic independent of Flutter and storage (`SimulatorEngine`, `RateConverter`, `AnnuityCalculator`)
- **presentation** — Flutter widgets consuming Riverpod providers (`AppSettingsNotifier` and `SavedCaseList` as `AsyncNotifier`; household and parameter providers as sync `Notifier`)
- **core** — Theme, constants, and shared utilities

Riverpod serves dual purposes: dependency injection and reactive state management. Overriding providers via `ProviderScope` makes swapping implementations straightforward.

![Home screen](/assets/images/posts/introducing-asset-trend-simulator/01_home.png)

## Compound Interest Simulation Logic

The core of the app is the monthly simulation engine. At each time step $t$, the engine computes:

**Income growth and expense inflation (applied independently)**

Income grows each month at the monthly equivalent of `defaultIncomeGrowthRate` (annual rate). Expenses grow at the monthly equivalent of `defaultInflationRate` (annual rate). The two rates are applied independently.

**Investment compounding**

$$I_{t} = I_{t-1} \times (1 + r_{\text{investment}} / 12) + \text{monthly\_investment}_{t}$$

When dividend reinvestment is on, the simulation folds dividends back into $I_t$ rather than routing them into cash.

**Loan repayment (equal-installment method)**

For each loan, the monthly repayment $P$ follows:

$$P = \frac{L_0 \cdot r_m}{1 - (1 + r_m)^{-n}}$$

where $L_0$ is the principal, $r_m$ is the monthly interest rate, and $n$ is the remaining term in months. Each loan disappears from the simulation when its term ends.

**Net worth at month $t$**

$$\text{NW}_t = C_t + I_t - \sum \text{LoanBalance}_t$$

where $C_t$ accumulates cash (income − expenses − loan payments) and $I_t$ is the investment asset value. The simulation runs from month 1 through `periodYears × 12`, producing a data series consumed directly by fl_chart.

![Simulation parameters](/assets/images/posts/introducing-asset-trend-simulator/04_params.png)

## Key Features

### Cash Flow and Asset Input

The app manages income, expenses, and investment items across three tabs. It splits assets into cash holdings and investment holdings; loans expose principal, interest rate, and remaining term for full equal-installment modelling.

![Cash flow input](/assets/images/posts/introducing-asset-trend-simulator/02_cash_flow.png)
![Asset input](/assets/images/posts/introducing-asset-trend-simulator/03_assets.png)

### Asset Trend Graph

Four data series — net worth, cash, investments, and liabilities — appear in a single fl_chart `LineChart`. A toggle switches between annual and monthly granularity.

![Result chart](/assets/images/posts/introducing-asset-trend-simulator/06_result_chart.png)
![Result table](/assets/images/posts/introducing-asset-trend-simulator/07_result_table.png)

### Scenario Save and Compare

Save any combination of inputs and parameters under a custom name as a `SavedCase` in Hive. Tapping a saved case instantly restores all inputs and re-runs the simulation. This makes it easy to compare “aggressive investment” vs “conservative” scenarios side by side.

![Saved cases](/assets/images/posts/introducing-asset-trend-simulator/05_saved_cases.png)

### Multi-Currency and Dark Mode

Switch the currency between JPY, USD, and EUR in settings; the app reformats all displayed values for the selected locale. The theme follows the system setting automatically.

![Settings](/assets/images/posts/introducing-asset-trend-simulator/08_settings.png)

## Summary

Asset Trend Simulator is available on the App Store.

- **App Store**: [Asset Trend Simulator](https://apps.apple.com/jp/app/%E8%B3%87%E7%94%A3%E6%8E%A8%E7%A7%BB%E3%82%B7%E3%83%A5%E3%83%9F%E3%83%AC%E3%83%BC%E3%82%BF%E3%83%BC/id6759601487)
