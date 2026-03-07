---
title: Easy Installation of Laravel with Installer
slug: laravel-installer-easy-installation
date: 2016-05-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
translation_key: laravel-installer-easy-installation
---

When installing Laravel, I used to type cumbersome commands with composer every time, but using the installer is much easier (better late than never).

# Preparation
- composer
- MacOS

# Global Installation of Installer

`composer global require "laravel/installer"`

# Set the Path

I think this will work on MacOS. (I don't know about Windows...)

`export PATH="~/.composer/vendor/bin:$PATH"`

# Create a New Project

`laravel new PROJECTNAME`

The latest version of Laravel will be installed in the current directory.

As mentioned in the documentation, it seems to run **faster than using composer**.

# Impressions

Fast and easy.

# Additional Information

Add `export PATH="~/.composer/vendor/bin:$PATH"` to the .bash_profile in your MacOS home directory. If you don't have a .bash_profile, create one. If you want to know the difference from .bashrc, just Google it.