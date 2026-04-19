---
title: Developed an Asset Trend Simulator
slug: asset-transition-simulator
date: 2025-05-25T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - JavaScript
description: Introducing a tool for personal asset formation, the Asset Trend Simulator.
translation_key: asset-transition-simulator
---



# Overview
As a tool for considering personal asset formation, I developed an asset trend simulator, which I would like to introduce.

[asset-trend-simulator](http://web.archive.org/web/20250622150217/https://asset-trend-simulator.netlify.app/)

# Background
I use Money Forward for asset management, but for future asset trend simulation, I had been using a spreadsheet.

While the spreadsheet allowed for a reasonable degree of accuracy (with discrepancies in actual asset trends being limited to several hundred thousand yen), updating it once a month was cumbersome. I was looking for a simpler solution, which led to this development.

Enhancing the accuracy of the simulation requires considerable calculations and variable settings, but I believe that an approximation is still useful for personal asset formation.

Ideally, the actual asset situation should exceed the simulation results, as a downward trend could affect life plans.

# About the Asset Trend Simulator
## Main Features
### 1. Simple Simulation Functionality
- Simulate asset trends for up to 50 years
- Supports multiple currencies, considering exchange rate fluctuations
- Reflects comprehensive asset situations, including income, expenses, investments, and loans

It is designed to provide rough estimates for life planning.

### 2. Intuitive User Interface
- Responsive design for mobile/desktop compatibility
- Visual results displayed through interactive graphs
- Easy data entry with accordion-style input forms
- Detailed data tables for numerical verification

The calculation results can be checked annually or monthly, making it easy to understand asset trend tendencies.

### 3. Calculation Engine
- Simulates investment returns through compound interest calculations
- Loan calculations using the equal principal and interest repayment method
- Adapts to exchange rate fluctuations
- Income predictions considering real wage growth rates
- Calculates asset trends considering the impact of liabilities
- Simulates the growth of managed assets

For managed assets, it would be ideal to reflect actual stock prices for accuracy, but predicting stock price trends is difficult. Therefore, it allows setting the growth rate of assets under management.

Since the rate of return on investments significantly impacts assets, setting this rate strictly allows for a conservative simulation.

## Technologies Used
- Frontend: HTML5, CSS3, JavaScript (ES6+)
- Graph Rendering: Chart.js
- Build/Development: npm, webpack
- Testing: Jest

The application is deployed on Netlify with ease.

## Details of Key Features
### Asset Status Input
- Set income and expenses (supports multiple currencies)
- Input cash and deposit balances
- Register investment assets
- Set liability information such as housing loans

### Simulation Parameters
- Set simulation period (up to 50 years)
- Set real wage growth rate
- Set investment returns
- Predict exchange rate fluctuations

### Result Display
- Net asset trend graph
- Asset allocation pie chart
- Summary cards of key indicators
- Detailed monthly data table

## How to Use
1. Enter Basic Information
   - Current income and expense status
   - Status of owned assets
   - Details of liabilities

2. Set Simulation Conditions
   - Select period
   - Adjust various parameters

3. Check and Analyze Results
   - Check trends with graphs and tables
   - Adjust parameters as needed
   - Simulate with different scenarios

# Future Prospects
There is room for improvement in simulation accuracy, so I would like to adjust it while maintaining simplicity.

Currently, it is completed with only the frontend, but if the number of users increases, it might be interesting to build a backend and develop it into a proper service. (Currently, there is no movement)

By visualizing the current asset situation and being aware of future asset trends, I believe it will be easier to plan asset formation, so please give it a try.

