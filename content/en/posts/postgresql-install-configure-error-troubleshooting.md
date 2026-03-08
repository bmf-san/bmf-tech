---
title: Troubleshooting 'configure' Errors During PostgreSQL Installation
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

When compiling the source code of PostgreSQL, an error occurred when running `make`: `You need to run the 'configure' program first.`

I will document the troubleshooting steps.

## 2. Error Details

Error message:

```sh
root@7a46fccdd51a:/postgresql-14.7# make -j$(nproc)
You need to run the 'configure' program first. See the file
'INSTALL' for installation instructions.
Makefile:19: recipe for target 'all' failed
make: *** [all] Error 1
```

This error occurs because the `configure` script has not been executed in the PostgreSQL source code directory.

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

To check available options, run the following command:

```sh
./configure --help
```

### Step 2: Re-run `make`

Once `configure` completes successfully, run the following command to build:

```sh
make -j$(nproc)
```

This will utilize the available CPU cores for faster compilation.

### Step 3: Run `make install` (if necessary)

If `make` succeeds, the next step is to install PostgreSQL:

```sh
make install
```

This will install PostgreSQL in the specified prefix directory (default is `/usr/local/pgsql`).

## 4. Additional Issues and Solutions When Running `configure`

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

If errors still occur, check if the required packages are installed.

#### For Ubuntu/Debian

```sh
apt update && apt install -y build-essential libreadline-dev zlib1g-dev flex bison
```

#### For CentOS/RHEL

```sh
yum groupinstall -y "Development Tools"
yum install -y readline-devel zlib-devel flex bison
```

### Issue 2: Library errors when running `configure`

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

The failure of `make` during PostgreSQL compilation can be attributed to the `configure` not being executed or missing required libraries. By following the steps outlined in this article, you may resolve the following issues:

1. Execute `configure` to properly set up the environment
2. Run `make` to build PostgreSQL
3. Execute `make install` if necessary to install
4. Properly resolve issues that arise during the execution of `configure` (such as missing libraries)