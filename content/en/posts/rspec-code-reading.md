---
title: Reading RSpec Code
slug: rspec-code-reading
date: 2024-10-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - RSpec
  - Ruby
translation_key: rspec-code-reading
---

# Overview
Reading the RSpec code.

# Preparation
1. Clone the RSpec repository.
- https://github.com/rspec/rspec-core

# Code Reading
1. RSpec Invocation
- [rspec/rspec-core/blob/main/exe/rspec#L4](https://github.com/rspec/rspec-core/blob/main/exe/rspec#L4)
  - Entry point
- [lib/rspec/core/runner.rb#L43](https://github.com/rspec/rspec-core/blob/main/lib/rspec/core/runner.rb#L43)
  - Calls the class method invoke of the Runner class
  - The `disable_autorun` method disables the auto-run feature
3. Running RSpec
- [lib/rspec/core/runner.rb#L64](https://github.com/rspec/rspec-core/blob/main/lib/rspec/core/runner.rb#L64)
  - Calls the class method run of the Runner
  - The `trap_interrupt` method handles interruptions like Ctrl+C
  - If there is a runner in options, `call`, otherwise `new.run`
- [lib/rspec/core/runner.rb#L85](https://github.com/rspec/rspec-core/blob/main/lib/rspec/core/runner.rb#L85)
  - If the test execution ends early, calls the reporting process
  - If not ending early, calls the `run_specs` method
- [lib/rspec/core/runner.rb#L113](https://github.com/rspec/rspec-core/blob/main/lib/rspec/core/runner.rb#L113)
  - Executes tests sequentially, reports results, and returns the exit code