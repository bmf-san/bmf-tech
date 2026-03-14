---
title: GoのHTTP Routerを比較するベンチマーカーを実装した
description: "複数のGoルーターをベンチマーク比較し、静的ルート・パスパラメータなどのテストケース設計を通してパフォーマンス差分を測定。"
slug: http-router-benchmark-go
date: 2022-12-20T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Golang
translation_key: http-router-benchmark-go
---


[Makuake Advent Calendar 2022](https://adventar.org/calendars/8496)の5日目の記事です！

# 概要
GoのHTTP Routerのパフォーマンスを比較するためのベンチマーカーを実装した。

[bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark)

現在のところ以下のHTTP Routerを比較対象としている。

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

一部のテストケースでは、Goの標準である[net/http#ServeMux](https://pkg.go.dev/net/http#ServeMux)も対象となっている。

# 動機
[bmf-san/goblin](https://github.com/bmf-san/goblin)というHTTP Routerを自作している。

[bmf-san/goblin](https://github.com/bmf-san/goblin)はTrie木をベースとしている必要最低限の機能を持ったシンプルなHTTP Routerである。

[bmf-san/goblin](https://github.com/bmf-san/goblin)と他のHTTP Routerのパフォーマンスを比較することで、[bmf-san/goblin](https://github.com/bmf-san/goblin)の改善のヒントを得たいというのが動機である。

別の理由としては、[julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) の代わりになるベンチマーカーを自分でメンテナンスできるようにしたいと考えていたという動機もある。

[julienschmidt/go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark)は[julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)のメンテナーが管理しているリポジトリだが、近年でメンテナンスが止まってしているように見えたので、それならば自分で作ってみようという気持ちになり、ベンチマーカーを実装することにした。

# ベンチマーカーのテスト設計について
前提として、[bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark)はHTTP Routerのパフォーマンスを完全に比較するものではないということを明言しておきたい。

理由としては次の通りである。

- HTTP Routerごとに機能や仕様が異なり、全てを完璧に網羅したテストケースを用意することが現実ではないため、一部の仕様のみ考慮した比較になるため
- 各HTTP Routerが持つデータ構造やアルゴリズムに依っては得意・不得意があるため、定義したルーティングテストケースにより性能が十分に評価されない可能性があるため

従って、ベンチマークテストはHTTP Routerの特定の機能・仕様のみを対象としたテストケースになってしまうが、一定の性能差を測り知ることはできる。

[bmf-san/go-router-benchmark](https://github.com/bmf-san/go-router-benchmark)では、ルーティングの処理部分を性能測定している。
具体的にいうと、[http#Handler](https://pkg.go.dev/net/http#Handler)の `ServeHTTP` 関数をテストしている。

[benchmark_test.go#L21](https://github.com/bmf-san/go-router-benchmark/blob/main/benchmark_test.go#L21)

ルーティングを定義する処理についてテスト対象としていない。
ルーティングを定義する処理とは、ルーティング処理に必要なデータを登録する処理である。

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

ルーティング処理のテストケースとしては次の2つのテストケースを用意している。

- 静的なルート
- パスパラメータを使ったルート

それぞれのテストケースについて説明する。

## 静的なルート
静的なルートとは、`/foo/bar`のような可変のパラメータを持たないルートのことを指す。

このルートのテストでは次の4パターンの入力を用意している。

- `/`
- `/foo`
- `/foo/bar/baz/qux/quux`
- `/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred`

このテストケースではGoの標準である[net/http#ServeMux](https://pkg.go.dev/net/http#ServeMux)も比較対象としている。

## パスパラメータを使ったルート
パスパラメータを使ったルートとは、`/foo/:bar`のうような可変のパラメータを持つルートのことを指す。

このルートのテストでは次の3パターンの入力を用意している。

- `/foo/:bar`
- `/foo/:bar/:baz/:qux/:quux/:corge`
- `/foo/:bar/:baz/:qux/:quux/:corge/:grault/:garply/:waldo/:fred/:plugh`

HTTP Routerによっては、可変のパラメータのシンタックスが異なるため、それぞれのシンタックスに対応することも考慮している。

ex. [pathparam.go#L15](https://github.com/bmf-san/go-router-benchmark/blob/main/pathparam.go#L15)

# ベンチマークテスト実施結果
ベンチマークテストの実施結果は[<Public>go-router-benchmark](https://docs.google.com/spreadsheets/d/1DrDNGJXfquw_PED3-eMqWqh7qbCTVBWZtqaKPXtngxg/edit#gid=1913830192)にて公開している。

ベンチマークテストの実行環境は次の通り。

- go version: go1.19
- goos: darwin
- goarch: amd64
- pkg: github.com/go-router-benchmark
- cpu: VirtualApple @ 2.50GHz

ベンチマーク結果の見方としては次の通り。
- time
  - 関数の実行回数
  - 実行回数が多いほどパフォーマンスが良いと考えられる
- ns/op
  - 関数の1回あたりの実行に要した時間
  - 時間が少ないほどパフォーマンスが良いと考えられる
- B/op
  - 関数の実行ごとに割当されたメモリのサイズ
  - 少なければ少ないほどパフォーマンスが良いと考えられる
- allocs/op
  - 関数の1回あたりの実行で行われたメモリアロケーションの回数
  - 少なければ少ないほどパフォーマンスが良いと考えられる

cf. [bmf-tech.com - Goで始めるコードのパフォーマンス改善](https://bmf-tech.com/posts/Go%e3%81%a7%e5%a7%8b%e3%82%81%e3%82%8b%e3%82%b3%e3%83%bc%e3%83%89%e3%81%ae%e3%83%91%e3%83%95%e3%82%a9%e3%83%bc%e3%83%9e%e3%83%b3%e3%82%b9%e6%94%b9%e5%96%84)

それぞれのテストケースの結果について記載する。

## 静的なルート
静的なルートについては、標準の[net/http#ServeMux](https://pkg.go.dev/net/http#ServeMux)よりも性能が良いか、等しいかが一つの比較ポイントであるように思う。
パフォーマンスの良さを謳っているHTTP Routerはやはり標準よりよい結果を出している。

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

## パスパラメータを使ったルート
パスパラメータを使ったルートについては、パラメータの数が増えるに従って性能が大きく劣化していくものと、控えめに劣化していくグループに別れた。

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

# 結論
性能の良いHTTP Routerは、各テストケースにおいて性能劣化が少ないことが分かる。
これは実装が最適化されていることを示す明確な傾向であると思われる。

パフォーマンスの良いHTTP Routerの実装をいくつか調べると、より高度な木構造を採用していることが分かった。
例えば、Echo,gin,httprouter,bon,chiはRadix tree (Patricia trie)を、dencoはdouble arrayを採用している。

[bmf-san/goblin](https://github.com/bmf-san/goblin)について、トライツリーを独自に拡張したもので、あまり最適化されておらず、他のHTTP Routerに比べて性能が低いことがよく分かった。(改善できるよう頑張る...)

一方で、性能が低いと思われるHTTP Routerの中には、多機能さが性能を落としている可能性がありそうであった。

テストケースを追加することでHTTP Routerごとの性能傾向が更に得られそうだと感じたため、時間があれば対応してみようと思う。

