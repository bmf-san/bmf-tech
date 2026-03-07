---
title: Differences Between Processes and Threads
slug: process-vs-thread
date: 2018-06-25T00:00:00Z
author: bmf-san
categories:
  - Operating Systems
tags:
  - OS
  - Thread
  - Process
  - Stack
  - Heap
description: A summary of the differences between processes and threads.
translation_key: process-vs-thread
---

# Overview
A summary of the differences between processes and threads.

# What is a Process?
- A unit of program execution
- Executed on the CPU
- Cannot share resources
- Uses dedicated memory space
  - Maintains a virtual address space
  - Switching processes requires mapping between virtual and physical addresses

## Data Structures Maintained by a Process
The data structures maintained in memory by a process are divided into two segments:

- **Text Segment**
    - Sequence of program instructions (the program itself to be executed)
        - Read-only

- **Data Segment**
    - **PDA (Processor Data Area)**
        - Contains processor information and data for process management
            - Includes the stack pointer and program counter
               - **Stack Pointer**
                   - When the CPU's internal registers (storage locations within the CPU) are insufficient during program execution, data may be temporarily stored in main memory.
                   - This temporary storage location is called the stack, and the stack pointer holds the address of this temporary storage.
               - **Program Counter**
                   - A register that holds the memory address of the next instruction to be executed
    - **Data Area**
        - **Static Area**
             - Stores constants and global variables
             - **Heap Area**
              - Stores regular variables
              - The process dynamically increases or decreases this area, so its size is unknown until runtime
    - **Stack Area**
        - Temporarily stores data such as arguments and local scope data

# What is a Thread?
- An execution unit generated from a single process
- Can share resources
- Uses shared memory space
  - Stores data within the memory space of the parent process

# Differences Between Processes and Threads During Parallel Execution
Differences between processes and threads when executing programs in parallel:

## Process
- Programs can be executed in parallel by forking the parent process to create multiple child processes
- When a child process is created, memory space is allocated for it
    - The allocated memory space includes a copy of the parent process's data segment, and a dedicated data segment is reserved for the child process
- Since the program instructions are the same as the parent process, the text segment refers to the same area as the parent process
- A child process is still a separate process and cannot directly access the memory of other processes

## Thread
- Programs can be executed in parallel by generating threads from a process
- When a thread is created, the following values are copied into the parent process's virtual address space:
    - Stack area
    - Stack pointer
    - Program counter
- Other data is shared with the parent process
- A thread only holds data about which part of the program it is executing
- Variables used by a thread may be modified by other threads
    - Programs that can operate without issues in such situations are called thread-safe

# Performance

## Process
- **Process Switching**
    - Requires clearing the cache held by the MMU
        - TLB flush

## Thread
- **Thread Switching**
    - Only requires switching the stack area, stack pointer, and program counter

# Miscellaneous

## Stack
- Automatically allocated and released by the OS or compiler
- The size is determined during program compilation and linking

## Heap
- Applications need to allocate and release memory as needed
- The size can be dynamically specified when allocating memory

# References
- [Differences Between Processes and Threads](https://imokuri123.com/blog/2013/12/difference-between-process-and-thread.html)
- [Let's Become a Cool Engineer Series ~ Memory, Processes, and Threads Edition ~](http://moro-archive.hatenablog.com/entry/2014/09/11/013520)
- [Heap and Stack](https://www.uquest.co.jp/embedded/learning/lecture16.html)
