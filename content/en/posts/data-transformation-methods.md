---
title: "Comparing Data Transformation Methods: Encoding, Serialization, Encryption, Hashing, Compression, Compilation, and Parsing"
slug: data-transformation-methods
date: 2026-06-23T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Data Transformation
  - Encoding
  - Serialization
  - Encryption
  - Hashing
  - Compression
description: "A concise comparison of common data transformation methods — encoding, serialization, encryption, hashing, compression, compilation, and parsing — across three axes: reversibility, whether a key is required, and purpose."
translation_key: data-transformation-methods
draft: false
---


# Introduction
"Data transformation" covers a lot of ground, and the goals vary widely.

The terms look alike, which invites confusion, but three axes make the differences clear.

- Can it be reversed (reversibility)?
- Does reversing it require a key?
- What is the transformation for (purpose)?

This post covers seven of them: encoding, serialization, encryption, hashing, compression, compilation, and parsing.

# Encoding
Encoding converts data into another format or representation according to a rule.

```
Hello → SGVsbG8= (Base64)
あ    → %E3%81%82 (URL encoding)
```

Common examples include Base64, URL encoding (percent-encoding), and character encodings such as UTF-8.

No key is involved, so anyone who knows the rule can reverse it (decode it).

Its main purpose is to fit data to another rule or environment and to make it safe to carry over a transport. A typical case is using Base64 to send binary data over a path that only accepts text.

Note that it is not a mechanism for secrecy. Base64 is not "encryption" — anyone can decode it.

# Serialization
Serialization converts an in-memory object or data structure into a form that can be stored or transmitted (a byte sequence or string). The reverse is called deserialization.

```json
{ "id": 1, "name": "bmf" }
```

Common examples include JSON, XML, YAML, Protocol Buffers, and MessagePack.

No key is required; anyone who knows the format can reverse it.

Its main purpose is to save or transfer the state of an object — for inter-process communication, exchanging data across a network, or persisting to a file.

It targets structured data and objects, and in a broad sense it can be seen as a kind of encoding.

# Encryption
Encryption uses a key to convert data into a form (ciphertext) that third parties cannot read. The reverse is called decryption.

It is broadly divided into symmetric-key encryption (such as AES) and public-key encryption (such as RSA).

With the key you can reverse it, but without the key you cannot reverse it in any practical amount of time. This "cannot be reversed without a key" property is the decisive difference from encoding.

Its main purpose is to prevent eavesdropping and ensure confidentiality — that is, security.

# Hashing
Hashing converts data of arbitrary length into a fixed-length value (a hash value).

```
password → 5e884898da... (SHA-256)
```

Its key trait is being one-way: you cannot recover the original data from the hash value (it is irreversible). The same input always produces the same output, and a tiny change in the input changes the output dramatically.

Common examples are SHA-256 and SHA-3. For storing passwords, the standard practice is to use bcrypt or Argon2, which add salting and stretching.

Its main purpose is to protect passwords, detect tampering, and verify data integrity (checksums and fingerprints).

It is easily confused with encryption, but the decisive difference is that it cannot be reversed. Collisions, where different inputs produce the same hash, exist in theory, but cryptographic hash functions are designed to make them impractical to produce.

# Compression
Compression removes redundancy from data to make it smaller.

There is lossless compression, which can be restored exactly (ZIP, gzip, PNG, and so on), and lossy compression, which cannot be fully restored (JPEG, MP3, H.264, and so on).

Lossless compression restores the original exactly, while lossy compression discards information that is hard for humans to notice, so it does not return to the original precisely.

Its main purpose is to reduce size and transfer volume — that is, to make things lighter.

# Compilation
Compilation converts source code into another form (machine code, bytecode, or another language).

It is essentially irreversible: optimization and loss of information make it hard to recover the exact original source. Decompilation is possible, but it does not restore the original down to variable names and comments.

Its main purpose is to run the program — converting code that humans read and write into a form a computer can execute.

Converting between languages at the same level, such as TypeScript to JavaScript, is sometimes distinguished as transpilation.

# Parsing
Parsing analyzes a string or byte sequence according to a grammar and turns it into structured data (such as a syntax tree).

Common examples are parsing JSON, parsing HTML, SQL parsers, and the syntax-analysis phase of a compiler.

Unlike the other methods, its main purpose is not the "transformation" itself but the analysis, extraction, and interpretation of data. In the sense of reading input and reconstructing it into structure, it has something in common with the reverse of serialization.

# Comparison table
| Method | Reversible? | Key required | Main purpose |
| --- | --- | --- | --- |
| Encoding | Yes (anyone) | No | Fit to another rule/environment, transport |
| Serialization | Yes (anyone) | No | Save/transfer object state |
| Encryption | Yes, with a key | Yes | Prevent eavesdropping, confidentiality (security) |
| Hashing | No (irreversible) | No | Password protection, tamper detection |
| Compression | Yes (if lossless) | No | Reduce size/transfer volume |
| Compilation | Essentially no | No | Run the program |
| Parsing | Different goal (analysis) | No | Analyze/extract data |

# Points that are easily confused
- Encoding ≠ encryption: Base64 can be reversed by anyone, so it offers no secrecy. Use encryption when you want to hide something.
- Hashing ≠ encryption: a hash cannot be reversed. It differs in both purpose and nature from encryption, which can be reversed with a key.
- Serialization ⊂ encoding: serialization can be seen as a kind of encoding aimed at structured data and objects.
- Compression ≠ encoding: compression aims at size reduction, and whether it is reversible depends on the method (broadly a kind of encoding, but with a different purpose).

# Conclusion
There are many data transformation methods with similar-sounding names, but lining them up along three axes — can it be reversed, does it need a key, and what is it for — makes them easy to sort out.

In a security context especially, it is easy to mistakenly use encoding in place of encryption or hashing. Choose by purpose: encryption when you want secrecy, and hashing when being irreversible is the value, as in password protection and tamper detection.

# References
- [Converting Between Binary, Decimal, and Hexadecimal](https://bmf-tech.com/posts/binary-decimal-hexadecimal-conversion)
- [Wikipedia - Base64](https://en.wikipedia.org/wiki/Base64)
- [Wikipedia - Serialization](https://en.wikipedia.org/wiki/Serialization)
- [Wikipedia - Encryption](https://en.wikipedia.org/wiki/Encryption)
- [Wikipedia - Hash function](https://en.wikipedia.org/wiki/Hash_function)
- [Wikipedia - Data compression](https://en.wikipedia.org/wiki/Data_compression)
- [Wikipedia - Compiler](https://en.wikipedia.org/wiki/Compiler)
- [Wikipedia - Parsing](https://en.wikipedia.org/wiki/Parsing)
