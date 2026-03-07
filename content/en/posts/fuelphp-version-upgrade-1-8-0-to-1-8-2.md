---
title: Upgraded FuelPHP from 1.8.0 to 1.8.2 and PHP from 5.6 to 7.3
slug: fuelphp-version-upgrade-1-8-0-to-1-8-2
date: 2019-10-05T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - PHP
  - FuelPHP
translation_key: fuelphp-version-upgrade-1-8-0-to-1-8-2
---

# Overview
Upgraded FuelPHP from 1.8.0 to 1.8.2 and PHP from 5.6 to 7.3. I summarized the efforts as we worked on the application upgrade for business purposes.

# Scope
- FuelPHP 1.8.0 → FuelPHP 1.8.2
- PHP 5.6 → PHP 7.3
- Target repositories
  - User-side application
  - Admin-side application
  - Package repository

*Middleware version details are omitted*
*OS is Amazon Linux (not version 2)*

FuelPHP 1.8.0 supports up to PHP 7.2, while 1.8.2 supports up to 7.3.

The release suddenly came about two weeks before we started the upgrade work.
[fuelphp.com - Fuel releases 1.8.2](https://fuelphp.com/blogs/2019/06/fuel-releases-1-8-2)

PHP 7.2's active support ended on November 30, 2019, and security support ended on November 30, 2020, while PHP 7.3's dates are December 6, 2020, and December 6, 2021.

I believe we should be very grateful for the release of FuelPHP 1.8.2, as it extends the support period.

# Duration
About a month and a half.

The first two weeks were spent working intensively away from the office in a retreat format. During this period, we had a code freeze, and in principle, releases other than emergency responses were halted.

# Team Composition
- 6 Engineers

# Why Upgrade?
The intention was more about security than performance requirements.

Not being able to ensure the product's security could pose a risk to the business.

# Work
## Preparation
- Caught up on upgrade information (changes) for FuelPHP and PHP. Read diffs and such.
- Searched for knowledge
  - Scoured GitHub for previous upgrade efforts, looking at issues and PRs that might be helpful.
  - Searched for useful documents from previous upgrade efforts and asked for guidance.
- Investigation (roughly identifying necessary fixes, checking the status of external libraries and middleware)
  - What fixes might be needed on the application side? What is the status of external libraries? What updates are needed for middleware?...etc
    - → I wanted to estimate the work but only did it roughly. I thought I could get a better estimate as I ran tests and confirmed operations, so I decided to keep it simple. In fact, the work went smoothly, so it might have been correct not to spend too much time on this.
- Formulated update work policies
  - Made a vow not to do refactoring unrelated to the upgrade.
    - → Since there seemed to be no discrepancies in understanding among the team, we did not establish strict rules.
  - If there were any points of confusion, I handled them as the project leader as needed.
- Created a schedule
  - Established a rough schedule for development work, QA, and release for the upgrade.
    - Rather than a detailed schedule, I focused on flexibly handling tasks that would arise, creating a GitHub project Kanban for task visualization and progress tracking.

## Branch Management Policy
Created a release branch for upgrade work derived from master. We set a 2-week code freeze period, but after the freeze was lifted, we allowed releases on the master branch. If there were functional modifications or additions to master, we rebased those changes into the release branch as needed.

## Progress
The approach for upgrading both FuelPHP and PHP was generally the same:

- First, raise the version
- Run tests and fix any failures
- Once everything is complete, conduct operational confirmation tests. Fix any failures found during operational confirmation one by one. *Note 1* *Note 2*

*Note 1: The operational confirmation tests utilized updated test items from previous upgrade efforts.*
*Note 2: After completing the FuelPHP upgrade, I conducted an operational confirmation, then upgraded PHP and conducted another confirmation.*
(At the stage where the FuelPHP confirmation was completed, it might have been okay to pause for a release.)

This was the flow of progress.

The difference was that for the PHP upgrade, it was necessary to prepare the staging environment and CI execution environment for PHP 7.3 in advance.

We also held morning meetings to share progress and any issues encountered, and we enjoyed lunch together. It was fun (just a casual remark).

Since there were a decent number of unit tests prepared, I don't think we faced hellish situations. (Testing is important.)

Bugs that arose after completing the work were mainly in areas not covered by tests (like fat controllers or areas that would likely be detected by E2E), and there were hardly any issues due to insufficient test cases.

# Changes Made During the Upgrade
I played a role similar to a project leader, managing progress and handling tasks while also doing actual work. However, I feel I spent more time on infrastructure work than on the development work for upgrading Fuel and PHP. (I might have mostly been involved in review-focused tasks for development work, although I did handle a few things.)

## FuelPHP
Unlike frameworks like Laravel, FuelPHP is not frequently updated, so I was somewhat skeptical about whether PHP 7.3 support was truly solid (due to the recent release, there was little track record), but it turned out to be fine.

- Changes in the session lifecycle broke login and logout.
    - → Changed affected areas to align with the new lifecycle specifications.
- A bug arose because the primary key could no longer be passed as an argument to the Model's forge method.
    - → Since there were many places passing the primary key as an argument and there was a need, I created a Model class for backward compatibility through inheritance.
- Changes to the ORM's to_array() specification
    - The order of values returned by to_array() is now guaranteed.
    - → Adjusted the test cases that were failing to align with the changes.
- Changes to the Pagination class specification
    - There was a change in the casting of values when obtaining the query string.
      - The pagination broke because it no longer cast to int.
        - → Adjusted to avoid passing unnecessary strings.
    - The existing functionality was impacted because it no longer urldecoded.
      - https://github.com/fuel/core/commit/4614349b243dc48e864a694b017d7356984d9f3c#diff-247dccd657c3e1348e064416f2b1e22bL201
      - → Added url_decode where necessary in the get() method.
- Impact on session retention due to changes in the internal encryption algorithm of Crypt
  - If you upgrade from 1.8.0 to 1.8.2 without any consideration, the login session becomes invalid, resulting in logout.
  - This is because the sodium library was standardized from PHP 7.2, and FuelPHP 1.8.2 supports it.
  - → Extended the core to allow the use of the 1.8.0 encryption method in the config.

## Smarty
There were changes that made me feel emotional, like "What was a minor update?"

Two issues arose during the minor update from v3.1.30 to v3.1.33.

- [Large template parsing error in smarty 3.1.33 #488](https://github.com/smarty-php/smarty/issues/488) 
    - Encountered a bug where the template parsing would not work correctly unless a {literal} tag was carelessly inserted, depending on the size of the tpl.
- A bug due to changes in the date_format tag specification
  - It used to return null when null was passed, but now it returns the current date and time, causing inappropriate values to be output in the UI.
    - Implemented a plugin to override the core date_format to address this.

## PHP
Having been away from PHP 7.x for about a year, I found I had forgotten some features and changes, leading to learning opportunities.
I feel I might have been somewhat haphazard in my responses. To respond more carefully, it might have been necessary to reconsider data structures and adjust method appearances.

- Applying an empty index operator to an empty string results in a fatal error.
    - [php.net - PHP Manual Language Reference Types](https://www.php.net/manual/ja/language.types.string.php)
    - A change that is not backward compatible since PHP 7.1.
- Errors due to changes in the sorting algorithm.
    - [php.net - Order of Equivalent Elements](https://www.php.net/manual/ja/migration70.incompatible.php#migration70.incompatible.other.sort-order)
    - Bugs appeared in areas dependent on order.
      - → Since I incorporated LaravelCollection, I used the collection to avoid dependency on the sorting algorithm.
- A 'v' representing milliseconds was added to the datetime format.
    - [php.net - Parameters](https://www.php.net/manual/ja/function.date.php#refsect1-function.date-parameters)
    - → Fixed tests that failed due to differences in milliseconds.
- An E_WARNING occurs when performing operations with non-numeric values.
    - [php.net - Notification of Arithmetic Operations with Invalid Strings](https://www.php.net/manual/ja/migration71.other-changes.php#migration71.other-changes.apprise-on-arithmetic-with-invalid-strings)
    - A change since PHP 7.1.
    - Addressed necessary operations by casting to int.
- Errors in applying non-array or non-object arguments to count().
  - `count(null)` also raises a warning.
  - [php.net - Countable Interface](https://www.php.net/manual/ja/class.countable.php) 
  - → Checked for null with `isset` or `!empty` and changed the arguments in the relevant areas to arrays or objects.

# Infrastructure
- Wrote Chef scripts to migrate from the existing PHP 5.6 environment to PHP 7.3.
- Instance construction
    - Ran Chef on a base instance to obtain an AMI. Built multiple instances from the AMI launch.
        - Constructed a total of over 20 instances for staging and production.
- Jenkins
  - Since we use Jenkins for deployment and release work, I adjusted jobs as necessary.

## Release
Handled with a canary release.
Gradually added PHP 7.3 instances to the target group where PHP 5.6 instances were hanging, running them in parallel, increasing the number of 7.3 instances while decreasing the number of 5.6 instances.
Once the switch was fully completed, we merged the release branch into master.

For some errors investigated in the production environment, we utilized the listener settings of the load balancer.
We specified the source IP as the internal IP and directed access from within the company to specific instances for verification. (Dark canary release?)

During the parallel operation period, errors frequently occurred, and we handled failures by disconnecting and reconnecting instances from the LB, stabilizing operations, which took about two weeks until complete switching.

# Challenges Faced
Several bugs occurred during the release phase that were not detected during testing or operational confirmation.

Some were difficult to resolve or troublesome to address, but when something happened, we quickly performed rollback operations (just disconnecting instances from the LB) while converging error logs, and it took nearly two weeks to stabilize operations.

From my perspective, the release work seemed more challenging than the upgrade work.

# Others
I searched for blogs and slides related to PHP version upgrades.

The areas requiring attention during upgrades likely vary by product, but I found that the basic workflow is quite similar.
GameWith's environment is very similar, which made me feel quite emotional.

[Upgraded GameWith running on FuelPHP, which had been operating on PHP 5 for over 5 years, to PHP 7.3! #GameWith #TechWith](https://tech.gamewith.co.jp/entry/2019/09/26/185515)
[Can't upgrade to PHP 7 due to lack of test code? Solved with a bot!](https://speakerdeck.com/sgeengineer/tesutokodogawu-kutephp7hefalsebaziyonatupugachu-lai-nai-botutodejie-jue-simasita)
[How we upgraded from PHP 5.5 to 7.2 in 3 months and our future approach](https://speakerdeck.com/kosa3/3keyue-dephp5-dot-5kara7-dot-2nibaziyonatupusitaxian-zai-tojin-hou-falsexiang-kihe-ifang-number-phperkaigi-2019?slide=2)

# Performance Improvement
There was a significant improvement in CPU and memory usage, but no major changes were observed in response times.
There are still parts that haven't been properly measured, so investigations are ongoing.

# Impressions
I believe that a large part of the upgrade work benefits from a manpower approach, so focusing on it might allow for relatively quick completion.
It was a valuable experience.

While there was a miraculous sudden release of FuelPHP 1.8.2, we also faced a rare event of a large-scale AWS outage towards the end of the retreat (which somewhat affected operational verification in the staging environment).