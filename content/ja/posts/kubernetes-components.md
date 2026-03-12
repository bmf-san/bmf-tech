---
title: "Kubernetesのコンポーネント解説：Pod・Node・コントロールプレーン"
description: 'Kubernetes の各コンポーネントを解説。Pod・Node・コントロールプレーン（API Server / etcd / Scheduler / Controller Manager）の役割を学べます。'
slug: kubernetes-components
date: 2024-03-27T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Kubernetes
translation_key: kubernetes-components
---


 # Kubernetesの各コンポーネントについて

![スクリーンショット 2024-03-27 22 40 42](https://github.com/bmf-san/bmf-tech-client/assets/13291041/cf2f9712-2bf0-4c5e-9212-c66c387110b9)

## Control Plane Components
### kube-apiserver
クラスタを操作するためのKubernetes APIを提供するサーバー。水平スケールが可能な形に設計されている。

### etcd
全てのクラスタの状態管理を行うための高可用性のキーバリューストア。

### kube-scheduler
ノードが未割り当ての新しく生成されたPodをどのノードにスケジューリングするかを決定する。

### kube-controller-manager
クラスタの状態を制御するためのプロセスで、複数の種類が存在する。

- Node controller
  - ノードがダウンしている場合に通知する
- Job controller
  - Job（1回限りのタスク）の監視と作成し、タスクを完了するためにPodを実行する
- EndpointSlice controller
  - EndpointSlice（ネットワークエンドポイントの集合へのリファレンス）を作成する
- Service Account controller
  - 新しいNamespaceにServiceAccountを作成する

### cloud-controller-manager
クラウドプロバイダー固有の制御ロジックを持つコントローラー。クラウドプロバイダーが提供するサービスをKubernetesに統合するために使用される。

## Node Components
### kubelet
各ノードで実行されるエージェント。kube-apiserverと通信し、Podの実行を管理する。

### kube-proxy
クラスターにおけるネットワークプロキシ。クラスター内部や外部との通信を制御する。

### Container Runtime
コンテナを実行するためのソフトウェア。

# 参考
- [kubernetes.io - Kubernetes Components](https://kubernetes.io/docs/concepts/overview/components/)
- [www.redhat.com - Kubernetes アーキテクチャの概要](https://www.redhat.com/ja/topics/containers/kubernetes-architecture)
- [www.rworks.jp - Kubernetes のアーキテクチャとは？特徴と基本コンポーネントからデータ保護の方法まで詳しく解説](https://www.rworks.jp/cloud/kubernetes-op-support/kubernetes-column/kubernetes-entry/29132/)
- [speakerdeck.com - アーキテクチャから学ぶKubernetesの全体像](https://speakerdeck.com/bells17/akitekutiyakaraxue-hukubernetesnoquan-ti-xiang)
