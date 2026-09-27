---
title: "The Lineage of Programming Languages Leading to Go"
slug: programming-language-lineage-to-go
date: 2026-09-27
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Design
description: "How programming languages evolved from Algol 60 to Go: Wirth's Pascal-to-Oberon path and Hoare's CSP-to-Alef concurrency path, and how both streams merge into Go."
translation_key: programming-language-lineage-to-go
draft: false
---


# Overview

The history of programming languages runs along two great lineages. Niklaus Wirth shaped one of them, the lineage of structure and modularity, from Pascal onward. Tony Hoare started the other with CSP, the root of a concurrency lineage. The two rivers cross over decades and finally merge into Go.

This article walks the path from Algol 60 to Go and organizes the traits and historical links of each language.

# Two great lineages

Start with the big picture. Go joins two streams into a single language.

- The Oberon line: a lean type system, object orientation without inheritance, module management through packages, and garbage collection
- The CSP / Alef line: process communication through channels, the `select` statement, and lightweight threads

Go builds on the simple syntax of C. It fuses Oberon's lean type system with the CSP-style concurrency of Newsqueak and Alef. Then a modern runtime solves the memory management that troubled Alef.

# 1. The Algol 60 lineage: the origin of structured programming

## Algol 60 (1960)

Nearly every procedural language today descends directly from Algol 60.

Algol 60 became the first language to introduce block structure through `begin ... end`, local variable scope, recursion, and grammar definition in BNF notation. From C and Pascal to Java and Go, the basic rules of syntax start here.

## Pascal (1970)

Niklaus Wirth designed Pascal for teaching and for writing systems.

Pascal simplified Algol 60 and added a strict type system, including the `record` type. It also earned a reputation for fast compilation. Pascal spread structured programming across the world and stood beside C as a standard language.

# 2. Modularity and the Oberon lineage: Wirth's pursuit

This part follows how Wirth resisted the growing complexity of software and pursued simplicity and modularity.

## Modula-2 (1978)

Modula-2 succeeded Pascal and introduced the concept of the module, exactly as its name suggests.

It split the definition part (`DEFINITION MODULE`) from the implementation part (`IMPLEMENTATION MODULE`), which enabled separate compilation and encapsulation. It also supported simple concurrency through coroutines.

## Oberon (1987)

Oberon stripped Modula-2 down to the extreme and shipped alongside a tiny operating system, Oberon OS.

Wirth removed the complex features and added only two things: type extension, which resembles single inheritance, and garbage collection. Oberon marks the peak of his philosophy: gain the most expressive power from the fewest features.

## Object Oberon (1989) and Oberon-2 (1991)

These two versions added full object orientation to Oberon and refined it. They introduced type-bound procedures that carry a receiver, that is, methods.

One choice stands out: they added no `class` construct. Instead, they attached methods directly to Oberon's type extension.

### The connection to Go

Go's design carries a clear mark of Oberon and Oberon-2.

- It holds no `class` and binds methods to structs
- It uses type embedding instead of inheritance
- It manages modules through packages

Each idea descends from the Oberon line.

# 3. CSP and the concurrency lineage: the direct ancestor of Go's concurrency

Goroutines and channels, the defining features of Go, grew out of a concurrency model that centered on Bell Labs.

## CSP (1978)

Tony Hoare proposed CSP (Communicating Sequential Processes) as a theory and model of concurrency. Note that CSP names a concept, not a language.

CSP set out a clear philosophy: do not share memory to communicate; instead, communicate to share memory. It frames concurrency around processes and channels rather than threads and locks.

## Squeak (1985)

Luca Cardelli and Rob Pike at Bell Labs built Squeak to describe the concurrency of a screen UI, that is, a window system.

Squeak stands as an early experiment that expressed the CSP model as a programming language. Note that it differs from the Smalltalk implementation of the same name.

## Newsqueak (1988)

Rob Pike built Newsqueak on top of Squeak as a more practical concurrency language.

Its syntax sits close to Oberon and C. Newsqueak realized first-class channels (`chan`) and the `select` statement for asynchronous work in a clear form for the first time.

## Alef (1995)

Bell Labs built Alef as a systems programming language for Plan 9, its next-generation operating system project. Phil Winterbottom and Rob Pike joined its development.

Alef combined C-like syntax with Newsqueak's CSP-style concurrency (`proc`, `task`, `chan`), an Oberon-style type system, and garbage collection. Yet Alef never spread widely, because pointer manipulation in the style of C sat poorly with GC and error handling proved hard.

That failure later shaped the design of Go.

# Conclusion: every lineage leads to Go

In short, these languages form the lineage of technology that Go's designers passed through and refined on their way to Go. Rob Pike, Ken Thompson, and their colleagues walked this exact road.

```
[Algol 60]
   │
   ├─►[Pascal]─►[Modula-2]─►[Oberon / Oberon-2]
   │                              │ (type system, methods on structs, GC)
   │                              ▼
   │                          ┌────────┐
   │                          │   Go   │
   │                          └────────┘
   │                              ▲
   ├─►[C]                         │ (CSP-style concurrency: goroutine / channel)
   │    │                         │
   └────┴─►[CSP]─►[Newsqueak]─►[Alef]
```

Go inherited a lean type system and module management from the Oberon line, and channel-based concurrency from the CSP / Alef line. On top of the simple syntax of C, a modern runtime solved the memory management that Alef could not.

Behind a single line of Go lies more than half a century of language design. Once you know the lineage, the intent behind Go's design comes into sharper view.

