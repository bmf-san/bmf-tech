---
title: Installing Previous Versions with Homebrew
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
Sometimes you want to specify and install a previous version of a package installed with Homebrew. It seems that Homebrew has adopted a policy of only keeping the latest version, so installing a previous version requires a bit of extra work, hence this note.

# Method
In this case, I wanted to downgrade from vim9 to vim8, so I will use that process as an example.

The steps are as follows:
```sh
brew tap-new bmf-san/vim8
brew extract vim bmf-san/vim8 --version 8.2.5150
brew install bmf-san/vim8/vim@8.2.5150
brew unlink vim
brew link vim@8.2.5150
vim --version // should show 8.2.5150
```

You need to prepare your own repository for tapping. You can use a GitHub repository, but you can also use the `tap-new` command, so I opted for that. The name is arbitrary.

`brew tap-new bmf-san/vim8`

To extract the old formula to the tap, you need to use `extract`. The version you specify might be anything, but it seems safer to fetch a version that was previously managed by Homebrew, so I searched for the commit of the version I wanted in the Homebrew repository. The commit I found is as follows:

[github.com/Homebrew/homebrew-core/blob/2e3d51340f2f9b47d680f656e712fbee77cbcf79/Formula/vim.rb](https://github.com/Homebrew/homebrew-core/blob/2e3d51340f2f9b47d680f656e712fbee77cbcf79/Formula/vim.rb)

`brew extract vim bmf-san/vim8 --version 8.2.5150`

Install from the tap and adjust the symbolic link with link & unlink.

`brew install bmf-san/vim8/vim@8.2.5150`

`brew unlink vim`

`brew link vim@8.2.5150`

# Thoughts
The reason I wanted to downgrade vim was that I sensed that vim-lsp does not seem to work with version 9...

# References
- [christina04.hatenablog.com - Want to Use Previous Versions with Homebrew [tap version]](https://christina04.hatenablog.com/entry/install-old-version-with-homebrew)