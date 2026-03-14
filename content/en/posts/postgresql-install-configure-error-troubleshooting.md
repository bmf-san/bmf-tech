---
title: Troubleshooting 'configure' Errors During PostgreSQL Installation
description: 'Fix the You need to run the configure program first error when compiling PostgreSQL from source. Covers ./configure options, re-running make, and common dependency issues on Ubuntu/CentOS.'
slug: postgresql-install-configure-error-troubleshooting
date: 2025-03-07T00:00:00Z
author: bmf-san
categories:
  - Database
tags:
  - PostgreSQL
  - Ubuntu
  - CentOS
translation_key: postgresql-install-configure-error-troubleshooting
---


# Troubleshooting 'configure' Errors During PostgreSQL Installation
## 1. Introduction

When compiling the PostgreSQL source code, an error occurred when running `make`: `You need to run the 'configure' program first.`

This post documents the troubleshooting steps.

## 2. Error Details

Error message:

```sh
root@7a46fccdd51a:/postgresql-14.7# make -j$(nproc)
You need to run the 'configure' program first. See the file
'INSTALL' for installation instructions.
Makefile:19: recipe for target 'all' failed
make: *** [all] Error 1
```

This error occurs because the `configure` script has not been run in the PostgreSQL source code directory.

## 3. Solution

### Step 1: Run `configure`

The `configure` script is necessary to set up the build environment for PostgreSQL. First, navigate to the PostgreSQL source directory and run the following command:

```sh
./configure
```

If additional options are needed when running `configure`, execute it as follows:

```sh
./configure --prefix=/usr/local/pgsql --enable-debug --enable-cassert
```

To check available options, run:

```sh
./configure --help
```

### Step 2: Re-run `make`

Once `configure` completes successfully, execute the build with the following command:

```sh
make -j$(nproc)
```

This will compile quickly by utilizing the available CPU cores.

### Step 3: Run `make install` (if necessary)

If `make` succeeds, proceed to install PostgreSQL:

```sh
make install
```

This will install PostgreSQL in the specified prefix directory (default is `/usr/local/pgsql`).

## 4. Additional Issues and Solutions During `configure`

### Issue 1: `configure: command not found`

#### Cause

- The `configure` script does not exist
- Required build tools are missing

#### Solution

First, check if the `configure` script exists:

```sh
ls -l configure
```

If `configure` is not found, regenerate it by running `autoconf`:

```sh
autoreconf -fi
./configure
```

If errors persist, verify that the necessary packages are installed.

#### For Ubuntu/Debian

```sh
apt update && apt install -y build-essential libreadline-dev zlib1g-dev flex bison
```

#### For CentOS/RHEL

```sh
yum groupinstall -y "Development Tools"
yum install -y readline-devel zlib-devel flex bison
```

### Issue 2: Library Errors During `configure`

Example error message:

```sh
configure: error: readline library not found
```

#### Solution

Install the missing library:

```sh
apt install -y libreadline-dev
```

or

```sh
yum install -y readline-devel
```

## 5. Conclusion

The failure of `make` during PostgreSQL compilation can be due to not running `configure` or missing necessary libraries. By following the steps outlined in this article, you may resolve the following issues:

1. Run `configure` to properly set up the environment
2. Execute `make` to build PostgreSQL
3. Run `make install` if necessary to install
4. Properly address issues that arise during `configure` (such as missing libraries)
