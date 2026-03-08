---
title: Prompt Confirmation for Direct Push to Master
slug: master-push-prompt-confirmation
date: 2019-05-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Git
  - Shell Script
translation_key: master-push-prompt-confirmation
---

# Overview
How to create a safety net to prevent direct pushes to the master branch in Git. While it would be ideal to prohibit pushes to the master branch on GitHub, there were issues with the GitHub settings due to DevOps requirements, so we opted to configure it using hooks.

# Introduction
To set this globally, create a `pre-push` file under `~/.git_template/hooks`. If the `.git_template` directory does not exist, create it.

Note that even if you set it globally, it will not reflect in existing repositories. To apply it to existing repositories, you need to prepare a `pre-push` file under `./git/hooks` of the existing repository and write the same source there.

Refer to [git's pre-push hook to prompt confirmation when pushing to master](https://qiita.com/mkiken/items/af5c40ce0d0c6d3530c7) for the content to be written in `pre-push`.

It might be safer to limit the answer part to just 'yes'.

Make sure to give execution permissions to the newly created `pre-push`:
`chmod +x pre-push`

Setup is complete.

# References
- [git's pre-push hook to prompt confirmation when pushing to master](https://qiita.com/mkiken/items/af5c40ce0d0c6d3530c7)