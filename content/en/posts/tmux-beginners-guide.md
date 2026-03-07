---
title: Getting Started with tmux
slug: tmux-beginners-guide
date: 2018-05-22T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - tmux
translation_key: tmux-beginners-guide
---

# Overview
Getting started with tmux, a terminal multiplexer software.

# Key Bindings
tmux launch
tmux or tmux new-session

Create a new session within a session
prefix+:new

Session list
tmux ls

Detach from session (exit tmux)
prefix+d

Attach to session
tmux attach(a)

Attach to a specific session
tmux attach(a) -t 0(name)

Delete session
tmux kill-session

Delete a specific session
tmux kill-session -t 0

Delete all sessions
tmux kill-server

Rename session
prefix+$

New window
prefix+c

Switch to next window
prefix+n

Switch to previous window
prefix+p

Switch to a specific window
prefix+0

Window list
prefix+w

Delete window
prefix+&

Delete pane
prefix+x

Swap panes (forward)
prefix+{

Swap panes (backward)
prefix+}

Copy mode
prefix+[

Select copy range (in copy mode)
v or space

Copy (in copy mode)
y or enter

# Tips
When using tmux in the Mac terminal, you cannot copy (cmd+c) by selecting text with the mouse. (You can use tmux shortcuts for vim-like copying.)
If you want to copy the selected range with the mouse, you need to toggle the terminal's Allow mouse reporting setting with cmd+r.

# References
- [github - bmf-san/dotfiles](https://github.com/bmf-san/dotfiles)
  - Contains tmux configuration.