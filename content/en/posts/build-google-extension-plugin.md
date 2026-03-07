---
title: Creating a Google Chrome Extension Plugin
slug: build-google-extension-plugin
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Applications
tags:
  - JavaScript
  - Google Chrome Extension
description: Exploring how to create a simple plugin for Google Chrome to open the English version of Google in a new tab.
translation_key: build-google-extension-plugin
---

When I wanted to search on the English version of Google, I used bookmarks to access it. However, I wondered if I could make it more convenient with a plugin, so I decided to create one.

For those niche users who hide their default bookmarks and use plugins like Bookolio (which makes bookmarks easier to view), this plugin might be somewhat useful—like for myself.

# Environment
* Google Chrome
* JavaScript

# Specifications
There are various types of plugins, but the one I created this time is as follows:

【Plugin Image】

When you click the plugin icon, it simply opens the [English version of Google](https://www.google.co.jp/?hl=en&gws_rd=cr&ei=O2OgV4jODcS30gSs1Ihw) in a new tab. It's a super simple functionality.

Its simplicity leaves room for improvement.

# Preparation
First, create the folder and files as follows:

```
└── search_by_english
    ├── background.js
    ├── icons
    │   ├── icon128.png
    │   ├── icon16.png
    │   └── icon48.png
    └── manifest.json
```
※ Prepare the icons as needed.

`background.js` is the JavaScript that runs in the background. For more details, refer to the [Developer's Guide](https://developer.chrome.com/extensions/devguide).

# Editing manifest.json

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

The description varies depending on the type of plugin. It's not particularly difficult, so for more details, refer to the [Developer's Guide](https://developer.chrome.com/extensions/devguide).

# Editing background.js

background.js

The code was referenced from [Chrome extension: open link in new tab?](http://stackoverflow.com/questions/16503879/chrome-extension-open-link-in-new-tab). You should be able to understand it somewhat by looking at it.

For more details, refer to the documentation (ry) [Developer's Guide](https://developer.chrome.com/extensions/devguide).

```
chrome.browserAction.onClicked.addListener(function(activeTab){
  var newURL = "https://www.google.co.jp/?hl=en&gws_rd=cr&ei=O2OgV4jODcS30gSs1Ihw";
  chrome.tabs.create({ url: newURL });
});
```

Basically, you write the code in JavaScript, utilizing the APIs provided by Google for browser operations.

# Thoughts
I’d like to try developing plugins for Atom as well.

# References
* [Chrome extension: open link in new tab?](http://stackoverflow.com/questions/16503879/chrome-extension-open-link-in-new-tab)
* [Developing original Chrome extensions (with source code)](https://liginc.co.jp/web/tool/browser/163575)
* [Chrome extension development: Injecting JavaScript into pages from extensions](http://qiita.com/suin/items/5e1aa942e654bce442f7)
* [Google Chrome extension implementation notes (Browser Action only)](http://qiita.com/dorachan1029/items/683a11d6f208e13f5b77)