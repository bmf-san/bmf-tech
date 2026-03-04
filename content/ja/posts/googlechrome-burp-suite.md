---
title: "GoogleChromeでBurp Suiteを使う手順"
slug: "googlechrome-burp-suite"
date: 2019-03-22
author: bmf-san
categories:
  - "ツール"
tags:
  - "Burp Suite"
  - "セキュリティ"
  - "脆弱性"
draft: false
---

# 概要
Burp SuiteをChormeで使う際の諸々の設定について。
脆弱性診断や対応時にburpをchromeで使えるようにしたかった。

# 環境
Mac OS

# 準備
- [burpsuite](https://support.portswigger.net/customer/portal/topics/718317-installing-and-configuring-burp/articles)
- [chrome extension - proxy switchsharp](https://chrome.google.com/webstore/detail/proxy-switchysharp/dpplabbmogkhghncfbfdeeokoefdjegm?hl=ja)

#  手順
## プロキシを設定する
[chrome extension - proxy switchsharp](https://chrome.google.com/webstore/detail/proxy-switchysharp/dpplabbmogkhghncfbfdeeokoefdjegm?hl=ja)でProxy Profilesを設定する。

Profile NameをBurp（何でもよいがBurp用のプロキシ設定なのでBurpにしておく）
Manual Configurationを選択して、HTTP Proxyに `127.0.0.1`を設定、Portは各自の環境でバッティングしないように設定。
Saveを押してProfileを保存。

proxy switchsharpを使うのは毎回プロキシ設定を変更するのが手間なので拡張機能でシャッとプロキシ設定を変更できるようにするため。

プロキシ設定の変更はChromeの右上の拡張機能が並んでいるところからproxy switchsharpを選択して任意のProfileを選択することで変更することができる。

特にプロキシ設定をいじる必要がない普段はDirect Connectionを選択。

## 証明書を設定する
 Burpを起動。

chromeのプロキシ設定が上で設定したProfileになっているか確認。（proxy switchsharpでプロキシ設定を保存しただけでは設定が有効になっていないので、chromeブラウザ右上の拡張機能が並んでいるところがproxy switchsharpを選択してProfileを選択して有効にする必要がある）

Burpをデフォルトの設定で起動した場合は、`http://127.0.0.1:8080`にアクセスする。

右上のCA Certificateをクリックし、証明書をダウンロード。

ダウンロードした証明書をKeychain accessで開き、証明書を`常に信頼`に設定
する。
`Port Swigger CA`という証明書名なはず。

以上の作業でChromeでBurpが使えるようになっているはず。

# 補足
Chromeだとlocalhostをinterceptするには、proxy switchsharpのProfile Detail > No Proxy Forで`<-loopback>`を追加する必要がある。

[Burp InterceptionがChromeのローカルホストで機能しない](https://stackoverflow.com/questions/55616614/burp-interception-does-not-work-for-localhost-in-chrome)

# 参考
- [Burp Suiteを使うためのGoogle Chromeでの設定の仕方](https://taiyakon.com/2018/05/burp-suite-macosgoogle-chrome.html)
