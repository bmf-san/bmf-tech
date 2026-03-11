---
title: Implemented a Benchmark for Comparing Go HTTP Routers
slug: http-router-benchmark-go
date: 2022-12-20T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Golang
description: A benchmark tool to compare the performance of various Go HTTP Routers.
translation_key: http-router-benchmark-go
---

[Makuake Advent Calendar 2022](https://adventar.org/calendars/8496) - Day 5!

# Overview
Implemented a benchmark tool to compare the performance of Go HTTP Routers.

[bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark)

Currently, the following HTTP Routers are included in the comparison:

- [bmf-san/goblin](https://github.com/bmf-san/goblin)
- [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
- [go-chi/chi](https://github.com/go-chi/chi)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [uptrace/bunrouter](https://github.com/uptrace/bunrouter)
- [dimfeld/httptreemux](https://github.com/dimfeld/httptreemux)
- [beego/mux](https://github.com/beego/mux)
- [gorilla/mux](https://github.com/gorilla/mux)
- [nissy/bon](https://github.com/nissy/bon)
- [naoina/denco](https://github.com/naoina/denco)
- [labstack/echo](https://github.com/labstack/echo/v4)
- [gocraft/web](https://github.com/gocraft/web)
- [vardius/gorouter](https://github.com/vardius/gorouter)
- [go-ozzo/ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
- [lkeix/techbook13-sample](https://github.com/lkeix/techbook13-sample)

In some test cases, Go's standard [net/http#ServeMux](https://pkg.go.dev/net/http#ServeMux) is also included.

# Motivation
I am developing my own HTTP Router called [bmf-san/goblin](https://github.com/bmf-san/goblin).

[bmf-san/goblin](https://github.com/bmf-san/goblin) is a simple HTTP Router based on a Trie tree with minimal necessary features.

The motivation is to gain insights for improving [bmf-san/goblin](https://github.com/bmf-san/goblin) by comparing its performance with other HTTP Routers.

Another reason is the desire to maintain a benchmark tool as an alternative to [julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark), which seems to have stopped being maintained recently.

# About the Benchmark Test Design
I want to clarify that [bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark) does not fully compare the performance of HTTP Routers.

Reasons include:

- Different features and specifications among HTTP Routers make it impractical to cover all test cases comprehensively, leading to comparisons based on limited specifications.
- The data structures and algorithms of each HTTP Router may have strengths and weaknesses, potentially leading to insufficient performance evaluation based on defined routing test cases.

Therefore, while the benchmark tests focus on specific features and specifications of HTTP Routers, they can still measure certain performance differences.

[bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark) measures the performance of the routing process, specifically testing the `ServeHTTP` function of [http#Handler](https://pkg.go.dev/net/http#Handler).

[benchmark_test.go#L21](https://github.com/bmf-san/go-router-benchmark/blob/main/benchmark_test.go#L21)

The process of defining routing is not included in the test target. The process of defining routing refers to registering the data necessary for routing processing.

```go
package main

import (
	"fmt"
	"net/http" )
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler) // here
	ListenAndServe(":8080", mux)
}
```

The routing process test cases include:

- Static routes
- Routes with path parameters

Each test case is explained below.

## Static Routes
Static routes refer to routes without variable parameters, like `/foo/bar`.

This test includes the following four input patterns:

- `/`
- `/foo`
- `/foo/bar/baz/qux/quux`
- `/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred`

This test case also includes Go's standard [net/http#ServeMux](https://pkg.go.dev/net/http#ServeMux) for comparison.

## Routes with Path Parameters
Routes with path parameters include variable parameters, like `/foo/:bar`.

This test includes the following three input patterns:

- `/foo/:bar`
- `/foo/:bar/:baz/:qux/:quux/:corge`
- `/foo/:bar/:baz/:qux/:quux/:corge/:grault/:garply/:waldo/:fred/:plugh`

Since the syntax for variable parameters differs among HTTP Routers, each syntax is considered.

ex. [pathparam.go#L15](https://github.com/bmf-san/go-router-benchmark/blob/main/pathparam.go#L15)

# Benchmark Test Results
The results of the benchmark tests are publicly available at [<Public>go-router-benchmark](https://docs.google.com/spreadsheets/d/1DrDNGJXfquw_PED3-eMqWqh7qbCTVBWZtqaKPXtngxg/edit#gid=1913830192).

The benchmark test environment is as follows:

- go version: go1.19
- goos: darwin
- goarch: amd64
- pkg: github.com/go-router-benchmark
- cpu: VirtualApple @ 2.50GHz

The benchmark results can be interpreted as follows:
- time
  - Number of function executions
  - More executions indicate better performance
- ns/op
  - Time taken per function execution
  - Less time indicates better performance
- B/op
  - Memory allocated per function execution
  - Less memory indicates better performance
- allocs/op
  - Number of memory allocations per function execution
  - Fewer allocations indicate better performance

cf. [bmf-tech.com - Improving Code Performance with Go](https://bmf-tech.com/posts/Go%e3%81%a7%e5%a7%8b%e3%82%81%e3%82%8b%e3%82%b3%e3%83%bc%e3%83%89%e3%81%ae%e3%83%91%e3%83%95%e3%82%a9%e3%83%bc%e3%83%9e%e3%83%b3%e3%82%b9%e6%94%b9%e5%96%84)

Results for each test case are detailed below.

## Static Routes
For static routes, a key comparison point is whether the performance is better or equal to the standard [net/http#ServeMux](https://pkg.go.dev/net/http#ServeMux). HTTP Routers claiming high performance indeed show better results than the standard.

### time
|       time        | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 24301910           | 22053468        | 13324357        | 8851803          |
| goblin            | 32296879           | 16738813        | 5753088         | 3111172          |
| httprouter        | 100000000          | 100000000       | 100000000       | 72498970         |
| chi               | 5396652            | 5350285         | 5353856         | 5415325          |
| gin               | 34933861           | 34088810        | 34136852        | 33966028         |
| bunrouter         | 63478486           | 54812665        | 53564055        | 54345159         |
| httptreemux       | 6669231            | 6219157         | 5278312         | 4300488          |
| beegomux          | 22320199           | 15369320        | 1000000         | 577272           |
| gorillamux        | 1807042            | 2104210         | 1904696         | 1869037          |
| bon               | 72425132           | 56830177        | 59573305        | 58364338         |
| denco             | 90249313           | 92561344        | 89325312        | 73905086         |
| echo              | 41742093           | 36207878        | 23962478        | 12379764         |
| gocraftweb        | 1284613            | 1262863         | 1000000         | 889360           |
| gorouter          | 21622920           | 28592134        | 15582778        | 9636147          |
| ozzorouting       | 31406931           | 34989970        | 24825552        | 19431296         |
| techbook13-sample | 8176849            | 6349896         | 2684418         | 1384840          |

<iframe width="467" height="741" seamless frameborder="0" scrolling="no" src="https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=800028423&amp;format=interactive"></iframe>

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=800028423&format=interactive)

### nsop
|       nsop        | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 50.44              | 54.97           | 89.81           | 135.2            |
| goblin            | 36.63              | 69.9            | 205.2           | 382.7            |
| httprouter        | 10.65              | 10.74           | 10.75           | 16.42            |
| chi               | 217.2              | 220.1           | 216.7           | 221.5            |
| gin               | 34.53              | 34.91           | 34.69           | 35.04            |
| bunrouter         | 18.77              | 21.78           | 22.41           | 22               |
| httptreemux       | 178.8              | 190.9           | 227.2           | 277.7            |
| beegomux          | 55.07              | 74.69           | 1080            | 2046             |
| gorillamux        | 595.7              | 572.8           | 626.5           | 643.3            |
| bon               | 15.75              | 20.17           | 18.87           | 19.16            |
| denco             | 14                 | 13.03           | 13.4            | 15.87            |
| echo              | 28.17              | 32.83           | 49.82           | 96.77            |
| gocraftweb        | 929.4              | 948.8           | 1078            | 1215             |
| gorouter          | 55.16              | 37.64           | 76.6            | 124.1            |
| ozzorouting       | 42.62              | 34.22           | 48.12           | 61.6             |
| techbook13-sample | 146.1              | 188.4           | 443.5           | 867.8            |

<iframe width="467" height="741" seamless frameborder="0" scrolling="no" src="https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1691114342&amp;format=interactive"></iframe>

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1691114342&format=interactive)

### bop
|        bop        | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 0                  | 0               | 0               | 0                |
| goblin            | 0                  | 16              | 80              | 160              |
| httprouter        | 0                  | 0               | 0               | 0                |
| chi               | 304                | 304             | 304             | 304              |
| gin               | 0                  | 0               | 0               | 0                |
| bunrouter         | 0                  | 0               | 0               | 0                |
| httptreemux       | 328                | 328             | 328             | 328              |
| beegomux          | 32                 | 32              | 32              | 32               |
| gorillamux        | 720                | 720             | 720             | 720              |
| bon               | 0                  | 0               | 0               | 0                |
| denco             | 0                  | 0               | 0               | 0                |
| echo              | 0                  | 0               | 0               | 0                |
| gocraftweb        | 288                | 288             | 352             | 432              |
| gorouter          | 0                  | 0               | 0               | 0                |
| ozzorouting       | 0                  | 0               | 0               | 0                |
| techbook13-sample | 304                | 308             | 432             | 872              |

<iframe width="467" height="741" seamless frameborder="0" scrolling="no" src="https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=675738282&amp;format=interactive"></iframe>

[Graph - bop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=675738282&format=interactive)

### allocs
|      allocs       | static-routes-root | static-routes-1 | static-routes-5 | static-routes-10 |
| ----------------- | ------------------ | --------------- | --------------- | ---------------- |
| servemux          | 0                  | 0               | 0               | 0                |
| goblin            | 0                  | 1               | 1               | 1                |
| httprouter        | 0                  | 0               | 0               | 0                |
| chi               | 2                  | 2               | 2               | 2                |
| gin               | 0                  | 0               | 0               | 0                |
| bunrouter         | 0                  | 0               | 0               | 0                |
| httptreemux       | 3                  | 3               | 3               | 3                |
| beegomux          | 1                  | 1               | 1               | 1                |
| gorillamux        | 7                  | 7               | 7               | 7                |
| bon               | 0                  | 0               | 0               | 0                |
| denco             | 0                  | 0               | 0               | 0                |
| echo              | 0                  | 0               | 0               | 0                |
| gocraftweb        | 6                  | 6               | 6               | 6                |
| gorouter          | 0                  | 0               | 0               | 0                |
| ozzorouting       | 0                  | 0               | 0               | 0                |
| techbook13-sample | 2                  | 3               | 11              | 21               |

<iframe width="467" height="741" seamless frameborder="0" scrolling="no" src="https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1054975173&amp;format=interactive"></iframe>

[Graph - allocs](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1054975173&format=interactive)

## Routes with Path Parameters
For routes with path parameters, there are groups that degrade significantly as the number of parameters increases, and those that degrade modestly.

### time
|       time        | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 1802690            | 492392             | 252274              |
| httprouter        | 25775940           | 10057874           | 6060843             |
| chi               | 4337922            | 2687157            | 1772881             |
| gin               | 29479381           | 15714673           | 9586220             |
| bunrouter         | 37098772           | 8479642            | 3747968             |
| httptreemux       | 2610324            | 1550306            | 706356              |
| beegomux          | 3177818            | 797472             | 343969              |
| gorillamux        | 1364386            | 470180             | 223627              |
| bon               | 6639216            | 4486780            | 3285571             |
| denco             | 20093167           | 8503317            | 4988640             |
| echo              | 30667137           | 12028713           | 6721176             |
| gocraftweb        | 921375             | 734821             | 466641              |
| gorouter          | 4678617            | 3038450            | 2136946             |
| ozzorouting       | 27126000           | 12228037           | 7923040             |
| techbook13-sample | 3019774            | 917042             | 522897              |

<iframe width="455" height="739" seamless frameborder="0" scrolling="no" src="https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1039813866&amp;format=interactive"></iframe>

[Graph - time](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1039813866&format=interactive)

### nsop
|       nsop        | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 652.4              | 2341               | 4504                |
| httprouter        | 45.73              | 117.4              | 204.2               |
| chi               | 276.4              | 442.8              | 677.6               |
| gin               | 40.21              | 76.39              | 124.3               |
| bunrouter         | 32.52              | 141.1              | 317.2               |
| httptreemux       | 399.7              | 778.5              | 1518                |
| beegomux          | 377.2              | 1446               | 3398                |
| gorillamux        | 850.3              | 2423               | 5264                |
| bon               | 186.5              | 269.6              | 364.4               |
| denco             | 60.47              | 139.4              | 238.7               |
| echo              | 39.36              | 99.6               | 175.7               |
| gocraftweb        | 1181               | 1540               | 2280                |
| gorouter          | 256.4              | 393                | 557.6               |
| ozzorouting       | 43.66              | 99.52              | 150.4               |
| techbook13-sample | 380.7              | 1154               | 2150                |

<iframe width="460" height="739" seamless frameborder="0" scrolling="no" src="https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1534246873&amp;format=interactive"></iframe>

[Graph - nsop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=1534246873&format=interactive)

### bop
|        bop        | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 409                | 962                | 1608                |
| httprouter        | 32                 | 160                | 320                 |
| chi               | 304                | 304                | 304                 |
| gin               | 0                  | 0                  | 0                   |
| bunrouter         | 0                  | 0                  | 0                   |
| httptreemux       | 680                | 904                | 1742                |
| beegomux          | 672                | 672                | 1254                |
| gorillamux        | 1024               | 1088               | 1751                |
| bon               | 304                | 304                | 304                 |
| denco             | 32                 | 160                | 320                 |
| echo              | 0                  | 0                  | 0                   |
| gocraftweb        | 656                | 944                | 1862                |
| gorouter          | 360                | 488                | 648                 |
| ozzorouting       | 0                  | 0                  | 0                   |
| techbook13-sample | 432                | 968                | 1792                |

<iframe width="460" height="739" seamless frameborder="0" scrolling="no" src="https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=73824357&amp;format=interactive"></iframe>

[Graph - bop](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=73824357&format=interactive)

### allocs
|      allocs       | pathparam-routes-1 | pathparam-routes-5 | pathparam-routes-10 |
| ----------------- | ------------------ | ------------------ | ------------------- |
| goblin            | 6                  | 13                 | 19                  |
| httprouter        | 1                  | 1                  | 1                   |
| chi               | 2                  | 2                  | 2                   |
| gin               | 0                  | 0                  | 0                   |
| bunrouter         | 0                  | 0                  | 0                   |
| httptreemux       | 6                  | 9                  | 11                  |
| beegomux          | 5                  | 5                  | 6                   |
| gorillamux        | 8                  | 8                  | 9                   |
| bon               | 2                  | 2                  | 2                   |
| denco             | 1                  | 1                  | 1                   |
| echo              | 0                  | 0                  | 0                   |
| gocraftweb        | 9                  | 12                 | 14                  |
| gorouter          | 4                  | 4                  | 4                   |
| ozzorouting       | 0                  | 0                  | 0                   |
| techbook13-sample | 10                 | 33                 | 59                  |

<iframe width="460" height="739" seamless frameborder="0" scrolling="no" src="https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=344550080&amp;format=interactive"></iframe>

[Graph - allocs](https://docs.google.com/spreadsheets/d/e/2PACX-1vRiWBjJim4v_XyoN45s4VVQXD-PIBHKjyVfOv5OX37376SZ9GvL5bmqQegLl-5arBpD-3hhTKTEgkIj/pubchart?oid=344550080&format=interactive)

# Conclusion
It is evident that high-performance HTTP Routers exhibit minimal performance degradation across test cases. This suggests a clear trend of optimized implementations.

Upon examining the implementations of high-performance HTTP Routers, it is found that they employ more advanced tree structures. For instance, Echo, gin, httprouter, bon, and chi use Radix trees (Patricia tries), while denco uses a double array.

Regarding [bmf-san/goblin](https://github.com/bmf-san/goblin), it is an independently extended trie tree that is not highly optimized, resulting in lower performance compared to other HTTP Routers. (I'll strive to improve it...)

On the other hand, some HTTP Routers with seemingly lower performance may have their performance affected by their multifunctionality.

I feel that adding more test cases could further reveal performance trends for each HTTP Router, so I plan to address this if time permits.