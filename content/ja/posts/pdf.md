---
title: "パスワード付きPDFの解除をするワンライナー"
slug: "pdf"
date: 2024-03-29
author: bmf-san
categories:
  - "ツール"
tags:
  - "PDF"
  - "Tips"
draft: false
---

1. [QPDF](https://texwiki.texjp.org/?QPDF)をインストールする。

`brew install qpdf`

2. 対象ファイルのパスワードを解除する。

```
qpdf --decrypt input.pdf --password=PASSWORD output.pdf
```

〜完〜
