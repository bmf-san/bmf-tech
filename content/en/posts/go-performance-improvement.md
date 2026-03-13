---
title: Improving Code Performance with Go
description: 'Learn practical Go performance optimization: pprof profiling, benchmark tests, escape analysis, memory allocation reduction, and real examples from optimizing a custom HTTP router.'
slug: go-performance-improvement
date: 2023-06-12T00:00:00Z
author: bmf-san
categories:
  - Performance
tags:
  - Golang
translation_key: go-performance-improvement
---

[Makuake Advent Calendar 2022](https://adventar.org/calendars/8496) - Day 9!

# Improving Code Performance with Go
When I decided to enhance the performance of my custom HTTP Router, [goblin](https://github.com/bmf-san/goblin), I delved into improving Go's performance. This post discusses the approaches and practices I implemented.

# Prerequisites
While deeper tuning requires more knowledge, here are the essentials:

- Garbage Collection
  - Automatically frees up memory areas that are no longer needed during program execution.
- Memory Areas
  - Text Area
    - Area where machine code-translated programs are stored.
  - Stack Area
    - Memory area allocated during program execution.
    - Targets data with a predetermined size at runtime.
    - Automatically freed (e.g., when a function execution ends and is no longer needed).
    - ex. Arguments, return values, temporary variables, etc.
  - Heap Area
    - Memory area allocated during program execution.
    - Targets data with dynamically determined size.
    - Subject to garbage collection.
  - Static Area
    - Memory area allocated during program execution.
    - Retained until the program ends.
    - ex. Global variables, static variables, etc.

# Performance Improvement Approach
Before improving performance, it's crucial to determine if it's necessary (is it worth sacrificing readability, is the application a bottleneck, etc.). Assuming it's necessary, here are some methods:

- Algorithm Optimization
- Data Structure Optimization
- Cache Utilization
- Parallel Processing
- Compile Optimization

Before implementing improvements, conduct measurement and analysis. (The need for performance improvement is assumed, but this varies by individual needs, so it won't be discussed here.)

Introducing packages and tools for measurement and analysis in Go.

## Benchmark
Go includes [Benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks) in its standard testing package for obtaining code benchmarks.

For example, you can execute the following code with the command `go test -bench=. -benchmem` to obtain benchmarks.

```go
package main

import (
	"math/rand"
	"testing"
)

func BenchmarkRandIn(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N automatically specifies a reliable number of benchmark iterations
		rand.Int() // Function to be measured
	}
}
```

The output will look like this:
```sh
goos: darwin
goarch: amd64
pkg: bmf-san/go-perfomance-tips
cpu: VirtualApple @ 2.50GHz
BenchmarkRandIn-8       87550500                13.53 ns/op            0 B/op          0 allocs/op
PASS
ok      bmf-san/go-perfomance-tips      1.381s
```

The benchmark results can be interpreted as follows:
- 87550500
  - Number of function executions.
  - More executions indicate better performance.
- 13.53 ns/op
  - Time taken per function execution.
  - Less time indicates better performance.
- 0 B/op
  - Memory size allocated per function execution.
  - Less memory indicates better performance.
- 0 allocs/op
  - Number of memory allocations per function execution.
  - Fewer allocations indicate better performance.

In Go, benchmarks can be easily obtained like this.

For more on Go's benchmark features, refer to the documentation.
[Benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)

Using a tool like [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat#section-readme) to compare benchmark results can show the percentage of improvement.

In my [bmf-san/goblin](https://github.com/bmf-san/goblin), I've integrated it into CI to compare results before and after commits.

```sh
// This is an example where nothing has improved...
go test -bench . -benchmem -count 1 > new.out
benchstat old.out new.out
name           old time/op    new time/op    delta
Static1-36        248ns ± 0%     246ns ± 0%   ~     (p=1.000 n=1+1)
Static5-36        502ns ± 0%     495ns ± 0%   ~     (p=1.000 n=1+1)
Static10-36       874ns ± 0%     881ns ± 0%   ~     (p=1.000 n=1+1)
WildCard1-36     1.60µs ± 0%    1.50µs ± 0%   ~     (p=1.000 n=1+1)
WildCard5-36     4.49µs ± 0%    4.92µs ± 0%   ~     (p=1.000 n=1+1)
WildCard10-36    7.68µs ± 0%    7.58µs ± 0%   ~     (p=1.000 n=1+1)
Regexp1-36       1.38µs ± 0%    1.48µs ± 0%   ~     (p=1.000 n=1+1)
Regexp5-36       4.30µs ± 0%    4.11µs ± 0%   ~     (p=1.000 n=1+1)
Regexp10-36      7.66µs ± 0%    7.18µs ± 0%   ~     (p=1.000 n=1+1)
```

If performance degradation is absolutely unacceptable, you might want to set up CI to fail in such cases.

To check actual memory allocation, you can specify build options and build. Increasing the number of `-m` in `-gcflags` provides more detailed results.

```go
package main

import "fmt"

// Run build with go build -o example -gcflags '-m' gcflagsexample.go
func main() {
	a := "hello"
	b := "world"
	fmt.Println(a + b)
}
```

Running `go build -o example -gcflags '-m' gcflagsexample.go` yields the following output:

```sh
# command-line-arguments
./gcflagsexample.go:9:13: inlining call to fmt.Println
./gcflagsexample.go:9:13: ... argument does not escape
./gcflagsexample.go:9:16: a + b escapes to heap
./gcflagsexample.go:9:16: a + b escapes to heap
```

This simple example clearly shows how heap allocation can be identified and reduced to improve memory allocation, making it a useful analysis method.

## Profiling
To analyze where bottlenecks exist at the function level, Go offers a tool called [pprof](https://pkg.go.dev/net/http/pprof).

```go
package main

import (
	"sort"
	"testing"
)

func sortAlphabetically() {
	s := []string{"abc", "aba", "cba", "acb"}
	sort.Strings(s)
}

func BenchmarkSortAlphabetically(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortAlphabetically()
	}
}
```

To view the CPU profile, execute:

`go test -test.bench=BenchmarkSortAlphabetically -cpuprofile cpu.out && go tool pprof -http=":8888" cpu.out`

![cpu_profile](/assets/images/posts/go-performance-improvement/206718659-bc8b2df8-30d6-4d3c-819f-2846fd3b2c71.png)

To view the memory profile, execute:

`go test -test.bench=BenchmarkSortAlphabetically profilingexample_test.go -memprofile mem.out && go tool pprof -http=":8889" mem.out`

![memory_profile](/assets/images/posts/go-performance-improvement/206716765-b62ab1a9-9bad-4cdb-8dd7-966c714fe940.png)

Using [pprof](https://pkg.go.dev/net/http/pprof)'s UI makes it easier to identify where the bottlenecks are.

# Practical Application
An example of improving my custom HTTP Router, [goblin](https://github.com/bmf-san/goblin).

The PR in question is here.
[Reduce the memory allocation by refactoring explodePath method #68](https://github.com/bmf-san/goblin/pull/68)

[goblin](https://github.com/bmf-san/goblin) is an HTTP Router based on a trie tree, compatible with the net/http interface.

It has the minimum necessary features for routing.
cf. [goblin#features](https://github.com/bmf-san/goblin#features)

## Benchmark
First, execute a benchmark test to measure performance.

```sh
go test -bench=. -cpu=1 -benchmem
```

The benchmark test includes static routing (e.g., /foo/bar), dynamic routing (e.g., /foo/:bar), and regex-based routing (e.g., /foo/:bar[^\d+$]) test cases.

The routing process involves:

1. Inserting data into the tree structure (≒ defining routing)
2. Searching for data in the tree structure (returning data based on the requested path)

This test case measures only the latter.

The output is as follows:

```sh
goos: darwin
goarch: amd64
pkg: github.com/bmf-san/goblin
cpu: VirtualApple @ 2.50GHz
BenchmarkStatic1         5072353               240.1 ns/op           128 B/op          4 allocs/op
BenchmarkStatic5         2491546               490.0 ns/op           384 B/op          6 allocs/op
BenchmarkStatic10        1653658               729.6 ns/op           720 B/op          7 allocs/op
BenchmarkWildCard1       1602606               747.3 ns/op           456 B/op          9 allocs/op
BenchmarkWildCard5        435784              2716 ns/op            1016 B/op         23 allocs/op
BenchmarkWildCard10       246729              5033 ns/op            1680 B/op         35 allocs/op
BenchmarkRegexp1         1647258               733.2 ns/op           456 B/op          9 allocs/op
BenchmarkRegexp5          456652              2641 ns/op            1016 B/op         23 allocs/op
BenchmarkRegexp10         251998              4780 ns/op            1680 B/op         35 allocs/op
PASS
ok      github.com/bmf-san/goblin       14.304s
```

Trends can be observed in execution count, time per execution, memory size per execution, and memory allocation count.

Even static routing incurs memory allocation, which is concerning (other HTTP Router benchmarks show 0 allocs).

## Profiling
Next, obtain a profile using pprof.

Focus on memory profiling this time.

```sh
go test -bench . -memprofile mem.out && go tool pprof -http=":8889" mem.out
```

Graph output:
![pprof_graph](/assets/images/posts/go-performance-improvement/206716778-8c5b2ad6-2e6a-444f-8f4c-7267a253446f.png)

The largest box (most memory usage) is the `explodePath` process.

The Top (list of longest execution times) also shows `explodePath` at the top.

![pprof_top](/assets/images/posts/go-performance-improvement/206716793-08c464a8-db4c-4838-b872-dc6b2c51b154.png)

Flat is the function's processing time, Cum is the processing time including wait time.

Further, check the Source to see which part of the function is heavy.

![pprof_source](/assets/images/posts/go-performance-improvement/206716787-c1be565d-9364-40d1-b555-70836d056832.png)

`Search` is the core process responsible for router matching, and `explodePath` is identified as the bottleneck within it.

## Tuning
`explodePath` splits a received string by `/` and returns it as a []string.

```go
// explodePath removes an empty value in slice.
func explodePath(path string) []string {
	s := strings.Split(path, pathDelimiter)
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
```

Test code for clarity:
```go
func TestExplodePath(t *testing.T) {
	cases := []struct {
		actual   []string
		expected []string
	}{
		{
			actual:   explodePath(""),
			expected: nil,
		},
		{
			actual:   explodePath("/"),
			expected: nil,
		},
		{
			actual:   explodePath("//"),
			expected: nil,
		},
		{
			actual:   explodePath("///"),
			expected: nil,
		},
		{
			actual:   explodePath("/foo"),
			expected: []string{"foo"},
		},
		{
			actual:   explodePath("/foo/bar"),
			expected: []string{"foo", "bar"},
		},
		{
			actual:   explodePath("/foo/bar/baz"),
			expected: []string{"foo", "bar", "baz"},
		},
		{
			actual:   explodePath("/foo/bar/baz/"),
			expected: []string{"foo", "bar", "baz"},
		},
	}

	for _, c := range cases {
		if !reflect.DeepEqual(c.actual, c.expected) {
			t.Errorf("actual:%v expected:%v", c.actual, c.expected)
		}
	}
}
```

The variable `r` defined as []string has no defined capacity, suggesting poor memory efficiency.

Below is a benchmark test for appending to a slice, prepared for verification, and its results:
```go
package main

import "testing"

func BenchmarkSliceLen0Cap0(b *testing.B) {
	var s []int

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
	b.StopTimer()
}

func BenchmarkSliceLenN(b *testing.B) {
	var s = make([]int, b.N)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
	b.StopTimer()
}

func BenchmarkSliceLen0CapN(b *testing.B) {
	var s = make([]int, 0, b.N)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s = append(s, i)
	}
	b.StopTimer()
}
```

```sh
goos: darwin
goarch: amd64
pkg: example.com
cpu: VirtualApple @ 2.50GHz
BenchmarkSliceLen0Cap0  100000000               11.88 ns/op           45 B/op          0 allocs/op
BenchmarkSliceLenN      78467056                23.60 ns/op           65 B/op          0 allocs/op
BenchmarkSliceLen0CapN  616491007                5.057 ns/op           8 B/op          0 allocs/op
PASS
ok      example.com     6.898s
bmf@bmfnoMacBook-Air:~/Desktop$
```

This result suggests that specifying capacity can lead to more efficient code.

Thus, `explodePath` is modified as follows:

```go
func explodePath(path string) []string {
	s := strings.Split(path, "/")
	// var r []string
	r := make([]string, 0, strings.Count(path, "/")) // Specify capacity
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
```

Further implementation improvements:
```go
func explodePath(path string) []string {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	return strings.FieldsFunc(path, splitFn)
}
```

Compare benchmarks for the original `explodePath`, the implementation with specified slice capacity, and the implementation using `strings.FieldFunc`.

```go
package main

import (
	"strings"
	"testing"
)

func explodePath(path string) []string {
	s := strings.Split(path, "/")
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func explodePathCap(path string) []string {
	s := strings.Split(path, "/")
	r := make([]string, 0, strings.Count(path, "/"))
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func explodePathFieldsFunc(path string) []string {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	return strings.FieldsFunc(path, splitFn)
}

func BenchmarkExplodePath(b *testing.B) {
	paths := []string{"", "/", "///", "/foo", "/foo/bar", "/foo/bar/baz"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range paths {
			explodePath(v)
		}
	}
	b.StopTimer()
}

func BenchmarkExplodePathCap(b *testing.B) {
	paths := []string{"", "/", "///", "/foo", "/foo/bar", "/foo/bar/baz"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range paths {
			explodePathCap(v)
		}
	}
	b.StopTimer()
}

func BenchmarkExplodePathFieldsFunc(b *testing.B) {
	paths := []string{"", "/", "///", "/foo", "/foo/bar", "/foo/bar/baz"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range paths {
			explodePathFieldsFunc(v)
		}
	}
	b.StopTimer()
}
```

```sh
goos: darwin
goarch: amd64
pkg: example.com
cpu: VirtualApple @ 2.50GHz
BenchmarkExplodePath             1690340               722.2 ns/op           432 B/op         12 allocs/op
BenchmarkExplodePathCap          1622161               729.5 ns/op           416 B/op         11 allocs/op
BenchmarkExplodePathFieldsFunc   4948364               239.5 ns/op            96 B/op          3 allocs/op
PASS
ok      example.com     5.685s
```

The implementation using `strings.PathFieldFunc` seems to perform best, so it is adopted.

## Effect Measurement
Check the results after improving the `explodePath` implementation.

### Benchmark
```sh
# Before Improvement
goos: darwin
goarch: amd64
pkg: github.com/bmf-san/goblin
cpu: VirtualApple @ 2.50GHz
BenchmarkStatic1         5072353               240.1 ns/op           128 B/op          4 allocs/op
BenchmarkStatic5         2491546               490.0 ns/op           384 B/op          6 allocs/op
BenchmarkStatic10        1653658               729.6 ns/op           720 B/op          7 allocs/op
BenchmarkWildCard1       1602606               747.3 ns/op           456 B/op          9 allocs/op
BenchmarkWildCard5        435784              2716 ns/op            1016 B/op         23 allocs/op
BenchmarkWildCard10       246729              5033 ns/op            1680 B/op         35 allocs/op
BenchmarkRegexp1         1647258               733.2 ns/op           456 B/op          9 allocs/op
BenchmarkRegexp5          456652              2641 ns/op            1016 B/op         23 allocs/op
BenchmarkRegexp10         251998              4780 ns/op            1680 B/op         35 allocs/op
PASS
ok      github.com/bmf-san/goblin       14.304s

# After Improvement
go test -bench=. -cpu=1 -benchmem -count=1
goos: darwin
goarch: amd64
pkg: github.com/bmf-san/goblin
cpu: VirtualApple @ 2.50GHz
BenchmarkStatic1        10310658               117.7 ns/op            32 B/op          1 allocs/op
BenchmarkStatic5         4774347               258.1 ns/op            96 B/op          1 allocs/op
BenchmarkStatic10        2816960               435.8 ns/op           176 B/op          1 allocs/op
BenchmarkWildCard1       1867770               653.4 ns/op           384 B/op          6 allocs/op
BenchmarkWildCard5        496778              2484 ns/op             928 B/op         13 allocs/op
BenchmarkWildCard10       258508              4538 ns/op            1560 B/op         19 allocs/op
BenchmarkRegexp1         1978704               608.4 ns/op           384 B/op          6 allocs/op
BenchmarkRegexp5          519240              2394 ns/op             928 B/op         13 allocs/op
BenchmarkRegexp10         280741              4309 ns/op            1560 B/op         19 allocs/op
PASS
ok      github.com/bmf-san/goblin       13.666s
```

Overall improvement trends can be observed when comparing before and after.

### Profiling
pprof Graph.

![pprof_graph_after](/assets/images/posts/go-performance-improvement/206716776-7cee0600-6cb4-4d82-b534-e9b2e2ff72ac.png)

pprof Top.

![pprof_top_after](/assets/images/posts/go-performance-improvement/206716789-1ef3cbae-c638-4935-a6fa-22907fe30633.png)

The bottleneck has shifted to `strings.FieldsFunc` called within `explodePath`.

## Further Improvements
Other improvements have been made to [goblin](https://github.com/bmf-san/goblin), resulting in this release tag.
[6.0.0](https://github.com/bmf-san/goblin/releases/tag/6.0.0)

These are minor improvements without major changes to data structures or algorithms, so unfortunately, no remarkable improvements are observed.

It seems challenging with the current data structures and algorithms. (Other routers use more advanced trees, so it's understandable.)

Although slightly off-topic, I created a benchmarker to compare with other routers for improvement hints.

[bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark)

It's interesting to compare, and it's clear that I'm losing badly. I cried.

I want to study other router implementations, advanced tree structures I previously struggled with, and improve further.

# Summary
- Benchmarking and profiling in Go are easy.
- Don't guess, measure.
- Minor improvements rarely yield significant results (that's true).

# References
- [github.com - google/pprof](https://github.com/google/pprof/blob/main/doc/README.md)
- [github.com - dgryski/go-perfbook](https://github.com/dgryski/go-perfbook)
- [dave.cheney.net - High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html)
- [go.dev - Profiling Go Programs](https://go.dev/blog/pprof)
- [go.dev - A Guide to the Go Garbage Collector](https://go.dev/doc/gc-guide)
- [developer.so-tech.co.jp](https://developer.so-tech.co.jp/entry/2022/07/28/112845)
- [segment.com - Allocation efficiency in high-performance Go services](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/)
- [blog.logrocket.com - Benchmarking in Golang: Improving function performance](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)
- [medium.com - Go code refactoring : the 23x performance hunt](https://medium.com/@val_deleplace/go-code-refactoring-the-23x-performance-hunt-156746b522f7)
- [medium.com - Go言語のプロファイリングツール、pprofのWeb UIがめちゃくちゃ便利なので紹介する](https://medium.com/eureka-engineering/go%E8%A8%80%E8%AA%9E%E3%81%AE%E3%83%97%E3%83%AD%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AA%E3%83%B3%E3%82%B0%E3%83%84%E3%83%BC%E3%83%AB-pprof%E3%81%AEweb-ui%E3%81%8C%E3%82%81%E3%81%A1%E3%82%83%E3%81%8F%E3%81%A1%E3%82%83%E4%BE%BF%E5%88%A9%E3%81%AA%E3%81%AE%E3%81%A7%E7%B4%B9%E4%BB%8B%E3%81%99%E3%82%8B-6a34a489c9ee)
- [teivah.medium.com - Good Code vs Bad Code in Golang
](https://teivah.medium.com/good-code-vs-bad-code-in-golang-84cb3c5da49d)
- [hnakamur.github.io - goで書いたコードがヒープ割り当てになるかを確認する方法](https://hnakamur.github.io/blog/2018/01/30/go-heap-allocations/)
- [glog.kazu69.net - Goのメモリ管理を眺めてみた](https://blog.kazu69.net/2017/08/20/memory-management-go/)
- [dsas.blog.klab.org - Goでアロケーションに気をつけたコードを書く方法](http://dsas.blog.klab.org/archives/52191778.html)
- [tech.speee.jp - Goのロギングライブラリから見たゼロアロケーション](https://tech.speee.jp/entry/2022/07/12/134605)
- [kawasin73.hatenablog.com - メモリアロケーションに対する罪悪感
](https://kawasin73.hatenablog.com/entry/2019/11/10/112301)