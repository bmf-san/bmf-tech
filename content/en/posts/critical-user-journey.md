---
title: What is a Critical User Journey?
slug: critical-user-journey
date: 2024-08-15T00:00:00Z
author: bmf-san
categories:
  - Operations
tags:
  - Critical User Journey
  - SLO
description: An overview of critical user journeys in the context of SLOs.
translation_key: critical-user-journey
---

# Overview
I wanted to learn more about critical user journeys in the context of SLOs.

This post summarizes what I found about critical user journeys.

# What is a Critical User Journey?
A critical user journey represents the most important experiences or key paths that users take to achieve specific goals within a service.

It helps identify key touchpoints and obstacles in the process of achieving user goals.

It clarifies the goals users need to achieve and identifies tasks and subtasks.

# Examples of Critical User Journeys
## Example 1: Purchase Process on an E-commerce Site

### User Goal:
To purchase a specific product online.

### Touchpoints:

#### Product Search and Discovery
- Users search for products using a search bar or browse categories.
- A recommendation engine displays related products.

#### Product Details Review
- Users access the product detail page to check price, reviews, and specifications.
- Users view product images or videos.

#### Adding to Cart
- Users add desired products to their cart.
- Users review the list of items in their cart on the cart page.

#### Checkout
- Users proceed to the payment page and enter shipping and payment information.
- Users apply discount codes for savings.

#### Order Confirmation
- Users review their order and confirm the purchase.
- Users receive an order confirmation email.

#### Product Delivery
- Users receive the product and check for any issues.
- If necessary, users leave reviews or contact customer support.

### Pain Points:
- Difficulty in finding products or irrelevant search results.
- Limited payment options or failed transactions during checkout.
- Delayed delivery or defective products.

## Example 2: Signing Up for a Subscription Service

### User Goal:
To sign up for a music streaming service, create playlists, and listen to music.

### Touchpoints:

#### Visiting the Website and Signing Up
- Users visit the service's website and review the service details.
- Users create an account and select a subscription plan.

#### Downloading and Installing the App
- Users download and install the mobile app.

#### Logging In and Setting Up
- Users log in to their account and set up their profile.
- Users select their favorite music genres and artists.

#### Searching and Playing Music
- Users search for their favorite music and start streaming.
- Users create custom playlists and add favorite songs.

#### Setting Up Offline Playback
- Users download music for offline playback.

### Pain Points:
- Complicated registration process with too many input fields.
- Difficult-to-use app interface or inability to find desired music.
- Unstable download functionality or issues with offline playback.

## Example 3: Transferring Money via a Banking App

### User Goal:
To transfer money to a specified account using a banking app.

### Touchpoints:

#### Launching the App and Logging In
- Users launch the banking app and log in using a security code or fingerprint authentication.

#### Selecting the Recipient
- Users select an existing recipient or enter new recipient details.

#### Entering the Transfer Amount
- Users input the transfer amount and add a purpose or note.

#### Reviewing and Approving
- Users review the transfer details and approve the transaction (e.g., using a one-time password).
- Users receive a notification confirming the transfer.

### Pain Points:
- Complicated login process or failed authentication.
- Unclear error messages when entering incorrect recipient details.
- Delays in real-time transfer processing.

# Critical User Journeys and SLOs
Critical user journeys are closely related to defining reliability.

Achieving SLOs is considered to support critical journeys.

By defining critical user journeys and using them as a basis for defining SLOs, the contribution of SLOs to reliability can be enhanced.

## Example 1: Purchase Process on an E-commerce Site

**CUJ: Adding products to the cart and completing checkout**

**SLO Settings:**

- **Product Search and Discovery**
  - **Page Load Time**: Product search result pages load within 2 seconds.
  - **Recommendation System Accuracy**: Related product suggestions are 80% accurate.

- **Product Details Review**
  - **Accuracy of Product Details**: Information on product pages (price, reviews, specifications) is 100% accurate.
  - **Image/Video Load Time**: Product images and videos load within 1 second.

- **Adding to Cart**
  - **Cart Update Time**: Adding a product to the cart reflects within 2 seconds.
  - **Cart Information Accuracy**: Items added to the cart are displayed correctly 99.9% of the time.

- **Checkout**
  - **Payment Processing Time**: Payment processing completes within 3 seconds.
  - **Payment Error Rate**: Payment error rate is below 0.1%.
  - **Coupon Code Application**: Coupon codes are applied in real-time.

- **Order Confirmation**
  - **Order Confirmation Email Delivery**: Confirmation emails are sent within 5 minutes of order completion.

- **Product Delivery**
  - **Delivery Time**: Average delivery time from shipment to completion is within 3 business days.
  - **Delivery Error Rate**: Delivery errors (delays, incorrect deliveries, etc.) are below 0.5%.

**Reliability Perspective:**

SLOs are set to ensure the reliability required for users to smoothly complete the purchase process. Metrics such as error rates and response times are critical indicators of reliability.

---

## Example 2: Signing Up for a Subscription Service

**CUJ: Signing up for a music streaming service, creating playlists, and listening to music**

**SLO Settings:**

- **Visiting the Website and Signing Up**
  - **Website Uptime**: 99.9% or higher uptime.
  - **Registration Form Response Time**: Form submission processes within 2 seconds.

- **Downloading and Installing the App**
  - **Download Success Rate**: Download and installation success rate is 99.9% or higher.
  - **Installation Time**: App installation completes within 1 minute.

- **Logging In and Setting Up**
  - **Login Success Rate**: Login attempts succeed 99.9% of the time.
  - **Profile Setup Response Time**: Profile settings reflect in real-time.

- **Searching and Playing Music**
  - **Search Response Time**: Search results display within 1 second.
  - **Playback Start Time**: Music playback starts within 3 seconds.
  - **Streaming Interruption Rate**: Music streaming interruption rate is below 0.05%.

- **Setting Up Offline Playback**
  - **Download Success Rate**: Music download success rate is 99.9% or higher.
  - **Offline Playback Error Rate**: Offline playback error rate is below 0.1%.

**Reliability Perspective:**

SLOs ensure users can complete the process from registration to music playback without issues. Key reliability indicators include installation and login success rates, as well as playback speed.

---

## Example 3: Transferring Money via a Banking App

**CUJ: Using a banking app to transfer money to a specified account**

**SLO Settings:**

- **Launching the App and Logging In**
  - **Login Success Rate**: Login attempts succeed 99.9% of the time.
  - **Login Processing Time**: Login processes complete within 3 seconds.

- **Selecting the Recipient**
  - **Recipient Selection Response Time**: Selecting a recipient completes within 2 seconds.

- **Entering the Transfer Amount**
  - **Amount Input Accuracy**: Entered amounts are processed with 100% accuracy.

- **Reviewing and Approving**
  - **Review Process Response Time**: Reviewing transfer details completes within 2 seconds.
  - **Approval Process Error Rate**: Transfer approval error rate is below 0.05%.
  - **One-Time Password Delivery Time**: OTP is received within 30 seconds.

- **Transfer Completion Notification**
  - **Notification Delivery Time**: Transfer completion notification is sent within 5 minutes.

**Reliability Perspective:**

SLOs ensure the accuracy and speed of the transfer process, enabling users to complete transactions without issues. Key reliability indicators include login and transfer success rates, as well as timely notification delivery.

# Thoughts
I realized the importance of defining critical user journeys when considering SLOs.

Organizing based on user experience, rather than user flows (expected user actions or screen transitions in the specifications), makes the definition of reliability clearer.

By being able to answer questions like "Why was this SLO defined?" or "Why is this SLO important?", it becomes easier to involve stakeholders outside of development in the design and operation of SLOs.

# References
- [sre.google - Modeling User Journeys](https://sre.google/workbook/implementing-slos/#modeling-user-journeys)
- [cloud.google.com - Standardizing the SLO Design Process](https://cloud.google.com/blog/ja/products/devops-sre/how-to-design-good-slos-according-to-google-sres)
- [popinsight.jp - In-depth Explanation of Critical User Journeys by a Google UX Manager](https://popinsight.jp/blog/?p=40141)
- [speakerdeck.com - Improving SLI/SLO Using Critical User Journeys / #mackerelio](https://speakerdeck.com/heleeen/slo-critical-user-journey)
- [u-site.jp - Comparison of User Journeys and User Flows](https://u-site.jp/alertbox/user-journeys-vs-user-flows)