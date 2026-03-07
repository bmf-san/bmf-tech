---
title: OS Memo
slug: os-memo
date: 2018-05-11T00:00:00Z
author: bmf-san
categories:
  - Operating Systems
tags:
  - os
translation_key: os-memo
---

# Overview
Things I researched while studying about OS.

# Swap
A feature that moves the contents of memory to the hard disk when there is not enough memory.

# Page Cache
- The CPU cannot directly read data from storage.
- Data must first be loaded into memory.
- Data loaded into memory can be reused as page cache.
- A cache used when accessing data on a file basis.

# Buffer Cache
- A cache used when directly accessing block devices.

# Slab Cache
- A memory area within the kernel that caches structures like dentry for directory metadata and inode for file metadata.
- Memory used by the kernel for inodes and dentries.
- The kernel has a mechanism to cache various memory resources within the kernel space to improve memory utilization efficiency = slab cache.
- Dentry Cache
  - Directory Entry Cache
    - A structure that associates hard links, parent-child relationships of directories, file names, directory names, and inodes.
  - Hard Link
    - The association of resources such as files or directories on a computer's file system with the names assigned to those resources, or the association itself.
    - Allows referencing an existing file with a different file name (path name).
    - Referenced by inode.
    - Limited to the same file system.
    - File (or path).
    - Moving the original file does not affect it.
    - Deleting the original file does not make it disappear.
  - Symbolic Link
    - Allows referencing an existing file by path instead of inode, enabling references to files from another file system.
    - Referenced by path.
    - Can reference files from another file system.
    - Can also be specified for directories.
    - Moving the original file makes it unreferenced.
    - Deleting the original file makes it disappear.
  - Directory Entry
    - A special file that lists the names of files within each directory, their actual storage locations, creation dates, update dates, file attributes, sizes, etc.
- Slab Allocator
  - An algorithm discovered by Jeff Bonwick.
- Inode
  - A storage area for file information created for each file.
  - Records the name, size, permissions, owner, and update information of files stored on the disk.