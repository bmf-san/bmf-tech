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
description: "A source-checked walk through the lineage of programming languages from Algol 60 to Go: Wirth's Pascal-to-Oberon path and Hoare's CSP path through Newsqueak, Alef, and Limbo, and how both streams merge into Go."
translation_key: programming-language-lineage-to-go
draft: false
---


# Overview

Go grew out of two great lineages. Niklaus Wirth shaped one of them, the lineage of structure and modularity, from Pascal through Oberon. Tony Hoare started the other with CSP, the root of a concurrency lineage. On the simple syntax of C, Go brings these two rivers together into a single language.

This article walks the path from Algol 60 to Go and organizes the traits and historical links of each language. Every claim rests on the sources at the end and matches the official Go FAQ and the primary references for each language.

# Two great lineages

Start with the big picture. The official Go FAQ describes Go's ancestry as three streams: most of the syntax comes from the C family; declarations and packages come from the Pascal/Modula/Oberon family; and concurrency comes from Newsqueak and Limbo, which trace back to Hoare's CSP. This article follows the same map.

- The Oberon line: a lean type system, object orientation without inheritance, module management through packages, and garbage collection
- The CSP line: process communication through channels, the `select` statement, and lightweight processes

# 1. The Algol 60 lineage: the origin of structured programming

## Algol 60 (1960)

Nearly every procedural language today descends directly from Algol 60.

Algol 60 offered block structure through `begin ... end`, local variable scope, and recursion. It also became renowned for defining its grammar formally in BNF (Backus–Naur Form), the notation Backus devised for Algol 58 and Naur refined for Algol 60. From C and Pascal to Java and Go, the basic rules of syntax start here.

## Pascal (1970)

Niklaus Wirth designed Pascal for teaching and for writing systems.

Pascal simplified Algol 60 and added a strict type system, including the `record` type. It also earned a reputation for fast compilation. Pascal spread structured programming across the world and stood beside C as a standard language.

# 2. Modularity and the Oberon lineage: Wirth's pursuit

This part follows how Wirth resisted the growing complexity of software and pursued simplicity and modularity.

## Modula-2 (1978)

Modula-2 succeeded Pascal and put the module at the center of the language, exactly as its name suggests.

It split the definition part (`DEFINITION MODULE`) from the implementation part (`IMPLEMENTATION MODULE`), which enabled separate compilation and encapsulation. It also supported simple concurrency through coroutines.

## Oberon (1987)

Niklaus Wirth designed Oberon by stripping Modula-2 down to the extreme. He and Jürg Gutknecht then built the matching tiny operating system, the Oberon System.

Wirth removed the complex features and added only two things: type extension, which resembles single inheritance, and garbage collection. Oberon marks the peak of his philosophy: gain the most expressive power from the fewest features.

## Object Oberon (1989) and Oberon-2 (1991)

Two extensions brought full object orientation to Oberon. Mössenböck, Templ, and Griesemer built Object Oberon; Mössenböck and Wirth built Oberon-2. Oberon-2 in particular settled on type-bound procedures with a receiver — that is, methods.

One choice stands out: they added no `class` construct. Instead, they attached methods directly to Oberon's type extension.

### The connection to Go

Go's design carries a clear mark of the Oberon line.

- It holds no `class` and binds methods to structs
- It uses type embedding instead of inheritance
- It manages modules through packages

The clearest link runs through a person. Robert Griesemer, one of Go's three creators, studied under Wirth and Mössenböck at ETH Zurich and co-authored Object Oberon (1989). An author of the Oberon line joined the Go design directly.

# 3. CSP and the concurrency lineage: the direct ancestor of Go's concurrency

Goroutines and channels, the defining features of Go, grew out of a concurrency model that centered on Bell Labs.

## CSP (1978)

Tony Hoare first described CSP (Communicating Sequential Processes) in a 1978 paper. CSP is a formal model of concurrency — a member of the process-calculus family. Hoare presented the 1978 version much like a concurrent programming language, then he and Roscoe later refined it into a full theory (his 1985 book). Rather than a language you code in day to day, CSP is a way to reason about concurrency.

Its essence is simple: instead of sharing memory, independent processes coordinate by passing messages. That stance — message passing over threads and locks — carried into Go as a proverb: "Do not communicate by sharing memory; instead, share memory by communicating."

## Squeak (1985)

Luca Cardelli and Rob Pike at Bell Labs built Squeak to express, in a programming language, the concurrency of a user interface that handles input devices such as mice and keyboards. They presented it as "a language for communicating with mice."

Squeak stands as an early experiment that expressed the CSP model as a programming language. Note that it differs from the Smalltalk implementation of the same name.

## Newsqueak (1989)

Rob Pike built Newsqueak on top of Squeak as a more practical concurrency language.

Its syntax sits close to C. Its key advance was to make channels first-class values: unlike in CSP and Squeak, a Newsqueak program can store a channel in a variable, pass it to a function, and even send it over another channel. It could also create processes and channels dynamically and choose among several channel communications at once — the direct ancestor of Go's `select`. Each of these traits appears in Go today.

## Alef (1992)

Phil Winterbottom at Bell Labs designed Alef for Plan 9, the next-generation operating system project. Alef appeared around 1992, and its language reference shipped with the second edition of Plan 9 in 1995. It expressed Newsqueak's channel-based CSP concurrency (`proc`, `task`, `chan`) in a compiled, C-like language.

Yet Alef carried a fatal weakness: it had no automatic memory management, that is, no garbage collection. Pike and others urged Winterbottom to add garbage collection, but it never happened. Manual memory management sits poorly with concurrency, and maintaining a variant language across many architectures proved hard. Plan 9 dropped Alef in its third edition, and its concurrency model carried over into a thread library for C (libthread).

## Limbo (1995)

Limbo followed as the direct successor of Alef. Sean Dorward, Phil Winterbottom, and Rob Pike built it for Inferno, a distributed operating system.

Limbo carried the CSP-style concurrency that Newsqueak and Alef had refined, and it added the automatic garbage collection that Alef lacked. Its typed channels, strong typing, and modularity show through clearly in Go's design. This is why the official Go FAQ names Newsqueak and Limbo as the ancestors of Go's concurrency.

# Conclusion: every lineage leads to Go

These languages form the lineage of technology that Go's designers passed through and refined on their way to Go.

The two lineages finally met, literally, in the Go team. From the Oberon line came Robert Griesemer; from CSP and Bell Labs came Rob Pike and Ken Thompson. They began designing Go in 2007.

Go built on the simple syntax of C. Onto it, the team brought a lean type system and package-based module management from Oberon and Oberon-2, and the CSP concurrency that grew from Newsqueak through Limbo. Then a modern runtime, with garbage collection, solved the memory management that Alef could not.

```mermaid
flowchart TB
    algol["Algol 60 (1960)"]
    c["C (1972)"]
    pascal["Pascal (1970)"]
    modula["Modula-2 (1978)"]
    oberon["Oberon (1987)"]
    oberon2["Object Oberon (1989) / Oberon-2 (1991)"]
    csp["CSP (1978)"]
    squeak["Squeak (1985)"]
    newsqueak["Newsqueak (1989)"]
    alef["Alef (1992)"]
    limbo["Limbo (1995)"]
    go["Go (2009)"]

    algol --> pascal --> modula --> oberon --> oberon2
    algol --> c
    csp --> squeak --> newsqueak --> alef --> limbo

    oberon2 -->|"types, methods, packages, GC"| go
    c -->|"syntax"| go
    limbo -->|"CSP-style concurrency (goroutine / channel)"| go
```

Behind a single line of Go lies more than half a century of language design. Once you know the lineage, the intent behind Go's design comes into sharper view.

# References

- [Go FAQ — What are Go's ancestors? / Why build concurrency on the ideas of CSP?](https://go.dev/doc/faq)
- [Rob Pike, "Origins of Go concurrency style" (OSCON 2010)](https://www.youtube.com/watch?v=3DtUzH3zoFo)
- [Russ Cox, "Bell Labs and CSP Threads"](https://swtch.com/~rsc/thread/)
- [Wikipedia: Communicating sequential processes](https://en.wikipedia.org/wiki/Communicating_sequential_processes)
- [Wikipedia: Newsqueak](https://en.wikipedia.org/wiki/Newsqueak)
- [Wikipedia: Alef (programming language)](https://en.wikipedia.org/wiki/Alef_(programming_language))
- [Wikipedia: Limbo (programming language)](https://en.wikipedia.org/wiki/Limbo_(programming_language))
- [Mössenböck, Templ, Griesemer, "Object Oberon: An Object-Oriented Extension of Oberon" (ETH TR 109, 1989)](https://www.research-collection.ethz.ch/handle/20.500.11850/68697)
- [Cardelli, Pike, "Squeak: a language for communicating with mice" (SIGGRAPH 1985)](http://ordiecole.com/squeak/cardelli_squeak1985.pdf)
