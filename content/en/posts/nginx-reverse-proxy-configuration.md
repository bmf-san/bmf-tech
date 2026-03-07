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
translation_key: nginx-reverse-proxy-configuration
---

I set up Nginx as a reverse proxy on my Sakura VPS that was previously running Apache, so here are my notes.

Since this was set up quite a while ago, I might not remember everything perfectly, so please bear with me.

# Environment
* Sakura VPS
* CentOS 6 series
* Apache 2.2.15
* Nginx 1.8.1

# Prerequisite Knowledge
* Understanding of Apache's virtual host mechanism and configuration

In general terms, Nginx accepts requests and forwards them to the specified port of Apache. The virtual host settings will be configured on the Apache side. Nginx simply acts as a conduit.

# Installing Nginx

Using wget and yum to download it was straightforward. I will skip the details here, so please install it according to your own setup.

Once the installation is complete, let's stop Apache and check if Nginx is working properly.

# Changing Apache's Port

We will use port 80 for Nginx. On the Apache side, we will specify a different port. Here, we will use port 8080.

/etc/httpd/conf/httpd.conf

```
NameVirtualHost *:8080

<VirtualHost *:8080>

hogehogehogehoge...

</VirtualHost>
```

Note:
To check iptables, use:
`iptables -L`

The location of iptables is:
/etc/sysconfig/iptables

If you have set up virtual hosts for your custom domain, make sure to change the port there as well.

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

# Configuring Reverse Proxy on Nginx

I referred to [Coexisting Apache and Nginx for Gradual Migration](http://concrete5.tomo.ac/developer/nginx%E3%81%A7concrete5/apache%E3%81%A8nginx%E3%82%92%E5%85%B1%E5%AD%98%E3%81%97%E3%81%A6%E5%BE%90%E3%80%85%E3%81%AB%E7%A7%BB%E8%A1%8C%E3%81%99%E3%82%8B) for this.

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
I wonder how the performance will be... it seems to have improved a bit, maybe?

There are many areas where I lack knowledge in infrastructure setup, so I plan to work on that in the future.