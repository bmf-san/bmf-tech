---
title: Created an Asset Trend Simulator
slug: asset-transition-simulator
date: 2025-05-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - JavaScript
translation_key: asset-transition-simulator
---

# Overview
I developed an asset trend simulator as a tool for personal asset formation, and I would like to introduce it.

[asset-trend-simulator](https://asset-trend-simulator.netlify.app/)

# Background
I use Money Forward for asset management, but I had been using a Spreadsheet for simulating future asset trends.

While the Spreadsheet provided a reasonable level of accuracy (with discrepancies in actual asset trends being around several hundred thousand), updating it monthly was cumbersome. I was looking for a simpler solution, which led to this development.

To enhance the accuracy of the simulation, certain calculations and variable settings are necessary, but I believe that even a rough estimate can be useful for personal asset formation.

Ideally, the actual asset situation should exceed the simulation results (as falling short could impact life plans).

# About the Asset Trend Simulator
## Main Features
### 1. Simple Simulation Functionality
- Simulates asset trends for up to 50 years
- Supports multiple currencies and considers exchange rate fluctuations
- Reflects comprehensive asset situations including income, expenses, investments, and loans

It is designed to provide a rough estimate for life planning.

### 2. Intuitive User Interface
- Responsive design for mobile/desktop compatibility
- Visual results displayed through interactive graphs
- Easy data entry with accordion-style input forms
- Detailed data tables for numerical verification

Users can check calculation results annually and monthly, making it easier to grasp asset trend patterns.

### 3. Calculation Engine
- Simulates investment returns using compound interest calculations
- Loan calculations using equal principal and interest repayment methods
- Accounts for exchange rate fluctuations
- Income forecasts considering real wage growth rates
- Calculates asset trends considering the impact of liabilities
- Simulates growth of managed assets

For managed assets, it would be ideal to reflect actual stock prices for accuracy, but predicting stock price trends is challenging. Therefore, users can set the growth rate of managed assets.

Since the return on investment has a significant impact on assets, tightening this return setting allows for a more conservative simulation.

## Technologies Used
- Frontend: HTML5, CSS3, JavaScript (ES6+)
- Graph Rendering: Chart.js
- Build/Development: npm, webpack
- Testing: Jest

The application is simply deployed on Netlify.

## Details of Main Features
### Asset Situation Input
- Settings for income and expenses (supports multiple currencies)
- Input for cash and savings balances
- Registration of investment assets
- Settings for liability information such as mortgages

### Simulation Parameters
- Setting the simulation period (up to 50 years)
- Setting the real wage growth rate
- Setting investment returns
- Forecasting exchange rate fluctuations

### Result Display
- Net asset trend graph
- Asset allocation pie chart
- Summary cards for key indicators
- Detailed monthly data tables

## How to Use
1. Input Basic Information
   - Current income and expense situation
   - Status of owned assets
   - Details of liabilities

2. Set Simulation Conditions
   - Select the period
   - Adjust various parameters

3. Review and Analyze Results
   - Check trends with graphs and tables
   - Adjust parameters as needed
   - Simulate different scenarios

# Future Prospects
There is room for improvement in the accuracy of the simulation, so I want to adjust the precision while maintaining simplicity.

Currently, it is simplified to the frontend only, but if the number of users increases, it might be interesting to build a backend and develop it into a more substantial service. (For now, it’s quiet)

I believe that visualizing the current asset situation and being aware of future asset trends will make it easier to plan for asset formation, so I hope you will give it a try.