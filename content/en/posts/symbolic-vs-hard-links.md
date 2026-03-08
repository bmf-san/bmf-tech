---
title: Differences Between Symbolic Links and Hard Links
slug: symbolic-vs-hard-links
date: 2018-05-24T00:00:00Z
author: bmf-san
categories:
  - OS
tags:
  - Symbolic Links
  - Hard Links
translation_key: symbolic-vs-hard-links
---

# Overview
Summarizing the differences between symbolic links and hard links

# Prerequisites
- inode
  - Data structure
  - Holds attribute information on the file system (creator, group, creation date, etc.) as data
  - You can check the inode number with `ls -i1 /` or `stat /`

# What is a Symbolic Link?
- Adds a directory entry that references the path of the original file or directory
- Experiment
  ```
  touch a.md
  ln -s a.md a_symbolic_link.md // Create a symbolic link
  ls -i1 a.md a_symbolic_link.md // Confirm that the inodes are different
  ```
- Cannot reference if the original file is moved
- Deleted when the original file is deleted
- Can reference across different file systems

# What is a Hard Link?
- Adds a directory entry that references the inode of the original file or directory
- Experiment
  ```
  touch a.md
  ln a.md a_hardlink.md // Set up a hard link
  ls -i1 a.md a_hardlink.md // Confirm that the inodes are the same
  ```
- No effect if the original file is moved
- Not deleted when the original file is deleted
- Can only reference within the same file system

# References
- [Differences Between Symbolic Links and Hard Links](https://qiita.com/katsuo5/items/fc57eaa9330d318ee342)
- [Practice Now! Linux System Administration](http://www.usupi.org/sysad/242.html)