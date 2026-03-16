---
title: punditのコードリーディング
description: punditのコードリーディング
slug: pundit-code-reading
date: 2024-10-22T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - pundit
  - Ruby
translation_key: pundit-code-reading
---


# 概要
punditのコードリーディングをする。

# 準備
1. punditのリポジトリをクローンする
   - `git clone git@github.com:varvet/pundit.git`

# コードリーディング
権限を適用するときに使う`authorize`メソッドを見ていく。

1. authorize
- [varvet/pundit/blob/main/lib/pundit.rb#L75](https://github.com/varvet/pundit/blob/main/lib/pundit.rb#L75)
  - モジュールのクラスメソッドとして定義されている
- [varvet/pundit/blob/main/lib/pundit/context.rb#L55](https://github.com/varvet/pundit/blob/main/lib/pundit/context.rb#L55)
  - 実際の処理は`Pundit::Context`クラスの`authorize`メソッドで行われている
  - ここでpolicyをチェックして、権限を判定している

拍子抜けするくらいシンプルな実装だった。
