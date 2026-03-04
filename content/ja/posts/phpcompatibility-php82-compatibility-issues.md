---
title: "PHPCompatibilityを使ってPHP8.2の互換性チェックを行う際にハマったこと"
slug: "phpcompatibility-php82-compatibility-issues"
date: 2024-04-14
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "PHP CodeSniffer"
  - "PHPCompatibility"
draft: false
---

[PHPCompatibilitiy](https://github.com/PHPCompatibility/PHPCompatibility)を使ってPHP8.2の互換性チェックを行う際にハマったことのメモ。

# ハマったこと
2023年10月現在、latestである[9.3.5](https://github.com/PHPCompatibility/PHPCompatibility/releases/tag/9.3.5)をインストールしてもPHP8.2の互換性チェックはできない。

9.3.5は2019年リリースで、最近のPHPのバージョンにまだ対応していないらしい・・・

# 解決策
もしかして開発止まってる？と一瞬思ったが、その様子はない。

developにはコミットが積まれているのでどうやらdevelopを使えばよいらしい。

cf. [Should I use develop or 9.3.5 sniffs? #1653](https://github.com/PHPCompatibility/PHPCompatibility/issues/1653)

# 所感
100%の互換性チェックをサポートしているわけではないと思うが、便利なツールなので今後も使っていきたい。


