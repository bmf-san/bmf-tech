---
title: 集合について
description: "理解する集合論の基礎。要素・部分集合・和集合・積集合、データ構造・RDB理論、論理表現への応用と問題解決力養成。"
slug: set-theory
date: 2024-07-06T00:00:00Z
author: bmf-san
categories:
  - 数学
tags:
  - 離散数学
  - 集合
translation_key: set-theory
---


# 概要
集合の基本についてまとめる。

# 集合とは
集合論における集合とは、特定条件を満たす要素の集まりのこと。

集合に含まれる要素のことを元という。（当記事内では要素と表記する。）

# ソフトウェアエンジニアと集合
ソフトウェアエンジニアにとって集合は、データ構造やアルゴリズムの基礎概念である。配列やマップ、グラフ理論や組み合わせ理論など集合の概念が関係している。

RDBにおいては、集合論が非常に重要な概念になっており、リレーションやタプル、SQLなどは集合そのものとも言える。

cf. [bmf-tech.com - 理論から学ぶデータベース実践入門 ~リレーショナルモデルによる効率的なSQL](https://bmf-tech.com/posts/%e7%90%86%e8%ab%96%e3%81%8b%e3%82%89%e5%ad%a6%e3%81%b6%e3%83%87%e3%83%bc%e3%82%bf%e3%83%99%e3%83%bc%e3%82%b9%e5%ae%9f%e8%b7%b5%e5%85%a5%e9%96%80%20~%e3%83%aa%e3%83%ac%e3%83%bc%e3%82%b7%e3%83%a7%e3%83%8a%e3%83%ab%e3%83%a2%e3%83%87%e3%83%ab%e3%81%ab%e3%82%88%e3%82%8b%e5%8a%b9%e7%8e%87%e7%9a%84%e3%81%aaSQL)

集合論は論理学にも関連しており、論理の表現として集合が利用されることもある。

また、問題に対する抽象的な思考整理にも役立つため、問題解決のための基礎力にも関係する。

集合はソフトウェアエンジニアリングの基礎概念であり、データ構造やアルゴリズムを最適に取り扱うことができる。また、課題解決のための1要素として役立てることで、課題解決力を養うこともできる。

# 基礎的な集合
## a ∈ A
aは集合Aの要素である。

```
A = {a, b, c, ...}
a ∈ A
```

![a ∈ A](https://github.com/bmf-san/bmf-tech-client/assets/13291041/f1cda391-2848-4c58-aa06-91671a162038)

## a ∉ A
aは集合Aの要素ではない。

```
A = {a, b, , ...}
a ∉ A
```

![a ∉ A](https://github.com/bmf-san/bmf-tech-client/assets/13291041/4921fef3-7268-433a-b14c-3ca1de2dc011)

## Ａ⊂Ｂ
集合Aは集合Bの部分集合である。A=Bも該当する。

```
A = {1, 2, 3}
B = {1, 2, 3, 4, 5}
A ⊂ B
```

![Ａ⊂Ｂ](https://github.com/bmf-san/bmf-tech-client/assets/13291041/3f9e7fcb-6645-4f11-9c13-34677c35aded)

## Ａ⊃Ｂ
集合Bは集合Aの部分集合である。B⊂Aに等しい。

```
A = {1, 2, 3, 4, 5}
B = {2, 3}
A ⊃ B
```

![Ａ⊃Ｂ](https://github.com/bmf-san/bmf-tech-client/assets/13291041/33ef3a21-f835-496d-b842-c913696e5b03)

## φ（空集合）
要素を持っていない集合。

```
φ = {}
```

![φ（空集合）](https://github.com/bmf-san/bmf-tech-client/assets/13291041/66542186-c95f-41f1-a555-b1b150b03803)

## Ａ∪Ｂ（和集合）
集合Aと集合Bを足し合わせた集合。要素は集合Aか集合Bのどちらか、あるいは両方に属している。（≒少なくとも片方の集合に属している。）

```
A = {1, 2, 3}
B = {3, 4, 5}
A ∪ B = {1, 2, 3, 4, 5}
```

![Ａ∪Ｂ（和集合）](https://github.com/bmf-san/bmf-tech-client/assets/13291041/d49588a3-e0ce-45a3-9353-c6399456459f)

## Ａ∩Ｂ（積集合）
集合Aと集合Bの共通集合。要素は集合Aと集合Bの両方に属している。

```
A = {1, 2, 3}
B = {3, 4, 5}
A ∩ B = {3}
```

![Ａ∩Ｂ（積集合）](https://github.com/bmf-san/bmf-tech-client/assets/13291041/c54c00fa-eb5b-436d-bb39-9aa53bf99c47)

## Ａ×Ｂ（直積集合）
集合Aと集合Bから要素を一つずつ取り出して組にしたもの。

```
A = {1, 2}
B = {x, y}
A × B = {(1, x), (1, y), (2, x), (2, y)}
```

## A\B（差集合）
集合Aから集合Bに属する要素を取り除いて得られる集合。

```
A = {1, 2, 3, 4}
B = {3, 4, 5}
A \ B = {1, 2}
```

![A B（差集合）](https://github.com/bmf-san/bmf-tech-client/assets/13291041/59ce55d8-3af3-4dd8-aac5-76c6c98da19c)

## 補集合
記号としては、集合をAとした場合、Aの上にバーがつく。

集合Aが全体集合Uの部分集合であるとき、全体集合Uから集合Aを取り除いて得られる集合。

```
U = {1, 2, 3, 4, 5}  # 全体集合
A = {1, 2, 3}
A' = {4, 5}
```

![補集合](https://github.com/bmf-san/bmf-tech-client/assets/13291041/4d55934d-4d7a-42ba-afb6-e7fea1083ae1)

# 参考
- [ja.wikipedia.org - 集合論](https://ja.wikipedia.org/wiki/%E9%9B%86%E5%90%88%E8%AB%96)
- ~~www2.toyo.ac.jp - 集合に関する記号~~
- [juken-mikata.net - 【集合】必ず覚えなくてはならない６つの記号と３つの法則](https://juken-mikata.net/how-to/mathematics/shugou.html)
