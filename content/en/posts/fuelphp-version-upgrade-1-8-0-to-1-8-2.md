---
title: Upgrading from FuelPHP 1.8.0 to 1.8.2 and PHP 5.6 to PHP 7.3
slug: fuelphp-version-upgrade-1-8-0-to-1-8-2
date: 2019-10-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
  - FuelPHP
description: Summary of the application version upgrade from FuelPHP 1.8.0 to 1.8.2 and PHP 5.6 to PHP 7.3.
translation_key: fuelphp-version-upgrade-1-8-0-to-1-8-2
---

# Overview
We upgraded FuelPHP from 1.8.0 to 1.8.2 and PHP from 5.6 to 7.3. Since this was part of our business application upgrade efforts, I wanted to summarize our approach.

# Scope
- FuelPHP 1.8.0 → FuelPHP 1.8.2
- PHP 5.6 → PHP 7.3
- Target repositories
  - User-side application
  - Admin-side application
  - Package repository

*Middleware version details are omitted.*
*The OS is Amazon Linux (not version 2).*

FuelPHP 1.8.0 supports up to PHP 7.2, while 1.8.2 supports up to PHP 7.3.

The release occurred suddenly about two weeks before we started the upgrade.
[fuelphp.com - Fuel releases 1.8.2](https://fuelphp.com/blogs/2019/06/fuel-releases-1-8-2)

PHP 7.2 has active support until November 30, 2019, and security support until November 30, 2020, while PHP 7.3 has support until December 6, 2020, and December 6, 2021.

I believe we should be very grateful for the release of FuelPHP 1.8.2, as it extends the support period.

# Duration
About one and a half months.

The first two weeks were spent working intensively in a retreat away from the office. During this period, we had a code freeze, stopping all releases except for emergency fixes.

# Team Composition
- 6 engineers

# Why We Upgraded
The primary motivation was security rather than performance requirements.

Failing to ensure the security of the product could pose a risk to the business.

# Tasks
## Preparations
- Caught up on upgrade information (changes) for FuelPHP and PHP. I read diffs and so on.
- Searched for knowledge
  - I scoured GitHub for previous upgrade efforts, looking at issues and PRs that might be helpful.
  - I searched for and received documents used in previous upgrades that could be useful.
- Research (roughly identifying what modifications might be necessary, checking the status of external libraries and middleware)
  - What modifications might be needed on the application side? What is the status of external libraries? What updates are needed for middleware? ...etc
    - → I wanted to estimate the work, but I only did a rough check. I thought I would get better estimates as I ran tests and confirmed functionality, so I decided to keep it simple. In fact, the work went smoothly, so it might have been correct not to spend too much time on this.
- Formulated an update work policy
  - I made a pledge not to do any refactoring unrelated to the upgrade.
    - → Since there seemed to be no discrepancies in understanding among the team, we did not establish any strict rules.
  - If there were any areas of uncertainty, I handled them as the project leader.
- Developed a schedule
  - Created a rough schedule for development work, QA, and release for the upgrade.
    - I focused more on flexibly handling tasks that would arise rather than on detailed scheduling, so I created a GitHub project kanban for task visualization and progress tracking.

## Branch Management Policy
I prepared a release branch for upgrade work derived from the master branch. We set a two-week code freeze period, but after the freeze was lifted, to allow for releases from the master branch, any functional modifications or additions to the master were rebased into the release branch as needed.

## Progress
The approach for upgrading both FuelPHP and PHP was generally the same:

- First, increase the version.
- Run tests and fix any failures.
- Once everything is complete, conduct functionality confirmation tests. Fix any failures encountered during these tests. *Note 1* *Note 2*

*Note 1: The functionality confirmation tests utilized past assets (test items used in previous upgrade efforts).*
*Note 2: After completing the FuelPHP upgrade, I conducted a functionality confirmation test, then performed the PHP upgrade and conducted another confirmation test.*
(I felt it might have been okay to do a release after completing the FuelPHP confirmation tests.)

This was the flow of progress.

The only difference was that for the PHP upgrade, it was necessary to prepare the staging environment and CI execution environment for PHP 7.3 in advance.

We also held morning meetings to share progress and any issues, and we enjoyed lunch together. It was fun (a simple pleasure).

Since there were a decent number of unit tests prepared, I don't think we faced any hellish situations. (Testing is important.)

The bugs that arose after the work was completed were mainly in areas not covered by tests (such as fat controllers or areas that would likely be detected by E2E), and there were hardly any issues related to insufficient test cases.

# Changes Made During the Upgrade
As I took on a role similar to that of a project leader, I managed progress and task handling while also engaging in actual work. However, I feel I spent more time on infrastructure tasks than on the development work for upgrading Fuel and PHP. (Most of my development work was centered around reviews, although I did handle a few tasks.)

## FuelPHP
Unlike frameworks like Laravel, FuelPHP is not frequently updated, so I had some skepticism about whether the PHP 7.3 support was truly solid (given that it had just been released and had little track record). However, it turned out to be fine.

- Due to changes in the session lifecycle, login and logout broke.
    - → I modified the affected areas to align with the new lifecycle specifications.
- A bug arose because the forge method of the Model could no longer accept the primary key as an argument.
    - → Since there were many places passing the primary key as an argument and there was a need for it, I created a Model class for backward compatibility through inheritance.
- Changes in the specification of ORM's to_array()
    - The order of values returned by to_array() is now guaranteed.
    - → I adjusted the areas where test cases were failing to align with the changes.
- Changes in the specification of the Pagination class
    - There was a change in the casting of values when obtaining the query string.
      - The pagination broke because it no longer cast to int.
        - → I handled this by ensuring that unnecessary strings were not passed.
    - The change of no longer using urldecode affected existing functionality.
      - https://github.com/fuel/core/commit/4614349b243dc48e864a694b017d7356984d9f3c#diff-247dccd657c3e1348e064416f2b1e22bL201
      - → I added url_decode where necessary in the areas where values were being picked up with get().
- Changes in the internal encryption algorithm of Crypt affected session retention.
  - If you upgrade from 1.8.0 to 1.8.2 without consideration, the login session becomes invalid and logs out.
  - This is due to the standardization of the sodium library in PHP 7.2, which FuelPHP 1.8.2 supports.
  - → I extended the core to allow the use of the 1.8.0 encryption method in the config.

## Smarty
There were changes that made me feel emotional, like "What was this minor update?"

Two issues arose during the minor update from v3.1.30 to v3.1.33.

- [Large template parsing error in smarty 3.1.33 #488](https://github.com/smarty-php/smarty/issues/488) 
    - I encountered a bug where, depending on the size of the tpl, you had to carelessly insert a {literal} tag hack to parse the template correctly.
- A bug due to a change in the specification of the date_format tag.
  - It was supposed to return null when null was passed, but it was changed to return the current date and time, causing a bug that outputted inappropriate values in the UI.
    - I implemented a plugin to override the core date_format to address this.

## PHP
Having been away from PHP 7 for about a year, I found that I had forgotten some features and changes, leading to learning opportunities. I feel that I may have taken a somewhat haphazard approach. To respond more carefully, I might have needed to reconsider data structures and adjust method appearances.

- Applying an empty index operator to an empty string results in a fatal error.
    - [php.net - PHP Manual Language Reference Types](https://www.php.net/manual/en/language.types.string.php)
    - This is a change that is not backward compatible since PHP 7.1.
- Errors due to changes in the sorting algorithm.
    - [php.net - Order of Equal Elements](https://www.php.net/manual/en/migration70.incompatible.php#migration70.incompatible.other.sort-order)
    - Bugs appeared in areas that depended on the order.
      - → Since I incorporated LaravelCollection, I utilized collections to avoid dependency on the sorting algorithm.
- The format for Datetime now includes a 'v' to represent milliseconds.
    - [php.net - Parameters](https://www.php.net/manual/en/function.date.php#refsect1-function.date-parameters)
    - → I fixed tests that failed due to differences in milliseconds.
- Performing operations with non-numeric values generates an E_WARNING.
    - [php.net - Notification of Arithmetic Operations with Invalid Strings](https://www.php.net/manual/en/migration71.other-changes.php#migration71.other-changes.apprise-on-arithmetic-with-invalid-strings)
    - This is a change since PHP 7.1.
    - I handled necessary operations by casting to int.
- Errors occurred when applying non-array or non-object arguments to count().
  - `count(null)` also generates a warning.
  - [php.net - Countable Interface](https://www.php.net/manual/en/class.countable.php) 
  - → I handled this by checking for null with `isset` or `!empty` and changing the arguments in the relevant areas to arrays or objects.

# Infrastructure
- I wrote Chef scripts to migrate from the existing PHP 5.6 environment to PHP 7.3.
- Instance construction
    - I ran Chef on a base instance to obtain an AMI. From the AMI launch, I constructed multiple instances.
        - A total of about 20 instances were constructed for staging and production.
- Jenkins
  - Since we use Jenkins for deployment and release tasks, I adjusted jobs as necessary.

## Release
We handled the release with a canary release strategy.
We gradually added PHP 7.3 instances to the target group where PHP 5.6 instances were hanging, running them in parallel, while increasing the number of 7.3 instances and decreasing the number of 5.6 instances.
Once the switch was fully completed, we performed a complete switch (merging the release branch into master).

For some errors investigated in the production environment, we utilized the listener settings of the load balancer.
We specified the source IP as the internal IP to route access from within the company to specific instances for verification. (Dark canary release?)

During the parallel operation period, errors frequently occurred, and we handled issues by disconnecting and reconnecting instances from the load balancer, requiring nearly two weeks to achieve stable operation and complete switching.

# Challenges Faced
Several bugs arose during the release phase that could not be detected through testing or functionality confirmation.

Some were difficult to resolve or cumbersome to address, but whenever something arose, we quickly performed rollback operations (just disconnecting instances from the load balancer) while converging error logs, and it took nearly two weeks to stabilize operations.

From my perspective, the release work seemed more challenging than the upgrade work.

# Others
I searched for blogs and slides related to PHP version upgrades.

The areas that need to be addressed during upgrades vary by product, but I found that the basic flow of work is quite similar.
I felt a strong emotional connection with GameWith, as their environment is very similar.

[We upgraded GameWith, which had been running on FuelPHP for over 5 years, from PHP 5 to 7.3! #GameWith #TechWith](https://tech.gamewith.co.jp/entry/2019/09/26/185515)
[Can't upgrade to PHP 7 due to lack of test code? We solved it with a bot!](https://speakerdeck.com/sgeengineer/tesutokodogawu-kutephp7hefalsebaziyonatupugachu-lai-nai-botutodejie-jue-simasita)
[How we upgraded from PHP 5.5 to 7.2 in 3 months and how we will approach it in the future](https://speakerdeck.com/kosa3/3keyue-dephp5-dot-5kara7-dot-2nibaziyonatupusitaxian-zai-tojin-hou-falsexiang-kihe-ifang-number-phperkaigi-2019?slide=2)

# Performance Improvement
There was a significant improvement in CPU and memory usage, but no major changes were observed in response times.
There are still some areas that have not been properly measured, so investigations are ongoing.

# Reflections
I believe that a manpower-intensive approach is highly effective for upgrade efforts, so focusing on it may allow for relatively quick completion.
It was a valuable experience.

While there was the miraculous release of FuelPHP 1.8.2, we also faced the rare event of a large-scale AWS outage toward the end of the retreat. (This had some impact on functionality verification in the staging environment.)
