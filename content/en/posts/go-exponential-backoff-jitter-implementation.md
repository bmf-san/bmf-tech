---
title: Implementing Exponential Backoff and Jitter in Go
description: 'Implement exponential backoff with jitter in Go for robust retry logic. Covers multiplying retry delays, adding randomness to prevent thundering herd, retry limits, and timeouts.'
slug: go-exponential-backoff-jitter-implementation
date: 2024-01-31T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Exponential Backoff
  - Retry
  - Jitter
translation_key: go-exponential-backoff-jitter-implementation
---

# Exponential Backoff
A method to periodically retry failed requests by multiplicatively increasing the delay of requests (delaying the retry interval).

ex. The first retry is after 1 second, the second is after 2 seconds, the third is after 4 seconds, the fourth is after 8 seconds...

In retry design, it is necessary to consider not only backoff but also retry limits and timeouts (connection timeout and request timeout).

# Jitter
A method to prevent requests that failed simultaneously from retrying at the same time by adding a random value to the retry interval of exponential backoff.

With simple exponential intervals, the retry intervals become the same, so jitter is introduced to add temporal variation.

# Implementing Exponential Backoff and Jitter
Here's a simple implementation.

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
			// Continue retrying due to error
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

I referred to the following article for the jitter algorithm, but I'm not entirely confident that it's correctly implemented. It's a bit fuzzy, so there might be some logical oversights.
cf. [aws.amazon.com - Exponential Backoff And Jitter](https://aws.amazon.com/jp/blogs/architecture/exponential-backoff-and-jitter/)

# Thoughts
The implementation is a bit rough, but I got the gist!

# References
- [en.wikipedia.org - Exponential backoff](https://en.wikipedia.org/wiki/Exponential_backoff)
- [ebc-2in2crc.hatenablog.jp - A memo on confirming and minimally implementing exponential backoff](https://ebc-2in2crc.hatenablog.jp/entry/2020/12/19/220801)
- [aws.amazon.com - Timeouts, retries, and backoff with jitter](https://aws.amazon.com/jp/builders-library/timeouts-retries-and-backoff-with-jitter/)
- [aws.amazon.com - Exponential Backoff And Jitter](https://aws.amazon.com/jp/blogs/architecture/exponential-backoff-and-jitter/)
- [qiita.com - Efficient approach to retry processing "Exponential Backoff" overview and implementation in Go](https://qiita.com/po3rin/items/c80dea298f16a2625dbe)
- [zenn.dev - Implementing Exponential Backoff And Jitter in Go](https://zenn.dev/sinozu/articles/5c0457876be42e)