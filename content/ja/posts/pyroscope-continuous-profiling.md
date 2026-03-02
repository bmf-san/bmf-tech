---
title: "PyroscopeでContinuous Profiling"
slug: "pyroscope-continuous-profiling"
date: 2023-05-07
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "プロファイリング"
draft: false
---

# 概要
[Pyroscope](https://pyroscope.io/)というContinous Profilingのツールを導入してみた。

Continous Profilingについてはこちら参照。
[What is continuous profiling?](https://www.cncf.io/blog/2022/05/31/what-is-continuous-profiling/)

今年に入ってからGrafanaが買収したらしい。

[Grafana Labs が Pyroscope を買収してコード プロファイリング機能を追加](https://devops.com/grafana-labs-acquires-pyroscope-to-add-code-profiling-capability/)

買収してからは、Grafana Pyroscopeが正式名称？ぽい。

Grafanaにプラグインが用意されているので連携することもできるが、Pyroscope単体にもUIが用意されている。

[Demo](https://demo.pyroscope.io/?name=hotrod.python.frontend%7B%7D&query=)が用意されているので触ってみると何が見れるか分かりやすいかも。

OSSとしてコードが公開されているので実装が気になる場合は見に行くこともできる。

[grafana/pyroscope](https://github.com/grafana/pyroscope)

# 導入
構成について先に目を通しておくと良い。

[Pyroscope Agent](https://pyroscope.io/docs/agent-overview/)

## 1. Pyroscope Serverの導入
DockerHubにイメージがプッシュされているのでこちらを利用することができる。

[pyroscope/pyroscope](https://hub.docker.com/r/pyroscope/pyroscope)

Dockerの導入ガイドはこちら。

[Docker Guide](https://pyroscope.io/docs/docker-guide/)

Kubernetesの導入ガイドもある。

[Kubernetes/Helm](https://pyroscope.io/docs/kubernetes-helm-chart/)

## 2. アプリケーション側でプロファイリングの有効化
アプリケーション側でプロファイリングの設定およびエージェントのインストール。

基本はPush型での対応になるが、Goの場合はPull型の対応があり、Pyroscopeサーバーでターゲットの管理ができる。

[Go Pull Mode](https://pyroscope.io/docs/golang-pull-mode/)

GoでPull型で導入する場合は、次のようにターゲットを管理することができる。

```yml
---
scrape-configs:
  - job-name: pyroscope
    scrape-interval: 60s
    enabled-profiles: [cpu, mem, goroutines, mutex, block]
    static-configs:
      - application: foo
        spy-name: gospy
        targets:
          - foo:80
      - application: bar
        spy-name: gospy
        targets:
          - bar:81
```

ここでは設定していないが、データ保持期間は設定したほうが良さそう。デフォルトでは無制限に保持するらしい。
cf. [Data retention](https://pyroscope.io/docs/data-retention/)

# サンプルコード
[examples](https://github.com/grafana/pyroscope/tree/main/examples)に色々例が用意されている。

自分の管理しているアプリケーションで導入した例は下記。

[gobel-example](https://github.com/bmf-san/gobel-example/tree/master/pyroscope)

# ハマったところ
## パスワード認証の初期設定
PyroscopeにはAPI KEYやOAuth２、パスワード認証の仕組みが用意されている。

パスワード認証で初期の認証情報設定をする際にドキュメントを読み違えて少しハマった。

初期の認証情報をセットするには次のような感じで設定ファイルに記載する。

```yml
auth:
  internal:
    admin:
      name: USERNAME
      password: PASSWORD
    enabled: true
```

ちゃんと読むと[configuring-built-in-admin-user](https://pyroscope.io/docs/auth-internal/#configuring-built-in-admin-user)のところで記載があるのだが、CLIでしか設定変更できないものだと勘違いしてムダにハマってしまった...。

## pprof
Pyroscope側の話ではなく、Goのアプリケーション側の話だが、pprofの設定でハマった。

それについては[DefaultServeMux以外でpprofを使う方法](https://bmf-tech.com/posts/DefaultServeMux%e4%bb%a5%e5%a4%96%e3%81%a7pprof%e3%82%92%e4%bd%bf%e3%81%86%e6%96%b9%e6%b3%95)に記事を書いたのでそちらを参照。

# 所感
OSSで利用できるプロファイルングツールを以前から探し求めていたのだが、Pyroscopeは導入がしやすく、使いやすいUIが用意されていて良さそう。GoならPull型も対応していて良い。
