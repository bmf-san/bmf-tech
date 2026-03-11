---
title: php7をインストールした時にmysqlドライバーがなかった
slug: mysql-driver-php7-installation-issue
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - CentOS
  - MySQL
  - PHP
  - PHP7
  - さくらのVPS
  - yum
translation_key: mysql-driver-php7-installation-issue
---


# php7をインストールした時にmysqlドライバーがなかった（）

php7のインストールについてはこちらを参照にしました。
* [CentOS6／CentOS7にPHP5.6／PHP7をyumでインストール](http://qiita.com/ozawan/items/caf6e7ddec7c6b31f01e)

ちなみに私の環境は・・・
* さくらVPS
* CentOS6.7


# ドライバーがないならインストールすればいいじゃない

```
yum install yum install --enablerepo=remi,remi-php70 php-mysqlnd
```


これで解決できたっぽいです_(:3」∠)_

# 補足：php7のインストールに関して
もしかしたら足りないパッケージもあるかもしれませんが、laravelをphp7でcomposerを使っていく分には不足ない気がします。

```
yum -y install --enablerepo=remi-php70 php php-mbstring php-pear php-fpm php-mcrypt php-devel php-xml
```

参考
* [PHP5.6からPHP7にアップデートしました](https://monochrome-design.jp/126)

php7にしたら気の所為かもしれないけど早くなった気がします。いやあきらかに体感速度が変わったような。。

