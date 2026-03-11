---
title: アルゴリズムとデータ構造 - 配列
slug: algorithms-data-structures-array
image: /assets/images/posts/post-202/67912042-f3e03200-fbcb-11e9-8a42-34f28fd474f4.jpg
date: 2019-10-31T00:00:00Z
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - 配列
translation_key: algorithms-data-structures-array
---


# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# 配列
- データを1列に並べたもの
- データへのアクセスは容易だが、追加や削除には時間がかかる
- 配列のデータはメモリの連続した領域に順番に格納される
  - 固定長のメモリを確保する
    - 宣言時に確保（静的確保）
    - 実行時に確保（動的確保）

# 計算時間
配列に格納されているデータ数をnとする。

## データへのアクセス
- O(1)
  - メモリアドレスが添字を使って計算することができるため、データに直接アクセスすることができる。（ランダムアクセス）

## データの追加
- O(n)
  - 追加する箇所より後ろのデータをすべて1つずつずらす必要がある。

## データの削除
- データの追加と同様

# 実装
```golang
package main

import (
	"errors"
	"fmt"
)

// A Array is array implemented by slice.
type Array struct {
	data   []string
	length int // Keep a array memory size
}

// Insert is insert a data to array.
func (a *Array) insert(index int, value string) error {
	if a.length == int(cap(a.data)) {
		return errors.New("a array is full")
	}

	if index != a.length && index >= a.length {
		return errors.New("out of index range")
	}

	// shift data
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}

	// insert a value to target index
	a.data[index] = value

	// update the length
	a.length++

	return nil
}

// delete is delete a target data by index.
func (a *Array) delete(index int) (string, error) {
	if index >= a.length {
		return "", errors.New("out of index range")
	}

	// target value for deleting
	v := a.data[index]

	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}

	// unset
	a.data[a.length-1] = ""

	// update the length
	a.length--

	return v, nil
}

// get is get a target data by index.
func (a *Array) get(index int) (string, error) {
	if index >= a.length {
		return "", errors.New("out of index range")
	}

	// random access
	return a.data[index], nil
}

func main() {
	a := &Array{
		data:   make([]string, 10, 10),
		length: 0,
	}

	cases := []struct {
		index int
		value string
	}{
		{
			index: 0,
			value: "foo",
		},
		{
			index: 1,
			value: "bar",
		},
		{
			index: 2,
			value: "foobar",
		},
	}

	for _, c := range cases {
		if err := a.insert(c.index, c.value); err != nil {
			fmt.Printf("index: %v value: %v is error. %v\n", c.index, c.value, err)
		}
	}

	if s, err := a.delete(2); err != nil {
		fmt.Printf("index: 0 is error. %v\n", err)
	} else {
		fmt.Printf("%v is deleted.", s)
	}

	if r, err := a.get(0); err != nil {
		fmt.Printf("index: 0 is error. %v", err)
	} else {
		fmt.Printf("%v", r)
	}
}
```

- 構造体Arrayには配列のデータ構造を定義する。
  - golangには配列があるが、ここではスライスを使って配列のデータ構造を実装する。
  - 配列は固定長なので長さ（length)を用意しておく。 
- insert
  - 条件分岐
    - 配列が満杯のとき（=data数がlengthと同じとき）
    - 配列の範囲外のアクセスが発生するとき
  - データをずらす
    - 配列の末尾から任意のindexまでを対象にループでデータをずらす
    - ずらし終わった後に配列にデータを追加し、lengthを更新する
- delete
  - 考え方としてinsertの逆に近い
  - 条件分岐
    - 配列の範囲外のアクセスだけ考慮すれば良い
    - データをへらすので配列が満杯かどうかは条件外
  - データをずらす
    - indexで指定したデータをunsetする
    - 末尾からではなく、任意のindexから末尾に向かってループでデータをずらす
    - ずらし終わったら後はlengthの更新をする
- get
  - 配列の範囲外へのアクセスだけ考慮してランダムアクセス（添字でデータを参照）
- ノート
![Image from iOS](/assets/images/posts/algorithms-data-structures-array/67912042-f3e03200-fbcb-11e9-8a42-34f28fd474f4.jpg)

# 参考
- [github -  TomorrowWu/golang-algorithms](https://github.com/TomorrowWu/golang-algorithms/blob/master/data-structures/array/array.go)
  - sliceでarrayを実装したわかりやすいコードだったので参考させて頂いた。

# 関連
- [bmf-tech.com - O（オーダー）記法とアルゴリズムの計算量の求め方](https://bmf-tech.com/posts/O（オーダー）記法とアルゴリズムの計算量の求め方)



