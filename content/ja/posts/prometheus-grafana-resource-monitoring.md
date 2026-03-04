---
title: "PrometheusとGrafanaでリソース監視環境を整える"
slug: "prometheus-grafana-resource-monitoring"
date: 2018-09-15
author: bmf-san
categories:
  - "インフラストラクチャ"
tags:
  - "Prometheus"
  - "Grafana"
draft: false
---

# 概要
PrometheusとGrafanaでリソース監視の環境を構築したのでメモ。

# 環境
- さくらVPS
- CentOS7系

# Prometheusをインストール
[Prometheus - Getting Started](https://prometheus.io/docs/prometheus/latest/getting_started/)に従ってprometheusをダウンロードする。

ダウンロードしたらprometheusを起動、ダッシュボードにアクセスできることを確認しておく。

ポートが開放されていない場合は開放しておく。

# node_exporterをインストール
[Prometheus - node_exporter
](https://prometheus.io/download/#node_exporter)からnode_exporterをダウンロード。
 
こちらも起動しておく。

ポートの確認も同様。

`Prometheus.yml`の`static_configs`の項目に以下を追加する。

```
- targets: ['localhost:9100']
```

`killall prometheus`してからprometheusを再起動する。

# Grafanaをインストール
[Grafana - Installing on RPM-based Linux (CentOS, Fedora, OpenSuse, RedHat)](http://docs.grafana.org/installation/debian/)に従ってgrafanaをダウンロードする。
Yum Repositoryを追加してインストールした。

インストールしたらgrafanaを起動、アクセスできることを確認しておく。

こちらもポートが開放されていない場合は開放しておく。

# PrometheusとGrafanaを連携
## ログイン
Grafanaのインストールが完了したら、Grafanaにアクセスして、まずはログインする。

初期のログイン情報はusernameがadmin、passwordがadmin。

ログイン後にログイン情報は変更できるので適宜調整。

## Data Sourceの設定
左側メニューにWindowsみたいなアイコンがあるので、それをクリックしてDashboards→Homeをクリック。

Data Sourceの設定をするのでAdd data sourceを選択。

設定方法は[さくらのナレッジ - PrometheusとGrafanaを組み合わせて監視用ダッシュボードを作る](https://knowledge.sakura.ad.jp/12058/)を参照。

HTTP settingsのURLがplacefolderのデザインのせいでデフォルトで指定されるものだと勘違いしていて、未設定のまま作業を進めていたらグラフがちゃんと生成されなかった。
設定するのを忘れずに。

## ダッシュボードのテンプレートを用意
Grapana LabsでPrometheus用のダッシュボードテンプレートを用意する。

[Prometheus systemby Thomas Cheronneau](https://grafana.com/dashboards/159)で`Copy ID to Clipboard`をクリック。

左側メニューの＋アイコン→Dashboards→Importを選択。

Grafana.com DashboardにIDをペース→Loadをクリック。

OptionsのData sourceでprometheus（PrometheusのData source）を選択。

# 所感
ざっと雑にまとめたがこれで監視ができるはず。
アラートとかもちゃんと設定できるらしいのでそのうちやってみたい。

# 参考
- [Qiita - Prometheus 環境構築手順](https://qiita.com/tSU_RooT/items/fec5b9217417758988ae)

