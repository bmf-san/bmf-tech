---
title: Implementing Fan-In and Fan-Out in Go
description: 'Implement fan-in and fan-out concurrent patterns in Go using channels and goroutines. Fan-in aggregates multiple producer outputs; fan-out distributes one input across multiple workers.'
slug: go-fan-in-fan-out-implementation
date: 2023-08-21T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - fan-in
  - fan-out
translation_key: go-fan-in-fan-out-implementation
---

# Overview
Implementing the fan-in and fan-out patterns in Go for concurrent processing.

# What are Fan-In and Fan-Out?
Fan-in is the process of aggregating multiple inputs into one, while fan-out is the process of distributing a single input into multiple outputs.

Fan-in aggregates data, and fan-out distributes data.

In Go, this can be achieved using channels and goroutines.

# Implementation
The source code is also available on [github](https://github.com/bmf-san/go-snippets/blob/master/channel/fan-in-and-fan-out.go).

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

Data is distributed in the fan-out process and aggregated in the fan-in process.

# Thoughts
I lack confidence in concurrent processing, so I need to study more...

# References
- [ludwig125.hatenablog.com - Go Language Pipeline, Fan-In, Fan-Out](https://ludwig125.hatenablog.com/entry/2019/10/01/052011)
- [devlights.hatenablog.com - Go Memo-73 (Channel for Aggregating Data with Fan-In Pattern)](https://devlights.hatenablog.com/entry/2020/03/23/015027)
- [devlights.hatenablog.com - Go Memo-79 (Function for Fan-Out with Specified Number of Workers)](https://devlights.hatenablog.com/entry/2020/03/27/165236)
- [tech-up.hatenablog.com - Fan-Out, Fan-In Pattern [Go]](https://tech-up.hatenablog.com/entry/2018/12/03/170013)
- [selfnote.work - [Golang] Challenge with Algorithms ~ Let's Implement Fan In/Fan Out with Channels! ~](https://selfnote.work/20211004/programming/golang-binary-tree/)
- [go.dev - Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
- [kapoorrahul.medium.com - Golang Fan-In Fan-Out Concurrency Pattern](https://kapoorrahul.medium.com/golang-fan-in-fan-out-concurrency-pattern-f5a29ff1f93b)
- [mariocarrion.com - LEARNING GO: FAN-IN/FAN-OUT CONCURRENCY PATTERN](https://mariocarrion.com/2021/08/19/learning-golang-concurrency-patterns-fan-in-fan-out.html)
- [www.golinuxcloud.com - Go Fan Out Fan In Concurrency Pattern Explained](https://www.golinuxcloud.com/go-fan-out-fan-in/)
