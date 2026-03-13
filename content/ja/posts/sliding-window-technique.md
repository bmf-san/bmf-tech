---
title: 尺取り法について
description: '尺取り法（Two Pointer Technique）とは何か。左右のインデックスを使った探索の仕組み、O(N²)からO(N log N)への計算量改善の仕組みをGoコードの例で解説します。'
slug: sliding-window-technique
date: 2023-08-01T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - 尺取り法
  - Two-Pointer Approach
translation_key: sliding-window-technique
---


# 概要
尺取り法についてまとめる。

英語だと、Two Pointer ApproachまたはTwo Pointer Techniqueと呼ばれる。

# 尺取り法とは
データセット（数列や文字列など）の右端と左端のインデックスを保持して、条件によって左右のインデックスを移動させることで、条件を満たすデータを探索するアルゴリズム。

特定の条件を満たすデータを区間の中から探索したいような時に役立つ。

# 例題
配列nの中から指定された数値m未満になる数字のペアがいくつあるかを求める関数を書く。

ex. n = [1, 3, -1, 2] m = 4 の場合、(1, -1）、（1, 2）、（3, -1）、（-1, 2）の4つのペアが該当するので4を返す。


関数を愚直に実装した場合のコードは次の通り。

```go
func findPairs(n []int, m int) int {
    rslt := 0
    for i := 0; i < len(n); i++ {
        for j := i+1; j < len(n); j++ {
            if (n[i] + n[j]) < m {
                rslt++
            }
        }
    }
    return rslt
}
```

これだとO(N^2)の計算量になってしまうので、尺取り法を使ってO(N log N)にする。

```go
func findPairs(n []int, m int) int {
    sort.Ints(n)
    cnt, l, r := 0, 0, len(n)-1
    for l < r {
        if n[l] + n[r] < m {
            cnt += r-l
            l++
            continue
        }
        r--
    }
    return cnt
}
```

左右のインデックスが重なり合うまで繰り返すことで、条件を満たすペアを全て探索できる。

ソートを行っているのは、効率的にペアを見つけやすいようにするためである。

r-lが直感に反するような気がするが、ペアを探すためのコードなので、配列内の値そのものではなくインデックスに注目してペアをカウントすれば良いため、このようなコードになる。

# 参照
- [algodaily.com - The Two Pointer Technique](https://algodaily.com/lessons/using-the-two-pointer-technique/go)
- [www.geekforgeeks.org - Two Pointers Technique](https://www.geeksforgeeks.org/two-pointers-technique/)
- [www.codingninjas.com - Two Pointer Approach](https://www.codingninjas.com/studio/library/what-is-a-two-pointer-technique)
- [jaigotec.com - [アルゴリズム(Ruby)]尺取り法の解説](https://jaigotec.com/algorithm_two-pointer-techinique/)
- [qiita.com - しゃくとり法 (尺取り法) の解説と、それを用いる問題のまとめ
](https://qiita.com/drken/items/ecd1a472d3a0e7db8dce)
- [na.fuis.u-fukui.ac.jp - しゃくとり法](https://na.fuis.u-fukui.ac.jp/~hirota/course/2022_Exp2_Programming/03-1_Shakutori.pdf)
- [paiza.hatenablog.com - 【累積和、しゃくとり法】初級者でも解るアルゴリズム図解](https://paiza.hatenablog.com/entry/2015/01/21/%E3%80%90%E7%B4%AF%E7%A9%8D%E5%92%8C%E3%80%81%E3%81%97%E3%82%83%E3%81%8F%E3%81%A8%E3%82%8A%E6%B3%95%E3%80%91%E5%88%9D%E7%B4%9A%E8%80%85%E3%81%A7%E3%82%82%E8%A7%A3%E3%82%8B%E3%82%A2%E3%83%AB%E3%82%B4)
- [www.kumilog.net - 尺取り法](https://www.kumilog.net/entry/two-pointers)
- [scrapbox.io - しゃくとり法](https://scrapbox.io/pocala-kyopro/%E3%81%97%E3%82%83%E3%81%8F%E3%81%A8%E3%82%8A%E6%B3%95)
