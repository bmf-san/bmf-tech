---
title: Code Reading of irb
slug: irb-code-reading
date: 2024-10-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - irb
  - Ruby
description: A guide to reading and understanding the source code of irb.
translation_key: irb-code-reading
---

# Overview
Perform a code reading of irb.

# Preparation
1. Clone the source code of irb
   - `git clone git@github.com:ruby/irb.git`
2. Modify the entry point
   By default, it executes the irb located in the Ruby installation directory. Modify it to execute the local irb instead.
   ```ruby
   #!/usr/bin/env ruby
   #
   #   irb.rb - interactive ruby
   #       by Keiju ISHITSUKA(keiju@ruby-lang.org)
   #

   $LOAD_PATH.unshift(File.expand_path("../lib", __dir__)) # Add this line
   require "irb"

   IRB.start(__FILE__)
   ```
3. Debug at any desired point

# Code Reading
1. Execution of the irb command
   - [ruby/irb/blob/master/exe/irb#L9](https://github.com/ruby/irb/blob/master/exe/irb#L9)
     - Executes `IRB.start(__FILE__)` with the filename as an argument
2. Setup process during startup
   - [ruby/irb/blob/master/lib/irb.rb#L895](https://github.com/ruby/irb/blob/master/lib/irb.rb#L895)
3. Starting irb
   - [ruby/irb/blob/master/lib/irb.rb#L1001](https://github.com/ruby/irb/blob/master/lib/irb.rb#L1001)
     - Loads the history
       - History refers to the command execution history, which is written to a file (`~/.irb_history`).
4. Execution of input values
   - [ruby/irb/blob/master/lib/irb.rb#L1041](https://github.com/ruby/irb/blob/master/lib/irb.rb#L1041)
   - [ruby/irb/blob/master/lib/irb/context.rb#L589](https://github.com/ruby/irb/blob/master/lib/irb/context.rb#L589)
     - Branches and executes based on the input value
     - In the case of an expression:
       - [ruby/irb/blob/master/lib/irb/context.rb#L589](https://github.com/ruby/irb/blob/master/lib/irb/context.rb#L589)
       - [ruby/irb/blob/master/lib/irb/context.rb#L609](https://github.com/ruby/irb/blob/master/lib/irb/context.rb#L609)
       - [ruby/irb/blob/master/lib/irb/workspace.rb#L120](https://github.com/ruby/irb/blob/master/lib/irb/workspace.rb#L120)
         - Evaluates the expression using `eval`
     - In the case of a command:
       - [ruby/irb/blob/master/lib/irb/command/base.rb#L55](https://github.com/ruby/irb/blob/master/lib/irb/command/base.rb#L55)
         - It was unclear what the command is, but it does not seem to be a user-facing feature.