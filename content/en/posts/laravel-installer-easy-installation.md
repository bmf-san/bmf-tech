---
title: Easy Laravel Installation with Installer
description: A step-by-step guide on Easy Laravel Installation with Installer, with practical examples and configuration tips.
slug: laravel-installer-easy-installation
date: 2016-05-15T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Laravel
translation_key: laravel-installer-easy-installation
---



When installing Laravel, I used to type tedious commands with composer every time, but using the installer was much easier (better late than never).


# Preparation
- composer
- MacOS


# Install the Installer Globally

`composer global require "laravel/installer"`


# Set the Path

If you're on MacOS, this should work. (I don't know about Windows...)

`export PATH="~/.composer/vendor/bin:$PATH"`

# Create a New Project

`laravel new PROJECTNAME`

The latest version of Laravel will be installed in the current directory.

As mentioned in the documentation, it **runs faster than using composer**.


# Thoughts

Fast and easy.

# Additional Notes
Add `export PATH="~/.composer/vendor/bin:$PATH"` to the .bash_profile in your MacOS home directory.
If .bash_profile doesn't exist, create it. If you want to know the difference from .bashrc, Google it.
