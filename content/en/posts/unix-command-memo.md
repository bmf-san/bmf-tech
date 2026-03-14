---
title: Unix Command Notes
description: "Master Unix commands including jq for JSON processing, lsof for monitoring, and nmap for network port analysis."
slug: unix-command-memo
date: 2018-07-07T00:00:00Z
author: bmf-san
categories:
  - Operating Systems
tags:
  - unix commands
  - jq
  - tee
translation_key: unix-command-memo
---



# Overview
Notes on Unix commands.

# jq
A command to process data in JSON format.

## JSON Pretty Print
```
echo '[{"name": "Tom", "age": 20}}]' | jq .
```

Besides Pretty Print, you can extract data by specifying properties from objects, get the length of objects, and use it in various other ways.


# tee
Outputs standard input to both standard output and a file. Can be used with sudo. Overwrites without options, appends with option -a.

```
echo 'hello world' | sudo tee ./sample.txt
```

In the case of redirection, > overwrites, >> appends. sudo cannot be used.


# at
Allows you to schedule the execution time of a command.

```
at -f ./sample.txt 2230
```

There are various formats for the date and time part.


# mktemp
Creates a file with a random name under the /tmp directory.

```
mktemp
```


# lsof
List of open files
Outputs files opened by processes.

-i
Filters output by port number.

-i tcp or udp
Filters output by TCP or UDP.

-P
Outputs port numbers as digits.

-n
Outputs without reverse DNS lookup of IP addresses to hostnames.

# nmap
Examines the port status of a target host over the network.

```
nmap 192.168.33.10
```

# fsfreeze
A command to temporarily suspend file system I/O.

Temporarily suspends file system I/O (freeze process).
```
fsfreeze -f /data
```

Releases file system I/O (unfreeze process).
```
fsfreeze -u /data
```

# findmnt
Outputs information about mounted file systems.


