---
title: "Goでfan-inとfan-outを実装する"
slug: "go-fan-in-fan-out-implementation"
date: 2023-08-21
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "fan-in"
  - "fan-out"
draft: false
---

# 概要
並行処理のパターンであるfan-in、fan-outをGoで実装する。

# fan-in/fan-outとは
fan-inは、複数の入力を1つにまとめる処理で、fan-outは、1つの入力を複数に分ける処理である。

fan-inはデータを集約させ、fan-outはデータを分散させる。

Goではchannelとgoroutineを使って実現することができる。

# 実装
ソースコードは[github](https://github.com/bmf-san/go-snippets/blob/master/channel/fan-in-and-fan-out.go)にも置いてある。

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(id int, out chan<- int) {
	for i := 0; i < 5; i++ {
		value := rand.Intn(100)
		fmt.Printf("Producer %d: Sending %d\n", id, value)
		out <- value
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	}
	close(out)
}

func fanIn(inputs []<-chan int, out chan<- int) {
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(ch <-chan int) {
			for value := range ch {
				out <- value
			}
			wg.Done()
		}(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Fan-Out
	numProducers := 3
	inputs := make([]chan int, numProducers)
	for i := 0; i < numProducers; i++ {
		inputs[i] = make(chan int)
		go producer(i+1, inputs[i])
	}

	// Convert channels to <-chan int
	inputChans := make([]<-chan int, numProducers)
	for i := 0; i < numProducers; i++ {
		inputChans[i] = inputs[i]
	}

	// Fan-In
	result := make(chan int)
	go fanIn(inputChans, result)

	// Consume the merged values
	for value := range result {
		fmt.Printf("Consumer: Received %d\n", value)
	}

	fmt.Println("All done!")
}
```

fan-outの処理でデータを分散して、fan-inの処理でデータを集約している。

# 所感
並行処理は自身がないので勉強しないといけない。。。

# 参考
- [ludwig125.hatenablog.com - go言語のpipeline、fan-in、fan-out](https://ludwig125.hatenablog.com/entry/2019/10/01/052011)
- [devlights.hatenablog.com - Goメモ-73 (fan-in パターンでデータを集約するチャネル, FanIn)](https://devlights.hatenablog.com/entry/2020/03/23/015027)
- [devlights.hatenablog.com - Goメモ-79 (指定されたワーカー数でファンアウトさせる関数, FanOut)
](https://devlights.hatenablog.com/entry/2020/03/27/165236)
- [tech-up.hatenablog.com - fan-out、fan-inパターン【Go】](https://tech-up.hatenablog.com/entry/2018/12/03/170013)
- [selfnote.work - [Golang]アルゴリズムにチャレンジ \~ChannelでFan In/Fan Outを実装しよう!\~](https://selfnote.work/20211004/programming/golang-binary-tree/)
- [go.dev - Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
- [kapoorrahul.medium.com - Golang Fan-In Fan-Out Concurrency Pattern](https://kapoorrahul.medium.com/golang-fan-in-fan-out-concurrency-pattern-f5a29ff1f93b)
- [mariocarrion.com - LEARNING GO: FAN-IN/FAN-OUT CONCURRENCY PATTERN](https://mariocarrion.com/2021/08/19/learning-golang-concurrency-patterns-fan-in-fan-out.html)
- [www.golinuxcloud.com - Go Fan Out Fan In Concurrency Pattern Explained](https://www.golinuxcloud.com/go-fan-out-fan-in/)
