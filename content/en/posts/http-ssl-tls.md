---
title: HTTP and SSL/TLS
slug: http-ssl-tls
date: 2018-04-18T00:00:00Z
author: bmf-san
categories:
  - Network
tags:
  - HTTP
  - HTTPS
  - SSL
  - TLS
translation_key: http-ssl-tls
---

# What is HTTPS

-   HTTP (Hyper Transfer Protocol) over TLS (Transport Layer Security)
-   Perform TLS handshake after TCP handshake
    -   After completion, exchange HTTP requests and responses in encrypted communication

# Features of TLS

Functions provided by the TLS protocol

-   Confidentiality
    -   Data contents cannot be seen on the network path
-   Integrity
    -   Prevention of tampering with communication data
        -   Ensured by Message Authentication Code (MAC)
-   Authenticity
    -   Prevention of impersonation
    -   Browsers use server certificates issued by certificate authorities to determine the legitimacy of the destination

# Background of Full HTTPS Adoption

-   State-level hacking undermines the reliability of the internet
-   Recommendation for HTTPS adoption

    -   Development of new technologies based on encryption
    -   Abolition of plaintext HTTP functionality in browsers
    -   SEO advantages for HTTPS sites recognized by browser vendors
    -   Rise of free server certificate issuance services like Let’s Encrypt

# Differences Between SSL and TLS

-   SSL (Secure Socket Layer)
-   TLS (Transport Layer Security)
-   TLS is the next-generation standard of SSL
    -   Vulnerability of SSL 3.0 released in 1995 (POODLE) was discovered in 2014.
        -   The only security measure is to disable SSL
    -   TLS 1.0 was released in 1999
    -   There are only slight differences from SSL
    -   TLS 1.1 was released in 2006

# References

-   [Why HTTPS Now? The History and Technical Background of TLS That Engineers Should Know for Internet Security](https://employment.en-japan.com/engineerhub/entry/2018/02/14/110000)
-   [What are the Differences Between SSL and TLS?](https://ssl.sakura.ad.jp/column/ssl_tls/)