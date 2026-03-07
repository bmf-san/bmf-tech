---
title: パスワード付きPDFの解除をするワンライナー
slug: unlock-password-protected-pdf
date: 2024-03-29T00:00:00Z
author: bmf-san
categories:
  - ツール
tags:
  - PDF
  - Tips
translation_key: unlock-password-protected-pdf
---


1. [QPDF](https://texwiki.texjp.org/?QPDF)をインストールする。

`brew install qpdf`

2. 対象ファイルのパスワードを解除する。

```
qpdf --decrypt input.pdf --password=PASSWORD output.pdf
```

〜完〜
