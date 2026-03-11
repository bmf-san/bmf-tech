---
title: Encountered Path Issues When Installing Bundler with rbenv via anyenv
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
 A discussion on the path issues encountered when installing bundler with rbenv installed via anyenv.

# The Issue
I installed rbenv using anyenv to use Ruby, but when I installed bundler with the command:

`gem install bundler`

without much thought, bundler ended up being placed under `/usr/local/bin/`.

Since this was not the intended path, attempts to use gems like chef would fail.

# Solution
`rbenv exec gem install bundler`

Specify to execute the gem with the Ruby introduced by rbenv.

# Thoughts
If I had calmly checked the path, I would have realized it quickly, even if I was not familiar with Ruby...

# References
- [github - rbenv/rbenv](https://github.com/rbenv/rbenv)
