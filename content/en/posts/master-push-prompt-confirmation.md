---
title: Prompt Confirmation When Pushing Directly to Master
slug: master-push-prompt-confirmation
date: 2019-05-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Git
  - Shell Script
description: How to create a safety net to prevent direct pushes to the master branch in git.
translation_key: master-push-prompt-confirmation
---

# Overview
How to create a safety net to prevent direct pushes to the master branch in git. Although it would be ideal to prohibit pushes to the master branch on GitHub, due to DevOps constraints, there were issues with GitHub settings, so we opted to use hooks for configuration.

# Introduction
Since we want to set this globally, create a `pre-push` file under `~/.git_template/hooks`. If the `.git_template` directory does not exist, create it.

Note that even if you set it globally, it will not be reflected in existing repositories. If you want to apply it to existing repositories, you need to prepare a `pre-push` under `./git/hooks` in the existing repository and write the same source there.

Refer to [Prompt Confirmation When Pushing to Master Branch with git pre-push Hook](https://qiita.com/mkiken/items/af5c40ce0d0c6d3530c7) for what to write in `pre-push`.

It might be safer to only allow 'yes' as the answer.

Grant execution permissions to the newly created `pre-push`
`chmod +x pre-push`

Setup is now complete.

# References
- [Prompt Confirmation When Pushing to Master Branch with git pre-push Hook](https://qiita.com/mkiken/items/af5c40ce0d0c6d3530c7)