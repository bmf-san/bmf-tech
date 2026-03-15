---
title: Golangでトライ木を実装する
description: "トライ木（プレフィックス木）のデータ構造をGoで実装し、O(m)の検索・挿入計算量とRadix Treeとの使い分けを解説。"
slug: golang-tries-implementation
date: 2019-09-24T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - アルゴリズムとデータ構造
tags:
  - Golang
  - 基数木
  - トライ木
translation_key: golang-tries-implementation
---


# 概要
トライ木のアルゴリズムと実装についてかく。

[bmf-san/road-to-algorithm-master](https://github.com/bmf-san/road-to-algorithm-master)　

# トライ木とは
トライ木（プレフィックス木ともいう。英語はそれぞれ、trie、prefix tree）は文字列の集合を扱う木構造の一種。

各ノードは単一または複数の文字列あるいは数値を持ち（ノードは必ずしも値を持つ必要はない）、根ノードから葉に向かって探索して値をつなげていくことで単語を表現する。

ネットワークのレイヤーならIPアドレスの探索、アプリケーションレイヤーならhttpのルーティングに、機械学習とかの文脈であれば形態素解析といったところでトライ木の応用が見受けられる。

言葉で説明するよりもビジュアルのほうが頭に入りやすい。

[Algorithm Visualizations - Trie (Prefix Tree)](https://www.cs.usfca.edu/~galles/visualization/Trie.html)

メモリ効率が悪いので、メモリ効率がボトルネックとなるような場合はRadix Tree（Patricia Trie）のような文字列のプレフィックスを効率的に格納する木構造を検討したほうが良さそう。

# 計算量
検索キーの長さをmとしたとき、最悪計算量はO(m)となる。検索も挿入も同じである。

# 実装

[github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/tree/trie](https://github.com/bmf-san/road-to-algorithm-master/tree/master/data_structures/tree/trie)においてある。

検索キーの挿入と検索だけ実装した。

```go
package main

import "fmt"

// Node is a node of tree.
type Node struct {
	key      string
	children map[rune]*Node
}

// NewTrie is create a root node.
func NewTrie() *Node {
	return &Node{
		key:      "",
		children: make(map[rune]*Node),
	}
}

// Insert is insert a word to tree.
func (n *Node) Insert(word string) {
	runes := []rune(word)
	curNode := n

	for _, r := range runes {
		if nextNode, ok := curNode.children[r]; ok {
			curNode = nextNode
		} else {
			curNode.children[r] = &Node{
				key:      string(r),
				children: make(map[rune]*Node),
			}
		}
	}
}

// Search is search a word from a tree.
func (n *Node) Search(word string) bool {
	if len(n.key) == 0 && len(n.children) == 0 {
		return false
	}

	runes := []rune(word)
	curNode := n

	for _, r := range runes {
		if nextNode, ok := curNode.children[r]; ok {
			curNode = nextNode
		} else {
			return false
		}
	}

	return true
}

func main() {
	t := NewTrie()

	t.Insert("word")
	t.Insert("wheel")
	t.Insert("world")
	t.Insert("hospital")
	t.Insert("mode")

	fmt.Printf("%v", t.Search("mo")) // true
}
```

構造体Nodeのchildrenはruneをmapのキーとしているが、stringでも問題ない気がする。

Insertのアルゴリズムは単純で、Insertしたい文字列の文字数分ループして、子ノードに一致するキーがあるかチェック、なければノードを追加するだけである。

子ノードに一致するキーがあるかどうかはGolangでは、`v, ok:=map[key]`のイディオムを使って書くことができる。（初心者過ぎて知らなかったせいでちょっとハマった）

SearchのアルゴリズムはInsertのアルゴリズムが書ければ理解できる。というより殆ど考え方は同じ。

# 所感
トライ木の応用であるRadix Treeという木構造の実装に取り組んでいたのだが、あと一歩みたいなところで何度も躓いて挫折した。。。
ちょっと遠回りしてまた再挑戦する。

GolangでTrieを使ってroutingを作ろうとしている。
[https://github.com/bmf-san/bmf-go-router](https://github.com/bmf-san/bmf-go-router)

トライ木はちょっとしたサジェスト機能なんかの実装とかにも使える気がするのでJavaScriptとかで実装してみたい。

## 関連記事

- [O（オーダー）記法とアルゴリズムの計算量の求め方](/ja/posts/big-o-notation-algorithms/)
- [アルゴリズムとデータ構造 - 二分探索木](/ja/posts/algorithms-data-structures-binary-search-tree/)
- [アルゴリズムとデータ構造の基本の復習](/ja/posts/algorithms-data-structures-review/)
- [自作HTTPルーターから新しいServeMuxへ](/ja/posts/custom-http-router-to-new-servemux/)

