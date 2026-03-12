---
title: Redirect with exec Command
description: An in-depth look at Redirect with exec Command, covering key concepts and practical insights.
slug: exec-command-redirection
date: 2019-05-08T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - bash
  - shell script
  - exec
translation_key: exec-command-redirection
---

# Overview
The exec command replaces the current process with a command, but when used without arguments, it allows for dynamic changes to redirection.

This was prompted by the code that appeared when confirming at the prompt during [direct push to master](https://bmf-tech.com/posts/master%E7%9B%B4%E3%83%97%E3%83%83%E3%82%B7%E3%83%A5%E3%81%AE%E9%9A%9B%E3%81%AB%E3%83%97%E3%83%AD%E3%83%B3%E3%83%97%E3%83%88%E3%81%A7%E7%A2%BA%E8%AA%8D%E3%81%99%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%99%E3%82%8B):

```sh
#!/bin/sh
eval < /dev/tty
read ANSWER
```

I didn't quite understand it, so I looked it up.

# Usage
```sh
#!/bin/sh
echo "Output to stdout" // Standard input
exec > redirect.txt // Change file descriptor
echo "Output to file" // Output to file
```

The code that appeared during [direct push to master](https://bmf-tech.com/posts/master%E7%9B%B4%E3%83%97%E3%83%83%E3%82%B7%E3%83%A5%E3%81%AE%E9%9A%9B%E3%81%AB%E3%83%97%E3%83%AD%E3%83%B3%E3%83%97%E3%83%88%E3%81%A7%E7%A2%BA%E8%AA%8D%E3%81%99%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%99%E3%82%8B) is:

```sh
#!/bin/sh
exec < /dev/tty
read ANSWER
```

This passes the input from the current terminal (`/dev/tty`) to the standard input of the exec command. (It's a bit confusing...)

# Thoughts
It seems useful when you can't read from standard input for some reason.

# References
- [SEEKPOINT - exec redirect](http://seekpoint.blogspot.com/2012/12/exec.html)
- [UNIX/Linux Room Command: exec](http://x68000.q-e-d.net/~68user/unix/pickup?exec#prgmemo-exec-basic)
- [exec - Replace with executed command](https://linuxcommand.net/exec/)
- [Qiita - Issues and solutions when processing screen output with Process Substitution and exec redirect, and the story of sleeping infinitely](https://qiita.com/takei-yuya@github/items/7afcb92cfe7e678b7f6d#%E3%81%AF%E3%81%98%E3%82%81%E3%81%AB2-exec-%E3%81%A8%E3%83%AA%E3%83%80%E3%82%A4%E3%83%AC%E3%82%AF%E3%83%88)
- [Qiita - Various input/output controls in shell](https://qiita.com/tag1216/items/7ce35b7c27d371165e56#%E6%A8%99%E6%BA%96%E5%87%BA%E5%8A%9B%E3%81%A8%E6%A8%99%E6%BA%96%E3%82%A8%E3%83%A9%E3%83%BC%E5%87%BA%E5%8A%9B%E3%82%92%E5%88%A5%E3%80%85%E3%81%AE%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%AB%E5%87%BA%E5%8A%9B%E3%81%99%E3%82%8B)