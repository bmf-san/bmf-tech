---
title: Revisiting the Development Environment to Make it Vim-like
slug: vim-like-development-environment
date: 2018-05-22T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - tmux
  - vim
  - Atom
  - iTerm
translation_key: vim-like-development-environment
---

# Overview
To improve development efficiency, I incorporated Vim and revamped my development environment. I will omit detailed settings for each tool and the specifics of the plugins used.

# Editor
- Atom
	- Main editor used for development
	- Introduced plugins to enable Vim keybindings
		- vim-mode-plus-ex-mode
		- vim-mode-plus
		- Adjusted keybindings slightly to allow for screen splitting and pane movement similar to Vim

- Vim
	- Used as a secondary editor for minor edits
	- Only the minimum necessary plugins are installed

# Command Line Tools
- iTerm2
	- Enabled vi-mode in bash
		- Allows the use of vi keybindings in the command line

# Terminal Multiplexer Software
- tmux
	- Might be essential for heavy command line usage

# References
- [github - bmf-san/my-dotfiles](https://github.com/bmf-san/my-dotfiles)
	- Current settings for Atom and Vim

# Thoughts
By embracing the Vim mindset and applying Vim keybindings across various tools, I feel a sense of happiness.