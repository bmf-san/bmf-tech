---
title: TerraformとAnsibleを使ってKubernetes環境構築
slug: kubernetes-setup-with-terraform-ansible
date: 2021-04-06T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Ansible
  - Kubernetes
  - Terraform
  - kubeadm
translation_key: kubernetes-setup-with-terraform-ansible
---


# 概要
TerraformとAnsibleを使ってKubernetes環境構築に取り組んだ。

自作アプリの運用をKubernetesに乗っけてみたいという気持ちから環境を構築するところから初めてみた。

# サーバー選定
プライベートでの開発なので、せいぜい月2000円前後くらいの予算に留めたいところ。

クラウドかVPSか、マネージドか、ノンマネージドかといったところが大きな観点だが、そのへんはコスト感と運用メリットを考慮しつつ決めれば良いのでそんなに悩まないと思う。後述するが、一番の悩みのタネはロードバランサーだった・・。

今回候補に上がったのは3つ。

## GCP
- GKE
  - そのまま使うとクラウド破産・・とまではいかないと思うが予算超過する可能性大。
  - プリエンプティブVMを活用した方法で安く利用ができるが、運用的にどうだろうかというところ。
    - [ludwig125.hatenablog.com - GKE を格安で使うためにやったこと](https://ludwig125.hatenablog.com/entry/2019/11/30/073458)
    - [sleepless-se.net - GKEで最安値のKubernetesクラスタを作る方法](https://sleepless-se.net/2018/12/11/gke-kubernetes/)
    - [blog.a-know.me - 安価なGKE（k8s）クラスタを作って趣味開発に活用する](https://blog.a-know.me/entry/2018/06/17/220222)

## Digital Ocean
- マネージドKubernetesが使えるVPS。
- masterノードが無料で、workerノードが従量課金。
- ダウンタイム許容するならworker1台で、月10$（$0.01/hour）。　
- ロードバランサーも従量課金。
- プロモリンクから新規登録すると$100くらいもらえるのでそれで色々試すことができる。
- 無料で使えるモニタリングが結構しっかりしていて良い。
- エコシステムが良い
  - [marketplace.digitaloceancom.com](https://marketplace.digitalocean.com/)
  - [community](https://www.digitalocean.com/community)
- Openstack API対応
- Kubernetes使う使わない関係なく使いたいかも

## Conoha VPS
- データ転送量課金なし。
- Openstack API対応
- UIがわわかりやすい。
- DBサーバー、オブジェクトストレージなどもある。

上記以外にマネージドk3sを提供している[civo.com](https://www.civo.com/)という選択肢も考えたが、k8sを触りたかったので検討外とした。

Digital OceanとConohaで迷ったが、従量課金なしの安心の料金体系に心奪われたのでConohaを選定した。

GKEやDigital OceanはKubernetesをさっと構築して勉強するにはちょうどよい環境が整っていると思うので、そういった目的で利用を検討していく判断をした。

# Conoha VPSにKubernetes環境を構築
マネージドKubernetesを利用しない選択をしたので、セルフでKubernetesを構築することにした。

構築のツールとしてはkubeadmを採用。

TerraformとAnsibleを使って、インスタンスの構築から初期セットアップ（ユーザー作成、ssh鍵調整など）、kubeadmを使ったKubernetesの構築までコード化したものがこちら↓

[github.com - bmf-san/setup-kubernetes-cluster-on-vps-boilerplate](https://github.com/bmf-san/setup-kubernetes-cluster-on-vps-boilerplate)

masterノード1台、workerノードは複数台想定になっている。

ConohaはOpenstackをサポートしたAPIを用意しているので、Openstackをサポートしている他サーバー（ex. Digital Ocean）であれば、書き換えも楽なはず。

kubeadmによるKubernetesの構築は、Kubernetesの公式ドキュメントを一読して構築の前提条件を把握しておけばそれほど難しくなかった。

# 解決できなかったこと
ロードバランサーの対応ができなかったので、アプリケーションを公開してKubernetes運用をするまでに至らなかった。

自前Kubernetesクラスタの場合、クラウドが用意しているロードバランサーが使えないのでOSSのものを自分で用意する必要があるのだが、そのセットアップが上手く行かず断念・・

1週間近く睡眠時間を削ったが歯が立たなかった..w

解決できなかった問題はこれ。
https://github.com/kubernetes/ingress-nginx/issues/5401

自作アプリは一旦docker-composeでの運用をする方向に転換して、Kubernetesの運用はもう少し理解を深めてからにしようと思う。。。
