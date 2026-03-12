---
title: Google ExtenstionのPluginをつくってみた
description: Google ExtenstionのPluginをつくってみたについて、基本的な概念から実践的な知見まで詳しく解説します。
slug: build-google-extension-plugin
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - アプリケーション
tags:
  - JavaScript
  - google chrome extension
translation_key: build-google-extension-plugin
---


英語版のGoogleで検索したい時には、ブックマークを利用してアクセスしていたのですが、プラグインで便利にできないかなーと思い、作ってみました。

デフォルトのブックマークが非表示で、Bookolio（ブクマを見やすくするやつ）とかいうプラグインを使っているニッチな人だと多少便利なプラグインかもしれません←自分


# 環境
* Google Chrome
* Javascript


# 仕様
プラグインの種類は色々ありますが、今回つくるのはこれです↓

【プラグインの画像】

プラグインのアイコンを押すと、[Googleの英語版](https://www.google.co.jp/?hl=en&gws_rd=cr&ei=O2OgV4jODcS30gSs1Ihw)を新規タブで開いてくれるだけの超単純な機能です。

超単純なだけに伸びしろのある仕様ですね（）



# 準備

先にフォルダとファイルを作っておきます。


```
└── search_by_english
    ├── background.js
    ├── icons
    │   ├── icon128.png
    │   ├── icon16.png
    │   └── icon48.png
    └── manifest.json
```
※アイコンは適宜用意してください。


background.jsというのはbackgroundで動作するJavaScriptです（）
詳しくは[Developer's Guide](https://developer.chrome.com/extensions/devguide)をご覧ください。



# manifest.jsonを編集

manifest.json

```
{
  "name": "Open A Google English Edition In A New Tab",
  "version": "1.0",
  "manifest_version": 2,
  "description": "Open a Google English Edition in a new tab.",
  "icons": {
    "16": "icons/icon16.png",
    "48": "icons/icon48.png",
    "128": "icons/icon128.png"
  },
  "browser_action": {
      "default_icon": "icons/icon48.png"
  },
  "background": {
    "scripts": [
      "background.js"
    ]
  }
}
```

プラグインの種類によって記述が変わります。これといって難しいものではないので詳細は[Developer's Guide](https://developer.chrome.com/extensions/devguide)をご覧ください。

# background.jsを編集

background.js

コードは[Chrome extension: open link in new tab?](http://stackoverflow.com/questions/16503879/chrome-extension-open-link-in-new-tab)を参考にさせて頂きました。
見ると何となくわかるかと思います。

詳しくはドキュメン(ry  [Developer's Guide](https://developer.chrome.com/extensions/devguide)

```
chrome.browserAction.onClicked.addListener(function(activeTab){
  var newURL = "https://www.google.co.jp/?hl=en&gws_rd=cr&ei=O2OgV4jODcS30gSs1Ihw";
  chrome.tabs.create({ url: newURL });
});
```

基本はjavascriptでホイホイかくわけですが、ブラウザの動作等に関してGoogleが用意するAPIを活用していく感じでしょうか。


# 所感
Atomのプラグイン開発とかもやってみたいです。


# 参考
* [Chrome extension: open link in new tab?](http://stackoverflow.com/questions/16503879/chrome-extension-open-link-in-new-tab)
* [Chromeのオリジナル拡張機能を開発しよう（ソースコードあり）](https://liginc.co.jp/web/tool/browser/163575)
* [Chrome拡張開発: 拡張からページにJavaScriptを送り込みたい](http://qiita.com/suin/items/5e1aa942e654bce442f7)
* [google chrome extention 実装メモ (Browser Actionだけ)](http://qiita.com/dorachan1029/items/683a11d6f208e13f5b77)

