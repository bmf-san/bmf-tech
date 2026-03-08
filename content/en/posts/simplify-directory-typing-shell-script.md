---
title: Got Tired of Typing Directories with the cd Command, So I Made It Easier with a Shell Script
slug: simplify-directory-typing-shell-script
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - bash
  - shell script
translation_key: simplify-directory-typing-shell-script
---



# Overview
I started feeling lazy about typing paths and directories like `cd hogehoge`, so I used a shell script to make it a bit easier.

# Source

```bash
#!/bin/sh

# cd by selecting numbers
function cdSelect() {
        dirs=`ls -a`

        PS3="Select directory > "
        echo 'Directory list:'

        select dir in ${dirs}
        do
                stty erase ^H
                cd ${dir}
                break
        done
}
alias cd-s=cdSelect
```

When you type `cd-s`, it looks like this:

```
Directory list:

1) .
2) ..
3) hoge_a
4) hoge_b
5) hoge_c

Select directory > 3
```

# Thoughts
It might be challenging when there are many directories, but it has reduced the stress of using the cd command. I'm considering making a vim version too.
