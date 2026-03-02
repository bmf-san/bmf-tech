---
title: "ステートレスとステートフル"
slug: "post-98"
date: 2018-04-17
author: bmf-san
categories:
  - "ネットワーク"
tags:
  - "ステートフル"
  - "ステートレス"
  - "セッション"
draft: false
---

# 前提

-   セッション
    -   HTTPは状態を持たないプロトコル
        -   リクエストが同一のクライアントからの通信かどうか判断しない
    -   接続確立から切断までの一連の通信

# ステートフル

-   サーバーがクライアントのセッション状態を保持している
-   システムが状態やデータを保持している
-   プロトコルの例
    -   FTP, TCP, BGP, OSPF, EIGRP, SMTP, SSH

# ステートレス

-   サーバーがクライアントのセッション情報を保持していない
-   システムが状態やデータを保持していない、入力値から出力を判断
-   プロトコルの例
    -   HTTP, UDP, IP, DNP

# 参考

-   [ステートレス ステートフルとはどういうことか](http://blog.sojiro.me/blog/2014/09/13/stateful-and-stateless/)
-   [ステートフル（Stateful）とステートレス（Stateless）の違い](https://milestone-of-se.nesuke.com/nw-basic/as-nw-engineer/stateful-and-stateless/)
-   [セッション状態](https://msdn.microsoft.com/ja-jp/library/aa720612)

