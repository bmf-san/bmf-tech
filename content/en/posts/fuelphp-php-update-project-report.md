---
title: FuelPHP and PHP Upgrade Project Report
slug: fuelphp-php-update-project-report
date: 2024-01-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
  - FuelPHP
description: A report on the project to upgrade PHP and FuelPHP for a service in operation for over 10 years.
translation_key: fuelphp-php-update-project-report
---

This article is for [Makuake Advent Calendar 2023](https://adventar.org/calendars/8992), Day 24.

# Overview
This is a report on the project to upgrade PHP and FuelPHP for a service that has been in operation for over 10 years.

- **Upgrade Targets**
  - Main service (the company's primary application)
    - Uses FuelPHP
  - Two internal libraries used by the main service
    - Built with PHP
  - Execution environment of the main service
    - Container environment using ECS
    - Uses PHP-FPM
- **Target Versions**
  - PHP 7.3 to PHP 8.1
  - FuelPHP 1.8.2 to FuelPHP 1.9-develop
- **Team Structure**
  - 2 in-house engineers
    - Responsible for project management, research, and implementation
  - 1 contract engineer
    - Responsible for research and implementation
      - Due to internal resource constraints, the main implementation work was assigned to the contract engineer
- **Timeline**
  - From May 2022 to April 2023
    - Completed in April 2023 with a one-month delay
      - The delay was mainly due to the high workload on in-house engineers who were involved in the project while also working in their respective teams, and the time required to validate the infrastructure configuration for the upgrade project.

# Background
PHP 7.3 reached its EOL in December 2021, ending security updates, making an upgrade necessary. FuelPHP 1.8.2, which depends on PHP, was in a similar situation.

In this project, it was necessary to evaluate whether simply upgrading was the optimal choice. Unlike previous upgrade projects, the development of FuelPHP had stagnated, and it was necessary to consider whether continuing to use FuelPHP was appropriate from a technical strategy perspective.

For more details on the state of FuelPHP, refer to [FuelPHP's Current Status as of March 2023](https://bmf-tech.com/posts/FuelPHP%E3%81%AE2023%E5%B9%B43%E6%9C%88%E7%8F%BE%E5%9C%A8%E3%81%AE%E7%8F%BE%E6%B3%81).

# Planning
## Preliminary Research
Before starting the project, the following research was conducted:

- Reviewed migration information from PHP 7.3 to PHP 8.1 on [www.php.net](https://www.php.net/docs.php)
- Investigated PHP compatibility in the codebase
  - Used tools like Rector and PHPCompatibility, and searched with grep based on documentation to identify compatibility issues for the upgrade from PHP 7.3 to PHP 8.1
    - Investigation items were based on [PHP Manual > Appendices](https://www.php.net/manual/ja/appendices.php)
      - New features/classes/interfaces, new functions, new global constants, backward-incompatible changes, deprecated features in PHP 8.3.x, other changes, Windows support
- Investigated runtime errors
  - Set up a PHP 8.2 environment, ran tests, and manually checked application behavior to identify runtime errors
- Checked the PHP 8.1 compatibility of dependent packages
  - Identified packages that were already compatible, needed updates, or required alternative solutions due to incompatibility
- Identified necessary adjustments for CI/CD
- Reviewed installed extensions in the execution environment and identified those needing updates or additions

Based on these investigations, we identified areas requiring fixes and estimated the overall project workload.

## Monolithic Application Architecture Strategy
The main service application is a monolith developed by multiple teams. A significant challenge has been determining how to modernize this monolith as part of the development organization’s strategy.

The update strategy considered the future of the technology stack, including PHP and FuelPHP, and the architecture of the main service. Several plans were evaluated:

1. Replace with another framework
2. Replace with another language
3. Service decomposition
4. Transition to a modular monolith architecture
5. Continue using FuelPHP

Each plan had its pros and cons. Ultimately, we chose to continue using FuelPHP for the following reasons:

- To mitigate security and compliance risks from prolonged use of EOL software
- Anticipation that some code would remain in the monolith even if decomposition was pursued, necessitating a future decision on FuelPHP
- Desire to complete the upgrade with minimal internal resources

## Selecting Target Versions for PHP and FuelPHP
Since we decided to continue using FuelPHP, we determined the target versions for PHP and FuelPHP based on research.

For PHP, we considered upgrading from PHP 7.3 to PHP 7.4, PHP 8.0, or PHP 8.1. PHP 7.4 was excluded as it reached EOL in November 2022, and PHP 8.0 was excluded as it would reach EOL shortly after the project’s completion. Thus, PHP 8.1 was chosen.

For FuelPHP, we considered two options:

- Forking FuelPHP to support PHP 8.1
- Using FuelPHP 1.9-develop

We chose FuelPHP 1.9-develop because:

- It was likely to work with PHP 8.1 based on codebase research and inquiries with developers
- It had ongoing, albeit irregular, development activity, suggesting potential for an official release
- Using 1.9-develop was less costly than forking FuelPHP

However, risks included low test coverage and uncertainty about future PHP 8.2 support.

# Upgrade Strategy
## Policies
The following policies guided the upgrade project:

- Avoid big bang releases
- Minimize code freeze periods
- Enable quick rollback in case of issues
- Avoid complicating development workflows during the project

## Upgrade Strategy
We planned a phased architecture change from PHP 7.3 to PHP 8.1, enabling parallel operation of both environments. The plan included five phases:

1. Prepare for parallel operation in the staging environment
2. Start parallel operation in the staging environment
3. Extend parallel operation to the production environment
4. Transition to PHP 8.1-only operation
5. Complete the transition to PHP 8.1

Each phase is detailed in the blog post with diagrams illustrating the infrastructure changes.

# Implementation
## Code Modifications
Application source code was categorized into two patterns:

1. Code that could be modified to work with both PHP 7.3 and PHP 8.1
2. Code that required version-specific modifications

For the latter, we used conditional statements based on PHP version:

```php
if (version_compare(PHP_VERSION, '7.4.0') < 0){
	// Code for PHP < 7.4
}

if (version_compare(PHP_VERSION, '8.1.0') >= 0) {
	// Code for PHP 8.1
}
```

Dependency packages were categorized into three patterns:

1. Compatible with both PHP versions
2. Not compatible with both versions but manageable with separate `composer.json` files
3. Incompatible and required forking for PHP 8.1 support

Forking was necessary for two cases, including the `ruflin/Elastica` library.

## Infrastructure
We modified the existing ALB+ECS environment to support parallel operation of PHP 7.3 and PHP 8.1. We used CloudFront’s Continuous Deployment feature to distribute traffic between environments.

Final architecture:

![Infrastructure Diagram](https://github.com/bmf-san/bmf-tech-client/assets/13291041/ab6725f2-7c6e-4ecc-91df-0fafa41f6cf6)

# Testing
## QA
We conducted QA in both PHP 7.3 and PHP 8.1 environments, using comprehensive test cases executed through the UI.

## Load Testing
We used `k6` to ensure no performance degradation with PHP 8.1. Tests showed a 25% improvement in response times.

# Results
The project was completed successfully with no major issues, thanks to careful planning.

# Reflections
Despite challenges such as team member transitions, thorough documentation and communication minimized project impact. However, areas for improvement include test coverage, reducing dead code, updating dependencies, and improving QA efficiency.

This was my second upgrade project at my current company, and it highlighted the increasing complexity of such projects as organizations and architectures evolve. While the future of FuelPHP remains uncertain, I aim to apply the lessons learned to future projects.