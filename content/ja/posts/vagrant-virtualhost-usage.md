---
title: "VagrantでVirtualHostを使う"
slug: "vagrant-virtualhost-usage"
date: 2017-10-01
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "Vagrant"
draft: false
---

VagrantBoxを作りっぱなしでほったらかしな仮想環境を維持してきたのですが、今更ながらVirtualHostの設定をしました。

# やること
* vagrant-hostupdaterのインストールとVagrantfileの設定
* Apacheのconfファイルの設定（VirtualHostの設定）


# やらないこと
* vagrantのインストール
* symbolic linkの設定とか面倒なこと

# 環境
* Vagrant
* vagrant-hostsupdater・・・今回インストールするvagrantのpluginです。
* CentOS6系
* Apache
* Hosts・・・インストールしておくとMACの環境設定からGUIでHosts確認できます。またHostsの設定もできます。今回はvagrantのpluginでHostsを設定するのでGUIから設定はしません。参考：[Macのhosts編集はHostsが良さげ【Mac】](http://blog.sou-lab.com/mac-hosts/)



# vagrant-hostupdaterのインストールとVagrantfileの編集

vagrant-hostupdaterのインストール（sudoまたはroot権限が必要かもです）
`vagrant plugin install vagrant-hostsupdater`


Vagrantfileの編集
ファイルの末尾に以下のような記述を追加。
```
# vagrant-hostupdater
 config.vm.network :private_network, ip: "192.168.33.10"
 config.vm.hostname = "localdev"
 config.hostsupdater.aliases = ["dev", "hoge"]
```

hostnameとaliasは自由に設定してください。aliasは後でapacheのconfファイルで利用します。ドメイン名といったところでしょうか。

設定が終わったらvagrantを起動するなり再起動するなりしてください。


ターミナルで/private/etc/hostsを見るか、Hostsをインストールしている方はHostsを見ると自動でHostが設定されているのを確認できると思います。


# VirtualHostの設定
VPSで手馴れている人もいるかと思います。同じ手順です。


/etc/httpd/confのhttpd.confの`#NameVirtualHost *:80`の#を外してコメントアウトしてください。

viまたはvimでhttpd.confを開いて/NameVとか打ってenter押してnを数回たたいたとこにあるやつです。


そのままhttpd.confにバーチャルホストの設定をかいてもいいのですが、/etc/httpd/conf.dに設定をまとめることにします。

aliasesにdevとhogeを設定したのでそれぞれ設定ファイルを作成します。



/etc/httpd/conf.d/dev.conf

```
<VirtualHost *:80>
        ServerName dev
        ServerAdmin localhost
        DocumentRoot /var/www/html/dev

        <Directory "/var/www/html/dev">
                AllowOverride All
        </Directory>
</VirtualHost>
```

/etc/httpd/conf.d/hoge.conf
```
<VirtualHost *:80>
        ServerName hoge
        ServerAdmin localhost
        DocumentRoot /var/www/html/hoge

        <Directory "/var/www/html/hoge">
                AllowOverride All
        </Directory>
</VirtualHost>
```


ログファイルの出力先は省略します。（省略するとデフォルトの出力先になります。）

/var/www/htmlは共有フォルダとかシンボリックリンクとかVagrantの厄介なやつです。各自の環境に合わせてください。

ブラウザでdev/とhoge/それぞれ表示できればOKです。


# 所感
もっと早くやっておけよと自分を責めました(´・ω・`)


# 参考
* [Vagrantでバーチャルホストを設定してみたｗ
](http://raichel.hatenablog.com/entry/2015/06/06/205958)
* [Vagrantを利用したVirtualHostの設定](http://sk51.jp/how-to-setup-virtualhost-using-vagrant/)

