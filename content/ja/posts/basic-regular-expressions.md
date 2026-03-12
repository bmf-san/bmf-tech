---
title: 正規表現の基本
description: 「正規表現の基本」のまとめと読書メモ。重要なポイントと実践的な知見を整理します。
slug: basic-regular-expressions
date: 2018-12-01T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - ERE
  - 正規表現
translation_key: basic-regular-expressions
---


# 概要
正規表現の基本。
ERE(Extended regular expression)で扱える記法の中でよく使いそうなやつをまとめる。

# 文字クラス
## [文字列]
- 文字列の中の任意の一文字と一致
- Ex:
  - [きつね]
    - きつねたぬきねこ
      - き,つ,ね,き,ね


## [^文字列]
- 文字列の中にない任意の一文字と一致
- Ex:
  - [^きつね]
    - きつねたぬきねこ
      - た,ね,こ

## [文字列-文字列]
- 任意の文字範囲にある一文字と一致
- Ex:
  - [あ-ん]
    - きつねたぬきcat
      - き,つ,ね,た,ぬ,き

## \d
- 10進数数字と一致
- Ex:
  - \d
    - りんごが10個
      - 1, 0

## \D
- 10進数数字以外の任意の文字と一致
- Ex:
  - \D
    - りんごが10個
      - り,ん,ご,が,個

## \w
- 全ての半角英数字とアンダースコアと一致
- Ex:
  - \w
    - abc_*
      - a,b,c,_

## \W
- 全ての半角英数字とアンダースコア以外と一致
- Ex:
  - \W
    - abc_*
      - *

## \s
- 空白文字と一致
- Ex:
  - \s
    - a b c
      - 空白文字2つに一致（aとbの間、bとcの間）

## \s
- 空白以外の文字と一致
- Ex:
  - \S
    - a b c
      - a,b,c

# アンカー
## ^
- 行の先頭の文字列で一致
- Ex:
  - ^ありがとう
    - ありがとう友よ
    - 昨日はありがとう ✗

## $
- 行の末尾の文字列で一致
- Ex:
  - ありがとう$
    - ありがとう友よ ✗
    - 昨日はありがとう ○

# グループ化構成体
## (副次式)
- 副次式に一致する文字列をキャプチャする
- Ex:
  - (りり){2}
    - ありりりりがとう ○
    - ありりがとう ✗

#量指定子
## *
- 直前の要素を0回以上繰り返しているときに一致
- 最長一致数量子
- Ex:
  - ab*
    - ab
      - ab
    - abab
      - ab, ab
    - aabb
      - ab
    - abbb
      - abbb
    - a
      - a
    - ba
      - a

## +
- 直前の要素を1回以上繰り返しているときに一致
- 最長一致数量子
- Ex:
  - ab+
    - ab
      - ab
    - abab
      - ab,ab
    - aabb
      - abb
    - abbb
      - abbb
    - a
      - 一致なし
    - ba
      - 一致なし

## ?
- 直前の要素を0回または1回繰り返しているときに一致
- 最長一致数量子
- Ex:
  - ab?
    - ab
      - ab
    - abab
      - ab,ab
    - aabb
      - a, ab
    - abbb
      - ab
    - a
      - a
    - ba
      - a

## *?
- 直前の要素を0回以上繰り返しているときに一致
- 最短一致数量子
- Ex:
  - ab*?
    - ab
      - a
    - abab
      - a,a
    - aabb
      - a,a
    - abbb
      - a
    - a
      - a
    - ba
      - a

## +?
- 直前の要素を1回以上繰り返しているときに一致
- 最短一致数量子
- Ex:
  - ab+?
    - ab
      - ab
    - abab
      - ab,ab
    - aabb
      - ab
    - abbb
      - ab
    - a
      - 一致なし
    - ba
      - 一致なし

## ??
- 直前の要素を0回または1回繰り返しているときに一致
- 最短一致数量子
- Ex:
  - ab??
    - ab
      - a
    - abab
      - a,a
    - aabb
      - a, a
    - abbb
      - a
    - a
      - a
    - ba
      - a

## {n}
- 直前の要素をn回繰り返しているときに一致
- Ex:
  - b{2}
    - abba
      - bb

## {n,}
- 直前の要素をn回以上繰り返しているときに一致
- Ex:
  - b{2,}
    - abbba
      - bbb

## {n, m}
- 直前の要素をn回以上m回以下繰り返しているときに一致
- Ex:
  - b{2,4}
    - abbba
      - bbb
    - abbbba
      - bbbb

# 代替構成体
## |
- 区切られた文字列のいずれか１つに一致
- Ex:
  - ab|cd
    - abcd
      - ab,cd
    - aaccd
      - cd

# 参考
- [正規表現言語 - クイック リファレンス](https://docs.microsoft.com/ja-jp/dotnet/standard/base-types/regular-expression-language-quick-reference)
- [正規表現メモ](http://www.kt.rim.or.jp/~kbk/regex/regex.html#BRE)
- [Qiita - どのUNIXコマンドでも使える正規表現 ](https://qiita.com/richmikan@github/items/b6fb641e5b2b9af3522e)
- [SE学院 - 正規表現](http://itref.fc2web.com/unix/regular-expression.html#BRE)

