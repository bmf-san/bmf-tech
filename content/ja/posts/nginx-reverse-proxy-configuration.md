---
title: Nginxをリバースプロキシとして設定する
description: Nginxをリバースプロキシとして設定するの手順と実践例を詳しく解説します。
slug: nginx-reverse-proxy-configuration
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - CentOS
  - Nginx
  - apache
  - さくらのVPS
translation_key: nginx-reverse-proxy-configuration
---


Apacheで運用していたさくらVPSにNginxをリバースプロキシとして設定したのでメモします。

随分前に設定したのでうろ覚えのところもあるかもしれませんがご了承ください。


# 環境
* さくらVPS
* CentOS６系
* Apache2.2.15
* Nginx1.8.1


# あると良い（いやあったほうが良い）前提知識
* Apacheのバーチャルホストの仕組み及び設定方法


おおまかに仕組みをいうと、
Nginxでリクエストを受け付けてApcheの指定ポートにリクエストを流すという感じです（）
バーチャルホストの設定はApache側で設定しておきます。Nginxが右から左へムーディ勝山するだけです。




# Nginxをインストール

wgetしてyumで落としてくるやり方がカンタンでした。ここでは割愛するので各自の艦橋に合わせてインストールしておいてください。


インストールが完了したら、一度Apacheを停止してNginxの動作確認をしておきましょう。


# Apache側のポートを変更する

Nginxでは80番ポートを使うことにします。
Apache側ではそれ以外のポートを指定しましょう。
ここでは8080番ポートを使用することにします。


/etc/httpd/conf/httpd.conf

```
NameVirtualHost *:8080

<VirtualHost *:8080>

hogehogehogehoge...

</VirtualHost>
```

補足
iptablesの確認は、
`iptables -L`

iptablesの場所は、
/etc/sysconfig/iptables


独自ドメインのバーチャルホストの設定をしている場合は、そっちの方のポートも変更しておきましょう。

Ex. /etc/httpd/conf.d/hoge.com.conf

```
# Domain
<VirtualHost *:8080>
  ServerName hoge.com
  DocumentRoot "/var/www/html/hoge"
  DirectoryIndex index.html index.php
  ErrorLog /var/log/httpd/error_log
  CustomLog /var/log/httpd/access_log combined
  AddDefaultCharset UTF-8
  <Directory "/var/www/html/hoge">
    AllowOverride All
  </Directory>
</VirtualHost>

# Sub Domain
<VirtualHost *:8080>
  ServerName sub-hoge.hoge.com
  DocumentRoot "/var/www/html/sub-hoge"
  DirectoryIndex index.html index.php
  ErrorLog /var/log/httpd/error_log
  CustomLog /var/log/httpd/access_log combined
  AddDefaultCharset UTF-8
  <Directory "/var/www/html/sub-hoge">
    AllowOverride All
  </Directory>
</VirtualHost>
```



# Nginx側でリバースプロキシの設定をする


[ApacheとNginxを共存して徐々に移行する](http://concrete5.tomo.ac/developer/nginx%E3%81%A7concrete5/apache%E3%81%A8nginx%E3%82%92%E5%85%B1%E5%AD%98%E3%81%97%E3%81%A6%E5%BE%90%E3%80%85%E3%81%AB%E7%A7%BB%E8%A1%8C%E3%81%99%E3%82%8B)を参考にさせていただきました。


/etc/nginx/conf.d/reverse_proxy.conf

```
"reverse_proxy.conf" 14L, 392C
server {
        listen 80;

        location / {
                proxy_pass http://127.0.0.1:8080;
                proxy_redirect                         off;
                proxy_set_header Host                  $host;
                proxy_set_header X-Real-IP             $remote_addr;
                proxy_set_header X-Forwarded-Host      $host;
                proxy_set_header X-Forwarded-Server    $host;
                proxy_set_header X-Forwarded-For       $proxy_add_x_forwarded_for;
        }
}

```


最後にApacheとNginxを再起動して完了です！



# 所感
気になるパフォーマンスはいかほどに・・・といった感じですが、多少は早くなったかな？という感じです。。


インフラの構築は不勉強なところが多々あるので今後頑張ってみようと思います。

