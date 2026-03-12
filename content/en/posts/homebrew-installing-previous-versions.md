---
title: Installing Older Versions with Homebrew
slug: homebrew-installing-previous-versions
date: 2022-10-30T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - homebrew
  - vim
translation_key: homebrew-installing-previous-versions
---

# Overview
Sometimes, you may want to install an older version of a package using Homebrew. Homebrew has adopted a policy of keeping only the latest versions, so installing older versions requires a bit of extra work, which I've noted here.

# Method
I wanted to downgrade from vim9 to vim8, so I'll use that as an example.

Here are the steps:
```sh
brew tap-new bmf-san/vim8
brew extract vim bmf-san/vim8 --version 8.2.5150
brew install bmf-san/vim8/vim@8.2.5150
brew unlink vim
brew link vim@8.2.5150
vim --version // Should be 8.2.5150
```

You need to prepare your own repository for the tap. You can use a GitHub repository, but since there's a command called `tap-new`, I'll use that. The name can be anything.

`brew tap-new bmf-san/vim8`

To extract the old formula into the tap, use extract. The version specified can be anything, but it's safer to use a version that Homebrew managed in the past, so I searched for the commit of the desired version in the Homebrew repository. The commit I found is below.

[github.com/Homebrew/homebrew-core/blob/2e3d51340f2f9b47d680f656e712fbee77cbcf79/Formula/vim.rb](https://github.com/Homebrew/homebrew-core/blob/2e3d51340f2f9b47d680f656e712fbee77cbcf79/Formula/vim.rb)

`brew extract vim bmf-san/vim8 --version 8.2.5150`

Install from the tap and adjust the symbolic links with link & unlink.

`brew install bmf-san/vim8/vim@8.2.5150`

`brew unlink vim`

`brew link vim@8.2.5150`

# Thoughts
I wanted to downgrade vim because it seemed like vim-lsp wasn't working properly with version 9...

# References
- [christina04.hatenablog.com - Using Older Versions with Homebrew [tap version]](https://christina04.hatenablog.com/entry/install-old-version-with-homebrew)