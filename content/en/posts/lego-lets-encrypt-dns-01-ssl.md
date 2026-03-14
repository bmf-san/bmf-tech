---
title: Obtaining SSL Certificates from Let's Encrypt Using DNS-01 with Lego
description: "Obtain Let's Encrypt SSL certificates efficiently using DNS-01 challenge method with Lego client tool and Conoha API provider integration."
slug: lego-lets-encrypt-dns-01-ssl
date: 2022-01-18T00:00:00Z
author: bmf-san
categories:
  - Infrastructure
tags:
  - Let's Encrypt
  - Tips
translation_key: lego-lets-encrypt-dns-01-ssl
---

I was trying to obtain SSL certificates from Let's Encrypt using Ansible on a Conoha VPS.

I wanted to acquire the certificate using the [DNS-01](https://datatracker.ietf.org/doc/html/draft-ietf-acme-acme-03#section-7.4) method (registering a one-time token issued by the certificate authority in the domain's TXT record for verification), so I created a script using the Conoha API to register and delete TXT records (cf. [github.com - k2snow/letsencrypt-dns-conoha](https://github.com/k2snow/letsencrypt-dns-conoha)). However, managing the script was cumbersome, so I was looking for a simpler way when I found the Let's Encrypt client tool [go-acme/lego](https://github.com/go-acme/lego) and decided to give it a try.

# go-acme/lego
Lego is a Go-based Let's Encrypt client and ACME library.

Various [DNS Providers](https://go-acme.github.io/lego/dns/) are available beyond Conoha.

You can install it using Docker, a package manager, or Go.

# Obtaining SSL Certificates from Let's Encrypt with Lego
If you're using Docker, the command looks like this (I directly took the Ansible code).

```bash
docker run --rm -e CONOHA_POLLING_INTERVAL=30 -e CONOHA_PROPAGATION_TIMEOUT=3600 -e CONOHA_TTL=3600 -e CONOHA_REGION={{ conoha_region }} -e CONOHA_TENANT_ID={{ conoha_tenant_id }} -e CONOHA_API_USERNAME={{ conoha_username }} -e CONOHA_API_PASSWORD={{ conoha_password }} -v /home/{{ ssh_user_name }}/lego:/lego goacme/lego --path /lego --email {{ email }} --dns conoha --domains *.{{ domain }} --domains {{ domain }} --accept-tos run
```

It seems that Conoha's DNS has some delays in reflecting TXT records, so with the default settings, it throws a timeout error. It's better to set `CONOHA_PROPAGATION_TIMEOUT`, `CONOHA_PROPAGATION_TIMEOUT`, and `CONOHA_TTL` as shown above.

When I was using the script, I often got stuck with the DNS behavior and couldn't get it to work. I wonder why...

# Thoughts
It's incredibly easy. Updating the certificate is also fine with Lego.