---
title: "Cryptography Fundamentals: Symmetric, Asymmetric, One-Way Functions, Hashing, and Digital Signatures"
slug: cryptography-fundamentals
date: 2026-06-25
author: bmf-san
categories:
  - Application
tags:
  - Encryption
  - Public Key Cryptography
  - Digital Signature
  - Hashing
  - RSA
  - Security
description: "From symmetric versus asymmetric encryption to the one-way and trapdoor functions behind RSA and ECC, plus hash functions and digital signatures, this article explains the building blocks of cryptography with primary sources. Part 1 of a three-part series that groups public-key uses into signing, encryption, and key exchange."
translation_key: cryptography-fundamentals
draft: false
---

# Introduction
HTTPS, SSH, and JWT all rest on cryptography for their security.

The topic looks complex, yet it reduces to a small set of building blocks.

This article opens the series and explains the foundations:

- Symmetric and asymmetric encryption
- One-way and trapdoor functions (the mathematical core of `RSA` and elliptic-curve cryptography)
- Hash functions
- Digital signatures

Across the series, we group the uses of a public key into three: signing, encryption, and key exchange.

- Part 1 (this article): the building blocks
- Part 2: [Key Exchange and PKI](https://bmf-tech.com/posts/cryptography-key-exchange-and-pki)
- Part 3: [Cryptography in Practice: TLS, JWT, and SSH](https://bmf-tech.com/posts/cryptography-in-practice)
- Practical: [Three Uses of a Public Key](https://bmf-tech.com/posts/public-key-three-uses)

To see how encryption and hashing compare with other data transformations, read [Comparing Data Transformation Methods](https://bmf-tech.com/posts/data-transformation-methods). Here we dig into how they work.

To let you verify the details, each section cites a primary source at the end.

# Symmetric and Asymmetric Encryption
Encryption splits into two families based on how it uses keys.

## Symmetric Encryption
Symmetric encryption uses the same key to encrypt and decrypt. `AES` is the leading example.

```
plaintext --[encrypt with key K]--> ciphertext --[decrypt with the same key K]--> plaintext
```

It runs fast and suits bulk data.

Its challenge is key distribution: how do two parties share the same key safely?

## Asymmetric Encryption
Asymmetric encryption uses a matched pair: a public key and a private key. `RSA` and elliptic-curve cryptography lead here.

You hand the public key to anyone, and you keep the private key to yourself.

The core property is asymmetry: whatever one key processes, only the other key reverses.

```
encrypt with the public key --> only the private key decrypts (confidentiality)
sign with the private key    --> the public key verifies (authenticity)
```

It runs slower than symmetric encryption, so it does not suit bulk data.

## Combining Them (Hybrid)
Real-world communication combines both.

A public-key step authenticates the peer and shares a symmetric key safely; from there, symmetric encryption carries the payload fast.

| Aspect | Symmetric | Asymmetric |
| --- | --- | --- |
| Keys | one (shared) | two (public and private) |
| Speed | fast | slow |
| Main challenge | key distribution | compute cost and key authenticity |
| Examples | `AES` | `RSA`, elliptic-curve cryptography |

# One-Way and Trapdoor Functions
The security of asymmetric encryption rests on a mathematical property: easy to compute, hard to reverse.

## One-Way Functions
A one-way function computes easily in the forward direction but resists reversal within any practical time.

Multiplying two primes takes an instant, yet recovering those primes from the product (factoring) grows hard as the numbers get large.

## Trapdoor Functions
A trapdoor function is a one-way function that you reverse only if you hold a secret (the trapdoor).

That secret is the private key. Asymmetric encryption builds on trapdoor functions.

## RSA: Hardness of Factoring
`RSA` roots its security in the difficulty of factoring a large composite number.

```
n = p * q        (p, q are large primes)
public key: (n, e)   private key: d
encrypt/verify: c = m^e mod n
decrypt/sign:   m = c^d mod n
```

`n` becomes public, yet without recovering `p` and `q` from it, no one derives the private key `d`.

## Elliptic-Curve Cryptography: Hardness of Discrete Logarithms
Elliptic-curve cryptography (`ECC`) roots its security in the hardness of the elliptic-curve discrete logarithm problem.

```
Q = k * G   (G is the base point, k is the secret scalar)
```

Even with `G` and `Q`, recovering `k` stays hard.

`ECC` reaches the same security with a shorter key than `RSA`, so its keys stay small.

## What "Hard" Means
Here, "hard" means that no known algorithm and computer solves it within practical time.

Mathematics has not proven it strictly impossible. Once quantum computers mature, they could break `RSA` and `ECC`, which is why post-quantum cryptography moves toward standardization.

# Hash Functions
A hash function maps an input of any length to a fixed-length output (a hash value) in one direction.

```
"hello" --> 2cf24dba5fb0a30e...   (SHA-256, 256 bits)
"hellp" --> 7c8e8b58a3b2...        (one character differs, the value changes completely)
```

Its key properties:

- Deterministic: the same input always yields the same output
- One-way: you cannot recover the input from the hash
- Collision-resistant: finding two inputs with the same hash stays hard
- Avalanche effect: a one-bit change in the input produces a completely different output

`SHA-256` (`SHA-2` family) and `SHA-3` lead here. `MD5` and `SHA-1` have known collisions, so you should not use them for signatures.

Hashes serve integrity checks, the pre-step of digital signatures, and password storage (paired with a salt and stretching).

A hash uses no key and never reverses, so it differs from encryption. [Comparing Data Transformation Methods](https://bmf-tech.com/posts/data-transformation-methods) lays out that distinction.

# Digital Signatures
A digital signature proves who produced the data (authenticity) and that no one altered it (integrity).

## Sign the Hash, Not the Data
A signature covers the hash of the data, not the data itself.

Signing a fixed-length hash keeps large data efficient to sign and strengthens security.

```
sign:
  hash = H(message)
  signature = Sign(private key, hash)

verify:
  Verify(public key, message, signature) -> true / false
```

The private key produces the signature, and the public key verifies it.

Only the holder of the private key produces a signature, while anyone with the public key verifies it.

## Contrast with Encryption
Signing and encryption run the keys in opposite directions.

- Encryption: the public key encrypts, the private key decrypts (confidentiality)
- Signing: the private key signs, the public key verifies (authenticity)

Common schemes include `RSA` signatures (such as `RSASSA-PSS`), `ECDSA`, and `EdDSA`.

# The Three Uses a Public Key Supports
With these blocks in place, the uses of a public key (a key pair) fall into three:

- Signing: the private key signs, the public key verifies. It serves authentication, authenticity, and tamper detection
- Encryption: the public key encrypts, the private key decrypts. It serves confidentiality (the mirror of signing)
- Key exchange: it shares a symmetric key safely. Part 2 covers this in depth

Almost every cryptographic application combines these three. This lens pays off in Part 3, on applications.

# Summary
This article laid out the building blocks of cryptography.

| Block | Role | Examples |
| --- | --- | --- |
| Symmetric encryption | fast encryption | `AES` |
| Asymmetric encryption | authentication and key sharing | `RSA`, `ECC` |
| One-way / trapdoor functions | mathematical base of asymmetric encryption | factoring, discrete logarithms |
| Hash functions | integrity and signature pre-step | `SHA-256`, `SHA-3` |
| Digital signatures | authenticity and integrity | `ECDSA`, `EdDSA` |

Part 2 turns to key exchange, which shares a symmetric key safely, and PKI, which vouches for the correctness of a public key.

# References
- Rivest, Shamir, Adleman, "A Method for Obtaining Digital Signatures and Public-Key Cryptosystems", Communications of the ACM 21(2), 1978. https://doi.org/10.1145/359340.359342
- RFC 8017: PKCS #1: RSA Cryptography Specifications Version 2.2. https://www.rfc-editor.org/rfc/rfc8017
- NIST SP 800-186: Recommendations for Discrete Logarithm-Based Cryptography: Elliptic Curve Domain Parameters. https://csrc.nist.gov/pubs/sp/800/186/final
- FIPS 180-4: Secure Hash Standard (SHS). https://csrc.nist.gov/pubs/fips/180-4/upd1/final
- FIPS 202: SHA-3 Standard: Permutation-Based Hash and Extendable-Output Functions. https://csrc.nist.gov/pubs/fips/202/final
- FIPS 197: Advanced Encryption Standard (AES). https://csrc.nist.gov/pubs/fips/197/final
- FIPS 186-5: Digital Signature Standard (DSS). https://csrc.nist.gov/pubs/fips/186/5/final
- RFC 8032: Edwards-Curve Digital Signature Algorithm (EdDSA). https://www.rfc-editor.org/rfc/rfc8032

