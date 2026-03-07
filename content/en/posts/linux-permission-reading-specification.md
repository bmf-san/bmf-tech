---
title: Understanding and Setting Linux Permissions
slug: linux-permission-reading-specification
date: 2018-07-18T00:00:00Z
author: bmf-san
categories:
  - OS
tags:
  - Linux
  - Permissions
translation_key: linux-permission-reading-specification
---

# Overview
This post summarizes Linux permissions.

# Reading Permissions

## File Type
The first character represents the file type.

- `-` File
- `d` Directory
- `l` Symbolic Link

The following three characters represent permissions for the file owner, group, and others.

- 2-4 characters: User permissions (owner)
- 5-7 characters: Group permissions
- 8-10 characters: Other permissions

## Types of Permissions
There are three types of permissions:

- `r` Read
- `w` Write
- `x` Execute

Be careful as the meaning changes depending on whether it is a file or a directory.

### For Files
- `r` Read: Can read the contents of the file
- `w` Write: Can edit the contents of the file
- `x` Execute: Can execute the file as a program

### For Directories
- `r` Read: Can list files in the directory
- `w` Write: Can create or delete files in the directory
- `x` Execute: Can navigate into the directory

*Note 1: Files in the directory can be deleted even without write permission.*
*Note 2: If the directory does not have execute permission, you cannot navigate into it.*

## Example

```
ls -l hoge.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 hoge.md
```

Permissions for hoge.md:
- File Type: File
- User Permissions: Read/Write
- Group Permissions: Read
- Other Permissions: Read

# Setting Permissions
Permissions can be set using either numeric or alphabetic methods. The `chmod` command is used to set permissions.

## Numeric Method
```
chmod mode target_filename
```

Use three digits to specify permissions. The hundreds place is for the user, the tens place is for the group, and the units place is for others.

- 4: `r` Read
- 2: `w` Write
- 1: `x` Execute

To grant multiple permissions, specify the sum of the numbers. For example, to grant read and write permissions, specify 6; to grant all permissions, specify 7.

```
ls -l hoge.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 hoge.md

chmod 766 hoge.md

-rwxrw-rw-  1 bmf  staff  1788 Jul 18 11:57 hoge.md
```

## Alphabetic Method
```
chmod target+method+permissions target_file
```

Specify the target, method, and permissions using letters and symbols.

### Target
- `u` User
- `g` Group
- `o` Others
- `a` All

Multiple targets can be specified. For example, to target both user and group, specify `ug`.

### Method
- `=` Set to specified permissions
- `+` Add specified permissions
- `-` Remove specified permissions

### Permissions
- `r` Read
- `w` Write
- `x` Execute

Multiple permissions can be specified. For example, to specify read and write, use `rw`.

```
ls -l hoge.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 hoge.md

chmod a+rw hoge.md

-rw-rw-rw-  1 k.takeuchi  staff  2877 Jul 18 12:24 hoge.md
```

# References
- [Checking and Changing Linux Permissions](Linuxの権限確認と変更（超初心者向け）)
- [Linux File Management: Setting Permissions](http://proengineer.internous.co.jp/content/columnfeature/8843)

# Note
```
ls -l hoge.md

-rw-r--r--  1 bmf  staff  652 Jul 18 11:45 hoge.md
```

The number `1` displayed after the permissions in the output of `ls -l` indicates the number of hard links.