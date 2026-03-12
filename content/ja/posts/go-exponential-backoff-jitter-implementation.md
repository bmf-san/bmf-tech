---
title: Goで指数バックオフとジッターを実装してみる
description: Goで指数バックオフとジッターを実装してみるの手順と実践例を詳しく解説します。
slug: go-exponential-backoff-jitter-implementation
date: 2024-01-31T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - 指数バックオフ
  - リトライ
  - ジッター
translation_key: go-exponential-backoff-jitter-implementation
---


# 指数バックオフ(Exponential backoff)
リクエストの遅延を乗算的に増加させる（リトライ間隔を遅延させていく）形で失敗したリクエストを定期的に再試行（リトライ）する手法。

ex. 1回目のリトライは1秒後、2回目は2秒後、3回目は4秒後、4回目は8秒後...

リトライ設計においては、バックオフのみに依存するのではなく、リトライ上限やタイムアウト（接続タイムアウトとリクエストタイムアウト）も考慮する必要がある。

# ジッター
指数バックオフのリトライ間隔にランダムな値を加えることで、同時に失敗したリクエストが同時に再試行するのを防ぐ手法。

単純な指数的間隔だとリトライ間隔が同じになってしまうため、時間的ゆらぎを持たせるために導入される。

# 指数バックオフとジッターの実装
簡易的に実装するとしたらこんな感じだろうか。

```go
package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

// Retryer is a retryer.
type Retryer struct {
	MaxRetryCount int
	RetryCount    int
	Jitter        *Jitter
}

func NewRetryer(mrc int) *Retryer {
	return &Retryer{
		MaxRetryCount: mrc,
		RetryCount:    0,
		Jitter: &Jitter{
			base:  10,
			cap:   100,
			sleep: 10,
		},
	}
}

func (r *Retryer) Retry(ja string, f func() error) {
	for i := r.RetryCount; i < r.MaxRetryCount; i++ {
		var d time.Duration
		switch ja {
		case jitterAlgoFull:
			d = r.Jitter.FullJitter(r.RetryCount)
		case jitterAlgoEqual:
			d = r.Jitter.EqualJitter(r.RetryCount)
		case jitterAlgoDecorrelated:
			d = r.Jitter.DecorrelatedJitter()
		}
		time.Sleep(d)
		err := f()
		log.Printf("retry %d times\n", i)
		if err != nil {
			log.Println(err)
			// エラーなのでretryを継続
			continue
		}
	}

	log.Println("retry done")
	return
}

const jitterAlgoFull = "full"
const jitterAlgoEqual = "equal"
const jitterAlgoDecorrelated = "decorrelated"

// Jitter is a retryer with jitter.
type Jitter struct {
	base  int
	cap   int
	sleep int /// for decorrelated jitter
}

// FullJitter is a full jitter algo.
// sleep = random_between(0 min(cap, base * 2 ** attempt))
// see: https://aws.typepad.com/sajp/2015/03/backoff.html
func (j *Jitter) FullJitter(retryCount int) time.Duration {
	sleep := rand.Intn(min(j.cap, (j.base * int(math.Pow(2, float64(retryCount))))))
	return time.Duration(sleep) * time.Second
}

// EqualJitter is a full equal algo.
// temp = min(cap, base * 2 ** attempt)
// sleep = temp / 2 + random_between(0, temp / 2)
// see: https://aws.typepad.com/sajp/2015/03/backoff.html
func (j *Jitter) EqualJitter(retryCount int) time.Duration {
	temp := rand.Intn(min(j.cap, (j.base * int(math.Pow(2, float64(retryCount))))))
	sleep := (int(math.Ceil(float64(temp/2))) + rand.Intn(int(math.Ceil(float64(temp/2)))))
	return time.Duration(sleep) * time.Second
}

// DecorrelatedJitter is a decorrelated jitter algo.
// sleep = min(cap, random_between(base, sleep * 3))
// see: https://aws.typepad.com/sajp/2015/03/backoff.html
func (j *Jitter) DecorrelatedJitter() time.Duration {
	randomBetween := func(min, max int) int {
		return rand.Intn(max-min) + min
	}
	sleep := min(j.cap, randomBetween(j.base, j.sleep*3))
	j.sleep = sleep
	return time.Duration(sleep) * time.Second
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	r := NewRetryer(5)
	r.Retry(jitterAlgoFull, func() error {
		return fmt.Errorf("retry error")
	})
	r.Retry(jitterAlgoEqual, func() error {
		return fmt.Errorf("retry error")
	})
	r.Retry(jitterAlgoDecorrelated, func() error {
		return fmt.Errorf("retry error")
	})
}
```

ジッターのアルゴリズムは下記記事を参考にしたが、ちゃんと正しく反映できているかちょっと自信がない。ファジーなのでロジックに考慮漏れがある。
cf. [aws.amazon.com - Exponential Backoff And Jitter](https://aws.amazon.com/jp/blogs/architecture/exponential-backoff-and-jitter/)

# 所感
実装雑にしてしまったが雰囲気はわかった！

# 参考
- [en.wikipedia.org - Exponential backoff](https://en.wikipedia.org/wiki/Exponential_backoff)
- [ebc-2in2crc.hatenablog.jp - いまさら指数関数的バックオフを確認と最小限の実装をしてみたメモ。](https://ebc-2in2crc.hatenablog.jp/entry/2020/12/19/220801)
- [aws.amazon.com - ジッターを伴うタイムアウト、再試行、およびバックオフ](https://aws.amazon.com/jp/builders-library/timeouts-retries-and-backoff-with-jitter/)
- [aws.amazon.com - Exponential Backoff And Jitter](https://aws.amazon.com/jp/blogs/architecture/exponential-backoff-and-jitter/)
- [qiita.com - リトライ処理の効率的アプローチ「Exponential Backoff」の概要とGoによる実装](https://qiita.com/po3rin/items/c80dea298f16a2625dbe)
- [zenn.dev - Exponential Backoff And JitterをGoで実装](https://zenn.dev/sinozu/articles/5c0457876be42e)
