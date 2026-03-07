---
title: Difference Between Orphan Processes and Zombie Processes
slug: orphan-process-vs-zombie-process
date: 2020-08-24T00:00:00Z
author: bmf-san
categories:
  - Computer Architecture
tags:
  - UNIX
  - Processes
translation_key: orphan-process-vs-zombie-process
---

# Overview
While working with Docker, I learned about the existence of orphan processes and researched the difference between them and zombie processes.

# Zombie Processes
- A child process that has finished execution
- Remains in the process table waiting for the parent process to wait
- Does not use system resources, but the PID is retained
- If there are many zombie processes, the available PIDs decrease, preventing other processes from starting
- How to check for zombie processes
  - A process with stat Z and ending with defunct can be found using `ps aux`
  - Output only zombie processes with `ps -ef | grep defunct`
- Killing zombie processes
  - Kill the parent process

# Orphan Processes
- A process that ends without the parent process waiting
- Becomes a child process of the init process (PID 1), with init as the parent (or foster parent)
- How to check for orphan processes
  - Use `ps -elf | head -1; ps -elf | awk '{if ($5 == 1 && $3 != "root") {print $0}}' | head`
- Killing orphan processes
  - Use the kill command

# References
- [Qiita - What are Zombie and Orphan Processes?](https://qiita.com/ninoko1995/items/582106e8507163b2c50b)
- [Hibariya - When forking processes](https://note.hibariya.org/articles/20120326/a0.html)
- [tutorialspoint.com - Zombie and Orphan Processes in Linux](https://www.tutorialspoint.com/zombie-and-orphan-processes-in-linux)
- [Nikkei XTECH - Zombie Processes](https://xtech.nikkei.com/it/article/Keyword/20070727/278487/)
- [geekride.com - Orphan Process](http://www.geekride.com/orphan-zombie-process/)
- [makiuchi-d.gihub.io - How to Ensure Killing Child Processes in Go](http://makiuchi-d.github.io/2020/05/10/go-kill-child-process.ja.html)
  - Contains code that reproduces orphan processes in Go.