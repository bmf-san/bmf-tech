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
-   After the TCP handshake, a TLS handshake is performed
    -   Once completed, HTTP requests and responses are exchanged in encrypted communication

# Features of TLS

Functions provided by the TLS protocol

-   Confidentiality
    -   Data cannot be viewed on the network path
-   Integrity
    -   Prevents tampering with communication data
        -   Ensures integrity through Message Authentication Code (MAC)
-   Authenticity
    -   Prevents impersonation
    -   Server certificates issued by certification authorities are used by browsers to determine the legitimacy of the accessed site

# Background of Comprehensive HTTPS Adoption

-   State-level hacking undermines the reliability of the internet
-   Recommendation for HTTPS adoption

    -   Development of new technologies based on encryption
    -   Discontinuation of plaintext HTTP features in browsers
    -   SEO advantages for HTTPS sites by browser vendors
    -   Rise of free server certificate issuance services like Let’s Encrypt

# Differences Between SSL and TLS

-   SSL (Secure Socket Layer)
-   TLS (Transport Layer Security)
-   TLS is the next-generation standard of SSL
    -   Vulnerabilities in SSL3.0 released in 1995 (POODLE) were discovered in 2014.
        -   The only security measure is to disable SSL
    -   TLS1.0 was released in 1999
    -   Only slight differences from SSL
    -   TLS1.1 was released in 2006

# References

-   [Why HTTPS Now? The History and Technical Background of TLS That Engineers Should Know for Internet Reliability](https://employment.en-japan.com/engineerhub/entry/2018/02/14/110000)
-   [What are the Differences Between SSL and TLS](https://ssl.sakura.ad.jp/column/ssl_tls/)
