---
title: Big O Notation and How to Calculate Algorithm Complexity
slug: big-o-notation-algorithms
date: 2018-04-18T00:00:00Z
author: bmf-san
categories:
  - Algorithms
  - Data Structures
tags:
  - Big O Notation
translation_key: big-o-notation-algorithms
---

# Overview
This post summarizes the prerequisites for understanding how to roughly calculate the computational performance of algorithms using Big O notation and complexity.

# What is Computational Complexity (Order)
- An indicator that represents **the ratio of how much the execution time increases in relation to the increase in data volume**.
    - Time Complexity
        - Processing Time
    - Space Complexity
        - Memory Usage

# Big O/Big θ/Big Ω
These describe computational time, but here we summarize the academic differences in meaning.

- Big O
  - Upper limit of computational time
- Big θ
  - Lower limit of computational time
- Big Ω
  - Both O and Ω

# Best/Worst/Average Case
Here we summarize the three patterns that represent the execution time of an algorithm.

- Best Case
  - Lower limit of computational complexity. A case where all element values are equal. Since it is often significantly different from the worst and average cases, it usually runs in O(₁) and is not often discussed.
- Worst Case
  - Upper limit of computational complexity.
- Average Case
  - Average computational complexity. A case with average element values.

In many algorithms, the worst case and average case tend to be the same.

# Big O Notation

Here are some representative complexities sorted by processing time.

| O Notation | Name in Computational Theory | Summary |
|---|---|---|
| O(₁) | Constant Time | Processing time does not increase with data volume |
| O(log n) | Logarithmic Time | Processing time hardly increases even as data volume increases. When it does, the increase is small. |
| O(n) | Linear Time | Processing time increases in proportion to the increase in data volume |
| O(n log n) | Quasi-linear, Linear Logarithmic Time | Slightly heavier than O(n) |
| O(n²) | Quadratic Time | Processing that examines all combinations of pairs from elements. The increase in computational complexity becomes larger as data volume increases |
| O(n³) | Polynomial Time | Triple loop |
| O(kⁿ) | Exponential Time | Processing that retrieves all combinations from elements |
| O(n!) | Factorial Time | Time proportional to n factorial |

# How to Calculate Complexity
Calculate the number of steps and derive the complexity based on the total. When determining complexity, the following two points are of low importance and can be omitted.

- Omit terms other than the highest degree term
  - O(n²+n)
    - Becomes O(n²)
- Omit coefficients
  - O(2n)
    - Becomes O(n)

When calculating steps, whether to add or multiply the execution times depends on whether the processes occur simultaneously or not.

If they do not occur simultaneously, it is a case of adding execution times.
```js
for (condition) {
  // do something
}

for (condition) {
  // do something
}
```

If they occur simultaneously, it is a case of multiplying execution times.
```js
for (condition) {
  for (condition) {
    // do something
  }
}
```

## Example

### Linear Search
```js
const targetData = 4; // 1 execution
const data = [1, 2, 3, 4, 5]; // 1 execution

for (let i = 0; i < data.length; i++) {
	if (targetData == data[i]) {
		console.log(`${i} index found the data`); // executed data.length times → n times
		return;
	}
}

console.log('Data not found'); // 1 execution
```
In the above code, the total number of steps is 1+1+n+1=3n. Omitting the coefficient, the complexity is O(n).

### Nested for Loop
```js
const data = [1, 2, 3, 4, 5]; // 1 execution

for (let i = 0; i < data.length; i++) {
	console.log(`${i} processing`); // executed 1 time
	for (let j = 0; j < data.length; j++) {
		console.log(j); // executed 4 * 4 times → n² times
	}
}
```
In this case, the number of steps is 1+1+n²=2n², so the complexity, excluding the coefficient, is O(n²).

### Algorithms with Logarithmic Complexity
```js
const n = 10;  // 1 execution

for (let i = 0; i < n; i = i * 2) {
  console.log(i++); // executed log2ⁿ times
}
```
When n is 1
Loop count 1

When n is 4
Loop count 2

When n is 8
Loop count 3

The loop count can be determined as log2ⁿ. The number of steps is 1+log2ⁿ, so omitting various factors, the complexity is log n.

# Reference Links
- [About Computational Complexity Order](https://qiita.com/asksaito/items/59e0d48408f1eab081b5)
- [[Beginner's Guide] How to Calculate Program Complexity](https://qiita.com/cotrpepe/items/1f4c38cc9d3e3a5f5e9c)
- [[Beginner's Guide] How to Calculate Program Complexity](https://qiita.com/cotrpepe/items/1f4c38cc9d3e3a5f5e9c)
- [Comprehensive Guide to Calculating Complexity Order! ~ Where log Comes From ~](https://qiita.com/drken/items/872ebc3a2b5caaa4a0d0)
- [techscore - A Guide to Basic Algorithms and Complexity for New Graduates](http://www.techscore.com/blog/2016/08/08/%E9%96%8B%E7%99%BA%E6%96%B0%E5%8D%92%E3%81%AB%E6%8D%A7%E3%81%90%E3%80%81%E5%9F%BA%E6%9C%AC%E3%81%AE%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0%E3%81%A8%E8%A8%88%E7%AE%97%E9%87%8F/)

# Reference Books
- [A Book to Train Programming Skills to Compete Globally ~ 189 Coding Interview Questions and Their Solutions~](https://amzn.to/2yOgQ08)
  - Contains discussions about Big O notation with easy-to-understand examples.
- [Algorithm Illustrated: 26 Algorithms Explained Visually](https://amzn.to/2SVXD3p)
- [Algorithms and Data Structures for Programming Contest Success](https://amzn.to/2WTJm8C)
- [Let's Start with Algorithms](https://amzn.to/35SiHxa)
- [Web+DB Press vol.91 2016](https://amzn.to/2y0zcea)