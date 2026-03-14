---
title: What is a Critical User Journey
description: 'Understand Critical User Journeys to identify essential user touchpoints, goals, and pain points for SLO-based service reliability.'
slug: critical-user-journey
date: 2024-08-15T00:00:00Z
author: bmf-san
categories:
  - Operations
tags:
  - Critical User Journey
  - SLO
translation_key: critical-user-journey
---

# Overview
I became interested in the Critical User Journey in the context of SLOs.

Here, I summarize what I found out about Critical User Journeys.

# What is a Critical User Journey
It represents the most important experiences for users of a service and the main paths to achieve specific goals.

It allows us to identify important touchpoints and obstacles in the process of users achieving their goals.

It clarifies the goals users should achieve and identifies tasks and subtasks.

# Examples of Critical User Journeys
## Example 1: Purchase Process on an E-commerce Site

### User Goal:
The user wants to purchase a specific product online.

### Touchpoints:

#### Product Search and Discovery
- The user searches for products using the search bar or browses categories.
- Related products are displayed by the recommendation engine.

#### Product Detail Confirmation
- The user accesses the product detail page to check the price, reviews, and specifications.
- The user views product images and videos.

#### Add to Cart
- The user adds the desired product to the cart.
- The user checks the list of products intended for purchase on the cart page.

#### Checkout
- The user proceeds to the payment page and enters shipping information and payment method.
- The user applies a coupon code for discounts.

#### Order Confirmation
- The user performs a final check and confirms the order.
- The user receives an order confirmation email.

#### Product Receipt
- The user receives the product and checks for any issues.
- The user may post a review or contact customer support if necessary.

### Pain Points:
- If product search is not smooth or search results do not meet expectations.
- Limited payment methods or payment failures during checkout.
- Delays in delivery or defects in the product.

## Example 2: Registration for a Subscription Service

### User Goal:
The user wants to register for a music streaming service, create playlists, and listen to music.

### Touchpoints:

#### Website Visit and Registration
- The user accesses the service's website and checks the service details.
- The user creates an account and selects a subscription plan.

#### App Download and Installation
- The user downloads and installs the mobile app.

#### Login and Setup
- The user logs into the account and sets up the profile.
- The user selects preferred music genres and artists.

#### Music Search and Playback
- The user searches for favorite music and starts streaming playback.
- The user creates custom playlists and adds favorite songs.

#### Offline Playback Setup
- The user downloads music and sets it up for offline playback.

### Pain Points:
- The registration process is complicated with too much information to input.
- The app interface is difficult to use, making it hard to find desired music.
- The download feature is unstable, causing issues with offline playback.

## Example 3: Money Transfer Using a Banking App

### User Goal:
The user wants to use the banking app to transfer money to a specified account.

### Touchpoints:

#### App Launch and Login
- The user launches the banking app and logs in using a security code or fingerprint authentication.

#### Selecting Transfer Destination
- The user selects a transfer destination or inputs new transfer destination information.

#### Entering Transfer Amount
- The user inputs the transfer amount and adds a purpose or memo for the transfer.

#### Confirmation and Approval
- The user reviews the transfer details and approves the transfer (using a one-time password, etc.).
- The user receives a notification of transfer completion.

### Pain Points:
- The login process is complicated, and authentication may fail.
- If the transfer destination information is entered incorrectly, the error message may be unclear.
- Transfers may not reflect in real-time.

# Critical User Journey and SLO
The Critical User Journey is related to the definition of reliability.

Achieving SLOs is thought to support the Critical User Journey.

By defining the Critical User Journey and then defining SLOs based on it, we can enhance the degree to which SLOs contribute to reliability.

## Example 1: Purchase Process on an E-commerce Site

**CUJ: Add product to cart and checkout**

**SLO Settings:**

- **Product Search and Discovery**
  - **Page Load Time**: The load time for the product search results page is within 2 seconds.
  - **Recommendation System Accuracy**: Related product suggestions display appropriate products with 80% accuracy.

- **Product Detail Confirmation**
  - **Accuracy of Product Details**: Information on the product page (price, reviews, specifications) is 100% accurate.
  - **Image and Video Load Time**: Product images and videos are displayed within 1 second.

- **Add to Cart**
  - **Cart Update Time**: The action of adding a product to the cart is reflected within 2 seconds.
  - **Cart Information Consistency**: The rate at which products added to the cart are displayed accurately is 99.9% or higher.

- **Checkout**
  - **Payment Processing Time**: The processing of the payment page is completed within 3 seconds.
  - **Payment Error Rate**: The error rate for payment processing is below 0.1%.
  - **Coupon Code Application**: The application of coupon codes is reflected in real-time.

- **Order Confirmation**
  - **Order Confirmation Email Delivery**: The time from order confirmation to the delivery of the confirmation email is within 5 minutes.

- **Product Receipt**
  - **Delivery Time**: The average time from shipment to delivery completion is within 3 business days.
  - **Delivery Error Rate**: The rate of delivery errors (delays, misdeliveries, etc.) is below 0.5%.

**Reliability Perspective:**

These settings are established to ensure the reliability necessary for users to smoothly complete the purchase process. In particular, error rates and response times are important indicators that directly demonstrate reliability.

---

## Example 2: Registration for a Subscription Service

**CUJ: Register for a music streaming service, create playlists, and listen to music**

**SLO Settings:**

- **Website Visit and Registration**
  - **Website Uptime**: 99.9% uptime.
  - **Registration Form Response Time**: Form submissions are processed within 2 seconds.

- **App Download and Installation**
  - **Download Success Rate**: The success rate for downloads and installations is 99.9% or higher.
  - **Installation Time**: The app installation is completed within 1 minute.

- **Login and Setup**
  - **Login Success Rate**: The success rate for login attempts is 99.9% or higher.
  - **Profile Setting Reflection Time**: Profile settings are reflected in real-time.

- **Music Search and Playback**
  - **Search Result Response Time**: Search results are displayed within 1 second.
  - **Playback Start Time**: Music playback starts within 3 seconds.
  - **Streaming Interruption Rate**: The music interruption rate is below 0.05%.

- **Offline Playback Setup**
  - **Download Success Rate**: The success rate for music downloads is 99.9% or higher.
  - **Offline Playback Error Rate**: The error rate during offline playback is below 0.1%.

**Reliability Perspective:**

This ensures that users can smoothly proceed from registration to music playback. In particular, the success rates for installation and login, as well as the speed of playback start, are key indicators of reliability.

---

## Example 3: Money Transfer Using a Banking App

**CUJ: Use the banking app to transfer money to a specified account**

**SLO Settings:**

- **App Launch and Login**
  - **Login Success Rate**: The success rate for login attempts is 99.9% or higher.
  - **Login Processing Time**: The login process is completed within 3 seconds.

- **Selecting Transfer Destination**
  - **Transfer Destination Selection Response Time**: Selecting the transfer destination is completed within 2 seconds.

- **Entering Transfer Amount**
  - **Amount Input Consistency**: The rate at which the entered amount is processed accurately is 100%.

- **Confirmation and Approval**
  - **Confirmation Process Response Time**: The confirmation of transfer details is completed within 2 seconds.
  - **Approval Process Error Rate**: The error rate during transfer approval is below 0.05%.
  - **One-Time Password Reception Time**: OTP is received within 30 seconds.

- **Transfer Completion Notification**
  - **Notification Delivery Time**: Notifications are delivered within 5 minutes after transfer completion.

**Reliability Perspective:**

SLOs guarantee the accuracy and speed of transfer processing, allowing users to complete the transfer process without issues. The success rates for login and transfer processing, as well as the timely delivery of notifications, are key indicators of reliability.

# Thoughts
I understood the importance of defining the Critical User Journey when considering SLOs.

I felt that organizing based on user experience, rather than user flow (the expected behavior flow for users, as per specifications), clarifies the definition of reliability.

Having clear answers to questions like "Why was this SLO defined?" and "Why is this SLO important?" makes it easier to involve non-developers in the design and operation of SLOs.

# References
- [sre.google - Modeling User Journeys](https://sre.google/workbook/implementing-slos/#modeling-user-journeys)
- [cloud.google.com - Standardizing the SLO Design Process](https://cloud.google.com/blog/ja/products/devops-sre/how-to-design-good-slos-according-to-google-sres)
- [popinsight.jp - Comprehensive Explanation of Critical User Journey by a Current Google UX Manager](https://popinsight.jp/blog/?p=40141)
- [speakerdeck.com - Improving SLI/SLO Using Critical User Journey](https://speakerdeck.com/heleeen/slo-critical-user-journey)
- [u-site.jp - Comparison of User Journeys and User Flows](https://u-site.jp/alertbox/user-journeys-vs-user-flows)