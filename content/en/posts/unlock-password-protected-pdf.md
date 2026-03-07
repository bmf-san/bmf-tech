---
title: One-liner to Unlock Password-Protected PDF
slug: unlock-password-protected-pdf
date: 2024-03-29T00:00:00Z
author: bmf-san
categories:
  - Tools
tags:
  - PDF
  - Tips
translation_key: unlock-password-protected-pdf
---

1. Install [QPDF](https://texwiki.texjp.org/?QPDF).

`brew install qpdf`

2. Unlock the password of the target file.

```
qpdf --decrypt input.pdf --password=PASSWORD output.pdf
```

〜The End〜