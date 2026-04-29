---
title: "パスキーのすべて: Adoption, UX Design, and Implementation"
description: "パスキーのすべて: Adoption, UX Design, and Implementation"
slug: passkeys-adoption-ux-implementation
date: 2026-04-29T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Book Review
  - Authentication
  - Passkeys
  - WebAuthn
  - FIDO
translation_key: passkeys-adoption-ux-implementation
books:
  - asin: "4297146533"
    title: "パスキーのすべて ―導入・UX設計・実装"
---


[I read パスキーのすべて ―導入・UX設計・実装](https://www.amazon.co.jp/dp/4297146533).

A book on passkeys by Eiji Kitamura (agektmr), Masaru Kurabayashi, and Kosuke Koiwai.

As the title suggests, it covers three layers in a single volume: "adoption (why / what)," "UX design," and "implementation." Its scope ranges from the context that brought passkeys onto the stage (the limits of passwords, phishing resistance), through the underlying specs like WebAuthn and FIDO, server-side / Web / iOS / Android implementations, login UX design, and operational pitfalls (device loss and account recovery, combination with security keys, etc.).

What stood out:

- It frames passkeys as a combination of "public-key cryptography + cross-device sync + login UX," so scattered knowledge from picking up WebAuthn in isolation lines up along a single axis.
- The UX chapters are substantial. Topics like Conditional UI (autofill prompts) are treated concretely, so it doesn't end at "just place a button."
- It walks through the bindings that are easy to miss when you only touch the Web side, such as `apple-app-site-association` and `assetlinks.json` for native apps.
- Toward the end, it is honest that passkeys do not solve everything on their own. Recovery flows, combining with security keys, and future extensions like the Digital Credentials API are kept in view.

This is not an entry-level book. It reads more as a reference book for someone who already has the basics of web authentication, public-key cryptography, and OAuth/OIDC, and is now adopting passkeys in a product.

Reading [Introduction to Digital Identity](https://amzn.to/4b2GqOg) first to get a map of the authentication/authorization landscape, and then going deep on passkeys with this book, meshes well.
