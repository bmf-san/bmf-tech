---
title: "Processes vs Threads: Key Differences Every Developer Should Know"
description: 'Understand the key differences between processes and threads—memory isolation, context switching costs, and when to use multi-processing vs multi-threading in your code.'
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
translation_key: process-vs-thread
---

# Overview
Summarizing the differences between processes and threads

# What is a Process
- The unit of execution for a program
- Executed on the CPU
- Cannot share resources
- Uses dedicated memory space
  - Holds a virtual address space
  - Switching processes requires mapping between virtual and physical addresses

## Data Structures Held by a Process
The data structures held in memory by a process are divided into two segments.

- Text Segment
    - Sequence of instructions for the program (the program itself)
        - Read-only

- Data Segment
    - PDA (Processor Data Area)
        - Contains processor information and data for process management
            - Holds stack pointer, program counter, etc.
               - Stack Pointer
                   - When the internal registers of the CPU (the storage for data within the CPU) are insufficient during program execution, data may be temporarily stored in main memory.
                   - This temporary storage area is called the stack, and the address that holds temporary information is maintained by the stack pointer.
             - Program Counter
                   - A register that holds the memory address of the next instruction to be executed.
    - Data Area
        - Static Area
             - Holds constants and global variables.
             - Heap Area
              - Holds regular variables, etc.
              - The size is unknown at runtime as the process dynamically increases or decreases the area.
    - Stack Area
        - Temporarily holds arguments and local scope data.

# What is a Thread
- An execution unit generated from a single process
- Can share resources
- Utilizes shared memory space
  - Holds data within the parent process's memory space.

# Differences Between Processes and Threads During Parallel Execution
The differences between processes and threads when executing a program in parallel.

## Process
- Can run programs in parallel by forking the parent process to launch multiple child processes.
- When launching a child process, memory space for the child process is allocated.
    - The allocated memory space copies the parent process's data segment and secures a dedicated data segment for the child process.
- Since the program instructions themselves are the same as the parent process, the text segment references the same area as the parent process.
- A child process is still a single process, so it cannot directly access the memory of other processes.

## Thread
- Can run programs in parallel by generating threads from a process.
- When generating a thread, the following values are copied into the parent process's virtual address space:
    - Stack Area
    - Stack Pointer
    - Program Counter
- Data other than the above is shared with the parent process.
- A thread only holds data about which part of the program it is executing.
- There is a possibility that a variable a thread intends to use may be modified by another thread.
    - Programs that operate without issues in such situations are called thread-safe.

# Performance

## Process
- Process switching
    - Requires clearing the cache held by the MMU.
        - TLB flush.

## Thread
- Thread switching
    - Only requires switching the stack area, stack pointer, and program counter.

# Aside

## Stack
- Automatically allocated and freed by the OS or compiler.
- The size is determined during program compilation and linking.

## Heap
- The application must handle the allocation when memory is needed and release it when it is no longer needed.
- It is possible to specify the size dynamically when allocating memory.

# References
- [Differences Between Processes and Threads](https://imokuri123.com/blog/2013/12/difference-between-process-and-thread.html)
- [Let's Become Cool Engineers Series - Memory, Processes, and Threads Edition](http://moro-archive.hatenablog.com/entry/2014/09/11/013520)
- [Heap and Stack](https://www.uquest.co.jp/embedded/learning/lecture16.html)