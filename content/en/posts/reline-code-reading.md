---
title: Code Reading of reline
description: 'A review and summary of "Code Reading of reline", covering key takeaways and practical insights.'
slug: reline-code-reading
date: 2024-10-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - reline
  - Ruby
translation_key: reline-code-reading
---



# Overview
Conducting a code reading of reline.

# Preparation
1. Clone reline
`git@github.com:ruby/reline.git`
2. Create a sample code
```ruby
$LOAD_PATH.unshift(File.expand_path('lib', __dir__))
require "reline"

prompt = 'prompt> '
use_history = true

begin
  while true
    text = Reline.readmultiline(prompt, use_history) do |multiline_input|
      # Accept the input until `end` is entered
      multiline_input.split.last == "end"
    end

    puts 'You entered:'
    puts text
  end
# If you want to exit, type Ctrl-C
rescue Interrupt
  puts '^C'
  exit 0
end
```
3. Insert `binding.irb` at any desired point

# Code Reading
Conduct code reading based on the sample code.

1. Call of readline
- [ruby/reline/blob/master/lib/reline.rb#L251](https://github.com/ruby/reline/blob/master/lib/reline.rb#L251)
- [ruby/reline/blob/master/lib/reline.rb#L294](https://github.com/ruby/reline/blob/master/lib/reline.rb#L294)
2. Processing of readline
- Omitted due to length
3. Output
- [ruby/reline/blob/master/lib/reline/line_editor.rb#L1323](https://github.com/ruby/reline/blob/master/lib/reline/line_editor.rb#L1323)
  - The input accumulated in the buffer is finally output here
