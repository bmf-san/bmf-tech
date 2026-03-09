---
title: Goで始めるコードのパフォーマンス改善
slug: go-performance-improvement
date: 2023-06-12T00:00:00Z
author: bmf-san
categories:
  - パフォーマンス
tags:
  - Golang
translation_key: go-performance-improvement
---


[Makuake Advent Calendar 2022](https://adventar.org/calendars/8496)の9日目の記事です！

# Goで始めるコードのパフォーマンス改善
自作HTTP Routerの[goblin](https://github.com/bmf-san/goblin)のパフォーマンス改善をしよう思った際に、Goのパフォーマンス改善について取り組んでみたので、その際のアプローチと実践した取り組みについて書く。

# 前提知識
より奥深いチューニングをする上ではもっと必要な知識があると思うが、最低限必要なことだけリストアップ。

- ガベージコレクション
  - プログラムが実行時に確保したメモリ領域のうち、不要になった領域を自動で解放する機能のこと
- メモリ領域
  - テキスト領域
    - 機械語に変換されたプログラムが可能される領域
  - スタック領域
    - プログラム実行時に確保されるメモリ領域
    - 実行時にサイズが決まっているデータが対象
	- 自動的に解放される（関数の実行が終了して不要になったときなど）
    - ex. 引数、戻り値、一時変数など
  - ヒープ領域
    - プログラム実行時に確保されるメモリ領域
    - 動的にサイズを決まるデータが対象
    - ガベージコレクションの対象
  - 静的領域
    - プログラム実行時に確保されるメモリ領域
    - プログラムが終了されるまで確保される
    - ex. グローバル変数や静的（static）変数など

# パフォーマンス改善のアプローチ
前提として、パフォーマンスを改善する必要性がある（可読性を犠牲にする価値があるか、そもそもアプリケーションがボトルネックだと断定できているのか、など改善すべき理由があるか）かどうかという部分があるが、必要性があるという前提のもとで話を進める。

コードのパフォーマンスを改善する方法として、

- アルゴリズムの最適化　
- データ構造の最適化
- キャッシュの利用
- 並列処理の適用
- コンパイルの最適化

などいくつか思い浮かぶことがあるが、改善策を講じる前に計測や分析を行う。
（計測よりもそもそもパフォーマンス改善が必要性というのが前提にあるが各々のニーズによるのでここでは触れない。）

Goで計測や分析を行うパッケージやツールの紹介をする。

## Benchmark
Goではコードのベンチマークを取得するための[Benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)が標準パッケージであるtestingに含まれている。

例えば次のようなコードを`go test -bench=. -benchmem`というコマンドで実行するとベンチマークを取得することできる。

```go
package main

import (
	"math/rand"
	"testing"
)

func BenchmarkRandIn(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.Nはベンチマークが信頼できる回数を自動的に指定する
		rand.Int() // 計測したい関数
	}
}
```

出力結果は次のような形になる。
```sh
goos: darwin
goarch: amd64
pkg: bmf-san/go-perfomance-tips
cpu: VirtualApple @ 2.50GHz
BenchmarkRandIn-8       87550500                13.53 ns/op            0 B/op          0 allocs/op
PASS
ok      bmf-san/go-perfomance-tips      1.381s
```

ここから読み取れるベンチマークの結果は次の通り。
- 87550500
  - 関数の実行回数
  - 実行回数が多いほどパフォーマンスが良いと考えられる
- 13.53 ns/op
  - 関数の1回あたりの実行に要した時間
  - 時間が少ないほどパフォーマンスが良いと考えられる
- 0 B/op
  - 関数の実行ごとに割当されたメモリのサイズ
  - 少なければ少ないほどパフォーマンスが良いと考えられる
- 0 allocs/op
  - 関数の1回あたりの実行で行われたメモリアロケーションの回数
  - 少なければ少ないほどパフォーマンスが良いと考えられる

Goではこのように簡単にベンチマークを取得することができる。

その他のGoのベンチマークの機能についてはドキュメント参照。
[Benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)

ベンチマークの結果を比較するツールとして[benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat#section-readme)というパッケージが使うと、ベンチマークの結果がどれくらい改善されたか割合を表示してくれるので良い。

自分が管理している[bmf-san/goblin](https://github.com/bmf-san/goblin)ではCIに組み込んでコミット前後の結果を比較できるようにしている。

```sh
// これは何も改善されていない例だが・・
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

パフォーマンス劣化を絶対に許さない！みたいな場合はCIをFailさせるような仕組みにすると良いかもしれない。

このようなベンチマークの結果を見て、実際のメモリ割り当ての様子を確認したい場合には、buildオプションを指定してビルドすることで確認することできる。
-gcflagsに指定する-mの数を増やすとより詳細な結果が得られる。

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

`go build -o example -gcflags '-m' gcflagsexample.go`と実行すると次のような出力が得られる。

```sh
# command-line-arguments
./gcflagsexample.go:9:13: inlining call to fmt.Println
./gcflagsexample.go:9:13: ... argument does not escape
./gcflagsexample.go:9:16: a + b escapes to heap
./gcflagsexample.go:9:16: a + b escapes to heap
```

これは単純な例なので一目瞭然だが、このようにしてヒープへの割当を特定し、ヒープ割当を減らすことによりメモリアロケーションを改善することができるため、分析の方法としても有用である。

## Profiling
関数レベルでどこにボトルネックがあるかというのを分析するためのツールとしてGoには[pprof](https://pkg.go.dev/net/http/pprof)というツールがある。

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

CPUのプロファイルがみたいとき以下を実行。

`go test -test.bench=BenchmarkSortAlphabetically -cpuprofile cpu.out && go tool pprof -http=":8888" cpu.out`

![cpu_profile](https://user-images.githubusercontent.com/13291041/206718659-bc8b2df8-30d6-4d3c-819f-2846fd3b2c71.png)

メモリのプロファイルがみたいときは以下を実行。

`go test -test.bench=BenchmarkSortAlphabetically profilingexample_test.go -memprofile mem.out && go tool pprof -http=":8889" mem.out`

![memory_profile](https://user-images.githubusercontent.com/13291041/206716765-b62ab1a9-9bad-4cdb-8dd7-966c714fe940.png)

[pprof](https://pkg.go.dev/net/http/pprof)のUIを活用することでどこの処理にボトルネックがあるか特定しやすくなる。

# 実践
自作HTTP Routerの[goblin](https://github.com/bmf-san/goblin)の改善例を上げる。

題材としているPRはこちら。
[Reduce the memory allocation by refactoring explodePath method #68](https://github.com/bmf-san/goblin/pull/68)

[goblin](https://github.com/bmf-san/goblin)はトライ木をベースとしたnet/httpのインターフェースと相性の良いHTTP Routerである。

機能としては、ルーティングに必要と思われる最低限のものは持っている。
cf. [goblin#features](https://github.com/bmf-san/goblin#features)

## ベンチマーク
まずはパフォーマンスを計測するためにベンチマークテストを実行する。

```sh
go test -bench=. -cpu=1 -benchmem
```

ベンチマークテストは、静的なルーティング（ex. /foo/bar）、動的なルーティング（ex. /foo/:bar）、正規表現を使ったルーティング（ex. /foo/:bar[^\d+$]）のテストケースをそれぞれ3パターンほど用意している。

ルーティングの処理として、

1. 木構造にデータを入れる（≒ルーティングを定義する）
2. 木構造からデータを探索する（リクエストされたパスを元にデータを返す）

といった流れになるが、このテストケースでは後者のみを計測するようになっている。

出力結果は以下の通り。

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

実行回数、1回あたりの実行回数、実行ごとのメモリサイズ、メモリアローケーション回数のそれぞれにいくつか傾向が読み取れる。

静的なルーティングであってもメモリアローケーションが発生しているのが個人的には気になるところである。（他のHTTP Routerのベンチマークを見ると0 allocsだったりする。）

## プロファイリング
次にpprofを使ってプロファイルを取得する。

今回はメモリだけにフォーカスしてプロファイルを取得。

```sh
go test -bench . -memprofile mem.out && go tool pprof -http=":8889" mem.out
```

Graphの出力結果。
![pprof_graph](https://user-images.githubusercontent.com/13291041/206716778-8c5b2ad6-2e6a-444f-8f4c-7267a253446f.png)

ボックスが一番大きい（メモリを一番使っている）処理が`explodePath`だと分かる。

Top（実行時間の長い順のリスト）を見ても`explodePath`が最上位にいる。

![pprof_top](https://user-images.githubusercontent.com/13291041/206716793-08c464a8-db4c-4838-b872-dc6b2c51b154.png)

Flatは関数の処理時間、Cumは待ち時間も含めた処理時間となる。

さらにSourceを実際に関数内のどのあたりの処理が重いかの確認。

![pprof_source](https://user-images.githubusercontent.com/13291041/206716787-c1be565d-9364-40d1-b555-70836d056832.png)

`Search`はルーターのマッチング処理を担う根幹の処理なので、そこが一番ネックだろうとは思っていたが、その中の特定の処理である`explodePath`がネックになっているということが分かった。

## チューニング
`explodePath`は受け取った文字列を`/`で分割して[]string型にして返すという処理になっている。

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

仕様が分かりやすいようにテストコードも記載。
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

[]string型で定義されている変数`r`は容量が定義されていないため、メモリ効率が悪そうなことが推測される。

以下は検証用に用意したsliceにappendを追加するベンチマークテストとその結果。
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

この結果から、容量を指定してあげることでいくらか効率の良いコードになりそうなことが伺える。

そこで`explodePath`を次のように修正。

```go
func explodePath(path string) []string {
	s := strings.Split(path, "/")
	// var r []string
	r := make([]string, 0, strings.Count(path, "/")) // 容量を指定
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
```

もう少し踏み込んで実装を改善。
```go
func explodePath(path string) []string {
	splitFn := func(c rune) bool {
		return c == '/'
	}
	return strings.FieldsFunc(path, splitFn)
}
```

元の`explodePath`の実装、sliceの容量を確保した実装、`strings.FieldFunc`を利用した実装の3パターンでベンチマークを比較してみる。

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

`strings.PathFieldFunc`を使った実装が一番パフォーマンスが良さそうなので採用。

## 効果測定
`explodePath`の実装を改善した後の結果を確認してみる。

### ベンチーマーク
```sh
# 改善前
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

# 改善後
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

改善前後を比較するに全体的に改善された傾向にあると言えそう。

### プロファイリング
pprofのGraph。

![pprof_graph_after](https://user-images.githubusercontent.com/13291041/206716776-7cee0600-6cb4-4d82-b534-e9b2e2ff72ac.png)

pprofのTop。

![pprof_top_after](https://user-images.githubusercontent.com/13291041/206716789-1ef3cbae-c638-4935-a6fa-22907fe30633.png)

ボトルネックが`explodePath`内で呼び出している`strings.FieldsFunc`に移動したのが分かる。

## さらなる改善
[goblin](https://github.com/bmf-san/goblin)に他にも改善を加えていってリリースされたタグがこちら。
[6.0.0](https://github.com/bmf-san/goblin/releases/tag/6.0.0)

データ構造やアルゴリズムの大きな改善をしていない、いわば小手先の改善ではあるので目を見張るほどの改善は残念ながら見られない。

なんとなく今採用しているデータ構造やアルゴリズムだとやはり難しいのだろうなという感じがする。（他所のルーター見ているともっと高度な木を採用しているのでそれはそうだという気がするが・・）

本題とはややずれるが、他のルーターとの比較をして改善のヒントが得られないかと思ってベンチマーカーを作成した。

[bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark)

比較してみると面白くて、ボロ負けなのがよく分かる。泣いた。

他ルーターの実装を研究したり、以前挫折した高度な木構造についての勉強等やって改善につなげていきたい。

# まとめ
- Goではベンチマークやプロファイリングが簡単にできる
- 推測するな、計測せよ
- 小手先の改善では大きな成果は出づらい（それはそう）

# 参考
- [github.com - google/pprof](https://github.com/google/pprof/blob/main/doc/README.md)
- [github.com - dgryski/go-perfbook](https://github.com/dgryski/go-perfbook)
- [dave.cheney.net - High Perfomance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html)
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
