---
title: Encountering an Unfamiliar Error When Requiring Components in React
slug: react-babel-code-generator-warning
date: 2017-10-01T00:00:00Z
author: bmf-san
categories:
  - Application
tags:
  - React
  - Tips
translation_key: react-babel-code-generator-warning
---

While requiring a component in React, I encountered an unfamiliar error.

# Environment
* Laravel
* Elixir
* Babel
* Browserify


# File is Too Large!

`[BABEL] Note: The code generator has deoptimised the styling of "D:/path/to/hoge.js" as it exceeds the max of "100KB"`

This error seems to occur when using Babel and the file is too large.

[What does “The code generator has deoptimised the styling of [some file] as it exceeds the max of "100KB"” mean?](http://stackoverflow.com/questions/29576341/what-does-the-code-generator-has-deoptimised-the-styling-of-some-file-as-it-e)

It seems that there’s nothing to worry about, but if you want to hide the warning, you can set the `compact` option in Babel to false.

[How to Edit Babel Plugins When Using Browserify with Laravel-Elixir](http://qiita.com/fagai/items/c4dbe5d2adeb79e42e40)

This link discusses how to edit plugins, but it may also be helpful regarding options.

However, even if I suppress the warning, the file remains large, so I addressed this by compressing the file.

`gulp --production`

I wasn't quite sure how to change the options. ヽ(´ー｀)ノ