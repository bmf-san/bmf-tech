---
title: TerraformとAnsibleを使ってVPSを構築する
description: "TerraformとAnsibleでConoHa VPS上にOpenStack IaC構築、サーバープロビジョニングとインフラ構成管理を自動化"
slug: terraform-ansible-vps-setup
date: 2020-12-20T00:00:00Z
lastmod: 2026-03-15
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Ansible
  - VPS
  - OpenStack
  - Terraform
translation_key: terraform-ansible-vps-setup
---


# 概要
この記事は[Makuake Advent Calendar 2020](https://adventar.org/calendars/5986)の20日目の記事です。

TerraformとAnsibleを組み合わせてVPS上でサーバー構築をしてみたのでその手順をまとめておこうと思います。

# 動機
趣味で開発しているアプリケーションのインフラ環境をIaCで整備したかったので、勉強を兼ねてTerraformを使ってみました。

# 環境
- Terraform v0.14.0
- Ansible 2.9.10
- ConoHa VPS
- macOS Catalina 10.15.5

ローカルでTerraformを実行してConoHa VPSにサーバーを建てたり、壊したりします。

# OpenStack
OpenStackはIaaS環境を構築するためのOSSです。

ConoHaのVPSはOpenStackを採用しており、OpenStack準拠のAPIが用意されています。

cf. [www.slideshare.net - "ConoHa" VPS-KVM; OpenStack Grizzly based service](https://www.slideshare.net/chroum/cono-ha-openstackgrizzlyvpsrelease)

TerraformでOpenStackのproviderを利用することでConoHa VPSにサーバーを構築することができます。

cf. [conoha.jp - API](https://www.conoha.jp/vps/function/api/)

今回はTerraformでOpenStackのproviderを使いますが、AnsibleにもOpenStack Ansibleモジュールというのがあるので、同様のことはAnsibleだけでも実現可能だとは思います。試してはいないですが・・

cf. [docs.ansible.com - OpenStack Ansible モジュール](https://docs.ansible.com/ansible/2.9_ja/dev_guide/platforms/openstack_guidelines.html)

# ソースコード
今回作成したソースコードは、[github.com - bmf-san/terraform-ansible-openstack-boilerplate](https://github.com/bmf-san/terraform-ansible-openstack-boilerplate)に置いてあります。

# 実装方針
Terraformでサーバー構築をして、Ansibleでサーバーの初期セットアップをします。

TerraformとAnsibleを両方使う場合は、TerraformからAnsibleを呼ぶのか、AnsibleからTerraformを呼び出すべきなのか迷う気がしますが、下記の記事ではどちらでも良い、正解不正解は特にないという見解でした。

cf. [www.redhat.com - HASHICORP TERRAFORM AND RED HAT ANSIBLE AUTOMATION](https://www.redhat.com/cms/managed-files/pa-terraform-and-ansible-overview-f14774wg-201811-en.pdf)

Terraformはインフラリソースの設定管理、Ansibleはサーバー内の構成管理にそれぞれ強みがあるイメージなので、それぞれが得意な領域を担当できるようには意識しつつ、Terraform内でAnsibleを実行する構成にしてみました。

Terraform、Ansibleそれぞれの役割を意識した上で、コードをどう管理していきたいかという方針によっては逆のパターンが良いという場合もあるのではないかなと思います。

# 準備
- Terraformのインストール
  - brewでtfenvをインストールして最新版を用意しました。
- Ansibleのインストール
  - brewでインストールしました。
- 公開鍵・秘密鍵のキーペアを2セット用意
  - rootユーザー用と作業ユーザー用です。
- ConoHaの管理画面でAPIユーザーを作成
  - 手順は[support.conoha.jp - APIユーザーを追加する](https://support.conoha.jp/v/addapiuser/)を参照してください。
  - 以下の情報が必要となるので事前に確認しておきます。
    - API情報
      - テナントID
      - テナント名
    - APIユーザー
      - ユーザー名
      - パスワード

# 構築手順
大まかな流れは以下の通りです。

ConoHaのAPI利用のためのAPIトークン取得

↓

利用したいイメージ、VMプランを決める

↓

Terraformのコードを書く

↓

Ansibleのコードを書く

## ConoHaのAPI利用のためのAPIトークン取得
まずはConoHaのAPIを利用するためのAPIトークンを取得します。
APIのエンドポイントはユーザーごとに異なるのでConoHaコントロールパネルのAPI情報にあるエンドポイントのリストを適宜参照してください。

cf. [conoha.jp - トークン発行](https://www.conoha.jp/docs/identity-post_tokens.php)

```sh
curl -X POST \
-H "Accept: application/json" \
-d '{"auth":{"passwordCredentials":{"username":"USER_NAME","password":"PASSWORD"},"tenantId":"TENANT_ID"}}' \
https://identity.tyo2.conoha.io/v2.0/tokens \
| jq ".access.token.id"
```

## 利用したいイメージ、VMプランを決める
取得したAPIトークンを使ってそれぞれの情報を取得して、利用したいイメージとVMプランを決めます。

### 利用可能なイメージ一覧を取得
利用可能なイメージ一覧を取得します。

cf. [conoha.jp - イメージ一覧取得](https://www.conoha.jp/docs/image-get_images_list.php)

```sh
curl -X GET \
-H 'Content-Type: application/json' \
-H "Accept: application/json" \
-H "X-Auth-Token: API_TOKEN" \
https://compute.tyo2.conoha.io/v2/TENANT_ID/images \
| jq ".images | sort_by(.name) | map(.name)"
```

今回は`vmi-ubuntu-20.04-amd64-30gb`を使いました。

### VMプラン一覧を取得 
利用可能なVMプラン一覧を取得します。

cf. [conoha.jp - VMプラン一覧を取得](https://www.conoha.jp/docs/compute-get_flavors_list.php)

```sh
curl -X GET \
-H 'Content-Type: application/json' \
-H "Accept: application/json" \
-H "X-Auth-Token: API_TOKEN" \
https://compute.tyo2.conoha.io/v2/TENANT_ID/flavors \
| jq ".flavors | sort_by(.name) | map(.name)"
```

今回は`g-1gb`を選択しました。

1gb以下のプランだとディスクサイズが足りずに構築エラーになるようです。（`g-512mb`で試しましたがダメでした。）

## Terraform
必要な情報が揃ったのでコードを書いていきます。

今回は以下のようなディレクトリ構成にしました。

```
.
├── ansible.cfg
├── main.tf
├── playbooks
├── templates
│   └── playbooks
│       ├── hosts.tpl
│       └── setup.tpl
├── terraform.tfvars
└── variable.tf

3 directories, 12 files
```

今回はやることが少ないのでtfファイルは特に細かく分割していません。

tfstateファイルの管理については、backendを使って外部ストレージで管理するのが良いかと思いますが、今回はローカルからの実行なので`.gitignore`対象に含めるだけになっています。（ローカルとはいえちゃんとやっておきたい部分ではありますが..）

後述しますが、`playbooks`にはterraformが`templates`から生成する`hosts`ファイルと`setup`ファイル（yml）が配置されます。

ゼロからインスタンスを構築するので、構築過程でIPアドレスの値を拾ってTerraformからAnsibleに値を渡してあげる必要があるため、`hosts`ファイルについてはテンプレ化しておく意義があるかなと思うのですが、`setup`ファイル（yml）についてはタスクと変数定義を分けて、変数定義をするファイルをテンプレ化したほうが良いかなと思います。今回は端折って分割していません。

Terraformに寄せすぎると後でAnsibleを切り出したいとなった時などに腰が重くなるような気がするので、この辺りは色んな事例を知りたいところです。

### main.tf
`main.tf`の中身はこんな感じです。

```
terraform {
  required_version = ">= 0.14"
  required_providers {
    openstack = {
      source = "terraform-provider-openstack/openstack"
      version = "1.33.0"
    }
  }
}

provider "openstack" {
  user_name   = (var.user_name)
  password    = (var.password)
  tenant_name = (var.tenant_name)
  auth_url = (var.auth_url)
}

resource "openstack_compute_keypair_v2" "example_keypair" {
  name       = (var.keypair_name)
  public_key = file(var.path_to_public_key_for_root)
}

resource "openstack_compute_instance_v2" "example_instance" {
  name        = (var.instance_name)
  image_name  = (var.image_name)
  flavor_name = (var.flavor_name)
  key_pair    = (var.keypair_name)

  security_groups = [
    "gncs-ipv4-ssh",
    "gncs-ipv4-web",
  ]

  metadata = {
    instance_name_tag = (var.instance_name_tag)
  }
}

data "template_file" "hosts" {
  template = file("./templates/playbooks/hosts.tpl")

  vars = {
    host = (var.host)
    ip = (openstack_compute_instance_v2.example_instance.access_ip_v4)
  }

  depends_on = [openstack_compute_instance_v2.example_instance]
}

resource "local_file" "save_hosts" {
  content  = (data.template_file.hosts.rendered)
  filename = "./playbooks/hosts"

  depends_on = [openstack_compute_instance_v2.example_instance]
}

data "template_file" "setup" {
  template = file("./templates/playbooks/setup.tpl")

  vars = {
    host = (var.host)
    new_user_name = (var.new_user_name)
    new_user_password = (var.new_user_password)
    shell = (var.shell)
    new_user_public_key = file(var.path_to_public_key)
    port = (var.port)
  }

  depends_on = [openstack_compute_instance_v2.example_instance]
}

resource "local_file" "save_setup" {
  content  = (data.template_file.setup.rendered)
  filename = "./playbooks/setup.yml"

  depends_on = [openstack_compute_instance_v2.example_instance]
}

resource "null_resource" "example_provisoner" {
  provisioner "local-exec" {
    command = "ansible-playbook ./playbooks/setup.yml -i ./playbooks/hosts --private-key=${var.path_to_private_key_for_root}"
  }

  depends_on = [openstack_compute_instance_v2.example_instance]
}
```

#### provider
公式でopenstackのproviderがあるのでそれを使っています。

```
provider "openstack" {
  user_name   = (var.user_name)
  password    = (var.password)
  tenant_name = (var.tenant_name)
  auth_url = (var.auth_url)
}
```

`user_name`、`password`はConoHaで作成したAPIユーザーの情報になります。
`tenant_name`は文字通りテナント名です。
`auth_url`はわかりづらいのですが、ここではConoHaのIdentity APIのエンドポイント（ex. `https://identity.tyo2.conoha.io/v2.0`）になります。

#### openstack_compute_keypair_v2
インスタンス構築時にrootユーザーが利用する公開鍵・秘密鍵のキーペアのセットアップです。

cf. [registry.terraform.io - openstack_compute_keypair_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_keypair_v2)

```
resource "openstack_compute_keypair_v2" "example_keypair" {
  name       = (var.keypair_name)
  public_key = file(var.path_to_public_key_for_root)
}
```

公開鍵を指定しない場合は公開鍵・秘密鍵のキーペアが自動で生成され仕組みになっています。

鍵情報はtfstateファイルに出力されるため、実環境で実行する場合はtfstateファイルを適切に管理する必要があります。

公開鍵認証が前提になっていますが、パスワード認証を可能にする方法も無いこともないみたいです。

cf. [noaboutsnote.hatenablog.com - 【Openstack】インスタンスOSにパスワードログインできるようする](http://noaboutsnote.hatenablog.com/entry/openstack_instance_password)

#### openstack_compute_instance_v2
インスタンスのイメージやVMプラン、ネットワーク構成などインスタンスを構築するためのセットアップです。

cf. [registry.terraform.io - openstack_compute_instance_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_instance_v2)

```
resource "openstack_compute_instance_v2" "example_instance" {
  name        = (var.instance_name)
  image_name  = (var.image_name)
  flavor_name = (var.flavor_name)
  key_pair    = (var.keypair_name)

  security_groups = [
    "gncs-ipv4-ssh",
    "gncs-ipv4-web",
  ]

  metadata = {
    instance_name_tag = (var.instance_name_tag)
  }
}
```

`instance_name`は任意の名前、`image_name`は文字通りイメージ名です。
`flavor_name`は初見だと察しが付きづらいですが、ここではVMプラン名になります。

`instance_name_tag`の部分は、ConoHaのコントロールパネルで表示されるネームタグになります。

今回は使用していませんが、user_dataを指定すればcloud-initを使うこともできます。

ex.
```
user_data = data.template_file.user_data.rendered

data "template_file" "user_data" {
  template = file("user_data.sh")
}
```

#### null_resource
null_resourceは他のresourceをトリガとしてプロビジョニングを行うresourceです。トリガはdepends_onで指定します。

構築したインスタンスにAnsibleでプロビジョニングを行いたいので、インスタンスの構築完了（Terraformの実行が完了というのが正確かもしれません。Terraformの実行が終了してもインスタンスの構築が完了しているわけではないので、後述しますがAnsibleでインスタンスの構築を待つ処理を用意しています。）をトリガとしています。

```
resource "null_resource" "example_provisoner" {
  provisioner "local-exec" {
    command = "ansible-playbook ./playbooks/setup.yml -i ./playbooks/hosts --private-key=${var.path_to_private_key_for_root}"
  }

  depends_on = [openstack_compute_instance_v2.example_instance]
}
```

今回はローカルで実行するので`local-exec`を使っています。

ちょうど良い感じのresourceがないかと調べたところ、[github.com - jonmorehouse/terraform-provisioner-ansible](https://github.com/jonmorehouse/terraform-provisioner-ansible)というのがありましたが、現在はメンテナンスされていないようでした。

## Ansible
今回は`templates/playbooks`配下にテンプレートを用意して、Terraform実行時にテンプレを元に実ファイルを生成、生成したファイルを使ってAnsibleを実行する形を取りました。

プロビジョニングの内容は、実行ユーザーの作成、ssh周りの設定調整くらいです。

インスタンスの疎通を待たずしてプロビジョニングしようとしてハマりました...（`wait_for_connection`を使って対応しました。）

# 実行
terraformコマンドオンリーで完結です。

```
terraform init
terraform plan
terraform apply
terraform show
ssh username@ipaddress -i path_to_private_key
terraform destroy
```

# 所感
初Terraformだったので良い勉強になりました。

TerraformはともかくOpenStackは面白い技術だなと思ったのでもう少し深堀りする機会を作りたいです。

# 参考
- [blog.mosuke.tech](https://blog.mosuke.tech/entry/2019/03/26/terraform-integrated-with-ansible/)
- [www.redhat.com - HASHICORP TERRAFORM AND RED HAT ANSIBLE AUTOMATION](https://www.redhat.com/cms/managed-files/pa-terraform-and-ansible-overview-f14774wg-201811-en.pdf)
- [note.com - AnsibleとTerraformと](https://note.com/santak/n/n22dbb35f6c01)
- [qiita.com - TerraformでConoHaのインスタンスを立ててみる](https://qiita.com/kaminchu/items/d0776c381213d54a3a69)
- [qiita.com - TerraformからAnsibleのplaybookを実行する](https://qiita.com/hayaosato/items/ee0d6eabb7b3d0a22136)
- [conoha.jp - ConoHa API Documentation](https://www.conoha.jp/docs/)
- [blog.adachin.me - [OpenStack]TerraformでConoHa VPSのインスタンスを立ち上げてみた！](https://blog.adachin.me/archives/10378)
- [gist.github.com - san-tak/tf-bookmark.md](https://gist.github.com/san-tak/9ef4d15eafb4c8a42af33ffa04464739)
- [github.com - maki0922/terraform_in_conoha](https://github.com/maki0922/terraform_in_conoha)
- [github.com - hobbby-kube](https://github.com/hobby-kube/provisioning)
- [github.com - r0b2g1t/terraform-openstack-ansible](https://github.com/r0b2g1t/terraform-openstack-ansible)
- [github.com - dan4ex/Terraform](https://github.com/dan4ex/Terraform)
- [registry.terraform.io - openstack_compute_keypair_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_keypair_v2)
- [registry.terraform.io - openstack_compute_instance_v2](https://registry.terraform.io/providers/terraform-provider-openstack/openstack/latest/docs/resources/compute_instance_v2)
