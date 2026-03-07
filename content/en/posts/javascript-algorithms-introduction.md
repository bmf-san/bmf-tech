---
title: Getting Started with Algorithms in JavaScript
slug: javascript-algorithms-introduction
date: 2018-07-13T00:00:00Z
author: bmf-san
categories:
  - Algorithms and Data Structures
tags:
  - Binary Search
  - Linear Search
  - Bubble Sort
  - Selection Sort
translation_key: javascript-algorithms-introduction
---

# Overview
Learn algorithms with JavaScript.

# Search Algorithms

## Linear Search

An algorithm that compares data in a list or array sequentially from the beginning.

The process is repeated for the length of the array, and it stops when the target data is found. The further the target data is located towards the end, the slower the process becomes.

```js
const targetData = 5;
const data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

(function () {
	for (let i = 0; i < data.length; i++) {
  	if (targetData == data[i]) {
    	alert(i + '番目でデータを発見');
      return;
    }
  }
  
  alert('データがありません');
}());
```

## Binary Search

An algorithm that searches for data in a sorted list or array by narrowing down the search range based on comparisons with the median.

First, the median is calculated, and the comparison between the target data and the median is repeated until the start of the search range exceeds the end of the search range.

```js
const targetData = 5;
const data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

let head = 0;  // Start of the search range
let tail = data.length;  // End of the search range

(function () { 
  while (head <= tail) {
    let center = Math.floor((head + tail) / 2);
    
    if (data[center] == targetData) {
      alert('配列の' + center + '番目でデータを発見'); 
      return;
    } else if (data[center] < targetData) {
      head = center + 1;
    } else {
      tail = center - 1;
    }
  }
  
  alert('データがありません');	
}());
```

If the median is smaller than the target data, the start of the range is updated to median + 1. If the median is larger, the end of the range is updated to median - 1. It might seem confusing at first, but following the process step by step (1st iteration, 2nd iteration, 3rd iteration, etc.) will make it easier to understand.

# Sorting Algorithms

## Selection Sort

An algorithm that sorts data sequentially from the beginning.

```js
const data = [10, 1, 5, 7, 8, 2];

for (let i = 0; i < data.length-1; i++) { // Repeat the process for the length of the array
  let min = data[i];
  let head = i;

  for (let headNext = i+1; headNext < data.length; headNext++) {
    if (min > data[headNext]) {
        min = data[headNext];
        head = headNext;
    }
  }

  let tmp = data[i];
  data[i] = data[head];
  data[head] = tmp;
}

console.log(data);
```

This algorithm is also easier to understand if you follow the process step by step (1st iteration, 2nd iteration, 3rd iteration, etc.).

The process involves repeatedly comparing the first value with the values from the first + 1 to the last. If a smaller value is found, it replaces the first value. This process is repeated for the length of the array.

## Bubble Sort

An algorithm that sorts data by comparing adjacent values and swapping them.

```js
const data = [9, 7, 1, 10, 5];

for (let i = 0; i < data.length; i++) {
  for (let dataNext = data.length-1; dataNext > i; dataNext--) { 
    if (data[dataNext] < data[dataNext-1]) {
      let tmp = data[dataNext];
      data[dataNext] = data[dataNext-1];
      data[dataNext-1] = tmp;
     }
   }
}

console.log(data);
```
