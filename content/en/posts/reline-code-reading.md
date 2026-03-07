---
title: Code Reading of Reline
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
Conducting a code reading of Reline.

# Preparation
1. Clone Reline
`git@github.com:ruby/reline.git`
2. Create sample code
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
3. Insert `binding.irb` at any point.

# Code Reading
Reading the code based on the sample code.

1. Calling readline
- [ruby/reline/blob/master/lib/reline.rb#L251](https://github.com/ruby/reline/blob/master/lib/reline.rb#L251)
- [ruby/reline/blob/master/lib/reline.rb#L294](https://github.com/ruby/reline/blob/master/lib/reline.rb#L294)
2. Processing readline
- Omitted due to length
3. Output
- [ruby/reline/blob/master/lib/reline/line_editor.rb#L1323](https://github.com/ruby/reline/blob/master/lib/reline/line_editor.rb#L1323)
  - Finally, the input accumulated in the buffer is output here.