---
title: Install anyenv on Mac
slug: installing-anyenv-on-mac
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - shell script
  - anyenv
description: Steps to install anyenv on Mac, including troubleshooting tips.
translation_key: installing-anyenv-on-mac
---

Steps to install anyenv on Mac. Documenting this as I encountered a few issues.

# Installation
In my environment, various tools are placed in `/usr/local/bin/`, so I decided to install anyenv there.

```bash
cd /usr/local/bin
git clone https://github.com/riywo/anyenv
```

# Add to Path

```shell-session:~/.bashrc
export PATH="/usr/local/bin/anyenv/bin:$PATH"
export ANYENV_ROOT=/usr/local/bin/anyenv
eval "$(anyenv init -)"
```

Since anyenv assumes installation directly under the root directory, the `ANYENV_ROOT` must be explicitly set to the desired directory; otherwise, the anyenv commands won't work correctly. Additionally, forgetting to include `eval "$(anyenv init -)"` can cause issues, such as being unable to execute commands from installed packages. Make sure not to skip this step.

# Conclusion
This completes the installation. You should now be able to use various anyenv commands.

# References
- [github - riywo/anyenv](https://github.com/riywo/anyenv)