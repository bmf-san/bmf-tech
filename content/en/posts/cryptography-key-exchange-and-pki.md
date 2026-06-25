---
title: "Key Exchange and PKI: Diffie-Hellman, Certificates, and Certificate Authorities"
slug: cryptography-key-exchange-and-pki
date: 2026-06-25
author: bmf-san
categories:
  - Application
tags:
  - Public Key Cryptography
  - PKI
  - Certificate
  - TLS
  - Security
description: "How Diffie-Hellman shares a symmetric key safely, how ECDHE adds forward secrecy, and how PKI binds a public key to its owner through certificates, CAs, and a chain of trust. Backed by primary sources such as RFCs. Part 2 of a three-part cryptography series."
translation_key: cryptography-key-exchange-and-pki
draft: false
---

# Introduction
This is Part 2 of the cryptography series.

Part 1 covered the building blocks: symmetric and asymmetric encryption, one-way functions, hashing, and digital signatures.

The series runs as follows.

- Part 1: [Cryptography Fundamentals](https://bmf-tech.com/posts/cryptography-fundamentals)
- Part 2 (this article): Key Exchange and PKI
- Part 3: [Cryptography in Practice: TLS, JWT, and SSH](https://bmf-tech.com/posts/cryptography-in-practice)
- Practical: [Three Uses of a Public Key](https://bmf-tech.com/posts/public-key-three-uses)

If you have not read Part 1, start there for an easier read.

This part explains two of the three uses of a public key: key exchange, which shares a symmetric key safely, and PKI, which vouches for a public key.

- Key exchange: a way to share a symmetric key safely (Diffie-Hellman)
- PKI: a way to prove that a public key belongs to its claimed owner

# What Key Exchange Solves
As Part 1 noted, symmetric encryption faces the key distribution problem.

Over a channel that someone could wiretap, how do two parties share the same symmetric key?

Asymmetric encryption could encrypt and send the key, but key exchange offers a cleaner answer.

Key exchange never sends the symmetric key itself. The two parties exchange public values, and each side computes the same symmetric key on its own.

# Diffie-Hellman Key Exchange
Diffie-Hellman (DH), published in 1976, introduced the first public-key idea.

## How It Works
DH relies on the hardness of the discrete logarithm problem. Like the elliptic curves in Part 1, it exploits one property: exponentiation runs easily, reversal runs hard.

```
public parameters: a prime p and a generator g
Alice: picks secret a, sends A = g^a mod p
Bob:   picks secret b, sends B = g^b mod p

shared key:
  Alice: K = B^a mod p
  Bob:   K = A^b mod p
  --> both reach K = g^(ab) mod p (the same value)
```

An eavesdropper sees `p`, `g`, `A`, and `B`, yet cannot recover the secrets `a` or `b`. So the shared key `K` stays out of reach.

## The Man-in-the-Middle Problem and Authentication
DH alone has a weakness: it never checks that the peer is genuine.

An attacker who sits in the middle runs a separate exchange with each side and intercepts the traffic. That is a man-in-the-middle attack.

So practice pairs key exchange with peer authentication. Authentication draws on digital signatures and the PKI described below.

## ECDHE and Forward Secrecy
Today the mainstream is the elliptic-curve form, `ECDHE`.

The trailing `E` stands for ephemeral. It generates a throwaway secret for every connection.

With a throwaway secret, a later leak of the private key does not decrypt past traffic. This property is forward secrecy.

# PKI (Public Key Infrastructure)
Key exchange and signatures assume the peer's public key is genuine.

But how do you trust that a public key you received truly belongs to the peer? PKI answers this.

## A Certificate Is a CA's Signature
A certificate sits at the heart of PKI.

A certificate is a statement that "this public key belongs to this subject (a domain, for example)", which a Certificate Authority (CA) signs.

```
certificate = Sign( CA private key, { subject name, subject public key, validity, ... } )
```

RFC 5280 puts it this way: a trusted CA asserts the binding by signing each certificate.

So PKI applies the digital signature from Part 1 (the signing use).

## The Chain of Trust
A higher CA signs the public key of each CA in turn. This sequence forms the chain of trust.

```
root CA --signs--> intermediate CA --signs--> server certificate
```

The root CA signs its own certificate, which anchors the chain.

## The Trust Store and the Anchor of Trust
So what makes a root CA trustworthy?

An OS or browser ships with a built-in list of trusted root CAs. That list is the trust store, the anchor of trust.

The verifier walks the certificate chain it received and accepts it once the chain reaches a root CA inside the trust store.

## Revocation
Even within the validity period, a leaked private key may force you to invalidate a certificate.

Revocation handles that. The common mechanisms are the CRL (a revocation list) and OCSP (an online query).

## ACME and Let's Encrypt
Obtaining a certificate once meant tedious manual steps.

The ACME protocol automated it, and Let's Encrypt drove its adoption.

It confirms control of a domain automatically and issues and renews certificates for free.

# Summary
This article covered key exchange and PKI.

| Mechanism | Role | Primary source |
| --- | --- | --- |
| Diffie-Hellman | shares a symmetric key safely | DH 1976 |
| ECDHE | key exchange plus forward secrecy | RFC 7919 |
| Certificates and CAs | vouch for a public key | RFC 5280 |
| ACME | issues certificates automatically | RFC 8555 |

Key exchange is the key-exchange use, and PKI applies the signing use; both combine the building blocks from Part 1.

Part 3 shows how these come together in real protocols: TLS, JWT, and SSH.

# References
- Diffie, Hellman, "New Directions in Cryptography", IEEE Transactions on Information Theory 22(6), 1976. https://doi.org/10.1109/TIT.1976.1055638
- NIST SP 800-56A Rev.3: Recommendation for Pair-Wise Key-Establishment Schemes Using Discrete Logarithm Cryptography. https://csrc.nist.gov/pubs/sp/800/56/a/r3/final
- RFC 7919: Negotiated Finite Field Diffie-Hellman Ephemeral Parameters for TLS. https://www.rfc-editor.org/rfc/rfc7919
- RFC 5280: X.509 Public Key Infrastructure Certificate and CRL Profile. https://www.rfc-editor.org/rfc/rfc5280
- RFC 6960: X.509 PKI Online Certificate Status Protocol (OCSP). https://www.rfc-editor.org/rfc/rfc6960
- RFC 8555: Automatic Certificate Management Environment (ACME). https://www.rfc-editor.org/rfc/rfc8555

