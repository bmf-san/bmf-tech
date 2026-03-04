---
title: "bmf-techを支える技術"
slug: "bmf-tech"
date: 2022-08-08
author: bmf-san
categories:
  - "アーキテクチャ"
tags:
  - "Docker"
  - "Docker Compose"
  - "VPS"
  - "Golang"
  - "Vue.js"
  - "Prometheus"
  - "Promtail"
  - "Loki"
  - "Grafana"
  - "Nginx"
draft: false
---

# bmf-techを支える技術
このブログ（bmf-tech.com）を支える技術スタックについてかく。

# 旧bmf-techの構成
まずは現行のbmf-techはよりも1世代前の構成について。

<img style="width:757px;!important" alt="old_architecture" src="/assets/images/posts/bmf-tech/183280770-84280c0f-e9ab-4cce-9f2d-0ea775e96ea5.png">

- アプリケーションはLaravelをベースとしたモノリシックな構成
  - APIはPHPで、管理画面はReactでSPAを構築していた
    - 当時触っていた技術を何となく採用しただけ
- Sakura VPSでホスティング
  - 構成管理ツール（Ansibleとか）は確か何も使ってなくて手動でミドルウェアをインストールしたり、構成を整えたりしてた
    - ”温かみのある”構成
- モニタリングツールは用意されていない
  - ログを見たいときはサーバーにsshして直接見に行く
- コンテナは使っていない
  - Vagrantが主流？流行っていた？時代だった
- デプロイはgitのhookで良しなにやる構成にしていた 

前世代運用していたアプリーケーションが初代自作CMSである[Rubel](https://github.com/bmf-san/Rubel)。

これは何年運用したか覚えていないけど、多分3~5年くらいだと思う。

[Rubel](https://github.com/bmf-san/Rubel)を運用する前は、Wordpressでオリジナルのテーマを作ってブログを運用していた。

Wordpress（オリジナルテーマ1）→Wordpress（オリジナルテーマ2）→[Rubel](https://github.com/bmf-san/Rubel)→今に至るといった感じ。

bmf-tech.comのドメイン年齢を確認してみたら2015年11月2日が取得日であった。

いつからブログ運用を始めたか忘れてしまったが、ドメイン年齢に基づくなら7年近く運用していることになる。

# 新bmf-techの構成
現行のbmf-techの構成についてかく。

同じ構成でサンプルコードを[gobel-example](https://github.com/bmf-san/gobel-example)にて公開している。

## 設計
システムをリプレースしたい理由がいくつかあった。

- 新しい技術に触れるきっかけを作りたかった
  - 当時はPHPを主に触っていたのでそれ以外の言語を触るきっかけが欲しかった
- [Rubel](https://github.com/bmf-san/Rubel)をメンテナンスする気力が沸かなかった
  - Laravelのアップデートサイクル早く、頻繁に対応が必要だった
    - フレームワークのアップデートよりできるだけビジネスロジックに集中したい
  - Reactは自分の手に余る
    - Reduxに泣いた
      - FLUXが必要な規模ではなかった
  - APIは切り出してフロントは捨てやすい構成にしたかった
  - ソースコードをもっと自分の管理下に治めたかった
    - フレームワークに乗っかりすぎて、フレームワーク依存に何となくの危機感を抱いた
  - システムの不具合を追いづらい
  - etc...

それらの理由から新しいシステムの設計方針をざっくり考えた。

- 長期的にメンテナンスしやすいシステムを構築する
  - サーバー構成管理
	- IaCをちゃんとやる
		- 冪等性を担保する
		- サーバーを乗り換えるようなことがあっても容易に乗り換えるできる
  - アプリケーションの設計
    - フレームワークへの依存を避けたい
    - フロントエンド以外は標準ライブラリや自作ライブラリへの依存のみ留める
  - モニタリング環境を構築する
    - システムメトリクスやログ等を監視したり、アラートを設定したりできるようにする

## アーキテクチャ構成
設計方針を元に構築したアーキテクチャ構成が以下。

![スクリーンショット 2022-11-22 22 53 31](/assets/images/posts/bmf-tech/203331548-95daeea8-8108-400a-91ae-35f8cddf899a.png)

- サーバーはさくらVPSではなくConoHaVPSを採用した。
  - スペック
    - CPU 2Core
    - メモリ 1GB
    - SSD 100GB
	- イメージタイプ Ubuntu
  - ConoHaはOpenStackに対応しているのでインスタンス構築はTerraformでコード管理しやすい
  - 安価で使いやすいのも良いポイント
- サーバーの構成管理はAnsibleを使っている
- SSLはLet's Encryptを使っている
  - 証明書取得・更新は[go-acme/lego](https://github.com/go-acme/lego)を使っている
    - cf. [legoでLet's encryptのSSL証明書をDNS-01方式で取得する](https://bmf-tech.com/posts/lego%E3%81%A7Let%27s%20encrypt%E3%81%AESSL%E8%A8%BC%E6%98%8E%E6%9B%B8%E3%82%92DNS-01%E6%96%B9%E5%BC%8F%E3%81%A7%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
  - 最初はlegoを使わず自前スクリプトでDNS-01方式の証明書取得を試みていたが、たまに上手く取得できないことがあったので断念した
    - cf. [k2snow/letsencrypt-dns-conoha](https://github.com/k2snow/letsencrypt-dns-conoha)
- 仮想化でDockerを採用
  - Docker Composeを使って複数コンテナを管理している
  - Kubernetesを使ってみようと検証したこともあった
    - [TerraformとAnsibleを使ってKubernetes環境構築](https://bmf-tech.com/posts/Terraform%E3%81%A8Ansible%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6Kubernetes%E7%92%B0%E5%A2%83%E6%A7%8B%E7%AF%89) 
	- 自前Kubernetesクラスタで運用してみたかったが、ベアメタルなロードバランサーを上手く扱えず断念
- Nginxは管理画面（SPA）、API（ヘッドレスCMS。Go）、Client（ユーザー側画面。Go）画面を配信するWebサーバー
  - Nginxのイメージはマルチステージビルドで管理画面のイメージを含むようにカスタマイズしているので、管理画面のビルド済みソースコードが内包されている形になっている
  - APIやClientはGoでビルドしたバイナリを内包するイメージ
- モニタリングはGrafanaをUIとして使っている
  - サーバーの各種メトリクス収集はPrometheusが、コンテナログ収集周りはPromtailとLokiが行っている
    - 収集されたものの可視化はGrafana
  - リリース当初はコンテナログ収集周りはEFKスタックを構築していたが、1GBのサーバースペックではリソースがかつかつだったので、LokiとPromtailを使う構成に変更した
    - 機能的には十分そうでかつシンプルに簡単に使えたので、自分の要件には十分そうに見えた

## デプロイ
デプロイは特に複雑なことはしていない。

<img style="width:820px;!important" alt="deploy" src="/assets/images/posts/bmf-tech/183280768-78484c56-5775-4691-898b-f12b42d573e3.png">

- コンテナ構成を管理するprivateリポジトリ（bmf-tech）を用意している
  - [gobel-example](https://github.com/bmf-san/gobel-example) をテンプレートとしている
    - [gobel-example](https://github.com/bmf-san/gobel-example) はEFKスタックだが、privateリポジトリ（bmf-tech）はpromtail×Lokiを使った構成にしている
  - 本番環境と同じコンテナ構成でローカル環境でもコンテナを動かすことができる
  - アプリケーションのソースコードは一切含まれていない
  - 環境変数はすべて外から注入できるように構成されている 
- デプロイはssh→bmf-techをpull→rsyncで環境変数のファイルアップロード→docker-compose build & upするだけのスクリプトを組んで実行するだけ
  - バージョン管理は特にしていない
  - 瞬間的なダウンタイムが発生するのがデメリットだが、現状のトラフィックから考えて大きな問題はない
    - 可用性には響くけど..
- 最初はdocker contextを使ったデプロイ方法を検証して試していたが、不採用になった（理由は忘れてしまった・・）

## ソースコード管理
コンテナ構成に基づいてソースコード管理がどういう形になっているか示した図が以下。

![NOTE - Source code management](https://github.com/bmf-san/bmf-tech-client/assets/13291041/1fb40523-cfc2-4030-82bd-10e7f38dafff)

- bmf-tech
  - [gobel-example](https://github.com/bmf-san/gobel-example) をテンプレートとしたプライベートリポジトリで、コンテナ構成を管理している
- bmf-tech-ops
  - [gobel-ops-example](https://github.com/bmf-san/gobel-ops-example)をテンプレートとしたプライベートリポジトリで、サーバー構成管理やデプロイスクリプトの管理など構築・運用のオペレーションに関わるコードを管理している
- [bmf-tech-client](https://github.com/bmf-san/bmf-tech-client)
  - [gobel-client-example](https://github.com/bmf-san/gobel-client-example)をテンプレートとしたリポジトリで、フロントエンド側のソースコードを管理している
  - DockerHubにイメージをプッシュして、パブリックイメージとして管理している
- gobel-api
  - [gobel-api](https://github.com/bmf-san/gobel-api)
  - ヘッドレスCMSのソースコードを管理している
  - Dockerhubにイメージをプッシュして、パブリックイメージとして管理している
- gobel-admin-client
  - [gobel-admin-client-example](https://github.com/bmf-san/gobel-admin-client-example)をそのまま使っている。
  - 管理画面のソースコードを管理している
  - Dockerhubにイメージをプッシュして、パブリックイメージとして管理している

## アプリケーション設計
API、Client（ユーザー側画面）、Admin（管理画面）それぞれのアプリケーションについて。

### API
- ヘッドレスCMSとしてAPIを構築
- Goを採用している
  - 言語仕様がシンプル、後方互換性がしっかり担保されている、コンパイルが速い、ビルドすればバイナリで動くという可搬性があるのでコンテナとの相性が良い、型がある、標準ライブラリが充実している etc...
- Clearn Architectureを採用した
  - GoとClean Architectureの相性について賛否両論あると思うが、取り組んでみた
  - Clean Architectureにおける”関心の分離”を守っていくことがアプリケーションを長期的にメンテナンスしやすい形に保つ1つの方法だと考えている
- 標準パッケージ以外への依存は少ない
  - cf. [go.mod](https://github.com/bmf-san/gobel-api/blob/master/app/go.mod)
- APIのプロトコルはREST
  - 今からつくるのだとしたらgRPCでも良かったかもしれない

### Client
- 画面をレスポンスするWebサーバー
- GoでAPIクライアントを実装
- テンプレートファイル（html）はembedでバイナリに含めている
- デザインは自作CSSフレームワークを使っている
  - cf. [sea.css](https://github.com/bmf-san/sea.css)

### Admin
- 管理画面
- 認証はJWTによるトークン認証
- セッション管理はRedis
- Vue.jsを採用した
  - Reactよりシンプルにかける、取り組みやすいという印象がある
    - 最近React触っていないのでわからないが、あちこちでVueを触るようになったので知見獲得も兼ねて採用した
  - Atomic Designにチャレンジしてみた
    - が、上手くできているのかはわからない
  - 開発途中、2系から3系にアップデートした 
  - TypeScriptも導入しようかと思ったが、後手に回っている
    - 管理画面は頻繁に機能追加したりしなそうだったので、ライブラリのアップデート程度のメンテナンスしか手をかけない想定だったりする
    - 数年運用する間にフロントエンドのトレンドも変わったり、次のフレームワークが出るかもしれないと考えたりすると、余り手を入れると大変になりそうだと思ったので、なるべくシンプルにしている
- SPA配信はNginxで対応している

## DB設計
基本的には[Rubel](https://github.com/bmf-san/Rubel)のDB設計をそのまま引き継いでいるが、論理・物理削除を見直して再設計した部分がいくつかある。
後はカラムのデータ型やサイズを見直したりといった感じ。

## 移行作業
データの移行は自前のデータ移行ツールを書いて対応した。

[migrate-rubel-to-gobel](https://github.com/bmf-san/migrate-rubel-to-gobel)

DB設計に大きな差がなかったので移行ツールは2~3日くらいで実装することができた。

サーバー移行については、大したことはしていない。
移行先のサーバー環境を整えて、動作確認するための検証用のドメインを取得し、移行先サーバー環境で各種動作確認を行った。
作っては壊しを繰り返してIaCに問題ないかもチェックした。

リリース作業時はDNSを切り替えるだけで新環境への移行が完了。

旧環境は新環境を1ヶ月くらい運用して問題ないことを確認してから削除、契約終了の手続きをした。

## 監視
監視ダッシュボードやアラートはGrafanaで作成・設定した。

監視ダッシュボードの方はjsonファイル形式でデータを管理し、プロビジョニングできるようになっているが、アラートはGrafanaのUIから設定している。
~（アラートの方はまだプロビジョンニングに対応していないため cf.  [github.com - grafana/issues/36153](https://github.com/grafana/grafana/issues/36153)~ →対応されたのでプロビジョニングできるように対応した。

## SLI・SLO
大してトラフィックがないのに設定するのは虚しい気もするが、トラフィックがどうこうというより、一定の可用性を安定して保つことができるか観測するという意味で設定したい。が、まだ未対応。

## 負荷試験
ちょっとやってみたいと思っているので検討中。

## 作ったものまとめ
新bmf-techをリリースするまでに作ったものをまとめる。

- アプリケーション
  - [gobel-api](https://github.com/bmf-san/gobel-api)
  - [gobel-admin-client-example](https://github.com/bmf-san/gobel-admin-client-example)
  - [gobel-client-example](https://github.com/bmf-san/gobel-client-example)
  - [gobel-example](https://github.com/bmf-san/gobel-example)
  - [gobel-ops-example](https://github.com/bmf-san/gobel-ops-example)
- ライブラリ
  - [goblin](https://github.com/bmf-san/goblin)
    - トライ木ベースのルーター
	- 結構時間使ったので色々ブログに書いている
  - [golem](https://github.com/bmf-san/golem)
	- 簡易JSONロガー。ログレベルを指定できる。
  - [goemon](https://github.com/bmf-san/goemon)
    - Go製dotEnv
    - 作ったけど結局使わなかった
- ツール
  - [migrate-rubel-to-gobel](https://github.com/bmf-san/migrate-rubel-to-gobel)
    - データ移行ツール
- ボイラープレート（検証過程でつくった）
  - [go-clean-architecture-web-application-boilerplate](https://github.com/bmf-san/go-clean-architecture-web-application-boilerplate)
    - GoでClean Architectureやるベースを用意するために作ったリポジトリ
  - [go-production-boilerplate](https://github.com/bmf-san/go-production-boilerplate)
    - docker contextを使ったデプロイを方法を検証したリポジトリ
  - [docker-based-monitoring-stack-boilerplate](https://github.com/bmf-san/docker-based-monitoring-stack-boilerplate)
    - モニタリング環境を作ったときに後で再利用できるように作ったリポジトリ
  - [terraform-ansible-openstack-boilerplate](https://github.com/bmf-san/terraform-ansible-openstack-boilerplate)
    - terraformをansibleを仲良く使うために検証したリポジトリ
  - [setup-kubernetes-cluster-on-vps-boilerplate](https://github.com/bmf-san/setup-kubernetes-cluster-on-vps-boilerplate)
    - VPSでKubernetesクラスタを構築するために検証したリポジトリ。ベアメタルのロードバランサーを扱えず断念。
  - [vue-js-boilerplate](https://github.com/bmf-san/vue-js-boilerplate)
    - 最近のVue.jsはどんな感じかキャッチアップするために作ったリポジトリ

上記色々作る中でブログに書いたり、LTしたり他に色々やったりしていたので、新bmf-techのリリースにはかなりの時間を使ってしまった。

## 今後の展望
偶に開発の手を止めたり、道を外れたり、何回か途中でWordpressとか別の既存システムに変えようかと、迷ったことが幾度があったが、無事運用できる形になって肩の荷が下りた。

やれていないことややりたいことは色々といくつかissueを積んでいるので、趣味程度に片手間に対応していきたい。

ただ作るのではなく、どう作るのか、どう運用するのかということに磨きをかけていきたいと思っているので、そのための投資をこのブログシステムを通じてやっていこうと考えている。

自分が自作ブログを運用している理由は学びの要素が強い。実際結構多くのことを学ぶ機会になったり、今後も更に学ぶことができそうである。

- 自分の学びのため
  - 学びを整理するために記事を書く
  - ブログシステム自体を自作して自分でホスティングして自分で利用したらまた学びを得る機会になるだろうと考えている
- システム構築・運用経験の獲得
  - 自分がすべてを把握をしているシステムを運用していくことが学び得ることがあると考えている
- お小遣いを稼ぐ
	- 広告収入で遊んで暮らしたい、というのは冗談だが多少のお小遣いくらい、せめてサーバー運用費くらいは稼げたらいいかなと思っている
	- おまけ程度にしか考えていないのでマネタイズを頑張るということはしていない
		- ここ1~2年くらいは年間の運用コストの何割か程度は利益が出たりしているので続けていくことに何か価値が出るかもしれないと都合良く考えている

当分は現行のシステムを運用し続けていくことができると考えているので気長にやっていこうと思っている。
