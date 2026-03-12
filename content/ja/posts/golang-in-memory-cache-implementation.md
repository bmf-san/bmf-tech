---
title: Golangでインメモリなキャッシュを実装する
description: Golangでインメモリなキャッシュを実装するについて調査・整理したメモ。基本概念と重要ポイントを解説します。
slug: golang-in-memory-cache-implementation
date: 2020-09-29T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - キャッシュ
translation_key: golang-in-memory-cache-implementation
---


# 概要
Golangのインメモリキャッシュのライブラリは良さそうなものが存在するが、軽量でシンプルなもので十分だったので自前で実装してみた。

# 実装
## 要件
- 複数のデータを保持することができる。
- 期限付きのデータをメモリに保持することができる。期限が来たらメモリから破棄されること。
- キャッシュへの同時参照、更新を考慮し、データのロックが意識されていること。

## 初期設計
※[github.com - bmf-san/go-snippets/architecture_design/cache/cache.go](https://github.com/bmf-san/go-snippets/blob/master/architecture_design/cache/cache.go)に置いてあるが、転載。

最初に思いついた感じで実装したもの。

```golang
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Cache is a struct for caching.
type Cache struct {
	value   sync.Map
	expires int64
}

// Expired determines if it has expired.
func (c *Cache) Expired(time int64) bool {
	if c.expires == 0 {
		return false
	}
	return time > c.expires
}

// Get gets a value from a cache. Returns an empty string if the value does not exist or has expired.
func (c *Cache) Get(key string) string {
	if c.Expired(time.Now().UnixNano()) {
		log.Printf("%s has expired", key)
		return ""
	}
	v, ok := c.value.Load(key)
	var s string
	if ok {
		s, ok = v.(string)
		if !ok {
			log.Printf("%s does not exists", key)
			return ""
		}
	}
	return s
}

// Put puts a value to a cache. If a key and value exists, overwrite it.
func (c *Cache) Put(key string, value string, expired int64) {
	c.value.Store(key, value)
	c.expires = expired
}

var cache = &Cache{}

func main() {
	fk := "first-key"
	sk := "second-key"

	cache.Put(fk, "first-value", time.Now().Add(2*time.Second).UnixNano())
	s := cache.Get(fk)
	fmt.Println(cache.Get(fk))

	time.Sleep(5 * time.Second)

	// fk should have expired
	s = cache.Get(fk)
	if len(s) == 0 {
		cache.Put(sk, "second-value", time.Now().Add(100*time.Second).UnixNano())
	}
	fmt.Println(cache.Get(sk))
}
```

ロックの処理を気にしなくてよい`sync.Map`が便利で良いなぁと思っていたのだが、データ構造や機能的に要件を満たせていないので却下。

## リリース版
※[github.com - bmf-san/go-snippets/architecture_design/cache/cache.go](https://github.com/bmf-san/go-snippets/blob/master/architecture_design/cache/cache_with_goroutine.go)に置いてあるが、転載。

要件を満たす実装をしたバージョン。

```golang
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// item is the data to be cached.
type item struct {
	value   string
	expires int64
}

// Cache is a struct for caching.
type Cache struct {
	items map[string]*item
	mu    sync.Mutex
}

func New() *Cache {
	c := &Cache{items: make(map[string]*item)}
	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				c.mu.Lock()
				for k, v := range c.items {
					if v.Expired(time.Now().UnixNano()) {
						log.Printf("%v has expires at %d", c.items, time.Now().UnixNano())
						delete(c.items, k)
					}
				}
				c.mu.Unlock()
			}
		}
	}()
	return c
}

// Expired determines if it has expires.
func (i *item) Expired(time int64) bool {
	if i.expires == 0 {
		return true
	}
	return time > i.expires
}

// Get gets a value from a cache.
func (c *Cache) Get(key string) string {
	c.mu.Lock()
	var s string
	if v, ok := c.items[key]; ok {
		s = v.value
	}
	c.mu.Unlock()
	return s
}

// Put puts a value to a cache. If a key and value exists, overwrite it.
func (c *Cache) Put(key string, value string, expires int64) {
	c.mu.Lock()
	if _, ok := c.items[key]; !ok {
		c.items[key] = &item{
			value:   value,
			expires: expires,
		}
	}
	c.mu.Unlock()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fk := "first-key"
		sk := "second-key"

		cache := New()

		cache.Put(fk, "first-value", time.Now().Add(2*time.Second).UnixNano())
		fmt.Println(cache.Get(fk))

		time.Sleep(10 * time.Second)

		if len(cache.Get(fk)) == 0 {
			cache.Put(sk, "second-value", time.Now().Add(100*time.Second).UnixNano())
		}
		fmt.Println(cache.Get(sk))
	})
	http.ListenAndServe(":8080", nil)
}
```

`sync.Map`が便利なので使いたかったのだが、`sync.Map`にキャッシュデータを保持させるとキャッシュキー指定なしにキャッシュデータの期限切れチェック・削除が難しいため、キャッシュデータの保持には`map`を採用することにした。

期限切れチェックは`ticker`を使ってインターバルを置いてチェックするようにしている。
上記では1秒毎のインターバルとなっている。
上記の実装ではキャッシュ期限＋1秒経過するまではキャッシュにアクセスすることができてしまうので、実際のキャッシュ期限はexpiresで指定した時間＋インターバルになっている。

# 所感
Golangでの並行処理やロックに入門する良い機会だった。

# 参考
- [github.com - patrickmn/go-cache](https://github.com/patrickmn/go-cache)
- [Stack overflow - Map with TTL option in Go](https://stackoverflow.com/questions/25484122/map-with-ttl-option-in-go)
- [groups.google.com - sync.Map for caching](https://groups.google.com/g/golang-nuts/c/avSIKqUKKAM?pli=1)
- [golang.org - pkg/sync/#Map](https://golang.org/pkg/sync/#Map)

