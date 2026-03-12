---
title: Open Addressing Hash Tables and Swiss Tables
description: An in-depth look at Open Addressing Hash Tables and Swiss Tables, covering key concepts and practical insights.
slug: open-address-hash-table-swizzle-table
date: 2025-02-27T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Open Addressing Hash Table
  - Swiss Table
  - Data Structures
translation_key: open-address-hash-table-swizzle-table
---

# Open Addressing Hash Tables and Swiss Tables
While reading [go.dev - Faster Go maps with Swiss Tables](https://go.dev/blog/swisstable), I came across an explanation of open addressing hash tables and Swiss tables, so I decided to investigate further.

## 1. What is an Open Addressing Hash Table?
**Open Addressing Hash Table** is a representative method for resolving hash collisions (when different keys have the same hash value) and is also known as **"Open Addressing"**. Unlike chaining (using lists, etc., to handle collisions), it is characterized by **continuing to search and store in another slot of the same array (hash table) when a collision occurs**.

### 1.1 Mechanism
1. **Determine the index with a hash function**
   Pass the key \( k \) through the hash function \( h(k) \) to determine which index of the table to store it in.

2. **Find another place if a collision occurs**
   If the specified index is already in use, find the next available slot and store the data there.

3. **Examples of probing methods**
   - **Linear Probing**
     In case of a collision, search continuously with \( h(k) + 1 \), \( h(k) + 2 \), … in one-step increments.
   - **Quadratic Probing**
     In case of a collision, search quadratically with \( h(k) + 1^2, h(k) + 2^2, h(k) + 3^2, … \).
   - **Double Hashing**
     Use another hash function \( h'(k) \) to determine the next probing step in case of a collision.

### 1.2 Features
- **Memory Efficient**
  It does not require pointers or lists used in chaining for collision resolution, thus using minimal extra memory.
- **Cache Efficient**
  Data is stored on a continuous array, which tends to increase cache hit rates.
- **Slightly Complex Deletion**
  If the deleted slot is emptied, the search may be interrupted, so it is necessary to manage it with a "deleted" mark (tombstone).
- **Increased Search Cost Under High Load**
  As the table usage rate increases, collisions occur more frequently, increasing the number of searches needed to find a slot.

## 2. What is a Swiss Table?
**Swiss Table** is a **high-performance hash table** developed by Google, adopted in C++'s **Abseil Library (`absl::flat_hash_map`)** and Rust's **`hashbrown`**. It is an implementation that thoroughly pursues **memory efficiency and speed** based on open addressing.

### 2.1 Features
1. **High-Speed Search Using SIMD**
   By using **SIMD (Single Instruction, Multiple Data)**, multiple buckets can be checked in parallel at once. This allows collision resolution and search comparisons to be processed together, making it faster than conventional open addressing.

2. **Bucket Grouping**
   Swiss Table manages in units called **bucket groups** that combine multiple slots. For example, if a group contains 8 or 16 slots, they can be scanned at once, efficiently resolving collisions and searches.

3. **Use of Control Bytes (Tags)**
   Each bucket group is equipped with **Control Bytes** (also called Control Block), which collectively holds the state of each slot, such as "in use" or "deleted," and part of the **full hash (tag)**. This allows narrowing down "potentially matching slots" before referring to the actual key, achieving high cache efficiency and search performance.

4. **Tombstone Method for Deletion**
   When deleting, instead of emptying the slot, a **tombstone** is placed. During rehashing or resizing, slots marked with tombstones are collectively rearranged and organized.

### 2.2 Mechanism (Basic Flow)
1. **Hash the Key and Record Partial Hash (Tag)**
   Obtain the hash value from the key and store part of it as a tag in the Control Bytes.

2. **SIMD Parallel Search Using Tags**
   Compare the tags of multiple slots in a bucket group at once using SIMD instructions to quickly narrow down candidates.

3. **Verify the Key of the Slot with Matching Tags**
   Once a slot with matching tags is found, retrieve the actual key to confirm a complete match.

4. **Move to the Next Bucket Group in Case of Collision**
   If there is no space in the same group or the key does not match, check the next bucket group through probing.

5. **Assign a Tombstone for Deletion**
   Instead of emptying the deleted slot, assign a tombstone. During resizing or reallocation, organize these unnecessary slots.

# Implementation
Refer to [bmf-san/road-to-algorithm-master/tree/master/data_structures/hash_table](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/hash_table).

The code was written as a practice for Cline, but there may still be some bugs.

## 3. Summary
Open addressing hash tables are a method that excels in memory and cache efficiency because they do not require extra pointer structures. However, they face challenges such as increased search times as load rates increase and complex deletion processes. On the other hand, **Swiss Table** achieves high performance by cleverly compensating for the shortcomings of open addressing through the use of SIMD, bucket grouping, and Control Bytes (tag management). It is adopted in implementations like Google's Abseil Library and Rust's hashbrown.