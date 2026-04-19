---
title: Setting Up Nginx as a Reverse Proxy
slug: nginx-reverse-proxy-configuration
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - CentOS
  - Nginx
  - Apache
  - Sakura VPS
description: Notes on configuring Nginx as a reverse proxy on a Sakura VPS previously running Apache.
translation_key: nginx-reverse-proxy-configuration
---

I configured Nginx as a reverse proxy on a Sakura VPS that was previously running Apache, so I'm jotting down some notes.

It's been a while since I set it up, so I might not remember everything perfectly. Please bear with me.

# Environment
* Sakura VPS
* CentOS 6 series
* Apache 2.2.15
* Nginx 1.8.1

# Recommended Knowledge
* Understanding and configuring Apache virtual hosts

In simple terms, Nginx receives requests and forwards them to a specified port on Apache. The virtual host settings are configured on the Apache side. Nginx just acts as a pass-through.

# Installing Nginx

Using `wget` and `yum` to download and install was straightforward. I won't cover the installation here, so please install it according to your environment.

Once installed, stop Apache and verify Nginx's operation.

# Changing Apache's Port

We'll use port 80 for Nginx. Specify a different port for Apache. Here, we'll use port 8080.

/etc/httpd/conf/httpd.conf

```
NameVirtualHost *:8080

<VirtualHost *:8080>

hogehogehogehoge...

</VirtualHost>
```

Note
Check iptables with:
`iptables -L`

The location of iptables is:
/etc/sysconfig/iptables

If you have virtual host settings for a custom domain, change those ports as well.

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

# Configuring Reverse Proxy in Nginx

I referred to [Coexisting Apache and Nginx and Gradually Migrating](http://web.archive.org/web/20150917084807/http://concrete5.tomo.ac:80/developer/nginx%E3%81%A7concrete5/apache%E3%81%A8nginx%E3%82%92%E5%85%B1%E5%AD%98%E3%81%97%E3%81%A6%E5%BE%90%E3%80%85%E3%81%AB%E7%A7%BB%E8%A1%8C%E3%81%99%E3%82%8B).

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

Finally, restart Apache and Nginx to complete the setup!

# Thoughts
The performance improvement is noticeable... somewhat faster, perhaps?

I have much to learn about infrastructure setup, so I'll strive to improve in the future.