---
title: Unix Command Notes
slug: unix-command-memo
date: 2018-07-07T00:00:00Z
author: bmf-san
categories:
  - OS
tags:
  - unix command
  - jq
  - tee
translation_key: unix-command-memo
---

# Overview
Notes on Unix commands.

# jq
A command for processing data in JSON format.

## JSON Pretty Print
```
echo '[{"name": "Tom", "age": 20}}]' | jq .
```

In addition to Pretty Print, there are various uses such as extracting data by specifying properties from an object or obtaining the length of an object.

# tee
Outputs standard input to both standard output and a file.
Can be used with sudo.
Overwrite without options, append with option -a.

```
echo 'hello world' | sudo tee ./sample.txt
```

In the case of redirection, `>` overwrites and `>>` appends. sudo cannot be used.

# at
Allows you to schedule the execution time of a command.

```
at -f ./sample.txt 2230
```

There are various formats for the date and time part.

# mktemp
Creates a file in the /tmp directory with a random name.

```
mktemp
```

# lsof
List of open files.
Outputs files that a process has opened.

-i
Filters output by port number.

-i tcp or udp
Filters output by TCP or UDP.

-P
Outputs port numbers as numbers.

-n
Outputs IP addresses without resolving to hostnames.

# nmap
Checks the port status of the target host over the network.

```
nmap 192.168.33.10
```

# fsfreeze
A command to temporarily stop I/O on a filesystem.

Temporarily stops I/O on the filesystem (freeze operation).
```
fsfreeze -f /data
```

Releases I/O on the filesystem (unfreeze operation).
```
fsfreeze -u /data
```

# findmnt
Outputs information about mounted filesystems.