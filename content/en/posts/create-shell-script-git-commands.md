---
title: Created a Shell Script to Simplify Git Commands
description: 'Build interactive shell scripts using select to manage git branch checkouts, remote branch pulls, and local branch deletion.'
slug: create-shell-script-git-commands
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Git
  - Shell Script
translation_key: create-shell-script-git-commands
---

# Overview
I created git alias commands in ~~.bashrc~~, but since it was incomplete, I improved it.

While the previous alias made using git commands somewhat comfortable, I thought it was necessary to improve the requirement of typing the branch name every time I executed a git command, so I solved it using `select`.

# Source
I thought it would be good to loop through the values of `git branch` with `select`, but since it also retrieves file names in addition to branch names, it needed some processing. I referred to an article that was doing something similar.

The commands I created are:

+ Checkout a local branch
+ Bring a remote branch to local
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

I wanted to change the color of the options output by `select`, but I didn't know how, so I postponed it.

# Thoughts
This source is available at [github - bmf-san/my-scripts](https://github.com/bmf-san/my-scripts).

# References
* [A tool to simplify switching git branches (checkout)](http://qiita.com/amichang/items/5f7e715801771214430e)