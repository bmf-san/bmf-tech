---
title: SPAを構築したときにハマったこと
description: "SPAで History API を使う際のnginx設定、try_filesでindex.htmlへルーティングしながらjsファイルのパス管理を実装する方法"
slug: spa-development-challenges
date: 2018-06-06T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - Nginx
  - SPA
translation_key: spa-development-challenges
draft: false
---


# 概要
以前、[LaravelにSPAを組み込む時に考えたディレクトリ構成とnginxのconfファイル](/ja/posts/laravel-spa-directory-structure-nginx-config/)というタイトルの記事を書いたが、そこで記載したnginxのconfが不十分だったため、改めて問題点を整理、解決した。

# 前提
- History API
- nginx

# SPAを構築したときにハマったこと

## nginxの設定
リロードしても常にindex.htmlを返すように設定する必要がある。
こんな感じでtry_filesを使ってconfを設定する。

```
location / {
        try_files $uri $uri/ /index.html;
}
```

## jsファイルなどのソースのパス
index.htmlでjsファイルのパスを

```javascript
<script type="text/javascript" src="./dist/bundle.js"></script>
```

と指定していため、`/dashboard/post`などにアクセスすると
 `/dashboard/post/dist/bundle.js`とリソースを返すようになってしまっていた。

URIに関係なく常にbundle.jsを参照できるように絶対パスを指定するようにした。

```javascript
<script type="text/javascript" src="/dist/bundle.js"></script>
```

# 所感
割と解決に時間がかかったが、nginx側なのか、アプリーケーション側なのか問題を切り分けて考えてみるとすぐに理解できた。

# 参考
- [react-router + 静的ファイル(css, js) の組み合わせの罠](https://qiita.com/rooooomania/items/c50acf84d56793de6318)


