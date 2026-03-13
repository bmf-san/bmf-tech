---
title: Challenges Faced When Building an SPA
description: "Fix SPA routing and resource path issues with nginx try_files configuration for single-page application development."
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
Previously, I wrote an article titled [Directory Structure and nginx conf File Considerations When Integrating SPA into Laravel](https://bmf-tech.com/posts/Laravel%E3%81%ABSPA%E3%82%92%E7%B5%84%E3%81%BF%E8%BE%BC%E3%82%80%E6%99%82%E3%81%AB%E8%80%83%E3%81%88%E3%81%9F%E3%83%86%E3%82%99%E3%82%A3%E3%83%AC%E3%82%AF%E3%83%88%E3%83%AA%E6%A7%8B%E6%88%90%E3%81%A8nginx%E3%81%AEconf%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB), but the nginx conf described there was insufficient, so I revisited and resolved the issues.

# Prerequisites
- History API
- nginx

# Challenges Faced When Building an SPA

## nginx Configuration
You need to configure it to always return `index.html` even after a reload. Use `try_files` in the conf like this:

```
location / {
        try_files $uri $uri/ /index.html;
}
```

## Paths for JS Files and Other Sources
In `index.html`, the path for JS files was specified as:

```javascript
<script type="text/javascript" src="./dist/bundle.js"></script>
```

This caused resources to be returned as `/dashboard/post/dist/bundle.js` when accessing `/dashboard/post`, etc. I changed it to an absolute path to always reference `bundle.js` regardless of the URI.

```javascript
<script type="text/javascript" src="/dist/bundle.js"></script>
```

# Thoughts
It took quite some time to resolve, but by separating the issues between nginx and the application, I was able to understand it quickly.

# References
- [react-router + Static Files (css, js) Combination Pitfalls](https://qiita.com/rooooomania/items/c50acf84d56793de6318)
