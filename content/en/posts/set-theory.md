---
title: About Sets
slug: set-theory
date: 2024-07-06T00:00:00Z
author: bmf-san
categories:
  - Mathematics
tags:
  - Discrete Mathematics
  - Sets
description: Summarizing the basics of sets.
translation_key: set-theory
---



# Overview
Summarizing the basics of sets.

# What is a Set
In set theory, a set is a collection of elements that satisfy specific conditions.

The elements contained in a set are referred to as members (in this article, they are referred to as elements).

# Software Engineers and Sets
For software engineers, sets are a fundamental concept in data structures and algorithms. Concepts of sets are related to arrays, maps, graph theory, combinatorial theory, and more.

In RDB, set theory is a very important concept, and relations, tuples, and SQL can be considered as sets themselves.

cf. [bmf-tech.com - Practical Database Introduction from Theory ~ Efficient SQL with Relational Model](https://bmf-tech.com/posts/%e7%90%86%e8%ab%96%e3%81%8b%e3%82%89%e5%ad%a6%e3%81%b6%e3%83%87%e3%83%bc%e3%82%bf%e3%83%99%e3%83%bc%e3%82%b9%e5%ae%9f%e8%b7%b5%e5%85%a5%e9%96%80%20~%e3%83%aa%e3%83%ac%e3%83%bc%e3%82%b7%e3%83%a7%e3%83%8a%e3%83%ab%e3%83%a2%e3%83%87%e3%83%ab%e3%81%ab%e3%82%88%e3%82%8b%e5%8a%b9%e7%8e%87%e7%9a%84%e3%81%aaSQL)

Set theory is also related to logic, and sets are sometimes used as expressions of logic.

Additionally, it is useful for organizing abstract thoughts about problems, thus relating to foundational skills for problem-solving.

Sets are a fundamental concept in software engineering, allowing optimal handling of data structures and algorithms. By utilizing them as an element for problem-solving, one can also develop problem-solving skills.

# Basic Sets
## a ∈ A
a is an element of set A.

```
A = {a, b, c, ...}
a ∈ A
```

![a ∈ A](https://github.com/bmf-san/bmf-tech-client/assets/13291041/f1cda391-2848-4c58-aa06-91671a162038)

## a ∉ A
a is not an element of set A.

```
A = {a, b, , ...}
a ∉ A
```

![a ∉ A](https://github.com/bmf-san/bmf-tech-client/assets/13291041/4921fef3-7268-433a-b14c-3ca1de2dc011)

## A ⊂ B
Set A is a subset of set B. A = B is also applicable.

```
A = {1, 2, 3}
B = {1, 2, 3, 4, 5}
A ⊂ B
```

![A ⊂ B](https://github.com/bmf-san/bmf-tech-client/assets/13291041/3f9e7fcb-6645-4f11-9c13-34677c35aded)

## A ⊃ B
Set B is a subset of set A. Equivalent to B ⊂ A.

```
A = {1, 2, 3, 4, 5}
B = {2, 3}
A ⊃ B
```

![A ⊃ B](https://github.com/bmf-san/bmf-tech-client/assets/13291041/33ef3a21-f835-496d-b842-c913696e5b03)

## φ (Empty Set)
A set with no elements.

```
φ = {}
```

![φ (Empty Set)](https://github.com/bmf-san/bmf-tech-client/assets/13291041/66542186-c95f-41f1-a555-b1b150b03803)

## A ∪ B (Union Set)
A set that combines sets A and B. Elements belong to either set A or set B, or both. (≒ Belong to at least one of the sets.)

```
A = {1, 2, 3}
B = {3, 4, 5}
A ∪ B = {1, 2, 3, 4, 5}
```

![A ∪ B (Union Set)](https://github.com/bmf-san/bmf-tech-client/assets/13291041/d49588a3-e0ce-45a3-9353-c6399456459f)

## A ∩ B (Intersection Set)
The common set of sets A and B. Elements belong to both set A and set B.

```
A = {1, 2, 3}
B = {3, 4, 5}
A ∩ B = {3}
```

![A ∩ B (Intersection Set)](https://github.com/bmf-san/bmf-tech-client/assets/13291041/c54c00fa-eb5b-436d-bb39-9aa53bf99c47)

## A × B (Cartesian Product)
Pairs formed by taking one element each from sets A and B.

```
A = {1, 2}
B = {x, y}
A × B = {(1, x), (1, y), (2, x), (2, y)}
```

## A \ B (Difference Set)
A set obtained by removing elements belonging to set B from set A.

```
A = {1, 2, 3, 4}
B = {3, 4, 5}
A \ B = {1, 2}
```

![A \ B (Difference Set)](https://github.com/bmf-san/bmf-tech-client/assets/13291041/59ce55d8-3af3-4dd8-aac5-76c6c98da19c)

## Complement Set
As a symbol, when a set is A, a bar is placed above A.

When set A is a subset of the universal set U, the set obtained by removing set A from the universal set U.

```
U = {1, 2, 3, 4, 5}  # Universal set
A = {1, 2, 3}
A' = {4, 5}
```

![Complement Set](https://github.com/bmf-san/bmf-tech-client/assets/13291041/4d55934d-4d7a-42ba-afb6-e7fea1083ae1)

# References
- [ja.wikipedia.org - Set Theory](https://ja.wikipedia.org/wiki/%E9%9B%86%E5%90%88%E8%AB%96)
- [www2.toyo.ac.jp - Symbols Related to Sets](http://www2.toyo.ac.jp/~y-mizuno/Lang/appendix/symbols.pdf)
- [juken-mikata.net - 6 Symbols and 3 Laws You Must Remember for Sets](https://juken-mikata.net/how-to/mathematics/shugou.html)
