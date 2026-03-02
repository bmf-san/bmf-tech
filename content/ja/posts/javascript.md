---
title: "JavaScriptで始めるアルゴリズム"
slug: "javascript"
date: 2018-07-13
author: bmf-san
categories:
  - "アルゴリズムとデータ構造"
tags:
  - "バイナリーサーチ"
  - "リニアサーチ"
  - "バブルソート"
  - "セレクションソート"
draft: false
---

# 概要
JavaScriptでアルゴリズムを学ぶ。

# サーチのアルゴリズム

## リニアサーチ

リストや配列のデータに対して、先頭から順番に比較を行っていくアルゴリズム。

配列の長さ分処理を繰り返し、目的のデータに到達したら処理を終了する。
目的とするデータが後ろにあるほど処理が遅くなる。

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

## バイナリーサーチ

ソート済みのリストや配列に対し、中央値との大小関係を判定条件とし、探索範囲を狭めながらデータを探索していく。

初めに中央値を求め、目的のデータと中央値の大小比較を探索範囲の先頭が探索範囲の後尾を上回るまで繰り返す。

```js
const targetData = 5;
const data = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

let head = 0;  // 探索範囲の先頭
let tail = data.length;  // 探索範囲の後尾

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

目的のデータより中央値のほうが小さい場合は、中央値＋１を先頭の値とし、大きい場合は、中央値ー１を後尾の値とする。
ちょっと混乱するが、処理の1回目、2回目、3回目...と順を追って考えてみるとすぐ理解できる。

# ソートのアルゴリズム

## セレクションソート

先頭から順番に並べ替えていくアルゴリズム。

```js
const data = [10, 1, 5, 7, 8, 2];

for (let i = 0; i < data.length-1; i++) { // 配列の長さ分処理を繰り返す
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

こちらも処理の1回目、2回目、3回目...と順を追って考えてみると理解しやすい。

先頭の値と先頭＋1から末尾までの値の比較を繰り返し、先頭の値より小さければ、先頭の値を置き換える、という処理を配列の長さ分繰り返す。

## バブルソート

隣合う値同士を比較してデータを並べ替えるアルゴリズム。

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
