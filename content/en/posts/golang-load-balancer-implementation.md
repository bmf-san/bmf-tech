---
title: Implementing a Load Balancer in Golang
slug: golang-load-balancer-implementation
date: 2022-01-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
  - Load Balancer
  - Round Robin
translation_key: golang-load-balancer-implementation
---

# Overview
This article is the 24th entry for the [Makuake Advent Calendar 2021](https://adventar.org/calendars/6822). (I am very late...)
It's about creating a custom load balancer in Golang that distributes load using round-robin.

# What is a Load Balancer?
A load balancer is a server that distributes requests to multiple servers to balance the load (load balancing).

![Screenshot 2022-01-01 23 05 20](https://user-images.githubusercontent.com/13291041/147852643-0d5a6fab-1d8f-4d60-81f6-cf95091ca974.png)

It is a type of reverse proxy that enhances service availability. There are two main types of load balancers: L7 load balancers that distribute load at the application layer and L4 load balancers that do so at the transport layer. Besides load balancing, load balancers also provide persistence (session maintenance) and health check functionalities.

# Types of Load Balancing
Load balancing can be static or dynamic. A representative static method is Round Robin, which distributes requests evenly. A representative dynamic method is Least Connection, which distributes requests to the server with the fewest unprocessed requests.

# Types of Persistence
Persistence is a feature that maintains sessions across multiple servers that a load balancer distributes to. There are two main types: Source address affinity persistence, which fixes the destination server based on the source IP address, and Cookie persistence, which issues a cookie for session maintenance and fixes the destination server based on the cookie.

# Types of Health Checks
Health checks are a feature of load balancers that check the operational status of destination servers. There are active health checks, where the load balancer checks the destination servers, and passive checks, which monitor responses to client requests. Active checks can be categorized into L3, L4, and L7 checks depending on the protocol used.

# Implementation
We will implement an L4 load balancer as a package. The load balancing method will be round-robin, and it will support both active and passive health checks. Persistence will not be supported.

The code implemented this time is available at [github.com/bmf-san/godon](https://github.com/bmf-san/godon).

# Implementing a Reverse Proxy
A load balancer is a type of reverse proxy. Let's start with a simple reverse proxy implementation.

In Golang, you can easily implement it using `httputil`.

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

I will omit the explanation here, but it would be good to read [pkg.go.dev/net/http/httputil#ReverseProxy](https://pkg.go.dev/net/http/httputil#ReverseProxy) thoroughly.

# Implementing Config
Since this is a simple load balancer, it does not have complex settings, but we will implement a feature to read settings from JSON.

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

# Implementing Round Robin
Next, we will implement round-robin.

We will implement it so that requests are evenly distributed to backend servers without considering the status of the backend servers.


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

The use of `sync.Mutex` is to avoid race conditions caused by multiple Goroutines accessing the variable.

Try removing `sync.Mutex` and start the server with `go run -race server.go`, then send requests from multiple terminals simultaneously to observe the race condition.

# Implementing Active Check
So far, the implementation allows the load balancer to forward requests even to abnormal backends.

In real use cases, you wouldn't want requests to be forwarded to abnormal backends, so we will detect abnormal backends and exclude them from the distribution.


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

We implement `ErrorHandler`, which is called when the load balancer detects an error while forwarding a request to a backend. In `ErrorHandler`, a flag is set for backends that do not return a normal response, and the load balancer is requested to forward the request again. The load balancer is adjusted so that it does not forward requests to backends with flags set.

# Implementing Passive Check
Finally, we will implement passive checks. Passive checks simply monitor the response of backend servers at specified intervals. Backends detected as abnormal are flagged the same way as in active checks.

The complete code after implementing passive checks is as follows.

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

# Thoughts
Although retry implementation and persistence support are not covered, I hope you found that implementing a load balancer in Golang is relatively straightforward.

# References
- [qiita.com - Introduction to net/http/httputil.ReverseProxy for creating a reverse proxy in Go](https://qiita.com/convto/items/64e8f090198a4cf7a4fc)
- [kasvith.me - Let's Create a Simple Load Balancer With Go](https://kasvith.me/posts/lets-create-a-simple-lb-go/)
- [dev.to - Build Load Balancer in Go](https://dev.to/b0r/build-load-balancer-in-go-1oo7)
- [en.wikipedia.org - Load_balancing](https://en.wikipedia.org/wiki/Load_balancing_(computing)#Others)
- [www.infraexpert.com - Introduction to Load Balancers](https://www.infraexpert.com/study/study24.html)
- [www.opensquare.co.jp - Module 5 – Persistence](https://www.opensquare.co.jp/lmfile/support/document/TraningG/Module-5.pdf)
- [ascii.jp - Basic Technologies of Load Balancers You Should Know](https://ascii.jp/elem/000/000/506/506272/)
- [www.f5.com - Health Check](https://www.f5.com/ja_jp/services/resources/glossary/health-check)
- [www.rworks.jp - What is a Load Balancer? Explaining the Mechanism and Differences from DNS Round Robin](https://www.rworks.jp/system/system-column/sys-entry/16305/)
- [docs.nginx.com - HTTP Load Balancing](https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/#choosing-a-load-balancing-method)
- [medium.com - Running multiple HTTP servers in Go](https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f)
- [news.mynavi.jp - Reviewing the Basic Roles of Load Balancers](https://news.mynavi.jp/techplus/article/load_balancer-1/)
- [github.com - yyyar/gobetween](https://github.com/yyyar/gobetween)
- [github.com - kasvith/simplelb](https://github.com/kasvith/simplelb)
- [github.com - arjunmahishi/loadbalancer-in-go](https://github.com/arjunmahishi/loadbalancer-in-go)
- [github.com - arbazsiddiqui/anabranch](https://github.com/arbazsiddiqui/anabranch)