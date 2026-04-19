---
title: Getting Started with tmux
description: "Discover tmux terminal multiplexer essentials with keybindings for sessions, windows, panes, and copy mode operations."
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
An introduction to tmux, a terminal multiplexer.

# Key Bindings
Start tmux
`tmux` or `tmux new-session`

Create a new session within a session
`prefix+:new`

List sessions
`tmux ls`

Detach from a session (exit tmux)
`prefix+d`

Attach to a session
`tmux attach(a)`

Attach to a specific session
`tmux attach(a) -t 0(name)`

Delete a session
`tmux kill-session`

Delete a specific session
`tmux kill-session -t 0`

Delete all sessions
`tmux kill-server`

Rename a session
`prefix+$`

Create a new window
`prefix+c`

Switch to the next window
`prefix+n`

Switch to the previous window
`prefix+p`

Switch to a specific window
`prefix+0`

List windows
`prefix+w`

Delete a window
`prefix+&`

Delete a pane
`prefix+x`

Swap panes (forward)
`prefix+{`

Swap panes (backward)
`prefix+}`

Copy mode
`prefix+~~`

Select copy range (in copy mode)
`v` or `space`

Copy (in copy mode)
`y` or `enter`

# Tips
When using tmux in macOS Terminal, you cannot copy text by selecting it with the mouse (cmd+c) due to tmux's vim-like shortcuts. To copy the selected range with the mouse, you need to toggle the Terminal's Allow mouse reporting setting with `cmd+r`.

# References
- [github - bmf-san/dotfiles~~
  - Contains tmux configurations
