---
title: Created a Google Extension Plugin
description: 'Build a Google Chrome extension plugin with manifest.json and JavaScript: Complete guide to creating your first plugin.'
slug: build-google-extension-plugin
date: 2017-09-26T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - JavaScript
  - Google Chrome Extension
translation_key: build-google-extension-plugin
---

When I wanted to search in English on Google, I used bookmarks to access it, but I thought it would be convenient to do it with a plugin, so I created one.

This plugin might be somewhat useful for niche users who hide the default bookmarks and use a plugin called Bookolio (which makes bookmarks easier to see) ← me.

# Environment
* Google Chrome
* JavaScript

# Specifications
There are various types of plugins, but what I am creating this time is this:

![Plugin Image](#)

When you click the plugin icon, it simply opens the [English version of Google](https://www.google.co.jp/?hl=en&gws_rd=cr&ei=O2OgV4jODcS30gSs1Ihw) in a new tab. 

Since it's super simple, there is room for improvement in the specifications. ()

# Preparation
First, create the folder and files.

```
└── search_by_english
    ├── background.js
    ├── icons
    │   ├── icon128.png
    │   ├── icon16.png
    │   └── icon48.png
    └── manifest.json
```

*Please prepare the icons as needed.*

The background.js is the JavaScript that runs in the background. For more details, please refer to the [Developer's Guide](https://developer.chrome.com/extensions/devguide).

# Edit manifest.json

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

The description varies depending on the type of plugin. It's not particularly difficult, so for details, please refer to the [Developer's Guide](https://developer.chrome.com/extensions/devguide).

# Edit background.js

background.js

I referred to the code from [Chrome extension: open link in new tab?](http://stackoverflow.com/questions/16503879/chrome-extension-open-link-in-new-tab). You should get a general idea by looking at it.

For more details, please refer to the documentation (ry) [Developer's Guide](https://developer.chrome.com/extensions/devguide).

```
chrome.browserAction.onClicked.addListener(function(activeTab){
  var newURL = "https://www.google.co.jp/?hl=en&gws_rd=cr&ei=O2OgV4jODcS30gSs1Ihw";
  chrome.tabs.create({ url: newURL });
});
```

Basically, you write in JavaScript, but you will be utilizing the API provided by Google regarding browser operations.

# Thoughts
I would also like to try developing plugins for Atom.

# References
* [Chrome extension: open link in new tab?](http://stackoverflow.com/questions/16503879/chrome-extension-open-link-in-new-tab)
* [Let's develop original Chrome extensions (with source code)](https://liginc.co.jp/web/tool/browser/163575)
* [Chrome extension development: Sending JavaScript from the extension to the page](http://qiita.com/suin/items/5e1aa942e654bce442f7)
* [Google Chrome extension implementation notes (Browser Action only)](http://qiita.com/dorachan1029/items/683a11d6f208e13f5b77)