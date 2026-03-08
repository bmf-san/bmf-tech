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
This article is the 24th entry of the [Makuake Advent Calendar 2021](https://adventar.org/calendars/6822). (I am very late..)
This is a story about creating a load balancer that performs load balancing using Round Robin in Golang.

# What is a Load Balancer?
A load balancer is a server that has the function of distributing requests to multiple servers for load balancing.

![Screenshot 2022-01-01 23 05 20](https://user-images.githubusercontent.com/13291041/147852643-0d5a6fab-1d8f-4d60-81f6-cf95091ca974.png)

It is a type of reverse proxy that enhances service availability. There are two main types of load balancers: L7 load balancers that perform load balancing at the application layer, and L4 load balancers that perform load balancing at the transport layer. In addition to load balancing, load balancers also have functions for persistence (session maintenance) and health checks.

# Types of Load Balancing
Load balancing can be static or dynamic, each with its own methods. A representative static method is Round Robin, which evenly distributes requests. A representative dynamic method is Least Connection, which distributes requests to the server with the fewest unprocessed requests.

# Types of Persistence
Persistence is a function that maintains sessions among multiple destination servers of the load balancer. There are two main types: Source address affinity persistence and Cookie persistence. Source address affinity persistence fixes the destination server based on the source IP address. Cookie persistence issues a cookie for session maintenance and fixes the destination server based on the cookie.

# Types of Health Checks
Health checks are functions that allow the load balancer to check the operational status of the destination servers. There are active health check methods where the load balancer actively checks the destination servers and methods that monitor responses to requests from clients. Active checks can be categorized into L3 checks, L4 checks, and L7 checks depending on the protocol used.

# Implementation
We will implement an L4 load balancer as a package. The type of load balancing will be Round Robin, and health checks will support both active and passive checks. Persistence will not be supported.

The code implemented this time can be found at [github.com/bmf-san/godon](https://github.com/bmf-san/godon).

# Implementing a Reverse Proxy
A load balancer is a type of reverse proxy. We will start with a simple implementation of a reverse proxy.

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

I will skip the explanation here, but it would be good to read [pkg.go.dev/net/http/httputil#ReverseProxy](https://pkg.go.dev/net/http/httputil#ReverseProxy).

# Implementing Config
Since it is a simple load balancer, it does not have complex settings, but we will implement a feature to read settings from JSON.

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
Next, we will implement Round Robin.

We will implement it in a way that evenly distributes requests to the backend servers without considering the health of the backend servers.

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

We use `sync.Mutex` to avoid race conditions caused by multiple Goroutines accessing the variable.

If you try removing `sync.Mutex` and run `go run -race server.go` to start the server, you can confirm the race condition by sending requests simultaneously from multiple terminals.

# Implementing Active Check
In the implementation so far, the load balancer has logic to forward requests to abnormal backends.

In actual use cases, we do not want to forward requests to abnormal backends, so we will detect abnormal backends and exclude them from the distribution.

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

We implement the `ErrorHandler` that is called when the load balancer detects an error while forwarding a request to the backend. In `ErrorHandler`, we flag the backend that does not return a normal response and ask it to forward the request to the load balancer again. The load balancer adjusts its logic not to forward requests to backends with the flag set.

# Implementing Passive Check
Finally, we will implement passive checks. Passive checks simply monitor the responses of backend servers at specified intervals. Abnormal backends are flagged just like in active checks.

The complete code after implementing passive checks is as follows:

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
Although I have not implemented retry logic or persistence support, I hope you can see that it is relatively easy to implement a load balancer in Golang.

# References
- [qiita.com - Introduction to net/http/httputil.ReverseProxy for creating a reverse proxy in Go](https://qiita.com/convto/items/64e8f090198a4cf7a4fc)
- [kasvith.me - Let's Create a Simple Load Balancer With Go](https://kasvith.me/posts/lets-create-a-simple-lb-go/)
- [dev.to - Build Load Balancer in Go](https://dev.to/b0r/build-load-balancer-in-go-1oo7)
- [en.wikipedia.org - Load_balancing](https://en.wikipedia.org/wiki/Load_balancing_(computing)#Others)
- [www.infraexpert.com - Starting from Load Balancer](https://www.infraexpert.com/study/study24.html)
- [www.opensquare.co.jp - Module 5 – Persistence](https://www.opensquare.co.jp/lmfile/support/document/TraningG/Module-5.pdf)
- [ascii.jp - Basic Technologies of Load Balancers You Should Know](https://ascii.jp/elem/000/000/506/506272/)
- [www.f5.com - Health Check](https://www.f5.com/ja_jp/services/resources/glossary/health-check)
- [www.rworks.jp - What is a Load Balancer? Explanation of its Mechanism and Differences from DNS Round Robin](https://www.rworks.jp/system/system-column/sys-entry/16305/)
- [docs.nginx.com - HTTP Load Balancing](https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/#choosing-a-load-balancing-method)
- [medium.com - Running multiple HTTP servers in Go](https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f)
- [news.mynavi.jp - A Review of the Basic Roles of Load Balancers](https://news.mynavi.jp/techplus/article/load_balancer-1/)
- [github.com - yyyar/gobetween](https://github.com/yyyar/gobetween)
- [github.com - kasvith/simplelb](https://github.com/kasvith/simplelb)
- [github.com - arjunmahishi/loadbalancer-in-go](https://github.com/arjunmahishi/loadbalancer-in-go)
- [github.com - arbazsiddiqui/anabranch](https://github.com/arbazsiddiqui/anabranch)