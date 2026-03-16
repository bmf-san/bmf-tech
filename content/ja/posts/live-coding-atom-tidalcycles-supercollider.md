---
title: Atom×TidalCycles×SuperColliderでライブコーディングをはじめる
description: Atom×TidalCycles×SuperColliderでライブコーディングをはじめる
slug: live-coding-atom-tidalcycles-supercollider
date: 2018-06-11T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Atom
  - Git
  - Haskell
  - homebrew
  - tidalcycles
translation_key: live-coding-atom-tidalcycles-supercollider
---


# 概要
プログラミングで音楽をつくってみたいと思い、音響プログラミングに手を出してみた。

#  準備
- Mac
- Atom
- Git
- Homebrew
- Haskell
- Tidalcycles
- SuperCollider

Mac、Git、Atom、Homebrewは既に用意されている前提で話を進める。

# HomebrewでHaskellとTidalCyclesをインストールする

```
brew install ghc
brew install cabal-install
cabal update
cabal install cabal-install
cabal install tidal
```

# Atomにtidalcyclesのパッケージをインストールする
[Atom - tidalcycles](https://atom.io/packages/tidalcycles)をAtomにインストール。

# SuperColliderをインストールする
[SuperCollider](https://supercollider.github.io/download)からCurrent Versionをインストール。

インストールが完了したら、SuperColliderを起動して以下のコマンドを実行する。（command+enter）

```
include("SuperDirt")
```

# 動作確認
## SuperColiderを起動
`SuperDirt.start`でSuperDirtを起動。

## TidalCyclesを起動
Atomでtidalcyclesを起動する。
`.tidal`を拡張子としてファイルを作成し、以下のコマンドを実行する。（tidalcyclesのeval）

```
d1 $ sound "bd sn"
```

# 参考
- [TidalCycles](https://tidalcycles.org/)
- [TidalCycles 0.8 リリース](https://qiita.com/yoppa/items/b195d4014de63686b2e0)

