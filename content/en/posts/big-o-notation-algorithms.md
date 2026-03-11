---
title: Order Notation and How to Determine Algorithm Complexity
slug: big-o-notation-algorithms
date: 2018-04-18T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Big O Notation
description: An overview of the basics of calculating algorithm performance using Big O notation and complexity.
translation_key: big-o-notation-algorithms
---



# Overview
This post summarizes the foundational knowledge of calculating algorithm performance using Big O notation and complexity.

# What is Complexity (Order)?
- A metric that represents the **rate at which execution time increases as data volume increases**.
    - Time Complexity
        - Processing time
    - Space Complexity
        - Memory usage

# Big O/Big θ/Big Ω
These notations describe computation time, but here we summarize their academic differences.

- Big O
  - Upper bound of computation time
- Big θ
  - Lower bound of computation time
- Big Ω
  - Both O and Ω

# Best/Worst/Average Case
A summary of the three patterns representing algorithm execution time.

- Best Case
  - The lower bound of complexity. A case where all element values are equal. It is rarely discussed as it is often executable in O(₁) and significantly differs from the worst or average case.
- Worst Case
  - The upper bound of complexity.
- Average Case
  - The average complexity. A case where element values are average.

In many algorithms, the worst case and average case often coincide.

# Big O Notation

Representative notations are summarized in order of shortest processing time.

|Big O Notation|Name in Computational Theory|Overview|
|---|---|---|
|O(₁)|Constant Time|Processing time does not increase even if data volume increases|
|O(log n)|Logarithmic Time|Computation time barely increases even if data volume increases. The rate of increase in complexity becomes smaller.|
|O(n)|Linear Time|Processing time increases in proportion to the increase in data volume|
|O(n log n)|Quasi-linear, Linear Logarithmic Time|Slightly heavier than O(n)|
|O(n²)|Quadratic Time|Processes that examine all combination pairs of elements. The rate of increase in complexity becomes larger as data volume increases|
|O(n³)|Polynomial Time|Triple loop|
|O(kⁿ)|Exponential Time|Processes that obtain all combinations of elements|
|O(n!)|Factorial Time|Time proportional to the factorial of n|

# How to Determine Complexity
Calculate the number of steps and determine complexity based on their sum.
When determining complexity, the following two points are of low importance and can be omitted.

- Omit non-maximum degree terms
  - O(n²+n)
    - Consider as O(n²)
- Omit coefficients
  - O(2n)
    - Consider as O(n)

Whether to add or multiply execution times in step calculations depends on whether the processes occur simultaneously or not.

If they do not occur simultaneously, add execution times.
```js
for (condition) {
  // do something
}

for (condition) {
  // do something
}
```

If they occur simultaneously, multiply execution times.
```js
for (condition) {
  for (condition) {
    // do something
  }
}
```

## Examples

### Linear Search

```js
const targetData = 4; // Executed once
const data = [1, 2, 3, 4, 5]; // Executed once

for (let i = 0; i < data.length; i++) {
	if (targetData == data[i]) {
  	console.log(`${i}番目でデータを発見した`); // Executed data.length times → n times
    return;
  }
}

console.log('目的のデータはない'); // Executed once
```

In the above code, the total number of steps is 1+1+n+1=3n.
Excluding coefficients, the complexity is O(n).

### Nested for Loop

```js
const data = [1, 2, 3, 4, 5]; // Executed once

for (let i = 0; i < data.length; i++) {
	console.log(`${i}回目の処理`); // Executed once
	for (let j = 0; j < data.length; j++) {
		console.log(j); // Executed 4 * 4 times → n² times
  }
}
```

In this case, the number of steps is 1+1+n²=2n², so excluding coefficients, the complexity is O(n²).

### Algorithms with Logarithmic Complexity

```js
const n = 10;  // Executed once

for (let i = 0; i < n; i = i * 2) {
  console.log(i++); // Executed log2ⁿ times
}
```

When n is 1
Loop count 1

When n is 4
Loop count 2

When n is 8
Loop count 3

The loop count is determined by log2ⁿ.
The number of steps is 1+log2ⁿ, so excluding various factors, the complexity is log n.


# Reference Links
- [About Complexity Order](https://qiita.com/asksaito/items/59e0d48408f1eab081b5)
- [[For Beginners] How to Determine Program Complexity](https://qiita.com/cotrpepe/items/1f4c38cc9d3e3a5f5e9c)
- [[For Beginners] How to Determine Program Complexity](https://qiita.com/cotrpepe/items/1f4c38cc9d3e3a5f5e9c)
- [Comprehensive Guide to Determining Complexity Order! ~ Where Does log Come From ~](https://qiita.com/drken/items/872ebc3a2b5caaa4a0d0)
- [techscore - A Basic Guide to Algorithms and Complexity for New Developers](http://www.techscore.com/blog/2016/08/08/%E9%96%8B%E7%99%BA%E6%96%B0%E5%8D%92%E3%81%AB%E6%8D%A7%E3%81%90%E3%80%81%E5%9F%BA%E6%9C%AC%E3%81%AE%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%A8%E8%A8%88%E7%AE%97%E9%87%8F/)

# Reference Books
- [Cracking the Coding Interview: 189 Programming Questions and Solutions](https://amzn.to/2yOgQ08)
  - Easy-to-understand examples about Big O notation.
- [Algorithm Picture Book: 26 Algorithms Explained with Pictures](https://amzn.to/2SVXD3p)
- [Algorithms and Data Structures for Programming Contests](https://amzn.to/2WTJm8C)
- [Let's Start with Algorithms](https://amzn.to/35SiHxa)
- [Web+DB Press vol.91 2016](https://amzn.to/2y0zcea)
