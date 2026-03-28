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

When thinking about life planning, I wanted a tool that could answer questions like "How much can I spend each year?", "How much will I have at a given point in the future?", and "How much will I have left in retirement?" without much friction.

I had been collecting financial data in a spreadsheet — some entered manually, some semi-automated — to forecast my future assets. The accuracy was fine, but the upkeep was tedious. I also had a financial planner draw up life-planning tables a few times, but the detailed setup was cumbersome and the bar was too high for casual what-if simulations.

So I built Asset Trend Simulator: enter your household finances and a few simulation parameters, and instantly see how your net worth evolves. The calculation is not perfectly precise, but roughly on par with what I was doing in my spreadsheet, and good enough for personal use.

If you want to casually simulate your own asset trends, the app is on the App Store — give it a try.

## Use Cases

- **Retirement estimate** — Enter current income, expenses, assets, and loans; check projected net worth at age 65.
- **Sensitivity analysis** — Compare net worth trajectories at 3%, 5%, and 7% annual investment returns.
- **Scenario comparison** — Save an "aggressive" and a "conservative" case, then switch between them on the chart.
- **Loan impact** — See how adding a mortgage or car loan shifts the net worth peak.

## Key Features

![Home screen](/assets/images/posts/introducing-asset-trend-simulator/01_home.png)

### Cash Flow and Asset Input

Income, expenses, and investments each support many line items registered as monthly amounts. Loans take principal, interest rate, and remaining term; the app computes the equal-installment monthly payment automatically. Initial cash balance and investment balance set the starting values.

![Cash flow input](/assets/images/posts/introducing-asset-trend-simulator/02_cash_flow.png)
![Asset input](/assets/images/posts/introducing-asset-trend-simulator/03_assets.png)

### Simulation Parameters

Inflation rate, income growth rate, investment return rate, dividend reinvestment toggle, and simulation period in years are adjustable with sliders.

![Simulation parameters](/assets/images/posts/introducing-asset-trend-simulator/04_params.png)

### Asset Trend Graph

Four data series — net worth, cash, investments, and liabilities — appear in a single fl_chart `LineChart`. A toggle switches between annual and monthly granularity.

![Result chart](/assets/images/posts/introducing-asset-trend-simulator/06_result_chart.png)
![Result table](/assets/images/posts/introducing-asset-trend-simulator/07_result_table.png)

### Scenario Save and Compare

Save any combination of inputs and parameters under a custom name as a `SavedCase` in Hive. Tapping a saved case instantly restores all inputs and re-runs the simulation. This makes it easy to compare an aggressive investment plan against a conservative one.

![Saved cases](/assets/images/posts/introducing-asset-trend-simulator/05_saved_cases.png)

### Multi-Currency and Dark Mode

Switch the currency between JPY, USD, and EUR in settings; the app reformats all displayed values for the selected locale. The theme follows the system setting automatically.

![Settings](/assets/images/posts/introducing-asset-trend-simulator/08_settings.png)

## Tech Stack

| Layer | Technology |
|---|---|
| UI | Flutter (iOS) |
| State management / DI | Riverpod + riverpod_generator |
| Models | Freezed + json_serializable |
| Persistence | Hive |
| Charts | fl_chart |
| Routing | go_router |
| Testing | flutter_test / mocktail |

The app follows a four-layer layout: data, domain, presentation, and core. The domain layer holds pure calculation logic with no dependency on Flutter or storage. Riverpod handles both dependency injection and reactive state.

## How the Simulation Engine Works

The simulation steps through one month at a time, repeating for the configured period. Each month involves four steps.

1. **Update income and expenses** — Income rises each month by a monthly slice of the annual growth rate; expenses rise by a monthly slice of the inflation rate. This captures the long-term shift in purchasing power.

2. **Compound the investment portfolio** — The previous month's investment balance earns the monthly fraction of the annual return, then the month's contribution gets added. With dividend reinvestment on, dividends fold back into the portfolio rather than flowing into cash.

3. **Deduct loan repayments** — Each loan's monthly payment follows the equal-installment formula and gets subtracted from income until the loan term ends.

4. **Compute end-of-month net worth** — Net worth for the month is cash accumulated (income − expenses − loan payments) plus investment balance minus outstanding loan balances.
## Summary

Asset Trend Simulator is available on the App Store.

- **App Store**: [Asset Trend Simulator](https://apps.apple.com/jp/app/%E8%B3%87%E7%94%A3%E6%8E%A8%E7%A7%BB%E3%82%B7%E3%83%A5%E3%83%9F%E3%83%AC%E3%83%BC%E3%82%BF%E3%83%BC/id6759601487)
