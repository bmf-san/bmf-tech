---
title: Starting Algorithms with JavaScript
description: 'Master fundamental JavaScript algorithms including linear and binary search, plus selection and bubble sort implementations.'
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

It repeats the process for the length of the array and stops when it reaches the target data. The further back the target data is, the slower the process becomes.

```js
const targetData = 5;
const data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

(function () {
	for (let i = 0; i < data.length; i++) {
		if (targetData == data[i]) {
			alert(i + ' found the data');
			return;
		}
	}
	
	alert('Data not found');
}());
```

## Binary Search

An algorithm that narrows down the search range by comparing the target data with the median in a sorted list or array.

First, find the median and repeat the comparison of the target data with the median until the start of the search range exceeds the end.

```js
const targetData = 5;
const data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

let head = 0;  // Start of search range
let tail = data.length;  // End of search range

(function () { 
	while (head <= tail) {
		let center = Math.floor((head + tail) / 2);
		
		if (data[center] == targetData) {
			alert('Found data at index ' + center); 
			return;
		} else if (data[center] < targetData) {
			head = center + 1;
		} else {
			tail = center - 1;
		}
	}
	
	alert('Data not found');	
}());
```

If the median is smaller than the target data, set the start value to median + 1; if larger, set the end value to median - 1. It can be a bit confusing, but it becomes clear when you think through the first, second, third... iterations.

# Sorting Algorithms

## Selection Sort

An algorithm that sorts data sequentially from the beginning.

```js
const data = [10, 1, 5, 7, 8, 2];

for (let i = 0; i < data.length-1; i++) { // Repeat for the length of the array
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

This is also easier to understand when you think through the first, second, third... iterations.

Repeat the comparison of the first value with the values from the first + 1 to the end, and if it's smaller than the first value, replace the first value, repeating this for the length of the array.

## Bubble Sort

An algorithm that sorts data by comparing adjacent values.

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