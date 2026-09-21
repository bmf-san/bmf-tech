---
title: パスワード付きPDFの解除をするワンライナー
description: "コマンドラインツールのQPDFを使って、パスワードで保護されたPDFの制限を安全に解除する方法を紹介する記事です。brewでQPDFをインストールし、qpdf --decryptコマンドで対象のPDFからパスワードによる制限を外す、シンプルなワンライナーの手順を解説します。"
slug: unlock-password-protected-pdf
date: 2024-03-29T00:00:00Z
author: bmf-san
categories:
  - ツール
tags:
  - PDF
  - Tips
translation_key: unlock-password-protected-pdf
draft: false
---


1. [QPDF](https://texwiki.texjp.org/?QPDF)をインストールする。

`brew install qpdf`

2. 対象ファイルのパスワードを解除する。

```
qpdf --decrypt input.pdf --password=PASSWORD output.pdf
```

〜完〜
