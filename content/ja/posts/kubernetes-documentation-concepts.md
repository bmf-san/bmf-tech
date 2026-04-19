---
title: Kubernetesドキュメントのリーディングーコンセプトのまとめ
description: "理解するKubernetesの概念。コンテナ化、クラスタ管理、自動スケーリング、ストレージオーケストレーション、マイクロサービスアーキテクチャを実践的に解説"
slug: kubernetes-documentation-concepts
date: 2020-10-20T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Docker
  - コンテナ
  - Kubernetes
translation_key: kubernetes-documentation-concepts
---


# 概要
Kubernetesを本格的にキャッチアップしていくためにドキュメントを読んだので、オレオレメモを残す。
全部は長いのでメモ書きはコンセプトの章だけにする。

[kubernetes.io](https://kubernetes.io/ja/docs/home/)

# Kubernetesとは何か？
cf. ~~Kubernetesとは何か？~~

## Kubernetesとは何か？
- 宣言的な構成管理
- 自動化の促進
- コンテナ化されたワークロードやサービスの管理のためのプラットフォーム

## 過去を振り返ってみると 
- 仮想化ができる前の時代におけるデプロイ (Traditional deployment)
  - 物理サーバー上でのアプリケーションのリソース制限がない
    - リソース割当問題がある
  - スケールしづらい
  - 維持費用がかかる
- 仮想化を使ったデプロイ (Virtualized deployment)
  - アプリケーションをVMごとに隔離できる
    - アプリケーション間でのデータアクセス制限
  - 仮想化による物理サーバー内のリソース使用率向上
  - アプリケーションの追加や更新が容易
    - ハードウェアコスト削減、スケーラビリティ向上 
- コンテナを使ったデプロイ (Container deployment)
  - アプリケーション間でOSを共有できる
    - 軽量
  - コンテナはファイルシステム、CPU、メモリー、プロセス空間等をそれぞれ持っている
  - クラウドやOSディストリビューションへ依存しない
  - コンテナのメリット
    - VMイメージよりもコンテナイメージは作成が容易、効率的
    - コンテナイメージのビルド・デプロイを継続的にしやすい
    - 開発と運用の関心の分離
      - アプリケーションコンテナイメージの作成ははビルド・リリース時に行う
    - 可観測性が高い
      - OSレベルの情報とメトリクスに加え、アプリケーションの稼働状態やその他の警告も
    - 環境の一貫性
      - 開発、テスト、本番同じように実行可能
    - クラウドとOSディストリビューションの可搬性
      - オンプレミスでもパブリッククラウドでもどんな環境でも実行可能
    - アプリケーションを中心とした管理
      - 仮想マシン上でOSを実行する形から、論理リソースを使用したOS上でのアプリケーションを実行する形に。
    - マイクロサービスとの親和性が高い
      - 疎結合、分散化、拡張性、柔軟性を持つマイクロサービスと親和性がある
    - リソースの分割
      - アプリケーションのパフォーマンスが予測可能
    - リソースの効率的利用と集約性

## Kubernetesが必要な理由と提供する機能
- サービスディスカバリーと負荷分散
  - DNS名またはIPアドレスでコンテナ公開できる
  - ネットワークトラフィックを分散させることができる
- ストレージオーケストレーション
  - マウントするストレージを自由に選択可能
- 自動化されたロールアウトとロールバック
  - デプロイするコンテナの状態を定義することができる 
- 自動ビンパッキング
  - コンテナが必要とするCPUやメモリ（RAM）を宣言することができる
  - ノードに合わせて調整することが可能。リソースを効率的に利用できる
- 自己修復
  - 起動に失敗したコンテナのし起動、入れ替え、強制終了ができる
- 機密情報と構成管理
  - コンテナイメージの再作成を必要とせず、アプリケーションの構成情報を更新することができる

## Kubernetesにないもの
- Kubernetesは...
  - サポートするアプリケーションの種類を制限しない
  - ソースコードのデプロイやアプリケーションのビルドはしない
  - アプリケーションレベル（ミドルウェア、データベース、キャッシュなど）の機能を組み込んで提供しない
  - ロギング、モニタリング、アラートといった機能を指定しない
  - 構成言語の提供も指定しない
  - マシン構成、メンテナンス、管理、自己修復を行うシステムは提供も採用もしない
  - オーケストレーションを前提としない

# Kubernetesのコンポーネント

cf. [Kubernetesのコンポーネント](https://kubernetes.io/ja/docs/concepts/overview/components/)
- Kubernetesをデプロイするとクラスターが展開される
  - クラスターはコンテナ化されたアプリケーションを実行する、ノードと呼ばれる集合
    - すべてのクラスターには少なくとも1つのワーカーノードがある
  - ワーカーノードはアプリケーションのコンポーネントであるPodをホストする
  - マスターノードはクラスター内のワーカーノードとPodを管理する
    - 複数のマスターノードを使用することでクラスターにフェイルオーバーと高可用性を提供することができる
  - コントロールプレーンはクラスター内のワーカーノードとPodを管理する
    - 本番環境では複数のノードを使用し、耐障害性や高可用性を提供することができる
  - [Kubernetesクラスターの図](/assets/images/posts/kubernetes-documentation-concepts/components-of-kubernetes.png)

## コントロールプレーンコンポーネント
- クラスターに関する全体的な決定（スケジューリングなど）を行う

### kube-apiservier
- Kubernetes APIを外部提供するコンポーネント
- 水平方向にスケールするように設計されている

### etcd
- 一貫性、高可用性を持ったキーバリューストア
- Kubernetesの全てのクラスター情報の保存場所
- Kubernetesのデータストアとして利用する場合は必ずバックアッププランを作成

### kube-scheduler
- 新規に作成されたPodにノードが割り当てられているか監視、割り当てられていなかった場合はそのPodを実行するノードを選択

### kube-controller-manager
- 複数のコントローラープロセスを実行する
- 単一プロセスとして動作する
  - 論理的には各コントローラーは個別のプロセスが1つの実行ファイルにまとめてコンパイルされる
- 以下のコントローラーを含む
  - ノードコントローラー
    - ノードがダウンした場合の通知と対応を行う
  - レプリケーションコントローラー
    -  全レプリケーションコントローラーオブジェクトについて、Podの数を正しく保つ
  - エンドポイントコントローラー
    - ServiceとPodを紐付ける
  - サービスアカウントとトークンコントローラー
    - 新規の名前空間に対してデフォルトアカウントとAPIアクセストークンを作成

### cloud-controller-manager
- 基盤であるクラウドプロバイダーと対話するコントローラーを実行
- 以下のコントローラーはクラウドプロバイダーへの依存関係がある
  - ノードコントローラー
    - ノードの応答停止後、クラウドで削除されたかどうかを判断するためにクラウドプロバイダーを確認する
  - ルーティングコントローラー
    - 基盤であるクラウドインフラでルーティングを設定
  - サービスコントローラー
    - クラウドプロバイダーのロードバランサーの作成・更新・削除
  - ボリュームコントローラー
    - ボリュームを作成・アタッチ・マウントしたり、クラウドプロバイダーとボリュームの調整を行う

## ノードコンポーネント
- 全てのノードにおいて稼働中のPodの管理やKuberetesの実行環境を提供

### kubelet
- クラスター内の各ノードで実行されるエージェント
- 各コンテナがPodで実行されていることを保証する

### kube-proxy
- クラスター内の各nodeで動作しているネットワークプロキシで、KubernetesのServiceコンセプトの一部を実装している

### コンテナランタイム
- コンテナ実行を担当するソフトウェア
- ex. Docker, containerd, CRI-O etc...

## アドオン
- クラスター機能を実装するためのKubernetesリソース（DaemonSet、Deploymentなど）を使用
- クラスターレベルの機能を提供している
  - アドオンのリソースで名前空間が必要なもんおはkube-system名前空間に属す
- いくつかのアドオン
  - DNS
  - Web UI
  - コンテナリソース監視
  - クラスターレベルログ

# Kubernetes API

cf. [Kubernetes API](https://kubernetes.io/ja/docs/concepts/overview/kubernetes-api/)

- [API Reference](https://kubernetes.io/docs/reference/)を参照

# Kubernetesのオブジェクトについて

cf. [Kubernetesのオブジェクトについて](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/)

## オブジェクトのspec（仕様）とstatus（状態）
- Kubernetesオブジェクトはオブジェクトの設定を管理する2つの入れ子になったオブジェクトのフィールドを持っている
  - spec
    - 望ましい状態としてオブジェクトに持たせたい特徴を記述する
  - status
    - オブジェクトの現在の状態を示す　

# Kubernetesオブジェクト管理

cf. [Kubernetesオブジェクト管理](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/object-management/)

## 管理手法
- 命令型コマンド
    - 現行のオブジェクトが対象
    - 開発用プロジェクトの環境を推奨
- 命令形オブジェクト設定
    - 個々のファイルが対象
    - 本番用プロジェクトの環境を推奨
- 宣言型オブジェクト設定
    - ファイルのディレクトリが対象
    - 本番用プロジェクトの環境を推奨

## 命令形コマンド
- ユーザーはクラスター内の現行のオブジェクトに対して処理は行う

## 命令形オブジェクト設定
- kubectlコマンドに処理内容、任意のフラグ、1つ以上のファイル名を指定

## 宣言型オブジェクト設定
- ユーザーはローカルに置かれている設定ファイルを操作する
- 操作内容はファイルに記載しない

# オブジェクトの名前とID

cf. [オブジェクトの名前とID](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/names/)

# Namespace（名前空間）

cf. [Namespace（名前空間](https://kubernetes.io/ja/docs/concepts/overview/working-with-objects/namespaces/)

- 同一の物理クラスター上で複数の仮想クラスターの動作をサポートする
  - この仮想クラスターをNamespaceと呼ぶ

# ノード

cf. [ノード](https://kubernetes.io/ja/docs/concepts/architecture/nodes/)

- ワーカーマシン
- 1つのノードはクラスターの性質にもよるが、1つのVMまたは物理的なマシン
- 各ノードにはPodを動かすために必要なサービスが含まれており、マスターコンポーネントによって管理されている

# Podの外観

cf. ~~Pooの外観~~

- Kubernetesのオブジェクトモデルにおいてデプロイ可能な最小単位のオブジェクト

## Podについて理解する
- PodはKubernetesアプリケーションの基本的な実行単位
- アプリケーションのコンテナ、ストレージリソース、ユニークなネットワークIP、コンテナの実行方法を管理するオプションをカプセル化

# ReplicaSet

cf. [ReplicaSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/replicaset/)

- いつでも安定したレプリカPodのセットを維持することを目的としている

## ReplicaSetを使うとき
- いつでも指定されたあ数のPodのレプリカが稼働することを保証する

# Deployment

cf. [Deployment](https://kubernetes.io/ja/docs/concepts/workloads/controllers/deployment/)

- PodとReplicaSetの宣言的なアップデート機能を提供

# StatefulSet

cf. [StatefulSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/statefulset/)

- ステートフルなアプリケーションを管理するためのワークロードAPI
- DeploymentとPodのセットのスケーリングを管理し、Podの順序と一意性を保証する

# DaemonSet

cf. [DaemonSet](https://kubernetes.io/ja/docs/concepts/workloads/controllers/statefulset/)

- 全て（またはいくつか）のノードが単一のPodのコピーを稼働させることを保証する

# Job

cf. [Job](https://kubernetes.io/ja/docs/concepts/workloads/controllers/statefulset/)

- Jobは1つ以上のPodを作成し、指定された数のPodが正常に終了することを保証する
- JobはPodの正常終了を追跡する

# Service

cf. [Service](https://kubernetes.io/ja/docs/concepts/services-networking/service/)

- Podの集合で実行されているアプリケーションをネットワークサービスとして公開する抽象的な方法

## Serviceを利用する動機
- Podは停止が想定されて設計されており、Podが作成され、停止した際、再作成はされない
  - Deployementをアプリケーション稼働のために利用すると、Podwo動的に作成・削除してくれる
- 各Podはそれ自信のIPアドレスを持つ

## Serviceリソース
- KuberunetesにおいてServiceはPodの論理的なセットやそのPodのセットにアクセスするためのポリシーを定義する

## Serviceの公開（Serviceのタイプ）
- ClusterIP
  - クラスター内部のIPでServiceを公開
- NodePort
  - 各NodeのIPにて静的なポート上でServiceを公開
- LoadBalancer
  - クラウドプロバイダーのロードバランサーを使用してえ、Serviceを外部に公開
- ExternalName
  - CNAMEレコードを返すことでexeternalNameフィールドに指定したコンテンツとServiceを紐付ける

# 設定

cf. [設定](https://kubernetes.io/ja/docs/concepts/configuration/)

## 設定のベストプラクティス
- 一般的な設定のTips
  - 最新の安定したAPIバージョンを使用する
  - 設定ファイルはバージョン管理システムに保存されるべき
  - JSONではなくYAMLを使う。互換性はあまり変わらないが、YAMLのほうがユーザーフレンドリー。
  - 意味が有る場合は関連オブジェクトを単一ファイルにグループ化する
  - kubectlコマンドの多くがディレクトリに対しても呼び出せることを覚えておく
  - 不必要にデフォルト値を指定しない。シンプルかつ最小限なほうがエラーが発生しにくい。
  - オブジェクトの説明にアノテーションを入れる

## ConfigMap
- 機密性のないデータをキーと値のペアで保存するために使用されるAPIオブジェクト。
  - ConfigMapは機密性や暗号化を提供しない。機密情報を扱う場合はSecretを使用するか、追加のサードパーティツールを使用。
- Podは環境変数、コマンドライン引数、またはボリューム内の設定ファイルとしてConfigMapを使用できる

## Secrets
- パスワード、OAuthトークン、SSHキーのような機密情報を保存し、管理できるようにする
- Podの定義やイメージに含めることができる

# セキュリティ

cf. https://kubernetes.io/ja/docs/concepts/security/

## クラウドネイティブセキュリティの概要

### クラウドネイティブセキュリティの4C
- セキュリティは改装で考えることができる。
- クラウドネイティブの4C
  - クラウド
  - クラスター
  - コンテナ
  - コード

### インフラのセキュリティ
- Kubernetesインフラに関する懸念事項
  - API Server（コントロールプレーン）へのネットワークアクセス
  - Nodeへのネットワークアクセス
  - KubernetesからのクラウドプロバイダーAPIへのアクセス
  - etcdへのアクセス
  - etcdの暗号化

### クラスター内のコンポーネント（アプリケーション）
- ワークロードセキュリティに関する懸念事項
  - RBAC認可（Kubernetes APIへのアクセス）
- 認証
- アプリケーションのSecret管理（およびetcdへの保存時に暗号化）
- PodSecurityPolicy
- Quality of Service（およびクラスターリソース管理）
- NetworkPolicy
- Kubernetes IngressのTLS

### コンテナ
- コンテナに関する懸念事項
  - コンテナの脆弱性スキャンとOS依存のセキュリティ
  - イメージの署名と実施
  - 特権ユーザーを許可しない

### コード
- コードに関する懸念事項
  - TLS経由のアクセスのみ
  - 通信ポートの範囲制限
  - サードパーティに依存するセキュリティ
  - 静的コード解析
  - 動的プロービング攻撃

# 所感
メモはかなり端折っている。ドキュメント読みきるのにそこそこ時間を使った。。。

# 参考
Kubernetesドキュメントを読み進めていく中で外部資料にも目を通したので参考になったものをメモ。

- [slideshare.net - Kubernetesのしくみ やさしく学ぶ 内部構造とアーキテクチャー](https://www.slideshare.net/ToruMakabe/kubernetes-120907020)
- [qiita.com - Kubernetes道場 Advent Calendar 2018](https://qiita.com/advent-calendar/2018/k8s-dojo)
- [qiita.com - kubernetes初心者のための入門ハンズオン](https://qiita.com/mihirat/items/ebb0833d50c882398b0f)
- [qiita.com - 数時間で完全理解！わりとゴツいKubernetesハンズオン！！](https://qiita.com/Kta-M/items/ce475c0063d3d3f36d5d)
- [www.netone.co.jp - Kubernetesネットワーク入門](https://www.netone.co.jp/knowledge-center/netone-blog/20191226-1/) 
- [www.slideshare.net - ”30分”ぐらいでわかる「Kubernetes」について](https://www.slideshare.net/YuyaOhara/30kubernetes-81054893) 
- [cloud.google.com - ネットワークの概要](https://cloud.google.com/kubernetes-engine/docs/concepts/network-overview?hl=ja)
