---
title: Current Status of FuelPHP as of March 2023
slug: fuelphp-status-march-2023
date: 2023-03-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - FuelPHP
translation_key: fuelphp-status-march-2023
---

# Overview
This post summarizes the current status of FuelPHP as of March 2023. The information is gathered from my own research, so it may lack accuracy.

While Laravel seems to be overwhelmingly popular among PHP frameworks, there may still be users actively using FuelPHP, and I hope this information can be of help.

# FuelPHP Version Information
Here is a summary of the information I found regarding FuelPHP.

|                           Version                            | Supported PHP Version |                                                                         Changelog                                                                          |                                                                                                    Note                                                                                                    |
| ------------------------------------------------------------ | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [1.8.0](https://github.com/fuel/fuel/tree/1.8.0)             | >=5.3.3               | [Changelog v1.8](https://github.com/fuel/core/wiki/Changelog-v1.8) [Changelog v1.8.0 hotfix1](https://github.com/fuel/core/wiki/Changelog-v1.8.0-hotfix-1) | Probably supports PHP up to 7.0                                                                                                                                                                       |
| [1.8.1](https://github.com/fuel/fuel/tree/1.8.1)             | >=5.3.3               | [Changelog v1.8.1](https://github.com/fuel/core/wiki/Changelog-v1.8.1)                                                                                     | Probably supports PHP up to 7.3                                                                                                                                                                                 |
| [1.8.2](https://github.com/fuel/fuel/tree/1.8.2)             | >=5.4                 | [Changelog v1.8.2](https://github.com/fuel/core/wiki/Changelog-v1.8.2)                                                                                     | According to [release information](https://fuelphp.com/blogs/2019/06/fuel-releases-1-8-2), it supports up to 7.3                                                                                                             |
| [1.9/develop](https://github.com/fuel/fuel/tree/1.9/develop) | >=5.4                 | N/A                                                                                                                                                        | Based on [Forum](https://fuelphp.com/forums/discussion/comment/28148#Comment_28148) and [commit](https://github.com/fuel/core/commit/ad982afacf66f56b756b805071be29d10a04dab6), it seems to support PHP 8.0. |
| 2.0 (private)                                                | N/A                   | N/A                                                                                                                                                        | According to [Forum](https://fuelphp.com/forums/discussion/comment/28148#Comment_28148), it is a redesigned version that supports PHP between 7.4 and 8.0.                                                    |

It seems that there is no clear indication of how long security support will last for each version of FuelPHP, and the EOL schedule is also unclear (it may be EOL for version 1).

The current latest version is FuelPHP 1.8.2, but the supported PHP version has already reached EOL.

# About FuelPHP 1.9/develop and 2.0
While the table above indicates support up to 8.0, it seems that there is also some level of support for 8.1.

[Forum](https://fuelphp.com/forums/discussion/15418/php-8-1/p1)

I have asked some questions to an insider (Harro) regarding this, and I will summarize the responses below. (The questions were asked around April 2022)
Please note that the original text was in English, so I may not have fully captured the nuances (my English skills are not very high).

Q. Is there any update information about FuelPHP 1.9?<br>
A. Our application (presumably the one Harro is working on) runs on PHP 8.1, so I think it is possible to use 1.9-dev.

Q. Does 1.9-dev run on PHP 8.1?<br>
A. We haven't thoroughly tested it on 8.1; we can only run tests on 8.x.

Q. Is there anything I can do to help with the release of 1.9-dev?<br>
A. We haven't tested the [PR](https://github.com/fuel/core/pull/2168), so we would appreciate help with that.
(From this, it seems that addressing this PR alone may not lead to a release, so I think they were looking for other tasks they could assist with.)

Q. Is financial support needed? Can donations help with FuelPHP development? (I was curious if it could be a means to expedite the release of 1.9/develop.)<br>
A. Not at this time. My (Harro's) company covers all the necessary infrastructure, and since all applications are developed with Fuel, bug fixes are implicitly supported. We are not lacking resources.

**While 1.9/develop is not an official release, it seems to be somewhat operational on PHP 8.1.**

I conducted my own personal research (upgrading the PHP and FuelPHP versions of applications running on FuelPHP) and found that it worked quite well.

However, the test coverage for 1.9/develop is quite low, which is concerning. But if you are using versions below 1.8.2 and absolutely need to upgrade your PHP version to 8.0 or 8.1, using 1.9/develop might be an option. (This is just my personal opinion. Use at your own risk.)

From the conversations, I learned that FuelPHP version 1 has already passed EOL. Perhaps for this reason, I heard that there is a policy not to focus on version 1 anymore.

Just a few days after that conversation, there were posts on [Forum](https://fuelphp.com/forums/discussion/15435/sneak-peak-#Item_3) and Twitter [post](https://twitter.com/fuelphp/status/1629873448460734468?s=20), indicating that the focus will shift from version 1 to version 2.0.

Previously, it seemed that the direction was to release 1.9/develop first because 2.0 would involve significant changes and take a considerable amount of time, especially since the EOL for PHP was approaching. However, it appears that my understanding was incorrect.

**It is unclear whether 1.9/develop will be officially released, but we might expect an alpha release of 2.0.**

Since 2.0 is still private and the source code is not public, the full picture is not yet clear. However, it seems that static interfaces will be removed and a DI container will be used, making it easier to write tests, so we might have some expectations regarding test coverage.

# Thoughts
I believe there are still many services in the world facing challenges on how to refresh applications that are locked into FuelPHP. I wonder what strategies they are taking... especially in cases where replacing with another application or architecture is not realistic and would involve significant pain, the choices available are likely quite limited...