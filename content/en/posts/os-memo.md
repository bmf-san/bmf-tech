---
title: OS Notes
slug: os-memo
date: 2018-05-11T00:00:00Z
author: bmf-san
categories:
  - Operating Systems
tags:
  - OS
description: "Explore OS fundamentals including swap memory, page cache, buffer cache, slab cache, dentry/inode structures, hard links, and symbolic links."
translation_key: os-memo
---



# Overview
Notes on what I researched while studying about operating systems.

# Swap
A feature that moves the contents of memory to the hard disk when memory is insufficient.

# What is Page Cache
- The CPU cannot read data directly from storage.
- Data needs to be loaded into memory first.
- Data loaded into memory can be reused as page cache.
- Cache used when accessing files.

# Buffer Cache
- Cache used when directly accessing block devices.

# Slab Cache
- A memory area within the kernel that caches structures like dentry, which stores directory metadata, and inode, which stores file metadata.
- Memory used by the kernel such as inode and dentry.
- The kernel has a mechanism to cache various memory resources within the kernel space by resource to improve memory utilization efficiency = slab cache.
- Dentry Cache
  - Directory entry cache
    - A structure that associates hard links, parent-child relationships of directories and file names, and inode with directory names and file names.
  - What is a Hard Link
    - The association of resources like files and directories on a computer's file system with the names given to those resources, or the association itself.
    - Allows an existing file to be referenced by a different file name (path name).
    - Referenced by inode.
    - Limited to the same file system.
    - File (or path)
    - Not affected by moving the original file.
    - Does not disappear even if the original file is deleted.
  - Symbolic Link
    - Specifies an existing file by path instead of inode, allowing files from different file systems to be referenced.
    - Referenced by path.
    - Can be referenced from different file systems.
    - Can be specified for directories as well.
    - Cannot be referenced if the original file is moved.
    - Disappears if the original file is deleted.
  - Directory Entry
    - A special file within each directory that lists the names of files, their actual storage locations, creation dates, update dates, file attributes, sizes, etc.
- Slab Allocator
  - An algorithm discovered by Jeff Bonwick.
- Inode
  - A storage area for file information created for each file.
  - Records the name, size, permissions, owner, update information, etc., of files stored on the disk.

