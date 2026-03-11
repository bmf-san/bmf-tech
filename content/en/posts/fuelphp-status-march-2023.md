---
title: Current Status of FuelPHP as of March 2023
slug: fuelphp-status-march-2023
date: 2023-03-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - FuelPHP
description: A summary of the current status of FuelPHP as of March 2023.
translation_key: fuelphp-status-march-2023
---

# Overview
This is a summary of the current status of FuelPHP as of March 2023, based on my own research. The information may lack accuracy as it is independently gathered.

While Laravel seems to be overwhelmingly popular among PHP frameworks, there might still be people actively using FuelPHP. I hope this information can be of some help.

# FuelPHP Version Information
Here is a summary of the information I gathered about FuelPHP.

|                           Version                            | Supported PHP Version |                                                                         Changelog                                                                          |                                                                                                    Notes                                                                                                    |
| ------------------------------------------------------------ | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [1.8.0](https://github.com/fuel/fuel/tree/1.8.0)             | >=5.3.3               | [Changelog v1.8](https://github.com/fuel/core/wiki/Changelog-v1.8) [Changelog v1.8.0 hotfix1](https://github.com/fuel/core/wiki/Changelog-v1.8.0-hotfix-1) | It seems to support PHP up to 7.0                                                                                                                                                                       |
| [1.8.1](https://github.com/fuel/fuel/tree/1.8.1)             | >=5.3.3               | [Changelog v1.8.1](https://github.com/fuel/core/wiki/Changelog-v1.8.1)                                                                                     | It seems to support PHP up to 7.3                                                                                                                                                                                 |
| [1.8.2](https://github.com/fuel/fuel/tree/1.8.2)             | >=5.4                 | [Changelog v1.8.2](https://github.com/fuel/core/wiki/Changelog-v1.8.2)                                                                                     | According to [release information](https://fuelphp.com/blogs/2019/06/fuel-releases-1-8-2), it supports up to 7.3                                                                                                             |
| [1.9/develop](https://github.com/fuel/fuel/tree/1.9/develop) | >=5.4                 | N/A                                                                                                                                                        | According to [Forum](https://fuelphp.com/forums/discussion/comment/28148#Comment_28148) and [commit](https://github.com/fuel/core/commit/ad982afacf66f56b756b805071be29d10a04dab6), it seems to support PHP 8.0. |
| 2.0 (private)                                                | N/A                   | N/A                                                                                                                                                        | According to [Forum](https://fuelphp.com/forums/discussion/comment/28148#Comment_28148), it is a redesigned version that supports PHP 7.4 and above but below 8.0                                                    |

It seems that there is no clear statement on how long security support will be provided for each version of FuelPHP, or any planned EOL (End of Life) dates (version 1.x might already be EOL).

The current latest version is FuelPHP 1.8.2, but the PHP versions it supports have already reached EOL.

# About FuelPHP 1.9/develop and 2.0
In the table above, I mentioned that it seems to support up to 8.0, but it appears that some support for 8.1 is also in place.

[Forum](https://fuelphp.com/forums/discussion/15418/php-8-1/p1)

I have asked some questions to an insider (Harro) about this, and I will summarize the responses below. (The questions were asked around April 2022)
Please note that the original text is in English, so there might be nuances that I couldn't fully capture (my English skills are not high, so please bear with me).

Q. Is there any update information about FuelPHP 1.9?<br>
A. Our application (presumably the application Harro handles at work) is running on PHP 8.1, so I think it is possible to use 1.9-dev.

Q. Does 1.9-dev work on PHP 8.1?<br>
A. We haven't been able to fully verify its operation on 8.1, and have only been able to run tests on version 8.

Q. Is there anything I can do to help with the release of 1.9-dev?<br>
A. We haven't been able to test [PR](https://github.com/fuel/core/pull/2168), so we would appreciate your help.
(It seems that addressing this PR alone won't make it releasable, so I think they suggested tasks I could help with.)

Q. Do you need financial support? Can donations help with the development of FuelPHP? (I was curious if it could expedite the release of 1.9/develop, so I asked)<br>
A. It is not necessary at the moment. My (Harro's) company provides all the necessary infrastructure, and since all applications are developed with Fuel, bug fixes are implicitly supported. We are not lacking resources.

**1.9/develop is not an officially released version, but it seems to be somewhat operational on PHP 8.1.**

I conducted a personal investigation (upgrading the PHP and FuelPHP versions of applications running on FuelPHP and verifying their operation), and they worked quite well.

The low test coverage of 1.9/develop is a bit concerning, but if you are using a version below 1.8.2 and absolutely need to upgrade the PHP version to 8.0-8.1, using 1.9/develop might be an option. (This is a personal opinion. Please proceed at your own risk.)

Through the conversation, I learned that FuelPHP 1.x is apparently already past EOL. As a result, I also heard that there is a policy not to focus on version 1.x anymore.

Shortly after having such a conversation, there was a [Forum](https://fuelphp.com/forums/discussion/15435/sneak-peak-#Item_3) post and a [Twitter post](https://twitter.com/fuelphp/status/1629873448460734468?s=20), revealing that the focus is on version 2.0 rather than version 1.x.

Previously, I thought the direction was "2.0 will be a major change, so it is taking quite some time, and since PHP's EOL has arrived, we will release 1.9/develop first," but it seems that understanding was incorrect.

**It is unclear whether 1.9/develop will be officially released, but we might be able to look forward to an alpha release of 2.0.**

Version 2.0 is still private, and the source code is not public, so the full picture is unknown. However, it seems that static interfaces will be removed, and a DI container will be used, making it easier to write tests, so we might be able to have expectations regarding test coverage as well.

# Impressions
I believe there are still many services in the world facing challenges on how to revamp applications locked into FuelPHP. I wonder what strategies various places are adopting...
Especially in cases where replacing with a different application or architecture is not realistic and involves significant pain, the choices to be made should be quite limited...