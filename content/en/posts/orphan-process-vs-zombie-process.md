---
title: Difference Between Orphan Processes and Zombie Processes
description: An in-depth look at Difference Between Orphan Processes and Zombie Processes, covering key concepts and practical insights.
slug: orphan-process-vs-zombie-process
date: 2020-08-24T00:00:00Z
author: bmf-san
categories:
  - Computer Architecture
tags:
  - UNIX
  - Process
translation_key: orphan-process-vs-zombie-process
---


# Overview
While working with Docker, I learned about the existence of orphan processes, so I decided to investigate the differences between them and zombie processes.

# What is a Zombie Process
- A child process that has completed execution
- Remains in the process table waiting for the parent process to perform a wait
- Does not use system resources, but retains its PID
- If many zombie processes accumulate, the available PIDs decrease, preventing other processes from starting
- How to check for zombie processes
  - Use `ps aux` to find processes with a stat of Z or those ending with defunct
  - Use `ps -ef | grep defunct` to output only zombie processes
- How to kill a zombie process
  - Kill the parent process

# Orphan Process
- A process whose parent process has exited without performing a wait
- Becomes a child process of the init process (PID 1), with the init process acting as the parent (foster parent)
- How to check for orphan processes
  - Use `ps -elf | head -1; ps -elf | awk '{if ($5 == 1 && $3 != "root") {print $0}}' | head`
- How to kill an orphan process
  - Use the kill command

# References
- [Qiita - What are Zombie and Orphan Processes in UNIX](https://qiita.com/ninoko1995/items/582106e8507163b2c50b)
- [Hibariya - About Forking Processes](https://note.hibariya.org/articles/20120326/a0.html)
- [tutorialspoint.com - Zombie and Orphan Processes in Linux](https://www.tutorialspoint.com/zombie-and-orphan-processes-in-linux)
- [Nikkei XTECH - Zombie Process](https://xtech.nikkei.com/it/article/Keyword/20070727/278487/)
- [geekride.com - Orphan Process](http://www.geekride.com/orphan-zombie-process/)
- [makiuchi-d.gihub.io - How to Ensure Killing Child Processes in Go](http://makiuchi-d.github.io/2020/05/10/go-kill-child-process.ja.html)
  - Includes code to reproduce orphan processes in Go
