---
title: Implementing an In-Memory Cache in Golang
slug: golang-in-memory-cache-implementation
date: 2020-09-29T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Cache
description: A lightweight and simple in-memory cache implementation in Golang.
translation_key: golang-in-memory-cache-implementation
---

# Overview
There are several good in-memory cache libraries for Golang, but I needed something lightweight and simple, so I decided to implement one myself.

# Implementation
## Requirements
- Should be able to store multiple data entries.
- Should store data in memory with an expiration time. Data should be discarded from memory once expired.
- Should handle concurrent access and updates to the cache, with proper locking mechanisms.

## Initial Design
*Note: The code is available at [github.com - bmf-san/go-snippets/architecture_design/cache/cache.go](https://github.com/bmf-san/go-snippets/blob/master/architecture_design/cache/cache.go).*

This is the initial implementation based on my first idea:

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

I initially thought `sync.Map` was convenient because it eliminates the need to handle locking manually. However, due to its data structure and functionality, it did not meet the requirements, so I decided against using it.

## Release Version
*Note: The code is available at [github.com - bmf-san/go-snippets/architecture_design/cache/cache_with_goroutine.go](https://github.com/bmf-san/go-snippets/blob/master/architecture_design/cache/cache_with_goroutine.go).*

This is the version that meets the requirements:

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

Although `sync.Map` is convenient, it became challenging to check and delete expired cache data without specifying cache keys. Therefore, I decided to use a `map` to store the cache data.

The expiration check is performed at intervals using a `ticker`. In the above implementation, the interval is set to one second. However, this means that cached data can still be accessed for up to one second after its expiration time, effectively making the expiration time equal to the specified time plus the interval.

# Thoughts
This was a good opportunity to get started with concurrency and locking in Golang.

# References
- [github.com - patrickmn/go-cache](https://github.com/patrickmn/go-cache)
- [Stack overflow - Map with TTL option in Go](https://stackoverflow.com/questions/25484122/map-with-ttl-option-in-go)
- [groups.google.com - sync.Map for caching](https://groups.google.com/g/golang-nuts/c/avSIKqUKKAM?pli=1)
- [golang.org - pkg/sync/#Map](https://golang.org/pkg/sync/#Map)