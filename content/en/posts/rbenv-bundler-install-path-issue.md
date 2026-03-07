---
title: Path Issues When Installing Bundler with rbenv Installed via anyenv
slug: rbenv-bundler-install-path-issue
date: 2018-12-04T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - anyenv
  - gem
  - rbenv
  - Ruby
  - Tips
translation_key: rbenv-bundler-install-path-issue
---

# Overview
I encountered a path issue when installing bundler with rbenv installed via anyenv.

# The Issue
I am using ruby installed with rbenv via anyenv, but when I tried to install bundler by simply running:

`gem install bundler`

it ended up placing bundler in `/usr/local/bin/`.

Since this is not the intended path, trying to use gems like chef installed via gem results in errors.

# Solution
Use the following command to specify that you want to execute the gem for the ruby installed with rbenv:

`rbenv exec gem install bundler`

# Thoughts
If I had calmly checked the path, I would have realized this quickly even as a beginner with ruby...

# References
- [github - rbenv/rbenv](https://github.com/rbenv/rbenv)