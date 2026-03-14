---
title: Created Git Shortcut Commands in .bashrc
description: 'Create powerful git aliases and bash functions for branch checkout, pull, and push operations using .bashrc configuration.'
slug: create-git-shortcut-bashrc
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - bash
  - Git
  - shell script
translation_key: create-git-shortcut-bashrc
---

git add hogehoge, git commit hogehoge, git push hogehoge....

I only use basic git commands, but typing them every time is cumbersome, so I decided to create some aliases.

# Script

```vim
#git branch
alias git-b='git branch'

#git checkout
function gitCheckout() {
         stty erase ^H
         echo -n "What is the new branch name"?
         stty echo
         read var1
         git checkout ${var1}
}
alias git-c=gitCheckout

#git checkout -b
function gitCheckoutBranch() {
         stty erase ^H
         echo -n "What is the new branch name for checkout?"
         stty echo
         read var1
         git checkout -b ${var1}
}
alias git-c-b=gitCheckoutBranch

#git pull
function gitPull() {
        stty erase ^H
        echo -n "What is the remote repository name?"
        stty echo
        read var1
        git pull origin ${var1}
}
alias git-p=gitPull

#git set
function gitSet() {
      stty erase ^H
      echo -n "What file name do you add?"
      read var1
      git add ${var1}
      echo -n "What is the commit message?"
      read var2
      git commit -m'${var2}'
      echo -n "What is the branch name?"
      stty echo
      read var3
      git push origin ${var3}
}
alias git-set=gitSet
```

# Command Descriptions
* git-b ・・・Check branches
* git-c ・・・Checkout
* git-c-b ・・・Create and checkout a new branch
* git-p ・・・Pull
* git-set ・・・Interactively perform add/commit/push. Please correct me as my English sounds awkward.

# Troubleshooting
* .bashrc is not being loaded
  * Check if .bash_profile loads .bashrc
     * cf. [[bash] Things to check when .bashrc is not loaded after starting the terminal](http://programming-log.tumblr.com/post/102419333247/bash-%E3%82%BF%E3%83%BC%E3%83%9F%E3%83%8A%E3%83%AB%E3%82%92%E8%B5%B7%E5%8B%95%E3%81%97%E3%81%A6%E3%82%82bashrc%E3%81%8C%E8%AA%AD%E3%81%BF%E8%BE%BC%E3%81%BE%E3%82%8C%E3%81%A6%E3%81%84%E3%81%AA%E3%81%84%E6%99%82%E3%81%AB%E7%A2%BA%E8%AA%8D%E3%81%99%E3%82%8B%E3%81%93%E3%81%A8)

* The backspace (delete) key is converted to strange characters after executing alias commands
  * Keyboard is US layout → Keep ^H in .bashrc
  * JIS → Change ^H to ^?

# Thoughts
* Development efficiency has slightly improved!
* How do people who use git usually do it? Do they create aliases?

# Addendum
I forgot there is a command called `git config` for setting git aliases.
[Set aliases for frequently used git commands to improve development efficiency](http://qiita.com/unsoluble_sugar/items/ce14e9ce20aa5ba34fe5)