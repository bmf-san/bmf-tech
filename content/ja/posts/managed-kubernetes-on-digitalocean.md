---
title: DigitalOceanでマネージドKubernetesを使ってみる
description: DigitalOceanでマネージドKubernetesを使ってみるについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: managed-kubernetes-on-digitalocean
date: 2021-03-07T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Kubernetes
  - VPS
  - DigitalOcean
translation_key: managed-kubernetes-on-digitalocean
---


# DigitalOceanでマネージドKubernetesを使ってみる
プライベートの開発で学習も兼ねてk8sを利用したく、色々検討した結果、DigitalOceanが良さそうだったので、利用してみた。

[www.digitalocean.com - The best managed Kubernetes service is the one that’s made for you](https://www.digitalocean.com/blog/best-managed-kubernetes/)

新規に始める場合は、プロモリンクやクーポンを使うと良いと思う。

自分はプロモリンクから登録するのを忘れて最初クレジットをもらい損ねたが、問い合わせしたら良きに図らってもらえた（クレジットもらった）。ありがたや。

一応リファラルリンクを貼っておく。

https://m.do.co/c/9fbf85c22695

# Digital Oceanの良いところ
マネージドKubernetesの話しの前に、Digital Oceanの良いところを書いておく。

Kubernetesを利用しなくともDigitalOceanを利用したくなる機能が充実している。

- 安い
  - Droplet（VM)はスペック毎に時間課金
  - 転送量はスペックごとに無償枠がある。超過すると$0.02/1GB。
    - Conohaだと転送量は従量課金がないので月額料金だけで済む。
    - 転送量をダッシュボードから確認することができない。これは気になる。vmstatを使って監視する体制を作ったほうが良さそう。
  - コストパフォーマンスは優れていると思う
- 予算アラートがある
  - 管理画面から設定できる。メール通知のみっぽいが、API利用でSlack通知とか自分で調整できたりするかも？（ちゃんと見ていないので推測）
- メトリクスが取れる
  - 基本的なメトリクスが充実している。
  - しかもアラートが設定できてしまう。slackやメール通知が可能。
- openstackベースのAPIがある
  - terraformとかにもproviderがあるのでIaCはやりやすい。
    - [developers.digitalocean.com - documentation](https://developers.digitalocean.com/documentation/)
  - 日本のVPSだとConohaもAPIが整っていますね。
- DNSサービス無料
- エコシステムが充実している
  - [marketplace.digitalocean.com](https://marketplace.digitalocean.com/)
  - 1clickで起動できる起動テンプレートみたいものが結構充実している。試用してみたいアプリケーションとかあれば便利かも。
- 各種ドキュメントが分かりやすく整備されている
  - 個人差あるかもだが、安心感がある。
- サポートがいい感じ（個人の感想）
  - 冒頭のクレジットの件でサポートのticketを切ったが、1日以内に対応してもらった。たまたまかもしれないが期待できそう。
  - サポートが不評な部分もあるっぽい。
     - [www.websiteplanet.com - digitalocean/#support](www.websiteplanet.com/ja/web-hosting/digitalocean/#support)
- サメのアイコンに愛着が湧く。（個人の感想）

# Digital OceanのマネージドKubernetes概要
[www.digitalocean.com - docs/kubernetes](https://www.digitalocean.com/docs/kubernetes/)を参考に概要をまとめておく。

- マスターノードは無料
- ワーカーノードはドロップレット（Digital OceanではスケーラブルなVMをdropletといっている）の料金レートで課金
  - [www.digitalocean.com - plans-and-pricing](https://www.digitalocean.com/docs/droplets/#plans-and-pricing)
- 最小構成は$10
  - k8sは2GB以上メモリを要求するので最小構成でやろうとワーカーは2GB以上を選ぶ必要がある。
  - 1vCPU/2GB Memoryのスペックで、ワーカー1台の構成だと$10/Monthyで済む
    - ワーカーは2台以上推奨。1台だとクラスターアップデートやメンテナンスなどでダウンタイムが発生する可能性がある。
- CNIはCiliumを採用しているぽい
- 日本にリージョンはなし
  - [www.digitalocean.com - Regional Availability Matrix](https://www.digitalocean.com/docs/platform/availability-matrix/)
- CNCFに準拠しているKubernetes製品
  - [github.com - cncf/k8s-conformance](https://github.com/cncf/k8s-conformance)
- ロードバランサーも用意されている
  - [www.digitalocean.com - docs/kubernetes/how-to/add-load-balancers](https://www.digitalocean.com/docs/kubernetes/how-to/add-load-balancers/)
- ログローテーション
  - クラスターログは10MBでローテーションされる。
  - アクティブログと最後の2つが保持される。
- 制限
  - [www.digitalocean.com - docs/kubernetes/#limits](https://www.digitalocean.com/docs/kubernetes/#limits)
  - 個人での利用用途してはそんなに気になるところはないかな
- 既知の問題
  - [www.digitalocean.com - docs/kubernetes/#known-issues](https://www.digitalocean.com/docs/kubernetes/#known-issues)
  - 致命的な問題はないと思うが、一度目を通しておいたほうがよいかな
- クラスターの自動アップグレード機能あり　
  - [www.digitalocean.com - docs/kubernetes/how-to/upgrade-cluster](https://www.digitalocean.com/docs/kubernetes/how-to/upgrade-cluster/)


# 触ってみる
実際にDigitalOceanのKubernetesマネージドサービスを触ってみる。

## Kubernetesクラスター起動
以下のような構成でKubernetesクラスターを起動した。

- Region
  - Singapore
  - 日本から一番近いところを選択した
- VPC Network：default-sgp1 DEFAULT
  - これはデフォルト。変更の余地なし。
- Cluster capacity
  - Node pool name
    - test-k8s-node-pool
      - 任意の名前
  - Machine type
    - Basic nodes
  - Node plan
    - 1GB RAM usable(2GB Total)/1vCPU
    - $10/Month per node($0.015/hr)
  - Number Nodes
    - 2
- Tags
  - test-k8s
  - クラスターに任意のタグを複数付けることができる。 
- Name
  - test-k8s-cluster
  - 任意のクラスター名を付けることができる。

月額料金はこれくらい。
MONTHLY RATE $20.00/month $0.03/hour

## kubectlとdoctlをローカル環境にインストール
kubectlは以下参照。
[kubernetes.io - install-kubectl](https://kubernetes.io/ja/docs/tasks/tools/install-kubectl/)

```sh
brew install doctl
```

## Personal access tokensを用意
ダッシュボードのAPIメニューからPersonal access tokensを確認できる。
初回だとREADのみなのでWRITEも付与して、tokenを控えておく。
※初回だとtokenが生成されていないようなので、Regenerate tokenをすることをtokenを払い出す必要があるぽい。あるいは新規トークン作成でも良いと思う。

## クラスターに接続
まずは認証をする。

```sh
doctl auth init
```

クラスター一覧を確認してみる。
```sh
doctl kubernetes cluster list
```

接続するクラスター名を指定してcontextを追加する。（./kube/configが更新される）
```sh
doctl kubernetes cluster kubeconfig save CLUSTER_NAME
```

ノードを確認してみる。
```sh
kubectl get no
```

これでサンプルのアプリをデプロイする準備ができたので、[github.com - digitalocean/doks-example](https://github.com/digitalocean/doks-example)あたり試してみると良さそう。ロードバランサーが作成されて、ロードバランサーの課金が発生するので注意が必要。

# 所感
GKEで最安構成のk8s環境を作るのも魅力的だが、個人利用でマネージド使うならDigital Oceanのほうが良いかも。

ちょっと神経質かもしれないけど、転送量が従量課金なのは気になっちゃうので、Conohaも検討を続ける。

# 参考
- [zenn.dev - DigitalOcean Kubernetesに入門した時のメモ](https://zenn.dev/gosarami/articles/94475cc82d73b5e3f453)
- [www.slideshare.net - 今日から始めるDigitalOcean](https://www.slideshare.net/zembutsu/all-about-degital-ocean-introduction-distribution-hbstudy)
