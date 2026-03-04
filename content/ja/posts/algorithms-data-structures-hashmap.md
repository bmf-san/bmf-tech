---
title: "アルゴリズムとデータ構造 - ハッシュマップ"
slug: "algorithms-data-structures-hashmap"
date: 2023-07-31
author: bmf-san
categories:
  - "アルゴリズムとデータ構造"
tags:
  - "ハッシュマップ"
draft: false
---

# 概要
[アルゴリズム図鑑](https://www.shoeisha.co.jp/book/detail/9784798149776)を参考に、アルゴリズムとデータ構造を学ぶ。

実装は[github - bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)にも置いてある。

# ハッシュマップ
- ハッシュ値を添え字とした配列
- ハッシュの衝突処理
  - 開番地法
    - 衝突が生じた際に、ハッシュ関数とは別の関数を使って別の番地を求める方法。
  - 連鎖法
    - 衝突が生じても新しい番地を求めずに、衝突した番地に衝突したキー同士をポインタでつないだリンクリストを格納することで対応する方式。

# 計算時間
## データへのアクセス
- O(1)
 - 添字を使ってランダムアクセスが可能。

## データの追加
- O(1)
 - 配列の場合は、線形探索で追加箇所を探索する必要があるため、O(n)だが、ハッシュテーブルは、データの追加’箇所をハッシュによって求めるのでO(1)で済む。
 - ハッシュが衝突した場合はこの限りではない。

# 実装
以下はハッシュ値の衝突を考慮していない粗雑なハッシュマップ。

```golang
package main

import "fmt"

// A HashMap is hash map.
type HashMap struct {
	data map[int]string
}

// hash is create a hash key.
func hash(key int) int {
	return key % 5
}

// put is add key to hash map.
func (h HashMap) put(key int, value string) {
	hash := hash(key)
	if h.data == nil {
		h.data = make(map[int]string)
	}
	h.data[hash] = value
}

// get is get a value from hash map.
func (h HashMap) get(key int) string {
	var hash int = hash(key)

	return h.data[hash]
}

func main() {
	h := &HashMap{
		data: make(map[int]string),
	}

	h.put(1, "foo")
	h.put(2, "bar")

	fmt.Printf("%#v\n", h.get(1))
	fmt.Printf("%#v\n", h.get(2))
}
```

- ハッシュ関数のアルゴリズムは色々ある。
  - cf. [wikipedia - ハッシュ関数](https://ja.wikipedia.org/wiki/%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E9%96%A2%E6%95%B0#:~:targetText=%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E9%96%A2%E6%95%B0%20(%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E3%81%8B%E3%82%93%E3%81%99%E3%81%86,%E5%80%A4%E3%81%BE%E3%81%9F%E3%81%AF%E5%8D%98%E3%81%AB%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E3%81%A8%E3%81%84%E3%81%86%E3%80%82)

# 参考
- [Wikipedia - ハッシュテーブル](https://ja.wikipedia.org/wiki/%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB)
- [github.com - flaviocopes/datastructures/blob/master/hashtable/hashtable.go](https://github.com/flaviocopes/datastructures/blob/master/hashtable/hashtable.go)
