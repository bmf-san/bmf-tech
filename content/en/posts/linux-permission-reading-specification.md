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
This post summarizes Linux permissions.

# How to Read Permissions

## File Types
The first character represents the file type.

- File
d Directory
l Symbolic link

Subsequent characters, in groups of three, represent permissions for different owners of the file.

2~4th characters: User permissions for the file owner
5~7th characters: Group permissions for the owning group
8~10th characters: Other permissions

## Types of Permissions
There are three types of permissions:

r Read
w Write
x Execute

The meaning changes depending on whether it's a file or a directory.

### For Files
r Read: Can read the file contents
w Write: Can edit the file contents
x Execute: Can execute the file as a program

### For Directories
r Read: Can display the list of files under the directory
w Write: Can create or delete files under the directory※1
x Execute: Can move into the directory※2

※1 Files under the directory can be deleted even without write permission
※2 If execute permission is not granted to the directory, you cannot move into it

## Example

```
ls -l hoge.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 hoge.md
```

Permissions for hoge.md

File type: File
User permissions: Read/Write
Group permissions: Read
Other permissions: Read

# How to Set Permissions
Permissions can be specified numerically or alphabetically. Use the chmod command to set permissions.

## Specifying Numerically
chmod mode target_filename

Use three digits to specify permissions. The hundreds place is for the user, the tens place is for the group, and the ones place is for others.

4 r Read
2 w Write
1 x Execute

To grant multiple permissions, specify the sum of the numbers. For example, to grant read and write permissions, specify 6; to grant all permissions, specify 7.

```
ls -l hoge.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 hoge.md

chmod 766 hoge.md

-rwxrw-rw-  1 bmf  staff  1788 Jul 18 11:57 hoge.md
```

## Specifying Alphabetically
chmod target+method+content target_file

Specify the target, method, and content using alphabets and symbols.

### Target
u User
g Group
o Others
a All

Multiple targets can be specified. For example, to target both user and group, specify ug.

### Method
= Set specified permissions
+ Add specified permissions
- Remove specified permissions

### Content
r Read
w Write
x Execute

Multiple contents can be specified. For example, to specify read and write, specify rw.

```
ls -l hoge.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 hoge.md

chmod a+rw hoge.md

-rw-rw-rw-  1 k.takeuchi  staff  2877 Jul 18 12:24 hoge.md
```

# References
- [Checking and Changing Linux Permissions](Linuxの権限確認と変更（超初心者向け）)
- [Linux File Management: Setting Permissions](http://proengineer.internous.co.jp/content/columnfeature/8843)

# Additional Notes
```
ls -l hoge.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 hoge.md
```

The number 1 displayed next to the permissions in the output of ls -l represents the number of hard links.
