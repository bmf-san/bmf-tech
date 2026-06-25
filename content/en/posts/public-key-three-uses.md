---
title: "Three Uses of a Public Key: Signing, Encryption, and Key Exchange"
slug: public-key-three-uses
date: 2026-06-25
author: bmf-san
categories:
  - Application
tags:
  - Public Key Cryptography
  - Digital Signature
  - Authentication
  - Security
  - JWT
  - WebAuthn
  - OAuth
  - PKI
description: "A public key has just three uses: signing, encryption, and key exchange. This article sorts real-world applications such as SSH public-key auth, WebAuthn, mTLS, private_key_jwt, code signing, and container signing into these three, and clarifies how they differ from shared-key schemes such as HMAC, with primary sources."
translation_key: public-key-three-uses
draft: false
---

# Introduction
Modern authentication and security, including HTTPS, SSH, JWT, and passkeys, rest on public-key cryptography.

These technologies look scattered, yet through one lens, the use of a public key (a key pair), they reduce to three:

- Signing: the private key signs, the public key verifies
- Encryption: the public key encrypts, the private key decrypts
- Key exchange: two sides share a symmetric key safely

Almost every application combines these three. This article sorts real technologies into the three.

For how public-key cryptography itself works, see the companion series.

- [Cryptography Fundamentals](https://bmf-tech.com/posts/cryptography-fundamentals) (symmetric/asymmetric keys, trapdoor functions, hashing, signatures)
- [Key Exchange and PKI](https://bmf-tech.com/posts/cryptography-key-exchange-and-pki)
- [Cryptography in Practice](https://bmf-tech.com/posts/cryptography-in-practice) (TLS, JWT, SSH)

This article serves as the practical follow-up, surveying applications by use.

# The Three Uses of a Public Key
Here is the overview.

| Use | Operation | What it protects |
| --- | --- | --- |
| Signing | the private key signs, the public key verifies | authenticity, authentication, tamper detection |
| Encryption | the public key encrypts, the private key decrypts | confidentiality |
| Key exchange | exchange public values to derive a symmetric key | safe sharing of a symmetric key |

Signing and encryption run the keys in opposite directions, as the fundamentals article explained.

Next, we walk through applications use by use.

# Signing: Authenticity and Authentication
Signing has the widest reach.

Only the holder of the private key signs, and anyone with the public key verifies. This asymmetry serves authentication, authenticity, and tamper detection.

The table sorts common applications by signer (private-key side) and verifier (public-key side).

| Pattern | Signer (private key) | Verifier (public key) | Main use |
| --- | --- | --- | --- |
| SSH public-key auth | client | server (`authorized_keys`) | remote login |
| WebAuthn / passkeys | user device | service (RP) | passwordless auth |
| mTLS | both client and server | the peer | M2M mutual auth |
| private_key_jwt | client | authorization server | OAuth client auth |
| service-account JWT | workload | IdP / token endpoint | M2M token retrieval |
| SAML / OIDC IdP signature | IdP | SP / RP | federation |
| code signing | publisher | OS / user | tamper detection |
| container signing (cosign) | publisher | deploy side | supply-chain protection |
| Git commit/tag signing | committer | reviewer / GitHub | authenticity |
| crypto wallet | owner | the network | transaction signing |

Each follows the same shape: the private key signs, the public key verifies.

For example, SSH public-key auth has the client sign with its private key, and the server verify with the public key in `authorized_keys`.

OAuth's private_key_jwt works the same way. The client signs a JWT with its private key, and the authorization server verifies it with a registered public key. This beats sending a client_secret.

# Encryption: Confidentiality
The public key encrypts, and the private key decrypts. This mirrors signing.

Because anyone uses the public key, you craft a message that only the holder of the private key reads.

Common applications:

- PGP / GPG: encrypting mail and files
- S/MIME: encrypting and signing mail
- SOPS + age: encrypting secrets as files for IaC

Communication rarely uses encryption directly, though. Public-key operations run slow, so a key exchange derives a symmetric key, and symmetric encryption carries the payload. Encryption shines for stored data or asynchronous messages, where the parties cannot exchange keys in real time.

# Key Exchange: Sharing a Symmetric Key Safely
Rather than send the symmetric key, both sides exchange public values and derive the same symmetric key locally.

Diffie-Hellman (DH) and its elliptic-curve form ECDHE handle this. The [Key Exchange and PKI](https://bmf-tech.com/posts/cryptography-key-exchange-and-pki) article covers the mechanism.

Common applications:

- TLS / HTTPS: server authentication (signing) plus ECDHE (key exchange)
- Signal and other end-to-end encryption: messaging E2EE
- WireGuard / SSH: encrypting the channel

Communication uses key exchange rather than encryption for forward secrecy. A throwaway key per connection keeps past traffic safe even after a later key leak.

# The Foundation of Trust: PKI
All three uses assume the peer's public key is genuine.

PKI guarantees that assumption. A Certificate Authority (CA) issues a certificate that signs the statement "this public key belongs to this subject."

PKI itself applies signing. The [Key Exchange and PKI](https://bmf-tech.com/posts/cryptography-key-exchange-and-pki) article covers it.

# A Caution: What Is Not Public-Key Cryptography
Finally, note the shared-key schemes that people confuse with public-key cryptography.

The following use a shared secret key between sender and receiver, not a key pair:

- HMAC (`HS256` JWTs, AWS SigV4, webhook signatures)
- client_secret client authentication

People call these "signatures" too, but whoever holds the key forges the same value. That breaks the asymmetry of a public-key signature.

A JWT's `alg` tells them apart:

- `HS*` (such as `HS256`): HMAC, a shared key
- `RS*`, `ES*`, `PS*`, `EdDSA`: a public-key signature

# Summary
A public key (a key pair) has, at its core, three uses:

- Signing: authentication, authenticity, tamper detection
- Encryption: confidentiality
- Key exchange: safe sharing of a symmetric key

From SSH and WebAuthn to mTLS, JWT, code signing, and crypto wallets, applications read as combinations of these three. When you meet a new technology, ask what the private key does and what the public key verifies, and the structure appears.

# References
- W3C Web Authentication (WebAuthn) Level 2. https://www.w3.org/TR/webauthn-2/
- RFC 8705: OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens. https://www.rfc-editor.org/rfc/rfc8705
- RFC 7523: JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants. https://www.rfc-editor.org/rfc/rfc7523
- OpenID Connect Core 1.0. https://openid.net/specs/openid-connect-core-1_0.html
- RFC 4880: OpenPGP Message Format. https://www.rfc-editor.org/rfc/rfc4880
- RFC 8551: S/MIME Version 4.0 Message Specification. https://www.rfc-editor.org/rfc/rfc8551
- RFC 2104: HMAC: Keyed-Hashing for Message Authentication. https://www.rfc-editor.org/rfc/rfc2104
- RFC 7518: JSON Web Algorithms (JWA). https://www.rfc-editor.org/rfc/rfc7518

