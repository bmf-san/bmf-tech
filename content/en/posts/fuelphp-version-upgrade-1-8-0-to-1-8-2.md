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
We upgraded from FuelPHP 1.8.0 to 1.8.2 and from PHP 5.6 to PHP 7.3. This post summarizes the efforts made for the application version upgrade in our work.

# Scope
- FuelPHP 1.8.0 → FuelPHP 1.8.2
- PHP 5.6 → PHP 7.3
- Target repositories
  - User-side application
  - Admin-side application
  - Package repository

*Middleware versions are omitted
*OS is Amazon Linux (not version 2)

FuelPHP 1.8.0 supports up to PHP 7.2, while 1.8.2 supports up to 7.3.

The upgrade was released unexpectedly about two weeks before we started.
[fuelphp.com - Fuel releases 1.8.2](https://fuelphp.com/blogs/2019/06/fuel-releases-1-8-2)

PHP 7.2's active support ended on November 30, 2019, and security support on November 30, 2020. PHP 7.3's dates are December 6, 2020, and December 6, 2021, respectively.

The release of FuelPHP 1.8.2, which extends the support period, was greatly appreciated.

# Duration
About one and a half months

The first two weeks were spent intensively working away from the office in a retreat format. During this period, we implemented a code freeze, halting all releases except for emergency responses.

# Team Structure
- 6 engineers

# Why Upgrade?
The primary reason was security rather than performance requirements.

Not being able to ensure the product's security could pose a risk to the business.

# Work
## Preparation
- Caught up on version upgrade information (changes) for FuelPHP and PHP. Read diffs.
- Searched for knowledge
  - Searched GitHub for previous upgrade work, looking at issues and PRs that might be helpful.
  - Sought and was informed about usable documents from previous upgrades.
- Investigation (roughly identified necessary fixes, external library and middleware support status)
  - What fixes are needed on the application side? What is the status of external library support? What middleware updates are needed?...etc.
    - → Tried to estimate the work but only did it roughly. Thought estimates would be clearer through testing and verification, so I was pragmatic. In the end, the work went smoothly, so it might have been right not to spend too much time.
- Formulated update work policy
  - Decided not to do refactoring unrelated to the update.
    - → There seemed to be no discrepancies in understanding within the team, so no strict rules were set.
  - If there were any difficult decisions, I handled them as the project leader.
- Scheduled
  - Created a rough schedule for upgrade development work, QA, and release.
    - Focused more on flexibly handling tasks that would arise rather than a detailed schedule, so I used GitHub projects to visualize tasks and track progress.

## Branch Operation Policy
Created a release branch for upgrade work derived from master. A two-week code freeze was set, but after the freeze was lifted, releases on the master branch were allowed. If there were feature modifications or additions to master, changes were rebased into the release branch.

## Progress
The approach to upgrading FuelPHP and PHP was generally the same:

- First, raise the version
- Run tests and fix any failures
- Once everything is complete, perform operational verification tests. Fix any issues found during verification. *1 *2

*1 Operational verification tests used updated past assets (test items used in previous upgrades).
*2 After completing FuelPHP upgrade support, operational verification was done once, then PHP upgrade support was followed by another verification.
(It might have been okay to release after FuelPHP verification was done.)

The difference with PHP support was the need to prepare the staging environment and CI execution environment for PHP 7.3 in advance.

We also held morning meetings to share progress and challenges, and enjoyed lunches together. It was fun (simple joy).

Having a decent amount of Unit Tests prepared prevented us from experiencing a nightmare (tests are important).

Bugs that occurred after work completion were mainly in areas not covered by tests (e.g., fat controllers or areas detectable by E2E), with few cases of insufficient test cases.

# Areas Addressed in the Upgrade
As a project leader, I managed progress and handled tasks while also doing some actual work. However, I felt I spent more time on infrastructure work than on development work for Fuel and PHP upgrades (most development work was review-focused, though I did handle some tasks).

## FuelPHP
Unlike Laravel, FuelPHP is not frequently updated, so there was some skepticism about whether it fully supported PHP 7.3 (since it was released recently and had little track record). However, it turned out to be fine.

- Changes in session lifecycle broke login and logout.
    - → Modified affected areas to match the new lifecycle specifications.
- Bug due to inability to pass primary key as an argument in Model's forge method.
    - → Created an inherited Model class for backward compatibility as there were many places passing primary keys as arguments.
- Change in ORM's to_array() specification
    - The order of values in the array returned by to_array() is now guaranteed.
    - → Adjusted test cases to match the changes.
- Pagination class specification change
    - Change in casting values when obtaining query strings.
      - Pagination broke due to no longer casting to int.
        - → Handled by casting to avoid passing unnecessary strings.
    - Impact on existing functionality due to no longer urldecoding
      - https://github.com/fuel/core/commit/4614349b243dc48e864a694b017d7356984d9f3c#diff-247dccd657c3e1348e064416f2b1e22bL201
      - → Added url_decode where needed when picking up values with get().
- Impact on session retention due to change in Crypt's internal encryption algorithm
  - Upgrading from 1.8.0 to 1.8.2 without consideration invalidated login sessions, causing logouts.
  - FuelPHP 1.8.2 supports the sodium library, standardized from PHP 7.2.
  - → Extended core to use 1.8.0 encryption method via config.

## Smarty
There were changes that made me feel "what was the minor update?" and brought emotional moments.

Two issues arose from the minor update from v3.1.30 to v3.1.33.

- [Large template parsing error in smarty 3.1.33 #488](https://github.com/smarty-php/smarty/issues/488)
    - Encountered a bug where templates wouldn't parse correctly without inserting {literal} tags in large tpl files.
- Bug due to change in date_format tag specification
  - Previously returned null when passed null, but now returns the current date, causing inappropriate UI output.
    - Addressed by implementing a plugin to override the core date_format.

## PHP
Having been away from PHP 7 for about a year, I had forgotten some features and changes, leading to some learning moments.

There seems to be a tendency for ad-hoc responses. To address it more thoroughly, it might have been necessary to review data structures and adjust method appearances.

- Error when applying an empty index operator to an empty string
    - [php.net - PHP Manual Language Reference Types](https://www.php.net/manual/ja/language.types.string.php)
    - Incompatible change from PHP 7.1
- Error due to change in sorting algorithm
    - [php.net - Order of equal elements](https://www.php.net/manual/ja/migration70.incompatible.php#migration70.incompatible.other.sort-order)
    - Bugs occurred in areas dependent on order
      - → Addressed by using LaravelCollection to avoid dependency on sorting algorithm
- Addition of v to Datetime format for milliseconds
    - [php.net - Parameters](https://www.php.net/manual/ja/function.date.php#refsect1-function.date-parameters)
    - → Fixed tests failing due to millisecond differences
- E_WARNING when performing arithmetic with non-numeric values
    - [php.net - Notification for arithmetic with invalid strings](https://www.php.net/manual/ja/migration71.other-changes.php#migration71.other-changes.apprise-on-arithmetic-with-invalid-strings)
    - Change from PHP 7.1
    - Addressed by casting to int where arithmetic is needed
- Error when applying count() to non-array or non-object arguments
  - `count(null)` now results in a Warning
  - [php.net - Countable Interface](https://www.php.net/manual/ja/class.countable.php)
  - → Addressed by checking null with `isset` or `!empty`, or changing arguments to arrays or objects

# Infrastructure
- Wrote chef to migrate from existing PHP 5.6 environment to PHP 7.3 environment.
- Instance construction
    - Ran chef on a base instance to obtain an AMI. Constructed multiple instances from AMI launch.
        - Constructed a total of 20+ instances for staging and production
- Jenkins
  - Adjusted jobs as needed since deployment and release work uses Jenkins.

## Release
Handled with canary release.
Gradually attached PHP 7.3 instances to the target group with PHP 5.6 instances, running in parallel, increasing 7.3 instances while decreasing 5.6 instances.
Once fully switched, completed the switch (merged release branch into master).

For some errors investigated in the production environment, we utilized load balancer listener settings.
Specified internal IP as the source IP to route internal access to specific instances for verification (dark canary release?).

During the parallel operation period, errors frequently occurred, requiring detachment and reattachment of instances from the LB, taking two weeks to stabilize and fully switch.

# Challenges
Several bugs that couldn't be detected through tests or operational verification arose during the release phase.

Some were difficult to resolve or troublesome to address, but we quickly performed rollback operations (simply detaching instances from the LB) while dealing with error logs, taking nearly two weeks to stabilize.

From my perspective, the release work was more challenging than the upgrade work.

# Miscellaneous
I explored blogs and slides related to PHP version upgrades.

The areas requiring attention during upgrades vary by product, but I found the basic workflow quite similar.
GameWith's environment was very similar, which was quite emotional.

[5年以上PHP5で運用されていたFuelPHPで動くGameWithをPHP7.3にバージョンアップしました！ #GameWith #TechWith](https://tech.gamewith.co.jp/entry/2019/09/26/185515)
[テストコードが無くてPHP7へのバージョンアップが出来ない？ボットで解決しました！](https://speakerdeck.com/sgeengineer/tesutokodogawu-kutephp7hefalsebaziyonatupugachu-lai-nai-botutodejie-jue-simasita)
[3ヶ月でphp5.5から7.2にバージョンアップした現在と今後の向き合い方](https://speakerdeck.com/kosa3/3keyue-dephp5-dot-5kara7-dot-2nibaziyonatupusitaxian-zai-tojin-hou-falsexiang-kihe-ifang-number-phperkaigi-2019?slide=2)

# Performance Improvement
There was a significant improvement in CPU and memory usage, but no major change in response time.
Some areas haven't been properly measured yet, so investigation is ongoing.

# Impressions
I think manpower is effective for version upgrade work, so focusing on it might allow completion in a relatively short period.
It was a good experience.

While there was a miracle of FuelPHP 1.8.2 being suddenly released, there was also a rare event of facing a large-scale AWS outage towards the end of the retreat (which slightly affected operational verification in the staging environment).
