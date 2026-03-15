---
title: Golangでロードバランサーを実装する
description: "L4ロードバランサーをGoで実装し、ラウンドロビンスケジューリング・ヘルスチェック・リバースプロキシの仕組みを構築例を通して解説。"
slug: golang-load-balancer-implementation
date: 2022-01-01T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
  - ロードバランサー
  - ラウンドロビン
translation_key: golang-load-balancer-implementation
---


# 概要
この記事は[Makuake Advent Calendar 2021](https://adventar.org/calendars/6822)の24日目の記事です。（大遅刻しました・・）
ラウンドロビンで負荷分散するロードバランサーをGolangで自作してみるという話です。

# ロードバランサーとは何か
ロードバランサーはリクエストを複数のサーバーへ振り分けて負荷分散する（ロードバランシング）機能を持ったサーバーです。

![スクリーンショット 2022-01-01 23 05 20](/assets/images/posts/golang-load-balancer-implementation/147852643-0d5a6fab-1d8f-4d60-81f6-cf95091ca974.png)

サービスの可用性を高めてくれるリバースプロキシの一種です。
ロードバランサーの種類は大きく分けて2種類あります。アプリケーション層で負荷分散するL7ロードバランサーと、トランスポート層で負荷分散するL4ロードバランサーです。
ロードバランサーは、ロードバランシングの他、パーシステンス（セッション維持）とヘルスチェックの機能を兼ね備えています。

# ロードバランシングの種類
負荷分散には静的な方式と動的な方式のものでそれぞれ種類があります。
静的なものの代表的な方式としては、リクエストを均等に振り分けるRound Robinという方式があります。
動的なものの代表的な方式としては、リクエストの未処理数が最小のサーバーに振り分けるLeast Connectionという方式があります。

# パーシステンスの種類
パーシステンスはロードバランサーの複数の振り分け先のサーバー間でセッションを維持するための機能です。
大きく分けてSource address affinity persistenceという方式とCookie persistenceという2つの方式があります。
Source address affinity persistenceは送信元IPアドレスを見て振り分け先のサーバーを固定する方式です。
Cookie persistenceはセッション維持のためのCookieを発行して、Cookieを見て振り分け先のサーバーを固定する方式です。

# ヘルスチェックの種類
ヘルスチェックはロードバランサーが振り分け先のサーバーの稼働状況を確認する機能です。
ロードバランサーから振り分け先のサーバーにヘルスチェックするアクティブ型のヘルスチェック方式と、クライアントからのリクエストに対するレスポンスを監視する方式です。
アクティブチェックは利用するプロトコルによってはL3チェック、L4チェック、L7チェックといった種類に分別することができます。

# 実装
L4ロードバランサーをパッケージとして実装します。
ロードバランシングの種類はラウンドロビンで、ヘルスチェックはアクティブチェック・パッシブチェックのそれぞれ対応します。
パーシステンスはの対応はしません。

今回実装したコードは[github.com/bmf-san/godon](https://github.com/bmf-san/godon)にあります。

# リバースプロキシを実装
ロードバランサーはリバースプロキシの一種です。まずは簡単なリバースプロキシの実装から始めます。

Golangでは`httputil`を利用することで簡単に実装することができます。

```golang
package godon

import (
	"log"
	"net/http"
	"net/http/httputil"
)


func Serve() {
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = ":8081"
	}

	rp := &httputil.ReverseProxy{
		Director: director,
	}

	s := http.Server{
		Addr:    ":8080",
		Handler: rp,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
```

ここでは説明を省きますが、[pkg.go.dev/net/http/httputil#ReverseProxy](https://pkg.go.dev/net/http/httputil#ReverseProxy)をよく読んでおくと良いかと思います。

# Configの実装
簡単なロードバランサーなので複雑な設定を持ちませんが、jsonから設定を読み込むような設定の機能を実装しておきます。

```json
{
	"proxy": {
		"port": "8080"
	},
	"backends": [
		{
			"url": "http://localhost:8081/"
		},
		{
			"url": "http://localhost:8082/"
		},
		{
			"url": "http://localhost:8083/"
		},
		{
			"url": "http://localhost:8084/"
		}
	]
}
```

```golang
// ...

// Config is a configuration.
type Config struct {
	Proxy    Proxy     `json:"proxy"`
	Backends []Backend `json:"backends"`
}

// Proxy is a reverse proxy, and means load balancer.
type Proxy struct {
	Port string `json:"port"`
}

// Backend is servers which load balancer is transferred.
type Backend struct {
	URL    string `json:"url"`
	IsDead bool
	mu     sync.RWMutex
}

var cfg Config

// Serve serves a loadbalancer.
func Serve() {
	// ...

	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(data, &cfg)

	// ...
}
```

# ラウンドロビンの実装
次にラウンドロビンの実装をします。

均等にバックエンドのサーバーにリクエストを振り分けるのみで、バックエンドのサーバーの生死は問わない形で実装します。


```golang
// ...

var mu sync.Mutex
var idx int = 0

// lbHandler is a handler for loadbalancing
func lbHandler(w http.ResponseWriter, r *http.Request) {
	maxLen := len(cfg.Backends)
	// Round Robin
	mu.Lock()
	currentBackend := cfg.Backends[idx%maxLen]
	targetURL, err := url.Parse(cfg.Backends[idx%maxLen].URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	idx++
	mu.Unlock()
	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ServeHTTP(w, r)
}

// ...

var cfg Config

// Serve serves a loadbalancer.
func Serve() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(data, &cfg)

	s := http.Server{
		Addr:    ":" + cfg.Proxy.Port,
		Handler: http.HandlerFunc(lbHandler),
	}
	if err = s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
```

`sync.Mutex`を利用しているのは、複数のGoroutineが変数にアクセスすることによる競合状態を回避するためです。

試しに`sync.Mutex`を外して`go run -race server.go`でサーバー起動、複数端末から同時にリクエストするとrace conditionを確認することができます。

# アクティブチェックの実装
ここまでの実装では、ロードバランサーは異常なバックエンドに対してもリクエストを転送するようなロジックとなっています。

実際のユースケースでは異常なバックエンドにわざわざリクエストを転送してほしくはないので、異常なバックエンドを検知して、振り分け先から外れるようにします。


```golang
// Backend is servers which load balancer is transferred.
type Backend struct {
	URL    string `json:"url"`
	IsDead bool
	mu     sync.RWMutex
}

// SetDead updates the value of IsDead in Backend.
func (backend *Backend) SetDead(b bool) {
	backend.mu.Lock()
	backend.IsDead = b
	backend.mu.Unlock()
}

// GetIsDead returns the value of IsDead in Backend.
func (backend *Backend) GetIsDead() bool {
	backend.mu.RLock()
	isAlive := backend.IsDead
	backend.mu.RUnlock()
	return isAlive
}

var mu sync.Mutex
var idx int = 0

// lbHandler is a handler for loadbalancing
func lbHandler(w http.ResponseWriter, r *http.Request) {
	maxLen := len(cfg.Backends)
	// Round Robin
	mu.Lock()
	currentBackend := cfg.Backends[idx%maxLen]
	if currentBackend.GetIsDead() {
		idx++
	}
	targetURL, err := url.Parse(cfg.Backends[idx%maxLen].URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	idx++
	mu.Unlock()
	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		// NOTE: It is better to implement retry.
		log.Printf("%v is dead.", targetURL)
		currentBackend.SetDead(true)
		lbHandler(w, r)
	}
	reverseProxy.ServeHTTP(w, r)
}

var cfg Config

// Serve serves a loadbalancer.
func Serve() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(data, &cfg)

	s := http.Server{
		Addr:    ":" + cfg.Proxy.Port,
		Handler: http.HandlerFunc(lbHandler),
	}
	if err = s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
```

ロードバランサーがバックエンドにリクエストを転送したときにエラーを検知すると呼び出される`ErrorHandler`を実装します。
`ErrorHandler`では、正常にレスポンスを返さないバックエンドにフラグを立てて、もう一度ロードバランサーにリクエストを転送してもらうような形にしています。
ロードバランサーはフラグの立っているバックエンドにはリクエストを転送しないようにロジックを調整しています。

# パッシブチェックの実装
最後にパッシブチェックの実装をします。
パッシブチェックは、インターバルを指定してバックエンドサーバーのレスポンスを監視するだけです。
異常が検知されたバックエンドは、アクティブチェックのときと同じようにフラグが立てられます。

パッシブチェックを実装し終えた全てのコードは以下になります。

```golang
package godon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

// Config is a configuration.
type Config struct {
	Proxy    Proxy     `json:"proxy"`
	Backends []Backend `json:"backends"`
}

// Proxy is a reverse proxy, and means load balancer.
type Proxy struct {
	Port string `json:"port"`
}

// Backend is servers which load balancer is transferred.
type Backend struct {
	URL    string `json:"url"`
	IsDead bool
	mu     sync.RWMutex
}

// SetDead updates the value of IsDead in Backend.
func (backend *Backend) SetDead(b bool) {
	backend.mu.Lock()
	backend.IsDead = b
	backend.mu.Unlock()
}

// GetIsDead returns the value of IsDead in Backend.
func (backend *Backend) GetIsDead() bool {
	backend.mu.RLock()
	isAlive := backend.IsDead
	backend.mu.RUnlock()
	return isAlive
}

var mu sync.Mutex
var idx int = 0

// lbHandler is a handler for loadbalancing
func lbHandler(w http.ResponseWriter, r *http.Request) {
	maxLen := len(cfg.Backends)
	// Round Robin
	mu.Lock()
	currentBackend := cfg.Backends[idx%maxLen]
	if currentBackend.GetIsDead() {
		idx++
	}
	targetURL, err := url.Parse(cfg.Backends[idx%maxLen].URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	idx++
	mu.Unlock()
	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		// NOTE: It is better to implement retry.
		log.Printf("%v is dead.", targetURL)
		currentBackend.SetDead(true)
		lbHandler(w, r)
	}
	reverseProxy.ServeHTTP(w, r)
}

// pingBackend checks if the backend is alive.
func isAlive(url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Minute*1)
	if err != nil {
		log.Printf("Unreachable to %v, error:", url.Host, err.Error())
		return false
	}
	defer conn.Close()
	return true
}

// healthCheck is a function for healthcheck
func healthCheck() {
	t := time.NewTicker(time.Minute * 1)
	for {
		select {
		case <-t.C:
			for _, backend := range cfg.Backends {
				pingURL, err := url.Parse(backend.URL)
				if err != nil {
					log.Fatal(err.Error())
				}
				isAlive := isAlive(pingURL)
				backend.SetDead(!isAlive)
				msg := "ok"
				if !isAlive {
					msg = "dead"
				}
				log.Printf("%v checked %v by healthcheck", backend.URL, msg)
			}
		}
	}

}

var cfg Config

// Serve serves a loadbalancer.
func Serve() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(data, &cfg)

	go healthCheck()

	s := http.Server{
		Addr:    ":" + cfg.Proxy.Port,
		Handler: http.HandlerFunc(lbHandler),
	}
	if err = s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
```

# 所感
リトライの実装やパーシステンスの対応などができていませんが、Golangでは比較的簡単にロードバランサーを実装できることが分かったかと思います。

## 関連記事

- [Goで始めるコードのパフォーマンス改善](/ja/posts/go-performance-improvement/)
- [Goの並行・並列処理モデルとgoroutineスケジューリング](/ja/posts/go-concurrency-parallelism-models/)
- [GoのHTTP Routerを比較するベンチマーカーを実装した](/ja/posts/http-router-benchmark-go/)
- [自作HTTPルーターから新しいServeMuxへ](/ja/posts/custom-http-router-to-new-servemux/)

# 参考
- [qiita.com - Goでリバースプロキシつくるときにつかえる net/http/httputil.ReverseProxy の紹介](https://qiita.com/convto/items/64e8f090198a4cf7a4fc)
- [kasvith.me - Let's Create a Simple Load Balancer With Go](https://kasvith.me/posts/lets-create-a-simple-lb-go/)
- [dev.to - Build Load Balancer in Go](https://dev.to/b0r/build-load-balancer-in-go-1oo7)
- [en.wikipedia.org - Load_balancing](https://en.wikipedia.org/wiki/Load_balancing_(computing)#Others)
- [www.infraexpert.com - ロードバランサをはじめから](https://www.infraexpert.com/study/study24.html)
- [www.opensquare.co.jp - Module 5 – パーシステンス（Persistence）](https://www.opensquare.co.jp/lmfile/support/document/TraningG/Module-5.pdf)
- [ascii.jp - 知っておきたいロードバランサーの基礎技術](https://ascii.jp/elem/000/000/506/506272/)
- [www.f5.com - ヘルスチェック](https://www.f5.com/ja_jp/services/resources/glossary/health-check)
- [www.rworks.jp - ロードバランサー（LB）とは？仕組みやDNSラウンドロビンとの違いについて解説](https://www.rworks.jp/system/system-column/sys-entry/16305/)
- [docs.nginx.com - HTTP Load Balancing](https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/#choosing-a-load-balancing-method)
- [medium.com - Running multiple HTTP servers in Go](https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f)
- [news.mynavi.jp - ロードバランサーの基本的な役割についてあらためておさらい](https://news.mynavi.jp/techplus/article/load_balancer-1/)
- [github.com - yyyar/gobetween](https://github.com/yyyar/gobetween)
- [github.com - kasvith/simplelb](https://github.com/kasvith/simplelb)
- [github.com - arjunmahishi/loadbalancer-in-go](https://github.com/arjunmahishi/loadbalancer-in-go)
- [github.com - arbazsiddiqui/anabranch](https://github.com/arbazsiddiqui/anabranch)
