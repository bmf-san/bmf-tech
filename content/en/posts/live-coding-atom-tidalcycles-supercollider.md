---
title: Starting Live Coding with Atom×TidalCycles×SuperCollider
description: An in-depth look at Starting Live Coding with Atom×TidalCycles×SuperCollider, covering key concepts and practical insights.
slug: live-coding-atom-tidalcycles-supercollider
date: 2018-06-11T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Atom
  - Git
  - Haskell
  - homebrew
  - tidalcycles
translation_key: live-coding-atom-tidalcycles-supercollider
---



# Overview
I wanted to create music through programming, so I decided to try out sound programming.

# Preparation
- Mac
- Atom
- Git
- Homebrew
- Haskell
- Tidalcycles
- SuperCollider

Assuming Mac, Git, Atom, and Homebrew are already set up.

# Install Haskell and TidalCycles with Homebrew

```
brew install ghc
brew install cabal-install
cabal update
cabal install cabal-install
cabal install tidal
```

# Install the tidalcycles package in Atom
Install [Atom - tidalcycles](https://atom.io/packages/tidalcycles) in Atom.

# Install SuperCollider
Download the Current Version from [SuperCollider](https://supercollider.github.io/download).

Once the installation is complete, launch SuperCollider and execute the following command (command+enter).

```
include("SuperDirt")
```

# Verification
## Launch SuperCollider
Start SuperDirt with `SuperDirt.start`.

## Launch TidalCycles
Launch tidalcycles in Atom.
Create a file with the extension `.tidal` and execute the following command (tidalcycles eval).

```
d1 $ sound "bd sn"
```

# References
- [TidalCycles](https://tidalcycles.org/)
- [TidalCycles 0.8 Release](https://qiita.com/yoppa/items/b195d4014de63686b2e0)
