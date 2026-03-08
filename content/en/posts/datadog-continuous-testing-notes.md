---
title: Notes on Datadog Continuous Testing
slug: datadog-continuous-testing-notes
date: 2023-01-31T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - Datadog
translation_key: datadog-continuous-testing-notes
---



# Overview
Notes on what I researched about Datadog continuous testing.

# What is Datadog Continuous Testing
- No-code
    - Tests can be set up with simple clicks
- Self-healing E2E
    - Tests with "resilience"
    - Minimizes false positives
- Supports cross-browser testing
- Covers major integrations
    - CircleCI, Github Actions, Jenkins, etc.

It seems to have mechanisms to ease the operation of E2E and maintain test reliability.

# Introduction
The [Datadog test recorder](https://chrome.google.com/webstore/detail/datadog-test-recorder/kkbncfpddhdmkfmalecgnphegacgejoa) Chrome extension is required.

No other preparation is needed, and you can start using it immediately.

# Creating Tests
1. Create a Browser Test
2. Record the test case

It seems better to open the recording in a popup (since the UI in the recording screen is an iframe).

# Test Configuration
Setting various options related to test execution.

You can handle Basic authentication, set cookies, request headers, etc.

cf. [Test Configuration](https://docs.datadoghq.com/ja/synthetics/browser_tests/?tab=%E3%83%AA%E3%82%AF%E3%82%A8%E3%82%B9%E3%83%88%E3%82%AA%E3%83%97%E3%82%B7%E3%83%A7%E3%83%B3#%E3%83%86%E3%82%B9%E3%83%88%E3%82%B3%E3%83%B3%E3%83%95%E3%82%A3%E3%82%AE%E3%83%A5%E3%83%AC%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3)

# Test Items
Let's look at the test items available in Browser Tests.

cf. https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B

## Assertion
Assertions to check around the DOM.

- [Test the content of an element](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E8%A6%81%E7%B4%A0%E3%81%AE%E3%82%B3%E3%83%B3%E3%83%86%E3%83%B3%E3%83%84%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B)
- [Test the attributes of an element](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E8%A6%81%E7%B4%A0%E3%81%AE%E5%B1%9E%E6%80%A7%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B)
- [Test if an element exists](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%81%82%E3%82%8B%E8%A6%81%E7%B4%A0%E3%81%8C%E5%AD%98%E5%9C%A8%E3%81%99%E3%82%8B%E3%81%8B%E3%81%A9%E3%81%86%E3%81%8B%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B)
    - You can choose CSS or XPath 1.0 to test the existence of an element

## Navigation
Transition-related.

- [Refresh a page](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%83%9A%E3%83%BC%E3%82%B8%E3%82%92%E6%9B%B4%E6%96%B0%E3%81%99%E3%82%8B)
    - There doesn't seem to be a super reload.
        - If you want a super reload, you might handle it with JavaScript
- [Navigate to an email and click a link](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%83%A1%E3%83%BC%E3%83%AB%E3%81%AB%E7%A7%BB%E5%8B%95%E3%81%97%E3%81%A6%E3%83%AA%E3%83%B3%E3%82%AF%E3%82%92%E3%82%AF%E3%83%AA%E3%83%83%E3%82%AF%E3%81%99%E3%82%8B)
- [Follow a specific link](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E7%89%B9%E5%AE%9A%E3%81%AE%E3%83%AA%E3%83%B3%E3%82%AF%E3%82%92%E3%81%9F%E3%81%A9%E3%82%8B)

## Special Actions
UI-related operations.

- [Hover](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%83%9B%E3%83%90%E3%83%BC)
- [Key press](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%AD%E3%83%BC%E3%81%AE%E6%8A%BC%E4%B8%8B)
- [Scroll](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%B9%E3%82%AF%E3%83%AD%E3%83%BC%E3%83%AB)
- [Wait](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E5%BE%85%E6%A9%9F)

## Variables
You can define arbitrary variables.

- [Pattern](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3)
- [Element](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E8%A6%81%E7%B4%A0)
    - Built-in options include numeric, alphabetic, alphanumeric, date, timestamp
    - There is an option to obfuscate local variable values of test results
- [JavaScript](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#javascript)
    - You can load and execute defined functions
    - Supports both synchronous and asynchronous
- [Global Variables](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%B0%E3%83%AD%E3%83%BC%E3%83%90%E3%83%AB%E5%A4%89%E6%95%B0)
    - You can use global variables defined in [Synthetic Monitoring Settings](https://docs.datadoghq.com/ja/synthetics/settings/?tab=specifyvalue)
- [Global Variables-MFA](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%B0%E3%83%AD%E3%83%BC%E3%83%90%E3%83%AB%E5%A4%89%E6%95%B0---mfa)
    - You can use MFA global variables defined in [Synthetic Monitoring Settings](https://docs.datadoghq.com/ja/synthetics/settings/?tab=specifyvalue)
        - Supports TOTP for MFA
            - [TOTP for Multi-Factor Authentication (MFA) in Browser Tests](https://docs.datadoghq.com/ja/synthetics/guide/browser-tests-totp/)
- [Email](https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#email)
    - You can create email addresses for testing email sending/receiving confirmation, email content verification, link clicking, etc.
    - A unique email address and mailbox are generated each time the test is executed

## Subtest
You can reuse existing browser tests within another browser test.
You can nest up to two levels for reuse.

## HTTP Request
You can execute HTTP requests within a browser test.

# Self-healing
It's not clear how much repair is actually possible.
It is stated that changes in the UI are automatically detected and explored.

 > If there is a UI change that modifies an element (e.g., moves it to another location), the test will automatically locate the element again based on the points of reference that were not affected by the change. Once the test completes successfully, Datadog will recompute, or “self-heal,” any broken locators with updated values. This ensures that your tests do not break because of a simple UI change and can automatically adapt to the evolution of your application’s UI.
In the next section, we’ll look at how you can fine-tune your test notifications to ensure that you are only notified of legitimate failures.

cf. https://www.datadoghq.com/ja/blog/test-maintenance-best-practices/

It's best not to raise expectations too high as it is limited to "simple UI changes."

> When the test runs successfully, the browser test will recalculate (or "self-heal") broken locators with updated values, ensuring that the test does not break due to a simple UI update and automatically adapts to the application's UI.

cf.  https://docs.datadoghq.com/ja/synthetics/browser_tests/advanced_options/

# Parallelization
You can set parallelization from Synthetic settings.

cf. https://docs.datadoghq.com/ja/continuous_testing/settings/

# Dashboard
## Explore
You can search for results of Synthetic Monitoring and Continuous Testing (CI Batches) and test execution results (Test Runs).
cf. https://docs.datadoghq.com/ja/continuous_testing/explorer/?tab=cibatches

## Test coverage
> Provides insights into the overall test coverage of RUM applications using browser data collected from RUM and Synthetic browser test results
cf. https://docs.datadoghq.com/ja/synthetics/test_coverage/

It seems possible to analyze what is not being tested. It could help improve the comprehensiveness of test cases.
Track the transition of coverage

# CI Integration
Synthetic tests can integrate with various CIs.

Available integrations:
- Azure DevOps Extension
- CircleCI Orb
- Github Actions
- GitLab
- Jenkins
- NPM package

cf. https://docs.datadoghq.com/continuous_testing/cicd_integrations/

# Codification
cf. https://registry.terraform.io/providers/DataDog/datadog/latest/docs/resources/synthetics_test

Since it is a DOM-based test, it seems difficult to create test cases from code...

# Notifications
Since it is managed as Synthetics, there are no particular concerns about notifications.

# Cost
1000 times/$12 (on-demand $18)
cf. https://www.datadoghq.com/ja/pricing/?product=continuous-testing#continuous-testing

I think it's cheap.
Is it free for less than 1000 times??

# Scenario Management Method
Will it break down when many scenarios are prepared?

As a policy for managing test scenarios:
- Views
- tag
It seems to be in the form of using. There doesn't seem to be any other management UI.

As a policy for managing test cases, it is recommended to create tests in a DRY manner, and it seems better to actively use **subtests**.
cf. https://docs.datadoghq.com/ja/synthetics/browser_tests/actions/?tab=%E3%82%A2%E3%82%AF%E3%83%86%E3%82%A3%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%A7%E8%A6%81%E7%B4%A0%E3%82%92%E3%83%86%E3%82%B9%E3%83%88%E3%81%99%E3%82%8B#%E3%82%B5%E3%83%96%E3%83%86%E3%82%B9%E3%83%88

# Notes on Points of Interest
- Since Browser Test can obtain the same information as RUM, it might be useful as reference information for debugging
- When creating a Browser Test, you need to set the test frequency, but it cannot be set to 0, so it is assumed to be executed regularly?
    - It seems okay if you set it to PAUSED
- Parallelization seems essential
    - It is not enabled, so it seems necessary to enable it?
        - It seems like an add-on
            - cf. https://www.datadoghq.com/ja/pricing/?product=continuous-testing#continuous-testing
- Be careful with data cleanup
    - In cases where data is generated during test execution, if the test fails midway, data will be generated in an incomplete state.
        - Even if you create test cases with cleanup in mind, data will remain if the test fails.
            - How to ensure cleanup?
- Things that seem untestable
  - Selecting the OS on the test execution device
    - You can select the type of browser and browser size, but not the OS
- SMS authentication is not supported
- Native apps are not supported
- Verification of push notifications is also difficult
- Not limited to E2E, but it seems important to have a mechanism to observe parts that change frequently or where tests are prone to breakage. The dashboard seems to allow for visualization of that.

# References
- [www.datadoghq.com - continuous-testing](https://www.datadoghq.com/ja/product/continuous-testing/)
- [docs.datadoghq.com](https://docs.datadoghq.com/ja/continuous_testing/)
- [www.datadoghq.com - Use Datadog Continuous Testing to release with confidence](https://www.datadoghq.com/ja/blog/release-confidently-with-datadog-continuous-testing/)
- [docs.datadoghq.com - Continuous Testing and CI/CD](https://docs.datadoghq.com/ja/continuous_testing/cicd_integrations/)
- [www.datadoghq.com - Best practices for creating end-to-end tests](https://www.datadoghq.com/ja/blog/test-creation-best-practices/)
- [www.datadoghq.com - Best practices for continuous testing with Datadog](https://www.datadoghq.com/ja/blog/best-practices-datadog-continuous-testing/)
