---
title: How to Read and Set Linux Permissions
slug: linux-permission-reading-specification
date: 2018-07-18T00:00:00Z
author: bmf-san
categories:
  - Operating Systems
tags:
  - Linux
  - Permissions
description: An overview of Linux permissions
translation_key: linux-permission-reading-specification
---

# Overview
A summary of Linux permissions.

# How to Read Permissions

## File Type
The first character represents the file type.

- - File
- d Directory
- l Symbolic link

The following three-character groups represent permissions for the file owner.

2nd to 4th characters: User - permissions for the file owner  
5th to 7th characters: Group - permissions for the file group  
8th to 10th characters: Others - permissions for others  

## Types of Permissions
There are three types of permissions.

- r Read
- w Write
- x Execute

The meaning changes depending on whether it is a file or a directory, so be careful.

### For Files
- r Read - can read the contents of the file
- w Write - can edit the contents of the file
- x Execute - can execute the file as a program

### For Directories
- r Read - can list the files within the directory
- w Write - can create or delete files within the directory*1
- x Execute - can navigate into the directory*2

*1 Files within the directory can be deleted even without write permission.  
*2 If execute permission is not granted to the directory, you cannot navigate into that directory.

## Example

```
ls -l example.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 example.md
```

Permissions for example.md:

File type: File  
User permissions: Read/Write  
Group permissions: Read  
Other permissions: Read  

# How to Set Permissions
There are two methods: specifying with numbers and specifying with letters. The chmod command is used to set permissions.

## Specifying with Numbers
chmod mode target_filename

Use three digits to specify permissions.  
The hundreds place is for the user, the tens place is for the group, and the units place is for others.

- 4 r Read
- 2 w Write  
- 1 x Execute

To grant multiple permissions, specify the sum of the numbers.  
For example, to grant read and write permissions, specify 6; to grant all permissions, specify 7.

```
ls -l example.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 example.md

chmod 766 example.md

-rwxrw-rw-  1 bmf  staff  1788 Jul 18 11:57 example.md
```

## Specifying with Letters
chmod target+operation+permissions target_file

Specify the target, operation, and permissions using letters and symbols.

### Target
- u User
- g Group
- o Others
- a All

Multiple targets can be specified.  
For example, if you want to target both user and group, specify ug.

### Operation
- = Set to specified permissions
- + Add specified permissions
- - Remove specified permissions

### Permissions
- r Read
- w Write
- x Execute

Multiple permissions can be specified.  
For example, if you want to specify read and write, specify rw.

```
ls -l example.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 example.md

chmod a+rw example.md

-rw-rw-rw-  1 k.takeuchi  staff  2877 Jul 18 12:24 example.md
```

# References
- [Checking and Changing Linux Permissions](https://example.com/linux-permissions-check-change)
- [Linux File Management: Setting Permissions](http://proengineer.internous.co.jp/content/columnfeature/8843)

# By the Way
```
ls -l example.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 example.md
```

The number 1 displayed after the permissions in the output of ls -l indicates the number of hard links.
