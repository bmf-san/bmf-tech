---
title: Installing anyenv on Mac
slug: installing-anyenv-on-mac
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - shell script
  - anyenv
translation_key: installing-anyenv-on-mac
---

This is the procedure for installing anyenv on Mac. I encountered a few issues, so here are my notes.

# Installation
In my environment, I have various things placed in `/usr/local/bin/`, so I will install anyenv there.

`cd /usr/local/bin`
`git clone https://github.com/riywo/anyenv`

# Setting the Path

```shell-session:~/.bashrc
export PATH="/usr/local/bin/anyenv/bin:$PATH"
export ANYENV_ROOT=/usr/local/bin/anyenv
eval "$(anyenv init -)"
```

It seems that anyenv is designed to be installed directly under the root directory, so if you don't specify `ANYENV_ROOT` to a specific directory, the anyenv command won't execute correctly. Also, if you forget to write `eval "$(anyenv init -)"`, you may encounter issues where the commands of the installed packages cannot be executed, so be sure not to forget this.

# Conclusion
The installation is now complete. You should be able to use various anyenv commands.

# Reference
- [github - riywo/anyenv](https://github.com/riywo/anyenv)