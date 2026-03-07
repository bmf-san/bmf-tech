---
title: Redirecting with the exec Command
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
The `exec` command replaces the current process with the specified command, but when used without arguments, it allows dynamic redirection.

For example, in the post [Prompt for confirmation when pushing directly to master](https://bmf-tech.com/posts/master%E7%9B%B4%E3%83%97%E3%83%83%E3%82%B7%E3%83%A5%E3%81%AE%E9%9A%9B%E3%81%AB%E3%83%97%E3%83%AD%E3%83%B3%E3%83%97%E3%83%88%E3%81%A7%E7%A2%BA%E8%AA%8D%E3%81%99%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%99%E3%82%8B), the following code appeared:

```sh
#!/bin/sh
exec < /dev/tty
read ANSWER
```

I didn't quite understand it at first, so I decided to look into it.

# Usage
```sh
#!/bin/sh
echo "Output to stdout" // Standard output
exec > redirect.txt // Change the file descriptor
echo "Output to file" // This will be written to the file
```

The code from [Prompt for confirmation when pushing directly to master](https://bmf-tech.com/posts/master%E7%9B%B4%E3%83%97%E3%83%83%E3%82%B7%E3%83%A5%E3%81%AE%E9%9A%9B%E3%81%AB%E3%83%97%E3%83%AD%E3%83%B3%E3%83%97%E3%83%88%E3%81%A7%E7%A2%BA%E8%AA%8D%E3%81%99%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%99%E3%82%8B) does the following:

```sh
#!/bin/sh
exec < /dev/tty
read ANSWER
```

It passes the input from the current terminal (`/dev/tty`) to the standard input of the `exec` command. (This is a bit hard to understand...)

# Thoughts
This seems useful in situations where standard input cannot be read for some reason.

# References
- [SEEKPOINT - exec Redirect](http://seekpoint.blogspot.com/2012/12/exec.html)
- [UNIX/Linux no Heya - Command: exec](http://x68000.q-e-d.net/~68user/unix/pickup?exec#prgmemo-exec-basic)
- [exec - Replace the shell with the executed command](https://linuxcommand.net/exec/)
- [Qiita - Issues and solutions when processing screen output with Process Substitution and exec redirect, and the story of sleeping infinitely](https://qiita.com/takei-yuya@github/items/7afcb92cfe7e678b7f6d#%E3%81%AF%E3%81%98%E3%82%81%E3%81%AB2-exec-%E3%81%A8%E3%83%AA%E3%83%80%E3%82%A4%E3%83%AC%E3%82%AF%E3%83%88)
- [Qiita - Various controls for shell input/output](https://qiita.com/tag1216/items/7ce35b7c27d371165e56#%E6%A8%99%E6%BA%96%E5%87%BA%E5%8A%9B%E3%81%A8%E6%A8%99%E6%BA%96%E3%82%A8%E3%83%A9%E3%83%BC%E5%87%BA%E5%8A%9B%E3%82%92%E5%88%A5%E3%80%85%E3%81%AE%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%AB%E5%87%BA%E5%8A%9B%E3%81%99%E3%82%8B)