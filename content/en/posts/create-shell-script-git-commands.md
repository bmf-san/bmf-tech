---
title: Wrote a Shell Script to Simplify Git Commands
slug: create-shell-script-git-commands
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Git
  - Shell Script
description: Improved a previous attempt at creating Git alias commands by using `select` for better usability.
translation_key: create-shell-script-git-commands
---

# Overview
I previously created [shortcut Git commands in .bashrc](http://qiita.com/bmf_san/items/d41bc8d5c5677c69a1e4), but it was incomplete, so I made improvements.

While the previous aliases made Git commands somewhat more convenient, the need to type the branch name every time I ran a Git command was a point that needed improvement. I resolved this issue by using `select`.

# Source
I initially thought I could simply use `select` to loop through the output of `git branch`, but it also included file names and other unnecessary data, so I had to process it. I found an article doing something similar and used it as a reference.

The commands I created are:

+ Checkout a local branch
+ Fetch a remote branch to the local environment
+ Delete a local branch

```sh
#!/bin/sh

#checkout a local branch
function gitCheckoutLocalBranch() {
        branches=`git branch | grep -v -e"^\*" | tr -d ' '`

        PS3="Select branch > "
        echo 'Branch list:'

        select branch in ${branches}
        do
                stty erase ^H
                git checkout ${branch}
                break
        done
}
alias g-c=gitCheckoutLocalBranch

# create a new branch and checktout a remote branch
function gitCreateAndCheckoutRemoteBranch() {
        branches=`git branch -r | grep -v -e"^\*" | tr -d ' '`

        PS3="Select branch > "
        echo 'Branch list:'

        select branch in ${branches}
        do
                stty erase ^H
                echo -n "What is the new branch name?"
                read new_branch_name
                git checkout -b ${new_branch_name} ${branch}
                break
        done
}
alias g-c-b-r=gitCreateAndCheckoutRemoteBranch

# delete a local branch
function gitDeleteLocalBranch() {
        branches=`git branch | grep -v -e"^\*" | tr -d ' '`

        PS3="Select branch > "
        echo 'Branch list:'

        select branch in ${branches}
        do
                stty erase ^H
                git branch -D ${branch}
                break
        done
}
alias g-b-d=gitDeleteLocalBranch
```

I wanted to change the color of the options displayed by `select`, but I couldn't figure it out, so I left it for later.

# Thoughts
The source code for this script is available on [GitHub - bmf-san/my-scripts](https://github.com/bmf-san/my-scripts).

# References
* [A tool to make Git branch switching (checkout) easier](http://qiita.com/amichang/items/5f7e715801771214430e)