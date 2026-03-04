---
title: "cdコマンドでディレクトリをタイピングするのが面倒くさくなってきたのでシェルスクリプトで楽をする"
slug: "simplify-directory-typing-shell-script"
date: 2017-09-26
author: bmf-san
categories:
  - "アプリケーション"
tags:
  - "bash"
  - "shellscript"
draft: false
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

