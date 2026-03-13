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
---


# 概要
以前、[LaravelにSPAを組み込む時に考えたディレクトリ構成とnginxのconfファイル](https://bmf-tech.com/posts/Laravel%E3%81%ABSPA%E3%82%92%E7%B5%84%E3%81%BF%E8%BE%BC%E3%82%80%E6%99%82%E3%81%AB%E8%80%83%E3%81%88%E3%81%9F%E3%83%86%E3%82%99%E3%82%A3%E3%83%AC%E3%82%AF%E3%83%88%E3%83%AA%E6%A7%8B%E6%88%90%E3%81%A8nginx%E3%81%AEconf%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB)というタイトルの記事を書いたが、そこで記載したnginxのconfが不十分だったため、改めて問題点を整理、解決した。

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


