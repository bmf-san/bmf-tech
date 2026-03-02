---
title: "ネットワーク用語まとめ"
slug: "post-245"
date: 2020-11-02
author: bmf-san
categories:
  - "ネットワーク"
tags:
  - "IP"
  - "ロードバランサー"
draft: false
---

# 概要
ネットワークに関して知識が曖昧なワードをまとめる。

## IPv4
- Internet Protocol version 4
- IP（Internet Protocol）の一種
- OSI参照モデルにおいてはネットワーク層に位置づけられるプロトコル
- 32ビットのアドレス空間を持つ
  - 2の32乗＝42億9496万7296個がIPv4のIPアドレス総数
- インターネットの普及とともに枯渇してきた

cf.
- [wikipedia -IPv4](https://ja.wikipedia.org/wiki/IPv4)
- [jprs.jp - 用語辞典 IPv4](https://jprs.jp/glossary/index.php?ID=0034)

## IPv6
- Internet Protocol version 6
- IP（Internet Protocol）の一種
- OSI参照モデルにおいてはネットワーク層に位置づけられるプロトコル
- 128ビットのアドレス空間を持つ
  - 2の128乗＝約3400澗（1澗は1兆*1兆*1兆）個がIPv6のIPアドレス総数
- IPSec（IPパケットのデータを暗号化する）が標準機能
  - IPv4ではオプション

cf.
- [wikipedia -IPv6](https://ja.wikipedia.org/wiki/IPv6)
- [jprs.jp - 用語辞典 IPv6](https://jprs.jp/glossary/index.php?ID=0035)
- [www.nic.ad.jp - IPv6](https://www.nic.ad.jp/ja/newsletter/No20/sec0700.html)

## NAT
- IPアドレスを別のIPアドレスに変換する技術
- ルーターや無線LANのアクセスポイントで利用されている
- 静的NAT
  - IPアドレスとIPアドレスを1:1で変換
  - 主にグローバルIPとプライベートIPの変換に利用
- 動的NAT
  - IPアドレスとIPアドレスを1:Nで変換
- NAT（NAPT）の使用例
  - プライベートIPアドレスを利用しているネットワークからインターネットに出る時
  - 企業内のプライベートIPアドレスのサーバーを外部公開するとき
  - 企業内でIPアドレスが重複してしまったとき

cf.
- [wikipedia - ネットワークアドレス変換](https://ja.wikipedia.org/wiki/%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9%E5%A4%89%E6%8F%9B)
- [www.infraexpert.com - NAT（Network Address Translation)](https://www.infraexpert.com/study/ip10.html)
- [locked.jp](https://locked.jp/blog/what-is-nat/)
- [milestone-of-se.nesuke.com - 【図解】初心者にも分かるNATの仕組みと種類 ~静的/動的NAT/NAPT(PAT),セキュリティ等メリット/デメリット～](https://milestone-of-se.nesuke.com/nw-basic/nat/nat-summary/)

## NAPT
- IPアドレスを別のIPアドレスに変換する技術
- ルーターや無線LANのアクセスポイントで利用されている
- Linuxにおける実装をIPマスカレードと呼ぶ
- 静的NAPT
  - 1つのIPアドレス+TCP/UDPポート番号を別のIPアドレス+ポート番号に1:1で変換
- 動的NAPT
  - IPアドレスとIPアドレスを1:Nで変換
  - 送信元IPアドレスを1つのグローバルIPにN:1で変換する
    - 送信元ポート番号を未使用のポート万能に変換する

cf.
- [wikipedia - ネットワークアドレス変換](https://ja.wikipedia.org/wiki/%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9%E5%A4%89%E6%8F%9B)
- [locked.jp](https://locked.jp/blog/what-is-nat/)
- [milestone-of-se.nesuke.com - 【図解】初心者にも分かるNATの仕組みと種類 ~静的/動的NAT/NAPT(PAT),セキュリティ等メリット/デメリット～](https://milestone-of-se.nesuke.com/nw-basic/nat/nat-summary/)

## CIDR
- Classless Inter-Domain Routing
- クラスを使用しないIPアドレスの割当と経路選択（ルーティング）を行う技術
- インターネット上のルーターにおけるルーティングテーブルの肥大化低減を低減させるための機構 
- IPアドレスをクラス分けして割当て、ルーティングする方式はスケーラブルではない
  - この問題に対処するために、IPアドレスブロックの割当方式を定義する規格が登場した。(RFC1518、RFC1519など)
- 1ビット単位で可変長のサブネットマスクを用いて、IPアドレスのネットワーク部の大きさを指定することで割り当てるブロックのサイズを変更できる
- CIDR表記
  - 192.168.1.0/24

cf. 
- [wikipedia](https://ja.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
- [www.nic.ad.jp - CIDRとは](https://www.nic.ad.jp/ja/basics/terms/cidr.html)

## サブネットマスク
- IPアドレスのうち、ネットワークアドレスとホストアドレスを識別するための数値
  - IPv4の場合は、32ビット
  - IPv6の場合は、128ビット

cf. 
- [wikipedia - サブネットマスク](https://ja.wikipedia.org/wiki/%E3%82%B5%E3%83%96%E3%83%8D%E3%83%83%E3%83%88%E3%83%9E%E3%82%B9%E3%82%AF)

## DHCP
- Dynamic Host Configuration Protocol
- IPv4ネットワークで使用される通信プロトコル
  - コンピューターがネットワークに接続するために必要な設定情報を自動的に割り当てる
  - IPv4でもIPv6でも使用されるが、プロコトルの詳細が異なるため別のプロトコルと見なされる
- OSI参照モデルではアプリケーション層に位置づけられる

cf. 
- [wikipedia - Dynamic Host Configuration Protocol](https://ja.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol)
- [www.nic.ad.jp - DHCPとは](https://www.nic.ad.jp/ja/basics/terms/dhcp.html)

## リンクローカルアドレス
- ホストが接続されているネットワークセグメント、またはブロードキャストドメイン（ブロードキャストできるネットワーク範囲内のこと。ブロードキャストは同じネットワークにいるユーザー全てに同じデータを送信すること。）内の通信のみ有効なアドレス
- DHCPサーバーが存在しないネットワークで使われる、自ら発行する特殊なIPアドレス
- 
cf.
- [e-words.jp](http://e-words.jp/w/%E3%83%AA%E3%83%B3%E3%82%AF%E3%83%AD%E3%83%BC%E3%82%AB%E3%83%AB%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9.html)

## ICMP
- Internet Control Message Protocol
- IPプロトコルのエラー通知や制御メッセージを転送するためのプロトコル
  - 通信状態を確認するために使用される
- pingコマンドで使われている

cf.
- [wikipedia - Internet Control Message Protocol](https://ja.wikipedia.org/wiki/Internet_Control_Message_Protocol)
- [www.infraexpert.com - TCP/IP-ICMP](https://www.infraexpert.com/study/tcpip4.html)

## L4ロードバランサー
- L4→トランスポート層
- IPアドレスとポート番号によるロードバランシングが可能

cf. 
- [faq.support.nifcloud.com - ロードバランサー（L4）とL7ロードバランサー（Pulse Secure Virtual Traffic Manager）の違いを教えてください](https://faq.support.nifcloud.com/faq/show/420?site_domain=default)
- [www.kimullaa.com - L7ロードバランサとL4ロードバランサ](https://www.kimullaa.com/entry/2019/12/01/135430)

## L7ロードバランサー
- L7→アプリケーション層
- URLやHTTPヘッダーによるロードバランシングが可能

cf. 
- [faq.support.nifcloud.com - ロードバランサー（L4）とL7ロードバランサー（Pulse Secure Virtual Traffic Manager）の違いを教えてください](https://faq.support.nifcloud.com/faq/show/420?site_domain=default)
- [www.kimullaa.com - L7ロードバランサとL4ロードバランサ](https://www.kimullaa.com/entry/2019/12/01/135430)

