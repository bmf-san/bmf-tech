---
title: Challenges Faced When Building an SPA
slug: spa-development-challenges
date: 2018-06-06T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - Nginx
  - SPA
translation_key: spa-development-challenges
---

# Overview
Previously, I wrote an article titled [Directory Structure and Nginx Configurations Considered When Integrating SPA into Laravel](https://bmf-tech.com/posts/Laravel%E3%81%ABSPA%E3%82%92%E7%B5%84%E3%81%BF%E8%BE%BC%E3%82%80%E6%99%82%E3%81%AB%E8%80%83%E3%81%88%E3%81%9F%E3%83%86%E3%82%99%E3%82%A3%E3%83%AC%E3%82%AF%E3%83%88%E3%83%AA%E6%A7%8B%E6%88%90%E3%81%A8nginx%E3%81%AEconf%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB), but the nginx configuration provided was insufficient, so I organized and resolved the issues again.

# Prerequisites
- History API
- nginx

# Challenges Faced When Building an SPA

## Nginx Configuration
It is necessary to configure nginx to always return index.html even after a reload. You can set up the conf using try_files like this:

```
location / {
        try_files $uri $uri/ /index.html;
}
```

## Source Path for JS Files
In index.html, the path for the js file was specified as

```javascript
<script type="text/javascript" src="./dist/bundle.js"></script>
```

As a result, when accessing `/dashboard/post`, it would return the resource as `/dashboard/post/dist/bundle.js`.

To ensure that bundle.js can always be referenced regardless of the URI, I changed it to use an absolute path:

```javascript
<script type="text/javascript" src="/dist/bundle.js"></script>
```

# Thoughts
It took quite a while to resolve, but by isolating whether the issue was on the nginx side or the application side, I was able to understand it quickly.

# References
- [The Trap of Combining react-router + Static Files (css, js)](https://qiita.com/rooooomania/items/c50acf84d56793de6318)