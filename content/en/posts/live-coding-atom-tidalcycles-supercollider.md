---
title: Getting Started with Live Coding using Atom, TidalCycles, and SuperCollider
slug: live-coding-atom-tidalcycles-supercollider
date: 2018-06-11T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Atom
  - Git
  - Haskell
  - Homebrew
  - TidalCycles
translation_key: live-coding-atom-tidalcycles-supercollider
---

# Overview
I wanted to create music through programming, so I ventured into sound programming.

# Preparation
- Mac
- Atom
- Git
- Homebrew
- Haskell
- TidalCycles
- SuperCollider

Assuming that Mac, Git, Atom, and Homebrew are already set up, let's proceed.

# Installing Haskell and TidalCycles with Homebrew

```
brew install ghc
brew install cabal-install
cabal update
cabal install cabal-install
cabal install tidal
```

# Installing the tidalcycles package in Atom
Install the [Atom - tidalcycles](https://atom.io/packages/tidalcycles) package in Atom.

# Installing SuperCollider
Install the Current Version from [SuperCollider](https://supercollider.github.io/download).

Once the installation is complete, launch SuperCollider and execute the following command. (command+enter)

```
include("SuperDirt")
```

# Functionality Check
## Launching SuperCollider
Start SuperDirt with `SuperDirt.start`.

## Launching TidalCycles
Launch tidalcycles in Atom. Create a file with the extension `.tidal` and execute the following command. (eval in tidalcycles)

```
d1 $ sound "bd sn"
```

# References
- [TidalCycles](https://tidalcycles.org/)
- [TidalCycles 0.8 Release](https://qiita.com/yoppa/items/b195d4014de63686b2e0)