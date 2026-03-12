---
title: cdコマンドでディレクトリをタイピングするのが面倒くさくなってきたのでシェルスクリプトで楽をする
description: cdコマンドでディレクトリをタイピングするのが面倒くさくなってきたのでシェルスクリプトで楽をするについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: simplify-directory-typing-shell-script
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - bash
  - shellscript
translation_key: simplify-directory-typing-shell-script
---


# 概要
`cd hogehoge`とかパスやらディレクトリやらタイプするのが面倒になるという怠惰っぷりを発揮してきたので、シェルスクリプトを使って少し楽できるようにしました。

# ソース

```bash
#!/bin/sh

# cd by selecting numbers
function cdSelect() {
        dirs=`ls -a`

        PS3="Select directory > "
        echo 'Directory list:'

        select dir in ${dirs}
        do
                stty erase ^H
                cd ${dir}
                break
        done
}
alias cd-s=cdSelect
```

`cd-s`と打つと、

```
Directory list:

1) .
2) ..
3) hoge_a
4) hoge_b
5) hoge_c

Select directory > 3
```

こんな感じになります。

# 所感
ディレクトリが多い時大変そうですが、cdコマンドのストレスが軽減されました。
vimバージョンもつくろうかと。

