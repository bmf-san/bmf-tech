---
title: "NGINXのバッファ系ディレクティブ解説：プロキシパフォーマンスの最適化"
description: 'NGINX のリバースプロキシ用バッファ系ディレクティブを解説。proxy_buffering・proxy_buffer_size などの設定がパフォーマンスに与える影響を学べます。'
slug: nginx-buffer-directives
date: 2024-03-02T00:00:00Z
author: bmf-san
categories:
  - インフラストラクチャ
tags:
  - Nginx
translation_key: nginx-buffer-directives-notes
---


bufferサイズを記録したかったときに調べていたのでメモ。

- [large_client_header_buffers](http://mogile.web.fc2.com/nginx/http/ngx_http_core_module.html#large_client_header_buffers)
    - ngx_http_core_module モジュールに含まれるディレクティブ
    - 構文:	large_client_header_buffers number size;
    - デフォルト: large_client_header_buffers 4 8k;
    - コンテキスト:	http, server
    - クライアントからのリクエストヘッダを読み込むのに使われるバッファのnumberとsizeを指定する
    - numberを増やすとバッファ数が増えて、使用するメモリも増える（たぶん）
    - number×sizeというような形ではなくて、最大のバッファサイズはあくまでsizeに依るはず・・
- [fastcgi_buffers](http://mogile.web.fc2.com/nginx/http/ngx_http_fastcgi_module.html#fastcgi_buffers)
    - ngx_http_fastcgi_module モジュールに含まれるディレクティブ
    - 構文:	fastcgi_buffers number size;
    - デフォルト:	fastcgi_buffers 8 4k|8k;
    - コンテキスト:	http, server, location
    - 接続ごとにFastCGIサーバからの応答を読み込むために使われるバッファのnumber とsizeを設定
    - size ≒ fastcgi_buffer_sizeではない。バッファが使うメモリサイズだと思われる・・
- [fastcgi_buffer_size](http://mogile.web.fc2.com/nginx/http/ngx_http_fastcgi_module.html#fastcgi_buffer_size)
    - ngx_http_fastcgi_module モジュールに含まれるディレクティブ
    - 構文:	fastcgi_buffer_size size;
    - デフォルト:	fastcgi_buffer_size 4k|8k;
    - コンテキスト:	http, server, location
    - FastCGIサーバから応答の最初の部分を読み込むために使われるバッファのsizeを設定
