---
title: "Cryptography in Practice: TLS, JWT, and SSH"
slug: cryptography-in-practice
date: 2026-06-25
author: bmf-san
categories:
  - Application
tags:
  - TLS
  - JWT
  - SSH
  - Authentication
  - Security
description: "How everyday protocols such as TLS, JWT/JWS, and SSH combine symmetric keys, public keys, signatures, key exchange, and PKI, viewed through the three uses of a public key and backed by RFCs. The final part of a three-part cryptography series."
translation_key: cryptography-in-practice
draft: false
---

# Introduction
This is Part 3, the final part of the cryptography series.

Part 1 covered the building blocks (symmetric and asymmetric keys, one-way functions, hashing, and signatures), and Part 2 covered key exchange and PKI.

This final part shows how they come together in real protocols: TLS, JWT, and SSH.

The key lens is the three uses of a public key from Part 1: signing, encryption, and key exchange.

- Part 1: [Cryptography Fundamentals](https://bmf-tech.com/posts/cryptography-fundamentals)
- Part 2: [Key Exchange and PKI](https://bmf-tech.com/posts/cryptography-key-exchange-and-pki)

# TLS (the Foundation of HTTPS)
TLS encrypts traffic for HTTPS and other protocols.

Its mechanism combines the building blocks we have seen.

## Three Parts of the Handshake
The TLS 1.3 handshake combines these:

- Key exchange (ECDHE): establishes a symmetric key
- Signing plus PKI: the server presents a certificate and signs the handshake; the client verifies the certificate chain
- Symmetric key plus hashing: the established key encrypts the rest of the traffic and guards its integrity

```
1. Establish a symmetric key through key exchange (ECDHE)
2. Authenticate the server with its certificate (a CA signature) plus a handshake signature
3. Encrypt the rest fast with the symmetric key
```

This is the hybrid from Part 1: a public key authenticates and shares a key, and a symmetric key carries the bulk fast.

## Forward Secrecy and TLS 1.3
Through TLS 1.2, one option encrypted the symmetric key with a public key and sent it (RSA key transport).

That option lacks forward secrecy. TLS 1.3 dropped it and settled on key exchange such as ECDHE.

# JWT / JWS
A JWT is a token format common in API access tokens.

## Structure
A JWT expresses claims (statements such as the issuer and the expiry) as JSON and signs them with JWS.

```
header.payload.signature   (each part uses base64url)
```

The signature covers a hash rather than the payload itself, as Part 1 explained.

## Read the alg
The header's `alg` names the signing method, and one distinction matters here.

- `HS256` and similar (`HS*`): HMAC. This uses a symmetric key, not public-key cryptography
- `RS256`, `ES256`, `PS256`, `EdDSA`: a public-key signature (the signing use)

Note that `HS*` shares one secret key between sender and receiver. Whoever holds that key forges a signature too.

## The alg:none Trap
JWS also allows an "unsecured" form with `alg` set to `none`.

If a verifier accepts it by mistake, anyone forges a token. A verifier should pin the allowed `alg`.

# SSH
SSH secures remote login and similar tasks.

Here too, key exchange and signing combine.

## The Transport Layer
On connection, SSH establishes a symmetric key through key exchange (DH).

At the same time, it signs with the server's host key to authenticate the server. RFC 4253 states that the protocol combines the key exchange with a host-key signature for server authentication.

From there, a symmetric key encrypts the traffic, and HMAC guards its integrity.

## Trust On First Use
SSH often skips PKI.

On the first connection, the client records the server's host key in `known_hosts`, then matches against it on later connections.

This "trust the first connection, then pin it" approach is TOFU (Trust On First Use), a trust model apart from PKI.

## Public-Key Authentication
For user authentication, the client signs with its own private key.

The server verifies the signature with the public key in `authorized_keys`. This is SSH public-key authentication, a prime example of the signing use from Part 1.

# A Map of the Three Uses Across Applications
Organizing these applications by the three uses of a public key gives this:

| | Signing | Key exchange | Symmetric key |
| --- | --- | --- | --- |
| TLS 1.3 | certificate and handshake signature | ECDHE | AEAD over the payload |
| JWT/JWS | RS256, ES256, etc. (HS256 is HMAC, a symmetric key) | ECDH-ES, etc. | HMAC, AES |
| SSH | host-key and user-key signatures | DH | encryption plus HMAC |

One caution: HMAC schemes (such as `HS256` or AWS SigV4) and a client_secret use a symmetric key, not public-key cryptography.

# Summary
Across three parts, we moved from the basics of cryptography to its applications.

Every real protocol combines the parts from Part 1 and Part 2.

- TLS: signing (certificate) plus key exchange (ECDHE) plus a symmetric key
- JWT/JWS: signing (though `HS*` uses a symmetric key)
- SSH: signing (host and user keys) plus key exchange plus a symmetric key

Once you view a public key through signing, encryption, and key exchange, even complex applications fall into place.

# References
- RFC 8446: The Transport Layer Security (TLS) Protocol Version 1.3. https://www.rfc-editor.org/rfc/rfc8446
- RFC 7519: JSON Web Token (JWT). https://www.rfc-editor.org/rfc/rfc7519
- RFC 7515: JSON Web Signature (JWS). https://www.rfc-editor.org/rfc/rfc7515
- RFC 4251: The Secure Shell (SSH) Protocol Architecture. https://www.rfc-editor.org/rfc/rfc4251
- RFC 4252: The Secure Shell (SSH) Authentication Protocol. https://www.rfc-editor.org/rfc/rfc4252
- RFC 4253: The Secure Shell (SSH) Transport Layer Protocol. https://www.rfc-editor.org/rfc/rfc4253

