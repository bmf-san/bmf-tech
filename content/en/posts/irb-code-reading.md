---
title: Code Reading of irb
description: 'A review and summary of "Code Reading of irb", covering key takeaways and practical insights.'
slug: irb-code-reading
date: 2024-10-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - irb
  - Ruby
translation_key: irb-code-reading
---

# Overview
Reading the code of irb.

# Preparation
1. Clone the irb source code
- `git clone git@github.com:ruby/irb.git`
2. Modify the entry point
Since the original state runs the irb located in the Ruby installation directory, we will run the local irb instead.
```ruby
#!/usr/bin/env ruby
#
#   irb.rb - interactive ruby
#    by Keiju ISHITSUKA(keiju@ruby-lang.org)
#

$LOAD_PATH.unshift(File.expand_path("../lib", __dir__)) # Add this line
require "irb"

IRB.start(__FILE__)
```
3. Debug at any point

# Code Reading
1. Executing the irb command
- [ruby/irb/blob/master/exe/irb#L9](https://github.com/ruby/irb/blob/master/exe/irb#L9)
  - Executes `IRB.start(__FILE__)` with the filename of the executing file as an argument.
2. Setup process at startup
- [ruby/irb/blob/master/lib/irb.rb#L895](https://github.com/ruby/irb/blob/master/lib/irb.rb#L895)
3. Starting irb
- [ruby/irb/blob/master/lib/irb.rb#L1001](https://github.com/ruby/irb/blob/master/lib/irb.rb#L1001)
  - Loads the history
    - History is the execution history of commands. It is written to a file (~/.irb_history).
4. Executing input values
- [ruby/irb/blob/master/lib/irb.rb#L1041](https://github.com/ruby/irb/blob/master/lib/irb.rb#L1041)
- [ruby/irb/blob/master/lib/irb/context.rb#L589](https://github.com/ruby/irb/blob/master/lib/irb/context.rb#L589)
  - Processes are branched and executed based on the input values.
  - For expressions:
    - [ruby/irb/blob/master/lib/irb/context.rb#L589](https://github.com/ruby/irb/blob/master/lib/irb/context.rb#L589)
    - [ruby/irb/blob/master/lib/irb/context.rb#L609](https://github.com/ruby/irb/blob/master/lib/irb/context.rb#L609)
    - [ruby/irb/blob/master/lib/irb/workspace.rb#L120](https://github.com/ruby/irb/blob/master/lib/irb/workspace.rb#L120)
      - Evaluates the expression with `eval`.
  - For commands:
    - [ruby/irb/blob/master/lib/irb/command/base.rb#L55](https://github.com/ruby/irb/blob/master/lib/irb/command/base.rb#L55)
      - It was unclear what the command was, but it does not seem to be a feature for users.