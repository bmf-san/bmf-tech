---
title: FuelPHP and PHP Update Project Report
slug: fuelphp-php-update-project-report
date: 2024-01-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
  - FuelPHP
translation_key: fuelphp-php-update-project-report
---

This article is the 24th entry of the [Makuake Advent Calendar 2023](https://adventar.org/calendars/8992).

# Overview
This report details a project to update PHP and FuelPHP for a service that has been in operation for over 10 years.

- Update Targets
  - Main service (the primary service operated by the company) application
    - Uses FuelPHP
  - Internal libraries used by the main service (two)
    - Written in PHP
  - Execution environment of the main service
    - Container environment using ECS
    - Uses PHP-FPM
- Target Versions for Update
  - From PHP 7.3 to PHP 8.1
  - From FuelPHP 1.8.2 to FuelPHP 1.9-develop
- Team Structure
  - 2 in-house engineers
    - Responsible for project progress, research, and implementation
  - 1 contracted engineer
    - Responsible for research and implementation
      - Due to internal resource adjustments, the main implementation work was requested from the contracted engineer.
- Duration
  - From May 2022 to April 2023
    - Ultimately completed in April 2023 with a one-month delay
      - The main reasons for the delay are believed to be the high workload of in-house engineers being involved in the update project while belonging to their respective teams, and the time taken to verify the infrastructure configuration for the update project.

# Background
PHP 7.3 reached EOL in December 2021, and security updates have ended, necessitating consideration for an update. FuelPHP 1.8.2, which depends on PHP, was in a similar situation.

In this project, it was necessary to consider whether simply updating was optimal. Unlike previous update projects, the development of FuelPHP had stagnated, and it was necessary to evaluate whether continuing to use FuelPHP was appropriate from a technical strategy perspective.

*For the current status of FuelPHP, please refer to the previous summary in [FuelPHP's Status as of March 2023](https://bmf-tech.com/posts/FuelPHP%E3%81%AE2023%E5%B9%B43%E6%9C%88%E7%8F%BE%E5%9C%A8%E3%81%AE%E7%8F%BE%E6%B3%81).*  

# Planning
## Preliminary Research
As preliminary research before starting the project, the following investigations were conducted:

- Gathered migration information from [www.php.net](https://www.php.net/docs.php) for migrating from PHP 7.3 to PHP 8.1.
- Conducted a PHP compatibility investigation of the codebase.
  - Utilized tools like Rector and PHPCompatibility, and searched through documentation using grep to investigate compatibility for the update from PHP 7.3 to PHP 8.1.
    - The investigation items were based on [PHP Manual > Appendices](https://www.php.net/manual/ja/appendices.php).
      - New features/new classes and interfaces/newly added functions/new global constants/incompatibilities/features deprecated in PHP 8.3.x/other changes/Windows support.
- Investigated runtime errors.
  - Prepared a PHP 8.2 execution environment, executed tests, and manually verified application behavior to identify runtime errors.
- Investigated the PHP 8.1 compatibility status of dependent packages.
  - Identified those that were already compatible, those needing updates, and those that required some action due to incompatibility.
- Identified areas needing adjustments related to CI/CD.
- Checked installed extension modules in the execution environment and identified those that might need updates or additions.

Based on the investigation items, identified areas needing modifications and estimated the overall effort required for the project.

## Monolithic Application Architecture Strategy
The application used by the main service is a so-called monolith, developed by multiple teams.

As a strategy for the development organization, how to refresh this monolith has been a significant challenge in recent years.

The strategy for updating the technology stack and architecture used in the main service, including PHP and FuelPHP, was established.

For this update, several plans were considered:

- 1. Replace with a different framework
  - Move away from FuelPHP and rewrite using another framework.
- 2. Replace with a different language
  - Move away from FuelPHP and PHP and rewrite in a different language, considering the use of a framework.
- 3. Service splitting
  - Maintain the monolith while extracting part of it as a separate service.
  - Delay updates to FuelPHP and PHP for the time being, and consider the next steps after minimizing the monolith.
- 4. Change architecture to a modular monolith
  - Maintain the monolith while performing module splitting.
  - Service splitting is a prerequisite, and after module splitting, consider extracting to separate services.
- 5. Continue using FuelPHP
  - Proceed with updates for PHP and FuelPHP.

Each of these has its pros and cons (details omitted), and after comparison, the plan to "continue using FuelPHP" was chosen.

Reasons for this choice include:

- Avoiding security and compliance risks associated with leaving EOL software unattended for too long.
- It is anticipated that even if the monolith is split, there will still be code remaining in the monolith, necessitating a reevaluation of what to do with FuelPHP, making it unreasonable to postpone the update.
- Aiming to complete the update with minimal internal resources.
  - Cannot impose a heavy burden of rewriting or splitting the monolith.

## Selection of Target Versions for PHP and FuelPHP Updates
Having decided to continue using FuelPHP, the versions to which FuelPHP and PHP would be updated were determined based on investigations.

During the project planning phase, the candidates for the update version from PHP 7.3 were PHP 7.4, PHP 8.0, and PHP 8.1. (PHP 8.2 had not yet been released.)

PHP 7.4 reached EOL in November 2022, so it was excluded as it would become EOL during the project period.

PHP 8.0 has an EOL in November 2023, but it would likely reach EOL shortly after the project completion, so it was similarly excluded.

According to the official [release information](https://fuelphp.com/blogs/2019/06/fuel-releases-1-8-2), FuelPHP 1.8.2 supports up to PHP 7.3, so it was necessary to update FuelPHP to a version that supports PHP 8.1. However, the version being used, PHP 1.8.2, was the latest, and the next version had not been released.

Thus, two plans were considered:

- Fork FuelPHP to support PHP 8.1.
- Use FuelPHP 1.9-develop.

As a result of the investigation, it was decided to use FuelPHP 1.9-develop for the following reasons:

- Based on codebase investigations and inquiries to FuelPHP developers, it appeared that 1.9-develop could run on PHP 8.1.
- Although the release schedule is undecided, commits have been made irregularly, and considering the developers' intentions, it seemed possible that it might be officially released in the future.

FuelPHP 1.9-develop has not yet been officially released, but it was determined that there was room for adoption based on investigations.

Using 1.9-develop was also considered to be more cost-effective than forking FuelPHP.

However, there are risks and disadvantages, such as "the test coverage of the framework itself being quite low" and "the likelihood of PHP 8.2 support not being expected in the near future even if released."

Another approach considered was to become a committer for FuelPHP, but it was deemed uncertain how much could be contributed to accelerating FuelPHP's release cycle, so this was abandoned.

# Update Strategy
## Policy
The following policies were established for the update project:

- Avoid big bang releases.
- Minimize the code freeze period.
- Quickly roll back when issues arise.
- During the update project period, avoid complicating the development flow or increasing implementation effort as much as possible.

## Update Strategy
Based on the policy, it was deemed desirable to gradually implement architectural changes from PHP 7.3 to PHP 8.1.

To achieve this, it was necessary to build a structure that could operate both PHP 7.3 and PHP 8.1 environments in parallel.

To realize such a structure, five phases were established to achieve gradual architectural changes.

### Phase 1: Production and Staging Environments Running PHP 7.3
This phase serves as a preparation period to create an environment where PHP 7.3 and PHP 8.1 can run in parallel in the staging environment.

During this phase, the following was done:

- Updated the application codebase and libraries to work on both PHP 7.3 and PHP 8.1.

<img width="594" alt="phase1" src="https://github.com/bmf-san/bmf-tech-client/assets/13291041/c7a797ec-6685-4a01-8957-2bf8b5ce5777">

### Phase 2: Start of Parallel Operation in Staging Environment
This phase marks the start of parallel operation of PHP 7.3 and PHP 8.1 in the staging environment only.

QA was conducted, and load testing was performed to validate the pre-production environment.

Particularly, the infrastructure configuration for parallel operation was verified to ensure no issues would arise in production.

<img width="883" alt="phase2" src="https://github.com/bmf-san/bmf-tech-client/assets/13291041/54c1d027-4a73-4875-a8b2-9a71e7d667ad">

### Phase 2.5: Start of Parallel Operation in Production Environment
This phase involves deploying the parallel operation environment from the staging environment to the production environment and starting operations.

Monitoring and bug fixes arising from the start of operations are conducted, aiming to stabilize operations in the production environment and prepare for a complete switch to PHP 8.1.

<img width="886" alt="phase2 5" src="https://github.com/bmf-san/bmf-tech-client/assets/13291041/013aa6ad-a05a-4ff3-80ed-ba5cf0c408d6">

### Phase 3: Start Switching Staging and Production Environments to PHP 8.1
This phase involves switching the staging and production environments, which are running in parallel with PHP 7.3 and PHP 8.1, to operate solely on PHP 8.1, validating operational stability.

It is anticipated that there will be a certain level of stability at the Phase 2.5 stage, but due to increased traffic when operating solely on PHP 8.1, this phase is established for caution.

In this phase, infrastructure resources related to the PHP 7.3 environment are retained to allow for a rollback to the PHP 7.3 environment.

<img width="882" alt="phase3" src="https://github.com/bmf-san/bmf-tech-client/assets/13291041/74fc21f4-302f-44f8-9180-59f5dd0ebec7">

### Phase 3.5: Complete Switch to PHP 8.1 Environment
In this phase, various infrastructure resources related to the PHP 7.3 environment and the code branching for parallel operation are removed, completing the switch to the PHP 8.1 environment.

In this phase, rolling back to the PHP 7.3 environment is fundamentally impossible (it can be done if necessary, but quick rollback is not feasible).

<img width="546" alt="phase3 5" src="https://github.com/bmf-san/bmf-tech-client/assets/13291041/815b6917-7f56-4482-bf78-03ca71e47ca8">

# Implementation
## Modification Policy
The application source code, excluding dependent packages, could be broadly categorized into two patterns:

- Those that could be modified to work correctly on both PHP 7.3 and PHP 8.1.
- Those that could not be modified to work correctly on both PHP 7.3 and PHP 8.1.

The former only required simple modifications, while the latter required conditional branching based on the PHP version to ensure correct operation on each version.

```php
// Code example
if (version_compare(PHP_VERSION, '7.4.0') < 0){
	// Code for versions below 7.4.0
}

if (version_compare(PHP_VERSION, '8.1.0') >= 0) {
	// Code for PHP 8.1 support
}
```

This conditional branching was defined as a helper function and utilized in a feature toggle-like manner at each modification point.

On the other hand, there were three patterns regarding dependent packages:
1. Those that could be updated to versions functioning correctly on both PHP 7.3 and PHP 8.1.
2. Those that could not be updated to versions functioning correctly on both PHP 7.3 and PHP 8.1.
3. Those that could not be updated to versions functioning correctly on both PHP 7.3 and PHP 8.1 and required forking for PHP 8.1 support.

The first pattern only required straightforward updates, while the other patterns required respective actions.

The second pattern was addressed by preparing separate composer.json files for each environment of PHP 7.3 and PHP 8.1.

Due to this approach, until Phase 3.0, both composer.json files needed to specify the same dependent packages, but since library additions did not occur frequently, it did not become a significant burden.

The third pattern had two cases.

One was the fork of the PHP Elasticsearch client library [ruflin/Elastica](https://github.com/ruflin/Elastica).

The version of Elasticsearch used by the service was quite old, and the PHP 8.1 compatible version of ruflin/Elastica could not be utilized.

Therefore, ruflin/Elastica was forked, and PHP 8.1 support was implemented. (About six months after the completion of the update project, some features using Elasticsearch no longer required the client library, rendering the forked repository obsolete.)

The other was the fork of an internal library.

The internal library was also used in another internal service operated with PHP 7.3, so it was necessary to ensure that the internal library operated correctly on both PHP 7.3 and PHP 8.1.

If the internal library could be modified to separate processing based on PHP version by dividing the composer.json files, forking would not have been necessary, but a good approach could not be devised, leading to the decision to fork.

As a result, until Phase 3.0, there was an additional burden to synchronize the specifications between the forked source and the forked destination, but since it was not a library that underwent frequent specification changes, it did not become a significant burden.

After the completion of the update project, the forked source and destination could be operated as separate entities, allowing synchronization efforts to be unnecessary after Phase 3.5.

## Infrastructure Construction
To build a parallel operating environment for PHP 7.3 and PHP 8.1, it was necessary to modify the existing execution environment.

The existing execution environment was composed of ALB + ECS, utilizing Nginx as the web server.

<img width="491" alt="before" src="https://github.com/bmf-san/bmf-tech-client/assets/13291041/a46f532a-72e9-4026-96c3-48adfcd5d13c">

The existing execution environment was modified to meet the following requirements:

- Quick switching between PHP 7.3 and PHP 8.1 environments.
- Flexible traffic distribution between PHP 7.3 and PHP 8.1 environments.
- Ability to collect logs separately for PHP 7.3 and PHP 8.1 environments (ensuring observability).

Several approaches were considered for the modifications:

- Using Nginx for traffic distribution.
- Using Route53 for traffic distribution.
- Using NLB for traffic distribution.

Ultimately, it was decided to utilize the newly released feature of CloudFront's Continuous Deployment.

cf. [Using CloudFront continuous deployment to safely test CDN configuration changes](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/continuous-deployment.html)

This feature met the requirements by allowing the distribution of traffic between two CloudFront Distributions: Primary and Staging.

The traffic distribution condition adopted was weight-based, which has the constraint of only being able to allocate 0-15% of all requests to the distribution target, but it was determined that accepting traffic at a maximum of 15% for a certain period would gather sufficient traffic necessary for switching decisions.

The final configuration was as follows:

<img width="883" alt="after" src="https://github.com/bmf-san/bmf-tech-client/assets/13291041/ab6725f2-7c6e-4ecc-91df-0fafa41f6cf6">

# Testing
## QA
Due to the operation of the parallel environments for PHP 7.3 and PHP 8.1, code modifications were made to execute the main branch in both PHP 7.3 and PHP 8.1 environments, necessitating QA in both execution environments.

For QA execution, comprehensive test cases for the entire service were prepared and executed on the UI.

## Load Testing
To verify that there was no performance degradation due to the update to PHP 8.1, load testing was conducted using k6.

Tests were conducted for each expected traffic pattern, and response times showed approximately a 25% improvement.

# Results
Thanks to the planned execution, no major issues arose (though some difficult-to-solve problems did occur), and the update project was successfully completed.

# Reflections
This project initially started with a team of three engineers, including myself, but two of them took parental leave and left at different times, leading to a relay-like project with handovers to other members.

Even with such a structure, thorough documentation and internal communication minimized the impact on project execution.

Updating a monolithic application touched by multiple teams tends to incur high communication costs, but I realized that proper planning allows for smooth updates.

On the other hand, I also felt that there were several areas that would require improvement in the future.

Insufficient test coverage, excessive dead code, outdated dependent libraries that are not updated, and the efficiency of QA execution are some areas that I felt need ongoing improvement through the update project.

This was my second update project in my current position, but the organizational structure, architectural composition, and application state were different from the previous one, leading to an increased number of considerations. (This is also a challenge...)

I wonder what the next update will be like (whether to consider a forked FuelPHP or take another strategy...), and the future of FuelPHP (whether there will be a next release...) remains uncertain, but I hope to leverage the insights from this update project in the next one.