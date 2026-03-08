---
title: Using Shell Script to Avoid Typing Directory with cd Command
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
Typing paths and directories like `cd hogehoge` has become tedious, so I decided to use a shell script to make it a bit easier.

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

When you type `cd-s`,

```
Directory list:

1) .
2) ..
3) hoge_a
4) hoge_b
5) hoge_c

Select directory > 3
```

It looks like this.

# Thoughts
It seems cumbersome when there are many directories, but it has reduced the stress of using the cd command. I might create a vim version as well.